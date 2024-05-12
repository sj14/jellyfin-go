/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AuthenticationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationResult{}

// AuthenticationResult struct for AuthenticationResult
type AuthenticationResult struct {
	User NullableAuthenticationResultUser `json:"User,omitempty"`
	SessionInfo NullableAuthenticationResultSessionInfo `json:"SessionInfo,omitempty"`
	AccessToken NullableString `json:"AccessToken,omitempty"`
	ServerId NullableString `json:"ServerId,omitempty"`
}

// NewAuthenticationResult instantiates a new AuthenticationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationResult() *AuthenticationResult {
	this := AuthenticationResult{}
	return &this
}

// NewAuthenticationResultWithDefaults instantiates a new AuthenticationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationResultWithDefaults() *AuthenticationResult {
	this := AuthenticationResult{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationResult) GetUser() AuthenticationResultUser {
	if o == nil || IsNil(o.User.Get()) {
		var ret AuthenticationResultUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationResult) GetUserOk() (*AuthenticationResultUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *AuthenticationResult) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableAuthenticationResultUser and assigns it to the User field.
func (o *AuthenticationResult) SetUser(v AuthenticationResultUser) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *AuthenticationResult) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *AuthenticationResult) UnsetUser() {
	o.User.Unset()
}

// GetSessionInfo returns the SessionInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationResult) GetSessionInfo() AuthenticationResultSessionInfo {
	if o == nil || IsNil(o.SessionInfo.Get()) {
		var ret AuthenticationResultSessionInfo
		return ret
	}
	return *o.SessionInfo.Get()
}

// GetSessionInfoOk returns a tuple with the SessionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationResult) GetSessionInfoOk() (*AuthenticationResultSessionInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.SessionInfo.Get(), o.SessionInfo.IsSet()
}

// HasSessionInfo returns a boolean if a field has been set.
func (o *AuthenticationResult) HasSessionInfo() bool {
	if o != nil && o.SessionInfo.IsSet() {
		return true
	}

	return false
}

// SetSessionInfo gets a reference to the given NullableAuthenticationResultSessionInfo and assigns it to the SessionInfo field.
func (o *AuthenticationResult) SetSessionInfo(v AuthenticationResultSessionInfo) {
	o.SessionInfo.Set(&v)
}
// SetSessionInfoNil sets the value for SessionInfo to be an explicit nil
func (o *AuthenticationResult) SetSessionInfoNil() {
	o.SessionInfo.Set(nil)
}

// UnsetSessionInfo ensures that no value is present for SessionInfo, not even an explicit nil
func (o *AuthenticationResult) UnsetSessionInfo() {
	o.SessionInfo.Unset()
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationResult) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationResult) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AuthenticationResult) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *AuthenticationResult) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *AuthenticationResult) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *AuthenticationResult) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetServerId returns the ServerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationResult) GetServerId() string {
	if o == nil || IsNil(o.ServerId.Get()) {
		var ret string
		return ret
	}
	return *o.ServerId.Get()
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationResult) GetServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerId.Get(), o.ServerId.IsSet()
}

// HasServerId returns a boolean if a field has been set.
func (o *AuthenticationResult) HasServerId() bool {
	if o != nil && o.ServerId.IsSet() {
		return true
	}

	return false
}

// SetServerId gets a reference to the given NullableString and assigns it to the ServerId field.
func (o *AuthenticationResult) SetServerId(v string) {
	o.ServerId.Set(&v)
}
// SetServerIdNil sets the value for ServerId to be an explicit nil
func (o *AuthenticationResult) SetServerIdNil() {
	o.ServerId.Set(nil)
}

// UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
func (o *AuthenticationResult) UnsetServerId() {
	o.ServerId.Unset()
}

func (o AuthenticationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.User.IsSet() {
		toSerialize["User"] = o.User.Get()
	}
	if o.SessionInfo.IsSet() {
		toSerialize["SessionInfo"] = o.SessionInfo.Get()
	}
	if o.AccessToken.IsSet() {
		toSerialize["AccessToken"] = o.AccessToken.Get()
	}
	if o.ServerId.IsSet() {
		toSerialize["ServerId"] = o.ServerId.Get()
	}
	return toSerialize, nil
}

type NullableAuthenticationResult struct {
	value *AuthenticationResult
	isSet bool
}

func (v NullableAuthenticationResult) Get() *AuthenticationResult {
	return v.value
}

func (v *NullableAuthenticationResult) Set(val *AuthenticationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationResult(val *AuthenticationResult) *NullableAuthenticationResult {
	return &NullableAuthenticationResult{value: val, isSet: true}
}

func (v NullableAuthenticationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


