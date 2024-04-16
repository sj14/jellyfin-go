/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.8.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TimerEventInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimerEventInfo{}

// TimerEventInfo struct for TimerEventInfo
type TimerEventInfo struct {
	Id *string `json:"Id,omitempty"`
	ProgramId NullableString `json:"ProgramId,omitempty"`
}

// NewTimerEventInfo instantiates a new TimerEventInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimerEventInfo() *TimerEventInfo {
	this := TimerEventInfo{}
	return &this
}

// NewTimerEventInfoWithDefaults instantiates a new TimerEventInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimerEventInfoWithDefaults() *TimerEventInfo {
	this := TimerEventInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TimerEventInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerEventInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TimerEventInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TimerEventInfo) SetId(v string) {
	o.Id = &v
}

// GetProgramId returns the ProgramId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimerEventInfo) GetProgramId() string {
	if o == nil || IsNil(o.ProgramId.Get()) {
		var ret string
		return ret
	}
	return *o.ProgramId.Get()
}

// GetProgramIdOk returns a tuple with the ProgramId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimerEventInfo) GetProgramIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProgramId.Get(), o.ProgramId.IsSet()
}

// HasProgramId returns a boolean if a field has been set.
func (o *TimerEventInfo) HasProgramId() bool {
	if o != nil && o.ProgramId.IsSet() {
		return true
	}

	return false
}

// SetProgramId gets a reference to the given NullableString and assigns it to the ProgramId field.
func (o *TimerEventInfo) SetProgramId(v string) {
	o.ProgramId.Set(&v)
}
// SetProgramIdNil sets the value for ProgramId to be an explicit nil
func (o *TimerEventInfo) SetProgramIdNil() {
	o.ProgramId.Set(nil)
}

// UnsetProgramId ensures that no value is present for ProgramId, not even an explicit nil
func (o *TimerEventInfo) UnsetProgramId() {
	o.ProgramId.Unset()
}

func (o TimerEventInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimerEventInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.ProgramId.IsSet() {
		toSerialize["ProgramId"] = o.ProgramId.Get()
	}
	return toSerialize, nil
}

type NullableTimerEventInfo struct {
	value *TimerEventInfo
	isSet bool
}

func (v NullableTimerEventInfo) Get() *TimerEventInfo {
	return v.value
}

func (v *NullableTimerEventInfo) Set(val *TimerEventInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTimerEventInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTimerEventInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimerEventInfo(val *TimerEventInfo) *NullableTimerEventInfo {
	return &NullableTimerEventInfo{value: val, isSet: true}
}

func (v NullableTimerEventInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimerEventInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


