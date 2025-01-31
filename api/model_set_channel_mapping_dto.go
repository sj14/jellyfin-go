/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SetChannelMappingDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetChannelMappingDto{}

// SetChannelMappingDto Set channel mapping dto.
type SetChannelMappingDto struct {
	// Gets or sets the provider id.
	ProviderId string `json:"ProviderId"`
	// Gets or sets the tuner channel id.
	TunerChannelId string `json:"TunerChannelId"`
	// Gets or sets the provider channel id.
	ProviderChannelId string `json:"ProviderChannelId"`
}

type _SetChannelMappingDto SetChannelMappingDto

// NewSetChannelMappingDto instantiates a new SetChannelMappingDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetChannelMappingDto(providerId string, tunerChannelId string, providerChannelId string) *SetChannelMappingDto {
	this := SetChannelMappingDto{}
	this.ProviderId = providerId
	this.TunerChannelId = tunerChannelId
	this.ProviderChannelId = providerChannelId
	return &this
}

// NewSetChannelMappingDtoWithDefaults instantiates a new SetChannelMappingDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetChannelMappingDtoWithDefaults() *SetChannelMappingDto {
	this := SetChannelMappingDto{}
	return &this
}

// GetProviderId returns the ProviderId field value
func (o *SetChannelMappingDto) GetProviderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value
// and a boolean to check if the value has been set.
func (o *SetChannelMappingDto) GetProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderId, true
}

// SetProviderId sets field value
func (o *SetChannelMappingDto) SetProviderId(v string) {
	o.ProviderId = v
}

// GetTunerChannelId returns the TunerChannelId field value
func (o *SetChannelMappingDto) GetTunerChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TunerChannelId
}

// GetTunerChannelIdOk returns a tuple with the TunerChannelId field value
// and a boolean to check if the value has been set.
func (o *SetChannelMappingDto) GetTunerChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TunerChannelId, true
}

// SetTunerChannelId sets field value
func (o *SetChannelMappingDto) SetTunerChannelId(v string) {
	o.TunerChannelId = v
}

// GetProviderChannelId returns the ProviderChannelId field value
func (o *SetChannelMappingDto) GetProviderChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderChannelId
}

// GetProviderChannelIdOk returns a tuple with the ProviderChannelId field value
// and a boolean to check if the value has been set.
func (o *SetChannelMappingDto) GetProviderChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderChannelId, true
}

// SetProviderChannelId sets field value
func (o *SetChannelMappingDto) SetProviderChannelId(v string) {
	o.ProviderChannelId = v
}

func (o SetChannelMappingDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetChannelMappingDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ProviderId"] = o.ProviderId
	toSerialize["TunerChannelId"] = o.TunerChannelId
	toSerialize["ProviderChannelId"] = o.ProviderChannelId
	return toSerialize, nil
}

func (o *SetChannelMappingDto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ProviderId",
		"TunerChannelId",
		"ProviderChannelId",
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

	varSetChannelMappingDto := _SetChannelMappingDto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetChannelMappingDto)

	if err != nil {
		return err
	}

	*o = SetChannelMappingDto(varSetChannelMappingDto)

	return err
}

type NullableSetChannelMappingDto struct {
	value *SetChannelMappingDto
	isSet bool
}

func (v NullableSetChannelMappingDto) Get() *SetChannelMappingDto {
	return v.value
}

func (v *NullableSetChannelMappingDto) Set(val *SetChannelMappingDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSetChannelMappingDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSetChannelMappingDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetChannelMappingDto(val *SetChannelMappingDto) *NullableSetChannelMappingDto {
	return &NullableSetChannelMappingDto{value: val, isSet: true}
}

func (v NullableSetChannelMappingDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetChannelMappingDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


