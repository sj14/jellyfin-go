/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MediaPathInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaPathInfo{}

// MediaPathInfo struct for MediaPathInfo
type MediaPathInfo struct {
	Path *string `json:"Path,omitempty"`
}

// NewMediaPathInfo instantiates a new MediaPathInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaPathInfo() *MediaPathInfo {
	this := MediaPathInfo{}
	return &this
}

// NewMediaPathInfoWithDefaults instantiates a new MediaPathInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaPathInfoWithDefaults() *MediaPathInfo {
	this := MediaPathInfo{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *MediaPathInfo) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaPathInfo) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *MediaPathInfo) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *MediaPathInfo) SetPath(v string) {
	o.Path = &v
}

func (o MediaPathInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaPathInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	return toSerialize, nil
}

type NullableMediaPathInfo struct {
	value *MediaPathInfo
	isSet bool
}

func (v NullableMediaPathInfo) Get() *MediaPathInfo {
	return v.value
}

func (v *NullableMediaPathInfo) Set(val *MediaPathInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaPathInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaPathInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaPathInfo(val *MediaPathInfo) *NullableMediaPathInfo {
	return &NullableMediaPathInfo{value: val, isSet: true}
}

func (v NullableMediaPathInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaPathInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


