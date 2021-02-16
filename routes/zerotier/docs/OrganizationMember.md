# OrganizationMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** | Organization ID | [optional] [readonly] 
**UserId** | Pointer to **string** | User ID | [optional] 
**Name** | Pointer to **NullableString** | Organization member display name | [optional] [readonly] 
**Email** | Pointer to **NullableString** | Organization member email address | [optional] [readonly] 

## Methods

### NewOrganizationMember

`func NewOrganizationMember() *OrganizationMember`

NewOrganizationMember instantiates a new OrganizationMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMemberWithDefaults

`func NewOrganizationMemberWithDefaults() *OrganizationMember`

NewOrganizationMemberWithDefaults instantiates a new OrganizationMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *OrganizationMember) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrganizationMember) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrganizationMember) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *OrganizationMember) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetUserId

`func (o *OrganizationMember) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationMember) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationMember) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OrganizationMember) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetName

`func (o *OrganizationMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationMember) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrganizationMember) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationMember) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *OrganizationMember) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationMember) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationMember) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationMember) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *OrganizationMember) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrganizationMember) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


