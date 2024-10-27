/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// GroupShuffleMode Enum GroupShuffleMode.
type GroupShuffleMode string

// List of GroupShuffleMode
const (
	GROUPSHUFFLEMODE_SORTED GroupShuffleMode = "Sorted"
	GROUPSHUFFLEMODE_SHUFFLE GroupShuffleMode = "Shuffle"
)

// All allowed values of GroupShuffleMode enum
var AllowedGroupShuffleModeEnumValues = []GroupShuffleMode{
	"Sorted",
	"Shuffle",
}

func (v *GroupShuffleMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupShuffleMode(value)
	for _, existing := range AllowedGroupShuffleModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupShuffleMode", value)
}

// NewGroupShuffleModeFromValue returns a pointer to a valid GroupShuffleMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupShuffleModeFromValue(v string) (*GroupShuffleMode, error) {
	ev := GroupShuffleMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupShuffleMode: valid values are %v", v, AllowedGroupShuffleModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupShuffleMode) IsValid() bool {
	for _, existing := range AllowedGroupShuffleModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupShuffleMode value
func (v GroupShuffleMode) Ptr() *GroupShuffleMode {
	return &v
}

type NullableGroupShuffleMode struct {
	value *GroupShuffleMode
	isSet bool
}

func (v NullableGroupShuffleMode) Get() *GroupShuffleMode {
	return v.value
}

func (v *NullableGroupShuffleMode) Set(val *GroupShuffleMode) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupShuffleMode) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupShuffleMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupShuffleMode(val *GroupShuffleMode) *NullableGroupShuffleMode {
	return &NullableGroupShuffleMode{value: val, isSet: true}
}

func (v NullableGroupShuffleMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupShuffleMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

