# Permissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**A** | Pointer to **bool** | Authorize permission | [optional] 
**D** | Pointer to **bool** | Delete permission | [optional] 
**M** | Pointer to **bool** | Modify network settings permission | [optional] 
**R** | Pointer to **bool** | Read network settings permission | [optional] 

## Methods

### NewPermissions

`func NewPermissions() *Permissions`

NewPermissions instantiates a new Permissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsWithDefaults

`func NewPermissionsWithDefaults() *Permissions`

NewPermissionsWithDefaults instantiates a new Permissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetA

`func (o *Permissions) GetA() bool`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *Permissions) GetAOk() (*bool, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *Permissions) SetA(v bool)`

SetA sets A field to given value.

### HasA

`func (o *Permissions) HasA() bool`

HasA returns a boolean if a field has been set.

### GetD

`func (o *Permissions) GetD() bool`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *Permissions) GetDOk() (*bool, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *Permissions) SetD(v bool)`

SetD sets D field to given value.

### HasD

`func (o *Permissions) HasD() bool`

HasD returns a boolean if a field has been set.

### GetM

`func (o *Permissions) GetM() bool`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *Permissions) GetMOk() (*bool, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *Permissions) SetM(v bool)`

SetM sets M field to given value.

### HasM

`func (o *Permissions) HasM() bool`

HasM returns a boolean if a field has been set.

### GetR

`func (o *Permissions) GetR() bool`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *Permissions) GetROk() (*bool, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *Permissions) SetR(v bool)`

SetR sets R field to given value.

### HasR

`func (o *Permissions) HasR() bool`

HasR returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


