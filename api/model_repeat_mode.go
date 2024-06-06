/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// RepeatMode the model 'RepeatMode'
type RepeatMode string

// List of RepeatMode
const (
	REPEATMODE_REPEAT_NONE RepeatMode = "RepeatNone"
	REPEATMODE_REPEAT_ALL RepeatMode = "RepeatAll"
	REPEATMODE_REPEAT_ONE RepeatMode = "RepeatOne"
)

// All allowed values of RepeatMode enum
var AllowedRepeatModeEnumValues = []RepeatMode{
	"RepeatNone",
	"RepeatAll",
	"RepeatOne",
}

func (v *RepeatMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RepeatMode(value)
	for _, existing := range AllowedRepeatModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RepeatMode", value)
}

// NewRepeatModeFromValue returns a pointer to a valid RepeatMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRepeatModeFromValue(v string) (*RepeatMode, error) {
	ev := RepeatMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RepeatMode: valid values are %v", v, AllowedRepeatModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RepeatMode) IsValid() bool {
	for _, existing := range AllowedRepeatModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RepeatMode value
func (v RepeatMode) Ptr() *RepeatMode {
	return &v
}

type NullableRepeatMode struct {
	value *RepeatMode
	isSet bool
}

func (v NullableRepeatMode) Get() *RepeatMode {
	return v.value
}

func (v *NullableRepeatMode) Set(val *RepeatMode) {
	v.value = val
	v.isSet = true
}

func (v NullableRepeatMode) IsSet() bool {
	return v.isSet
}

func (v *NullableRepeatMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepeatMode(val *RepeatMode) *NullableRepeatMode {
	return &NullableRepeatMode{value: val, isSet: true}
}

func (v NullableRepeatMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepeatMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

