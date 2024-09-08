/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdatePlaylistUserDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePlaylistUserDto{}

// UpdatePlaylistUserDto Update existing playlist user dto. Fields set to `null` will not be updated and keep their current values.
type UpdatePlaylistUserDto struct {
	// Gets or sets a value indicating whether the user can edit the playlist.
	CanEdit NullableBool `json:"CanEdit,omitempty"`
}

// NewUpdatePlaylistUserDto instantiates a new UpdatePlaylistUserDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePlaylistUserDto() *UpdatePlaylistUserDto {
	this := UpdatePlaylistUserDto{}
	return &this
}

// NewUpdatePlaylistUserDtoWithDefaults instantiates a new UpdatePlaylistUserDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePlaylistUserDtoWithDefaults() *UpdatePlaylistUserDto {
	this := UpdatePlaylistUserDto{}
	return &this
}

// GetCanEdit returns the CanEdit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdatePlaylistUserDto) GetCanEdit() bool {
	if o == nil || IsNil(o.CanEdit.Get()) {
		var ret bool
		return ret
	}
	return *o.CanEdit.Get()
}

// GetCanEditOk returns a tuple with the CanEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdatePlaylistUserDto) GetCanEditOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CanEdit.Get(), o.CanEdit.IsSet()
}

// HasCanEdit returns a boolean if a field has been set.
func (o *UpdatePlaylistUserDto) HasCanEdit() bool {
	if o != nil && o.CanEdit.IsSet() {
		return true
	}

	return false
}

// SetCanEdit gets a reference to the given NullableBool and assigns it to the CanEdit field.
func (o *UpdatePlaylistUserDto) SetCanEdit(v bool) {
	o.CanEdit.Set(&v)
}
// SetCanEditNil sets the value for CanEdit to be an explicit nil
func (o *UpdatePlaylistUserDto) SetCanEditNil() {
	o.CanEdit.Set(nil)
}

// UnsetCanEdit ensures that no value is present for CanEdit, not even an explicit nil
func (o *UpdatePlaylistUserDto) UnsetCanEdit() {
	o.CanEdit.Unset()
}

func (o UpdatePlaylistUserDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePlaylistUserDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CanEdit.IsSet() {
		toSerialize["CanEdit"] = o.CanEdit.Get()
	}
	return toSerialize, nil
}

type NullableUpdatePlaylistUserDto struct {
	value *UpdatePlaylistUserDto
	isSet bool
}

func (v NullableUpdatePlaylistUserDto) Get() *UpdatePlaylistUserDto {
	return v.value
}

func (v *NullableUpdatePlaylistUserDto) Set(val *UpdatePlaylistUserDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePlaylistUserDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePlaylistUserDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePlaylistUserDto(val *UpdatePlaylistUserDto) *NullableUpdatePlaylistUserDto {
	return &NullableUpdatePlaylistUserDto{value: val, isSet: true}
}

func (v NullableUpdatePlaylistUserDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePlaylistUserDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


