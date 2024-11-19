/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// TonemappingMode Enum containing tonemapping modes.
type TonemappingMode string

// List of TonemappingMode
const (
	TONEMAPPINGMODE_AUTO TonemappingMode = "auto"
	TONEMAPPINGMODE_MAX TonemappingMode = "max"
	TONEMAPPINGMODE_RGB TonemappingMode = "rgb"
	TONEMAPPINGMODE_LUM TonemappingMode = "lum"
	TONEMAPPINGMODE_ITP TonemappingMode = "itp"
)

// All allowed values of TonemappingMode enum
var AllowedTonemappingModeEnumValues = []TonemappingMode{
	"auto",
	"max",
	"rgb",
	"lum",
	"itp",
}

func (v *TonemappingMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TonemappingMode(value)
	for _, existing := range AllowedTonemappingModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TonemappingMode", value)
}

// NewTonemappingModeFromValue returns a pointer to a valid TonemappingMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTonemappingModeFromValue(v string) (*TonemappingMode, error) {
	ev := TonemappingMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TonemappingMode: valid values are %v", v, AllowedTonemappingModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TonemappingMode) IsValid() bool {
	for _, existing := range AllowedTonemappingModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TonemappingMode value
func (v TonemappingMode) Ptr() *TonemappingMode {
	return &v
}

type NullableTonemappingMode struct {
	value *TonemappingMode
	isSet bool
}

func (v NullableTonemappingMode) Get() *TonemappingMode {
	return v.value
}

func (v *NullableTonemappingMode) Set(val *TonemappingMode) {
	v.value = val
	v.isSet = true
}

func (v NullableTonemappingMode) IsSet() bool {
	return v.isSet
}

func (v *NullableTonemappingMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTonemappingMode(val *TonemappingMode) *NullableTonemappingMode {
	return &NullableTonemappingMode{value: val, isSet: true}
}

func (v NullableTonemappingMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTonemappingMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

