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

// checks if the TaskTriggerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskTriggerInfo{}

// TaskTriggerInfo Class TaskTriggerInfo.
type TaskTriggerInfo struct {
	// Gets or sets the type.
	Type *TaskTriggerInfoType `json:"Type,omitempty"`
	// Gets or sets the time of day.
	TimeOfDayTicks NullableInt64 `json:"TimeOfDayTicks,omitempty"`
	// Gets or sets the interval.
	IntervalTicks NullableInt64 `json:"IntervalTicks,omitempty"`
	// Gets or sets the day of week.
	DayOfWeek NullableDayOfWeek `json:"DayOfWeek,omitempty"`
	// Gets or sets the maximum runtime ticks.
	MaxRuntimeTicks NullableInt64 `json:"MaxRuntimeTicks,omitempty"`
}

// NewTaskTriggerInfo instantiates a new TaskTriggerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskTriggerInfo() *TaskTriggerInfo {
	this := TaskTriggerInfo{}
	return &this
}

// NewTaskTriggerInfoWithDefaults instantiates a new TaskTriggerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskTriggerInfoWithDefaults() *TaskTriggerInfo {
	this := TaskTriggerInfo{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaskTriggerInfo) GetType() TaskTriggerInfoType {
	if o == nil || IsNil(o.Type) {
		var ret TaskTriggerInfoType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskTriggerInfo) GetTypeOk() (*TaskTriggerInfoType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaskTriggerInfo) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given TaskTriggerInfoType and assigns it to the Type field.
func (o *TaskTriggerInfo) SetType(v TaskTriggerInfoType) {
	o.Type = &v
}

// GetTimeOfDayTicks returns the TimeOfDayTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskTriggerInfo) GetTimeOfDayTicks() int64 {
	if o == nil || IsNil(o.TimeOfDayTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.TimeOfDayTicks.Get()
}

// GetTimeOfDayTicksOk returns a tuple with the TimeOfDayTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskTriggerInfo) GetTimeOfDayTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeOfDayTicks.Get(), o.TimeOfDayTicks.IsSet()
}

// HasTimeOfDayTicks returns a boolean if a field has been set.
func (o *TaskTriggerInfo) HasTimeOfDayTicks() bool {
	if o != nil && o.TimeOfDayTicks.IsSet() {
		return true
	}

	return false
}

// SetTimeOfDayTicks gets a reference to the given NullableInt64 and assigns it to the TimeOfDayTicks field.
func (o *TaskTriggerInfo) SetTimeOfDayTicks(v int64) {
	o.TimeOfDayTicks.Set(&v)
}
// SetTimeOfDayTicksNil sets the value for TimeOfDayTicks to be an explicit nil
func (o *TaskTriggerInfo) SetTimeOfDayTicksNil() {
	o.TimeOfDayTicks.Set(nil)
}

// UnsetTimeOfDayTicks ensures that no value is present for TimeOfDayTicks, not even an explicit nil
func (o *TaskTriggerInfo) UnsetTimeOfDayTicks() {
	o.TimeOfDayTicks.Unset()
}

// GetIntervalTicks returns the IntervalTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskTriggerInfo) GetIntervalTicks() int64 {
	if o == nil || IsNil(o.IntervalTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.IntervalTicks.Get()
}

// GetIntervalTicksOk returns a tuple with the IntervalTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskTriggerInfo) GetIntervalTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IntervalTicks.Get(), o.IntervalTicks.IsSet()
}

// HasIntervalTicks returns a boolean if a field has been set.
func (o *TaskTriggerInfo) HasIntervalTicks() bool {
	if o != nil && o.IntervalTicks.IsSet() {
		return true
	}

	return false
}

// SetIntervalTicks gets a reference to the given NullableInt64 and assigns it to the IntervalTicks field.
func (o *TaskTriggerInfo) SetIntervalTicks(v int64) {
	o.IntervalTicks.Set(&v)
}
// SetIntervalTicksNil sets the value for IntervalTicks to be an explicit nil
func (o *TaskTriggerInfo) SetIntervalTicksNil() {
	o.IntervalTicks.Set(nil)
}

// UnsetIntervalTicks ensures that no value is present for IntervalTicks, not even an explicit nil
func (o *TaskTriggerInfo) UnsetIntervalTicks() {
	o.IntervalTicks.Unset()
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskTriggerInfo) GetDayOfWeek() DayOfWeek {
	if o == nil || IsNil(o.DayOfWeek.Get()) {
		var ret DayOfWeek
		return ret
	}
	return *o.DayOfWeek.Get()
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskTriggerInfo) GetDayOfWeekOk() (*DayOfWeek, bool) {
	if o == nil {
		return nil, false
	}
	return o.DayOfWeek.Get(), o.DayOfWeek.IsSet()
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *TaskTriggerInfo) HasDayOfWeek() bool {
	if o != nil && o.DayOfWeek.IsSet() {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given NullableDayOfWeek and assigns it to the DayOfWeek field.
func (o *TaskTriggerInfo) SetDayOfWeek(v DayOfWeek) {
	o.DayOfWeek.Set(&v)
}
// SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil
func (o *TaskTriggerInfo) SetDayOfWeekNil() {
	o.DayOfWeek.Set(nil)
}

// UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
func (o *TaskTriggerInfo) UnsetDayOfWeek() {
	o.DayOfWeek.Unset()
}

// GetMaxRuntimeTicks returns the MaxRuntimeTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskTriggerInfo) GetMaxRuntimeTicks() int64 {
	if o == nil || IsNil(o.MaxRuntimeTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.MaxRuntimeTicks.Get()
}

// GetMaxRuntimeTicksOk returns a tuple with the MaxRuntimeTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskTriggerInfo) GetMaxRuntimeTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRuntimeTicks.Get(), o.MaxRuntimeTicks.IsSet()
}

// HasMaxRuntimeTicks returns a boolean if a field has been set.
func (o *TaskTriggerInfo) HasMaxRuntimeTicks() bool {
	if o != nil && o.MaxRuntimeTicks.IsSet() {
		return true
	}

	return false
}

// SetMaxRuntimeTicks gets a reference to the given NullableInt64 and assigns it to the MaxRuntimeTicks field.
func (o *TaskTriggerInfo) SetMaxRuntimeTicks(v int64) {
	o.MaxRuntimeTicks.Set(&v)
}
// SetMaxRuntimeTicksNil sets the value for MaxRuntimeTicks to be an explicit nil
func (o *TaskTriggerInfo) SetMaxRuntimeTicksNil() {
	o.MaxRuntimeTicks.Set(nil)
}

// UnsetMaxRuntimeTicks ensures that no value is present for MaxRuntimeTicks, not even an explicit nil
func (o *TaskTriggerInfo) UnsetMaxRuntimeTicks() {
	o.MaxRuntimeTicks.Unset()
}

func (o TaskTriggerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskTriggerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.TimeOfDayTicks.IsSet() {
		toSerialize["TimeOfDayTicks"] = o.TimeOfDayTicks.Get()
	}
	if o.IntervalTicks.IsSet() {
		toSerialize["IntervalTicks"] = o.IntervalTicks.Get()
	}
	if o.DayOfWeek.IsSet() {
		toSerialize["DayOfWeek"] = o.DayOfWeek.Get()
	}
	if o.MaxRuntimeTicks.IsSet() {
		toSerialize["MaxRuntimeTicks"] = o.MaxRuntimeTicks.Get()
	}
	return toSerialize, nil
}

type NullableTaskTriggerInfo struct {
	value *TaskTriggerInfo
	isSet bool
}

func (v NullableTaskTriggerInfo) Get() *TaskTriggerInfo {
	return v.value
}

func (v *NullableTaskTriggerInfo) Set(val *TaskTriggerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskTriggerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskTriggerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskTriggerInfo(val *TaskTriggerInfo) *NullableTaskTriggerInfo {
	return &NullableTaskTriggerInfo{value: val, isSet: true}
}

func (v NullableTaskTriggerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskTriggerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


