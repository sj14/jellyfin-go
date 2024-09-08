/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ProfileConditionType the model 'ProfileConditionType'
type ProfileConditionType string

// List of ProfileConditionType
const (
	PROFILECONDITIONTYPE_EQUALS ProfileConditionType = "Equals"
	PROFILECONDITIONTYPE_NOT_EQUALS ProfileConditionType = "NotEquals"
	PROFILECONDITIONTYPE_LESS_THAN_EQUAL ProfileConditionType = "LessThanEqual"
	PROFILECONDITIONTYPE_GREATER_THAN_EQUAL ProfileConditionType = "GreaterThanEqual"
	PROFILECONDITIONTYPE_EQUALS_ANY ProfileConditionType = "EqualsAny"
)

// All allowed values of ProfileConditionType enum
var AllowedProfileConditionTypeEnumValues = []ProfileConditionType{
	"Equals",
	"NotEquals",
	"LessThanEqual",
	"GreaterThanEqual",
	"EqualsAny",
}

func (v *ProfileConditionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProfileConditionType(value)
	for _, existing := range AllowedProfileConditionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProfileConditionType", value)
}

// NewProfileConditionTypeFromValue returns a pointer to a valid ProfileConditionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProfileConditionTypeFromValue(v string) (*ProfileConditionType, error) {
	ev := ProfileConditionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProfileConditionType: valid values are %v", v, AllowedProfileConditionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProfileConditionType) IsValid() bool {
	for _, existing := range AllowedProfileConditionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProfileConditionType value
func (v ProfileConditionType) Ptr() *ProfileConditionType {
	return &v
}

type NullableProfileConditionType struct {
	value *ProfileConditionType
	isSet bool
}

func (v NullableProfileConditionType) Get() *ProfileConditionType {
	return v.value
}

func (v *NullableProfileConditionType) Set(val *ProfileConditionType) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileConditionType) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileConditionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileConditionType(val *ProfileConditionType) *NullableProfileConditionType {
	return &NullableProfileConditionType{value: val, isSet: true}
}

func (v NullableProfileConditionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileConditionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

