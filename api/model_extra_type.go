/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ExtraType the model 'ExtraType'
type ExtraType string

// List of ExtraType
const (
	EXTRATYPE_UNKNOWN ExtraType = "Unknown"
	EXTRATYPE_CLIP ExtraType = "Clip"
	EXTRATYPE_TRAILER ExtraType = "Trailer"
	EXTRATYPE_BEHIND_THE_SCENES ExtraType = "BehindTheScenes"
	EXTRATYPE_DELETED_SCENE ExtraType = "DeletedScene"
	EXTRATYPE_INTERVIEW ExtraType = "Interview"
	EXTRATYPE_SCENE ExtraType = "Scene"
	EXTRATYPE_SAMPLE ExtraType = "Sample"
	EXTRATYPE_THEME_SONG ExtraType = "ThemeSong"
	EXTRATYPE_THEME_VIDEO ExtraType = "ThemeVideo"
	EXTRATYPE_FEATURETTE ExtraType = "Featurette"
	EXTRATYPE_SHORT ExtraType = "Short"
)

// All allowed values of ExtraType enum
var AllowedExtraTypeEnumValues = []ExtraType{
	"Unknown",
	"Clip",
	"Trailer",
	"BehindTheScenes",
	"DeletedScene",
	"Interview",
	"Scene",
	"Sample",
	"ThemeSong",
	"ThemeVideo",
	"Featurette",
	"Short",
}

func (v *ExtraType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExtraType(value)
	for _, existing := range AllowedExtraTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExtraType", value)
}

// NewExtraTypeFromValue returns a pointer to a valid ExtraType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExtraTypeFromValue(v string) (*ExtraType, error) {
	ev := ExtraType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExtraType: valid values are %v", v, AllowedExtraTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExtraType) IsValid() bool {
	for _, existing := range AllowedExtraTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExtraType value
func (v ExtraType) Ptr() *ExtraType {
	return &v
}

type NullableExtraType struct {
	value *ExtraType
	isSet bool
}

func (v NullableExtraType) Get() *ExtraType {
	return v.value
}

func (v *NullableExtraType) Set(val *ExtraType) {
	v.value = val
	v.isSet = true
}

func (v NullableExtraType) IsSet() bool {
	return v.isSet
}

func (v *NullableExtraType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtraType(val *ExtraType) *NullableExtraType {
	return &NullableExtraType{value: val, isSet: true}
}

func (v NullableExtraType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtraType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

