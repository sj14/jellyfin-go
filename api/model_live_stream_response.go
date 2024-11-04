/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the LiveStreamResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveStreamResponse{}

// LiveStreamResponse struct for LiveStreamResponse
type LiveStreamResponse struct {
	MediaSource *MediaSourceInfo `json:"MediaSource,omitempty"`
}

// NewLiveStreamResponse instantiates a new LiveStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveStreamResponse() *LiveStreamResponse {
	this := LiveStreamResponse{}
	return &this
}

// NewLiveStreamResponseWithDefaults instantiates a new LiveStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveStreamResponseWithDefaults() *LiveStreamResponse {
	this := LiveStreamResponse{}
	return &this
}

// GetMediaSource returns the MediaSource field value if set, zero value otherwise.
func (o *LiveStreamResponse) GetMediaSource() MediaSourceInfo {
	if o == nil || IsNil(o.MediaSource) {
		var ret MediaSourceInfo
		return ret
	}
	return *o.MediaSource
}

// GetMediaSourceOk returns a tuple with the MediaSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveStreamResponse) GetMediaSourceOk() (*MediaSourceInfo, bool) {
	if o == nil || IsNil(o.MediaSource) {
		return nil, false
	}
	return o.MediaSource, true
}

// HasMediaSource returns a boolean if a field has been set.
func (o *LiveStreamResponse) HasMediaSource() bool {
	if o != nil && !IsNil(o.MediaSource) {
		return true
	}

	return false
}

// SetMediaSource gets a reference to the given MediaSourceInfo and assigns it to the MediaSource field.
func (o *LiveStreamResponse) SetMediaSource(v MediaSourceInfo) {
	o.MediaSource = &v
}

func (o LiveStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MediaSource) {
		toSerialize["MediaSource"] = o.MediaSource
	}
	return toSerialize, nil
}

type NullableLiveStreamResponse struct {
	value *LiveStreamResponse
	isSet bool
}

func (v NullableLiveStreamResponse) Get() *LiveStreamResponse {
	return v.value
}

func (v *NullableLiveStreamResponse) Set(val *LiveStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveStreamResponse(val *LiveStreamResponse) *NullableLiveStreamResponse {
	return &NullableLiveStreamResponse{value: val, isSet: true}
}

func (v NullableLiveStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


