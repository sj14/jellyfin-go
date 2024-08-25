/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// VideoType Enum VideoType.
type VideoType string

// List of VideoType
const (
	VIDEOTYPE_VIDEO_FILE VideoType = "VideoFile"
	VIDEOTYPE_ISO VideoType = "Iso"
	VIDEOTYPE_DVD VideoType = "Dvd"
	VIDEOTYPE_BLU_RAY VideoType = "BluRay"
)

// All allowed values of VideoType enum
var AllowedVideoTypeEnumValues = []VideoType{
	"VideoFile",
	"Iso",
	"Dvd",
	"BluRay",
}

func (v *VideoType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VideoType(value)
	for _, existing := range AllowedVideoTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VideoType", value)
}

// NewVideoTypeFromValue returns a pointer to a valid VideoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVideoTypeFromValue(v string) (*VideoType, error) {
	ev := VideoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VideoType: valid values are %v", v, AllowedVideoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VideoType) IsValid() bool {
	for _, existing := range AllowedVideoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VideoType value
func (v VideoType) Ptr() *VideoType {
	return &v
}

type NullableVideoType struct {
	value *VideoType
	isSet bool
}

func (v NullableVideoType) Get() *VideoType {
	return v.value
}

func (v *NullableVideoType) Set(val *VideoType) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoType) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoType(val *VideoType) *NullableVideoType {
	return &NullableVideoType{value: val, isSet: true}
}

func (v NullableVideoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

