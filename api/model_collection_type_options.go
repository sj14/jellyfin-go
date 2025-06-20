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

// CollectionTypeOptions The collection type options.
type CollectionTypeOptions string

// List of CollectionTypeOptions
const (
	COLLECTIONTYPEOPTIONS_MOVIES CollectionTypeOptions = "movies"
	COLLECTIONTYPEOPTIONS_TVSHOWS CollectionTypeOptions = "tvshows"
	COLLECTIONTYPEOPTIONS_MUSIC CollectionTypeOptions = "music"
	COLLECTIONTYPEOPTIONS_MUSICVIDEOS CollectionTypeOptions = "musicvideos"
	COLLECTIONTYPEOPTIONS_HOMEVIDEOS CollectionTypeOptions = "homevideos"
	COLLECTIONTYPEOPTIONS_BOXSETS CollectionTypeOptions = "boxsets"
	COLLECTIONTYPEOPTIONS_BOOKS CollectionTypeOptions = "books"
	COLLECTIONTYPEOPTIONS_MIXED CollectionTypeOptions = "mixed"
)

// All allowed values of CollectionTypeOptions enum
var AllowedCollectionTypeOptionsEnumValues = []CollectionTypeOptions{
	"movies",
	"tvshows",
	"music",
	"musicvideos",
	"homevideos",
	"boxsets",
	"books",
	"mixed",
}

func (v *CollectionTypeOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionTypeOptions(value)
	for _, existing := range AllowedCollectionTypeOptionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionTypeOptions", value)
}

// NewCollectionTypeOptionsFromValue returns a pointer to a valid CollectionTypeOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionTypeOptionsFromValue(v string) (*CollectionTypeOptions, error) {
	ev := CollectionTypeOptions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionTypeOptions: valid values are %v", v, AllowedCollectionTypeOptionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionTypeOptions) IsValid() bool {
	for _, existing := range AllowedCollectionTypeOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CollectionTypeOptions value
func (v CollectionTypeOptions) Ptr() *CollectionTypeOptions {
	return &v
}

type NullableCollectionTypeOptions struct {
	value *CollectionTypeOptions
	isSet bool
}

func (v NullableCollectionTypeOptions) Get() *CollectionTypeOptions {
	return v.value
}

func (v *NullableCollectionTypeOptions) Set(val *CollectionTypeOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionTypeOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionTypeOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionTypeOptions(val *CollectionTypeOptions) *NullableCollectionTypeOptions {
	return &NullableCollectionTypeOptions{value: val, isSet: true}
}

func (v NullableCollectionTypeOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionTypeOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

