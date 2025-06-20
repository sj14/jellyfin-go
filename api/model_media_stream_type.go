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

// MediaStreamType Enum MediaStreamType.
type MediaStreamType string

// List of MediaStreamType
const (
	MEDIASTREAMTYPE_AUDIO MediaStreamType = "Audio"
	MEDIASTREAMTYPE_VIDEO MediaStreamType = "Video"
	MEDIASTREAMTYPE_SUBTITLE MediaStreamType = "Subtitle"
	MEDIASTREAMTYPE_EMBEDDED_IMAGE MediaStreamType = "EmbeddedImage"
	MEDIASTREAMTYPE_DATA MediaStreamType = "Data"
	MEDIASTREAMTYPE_LYRIC MediaStreamType = "Lyric"
)

// All allowed values of MediaStreamType enum
var AllowedMediaStreamTypeEnumValues = []MediaStreamType{
	"Audio",
	"Video",
	"Subtitle",
	"EmbeddedImage",
	"Data",
	"Lyric",
}

func (v *MediaStreamType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MediaStreamType(value)
	for _, existing := range AllowedMediaStreamTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MediaStreamType", value)
}

// NewMediaStreamTypeFromValue returns a pointer to a valid MediaStreamType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMediaStreamTypeFromValue(v string) (*MediaStreamType, error) {
	ev := MediaStreamType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MediaStreamType: valid values are %v", v, AllowedMediaStreamTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MediaStreamType) IsValid() bool {
	for _, existing := range AllowedMediaStreamTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MediaStreamType value
func (v MediaStreamType) Ptr() *MediaStreamType {
	return &v
}

type NullableMediaStreamType struct {
	value *MediaStreamType
	isSet bool
}

func (v NullableMediaStreamType) Get() *MediaStreamType {
	return v.value
}

func (v *NullableMediaStreamType) Set(val *MediaStreamType) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaStreamType) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaStreamType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaStreamType(val *MediaStreamType) *NullableMediaStreamType {
	return &NullableMediaStreamType{value: val, isSet: true}
}

func (v NullableMediaStreamType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaStreamType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

