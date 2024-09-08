/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ImageOrientation the model 'ImageOrientation'
type ImageOrientation string

// List of ImageOrientation
const (
	IMAGEORIENTATION_TOP_LEFT ImageOrientation = "TopLeft"
	IMAGEORIENTATION_TOP_RIGHT ImageOrientation = "TopRight"
	IMAGEORIENTATION_BOTTOM_RIGHT ImageOrientation = "BottomRight"
	IMAGEORIENTATION_BOTTOM_LEFT ImageOrientation = "BottomLeft"
	IMAGEORIENTATION_LEFT_TOP ImageOrientation = "LeftTop"
	IMAGEORIENTATION_RIGHT_TOP ImageOrientation = "RightTop"
	IMAGEORIENTATION_RIGHT_BOTTOM ImageOrientation = "RightBottom"
	IMAGEORIENTATION_LEFT_BOTTOM ImageOrientation = "LeftBottom"
)

// All allowed values of ImageOrientation enum
var AllowedImageOrientationEnumValues = []ImageOrientation{
	"TopLeft",
	"TopRight",
	"BottomRight",
	"BottomLeft",
	"LeftTop",
	"RightTop",
	"RightBottom",
	"LeftBottom",
}

func (v *ImageOrientation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImageOrientation(value)
	for _, existing := range AllowedImageOrientationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImageOrientation", value)
}

// NewImageOrientationFromValue returns a pointer to a valid ImageOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageOrientationFromValue(v string) (*ImageOrientation, error) {
	ev := ImageOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImageOrientation: valid values are %v", v, AllowedImageOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageOrientation) IsValid() bool {
	for _, existing := range AllowedImageOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageOrientation value
func (v ImageOrientation) Ptr() *ImageOrientation {
	return &v
}

type NullableImageOrientation struct {
	value *ImageOrientation
	isSet bool
}

func (v NullableImageOrientation) Get() *ImageOrientation {
	return v.value
}

func (v *NullableImageOrientation) Set(val *ImageOrientation) {
	v.value = val
	v.isSet = true
}

func (v NullableImageOrientation) IsSet() bool {
	return v.isSet
}

func (v *NullableImageOrientation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageOrientation(val *ImageOrientation) *NullableImageOrientation {
	return &NullableImageOrientation{value: val, isSet: true}
}

func (v NullableImageOrientation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageOrientation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

