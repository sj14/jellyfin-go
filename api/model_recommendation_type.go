/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// RecommendationType the model 'RecommendationType'
type RecommendationType string

// List of RecommendationType
const (
	RECOMMENDATIONTYPE_SIMILAR_TO_RECENTLY_PLAYED RecommendationType = "SimilarToRecentlyPlayed"
	RECOMMENDATIONTYPE_SIMILAR_TO_LIKED_ITEM RecommendationType = "SimilarToLikedItem"
	RECOMMENDATIONTYPE_HAS_DIRECTOR_FROM_RECENTLY_PLAYED RecommendationType = "HasDirectorFromRecentlyPlayed"
	RECOMMENDATIONTYPE_HAS_ACTOR_FROM_RECENTLY_PLAYED RecommendationType = "HasActorFromRecentlyPlayed"
	RECOMMENDATIONTYPE_HAS_LIKED_DIRECTOR RecommendationType = "HasLikedDirector"
	RECOMMENDATIONTYPE_HAS_LIKED_ACTOR RecommendationType = "HasLikedActor"
)

// All allowed values of RecommendationType enum
var AllowedRecommendationTypeEnumValues = []RecommendationType{
	"SimilarToRecentlyPlayed",
	"SimilarToLikedItem",
	"HasDirectorFromRecentlyPlayed",
	"HasActorFromRecentlyPlayed",
	"HasLikedDirector",
	"HasLikedActor",
}

func (v *RecommendationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecommendationType(value)
	for _, existing := range AllowedRecommendationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecommendationType", value)
}

// NewRecommendationTypeFromValue returns a pointer to a valid RecommendationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecommendationTypeFromValue(v string) (*RecommendationType, error) {
	ev := RecommendationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecommendationType: valid values are %v", v, AllowedRecommendationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecommendationType) IsValid() bool {
	for _, existing := range AllowedRecommendationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RecommendationType value
func (v RecommendationType) Ptr() *RecommendationType {
	return &v
}

type NullableRecommendationType struct {
	value *RecommendationType
	isSet bool
}

func (v NullableRecommendationType) Get() *RecommendationType {
	return v.value
}

func (v *NullableRecommendationType) Set(val *RecommendationType) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationType) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationType(val *RecommendationType) *NullableRecommendationType {
	return &NullableRecommendationType{value: val, isSet: true}
}

func (v NullableRecommendationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

