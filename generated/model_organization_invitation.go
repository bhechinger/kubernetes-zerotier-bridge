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

// OrganizationInvitation struct for OrganizationInvitation
type OrganizationInvitation struct {
	// Organization ID
	OrgId *string `json:"orgId,omitempty"`
	// Email address of invitee
	Email *string `json:"email,omitempty"`
	// Invitation ID
	Id *string `json:"id,omitempty"`
	// Creation time of the invite
	CreationTime *int32 `json:"creation_time,omitempty"`
	// Invitation status
	Status *InviteStatus `json:"status,omitempty"`
	// Last updated time of the invitation
	UpdateTime *int32 `json:"update_time,omitempty"`
	// Organization owner email address
	OwnerEmail *string `json:"ownerEmail,omitempty"`
}

// NewOrganizationInvitation instantiates a new OrganizationInvitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationInvitation() *OrganizationInvitation {
	this := OrganizationInvitation{}
	return &this
}

// NewOrganizationInvitationWithDefaults instantiates a new OrganizationInvitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationInvitationWithDefaults() *OrganizationInvitation {
	this := OrganizationInvitation{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *OrganizationInvitation) SetOrgId(v string) {
	o.OrgId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *OrganizationInvitation) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationInvitation) SetId(v string) {
	o.Id = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetCreationTime() int32 {
	if o == nil || o.CreationTime == nil {
		var ret int32
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetCreationTimeOk() (*int32, bool) {
	if o == nil || o.CreationTime == nil {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasCreationTime() bool {
	if o != nil && o.CreationTime != nil {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given int32 and assigns it to the CreationTime field.
func (o *OrganizationInvitation) SetCreationTime(v int32) {
	o.CreationTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetStatus() InviteStatus {
	if o == nil || o.Status == nil {
		var ret InviteStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetStatusOk() (*InviteStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given InviteStatus and assigns it to the Status field.
func (o *OrganizationInvitation) SetStatus(v InviteStatus) {
	o.Status = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetUpdateTime() int32 {
	if o == nil || o.UpdateTime == nil {
		var ret int32
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetUpdateTimeOk() (*int32, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int32 and assigns it to the UpdateTime field.
func (o *OrganizationInvitation) SetUpdateTime(v int32) {
	o.UpdateTime = &v
}

// GetOwnerEmail returns the OwnerEmail field value if set, zero value otherwise.
func (o *OrganizationInvitation) GetOwnerEmail() string {
	if o == nil || o.OwnerEmail == nil {
		var ret string
		return ret
	}
	return *o.OwnerEmail
}

// GetOwnerEmailOk returns a tuple with the OwnerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationInvitation) GetOwnerEmailOk() (*string, bool) {
	if o == nil || o.OwnerEmail == nil {
		return nil, false
	}
	return o.OwnerEmail, true
}

// HasOwnerEmail returns a boolean if a field has been set.
func (o *OrganizationInvitation) HasOwnerEmail() bool {
	if o != nil && o.OwnerEmail != nil {
		return true
	}

	return false
}

// SetOwnerEmail gets a reference to the given string and assigns it to the OwnerEmail field.
func (o *OrganizationInvitation) SetOwnerEmail(v string) {
	o.OwnerEmail = &v
}

func (o OrganizationInvitation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OrgId != nil {
		toSerialize["orgId"] = o.OrgId
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreationTime != nil {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdateTime != nil {
		toSerialize["update_time"] = o.UpdateTime
	}
	if o.OwnerEmail != nil {
		toSerialize["ownerEmail"] = o.OwnerEmail
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationInvitation struct {
	value *OrganizationInvitation
	isSet bool
}

func (v NullableOrganizationInvitation) Get() *OrganizationInvitation {
	return v.value
}

func (v *NullableOrganizationInvitation) Set(val *OrganizationInvitation) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationInvitation(val *OrganizationInvitation) *NullableOrganizationInvitation {
	return &NullableOrganizationInvitation{value: val, isSet: true}
}

func (v NullableOrganizationInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


