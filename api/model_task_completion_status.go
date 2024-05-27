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

// TaskCompletionStatus Enum TaskCompletionStatus.
type TaskCompletionStatus string

// List of TaskCompletionStatus
const (
	TASKCOMPLETIONSTATUS_COMPLETED TaskCompletionStatus = "Completed"
	TASKCOMPLETIONSTATUS_FAILED TaskCompletionStatus = "Failed"
	TASKCOMPLETIONSTATUS_CANCELLED TaskCompletionStatus = "Cancelled"
	TASKCOMPLETIONSTATUS_ABORTED TaskCompletionStatus = "Aborted"
)

// All allowed values of TaskCompletionStatus enum
var AllowedTaskCompletionStatusEnumValues = []TaskCompletionStatus{
	"Completed",
	"Failed",
	"Cancelled",
	"Aborted",
}

func (v *TaskCompletionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TaskCompletionStatus(value)
	for _, existing := range AllowedTaskCompletionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TaskCompletionStatus", value)
}

// NewTaskCompletionStatusFromValue returns a pointer to a valid TaskCompletionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTaskCompletionStatusFromValue(v string) (*TaskCompletionStatus, error) {
	ev := TaskCompletionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TaskCompletionStatus: valid values are %v", v, AllowedTaskCompletionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TaskCompletionStatus) IsValid() bool {
	for _, existing := range AllowedTaskCompletionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TaskCompletionStatus value
func (v TaskCompletionStatus) Ptr() *TaskCompletionStatus {
	return &v
}

type NullableTaskCompletionStatus struct {
	value *TaskCompletionStatus
	isSet bool
}

func (v NullableTaskCompletionStatus) Get() *TaskCompletionStatus {
	return v.value
}

func (v *NullableTaskCompletionStatus) Set(val *TaskCompletionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskCompletionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskCompletionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskCompletionStatus(val *TaskCompletionStatus) *NullableTaskCompletionStatus {
	return &NullableTaskCompletionStatus{value: val, isSet: true}
}

func (v NullableTaskCompletionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskCompletionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

