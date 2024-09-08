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

// ProgramAudio the model 'ProgramAudio'
type ProgramAudio string

// List of ProgramAudio
const (
	PROGRAMAUDIO_MONO ProgramAudio = "Mono"
	PROGRAMAUDIO_STEREO ProgramAudio = "Stereo"
	PROGRAMAUDIO_DOLBY ProgramAudio = "Dolby"
	PROGRAMAUDIO_DOLBY_DIGITAL ProgramAudio = "DolbyDigital"
	PROGRAMAUDIO_THX ProgramAudio = "Thx"
	PROGRAMAUDIO_ATMOS ProgramAudio = "Atmos"
)

// All allowed values of ProgramAudio enum
var AllowedProgramAudioEnumValues = []ProgramAudio{
	"Mono",
	"Stereo",
	"Dolby",
	"DolbyDigital",
	"Thx",
	"Atmos",
}

func (v *ProgramAudio) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProgramAudio(value)
	for _, existing := range AllowedProgramAudioEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProgramAudio", value)
}

// NewProgramAudioFromValue returns a pointer to a valid ProgramAudio
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProgramAudioFromValue(v string) (*ProgramAudio, error) {
	ev := ProgramAudio(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProgramAudio: valid values are %v", v, AllowedProgramAudioEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProgramAudio) IsValid() bool {
	for _, existing := range AllowedProgramAudioEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProgramAudio value
func (v ProgramAudio) Ptr() *ProgramAudio {
	return &v
}

type NullableProgramAudio struct {
	value *ProgramAudio
	isSet bool
}

func (v NullableProgramAudio) Get() *ProgramAudio {
	return v.value
}

func (v *NullableProgramAudio) Set(val *ProgramAudio) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramAudio) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramAudio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramAudio(val *ProgramAudio) *NullableProgramAudio {
	return &NullableProgramAudio{value: val, isSet: true}
}

func (v NullableProgramAudio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramAudio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

