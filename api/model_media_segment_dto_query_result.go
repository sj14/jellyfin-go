/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MediaSegmentDtoQueryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaSegmentDtoQueryResult{}

// MediaSegmentDtoQueryResult Query result container.
type MediaSegmentDtoQueryResult struct {
	// Gets or sets the items.
	Items []MediaSegmentDto `json:"Items,omitempty"`
	// Gets or sets the total number of records available.
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty"`
	// Gets or sets the index of the first record in Items.
	StartIndex *int32 `json:"StartIndex,omitempty"`
}

// NewMediaSegmentDtoQueryResult instantiates a new MediaSegmentDtoQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaSegmentDtoQueryResult() *MediaSegmentDtoQueryResult {
	this := MediaSegmentDtoQueryResult{}
	return &this
}

// NewMediaSegmentDtoQueryResultWithDefaults instantiates a new MediaSegmentDtoQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaSegmentDtoQueryResultWithDefaults() *MediaSegmentDtoQueryResult {
	this := MediaSegmentDtoQueryResult{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *MediaSegmentDtoQueryResult) GetItems() []MediaSegmentDto {
	if o == nil || IsNil(o.Items) {
		var ret []MediaSegmentDto
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSegmentDtoQueryResult) GetItemsOk() ([]MediaSegmentDto, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *MediaSegmentDtoQueryResult) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []MediaSegmentDto and assigns it to the Items field.
func (o *MediaSegmentDtoQueryResult) SetItems(v []MediaSegmentDto) {
	o.Items = v
}

// GetTotalRecordCount returns the TotalRecordCount field value if set, zero value otherwise.
func (o *MediaSegmentDtoQueryResult) GetTotalRecordCount() int32 {
	if o == nil || IsNil(o.TotalRecordCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordCount
}

// GetTotalRecordCountOk returns a tuple with the TotalRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSegmentDtoQueryResult) GetTotalRecordCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordCount) {
		return nil, false
	}
	return o.TotalRecordCount, true
}

// HasTotalRecordCount returns a boolean if a field has been set.
func (o *MediaSegmentDtoQueryResult) HasTotalRecordCount() bool {
	if o != nil && !IsNil(o.TotalRecordCount) {
		return true
	}

	return false
}

// SetTotalRecordCount gets a reference to the given int32 and assigns it to the TotalRecordCount field.
func (o *MediaSegmentDtoQueryResult) SetTotalRecordCount(v int32) {
	o.TotalRecordCount = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *MediaSegmentDtoQueryResult) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex) {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSegmentDtoQueryResult) GetStartIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.StartIndex) {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *MediaSegmentDtoQueryResult) HasStartIndex() bool {
	if o != nil && !IsNil(o.StartIndex) {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *MediaSegmentDtoQueryResult) SetStartIndex(v int32) {
	o.StartIndex = &v
}

func (o MediaSegmentDtoQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaSegmentDtoQueryResult) ToMap() (map[string]interface{}, error) {
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

type NullableMediaSegmentDtoQueryResult struct {
	value *MediaSegmentDtoQueryResult
	isSet bool
}

func (v NullableMediaSegmentDtoQueryResult) Get() *MediaSegmentDtoQueryResult {
	return v.value
}

func (v *NullableMediaSegmentDtoQueryResult) Set(val *MediaSegmentDtoQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaSegmentDtoQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaSegmentDtoQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaSegmentDtoQueryResult(val *MediaSegmentDtoQueryResult) *NullableMediaSegmentDtoQueryResult {
	return &NullableMediaSegmentDtoQueryResult{value: val, isSet: true}
}

func (v NullableMediaSegmentDtoQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaSegmentDtoQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


