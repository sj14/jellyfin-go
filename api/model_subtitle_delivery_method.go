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

// SubtitleDeliveryMethod Delivery method to use during playback of a specific subtitle format.
type SubtitleDeliveryMethod string

// List of SubtitleDeliveryMethod
const (
	SUBTITLEDELIVERYMETHOD_ENCODE SubtitleDeliveryMethod = "Encode"
	SUBTITLEDELIVERYMETHOD_EMBED SubtitleDeliveryMethod = "Embed"
	SUBTITLEDELIVERYMETHOD_EXTERNAL SubtitleDeliveryMethod = "External"
	SUBTITLEDELIVERYMETHOD_HLS SubtitleDeliveryMethod = "Hls"
	SUBTITLEDELIVERYMETHOD_DROP SubtitleDeliveryMethod = "Drop"
)

// All allowed values of SubtitleDeliveryMethod enum
var AllowedSubtitleDeliveryMethodEnumValues = []SubtitleDeliveryMethod{
	"Encode",
	"Embed",
	"External",
	"Hls",
	"Drop",
}

func (v *SubtitleDeliveryMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubtitleDeliveryMethod(value)
	for _, existing := range AllowedSubtitleDeliveryMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubtitleDeliveryMethod", value)
}

// NewSubtitleDeliveryMethodFromValue returns a pointer to a valid SubtitleDeliveryMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubtitleDeliveryMethodFromValue(v string) (*SubtitleDeliveryMethod, error) {
	ev := SubtitleDeliveryMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubtitleDeliveryMethod: valid values are %v", v, AllowedSubtitleDeliveryMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubtitleDeliveryMethod) IsValid() bool {
	for _, existing := range AllowedSubtitleDeliveryMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubtitleDeliveryMethod value
func (v SubtitleDeliveryMethod) Ptr() *SubtitleDeliveryMethod {
	return &v
}

type NullableSubtitleDeliveryMethod struct {
	value *SubtitleDeliveryMethod
	isSet bool
}

func (v NullableSubtitleDeliveryMethod) Get() *SubtitleDeliveryMethod {
	return v.value
}

func (v *NullableSubtitleDeliveryMethod) Set(val *SubtitleDeliveryMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSubtitleDeliveryMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSubtitleDeliveryMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubtitleDeliveryMethod(val *SubtitleDeliveryMethod) *NullableSubtitleDeliveryMethod {
	return &NullableSubtitleDeliveryMethod{value: val, isSet: true}
}

func (v NullableSubtitleDeliveryMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubtitleDeliveryMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

