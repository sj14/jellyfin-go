/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// EncodingContext the model 'EncodingContext'
type EncodingContext string

// List of EncodingContext
const (
	ENCODINGCONTEXT_STREAMING EncodingContext = "Streaming"
	ENCODINGCONTEXT_STATIC EncodingContext = "Static"
)

// All allowed values of EncodingContext enum
var AllowedEncodingContextEnumValues = []EncodingContext{
	"Streaming",
	"Static",
}

func (v *EncodingContext) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EncodingContext(value)
	for _, existing := range AllowedEncodingContextEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EncodingContext", value)
}

// NewEncodingContextFromValue returns a pointer to a valid EncodingContext
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEncodingContextFromValue(v string) (*EncodingContext, error) {
	ev := EncodingContext(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EncodingContext: valid values are %v", v, AllowedEncodingContextEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EncodingContext) IsValid() bool {
	for _, existing := range AllowedEncodingContextEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EncodingContext value
func (v EncodingContext) Ptr() *EncodingContext {
	return &v
}

type NullableEncodingContext struct {
	value *EncodingContext
	isSet bool
}

func (v NullableEncodingContext) Get() *EncodingContext {
	return v.value
}

func (v *NullableEncodingContext) Set(val *EncodingContext) {
	v.value = val
	v.isSet = true
}

func (v NullableEncodingContext) IsSet() bool {
	return v.isSet
}

func (v *NullableEncodingContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncodingContext(val *EncodingContext) *NullableEncodingContext {
	return &NullableEncodingContext{value: val, isSet: true}
}

func (v NullableEncodingContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncodingContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

