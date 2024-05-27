/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MediaUpdateInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaUpdateInfoDto{}

// MediaUpdateInfoDto Media Update Info Dto.
type MediaUpdateInfoDto struct {
	// Gets or sets the list of updates.
	Updates []MediaUpdateInfoPathDto `json:"Updates,omitempty"`
}

// NewMediaUpdateInfoDto instantiates a new MediaUpdateInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaUpdateInfoDto() *MediaUpdateInfoDto {
	this := MediaUpdateInfoDto{}
	return &this
}

// NewMediaUpdateInfoDtoWithDefaults instantiates a new MediaUpdateInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaUpdateInfoDtoWithDefaults() *MediaUpdateInfoDto {
	this := MediaUpdateInfoDto{}
	return &this
}

// GetUpdates returns the Updates field value if set, zero value otherwise.
func (o *MediaUpdateInfoDto) GetUpdates() []MediaUpdateInfoPathDto {
	if o == nil || IsNil(o.Updates) {
		var ret []MediaUpdateInfoPathDto
		return ret
	}
	return o.Updates
}

// GetUpdatesOk returns a tuple with the Updates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaUpdateInfoDto) GetUpdatesOk() ([]MediaUpdateInfoPathDto, bool) {
	if o == nil || IsNil(o.Updates) {
		return nil, false
	}
	return o.Updates, true
}

// HasUpdates returns a boolean if a field has been set.
func (o *MediaUpdateInfoDto) HasUpdates() bool {
	if o != nil && !IsNil(o.Updates) {
		return true
	}

	return false
}

// SetUpdates gets a reference to the given []MediaUpdateInfoPathDto and assigns it to the Updates field.
func (o *MediaUpdateInfoDto) SetUpdates(v []MediaUpdateInfoPathDto) {
	o.Updates = v
}

func (o MediaUpdateInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaUpdateInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Updates) {
		toSerialize["Updates"] = o.Updates
	}
	return toSerialize, nil
}

type NullableMediaUpdateInfoDto struct {
	value *MediaUpdateInfoDto
	isSet bool
}

func (v NullableMediaUpdateInfoDto) Get() *MediaUpdateInfoDto {
	return v.value
}

func (v *NullableMediaUpdateInfoDto) Set(val *MediaUpdateInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaUpdateInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaUpdateInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaUpdateInfoDto(val *MediaUpdateInfoDto) *NullableMediaUpdateInfoDto {
	return &NullableMediaUpdateInfoDto{value: val, isSet: true}
}

func (v NullableMediaUpdateInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaUpdateInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


