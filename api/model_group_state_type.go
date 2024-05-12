/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// GroupStateType Enum GroupState.
type GroupStateType string

// List of GroupStateType
const (
	GROUPSTATETYPE_IDLE GroupStateType = "Idle"
	GROUPSTATETYPE_WAITING GroupStateType = "Waiting"
	GROUPSTATETYPE_PAUSED GroupStateType = "Paused"
	GROUPSTATETYPE_PLAYING GroupStateType = "Playing"
)

// All allowed values of GroupStateType enum
var AllowedGroupStateTypeEnumValues = []GroupStateType{
	"Idle",
	"Waiting",
	"Paused",
	"Playing",
}

func (v *GroupStateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupStateType(value)
	for _, existing := range AllowedGroupStateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupStateType", value)
}

// NewGroupStateTypeFromValue returns a pointer to a valid GroupStateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupStateTypeFromValue(v string) (*GroupStateType, error) {
	ev := GroupStateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupStateType: valid values are %v", v, AllowedGroupStateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupStateType) IsValid() bool {
	for _, existing := range AllowedGroupStateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupStateType value
func (v GroupStateType) Ptr() *GroupStateType {
	return &v
}

type NullableGroupStateType struct {
	value *GroupStateType
	isSet bool
}

func (v NullableGroupStateType) Get() *GroupStateType {
	return v.value
}

func (v *NullableGroupStateType) Set(val *GroupStateType) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupStateType) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupStateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupStateType(val *GroupStateType) *NullableGroupStateType {
	return &NullableGroupStateType{value: val, isSet: true}
}

func (v NullableGroupStateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupStateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

