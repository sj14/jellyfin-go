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

// checks if the PlaylistUserPermissions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaylistUserPermissions{}

// PlaylistUserPermissions Class to hold data on user permissions for playlists.
type PlaylistUserPermissions struct {
	// Gets or sets the user id.
	UserId *string `json:"UserId,omitempty"`
	// Gets or sets a value indicating whether the user has edit permissions.
	CanEdit *bool `json:"CanEdit,omitempty"`
}

// NewPlaylistUserPermissions instantiates a new PlaylistUserPermissions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaylistUserPermissions() *PlaylistUserPermissions {
	this := PlaylistUserPermissions{}
	return &this
}

// NewPlaylistUserPermissionsWithDefaults instantiates a new PlaylistUserPermissions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaylistUserPermissionsWithDefaults() *PlaylistUserPermissions {
	this := PlaylistUserPermissions{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PlaylistUserPermissions) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistUserPermissions) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PlaylistUserPermissions) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PlaylistUserPermissions) SetUserId(v string) {
	o.UserId = &v
}

// GetCanEdit returns the CanEdit field value if set, zero value otherwise.
func (o *PlaylistUserPermissions) GetCanEdit() bool {
	if o == nil || IsNil(o.CanEdit) {
		var ret bool
		return ret
	}
	return *o.CanEdit
}

// GetCanEditOk returns a tuple with the CanEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaylistUserPermissions) GetCanEditOk() (*bool, bool) {
	if o == nil || IsNil(o.CanEdit) {
		return nil, false
	}
	return o.CanEdit, true
}

// HasCanEdit returns a boolean if a field has been set.
func (o *PlaylistUserPermissions) HasCanEdit() bool {
	if o != nil && !IsNil(o.CanEdit) {
		return true
	}

	return false
}

// SetCanEdit gets a reference to the given bool and assigns it to the CanEdit field.
func (o *PlaylistUserPermissions) SetCanEdit(v bool) {
	o.CanEdit = &v
}

func (o PlaylistUserPermissions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaylistUserPermissions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["UserId"] = o.UserId
	}
	if !IsNil(o.CanEdit) {
		toSerialize["CanEdit"] = o.CanEdit
	}
	return toSerialize, nil
}

type NullablePlaylistUserPermissions struct {
	value *PlaylistUserPermissions
	isSet bool
}

func (v NullablePlaylistUserPermissions) Get() *PlaylistUserPermissions {
	return v.value
}

func (v *NullablePlaylistUserPermissions) Set(val *PlaylistUserPermissions) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaylistUserPermissions) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaylistUserPermissions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaylistUserPermissions(val *PlaylistUserPermissions) *NullablePlaylistUserPermissions {
	return &NullablePlaylistUserPermissions{value: val, isSet: true}
}

func (v NullablePlaylistUserPermissions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaylistUserPermissions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


