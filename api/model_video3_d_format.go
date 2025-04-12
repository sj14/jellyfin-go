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

// Video3DFormat the model 'Video3DFormat'
type Video3DFormat string

// List of Video3DFormat
const (
	VIDEO3DFORMAT_HALF_SIDE_BY_SIDE Video3DFormat = "HalfSideBySide"
	VIDEO3DFORMAT_FULL_SIDE_BY_SIDE Video3DFormat = "FullSideBySide"
	VIDEO3DFORMAT_FULL_TOP_AND_BOTTOM Video3DFormat = "FullTopAndBottom"
	VIDEO3DFORMAT_HALF_TOP_AND_BOTTOM Video3DFormat = "HalfTopAndBottom"
	VIDEO3DFORMAT_MVC Video3DFormat = "MVC"
)

// All allowed values of Video3DFormat enum
var AllowedVideo3DFormatEnumValues = []Video3DFormat{
	"HalfSideBySide",
	"FullSideBySide",
	"FullTopAndBottom",
	"HalfTopAndBottom",
	"MVC",
}

func (v *Video3DFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Video3DFormat(value)
	for _, existing := range AllowedVideo3DFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Video3DFormat", value)
}

// NewVideo3DFormatFromValue returns a pointer to a valid Video3DFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVideo3DFormatFromValue(v string) (*Video3DFormat, error) {
	ev := Video3DFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Video3DFormat: valid values are %v", v, AllowedVideo3DFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Video3DFormat) IsValid() bool {
	for _, existing := range AllowedVideo3DFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Video3DFormat value
func (v Video3DFormat) Ptr() *Video3DFormat {
	return &v
}

type NullableVideo3DFormat struct {
	value *Video3DFormat
	isSet bool
}

func (v NullableVideo3DFormat) Get() *Video3DFormat {
	return v.value
}

func (v *NullableVideo3DFormat) Set(val *Video3DFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableVideo3DFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableVideo3DFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideo3DFormat(val *Video3DFormat) *NullableVideo3DFormat {
	return &NullableVideo3DFormat{value: val, isSet: true}
}

func (v NullableVideo3DFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideo3DFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

