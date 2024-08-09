/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// HardwareEncodingType Enum HardwareEncodingType.
type HardwareEncodingType string

// List of HardwareEncodingType
const (
	HARDWAREENCODINGTYPE_AMF HardwareEncodingType = "AMF"
	HARDWAREENCODINGTYPE_QSV HardwareEncodingType = "QSV"
	HARDWAREENCODINGTYPE_NVENC HardwareEncodingType = "NVENC"
	HARDWAREENCODINGTYPE_V4_L2_M2_M HardwareEncodingType = "V4L2M2M"
	HARDWAREENCODINGTYPE_VAAPI HardwareEncodingType = "VAAPI"
	HARDWAREENCODINGTYPE_VIDEO_TOOL_BOX HardwareEncodingType = "VideoToolBox"
	HARDWAREENCODINGTYPE_RKMPP HardwareEncodingType = "RKMPP"
)

// All allowed values of HardwareEncodingType enum
var AllowedHardwareEncodingTypeEnumValues = []HardwareEncodingType{
	"AMF",
	"QSV",
	"NVENC",
	"V4L2M2M",
	"VAAPI",
	"VideoToolBox",
	"RKMPP",
}

func (v *HardwareEncodingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HardwareEncodingType(value)
	for _, existing := range AllowedHardwareEncodingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HardwareEncodingType", value)
}

// NewHardwareEncodingTypeFromValue returns a pointer to a valid HardwareEncodingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHardwareEncodingTypeFromValue(v string) (*HardwareEncodingType, error) {
	ev := HardwareEncodingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HardwareEncodingType: valid values are %v", v, AllowedHardwareEncodingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HardwareEncodingType) IsValid() bool {
	for _, existing := range AllowedHardwareEncodingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HardwareEncodingType value
func (v HardwareEncodingType) Ptr() *HardwareEncodingType {
	return &v
}

type NullableHardwareEncodingType struct {
	value *HardwareEncodingType
	isSet bool
}

func (v NullableHardwareEncodingType) Get() *HardwareEncodingType {
	return v.value
}

func (v *NullableHardwareEncodingType) Set(val *HardwareEncodingType) {
	v.value = val
	v.isSet = true
}

func (v NullableHardwareEncodingType) IsSet() bool {
	return v.isSet
}

func (v *NullableHardwareEncodingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHardwareEncodingType(val *HardwareEncodingType) *NullableHardwareEncodingType {
	return &NullableHardwareEncodingType{value: val, isSet: true}
}

func (v NullableHardwareEncodingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHardwareEncodingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

