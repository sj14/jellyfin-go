/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// EmbeddedSubtitleOptions An enum representing the options to disable embedded subs.
type EmbeddedSubtitleOptions string

// List of EmbeddedSubtitleOptions
const (
	EMBEDDEDSUBTITLEOPTIONS_ALLOW_ALL EmbeddedSubtitleOptions = "AllowAll"
	EMBEDDEDSUBTITLEOPTIONS_ALLOW_TEXT EmbeddedSubtitleOptions = "AllowText"
	EMBEDDEDSUBTITLEOPTIONS_ALLOW_IMAGE EmbeddedSubtitleOptions = "AllowImage"
	EMBEDDEDSUBTITLEOPTIONS_ALLOW_NONE EmbeddedSubtitleOptions = "AllowNone"
)

// All allowed values of EmbeddedSubtitleOptions enum
var AllowedEmbeddedSubtitleOptionsEnumValues = []EmbeddedSubtitleOptions{
	"AllowAll",
	"AllowText",
	"AllowImage",
	"AllowNone",
}

func (v *EmbeddedSubtitleOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EmbeddedSubtitleOptions(value)
	for _, existing := range AllowedEmbeddedSubtitleOptionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EmbeddedSubtitleOptions", value)
}

// NewEmbeddedSubtitleOptionsFromValue returns a pointer to a valid EmbeddedSubtitleOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEmbeddedSubtitleOptionsFromValue(v string) (*EmbeddedSubtitleOptions, error) {
	ev := EmbeddedSubtitleOptions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EmbeddedSubtitleOptions: valid values are %v", v, AllowedEmbeddedSubtitleOptionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EmbeddedSubtitleOptions) IsValid() bool {
	for _, existing := range AllowedEmbeddedSubtitleOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EmbeddedSubtitleOptions value
func (v EmbeddedSubtitleOptions) Ptr() *EmbeddedSubtitleOptions {
	return &v
}

type NullableEmbeddedSubtitleOptions struct {
	value *EmbeddedSubtitleOptions
	isSet bool
}

func (v NullableEmbeddedSubtitleOptions) Get() *EmbeddedSubtitleOptions {
	return v.value
}

func (v *NullableEmbeddedSubtitleOptions) Set(val *EmbeddedSubtitleOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedSubtitleOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedSubtitleOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedSubtitleOptions(val *EmbeddedSubtitleOptions) *NullableEmbeddedSubtitleOptions {
	return &NullableEmbeddedSubtitleOptions{value: val, isSet: true}
}

func (v NullableEmbeddedSubtitleOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedSubtitleOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

