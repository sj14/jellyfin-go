/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// MediaSourceType the model 'MediaSourceType'
type MediaSourceType string

// List of MediaSourceType
const (
	MEDIASOURCETYPE_DEFAULT MediaSourceType = "Default"
	MEDIASOURCETYPE_GROUPING MediaSourceType = "Grouping"
	MEDIASOURCETYPE_PLACEHOLDER MediaSourceType = "Placeholder"
)

// All allowed values of MediaSourceType enum
var AllowedMediaSourceTypeEnumValues = []MediaSourceType{
	"Default",
	"Grouping",
	"Placeholder",
}

func (v *MediaSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MediaSourceType(value)
	for _, existing := range AllowedMediaSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MediaSourceType", value)
}

// NewMediaSourceTypeFromValue returns a pointer to a valid MediaSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMediaSourceTypeFromValue(v string) (*MediaSourceType, error) {
	ev := MediaSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MediaSourceType: valid values are %v", v, AllowedMediaSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MediaSourceType) IsValid() bool {
	for _, existing := range AllowedMediaSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MediaSourceType value
func (v MediaSourceType) Ptr() *MediaSourceType {
	return &v
}

type NullableMediaSourceType struct {
	value *MediaSourceType
	isSet bool
}

func (v NullableMediaSourceType) Get() *MediaSourceType {
	return v.value
}

func (v *NullableMediaSourceType) Set(val *MediaSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaSourceType(val *MediaSourceType) *NullableMediaSourceType {
	return &NullableMediaSourceType{value: val, isSet: true}
}

func (v NullableMediaSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

