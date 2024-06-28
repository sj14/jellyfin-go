/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// GroupRepeatMode Enum GroupRepeatMode.
type GroupRepeatMode string

// List of GroupRepeatMode
const (
	GROUPREPEATMODE_REPEAT_ONE GroupRepeatMode = "RepeatOne"
	GROUPREPEATMODE_REPEAT_ALL GroupRepeatMode = "RepeatAll"
	GROUPREPEATMODE_REPEAT_NONE GroupRepeatMode = "RepeatNone"
)

// All allowed values of GroupRepeatMode enum
var AllowedGroupRepeatModeEnumValues = []GroupRepeatMode{
	"RepeatOne",
	"RepeatAll",
	"RepeatNone",
}

func (v *GroupRepeatMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupRepeatMode(value)
	for _, existing := range AllowedGroupRepeatModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupRepeatMode", value)
}

// NewGroupRepeatModeFromValue returns a pointer to a valid GroupRepeatMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupRepeatModeFromValue(v string) (*GroupRepeatMode, error) {
	ev := GroupRepeatMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupRepeatMode: valid values are %v", v, AllowedGroupRepeatModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupRepeatMode) IsValid() bool {
	for _, existing := range AllowedGroupRepeatModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupRepeatMode value
func (v GroupRepeatMode) Ptr() *GroupRepeatMode {
	return &v
}

type NullableGroupRepeatMode struct {
	value *GroupRepeatMode
	isSet bool
}

func (v NullableGroupRepeatMode) Get() *GroupRepeatMode {
	return v.value
}

func (v *NullableGroupRepeatMode) Set(val *GroupRepeatMode) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRepeatMode) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRepeatMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRepeatMode(val *GroupRepeatMode) *NullableGroupRepeatMode {
	return &NullableGroupRepeatMode{value: val, isSet: true}
}

func (v NullableGroupRepeatMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRepeatMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

