/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// DayPattern the model 'DayPattern'
type DayPattern string

// List of DayPattern
const (
	DAYPATTERN_DAILY DayPattern = "Daily"
	DAYPATTERN_WEEKDAYS DayPattern = "Weekdays"
	DAYPATTERN_WEEKENDS DayPattern = "Weekends"
)

// All allowed values of DayPattern enum
var AllowedDayPatternEnumValues = []DayPattern{
	"Daily",
	"Weekdays",
	"Weekends",
}

func (v *DayPattern) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DayPattern(value)
	for _, existing := range AllowedDayPatternEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DayPattern", value)
}

// NewDayPatternFromValue returns a pointer to a valid DayPattern
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDayPatternFromValue(v string) (*DayPattern, error) {
	ev := DayPattern(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DayPattern: valid values are %v", v, AllowedDayPatternEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DayPattern) IsValid() bool {
	for _, existing := range AllowedDayPatternEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DayPattern value
func (v DayPattern) Ptr() *DayPattern {
	return &v
}

type NullableDayPattern struct {
	value *DayPattern
	isSet bool
}

func (v NullableDayPattern) Get() *DayPattern {
	return v.value
}

func (v *NullableDayPattern) Set(val *DayPattern) {
	v.value = val
	v.isSet = true
}

func (v NullableDayPattern) IsSet() bool {
	return v.isSet
}

func (v *NullableDayPattern) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDayPattern(val *DayPattern) *NullableDayPattern {
	return &NullableDayPattern{value: val, isSet: true}
}

func (v NullableDayPattern) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDayPattern) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

