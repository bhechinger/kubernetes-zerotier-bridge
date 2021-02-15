# Network

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Clock** | Pointer to **int32** |  | [optional] [readonly] 
**Config** | Pointer to [**NetworkConfig**](NetworkConfig.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RulesSource** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**PermissionsMap**](PermissionsMap.md) |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OnlineMemberCount** | Pointer to **int32** |  | [optional] [readonly] 
**AuthorizedMemberCount** | Pointer to **int32** |  | [optional] [readonly] 
**TotalMemberCount** | Pointer to **int32** |  | [optional] [readonly] 
**CapabilitiesByName** | Pointer to **map[string]interface{}** |  | [optional] 
**TagsByName** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewNetwork

`func NewNetwork() *Network`

NewNetwork instantiates a new Network object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkWithDefaults

`func NewNetworkWithDefaults() *Network`

NewNetworkWithDefaults instantiates a new Network object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Network) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Network) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Network) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Network) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClock

`func (o *Network) GetClock() int32`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *Network) GetClockOk() (*int32, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *Network) SetClock(v int32)`

SetClock sets Clock field to given value.

### HasClock

`func (o *Network) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetConfig

`func (o *Network) GetConfig() NetworkConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Network) GetConfigOk() (*NetworkConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Network) SetConfig(v NetworkConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Network) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *Network) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Network) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Network) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Network) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRulesSource

`func (o *Network) GetRulesSource() string`

GetRulesSource returns the RulesSource field if non-nil, zero value otherwise.

### GetRulesSourceOk

`func (o *Network) GetRulesSourceOk() (*string, bool)`

GetRulesSourceOk returns a tuple with the RulesSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesSource

`func (o *Network) SetRulesSource(v string)`

SetRulesSource sets RulesSource field to given value.

### HasRulesSource

`func (o *Network) HasRulesSource() bool`

HasRulesSource returns a boolean if a field has been set.

### GetPermissions

`func (o *Network) GetPermissions() PermissionsMap`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Network) GetPermissionsOk() (*PermissionsMap, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Network) SetPermissions(v PermissionsMap)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *Network) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetOwnerId

`func (o *Network) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Network) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Network) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Network) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOnlineMemberCount

`func (o *Network) GetOnlineMemberCount() int32`

GetOnlineMemberCount returns the OnlineMemberCount field if non-nil, zero value otherwise.

### GetOnlineMemberCountOk

`func (o *Network) GetOnlineMemberCountOk() (*int32, bool)`

GetOnlineMemberCountOk returns a tuple with the OnlineMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineMemberCount

`func (o *Network) SetOnlineMemberCount(v int32)`

SetOnlineMemberCount sets OnlineMemberCount field to given value.

### HasOnlineMemberCount

`func (o *Network) HasOnlineMemberCount() bool`

HasOnlineMemberCount returns a boolean if a field has been set.

### GetAuthorizedMemberCount

`func (o *Network) GetAuthorizedMemberCount() int32`

GetAuthorizedMemberCount returns the AuthorizedMemberCount field if non-nil, zero value otherwise.

### GetAuthorizedMemberCountOk

`func (o *Network) GetAuthorizedMemberCountOk() (*int32, bool)`

GetAuthorizedMemberCountOk returns a tuple with the AuthorizedMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedMemberCount

`func (o *Network) SetAuthorizedMemberCount(v int32)`

SetAuthorizedMemberCount sets AuthorizedMemberCount field to given value.

### HasAuthorizedMemberCount

`func (o *Network) HasAuthorizedMemberCount() bool`

HasAuthorizedMemberCount returns a boolean if a field has been set.

### GetTotalMemberCount

`func (o *Network) GetTotalMemberCount() int32`

GetTotalMemberCount returns the TotalMemberCount field if non-nil, zero value otherwise.

### GetTotalMemberCountOk

`func (o *Network) GetTotalMemberCountOk() (*int32, bool)`

GetTotalMemberCountOk returns a tuple with the TotalMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMemberCount

`func (o *Network) SetTotalMemberCount(v int32)`

SetTotalMemberCount sets TotalMemberCount field to given value.

### HasTotalMemberCount

`func (o *Network) HasTotalMemberCount() bool`

HasTotalMemberCount returns a boolean if a field has been set.

### GetCapabilitiesByName

`func (o *Network) GetCapabilitiesByName() map[string]interface{}`

GetCapabilitiesByName returns the CapabilitiesByName field if non-nil, zero value otherwise.

### GetCapabilitiesByNameOk

`func (o *Network) GetCapabilitiesByNameOk() (*map[string]interface{}, bool)`

GetCapabilitiesByNameOk returns a tuple with the CapabilitiesByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilitiesByName

`func (o *Network) SetCapabilitiesByName(v map[string]interface{})`

SetCapabilitiesByName sets CapabilitiesByName field to given value.

### HasCapabilitiesByName

`func (o *Network) HasCapabilitiesByName() bool`

HasCapabilitiesByName returns a boolean if a field has been set.

### GetTagsByName

`func (o *Network) GetTagsByName() map[string]interface{}`

GetTagsByName returns the TagsByName field if non-nil, zero value otherwise.

### GetTagsByNameOk

`func (o *Network) GetTagsByNameOk() (*map[string]interface{}, bool)`

GetTagsByNameOk returns a tuple with the TagsByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsByName

`func (o *Network) SetTagsByName(v map[string]interface{})`

SetTagsByName sets TagsByName field to given value.

### HasTagsByName

`func (o *Network) HasTagsByName() bool`

HasTagsByName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


