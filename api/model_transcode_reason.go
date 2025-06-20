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

// TranscodeReason the model 'TranscodeReason'
type TranscodeReason string

// List of TranscodeReason
const (
	TRANSCODEREASON_CONTAINER_NOT_SUPPORTED TranscodeReason = "ContainerNotSupported"
	TRANSCODEREASON_VIDEO_CODEC_NOT_SUPPORTED TranscodeReason = "VideoCodecNotSupported"
	TRANSCODEREASON_AUDIO_CODEC_NOT_SUPPORTED TranscodeReason = "AudioCodecNotSupported"
	TRANSCODEREASON_SUBTITLE_CODEC_NOT_SUPPORTED TranscodeReason = "SubtitleCodecNotSupported"
	TRANSCODEREASON_AUDIO_IS_EXTERNAL TranscodeReason = "AudioIsExternal"
	TRANSCODEREASON_SECONDARY_AUDIO_NOT_SUPPORTED TranscodeReason = "SecondaryAudioNotSupported"
	TRANSCODEREASON_VIDEO_PROFILE_NOT_SUPPORTED TranscodeReason = "VideoProfileNotSupported"
	TRANSCODEREASON_VIDEO_LEVEL_NOT_SUPPORTED TranscodeReason = "VideoLevelNotSupported"
	TRANSCODEREASON_VIDEO_RESOLUTION_NOT_SUPPORTED TranscodeReason = "VideoResolutionNotSupported"
	TRANSCODEREASON_VIDEO_BIT_DEPTH_NOT_SUPPORTED TranscodeReason = "VideoBitDepthNotSupported"
	TRANSCODEREASON_VIDEO_FRAMERATE_NOT_SUPPORTED TranscodeReason = "VideoFramerateNotSupported"
	TRANSCODEREASON_REF_FRAMES_NOT_SUPPORTED TranscodeReason = "RefFramesNotSupported"
	TRANSCODEREASON_ANAMORPHIC_VIDEO_NOT_SUPPORTED TranscodeReason = "AnamorphicVideoNotSupported"
	TRANSCODEREASON_INTERLACED_VIDEO_NOT_SUPPORTED TranscodeReason = "InterlacedVideoNotSupported"
	TRANSCODEREASON_AUDIO_CHANNELS_NOT_SUPPORTED TranscodeReason = "AudioChannelsNotSupported"
	TRANSCODEREASON_AUDIO_PROFILE_NOT_SUPPORTED TranscodeReason = "AudioProfileNotSupported"
	TRANSCODEREASON_AUDIO_SAMPLE_RATE_NOT_SUPPORTED TranscodeReason = "AudioSampleRateNotSupported"
	TRANSCODEREASON_AUDIO_BIT_DEPTH_NOT_SUPPORTED TranscodeReason = "AudioBitDepthNotSupported"
	TRANSCODEREASON_CONTAINER_BITRATE_EXCEEDS_LIMIT TranscodeReason = "ContainerBitrateExceedsLimit"
	TRANSCODEREASON_VIDEO_BITRATE_NOT_SUPPORTED TranscodeReason = "VideoBitrateNotSupported"
	TRANSCODEREASON_AUDIO_BITRATE_NOT_SUPPORTED TranscodeReason = "AudioBitrateNotSupported"
	TRANSCODEREASON_UNKNOWN_VIDEO_STREAM_INFO TranscodeReason = "UnknownVideoStreamInfo"
	TRANSCODEREASON_UNKNOWN_AUDIO_STREAM_INFO TranscodeReason = "UnknownAudioStreamInfo"
	TRANSCODEREASON_DIRECT_PLAY_ERROR TranscodeReason = "DirectPlayError"
	TRANSCODEREASON_VIDEO_RANGE_TYPE_NOT_SUPPORTED TranscodeReason = "VideoRangeTypeNotSupported"
	TRANSCODEREASON_VIDEO_CODEC_TAG_NOT_SUPPORTED TranscodeReason = "VideoCodecTagNotSupported"
	TRANSCODEREASON_STREAM_COUNT_EXCEEDS_LIMIT TranscodeReason = "StreamCountExceedsLimit"
)

// All allowed values of TranscodeReason enum
var AllowedTranscodeReasonEnumValues = []TranscodeReason{
	"ContainerNotSupported",
	"VideoCodecNotSupported",
	"AudioCodecNotSupported",
	"SubtitleCodecNotSupported",
	"AudioIsExternal",
	"SecondaryAudioNotSupported",
	"VideoProfileNotSupported",
	"VideoLevelNotSupported",
	"VideoResolutionNotSupported",
	"VideoBitDepthNotSupported",
	"VideoFramerateNotSupported",
	"RefFramesNotSupported",
	"AnamorphicVideoNotSupported",
	"InterlacedVideoNotSupported",
	"AudioChannelsNotSupported",
	"AudioProfileNotSupported",
	"AudioSampleRateNotSupported",
	"AudioBitDepthNotSupported",
	"ContainerBitrateExceedsLimit",
	"VideoBitrateNotSupported",
	"AudioBitrateNotSupported",
	"UnknownVideoStreamInfo",
	"UnknownAudioStreamInfo",
	"DirectPlayError",
	"VideoRangeTypeNotSupported",
	"VideoCodecTagNotSupported",
	"StreamCountExceedsLimit",
}

func (v *TranscodeReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TranscodeReason(value)
	for _, existing := range AllowedTranscodeReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TranscodeReason", value)
}

// NewTranscodeReasonFromValue returns a pointer to a valid TranscodeReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTranscodeReasonFromValue(v string) (*TranscodeReason, error) {
	ev := TranscodeReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TranscodeReason: valid values are %v", v, AllowedTranscodeReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TranscodeReason) IsValid() bool {
	for _, existing := range AllowedTranscodeReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TranscodeReason value
func (v TranscodeReason) Ptr() *TranscodeReason {
	return &v
}

type NullableTranscodeReason struct {
	value *TranscodeReason
	isSet bool
}

func (v NullableTranscodeReason) Get() *TranscodeReason {
	return v.value
}

func (v *NullableTranscodeReason) Set(val *TranscodeReason) {
	v.value = val
	v.isSet = true
}

func (v NullableTranscodeReason) IsSet() bool {
	return v.isSet
}

func (v *NullableTranscodeReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranscodeReason(val *TranscodeReason) *NullableTranscodeReason {
	return &NullableTranscodeReason{value: val, isSet: true}
}

func (v NullableTranscodeReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranscodeReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

