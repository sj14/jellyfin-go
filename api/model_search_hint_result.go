/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SearchHintResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchHintResult{}

// SearchHintResult Class SearchHintResult.
type SearchHintResult struct {
	// Gets the search hints.
	SearchHints []SearchHint `json:"SearchHints,omitempty"`
	// Gets the total record count.
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty"`
}

// NewSearchHintResult instantiates a new SearchHintResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchHintResult() *SearchHintResult {
	this := SearchHintResult{}
	return &this
}

// NewSearchHintResultWithDefaults instantiates a new SearchHintResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchHintResultWithDefaults() *SearchHintResult {
	this := SearchHintResult{}
	return &this
}

// GetSearchHints returns the SearchHints field value if set, zero value otherwise.
func (o *SearchHintResult) GetSearchHints() []SearchHint {
	if o == nil || IsNil(o.SearchHints) {
		var ret []SearchHint
		return ret
	}
	return o.SearchHints
}

// GetSearchHintsOk returns a tuple with the SearchHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHintResult) GetSearchHintsOk() ([]SearchHint, bool) {
	if o == nil || IsNil(o.SearchHints) {
		return nil, false
	}
	return o.SearchHints, true
}

// HasSearchHints returns a boolean if a field has been set.
func (o *SearchHintResult) HasSearchHints() bool {
	if o != nil && !IsNil(o.SearchHints) {
		return true
	}

	return false
}

// SetSearchHints gets a reference to the given []SearchHint and assigns it to the SearchHints field.
func (o *SearchHintResult) SetSearchHints(v []SearchHint) {
	o.SearchHints = v
}

// GetTotalRecordCount returns the TotalRecordCount field value if set, zero value otherwise.
func (o *SearchHintResult) GetTotalRecordCount() int32 {
	if o == nil || IsNil(o.TotalRecordCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordCount
}

// GetTotalRecordCountOk returns a tuple with the TotalRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchHintResult) GetTotalRecordCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordCount) {
		return nil, false
	}
	return o.TotalRecordCount, true
}

// HasTotalRecordCount returns a boolean if a field has been set.
func (o *SearchHintResult) HasTotalRecordCount() bool {
	if o != nil && !IsNil(o.TotalRecordCount) {
		return true
	}

	return false
}

// SetTotalRecordCount gets a reference to the given int32 and assigns it to the TotalRecordCount field.
func (o *SearchHintResult) SetTotalRecordCount(v int32) {
	o.TotalRecordCount = &v
}

func (o SearchHintResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchHintResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SearchHints) {
		toSerialize["SearchHints"] = o.SearchHints
	}
	if !IsNil(o.TotalRecordCount) {
		toSerialize["TotalRecordCount"] = o.TotalRecordCount
	}
	return toSerialize, nil
}

type NullableSearchHintResult struct {
	value *SearchHintResult
	isSet bool
}

func (v NullableSearchHintResult) Get() *SearchHintResult {
	return v.value
}

func (v *NullableSearchHintResult) Set(val *SearchHintResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchHintResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchHintResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchHintResult(val *SearchHintResult) *NullableSearchHintResult {
	return &NullableSearchHintResult{value: val, isSet: true}
}

func (v NullableSearchHintResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchHintResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


