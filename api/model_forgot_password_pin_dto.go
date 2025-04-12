/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ForgotPasswordPinDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForgotPasswordPinDto{}

// ForgotPasswordPinDto Forgot Password Pin enter request body DTO.
type ForgotPasswordPinDto struct {
	// Gets or sets the entered pin to have the password reset.
	Pin string `json:"Pin"`
}

type _ForgotPasswordPinDto ForgotPasswordPinDto

// NewForgotPasswordPinDto instantiates a new ForgotPasswordPinDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForgotPasswordPinDto(pin string) *ForgotPasswordPinDto {
	this := ForgotPasswordPinDto{}
	this.Pin = pin
	return &this
}

// NewForgotPasswordPinDtoWithDefaults instantiates a new ForgotPasswordPinDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForgotPasswordPinDtoWithDefaults() *ForgotPasswordPinDto {
	this := ForgotPasswordPinDto{}
	return &this
}

// GetPin returns the Pin field value
func (o *ForgotPasswordPinDto) GetPin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pin
}

// GetPinOk returns a tuple with the Pin field value
// and a boolean to check if the value has been set.
func (o *ForgotPasswordPinDto) GetPinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pin, true
}

// SetPin sets field value
func (o *ForgotPasswordPinDto) SetPin(v string) {
	o.Pin = v
}

func (o ForgotPasswordPinDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForgotPasswordPinDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Pin"] = o.Pin
	return toSerialize, nil
}

func (o *ForgotPasswordPinDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Pin",
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

	varForgotPasswordPinDto := _ForgotPasswordPinDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varForgotPasswordPinDto)

	if err != nil {
		return err
	}

	*o = ForgotPasswordPinDto(varForgotPasswordPinDto)

	return err
}

type NullableForgotPasswordPinDto struct {
	value *ForgotPasswordPinDto
	isSet bool
}

func (v NullableForgotPasswordPinDto) Get() *ForgotPasswordPinDto {
	return v.value
}

func (v *NullableForgotPasswordPinDto) Set(val *ForgotPasswordPinDto) {
	v.value = val
	v.isSet = true
}

func (v NullableForgotPasswordPinDto) IsSet() bool {
	return v.isSet
}

func (v *NullableForgotPasswordPinDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForgotPasswordPinDto(val *ForgotPasswordPinDto) *NullableForgotPasswordPinDto {
	return &NullableForgotPasswordPinDto{value: val, isSet: true}
}

func (v NullableForgotPasswordPinDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForgotPasswordPinDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


