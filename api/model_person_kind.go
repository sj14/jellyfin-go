/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// PersonKind The person kind.
type PersonKind string

// List of PersonKind
const (
	PERSONKIND_UNKNOWN PersonKind = "Unknown"
	PERSONKIND_ACTOR PersonKind = "Actor"
	PERSONKIND_DIRECTOR PersonKind = "Director"
	PERSONKIND_COMPOSER PersonKind = "Composer"
	PERSONKIND_WRITER PersonKind = "Writer"
	PERSONKIND_GUEST_STAR PersonKind = "GuestStar"
	PERSONKIND_PRODUCER PersonKind = "Producer"
	PERSONKIND_CONDUCTOR PersonKind = "Conductor"
	PERSONKIND_LYRICIST PersonKind = "Lyricist"
	PERSONKIND_ARRANGER PersonKind = "Arranger"
	PERSONKIND_ENGINEER PersonKind = "Engineer"
	PERSONKIND_MIXER PersonKind = "Mixer"
	PERSONKIND_REMIXER PersonKind = "Remixer"
	PERSONKIND_CREATOR PersonKind = "Creator"
	PERSONKIND_ARTIST PersonKind = "Artist"
	PERSONKIND_ALBUM_ARTIST PersonKind = "AlbumArtist"
	PERSONKIND_AUTHOR PersonKind = "Author"
	PERSONKIND_ILLUSTRATOR PersonKind = "Illustrator"
	PERSONKIND_PENCILLER PersonKind = "Penciller"
	PERSONKIND_INKER PersonKind = "Inker"
	PERSONKIND_COLORIST PersonKind = "Colorist"
	PERSONKIND_LETTERER PersonKind = "Letterer"
	PERSONKIND_COVER_ARTIST PersonKind = "CoverArtist"
	PERSONKIND_EDITOR PersonKind = "Editor"
	PERSONKIND_TRANSLATOR PersonKind = "Translator"
)

// All allowed values of PersonKind enum
var AllowedPersonKindEnumValues = []PersonKind{
	"Unknown",
	"Actor",
	"Director",
	"Composer",
	"Writer",
	"GuestStar",
	"Producer",
	"Conductor",
	"Lyricist",
	"Arranger",
	"Engineer",
	"Mixer",
	"Remixer",
	"Creator",
	"Artist",
	"AlbumArtist",
	"Author",
	"Illustrator",
	"Penciller",
	"Inker",
	"Colorist",
	"Letterer",
	"CoverArtist",
	"Editor",
	"Translator",
}

func (v *PersonKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PersonKind(value)
	for _, existing := range AllowedPersonKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PersonKind", value)
}

// NewPersonKindFromValue returns a pointer to a valid PersonKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPersonKindFromValue(v string) (*PersonKind, error) {
	ev := PersonKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PersonKind: valid values are %v", v, AllowedPersonKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PersonKind) IsValid() bool {
	for _, existing := range AllowedPersonKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PersonKind value
func (v PersonKind) Ptr() *PersonKind {
	return &v
}

type NullablePersonKind struct {
	value *PersonKind
	isSet bool
}

func (v NullablePersonKind) Get() *PersonKind {
	return v.value
}

func (v *NullablePersonKind) Set(val *PersonKind) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonKind) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonKind(val *PersonKind) *NullablePersonKind {
	return &NullablePersonKind{value: val, isSet: true}
}

func (v NullablePersonKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

