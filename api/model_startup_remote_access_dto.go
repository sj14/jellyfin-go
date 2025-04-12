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

// checks if the StartupRemoteAccessDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartupRemoteAccessDto{}

// StartupRemoteAccessDto Startup remote access dto.
type StartupRemoteAccessDto struct {
	// Gets or sets a value indicating whether enable remote access.
	EnableRemoteAccess bool `json:"EnableRemoteAccess"`
	// Gets or sets a value indicating whether enable automatic port mapping.
	EnableAutomaticPortMapping bool `json:"EnableAutomaticPortMapping"`
}

type _StartupRemoteAccessDto StartupRemoteAccessDto

// NewStartupRemoteAccessDto instantiates a new StartupRemoteAccessDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartupRemoteAccessDto(enableRemoteAccess bool, enableAutomaticPortMapping bool) *StartupRemoteAccessDto {
	this := StartupRemoteAccessDto{}
	this.EnableRemoteAccess = enableRemoteAccess
	this.EnableAutomaticPortMapping = enableAutomaticPortMapping
	return &this
}

// NewStartupRemoteAccessDtoWithDefaults instantiates a new StartupRemoteAccessDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartupRemoteAccessDtoWithDefaults() *StartupRemoteAccessDto {
	this := StartupRemoteAccessDto{}
	return &this
}

// GetEnableRemoteAccess returns the EnableRemoteAccess field value
func (o *StartupRemoteAccessDto) GetEnableRemoteAccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableRemoteAccess
}

// GetEnableRemoteAccessOk returns a tuple with the EnableRemoteAccess field value
// and a boolean to check if the value has been set.
func (o *StartupRemoteAccessDto) GetEnableRemoteAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableRemoteAccess, true
}

// SetEnableRemoteAccess sets field value
func (o *StartupRemoteAccessDto) SetEnableRemoteAccess(v bool) {
	o.EnableRemoteAccess = v
}

// GetEnableAutomaticPortMapping returns the EnableAutomaticPortMapping field value
func (o *StartupRemoteAccessDto) GetEnableAutomaticPortMapping() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableAutomaticPortMapping
}

// GetEnableAutomaticPortMappingOk returns a tuple with the EnableAutomaticPortMapping field value
// and a boolean to check if the value has been set.
func (o *StartupRemoteAccessDto) GetEnableAutomaticPortMappingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableAutomaticPortMapping, true
}

// SetEnableAutomaticPortMapping sets field value
func (o *StartupRemoteAccessDto) SetEnableAutomaticPortMapping(v bool) {
	o.EnableAutomaticPortMapping = v
}

func (o StartupRemoteAccessDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartupRemoteAccessDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["EnableRemoteAccess"] = o.EnableRemoteAccess
	toSerialize["EnableAutomaticPortMapping"] = o.EnableAutomaticPortMapping
	return toSerialize, nil
}

func (o *StartupRemoteAccessDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"EnableRemoteAccess",
		"EnableAutomaticPortMapping",
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

	varStartupRemoteAccessDto := _StartupRemoteAccessDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStartupRemoteAccessDto)

	if err != nil {
		return err
	}

	*o = StartupRemoteAccessDto(varStartupRemoteAccessDto)

	return err
}

type NullableStartupRemoteAccessDto struct {
	value *StartupRemoteAccessDto
	isSet bool
}

func (v NullableStartupRemoteAccessDto) Get() *StartupRemoteAccessDto {
	return v.value
}

func (v *NullableStartupRemoteAccessDto) Set(val *StartupRemoteAccessDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStartupRemoteAccessDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStartupRemoteAccessDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartupRemoteAccessDto(val *StartupRemoteAccessDto) *NullableStartupRemoteAccessDto {
	return &NullableStartupRemoteAccessDto{value: val, isSet: true}
}

func (v NullableStartupRemoteAccessDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartupRemoteAccessDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


