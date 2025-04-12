/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// GroupUpdateType Enum GroupUpdateType.
type GroupUpdateType string

// List of GroupUpdateType
const (
	GROUPUPDATETYPE_USER_JOINED GroupUpdateType = "UserJoined"
	GROUPUPDATETYPE_USER_LEFT GroupUpdateType = "UserLeft"
	GROUPUPDATETYPE_GROUP_JOINED GroupUpdateType = "GroupJoined"
	GROUPUPDATETYPE_GROUP_LEFT GroupUpdateType = "GroupLeft"
	GROUPUPDATETYPE_STATE_UPDATE GroupUpdateType = "StateUpdate"
	GROUPUPDATETYPE_PLAY_QUEUE GroupUpdateType = "PlayQueue"
	GROUPUPDATETYPE_NOT_IN_GROUP GroupUpdateType = "NotInGroup"
	GROUPUPDATETYPE_GROUP_DOES_NOT_EXIST GroupUpdateType = "GroupDoesNotExist"
	GROUPUPDATETYPE_CREATE_GROUP_DENIED GroupUpdateType = "CreateGroupDenied"
	GROUPUPDATETYPE_JOIN_GROUP_DENIED GroupUpdateType = "JoinGroupDenied"
	GROUPUPDATETYPE_LIBRARY_ACCESS_DENIED GroupUpdateType = "LibraryAccessDenied"
)

// All allowed values of GroupUpdateType enum
var AllowedGroupUpdateTypeEnumValues = []GroupUpdateType{
	"UserJoined",
	"UserLeft",
	"GroupJoined",
	"GroupLeft",
	"StateUpdate",
	"PlayQueue",
	"NotInGroup",
	"GroupDoesNotExist",
	"CreateGroupDenied",
	"JoinGroupDenied",
	"LibraryAccessDenied",
}

func (v *GroupUpdateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GroupUpdateType(value)
	for _, existing := range AllowedGroupUpdateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GroupUpdateType", value)
}

// NewGroupUpdateTypeFromValue returns a pointer to a valid GroupUpdateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGroupUpdateTypeFromValue(v string) (*GroupUpdateType, error) {
	ev := GroupUpdateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GroupUpdateType: valid values are %v", v, AllowedGroupUpdateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GroupUpdateType) IsValid() bool {
	for _, existing := range AllowedGroupUpdateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GroupUpdateType value
func (v GroupUpdateType) Ptr() *GroupUpdateType {
	return &v
}

type NullableGroupUpdateType struct {
	value *GroupUpdateType
	isSet bool
}

func (v NullableGroupUpdateType) Get() *GroupUpdateType {
	return v.value
}

func (v *NullableGroupUpdateType) Set(val *GroupUpdateType) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUpdateType) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUpdateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUpdateType(val *GroupUpdateType) *NullableGroupUpdateType {
	return &NullableGroupUpdateType{value: val, isSet: true}
}

func (v NullableGroupUpdateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUpdateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

