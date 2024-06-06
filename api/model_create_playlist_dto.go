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

// checks if the CreatePlaylistDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePlaylistDto{}

// CreatePlaylistDto Create new playlist dto.
type CreatePlaylistDto struct {
	// Gets or sets the name of the new playlist.
	Name *string `json:"Name,omitempty"`
	// Gets or sets item ids to add to the playlist.
	Ids []string `json:"Ids,omitempty"`
	// Gets or sets the user id.
	UserId NullableString `json:"UserId,omitempty"`
	// Gets or sets the media type.
	MediaType NullableMediaType `json:"MediaType,omitempty"`
	// Gets or sets the playlist users.
	Users []PlaylistUserPermissions `json:"Users,omitempty"`
	// Gets or sets a value indicating whether the playlist is public.
	IsPublic *bool `json:"IsPublic,omitempty"`
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

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreatePlaylistDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlaylistDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreatePlaylistDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreatePlaylistDto) SetName(v string) {
	o.Name = &v
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
func (o *CreatePlaylistDto) GetMediaType() MediaType {
	if o == nil || IsNil(o.MediaType.Get()) {
		var ret MediaType
		return ret
	}
	return *o.MediaType.Get()
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePlaylistDto) GetMediaTypeOk() (*MediaType, bool) {
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

// SetMediaType gets a reference to the given NullableMediaType and assigns it to the MediaType field.
func (o *CreatePlaylistDto) SetMediaType(v MediaType) {
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

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *CreatePlaylistDto) GetUsers() []PlaylistUserPermissions {
	if o == nil || IsNil(o.Users) {
		var ret []PlaylistUserPermissions
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlaylistDto) GetUsersOk() ([]PlaylistUserPermissions, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *CreatePlaylistDto) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []PlaylistUserPermissions and assigns it to the Users field.
func (o *CreatePlaylistDto) SetUsers(v []PlaylistUserPermissions) {
	o.Users = v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *CreatePlaylistDto) GetIsPublic() bool {
	if o == nil || IsNil(o.IsPublic) {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePlaylistDto) GetIsPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPublic) {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *CreatePlaylistDto) HasIsPublic() bool {
	if o != nil && !IsNil(o.IsPublic) {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *CreatePlaylistDto) SetIsPublic(v bool) {
	o.IsPublic = &v
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
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
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
	if !IsNil(o.Users) {
		toSerialize["Users"] = o.Users
	}
	if !IsNil(o.IsPublic) {
		toSerialize["IsPublic"] = o.IsPublic
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


