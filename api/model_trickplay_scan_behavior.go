/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// TrickplayScanBehavior Enum TrickplayScanBehavior.
type TrickplayScanBehavior string

// List of TrickplayScanBehavior
const (
	TRICKPLAYSCANBEHAVIOR_BLOCKING TrickplayScanBehavior = "Blocking"
	TRICKPLAYSCANBEHAVIOR_NON_BLOCKING TrickplayScanBehavior = "NonBlocking"
)

// All allowed values of TrickplayScanBehavior enum
var AllowedTrickplayScanBehaviorEnumValues = []TrickplayScanBehavior{
	"Blocking",
	"NonBlocking",
}

func (v *TrickplayScanBehavior) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TrickplayScanBehavior(value)
	for _, existing := range AllowedTrickplayScanBehaviorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TrickplayScanBehavior", value)
}

// NewTrickplayScanBehaviorFromValue returns a pointer to a valid TrickplayScanBehavior
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTrickplayScanBehaviorFromValue(v string) (*TrickplayScanBehavior, error) {
	ev := TrickplayScanBehavior(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TrickplayScanBehavior: valid values are %v", v, AllowedTrickplayScanBehaviorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TrickplayScanBehavior) IsValid() bool {
	for _, existing := range AllowedTrickplayScanBehaviorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TrickplayScanBehavior value
func (v TrickplayScanBehavior) Ptr() *TrickplayScanBehavior {
	return &v
}

type NullableTrickplayScanBehavior struct {
	value *TrickplayScanBehavior
	isSet bool
}

func (v NullableTrickplayScanBehavior) Get() *TrickplayScanBehavior {
	return v.value
}

func (v *NullableTrickplayScanBehavior) Set(val *TrickplayScanBehavior) {
	v.value = val
	v.isSet = true
}

func (v NullableTrickplayScanBehavior) IsSet() bool {
	return v.isSet
}

func (v *NullableTrickplayScanBehavior) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrickplayScanBehavior(val *TrickplayScanBehavior) *NullableTrickplayScanBehavior {
	return &NullableTrickplayScanBehavior{value: val, isSet: true}
}

func (v NullableTrickplayScanBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrickplayScanBehavior) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

