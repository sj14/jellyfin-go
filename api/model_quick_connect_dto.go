/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the QuickConnectDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickConnectDto{}

// QuickConnectDto The quick connect request body.
type QuickConnectDto struct {
	// Gets or sets the quick connect secret.
	Secret string `json:"Secret"`
}

type _QuickConnectDto QuickConnectDto

// NewQuickConnectDto instantiates a new QuickConnectDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickConnectDto(secret string) *QuickConnectDto {
	this := QuickConnectDto{}
	this.Secret = secret
	return &this
}

// NewQuickConnectDtoWithDefaults instantiates a new QuickConnectDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickConnectDtoWithDefaults() *QuickConnectDto {
	this := QuickConnectDto{}
	return &this
}

// GetSecret returns the Secret field value
func (o *QuickConnectDto) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *QuickConnectDto) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *QuickConnectDto) SetSecret(v string) {
	o.Secret = v
}

func (o QuickConnectDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickConnectDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Secret"] = o.Secret
	return toSerialize, nil
}

func (o *QuickConnectDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Secret",
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

	varQuickConnectDto := _QuickConnectDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQuickConnectDto)

	if err != nil {
		return err
	}

	*o = QuickConnectDto(varQuickConnectDto)

	return err
}

type NullableQuickConnectDto struct {
	value *QuickConnectDto
	isSet bool
}

func (v NullableQuickConnectDto) Get() *QuickConnectDto {
	return v.value
}

func (v *NullableQuickConnectDto) Set(val *QuickConnectDto) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickConnectDto) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickConnectDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickConnectDto(val *QuickConnectDto) *NullableQuickConnectDto {
	return &NullableQuickConnectDto{value: val, isSet: true}
}

func (v NullableQuickConnectDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickConnectDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


