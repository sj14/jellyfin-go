/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// TonemappingRange Enum containing tonemapping ranges.
type TonemappingRange string

// List of TonemappingRange
const (
	TONEMAPPINGRANGE_AUTO TonemappingRange = "auto"
	TONEMAPPINGRANGE_TV TonemappingRange = "tv"
	TONEMAPPINGRANGE_PC TonemappingRange = "pc"
)

// All allowed values of TonemappingRange enum
var AllowedTonemappingRangeEnumValues = []TonemappingRange{
	"auto",
	"tv",
	"pc",
}

func (v *TonemappingRange) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TonemappingRange(value)
	for _, existing := range AllowedTonemappingRangeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TonemappingRange", value)
}

// NewTonemappingRangeFromValue returns a pointer to a valid TonemappingRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTonemappingRangeFromValue(v string) (*TonemappingRange, error) {
	ev := TonemappingRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TonemappingRange: valid values are %v", v, AllowedTonemappingRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TonemappingRange) IsValid() bool {
	for _, existing := range AllowedTonemappingRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TonemappingRange value
func (v TonemappingRange) Ptr() *TonemappingRange {
	return &v
}

type NullableTonemappingRange struct {
	value *TonemappingRange
	isSet bool
}

func (v NullableTonemappingRange) Get() *TonemappingRange {
	return v.value
}

func (v *NullableTonemappingRange) Set(val *TonemappingRange) {
	v.value = val
	v.isSet = true
}

func (v NullableTonemappingRange) IsSet() bool {
	return v.isSet
}

func (v *NullableTonemappingRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTonemappingRange(val *TonemappingRange) *NullableTonemappingRange {
	return &NullableTonemappingRange{value: val, isSet: true}
}

func (v NullableTonemappingRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTonemappingRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

