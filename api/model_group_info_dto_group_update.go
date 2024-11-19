/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the GroupInfoDtoGroupUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupInfoDtoGroupUpdate{}

// GroupInfoDtoGroupUpdate Class GroupUpdate.
type GroupInfoDtoGroupUpdate struct {
	// Gets the group identifier.
	GroupId *string `json:"GroupId,omitempty"`
	// Gets the update type.
	Type *GroupUpdateType `json:"Type,omitempty"`
	// Gets the update data.
	Data *GroupInfoDto `json:"Data,omitempty"`
}

// NewGroupInfoDtoGroupUpdate instantiates a new GroupInfoDtoGroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupInfoDtoGroupUpdate() *GroupInfoDtoGroupUpdate {
	this := GroupInfoDtoGroupUpdate{}
	return &this
}

// NewGroupInfoDtoGroupUpdateWithDefaults instantiates a new GroupInfoDtoGroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupInfoDtoGroupUpdateWithDefaults() *GroupInfoDtoGroupUpdate {
	this := GroupInfoDtoGroupUpdate{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GroupInfoDtoGroupUpdate) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInfoDtoGroupUpdate) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupInfoDtoGroupUpdate) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupInfoDtoGroupUpdate) SetGroupId(v string) {
	o.GroupId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GroupInfoDtoGroupUpdate) GetType() GroupUpdateType {
	if o == nil || IsNil(o.Type) {
		var ret GroupUpdateType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInfoDtoGroupUpdate) GetTypeOk() (*GroupUpdateType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GroupInfoDtoGroupUpdate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given GroupUpdateType and assigns it to the Type field.
func (o *GroupInfoDtoGroupUpdate) SetType(v GroupUpdateType) {
	o.Type = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GroupInfoDtoGroupUpdate) GetData() GroupInfoDto {
	if o == nil || IsNil(o.Data) {
		var ret GroupInfoDto
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInfoDtoGroupUpdate) GetDataOk() (*GroupInfoDto, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GroupInfoDtoGroupUpdate) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GroupInfoDto and assigns it to the Data field.
func (o *GroupInfoDtoGroupUpdate) SetData(v GroupInfoDto) {
	o.Data = &v
}

func (o GroupInfoDtoGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupInfoDtoGroupUpdate) ToMap() (map[string]interface{}, error) {
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

type NullableGroupInfoDtoGroupUpdate struct {
	value *GroupInfoDtoGroupUpdate
	isSet bool
}

func (v NullableGroupInfoDtoGroupUpdate) Get() *GroupInfoDtoGroupUpdate {
	return v.value
}

func (v *NullableGroupInfoDtoGroupUpdate) Set(val *GroupInfoDtoGroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupInfoDtoGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupInfoDtoGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupInfoDtoGroupUpdate(val *GroupInfoDtoGroupUpdate) *NullableGroupInfoDtoGroupUpdate {
	return &NullableGroupInfoDtoGroupUpdate{value: val, isSet: true}
}

func (v NullableGroupInfoDtoGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupInfoDtoGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


