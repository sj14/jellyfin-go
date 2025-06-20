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

// SyncPlayUserAccessType Enum SyncPlayUserAccessType.
type SyncPlayUserAccessType string

// List of SyncPlayUserAccessType
const (
	SYNCPLAYUSERACCESSTYPE_CREATE_AND_JOIN_GROUPS SyncPlayUserAccessType = "CreateAndJoinGroups"
	SYNCPLAYUSERACCESSTYPE_JOIN_GROUPS SyncPlayUserAccessType = "JoinGroups"
	SYNCPLAYUSERACCESSTYPE_NONE SyncPlayUserAccessType = "None"
)

// All allowed values of SyncPlayUserAccessType enum
var AllowedSyncPlayUserAccessTypeEnumValues = []SyncPlayUserAccessType{
	"CreateAndJoinGroups",
	"JoinGroups",
	"None",
}

func (v *SyncPlayUserAccessType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SyncPlayUserAccessType(value)
	for _, existing := range AllowedSyncPlayUserAccessTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SyncPlayUserAccessType", value)
}

// NewSyncPlayUserAccessTypeFromValue returns a pointer to a valid SyncPlayUserAccessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSyncPlayUserAccessTypeFromValue(v string) (*SyncPlayUserAccessType, error) {
	ev := SyncPlayUserAccessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SyncPlayUserAccessType: valid values are %v", v, AllowedSyncPlayUserAccessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SyncPlayUserAccessType) IsValid() bool {
	for _, existing := range AllowedSyncPlayUserAccessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SyncPlayUserAccessType value
func (v SyncPlayUserAccessType) Ptr() *SyncPlayUserAccessType {
	return &v
}

type NullableSyncPlayUserAccessType struct {
	value *SyncPlayUserAccessType
	isSet bool
}

func (v NullableSyncPlayUserAccessType) Get() *SyncPlayUserAccessType {
	return v.value
}

func (v *NullableSyncPlayUserAccessType) Set(val *SyncPlayUserAccessType) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncPlayUserAccessType) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncPlayUserAccessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncPlayUserAccessType(val *SyncPlayUserAccessType) *NullableSyncPlayUserAccessType {
	return &NullableSyncPlayUserAccessType{value: val, isSet: true}
}

func (v NullableSyncPlayUserAccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncPlayUserAccessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

