# \UserApi

All URIs are relative to *https://my.zerotier.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAPIToken**](UserApi.md#AddAPIToken) | **Post** /user/{userID}/token | Add an API token
[**DeleteAPIToken**](UserApi.md#DeleteAPIToken) | **Delete** /user/{userID}/token/{tokenName} | Delete API Token
[**DeleteUserByID**](UserApi.md#DeleteUserByID) | **Delete** /user/{userID} | Delete user
[**GetUserByID**](UserApi.md#GetUserByID) | **Get** /user/{userID} | Get user record
[**UpdateUserByID**](UserApi.md#UpdateUserByID) | **Post** /user/{userID} | Update user record (SMS number or Display Name only)



## AddAPIToken

> APIToken AddAPIToken(ctx, userID).APIToken(aPIToken).Execute()

Add an API token

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
    userID := "userID_example" // string | User ID
    aPIToken := *openapiclient.NewAPIToken() // APIToken | APIToken JSON object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.AddAPIToken(context.Background(), userID).APIToken(aPIToken).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.AddAPIToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAPIToken`: APIToken
    fmt.Fprintf(os.Stdout, "Response from `UserApi.AddAPIToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAPITokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **aPIToken** | [**APIToken**](APIToken.md) | APIToken JSON object | 

### Return type

[**APIToken**](APIToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAPIToken

> DeleteAPIToken(ctx, userID, tokenName).Execute()

Delete API Token

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
    userID := "userID_example" // string | User ID
    tokenName := "tokenName_example" // string | Token Name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.DeleteAPIToken(context.Background(), userID, tokenName).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeleteAPIToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | User ID | 
**tokenName** | **string** | Token Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPITokenRequest struct via the builder pattern


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


## DeleteUserByID

> DeleteUserByID(ctx, userID).Execute()

Delete user



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
    userID := "userID_example" // string | User ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.DeleteUserByID(context.Background(), userID).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeleteUserByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserByIDRequest struct via the builder pattern


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


## GetUserByID

> User GetUserByID(ctx, userID).Execute()

Get user record

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
    userID := "userID_example" // string | User ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.GetUserByID(context.Background(), userID).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUserByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserByID`: User
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUserByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserByID

> User UpdateUserByID(ctx, userID).User(user).Execute()

Update user record (SMS number or Display Name only)

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
    userID := "userID_example" // string | User ID
    user := *openapiclient.NewUser() // User | User object JSON

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserApi.UpdateUserByID(context.Background(), userID).User(user).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UpdateUserByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserByID`: User
    fmt.Fprintf(os.Stdout, "Response from `UserApi.UpdateUserByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userID** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) | User object JSON | 

### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

