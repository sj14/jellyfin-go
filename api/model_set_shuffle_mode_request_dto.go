/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SetShuffleModeRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetShuffleModeRequestDto{}

// SetShuffleModeRequestDto Class SetShuffleModeRequestDto.
type SetShuffleModeRequestDto struct {
	// Enum GroupShuffleMode.
	Mode *GroupShuffleMode `json:"Mode,omitempty"`
}

// NewSetShuffleModeRequestDto instantiates a new SetShuffleModeRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetShuffleModeRequestDto() *SetShuffleModeRequestDto {
	this := SetShuffleModeRequestDto{}
	return &this
}

// NewSetShuffleModeRequestDtoWithDefaults instantiates a new SetShuffleModeRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetShuffleModeRequestDtoWithDefaults() *SetShuffleModeRequestDto {
	this := SetShuffleModeRequestDto{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *SetShuffleModeRequestDto) GetMode() GroupShuffleMode {
	if o == nil || IsNil(o.Mode) {
		var ret GroupShuffleMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetShuffleModeRequestDto) GetModeOk() (*GroupShuffleMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *SetShuffleModeRequestDto) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given GroupShuffleMode and assigns it to the Mode field.
func (o *SetShuffleModeRequestDto) SetMode(v GroupShuffleMode) {
	o.Mode = &v
}

func (o SetShuffleModeRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetShuffleModeRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["Mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableSetShuffleModeRequestDto struct {
	value *SetShuffleModeRequestDto
	isSet bool
}

func (v NullableSetShuffleModeRequestDto) Get() *SetShuffleModeRequestDto {
	return v.value
}

func (v *NullableSetShuffleModeRequestDto) Set(val *SetShuffleModeRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSetShuffleModeRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSetShuffleModeRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetShuffleModeRequestDto(val *SetShuffleModeRequestDto) *NullableSetShuffleModeRequestDto {
	return &NullableSetShuffleModeRequestDto{value: val, isSet: true}
}

func (v NullableSetShuffleModeRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetShuffleModeRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


