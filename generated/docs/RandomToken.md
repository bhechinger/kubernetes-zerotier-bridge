# RandomToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clock** | Pointer to **int32** | Current time on server | [optional] [readonly] 
**Hex** | Pointer to **string** | hex encoded random bytes of the token | [optional] [readonly] 
**Token** | Pointer to **string** | Random 32 character token | [optional] [readonly] 

## Methods

### NewRandomToken

`func NewRandomToken() *RandomToken`

NewRandomToken instantiates a new RandomToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRandomTokenWithDefaults

`func NewRandomTokenWithDefaults() *RandomToken`

NewRandomTokenWithDefaults instantiates a new RandomToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClock

`func (o *RandomToken) GetClock() int32`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *RandomToken) GetClockOk() (*int32, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *RandomToken) SetClock(v int32)`

SetClock sets Clock field to given value.

### HasClock

`func (o *RandomToken) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetHex

`func (o *RandomToken) GetHex() string`

GetHex returns the Hex field if non-nil, zero value otherwise.

### GetHexOk

`func (o *RandomToken) GetHexOk() (*string, bool)`

GetHexOk returns a tuple with the Hex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHex

`func (o *RandomToken) SetHex(v string)`

SetHex sets Hex field to given value.

### HasHex

`func (o *RandomToken) HasHex() bool`

HasHex returns a boolean if a field has been set.

### GetToken

`func (o *RandomToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RandomToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RandomToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RandomToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


