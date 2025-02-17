/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MediaPathDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaPathDto{}

// MediaPathDto Media Path dto.
type MediaPathDto struct {
	// Gets or sets the name of the library.
	Name string `json:"Name"`
	// Gets or sets the path to add.
	Path NullableString `json:"Path,omitempty"`
	// Gets or sets the path info.
	PathInfo NullableMediaPathInfo `json:"PathInfo,omitempty"`
}

type _MediaPathDto MediaPathDto

// NewMediaPathDto instantiates a new MediaPathDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaPathDto(name string) *MediaPathDto {
	this := MediaPathDto{}
	this.Name = name
	return &this
}

// NewMediaPathDtoWithDefaults instantiates a new MediaPathDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaPathDtoWithDefaults() *MediaPathDto {
	this := MediaPathDto{}
	return &this
}

// GetName returns the Name field value
func (o *MediaPathDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MediaPathDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MediaPathDto) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaPathDto) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaPathDto) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *MediaPathDto) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *MediaPathDto) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *MediaPathDto) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *MediaPathDto) UnsetPath() {
	o.Path.Unset()
}

// GetPathInfo returns the PathInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaPathDto) GetPathInfo() MediaPathInfo {
	if o == nil || IsNil(o.PathInfo.Get()) {
		var ret MediaPathInfo
		return ret
	}
	return *o.PathInfo.Get()
}

// GetPathInfoOk returns a tuple with the PathInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaPathDto) GetPathInfoOk() (*MediaPathInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.PathInfo.Get(), o.PathInfo.IsSet()
}

// HasPathInfo returns a boolean if a field has been set.
func (o *MediaPathDto) HasPathInfo() bool {
	if o != nil && o.PathInfo.IsSet() {
		return true
	}

	return false
}

// SetPathInfo gets a reference to the given NullableMediaPathInfo and assigns it to the PathInfo field.
func (o *MediaPathDto) SetPathInfo(v MediaPathInfo) {
	o.PathInfo.Set(&v)
}
// SetPathInfoNil sets the value for PathInfo to be an explicit nil
func (o *MediaPathDto) SetPathInfoNil() {
	o.PathInfo.Set(nil)
}

// UnsetPathInfo ensures that no value is present for PathInfo, not even an explicit nil
func (o *MediaPathDto) UnsetPathInfo() {
	o.PathInfo.Unset()
}

func (o MediaPathDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaPathDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Name"] = o.Name
	if o.Path.IsSet() {
		toSerialize["Path"] = o.Path.Get()
	}
	if o.PathInfo.IsSet() {
		toSerialize["PathInfo"] = o.PathInfo.Get()
	}
	return toSerialize, nil
}

func (o *MediaPathDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMediaPathDto := _MediaPathDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMediaPathDto)

	if err != nil {
		return err
	}

	*o = MediaPathDto(varMediaPathDto)

	return err
}

type NullableMediaPathDto struct {
	value *MediaPathDto
	isSet bool
}

func (v NullableMediaPathDto) Get() *MediaPathDto {
	return v.value
}

func (v *NullableMediaPathDto) Set(val *MediaPathDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaPathDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaPathDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaPathDto(val *MediaPathDto) *NullableMediaPathDto {
	return &NullableMediaPathDto{value: val, isSet: true}
}

func (v NullableMediaPathDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaPathDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


