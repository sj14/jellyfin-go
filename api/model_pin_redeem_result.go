/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.8.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PinRedeemResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PinRedeemResult{}

// PinRedeemResult struct for PinRedeemResult
type PinRedeemResult struct {
	// Gets or sets a value indicating whether this MediaBrowser.Model.Users.PinRedeemResult is success.
	Success *bool `json:"Success,omitempty"`
	// Gets or sets the users reset.
	UsersReset []string `json:"UsersReset,omitempty"`
}

// NewPinRedeemResult instantiates a new PinRedeemResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPinRedeemResult() *PinRedeemResult {
	this := PinRedeemResult{}
	return &this
}

// NewPinRedeemResultWithDefaults instantiates a new PinRedeemResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPinRedeemResultWithDefaults() *PinRedeemResult {
	this := PinRedeemResult{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *PinRedeemResult) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinRedeemResult) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *PinRedeemResult) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *PinRedeemResult) SetSuccess(v bool) {
	o.Success = &v
}

// GetUsersReset returns the UsersReset field value if set, zero value otherwise.
func (o *PinRedeemResult) GetUsersReset() []string {
	if o == nil || IsNil(o.UsersReset) {
		var ret []string
		return ret
	}
	return o.UsersReset
}

// GetUsersResetOk returns a tuple with the UsersReset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PinRedeemResult) GetUsersResetOk() ([]string, bool) {
	if o == nil || IsNil(o.UsersReset) {
		return nil, false
	}
	return o.UsersReset, true
}

// HasUsersReset returns a boolean if a field has been set.
func (o *PinRedeemResult) HasUsersReset() bool {
	if o != nil && !IsNil(o.UsersReset) {
		return true
	}

	return false
}

// SetUsersReset gets a reference to the given []string and assigns it to the UsersReset field.
func (o *PinRedeemResult) SetUsersReset(v []string) {
	o.UsersReset = v
}

func (o PinRedeemResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PinRedeemResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["Success"] = o.Success
	}
	if !IsNil(o.UsersReset) {
		toSerialize["UsersReset"] = o.UsersReset
	}
	return toSerialize, nil
}

type NullablePinRedeemResult struct {
	value *PinRedeemResult
	isSet bool
}

func (v NullablePinRedeemResult) Get() *PinRedeemResult {
	return v.value
}

func (v *NullablePinRedeemResult) Set(val *PinRedeemResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePinRedeemResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePinRedeemResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePinRedeemResult(val *PinRedeemResult) *NullablePinRedeemResult {
	return &NullablePinRedeemResult{value: val, isSet: true}
}

func (v NullablePinRedeemResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePinRedeemResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


