/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.8.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ExternalIdMediaType The specific media type of an MediaBrowser.Model.Providers.ExternalIdInfo.
type ExternalIdMediaType string

// List of ExternalIdMediaType
const (
	EXTERNALIDMEDIATYPE_ALBUM ExternalIdMediaType = "Album"
	EXTERNALIDMEDIATYPE_ALBUM_ARTIST ExternalIdMediaType = "AlbumArtist"
	EXTERNALIDMEDIATYPE_ARTIST ExternalIdMediaType = "Artist"
	EXTERNALIDMEDIATYPE_BOX_SET ExternalIdMediaType = "BoxSet"
	EXTERNALIDMEDIATYPE_EPISODE ExternalIdMediaType = "Episode"
	EXTERNALIDMEDIATYPE_MOVIE ExternalIdMediaType = "Movie"
	EXTERNALIDMEDIATYPE_OTHER_ARTIST ExternalIdMediaType = "OtherArtist"
	EXTERNALIDMEDIATYPE_PERSON ExternalIdMediaType = "Person"
	EXTERNALIDMEDIATYPE_RELEASE_GROUP ExternalIdMediaType = "ReleaseGroup"
	EXTERNALIDMEDIATYPE_SEASON ExternalIdMediaType = "Season"
	EXTERNALIDMEDIATYPE_SERIES ExternalIdMediaType = "Series"
	EXTERNALIDMEDIATYPE_TRACK ExternalIdMediaType = "Track"
)

// All allowed values of ExternalIdMediaType enum
var AllowedExternalIdMediaTypeEnumValues = []ExternalIdMediaType{
	"Album",
	"AlbumArtist",
	"Artist",
	"BoxSet",
	"Episode",
	"Movie",
	"OtherArtist",
	"Person",
	"ReleaseGroup",
	"Season",
	"Series",
	"Track",
}

func (v *ExternalIdMediaType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExternalIdMediaType(value)
	for _, existing := range AllowedExternalIdMediaTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExternalIdMediaType", value)
}

// NewExternalIdMediaTypeFromValue returns a pointer to a valid ExternalIdMediaType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExternalIdMediaTypeFromValue(v string) (*ExternalIdMediaType, error) {
	ev := ExternalIdMediaType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExternalIdMediaType: valid values are %v", v, AllowedExternalIdMediaTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExternalIdMediaType) IsValid() bool {
	for _, existing := range AllowedExternalIdMediaTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExternalIdMediaType value
func (v ExternalIdMediaType) Ptr() *ExternalIdMediaType {
	return &v
}

type NullableExternalIdMediaType struct {
	value *ExternalIdMediaType
	isSet bool
}

func (v NullableExternalIdMediaType) Get() *ExternalIdMediaType {
	return v.value
}

func (v *NullableExternalIdMediaType) Set(val *ExternalIdMediaType) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalIdMediaType) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalIdMediaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalIdMediaType(val *ExternalIdMediaType) *NullableExternalIdMediaType {
	return &NullableExternalIdMediaType{value: val, isSet: true}
}

func (v NullableExternalIdMediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalIdMediaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

