/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// ForgotPasswordAction the model 'ForgotPasswordAction'
type ForgotPasswordAction string

// List of ForgotPasswordAction
const (
	FORGOTPASSWORDACTION_CONTACT_ADMIN ForgotPasswordAction = "ContactAdmin"
	FORGOTPASSWORDACTION_PIN_CODE ForgotPasswordAction = "PinCode"
	FORGOTPASSWORDACTION_IN_NETWORK_REQUIRED ForgotPasswordAction = "InNetworkRequired"
)

// All allowed values of ForgotPasswordAction enum
var AllowedForgotPasswordActionEnumValues = []ForgotPasswordAction{
	"ContactAdmin",
	"PinCode",
	"InNetworkRequired",
}

func (v *ForgotPasswordAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ForgotPasswordAction(value)
	for _, existing := range AllowedForgotPasswordActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ForgotPasswordAction", value)
}

// NewForgotPasswordActionFromValue returns a pointer to a valid ForgotPasswordAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewForgotPasswordActionFromValue(v string) (*ForgotPasswordAction, error) {
	ev := ForgotPasswordAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ForgotPasswordAction: valid values are %v", v, AllowedForgotPasswordActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ForgotPasswordAction) IsValid() bool {
	for _, existing := range AllowedForgotPasswordActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ForgotPasswordAction value
func (v ForgotPasswordAction) Ptr() *ForgotPasswordAction {
	return &v
}

type NullableForgotPasswordAction struct {
	value *ForgotPasswordAction
	isSet bool
}

func (v NullableForgotPasswordAction) Get() *ForgotPasswordAction {
	return v.value
}

func (v *NullableForgotPasswordAction) Set(val *ForgotPasswordAction) {
	v.value = val
	v.isSet = true
}

func (v NullableForgotPasswordAction) IsSet() bool {
	return v.isSet
}

func (v *NullableForgotPasswordAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForgotPasswordAction(val *ForgotPasswordAction) *NullableForgotPasswordAction {
	return &NullableForgotPasswordAction{value: val, isSet: true}
}

func (v NullableForgotPasswordAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForgotPasswordAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

