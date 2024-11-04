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

// SeriesStatus The status of a series.
type SeriesStatus string

// List of SeriesStatus
const (
	SERIESSTATUS_CONTINUING SeriesStatus = "Continuing"
	SERIESSTATUS_ENDED SeriesStatus = "Ended"
	SERIESSTATUS_UNRELEASED SeriesStatus = "Unreleased"
)

// All allowed values of SeriesStatus enum
var AllowedSeriesStatusEnumValues = []SeriesStatus{
	"Continuing",
	"Ended",
	"Unreleased",
}

func (v *SeriesStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SeriesStatus(value)
	for _, existing := range AllowedSeriesStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SeriesStatus", value)
}

// NewSeriesStatusFromValue returns a pointer to a valid SeriesStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSeriesStatusFromValue(v string) (*SeriesStatus, error) {
	ev := SeriesStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SeriesStatus: valid values are %v", v, AllowedSeriesStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SeriesStatus) IsValid() bool {
	for _, existing := range AllowedSeriesStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SeriesStatus value
func (v SeriesStatus) Ptr() *SeriesStatus {
	return &v
}

type NullableSeriesStatus struct {
	value *SeriesStatus
	isSet bool
}

func (v NullableSeriesStatus) Get() *SeriesStatus {
	return v.value
}

func (v *NullableSeriesStatus) Set(val *SeriesStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesStatus(val *SeriesStatus) *NullableSeriesStatus {
	return &NullableSeriesStatus{value: val, isSet: true}
}

func (v NullableSeriesStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

