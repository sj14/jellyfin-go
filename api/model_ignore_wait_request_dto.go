/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the IgnoreWaitRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IgnoreWaitRequestDto{}

// IgnoreWaitRequestDto Class IgnoreWaitRequestDto.
type IgnoreWaitRequestDto struct {
	// Gets or sets a value indicating whether the client should be ignored.
	IgnoreWait *bool `json:"IgnoreWait,omitempty"`
}

// NewIgnoreWaitRequestDto instantiates a new IgnoreWaitRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIgnoreWaitRequestDto() *IgnoreWaitRequestDto {
	this := IgnoreWaitRequestDto{}
	return &this
}

// NewIgnoreWaitRequestDtoWithDefaults instantiates a new IgnoreWaitRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIgnoreWaitRequestDtoWithDefaults() *IgnoreWaitRequestDto {
	this := IgnoreWaitRequestDto{}
	return &this
}

// GetIgnoreWait returns the IgnoreWait field value if set, zero value otherwise.
func (o *IgnoreWaitRequestDto) GetIgnoreWait() bool {
	if o == nil || IsNil(o.IgnoreWait) {
		var ret bool
		return ret
	}
	return *o.IgnoreWait
}

// GetIgnoreWaitOk returns a tuple with the IgnoreWait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IgnoreWaitRequestDto) GetIgnoreWaitOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreWait) {
		return nil, false
	}
	return o.IgnoreWait, true
}

// HasIgnoreWait returns a boolean if a field has been set.
func (o *IgnoreWaitRequestDto) HasIgnoreWait() bool {
	if o != nil && !IsNil(o.IgnoreWait) {
		return true
	}

	return false
}

// SetIgnoreWait gets a reference to the given bool and assigns it to the IgnoreWait field.
func (o *IgnoreWaitRequestDto) SetIgnoreWait(v bool) {
	o.IgnoreWait = &v
}

func (o IgnoreWaitRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IgnoreWaitRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IgnoreWait) {
		toSerialize["IgnoreWait"] = o.IgnoreWait
	}
	return toSerialize, nil
}

type NullableIgnoreWaitRequestDto struct {
	value *IgnoreWaitRequestDto
	isSet bool
}

func (v NullableIgnoreWaitRequestDto) Get() *IgnoreWaitRequestDto {
	return v.value
}

func (v *NullableIgnoreWaitRequestDto) Set(val *IgnoreWaitRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableIgnoreWaitRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableIgnoreWaitRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIgnoreWaitRequestDto(val *IgnoreWaitRequestDto) *NullableIgnoreWaitRequestDto {
	return &NullableIgnoreWaitRequestDto{value: val, isSet: true}
}

func (v NullableIgnoreWaitRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIgnoreWaitRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


