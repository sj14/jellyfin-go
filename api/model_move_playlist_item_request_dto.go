/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MovePlaylistItemRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MovePlaylistItemRequestDto{}

// MovePlaylistItemRequestDto Class MovePlaylistItemRequestDto.
type MovePlaylistItemRequestDto struct {
	// Gets or sets the playlist identifier of the item.
	PlaylistItemId *string `json:"PlaylistItemId,omitempty"`
	// Gets or sets the new position.
	NewIndex *int32 `json:"NewIndex,omitempty"`
}

// NewMovePlaylistItemRequestDto instantiates a new MovePlaylistItemRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovePlaylistItemRequestDto() *MovePlaylistItemRequestDto {
	this := MovePlaylistItemRequestDto{}
	return &this
}

// NewMovePlaylistItemRequestDtoWithDefaults instantiates a new MovePlaylistItemRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovePlaylistItemRequestDtoWithDefaults() *MovePlaylistItemRequestDto {
	this := MovePlaylistItemRequestDto{}
	return &this
}

// GetPlaylistItemId returns the PlaylistItemId field value if set, zero value otherwise.
func (o *MovePlaylistItemRequestDto) GetPlaylistItemId() string {
	if o == nil || IsNil(o.PlaylistItemId) {
		var ret string
		return ret
	}
	return *o.PlaylistItemId
}

// GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePlaylistItemRequestDto) GetPlaylistItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaylistItemId) {
		return nil, false
	}
	return o.PlaylistItemId, true
}

// HasPlaylistItemId returns a boolean if a field has been set.
func (o *MovePlaylistItemRequestDto) HasPlaylistItemId() bool {
	if o != nil && !IsNil(o.PlaylistItemId) {
		return true
	}

	return false
}

// SetPlaylistItemId gets a reference to the given string and assigns it to the PlaylistItemId field.
func (o *MovePlaylistItemRequestDto) SetPlaylistItemId(v string) {
	o.PlaylistItemId = &v
}

// GetNewIndex returns the NewIndex field value if set, zero value otherwise.
func (o *MovePlaylistItemRequestDto) GetNewIndex() int32 {
	if o == nil || IsNil(o.NewIndex) {
		var ret int32
		return ret
	}
	return *o.NewIndex
}

// GetNewIndexOk returns a tuple with the NewIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePlaylistItemRequestDto) GetNewIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.NewIndex) {
		return nil, false
	}
	return o.NewIndex, true
}

// HasNewIndex returns a boolean if a field has been set.
func (o *MovePlaylistItemRequestDto) HasNewIndex() bool {
	if o != nil && !IsNil(o.NewIndex) {
		return true
	}

	return false
}

// SetNewIndex gets a reference to the given int32 and assigns it to the NewIndex field.
func (o *MovePlaylistItemRequestDto) SetNewIndex(v int32) {
	o.NewIndex = &v
}

func (o MovePlaylistItemRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MovePlaylistItemRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlaylistItemId) {
		toSerialize["PlaylistItemId"] = o.PlaylistItemId
	}
	if !IsNil(o.NewIndex) {
		toSerialize["NewIndex"] = o.NewIndex
	}
	return toSerialize, nil
}

type NullableMovePlaylistItemRequestDto struct {
	value *MovePlaylistItemRequestDto
	isSet bool
}

func (v NullableMovePlaylistItemRequestDto) Get() *MovePlaylistItemRequestDto {
	return v.value
}

func (v *NullableMovePlaylistItemRequestDto) Set(val *MovePlaylistItemRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMovePlaylistItemRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMovePlaylistItemRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovePlaylistItemRequestDto(val *MovePlaylistItemRequestDto) *NullableMovePlaylistItemRequestDto {
	return &NullableMovePlaylistItemRequestDto{value: val, isSet: true}
}

func (v NullableMovePlaylistItemRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovePlaylistItemRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


