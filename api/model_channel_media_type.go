/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ChannelMediaType the model 'ChannelMediaType'
type ChannelMediaType string

// List of ChannelMediaType
const (
	CHANNELMEDIATYPE_AUDIO ChannelMediaType = "Audio"
	CHANNELMEDIATYPE_VIDEO ChannelMediaType = "Video"
	CHANNELMEDIATYPE_PHOTO ChannelMediaType = "Photo"
)

// All allowed values of ChannelMediaType enum
var AllowedChannelMediaTypeEnumValues = []ChannelMediaType{
	"Audio",
	"Video",
	"Photo",
}

func (v *ChannelMediaType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChannelMediaType(value)
	for _, existing := range AllowedChannelMediaTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChannelMediaType", value)
}

// NewChannelMediaTypeFromValue returns a pointer to a valid ChannelMediaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChannelMediaTypeFromValue(v string) (*ChannelMediaType, error) {
	ev := ChannelMediaType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChannelMediaType: valid values are %v", v, AllowedChannelMediaTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelMediaType) IsValid() bool {
	for _, existing := range AllowedChannelMediaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChannelMediaType value
func (v ChannelMediaType) Ptr() *ChannelMediaType {
	return &v
}

type NullableChannelMediaType struct {
	value *ChannelMediaType
	isSet bool
}

func (v NullableChannelMediaType) Get() *ChannelMediaType {
	return v.value
}

func (v *NullableChannelMediaType) Set(val *ChannelMediaType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelMediaType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelMediaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelMediaType(val *ChannelMediaType) *NullableChannelMediaType {
	return &NullableChannelMediaType{value: val, isSet: true}
}

func (v NullableChannelMediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelMediaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

