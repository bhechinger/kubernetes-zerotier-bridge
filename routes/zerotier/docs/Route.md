# Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | Pointer to **string** |  | [optional] 
**Via** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRoute

`func NewRoute() *Route`

NewRoute instantiates a new Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteWithDefaults

`func NewRouteWithDefaults() *Route`

NewRouteWithDefaults instantiates a new Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *Route) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Route) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Route) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Route) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetVia

`func (o *Route) GetVia() string`

GetVia returns the Via field if non-nil, zero value otherwise.

### GetViaOk

`func (o *Route) GetViaOk() (*string, bool)`

GetViaOk returns a tuple with the Via field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVia

`func (o *Route) SetVia(v string)`

SetVia sets Via field to given value.

### HasVia

`func (o *Route) HasVia() bool`

HasVia returns a boolean if a field has been set.

### SetViaNil

`func (o *Route) SetViaNil(b bool)`

 SetViaNil sets the value for Via to be an explicit nil

### UnsetVia
`func (o *Route) UnsetVia()`

UnsetVia ensures that no value is present for Via, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


