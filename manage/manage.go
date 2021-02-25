package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/coreos/go-iptables/iptables"
	"github.com/hashicorp/consul/api"
	"github.com/lorenzosaino/go-sysctl"
	"github.com/zerotier/go-ztcentral"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type ZTStatus struct {
	Address string `json:"address"`
	Clock   int64  `json:"clock"`
	Config  struct {
		Physical interface{} `json:"physical"`
		Settings struct {
			AllowTCPFallbackRelay bool   `json:"allowTcpFallbackRelay"`
			PortMappingEnabled    bool   `json:"portMappingEnabled"`
			PrimaryPort           int    `json:"primaryPort"`
			SoftwareUpdate        string `json:"softwareUpdate"`
			SoftwareUpdateChannel string `json:"softwareUpdateChannel"`
		} `json:"settings"`
	} `json:"config"`
	Online               bool   `json:"online"`
	PlanetWorldID        int    `json:"planetWorldId"`
	PlanetWorldTimestamp int64  `json:"planetWorldTimestamp"`
	PublicIdentity       string `json:"publicIdentity"`
	TCPFallbackActive    bool   `json:"tcpFallbackActive"`
	Version              string `json:"version"`
	VersionBuild         int    `json:"versionBuild"`
	VersionMajor         int    `json:"versionMajor"`
	VersionMinor         int    `json:"versionMinor"`
	VersionRev           int    `json:"versionRev"`
}

func getIPFromInterface(ifaceName string) (string, error) {
	var (
		addresses  []net.Addr
		maxRetries int
	)

	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		return "", err
	}

	// We only care about IPv4 right now
	for {
		maxRetries += 1
		log.Printf("maxRetries: %d\n", maxRetries)
		if maxRetries > 300 {
			return "", errors.New("timeout exceeded waiting for IPv4")
		}

		addresses, err = iface.Addrs()
		log.Println(addresses)
		if err != nil {
			return "", err
		}

		if len(addresses) == 0 {
			// I *think* this is always bad?
			return "", nil
		}

		for _, addr := range addresses {
			log.Printf("Checking (%d): %+v\n", len(addr.(*net.IPNet).IP.To4()), addr)
			if len(addr.(*net.IPNet).IP.To4()) == net.IPv4len {
				return addr.(*net.IPNet).IP.String(), nil
			}
		}

		time.Sleep(time.Second)
	}
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

type ztLocal struct {
	authToken  []byte
	statusPort []byte
}

func newZTLocal() ztLocal {
	authTokenFile := "/var/lib/zerotier-one/authtoken.secret"
	statusPortFile := "/var/lib/zerotier-one/zerotier-one.port"

	// Wait for auth token file to get created
	for !fileExists(authTokenFile) {
	}

	authToken, err := ioutil.ReadFile(authTokenFile)
	if err != nil {
		panic(err)
	}

	// Wait for status port file to get created
	for !fileExists(statusPortFile) {
	}

	statusPort, err := ioutil.ReadFile(statusPortFile)
	if err != nil {
		panic(err)
	}

	// Wait for ZT http service to start
	for {
		_, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", statusPort))
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	return ztLocal{authToken: authToken, statusPort: statusPort}
}

func (zt ztLocal) getZTStatus() ZTStatus {
	client := &http.Client{}
	var ztStatus ZTStatus

	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:%s/status", zt.statusPort), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-ZT1-Auth", string(zt.authToken))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &ztStatus)
	return ztStatus
}

func (zt ztLocal) joinNetwork(networkID string) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:%s/network/%s", zt.statusPort, networkID), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-ZT1-Auth", string(zt.authToken))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}

func main() {
	var (
		podCIDR      string
		routeUpdated bool
		viaIP        string
		err          error
		maxRetries   int

		nodeInterface = "eth0"       // We may need to find this
		ztInterface   = "ztrtay6555" // We may need to find this
		networkID     = os.Getenv("NETWORK_ID")
		ctx           = context.Background()
		ztCtl         = newZTLocal()
		ztStatus      = ztCtl.getZTStatus()
		ztClient      = ztcentral.NewClient(os.Getenv("ZTAUTHTOKEN"))
		autoJoin      = os.Getenv("AUTOJOIN")
		consul        = os.Getenv("CONSUL")
	)

	log.Println("Joining Network")
	ztCtl.joinNetwork(networkID)

	// Set hostname
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	log.Println("Creating Authorized Member")
	if autoJoin == "true" {
		ztClient.CreateAuthorizedMember(ctx, networkID, ztStatus.Address, hostname)
	}

	// Get node IP so we can find ourselves in kube later
	log.Println("Getting nodeIP")
	nodeIP, err := getIPFromInterface(nodeInterface)
	if err != nil {
		log.Println("error:", err.Error())
		return
	}
	log.Printf("Using nodeIP: %s\n", nodeIP)

	log.Println("Getting viaIP")
	// Get the zerotier IP as that's our gateway address. Need to wait for it to exist.
	for viaIP == "" {
		maxRetries += 1
		if maxRetries > 300 { // 5 minutes
			panic("timeout waiting for ZT IP")
		}

		viaIP, err = getIPFromInterface(ztInterface)
		if err != nil {
			if err.Error() != "route ip+net: no such network interface" {
				log.Println("error:", err.Error())
				return
			}
		}
		time.Sleep(time.Second)

		msg := fmt.Sprintf("Waiting for a ZeroTier IP on %s interface...", ztInterface)
		if autoJoin != "true" {
			msg += " Accept the new host on my.zerotier.com"
		}
		fmt.Println(msg)
	}
	log.Printf("Using viaIP: %s\n", viaIP)

	// get podCIDR from kubernetes

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Println("error:", err.Error())
		return
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println("error:", err.Error())
		return
	}

	// Get the node list from kubernetes
	nodes, err := clientSet.CoreV1().Nodes().List(ctx, v1.ListOptions{})
	if err != nil {
		log.Println("error:", err.Error())
		return
	}

	// Find the podCIDR from our node using the ExternalIP we found earlier
	for _, node := range nodes.Items {
		fmt.Println(node.Name)
		for _, address := range node.Status.Addresses {
			if address.Type == "ExternalIP" {
				if address.Address == nodeIP {
					podCIDR = node.Spec.PodCIDR
					log.Printf("Using podCIDR: %s\n", podCIDR)
				}
			}
		}
	}

	// update zerotier routing table

	// get consul client so we can do locking
	client, err := api.NewClient(&api.Config{Address: consul})
	lock, err := client.LockKey("zerotier-route-table/1")

	log.Println("Acquiring lock")
	lock.Lock(nil)

	network, err := ztClient.GetNetwork(ctx, networkID)
	if err != nil {
		log.Println("error:", err.Error())
		return
	}

	// look for our podCIDR and update it (new container on existing node)
	for idx := range network.Config.Routes {
		if network.Config.Routes[idx].Target == podCIDR {
			network.Config.Routes[idx].Via = viaIP
			routeUpdated = true
			log.Println("Found route entry, updating")
		}
	}

	// If we didn't find our podCIDR in the routing table we're a new node
	if !routeUpdated {
		newRoute := ztcentral.Route{
			Target: podCIDR,
			Via:    viaIP,
		}
		network.Config.Routes = append(network.Config.Routes, newRoute)
		log.Println("No existing route entry, adding new route")
	}

	_, err = ztClient.UpdateNetwork(ctx, network)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("Releasing lock")
	lock.Unlock()

	err = sysctl.Set("net.ipv4.ip_forward", "1")
	ipTables, err := iptables.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	ipTables.Append("nat", "POSTROUTING", fmt.Sprintf("-o %s -j MASQUERADE", nodeInterface))

	// Wait forever
	select {}
}
