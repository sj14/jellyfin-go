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

// checks if the TimerInfoDtoQueryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimerInfoDtoQueryResult{}

// TimerInfoDtoQueryResult struct for TimerInfoDtoQueryResult
type TimerInfoDtoQueryResult struct {
	// Gets or sets the items.
	Items []TimerInfoDto `json:"Items,omitempty"`
	// Gets or sets the total number of records available.
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty"`
	// Gets or sets the index of the first record in Items.
	StartIndex *int32 `json:"StartIndex,omitempty"`
}

// NewTimerInfoDtoQueryResult instantiates a new TimerInfoDtoQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimerInfoDtoQueryResult() *TimerInfoDtoQueryResult {
	this := TimerInfoDtoQueryResult{}
	return &this
}

// NewTimerInfoDtoQueryResultWithDefaults instantiates a new TimerInfoDtoQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimerInfoDtoQueryResultWithDefaults() *TimerInfoDtoQueryResult {
	this := TimerInfoDtoQueryResult{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimerInfoDtoQueryResult) GetItems() []TimerInfoDto {
	if o == nil {
		var ret []TimerInfoDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimerInfoDtoQueryResult) GetItemsOk() ([]TimerInfoDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *TimerInfoDtoQueryResult) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []TimerInfoDto and assigns it to the Items field.
func (o *TimerInfoDtoQueryResult) SetItems(v []TimerInfoDto) {
	o.Items = v
}

// GetTotalRecordCount returns the TotalRecordCount field value if set, zero value otherwise.
func (o *TimerInfoDtoQueryResult) GetTotalRecordCount() int32 {
	if o == nil || IsNil(o.TotalRecordCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordCount
}

// GetTotalRecordCountOk returns a tuple with the TotalRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerInfoDtoQueryResult) GetTotalRecordCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordCount) {
		return nil, false
	}
	return o.TotalRecordCount, true
}

// HasTotalRecordCount returns a boolean if a field has been set.
func (o *TimerInfoDtoQueryResult) HasTotalRecordCount() bool {
	if o != nil && !IsNil(o.TotalRecordCount) {
		return true
	}

	return false
}

// SetTotalRecordCount gets a reference to the given int32 and assigns it to the TotalRecordCount field.
func (o *TimerInfoDtoQueryResult) SetTotalRecordCount(v int32) {
	o.TotalRecordCount = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *TimerInfoDtoQueryResult) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex) {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerInfoDtoQueryResult) GetStartIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.StartIndex) {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *TimerInfoDtoQueryResult) HasStartIndex() bool {
	if o != nil && !IsNil(o.StartIndex) {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *TimerInfoDtoQueryResult) SetStartIndex(v int32) {
	o.StartIndex = &v
}

func (o TimerInfoDtoQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimerInfoDtoQueryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["Items"] = o.Items
	}
	if !IsNil(o.TotalRecordCount) {
		toSerialize["TotalRecordCount"] = o.TotalRecordCount
	}
	if !IsNil(o.StartIndex) {
		toSerialize["StartIndex"] = o.StartIndex
	}
	return toSerialize, nil
}

type NullableTimerInfoDtoQueryResult struct {
	value *TimerInfoDtoQueryResult
	isSet bool
}

func (v NullableTimerInfoDtoQueryResult) Get() *TimerInfoDtoQueryResult {
	return v.value
}

func (v *NullableTimerInfoDtoQueryResult) Set(val *TimerInfoDtoQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTimerInfoDtoQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTimerInfoDtoQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimerInfoDtoQueryResult(val *TimerInfoDtoQueryResult) *NullableTimerInfoDtoQueryResult {
	return &NullableTimerInfoDtoQueryResult{value: val, isSet: true}
}

func (v NullableTimerInfoDtoQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimerInfoDtoQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


