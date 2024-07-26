/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the NameGuidPair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NameGuidPair{}

// NameGuidPair struct for NameGuidPair
type NameGuidPair struct {
	Name NullableString `json:"Name,omitempty"`
	Id *string `json:"Id,omitempty"`
}

// NewNameGuidPair instantiates a new NameGuidPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameGuidPair() *NameGuidPair {
	this := NameGuidPair{}
	return &this
}

// NewNameGuidPairWithDefaults instantiates a new NameGuidPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameGuidPairWithDefaults() *NameGuidPair {
	this := NameGuidPair{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NameGuidPair) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NameGuidPair) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *NameGuidPair) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *NameGuidPair) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *NameGuidPair) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *NameGuidPair) UnsetName() {
	o.Name.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NameGuidPair) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameGuidPair) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NameGuidPair) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NameGuidPair) SetId(v string) {
	o.Id = &v
}

func (o NameGuidPair) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NameGuidPair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	return toSerialize, nil
}

type NullableNameGuidPair struct {
	value *NameGuidPair
	isSet bool
}

func (v NullableNameGuidPair) Get() *NameGuidPair {
	return v.value
}

func (v *NullableNameGuidPair) Set(val *NameGuidPair) {
	v.value = val
	v.isSet = true
}

func (v NullableNameGuidPair) IsSet() bool {
	return v.isSet
}

func (v *NullableNameGuidPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameGuidPair(val *NameGuidPair) *NullableNameGuidPair {
	return &NullableNameGuidPair{value: val, isSet: true}
}

func (v NullableNameGuidPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameGuidPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


