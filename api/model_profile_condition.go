/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ProfileCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileCondition{}

// ProfileCondition struct for ProfileCondition
type ProfileCondition struct {
	Condition *ProfileConditionType `json:"Condition,omitempty"`
	Property *ProfileConditionValue `json:"Property,omitempty"`
	Value NullableString `json:"Value,omitempty"`
	IsRequired *bool `json:"IsRequired,omitempty"`
}

// NewProfileCondition instantiates a new ProfileCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileCondition() *ProfileCondition {
	this := ProfileCondition{}
	return &this
}

// NewProfileConditionWithDefaults instantiates a new ProfileCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileConditionWithDefaults() *ProfileCondition {
	this := ProfileCondition{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *ProfileCondition) GetCondition() ProfileConditionType {
	if o == nil || IsNil(o.Condition) {
		var ret ProfileConditionType
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileCondition) GetConditionOk() (*ProfileConditionType, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ProfileCondition) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given ProfileConditionType and assigns it to the Condition field.
func (o *ProfileCondition) SetCondition(v ProfileConditionType) {
	o.Condition = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *ProfileCondition) GetProperty() ProfileConditionValue {
	if o == nil || IsNil(o.Property) {
		var ret ProfileConditionValue
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileCondition) GetPropertyOk() (*ProfileConditionValue, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *ProfileCondition) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given ProfileConditionValue and assigns it to the Property field.
func (o *ProfileCondition) SetProperty(v ProfileConditionValue) {
	o.Property = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProfileCondition) GetValue() string {
	if o == nil || IsNil(o.Value.Get()) {
		var ret string
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProfileCondition) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ProfileCondition) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableString and assigns it to the Value field.
func (o *ProfileCondition) SetValue(v string) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ProfileCondition) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ProfileCondition) UnsetValue() {
	o.Value.Unset()
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *ProfileCondition) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileCondition) GetIsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRequired) {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *ProfileCondition) HasIsRequired() bool {
	if o != nil && !IsNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *ProfileCondition) SetIsRequired(v bool) {
	o.IsRequired = &v
}

func (o ProfileCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Condition) {
		toSerialize["Condition"] = o.Condition
	}
	if !IsNil(o.Property) {
		toSerialize["Property"] = o.Property
	}
	if o.Value.IsSet() {
		toSerialize["Value"] = o.Value.Get()
	}
	if !IsNil(o.IsRequired) {
		toSerialize["IsRequired"] = o.IsRequired
	}
	return toSerialize, nil
}

type NullableProfileCondition struct {
	value *ProfileCondition
	isSet bool
}

func (v NullableProfileCondition) Get() *ProfileCondition {
	return v.value
}

func (v *NullableProfileCondition) Set(val *ProfileCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileCondition(val *ProfileCondition) *NullableProfileCondition {
	return &NullableProfileCondition{value: val, isSet: true}
}

func (v NullableProfileCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


