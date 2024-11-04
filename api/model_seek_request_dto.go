/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SeekRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeekRequestDto{}

// SeekRequestDto Class SeekRequestDto.
type SeekRequestDto struct {
	// Gets or sets the position ticks.
	PositionTicks *int64 `json:"PositionTicks,omitempty"`
}

// NewSeekRequestDto instantiates a new SeekRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeekRequestDto() *SeekRequestDto {
	this := SeekRequestDto{}
	return &this
}

// NewSeekRequestDtoWithDefaults instantiates a new SeekRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeekRequestDtoWithDefaults() *SeekRequestDto {
	this := SeekRequestDto{}
	return &this
}

// GetPositionTicks returns the PositionTicks field value if set, zero value otherwise.
func (o *SeekRequestDto) GetPositionTicks() int64 {
	if o == nil || IsNil(o.PositionTicks) {
		var ret int64
		return ret
	}
	return *o.PositionTicks
}

// GetPositionTicksOk returns a tuple with the PositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeekRequestDto) GetPositionTicksOk() (*int64, bool) {
	if o == nil || IsNil(o.PositionTicks) {
		return nil, false
	}
	return o.PositionTicks, true
}

// HasPositionTicks returns a boolean if a field has been set.
func (o *SeekRequestDto) HasPositionTicks() bool {
	if o != nil && !IsNil(o.PositionTicks) {
		return true
	}

	return false
}

// SetPositionTicks gets a reference to the given int64 and assigns it to the PositionTicks field.
func (o *SeekRequestDto) SetPositionTicks(v int64) {
	o.PositionTicks = &v
}

func (o SeekRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeekRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PositionTicks) {
		toSerialize["PositionTicks"] = o.PositionTicks
	}
	return toSerialize, nil
}

type NullableSeekRequestDto struct {
	value *SeekRequestDto
	isSet bool
}

func (v NullableSeekRequestDto) Get() *SeekRequestDto {
	return v.value
}

func (v *NullableSeekRequestDto) Set(val *SeekRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSeekRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSeekRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeekRequestDto(val *SeekRequestDto) *NullableSeekRequestDto {
	return &NullableSeekRequestDto{value: val, isSet: true}
}

func (v NullableSeekRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeekRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


