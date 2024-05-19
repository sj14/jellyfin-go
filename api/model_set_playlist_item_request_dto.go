/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SetPlaylistItemRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetPlaylistItemRequestDto{}

// SetPlaylistItemRequestDto Class SetPlaylistItemRequestDto.
type SetPlaylistItemRequestDto struct {
	// Gets or sets the playlist identifier of the playing item.
	PlaylistItemId *string `json:"PlaylistItemId,omitempty"`
}

// NewSetPlaylistItemRequestDto instantiates a new SetPlaylistItemRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetPlaylistItemRequestDto() *SetPlaylistItemRequestDto {
	this := SetPlaylistItemRequestDto{}
	return &this
}

// NewSetPlaylistItemRequestDtoWithDefaults instantiates a new SetPlaylistItemRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetPlaylistItemRequestDtoWithDefaults() *SetPlaylistItemRequestDto {
	this := SetPlaylistItemRequestDto{}
	return &this
}

// GetPlaylistItemId returns the PlaylistItemId field value if set, zero value otherwise.
func (o *SetPlaylistItemRequestDto) GetPlaylistItemId() string {
	if o == nil || IsNil(o.PlaylistItemId) {
		var ret string
		return ret
	}
	return *o.PlaylistItemId
}

// GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetPlaylistItemRequestDto) GetPlaylistItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaylistItemId) {
		return nil, false
	}
	return o.PlaylistItemId, true
}

// HasPlaylistItemId returns a boolean if a field has been set.
func (o *SetPlaylistItemRequestDto) HasPlaylistItemId() bool {
	if o != nil && !IsNil(o.PlaylistItemId) {
		return true
	}

	return false
}

// SetPlaylistItemId gets a reference to the given string and assigns it to the PlaylistItemId field.
func (o *SetPlaylistItemRequestDto) SetPlaylistItemId(v string) {
	o.PlaylistItemId = &v
}

func (o SetPlaylistItemRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetPlaylistItemRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlaylistItemId) {
		toSerialize["PlaylistItemId"] = o.PlaylistItemId
	}
	return toSerialize, nil
}

type NullableSetPlaylistItemRequestDto struct {
	value *SetPlaylistItemRequestDto
	isSet bool
}

func (v NullableSetPlaylistItemRequestDto) Get() *SetPlaylistItemRequestDto {
	return v.value
}

func (v *NullableSetPlaylistItemRequestDto) Set(val *SetPlaylistItemRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSetPlaylistItemRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSetPlaylistItemRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetPlaylistItemRequestDto(val *SetPlaylistItemRequestDto) *NullableSetPlaylistItemRequestDto {
	return &NullableSetPlaylistItemRequestDto{value: val, isSet: true}
}

func (v NullableSetPlaylistItemRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetPlaylistItemRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


