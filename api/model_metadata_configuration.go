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

// checks if the MetadataConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataConfiguration{}

// MetadataConfiguration struct for MetadataConfiguration
type MetadataConfiguration struct {
	UseFileCreationTimeForDateAdded *bool `json:"UseFileCreationTimeForDateAdded,omitempty"`
}

// NewMetadataConfiguration instantiates a new MetadataConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataConfiguration() *MetadataConfiguration {
	this := MetadataConfiguration{}
	return &this
}

// NewMetadataConfigurationWithDefaults instantiates a new MetadataConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataConfigurationWithDefaults() *MetadataConfiguration {
	this := MetadataConfiguration{}
	return &this
}

// GetUseFileCreationTimeForDateAdded returns the UseFileCreationTimeForDateAdded field value if set, zero value otherwise.
func (o *MetadataConfiguration) GetUseFileCreationTimeForDateAdded() bool {
	if o == nil || IsNil(o.UseFileCreationTimeForDateAdded) {
		var ret bool
		return ret
	}
	return *o.UseFileCreationTimeForDateAdded
}

// GetUseFileCreationTimeForDateAddedOk returns a tuple with the UseFileCreationTimeForDateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataConfiguration) GetUseFileCreationTimeForDateAddedOk() (*bool, bool) {
	if o == nil || IsNil(o.UseFileCreationTimeForDateAdded) {
		return nil, false
	}
	return o.UseFileCreationTimeForDateAdded, true
}

// HasUseFileCreationTimeForDateAdded returns a boolean if a field has been set.
func (o *MetadataConfiguration) HasUseFileCreationTimeForDateAdded() bool {
	if o != nil && !IsNil(o.UseFileCreationTimeForDateAdded) {
		return true
	}

	return false
}

// SetUseFileCreationTimeForDateAdded gets a reference to the given bool and assigns it to the UseFileCreationTimeForDateAdded field.
func (o *MetadataConfiguration) SetUseFileCreationTimeForDateAdded(v bool) {
	o.UseFileCreationTimeForDateAdded = &v
}

func (o MetadataConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UseFileCreationTimeForDateAdded) {
		toSerialize["UseFileCreationTimeForDateAdded"] = o.UseFileCreationTimeForDateAdded
	}
	return toSerialize, nil
}

type NullableMetadataConfiguration struct {
	value *MetadataConfiguration
	isSet bool
}

func (v NullableMetadataConfiguration) Get() *MetadataConfiguration {
	return v.value
}

func (v *NullableMetadataConfiguration) Set(val *MetadataConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataConfiguration(val *MetadataConfiguration) *NullableMetadataConfiguration {
	return &NullableMetadataConfiguration{value: val, isSet: true}
}

func (v NullableMetadataConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


