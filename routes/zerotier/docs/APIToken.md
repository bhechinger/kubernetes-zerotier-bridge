# APIToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenName** | Pointer to **string** | user specified token name | [optional] 
**Token** | Pointer to **string** | API Token.  Minimum 32 characters. This token is encrypted in the database and can not be retrieved once set | [optional] 

## Methods

### NewAPIToken

`func NewAPIToken() *APIToken`

NewAPIToken instantiates a new APIToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPITokenWithDefaults

`func NewAPITokenWithDefaults() *APIToken`

NewAPITokenWithDefaults instantiates a new APIToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenName

`func (o *APIToken) GetTokenName() string`

GetTokenName returns the TokenName field if non-nil, zero value otherwise.

### GetTokenNameOk

`func (o *APIToken) GetTokenNameOk() (*string, bool)`

GetTokenNameOk returns a tuple with the TokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenName

`func (o *APIToken) SetTokenName(v string)`

SetTokenName sets TokenName field to given value.

### HasTokenName

`func (o *APIToken) HasTokenName() bool`

HasTokenName returns a boolean if a field has been set.

### GetToken

`func (o *APIToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *APIToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *APIToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *APIToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


