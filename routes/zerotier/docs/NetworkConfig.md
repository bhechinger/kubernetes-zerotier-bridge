# NetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Network ID | [optional] 
**CreationTime** | Pointer to **int32** | Time the network was created | [optional] 
**Capabilities** | Pointer to **[]map[string]interface{}** | Array of network capabilities | [optional] 
**EnableBroadcast** | Pointer to **bool** | Enable broadcast packets on the network | [optional] 
**IpAssignmentPools** | Pointer to [**[]IPRange**](IPRange.md) | Range of IP addresses for the auto assign pool | [optional] 
**LastModified** | Pointer to **int32** | Time the network was last modified | [optional] [readonly] 
**Mtu** | Pointer to **int32** | MTU to set on the client virtual network adapter | [optional] 
**MulticastLimit** | Pointer to **int32** | Maximum number of recipients per multicast or broadcast. Warning - Setting this to 0 will disable IPv4 communication on your network! | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **bool** | Whether or not the network is private.  If false, members will *NOT* need to be authorized to join. | [optional] 
**Routes** | Pointer to [**[]Route**](Route.md) |  | [optional] 
**Rules** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**V4AssignMode** | Pointer to [**IPV4AssignMode**](IPV4AssignMode.md) |  | [optional] 
**V6AssignMode** | Pointer to [**IPV6AssignMode**](IPV6AssignMode.md) |  | [optional] 

## Methods

### NewNetworkConfig

`func NewNetworkConfig() *NetworkConfig`

NewNetworkConfig instantiates a new NetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkConfigWithDefaults

`func NewNetworkConfigWithDefaults() *NetworkConfig`

NewNetworkConfigWithDefaults instantiates a new NetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreationTime

`func (o *NetworkConfig) GetCreationTime() int32`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *NetworkConfig) GetCreationTimeOk() (*int32, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *NetworkConfig) SetCreationTime(v int32)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *NetworkConfig) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCapabilities

`func (o *NetworkConfig) GetCapabilities() []map[string]interface{}`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *NetworkConfig) GetCapabilitiesOk() (*[]map[string]interface{}, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *NetworkConfig) SetCapabilities(v []map[string]interface{})`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *NetworkConfig) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetEnableBroadcast

`func (o *NetworkConfig) GetEnableBroadcast() bool`

GetEnableBroadcast returns the EnableBroadcast field if non-nil, zero value otherwise.

### GetEnableBroadcastOk

`func (o *NetworkConfig) GetEnableBroadcastOk() (*bool, bool)`

GetEnableBroadcastOk returns a tuple with the EnableBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableBroadcast

`func (o *NetworkConfig) SetEnableBroadcast(v bool)`

SetEnableBroadcast sets EnableBroadcast field to given value.

### HasEnableBroadcast

`func (o *NetworkConfig) HasEnableBroadcast() bool`

HasEnableBroadcast returns a boolean if a field has been set.

### GetIpAssignmentPools

`func (o *NetworkConfig) GetIpAssignmentPools() []IPRange`

GetIpAssignmentPools returns the IpAssignmentPools field if non-nil, zero value otherwise.

### GetIpAssignmentPoolsOk

`func (o *NetworkConfig) GetIpAssignmentPoolsOk() (*[]IPRange, bool)`

GetIpAssignmentPoolsOk returns a tuple with the IpAssignmentPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAssignmentPools

`func (o *NetworkConfig) SetIpAssignmentPools(v []IPRange)`

SetIpAssignmentPools sets IpAssignmentPools field to given value.

### HasIpAssignmentPools

`func (o *NetworkConfig) HasIpAssignmentPools() bool`

HasIpAssignmentPools returns a boolean if a field has been set.

### GetLastModified

`func (o *NetworkConfig) GetLastModified() int32`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *NetworkConfig) GetLastModifiedOk() (*int32, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *NetworkConfig) SetLastModified(v int32)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *NetworkConfig) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetMtu

`func (o *NetworkConfig) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *NetworkConfig) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *NetworkConfig) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *NetworkConfig) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetMulticastLimit

`func (o *NetworkConfig) GetMulticastLimit() int32`

GetMulticastLimit returns the MulticastLimit field if non-nil, zero value otherwise.

### GetMulticastLimitOk

`func (o *NetworkConfig) GetMulticastLimitOk() (*int32, bool)`

GetMulticastLimitOk returns a tuple with the MulticastLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastLimit

`func (o *NetworkConfig) SetMulticastLimit(v int32)`

SetMulticastLimit sets MulticastLimit field to given value.

### HasMulticastLimit

`func (o *NetworkConfig) HasMulticastLimit() bool`

HasMulticastLimit returns a boolean if a field has been set.

### GetName

`func (o *NetworkConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivate

`func (o *NetworkConfig) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *NetworkConfig) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *NetworkConfig) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *NetworkConfig) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetRoutes

`func (o *NetworkConfig) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *NetworkConfig) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *NetworkConfig) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.

### HasRoutes

`func (o *NetworkConfig) HasRoutes() bool`

HasRoutes returns a boolean if a field has been set.

### GetRules

`func (o *NetworkConfig) GetRules() []map[string]interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *NetworkConfig) GetRulesOk() (*[]map[string]interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *NetworkConfig) SetRules(v []map[string]interface{})`

SetRules sets Rules field to given value.

### HasRules

`func (o *NetworkConfig) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetTags

`func (o *NetworkConfig) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NetworkConfig) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NetworkConfig) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *NetworkConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetV4AssignMode

`func (o *NetworkConfig) GetV4AssignMode() IPV4AssignMode`

GetV4AssignMode returns the V4AssignMode field if non-nil, zero value otherwise.

### GetV4AssignModeOk

`func (o *NetworkConfig) GetV4AssignModeOk() (*IPV4AssignMode, bool)`

GetV4AssignModeOk returns a tuple with the V4AssignMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV4AssignMode

`func (o *NetworkConfig) SetV4AssignMode(v IPV4AssignMode)`

SetV4AssignMode sets V4AssignMode field to given value.

### HasV4AssignMode

`func (o *NetworkConfig) HasV4AssignMode() bool`

HasV4AssignMode returns a boolean if a field has been set.

### GetV6AssignMode

`func (o *NetworkConfig) GetV6AssignMode() IPV6AssignMode`

GetV6AssignMode returns the V6AssignMode field if non-nil, zero value otherwise.

### GetV6AssignModeOk

`func (o *NetworkConfig) GetV6AssignModeOk() (*IPV6AssignMode, bool)`

GetV6AssignModeOk returns a tuple with the V6AssignMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6AssignMode

`func (o *NetworkConfig) SetV6AssignMode(v IPV6AssignMode)`

SetV6AssignMode sets V6AssignMode field to given value.

### HasV6AssignMode

`func (o *NetworkConfig) HasV6AssignMode() bool`

HasV6AssignMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


