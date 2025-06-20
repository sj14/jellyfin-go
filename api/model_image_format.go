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

// ImageFormat Enum ImageOutputFormat.
type ImageFormat string

// List of ImageFormat
const (
	IMAGEFORMAT_BMP ImageFormat = "Bmp"
	IMAGEFORMAT_GIF ImageFormat = "Gif"
	IMAGEFORMAT_JPG ImageFormat = "Jpg"
	IMAGEFORMAT_PNG ImageFormat = "Png"
	IMAGEFORMAT_WEBP ImageFormat = "Webp"
	IMAGEFORMAT_SVG ImageFormat = "Svg"
)

// All allowed values of ImageFormat enum
var AllowedImageFormatEnumValues = []ImageFormat{
	"Bmp",
	"Gif",
	"Jpg",
	"Png",
	"Webp",
	"Svg",
}

func (v *ImageFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImageFormat(value)
	for _, existing := range AllowedImageFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImageFormat", value)
}

// NewImageFormatFromValue returns a pointer to a valid ImageFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageFormatFromValue(v string) (*ImageFormat, error) {
	ev := ImageFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImageFormat: valid values are %v", v, AllowedImageFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageFormat) IsValid() bool {
	for _, existing := range AllowedImageFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageFormat value
func (v ImageFormat) Ptr() *ImageFormat {
	return &v
}

type NullableImageFormat struct {
	value *ImageFormat
	isSet bool
}

func (v NullableImageFormat) Get() *ImageFormat {
	return v.value
}

func (v *NullableImageFormat) Set(val *ImageFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableImageFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableImageFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageFormat(val *ImageFormat) *NullableImageFormat {
	return &NullableImageFormat{value: val, isSet: true}
}

func (v NullableImageFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

