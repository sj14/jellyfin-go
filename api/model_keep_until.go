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

// KeepUntil the model 'KeepUntil'
type KeepUntil string

// List of KeepUntil
const (
	KEEPUNTIL_UNTIL_DELETED KeepUntil = "UntilDeleted"
	KEEPUNTIL_UNTIL_SPACE_NEEDED KeepUntil = "UntilSpaceNeeded"
	KEEPUNTIL_UNTIL_WATCHED KeepUntil = "UntilWatched"
	KEEPUNTIL_UNTIL_DATE KeepUntil = "UntilDate"
)

// All allowed values of KeepUntil enum
var AllowedKeepUntilEnumValues = []KeepUntil{
	"UntilDeleted",
	"UntilSpaceNeeded",
	"UntilWatched",
	"UntilDate",
}

func (v *KeepUntil) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := KeepUntil(value)
	for _, existing := range AllowedKeepUntilEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid KeepUntil", value)
}

// NewKeepUntilFromValue returns a pointer to a valid KeepUntil
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeepUntilFromValue(v string) (*KeepUntil, error) {
	ev := KeepUntil(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeepUntil: valid values are %v", v, AllowedKeepUntilEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeepUntil) IsValid() bool {
	for _, existing := range AllowedKeepUntilEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to KeepUntil value
func (v KeepUntil) Ptr() *KeepUntil {
	return &v
}

type NullableKeepUntil struct {
	value *KeepUntil
	isSet bool
}

func (v NullableKeepUntil) Get() *KeepUntil {
	return v.value
}

func (v *NullableKeepUntil) Set(val *KeepUntil) {
	v.value = val
	v.isSet = true
}

func (v NullableKeepUntil) IsSet() bool {
	return v.isSet
}

func (v *NullableKeepUntil) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeepUntil(val *KeepUntil) *NullableKeepUntil {
	return &NullableKeepUntil{value: val, isSet: true}
}

func (v NullableKeepUntil) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeepUntil) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

