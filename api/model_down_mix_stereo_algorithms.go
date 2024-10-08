/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// DownMixStereoAlgorithms An enum representing an algorithm to downmix 6ch+ to stereo.  Algorithms sourced from https://superuser.com/questions/852400/properly-downmix-5-1-to-stereo-using-ffmpeg/1410620#1410620.
type DownMixStereoAlgorithms string

// List of DownMixStereoAlgorithms
const (
	DOWNMIXSTEREOALGORITHMS_NONE DownMixStereoAlgorithms = "None"
	DOWNMIXSTEREOALGORITHMS_DAVE750 DownMixStereoAlgorithms = "Dave750"
	DOWNMIXSTEREOALGORITHMS_NIGHTMODE_DIALOGUE DownMixStereoAlgorithms = "NightmodeDialogue"
)

// All allowed values of DownMixStereoAlgorithms enum
var AllowedDownMixStereoAlgorithmsEnumValues = []DownMixStereoAlgorithms{
	"None",
	"Dave750",
	"NightmodeDialogue",
}

func (v *DownMixStereoAlgorithms) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DownMixStereoAlgorithms(value)
	for _, existing := range AllowedDownMixStereoAlgorithmsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DownMixStereoAlgorithms", value)
}

// NewDownMixStereoAlgorithmsFromValue returns a pointer to a valid DownMixStereoAlgorithms
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDownMixStereoAlgorithmsFromValue(v string) (*DownMixStereoAlgorithms, error) {
	ev := DownMixStereoAlgorithms(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DownMixStereoAlgorithms: valid values are %v", v, AllowedDownMixStereoAlgorithmsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DownMixStereoAlgorithms) IsValid() bool {
	for _, existing := range AllowedDownMixStereoAlgorithmsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DownMixStereoAlgorithms value
func (v DownMixStereoAlgorithms) Ptr() *DownMixStereoAlgorithms {
	return &v
}

type NullableDownMixStereoAlgorithms struct {
	value *DownMixStereoAlgorithms
	isSet bool
}

func (v NullableDownMixStereoAlgorithms) Get() *DownMixStereoAlgorithms {
	return v.value
}

func (v *NullableDownMixStereoAlgorithms) Set(val *DownMixStereoAlgorithms) {
	v.value = val
	v.isSet = true
}

func (v NullableDownMixStereoAlgorithms) IsSet() bool {
	return v.isSet
}

func (v *NullableDownMixStereoAlgorithms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDownMixStereoAlgorithms(val *DownMixStereoAlgorithms) *NullableDownMixStereoAlgorithms {
	return &NullableDownMixStereoAlgorithms{value: val, isSet: true}
}

func (v NullableDownMixStereoAlgorithms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDownMixStereoAlgorithms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

