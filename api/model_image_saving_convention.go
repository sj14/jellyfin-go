/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ImageSavingConvention the model 'ImageSavingConvention'
type ImageSavingConvention string

// List of ImageSavingConvention
const (
	IMAGESAVINGCONVENTION_LEGACY ImageSavingConvention = "Legacy"
	IMAGESAVINGCONVENTION_COMPATIBLE ImageSavingConvention = "Compatible"
)

// All allowed values of ImageSavingConvention enum
var AllowedImageSavingConventionEnumValues = []ImageSavingConvention{
	"Legacy",
	"Compatible",
}

func (v *ImageSavingConvention) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImageSavingConvention(value)
	for _, existing := range AllowedImageSavingConventionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImageSavingConvention", value)
}

// NewImageSavingConventionFromValue returns a pointer to a valid ImageSavingConvention
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageSavingConventionFromValue(v string) (*ImageSavingConvention, error) {
	ev := ImageSavingConvention(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImageSavingConvention: valid values are %v", v, AllowedImageSavingConventionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageSavingConvention) IsValid() bool {
	for _, existing := range AllowedImageSavingConventionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageSavingConvention value
func (v ImageSavingConvention) Ptr() *ImageSavingConvention {
	return &v
}

type NullableImageSavingConvention struct {
	value *ImageSavingConvention
	isSet bool
}

func (v NullableImageSavingConvention) Get() *ImageSavingConvention {
	return v.value
}

func (v *NullableImageSavingConvention) Set(val *ImageSavingConvention) {
	v.value = val
	v.isSet = true
}

func (v NullableImageSavingConvention) IsSet() bool {
	return v.isSet
}

func (v *NullableImageSavingConvention) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageSavingConvention(val *ImageSavingConvention) *NullableImageSavingConvention {
	return &NullableImageSavingConvention{value: val, isSet: true}
}

func (v NullableImageSavingConvention) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageSavingConvention) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

