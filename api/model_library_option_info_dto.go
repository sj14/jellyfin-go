/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the LibraryOptionInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibraryOptionInfoDto{}

// LibraryOptionInfoDto Library option info dto.
type LibraryOptionInfoDto struct {
	// Gets or sets name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets a value indicating whether default enabled.
	DefaultEnabled *bool `json:"DefaultEnabled,omitempty"`
}

// NewLibraryOptionInfoDto instantiates a new LibraryOptionInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibraryOptionInfoDto() *LibraryOptionInfoDto {
	this := LibraryOptionInfoDto{}
	return &this
}

// NewLibraryOptionInfoDtoWithDefaults instantiates a new LibraryOptionInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibraryOptionInfoDtoWithDefaults() *LibraryOptionInfoDto {
	this := LibraryOptionInfoDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryOptionInfoDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOptionInfoDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *LibraryOptionInfoDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *LibraryOptionInfoDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *LibraryOptionInfoDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *LibraryOptionInfoDto) UnsetName() {
	o.Name.Unset()
}

// GetDefaultEnabled returns the DefaultEnabled field value if set, zero value otherwise.
func (o *LibraryOptionInfoDto) GetDefaultEnabled() bool {
	if o == nil || IsNil(o.DefaultEnabled) {
		var ret bool
		return ret
	}
	return *o.DefaultEnabled
}

// GetDefaultEnabledOk returns a tuple with the DefaultEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptionInfoDto) GetDefaultEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultEnabled) {
		return nil, false
	}
	return o.DefaultEnabled, true
}

// HasDefaultEnabled returns a boolean if a field has been set.
func (o *LibraryOptionInfoDto) HasDefaultEnabled() bool {
	if o != nil && !IsNil(o.DefaultEnabled) {
		return true
	}

	return false
}

// SetDefaultEnabled gets a reference to the given bool and assigns it to the DefaultEnabled field.
func (o *LibraryOptionInfoDto) SetDefaultEnabled(v bool) {
	o.DefaultEnabled = &v
}

func (o LibraryOptionInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibraryOptionInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.DefaultEnabled) {
		toSerialize["DefaultEnabled"] = o.DefaultEnabled
	}
	return toSerialize, nil
}

type NullableLibraryOptionInfoDto struct {
	value *LibraryOptionInfoDto
	isSet bool
}

func (v NullableLibraryOptionInfoDto) Get() *LibraryOptionInfoDto {
	return v.value
}

func (v *NullableLibraryOptionInfoDto) Set(val *LibraryOptionInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLibraryOptionInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLibraryOptionInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibraryOptionInfoDto(val *LibraryOptionInfoDto) *NullableLibraryOptionInfoDto {
	return &NullableLibraryOptionInfoDto{value: val, isSet: true}
}

func (v NullableLibraryOptionInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibraryOptionInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


