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

// MemberConfig struct for MemberConfig
type MemberConfig struct {
	// Allow the member to be a bridge on the network
	ActiveBridge *bool `json:"activeBridge,omitempty"`
	// Is the member authorized on the network
	Authorized *bool `json:"authorized,omitempty"`
	Capabilities *[]int32 `json:"capabilities,omitempty"`
	// Time the member was created or first tried to join the network
	CreationTime *int32 `json:"creationTime,omitempty"`
	// ID of the member node.  This is the 10 digit identifier that identifies a ZeroTier node.
	Id *string `json:"id,omitempty"`
	// Public Key of the member's Identity
	Identity *string `json:"identity,omitempty"`
	// List of assigned IP addresses
	IpAssignments *[]string `json:"ipAssignments,omitempty"`
	// Time the member was authorized on the network
	LastAuthorizedTime *int32 `json:"lastAuthorizedTime,omitempty"`
	// Time the member was deauthorized on the network
	LastDeauthorizedTime *int32 `json:"lastDeauthorizedTime,omitempty"`
	// Exempt this member from the IP auto assignment pool on a Network
	NoAutoAssignIps *bool `json:"noAutoAssignIps,omitempty"`
	// Member record revision count
	Revision *int32 `json:"revision,omitempty"`
	// Array of 2 member tuples of tag [ID, tag value]
	Tags *[][]int32 `json:"tags,omitempty"`
	// Major version of the client
	VMajor *int32 `json:"vMajor,omitempty"`
	// Minor version of the client
	VMinor *int32 `json:"vMinor,omitempty"`
	// Revision number of the client
	VRev *int32 `json:"vRev,omitempty"`
	// Protocol version of the client
	VProto *int32 `json:"vProto,omitempty"`
}

// NewMemberConfig instantiates a new MemberConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberConfig() *MemberConfig {
	this := MemberConfig{}
	return &this
}

// NewMemberConfigWithDefaults instantiates a new MemberConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberConfigWithDefaults() *MemberConfig {
	this := MemberConfig{}
	return &this
}

// GetActiveBridge returns the ActiveBridge field value if set, zero value otherwise.
func (o *MemberConfig) GetActiveBridge() bool {
	if o == nil || o.ActiveBridge == nil {
		var ret bool
		return ret
	}
	return *o.ActiveBridge
}

// GetActiveBridgeOk returns a tuple with the ActiveBridge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetActiveBridgeOk() (*bool, bool) {
	if o == nil || o.ActiveBridge == nil {
		return nil, false
	}
	return o.ActiveBridge, true
}

// HasActiveBridge returns a boolean if a field has been set.
func (o *MemberConfig) HasActiveBridge() bool {
	if o != nil && o.ActiveBridge != nil {
		return true
	}

	return false
}

// SetActiveBridge gets a reference to the given bool and assigns it to the ActiveBridge field.
func (o *MemberConfig) SetActiveBridge(v bool) {
	o.ActiveBridge = &v
}

// GetAuthorized returns the Authorized field value if set, zero value otherwise.
func (o *MemberConfig) GetAuthorized() bool {
	if o == nil || o.Authorized == nil {
		var ret bool
		return ret
	}
	return *o.Authorized
}

// GetAuthorizedOk returns a tuple with the Authorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetAuthorizedOk() (*bool, bool) {
	if o == nil || o.Authorized == nil {
		return nil, false
	}
	return o.Authorized, true
}

// HasAuthorized returns a boolean if a field has been set.
func (o *MemberConfig) HasAuthorized() bool {
	if o != nil && o.Authorized != nil {
		return true
	}

	return false
}

// SetAuthorized gets a reference to the given bool and assigns it to the Authorized field.
func (o *MemberConfig) SetAuthorized(v bool) {
	o.Authorized = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *MemberConfig) GetCapabilities() []int32 {
	if o == nil || o.Capabilities == nil {
		var ret []int32
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetCapabilitiesOk() (*[]int32, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *MemberConfig) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []int32 and assigns it to the Capabilities field.
func (o *MemberConfig) SetCapabilities(v []int32) {
	o.Capabilities = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *MemberConfig) GetCreationTime() int32 {
	if o == nil || o.CreationTime == nil {
		var ret int32
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetCreationTimeOk() (*int32, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *MemberConfig) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given int32 and assigns it to the CreationTime field.
func (o *MemberConfig) SetCreationTime(v int32) {
	o.CreationTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MemberConfig) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MemberConfig) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MemberConfig) SetId(v string) {
	o.Id = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *MemberConfig) GetIdentity() string {
	if o == nil || o.Identity == nil {
		var ret string
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetIdentityOk() (*string, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *MemberConfig) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given string and assigns it to the Identity field.
func (o *MemberConfig) SetIdentity(v string) {
	o.Identity = &v
}

// GetIpAssignments returns the IpAssignments field value if set, zero value otherwise.
func (o *MemberConfig) GetIpAssignments() []string {
	if o == nil || o.IpAssignments == nil {
		var ret []string
		return ret
	}
	return *o.IpAssignments
}

// GetIpAssignmentsOk returns a tuple with the IpAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetIpAssignmentsOk() (*[]string, bool) {
	if o == nil || o.IpAssignments == nil {
		return nil, false
	}
	return o.IpAssignments, true
}

// HasIpAssignments returns a boolean if a field has been set.
func (o *MemberConfig) HasIpAssignments() bool {
	if o != nil && o.IpAssignments != nil {
		return true
	}

	return false
}

// SetIpAssignments gets a reference to the given []string and assigns it to the IpAssignments field.
func (o *MemberConfig) SetIpAssignments(v []string) {
	o.IpAssignments = &v
}

// GetLastAuthorizedTime returns the LastAuthorizedTime field value if set, zero value otherwise.
func (o *MemberConfig) GetLastAuthorizedTime() int32 {
	if o == nil || o.LastAuthorizedTime == nil {
		var ret int32
		return ret
	}
	return *o.LastAuthorizedTime
}

// GetLastAuthorizedTimeOk returns a tuple with the LastAuthorizedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetLastAuthorizedTimeOk() (*int32, bool) {
	if o == nil || o.LastAuthorizedTime == nil {
		return nil, false
	}
	return o.LastAuthorizedTime, true
}

// HasLastAuthorizedTime returns a boolean if a field has been set.
func (o *MemberConfig) HasLastAuthorizedTime() bool {
	if o != nil && o.LastAuthorizedTime != nil {
		return true
	}

	return false
}

// SetLastAuthorizedTime gets a reference to the given int32 and assigns it to the LastAuthorizedTime field.
func (o *MemberConfig) SetLastAuthorizedTime(v int32) {
	o.LastAuthorizedTime = &v
}

// GetLastDeauthorizedTime returns the LastDeauthorizedTime field value if set, zero value otherwise.
func (o *MemberConfig) GetLastDeauthorizedTime() int32 {
	if o == nil || o.LastDeauthorizedTime == nil {
		var ret int32
		return ret
	}
	return *o.LastDeauthorizedTime
}

// GetLastDeauthorizedTimeOk returns a tuple with the LastDeauthorizedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetLastDeauthorizedTimeOk() (*int32, bool) {
	if o == nil || o.LastDeauthorizedTime == nil {
		return nil, false
	}
	return o.LastDeauthorizedTime, true
}

// HasLastDeauthorizedTime returns a boolean if a field has been set.
func (o *MemberConfig) HasLastDeauthorizedTime() bool {
	if o != nil && o.LastDeauthorizedTime != nil {
		return true
	}

	return false
}

// SetLastDeauthorizedTime gets a reference to the given int32 and assigns it to the LastDeauthorizedTime field.
func (o *MemberConfig) SetLastDeauthorizedTime(v int32) {
	o.LastDeauthorizedTime = &v
}

// GetNoAutoAssignIps returns the NoAutoAssignIps field value if set, zero value otherwise.
func (o *MemberConfig) GetNoAutoAssignIps() bool {
	if o == nil || o.NoAutoAssignIps == nil {
		var ret bool
		return ret
	}
	return *o.NoAutoAssignIps
}

// GetNoAutoAssignIpsOk returns a tuple with the NoAutoAssignIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetNoAutoAssignIpsOk() (*bool, bool) {
	if o == nil || o.NoAutoAssignIps == nil {
		return nil, false
	}
	return o.NoAutoAssignIps, true
}

// HasNoAutoAssignIps returns a boolean if a field has been set.
func (o *MemberConfig) HasNoAutoAssignIps() bool {
	if o != nil && o.NoAutoAssignIps != nil {
		return true
	}

	return false
}

// SetNoAutoAssignIps gets a reference to the given bool and assigns it to the NoAutoAssignIps field.
func (o *MemberConfig) SetNoAutoAssignIps(v bool) {
	o.NoAutoAssignIps = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *MemberConfig) GetRevision() int32 {
	if o == nil || o.Revision == nil {
		var ret int32
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetRevisionOk() (*int32, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *MemberConfig) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given int32 and assigns it to the Revision field.
func (o *MemberConfig) SetRevision(v int32) {
	o.Revision = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MemberConfig) GetTags() [][]int32 {
	if o == nil || o.Tags == nil {
		var ret [][]int32
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetTagsOk() (*[][]int32, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MemberConfig) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given [][]int32 and assigns it to the Tags field.
func (o *MemberConfig) SetTags(v [][]int32) {
	o.Tags = &v
}

// GetVMajor returns the VMajor field value if set, zero value otherwise.
func (o *MemberConfig) GetVMajor() int32 {
	if o == nil || o.VMajor == nil {
		var ret int32
		return ret
	}
	return *o.VMajor
}

// GetVMajorOk returns a tuple with the VMajor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetVMajorOk() (*int32, bool) {
	if o == nil || o.VMajor == nil {
		return nil, false
	}
	return o.VMajor, true
}

// HasVMajor returns a boolean if a field has been set.
func (o *MemberConfig) HasVMajor() bool {
	if o != nil && o.VMajor != nil {
		return true
	}

	return false
}

// SetVMajor gets a reference to the given int32 and assigns it to the VMajor field.
func (o *MemberConfig) SetVMajor(v int32) {
	o.VMajor = &v
}

// GetVMinor returns the VMinor field value if set, zero value otherwise.
func (o *MemberConfig) GetVMinor() int32 {
	if o == nil || o.VMinor == nil {
		var ret int32
		return ret
	}
	return *o.VMinor
}

// GetVMinorOk returns a tuple with the VMinor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetVMinorOk() (*int32, bool) {
	if o == nil || o.VMinor == nil {
		return nil, false
	}
	return o.VMinor, true
}

// HasVMinor returns a boolean if a field has been set.
func (o *MemberConfig) HasVMinor() bool {
	if o != nil && o.VMinor != nil {
		return true
	}

	return false
}

// SetVMinor gets a reference to the given int32 and assigns it to the VMinor field.
func (o *MemberConfig) SetVMinor(v int32) {
	o.VMinor = &v
}

// GetVRev returns the VRev field value if set, zero value otherwise.
func (o *MemberConfig) GetVRev() int32 {
	if o == nil || o.VRev == nil {
		var ret int32
		return ret
	}
	return *o.VRev
}

// GetVRevOk returns a tuple with the VRev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetVRevOk() (*int32, bool) {
	if o == nil || o.VRev == nil {
		return nil, false
	}
	return o.VRev, true
}

// HasVRev returns a boolean if a field has been set.
func (o *MemberConfig) HasVRev() bool {
	if o != nil && o.VRev != nil {
		return true
	}

	return false
}

// SetVRev gets a reference to the given int32 and assigns it to the VRev field.
func (o *MemberConfig) SetVRev(v int32) {
	o.VRev = &v
}

// GetVProto returns the VProto field value if set, zero value otherwise.
func (o *MemberConfig) GetVProto() int32 {
	if o == nil || o.VProto == nil {
		var ret int32
		return ret
	}
	return *o.VProto
}

// GetVProtoOk returns a tuple with the VProto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberConfig) GetVProtoOk() (*int32, bool) {
	if o == nil || o.VProto == nil {
		return nil, false
	}
	return o.VProto, true
}

// HasVProto returns a boolean if a field has been set.
func (o *MemberConfig) HasVProto() bool {
	if o != nil && o.VProto != nil {
		return true
	}

	return false
}

// SetVProto gets a reference to the given int32 and assigns it to the VProto field.
func (o *MemberConfig) SetVProto(v int32) {
	o.VProto = &v
}

func (o MemberConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveBridge != nil {
		toSerialize["activeBridge"] = o.ActiveBridge
	}
	if o.Authorized != nil {
		toSerialize["authorized"] = o.Authorized
	}
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}
	if o.CreationTime != nil {
		toSerialize["creationTime"] = o.CreationTime
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.IpAssignments != nil {
		toSerialize["ipAssignments"] = o.IpAssignments
	}
	if o.LastAuthorizedTime != nil {
		toSerialize["lastAuthorizedTime"] = o.LastAuthorizedTime
	}
	if o.LastDeauthorizedTime != nil {
		toSerialize["lastDeauthorizedTime"] = o.LastDeauthorizedTime
	}
	if o.NoAutoAssignIps != nil {
		toSerialize["noAutoAssignIps"] = o.NoAutoAssignIps
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.VMajor != nil {
		toSerialize["vMajor"] = o.VMajor
	}
	if o.VMinor != nil {
		toSerialize["vMinor"] = o.VMinor
	}
	if o.VRev != nil {
		toSerialize["vRev"] = o.VRev
	}
	if o.VProto != nil {
		toSerialize["vProto"] = o.VProto
	}
	return json.Marshal(toSerialize)
}

type NullableMemberConfig struct {
	value *MemberConfig
	isSet bool
}

func (v NullableMemberConfig) Get() *MemberConfig {
	return v.value
}

func (v *NullableMemberConfig) Set(val *MemberConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberConfig(val *MemberConfig) *NullableMemberConfig {
	return &NullableMemberConfig{value: val, isSet: true}
}

func (v NullableMemberConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


