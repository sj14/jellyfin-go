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

// checks if the CustomDatabaseOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDatabaseOption{}

// CustomDatabaseOption The custom value option for custom database providers.
type CustomDatabaseOption struct {
	// Gets or sets the key of the value.
	Key *string `json:"Key,omitempty"`
	// Gets or sets the value.
	Value *string `json:"Value,omitempty"`
}

// NewCustomDatabaseOption instantiates a new CustomDatabaseOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDatabaseOption() *CustomDatabaseOption {
	this := CustomDatabaseOption{}
	return &this
}

// NewCustomDatabaseOptionWithDefaults instantiates a new CustomDatabaseOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDatabaseOptionWithDefaults() *CustomDatabaseOption {
	this := CustomDatabaseOption{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CustomDatabaseOption) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDatabaseOption) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CustomDatabaseOption) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CustomDatabaseOption) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CustomDatabaseOption) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDatabaseOption) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CustomDatabaseOption) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CustomDatabaseOption) SetValue(v string) {
	o.Value = &v
}

func (o CustomDatabaseOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDatabaseOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	return toSerialize, nil
}

type NullableCustomDatabaseOption struct {
	value *CustomDatabaseOption
	isSet bool
}

func (v NullableCustomDatabaseOption) Get() *CustomDatabaseOption {
	return v.value
}

func (v *NullableCustomDatabaseOption) Set(val *CustomDatabaseOption) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDatabaseOption) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDatabaseOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDatabaseOption(val *CustomDatabaseOption) *NullableCustomDatabaseOption {
	return &NullableCustomDatabaseOption{value: val, isSet: true}
}

func (v NullableCustomDatabaseOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDatabaseOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


