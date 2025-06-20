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

// ScrollDirection An enum representing the axis that should be scrolled.
type ScrollDirection string

// List of ScrollDirection
const (
	SCROLLDIRECTION_HORIZONTAL ScrollDirection = "Horizontal"
	SCROLLDIRECTION_VERTICAL ScrollDirection = "Vertical"
)

// All allowed values of ScrollDirection enum
var AllowedScrollDirectionEnumValues = []ScrollDirection{
	"Horizontal",
	"Vertical",
}

func (v *ScrollDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ScrollDirection(value)
	for _, existing := range AllowedScrollDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ScrollDirection", value)
}

// NewScrollDirectionFromValue returns a pointer to a valid ScrollDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewScrollDirectionFromValue(v string) (*ScrollDirection, error) {
	ev := ScrollDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ScrollDirection: valid values are %v", v, AllowedScrollDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ScrollDirection) IsValid() bool {
	for _, existing := range AllowedScrollDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ScrollDirection value
func (v ScrollDirection) Ptr() *ScrollDirection {
	return &v
}

type NullableScrollDirection struct {
	value *ScrollDirection
	isSet bool
}

func (v NullableScrollDirection) Get() *ScrollDirection {
	return v.value
}

func (v *NullableScrollDirection) Set(val *ScrollDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableScrollDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableScrollDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScrollDirection(val *ScrollDirection) *NullableScrollDirection {
	return &NullableScrollDirection{value: val, isSet: true}
}

func (v NullableScrollDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScrollDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

