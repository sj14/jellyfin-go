/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SeriesTimerCancelledMessageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesTimerCancelledMessageData{}

// SeriesTimerCancelledMessageData Gets or sets the data.
type SeriesTimerCancelledMessageData struct {
	Id *string `json:"Id,omitempty"`
	ProgramId NullableString `json:"ProgramId,omitempty"`
}

// NewSeriesTimerCancelledMessageData instantiates a new SeriesTimerCancelledMessageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesTimerCancelledMessageData() *SeriesTimerCancelledMessageData {
	this := SeriesTimerCancelledMessageData{}
	return &this
}

// NewSeriesTimerCancelledMessageDataWithDefaults instantiates a new SeriesTimerCancelledMessageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesTimerCancelledMessageDataWithDefaults() *SeriesTimerCancelledMessageData {
	this := SeriesTimerCancelledMessageData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SeriesTimerCancelledMessageData) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesTimerCancelledMessageData) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SeriesTimerCancelledMessageData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SeriesTimerCancelledMessageData) SetId(v string) {
	o.Id = &v
}

// GetProgramId returns the ProgramId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesTimerCancelledMessageData) GetProgramId() string {
	if o == nil || IsNil(o.ProgramId.Get()) {
		var ret string
		return ret
	}
	return *o.ProgramId.Get()
}

// GetProgramIdOk returns a tuple with the ProgramId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesTimerCancelledMessageData) GetProgramIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProgramId.Get(), o.ProgramId.IsSet()
}

// HasProgramId returns a boolean if a field has been set.
func (o *SeriesTimerCancelledMessageData) HasProgramId() bool {
	if o != nil && o.ProgramId.IsSet() {
		return true
	}

	return false
}

// SetProgramId gets a reference to the given NullableString and assigns it to the ProgramId field.
func (o *SeriesTimerCancelledMessageData) SetProgramId(v string) {
	o.ProgramId.Set(&v)
}
// SetProgramIdNil sets the value for ProgramId to be an explicit nil
func (o *SeriesTimerCancelledMessageData) SetProgramIdNil() {
	o.ProgramId.Set(nil)
}

// UnsetProgramId ensures that no value is present for ProgramId, not even an explicit nil
func (o *SeriesTimerCancelledMessageData) UnsetProgramId() {
	o.ProgramId.Unset()
}

func (o SeriesTimerCancelledMessageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesTimerCancelledMessageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.ProgramId.IsSet() {
		toSerialize["ProgramId"] = o.ProgramId.Get()
	}
	return toSerialize, nil
}

type NullableSeriesTimerCancelledMessageData struct {
	value *SeriesTimerCancelledMessageData
	isSet bool
}

func (v NullableSeriesTimerCancelledMessageData) Get() *SeriesTimerCancelledMessageData {
	return v.value
}

func (v *NullableSeriesTimerCancelledMessageData) Set(val *SeriesTimerCancelledMessageData) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesTimerCancelledMessageData) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesTimerCancelledMessageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesTimerCancelledMessageData(val *SeriesTimerCancelledMessageData) *NullableSeriesTimerCancelledMessageData {
	return &NullableSeriesTimerCancelledMessageData{value: val, isSet: true}
}

func (v NullableSeriesTimerCancelledMessageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesTimerCancelledMessageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


