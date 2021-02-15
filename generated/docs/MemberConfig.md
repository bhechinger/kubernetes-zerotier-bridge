# MemberConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveBridge** | Pointer to **bool** | Allow the member to be a bridge on the network | [optional] 
**Authorized** | Pointer to **bool** | Is the member authorized on the network | [optional] 
**Capabilities** | Pointer to **[]int32** |  | [optional] 
**CreationTime** | Pointer to **int32** | Time the member was created or first tried to join the network | [optional] [readonly] 
**Id** | Pointer to **string** | ID of the member node.  This is the 10 digit identifier that identifies a ZeroTier node. | [optional] [readonly] 
**Identity** | Pointer to **string** | Public Key of the member&#39;s Identity | [optional] [readonly] 
**IpAssignments** | Pointer to **[]string** | List of assigned IP addresses | [optional] 
**LastAuthorizedTime** | Pointer to **int32** | Time the member was authorized on the network | [optional] [readonly] 
**LastDeauthorizedTime** | Pointer to **int32** | Time the member was deauthorized on the network | [optional] [readonly] 
**NoAutoAssignIps** | Pointer to **bool** | Exempt this member from the IP auto assignment pool on a Network | [optional] 
**Revision** | Pointer to **int32** | Member record revision count | [optional] [readonly] 
**Tags** | Pointer to **[][]int32** | Array of 2 member tuples of tag [ID, tag value] | [optional] 
**VMajor** | Pointer to **int32** | Major version of the client | [optional] [readonly] 
**VMinor** | Pointer to **int32** | Minor version of the client | [optional] [readonly] 
**VRev** | Pointer to **int32** | Revision number of the client | [optional] [readonly] 
**VProto** | Pointer to **int32** | Protocol version of the client | [optional] [readonly] 

## Methods

### NewMemberConfig

`func NewMemberConfig() *MemberConfig`

NewMemberConfig instantiates a new MemberConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberConfigWithDefaults

`func NewMemberConfigWithDefaults() *MemberConfig`

NewMemberConfigWithDefaults instantiates a new MemberConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveBridge

`func (o *MemberConfig) GetActiveBridge() bool`

GetActiveBridge returns the ActiveBridge field if non-nil, zero value otherwise.

### GetActiveBridgeOk

`func (o *MemberConfig) GetActiveBridgeOk() (*bool, bool)`

GetActiveBridgeOk returns a tuple with the ActiveBridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveBridge

`func (o *MemberConfig) SetActiveBridge(v bool)`

SetActiveBridge sets ActiveBridge field to given value.

### HasActiveBridge

`func (o *MemberConfig) HasActiveBridge() bool`

HasActiveBridge returns a boolean if a field has been set.

### GetAuthorized

`func (o *MemberConfig) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *MemberConfig) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *MemberConfig) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *MemberConfig) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### GetCapabilities

`func (o *MemberConfig) GetCapabilities() []int32`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *MemberConfig) GetCapabilitiesOk() (*[]int32, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *MemberConfig) SetCapabilities(v []int32)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *MemberConfig) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCreationTime

`func (o *MemberConfig) GetCreationTime() int32`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *MemberConfig) GetCreationTimeOk() (*int32, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *MemberConfig) SetCreationTime(v int32)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *MemberConfig) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetId

`func (o *MemberConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MemberConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentity

`func (o *MemberConfig) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *MemberConfig) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *MemberConfig) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *MemberConfig) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIpAssignments

`func (o *MemberConfig) GetIpAssignments() []string`

GetIpAssignments returns the IpAssignments field if non-nil, zero value otherwise.

### GetIpAssignmentsOk

`func (o *MemberConfig) GetIpAssignmentsOk() (*[]string, bool)`

GetIpAssignmentsOk returns a tuple with the IpAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAssignments

`func (o *MemberConfig) SetIpAssignments(v []string)`

SetIpAssignments sets IpAssignments field to given value.

### HasIpAssignments

`func (o *MemberConfig) HasIpAssignments() bool`

HasIpAssignments returns a boolean if a field has been set.

### GetLastAuthorizedTime

`func (o *MemberConfig) GetLastAuthorizedTime() int32`

GetLastAuthorizedTime returns the LastAuthorizedTime field if non-nil, zero value otherwise.

### GetLastAuthorizedTimeOk

`func (o *MemberConfig) GetLastAuthorizedTimeOk() (*int32, bool)`

GetLastAuthorizedTimeOk returns a tuple with the LastAuthorizedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthorizedTime

`func (o *MemberConfig) SetLastAuthorizedTime(v int32)`

SetLastAuthorizedTime sets LastAuthorizedTime field to given value.

### HasLastAuthorizedTime

`func (o *MemberConfig) HasLastAuthorizedTime() bool`

HasLastAuthorizedTime returns a boolean if a field has been set.

### GetLastDeauthorizedTime

`func (o *MemberConfig) GetLastDeauthorizedTime() int32`

GetLastDeauthorizedTime returns the LastDeauthorizedTime field if non-nil, zero value otherwise.

### GetLastDeauthorizedTimeOk

`func (o *MemberConfig) GetLastDeauthorizedTimeOk() (*int32, bool)`

GetLastDeauthorizedTimeOk returns a tuple with the LastDeauthorizedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeauthorizedTime

`func (o *MemberConfig) SetLastDeauthorizedTime(v int32)`

SetLastDeauthorizedTime sets LastDeauthorizedTime field to given value.

### HasLastDeauthorizedTime

`func (o *MemberConfig) HasLastDeauthorizedTime() bool`

HasLastDeauthorizedTime returns a boolean if a field has been set.

### GetNoAutoAssignIps

`func (o *MemberConfig) GetNoAutoAssignIps() bool`

GetNoAutoAssignIps returns the NoAutoAssignIps field if non-nil, zero value otherwise.

### GetNoAutoAssignIpsOk

`func (o *MemberConfig) GetNoAutoAssignIpsOk() (*bool, bool)`

GetNoAutoAssignIpsOk returns a tuple with the NoAutoAssignIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAutoAssignIps

`func (o *MemberConfig) SetNoAutoAssignIps(v bool)`

SetNoAutoAssignIps sets NoAutoAssignIps field to given value.

### HasNoAutoAssignIps

`func (o *MemberConfig) HasNoAutoAssignIps() bool`

HasNoAutoAssignIps returns a boolean if a field has been set.

### GetRevision

`func (o *MemberConfig) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *MemberConfig) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *MemberConfig) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *MemberConfig) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetTags

`func (o *MemberConfig) GetTags() [][]int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MemberConfig) GetTagsOk() (*[][]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MemberConfig) SetTags(v [][]int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MemberConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVMajor

`func (o *MemberConfig) GetVMajor() int32`

GetVMajor returns the VMajor field if non-nil, zero value otherwise.

### GetVMajorOk

`func (o *MemberConfig) GetVMajorOk() (*int32, bool)`

GetVMajorOk returns a tuple with the VMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMajor

`func (o *MemberConfig) SetVMajor(v int32)`

SetVMajor sets VMajor field to given value.

### HasVMajor

`func (o *MemberConfig) HasVMajor() bool`

HasVMajor returns a boolean if a field has been set.

### GetVMinor

`func (o *MemberConfig) GetVMinor() int32`

GetVMinor returns the VMinor field if non-nil, zero value otherwise.

### GetVMinorOk

`func (o *MemberConfig) GetVMinorOk() (*int32, bool)`

GetVMinorOk returns a tuple with the VMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVMinor

`func (o *MemberConfig) SetVMinor(v int32)`

SetVMinor sets VMinor field to given value.

### HasVMinor

`func (o *MemberConfig) HasVMinor() bool`

HasVMinor returns a boolean if a field has been set.

### GetVRev

`func (o *MemberConfig) GetVRev() int32`

GetVRev returns the VRev field if non-nil, zero value otherwise.

### GetVRevOk

`func (o *MemberConfig) GetVRevOk() (*int32, bool)`

GetVRevOk returns a tuple with the VRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVRev

`func (o *MemberConfig) SetVRev(v int32)`

SetVRev sets VRev field to given value.

### HasVRev

`func (o *MemberConfig) HasVRev() bool`

HasVRev returns a boolean if a field has been set.

### GetVProto

`func (o *MemberConfig) GetVProto() int32`

GetVProto returns the VProto field if non-nil, zero value otherwise.

### GetVProtoOk

`func (o *MemberConfig) GetVProtoOk() (*int32, bool)`

GetVProtoOk returns a tuple with the VProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVProto

`func (o *MemberConfig) SetVProto(v int32)`

SetVProto sets VProto field to given value.

### HasVProto

`func (o *MemberConfig) HasVProto() bool`

HasVProto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


