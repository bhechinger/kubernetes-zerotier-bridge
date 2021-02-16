# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | concatenation of network ID and member ID | [optional] [readonly] 
**Clock** | Pointer to **int32** |  | [optional] [readonly] 
**NetworkId** | Pointer to **string** |  | [optional] [readonly] 
**NodeId** | Pointer to **string** | ZeroTier ID of the member | [optional] [readonly] 
**ControllerId** | Pointer to **string** |  | [optional] [readonly] 
**Hidden** | Pointer to **bool** | Whether or not the member is hidden in the UI | [optional] 
**Name** | Pointer to **string** | User defined name of the member | [optional] 
**Description** | Pointer to **string** | User defined description of the member | [optional] 
**Config** | Pointer to [**MemberConfig**](MemberConfig.md) |  | [optional] 
**LastOnline** | Pointer to **int32** | Last seen time of the member | [optional] [readonly] 
**PhysicalAddress** | Pointer to **string** | IP address the member last spoke to the controller via | [optional] [readonly] 
**ClientVersion** | Pointer to **string** | ZeroTier version the member is running | [optional] [readonly] 
**ProtocolVersion** | Pointer to **int32** | ZeroTier protocol version | [optional] [readonly] 
**SupportsRulesEngine** | Pointer to **bool** | Whether or not the client version is new enough to support the rules engine (1.4.0+) | [optional] [readonly] 

## Methods

### NewMember

`func NewMember() *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Member) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Member) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Member) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Member) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClock

`func (o *Member) GetClock() int32`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *Member) GetClockOk() (*int32, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *Member) SetClock(v int32)`

SetClock sets Clock field to given value.

### HasClock

`func (o *Member) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetNetworkId

`func (o *Member) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *Member) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *Member) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *Member) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNodeId

`func (o *Member) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Member) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Member) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *Member) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetControllerId

`func (o *Member) GetControllerId() string`

GetControllerId returns the ControllerId field if non-nil, zero value otherwise.

### GetControllerIdOk

`func (o *Member) GetControllerIdOk() (*string, bool)`

GetControllerIdOk returns a tuple with the ControllerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerId

`func (o *Member) SetControllerId(v string)`

SetControllerId sets ControllerId field to given value.

### HasControllerId

`func (o *Member) HasControllerId() bool`

HasControllerId returns a boolean if a field has been set.

### GetHidden

`func (o *Member) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *Member) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *Member) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *Member) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetName

`func (o *Member) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Member) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Member) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Member) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Member) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Member) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Member) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Member) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfig

`func (o *Member) GetConfig() MemberConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Member) GetConfigOk() (*MemberConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Member) SetConfig(v MemberConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Member) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLastOnline

`func (o *Member) GetLastOnline() int32`

GetLastOnline returns the LastOnline field if non-nil, zero value otherwise.

### GetLastOnlineOk

`func (o *Member) GetLastOnlineOk() (*int32, bool)`

GetLastOnlineOk returns a tuple with the LastOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnline

`func (o *Member) SetLastOnline(v int32)`

SetLastOnline sets LastOnline field to given value.

### HasLastOnline

`func (o *Member) HasLastOnline() bool`

HasLastOnline returns a boolean if a field has been set.

### GetPhysicalAddress

`func (o *Member) GetPhysicalAddress() string`

GetPhysicalAddress returns the PhysicalAddress field if non-nil, zero value otherwise.

### GetPhysicalAddressOk

`func (o *Member) GetPhysicalAddressOk() (*string, bool)`

GetPhysicalAddressOk returns a tuple with the PhysicalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalAddress

`func (o *Member) SetPhysicalAddress(v string)`

SetPhysicalAddress sets PhysicalAddress field to given value.

### HasPhysicalAddress

`func (o *Member) HasPhysicalAddress() bool`

HasPhysicalAddress returns a boolean if a field has been set.

### GetClientVersion

`func (o *Member) GetClientVersion() string`

GetClientVersion returns the ClientVersion field if non-nil, zero value otherwise.

### GetClientVersionOk

`func (o *Member) GetClientVersionOk() (*string, bool)`

GetClientVersionOk returns a tuple with the ClientVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVersion

`func (o *Member) SetClientVersion(v string)`

SetClientVersion sets ClientVersion field to given value.

### HasClientVersion

`func (o *Member) HasClientVersion() bool`

HasClientVersion returns a boolean if a field has been set.

### GetProtocolVersion

`func (o *Member) GetProtocolVersion() int32`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *Member) GetProtocolVersionOk() (*int32, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *Member) SetProtocolVersion(v int32)`

SetProtocolVersion sets ProtocolVersion field to given value.

### HasProtocolVersion

`func (o *Member) HasProtocolVersion() bool`

HasProtocolVersion returns a boolean if a field has been set.

### GetSupportsRulesEngine

`func (o *Member) GetSupportsRulesEngine() bool`

GetSupportsRulesEngine returns the SupportsRulesEngine field if non-nil, zero value otherwise.

### GetSupportsRulesEngineOk

`func (o *Member) GetSupportsRulesEngineOk() (*bool, bool)`

GetSupportsRulesEngineOk returns a tuple with the SupportsRulesEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRulesEngine

`func (o *Member) SetSupportsRulesEngine(v bool)`

SetSupportsRulesEngine sets SupportsRulesEngine field to given value.

### HasSupportsRulesEngine

`func (o *Member) HasSupportsRulesEngine() bool`

HasSupportsRulesEngine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


