/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ChannelMediaContentType the model 'ChannelMediaContentType'
type ChannelMediaContentType string

// List of ChannelMediaContentType
const (
	CHANNELMEDIACONTENTTYPE_CLIP ChannelMediaContentType = "Clip"
	CHANNELMEDIACONTENTTYPE_PODCAST ChannelMediaContentType = "Podcast"
	CHANNELMEDIACONTENTTYPE_TRAILER ChannelMediaContentType = "Trailer"
	CHANNELMEDIACONTENTTYPE_MOVIE ChannelMediaContentType = "Movie"
	CHANNELMEDIACONTENTTYPE_EPISODE ChannelMediaContentType = "Episode"
	CHANNELMEDIACONTENTTYPE_SONG ChannelMediaContentType = "Song"
	CHANNELMEDIACONTENTTYPE_MOVIE_EXTRA ChannelMediaContentType = "MovieExtra"
	CHANNELMEDIACONTENTTYPE_TV_EXTRA ChannelMediaContentType = "TvExtra"
)

// All allowed values of ChannelMediaContentType enum
var AllowedChannelMediaContentTypeEnumValues = []ChannelMediaContentType{
	"Clip",
	"Podcast",
	"Trailer",
	"Movie",
	"Episode",
	"Song",
	"MovieExtra",
	"TvExtra",
}

func (v *ChannelMediaContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ChannelMediaContentType(value)
	for _, existing := range AllowedChannelMediaContentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ChannelMediaContentType", value)
}

// NewChannelMediaContentTypeFromValue returns a pointer to a valid ChannelMediaContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewChannelMediaContentTypeFromValue(v string) (*ChannelMediaContentType, error) {
	ev := ChannelMediaContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ChannelMediaContentType: valid values are %v", v, AllowedChannelMediaContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ChannelMediaContentType) IsValid() bool {
	for _, existing := range AllowedChannelMediaContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ChannelMediaContentType value
func (v ChannelMediaContentType) Ptr() *ChannelMediaContentType {
	return &v
}

type NullableChannelMediaContentType struct {
	value *ChannelMediaContentType
	isSet bool
}

func (v NullableChannelMediaContentType) Get() *ChannelMediaContentType {
	return v.value
}

func (v *NullableChannelMediaContentType) Set(val *ChannelMediaContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelMediaContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelMediaContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelMediaContentType(val *ChannelMediaContentType) *NullableChannelMediaContentType {
	return &NullableChannelMediaContentType{value: val, isSet: true}
}

func (v NullableChannelMediaContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelMediaContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

