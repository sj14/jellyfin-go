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

// checks if the StringGroupUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringGroupUpdate{}

// StringGroupUpdate Class GroupUpdate.
type StringGroupUpdate struct {
	// Gets the group identifier.
	GroupId *string `json:"GroupId,omitempty"`
	// Gets the update type.
	Type *GroupUpdateType `json:"Type,omitempty"`
	// Gets the update data.
	Data *string `json:"Data,omitempty"`
}

// NewStringGroupUpdate instantiates a new StringGroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringGroupUpdate() *StringGroupUpdate {
	this := StringGroupUpdate{}
	return &this
}

// NewStringGroupUpdateWithDefaults instantiates a new StringGroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringGroupUpdateWithDefaults() *StringGroupUpdate {
	this := StringGroupUpdate{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *StringGroupUpdate) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringGroupUpdate) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *StringGroupUpdate) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *StringGroupUpdate) SetGroupId(v string) {
	o.GroupId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StringGroupUpdate) GetType() GroupUpdateType {
	if o == nil || IsNil(o.Type) {
		var ret GroupUpdateType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringGroupUpdate) GetTypeOk() (*GroupUpdateType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StringGroupUpdate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given GroupUpdateType and assigns it to the Type field.
func (o *StringGroupUpdate) SetType(v GroupUpdateType) {
	o.Type = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *StringGroupUpdate) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringGroupUpdate) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *StringGroupUpdate) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *StringGroupUpdate) SetData(v string) {
	o.Data = &v
}

func (o StringGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringGroupUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["GroupId"] = o.GroupId
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	return toSerialize, nil
}

type NullableStringGroupUpdate struct {
	value *StringGroupUpdate
	isSet bool
}

func (v NullableStringGroupUpdate) Get() *StringGroupUpdate {
	return v.value
}

func (v *NullableStringGroupUpdate) Set(val *StringGroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStringGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStringGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringGroupUpdate(val *StringGroupUpdate) *NullableStringGroupUpdate {
	return &NullableStringGroupUpdate{value: val, isSet: true}
}

func (v NullableStringGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


