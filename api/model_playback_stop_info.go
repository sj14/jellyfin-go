/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PlaybackStopInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaybackStopInfo{}

// PlaybackStopInfo Class PlaybackStopInfo.
type PlaybackStopInfo struct {
	// Gets or sets the item.
	Item NullableBaseItemDto `json:"Item,omitempty"`
	// Gets or sets the item identifier.
	ItemId *string `json:"ItemId,omitempty"`
	// Gets or sets the session id.
	SessionId NullableString `json:"SessionId,omitempty"`
	// Gets or sets the media version identifier.
	MediaSourceId NullableString `json:"MediaSourceId,omitempty"`
	// Gets or sets the position ticks.
	PositionTicks NullableInt64 `json:"PositionTicks,omitempty"`
	// Gets or sets the live stream identifier.
	LiveStreamId NullableString `json:"LiveStreamId,omitempty"`
	// Gets or sets the play session identifier.
	PlaySessionId NullableString `json:"PlaySessionId,omitempty"`
	// Gets or sets a value indicating whether this MediaBrowser.Model.Session.PlaybackStopInfo is failed.
	Failed *bool `json:"Failed,omitempty"`
	NextMediaType NullableString `json:"NextMediaType,omitempty"`
	PlaylistItemId NullableString `json:"PlaylistItemId,omitempty"`
	NowPlayingQueue []QueueItem `json:"NowPlayingQueue,omitempty"`
}

// NewPlaybackStopInfo instantiates a new PlaybackStopInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaybackStopInfo() *PlaybackStopInfo {
	this := PlaybackStopInfo{}
	return &this
}

// NewPlaybackStopInfoWithDefaults instantiates a new PlaybackStopInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaybackStopInfoWithDefaults() *PlaybackStopInfo {
	this := PlaybackStopInfo{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackStopInfo) GetItem() BaseItemDto {
	if o == nil || IsNil(o.Item.Get()) {
		var ret BaseItemDto
		return ret
	}
	return *o.Item.Get()
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackStopInfo) GetItemOk() (*BaseItemDto, bool) {
	if o == nil {
		return nil, false
	}
	return o.Item.Get(), o.Item.IsSet()
}

// HasItem returns a boolean if a field has been set.
func (o *PlaybackStopInfo) HasItem() bool {
	if o != nil && o.Item.IsSet() {
		return true
	}

	return false
}

// SetItem gets a reference to the given NullableBaseItemDto and assigns it to the Item field.
func (o *PlaybackStopInfo) SetItem(v BaseItemDto) {
	o.Item.Set(&v)
}
// SetItemNil sets the value for Item to be an explicit nil
func (o *PlaybackStopInfo) SetItemNil() {
	o.Item.Set(nil)
}

// UnsetItem ensures that no value is present for Item, not even an explicit nil
func (o *PlaybackStopInfo) UnsetItem() {
	o.Item.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *PlaybackStopInfo) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackStopInfo) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *PlaybackStopInfo) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *PlaybackStopInfo) SetItemId(v string) {
	o.ItemId = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackStopInfo) GetSessionId() string {
	if o == nil || IsNil(o.SessionId.Get()) {
		var ret string
		return ret
	}
	return *o.SessionId.Get()
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackStopInfo) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionId.Get(), o.SessionId.IsSet()
}

// HasSessionId returns a boolean if a field has been set.
func (o *PlaybackStopInfo) HasSessionId() bool {
	if o != nil && o.SessionId.IsSet() {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given NullableString and assigns it to the SessionId field.
func (o *PlaybackStopInfo) SetSessionId(v string) {
	o.SessionId.Set(&v)
}
// SetSessionIdNil sets the value for SessionId to be an explicit nil
func (o *PlaybackStopInfo) SetSessionIdNil() {
	o.SessionId.Set(nil)
}

// UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
func (o *PlaybackStopInfo) UnsetSessionId() {
	o.SessionId.Unset()
}

// GetMediaSourceId returns the MediaSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackStopInfo) GetMediaSourceId() string {
	if o == nil || IsNil(o.MediaSourceId.Get()) {
		var ret string
		return ret
	}
	return *o.MediaSourceId.Get()
}

// GetMediaSourceIdOk returns a tuple with the MediaSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackStopInfo) GetMediaSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaSourceId.Get(), o.MediaSourceId.IsSet()
}

// HasMediaSourceId returns a boolean if a field has been set.
func (o *PlaybackStopInfo) HasMediaSourceId() bool {
	if o != nil && o.MediaSourceId.IsSet() {
		return true
	}

	return false
}

// SetMediaSourceId gets a reference to the given NullableString and assigns it to the MediaSourceId field.
func (o *PlaybackStopInfo) SetMediaSourceId(v string) {
	o.MediaSourceId.Set(&v)
}
// SetMediaSourceIdNil sets the value for MediaSourceId to be an explicit nil
func (o *PlaybackStopInfo) SetMediaSourceIdNil() {
	o.MediaSourceId.Set(nil)
}

// UnsetMediaSourceId ensures that no value is present for MediaSourceId, not even an explicit nil
func (o *PlaybackStopInfo) UnsetMediaSourceId() {
	o.MediaSourceId.Unset()
}

// GetPositionTicks returns the PositionTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackStopInfo) GetPositionTicks() int64 {
	if o == nil || IsNil(o.PositionTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.PositionTicks.Get()
}

// GetPositionTicksOk returns a tuple with the PositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackStopInfo) GetPositionTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PositionTicks.Get(), o.PositionTicks.IsSet()
}

// HasPositionTicks returns a boolean if a field has been set.
func (o *PlaybackStopInfo) HasPositionTicks() bool {
	if o != nil && o.PositionTicks.IsSet() {
		return true
	}

	return false
}

// SetPositionTicks gets a reference to the given NullableInt64 and assigns it to the PositionTicks field.
func (o *PlaybackStopInfo) SetPositionTicks(v int64) {
	o.PositionTicks.Set(&v)
}
// SetPositionTicksNil sets the value for PositionTicks to be an explicit nil
func (o *PlaybackStopInfo) SetPositionTicksNil() {
	o.PositionTicks.Set(nil)
}

// UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
func (o *PlaybackStopInfo) UnsetPositionTicks() {
	o.PositionTicks.Unset()
}

// GetLiveStreamId returns the LiveStreamId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackStopInfo) GetLiveStreamId() string {
	if o == nil || IsNil(o.LiveStreamId.Get()) {
		var ret string
		return ret
	}
	return *o.LiveStreamId.Get()
}

// GetLiveStreamIdOk returns a tuple with the LiveStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackStopInfo) GetLiveStreamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LiveStreamId.Get(), o.LiveStreamId.IsSet()
}

// HasLiveStreamId returns a boolean if a field has been set.
func (o *PlaybackStopInfo) HasLiveStreamId() bool {
	if o != nil && o.LiveStreamId.IsSet() {
		return true
	}

	return false
}

// SetLiveStreamId gets a reference to the given NullableString and assigns it to the LiveStreamId field.
func (o *PlaybackStopInfo) SetLiveStreamId(v string) {
	o.LiveStreamId.Set(&v)
}
// SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil
func (o *PlaybackStopInfo) SetLiveStreamIdNil() {
	o.LiveStreamId.Set(nil)
}

// UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil
func (o *PlaybackStopInfo) UnsetLiveStreamId() {
	o.LiveStreamId.Unset()
}

// GetPlaySessionId returns the PlaySessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackStopInfo) GetPlaySessionId() string {
	if o == nil || IsNil(o.PlaySessionId.Get()) {
		var ret string
		return ret
	}
	return *o.PlaySessionId.Get()
}

// GetPlaySessionIdOk returns a tuple with the PlaySessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackStopInfo) GetPlaySessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaySessionId.Get(), o.PlaySessionId.IsSet()
}

// HasPlaySessionId returns a boolean if a field has been set.
func (o *PlaybackStopInfo) HasPlaySessionId() bool {
	if o != nil && o.PlaySessionId.IsSet() {
		return true
	}

	return false
}

// SetPlaySessionId gets a reference to the given NullableString and assigns it to the PlaySessionId field.
func (o *PlaybackStopInfo) SetPlaySessionId(v string) {
	o.PlaySessionId.Set(&v)
}
// SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil
func (o *PlaybackStopInfo) SetPlaySessionIdNil() {
	o.PlaySessionId.Set(nil)
}

// UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
func (o *PlaybackStopInfo) UnsetPlaySessionId() {
	o.PlaySessionId.Unset()
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *PlaybackStopInfo) GetFailed() bool {
	if o == nil || IsNil(o.Failed) {
		var ret bool
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackStopInfo) GetFailedOk() (*bool, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *PlaybackStopInfo) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given bool and assigns it to the Failed field.
func (o *PlaybackStopInfo) SetFailed(v bool) {
	o.Failed = &v
}

// GetNextMediaType returns the NextMediaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackStopInfo) GetNextMediaType() string {
	if o == nil || IsNil(o.NextMediaType.Get()) {
		var ret string
		return ret
	}
	return *o.NextMediaType.Get()
}

// GetNextMediaTypeOk returns a tuple with the NextMediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackStopInfo) GetNextMediaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextMediaType.Get(), o.NextMediaType.IsSet()
}

// HasNextMediaType returns a boolean if a field has been set.
func (o *PlaybackStopInfo) HasNextMediaType() bool {
	if o != nil && o.NextMediaType.IsSet() {
		return true
	}

	return false
}

// SetNextMediaType gets a reference to the given NullableString and assigns it to the NextMediaType field.
func (o *PlaybackStopInfo) SetNextMediaType(v string) {
	o.NextMediaType.Set(&v)
}
// SetNextMediaTypeNil sets the value for NextMediaType to be an explicit nil
func (o *PlaybackStopInfo) SetNextMediaTypeNil() {
	o.NextMediaType.Set(nil)
}

// UnsetNextMediaType ensures that no value is present for NextMediaType, not even an explicit nil
func (o *PlaybackStopInfo) UnsetNextMediaType() {
	o.NextMediaType.Unset()
}

// GetPlaylistItemId returns the PlaylistItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackStopInfo) GetPlaylistItemId() string {
	if o == nil || IsNil(o.PlaylistItemId.Get()) {
		var ret string
		return ret
	}
	return *o.PlaylistItemId.Get()
}

// GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackStopInfo) GetPlaylistItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaylistItemId.Get(), o.PlaylistItemId.IsSet()
}

// HasPlaylistItemId returns a boolean if a field has been set.
func (o *PlaybackStopInfo) HasPlaylistItemId() bool {
	if o != nil && o.PlaylistItemId.IsSet() {
		return true
	}

	return false
}

// SetPlaylistItemId gets a reference to the given NullableString and assigns it to the PlaylistItemId field.
func (o *PlaybackStopInfo) SetPlaylistItemId(v string) {
	o.PlaylistItemId.Set(&v)
}
// SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil
func (o *PlaybackStopInfo) SetPlaylistItemIdNil() {
	o.PlaylistItemId.Set(nil)
}

// UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil
func (o *PlaybackStopInfo) UnsetPlaylistItemId() {
	o.PlaylistItemId.Unset()
}

// GetNowPlayingQueue returns the NowPlayingQueue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackStopInfo) GetNowPlayingQueue() []QueueItem {
	if o == nil {
		var ret []QueueItem
		return ret
	}
	return o.NowPlayingQueue
}

// GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackStopInfo) GetNowPlayingQueueOk() ([]QueueItem, bool) {
	if o == nil || IsNil(o.NowPlayingQueue) {
		return nil, false
	}
	return o.NowPlayingQueue, true
}

// HasNowPlayingQueue returns a boolean if a field has been set.
func (o *PlaybackStopInfo) HasNowPlayingQueue() bool {
	if o != nil && !IsNil(o.NowPlayingQueue) {
		return true
	}

	return false
}

// SetNowPlayingQueue gets a reference to the given []QueueItem and assigns it to the NowPlayingQueue field.
func (o *PlaybackStopInfo) SetNowPlayingQueue(v []QueueItem) {
	o.NowPlayingQueue = v
}

func (o PlaybackStopInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaybackStopInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Item.IsSet() {
		toSerialize["Item"] = o.Item.Get()
	}
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
	}
	if o.SessionId.IsSet() {
		toSerialize["SessionId"] = o.SessionId.Get()
	}
	if o.MediaSourceId.IsSet() {
		toSerialize["MediaSourceId"] = o.MediaSourceId.Get()
	}
	if o.PositionTicks.IsSet() {
		toSerialize["PositionTicks"] = o.PositionTicks.Get()
	}
	if o.LiveStreamId.IsSet() {
		toSerialize["LiveStreamId"] = o.LiveStreamId.Get()
	}
	if o.PlaySessionId.IsSet() {
		toSerialize["PlaySessionId"] = o.PlaySessionId.Get()
	}
	if !IsNil(o.Failed) {
		toSerialize["Failed"] = o.Failed
	}
	if o.NextMediaType.IsSet() {
		toSerialize["NextMediaType"] = o.NextMediaType.Get()
	}
	if o.PlaylistItemId.IsSet() {
		toSerialize["PlaylistItemId"] = o.PlaylistItemId.Get()
	}
	if o.NowPlayingQueue != nil {
		toSerialize["NowPlayingQueue"] = o.NowPlayingQueue
	}
	return toSerialize, nil
}

type NullablePlaybackStopInfo struct {
	value *PlaybackStopInfo
	isSet bool
}

func (v NullablePlaybackStopInfo) Get() *PlaybackStopInfo {
	return v.value
}

func (v *NullablePlaybackStopInfo) Set(val *PlaybackStopInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaybackStopInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaybackStopInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaybackStopInfo(val *PlaybackStopInfo) *NullablePlaybackStopInfo {
	return &NullablePlaybackStopInfo{value: val, isSet: true}
}

func (v NullablePlaybackStopInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaybackStopInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


