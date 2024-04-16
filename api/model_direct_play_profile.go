/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.8.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DirectPlayProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectPlayProfile{}

// DirectPlayProfile struct for DirectPlayProfile
type DirectPlayProfile struct {
	Container NullableString `json:"Container,omitempty"`
	AudioCodec NullableString `json:"AudioCodec,omitempty"`
	VideoCodec NullableString `json:"VideoCodec,omitempty"`
	Type *DlnaProfileType `json:"Type,omitempty"`
}

// NewDirectPlayProfile instantiates a new DirectPlayProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectPlayProfile() *DirectPlayProfile {
	this := DirectPlayProfile{}
	return &this
}

// NewDirectPlayProfileWithDefaults instantiates a new DirectPlayProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectPlayProfileWithDefaults() *DirectPlayProfile {
	this := DirectPlayProfile{}
	return &this
}

// GetContainer returns the Container field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectPlayProfile) GetContainer() string {
	if o == nil || IsNil(o.Container.Get()) {
		var ret string
		return ret
	}
	return *o.Container.Get()
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectPlayProfile) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Container.Get(), o.Container.IsSet()
}

// HasContainer returns a boolean if a field has been set.
func (o *DirectPlayProfile) HasContainer() bool {
	if o != nil && o.Container.IsSet() {
		return true
	}

	return false
}

// SetContainer gets a reference to the given NullableString and assigns it to the Container field.
func (o *DirectPlayProfile) SetContainer(v string) {
	o.Container.Set(&v)
}
// SetContainerNil sets the value for Container to be an explicit nil
func (o *DirectPlayProfile) SetContainerNil() {
	o.Container.Set(nil)
}

// UnsetContainer ensures that no value is present for Container, not even an explicit nil
func (o *DirectPlayProfile) UnsetContainer() {
	o.Container.Unset()
}

// GetAudioCodec returns the AudioCodec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectPlayProfile) GetAudioCodec() string {
	if o == nil || IsNil(o.AudioCodec.Get()) {
		var ret string
		return ret
	}
	return *o.AudioCodec.Get()
}

// GetAudioCodecOk returns a tuple with the AudioCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectPlayProfile) GetAudioCodecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AudioCodec.Get(), o.AudioCodec.IsSet()
}

// HasAudioCodec returns a boolean if a field has been set.
func (o *DirectPlayProfile) HasAudioCodec() bool {
	if o != nil && o.AudioCodec.IsSet() {
		return true
	}

	return false
}

// SetAudioCodec gets a reference to the given NullableString and assigns it to the AudioCodec field.
func (o *DirectPlayProfile) SetAudioCodec(v string) {
	o.AudioCodec.Set(&v)
}
// SetAudioCodecNil sets the value for AudioCodec to be an explicit nil
func (o *DirectPlayProfile) SetAudioCodecNil() {
	o.AudioCodec.Set(nil)
}

// UnsetAudioCodec ensures that no value is present for AudioCodec, not even an explicit nil
func (o *DirectPlayProfile) UnsetAudioCodec() {
	o.AudioCodec.Unset()
}

// GetVideoCodec returns the VideoCodec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DirectPlayProfile) GetVideoCodec() string {
	if o == nil || IsNil(o.VideoCodec.Get()) {
		var ret string
		return ret
	}
	return *o.VideoCodec.Get()
}

// GetVideoCodecOk returns a tuple with the VideoCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DirectPlayProfile) GetVideoCodecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VideoCodec.Get(), o.VideoCodec.IsSet()
}

// HasVideoCodec returns a boolean if a field has been set.
func (o *DirectPlayProfile) HasVideoCodec() bool {
	if o != nil && o.VideoCodec.IsSet() {
		return true
	}

	return false
}

// SetVideoCodec gets a reference to the given NullableString and assigns it to the VideoCodec field.
func (o *DirectPlayProfile) SetVideoCodec(v string) {
	o.VideoCodec.Set(&v)
}
// SetVideoCodecNil sets the value for VideoCodec to be an explicit nil
func (o *DirectPlayProfile) SetVideoCodecNil() {
	o.VideoCodec.Set(nil)
}

// UnsetVideoCodec ensures that no value is present for VideoCodec, not even an explicit nil
func (o *DirectPlayProfile) UnsetVideoCodec() {
	o.VideoCodec.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DirectPlayProfile) GetType() DlnaProfileType {
	if o == nil || IsNil(o.Type) {
		var ret DlnaProfileType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectPlayProfile) GetTypeOk() (*DlnaProfileType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DirectPlayProfile) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DlnaProfileType and assigns it to the Type field.
func (o *DirectPlayProfile) SetType(v DlnaProfileType) {
	o.Type = &v
}

func (o DirectPlayProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectPlayProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Container.IsSet() {
		toSerialize["Container"] = o.Container.Get()
	}
	if o.AudioCodec.IsSet() {
		toSerialize["AudioCodec"] = o.AudioCodec.Get()
	}
	if o.VideoCodec.IsSet() {
		toSerialize["VideoCodec"] = o.VideoCodec.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDirectPlayProfile struct {
	value *DirectPlayProfile
	isSet bool
}

func (v NullableDirectPlayProfile) Get() *DirectPlayProfile {
	return v.value
}

func (v *NullableDirectPlayProfile) Set(val *DirectPlayProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectPlayProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectPlayProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectPlayProfile(val *DirectPlayProfile) *NullableDirectPlayProfile {
	return &NullableDirectPlayProfile{value: val, isSet: true}
}

func (v NullableDirectPlayProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectPlayProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


