/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// VideoRange An enum representing video ranges.
type VideoRange string

// List of VideoRange
const (
	VIDEORANGE_UNKNOWN VideoRange = "Unknown"
	VIDEORANGE_SDR VideoRange = "SDR"
	VIDEORANGE_HDR VideoRange = "HDR"
)

// All allowed values of VideoRange enum
var AllowedVideoRangeEnumValues = []VideoRange{
	"Unknown",
	"SDR",
	"HDR",
}

func (v *VideoRange) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VideoRange(value)
	for _, existing := range AllowedVideoRangeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VideoRange", value)
}

// NewVideoRangeFromValue returns a pointer to a valid VideoRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVideoRangeFromValue(v string) (*VideoRange, error) {
	ev := VideoRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VideoRange: valid values are %v", v, AllowedVideoRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VideoRange) IsValid() bool {
	for _, existing := range AllowedVideoRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VideoRange value
func (v VideoRange) Ptr() *VideoRange {
	return &v
}

type NullableVideoRange struct {
	value *VideoRange
	isSet bool
}

func (v NullableVideoRange) Get() *VideoRange {
	return v.value
}

func (v *NullableVideoRange) Set(val *VideoRange) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoRange) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoRange(val *VideoRange) *NullableVideoRange {
	return &NullableVideoRange{value: val, isSet: true}
}

func (v NullableVideoRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

