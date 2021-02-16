package main

import (
	"context"
	"fmt"
	"os"

	zerotier "github.com/bhechinger/kubernetes-zerotier-bridge/routes/zerotier"
)

const (
	NetworkID = "8286ac0e472379fa"
	AuthToken = "JnT3HHRnSZWSP1pcCd2nGAc134TKxk4u"
)

func main() {
	ctx := context.WithValue(context.Background(), zerotier.ContextAccessToken, AuthToken)

	configuration := zerotier.NewConfiguration()
	apiClient := zerotier.NewAPIClient(configuration)

	network, r, err := apiClient.NetworkApi.GetNetworkByID(ctx, NetworkID).Execute()
	if err.Error() != "" {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.GetNetworkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}


	fmt.Printf("Id: %s\n", *network.Id)
	config := network.GetConfig()
	routes := config.GetRoutes()
	var newRoutes []zerotier.Route
	for _, route := range routes {
		fmt.Printf("Target: %v\n", route.GetTarget())
		fmt.Printf("Via: %v\n", route.GetVia())
		newRoute := zerotier.NewRoute()
		newRoute.SetTarget("1.2.3.4")
		newRoute.SetVia("3.3.3.3")
		newRoutes = append(newRoutes, *newRoute)
	}

	newNetwork := zerotier.NewNetwork()
	newNetwork.Config = zerotier.NewNetworkConfig()
	newNetwork.Config.SetRoutes(newRoutes)

	resp, r, err := apiClient.NetworkApi.UpdateNetwork(ctx, NetworkID).Network(*newNetwork).Execute()
	if err.Error() != "" {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.UpdateNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetwork`: Network
	fmt.Fprintf(os.Stdout, "Response from `NetworkApi.UpdateNetwork`: %v\n", resp)
}
