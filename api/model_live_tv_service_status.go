/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// LiveTvServiceStatus the model 'LiveTvServiceStatus'
type LiveTvServiceStatus string

// List of LiveTvServiceStatus
const (
	LIVETVSERVICESTATUS_OK LiveTvServiceStatus = "Ok"
	LIVETVSERVICESTATUS_UNAVAILABLE LiveTvServiceStatus = "Unavailable"
)

// All allowed values of LiveTvServiceStatus enum
var AllowedLiveTvServiceStatusEnumValues = []LiveTvServiceStatus{
	"Ok",
	"Unavailable",
}

func (v *LiveTvServiceStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LiveTvServiceStatus(value)
	for _, existing := range AllowedLiveTvServiceStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LiveTvServiceStatus", value)
}

// NewLiveTvServiceStatusFromValue returns a pointer to a valid LiveTvServiceStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLiveTvServiceStatusFromValue(v string) (*LiveTvServiceStatus, error) {
	ev := LiveTvServiceStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LiveTvServiceStatus: valid values are %v", v, AllowedLiveTvServiceStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LiveTvServiceStatus) IsValid() bool {
	for _, existing := range AllowedLiveTvServiceStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LiveTvServiceStatus value
func (v LiveTvServiceStatus) Ptr() *LiveTvServiceStatus {
	return &v
}

type NullableLiveTvServiceStatus struct {
	value *LiveTvServiceStatus
	isSet bool
}

func (v NullableLiveTvServiceStatus) Get() *LiveTvServiceStatus {
	return v.value
}

func (v *NullableLiveTvServiceStatus) Set(val *LiveTvServiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveTvServiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveTvServiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveTvServiceStatus(val *LiveTvServiceStatus) *NullableLiveTvServiceStatus {
	return &NullableLiveTvServiceStatus{value: val, isSet: true}
}

func (v NullableLiveTvServiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveTvServiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

