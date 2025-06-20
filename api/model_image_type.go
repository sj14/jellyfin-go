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

// ImageType Enum ImageType.
type ImageType string

// List of ImageType
const (
	IMAGETYPE_PRIMARY ImageType = "Primary"
	IMAGETYPE_ART ImageType = "Art"
	IMAGETYPE_BACKDROP ImageType = "Backdrop"
	IMAGETYPE_BANNER ImageType = "Banner"
	IMAGETYPE_LOGO ImageType = "Logo"
	IMAGETYPE_THUMB ImageType = "Thumb"
	IMAGETYPE_DISC ImageType = "Disc"
	IMAGETYPE_BOX ImageType = "Box"
	IMAGETYPE_SCREENSHOT ImageType = "Screenshot"
	IMAGETYPE_MENU ImageType = "Menu"
	IMAGETYPE_CHAPTER ImageType = "Chapter"
	IMAGETYPE_BOX_REAR ImageType = "BoxRear"
	IMAGETYPE_PROFILE ImageType = "Profile"
)

// All allowed values of ImageType enum
var AllowedImageTypeEnumValues = []ImageType{
	"Primary",
	"Art",
	"Backdrop",
	"Banner",
	"Logo",
	"Thumb",
	"Disc",
	"Box",
	"Screenshot",
	"Menu",
	"Chapter",
	"BoxRear",
	"Profile",
}

func (v *ImageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImageType(value)
	for _, existing := range AllowedImageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImageType", value)
}

// NewImageTypeFromValue returns a pointer to a valid ImageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageTypeFromValue(v string) (*ImageType, error) {
	ev := ImageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImageType: valid values are %v", v, AllowedImageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageType) IsValid() bool {
	for _, existing := range AllowedImageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageType value
func (v ImageType) Ptr() *ImageType {
	return &v
}

type NullableImageType struct {
	value *ImageType
	isSet bool
}

func (v NullableImageType) Get() *ImageType {
	return v.value
}

func (v *NullableImageType) Set(val *ImageType) {
	v.value = val
	v.isSet = true
}

func (v NullableImageType) IsSet() bool {
	return v.isSet
}

func (v *NullableImageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageType(val *ImageType) *NullableImageType {
	return &NullableImageType{value: val, isSet: true}
}

func (v NullableImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

