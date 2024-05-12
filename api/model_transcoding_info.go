/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TranscodingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TranscodingInfo{}

// TranscodingInfo struct for TranscodingInfo
type TranscodingInfo struct {
	AudioCodec NullableString `json:"AudioCodec,omitempty"`
	VideoCodec NullableString `json:"VideoCodec,omitempty"`
	Container NullableString `json:"Container,omitempty"`
	IsVideoDirect *bool `json:"IsVideoDirect,omitempty"`
	IsAudioDirect *bool `json:"IsAudioDirect,omitempty"`
	Bitrate NullableInt32 `json:"Bitrate,omitempty"`
	Framerate NullableFloat32 `json:"Framerate,omitempty"`
	CompletionPercentage NullableFloat64 `json:"CompletionPercentage,omitempty"`
	Width NullableInt32 `json:"Width,omitempty"`
	Height NullableInt32 `json:"Height,omitempty"`
	AudioChannels NullableInt32 `json:"AudioChannels,omitempty"`
	HardwareAccelerationType NullableHardwareEncodingType `json:"HardwareAccelerationType,omitempty"`
	TranscodeReasons []TranscodeReason `json:"TranscodeReasons,omitempty"`
}

// NewTranscodingInfo instantiates a new TranscodingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranscodingInfo() *TranscodingInfo {
	this := TranscodingInfo{}
	return &this
}

// NewTranscodingInfoWithDefaults instantiates a new TranscodingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranscodingInfoWithDefaults() *TranscodingInfo {
	this := TranscodingInfo{}
	return &this
}

// GetAudioCodec returns the AudioCodec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranscodingInfo) GetAudioCodec() string {
	if o == nil || IsNil(o.AudioCodec.Get()) {
		var ret string
		return ret
	}
	return *o.AudioCodec.Get()
}

// GetAudioCodecOk returns a tuple with the AudioCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranscodingInfo) GetAudioCodecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioCodec.Get(), o.AudioCodec.IsSet()
}

// HasAudioCodec returns a boolean if a field has been set.
func (o *TranscodingInfo) HasAudioCodec() bool {
	if o != nil && o.AudioCodec.IsSet() {
		return true
	}

	return false
}

// SetAudioCodec gets a reference to the given NullableString and assigns it to the AudioCodec field.
func (o *TranscodingInfo) SetAudioCodec(v string) {
	o.AudioCodec.Set(&v)
}
// SetAudioCodecNil sets the value for AudioCodec to be an explicit nil
func (o *TranscodingInfo) SetAudioCodecNil() {
	o.AudioCodec.Set(nil)
}

// UnsetAudioCodec ensures that no value is present for AudioCodec, not even an explicit nil
func (o *TranscodingInfo) UnsetAudioCodec() {
	o.AudioCodec.Unset()
}

// GetVideoCodec returns the VideoCodec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranscodingInfo) GetVideoCodec() string {
	if o == nil || IsNil(o.VideoCodec.Get()) {
		var ret string
		return ret
	}
	return *o.VideoCodec.Get()
}

// GetVideoCodecOk returns a tuple with the VideoCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranscodingInfo) GetVideoCodecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoCodec.Get(), o.VideoCodec.IsSet()
}

// HasVideoCodec returns a boolean if a field has been set.
func (o *TranscodingInfo) HasVideoCodec() bool {
	if o != nil && o.VideoCodec.IsSet() {
		return true
	}

	return false
}

// SetVideoCodec gets a reference to the given NullableString and assigns it to the VideoCodec field.
func (o *TranscodingInfo) SetVideoCodec(v string) {
	o.VideoCodec.Set(&v)
}
// SetVideoCodecNil sets the value for VideoCodec to be an explicit nil
func (o *TranscodingInfo) SetVideoCodecNil() {
	o.VideoCodec.Set(nil)
}

// UnsetVideoCodec ensures that no value is present for VideoCodec, not even an explicit nil
func (o *TranscodingInfo) UnsetVideoCodec() {
	o.VideoCodec.Unset()
}

// GetContainer returns the Container field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranscodingInfo) GetContainer() string {
	if o == nil || IsNil(o.Container.Get()) {
		var ret string
		return ret
	}
	return *o.Container.Get()
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranscodingInfo) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Container.Get(), o.Container.IsSet()
}

// HasContainer returns a boolean if a field has been set.
func (o *TranscodingInfo) HasContainer() bool {
	if o != nil && o.Container.IsSet() {
		return true
	}

	return false
}

// SetContainer gets a reference to the given NullableString and assigns it to the Container field.
func (o *TranscodingInfo) SetContainer(v string) {
	o.Container.Set(&v)
}
// SetContainerNil sets the value for Container to be an explicit nil
func (o *TranscodingInfo) SetContainerNil() {
	o.Container.Set(nil)
}

// UnsetContainer ensures that no value is present for Container, not even an explicit nil
func (o *TranscodingInfo) UnsetContainer() {
	o.Container.Unset()
}

// GetIsVideoDirect returns the IsVideoDirect field value if set, zero value otherwise.
func (o *TranscodingInfo) GetIsVideoDirect() bool {
	if o == nil || IsNil(o.IsVideoDirect) {
		var ret bool
		return ret
	}
	return *o.IsVideoDirect
}

// GetIsVideoDirectOk returns a tuple with the IsVideoDirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingInfo) GetIsVideoDirectOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVideoDirect) {
		return nil, false
	}
	return o.IsVideoDirect, true
}

// HasIsVideoDirect returns a boolean if a field has been set.
func (o *TranscodingInfo) HasIsVideoDirect() bool {
	if o != nil && !IsNil(o.IsVideoDirect) {
		return true
	}

	return false
}

// SetIsVideoDirect gets a reference to the given bool and assigns it to the IsVideoDirect field.
func (o *TranscodingInfo) SetIsVideoDirect(v bool) {
	o.IsVideoDirect = &v
}

// GetIsAudioDirect returns the IsAudioDirect field value if set, zero value otherwise.
func (o *TranscodingInfo) GetIsAudioDirect() bool {
	if o == nil || IsNil(o.IsAudioDirect) {
		var ret bool
		return ret
	}
	return *o.IsAudioDirect
}

// GetIsAudioDirectOk returns a tuple with the IsAudioDirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingInfo) GetIsAudioDirectOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAudioDirect) {
		return nil, false
	}
	return o.IsAudioDirect, true
}

// HasIsAudioDirect returns a boolean if a field has been set.
func (o *TranscodingInfo) HasIsAudioDirect() bool {
	if o != nil && !IsNil(o.IsAudioDirect) {
		return true
	}

	return false
}

// SetIsAudioDirect gets a reference to the given bool and assigns it to the IsAudioDirect field.
func (o *TranscodingInfo) SetIsAudioDirect(v bool) {
	o.IsAudioDirect = &v
}

// GetBitrate returns the Bitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranscodingInfo) GetBitrate() int32 {
	if o == nil || IsNil(o.Bitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.Bitrate.Get()
}

// GetBitrateOk returns a tuple with the Bitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranscodingInfo) GetBitrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bitrate.Get(), o.Bitrate.IsSet()
}

// HasBitrate returns a boolean if a field has been set.
func (o *TranscodingInfo) HasBitrate() bool {
	if o != nil && o.Bitrate.IsSet() {
		return true
	}

	return false
}

// SetBitrate gets a reference to the given NullableInt32 and assigns it to the Bitrate field.
func (o *TranscodingInfo) SetBitrate(v int32) {
	o.Bitrate.Set(&v)
}
// SetBitrateNil sets the value for Bitrate to be an explicit nil
func (o *TranscodingInfo) SetBitrateNil() {
	o.Bitrate.Set(nil)
}

// UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
func (o *TranscodingInfo) UnsetBitrate() {
	o.Bitrate.Unset()
}

// GetFramerate returns the Framerate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranscodingInfo) GetFramerate() float32 {
	if o == nil || IsNil(o.Framerate.Get()) {
		var ret float32
		return ret
	}
	return *o.Framerate.Get()
}

// GetFramerateOk returns a tuple with the Framerate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranscodingInfo) GetFramerateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Framerate.Get(), o.Framerate.IsSet()
}

// HasFramerate returns a boolean if a field has been set.
func (o *TranscodingInfo) HasFramerate() bool {
	if o != nil && o.Framerate.IsSet() {
		return true
	}

	return false
}

// SetFramerate gets a reference to the given NullableFloat32 and assigns it to the Framerate field.
func (o *TranscodingInfo) SetFramerate(v float32) {
	o.Framerate.Set(&v)
}
// SetFramerateNil sets the value for Framerate to be an explicit nil
func (o *TranscodingInfo) SetFramerateNil() {
	o.Framerate.Set(nil)
}

// UnsetFramerate ensures that no value is present for Framerate, not even an explicit nil
func (o *TranscodingInfo) UnsetFramerate() {
	o.Framerate.Unset()
}

// GetCompletionPercentage returns the CompletionPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranscodingInfo) GetCompletionPercentage() float64 {
	if o == nil || IsNil(o.CompletionPercentage.Get()) {
		var ret float64
		return ret
	}
	return *o.CompletionPercentage.Get()
}

// GetCompletionPercentageOk returns a tuple with the CompletionPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranscodingInfo) GetCompletionPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletionPercentage.Get(), o.CompletionPercentage.IsSet()
}

// HasCompletionPercentage returns a boolean if a field has been set.
func (o *TranscodingInfo) HasCompletionPercentage() bool {
	if o != nil && o.CompletionPercentage.IsSet() {
		return true
	}

	return false
}

// SetCompletionPercentage gets a reference to the given NullableFloat64 and assigns it to the CompletionPercentage field.
func (o *TranscodingInfo) SetCompletionPercentage(v float64) {
	o.CompletionPercentage.Set(&v)
}
// SetCompletionPercentageNil sets the value for CompletionPercentage to be an explicit nil
func (o *TranscodingInfo) SetCompletionPercentageNil() {
	o.CompletionPercentage.Set(nil)
}

// UnsetCompletionPercentage ensures that no value is present for CompletionPercentage, not even an explicit nil
func (o *TranscodingInfo) UnsetCompletionPercentage() {
	o.CompletionPercentage.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranscodingInfo) GetWidth() int32 {
	if o == nil || IsNil(o.Width.Get()) {
		var ret int32
		return ret
	}
	return *o.Width.Get()
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranscodingInfo) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Width.Get(), o.Width.IsSet()
}

// HasWidth returns a boolean if a field has been set.
func (o *TranscodingInfo) HasWidth() bool {
	if o != nil && o.Width.IsSet() {
		return true
	}

	return false
}

// SetWidth gets a reference to the given NullableInt32 and assigns it to the Width field.
func (o *TranscodingInfo) SetWidth(v int32) {
	o.Width.Set(&v)
}
// SetWidthNil sets the value for Width to be an explicit nil
func (o *TranscodingInfo) SetWidthNil() {
	o.Width.Set(nil)
}

// UnsetWidth ensures that no value is present for Width, not even an explicit nil
func (o *TranscodingInfo) UnsetWidth() {
	o.Width.Unset()
}

// GetHeight returns the Height field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranscodingInfo) GetHeight() int32 {
	if o == nil || IsNil(o.Height.Get()) {
		var ret int32
		return ret
	}
	return *o.Height.Get()
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranscodingInfo) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Height.Get(), o.Height.IsSet()
}

// HasHeight returns a boolean if a field has been set.
func (o *TranscodingInfo) HasHeight() bool {
	if o != nil && o.Height.IsSet() {
		return true
	}

	return false
}

// SetHeight gets a reference to the given NullableInt32 and assigns it to the Height field.
func (o *TranscodingInfo) SetHeight(v int32) {
	o.Height.Set(&v)
}
// SetHeightNil sets the value for Height to be an explicit nil
func (o *TranscodingInfo) SetHeightNil() {
	o.Height.Set(nil)
}

// UnsetHeight ensures that no value is present for Height, not even an explicit nil
func (o *TranscodingInfo) UnsetHeight() {
	o.Height.Unset()
}

// GetAudioChannels returns the AudioChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranscodingInfo) GetAudioChannels() int32 {
	if o == nil || IsNil(o.AudioChannels.Get()) {
		var ret int32
		return ret
	}
	return *o.AudioChannels.Get()
}

// GetAudioChannelsOk returns a tuple with the AudioChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranscodingInfo) GetAudioChannelsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioChannels.Get(), o.AudioChannels.IsSet()
}

// HasAudioChannels returns a boolean if a field has been set.
func (o *TranscodingInfo) HasAudioChannels() bool {
	if o != nil && o.AudioChannels.IsSet() {
		return true
	}

	return false
}

// SetAudioChannels gets a reference to the given NullableInt32 and assigns it to the AudioChannels field.
func (o *TranscodingInfo) SetAudioChannels(v int32) {
	o.AudioChannels.Set(&v)
}
// SetAudioChannelsNil sets the value for AudioChannels to be an explicit nil
func (o *TranscodingInfo) SetAudioChannelsNil() {
	o.AudioChannels.Set(nil)
}

// UnsetAudioChannels ensures that no value is present for AudioChannels, not even an explicit nil
func (o *TranscodingInfo) UnsetAudioChannels() {
	o.AudioChannels.Unset()
}

// GetHardwareAccelerationType returns the HardwareAccelerationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranscodingInfo) GetHardwareAccelerationType() HardwareEncodingType {
	if o == nil || IsNil(o.HardwareAccelerationType.Get()) {
		var ret HardwareEncodingType
		return ret
	}
	return *o.HardwareAccelerationType.Get()
}

// GetHardwareAccelerationTypeOk returns a tuple with the HardwareAccelerationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranscodingInfo) GetHardwareAccelerationTypeOk() (*HardwareEncodingType, bool) {
	if o == nil {
		return nil, false
	}
	return o.HardwareAccelerationType.Get(), o.HardwareAccelerationType.IsSet()
}

// HasHardwareAccelerationType returns a boolean if a field has been set.
func (o *TranscodingInfo) HasHardwareAccelerationType() bool {
	if o != nil && o.HardwareAccelerationType.IsSet() {
		return true
	}

	return false
}

// SetHardwareAccelerationType gets a reference to the given NullableHardwareEncodingType and assigns it to the HardwareAccelerationType field.
func (o *TranscodingInfo) SetHardwareAccelerationType(v HardwareEncodingType) {
	o.HardwareAccelerationType.Set(&v)
}
// SetHardwareAccelerationTypeNil sets the value for HardwareAccelerationType to be an explicit nil
func (o *TranscodingInfo) SetHardwareAccelerationTypeNil() {
	o.HardwareAccelerationType.Set(nil)
}

// UnsetHardwareAccelerationType ensures that no value is present for HardwareAccelerationType, not even an explicit nil
func (o *TranscodingInfo) UnsetHardwareAccelerationType() {
	o.HardwareAccelerationType.Unset()
}

// GetTranscodeReasons returns the TranscodeReasons field value if set, zero value otherwise.
func (o *TranscodingInfo) GetTranscodeReasons() []TranscodeReason {
	if o == nil || IsNil(o.TranscodeReasons) {
		var ret []TranscodeReason
		return ret
	}
	return o.TranscodeReasons
}

// GetTranscodeReasonsOk returns a tuple with the TranscodeReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingInfo) GetTranscodeReasonsOk() ([]TranscodeReason, bool) {
	if o == nil || IsNil(o.TranscodeReasons) {
		return nil, false
	}
	return o.TranscodeReasons, true
}

// HasTranscodeReasons returns a boolean if a field has been set.
func (o *TranscodingInfo) HasTranscodeReasons() bool {
	if o != nil && !IsNil(o.TranscodeReasons) {
		return true
	}

	return false
}

// SetTranscodeReasons gets a reference to the given []TranscodeReason and assigns it to the TranscodeReasons field.
func (o *TranscodingInfo) SetTranscodeReasons(v []TranscodeReason) {
	o.TranscodeReasons = v
}

func (o TranscodingInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TranscodingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AudioCodec.IsSet() {
		toSerialize["AudioCodec"] = o.AudioCodec.Get()
	}
	if o.VideoCodec.IsSet() {
		toSerialize["VideoCodec"] = o.VideoCodec.Get()
	}
	if o.Container.IsSet() {
		toSerialize["Container"] = o.Container.Get()
	}
	if !IsNil(o.IsVideoDirect) {
		toSerialize["IsVideoDirect"] = o.IsVideoDirect
	}
	if !IsNil(o.IsAudioDirect) {
		toSerialize["IsAudioDirect"] = o.IsAudioDirect
	}
	if o.Bitrate.IsSet() {
		toSerialize["Bitrate"] = o.Bitrate.Get()
	}
	if o.Framerate.IsSet() {
		toSerialize["Framerate"] = o.Framerate.Get()
	}
	if o.CompletionPercentage.IsSet() {
		toSerialize["CompletionPercentage"] = o.CompletionPercentage.Get()
	}
	if o.Width.IsSet() {
		toSerialize["Width"] = o.Width.Get()
	}
	if o.Height.IsSet() {
		toSerialize["Height"] = o.Height.Get()
	}
	if o.AudioChannels.IsSet() {
		toSerialize["AudioChannels"] = o.AudioChannels.Get()
	}
	if o.HardwareAccelerationType.IsSet() {
		toSerialize["HardwareAccelerationType"] = o.HardwareAccelerationType.Get()
	}
	if !IsNil(o.TranscodeReasons) {
		toSerialize["TranscodeReasons"] = o.TranscodeReasons
	}
	return toSerialize, nil
}

type NullableTranscodingInfo struct {
	value *TranscodingInfo
	isSet bool
}

func (v NullableTranscodingInfo) Get() *TranscodingInfo {
	return v.value
}

func (v *NullableTranscodingInfo) Set(val *TranscodingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTranscodingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTranscodingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranscodingInfo(val *TranscodingInfo) *NullableTranscodingInfo {
	return &NullableTranscodingInfo{value: val, isSet: true}
}

func (v NullableTranscodingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranscodingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


