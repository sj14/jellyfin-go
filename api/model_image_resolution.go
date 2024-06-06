/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ImageResolution Enum ImageResolution.
type ImageResolution string

// List of ImageResolution
const (
	IMAGERESOLUTION_MATCH_SOURCE ImageResolution = "MatchSource"
	IMAGERESOLUTION_P144 ImageResolution = "P144"
	IMAGERESOLUTION_P240 ImageResolution = "P240"
	IMAGERESOLUTION_P360 ImageResolution = "P360"
	IMAGERESOLUTION_P480 ImageResolution = "P480"
	IMAGERESOLUTION_P720 ImageResolution = "P720"
	IMAGERESOLUTION_P1080 ImageResolution = "P1080"
	IMAGERESOLUTION_P1440 ImageResolution = "P1440"
	IMAGERESOLUTION_P2160 ImageResolution = "P2160"
)

// All allowed values of ImageResolution enum
var AllowedImageResolutionEnumValues = []ImageResolution{
	"MatchSource",
	"P144",
	"P240",
	"P360",
	"P480",
	"P720",
	"P1080",
	"P1440",
	"P2160",
}

func (v *ImageResolution) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImageResolution(value)
	for _, existing := range AllowedImageResolutionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImageResolution", value)
}

// NewImageResolutionFromValue returns a pointer to a valid ImageResolution
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageResolutionFromValue(v string) (*ImageResolution, error) {
	ev := ImageResolution(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImageResolution: valid values are %v", v, AllowedImageResolutionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageResolution) IsValid() bool {
	for _, existing := range AllowedImageResolutionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageResolution value
func (v ImageResolution) Ptr() *ImageResolution {
	return &v
}

type NullableImageResolution struct {
	value *ImageResolution
	isSet bool
}

func (v NullableImageResolution) Get() *ImageResolution {
	return v.value
}

func (v *NullableImageResolution) Set(val *ImageResolution) {
	v.value = val
	v.isSet = true
}

func (v NullableImageResolution) IsSet() bool {
	return v.isSet
}

func (v *NullableImageResolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageResolution(val *ImageResolution) *NullableImageResolution {
	return &NullableImageResolution{value: val, isSet: true}
}

func (v NullableImageResolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageResolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

