# \NetworkApi

All URIs are relative to *https://my.zerotier.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNetwork**](NetworkApi.md#DeleteNetwork) | **Delete** /network/{networkID} | delete network
[**GetNetworkByID**](NetworkApi.md#GetNetworkByID) | **Get** /network/{networkID} | Get network by ID
[**GetNetworkList**](NetworkApi.md#GetNetworkList) | **Get** /network | Returns a list of Networks you have access to.
[**NewNetwork**](NetworkApi.md#NewNetwork) | **Post** /network | Create a new network.
[**UpdateNetwork**](NetworkApi.md#UpdateNetwork) | **Post** /network/{networkID} | update network configuration



## DeleteNetwork

> DeleteNetwork(ctx, networkID).Execute()

delete network

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkID := "networkID_example" // string | ID of the network

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkApi.DeleteNetwork(context.Background(), networkID).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.DeleteNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkID** | **string** | ID of the network | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkByID

> Network GetNetworkByID(ctx, networkID).Execute()

Get network by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkID := "networkID_example" // string | ID of the network to return

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkApi.GetNetworkByID(context.Background(), networkID).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.GetNetworkByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkByID`: Network
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.GetNetworkByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkID** | **string** | ID of the network to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Network**](Network.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkList

> []Network GetNetworkList(ctx).Execute()

Returns a list of Networks you have access to.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkApi.GetNetworkList(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.GetNetworkList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkList`: []Network
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.GetNetworkList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkListRequest struct via the builder pattern


### Return type

[**[]Network**](Network.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewNetwork

> Network NewNetwork(ctx).Execute()

Create a new network.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkApi.NewNetwork(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.NewNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewNetwork`: Network
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.NewNetwork`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNewNetworkRequest struct via the builder pattern


### Return type

[**Network**](Network.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetwork

> Network UpdateNetwork(ctx, networkID).Network(network).Execute()

update network configuration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    networkID := "networkID_example" // string | ID of the network to change
    network := *openapiclient.NewNetwork() // Network | Network object JSON

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkApi.UpdateNetwork(context.Background(), networkID).Network(network).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.UpdateNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetwork`: Network
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.UpdateNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkID** | **string** | ID of the network to change | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **network** | [**Network**](Network.md) | Network object JSON | 

### Return type

[**Network**](Network.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

