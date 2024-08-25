/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AuthenticateUserByName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticateUserByName{}

// AuthenticateUserByName The authenticate user by name request body.
type AuthenticateUserByName struct {
	// Gets or sets the username.
	Username NullableString `json:"Username,omitempty"`
	// Gets or sets the plain text password.
	Pw NullableString `json:"Pw,omitempty"`
}

// NewAuthenticateUserByName instantiates a new AuthenticateUserByName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticateUserByName() *AuthenticateUserByName {
	this := AuthenticateUserByName{}
	return &this
}

// NewAuthenticateUserByNameWithDefaults instantiates a new AuthenticateUserByName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticateUserByNameWithDefaults() *AuthenticateUserByName {
	this := AuthenticateUserByName{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticateUserByName) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticateUserByName) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *AuthenticateUserByName) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *AuthenticateUserByName) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *AuthenticateUserByName) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *AuthenticateUserByName) UnsetUsername() {
	o.Username.Unset()
}

// GetPw returns the Pw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticateUserByName) GetPw() string {
	if o == nil || IsNil(o.Pw.Get()) {
		var ret string
		return ret
	}
	return *o.Pw.Get()
}

// GetPwOk returns a tuple with the Pw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticateUserByName) GetPwOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pw.Get(), o.Pw.IsSet()
}

// HasPw returns a boolean if a field has been set.
func (o *AuthenticateUserByName) HasPw() bool {
	if o != nil && o.Pw.IsSet() {
		return true
	}

	return false
}

// SetPw gets a reference to the given NullableString and assigns it to the Pw field.
func (o *AuthenticateUserByName) SetPw(v string) {
	o.Pw.Set(&v)
}
// SetPwNil sets the value for Pw to be an explicit nil
func (o *AuthenticateUserByName) SetPwNil() {
	o.Pw.Set(nil)
}

// UnsetPw ensures that no value is present for Pw, not even an explicit nil
func (o *AuthenticateUserByName) UnsetPw() {
	o.Pw.Unset()
}

func (o AuthenticateUserByName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticateUserByName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Username.IsSet() {
		toSerialize["Username"] = o.Username.Get()
	}
	if o.Pw.IsSet() {
		toSerialize["Pw"] = o.Pw.Get()
	}
	return toSerialize, nil
}

type NullableAuthenticateUserByName struct {
	value *AuthenticateUserByName
	isSet bool
}

func (v NullableAuthenticateUserByName) Get() *AuthenticateUserByName {
	return v.value
}

func (v *NullableAuthenticateUserByName) Set(val *AuthenticateUserByName) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticateUserByName) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticateUserByName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticateUserByName(val *AuthenticateUserByName) *NullableAuthenticateUserByName {
	return &NullableAuthenticateUserByName{value: val, isSet: true}
}

func (v NullableAuthenticateUserByName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticateUserByName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


