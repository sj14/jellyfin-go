/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SeriesTimerInfoDtoQueryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesTimerInfoDtoQueryResult{}

// SeriesTimerInfoDtoQueryResult Query result container.
type SeriesTimerInfoDtoQueryResult struct {
	// Gets or sets the items.
	Items []SeriesTimerInfoDto `json:"Items,omitempty"`
	// Gets or sets the total number of records available.
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty"`
	// Gets or sets the index of the first record in Items.
	StartIndex *int32 `json:"StartIndex,omitempty"`
}

// NewSeriesTimerInfoDtoQueryResult instantiates a new SeriesTimerInfoDtoQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesTimerInfoDtoQueryResult() *SeriesTimerInfoDtoQueryResult {
	this := SeriesTimerInfoDtoQueryResult{}
	return &this
}

// NewSeriesTimerInfoDtoQueryResultWithDefaults instantiates a new SeriesTimerInfoDtoQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesTimerInfoDtoQueryResultWithDefaults() *SeriesTimerInfoDtoQueryResult {
	this := SeriesTimerInfoDtoQueryResult{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SeriesTimerInfoDtoQueryResult) GetItems() []SeriesTimerInfoDto {
	if o == nil || IsNil(o.Items) {
		var ret []SeriesTimerInfoDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesTimerInfoDtoQueryResult) GetItemsOk() ([]SeriesTimerInfoDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SeriesTimerInfoDtoQueryResult) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SeriesTimerInfoDto and assigns it to the Items field.
func (o *SeriesTimerInfoDtoQueryResult) SetItems(v []SeriesTimerInfoDto) {
	o.Items = v
}

// GetTotalRecordCount returns the TotalRecordCount field value if set, zero value otherwise.
func (o *SeriesTimerInfoDtoQueryResult) GetTotalRecordCount() int32 {
	if o == nil || IsNil(o.TotalRecordCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordCount
}

// GetTotalRecordCountOk returns a tuple with the TotalRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesTimerInfoDtoQueryResult) GetTotalRecordCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordCount) {
		return nil, false
	}
	return o.TotalRecordCount, true
}

// HasTotalRecordCount returns a boolean if a field has been set.
func (o *SeriesTimerInfoDtoQueryResult) HasTotalRecordCount() bool {
	if o != nil && !IsNil(o.TotalRecordCount) {
		return true
	}

	return false
}

// SetTotalRecordCount gets a reference to the given int32 and assigns it to the TotalRecordCount field.
func (o *SeriesTimerInfoDtoQueryResult) SetTotalRecordCount(v int32) {
	o.TotalRecordCount = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *SeriesTimerInfoDtoQueryResult) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex) {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesTimerInfoDtoQueryResult) GetStartIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.StartIndex) {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *SeriesTimerInfoDtoQueryResult) HasStartIndex() bool {
	if o != nil && !IsNil(o.StartIndex) {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *SeriesTimerInfoDtoQueryResult) SetStartIndex(v int32) {
	o.StartIndex = &v
}

func (o SeriesTimerInfoDtoQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesTimerInfoDtoQueryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
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

type NullableSeriesTimerInfoDtoQueryResult struct {
	value *SeriesTimerInfoDtoQueryResult
	isSet bool
}

func (v NullableSeriesTimerInfoDtoQueryResult) Get() *SeriesTimerInfoDtoQueryResult {
	return v.value
}

func (v *NullableSeriesTimerInfoDtoQueryResult) Set(val *SeriesTimerInfoDtoQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesTimerInfoDtoQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesTimerInfoDtoQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesTimerInfoDtoQueryResult(val *SeriesTimerInfoDtoQueryResult) *NullableSeriesTimerInfoDtoQueryResult {
	return &NullableSeriesTimerInfoDtoQueryResult{value: val, isSet: true}
}

func (v NullableSeriesTimerInfoDtoQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesTimerInfoDtoQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


