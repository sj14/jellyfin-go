/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ProfileConditionValue the model 'ProfileConditionValue'
type ProfileConditionValue string

// List of ProfileConditionValue
const (
	PROFILECONDITIONVALUE_AUDIO_CHANNELS ProfileConditionValue = "AudioChannels"
	PROFILECONDITIONVALUE_AUDIO_BITRATE ProfileConditionValue = "AudioBitrate"
	PROFILECONDITIONVALUE_AUDIO_PROFILE ProfileConditionValue = "AudioProfile"
	PROFILECONDITIONVALUE_WIDTH ProfileConditionValue = "Width"
	PROFILECONDITIONVALUE_HEIGHT ProfileConditionValue = "Height"
	PROFILECONDITIONVALUE_HAS64_BIT_OFFSETS ProfileConditionValue = "Has64BitOffsets"
	PROFILECONDITIONVALUE_PACKET_LENGTH ProfileConditionValue = "PacketLength"
	PROFILECONDITIONVALUE_VIDEO_BIT_DEPTH ProfileConditionValue = "VideoBitDepth"
	PROFILECONDITIONVALUE_VIDEO_BITRATE ProfileConditionValue = "VideoBitrate"
	PROFILECONDITIONVALUE_VIDEO_FRAMERATE ProfileConditionValue = "VideoFramerate"
	PROFILECONDITIONVALUE_VIDEO_LEVEL ProfileConditionValue = "VideoLevel"
	PROFILECONDITIONVALUE_VIDEO_PROFILE ProfileConditionValue = "VideoProfile"
	PROFILECONDITIONVALUE_VIDEO_TIMESTAMP ProfileConditionValue = "VideoTimestamp"
	PROFILECONDITIONVALUE_IS_ANAMORPHIC ProfileConditionValue = "IsAnamorphic"
	PROFILECONDITIONVALUE_REF_FRAMES ProfileConditionValue = "RefFrames"
	PROFILECONDITIONVALUE_NUM_AUDIO_STREAMS ProfileConditionValue = "NumAudioStreams"
	PROFILECONDITIONVALUE_NUM_VIDEO_STREAMS ProfileConditionValue = "NumVideoStreams"
	PROFILECONDITIONVALUE_IS_SECONDARY_AUDIO ProfileConditionValue = "IsSecondaryAudio"
	PROFILECONDITIONVALUE_VIDEO_CODEC_TAG ProfileConditionValue = "VideoCodecTag"
	PROFILECONDITIONVALUE_IS_AVC ProfileConditionValue = "IsAvc"
	PROFILECONDITIONVALUE_IS_INTERLACED ProfileConditionValue = "IsInterlaced"
	PROFILECONDITIONVALUE_AUDIO_SAMPLE_RATE ProfileConditionValue = "AudioSampleRate"
	PROFILECONDITIONVALUE_AUDIO_BIT_DEPTH ProfileConditionValue = "AudioBitDepth"
	PROFILECONDITIONVALUE_VIDEO_RANGE_TYPE ProfileConditionValue = "VideoRangeType"
)

// All allowed values of ProfileConditionValue enum
var AllowedProfileConditionValueEnumValues = []ProfileConditionValue{
	"AudioChannels",
	"AudioBitrate",
	"AudioProfile",
	"Width",
	"Height",
	"Has64BitOffsets",
	"PacketLength",
	"VideoBitDepth",
	"VideoBitrate",
	"VideoFramerate",
	"VideoLevel",
	"VideoProfile",
	"VideoTimestamp",
	"IsAnamorphic",
	"RefFrames",
	"NumAudioStreams",
	"NumVideoStreams",
	"IsSecondaryAudio",
	"VideoCodecTag",
	"IsAvc",
	"IsInterlaced",
	"AudioSampleRate",
	"AudioBitDepth",
	"VideoRangeType",
}

func (v *ProfileConditionValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProfileConditionValue(value)
	for _, existing := range AllowedProfileConditionValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProfileConditionValue", value)
}

// NewProfileConditionValueFromValue returns a pointer to a valid ProfileConditionValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProfileConditionValueFromValue(v string) (*ProfileConditionValue, error) {
	ev := ProfileConditionValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProfileConditionValue: valid values are %v", v, AllowedProfileConditionValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProfileConditionValue) IsValid() bool {
	for _, existing := range AllowedProfileConditionValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProfileConditionValue value
func (v ProfileConditionValue) Ptr() *ProfileConditionValue {
	return &v
}

type NullableProfileConditionValue struct {
	value *ProfileConditionValue
	isSet bool
}

func (v NullableProfileConditionValue) Get() *ProfileConditionValue {
	return v.value
}

func (v *NullableProfileConditionValue) Set(val *ProfileConditionValue) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileConditionValue) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileConditionValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileConditionValue(val *ProfileConditionValue) *NullableProfileConditionValue {
	return &NullableProfileConditionValue{value: val, isSet: true}
}

func (v NullableProfileConditionValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileConditionValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

