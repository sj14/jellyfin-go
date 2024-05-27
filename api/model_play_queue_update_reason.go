/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// PlayQueueUpdateReason Enum PlayQueueUpdateReason.
type PlayQueueUpdateReason string

// List of PlayQueueUpdateReason
const (
	PLAYQUEUEUPDATEREASON_NEW_PLAYLIST PlayQueueUpdateReason = "NewPlaylist"
	PLAYQUEUEUPDATEREASON_SET_CURRENT_ITEM PlayQueueUpdateReason = "SetCurrentItem"
	PLAYQUEUEUPDATEREASON_REMOVE_ITEMS PlayQueueUpdateReason = "RemoveItems"
	PLAYQUEUEUPDATEREASON_MOVE_ITEM PlayQueueUpdateReason = "MoveItem"
	PLAYQUEUEUPDATEREASON_QUEUE PlayQueueUpdateReason = "Queue"
	PLAYQUEUEUPDATEREASON_QUEUE_NEXT PlayQueueUpdateReason = "QueueNext"
	PLAYQUEUEUPDATEREASON_NEXT_ITEM PlayQueueUpdateReason = "NextItem"
	PLAYQUEUEUPDATEREASON_PREVIOUS_ITEM PlayQueueUpdateReason = "PreviousItem"
	PLAYQUEUEUPDATEREASON_REPEAT_MODE PlayQueueUpdateReason = "RepeatMode"
	PLAYQUEUEUPDATEREASON_SHUFFLE_MODE PlayQueueUpdateReason = "ShuffleMode"
)

// All allowed values of PlayQueueUpdateReason enum
var AllowedPlayQueueUpdateReasonEnumValues = []PlayQueueUpdateReason{
	"NewPlaylist",
	"SetCurrentItem",
	"RemoveItems",
	"MoveItem",
	"Queue",
	"QueueNext",
	"NextItem",
	"PreviousItem",
	"RepeatMode",
	"ShuffleMode",
}

func (v *PlayQueueUpdateReason) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlayQueueUpdateReason(value)
	for _, existing := range AllowedPlayQueueUpdateReasonEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlayQueueUpdateReason", value)
}

// NewPlayQueueUpdateReasonFromValue returns a pointer to a valid PlayQueueUpdateReason
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlayQueueUpdateReasonFromValue(v string) (*PlayQueueUpdateReason, error) {
	ev := PlayQueueUpdateReason(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlayQueueUpdateReason: valid values are %v", v, AllowedPlayQueueUpdateReasonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlayQueueUpdateReason) IsValid() bool {
	for _, existing := range AllowedPlayQueueUpdateReasonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlayQueueUpdateReason value
func (v PlayQueueUpdateReason) Ptr() *PlayQueueUpdateReason {
	return &v
}

type NullablePlayQueueUpdateReason struct {
	value *PlayQueueUpdateReason
	isSet bool
}

func (v NullablePlayQueueUpdateReason) Get() *PlayQueueUpdateReason {
	return v.value
}

func (v *NullablePlayQueueUpdateReason) Set(val *PlayQueueUpdateReason) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayQueueUpdateReason) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayQueueUpdateReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayQueueUpdateReason(val *PlayQueueUpdateReason) *NullablePlayQueueUpdateReason {
	return &NullablePlayQueueUpdateReason{value: val, isSet: true}
}

func (v NullablePlayQueueUpdateReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayQueueUpdateReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

