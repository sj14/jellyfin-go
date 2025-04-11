/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the RemoveFromPlaylistRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveFromPlaylistRequestDto{}

// RemoveFromPlaylistRequestDto Class RemoveFromPlaylistRequestDto.
type RemoveFromPlaylistRequestDto struct {
	// Gets or sets the playlist identifiers of the items. Ignored when clearing the playlist.
	PlaylistItemIds []string `json:"PlaylistItemIds,omitempty"`
	// Gets or sets a value indicating whether the entire playlist should be cleared.
	ClearPlaylist *bool `json:"ClearPlaylist,omitempty"`
	// Gets or sets a value indicating whether the playing item should be removed as well. Used only when clearing the playlist.
	ClearPlayingItem *bool `json:"ClearPlayingItem,omitempty"`
}

// NewRemoveFromPlaylistRequestDto instantiates a new RemoveFromPlaylistRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveFromPlaylistRequestDto() *RemoveFromPlaylistRequestDto {
	this := RemoveFromPlaylistRequestDto{}
	return &this
}

// NewRemoveFromPlaylistRequestDtoWithDefaults instantiates a new RemoveFromPlaylistRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveFromPlaylistRequestDtoWithDefaults() *RemoveFromPlaylistRequestDto {
	this := RemoveFromPlaylistRequestDto{}
	return &this
}

// GetPlaylistItemIds returns the PlaylistItemIds field value if set, zero value otherwise.
func (o *RemoveFromPlaylistRequestDto) GetPlaylistItemIds() []string {
	if o == nil || IsNil(o.PlaylistItemIds) {
		var ret []string
		return ret
	}
	return o.PlaylistItemIds
}

// GetPlaylistItemIdsOk returns a tuple with the PlaylistItemIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveFromPlaylistRequestDto) GetPlaylistItemIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.PlaylistItemIds) {
		return nil, false
	}
	return o.PlaylistItemIds, true
}

// HasPlaylistItemIds returns a boolean if a field has been set.
func (o *RemoveFromPlaylistRequestDto) HasPlaylistItemIds() bool {
	if o != nil && !IsNil(o.PlaylistItemIds) {
		return true
	}

	return false
}

// SetPlaylistItemIds gets a reference to the given []string and assigns it to the PlaylistItemIds field.
func (o *RemoveFromPlaylistRequestDto) SetPlaylistItemIds(v []string) {
	o.PlaylistItemIds = v
}

// GetClearPlaylist returns the ClearPlaylist field value if set, zero value otherwise.
func (o *RemoveFromPlaylistRequestDto) GetClearPlaylist() bool {
	if o == nil || IsNil(o.ClearPlaylist) {
		var ret bool
		return ret
	}
	return *o.ClearPlaylist
}

// GetClearPlaylistOk returns a tuple with the ClearPlaylist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveFromPlaylistRequestDto) GetClearPlaylistOk() (*bool, bool) {
	if o == nil || IsNil(o.ClearPlaylist) {
		return nil, false
	}
	return o.ClearPlaylist, true
}

// HasClearPlaylist returns a boolean if a field has been set.
func (o *RemoveFromPlaylistRequestDto) HasClearPlaylist() bool {
	if o != nil && !IsNil(o.ClearPlaylist) {
		return true
	}

	return false
}

// SetClearPlaylist gets a reference to the given bool and assigns it to the ClearPlaylist field.
func (o *RemoveFromPlaylistRequestDto) SetClearPlaylist(v bool) {
	o.ClearPlaylist = &v
}

// GetClearPlayingItem returns the ClearPlayingItem field value if set, zero value otherwise.
func (o *RemoveFromPlaylistRequestDto) GetClearPlayingItem() bool {
	if o == nil || IsNil(o.ClearPlayingItem) {
		var ret bool
		return ret
	}
	return *o.ClearPlayingItem
}

// GetClearPlayingItemOk returns a tuple with the ClearPlayingItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveFromPlaylistRequestDto) GetClearPlayingItemOk() (*bool, bool) {
	if o == nil || IsNil(o.ClearPlayingItem) {
		return nil, false
	}
	return o.ClearPlayingItem, true
}

// HasClearPlayingItem returns a boolean if a field has been set.
func (o *RemoveFromPlaylistRequestDto) HasClearPlayingItem() bool {
	if o != nil && !IsNil(o.ClearPlayingItem) {
		return true
	}

	return false
}

// SetClearPlayingItem gets a reference to the given bool and assigns it to the ClearPlayingItem field.
func (o *RemoveFromPlaylistRequestDto) SetClearPlayingItem(v bool) {
	o.ClearPlayingItem = &v
}

func (o RemoveFromPlaylistRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveFromPlaylistRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlaylistItemIds) {
		toSerialize["PlaylistItemIds"] = o.PlaylistItemIds
	}
	if !IsNil(o.ClearPlaylist) {
		toSerialize["ClearPlaylist"] = o.ClearPlaylist
	}
	if !IsNil(o.ClearPlayingItem) {
		toSerialize["ClearPlayingItem"] = o.ClearPlayingItem
	}
	return toSerialize, nil
}

type NullableRemoveFromPlaylistRequestDto struct {
	value *RemoveFromPlaylistRequestDto
	isSet bool
}

func (v NullableRemoveFromPlaylistRequestDto) Get() *RemoveFromPlaylistRequestDto {
	return v.value
}

func (v *NullableRemoveFromPlaylistRequestDto) Set(val *RemoveFromPlaylistRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveFromPlaylistRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveFromPlaylistRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveFromPlaylistRequestDto(val *RemoveFromPlaylistRequestDto) *NullableRemoveFromPlaylistRequestDto {
	return &NullableRemoveFromPlaylistRequestDto{value: val, isSet: true}
}

func (v NullableRemoveFromPlaylistRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveFromPlaylistRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


