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

// checks if the SessionUserInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionUserInfo{}

// SessionUserInfo Class SessionUserInfo.
type SessionUserInfo struct {
	// Gets or sets the user identifier.
	UserId *string `json:"UserId,omitempty"`
	// Gets or sets the name of the user.
	UserName NullableString `json:"UserName,omitempty"`
}

// NewSessionUserInfo instantiates a new SessionUserInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionUserInfo() *SessionUserInfo {
	this := SessionUserInfo{}
	return &this
}

// NewSessionUserInfoWithDefaults instantiates a new SessionUserInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionUserInfoWithDefaults() *SessionUserInfo {
	this := SessionUserInfo{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *SessionUserInfo) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionUserInfo) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *SessionUserInfo) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *SessionUserInfo) SetUserId(v string) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionUserInfo) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionUserInfo) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *SessionUserInfo) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *SessionUserInfo) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *SessionUserInfo) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *SessionUserInfo) UnsetUserName() {
	o.UserName.Unset()
}

func (o SessionUserInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionUserInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["UserId"] = o.UserId
	}
	if o.UserName.IsSet() {
		toSerialize["UserName"] = o.UserName.Get()
	}
	return toSerialize, nil
}

type NullableSessionUserInfo struct {
	value *SessionUserInfo
	isSet bool
}

func (v NullableSessionUserInfo) Get() *SessionUserInfo {
	return v.value
}

func (v *NullableSessionUserInfo) Set(val *SessionUserInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionUserInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionUserInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionUserInfo(val *SessionUserInfo) *NullableSessionUserInfo {
	return &NullableSessionUserInfo{value: val, isSet: true}
}

func (v NullableSessionUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionUserInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


