/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ParentalRating type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParentalRating{}

// ParentalRating Class ParentalRating.
type ParentalRating struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the value.
	Value NullableInt32 `json:"Value,omitempty"`
}

// NewParentalRating instantiates a new ParentalRating object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParentalRating() *ParentalRating {
	this := ParentalRating{}
	return &this
}

// NewParentalRatingWithDefaults instantiates a new ParentalRating object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParentalRatingWithDefaults() *ParentalRating {
	this := ParentalRating{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentalRating) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentalRating) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ParentalRating) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ParentalRating) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ParentalRating) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ParentalRating) UnsetName() {
	o.Name.Unset()
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParentalRating) GetValue() int32 {
	if o == nil || IsNil(o.Value.Get()) {
		var ret int32
		return ret
	}
	return *o.Value.Get()
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParentalRating) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value.Get(), o.Value.IsSet()
}

// HasValue returns a boolean if a field has been set.
func (o *ParentalRating) HasValue() bool {
	if o != nil && o.Value.IsSet() {
		return true
	}

	return false
}

// SetValue gets a reference to the given NullableInt32 and assigns it to the Value field.
func (o *ParentalRating) SetValue(v int32) {
	o.Value.Set(&v)
}
// SetValueNil sets the value for Value to be an explicit nil
func (o *ParentalRating) SetValueNil() {
	o.Value.Set(nil)
}

// UnsetValue ensures that no value is present for Value, not even an explicit nil
func (o *ParentalRating) UnsetValue() {
	o.Value.Unset()
}

func (o ParentalRating) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParentalRating) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Value.IsSet() {
		toSerialize["Value"] = o.Value.Get()
	}
	return toSerialize, nil
}

type NullableParentalRating struct {
	value *ParentalRating
	isSet bool
}

func (v NullableParentalRating) Get() *ParentalRating {
	return v.value
}

func (v *NullableParentalRating) Set(val *ParentalRating) {
	v.value = val
	v.isSet = true
}

func (v NullableParentalRating) IsSet() bool {
	return v.isSet
}

func (v *NullableParentalRating) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParentalRating(val *ParentalRating) *NullableParentalRating {
	return &NullableParentalRating{value: val, isSet: true}
}

func (v NullableParentalRating) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParentalRating) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


