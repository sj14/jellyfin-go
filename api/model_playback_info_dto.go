/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PlaybackInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaybackInfoDto{}

// PlaybackInfoDto Plabyback info dto.
type PlaybackInfoDto struct {
	// Gets or sets the playback userId.
	UserId NullableString `json:"UserId,omitempty"`
	// Gets or sets the max streaming bitrate.
	MaxStreamingBitrate NullableInt32 `json:"MaxStreamingBitrate,omitempty"`
	// Gets or sets the start time in ticks.
	StartTimeTicks NullableInt64 `json:"StartTimeTicks,omitempty"`
	// Gets or sets the audio stream index.
	AudioStreamIndex NullableInt32 `json:"AudioStreamIndex,omitempty"`
	// Gets or sets the subtitle stream index.
	SubtitleStreamIndex NullableInt32 `json:"SubtitleStreamIndex,omitempty"`
	// Gets or sets the max audio channels.
	MaxAudioChannels NullableInt32 `json:"MaxAudioChannels,omitempty"`
	// Gets or sets the media source id.
	MediaSourceId NullableString `json:"MediaSourceId,omitempty"`
	// Gets or sets the live stream id.
	LiveStreamId NullableString `json:"LiveStreamId,omitempty"`
	DeviceProfile NullableClientCapabilitiesDeviceProfile `json:"DeviceProfile,omitempty"`
	// Gets or sets a value indicating whether to enable direct play.
	EnableDirectPlay NullableBool `json:"EnableDirectPlay,omitempty"`
	// Gets or sets a value indicating whether to enable direct stream.
	EnableDirectStream NullableBool `json:"EnableDirectStream,omitempty"`
	// Gets or sets a value indicating whether to enable transcoding.
	EnableTranscoding NullableBool `json:"EnableTranscoding,omitempty"`
	// Gets or sets a value indicating whether to enable video stream copy.
	AllowVideoStreamCopy NullableBool `json:"AllowVideoStreamCopy,omitempty"`
	// Gets or sets a value indicating whether to allow audio stream copy.
	AllowAudioStreamCopy NullableBool `json:"AllowAudioStreamCopy,omitempty"`
	// Gets or sets a value indicating whether to auto open the live stream.
	AutoOpenLiveStream NullableBool `json:"AutoOpenLiveStream,omitempty"`
}

// NewPlaybackInfoDto instantiates a new PlaybackInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaybackInfoDto() *PlaybackInfoDto {
	this := PlaybackInfoDto{}
	return &this
}

// NewPlaybackInfoDtoWithDefaults instantiates a new PlaybackInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaybackInfoDtoWithDefaults() *PlaybackInfoDto {
	this := PlaybackInfoDto{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *PlaybackInfoDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *PlaybackInfoDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *PlaybackInfoDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetMaxStreamingBitrate returns the MaxStreamingBitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetMaxStreamingBitrate() int32 {
	if o == nil || IsNil(o.MaxStreamingBitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxStreamingBitrate.Get()
}

// GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetMaxStreamingBitrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStreamingBitrate.Get(), o.MaxStreamingBitrate.IsSet()
}

// HasMaxStreamingBitrate returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasMaxStreamingBitrate() bool {
	if o != nil && o.MaxStreamingBitrate.IsSet() {
		return true
	}

	return false
}

// SetMaxStreamingBitrate gets a reference to the given NullableInt32 and assigns it to the MaxStreamingBitrate field.
func (o *PlaybackInfoDto) SetMaxStreamingBitrate(v int32) {
	o.MaxStreamingBitrate.Set(&v)
}
// SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil
func (o *PlaybackInfoDto) SetMaxStreamingBitrateNil() {
	o.MaxStreamingBitrate.Set(nil)
}

// UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
func (o *PlaybackInfoDto) UnsetMaxStreamingBitrate() {
	o.MaxStreamingBitrate.Unset()
}

// GetStartTimeTicks returns the StartTimeTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetStartTimeTicks() int64 {
	if o == nil || IsNil(o.StartTimeTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.StartTimeTicks.Get()
}

// GetStartTimeTicksOk returns a tuple with the StartTimeTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetStartTimeTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTimeTicks.Get(), o.StartTimeTicks.IsSet()
}

// HasStartTimeTicks returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasStartTimeTicks() bool {
	if o != nil && o.StartTimeTicks.IsSet() {
		return true
	}

	return false
}

// SetStartTimeTicks gets a reference to the given NullableInt64 and assigns it to the StartTimeTicks field.
func (o *PlaybackInfoDto) SetStartTimeTicks(v int64) {
	o.StartTimeTicks.Set(&v)
}
// SetStartTimeTicksNil sets the value for StartTimeTicks to be an explicit nil
func (o *PlaybackInfoDto) SetStartTimeTicksNil() {
	o.StartTimeTicks.Set(nil)
}

// UnsetStartTimeTicks ensures that no value is present for StartTimeTicks, not even an explicit nil
func (o *PlaybackInfoDto) UnsetStartTimeTicks() {
	o.StartTimeTicks.Unset()
}

// GetAudioStreamIndex returns the AudioStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetAudioStreamIndex() int32 {
	if o == nil || IsNil(o.AudioStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.AudioStreamIndex.Get()
}

// GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetAudioStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioStreamIndex.Get(), o.AudioStreamIndex.IsSet()
}

// HasAudioStreamIndex returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasAudioStreamIndex() bool {
	if o != nil && o.AudioStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetAudioStreamIndex gets a reference to the given NullableInt32 and assigns it to the AudioStreamIndex field.
func (o *PlaybackInfoDto) SetAudioStreamIndex(v int32) {
	o.AudioStreamIndex.Set(&v)
}
// SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil
func (o *PlaybackInfoDto) SetAudioStreamIndexNil() {
	o.AudioStreamIndex.Set(nil)
}

// UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
func (o *PlaybackInfoDto) UnsetAudioStreamIndex() {
	o.AudioStreamIndex.Unset()
}

// GetSubtitleStreamIndex returns the SubtitleStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetSubtitleStreamIndex() int32 {
	if o == nil || IsNil(o.SubtitleStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.SubtitleStreamIndex.Get()
}

// GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetSubtitleStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubtitleStreamIndex.Get(), o.SubtitleStreamIndex.IsSet()
}

// HasSubtitleStreamIndex returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasSubtitleStreamIndex() bool {
	if o != nil && o.SubtitleStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetSubtitleStreamIndex gets a reference to the given NullableInt32 and assigns it to the SubtitleStreamIndex field.
func (o *PlaybackInfoDto) SetSubtitleStreamIndex(v int32) {
	o.SubtitleStreamIndex.Set(&v)
}
// SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil
func (o *PlaybackInfoDto) SetSubtitleStreamIndexNil() {
	o.SubtitleStreamIndex.Set(nil)
}

// UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
func (o *PlaybackInfoDto) UnsetSubtitleStreamIndex() {
	o.SubtitleStreamIndex.Unset()
}

// GetMaxAudioChannels returns the MaxAudioChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetMaxAudioChannels() int32 {
	if o == nil || IsNil(o.MaxAudioChannels.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxAudioChannels.Get()
}

// GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetMaxAudioChannelsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAudioChannels.Get(), o.MaxAudioChannels.IsSet()
}

// HasMaxAudioChannels returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasMaxAudioChannels() bool {
	if o != nil && o.MaxAudioChannels.IsSet() {
		return true
	}

	return false
}

// SetMaxAudioChannels gets a reference to the given NullableInt32 and assigns it to the MaxAudioChannels field.
func (o *PlaybackInfoDto) SetMaxAudioChannels(v int32) {
	o.MaxAudioChannels.Set(&v)
}
// SetMaxAudioChannelsNil sets the value for MaxAudioChannels to be an explicit nil
func (o *PlaybackInfoDto) SetMaxAudioChannelsNil() {
	o.MaxAudioChannels.Set(nil)
}

// UnsetMaxAudioChannels ensures that no value is present for MaxAudioChannels, not even an explicit nil
func (o *PlaybackInfoDto) UnsetMaxAudioChannels() {
	o.MaxAudioChannels.Unset()
}

// GetMediaSourceId returns the MediaSourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetMediaSourceId() string {
	if o == nil || IsNil(o.MediaSourceId.Get()) {
		var ret string
		return ret
	}
	return *o.MediaSourceId.Get()
}

// GetMediaSourceIdOk returns a tuple with the MediaSourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetMediaSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaSourceId.Get(), o.MediaSourceId.IsSet()
}

// HasMediaSourceId returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasMediaSourceId() bool {
	if o != nil && o.MediaSourceId.IsSet() {
		return true
	}

	return false
}

// SetMediaSourceId gets a reference to the given NullableString and assigns it to the MediaSourceId field.
func (o *PlaybackInfoDto) SetMediaSourceId(v string) {
	o.MediaSourceId.Set(&v)
}
// SetMediaSourceIdNil sets the value for MediaSourceId to be an explicit nil
func (o *PlaybackInfoDto) SetMediaSourceIdNil() {
	o.MediaSourceId.Set(nil)
}

// UnsetMediaSourceId ensures that no value is present for MediaSourceId, not even an explicit nil
func (o *PlaybackInfoDto) UnsetMediaSourceId() {
	o.MediaSourceId.Unset()
}

// GetLiveStreamId returns the LiveStreamId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetLiveStreamId() string {
	if o == nil || IsNil(o.LiveStreamId.Get()) {
		var ret string
		return ret
	}
	return *o.LiveStreamId.Get()
}

// GetLiveStreamIdOk returns a tuple with the LiveStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetLiveStreamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LiveStreamId.Get(), o.LiveStreamId.IsSet()
}

// HasLiveStreamId returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasLiveStreamId() bool {
	if o != nil && o.LiveStreamId.IsSet() {
		return true
	}

	return false
}

// SetLiveStreamId gets a reference to the given NullableString and assigns it to the LiveStreamId field.
func (o *PlaybackInfoDto) SetLiveStreamId(v string) {
	o.LiveStreamId.Set(&v)
}
// SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil
func (o *PlaybackInfoDto) SetLiveStreamIdNil() {
	o.LiveStreamId.Set(nil)
}

// UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil
func (o *PlaybackInfoDto) UnsetLiveStreamId() {
	o.LiveStreamId.Unset()
}

// GetDeviceProfile returns the DeviceProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetDeviceProfile() ClientCapabilitiesDeviceProfile {
	if o == nil || IsNil(o.DeviceProfile.Get()) {
		var ret ClientCapabilitiesDeviceProfile
		return ret
	}
	return *o.DeviceProfile.Get()
}

// GetDeviceProfileOk returns a tuple with the DeviceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetDeviceProfileOk() (*ClientCapabilitiesDeviceProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceProfile.Get(), o.DeviceProfile.IsSet()
}

// HasDeviceProfile returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasDeviceProfile() bool {
	if o != nil && o.DeviceProfile.IsSet() {
		return true
	}

	return false
}

// SetDeviceProfile gets a reference to the given NullableClientCapabilitiesDeviceProfile and assigns it to the DeviceProfile field.
func (o *PlaybackInfoDto) SetDeviceProfile(v ClientCapabilitiesDeviceProfile) {
	o.DeviceProfile.Set(&v)
}
// SetDeviceProfileNil sets the value for DeviceProfile to be an explicit nil
func (o *PlaybackInfoDto) SetDeviceProfileNil() {
	o.DeviceProfile.Set(nil)
}

// UnsetDeviceProfile ensures that no value is present for DeviceProfile, not even an explicit nil
func (o *PlaybackInfoDto) UnsetDeviceProfile() {
	o.DeviceProfile.Unset()
}

// GetEnableDirectPlay returns the EnableDirectPlay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetEnableDirectPlay() bool {
	if o == nil || IsNil(o.EnableDirectPlay.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableDirectPlay.Get()
}

// GetEnableDirectPlayOk returns a tuple with the EnableDirectPlay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetEnableDirectPlayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableDirectPlay.Get(), o.EnableDirectPlay.IsSet()
}

// HasEnableDirectPlay returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasEnableDirectPlay() bool {
	if o != nil && o.EnableDirectPlay.IsSet() {
		return true
	}

	return false
}

// SetEnableDirectPlay gets a reference to the given NullableBool and assigns it to the EnableDirectPlay field.
func (o *PlaybackInfoDto) SetEnableDirectPlay(v bool) {
	o.EnableDirectPlay.Set(&v)
}
// SetEnableDirectPlayNil sets the value for EnableDirectPlay to be an explicit nil
func (o *PlaybackInfoDto) SetEnableDirectPlayNil() {
	o.EnableDirectPlay.Set(nil)
}

// UnsetEnableDirectPlay ensures that no value is present for EnableDirectPlay, not even an explicit nil
func (o *PlaybackInfoDto) UnsetEnableDirectPlay() {
	o.EnableDirectPlay.Unset()
}

// GetEnableDirectStream returns the EnableDirectStream field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetEnableDirectStream() bool {
	if o == nil || IsNil(o.EnableDirectStream.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableDirectStream.Get()
}

// GetEnableDirectStreamOk returns a tuple with the EnableDirectStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetEnableDirectStreamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableDirectStream.Get(), o.EnableDirectStream.IsSet()
}

// HasEnableDirectStream returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasEnableDirectStream() bool {
	if o != nil && o.EnableDirectStream.IsSet() {
		return true
	}

	return false
}

// SetEnableDirectStream gets a reference to the given NullableBool and assigns it to the EnableDirectStream field.
func (o *PlaybackInfoDto) SetEnableDirectStream(v bool) {
	o.EnableDirectStream.Set(&v)
}
// SetEnableDirectStreamNil sets the value for EnableDirectStream to be an explicit nil
func (o *PlaybackInfoDto) SetEnableDirectStreamNil() {
	o.EnableDirectStream.Set(nil)
}

// UnsetEnableDirectStream ensures that no value is present for EnableDirectStream, not even an explicit nil
func (o *PlaybackInfoDto) UnsetEnableDirectStream() {
	o.EnableDirectStream.Unset()
}

// GetEnableTranscoding returns the EnableTranscoding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetEnableTranscoding() bool {
	if o == nil || IsNil(o.EnableTranscoding.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableTranscoding.Get()
}

// GetEnableTranscodingOk returns a tuple with the EnableTranscoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetEnableTranscodingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableTranscoding.Get(), o.EnableTranscoding.IsSet()
}

// HasEnableTranscoding returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasEnableTranscoding() bool {
	if o != nil && o.EnableTranscoding.IsSet() {
		return true
	}

	return false
}

// SetEnableTranscoding gets a reference to the given NullableBool and assigns it to the EnableTranscoding field.
func (o *PlaybackInfoDto) SetEnableTranscoding(v bool) {
	o.EnableTranscoding.Set(&v)
}
// SetEnableTranscodingNil sets the value for EnableTranscoding to be an explicit nil
func (o *PlaybackInfoDto) SetEnableTranscodingNil() {
	o.EnableTranscoding.Set(nil)
}

// UnsetEnableTranscoding ensures that no value is present for EnableTranscoding, not even an explicit nil
func (o *PlaybackInfoDto) UnsetEnableTranscoding() {
	o.EnableTranscoding.Unset()
}

// GetAllowVideoStreamCopy returns the AllowVideoStreamCopy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetAllowVideoStreamCopy() bool {
	if o == nil || IsNil(o.AllowVideoStreamCopy.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowVideoStreamCopy.Get()
}

// GetAllowVideoStreamCopyOk returns a tuple with the AllowVideoStreamCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetAllowVideoStreamCopyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowVideoStreamCopy.Get(), o.AllowVideoStreamCopy.IsSet()
}

// HasAllowVideoStreamCopy returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasAllowVideoStreamCopy() bool {
	if o != nil && o.AllowVideoStreamCopy.IsSet() {
		return true
	}

	return false
}

// SetAllowVideoStreamCopy gets a reference to the given NullableBool and assigns it to the AllowVideoStreamCopy field.
func (o *PlaybackInfoDto) SetAllowVideoStreamCopy(v bool) {
	o.AllowVideoStreamCopy.Set(&v)
}
// SetAllowVideoStreamCopyNil sets the value for AllowVideoStreamCopy to be an explicit nil
func (o *PlaybackInfoDto) SetAllowVideoStreamCopyNil() {
	o.AllowVideoStreamCopy.Set(nil)
}

// UnsetAllowVideoStreamCopy ensures that no value is present for AllowVideoStreamCopy, not even an explicit nil
func (o *PlaybackInfoDto) UnsetAllowVideoStreamCopy() {
	o.AllowVideoStreamCopy.Unset()
}

// GetAllowAudioStreamCopy returns the AllowAudioStreamCopy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetAllowAudioStreamCopy() bool {
	if o == nil || IsNil(o.AllowAudioStreamCopy.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowAudioStreamCopy.Get()
}

// GetAllowAudioStreamCopyOk returns a tuple with the AllowAudioStreamCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetAllowAudioStreamCopyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowAudioStreamCopy.Get(), o.AllowAudioStreamCopy.IsSet()
}

// HasAllowAudioStreamCopy returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasAllowAudioStreamCopy() bool {
	if o != nil && o.AllowAudioStreamCopy.IsSet() {
		return true
	}

	return false
}

// SetAllowAudioStreamCopy gets a reference to the given NullableBool and assigns it to the AllowAudioStreamCopy field.
func (o *PlaybackInfoDto) SetAllowAudioStreamCopy(v bool) {
	o.AllowAudioStreamCopy.Set(&v)
}
// SetAllowAudioStreamCopyNil sets the value for AllowAudioStreamCopy to be an explicit nil
func (o *PlaybackInfoDto) SetAllowAudioStreamCopyNil() {
	o.AllowAudioStreamCopy.Set(nil)
}

// UnsetAllowAudioStreamCopy ensures that no value is present for AllowAudioStreamCopy, not even an explicit nil
func (o *PlaybackInfoDto) UnsetAllowAudioStreamCopy() {
	o.AllowAudioStreamCopy.Unset()
}

// GetAutoOpenLiveStream returns the AutoOpenLiveStream field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoDto) GetAutoOpenLiveStream() bool {
	if o == nil || IsNil(o.AutoOpenLiveStream.Get()) {
		var ret bool
		return ret
	}
	return *o.AutoOpenLiveStream.Get()
}

// GetAutoOpenLiveStreamOk returns a tuple with the AutoOpenLiveStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoDto) GetAutoOpenLiveStreamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoOpenLiveStream.Get(), o.AutoOpenLiveStream.IsSet()
}

// HasAutoOpenLiveStream returns a boolean if a field has been set.
func (o *PlaybackInfoDto) HasAutoOpenLiveStream() bool {
	if o != nil && o.AutoOpenLiveStream.IsSet() {
		return true
	}

	return false
}

// SetAutoOpenLiveStream gets a reference to the given NullableBool and assigns it to the AutoOpenLiveStream field.
func (o *PlaybackInfoDto) SetAutoOpenLiveStream(v bool) {
	o.AutoOpenLiveStream.Set(&v)
}
// SetAutoOpenLiveStreamNil sets the value for AutoOpenLiveStream to be an explicit nil
func (o *PlaybackInfoDto) SetAutoOpenLiveStreamNil() {
	o.AutoOpenLiveStream.Set(nil)
}

// UnsetAutoOpenLiveStream ensures that no value is present for AutoOpenLiveStream, not even an explicit nil
func (o *PlaybackInfoDto) UnsetAutoOpenLiveStream() {
	o.AutoOpenLiveStream.Unset()
}

func (o PlaybackInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaybackInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["UserId"] = o.UserId.Get()
	}
	if o.MaxStreamingBitrate.IsSet() {
		toSerialize["MaxStreamingBitrate"] = o.MaxStreamingBitrate.Get()
	}
	if o.StartTimeTicks.IsSet() {
		toSerialize["StartTimeTicks"] = o.StartTimeTicks.Get()
	}
	if o.AudioStreamIndex.IsSet() {
		toSerialize["AudioStreamIndex"] = o.AudioStreamIndex.Get()
	}
	if o.SubtitleStreamIndex.IsSet() {
		toSerialize["SubtitleStreamIndex"] = o.SubtitleStreamIndex.Get()
	}
	if o.MaxAudioChannels.IsSet() {
		toSerialize["MaxAudioChannels"] = o.MaxAudioChannels.Get()
	}
	if o.MediaSourceId.IsSet() {
		toSerialize["MediaSourceId"] = o.MediaSourceId.Get()
	}
	if o.LiveStreamId.IsSet() {
		toSerialize["LiveStreamId"] = o.LiveStreamId.Get()
	}
	if o.DeviceProfile.IsSet() {
		toSerialize["DeviceProfile"] = o.DeviceProfile.Get()
	}
	if o.EnableDirectPlay.IsSet() {
		toSerialize["EnableDirectPlay"] = o.EnableDirectPlay.Get()
	}
	if o.EnableDirectStream.IsSet() {
		toSerialize["EnableDirectStream"] = o.EnableDirectStream.Get()
	}
	if o.EnableTranscoding.IsSet() {
		toSerialize["EnableTranscoding"] = o.EnableTranscoding.Get()
	}
	if o.AllowVideoStreamCopy.IsSet() {
		toSerialize["AllowVideoStreamCopy"] = o.AllowVideoStreamCopy.Get()
	}
	if o.AllowAudioStreamCopy.IsSet() {
		toSerialize["AllowAudioStreamCopy"] = o.AllowAudioStreamCopy.Get()
	}
	if o.AutoOpenLiveStream.IsSet() {
		toSerialize["AutoOpenLiveStream"] = o.AutoOpenLiveStream.Get()
	}
	return toSerialize, nil
}

type NullablePlaybackInfoDto struct {
	value *PlaybackInfoDto
	isSet bool
}

func (v NullablePlaybackInfoDto) Get() *PlaybackInfoDto {
	return v.value
}

func (v *NullablePlaybackInfoDto) Set(val *PlaybackInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaybackInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaybackInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaybackInfoDto(val *PlaybackInfoDto) *NullablePlaybackInfoDto {
	return &NullablePlaybackInfoDto{value: val, isSet: true}
}

func (v NullablePlaybackInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaybackInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


