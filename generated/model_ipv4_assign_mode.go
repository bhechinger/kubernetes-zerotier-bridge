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

// IPV4AssignMode struct for IPV4AssignMode
type IPV4AssignMode struct {
	Zt *bool `json:"zt,omitempty"`
}

// NewIPV4AssignMode instantiates a new IPV4AssignMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPV4AssignMode() *IPV4AssignMode {
	this := IPV4AssignMode{}
	return &this
}

// NewIPV4AssignModeWithDefaults instantiates a new IPV4AssignMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPV4AssignModeWithDefaults() *IPV4AssignMode {
	this := IPV4AssignMode{}
	return &this
}

// GetZt returns the Zt field value if set, zero value otherwise.
func (o *IPV4AssignMode) GetZt() bool {
	if o == nil || o.Zt == nil {
		var ret bool
		return ret
	}
	return *o.Zt
}

// GetZtOk returns a tuple with the Zt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPV4AssignMode) GetZtOk() (*bool, bool) {
	if o == nil || o.Zt == nil {
		return nil, false
	}
	return o.Zt, true
}

// HasZt returns a boolean if a field has been set.
func (o *IPV4AssignMode) HasZt() bool {
	if o != nil && o.Zt != nil {
		return true
	}

	return false
}

// SetZt gets a reference to the given bool and assigns it to the Zt field.
func (o *IPV4AssignMode) SetZt(v bool) {
	o.Zt = &v
}

func (o IPV4AssignMode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Zt != nil {
		toSerialize["zt"] = o.Zt
	}
	return json.Marshal(toSerialize)
}

type NullableIPV4AssignMode struct {
	value *IPV4AssignMode
	isSet bool
}

func (v NullableIPV4AssignMode) Get() *IPV4AssignMode {
	return v.value
}

func (v *NullableIPV4AssignMode) Set(val *IPV4AssignMode) {
	v.value = val
	v.isSet = true
}

func (v NullableIPV4AssignMode) IsSet() bool {
	return v.isSet
}

func (v *NullableIPV4AssignMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPV4AssignMode(val *IPV4AssignMode) *NullableIPV4AssignMode {
	return &NullableIPV4AssignMode{value: val, isSet: true}
}

func (v NullableIPV4AssignMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPV4AssignMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

