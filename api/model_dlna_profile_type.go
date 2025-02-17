/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// DlnaProfileType the model 'DlnaProfileType'
type DlnaProfileType string

// List of DlnaProfileType
const (
	DLNAPROFILETYPE_AUDIO DlnaProfileType = "Audio"
	DLNAPROFILETYPE_VIDEO DlnaProfileType = "Video"
	DLNAPROFILETYPE_PHOTO DlnaProfileType = "Photo"
	DLNAPROFILETYPE_SUBTITLE DlnaProfileType = "Subtitle"
	DLNAPROFILETYPE_LYRIC DlnaProfileType = "Lyric"
)

// All allowed values of DlnaProfileType enum
var AllowedDlnaProfileTypeEnumValues = []DlnaProfileType{
	"Audio",
	"Video",
	"Photo",
	"Subtitle",
	"Lyric",
}

func (v *DlnaProfileType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DlnaProfileType(value)
	for _, existing := range AllowedDlnaProfileTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DlnaProfileType", value)
}

// NewDlnaProfileTypeFromValue returns a pointer to a valid DlnaProfileType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDlnaProfileTypeFromValue(v string) (*DlnaProfileType, error) {
	ev := DlnaProfileType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DlnaProfileType: valid values are %v", v, AllowedDlnaProfileTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DlnaProfileType) IsValid() bool {
	for _, existing := range AllowedDlnaProfileTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DlnaProfileType value
func (v DlnaProfileType) Ptr() *DlnaProfileType {
	return &v
}

type NullableDlnaProfileType struct {
	value *DlnaProfileType
	isSet bool
}

func (v NullableDlnaProfileType) Get() *DlnaProfileType {
	return v.value
}

func (v *NullableDlnaProfileType) Set(val *DlnaProfileType) {
	v.value = val
	v.isSet = true
}

func (v NullableDlnaProfileType) IsSet() bool {
	return v.isSet
}

func (v *NullableDlnaProfileType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDlnaProfileType(val *DlnaProfileType) *NullableDlnaProfileType {
	return &NullableDlnaProfileType{value: val, isSet: true}
}

func (v NullableDlnaProfileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDlnaProfileType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

