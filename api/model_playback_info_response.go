/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PlaybackInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaybackInfoResponse{}

// PlaybackInfoResponse Class PlaybackInfoResponse.
type PlaybackInfoResponse struct {
	// Gets or sets the media sources.
	MediaSources []MediaSourceInfo `json:"MediaSources,omitempty"`
	// Gets or sets the play session identifier.
	PlaySessionId NullableString `json:"PlaySessionId,omitempty"`
	// Gets or sets the error code.
	ErrorCode NullablePlaybackErrorCode `json:"ErrorCode,omitempty"`
}

// NewPlaybackInfoResponse instantiates a new PlaybackInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaybackInfoResponse() *PlaybackInfoResponse {
	this := PlaybackInfoResponse{}
	return &this
}

// NewPlaybackInfoResponseWithDefaults instantiates a new PlaybackInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaybackInfoResponseWithDefaults() *PlaybackInfoResponse {
	this := PlaybackInfoResponse{}
	return &this
}

// GetMediaSources returns the MediaSources field value if set, zero value otherwise.
func (o *PlaybackInfoResponse) GetMediaSources() []MediaSourceInfo {
	if o == nil || IsNil(o.MediaSources) {
		var ret []MediaSourceInfo
		return ret
	}
	return o.MediaSources
}

// GetMediaSourcesOk returns a tuple with the MediaSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackInfoResponse) GetMediaSourcesOk() ([]MediaSourceInfo, bool) {
	if o == nil || IsNil(o.MediaSources) {
		return nil, false
	}
	return o.MediaSources, true
}

// HasMediaSources returns a boolean if a field has been set.
func (o *PlaybackInfoResponse) HasMediaSources() bool {
	if o != nil && !IsNil(o.MediaSources) {
		return true
	}

	return false
}

// SetMediaSources gets a reference to the given []MediaSourceInfo and assigns it to the MediaSources field.
func (o *PlaybackInfoResponse) SetMediaSources(v []MediaSourceInfo) {
	o.MediaSources = v
}

// GetPlaySessionId returns the PlaySessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoResponse) GetPlaySessionId() string {
	if o == nil || IsNil(o.PlaySessionId.Get()) {
		var ret string
		return ret
	}
	return *o.PlaySessionId.Get()
}

// GetPlaySessionIdOk returns a tuple with the PlaySessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoResponse) GetPlaySessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaySessionId.Get(), o.PlaySessionId.IsSet()
}

// HasPlaySessionId returns a boolean if a field has been set.
func (o *PlaybackInfoResponse) HasPlaySessionId() bool {
	if o != nil && o.PlaySessionId.IsSet() {
		return true
	}

	return false
}

// SetPlaySessionId gets a reference to the given NullableString and assigns it to the PlaySessionId field.
func (o *PlaybackInfoResponse) SetPlaySessionId(v string) {
	o.PlaySessionId.Set(&v)
}
// SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil
func (o *PlaybackInfoResponse) SetPlaySessionIdNil() {
	o.PlaySessionId.Set(nil)
}

// UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
func (o *PlaybackInfoResponse) UnsetPlaySessionId() {
	o.PlaySessionId.Unset()
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaybackInfoResponse) GetErrorCode() PlaybackErrorCode {
	if o == nil || IsNil(o.ErrorCode.Get()) {
		var ret PlaybackErrorCode
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaybackInfoResponse) GetErrorCodeOk() (*PlaybackErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *PlaybackInfoResponse) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullablePlaybackErrorCode and assigns it to the ErrorCode field.
func (o *PlaybackInfoResponse) SetErrorCode(v PlaybackErrorCode) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *PlaybackInfoResponse) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *PlaybackInfoResponse) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

func (o PlaybackInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaybackInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MediaSources) {
		toSerialize["MediaSources"] = o.MediaSources
	}
	if o.PlaySessionId.IsSet() {
		toSerialize["PlaySessionId"] = o.PlaySessionId.Get()
	}
	if o.ErrorCode.IsSet() {
		toSerialize["ErrorCode"] = o.ErrorCode.Get()
	}
	return toSerialize, nil
}

type NullablePlaybackInfoResponse struct {
	value *PlaybackInfoResponse
	isSet bool
}

func (v NullablePlaybackInfoResponse) Get() *PlaybackInfoResponse {
	return v.value
}

func (v *NullablePlaybackInfoResponse) Set(val *PlaybackInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaybackInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaybackInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaybackInfoResponse(val *PlaybackInfoResponse) *NullablePlaybackInfoResponse {
	return &NullablePlaybackInfoResponse{value: val, isSet: true}
}

func (v NullablePlaybackInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaybackInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


