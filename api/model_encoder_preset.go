/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// EncoderPreset Enum containing encoder presets.
type EncoderPreset string

// List of EncoderPreset
const (
	ENCODERPRESET_AUTO EncoderPreset = "auto"
	ENCODERPRESET_PLACEBO EncoderPreset = "placebo"
	ENCODERPRESET_VERYSLOW EncoderPreset = "veryslow"
	ENCODERPRESET_SLOWER EncoderPreset = "slower"
	ENCODERPRESET_SLOW EncoderPreset = "slow"
	ENCODERPRESET_MEDIUM EncoderPreset = "medium"
	ENCODERPRESET_FAST EncoderPreset = "fast"
	ENCODERPRESET_FASTER EncoderPreset = "faster"
	ENCODERPRESET_VERYFAST EncoderPreset = "veryfast"
	ENCODERPRESET_SUPERFAST EncoderPreset = "superfast"
	ENCODERPRESET_ULTRAFAST EncoderPreset = "ultrafast"
)

// All allowed values of EncoderPreset enum
var AllowedEncoderPresetEnumValues = []EncoderPreset{
	"auto",
	"placebo",
	"veryslow",
	"slower",
	"slow",
	"medium",
	"fast",
	"faster",
	"veryfast",
	"superfast",
	"ultrafast",
}

func (v *EncoderPreset) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EncoderPreset(value)
	for _, existing := range AllowedEncoderPresetEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EncoderPreset", value)
}

// NewEncoderPresetFromValue returns a pointer to a valid EncoderPreset
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEncoderPresetFromValue(v string) (*EncoderPreset, error) {
	ev := EncoderPreset(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EncoderPreset: valid values are %v", v, AllowedEncoderPresetEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EncoderPreset) IsValid() bool {
	for _, existing := range AllowedEncoderPresetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EncoderPreset value
func (v EncoderPreset) Ptr() *EncoderPreset {
	return &v
}

type NullableEncoderPreset struct {
	value *EncoderPreset
	isSet bool
}

func (v NullableEncoderPreset) Get() *EncoderPreset {
	return v.value
}

func (v *NullableEncoderPreset) Set(val *EncoderPreset) {
	v.value = val
	v.isSet = true
}

func (v NullableEncoderPreset) IsSet() bool {
	return v.isSet
}

func (v *NullableEncoderPreset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncoderPreset(val *EncoderPreset) *NullableEncoderPreset {
	return &NullableEncoderPreset{value: val, isSet: true}
}

func (v NullableEncoderPreset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncoderPreset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

