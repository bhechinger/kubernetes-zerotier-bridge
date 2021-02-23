package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/zerotier/go-ztcentral"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func getIPFromInterface(ifaceName string) (string, error) {
	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		return "", err
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return "", err
	}

	return addrs[0].(*net.IPNet).IP.String(), nil
}

func main() {
	nodeIP, err := getIPFromInterface("eth0")
	if err != nil {
		log.Println("error:", err.Error())
		return
	}
	log.Printf("Using nodeIP: %s\n", nodeIP)

	viaIP, err := getIPFromInterface("ztrtay6555")
	if err != nil {
		log.Println("error:", err.Error())
		return
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

	ctx := context.Background()

	var podCIDR string
	nodes, err := clientSet.CoreV1().Nodes().List(ctx, v1.ListOptions{})
	if err != nil {
		log.Println("error:", err.Error())
		return
	}

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
	c := ztcentral.NewClient(os.Getenv("ZTAUTHTOKEN"))

	network, err := c.GetNetwork(ctx, os.Getenv("NETWORK_IDS"))
	if err != nil {
		log.Println("error:", err.Error())
		return
	}

	var routeUpdated bool
	for idx := range network.Config.Routes {
		if network.Config.Routes[idx].Target == podCIDR {
			network.Config.Routes[idx].Via = viaIP
			routeUpdated = true
		}
	}

	if !routeUpdated {
		newRoute := ztcentral.Route{
			Target: podCIDR,
			Via:    viaIP,
		}
		network.Config.Routes = append(network.Config.Routes, newRoute)
	}

	fmt.Println(network.Config.Routes)

	_, err = c.UpdateNetwork(ctx, network)
	if err != nil {
		fmt.Println(err)
		return
	}
}
