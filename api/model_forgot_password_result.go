/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the ForgotPasswordResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForgotPasswordResult{}

// ForgotPasswordResult struct for ForgotPasswordResult
type ForgotPasswordResult struct {
	// Gets or sets the action.
	Action *ForgotPasswordAction `json:"Action,omitempty"`
	// Gets or sets the pin file.
	PinFile NullableString `json:"PinFile,omitempty"`
	// Gets or sets the pin expiration date.
	PinExpirationDate NullableTime `json:"PinExpirationDate,omitempty"`
}

// NewForgotPasswordResult instantiates a new ForgotPasswordResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForgotPasswordResult() *ForgotPasswordResult {
	this := ForgotPasswordResult{}
	return &this
}

// NewForgotPasswordResultWithDefaults instantiates a new ForgotPasswordResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForgotPasswordResultWithDefaults() *ForgotPasswordResult {
	this := ForgotPasswordResult{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ForgotPasswordResult) GetAction() ForgotPasswordAction {
	if o == nil || IsNil(o.Action) {
		var ret ForgotPasswordAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForgotPasswordResult) GetActionOk() (*ForgotPasswordAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ForgotPasswordResult) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given ForgotPasswordAction and assigns it to the Action field.
func (o *ForgotPasswordResult) SetAction(v ForgotPasswordAction) {
	o.Action = &v
}

// GetPinFile returns the PinFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForgotPasswordResult) GetPinFile() string {
	if o == nil || IsNil(o.PinFile.Get()) {
		var ret string
		return ret
	}
	return *o.PinFile.Get()
}

// GetPinFileOk returns a tuple with the PinFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForgotPasswordResult) GetPinFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PinFile.Get(), o.PinFile.IsSet()
}

// HasPinFile returns a boolean if a field has been set.
func (o *ForgotPasswordResult) HasPinFile() bool {
	if o != nil && o.PinFile.IsSet() {
		return true
	}

	return false
}

// SetPinFile gets a reference to the given NullableString and assigns it to the PinFile field.
func (o *ForgotPasswordResult) SetPinFile(v string) {
	o.PinFile.Set(&v)
}
// SetPinFileNil sets the value for PinFile to be an explicit nil
func (o *ForgotPasswordResult) SetPinFileNil() {
	o.PinFile.Set(nil)
}

// UnsetPinFile ensures that no value is present for PinFile, not even an explicit nil
func (o *ForgotPasswordResult) UnsetPinFile() {
	o.PinFile.Unset()
}

// GetPinExpirationDate returns the PinExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForgotPasswordResult) GetPinExpirationDate() time.Time {
	if o == nil || IsNil(o.PinExpirationDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PinExpirationDate.Get()
}

// GetPinExpirationDateOk returns a tuple with the PinExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForgotPasswordResult) GetPinExpirationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PinExpirationDate.Get(), o.PinExpirationDate.IsSet()
}

// HasPinExpirationDate returns a boolean if a field has been set.
func (o *ForgotPasswordResult) HasPinExpirationDate() bool {
	if o != nil && o.PinExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetPinExpirationDate gets a reference to the given NullableTime and assigns it to the PinExpirationDate field.
func (o *ForgotPasswordResult) SetPinExpirationDate(v time.Time) {
	o.PinExpirationDate.Set(&v)
}
// SetPinExpirationDateNil sets the value for PinExpirationDate to be an explicit nil
func (o *ForgotPasswordResult) SetPinExpirationDateNil() {
	o.PinExpirationDate.Set(nil)
}

// UnsetPinExpirationDate ensures that no value is present for PinExpirationDate, not even an explicit nil
func (o *ForgotPasswordResult) UnsetPinExpirationDate() {
	o.PinExpirationDate.Unset()
}

func (o ForgotPasswordResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForgotPasswordResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["Action"] = o.Action
	}
	if o.PinFile.IsSet() {
		toSerialize["PinFile"] = o.PinFile.Get()
	}
	if o.PinExpirationDate.IsSet() {
		toSerialize["PinExpirationDate"] = o.PinExpirationDate.Get()
	}
	return toSerialize, nil
}

type NullableForgotPasswordResult struct {
	value *ForgotPasswordResult
	isSet bool
}

func (v NullableForgotPasswordResult) Get() *ForgotPasswordResult {
	return v.value
}

func (v *NullableForgotPasswordResult) Set(val *ForgotPasswordResult) {
	v.value = val
	v.isSet = true
}

func (v NullableForgotPasswordResult) IsSet() bool {
	return v.isSet
}

func (v *NullableForgotPasswordResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForgotPasswordResult(val *ForgotPasswordResult) *NullableForgotPasswordResult {
	return &NullableForgotPasswordResult{value: val, isSet: true}
}

func (v NullableForgotPasswordResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForgotPasswordResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


