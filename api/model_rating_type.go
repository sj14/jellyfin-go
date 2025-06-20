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

// RatingType the model 'RatingType'
type RatingType string

// List of RatingType
const (
	RATINGTYPE_SCORE RatingType = "Score"
	RATINGTYPE_LIKES RatingType = "Likes"
)

// All allowed values of RatingType enum
var AllowedRatingTypeEnumValues = []RatingType{
	"Score",
	"Likes",
}

func (v *RatingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RatingType(value)
	for _, existing := range AllowedRatingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RatingType", value)
}

// NewRatingTypeFromValue returns a pointer to a valid RatingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRatingTypeFromValue(v string) (*RatingType, error) {
	ev := RatingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RatingType: valid values are %v", v, AllowedRatingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RatingType) IsValid() bool {
	for _, existing := range AllowedRatingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RatingType value
func (v RatingType) Ptr() *RatingType {
	return &v
}

type NullableRatingType struct {
	value *RatingType
	isSet bool
}

func (v NullableRatingType) Get() *RatingType {
	return v.value
}

func (v *NullableRatingType) Set(val *RatingType) {
	v.value = val
	v.isSet = true
}

func (v NullableRatingType) IsSet() bool {
	return v.isSet
}

func (v *NullableRatingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatingType(val *RatingType) *NullableRatingType {
	return &NullableRatingType{value: val, isSet: true}
}

func (v NullableRatingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

