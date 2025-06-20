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

// SendCommandType Enum SendCommandType.
type SendCommandType string

// List of SendCommandType
const (
	SENDCOMMANDTYPE_UNPAUSE SendCommandType = "Unpause"
	SENDCOMMANDTYPE_PAUSE SendCommandType = "Pause"
	SENDCOMMANDTYPE_STOP SendCommandType = "Stop"
	SENDCOMMANDTYPE_SEEK SendCommandType = "Seek"
)

// All allowed values of SendCommandType enum
var AllowedSendCommandTypeEnumValues = []SendCommandType{
	"Unpause",
	"Pause",
	"Stop",
	"Seek",
}

func (v *SendCommandType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SendCommandType(value)
	for _, existing := range AllowedSendCommandTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SendCommandType", value)
}

// NewSendCommandTypeFromValue returns a pointer to a valid SendCommandType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSendCommandTypeFromValue(v string) (*SendCommandType, error) {
	ev := SendCommandType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SendCommandType: valid values are %v", v, AllowedSendCommandTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SendCommandType) IsValid() bool {
	for _, existing := range AllowedSendCommandTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SendCommandType value
func (v SendCommandType) Ptr() *SendCommandType {
	return &v
}

type NullableSendCommandType struct {
	value *SendCommandType
	isSet bool
}

func (v NullableSendCommandType) Get() *SendCommandType {
	return v.value
}

func (v *NullableSendCommandType) Set(val *SendCommandType) {
	v.value = val
	v.isSet = true
}

func (v NullableSendCommandType) IsSet() bool {
	return v.isSet
}

func (v *NullableSendCommandType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendCommandType(val *SendCommandType) *NullableSendCommandType {
	return &NullableSendCommandType{value: val, isSet: true}
}

func (v NullableSendCommandType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendCommandType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

