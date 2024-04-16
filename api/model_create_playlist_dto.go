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

// checks if the CreatePlaylistDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePlaylistDto{}

// CreatePlaylistDto Create new playlist dto.
type CreatePlaylistDto struct {
	// Gets or sets the name of the new playlist.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets item ids to add to the playlist.
	Ids []string `json:"Ids,omitempty"`
	// Gets or sets the user id.
	UserId NullableString `json:"UserId,omitempty"`
	// Gets or sets the media type.
	MediaType NullableString `json:"MediaType,omitempty"`
}

// NewCreatePlaylistDto instantiates a new CreatePlaylistDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePlaylistDto() *CreatePlaylistDto {
	this := CreatePlaylistDto{}
	return &this
}

// NewCreatePlaylistDtoWithDefaults instantiates a new CreatePlaylistDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePlaylistDtoWithDefaults() *CreatePlaylistDto {
	this := CreatePlaylistDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePlaylistDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePlaylistDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CreatePlaylistDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CreatePlaylistDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CreatePlaylistDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CreatePlaylistDto) UnsetName() {
	o.Name.Unset()
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *CreatePlaylistDto) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlaylistDto) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *CreatePlaylistDto) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *CreatePlaylistDto) SetIds(v []string) {
	o.Ids = v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePlaylistDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePlaylistDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *CreatePlaylistDto) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *CreatePlaylistDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *CreatePlaylistDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *CreatePlaylistDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetMediaType returns the MediaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePlaylistDto) GetMediaType() string {
	if o == nil || IsNil(o.MediaType.Get()) {
		var ret string
		return ret
	}
	return *o.MediaType.Get()
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePlaylistDto) GetMediaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaType.Get(), o.MediaType.IsSet()
}

// HasMediaType returns a boolean if a field has been set.
func (o *CreatePlaylistDto) HasMediaType() bool {
	if o != nil && o.MediaType.IsSet() {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given NullableString and assigns it to the MediaType field.
func (o *CreatePlaylistDto) SetMediaType(v string) {
	o.MediaType.Set(&v)
}
// SetMediaTypeNil sets the value for MediaType to be an explicit nil
func (o *CreatePlaylistDto) SetMediaTypeNil() {
	o.MediaType.Set(nil)
}

// UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
func (o *CreatePlaylistDto) UnsetMediaType() {
	o.MediaType.Unset()
}

func (o CreatePlaylistDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePlaylistDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.Ids) {
		toSerialize["Ids"] = o.Ids
	}
	if o.UserId.IsSet() {
		toSerialize["UserId"] = o.UserId.Get()
	}
	if o.MediaType.IsSet() {
		toSerialize["MediaType"] = o.MediaType.Get()
	}
	return toSerialize, nil
}

type NullableCreatePlaylistDto struct {
	value *CreatePlaylistDto
	isSet bool
}

func (v NullableCreatePlaylistDto) Get() *CreatePlaylistDto {
	return v.value
}

func (v *NullableCreatePlaylistDto) Set(val *CreatePlaylistDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePlaylistDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePlaylistDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePlaylistDto(val *CreatePlaylistDto) *NullableCreatePlaylistDto {
	return &NullableCreatePlaylistDto{value: val, isSet: true}
}

func (v NullableCreatePlaylistDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePlaylistDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


