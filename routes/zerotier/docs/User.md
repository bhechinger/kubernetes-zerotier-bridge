# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | User ID | [optional] [readonly] 
**OrgId** | Pointer to **string** | Organization ID | [optional] [readonly] 
**GlobalPermissions** | Pointer to [**Permissions**](Permissions.md) |  | [optional] [readonly] 
**DisplayName** | Pointer to **string** | Display Name | [optional] 
**Email** | Pointer to **string** | User email address | [optional] [readonly] 
**Auth** | Pointer to [**AuthMethods**](AuthMethods.md) |  | [optional] [readonly] 
**SmsNumber** | Pointer to **string** | SMS number | [optional] 
**Tokens** | Pointer to **[]string** | List of API token names. | [optional] [readonly] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *User) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *User) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *User) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *User) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetGlobalPermissions

`func (o *User) GetGlobalPermissions() Permissions`

GetGlobalPermissions returns the GlobalPermissions field if non-nil, zero value otherwise.

### GetGlobalPermissionsOk

`func (o *User) GetGlobalPermissionsOk() (*Permissions, bool)`

GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPermissions

`func (o *User) SetGlobalPermissions(v Permissions)`

SetGlobalPermissions sets GlobalPermissions field to given value.

### HasGlobalPermissions

`func (o *User) HasGlobalPermissions() bool`

HasGlobalPermissions returns a boolean if a field has been set.

### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *User) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAuth

`func (o *User) GetAuth() AuthMethods`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *User) GetAuthOk() (*AuthMethods, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *User) SetAuth(v AuthMethods)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *User) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetSmsNumber

`func (o *User) GetSmsNumber() string`

GetSmsNumber returns the SmsNumber field if non-nil, zero value otherwise.

### GetSmsNumberOk

`func (o *User) GetSmsNumberOk() (*string, bool)`

GetSmsNumberOk returns a tuple with the SmsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsNumber

`func (o *User) SetSmsNumber(v string)`

SetSmsNumber sets SmsNumber field to given value.

### HasSmsNumber

`func (o *User) HasSmsNumber() bool`

HasSmsNumber returns a boolean if a field has been set.

### GetTokens

`func (o *User) GetTokens() []string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *User) GetTokensOk() (*[]string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *User) SetTokens(v []string)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *User) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


