/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TranscodingProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TranscodingProfile{}

// TranscodingProfile struct for TranscodingProfile
type TranscodingProfile struct {
	Container *string `json:"Container,omitempty"`
	Type *DlnaProfileType `json:"Type,omitempty"`
	VideoCodec *string `json:"VideoCodec,omitempty"`
	AudioCodec *string `json:"AudioCodec,omitempty"`
	// Media streaming protocol.  Lowercase for backwards compatibility.
	Protocol *MediaStreamProtocol `json:"Protocol,omitempty"`
	EstimateContentLength *bool `json:"EstimateContentLength,omitempty"`
	EnableMpegtsM2TsMode *bool `json:"EnableMpegtsM2TsMode,omitempty"`
	TranscodeSeekInfo *TranscodeSeekInfo `json:"TranscodeSeekInfo,omitempty"`
	CopyTimestamps *bool `json:"CopyTimestamps,omitempty"`
	Context *EncodingContext `json:"Context,omitempty"`
	EnableSubtitlesInManifest *bool `json:"EnableSubtitlesInManifest,omitempty"`
	MaxAudioChannels NullableString `json:"MaxAudioChannels,omitempty"`
	MinSegments *int32 `json:"MinSegments,omitempty"`
	SegmentLength *int32 `json:"SegmentLength,omitempty"`
	BreakOnNonKeyFrames *bool `json:"BreakOnNonKeyFrames,omitempty"`
	Conditions []ProfileCondition `json:"Conditions,omitempty"`
}

// NewTranscodingProfile instantiates a new TranscodingProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranscodingProfile() *TranscodingProfile {
	this := TranscodingProfile{}
	var estimateContentLength bool = false
	this.EstimateContentLength = &estimateContentLength
	var enableMpegtsM2TsMode bool = false
	this.EnableMpegtsM2TsMode = &enableMpegtsM2TsMode
	var transcodeSeekInfo TranscodeSeekInfo = TRANSCODESEEKINFO_AUTO
	this.TranscodeSeekInfo = &transcodeSeekInfo
	var copyTimestamps bool = false
	this.CopyTimestamps = &copyTimestamps
	var context EncodingContext = ENCODINGCONTEXT_STREAMING
	this.Context = &context
	var enableSubtitlesInManifest bool = false
	this.EnableSubtitlesInManifest = &enableSubtitlesInManifest
	var minSegments int32 = 0
	this.MinSegments = &minSegments
	var segmentLength int32 = 0
	this.SegmentLength = &segmentLength
	var breakOnNonKeyFrames bool = false
	this.BreakOnNonKeyFrames = &breakOnNonKeyFrames
	return &this
}

// NewTranscodingProfileWithDefaults instantiates a new TranscodingProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranscodingProfileWithDefaults() *TranscodingProfile {
	this := TranscodingProfile{}
	var estimateContentLength bool = false
	this.EstimateContentLength = &estimateContentLength
	var enableMpegtsM2TsMode bool = false
	this.EnableMpegtsM2TsMode = &enableMpegtsM2TsMode
	var transcodeSeekInfo TranscodeSeekInfo = TRANSCODESEEKINFO_AUTO
	this.TranscodeSeekInfo = &transcodeSeekInfo
	var copyTimestamps bool = false
	this.CopyTimestamps = &copyTimestamps
	var context EncodingContext = ENCODINGCONTEXT_STREAMING
	this.Context = &context
	var enableSubtitlesInManifest bool = false
	this.EnableSubtitlesInManifest = &enableSubtitlesInManifest
	var minSegments int32 = 0
	this.MinSegments = &minSegments
	var segmentLength int32 = 0
	this.SegmentLength = &segmentLength
	var breakOnNonKeyFrames bool = false
	this.BreakOnNonKeyFrames = &breakOnNonKeyFrames
	return &this
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *TranscodingProfile) GetContainer() string {
	if o == nil || IsNil(o.Container) {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetContainerOk() (*string, bool) {
	if o == nil || IsNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *TranscodingProfile) HasContainer() bool {
	if o != nil && !IsNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *TranscodingProfile) SetContainer(v string) {
	o.Container = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TranscodingProfile) GetType() DlnaProfileType {
	if o == nil || IsNil(o.Type) {
		var ret DlnaProfileType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetTypeOk() (*DlnaProfileType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TranscodingProfile) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DlnaProfileType and assigns it to the Type field.
func (o *TranscodingProfile) SetType(v DlnaProfileType) {
	o.Type = &v
}

// GetVideoCodec returns the VideoCodec field value if set, zero value otherwise.
func (o *TranscodingProfile) GetVideoCodec() string {
	if o == nil || IsNil(o.VideoCodec) {
		var ret string
		return ret
	}
	return *o.VideoCodec
}

// GetVideoCodecOk returns a tuple with the VideoCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetVideoCodecOk() (*string, bool) {
	if o == nil || IsNil(o.VideoCodec) {
		return nil, false
	}
	return o.VideoCodec, true
}

// HasVideoCodec returns a boolean if a field has been set.
func (o *TranscodingProfile) HasVideoCodec() bool {
	if o != nil && !IsNil(o.VideoCodec) {
		return true
	}

	return false
}

// SetVideoCodec gets a reference to the given string and assigns it to the VideoCodec field.
func (o *TranscodingProfile) SetVideoCodec(v string) {
	o.VideoCodec = &v
}

// GetAudioCodec returns the AudioCodec field value if set, zero value otherwise.
func (o *TranscodingProfile) GetAudioCodec() string {
	if o == nil || IsNil(o.AudioCodec) {
		var ret string
		return ret
	}
	return *o.AudioCodec
}

// GetAudioCodecOk returns a tuple with the AudioCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetAudioCodecOk() (*string, bool) {
	if o == nil || IsNil(o.AudioCodec) {
		return nil, false
	}
	return o.AudioCodec, true
}

// HasAudioCodec returns a boolean if a field has been set.
func (o *TranscodingProfile) HasAudioCodec() bool {
	if o != nil && !IsNil(o.AudioCodec) {
		return true
	}

	return false
}

// SetAudioCodec gets a reference to the given string and assigns it to the AudioCodec field.
func (o *TranscodingProfile) SetAudioCodec(v string) {
	o.AudioCodec = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *TranscodingProfile) GetProtocol() MediaStreamProtocol {
	if o == nil || IsNil(o.Protocol) {
		var ret MediaStreamProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetProtocolOk() (*MediaStreamProtocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *TranscodingProfile) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given MediaStreamProtocol and assigns it to the Protocol field.
func (o *TranscodingProfile) SetProtocol(v MediaStreamProtocol) {
	o.Protocol = &v
}

// GetEstimateContentLength returns the EstimateContentLength field value if set, zero value otherwise.
func (o *TranscodingProfile) GetEstimateContentLength() bool {
	if o == nil || IsNil(o.EstimateContentLength) {
		var ret bool
		return ret
	}
	return *o.EstimateContentLength
}

// GetEstimateContentLengthOk returns a tuple with the EstimateContentLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetEstimateContentLengthOk() (*bool, bool) {
	if o == nil || IsNil(o.EstimateContentLength) {
		return nil, false
	}
	return o.EstimateContentLength, true
}

// HasEstimateContentLength returns a boolean if a field has been set.
func (o *TranscodingProfile) HasEstimateContentLength() bool {
	if o != nil && !IsNil(o.EstimateContentLength) {
		return true
	}

	return false
}

// SetEstimateContentLength gets a reference to the given bool and assigns it to the EstimateContentLength field.
func (o *TranscodingProfile) SetEstimateContentLength(v bool) {
	o.EstimateContentLength = &v
}

// GetEnableMpegtsM2TsMode returns the EnableMpegtsM2TsMode field value if set, zero value otherwise.
func (o *TranscodingProfile) GetEnableMpegtsM2TsMode() bool {
	if o == nil || IsNil(o.EnableMpegtsM2TsMode) {
		var ret bool
		return ret
	}
	return *o.EnableMpegtsM2TsMode
}

// GetEnableMpegtsM2TsModeOk returns a tuple with the EnableMpegtsM2TsMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetEnableMpegtsM2TsModeOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableMpegtsM2TsMode) {
		return nil, false
	}
	return o.EnableMpegtsM2TsMode, true
}

// HasEnableMpegtsM2TsMode returns a boolean if a field has been set.
func (o *TranscodingProfile) HasEnableMpegtsM2TsMode() bool {
	if o != nil && !IsNil(o.EnableMpegtsM2TsMode) {
		return true
	}

	return false
}

// SetEnableMpegtsM2TsMode gets a reference to the given bool and assigns it to the EnableMpegtsM2TsMode field.
func (o *TranscodingProfile) SetEnableMpegtsM2TsMode(v bool) {
	o.EnableMpegtsM2TsMode = &v
}

// GetTranscodeSeekInfo returns the TranscodeSeekInfo field value if set, zero value otherwise.
func (o *TranscodingProfile) GetTranscodeSeekInfo() TranscodeSeekInfo {
	if o == nil || IsNil(o.TranscodeSeekInfo) {
		var ret TranscodeSeekInfo
		return ret
	}
	return *o.TranscodeSeekInfo
}

// GetTranscodeSeekInfoOk returns a tuple with the TranscodeSeekInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetTranscodeSeekInfoOk() (*TranscodeSeekInfo, bool) {
	if o == nil || IsNil(o.TranscodeSeekInfo) {
		return nil, false
	}
	return o.TranscodeSeekInfo, true
}

// HasTranscodeSeekInfo returns a boolean if a field has been set.
func (o *TranscodingProfile) HasTranscodeSeekInfo() bool {
	if o != nil && !IsNil(o.TranscodeSeekInfo) {
		return true
	}

	return false
}

// SetTranscodeSeekInfo gets a reference to the given TranscodeSeekInfo and assigns it to the TranscodeSeekInfo field.
func (o *TranscodingProfile) SetTranscodeSeekInfo(v TranscodeSeekInfo) {
	o.TranscodeSeekInfo = &v
}

// GetCopyTimestamps returns the CopyTimestamps field value if set, zero value otherwise.
func (o *TranscodingProfile) GetCopyTimestamps() bool {
	if o == nil || IsNil(o.CopyTimestamps) {
		var ret bool
		return ret
	}
	return *o.CopyTimestamps
}

// GetCopyTimestampsOk returns a tuple with the CopyTimestamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetCopyTimestampsOk() (*bool, bool) {
	if o == nil || IsNil(o.CopyTimestamps) {
		return nil, false
	}
	return o.CopyTimestamps, true
}

// HasCopyTimestamps returns a boolean if a field has been set.
func (o *TranscodingProfile) HasCopyTimestamps() bool {
	if o != nil && !IsNil(o.CopyTimestamps) {
		return true
	}

	return false
}

// SetCopyTimestamps gets a reference to the given bool and assigns it to the CopyTimestamps field.
func (o *TranscodingProfile) SetCopyTimestamps(v bool) {
	o.CopyTimestamps = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TranscodingProfile) GetContext() EncodingContext {
	if o == nil || IsNil(o.Context) {
		var ret EncodingContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetContextOk() (*EncodingContext, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TranscodingProfile) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given EncodingContext and assigns it to the Context field.
func (o *TranscodingProfile) SetContext(v EncodingContext) {
	o.Context = &v
}

// GetEnableSubtitlesInManifest returns the EnableSubtitlesInManifest field value if set, zero value otherwise.
func (o *TranscodingProfile) GetEnableSubtitlesInManifest() bool {
	if o == nil || IsNil(o.EnableSubtitlesInManifest) {
		var ret bool
		return ret
	}
	return *o.EnableSubtitlesInManifest
}

// GetEnableSubtitlesInManifestOk returns a tuple with the EnableSubtitlesInManifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetEnableSubtitlesInManifestOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableSubtitlesInManifest) {
		return nil, false
	}
	return o.EnableSubtitlesInManifest, true
}

// HasEnableSubtitlesInManifest returns a boolean if a field has been set.
func (o *TranscodingProfile) HasEnableSubtitlesInManifest() bool {
	if o != nil && !IsNil(o.EnableSubtitlesInManifest) {
		return true
	}

	return false
}

// SetEnableSubtitlesInManifest gets a reference to the given bool and assigns it to the EnableSubtitlesInManifest field.
func (o *TranscodingProfile) SetEnableSubtitlesInManifest(v bool) {
	o.EnableSubtitlesInManifest = &v
}

// GetMaxAudioChannels returns the MaxAudioChannels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TranscodingProfile) GetMaxAudioChannels() string {
	if o == nil || IsNil(o.MaxAudioChannels.Get()) {
		var ret string
		return ret
	}
	return *o.MaxAudioChannels.Get()
}

// GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TranscodingProfile) GetMaxAudioChannelsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAudioChannels.Get(), o.MaxAudioChannels.IsSet()
}

// HasMaxAudioChannels returns a boolean if a field has been set.
func (o *TranscodingProfile) HasMaxAudioChannels() bool {
	if o != nil && o.MaxAudioChannels.IsSet() {
		return true
	}

	return false
}

// SetMaxAudioChannels gets a reference to the given NullableString and assigns it to the MaxAudioChannels field.
func (o *TranscodingProfile) SetMaxAudioChannels(v string) {
	o.MaxAudioChannels.Set(&v)
}
// SetMaxAudioChannelsNil sets the value for MaxAudioChannels to be an explicit nil
func (o *TranscodingProfile) SetMaxAudioChannelsNil() {
	o.MaxAudioChannels.Set(nil)
}

// UnsetMaxAudioChannels ensures that no value is present for MaxAudioChannels, not even an explicit nil
func (o *TranscodingProfile) UnsetMaxAudioChannels() {
	o.MaxAudioChannels.Unset()
}

// GetMinSegments returns the MinSegments field value if set, zero value otherwise.
func (o *TranscodingProfile) GetMinSegments() int32 {
	if o == nil || IsNil(o.MinSegments) {
		var ret int32
		return ret
	}
	return *o.MinSegments
}

// GetMinSegmentsOk returns a tuple with the MinSegments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetMinSegmentsOk() (*int32, bool) {
	if o == nil || IsNil(o.MinSegments) {
		return nil, false
	}
	return o.MinSegments, true
}

// HasMinSegments returns a boolean if a field has been set.
func (o *TranscodingProfile) HasMinSegments() bool {
	if o != nil && !IsNil(o.MinSegments) {
		return true
	}

	return false
}

// SetMinSegments gets a reference to the given int32 and assigns it to the MinSegments field.
func (o *TranscodingProfile) SetMinSegments(v int32) {
	o.MinSegments = &v
}

// GetSegmentLength returns the SegmentLength field value if set, zero value otherwise.
func (o *TranscodingProfile) GetSegmentLength() int32 {
	if o == nil || IsNil(o.SegmentLength) {
		var ret int32
		return ret
	}
	return *o.SegmentLength
}

// GetSegmentLengthOk returns a tuple with the SegmentLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetSegmentLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.SegmentLength) {
		return nil, false
	}
	return o.SegmentLength, true
}

// HasSegmentLength returns a boolean if a field has been set.
func (o *TranscodingProfile) HasSegmentLength() bool {
	if o != nil && !IsNil(o.SegmentLength) {
		return true
	}

	return false
}

// SetSegmentLength gets a reference to the given int32 and assigns it to the SegmentLength field.
func (o *TranscodingProfile) SetSegmentLength(v int32) {
	o.SegmentLength = &v
}

// GetBreakOnNonKeyFrames returns the BreakOnNonKeyFrames field value if set, zero value otherwise.
func (o *TranscodingProfile) GetBreakOnNonKeyFrames() bool {
	if o == nil || IsNil(o.BreakOnNonKeyFrames) {
		var ret bool
		return ret
	}
	return *o.BreakOnNonKeyFrames
}

// GetBreakOnNonKeyFramesOk returns a tuple with the BreakOnNonKeyFrames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetBreakOnNonKeyFramesOk() (*bool, bool) {
	if o == nil || IsNil(o.BreakOnNonKeyFrames) {
		return nil, false
	}
	return o.BreakOnNonKeyFrames, true
}

// HasBreakOnNonKeyFrames returns a boolean if a field has been set.
func (o *TranscodingProfile) HasBreakOnNonKeyFrames() bool {
	if o != nil && !IsNil(o.BreakOnNonKeyFrames) {
		return true
	}

	return false
}

// SetBreakOnNonKeyFrames gets a reference to the given bool and assigns it to the BreakOnNonKeyFrames field.
func (o *TranscodingProfile) SetBreakOnNonKeyFrames(v bool) {
	o.BreakOnNonKeyFrames = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *TranscodingProfile) GetConditions() []ProfileCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []ProfileCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscodingProfile) GetConditionsOk() ([]ProfileCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *TranscodingProfile) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ProfileCondition and assigns it to the Conditions field.
func (o *TranscodingProfile) SetConditions(v []ProfileCondition) {
	o.Conditions = v
}

func (o TranscodingProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TranscodingProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Container) {
		toSerialize["Container"] = o.Container
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.VideoCodec) {
		toSerialize["VideoCodec"] = o.VideoCodec
	}
	if !IsNil(o.AudioCodec) {
		toSerialize["AudioCodec"] = o.AudioCodec
	}
	if !IsNil(o.Protocol) {
		toSerialize["Protocol"] = o.Protocol
	}
	if !IsNil(o.EstimateContentLength) {
		toSerialize["EstimateContentLength"] = o.EstimateContentLength
	}
	if !IsNil(o.EnableMpegtsM2TsMode) {
		toSerialize["EnableMpegtsM2TsMode"] = o.EnableMpegtsM2TsMode
	}
	if !IsNil(o.TranscodeSeekInfo) {
		toSerialize["TranscodeSeekInfo"] = o.TranscodeSeekInfo
	}
	if !IsNil(o.CopyTimestamps) {
		toSerialize["CopyTimestamps"] = o.CopyTimestamps
	}
	if !IsNil(o.Context) {
		toSerialize["Context"] = o.Context
	}
	if !IsNil(o.EnableSubtitlesInManifest) {
		toSerialize["EnableSubtitlesInManifest"] = o.EnableSubtitlesInManifest
	}
	if o.MaxAudioChannels.IsSet() {
		toSerialize["MaxAudioChannels"] = o.MaxAudioChannels.Get()
	}
	if !IsNil(o.MinSegments) {
		toSerialize["MinSegments"] = o.MinSegments
	}
	if !IsNil(o.SegmentLength) {
		toSerialize["SegmentLength"] = o.SegmentLength
	}
	if !IsNil(o.BreakOnNonKeyFrames) {
		toSerialize["BreakOnNonKeyFrames"] = o.BreakOnNonKeyFrames
	}
	if !IsNil(o.Conditions) {
		toSerialize["Conditions"] = o.Conditions
	}
	return toSerialize, nil
}

type NullableTranscodingProfile struct {
	value *TranscodingProfile
	isSet bool
}

func (v NullableTranscodingProfile) Get() *TranscodingProfile {
	return v.value
}

func (v *NullableTranscodingProfile) Set(val *TranscodingProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableTranscodingProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableTranscodingProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranscodingProfile(val *TranscodingProfile) *NullableTranscodingProfile {
	return &NullableTranscodingProfile{value: val, isSet: true}
}

func (v NullableTranscodingProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranscodingProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


