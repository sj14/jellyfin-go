/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// SubtitlePlaybackMode An enum representing a subtitle playback mode.
type SubtitlePlaybackMode string

// List of SubtitlePlaybackMode
const (
	SUBTITLEPLAYBACKMODE_DEFAULT SubtitlePlaybackMode = "Default"
	SUBTITLEPLAYBACKMODE_ALWAYS SubtitlePlaybackMode = "Always"
	SUBTITLEPLAYBACKMODE_ONLY_FORCED SubtitlePlaybackMode = "OnlyForced"
	SUBTITLEPLAYBACKMODE_NONE SubtitlePlaybackMode = "None"
	SUBTITLEPLAYBACKMODE_SMART SubtitlePlaybackMode = "Smart"
)

// All allowed values of SubtitlePlaybackMode enum
var AllowedSubtitlePlaybackModeEnumValues = []SubtitlePlaybackMode{
	"Default",
	"Always",
	"OnlyForced",
	"None",
	"Smart",
}

func (v *SubtitlePlaybackMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubtitlePlaybackMode(value)
	for _, existing := range AllowedSubtitlePlaybackModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubtitlePlaybackMode", value)
}

// NewSubtitlePlaybackModeFromValue returns a pointer to a valid SubtitlePlaybackMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubtitlePlaybackModeFromValue(v string) (*SubtitlePlaybackMode, error) {
	ev := SubtitlePlaybackMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubtitlePlaybackMode: valid values are %v", v, AllowedSubtitlePlaybackModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubtitlePlaybackMode) IsValid() bool {
	for _, existing := range AllowedSubtitlePlaybackModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubtitlePlaybackMode value
func (v SubtitlePlaybackMode) Ptr() *SubtitlePlaybackMode {
	return &v
}

type NullableSubtitlePlaybackMode struct {
	value *SubtitlePlaybackMode
	isSet bool
}

func (v NullableSubtitlePlaybackMode) Get() *SubtitlePlaybackMode {
	return v.value
}

func (v *NullableSubtitlePlaybackMode) Set(val *SubtitlePlaybackMode) {
	v.value = val
	v.isSet = true
}

func (v NullableSubtitlePlaybackMode) IsSet() bool {
	return v.isSet
}

func (v *NullableSubtitlePlaybackMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubtitlePlaybackMode(val *SubtitlePlaybackMode) *NullableSubtitlePlaybackMode {
	return &NullableSubtitlePlaybackMode{value: val, isSet: true}
}

func (v NullableSubtitlePlaybackMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubtitlePlaybackMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

