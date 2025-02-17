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

// checks if the ForgotPasswordDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForgotPasswordDto{}

// ForgotPasswordDto Forgot Password request body DTO.
type ForgotPasswordDto struct {
	// Gets or sets the entered username to have its password reset.
	EnteredUsername string `json:"EnteredUsername"`
}

type _ForgotPasswordDto ForgotPasswordDto

// NewForgotPasswordDto instantiates a new ForgotPasswordDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForgotPasswordDto(enteredUsername string) *ForgotPasswordDto {
	this := ForgotPasswordDto{}
	this.EnteredUsername = enteredUsername
	return &this
}

// NewForgotPasswordDtoWithDefaults instantiates a new ForgotPasswordDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForgotPasswordDtoWithDefaults() *ForgotPasswordDto {
	this := ForgotPasswordDto{}
	return &this
}

// GetEnteredUsername returns the EnteredUsername field value
func (o *ForgotPasswordDto) GetEnteredUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnteredUsername
}

// GetEnteredUsernameOk returns a tuple with the EnteredUsername field value
// and a boolean to check if the value has been set.
func (o *ForgotPasswordDto) GetEnteredUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnteredUsername, true
}

// SetEnteredUsername sets field value
func (o *ForgotPasswordDto) SetEnteredUsername(v string) {
	o.EnteredUsername = v
}

func (o ForgotPasswordDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForgotPasswordDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["EnteredUsername"] = o.EnteredUsername
	return toSerialize, nil
}

func (o *ForgotPasswordDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"EnteredUsername",
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

	varForgotPasswordDto := _ForgotPasswordDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varForgotPasswordDto)

	if err != nil {
		return err
	}

	*o = ForgotPasswordDto(varForgotPasswordDto)

	return err
}

type NullableForgotPasswordDto struct {
	value *ForgotPasswordDto
	isSet bool
}

func (v NullableForgotPasswordDto) Get() *ForgotPasswordDto {
	return v.value
}

func (v *NullableForgotPasswordDto) Set(val *ForgotPasswordDto) {
	v.value = val
	v.isSet = true
}

func (v NullableForgotPasswordDto) IsSet() bool {
	return v.isSet
}

func (v *NullableForgotPasswordDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForgotPasswordDto(val *ForgotPasswordDto) *NullableForgotPasswordDto {
	return &NullableForgotPasswordDto{value: val, isSet: true}
}

func (v NullableForgotPasswordDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForgotPasswordDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


