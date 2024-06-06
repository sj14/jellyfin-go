/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PreviousItemRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreviousItemRequestDto{}

// PreviousItemRequestDto Class PreviousItemRequestDto.
type PreviousItemRequestDto struct {
	// Gets or sets the playing item identifier.
	PlaylistItemId *string `json:"PlaylistItemId,omitempty"`
}

// NewPreviousItemRequestDto instantiates a new PreviousItemRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviousItemRequestDto() *PreviousItemRequestDto {
	this := PreviousItemRequestDto{}
	return &this
}

// NewPreviousItemRequestDtoWithDefaults instantiates a new PreviousItemRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviousItemRequestDtoWithDefaults() *PreviousItemRequestDto {
	this := PreviousItemRequestDto{}
	return &this
}

// GetPlaylistItemId returns the PlaylistItemId field value if set, zero value otherwise.
func (o *PreviousItemRequestDto) GetPlaylistItemId() string {
	if o == nil || IsNil(o.PlaylistItemId) {
		var ret string
		return ret
	}
	return *o.PlaylistItemId
}

// GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviousItemRequestDto) GetPlaylistItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaylistItemId) {
		return nil, false
	}
	return o.PlaylistItemId, true
}

// HasPlaylistItemId returns a boolean if a field has been set.
func (o *PreviousItemRequestDto) HasPlaylistItemId() bool {
	if o != nil && !IsNil(o.PlaylistItemId) {
		return true
	}

	return false
}

// SetPlaylistItemId gets a reference to the given string and assigns it to the PlaylistItemId field.
func (o *PreviousItemRequestDto) SetPlaylistItemId(v string) {
	o.PlaylistItemId = &v
}

func (o PreviousItemRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreviousItemRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlaylistItemId) {
		toSerialize["PlaylistItemId"] = o.PlaylistItemId
	}
	return toSerialize, nil
}

type NullablePreviousItemRequestDto struct {
	value *PreviousItemRequestDto
	isSet bool
}

func (v NullablePreviousItemRequestDto) Get() *PreviousItemRequestDto {
	return v.value
}

func (v *NullablePreviousItemRequestDto) Set(val *PreviousItemRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviousItemRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviousItemRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviousItemRequestDto(val *PreviousItemRequestDto) *NullablePreviousItemRequestDto {
	return &NullablePreviousItemRequestDto{value: val, isSet: true}
}

func (v NullablePreviousItemRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviousItemRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


