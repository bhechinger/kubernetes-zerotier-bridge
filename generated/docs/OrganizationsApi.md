# \OrganizationsApi

All URIs are relative to *https://my.zerotier.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptInvitation**](OrganizationsApi.md#AcceptInvitation) | **Post** /org-invitation/{inviteID} | Accept organization invitation
[**DeclineInvitation**](OrganizationsApi.md#DeclineInvitation) | **Delete** /org-invitation/{inviteID} | Decline organization invitation
[**GetInvitationByID**](OrganizationsApi.md#GetInvitationByID) | **Get** /org-invitation/{inviteID} | Get organization invitation
[**GetOrganization**](OrganizationsApi.md#GetOrganization) | **Get** /org | Get the current user&#39;s organization
[**GetOrganizationByID**](OrganizationsApi.md#GetOrganizationByID) | **Get** /org/{orgID} | Get organization by ID
[**GetOrganizationInvitationList**](OrganizationsApi.md#GetOrganizationInvitationList) | **Get** /org-invitation | Get list of organization invitations
[**GetOrganizationMembers**](OrganizationsApi.md#GetOrganizationMembers) | **Get** /org/{orgID}/user | Get list of organization members
[**InviteUserByEmail**](OrganizationsApi.md#InviteUserByEmail) | **Post** /org-invitation | Invite a user to your organization by email



## AcceptInvitation

> OrganizationInvitation AcceptInvitation(ctx, inviteID).Execute()

Accept organization invitation

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
    inviteID := "inviteID_example" // string | Invitation ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.AcceptInvitation(context.Background(), inviteID).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.AcceptInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AcceptInvitation`: OrganizationInvitation
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.AcceptInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteID** | **string** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationInvitation**](OrganizationInvitation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeclineInvitation

> DeclineInvitation(ctx, inviteID).Execute()

Decline organization invitation

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
    inviteID := "inviteID_example" // string | Invitation ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.DeclineInvitation(context.Background(), inviteID).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.DeclineInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteID** | **string** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclineInvitationRequest struct via the builder pattern


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


## GetInvitationByID

> OrganizationInvitation GetInvitationByID(ctx, inviteID).Execute()

Get organization invitation

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
    inviteID := "inviteID_example" // string | Invitation ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.GetInvitationByID(context.Background(), inviteID).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetInvitationByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvitationByID`: OrganizationInvitation
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetInvitationByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteID** | **string** | Invitation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvitationByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationInvitation**](OrganizationInvitation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> Organization GetOrganization(ctx).Execute()

Get the current user's organization

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
    resp, r, err := api_client.OrganizationsApi.GetOrganization(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganization`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganization`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


### Return type

[**Organization**](Organization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationByID

> Organization GetOrganizationByID(ctx, orgID).Execute()

Get organization by ID

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
    orgID := "orgID_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.GetOrganizationByID(context.Background(), orgID).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationByID`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgID** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organization**](Organization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInvitationList

> []OrganizationInvitation GetOrganizationInvitationList(ctx).Execute()

Get list of organization invitations

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
    resp, r, err := api_client.OrganizationsApi.GetOrganizationInvitationList(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationInvitationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInvitationList`: []OrganizationInvitation
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationInvitationList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInvitationListRequest struct via the builder pattern


### Return type

[**[]OrganizationInvitation**](OrganizationInvitation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationMembers

> OrganizationMember GetOrganizationMembers(ctx, orgID).Execute()

Get list of organization members

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
    orgID := "orgID_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.GetOrganizationMembers(context.Background(), orgID).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationMembers`: OrganizationMember
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgID** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrganizationMember**](OrganizationMember.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUserByEmail

> OrganizationInvitation InviteUserByEmail(ctx).OrganizationInvitation(organizationInvitation).Execute()

Invite a user to your organization by email

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
    organizationInvitation := *openapiclient.NewOrganizationInvitation() // OrganizationInvitation | Organization Invitation JSON object

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrganizationsApi.InviteUserByEmail(context.Background()).OrganizationInvitation(organizationInvitation).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.InviteUserByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InviteUserByEmail`: OrganizationInvitation
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.InviteUserByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteUserByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationInvitation** | [**OrganizationInvitation**](OrganizationInvitation.md) | Organization Invitation JSON object | 

### Return type

[**OrganizationInvitation**](OrganizationInvitation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

