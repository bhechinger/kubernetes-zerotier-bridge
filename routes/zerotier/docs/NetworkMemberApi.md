# \NetworkMemberApi

All URIs are relative to *https://my.zerotier.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNetworkMember**](NetworkMemberApi.md#DeleteNetworkMember) | **Delete** /network/{networkID}/member/{memberID} | Delete a network member
[**GetNetworkMember**](NetworkMemberApi.md#GetNetworkMember) | **Get** /network/{networkID}/member/{memberID} | Return an individual member on a network
[**GetNetworkMemberList**](NetworkMemberApi.md#GetNetworkMemberList) | **Get** /network/{networkID}/member | Returns a list of Members on the network.
[**UpdateNetworkMember**](NetworkMemberApi.md#UpdateNetworkMember) | **Post** /network/{networkID}/member/{memberID} | Modify a network member



## DeleteNetworkMember

> DeleteNetworkMember(ctx, networkID, memberID).Execute()

Delete a network member

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
    memberID := "memberID_example" // string | ID of the member

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkMemberApi.DeleteNetworkMember(context.Background(), networkID, memberID).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkMemberApi.DeleteNetworkMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkID** | **string** | ID of the network | 
**memberID** | **string** | ID of the member | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkMemberRequest struct via the builder pattern


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


## GetNetworkMember

> Member GetNetworkMember(ctx, networkID, memberID).Execute()

Return an individual member on a network

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
    memberID := "memberID_example" // string | ID of the member

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkMemberApi.GetNetworkMember(context.Background(), networkID, memberID).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkMemberApi.GetNetworkMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkMember`: Member
    fmt.Fprintf(os.Stdout, "Response from `NetworkMemberApi.GetNetworkMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkID** | **string** | ID of the network | 
**memberID** | **string** | ID of the member | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Member**](Member.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkMemberList

> []Member GetNetworkMemberList(ctx, networkID).Execute()

Returns a list of Members on the network.

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
    resp, r, err := api_client.NetworkMemberApi.GetNetworkMemberList(context.Background(), networkID).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkMemberApi.GetNetworkMemberList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkMemberList`: []Member
    fmt.Fprintf(os.Stdout, "Response from `NetworkMemberApi.GetNetworkMemberList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkID** | **string** | ID of the network to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMemberListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Member**](Member.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkMember

> Member UpdateNetworkMember(ctx, networkID, memberID).Member(member).Execute()

Modify a network member

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
    memberID := "memberID_example" // string | ID of the member
    member := *openapiclient.NewMember() // Member | Member object JSON

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NetworkMemberApi.UpdateNetworkMember(context.Background(), networkID, memberID).Member(member).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkMemberApi.UpdateNetworkMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkMember`: Member
    fmt.Fprintf(os.Stdout, "Response from `NetworkMemberApi.UpdateNetworkMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkID** | **string** | ID of the network | 
**memberID** | **string** | ID of the member | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **member** | [**Member**](Member.md) | Member object JSON | 

### Return type

[**Member**](Member.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

