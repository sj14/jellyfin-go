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

// UnratedItem An enum representing an unrated item.
type UnratedItem string

// List of UnratedItem
const (
	UNRATEDITEM_MOVIE UnratedItem = "Movie"
	UNRATEDITEM_TRAILER UnratedItem = "Trailer"
	UNRATEDITEM_SERIES UnratedItem = "Series"
	UNRATEDITEM_MUSIC UnratedItem = "Music"
	UNRATEDITEM_BOOK UnratedItem = "Book"
	UNRATEDITEM_LIVE_TV_CHANNEL UnratedItem = "LiveTvChannel"
	UNRATEDITEM_LIVE_TV_PROGRAM UnratedItem = "LiveTvProgram"
	UNRATEDITEM_CHANNEL_CONTENT UnratedItem = "ChannelContent"
	UNRATEDITEM_OTHER UnratedItem = "Other"
)

// All allowed values of UnratedItem enum
var AllowedUnratedItemEnumValues = []UnratedItem{
	"Movie",
	"Trailer",
	"Series",
	"Music",
	"Book",
	"LiveTvChannel",
	"LiveTvProgram",
	"ChannelContent",
	"Other",
}

func (v *UnratedItem) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UnratedItem(value)
	for _, existing := range AllowedUnratedItemEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UnratedItem", value)
}

// NewUnratedItemFromValue returns a pointer to a valid UnratedItem
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUnratedItemFromValue(v string) (*UnratedItem, error) {
	ev := UnratedItem(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UnratedItem: valid values are %v", v, AllowedUnratedItemEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UnratedItem) IsValid() bool {
	for _, existing := range AllowedUnratedItemEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UnratedItem value
func (v UnratedItem) Ptr() *UnratedItem {
	return &v
}

type NullableUnratedItem struct {
	value *UnratedItem
	isSet bool
}

func (v NullableUnratedItem) Get() *UnratedItem {
	return v.value
}

func (v *NullableUnratedItem) Set(val *UnratedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUnratedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUnratedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnratedItem(val *UnratedItem) *NullableUnratedItem {
	return &NullableUnratedItem{value: val, isSet: true}
}

func (v NullableUnratedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnratedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

