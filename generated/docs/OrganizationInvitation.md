# OrganizationInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | Pointer to **string** | Organization ID | [optional] [readonly] 
**Email** | Pointer to **string** | Email address of invitee | [optional] 
**Id** | Pointer to **string** | Invitation ID | [optional] [readonly] 
**CreationTime** | Pointer to **int32** | Creation time of the invite | [optional] [readonly] 
**Status** | Pointer to [**InviteStatus**](InviteStatus.md) | Invitation status | [optional] [readonly] 
**UpdateTime** | Pointer to **int32** | Last updated time of the invitation | [optional] [readonly] 
**OwnerEmail** | Pointer to **string** | Organization owner email address | [optional] [readonly] 

## Methods

### NewOrganizationInvitation

`func NewOrganizationInvitation() *OrganizationInvitation`

NewOrganizationInvitation instantiates a new OrganizationInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInvitationWithDefaults

`func NewOrganizationInvitationWithDefaults() *OrganizationInvitation`

NewOrganizationInvitationWithDefaults instantiates a new OrganizationInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *OrganizationInvitation) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *OrganizationInvitation) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *OrganizationInvitation) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *OrganizationInvitation) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetEmail

`func (o *OrganizationInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationInvitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationInvitation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrganizationInvitation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetId

`func (o *OrganizationInvitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationInvitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationInvitation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrganizationInvitation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *OrganizationInvitation) GetCreationTime() int32`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *OrganizationInvitation) GetCreationTimeOk() (*int32, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *OrganizationInvitation) SetCreationTime(v int32)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *OrganizationInvitation) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetStatus

`func (o *OrganizationInvitation) GetStatus() InviteStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrganizationInvitation) GetStatusOk() (*InviteStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrganizationInvitation) SetStatus(v InviteStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrganizationInvitation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdateTime

`func (o *OrganizationInvitation) GetUpdateTime() int32`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *OrganizationInvitation) GetUpdateTimeOk() (*int32, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *OrganizationInvitation) SetUpdateTime(v int32)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *OrganizationInvitation) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *OrganizationInvitation) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *OrganizationInvitation) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *OrganizationInvitation) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *OrganizationInvitation) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


