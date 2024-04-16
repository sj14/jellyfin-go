/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.8.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ItemFilter Enum ItemFilter.
type ItemFilter string

// List of ItemFilter
const (
	ITEMFILTER_IS_FOLDER ItemFilter = "IsFolder"
	ITEMFILTER_IS_NOT_FOLDER ItemFilter = "IsNotFolder"
	ITEMFILTER_IS_UNPLAYED ItemFilter = "IsUnplayed"
	ITEMFILTER_IS_PLAYED ItemFilter = "IsPlayed"
	ITEMFILTER_IS_FAVORITE ItemFilter = "IsFavorite"
	ITEMFILTER_IS_RESUMABLE ItemFilter = "IsResumable"
	ITEMFILTER_LIKES ItemFilter = "Likes"
	ITEMFILTER_DISLIKES ItemFilter = "Dislikes"
	ITEMFILTER_IS_FAVORITE_OR_LIKES ItemFilter = "IsFavoriteOrLikes"
)

// All allowed values of ItemFilter enum
var AllowedItemFilterEnumValues = []ItemFilter{
	"IsFolder",
	"IsNotFolder",
	"IsUnplayed",
	"IsPlayed",
	"IsFavorite",
	"IsResumable",
	"Likes",
	"Dislikes",
	"IsFavoriteOrLikes",
}

func (v *ItemFilter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ItemFilter(value)
	for _, existing := range AllowedItemFilterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ItemFilter", value)
}

// NewItemFilterFromValue returns a pointer to a valid ItemFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewItemFilterFromValue(v string) (*ItemFilter, error) {
	ev := ItemFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ItemFilter: valid values are %v", v, AllowedItemFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ItemFilter) IsValid() bool {
	for _, existing := range AllowedItemFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ItemFilter value
func (v ItemFilter) Ptr() *ItemFilter {
	return &v
}

type NullableItemFilter struct {
	value *ItemFilter
	isSet bool
}

func (v NullableItemFilter) Get() *ItemFilter {
	return v.value
}

func (v *NullableItemFilter) Set(val *ItemFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableItemFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableItemFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemFilter(val *ItemFilter) *NullableItemFilter {
	return &NullableItemFilter{value: val, isSet: true}
}

func (v NullableItemFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

