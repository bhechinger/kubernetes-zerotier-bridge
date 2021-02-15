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

// User struct for User
type User struct {
	// User ID
	Id *string `json:"id,omitempty"`
	// Organization ID
	OrgId *string `json:"orgId,omitempty"`
	GlobalPermissions *Permissions `json:"globalPermissions,omitempty"`
	// Display Name
	DisplayName *string `json:"displayName,omitempty"`
	// User email address
	Email *string `json:"email,omitempty"`
	Auth *AuthMethods `json:"auth,omitempty"`
	// SMS number
	SmsNumber *string `json:"smsNumber,omitempty"`
	// List of API token names.
	Tokens *[]string `json:"tokens,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *User) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *User) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *User) SetOrgId(v string) {
	o.OrgId = &v
}

// GetGlobalPermissions returns the GlobalPermissions field value if set, zero value otherwise.
func (o *User) GetGlobalPermissions() Permissions {
	if o == nil || o.GlobalPermissions == nil {
		var ret Permissions
		return ret
	}
	return *o.GlobalPermissions
}

// GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetGlobalPermissionsOk() (*Permissions, bool) {
	if o == nil || o.GlobalPermissions == nil {
		return nil, false
	}
	return o.GlobalPermissions, true
}

// HasGlobalPermissions returns a boolean if a field has been set.
func (o *User) HasGlobalPermissions() bool {
	if o != nil && o.GlobalPermissions != nil {
		return true
	}

	return false
}

// SetGlobalPermissions gets a reference to the given Permissions and assigns it to the GlobalPermissions field.
func (o *User) SetGlobalPermissions(v Permissions) {
	o.GlobalPermissions = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *User) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *User) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *User) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email = &v
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *User) GetAuth() AuthMethods {
	if o == nil || o.Auth == nil {
		var ret AuthMethods
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAuthOk() (*AuthMethods, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *User) HasAuth() bool {
	if o != nil && o.Auth != nil {
		return true
	}

	return false
}

// SetAuth gets a reference to the given AuthMethods and assigns it to the Auth field.
func (o *User) SetAuth(v AuthMethods) {
	o.Auth = &v
}

// GetSmsNumber returns the SmsNumber field value if set, zero value otherwise.
func (o *User) GetSmsNumber() string {
	if o == nil || o.SmsNumber == nil {
		var ret string
		return ret
	}
	return *o.SmsNumber
}

// GetSmsNumberOk returns a tuple with the SmsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetSmsNumberOk() (*string, bool) {
	if o == nil || o.SmsNumber == nil {
		return nil, false
	}
	return o.SmsNumber, true
}

// HasSmsNumber returns a boolean if a field has been set.
func (o *User) HasSmsNumber() bool {
	if o != nil && o.SmsNumber != nil {
		return true
	}

	return false
}

// SetSmsNumber gets a reference to the given string and assigns it to the SmsNumber field.
func (o *User) SetSmsNumber(v string) {
	o.SmsNumber = &v
}

// GetTokens returns the Tokens field value if set, zero value otherwise.
func (o *User) GetTokens() []string {
	if o == nil || o.Tokens == nil {
		var ret []string
		return ret
	}
	return *o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTokensOk() (*[]string, bool) {
	if o == nil || o.Tokens == nil {
		return nil, false
	}
	return o.Tokens, true
}

// HasTokens returns a boolean if a field has been set.
func (o *User) HasTokens() bool {
	if o != nil && o.Tokens != nil {
		return true
	}

	return false
}

// SetTokens gets a reference to the given []string and assigns it to the Tokens field.
func (o *User) SetTokens(v []string) {
	o.Tokens = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrgId != nil {
		toSerialize["orgId"] = o.OrgId
	}
	if o.GlobalPermissions != nil {
		toSerialize["globalPermissions"] = o.GlobalPermissions
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	if o.SmsNumber != nil {
		toSerialize["smsNumber"] = o.SmsNumber
	}
	if o.Tokens != nil {
		toSerialize["tokens"] = o.Tokens
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


