/*
 * ZeroTier Central API
 *
 * ZeroTier Central Network Management Portal API.<p>All API requests must have an API token header specified in the <code>Authorization: Bearer xxxxx</code> format.  You can generate your API key by logging into <a href=\"https://my.zerotier.com\">ZeroTier Central</a> and creating a token on the Account page.</p><p>eg. <code>curl -X GET -H \"Authorization: bearer xxxxx\" https://my.zerotier.com/api/network</code></p>
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Network Network object
type Network struct {
	Id *string `json:"id,omitempty"`
	Clock *int32 `json:"clock,omitempty"`
	Config *NetworkConfig `json:"config,omitempty"`
	Description *string `json:"description,omitempty"`
	RulesSource *string `json:"rulesSource,omitempty"`
	Permissions *PermissionsMap `json:"permissions,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	OnlineMemberCount *int32 `json:"onlineMemberCount,omitempty"`
	AuthorizedMemberCount *int32 `json:"authorizedMemberCount,omitempty"`
	TotalMemberCount *int32 `json:"totalMemberCount,omitempty"`
	CapabilitiesByName *map[string]interface{} `json:"capabilitiesByName,omitempty"`
	TagsByName *map[string]interface{} `json:"tagsByName,omitempty"`
}

// NewNetwork instantiates a new Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetwork() *Network {
	this := Network{}
	return &this
}

// NewNetworkWithDefaults instantiates a new Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkWithDefaults() *Network {
	this := Network{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Network) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Network) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Network) SetId(v string) {
	o.Id = &v
}

// GetClock returns the Clock field value if set, zero value otherwise.
func (o *Network) GetClock() int32 {
	if o == nil || o.Clock == nil {
		var ret int32
		return ret
	}
	return *o.Clock
}

// GetClockOk returns a tuple with the Clock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetClockOk() (*int32, bool) {
	if o == nil || o.Clock == nil {
		return nil, false
	}
	return o.Clock, true
}

// HasClock returns a boolean if a field has been set.
func (o *Network) HasClock() bool {
	if o != nil && o.Clock != nil {
		return true
	}

	return false
}

// SetClock gets a reference to the given int32 and assigns it to the Clock field.
func (o *Network) SetClock(v int32) {
	o.Clock = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Network) GetConfig() NetworkConfig {
	if o == nil || o.Config == nil {
		var ret NetworkConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetConfigOk() (*NetworkConfig, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Network) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given NetworkConfig and assigns it to the Config field.
func (o *Network) SetConfig(v NetworkConfig) {
	o.Config = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Network) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Network) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Network) SetDescription(v string) {
	o.Description = &v
}

// GetRulesSource returns the RulesSource field value if set, zero value otherwise.
func (o *Network) GetRulesSource() string {
	if o == nil || o.RulesSource == nil {
		var ret string
		return ret
	}
	return *o.RulesSource
}

// GetRulesSourceOk returns a tuple with the RulesSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetRulesSourceOk() (*string, bool) {
	if o == nil || o.RulesSource == nil {
		return nil, false
	}
	return o.RulesSource, true
}

// HasRulesSource returns a boolean if a field has been set.
func (o *Network) HasRulesSource() bool {
	if o != nil && o.RulesSource != nil {
		return true
	}

	return false
}

// SetRulesSource gets a reference to the given string and assigns it to the RulesSource field.
func (o *Network) SetRulesSource(v string) {
	o.RulesSource = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *Network) GetPermissions() PermissionsMap {
	if o == nil || o.Permissions == nil {
		var ret PermissionsMap
		return ret
	}
	return *o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetPermissionsOk() (*PermissionsMap, bool) {
	if o == nil || o.Permissions == nil {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *Network) HasPermissions() bool {
	if o != nil && o.Permissions != nil {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given PermissionsMap and assigns it to the Permissions field.
func (o *Network) SetPermissions(v PermissionsMap) {
	o.Permissions = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *Network) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *Network) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *Network) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOnlineMemberCount returns the OnlineMemberCount field value if set, zero value otherwise.
func (o *Network) GetOnlineMemberCount() int32 {
	if o == nil || o.OnlineMemberCount == nil {
		var ret int32
		return ret
	}
	return *o.OnlineMemberCount
}

// GetOnlineMemberCountOk returns a tuple with the OnlineMemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetOnlineMemberCountOk() (*int32, bool) {
	if o == nil || o.OnlineMemberCount == nil {
		return nil, false
	}
	return o.OnlineMemberCount, true
}

// HasOnlineMemberCount returns a boolean if a field has been set.
func (o *Network) HasOnlineMemberCount() bool {
	if o != nil && o.OnlineMemberCount != nil {
		return true
	}

	return false
}

// SetOnlineMemberCount gets a reference to the given int32 and assigns it to the OnlineMemberCount field.
func (o *Network) SetOnlineMemberCount(v int32) {
	o.OnlineMemberCount = &v
}

// GetAuthorizedMemberCount returns the AuthorizedMemberCount field value if set, zero value otherwise.
func (o *Network) GetAuthorizedMemberCount() int32 {
	if o == nil || o.AuthorizedMemberCount == nil {
		var ret int32
		return ret
	}
	return *o.AuthorizedMemberCount
}

// GetAuthorizedMemberCountOk returns a tuple with the AuthorizedMemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetAuthorizedMemberCountOk() (*int32, bool) {
	if o == nil || o.AuthorizedMemberCount == nil {
		return nil, false
	}
	return o.AuthorizedMemberCount, true
}

// HasAuthorizedMemberCount returns a boolean if a field has been set.
func (o *Network) HasAuthorizedMemberCount() bool {
	if o != nil && o.AuthorizedMemberCount != nil {
		return true
	}

	return false
}

// SetAuthorizedMemberCount gets a reference to the given int32 and assigns it to the AuthorizedMemberCount field.
func (o *Network) SetAuthorizedMemberCount(v int32) {
	o.AuthorizedMemberCount = &v
}

// GetTotalMemberCount returns the TotalMemberCount field value if set, zero value otherwise.
func (o *Network) GetTotalMemberCount() int32 {
	if o == nil || o.TotalMemberCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalMemberCount
}

// GetTotalMemberCountOk returns a tuple with the TotalMemberCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetTotalMemberCountOk() (*int32, bool) {
	if o == nil || o.TotalMemberCount == nil {
		return nil, false
	}
	return o.TotalMemberCount, true
}

// HasTotalMemberCount returns a boolean if a field has been set.
func (o *Network) HasTotalMemberCount() bool {
	if o != nil && o.TotalMemberCount != nil {
		return true
	}

	return false
}

// SetTotalMemberCount gets a reference to the given int32 and assigns it to the TotalMemberCount field.
func (o *Network) SetTotalMemberCount(v int32) {
	o.TotalMemberCount = &v
}

// GetCapabilitiesByName returns the CapabilitiesByName field value if set, zero value otherwise.
func (o *Network) GetCapabilitiesByName() map[string]interface{} {
	if o == nil || o.CapabilitiesByName == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CapabilitiesByName
}

// GetCapabilitiesByNameOk returns a tuple with the CapabilitiesByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetCapabilitiesByNameOk() (*map[string]interface{}, bool) {
	if o == nil || o.CapabilitiesByName == nil {
		return nil, false
	}
	return o.CapabilitiesByName, true
}

// HasCapabilitiesByName returns a boolean if a field has been set.
func (o *Network) HasCapabilitiesByName() bool {
	if o != nil && o.CapabilitiesByName != nil {
		return true
	}

	return false
}

// SetCapabilitiesByName gets a reference to the given map[string]interface{} and assigns it to the CapabilitiesByName field.
func (o *Network) SetCapabilitiesByName(v map[string]interface{}) {
	o.CapabilitiesByName = &v
}

// GetTagsByName returns the TagsByName field value if set, zero value otherwise.
func (o *Network) GetTagsByName() map[string]interface{} {
	if o == nil || o.TagsByName == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.TagsByName
}

// GetTagsByNameOk returns a tuple with the TagsByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetTagsByNameOk() (*map[string]interface{}, bool) {
	if o == nil || o.TagsByName == nil {
		return nil, false
	}
	return o.TagsByName, true
}

// HasTagsByName returns a boolean if a field has been set.
func (o *Network) HasTagsByName() bool {
	if o != nil && o.TagsByName != nil {
		return true
	}

	return false
}

// SetTagsByName gets a reference to the given map[string]interface{} and assigns it to the TagsByName field.
func (o *Network) SetTagsByName(v map[string]interface{}) {
	o.TagsByName = &v
}

func (o Network) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Clock != nil {
		toSerialize["clock"] = o.Clock
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.RulesSource != nil {
		toSerialize["rulesSource"] = o.RulesSource
	}
	if o.Permissions != nil {
		toSerialize["permissions"] = o.Permissions
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.OnlineMemberCount != nil {
		toSerialize["onlineMemberCount"] = o.OnlineMemberCount
	}
	if o.AuthorizedMemberCount != nil {
		toSerialize["authorizedMemberCount"] = o.AuthorizedMemberCount
	}
	if o.TotalMemberCount != nil {
		toSerialize["totalMemberCount"] = o.TotalMemberCount
	}
	if o.CapabilitiesByName != nil {
		toSerialize["capabilitiesByName"] = o.CapabilitiesByName
	}
	if o.TagsByName != nil {
		toSerialize["tagsByName"] = o.TagsByName
	}
	return json.Marshal(toSerialize)
}

type NullableNetwork struct {
	value *Network
	isSet bool
}

func (v NullableNetwork) Get() *Network {
	return v.value
}

func (v *NullableNetwork) Set(val *Network) {
	v.value = val
	v.isSet = true
}

func (v NullableNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetwork(val *Network) *NullableNetwork {
	return &NullableNetwork{value: val, isSet: true}
}

func (v NullableNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


