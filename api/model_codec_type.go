/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// CodecType the model 'CodecType'
type CodecType string

// List of CodecType
const (
	CODECTYPE_VIDEO CodecType = "Video"
	CODECTYPE_VIDEO_AUDIO CodecType = "VideoAudio"
	CODECTYPE_AUDIO CodecType = "Audio"
)

// All allowed values of CodecType enum
var AllowedCodecTypeEnumValues = []CodecType{
	"Video",
	"VideoAudio",
	"Audio",
}

func (v *CodecType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CodecType(value)
	for _, existing := range AllowedCodecTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CodecType", value)
}

// NewCodecTypeFromValue returns a pointer to a valid CodecType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCodecTypeFromValue(v string) (*CodecType, error) {
	ev := CodecType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CodecType: valid values are %v", v, AllowedCodecTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CodecType) IsValid() bool {
	for _, existing := range AllowedCodecTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CodecType value
func (v CodecType) Ptr() *CodecType {
	return &v
}

type NullableCodecType struct {
	value *CodecType
	isSet bool
}

func (v NullableCodecType) Get() *CodecType {
	return v.value
}

func (v *NullableCodecType) Set(val *CodecType) {
	v.value = val
	v.isSet = true
}

func (v NullableCodecType) IsSet() bool {
	return v.isSet
}

func (v *NullableCodecType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodecType(val *CodecType) *NullableCodecType {
	return &NullableCodecType{value: val, isSet: true}
}

func (v NullableCodecType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodecType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

