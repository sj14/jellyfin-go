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

// PlayMethod the model 'PlayMethod'
type PlayMethod string

// List of PlayMethod
const (
	PLAYMETHOD_TRANSCODE PlayMethod = "Transcode"
	PLAYMETHOD_DIRECT_STREAM PlayMethod = "DirectStream"
	PLAYMETHOD_DIRECT_PLAY PlayMethod = "DirectPlay"
)

// All allowed values of PlayMethod enum
var AllowedPlayMethodEnumValues = []PlayMethod{
	"Transcode",
	"DirectStream",
	"DirectPlay",
}

func (v *PlayMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlayMethod(value)
	for _, existing := range AllowedPlayMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlayMethod", value)
}

// NewPlayMethodFromValue returns a pointer to a valid PlayMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlayMethodFromValue(v string) (*PlayMethod, error) {
	ev := PlayMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlayMethod: valid values are %v", v, AllowedPlayMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlayMethod) IsValid() bool {
	for _, existing := range AllowedPlayMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlayMethod value
func (v PlayMethod) Ptr() *PlayMethod {
	return &v
}

type NullablePlayMethod struct {
	value *PlayMethod
	isSet bool
}

func (v NullablePlayMethod) Get() *PlayMethod {
	return v.value
}

func (v *NullablePlayMethod) Set(val *PlayMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayMethod(val *PlayMethod) *NullablePlayMethod {
	return &NullablePlayMethod{value: val, isSet: true}
}

func (v NullablePlayMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

