/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// MediaType Media types.
type MediaType string

// List of MediaType
const (
	MEDIATYPE_UNKNOWN MediaType = "Unknown"
	MEDIATYPE_VIDEO MediaType = "Video"
	MEDIATYPE_AUDIO MediaType = "Audio"
	MEDIATYPE_PHOTO MediaType = "Photo"
	MEDIATYPE_BOOK MediaType = "Book"
)

// All allowed values of MediaType enum
var AllowedMediaTypeEnumValues = []MediaType{
	"Unknown",
	"Video",
	"Audio",
	"Photo",
	"Book",
}

func (v *MediaType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MediaType(value)
	for _, existing := range AllowedMediaTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MediaType", value)
}

// NewMediaTypeFromValue returns a pointer to a valid MediaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMediaTypeFromValue(v string) (*MediaType, error) {
	ev := MediaType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MediaType: valid values are %v", v, AllowedMediaTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MediaType) IsValid() bool {
	for _, existing := range AllowedMediaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MediaType value
func (v MediaType) Ptr() *MediaType {
	return &v
}

type NullableMediaType struct {
	value *MediaType
	isSet bool
}

func (v NullableMediaType) Get() *MediaType {
	return v.value
}

func (v *NullableMediaType) Set(val *MediaType) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaType) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaType(val *MediaType) *NullableMediaType {
	return &NullableMediaType{value: val, isSet: true}
}

func (v NullableMediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

