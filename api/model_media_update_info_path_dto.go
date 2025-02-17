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

// checks if the MediaUpdateInfoPathDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaUpdateInfoPathDto{}

// MediaUpdateInfoPathDto The media update info path.
type MediaUpdateInfoPathDto struct {
	// Gets or sets media path.
	Path NullableString `json:"Path,omitempty"`
	// Gets or sets media update type.  Created, Modified, Deleted.
	UpdateType NullableString `json:"UpdateType,omitempty"`
}

// NewMediaUpdateInfoPathDto instantiates a new MediaUpdateInfoPathDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaUpdateInfoPathDto() *MediaUpdateInfoPathDto {
	this := MediaUpdateInfoPathDto{}
	return &this
}

// NewMediaUpdateInfoPathDtoWithDefaults instantiates a new MediaUpdateInfoPathDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaUpdateInfoPathDtoWithDefaults() *MediaUpdateInfoPathDto {
	this := MediaUpdateInfoPathDto{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaUpdateInfoPathDto) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaUpdateInfoPathDto) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *MediaUpdateInfoPathDto) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *MediaUpdateInfoPathDto) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *MediaUpdateInfoPathDto) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *MediaUpdateInfoPathDto) UnsetPath() {
	o.Path.Unset()
}

// GetUpdateType returns the UpdateType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaUpdateInfoPathDto) GetUpdateType() string {
	if o == nil || IsNil(o.UpdateType.Get()) {
		var ret string
		return ret
	}
	return *o.UpdateType.Get()
}

// GetUpdateTypeOk returns a tuple with the UpdateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaUpdateInfoPathDto) GetUpdateTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdateType.Get(), o.UpdateType.IsSet()
}

// HasUpdateType returns a boolean if a field has been set.
func (o *MediaUpdateInfoPathDto) HasUpdateType() bool {
	if o != nil && o.UpdateType.IsSet() {
		return true
	}

	return false
}

// SetUpdateType gets a reference to the given NullableString and assigns it to the UpdateType field.
func (o *MediaUpdateInfoPathDto) SetUpdateType(v string) {
	o.UpdateType.Set(&v)
}
// SetUpdateTypeNil sets the value for UpdateType to be an explicit nil
func (o *MediaUpdateInfoPathDto) SetUpdateTypeNil() {
	o.UpdateType.Set(nil)
}

// UnsetUpdateType ensures that no value is present for UpdateType, not even an explicit nil
func (o *MediaUpdateInfoPathDto) UnsetUpdateType() {
	o.UpdateType.Unset()
}

func (o MediaUpdateInfoPathDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaUpdateInfoPathDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Path.IsSet() {
		toSerialize["Path"] = o.Path.Get()
	}
	if o.UpdateType.IsSet() {
		toSerialize["UpdateType"] = o.UpdateType.Get()
	}
	return toSerialize, nil
}

type NullableMediaUpdateInfoPathDto struct {
	value *MediaUpdateInfoPathDto
	isSet bool
}

func (v NullableMediaUpdateInfoPathDto) Get() *MediaUpdateInfoPathDto {
	return v.value
}

func (v *NullableMediaUpdateInfoPathDto) Set(val *MediaUpdateInfoPathDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaUpdateInfoPathDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaUpdateInfoPathDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaUpdateInfoPathDto(val *MediaUpdateInfoPathDto) *NullableMediaUpdateInfoPathDto {
	return &NullableMediaUpdateInfoPathDto{value: val, isSet: true}
}

func (v NullableMediaUpdateInfoPathDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaUpdateInfoPathDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


