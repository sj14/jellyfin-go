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

// FileSystemEntryType Enum FileSystemEntryType.
type FileSystemEntryType string

// List of FileSystemEntryType
const (
	FILESYSTEMENTRYTYPE_FILE FileSystemEntryType = "File"
	FILESYSTEMENTRYTYPE_DIRECTORY FileSystemEntryType = "Directory"
	FILESYSTEMENTRYTYPE_NETWORK_COMPUTER FileSystemEntryType = "NetworkComputer"
	FILESYSTEMENTRYTYPE_NETWORK_SHARE FileSystemEntryType = "NetworkShare"
)

// All allowed values of FileSystemEntryType enum
var AllowedFileSystemEntryTypeEnumValues = []FileSystemEntryType{
	"File",
	"Directory",
	"NetworkComputer",
	"NetworkShare",
}

func (v *FileSystemEntryType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FileSystemEntryType(value)
	for _, existing := range AllowedFileSystemEntryTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FileSystemEntryType", value)
}

// NewFileSystemEntryTypeFromValue returns a pointer to a valid FileSystemEntryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileSystemEntryTypeFromValue(v string) (*FileSystemEntryType, error) {
	ev := FileSystemEntryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileSystemEntryType: valid values are %v", v, AllowedFileSystemEntryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileSystemEntryType) IsValid() bool {
	for _, existing := range AllowedFileSystemEntryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FileSystemEntryType value
func (v FileSystemEntryType) Ptr() *FileSystemEntryType {
	return &v
}

type NullableFileSystemEntryType struct {
	value *FileSystemEntryType
	isSet bool
}

func (v NullableFileSystemEntryType) Get() *FileSystemEntryType {
	return v.value
}

func (v *NullableFileSystemEntryType) Set(val *FileSystemEntryType) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSystemEntryType) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSystemEntryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSystemEntryType(val *FileSystemEntryType) *NullableFileSystemEntryType {
	return &NullableFileSystemEntryType{value: val, isSet: true}
}

func (v NullableFileSystemEntryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSystemEntryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

