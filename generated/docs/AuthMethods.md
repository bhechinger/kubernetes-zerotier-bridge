# AuthMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Local** | Pointer to **NullableString** | email address for built-in authentication | [optional] [readonly] 
**Google** | Pointer to **NullableString** | Google OIDC ID | [optional] [readonly] 
**Oidc** | Pointer to **NullableString** | Generic OIDC ID | [optional] [readonly] 

## Methods

### NewAuthMethods

`func NewAuthMethods() *AuthMethods`

NewAuthMethods instantiates a new AuthMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodsWithDefaults

`func NewAuthMethodsWithDefaults() *AuthMethods`

NewAuthMethodsWithDefaults instantiates a new AuthMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocal

`func (o *AuthMethods) GetLocal() string`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *AuthMethods) GetLocalOk() (*string, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *AuthMethods) SetLocal(v string)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *AuthMethods) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### SetLocalNil

`func (o *AuthMethods) SetLocalNil(b bool)`

 SetLocalNil sets the value for Local to be an explicit nil

### UnsetLocal
`func (o *AuthMethods) UnsetLocal()`

UnsetLocal ensures that no value is present for Local, not even an explicit nil
### GetGoogle

`func (o *AuthMethods) GetGoogle() string`

GetGoogle returns the Google field if non-nil, zero value otherwise.

### GetGoogleOk

`func (o *AuthMethods) GetGoogleOk() (*string, bool)`

GetGoogleOk returns a tuple with the Google field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogle

`func (o *AuthMethods) SetGoogle(v string)`

SetGoogle sets Google field to given value.

### HasGoogle

`func (o *AuthMethods) HasGoogle() bool`

HasGoogle returns a boolean if a field has been set.

### SetGoogleNil

`func (o *AuthMethods) SetGoogleNil(b bool)`

 SetGoogleNil sets the value for Google to be an explicit nil

### UnsetGoogle
`func (o *AuthMethods) UnsetGoogle()`

UnsetGoogle ensures that no value is present for Google, not even an explicit nil
### GetOidc

`func (o *AuthMethods) GetOidc() string`

GetOidc returns the Oidc field if non-nil, zero value otherwise.

### GetOidcOk

`func (o *AuthMethods) GetOidcOk() (*string, bool)`

GetOidcOk returns a tuple with the Oidc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidc

`func (o *AuthMethods) SetOidc(v string)`

SetOidc sets Oidc field to given value.

### HasOidc

`func (o *AuthMethods) HasOidc() bool`

HasOidc returns a boolean if a field has been set.

### SetOidcNil

`func (o *AuthMethods) SetOidcNil(b bool)`

 SetOidcNil sets the value for Oidc to be an explicit nil

### UnsetOidc
`func (o *AuthMethods) UnsetOidc()`

UnsetOidc ensures that no value is present for Oidc, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


