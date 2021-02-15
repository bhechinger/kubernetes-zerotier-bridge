# \UtilApi

All URIs are relative to *https://my.zerotier.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRandomToken**](UtilApi.md#GetRandomToken) | **Get** /randomToken | Get a random 32 character token



## GetRandomToken

> RandomToken GetRandomToken(ctx).Execute()

Get a random 32 character token



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
    resp, r, err := api_client.UtilApi.GetRandomToken(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `UtilApi.GetRandomToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRandomToken`: RandomToken
    fmt.Fprintf(os.Stdout, "Response from `UtilApi.GetRandomToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRandomTokenRequest struct via the builder pattern


### Return type

[**RandomToken**](RandomToken.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

