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

// RecordingStatus the model 'RecordingStatus'
type RecordingStatus string

// List of RecordingStatus
const (
	RECORDINGSTATUS_NEW RecordingStatus = "New"
	RECORDINGSTATUS_IN_PROGRESS RecordingStatus = "InProgress"
	RECORDINGSTATUS_COMPLETED RecordingStatus = "Completed"
	RECORDINGSTATUS_CANCELLED RecordingStatus = "Cancelled"
	RECORDINGSTATUS_CONFLICTED_OK RecordingStatus = "ConflictedOk"
	RECORDINGSTATUS_CONFLICTED_NOT_OK RecordingStatus = "ConflictedNotOk"
	RECORDINGSTATUS_ERROR RecordingStatus = "Error"
)

// All allowed values of RecordingStatus enum
var AllowedRecordingStatusEnumValues = []RecordingStatus{
	"New",
	"InProgress",
	"Completed",
	"Cancelled",
	"ConflictedOk",
	"ConflictedNotOk",
	"Error",
}

func (v *RecordingStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecordingStatus(value)
	for _, existing := range AllowedRecordingStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecordingStatus", value)
}

// NewRecordingStatusFromValue returns a pointer to a valid RecordingStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecordingStatusFromValue(v string) (*RecordingStatus, error) {
	ev := RecordingStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecordingStatus: valid values are %v", v, AllowedRecordingStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecordingStatus) IsValid() bool {
	for _, existing := range AllowedRecordingStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecordingStatus value
func (v RecordingStatus) Ptr() *RecordingStatus {
	return &v
}

type NullableRecordingStatus struct {
	value *RecordingStatus
	isSet bool
}

func (v NullableRecordingStatus) Get() *RecordingStatus {
	return v.value
}

func (v *NullableRecordingStatus) Set(val *RecordingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordingStatus(val *RecordingStatus) *NullableRecordingStatus {
	return &NullableRecordingStatus{value: val, isSet: true}
}

func (v NullableRecordingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

