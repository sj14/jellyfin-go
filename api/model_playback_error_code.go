/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// PlaybackErrorCode the model 'PlaybackErrorCode'
type PlaybackErrorCode string

// List of PlaybackErrorCode
const (
	PLAYBACKERRORCODE_NOT_ALLOWED PlaybackErrorCode = "NotAllowed"
	PLAYBACKERRORCODE_NO_COMPATIBLE_STREAM PlaybackErrorCode = "NoCompatibleStream"
	PLAYBACKERRORCODE_RATE_LIMIT_EXCEEDED PlaybackErrorCode = "RateLimitExceeded"
)

// All allowed values of PlaybackErrorCode enum
var AllowedPlaybackErrorCodeEnumValues = []PlaybackErrorCode{
	"NotAllowed",
	"NoCompatibleStream",
	"RateLimitExceeded",
}

func (v *PlaybackErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaybackErrorCode(value)
	for _, existing := range AllowedPlaybackErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaybackErrorCode", value)
}

// NewPlaybackErrorCodeFromValue returns a pointer to a valid PlaybackErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaybackErrorCodeFromValue(v string) (*PlaybackErrorCode, error) {
	ev := PlaybackErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaybackErrorCode: valid values are %v", v, AllowedPlaybackErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaybackErrorCode) IsValid() bool {
	for _, existing := range AllowedPlaybackErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlaybackErrorCode value
func (v PlaybackErrorCode) Ptr() *PlaybackErrorCode {
	return &v
}

type NullablePlaybackErrorCode struct {
	value *PlaybackErrorCode
	isSet bool
}

func (v NullablePlaybackErrorCode) Get() *PlaybackErrorCode {
	return v.value
}

func (v *NullablePlaybackErrorCode) Set(val *PlaybackErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaybackErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaybackErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaybackErrorCode(val *PlaybackErrorCode) *NullablePlaybackErrorCode {
	return &NullablePlaybackErrorCode{value: val, isSet: true}
}

func (v NullablePlaybackErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaybackErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

