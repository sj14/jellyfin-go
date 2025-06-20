/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SyncPlayQueueItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncPlayQueueItem{}

// SyncPlayQueueItem Class QueueItem.
type SyncPlayQueueItem struct {
	// Gets the item identifier.
	ItemId *string `json:"ItemId,omitempty"`
	// Gets the playlist identifier of the item.
	PlaylistItemId *string `json:"PlaylistItemId,omitempty"`
}

// NewSyncPlayQueueItem instantiates a new SyncPlayQueueItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncPlayQueueItem() *SyncPlayQueueItem {
	this := SyncPlayQueueItem{}
	return &this
}

// NewSyncPlayQueueItemWithDefaults instantiates a new SyncPlayQueueItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncPlayQueueItemWithDefaults() *SyncPlayQueueItem {
	this := SyncPlayQueueItem{}
	return &this
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *SyncPlayQueueItem) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncPlayQueueItem) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *SyncPlayQueueItem) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *SyncPlayQueueItem) SetItemId(v string) {
	o.ItemId = &v
}

// GetPlaylistItemId returns the PlaylistItemId field value if set, zero value otherwise.
func (o *SyncPlayQueueItem) GetPlaylistItemId() string {
	if o == nil || IsNil(o.PlaylistItemId) {
		var ret string
		return ret
	}
	return *o.PlaylistItemId
}

// GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncPlayQueueItem) GetPlaylistItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaylistItemId) {
		return nil, false
	}
	return o.PlaylistItemId, true
}

// HasPlaylistItemId returns a boolean if a field has been set.
func (o *SyncPlayQueueItem) HasPlaylistItemId() bool {
	if o != nil && !IsNil(o.PlaylistItemId) {
		return true
	}

	return false
}

// SetPlaylistItemId gets a reference to the given string and assigns it to the PlaylistItemId field.
func (o *SyncPlayQueueItem) SetPlaylistItemId(v string) {
	o.PlaylistItemId = &v
}

func (o SyncPlayQueueItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncPlayQueueItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
	}
	if !IsNil(o.PlaylistItemId) {
		toSerialize["PlaylistItemId"] = o.PlaylistItemId
	}
	return toSerialize, nil
}

type NullableSyncPlayQueueItem struct {
	value *SyncPlayQueueItem
	isSet bool
}

func (v NullableSyncPlayQueueItem) Get() *SyncPlayQueueItem {
	return v.value
}

func (v *NullableSyncPlayQueueItem) Set(val *SyncPlayQueueItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncPlayQueueItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncPlayQueueItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncPlayQueueItem(val *SyncPlayQueueItem) *NullableSyncPlayQueueItem {
	return &NullableSyncPlayQueueItem{value: val, isSet: true}
}

func (v NullableSyncPlayQueueItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncPlayQueueItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


