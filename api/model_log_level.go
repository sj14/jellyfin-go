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

// LogLevel the model 'LogLevel'
type LogLevel string

// List of LogLevel
const (
	LOGLEVEL_TRACE LogLevel = "Trace"
	LOGLEVEL_DEBUG LogLevel = "Debug"
	LOGLEVEL_INFORMATION LogLevel = "Information"
	LOGLEVEL_WARNING LogLevel = "Warning"
	LOGLEVEL_ERROR LogLevel = "Error"
	LOGLEVEL_CRITICAL LogLevel = "Critical"
	LOGLEVEL_NONE LogLevel = "None"
)

// All allowed values of LogLevel enum
var AllowedLogLevelEnumValues = []LogLevel{
	"Trace",
	"Debug",
	"Information",
	"Warning",
	"Error",
	"Critical",
	"None",
}

func (v *LogLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogLevel(value)
	for _, existing := range AllowedLogLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogLevel", value)
}

// NewLogLevelFromValue returns a pointer to a valid LogLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLogLevelFromValue(v string) (*LogLevel, error) {
	ev := LogLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LogLevel: valid values are %v", v, AllowedLogLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LogLevel) IsValid() bool {
	for _, existing := range AllowedLogLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LogLevel value
func (v LogLevel) Ptr() *LogLevel {
	return &v
}

type NullableLogLevel struct {
	value *LogLevel
	isSet bool
}

func (v NullableLogLevel) Get() *LogLevel {
	return v.value
}

func (v *NullableLogLevel) Set(val *LogLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableLogLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableLogLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogLevel(val *LogLevel) *NullableLogLevel {
	return &NullableLogLevel{value: val, isSet: true}
}

func (v NullableLogLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

