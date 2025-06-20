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

// checks if the PlayerStateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayerStateInfo{}

// PlayerStateInfo struct for PlayerStateInfo
type PlayerStateInfo struct {
	// Gets or sets the now playing position ticks.
	PositionTicks NullableInt64 `json:"PositionTicks,omitempty"`
	// Gets or sets a value indicating whether this instance can seek.
	CanSeek *bool `json:"CanSeek,omitempty"`
	// Gets or sets a value indicating whether this instance is paused.
	IsPaused *bool `json:"IsPaused,omitempty"`
	// Gets or sets a value indicating whether this instance is muted.
	IsMuted *bool `json:"IsMuted,omitempty"`
	// Gets or sets the volume level.
	VolumeLevel NullableInt32 `json:"VolumeLevel,omitempty"`
	// Gets or sets the index of the now playing audio stream.
	AudioStreamIndex NullableInt32 `json:"AudioStreamIndex,omitempty"`
	// Gets or sets the index of the now playing subtitle stream.
	SubtitleStreamIndex NullableInt32 `json:"SubtitleStreamIndex,omitempty"`
	// Gets or sets the now playing media version identifier.
	MediaSourceId NullableString `json:"MediaSourceId,omitempty"`
	// Gets or sets the play method.
	PlayMethod NullablePlayMethod `json:"PlayMethod,omitempty"`
	// Gets or sets the repeat mode.
	RepeatMode *RepeatMode `json:"RepeatMode,omitempty"`
	// Gets or sets the playback order.
	PlaybackOrder *PlaybackOrder `json:"PlaybackOrder,omitempty"`
	// Gets or sets the now playing live stream identifier.
	LiveStreamId NullableString `json:"LiveStreamId,omitempty"`
}

// NewPlayerStateInfo instantiates a new PlayerStateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayerStateInfo() *PlayerStateInfo {
	this := PlayerStateInfo{}
	return &this
}

// NewPlayerStateInfoWithDefaults instantiates a new PlayerStateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayerStateInfoWithDefaults() *PlayerStateInfo {
	this := PlayerStateInfo{}
	return &this
}

// GetPositionTicks returns the PositionTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerStateInfo) GetPositionTicks() int64 {
	if o == nil || IsNil(o.PositionTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.PositionTicks.Get()
}

// GetPositionTicksOk returns a tuple with the PositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerStateInfo) GetPositionTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PositionTicks.Get(), o.PositionTicks.IsSet()
}

// HasPositionTicks returns a boolean if a field has been set.
func (o *PlayerStateInfo) HasPositionTicks() bool {
	if o != nil && o.PositionTicks.IsSet() {
		return true
	}

	return false
}

// SetPositionTicks gets a reference to the given NullableInt64 and assigns it to the PositionTicks field.
func (o *PlayerStateInfo) SetPositionTicks(v int64) {
	o.PositionTicks.Set(&v)
}
// SetPositionTicksNil sets the value for PositionTicks to be an explicit nil
func (o *PlayerStateInfo) SetPositionTicksNil() {
	o.PositionTicks.Set(nil)
}

// UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
func (o *PlayerStateInfo) UnsetPositionTicks() {
	o.PositionTicks.Unset()
}

// GetCanSeek returns the CanSeek field value if set, zero value otherwise.
func (o *PlayerStateInfo) GetCanSeek() bool {
	if o == nil || IsNil(o.CanSeek) {
		var ret bool
		return ret
	}
	return *o.CanSeek
}

// GetCanSeekOk returns a tuple with the CanSeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerStateInfo) GetCanSeekOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSeek) {
		return nil, false
	}
	return o.CanSeek, true
}

// HasCanSeek returns a boolean if a field has been set.
func (o *PlayerStateInfo) HasCanSeek() bool {
	if o != nil && !IsNil(o.CanSeek) {
		return true
	}

	return false
}

// SetCanSeek gets a reference to the given bool and assigns it to the CanSeek field.
func (o *PlayerStateInfo) SetCanSeek(v bool) {
	o.CanSeek = &v
}

// GetIsPaused returns the IsPaused field value if set, zero value otherwise.
func (o *PlayerStateInfo) GetIsPaused() bool {
	if o == nil || IsNil(o.IsPaused) {
		var ret bool
		return ret
	}
	return *o.IsPaused
}

// GetIsPausedOk returns a tuple with the IsPaused field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerStateInfo) GetIsPausedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPaused) {
		return nil, false
	}
	return o.IsPaused, true
}

// HasIsPaused returns a boolean if a field has been set.
func (o *PlayerStateInfo) HasIsPaused() bool {
	if o != nil && !IsNil(o.IsPaused) {
		return true
	}

	return false
}

// SetIsPaused gets a reference to the given bool and assigns it to the IsPaused field.
func (o *PlayerStateInfo) SetIsPaused(v bool) {
	o.IsPaused = &v
}

// GetIsMuted returns the IsMuted field value if set, zero value otherwise.
func (o *PlayerStateInfo) GetIsMuted() bool {
	if o == nil || IsNil(o.IsMuted) {
		var ret bool
		return ret
	}
	return *o.IsMuted
}

// GetIsMutedOk returns a tuple with the IsMuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerStateInfo) GetIsMutedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMuted) {
		return nil, false
	}
	return o.IsMuted, true
}

// HasIsMuted returns a boolean if a field has been set.
func (o *PlayerStateInfo) HasIsMuted() bool {
	if o != nil && !IsNil(o.IsMuted) {
		return true
	}

	return false
}

// SetIsMuted gets a reference to the given bool and assigns it to the IsMuted field.
func (o *PlayerStateInfo) SetIsMuted(v bool) {
	o.IsMuted = &v
}

// GetVolumeLevel returns the VolumeLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerStateInfo) GetVolumeLevel() int32 {
	if o == nil || IsNil(o.VolumeLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.VolumeLevel.Get()
}

// GetVolumeLevelOk returns a tuple with the VolumeLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerStateInfo) GetVolumeLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VolumeLevel.Get(), o.VolumeLevel.IsSet()
}

// HasVolumeLevel returns a boolean if a field has been set.
func (o *PlayerStateInfo) HasVolumeLevel() bool {
	if o != nil && o.VolumeLevel.IsSet() {
		return true
	}

	return false
}

// SetVolumeLevel gets a reference to the given NullableInt32 and assigns it to the VolumeLevel field.
func (o *PlayerStateInfo) SetVolumeLevel(v int32) {
	o.VolumeLevel.Set(&v)
}
// SetVolumeLevelNil sets the value for VolumeLevel to be an explicit nil
func (o *PlayerStateInfo) SetVolumeLevelNil() {
	o.VolumeLevel.Set(nil)
}

// UnsetVolumeLevel ensures that no value is present for VolumeLevel, not even an explicit nil
func (o *PlayerStateInfo) UnsetVolumeLevel() {
	o.VolumeLevel.Unset()
}

// GetAudioStreamIndex returns the AudioStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerStateInfo) GetAudioStreamIndex() int32 {
	if o == nil || IsNil(o.AudioStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.AudioStreamIndex.Get()
}

// GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerStateInfo) GetAudioStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioStreamIndex.Get(), o.AudioStreamIndex.IsSet()
}

// HasAudioStreamIndex returns a boolean if a field has been set.
func (o *PlayerStateInfo) HasAudioStreamIndex() bool {
	if o != nil && o.AudioStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetAudioStreamIndex gets a reference to the given NullableInt32 and assigns it to the AudioStreamIndex field.
func (o *PlayerStateInfo) SetAudioStreamIndex(v int32) {
	o.AudioStreamIndex.Set(&v)
}
// SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil
func (o *PlayerStateInfo) SetAudioStreamIndexNil() {
	o.AudioStreamIndex.Set(nil)
}

// UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
func (o *PlayerStateInfo) UnsetAudioStreamIndex() {
	o.AudioStreamIndex.Unset()
}

// GetSubtitleStreamIndex returns the SubtitleStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerStateInfo) GetSubtitleStreamIndex() int32 {
	if o == nil || IsNil(o.SubtitleStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.SubtitleStreamIndex.Get()
}

// GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerStateInfo) GetSubtitleStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubtitleStreamIndex.Get(), o.SubtitleStreamIndex.IsSet()
}

// HasSubtitleStreamIndex returns a boolean if a field has been set.
func (o *PlayerStateInfo) HasSubtitleStreamIndex() bool {
	if o != nil && o.SubtitleStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetSubtitleStreamIndex gets a reference to the given NullableInt32 and assigns it to the SubtitleStreamIndex field.
func (o *PlayerStateInfo) SetSubtitleStreamIndex(v int32) {
	o.SubtitleStreamIndex.Set(&v)
}
// SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil
func (o *PlayerStateInfo) SetSubtitleStreamIndexNil() {
	o.SubtitleStreamIndex.Set(nil)
}

// UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
func (o *PlayerStateInfo) UnsetSubtitleStreamIndex() {
	o.SubtitleStreamIndex.Unset()
}

// GetMediaSourceId returns the MediaSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerStateInfo) GetMediaSourceId() string {
	if o == nil || IsNil(o.MediaSourceId.Get()) {
		var ret string
		return ret
	}
	return *o.MediaSourceId.Get()
}

// GetMediaSourceIdOk returns a tuple with the MediaSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerStateInfo) GetMediaSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaSourceId.Get(), o.MediaSourceId.IsSet()
}

// HasMediaSourceId returns a boolean if a field has been set.
func (o *PlayerStateInfo) HasMediaSourceId() bool {
	if o != nil && o.MediaSourceId.IsSet() {
		return true
	}

	return false
}

// SetMediaSourceId gets a reference to the given NullableString and assigns it to the MediaSourceId field.
func (o *PlayerStateInfo) SetMediaSourceId(v string) {
	o.MediaSourceId.Set(&v)
}
// SetMediaSourceIdNil sets the value for MediaSourceId to be an explicit nil
func (o *PlayerStateInfo) SetMediaSourceIdNil() {
	o.MediaSourceId.Set(nil)
}

// UnsetMediaSourceId ensures that no value is present for MediaSourceId, not even an explicit nil
func (o *PlayerStateInfo) UnsetMediaSourceId() {
	o.MediaSourceId.Unset()
}

// GetPlayMethod returns the PlayMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerStateInfo) GetPlayMethod() PlayMethod {
	if o == nil || IsNil(o.PlayMethod.Get()) {
		var ret PlayMethod
		return ret
	}
	return *o.PlayMethod.Get()
}

// GetPlayMethodOk returns a tuple with the PlayMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerStateInfo) GetPlayMethodOk() (*PlayMethod, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlayMethod.Get(), o.PlayMethod.IsSet()
}

// HasPlayMethod returns a boolean if a field has been set.
func (o *PlayerStateInfo) HasPlayMethod() bool {
	if o != nil && o.PlayMethod.IsSet() {
		return true
	}

	return false
}

// SetPlayMethod gets a reference to the given NullablePlayMethod and assigns it to the PlayMethod field.
func (o *PlayerStateInfo) SetPlayMethod(v PlayMethod) {
	o.PlayMethod.Set(&v)
}
// SetPlayMethodNil sets the value for PlayMethod to be an explicit nil
func (o *PlayerStateInfo) SetPlayMethodNil() {
	o.PlayMethod.Set(nil)
}

// UnsetPlayMethod ensures that no value is present for PlayMethod, not even an explicit nil
func (o *PlayerStateInfo) UnsetPlayMethod() {
	o.PlayMethod.Unset()
}

// GetRepeatMode returns the RepeatMode field value if set, zero value otherwise.
func (o *PlayerStateInfo) GetRepeatMode() RepeatMode {
	if o == nil || IsNil(o.RepeatMode) {
		var ret RepeatMode
		return ret
	}
	return *o.RepeatMode
}

// GetRepeatModeOk returns a tuple with the RepeatMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerStateInfo) GetRepeatModeOk() (*RepeatMode, bool) {
	if o == nil || IsNil(o.RepeatMode) {
		return nil, false
	}
	return o.RepeatMode, true
}

// HasRepeatMode returns a boolean if a field has been set.
func (o *PlayerStateInfo) HasRepeatMode() bool {
	if o != nil && !IsNil(o.RepeatMode) {
		return true
	}

	return false
}

// SetRepeatMode gets a reference to the given RepeatMode and assigns it to the RepeatMode field.
func (o *PlayerStateInfo) SetRepeatMode(v RepeatMode) {
	o.RepeatMode = &v
}

// GetPlaybackOrder returns the PlaybackOrder field value if set, zero value otherwise.
func (o *PlayerStateInfo) GetPlaybackOrder() PlaybackOrder {
	if o == nil || IsNil(o.PlaybackOrder) {
		var ret PlaybackOrder
		return ret
	}
	return *o.PlaybackOrder
}

// GetPlaybackOrderOk returns a tuple with the PlaybackOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayerStateInfo) GetPlaybackOrderOk() (*PlaybackOrder, bool) {
	if o == nil || IsNil(o.PlaybackOrder) {
		return nil, false
	}
	return o.PlaybackOrder, true
}

// HasPlaybackOrder returns a boolean if a field has been set.
func (o *PlayerStateInfo) HasPlaybackOrder() bool {
	if o != nil && !IsNil(o.PlaybackOrder) {
		return true
	}

	return false
}

// SetPlaybackOrder gets a reference to the given PlaybackOrder and assigns it to the PlaybackOrder field.
func (o *PlayerStateInfo) SetPlaybackOrder(v PlaybackOrder) {
	o.PlaybackOrder = &v
}

// GetLiveStreamId returns the LiveStreamId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlayerStateInfo) GetLiveStreamId() string {
	if o == nil || IsNil(o.LiveStreamId.Get()) {
		var ret string
		return ret
	}
	return *o.LiveStreamId.Get()
}

// GetLiveStreamIdOk returns a tuple with the LiveStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlayerStateInfo) GetLiveStreamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LiveStreamId.Get(), o.LiveStreamId.IsSet()
}

// HasLiveStreamId returns a boolean if a field has been set.
func (o *PlayerStateInfo) HasLiveStreamId() bool {
	if o != nil && o.LiveStreamId.IsSet() {
		return true
	}

	return false
}

// SetLiveStreamId gets a reference to the given NullableString and assigns it to the LiveStreamId field.
func (o *PlayerStateInfo) SetLiveStreamId(v string) {
	o.LiveStreamId.Set(&v)
}
// SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil
func (o *PlayerStateInfo) SetLiveStreamIdNil() {
	o.LiveStreamId.Set(nil)
}

// UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil
func (o *PlayerStateInfo) UnsetLiveStreamId() {
	o.LiveStreamId.Unset()
}

func (o PlayerStateInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayerStateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PositionTicks.IsSet() {
		toSerialize["PositionTicks"] = o.PositionTicks.Get()
	}
	if !IsNil(o.CanSeek) {
		toSerialize["CanSeek"] = o.CanSeek
	}
	if !IsNil(o.IsPaused) {
		toSerialize["IsPaused"] = o.IsPaused
	}
	if !IsNil(o.IsMuted) {
		toSerialize["IsMuted"] = o.IsMuted
	}
	if o.VolumeLevel.IsSet() {
		toSerialize["VolumeLevel"] = o.VolumeLevel.Get()
	}
	if o.AudioStreamIndex.IsSet() {
		toSerialize["AudioStreamIndex"] = o.AudioStreamIndex.Get()
	}
	if o.SubtitleStreamIndex.IsSet() {
		toSerialize["SubtitleStreamIndex"] = o.SubtitleStreamIndex.Get()
	}
	if o.MediaSourceId.IsSet() {
		toSerialize["MediaSourceId"] = o.MediaSourceId.Get()
	}
	if o.PlayMethod.IsSet() {
		toSerialize["PlayMethod"] = o.PlayMethod.Get()
	}
	if !IsNil(o.RepeatMode) {
		toSerialize["RepeatMode"] = o.RepeatMode
	}
	if !IsNil(o.PlaybackOrder) {
		toSerialize["PlaybackOrder"] = o.PlaybackOrder
	}
	if o.LiveStreamId.IsSet() {
		toSerialize["LiveStreamId"] = o.LiveStreamId.Get()
	}
	return toSerialize, nil
}

type NullablePlayerStateInfo struct {
	value *PlayerStateInfo
	isSet bool
}

func (v NullablePlayerStateInfo) Get() *PlayerStateInfo {
	return v.value
}

func (v *NullablePlayerStateInfo) Set(val *PlayerStateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayerStateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayerStateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayerStateInfo(val *PlayerStateInfo) *NullablePlayerStateInfo {
	return &NullablePlayerStateInfo{value: val, isSet: true}
}

func (v NullablePlayerStateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayerStateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


