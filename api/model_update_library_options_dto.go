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

// checks if the UpdateLibraryOptionsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLibraryOptionsDto{}

// UpdateLibraryOptionsDto Update library options dto.
type UpdateLibraryOptionsDto struct {
	// Gets or sets the library item id.
	Id *string `json:"Id,omitempty"`
	// Gets or sets library options.
	LibraryOptions NullableLibraryOptions `json:"LibraryOptions,omitempty"`
}

// NewUpdateLibraryOptionsDto instantiates a new UpdateLibraryOptionsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLibraryOptionsDto() *UpdateLibraryOptionsDto {
	this := UpdateLibraryOptionsDto{}
	return &this
}

// NewUpdateLibraryOptionsDtoWithDefaults instantiates a new UpdateLibraryOptionsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLibraryOptionsDtoWithDefaults() *UpdateLibraryOptionsDto {
	this := UpdateLibraryOptionsDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateLibraryOptionsDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLibraryOptionsDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateLibraryOptionsDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateLibraryOptionsDto) SetId(v string) {
	o.Id = &v
}

// GetLibraryOptions returns the LibraryOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateLibraryOptionsDto) GetLibraryOptions() LibraryOptions {
	if o == nil || IsNil(o.LibraryOptions.Get()) {
		var ret LibraryOptions
		return ret
	}
	return *o.LibraryOptions.Get()
}

// GetLibraryOptionsOk returns a tuple with the LibraryOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateLibraryOptionsDto) GetLibraryOptionsOk() (*LibraryOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.LibraryOptions.Get(), o.LibraryOptions.IsSet()
}

// HasLibraryOptions returns a boolean if a field has been set.
func (o *UpdateLibraryOptionsDto) HasLibraryOptions() bool {
	if o != nil && o.LibraryOptions.IsSet() {
		return true
	}

	return false
}

// SetLibraryOptions gets a reference to the given NullableLibraryOptions and assigns it to the LibraryOptions field.
func (o *UpdateLibraryOptionsDto) SetLibraryOptions(v LibraryOptions) {
	o.LibraryOptions.Set(&v)
}
// SetLibraryOptionsNil sets the value for LibraryOptions to be an explicit nil
func (o *UpdateLibraryOptionsDto) SetLibraryOptionsNil() {
	o.LibraryOptions.Set(nil)
}

// UnsetLibraryOptions ensures that no value is present for LibraryOptions, not even an explicit nil
func (o *UpdateLibraryOptionsDto) UnsetLibraryOptions() {
	o.LibraryOptions.Unset()
}

func (o UpdateLibraryOptionsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLibraryOptionsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.LibraryOptions.IsSet() {
		toSerialize["LibraryOptions"] = o.LibraryOptions.Get()
	}
	return toSerialize, nil
}

type NullableUpdateLibraryOptionsDto struct {
	value *UpdateLibraryOptionsDto
	isSet bool
}

func (v NullableUpdateLibraryOptionsDto) Get() *UpdateLibraryOptionsDto {
	return v.value
}

func (v *NullableUpdateLibraryOptionsDto) Set(val *UpdateLibraryOptionsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLibraryOptionsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLibraryOptionsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLibraryOptionsDto(val *UpdateLibraryOptionsDto) *NullableUpdateLibraryOptionsDto {
	return &NullableUpdateLibraryOptionsDto{value: val, isSet: true}
}

func (v NullableUpdateLibraryOptionsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLibraryOptionsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


