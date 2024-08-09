/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the QueryFiltersLegacy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryFiltersLegacy{}

// QueryFiltersLegacy struct for QueryFiltersLegacy
type QueryFiltersLegacy struct {
	Genres []string `json:"Genres,omitempty"`
	Tags []string `json:"Tags,omitempty"`
	OfficialRatings []string `json:"OfficialRatings,omitempty"`
	Years []int32 `json:"Years,omitempty"`
}

// NewQueryFiltersLegacy instantiates a new QueryFiltersLegacy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryFiltersLegacy() *QueryFiltersLegacy {
	this := QueryFiltersLegacy{}
	return &this
}

// NewQueryFiltersLegacyWithDefaults instantiates a new QueryFiltersLegacy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryFiltersLegacyWithDefaults() *QueryFiltersLegacy {
	this := QueryFiltersLegacy{}
	return &this
}

// GetGenres returns the Genres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryFiltersLegacy) GetGenres() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryFiltersLegacy) GetGenresOk() ([]string, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *QueryFiltersLegacy) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []string and assigns it to the Genres field.
func (o *QueryFiltersLegacy) SetGenres(v []string) {
	o.Genres = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryFiltersLegacy) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryFiltersLegacy) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *QueryFiltersLegacy) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *QueryFiltersLegacy) SetTags(v []string) {
	o.Tags = v
}

// GetOfficialRatings returns the OfficialRatings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryFiltersLegacy) GetOfficialRatings() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OfficialRatings
}

// GetOfficialRatingsOk returns a tuple with the OfficialRatings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryFiltersLegacy) GetOfficialRatingsOk() ([]string, bool) {
	if o == nil || IsNil(o.OfficialRatings) {
		return nil, false
	}
	return o.OfficialRatings, true
}

// HasOfficialRatings returns a boolean if a field has been set.
func (o *QueryFiltersLegacy) HasOfficialRatings() bool {
	if o != nil && !IsNil(o.OfficialRatings) {
		return true
	}

	return false
}

// SetOfficialRatings gets a reference to the given []string and assigns it to the OfficialRatings field.
func (o *QueryFiltersLegacy) SetOfficialRatings(v []string) {
	o.OfficialRatings = v
}

// GetYears returns the Years field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryFiltersLegacy) GetYears() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Years
}

// GetYearsOk returns a tuple with the Years field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryFiltersLegacy) GetYearsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Years) {
		return nil, false
	}
	return o.Years, true
}

// HasYears returns a boolean if a field has been set.
func (o *QueryFiltersLegacy) HasYears() bool {
	if o != nil && !IsNil(o.Years) {
		return true
	}

	return false
}

// SetYears gets a reference to the given []int32 and assigns it to the Years field.
func (o *QueryFiltersLegacy) SetYears(v []int32) {
	o.Years = v
}

func (o QueryFiltersLegacy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryFiltersLegacy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Genres != nil {
		toSerialize["Genres"] = o.Genres
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	if o.OfficialRatings != nil {
		toSerialize["OfficialRatings"] = o.OfficialRatings
	}
	if o.Years != nil {
		toSerialize["Years"] = o.Years
	}
	return toSerialize, nil
}

type NullableQueryFiltersLegacy struct {
	value *QueryFiltersLegacy
	isSet bool
}

func (v NullableQueryFiltersLegacy) Get() *QueryFiltersLegacy {
	return v.value
}

func (v *NullableQueryFiltersLegacy) Set(val *QueryFiltersLegacy) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryFiltersLegacy) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryFiltersLegacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryFiltersLegacy(val *QueryFiltersLegacy) *NullableQueryFiltersLegacy {
	return &NullableQueryFiltersLegacy{value: val, isSet: true}
}

func (v NullableQueryFiltersLegacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryFiltersLegacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


