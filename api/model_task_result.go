/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the TaskResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskResult{}

// TaskResult Class TaskExecutionInfo.
type TaskResult struct {
	// Gets or sets the start time UTC.
	StartTimeUtc *time.Time `json:"StartTimeUtc,omitempty"`
	// Gets or sets the end time UTC.
	EndTimeUtc *time.Time `json:"EndTimeUtc,omitempty"`
	// Gets or sets the status.
	Status *TaskCompletionStatus `json:"Status,omitempty"`
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the key.
	Key NullableString `json:"Key,omitempty"`
	// Gets or sets the id.
	Id NullableString `json:"Id,omitempty"`
	// Gets or sets the error message.
	ErrorMessage NullableString `json:"ErrorMessage,omitempty"`
	// Gets or sets the long error message.
	LongErrorMessage NullableString `json:"LongErrorMessage,omitempty"`
}

// NewTaskResult instantiates a new TaskResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskResult() *TaskResult {
	this := TaskResult{}
	return &this
}

// NewTaskResultWithDefaults instantiates a new TaskResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskResultWithDefaults() *TaskResult {
	this := TaskResult{}
	return &this
}

// GetStartTimeUtc returns the StartTimeUtc field value if set, zero value otherwise.
func (o *TaskResult) GetStartTimeUtc() time.Time {
	if o == nil || IsNil(o.StartTimeUtc) {
		var ret time.Time
		return ret
	}
	return *o.StartTimeUtc
}

// GetStartTimeUtcOk returns a tuple with the StartTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResult) GetStartTimeUtcOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTimeUtc) {
		return nil, false
	}
	return o.StartTimeUtc, true
}

// HasStartTimeUtc returns a boolean if a field has been set.
func (o *TaskResult) HasStartTimeUtc() bool {
	if o != nil && !IsNil(o.StartTimeUtc) {
		return true
	}

	return false
}

// SetStartTimeUtc gets a reference to the given time.Time and assigns it to the StartTimeUtc field.
func (o *TaskResult) SetStartTimeUtc(v time.Time) {
	o.StartTimeUtc = &v
}

// GetEndTimeUtc returns the EndTimeUtc field value if set, zero value otherwise.
func (o *TaskResult) GetEndTimeUtc() time.Time {
	if o == nil || IsNil(o.EndTimeUtc) {
		var ret time.Time
		return ret
	}
	return *o.EndTimeUtc
}

// GetEndTimeUtcOk returns a tuple with the EndTimeUtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResult) GetEndTimeUtcOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTimeUtc) {
		return nil, false
	}
	return o.EndTimeUtc, true
}

// HasEndTimeUtc returns a boolean if a field has been set.
func (o *TaskResult) HasEndTimeUtc() bool {
	if o != nil && !IsNil(o.EndTimeUtc) {
		return true
	}

	return false
}

// SetEndTimeUtc gets a reference to the given time.Time and assigns it to the EndTimeUtc field.
func (o *TaskResult) SetEndTimeUtc(v time.Time) {
	o.EndTimeUtc = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TaskResult) GetStatus() TaskCompletionStatus {
	if o == nil || IsNil(o.Status) {
		var ret TaskCompletionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResult) GetStatusOk() (*TaskCompletionStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TaskResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given TaskCompletionStatus and assigns it to the Status field.
func (o *TaskResult) SetStatus(v TaskCompletionStatus) {
	o.Status = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskResult) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskResult) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TaskResult) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TaskResult) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TaskResult) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TaskResult) UnsetName() {
	o.Name.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskResult) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskResult) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *TaskResult) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *TaskResult) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *TaskResult) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *TaskResult) UnsetKey() {
	o.Key.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskResult) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskResult) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TaskResult) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *TaskResult) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *TaskResult) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TaskResult) UnsetId() {
	o.Id.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskResult) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskResult) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *TaskResult) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given NullableString and assigns it to the ErrorMessage field.
func (o *TaskResult) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}
// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil
func (o *TaskResult) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
func (o *TaskResult) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetLongErrorMessage returns the LongErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TaskResult) GetLongErrorMessage() string {
	if o == nil || IsNil(o.LongErrorMessage.Get()) {
		var ret string
		return ret
	}
	return *o.LongErrorMessage.Get()
}

// GetLongErrorMessageOk returns a tuple with the LongErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TaskResult) GetLongErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LongErrorMessage.Get(), o.LongErrorMessage.IsSet()
}

// HasLongErrorMessage returns a boolean if a field has been set.
func (o *TaskResult) HasLongErrorMessage() bool {
	if o != nil && o.LongErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetLongErrorMessage gets a reference to the given NullableString and assigns it to the LongErrorMessage field.
func (o *TaskResult) SetLongErrorMessage(v string) {
	o.LongErrorMessage.Set(&v)
}
// SetLongErrorMessageNil sets the value for LongErrorMessage to be an explicit nil
func (o *TaskResult) SetLongErrorMessageNil() {
	o.LongErrorMessage.Set(nil)
}

// UnsetLongErrorMessage ensures that no value is present for LongErrorMessage, not even an explicit nil
func (o *TaskResult) UnsetLongErrorMessage() {
	o.LongErrorMessage.Unset()
}

func (o TaskResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTimeUtc) {
		toSerialize["StartTimeUtc"] = o.StartTimeUtc
	}
	if !IsNil(o.EndTimeUtc) {
		toSerialize["EndTimeUtc"] = o.EndTimeUtc
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Key.IsSet() {
		toSerialize["Key"] = o.Key.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["ErrorMessage"] = o.ErrorMessage.Get()
	}
	if o.LongErrorMessage.IsSet() {
		toSerialize["LongErrorMessage"] = o.LongErrorMessage.Get()
	}
	return toSerialize, nil
}

type NullableTaskResult struct {
	value *TaskResult
	isSet bool
}

func (v NullableTaskResult) Get() *TaskResult {
	return v.value
}

func (v *NullableTaskResult) Set(val *TaskResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskResult(val *TaskResult) *NullableTaskResult {
	return &NullableTaskResult{value: val, isSet: true}
}

func (v NullableTaskResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


