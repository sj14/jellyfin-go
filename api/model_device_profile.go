/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeviceProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceProfile{}

// DeviceProfile A MediaBrowser.Model.Dlna.DeviceProfile represents a set of metadata which determines which content a certain device is able to play.  <br />  Specifically, it defines the supported <see cref=\"P:MediaBrowser.Model.Dlna.DeviceProfile.ContainerProfiles\">containers</see> and  <see cref=\"P:MediaBrowser.Model.Dlna.DeviceProfile.CodecProfiles\">codecs</see> (video and/or audio, including codec profiles and levels)  the device is able to direct play (without transcoding or remuxing),  as well as which <see cref=\"P:MediaBrowser.Model.Dlna.DeviceProfile.TranscodingProfiles\">containers/codecs to transcode to</see> in case it isn't.
type DeviceProfile struct {
	// Gets or sets the name of this device profile. User profiles must have a unique name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the unique internal identifier.
	Id NullableString `json:"Id,omitempty"`
	// Gets or sets the maximum allowed bitrate for all streamed content.
	MaxStreamingBitrate NullableInt32 `json:"MaxStreamingBitrate,omitempty"`
	// Gets or sets the maximum allowed bitrate for statically streamed content (= direct played files).
	MaxStaticBitrate NullableInt32 `json:"MaxStaticBitrate,omitempty"`
	// Gets or sets the maximum allowed bitrate for transcoded music streams.
	MusicStreamingTranscodingBitrate NullableInt32 `json:"MusicStreamingTranscodingBitrate,omitempty"`
	// Gets or sets the maximum allowed bitrate for statically streamed (= direct played) music files.
	MaxStaticMusicBitrate NullableInt32 `json:"MaxStaticMusicBitrate,omitempty"`
	// Gets or sets the direct play profiles.
	DirectPlayProfiles []DirectPlayProfile `json:"DirectPlayProfiles,omitempty"`
	// Gets or sets the transcoding profiles.
	TranscodingProfiles []TranscodingProfile `json:"TranscodingProfiles,omitempty"`
	// Gets or sets the container profiles. Failing to meet these optional conditions causes transcoding to occur.
	ContainerProfiles []ContainerProfile `json:"ContainerProfiles,omitempty"`
	// Gets or sets the codec profiles.
	CodecProfiles []CodecProfile `json:"CodecProfiles,omitempty"`
	// Gets or sets the subtitle profiles.
	SubtitleProfiles []SubtitleProfile `json:"SubtitleProfiles,omitempty"`
}

// NewDeviceProfile instantiates a new DeviceProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceProfile() *DeviceProfile {
	this := DeviceProfile{}
	return &this
}

// NewDeviceProfileWithDefaults instantiates a new DeviceProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceProfileWithDefaults() *DeviceProfile {
	this := DeviceProfile{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceProfile) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DeviceProfile) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DeviceProfile) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DeviceProfile) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DeviceProfile) UnsetName() {
	o.Name.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceProfile) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceProfile) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DeviceProfile) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DeviceProfile) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DeviceProfile) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DeviceProfile) UnsetId() {
	o.Id.Unset()
}

// GetMaxStreamingBitrate returns the MaxStreamingBitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceProfile) GetMaxStreamingBitrate() int32 {
	if o == nil || IsNil(o.MaxStreamingBitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxStreamingBitrate.Get()
}

// GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceProfile) GetMaxStreamingBitrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStreamingBitrate.Get(), o.MaxStreamingBitrate.IsSet()
}

// HasMaxStreamingBitrate returns a boolean if a field has been set.
func (o *DeviceProfile) HasMaxStreamingBitrate() bool {
	if o != nil && o.MaxStreamingBitrate.IsSet() {
		return true
	}

	return false
}

// SetMaxStreamingBitrate gets a reference to the given NullableInt32 and assigns it to the MaxStreamingBitrate field.
func (o *DeviceProfile) SetMaxStreamingBitrate(v int32) {
	o.MaxStreamingBitrate.Set(&v)
}
// SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil
func (o *DeviceProfile) SetMaxStreamingBitrateNil() {
	o.MaxStreamingBitrate.Set(nil)
}

// UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
func (o *DeviceProfile) UnsetMaxStreamingBitrate() {
	o.MaxStreamingBitrate.Unset()
}

// GetMaxStaticBitrate returns the MaxStaticBitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceProfile) GetMaxStaticBitrate() int32 {
	if o == nil || IsNil(o.MaxStaticBitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxStaticBitrate.Get()
}

// GetMaxStaticBitrateOk returns a tuple with the MaxStaticBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceProfile) GetMaxStaticBitrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStaticBitrate.Get(), o.MaxStaticBitrate.IsSet()
}

// HasMaxStaticBitrate returns a boolean if a field has been set.
func (o *DeviceProfile) HasMaxStaticBitrate() bool {
	if o != nil && o.MaxStaticBitrate.IsSet() {
		return true
	}

	return false
}

// SetMaxStaticBitrate gets a reference to the given NullableInt32 and assigns it to the MaxStaticBitrate field.
func (o *DeviceProfile) SetMaxStaticBitrate(v int32) {
	o.MaxStaticBitrate.Set(&v)
}
// SetMaxStaticBitrateNil sets the value for MaxStaticBitrate to be an explicit nil
func (o *DeviceProfile) SetMaxStaticBitrateNil() {
	o.MaxStaticBitrate.Set(nil)
}

// UnsetMaxStaticBitrate ensures that no value is present for MaxStaticBitrate, not even an explicit nil
func (o *DeviceProfile) UnsetMaxStaticBitrate() {
	o.MaxStaticBitrate.Unset()
}

// GetMusicStreamingTranscodingBitrate returns the MusicStreamingTranscodingBitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceProfile) GetMusicStreamingTranscodingBitrate() int32 {
	if o == nil || IsNil(o.MusicStreamingTranscodingBitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.MusicStreamingTranscodingBitrate.Get()
}

// GetMusicStreamingTranscodingBitrateOk returns a tuple with the MusicStreamingTranscodingBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceProfile) GetMusicStreamingTranscodingBitrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MusicStreamingTranscodingBitrate.Get(), o.MusicStreamingTranscodingBitrate.IsSet()
}

// HasMusicStreamingTranscodingBitrate returns a boolean if a field has been set.
func (o *DeviceProfile) HasMusicStreamingTranscodingBitrate() bool {
	if o != nil && o.MusicStreamingTranscodingBitrate.IsSet() {
		return true
	}

	return false
}

// SetMusicStreamingTranscodingBitrate gets a reference to the given NullableInt32 and assigns it to the MusicStreamingTranscodingBitrate field.
func (o *DeviceProfile) SetMusicStreamingTranscodingBitrate(v int32) {
	o.MusicStreamingTranscodingBitrate.Set(&v)
}
// SetMusicStreamingTranscodingBitrateNil sets the value for MusicStreamingTranscodingBitrate to be an explicit nil
func (o *DeviceProfile) SetMusicStreamingTranscodingBitrateNil() {
	o.MusicStreamingTranscodingBitrate.Set(nil)
}

// UnsetMusicStreamingTranscodingBitrate ensures that no value is present for MusicStreamingTranscodingBitrate, not even an explicit nil
func (o *DeviceProfile) UnsetMusicStreamingTranscodingBitrate() {
	o.MusicStreamingTranscodingBitrate.Unset()
}

// GetMaxStaticMusicBitrate returns the MaxStaticMusicBitrate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceProfile) GetMaxStaticMusicBitrate() int32 {
	if o == nil || IsNil(o.MaxStaticMusicBitrate.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxStaticMusicBitrate.Get()
}

// GetMaxStaticMusicBitrateOk returns a tuple with the MaxStaticMusicBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceProfile) GetMaxStaticMusicBitrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStaticMusicBitrate.Get(), o.MaxStaticMusicBitrate.IsSet()
}

// HasMaxStaticMusicBitrate returns a boolean if a field has been set.
func (o *DeviceProfile) HasMaxStaticMusicBitrate() bool {
	if o != nil && o.MaxStaticMusicBitrate.IsSet() {
		return true
	}

	return false
}

// SetMaxStaticMusicBitrate gets a reference to the given NullableInt32 and assigns it to the MaxStaticMusicBitrate field.
func (o *DeviceProfile) SetMaxStaticMusicBitrate(v int32) {
	o.MaxStaticMusicBitrate.Set(&v)
}
// SetMaxStaticMusicBitrateNil sets the value for MaxStaticMusicBitrate to be an explicit nil
func (o *DeviceProfile) SetMaxStaticMusicBitrateNil() {
	o.MaxStaticMusicBitrate.Set(nil)
}

// UnsetMaxStaticMusicBitrate ensures that no value is present for MaxStaticMusicBitrate, not even an explicit nil
func (o *DeviceProfile) UnsetMaxStaticMusicBitrate() {
	o.MaxStaticMusicBitrate.Unset()
}

// GetDirectPlayProfiles returns the DirectPlayProfiles field value if set, zero value otherwise.
func (o *DeviceProfile) GetDirectPlayProfiles() []DirectPlayProfile {
	if o == nil || IsNil(o.DirectPlayProfiles) {
		var ret []DirectPlayProfile
		return ret
	}
	return o.DirectPlayProfiles
}

// GetDirectPlayProfilesOk returns a tuple with the DirectPlayProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetDirectPlayProfilesOk() ([]DirectPlayProfile, bool) {
	if o == nil || IsNil(o.DirectPlayProfiles) {
		return nil, false
	}
	return o.DirectPlayProfiles, true
}

// HasDirectPlayProfiles returns a boolean if a field has been set.
func (o *DeviceProfile) HasDirectPlayProfiles() bool {
	if o != nil && !IsNil(o.DirectPlayProfiles) {
		return true
	}

	return false
}

// SetDirectPlayProfiles gets a reference to the given []DirectPlayProfile and assigns it to the DirectPlayProfiles field.
func (o *DeviceProfile) SetDirectPlayProfiles(v []DirectPlayProfile) {
	o.DirectPlayProfiles = v
}

// GetTranscodingProfiles returns the TranscodingProfiles field value if set, zero value otherwise.
func (o *DeviceProfile) GetTranscodingProfiles() []TranscodingProfile {
	if o == nil || IsNil(o.TranscodingProfiles) {
		var ret []TranscodingProfile
		return ret
	}
	return o.TranscodingProfiles
}

// GetTranscodingProfilesOk returns a tuple with the TranscodingProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetTranscodingProfilesOk() ([]TranscodingProfile, bool) {
	if o == nil || IsNil(o.TranscodingProfiles) {
		return nil, false
	}
	return o.TranscodingProfiles, true
}

// HasTranscodingProfiles returns a boolean if a field has been set.
func (o *DeviceProfile) HasTranscodingProfiles() bool {
	if o != nil && !IsNil(o.TranscodingProfiles) {
		return true
	}

	return false
}

// SetTranscodingProfiles gets a reference to the given []TranscodingProfile and assigns it to the TranscodingProfiles field.
func (o *DeviceProfile) SetTranscodingProfiles(v []TranscodingProfile) {
	o.TranscodingProfiles = v
}

// GetContainerProfiles returns the ContainerProfiles field value if set, zero value otherwise.
func (o *DeviceProfile) GetContainerProfiles() []ContainerProfile {
	if o == nil || IsNil(o.ContainerProfiles) {
		var ret []ContainerProfile
		return ret
	}
	return o.ContainerProfiles
}

// GetContainerProfilesOk returns a tuple with the ContainerProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetContainerProfilesOk() ([]ContainerProfile, bool) {
	if o == nil || IsNil(o.ContainerProfiles) {
		return nil, false
	}
	return o.ContainerProfiles, true
}

// HasContainerProfiles returns a boolean if a field has been set.
func (o *DeviceProfile) HasContainerProfiles() bool {
	if o != nil && !IsNil(o.ContainerProfiles) {
		return true
	}

	return false
}

// SetContainerProfiles gets a reference to the given []ContainerProfile and assigns it to the ContainerProfiles field.
func (o *DeviceProfile) SetContainerProfiles(v []ContainerProfile) {
	o.ContainerProfiles = v
}

// GetCodecProfiles returns the CodecProfiles field value if set, zero value otherwise.
func (o *DeviceProfile) GetCodecProfiles() []CodecProfile {
	if o == nil || IsNil(o.CodecProfiles) {
		var ret []CodecProfile
		return ret
	}
	return o.CodecProfiles
}

// GetCodecProfilesOk returns a tuple with the CodecProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetCodecProfilesOk() ([]CodecProfile, bool) {
	if o == nil || IsNil(o.CodecProfiles) {
		return nil, false
	}
	return o.CodecProfiles, true
}

// HasCodecProfiles returns a boolean if a field has been set.
func (o *DeviceProfile) HasCodecProfiles() bool {
	if o != nil && !IsNil(o.CodecProfiles) {
		return true
	}

	return false
}

// SetCodecProfiles gets a reference to the given []CodecProfile and assigns it to the CodecProfiles field.
func (o *DeviceProfile) SetCodecProfiles(v []CodecProfile) {
	o.CodecProfiles = v
}

// GetSubtitleProfiles returns the SubtitleProfiles field value if set, zero value otherwise.
func (o *DeviceProfile) GetSubtitleProfiles() []SubtitleProfile {
	if o == nil || IsNil(o.SubtitleProfiles) {
		var ret []SubtitleProfile
		return ret
	}
	return o.SubtitleProfiles
}

// GetSubtitleProfilesOk returns a tuple with the SubtitleProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceProfile) GetSubtitleProfilesOk() ([]SubtitleProfile, bool) {
	if o == nil || IsNil(o.SubtitleProfiles) {
		return nil, false
	}
	return o.SubtitleProfiles, true
}

// HasSubtitleProfiles returns a boolean if a field has been set.
func (o *DeviceProfile) HasSubtitleProfiles() bool {
	if o != nil && !IsNil(o.SubtitleProfiles) {
		return true
	}

	return false
}

// SetSubtitleProfiles gets a reference to the given []SubtitleProfile and assigns it to the SubtitleProfiles field.
func (o *DeviceProfile) SetSubtitleProfiles(v []SubtitleProfile) {
	o.SubtitleProfiles = v
}

func (o DeviceProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.MaxStreamingBitrate.IsSet() {
		toSerialize["MaxStreamingBitrate"] = o.MaxStreamingBitrate.Get()
	}
	if o.MaxStaticBitrate.IsSet() {
		toSerialize["MaxStaticBitrate"] = o.MaxStaticBitrate.Get()
	}
	if o.MusicStreamingTranscodingBitrate.IsSet() {
		toSerialize["MusicStreamingTranscodingBitrate"] = o.MusicStreamingTranscodingBitrate.Get()
	}
	if o.MaxStaticMusicBitrate.IsSet() {
		toSerialize["MaxStaticMusicBitrate"] = o.MaxStaticMusicBitrate.Get()
	}
	if !IsNil(o.DirectPlayProfiles) {
		toSerialize["DirectPlayProfiles"] = o.DirectPlayProfiles
	}
	if !IsNil(o.TranscodingProfiles) {
		toSerialize["TranscodingProfiles"] = o.TranscodingProfiles
	}
	if !IsNil(o.ContainerProfiles) {
		toSerialize["ContainerProfiles"] = o.ContainerProfiles
	}
	if !IsNil(o.CodecProfiles) {
		toSerialize["CodecProfiles"] = o.CodecProfiles
	}
	if !IsNil(o.SubtitleProfiles) {
		toSerialize["SubtitleProfiles"] = o.SubtitleProfiles
	}
	return toSerialize, nil
}

type NullableDeviceProfile struct {
	value *DeviceProfile
	isSet bool
}

func (v NullableDeviceProfile) Get() *DeviceProfile {
	return v.value
}

func (v *NullableDeviceProfile) Set(val *DeviceProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceProfile(val *DeviceProfile) *NullableDeviceProfile {
	return &NullableDeviceProfile{value: val, isSet: true}
}

func (v NullableDeviceProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


