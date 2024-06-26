/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MediaPathDtoPathInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaPathDtoPathInfo{}

// MediaPathDtoPathInfo Gets or sets the path info.
type MediaPathDtoPathInfo struct {
	Path *string `json:"Path,omitempty"`
	NetworkPath NullableString `json:"NetworkPath,omitempty"`
}

// NewMediaPathDtoPathInfo instantiates a new MediaPathDtoPathInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaPathDtoPathInfo() *MediaPathDtoPathInfo {
	this := MediaPathDtoPathInfo{}
	return &this
}

// NewMediaPathDtoPathInfoWithDefaults instantiates a new MediaPathDtoPathInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaPathDtoPathInfoWithDefaults() *MediaPathDtoPathInfo {
	this := MediaPathDtoPathInfo{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *MediaPathDtoPathInfo) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaPathDtoPathInfo) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *MediaPathDtoPathInfo) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *MediaPathDtoPathInfo) SetPath(v string) {
	o.Path = &v
}

// GetNetworkPath returns the NetworkPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaPathDtoPathInfo) GetNetworkPath() string {
	if o == nil || IsNil(o.NetworkPath.Get()) {
		var ret string
		return ret
	}
	return *o.NetworkPath.Get()
}

// GetNetworkPathOk returns a tuple with the NetworkPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaPathDtoPathInfo) GetNetworkPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NetworkPath.Get(), o.NetworkPath.IsSet()
}

// HasNetworkPath returns a boolean if a field has been set.
func (o *MediaPathDtoPathInfo) HasNetworkPath() bool {
	if o != nil && o.NetworkPath.IsSet() {
		return true
	}

	return false
}

// SetNetworkPath gets a reference to the given NullableString and assigns it to the NetworkPath field.
func (o *MediaPathDtoPathInfo) SetNetworkPath(v string) {
	o.NetworkPath.Set(&v)
}
// SetNetworkPathNil sets the value for NetworkPath to be an explicit nil
func (o *MediaPathDtoPathInfo) SetNetworkPathNil() {
	o.NetworkPath.Set(nil)
}

// UnsetNetworkPath ensures that no value is present for NetworkPath, not even an explicit nil
func (o *MediaPathDtoPathInfo) UnsetNetworkPath() {
	o.NetworkPath.Unset()
}

func (o MediaPathDtoPathInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaPathDtoPathInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	if o.NetworkPath.IsSet() {
		toSerialize["NetworkPath"] = o.NetworkPath.Get()
	}
	return toSerialize, nil
}

type NullableMediaPathDtoPathInfo struct {
	value *MediaPathDtoPathInfo
	isSet bool
}

func (v NullableMediaPathDtoPathInfo) Get() *MediaPathDtoPathInfo {
	return v.value
}

func (v *NullableMediaPathDtoPathInfo) Set(val *MediaPathDtoPathInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaPathDtoPathInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaPathDtoPathInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaPathDtoPathInfo(val *MediaPathDtoPathInfo) *NullableMediaPathDtoPathInfo {
	return &NullableMediaPathDtoPathInfo{value: val, isSet: true}
}

func (v NullableMediaPathDtoPathInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaPathDtoPathInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


