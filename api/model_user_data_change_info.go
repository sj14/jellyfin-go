/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UserDataChangeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDataChangeInfo{}

// UserDataChangeInfo Class UserDataChangeInfo.
type UserDataChangeInfo struct {
	// Gets or sets the user id.
	UserId NullableString `json:"UserId,omitempty"`
	// Gets or sets the user data list.
	UserDataList []UserItemDataDto `json:"UserDataList,omitempty"`
}

// NewUserDataChangeInfo instantiates a new UserDataChangeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDataChangeInfo() *UserDataChangeInfo {
	this := UserDataChangeInfo{}
	return &this
}

// NewUserDataChangeInfoWithDefaults instantiates a new UserDataChangeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataChangeInfoWithDefaults() *UserDataChangeInfo {
	this := UserDataChangeInfo{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDataChangeInfo) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDataChangeInfo) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *UserDataChangeInfo) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *UserDataChangeInfo) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *UserDataChangeInfo) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *UserDataChangeInfo) UnsetUserId() {
	o.UserId.Unset()
}

// GetUserDataList returns the UserDataList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDataChangeInfo) GetUserDataList() []UserItemDataDto {
	if o == nil {
		var ret []UserItemDataDto
		return ret
	}
	return o.UserDataList
}

// GetUserDataListOk returns a tuple with the UserDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDataChangeInfo) GetUserDataListOk() ([]UserItemDataDto, bool) {
	if o == nil || IsNil(o.UserDataList) {
		return nil, false
	}
	return o.UserDataList, true
}

// HasUserDataList returns a boolean if a field has been set.
func (o *UserDataChangeInfo) HasUserDataList() bool {
	if o != nil && !IsNil(o.UserDataList) {
		return true
	}

	return false
}

// SetUserDataList gets a reference to the given []UserItemDataDto and assigns it to the UserDataList field.
func (o *UserDataChangeInfo) SetUserDataList(v []UserItemDataDto) {
	o.UserDataList = v
}

func (o UserDataChangeInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDataChangeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["UserId"] = o.UserId.Get()
	}
	if o.UserDataList != nil {
		toSerialize["UserDataList"] = o.UserDataList
	}
	return toSerialize, nil
}

type NullableUserDataChangeInfo struct {
	value *UserDataChangeInfo
	isSet bool
}

func (v NullableUserDataChangeInfo) Get() *UserDataChangeInfo {
	return v.value
}

func (v *NullableUserDataChangeInfo) Set(val *UserDataChangeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataChangeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataChangeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataChangeInfo(val *UserDataChangeInfo) *NullableUserDataChangeInfo {
	return &NullableUserDataChangeInfo{value: val, isSet: true}
}

func (v NullableUserDataChangeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataChangeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


