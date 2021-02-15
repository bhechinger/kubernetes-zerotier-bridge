# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Organization ID | [optional] 
**OwnerId** | Pointer to **string** | User ID of the organization owner | [optional] [readonly] 
**OwnerEmail** | Pointer to **string** | Organization owner&#39;s email address | [optional] [readonly] 
**Members** | Pointer to [**[]OrganizationMember**](OrganizationMember.md) | List of organization members | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Organization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwnerId

`func (o *Organization) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Organization) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Organization) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Organization) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *Organization) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *Organization) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *Organization) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *Organization) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetMembers

`func (o *Organization) GetMembers() []OrganizationMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Organization) GetMembersOk() (*[]OrganizationMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Organization) SetMembers(v []OrganizationMember)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Organization) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


