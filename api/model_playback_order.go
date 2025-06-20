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

// PlaybackOrder Enum PlaybackOrder.
type PlaybackOrder string

// List of PlaybackOrder
const (
	PLAYBACKORDER_DEFAULT PlaybackOrder = "Default"
	PLAYBACKORDER_SHUFFLE PlaybackOrder = "Shuffle"
)

// All allowed values of PlaybackOrder enum
var AllowedPlaybackOrderEnumValues = []PlaybackOrder{
	"Default",
	"Shuffle",
}

func (v *PlaybackOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlaybackOrder(value)
	for _, existing := range AllowedPlaybackOrderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlaybackOrder", value)
}

// NewPlaybackOrderFromValue returns a pointer to a valid PlaybackOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlaybackOrderFromValue(v string) (*PlaybackOrder, error) {
	ev := PlaybackOrder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlaybackOrder: valid values are %v", v, AllowedPlaybackOrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlaybackOrder) IsValid() bool {
	for _, existing := range AllowedPlaybackOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlaybackOrder value
func (v PlaybackOrder) Ptr() *PlaybackOrder {
	return &v
}

type NullablePlaybackOrder struct {
	value *PlaybackOrder
	isSet bool
}

func (v NullablePlaybackOrder) Get() *PlaybackOrder {
	return v.value
}

func (v *NullablePlaybackOrder) Set(val *PlaybackOrder) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaybackOrder) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaybackOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaybackOrder(val *PlaybackOrder) *NullablePlaybackOrder {
	return &NullablePlaybackOrder{value: val, isSet: true}
}

func (v NullablePlaybackOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaybackOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

