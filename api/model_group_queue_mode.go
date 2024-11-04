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

// GroupQueueMode Enum GroupQueueMode.
type GroupQueueMode string

// List of GroupQueueMode
const (
	GROUPQUEUEMODE_QUEUE GroupQueueMode = "Queue"
	GROUPQUEUEMODE_QUEUE_NEXT GroupQueueMode = "QueueNext"
)

// All allowed values of GroupQueueMode enum
var AllowedGroupQueueModeEnumValues = []GroupQueueMode{
	"Queue",
	"QueueNext",
}

func (v *GroupQueueMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupQueueMode(value)
	for _, existing := range AllowedGroupQueueModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupQueueMode", value)
}

// NewGroupQueueModeFromValue returns a pointer to a valid GroupQueueMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupQueueModeFromValue(v string) (*GroupQueueMode, error) {
	ev := GroupQueueMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupQueueMode: valid values are %v", v, AllowedGroupQueueModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupQueueMode) IsValid() bool {
	for _, existing := range AllowedGroupQueueModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupQueueMode value
func (v GroupQueueMode) Ptr() *GroupQueueMode {
	return &v
}

type NullableGroupQueueMode struct {
	value *GroupQueueMode
	isSet bool
}

func (v NullableGroupQueueMode) Get() *GroupQueueMode {
	return v.value
}

func (v *NullableGroupQueueMode) Set(val *GroupQueueMode) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupQueueMode) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupQueueMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupQueueMode(val *GroupQueueMode) *NullableGroupQueueMode {
	return &NullableGroupQueueMode{value: val, isSet: true}
}

func (v NullableGroupQueueMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupQueueMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

