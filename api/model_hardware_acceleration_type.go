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

// HardwareAccelerationType Enum containing hardware acceleration types.
type HardwareAccelerationType string

// List of HardwareAccelerationType
const (
	HARDWAREACCELERATIONTYPE_NONE HardwareAccelerationType = "none"
	HARDWAREACCELERATIONTYPE_AMF HardwareAccelerationType = "amf"
	HARDWAREACCELERATIONTYPE_QSV HardwareAccelerationType = "qsv"
	HARDWAREACCELERATIONTYPE_NVENC HardwareAccelerationType = "nvenc"
	HARDWAREACCELERATIONTYPE_V4L2M2M HardwareAccelerationType = "v4l2m2m"
	HARDWAREACCELERATIONTYPE_VAAPI HardwareAccelerationType = "vaapi"
	HARDWAREACCELERATIONTYPE_VIDEOTOOLBOX HardwareAccelerationType = "videotoolbox"
	HARDWAREACCELERATIONTYPE_RKMPP HardwareAccelerationType = "rkmpp"
)

// All allowed values of HardwareAccelerationType enum
var AllowedHardwareAccelerationTypeEnumValues = []HardwareAccelerationType{
	"none",
	"amf",
	"qsv",
	"nvenc",
	"v4l2m2m",
	"vaapi",
	"videotoolbox",
	"rkmpp",
}

func (v *HardwareAccelerationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HardwareAccelerationType(value)
	for _, existing := range AllowedHardwareAccelerationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HardwareAccelerationType", value)
}

// NewHardwareAccelerationTypeFromValue returns a pointer to a valid HardwareAccelerationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHardwareAccelerationTypeFromValue(v string) (*HardwareAccelerationType, error) {
	ev := HardwareAccelerationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HardwareAccelerationType: valid values are %v", v, AllowedHardwareAccelerationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HardwareAccelerationType) IsValid() bool {
	for _, existing := range AllowedHardwareAccelerationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HardwareAccelerationType value
func (v HardwareAccelerationType) Ptr() *HardwareAccelerationType {
	return &v
}

type NullableHardwareAccelerationType struct {
	value *HardwareAccelerationType
	isSet bool
}

func (v NullableHardwareAccelerationType) Get() *HardwareAccelerationType {
	return v.value
}

func (v *NullableHardwareAccelerationType) Set(val *HardwareAccelerationType) {
	v.value = val
	v.isSet = true
}

func (v NullableHardwareAccelerationType) IsSet() bool {
	return v.isSet
}

func (v *NullableHardwareAccelerationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHardwareAccelerationType(val *HardwareAccelerationType) *NullableHardwareAccelerationType {
	return &NullableHardwareAccelerationType{value: val, isSet: true}
}

func (v NullableHardwareAccelerationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHardwareAccelerationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

