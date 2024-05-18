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

// checks if the UserDataChangedMessageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDataChangedMessageData{}

// UserDataChangedMessageData Class UserDataChangeInfo.
type UserDataChangedMessageData struct {
	// Gets or sets the user id.
	UserId NullableString `json:"UserId,omitempty"`
	// Gets or sets the user data list.
	UserDataList []UserItemDataDto `json:"UserDataList,omitempty"`
}

// NewUserDataChangedMessageData instantiates a new UserDataChangedMessageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDataChangedMessageData() *UserDataChangedMessageData {
	this := UserDataChangedMessageData{}
	return &this
}

// NewUserDataChangedMessageDataWithDefaults instantiates a new UserDataChangedMessageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataChangedMessageDataWithDefaults() *UserDataChangedMessageData {
	this := UserDataChangedMessageData{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDataChangedMessageData) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDataChangedMessageData) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *UserDataChangedMessageData) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *UserDataChangedMessageData) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *UserDataChangedMessageData) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *UserDataChangedMessageData) UnsetUserId() {
	o.UserId.Unset()
}

// GetUserDataList returns the UserDataList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDataChangedMessageData) GetUserDataList() []UserItemDataDto {
	if o == nil {
		var ret []UserItemDataDto
		return ret
	}
	return o.UserDataList
}

// GetUserDataListOk returns a tuple with the UserDataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDataChangedMessageData) GetUserDataListOk() ([]UserItemDataDto, bool) {
	if o == nil || IsNil(o.UserDataList) {
		return nil, false
	}
	return o.UserDataList, true
}

// HasUserDataList returns a boolean if a field has been set.
func (o *UserDataChangedMessageData) HasUserDataList() bool {
	if o != nil && !IsNil(o.UserDataList) {
		return true
	}

	return false
}

// SetUserDataList gets a reference to the given []UserItemDataDto and assigns it to the UserDataList field.
func (o *UserDataChangedMessageData) SetUserDataList(v []UserItemDataDto) {
	o.UserDataList = v
}

func (o UserDataChangedMessageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDataChangedMessageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["UserId"] = o.UserId.Get()
	}
	if o.UserDataList != nil {
		toSerialize["UserDataList"] = o.UserDataList
	}
	return toSerialize, nil
}

type NullableUserDataChangedMessageData struct {
	value *UserDataChangedMessageData
	isSet bool
}

func (v NullableUserDataChangedMessageData) Get() *UserDataChangedMessageData {
	return v.value
}

func (v *NullableUserDataChangedMessageData) Set(val *UserDataChangedMessageData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataChangedMessageData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataChangedMessageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataChangedMessageData(val *UserDataChangedMessageData) *NullableUserDataChangedMessageData {
	return &NullableUserDataChangedMessageData{value: val, isSet: true}
}

func (v NullableUserDataChangedMessageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataChangedMessageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


