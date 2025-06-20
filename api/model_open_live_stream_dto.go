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

// checks if the OpenLiveStreamDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpenLiveStreamDto{}

// OpenLiveStreamDto Open live stream dto.
type OpenLiveStreamDto struct {
	// Gets or sets the open token.
	OpenToken NullableString `json:"OpenToken,omitempty"`
	// Gets or sets the user id.
	UserId NullableString `json:"UserId,omitempty"`
	// Gets or sets the play session id.
	PlaySessionId NullableString `json:"PlaySessionId,omitempty"`
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
	// Gets or sets the item id.
	ItemId NullableString `json:"ItemId,omitempty"`
	// Gets or sets a value indicating whether to enable direct play.
	EnableDirectPlay NullableBool `json:"EnableDirectPlay,omitempty"`
	// Gets or sets a value indicating whether to enable direct stream.
	EnableDirectStream NullableBool `json:"EnableDirectStream,omitempty"`
	// Gets or sets a value indicating whether always burn in subtitles when transcoding.
	AlwaysBurnInSubtitleWhenTranscoding NullableBool `json:"AlwaysBurnInSubtitleWhenTranscoding,omitempty"`
	// Gets or sets the device profile.
	DeviceProfile NullableDeviceProfile `json:"DeviceProfile,omitempty"`
	// Gets or sets the device play protocols.
	DirectPlayProtocols []MediaProtocol `json:"DirectPlayProtocols,omitempty"`
}

// NewOpenLiveStreamDto instantiates a new OpenLiveStreamDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpenLiveStreamDto() *OpenLiveStreamDto {
	this := OpenLiveStreamDto{}
	return &this
}

// NewOpenLiveStreamDtoWithDefaults instantiates a new OpenLiveStreamDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpenLiveStreamDtoWithDefaults() *OpenLiveStreamDto {
	this := OpenLiveStreamDto{}
	return &this
}

// GetOpenToken returns the OpenToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenLiveStreamDto) GetOpenToken() string {
	if o == nil || IsNil(o.OpenToken.Get()) {
		var ret string
		return ret
	}
	return *o.OpenToken.Get()
}

// GetOpenTokenOk returns a tuple with the OpenToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenLiveStreamDto) GetOpenTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenToken.Get(), o.OpenToken.IsSet()
}

// HasOpenToken returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasOpenToken() bool {
	if o != nil && o.OpenToken.IsSet() {
		return true
	}

	return false
}

// SetOpenToken gets a reference to the given NullableString and assigns it to the OpenToken field.
func (o *OpenLiveStreamDto) SetOpenToken(v string) {
	o.OpenToken.Set(&v)
}
// SetOpenTokenNil sets the value for OpenToken to be an explicit nil
func (o *OpenLiveStreamDto) SetOpenTokenNil() {
	o.OpenToken.Set(nil)
}

// UnsetOpenToken ensures that no value is present for OpenToken, not even an explicit nil
func (o *OpenLiveStreamDto) UnsetOpenToken() {
	o.OpenToken.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenLiveStreamDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenLiveStreamDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *OpenLiveStreamDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *OpenLiveStreamDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *OpenLiveStreamDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetPlaySessionId returns the PlaySessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenLiveStreamDto) GetPlaySessionId() string {
	if o == nil || IsNil(o.PlaySessionId.Get()) {
		var ret string
		return ret
	}
	return *o.PlaySessionId.Get()
}

// GetPlaySessionIdOk returns a tuple with the PlaySessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenLiveStreamDto) GetPlaySessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaySessionId.Get(), o.PlaySessionId.IsSet()
}

// HasPlaySessionId returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasPlaySessionId() bool {
	if o != nil && o.PlaySessionId.IsSet() {
		return true
	}

	return false
}

// SetPlaySessionId gets a reference to the given NullableString and assigns it to the PlaySessionId field.
func (o *OpenLiveStreamDto) SetPlaySessionId(v string) {
	o.PlaySessionId.Set(&v)
}
// SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil
func (o *OpenLiveStreamDto) SetPlaySessionIdNil() {
	o.PlaySessionId.Set(nil)
}

// UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
func (o *OpenLiveStreamDto) UnsetPlaySessionId() {
	o.PlaySessionId.Unset()
}

// GetMaxStreamingBitrate returns the MaxStreamingBitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenLiveStreamDto) GetMaxStreamingBitrate() int32 {
	if o == nil || IsNil(o.MaxStreamingBitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxStreamingBitrate.Get()
}

// GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenLiveStreamDto) GetMaxStreamingBitrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStreamingBitrate.Get(), o.MaxStreamingBitrate.IsSet()
}

// HasMaxStreamingBitrate returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasMaxStreamingBitrate() bool {
	if o != nil && o.MaxStreamingBitrate.IsSet() {
		return true
	}

	return false
}

// SetMaxStreamingBitrate gets a reference to the given NullableInt32 and assigns it to the MaxStreamingBitrate field.
func (o *OpenLiveStreamDto) SetMaxStreamingBitrate(v int32) {
	o.MaxStreamingBitrate.Set(&v)
}
// SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil
func (o *OpenLiveStreamDto) SetMaxStreamingBitrateNil() {
	o.MaxStreamingBitrate.Set(nil)
}

// UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
func (o *OpenLiveStreamDto) UnsetMaxStreamingBitrate() {
	o.MaxStreamingBitrate.Unset()
}

// GetStartTimeTicks returns the StartTimeTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenLiveStreamDto) GetStartTimeTicks() int64 {
	if o == nil || IsNil(o.StartTimeTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.StartTimeTicks.Get()
}

// GetStartTimeTicksOk returns a tuple with the StartTimeTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenLiveStreamDto) GetStartTimeTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartTimeTicks.Get(), o.StartTimeTicks.IsSet()
}

// HasStartTimeTicks returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasStartTimeTicks() bool {
	if o != nil && o.StartTimeTicks.IsSet() {
		return true
	}

	return false
}

// SetStartTimeTicks gets a reference to the given NullableInt64 and assigns it to the StartTimeTicks field.
func (o *OpenLiveStreamDto) SetStartTimeTicks(v int64) {
	o.StartTimeTicks.Set(&v)
}
// SetStartTimeTicksNil sets the value for StartTimeTicks to be an explicit nil
func (o *OpenLiveStreamDto) SetStartTimeTicksNil() {
	o.StartTimeTicks.Set(nil)
}

// UnsetStartTimeTicks ensures that no value is present for StartTimeTicks, not even an explicit nil
func (o *OpenLiveStreamDto) UnsetStartTimeTicks() {
	o.StartTimeTicks.Unset()
}

// GetAudioStreamIndex returns the AudioStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenLiveStreamDto) GetAudioStreamIndex() int32 {
	if o == nil || IsNil(o.AudioStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.AudioStreamIndex.Get()
}

// GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenLiveStreamDto) GetAudioStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioStreamIndex.Get(), o.AudioStreamIndex.IsSet()
}

// HasAudioStreamIndex returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasAudioStreamIndex() bool {
	if o != nil && o.AudioStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetAudioStreamIndex gets a reference to the given NullableInt32 and assigns it to the AudioStreamIndex field.
func (o *OpenLiveStreamDto) SetAudioStreamIndex(v int32) {
	o.AudioStreamIndex.Set(&v)
}
// SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil
func (o *OpenLiveStreamDto) SetAudioStreamIndexNil() {
	o.AudioStreamIndex.Set(nil)
}

// UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
func (o *OpenLiveStreamDto) UnsetAudioStreamIndex() {
	o.AudioStreamIndex.Unset()
}

// GetSubtitleStreamIndex returns the SubtitleStreamIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenLiveStreamDto) GetSubtitleStreamIndex() int32 {
	if o == nil || IsNil(o.SubtitleStreamIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.SubtitleStreamIndex.Get()
}

// GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenLiveStreamDto) GetSubtitleStreamIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubtitleStreamIndex.Get(), o.SubtitleStreamIndex.IsSet()
}

// HasSubtitleStreamIndex returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasSubtitleStreamIndex() bool {
	if o != nil && o.SubtitleStreamIndex.IsSet() {
		return true
	}

	return false
}

// SetSubtitleStreamIndex gets a reference to the given NullableInt32 and assigns it to the SubtitleStreamIndex field.
func (o *OpenLiveStreamDto) SetSubtitleStreamIndex(v int32) {
	o.SubtitleStreamIndex.Set(&v)
}
// SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil
func (o *OpenLiveStreamDto) SetSubtitleStreamIndexNil() {
	o.SubtitleStreamIndex.Set(nil)
}

// UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
func (o *OpenLiveStreamDto) UnsetSubtitleStreamIndex() {
	o.SubtitleStreamIndex.Unset()
}

// GetMaxAudioChannels returns the MaxAudioChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenLiveStreamDto) GetMaxAudioChannels() int32 {
	if o == nil || IsNil(o.MaxAudioChannels.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxAudioChannels.Get()
}

// GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenLiveStreamDto) GetMaxAudioChannelsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAudioChannels.Get(), o.MaxAudioChannels.IsSet()
}

// HasMaxAudioChannels returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasMaxAudioChannels() bool {
	if o != nil && o.MaxAudioChannels.IsSet() {
		return true
	}

	return false
}

// SetMaxAudioChannels gets a reference to the given NullableInt32 and assigns it to the MaxAudioChannels field.
func (o *OpenLiveStreamDto) SetMaxAudioChannels(v int32) {
	o.MaxAudioChannels.Set(&v)
}
// SetMaxAudioChannelsNil sets the value for MaxAudioChannels to be an explicit nil
func (o *OpenLiveStreamDto) SetMaxAudioChannelsNil() {
	o.MaxAudioChannels.Set(nil)
}

// UnsetMaxAudioChannels ensures that no value is present for MaxAudioChannels, not even an explicit nil
func (o *OpenLiveStreamDto) UnsetMaxAudioChannels() {
	o.MaxAudioChannels.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenLiveStreamDto) GetItemId() string {
	if o == nil || IsNil(o.ItemId.Get()) {
		var ret string
		return ret
	}
	return *o.ItemId.Get()
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenLiveStreamDto) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemId.Get(), o.ItemId.IsSet()
}

// HasItemId returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasItemId() bool {
	if o != nil && o.ItemId.IsSet() {
		return true
	}

	return false
}

// SetItemId gets a reference to the given NullableString and assigns it to the ItemId field.
func (o *OpenLiveStreamDto) SetItemId(v string) {
	o.ItemId.Set(&v)
}
// SetItemIdNil sets the value for ItemId to be an explicit nil
func (o *OpenLiveStreamDto) SetItemIdNil() {
	o.ItemId.Set(nil)
}

// UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
func (o *OpenLiveStreamDto) UnsetItemId() {
	o.ItemId.Unset()
}

// GetEnableDirectPlay returns the EnableDirectPlay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenLiveStreamDto) GetEnableDirectPlay() bool {
	if o == nil || IsNil(o.EnableDirectPlay.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableDirectPlay.Get()
}

// GetEnableDirectPlayOk returns a tuple with the EnableDirectPlay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenLiveStreamDto) GetEnableDirectPlayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableDirectPlay.Get(), o.EnableDirectPlay.IsSet()
}

// HasEnableDirectPlay returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasEnableDirectPlay() bool {
	if o != nil && o.EnableDirectPlay.IsSet() {
		return true
	}

	return false
}

// SetEnableDirectPlay gets a reference to the given NullableBool and assigns it to the EnableDirectPlay field.
func (o *OpenLiveStreamDto) SetEnableDirectPlay(v bool) {
	o.EnableDirectPlay.Set(&v)
}
// SetEnableDirectPlayNil sets the value for EnableDirectPlay to be an explicit nil
func (o *OpenLiveStreamDto) SetEnableDirectPlayNil() {
	o.EnableDirectPlay.Set(nil)
}

// UnsetEnableDirectPlay ensures that no value is present for EnableDirectPlay, not even an explicit nil
func (o *OpenLiveStreamDto) UnsetEnableDirectPlay() {
	o.EnableDirectPlay.Unset()
}

// GetEnableDirectStream returns the EnableDirectStream field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenLiveStreamDto) GetEnableDirectStream() bool {
	if o == nil || IsNil(o.EnableDirectStream.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableDirectStream.Get()
}

// GetEnableDirectStreamOk returns a tuple with the EnableDirectStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenLiveStreamDto) GetEnableDirectStreamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableDirectStream.Get(), o.EnableDirectStream.IsSet()
}

// HasEnableDirectStream returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasEnableDirectStream() bool {
	if o != nil && o.EnableDirectStream.IsSet() {
		return true
	}

	return false
}

// SetEnableDirectStream gets a reference to the given NullableBool and assigns it to the EnableDirectStream field.
func (o *OpenLiveStreamDto) SetEnableDirectStream(v bool) {
	o.EnableDirectStream.Set(&v)
}
// SetEnableDirectStreamNil sets the value for EnableDirectStream to be an explicit nil
func (o *OpenLiveStreamDto) SetEnableDirectStreamNil() {
	o.EnableDirectStream.Set(nil)
}

// UnsetEnableDirectStream ensures that no value is present for EnableDirectStream, not even an explicit nil
func (o *OpenLiveStreamDto) UnsetEnableDirectStream() {
	o.EnableDirectStream.Unset()
}

// GetAlwaysBurnInSubtitleWhenTranscoding returns the AlwaysBurnInSubtitleWhenTranscoding field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenLiveStreamDto) GetAlwaysBurnInSubtitleWhenTranscoding() bool {
	if o == nil || IsNil(o.AlwaysBurnInSubtitleWhenTranscoding.Get()) {
		var ret bool
		return ret
	}
	return *o.AlwaysBurnInSubtitleWhenTranscoding.Get()
}

// GetAlwaysBurnInSubtitleWhenTranscodingOk returns a tuple with the AlwaysBurnInSubtitleWhenTranscoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenLiveStreamDto) GetAlwaysBurnInSubtitleWhenTranscodingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlwaysBurnInSubtitleWhenTranscoding.Get(), o.AlwaysBurnInSubtitleWhenTranscoding.IsSet()
}

// HasAlwaysBurnInSubtitleWhenTranscoding returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasAlwaysBurnInSubtitleWhenTranscoding() bool {
	if o != nil && o.AlwaysBurnInSubtitleWhenTranscoding.IsSet() {
		return true
	}

	return false
}

// SetAlwaysBurnInSubtitleWhenTranscoding gets a reference to the given NullableBool and assigns it to the AlwaysBurnInSubtitleWhenTranscoding field.
func (o *OpenLiveStreamDto) SetAlwaysBurnInSubtitleWhenTranscoding(v bool) {
	o.AlwaysBurnInSubtitleWhenTranscoding.Set(&v)
}
// SetAlwaysBurnInSubtitleWhenTranscodingNil sets the value for AlwaysBurnInSubtitleWhenTranscoding to be an explicit nil
func (o *OpenLiveStreamDto) SetAlwaysBurnInSubtitleWhenTranscodingNil() {
	o.AlwaysBurnInSubtitleWhenTranscoding.Set(nil)
}

// UnsetAlwaysBurnInSubtitleWhenTranscoding ensures that no value is present for AlwaysBurnInSubtitleWhenTranscoding, not even an explicit nil
func (o *OpenLiveStreamDto) UnsetAlwaysBurnInSubtitleWhenTranscoding() {
	o.AlwaysBurnInSubtitleWhenTranscoding.Unset()
}

// GetDeviceProfile returns the DeviceProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OpenLiveStreamDto) GetDeviceProfile() DeviceProfile {
	if o == nil || IsNil(o.DeviceProfile.Get()) {
		var ret DeviceProfile
		return ret
	}
	return *o.DeviceProfile.Get()
}

// GetDeviceProfileOk returns a tuple with the DeviceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpenLiveStreamDto) GetDeviceProfileOk() (*DeviceProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceProfile.Get(), o.DeviceProfile.IsSet()
}

// HasDeviceProfile returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasDeviceProfile() bool {
	if o != nil && o.DeviceProfile.IsSet() {
		return true
	}

	return false
}

// SetDeviceProfile gets a reference to the given NullableDeviceProfile and assigns it to the DeviceProfile field.
func (o *OpenLiveStreamDto) SetDeviceProfile(v DeviceProfile) {
	o.DeviceProfile.Set(&v)
}
// SetDeviceProfileNil sets the value for DeviceProfile to be an explicit nil
func (o *OpenLiveStreamDto) SetDeviceProfileNil() {
	o.DeviceProfile.Set(nil)
}

// UnsetDeviceProfile ensures that no value is present for DeviceProfile, not even an explicit nil
func (o *OpenLiveStreamDto) UnsetDeviceProfile() {
	o.DeviceProfile.Unset()
}

// GetDirectPlayProtocols returns the DirectPlayProtocols field value if set, zero value otherwise.
func (o *OpenLiveStreamDto) GetDirectPlayProtocols() []MediaProtocol {
	if o == nil || IsNil(o.DirectPlayProtocols) {
		var ret []MediaProtocol
		return ret
	}
	return o.DirectPlayProtocols
}

// GetDirectPlayProtocolsOk returns a tuple with the DirectPlayProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpenLiveStreamDto) GetDirectPlayProtocolsOk() ([]MediaProtocol, bool) {
	if o == nil || IsNil(o.DirectPlayProtocols) {
		return nil, false
	}
	return o.DirectPlayProtocols, true
}

// HasDirectPlayProtocols returns a boolean if a field has been set.
func (o *OpenLiveStreamDto) HasDirectPlayProtocols() bool {
	if o != nil && !IsNil(o.DirectPlayProtocols) {
		return true
	}

	return false
}

// SetDirectPlayProtocols gets a reference to the given []MediaProtocol and assigns it to the DirectPlayProtocols field.
func (o *OpenLiveStreamDto) SetDirectPlayProtocols(v []MediaProtocol) {
	o.DirectPlayProtocols = v
}

func (o OpenLiveStreamDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpenLiveStreamDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OpenToken.IsSet() {
		toSerialize["OpenToken"] = o.OpenToken.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["UserId"] = o.UserId.Get()
	}
	if o.PlaySessionId.IsSet() {
		toSerialize["PlaySessionId"] = o.PlaySessionId.Get()
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
	if o.ItemId.IsSet() {
		toSerialize["ItemId"] = o.ItemId.Get()
	}
	if o.EnableDirectPlay.IsSet() {
		toSerialize["EnableDirectPlay"] = o.EnableDirectPlay.Get()
	}
	if o.EnableDirectStream.IsSet() {
		toSerialize["EnableDirectStream"] = o.EnableDirectStream.Get()
	}
	if o.AlwaysBurnInSubtitleWhenTranscoding.IsSet() {
		toSerialize["AlwaysBurnInSubtitleWhenTranscoding"] = o.AlwaysBurnInSubtitleWhenTranscoding.Get()
	}
	if o.DeviceProfile.IsSet() {
		toSerialize["DeviceProfile"] = o.DeviceProfile.Get()
	}
	if !IsNil(o.DirectPlayProtocols) {
		toSerialize["DirectPlayProtocols"] = o.DirectPlayProtocols
	}
	return toSerialize, nil
}

type NullableOpenLiveStreamDto struct {
	value *OpenLiveStreamDto
	isSet bool
}

func (v NullableOpenLiveStreamDto) Get() *OpenLiveStreamDto {
	return v.value
}

func (v *NullableOpenLiveStreamDto) Set(val *OpenLiveStreamDto) {
	v.value = val
	v.isSet = true
}

func (v NullableOpenLiveStreamDto) IsSet() bool {
	return v.isSet
}

func (v *NullableOpenLiveStreamDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpenLiveStreamDto(val *OpenLiveStreamDto) *NullableOpenLiveStreamDto {
	return &NullableOpenLiveStreamDto{value: val, isSet: true}
}

func (v NullableOpenLiveStreamDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpenLiveStreamDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


