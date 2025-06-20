/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UpdateUserPassword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserPassword{}

// UpdateUserPassword The update user password request body.
type UpdateUserPassword struct {
	// Gets or sets the current sha1-hashed password.
	CurrentPassword NullableString `json:"CurrentPassword,omitempty"`
	// Gets or sets the current plain text password.
	CurrentPw NullableString `json:"CurrentPw,omitempty"`
	// Gets or sets the new plain text password.
	NewPw NullableString `json:"NewPw,omitempty"`
	// Gets or sets a value indicating whether to reset the password.
	ResetPassword *bool `json:"ResetPassword,omitempty"`
}

// NewUpdateUserPassword instantiates a new UpdateUserPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserPassword() *UpdateUserPassword {
	this := UpdateUserPassword{}
	return &this
}

// NewUpdateUserPasswordWithDefaults instantiates a new UpdateUserPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserPasswordWithDefaults() *UpdateUserPassword {
	this := UpdateUserPassword{}
	return &this
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserPassword) GetCurrentPassword() string {
	if o == nil || IsNil(o.CurrentPassword.Get()) {
		var ret string
		return ret
	}
	return *o.CurrentPassword.Get()
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserPassword) GetCurrentPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPassword.Get(), o.CurrentPassword.IsSet()
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *UpdateUserPassword) HasCurrentPassword() bool {
	if o != nil && o.CurrentPassword.IsSet() {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given NullableString and assigns it to the CurrentPassword field.
func (o *UpdateUserPassword) SetCurrentPassword(v string) {
	o.CurrentPassword.Set(&v)
}
// SetCurrentPasswordNil sets the value for CurrentPassword to be an explicit nil
func (o *UpdateUserPassword) SetCurrentPasswordNil() {
	o.CurrentPassword.Set(nil)
}

// UnsetCurrentPassword ensures that no value is present for CurrentPassword, not even an explicit nil
func (o *UpdateUserPassword) UnsetCurrentPassword() {
	o.CurrentPassword.Unset()
}

// GetCurrentPw returns the CurrentPw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserPassword) GetCurrentPw() string {
	if o == nil || IsNil(o.CurrentPw.Get()) {
		var ret string
		return ret
	}
	return *o.CurrentPw.Get()
}

// GetCurrentPwOk returns a tuple with the CurrentPw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserPassword) GetCurrentPwOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPw.Get(), o.CurrentPw.IsSet()
}

// HasCurrentPw returns a boolean if a field has been set.
func (o *UpdateUserPassword) HasCurrentPw() bool {
	if o != nil && o.CurrentPw.IsSet() {
		return true
	}

	return false
}

// SetCurrentPw gets a reference to the given NullableString and assigns it to the CurrentPw field.
func (o *UpdateUserPassword) SetCurrentPw(v string) {
	o.CurrentPw.Set(&v)
}
// SetCurrentPwNil sets the value for CurrentPw to be an explicit nil
func (o *UpdateUserPassword) SetCurrentPwNil() {
	o.CurrentPw.Set(nil)
}

// UnsetCurrentPw ensures that no value is present for CurrentPw, not even an explicit nil
func (o *UpdateUserPassword) UnsetCurrentPw() {
	o.CurrentPw.Unset()
}

// GetNewPw returns the NewPw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserPassword) GetNewPw() string {
	if o == nil || IsNil(o.NewPw.Get()) {
		var ret string
		return ret
	}
	return *o.NewPw.Get()
}

// GetNewPwOk returns a tuple with the NewPw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserPassword) GetNewPwOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewPw.Get(), o.NewPw.IsSet()
}

// HasNewPw returns a boolean if a field has been set.
func (o *UpdateUserPassword) HasNewPw() bool {
	if o != nil && o.NewPw.IsSet() {
		return true
	}

	return false
}

// SetNewPw gets a reference to the given NullableString and assigns it to the NewPw field.
func (o *UpdateUserPassword) SetNewPw(v string) {
	o.NewPw.Set(&v)
}
// SetNewPwNil sets the value for NewPw to be an explicit nil
func (o *UpdateUserPassword) SetNewPwNil() {
	o.NewPw.Set(nil)
}

// UnsetNewPw ensures that no value is present for NewPw, not even an explicit nil
func (o *UpdateUserPassword) UnsetNewPw() {
	o.NewPw.Unset()
}

// GetResetPassword returns the ResetPassword field value if set, zero value otherwise.
func (o *UpdateUserPassword) GetResetPassword() bool {
	if o == nil || IsNil(o.ResetPassword) {
		var ret bool
		return ret
	}
	return *o.ResetPassword
}

// GetResetPasswordOk returns a tuple with the ResetPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserPassword) GetResetPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.ResetPassword) {
		return nil, false
	}
	return o.ResetPassword, true
}

// HasResetPassword returns a boolean if a field has been set.
func (o *UpdateUserPassword) HasResetPassword() bool {
	if o != nil && !IsNil(o.ResetPassword) {
		return true
	}

	return false
}

// SetResetPassword gets a reference to the given bool and assigns it to the ResetPassword field.
func (o *UpdateUserPassword) SetResetPassword(v bool) {
	o.ResetPassword = &v
}

func (o UpdateUserPassword) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserPassword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CurrentPassword.IsSet() {
		toSerialize["CurrentPassword"] = o.CurrentPassword.Get()
	}
	if o.CurrentPw.IsSet() {
		toSerialize["CurrentPw"] = o.CurrentPw.Get()
	}
	if o.NewPw.IsSet() {
		toSerialize["NewPw"] = o.NewPw.Get()
	}
	if !IsNil(o.ResetPassword) {
		toSerialize["ResetPassword"] = o.ResetPassword
	}
	return toSerialize, nil
}

type NullableUpdateUserPassword struct {
	value *UpdateUserPassword
	isSet bool
}

func (v NullableUpdateUserPassword) Get() *UpdateUserPassword {
	return v.value
}

func (v *NullableUpdateUserPassword) Set(val *UpdateUserPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserPassword(val *UpdateUserPassword) *NullableUpdateUserPassword {
	return &NullableUpdateUserPassword{value: val, isSet: true}
}

func (v NullableUpdateUserPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


