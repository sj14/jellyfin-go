/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AddVirtualFolderDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddVirtualFolderDto{}

// AddVirtualFolderDto Add virtual folder dto.
type AddVirtualFolderDto struct {
	LibraryOptions NullableAddVirtualFolderDtoLibraryOptions `json:"LibraryOptions,omitempty"`
}

// NewAddVirtualFolderDto instantiates a new AddVirtualFolderDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddVirtualFolderDto() *AddVirtualFolderDto {
	this := AddVirtualFolderDto{}
	return &this
}

// NewAddVirtualFolderDtoWithDefaults instantiates a new AddVirtualFolderDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddVirtualFolderDtoWithDefaults() *AddVirtualFolderDto {
	this := AddVirtualFolderDto{}
	return &this
}

// GetLibraryOptions returns the LibraryOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddVirtualFolderDto) GetLibraryOptions() AddVirtualFolderDtoLibraryOptions {
	if o == nil || IsNil(o.LibraryOptions.Get()) {
		var ret AddVirtualFolderDtoLibraryOptions
		return ret
	}
	return *o.LibraryOptions.Get()
}

// GetLibraryOptionsOk returns a tuple with the LibraryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddVirtualFolderDto) GetLibraryOptionsOk() (*AddVirtualFolderDtoLibraryOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.LibraryOptions.Get(), o.LibraryOptions.IsSet()
}

// HasLibraryOptions returns a boolean if a field has been set.
func (o *AddVirtualFolderDto) HasLibraryOptions() bool {
	if o != nil && o.LibraryOptions.IsSet() {
		return true
	}

	return false
}

// SetLibraryOptions gets a reference to the given NullableAddVirtualFolderDtoLibraryOptions and assigns it to the LibraryOptions field.
func (o *AddVirtualFolderDto) SetLibraryOptions(v AddVirtualFolderDtoLibraryOptions) {
	o.LibraryOptions.Set(&v)
}
// SetLibraryOptionsNil sets the value for LibraryOptions to be an explicit nil
func (o *AddVirtualFolderDto) SetLibraryOptionsNil() {
	o.LibraryOptions.Set(nil)
}

// UnsetLibraryOptions ensures that no value is present for LibraryOptions, not even an explicit nil
func (o *AddVirtualFolderDto) UnsetLibraryOptions() {
	o.LibraryOptions.Unset()
}

func (o AddVirtualFolderDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddVirtualFolderDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LibraryOptions.IsSet() {
		toSerialize["LibraryOptions"] = o.LibraryOptions.Get()
	}
	return toSerialize, nil
}

type NullableAddVirtualFolderDto struct {
	value *AddVirtualFolderDto
	isSet bool
}

func (v NullableAddVirtualFolderDto) Get() *AddVirtualFolderDto {
	return v.value
}

func (v *NullableAddVirtualFolderDto) Set(val *AddVirtualFolderDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAddVirtualFolderDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAddVirtualFolderDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddVirtualFolderDto(val *AddVirtualFolderDto) *NullableAddVirtualFolderDto {
	return &NullableAddVirtualFolderDto{value: val, isSet: true}
}

func (v NullableAddVirtualFolderDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddVirtualFolderDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


