/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// DeinterlaceMethod Enum containing deinterlace methods.
type DeinterlaceMethod string

// List of DeinterlaceMethod
const (
	DEINTERLACEMETHOD_YADIF DeinterlaceMethod = "yadif"
	DEINTERLACEMETHOD_BWDIF DeinterlaceMethod = "bwdif"
)

// All allowed values of DeinterlaceMethod enum
var AllowedDeinterlaceMethodEnumValues = []DeinterlaceMethod{
	"yadif",
	"bwdif",
}

func (v *DeinterlaceMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DeinterlaceMethod(value)
	for _, existing := range AllowedDeinterlaceMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DeinterlaceMethod", value)
}

// NewDeinterlaceMethodFromValue returns a pointer to a valid DeinterlaceMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDeinterlaceMethodFromValue(v string) (*DeinterlaceMethod, error) {
	ev := DeinterlaceMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DeinterlaceMethod: valid values are %v", v, AllowedDeinterlaceMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DeinterlaceMethod) IsValid() bool {
	for _, existing := range AllowedDeinterlaceMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DeinterlaceMethod value
func (v DeinterlaceMethod) Ptr() *DeinterlaceMethod {
	return &v
}

type NullableDeinterlaceMethod struct {
	value *DeinterlaceMethod
	isSet bool
}

func (v NullableDeinterlaceMethod) Get() *DeinterlaceMethod {
	return v.value
}

func (v *NullableDeinterlaceMethod) Set(val *DeinterlaceMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableDeinterlaceMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableDeinterlaceMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeinterlaceMethod(val *DeinterlaceMethod) *NullableDeinterlaceMethod {
	return &NullableDeinterlaceMethod{value: val, isSet: true}
}

func (v NullableDeinterlaceMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeinterlaceMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

