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

// TransportStreamTimestamp the model 'TransportStreamTimestamp'
type TransportStreamTimestamp string

// List of TransportStreamTimestamp
const (
	TRANSPORTSTREAMTIMESTAMP_NONE TransportStreamTimestamp = "None"
	TRANSPORTSTREAMTIMESTAMP_ZERO TransportStreamTimestamp = "Zero"
	TRANSPORTSTREAMTIMESTAMP_VALID TransportStreamTimestamp = "Valid"
)

// All allowed values of TransportStreamTimestamp enum
var AllowedTransportStreamTimestampEnumValues = []TransportStreamTimestamp{
	"None",
	"Zero",
	"Valid",
}

func (v *TransportStreamTimestamp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransportStreamTimestamp(value)
	for _, existing := range AllowedTransportStreamTimestampEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransportStreamTimestamp", value)
}

// NewTransportStreamTimestampFromValue returns a pointer to a valid TransportStreamTimestamp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTransportStreamTimestampFromValue(v string) (*TransportStreamTimestamp, error) {
	ev := TransportStreamTimestamp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TransportStreamTimestamp: valid values are %v", v, AllowedTransportStreamTimestampEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TransportStreamTimestamp) IsValid() bool {
	for _, existing := range AllowedTransportStreamTimestampEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TransportStreamTimestamp value
func (v TransportStreamTimestamp) Ptr() *TransportStreamTimestamp {
	return &v
}

type NullableTransportStreamTimestamp struct {
	value *TransportStreamTimestamp
	isSet bool
}

func (v NullableTransportStreamTimestamp) Get() *TransportStreamTimestamp {
	return v.value
}

func (v *NullableTransportStreamTimestamp) Set(val *TransportStreamTimestamp) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportStreamTimestamp) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportStreamTimestamp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportStreamTimestamp(val *TransportStreamTimestamp) *NullableTransportStreamTimestamp {
	return &NullableTransportStreamTimestamp{value: val, isSet: true}
}

func (v NullableTransportStreamTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportStreamTimestamp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

