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

// checks if the ChannelMappingOptionsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelMappingOptionsDto{}

// ChannelMappingOptionsDto Channel mapping options dto.
type ChannelMappingOptionsDto struct {
	// Gets or sets list of tuner channels.
	TunerChannels []TunerChannelMapping `json:"TunerChannels,omitempty"`
	// Gets or sets list of provider channels.
	ProviderChannels []NameIdPair `json:"ProviderChannels,omitempty"`
	// Gets or sets list of mappings.
	Mappings []NameValuePair `json:"Mappings,omitempty"`
	// Gets or sets provider name.
	ProviderName NullableString `json:"ProviderName,omitempty"`
}

// NewChannelMappingOptionsDto instantiates a new ChannelMappingOptionsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelMappingOptionsDto() *ChannelMappingOptionsDto {
	this := ChannelMappingOptionsDto{}
	return &this
}

// NewChannelMappingOptionsDtoWithDefaults instantiates a new ChannelMappingOptionsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelMappingOptionsDtoWithDefaults() *ChannelMappingOptionsDto {
	this := ChannelMappingOptionsDto{}
	return &this
}

// GetTunerChannels returns the TunerChannels field value if set, zero value otherwise.
func (o *ChannelMappingOptionsDto) GetTunerChannels() []TunerChannelMapping {
	if o == nil || IsNil(o.TunerChannels) {
		var ret []TunerChannelMapping
		return ret
	}
	return o.TunerChannels
}

// GetTunerChannelsOk returns a tuple with the TunerChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelMappingOptionsDto) GetTunerChannelsOk() ([]TunerChannelMapping, bool) {
	if o == nil || IsNil(o.TunerChannels) {
		return nil, false
	}
	return o.TunerChannels, true
}

// HasTunerChannels returns a boolean if a field has been set.
func (o *ChannelMappingOptionsDto) HasTunerChannels() bool {
	if o != nil && !IsNil(o.TunerChannels) {
		return true
	}

	return false
}

// SetTunerChannels gets a reference to the given []TunerChannelMapping and assigns it to the TunerChannels field.
func (o *ChannelMappingOptionsDto) SetTunerChannels(v []TunerChannelMapping) {
	o.TunerChannels = v
}

// GetProviderChannels returns the ProviderChannels field value if set, zero value otherwise.
func (o *ChannelMappingOptionsDto) GetProviderChannels() []NameIdPair {
	if o == nil || IsNil(o.ProviderChannels) {
		var ret []NameIdPair
		return ret
	}
	return o.ProviderChannels
}

// GetProviderChannelsOk returns a tuple with the ProviderChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelMappingOptionsDto) GetProviderChannelsOk() ([]NameIdPair, bool) {
	if o == nil || IsNil(o.ProviderChannels) {
		return nil, false
	}
	return o.ProviderChannels, true
}

// HasProviderChannels returns a boolean if a field has been set.
func (o *ChannelMappingOptionsDto) HasProviderChannels() bool {
	if o != nil && !IsNil(o.ProviderChannels) {
		return true
	}

	return false
}

// SetProviderChannels gets a reference to the given []NameIdPair and assigns it to the ProviderChannels field.
func (o *ChannelMappingOptionsDto) SetProviderChannels(v []NameIdPair) {
	o.ProviderChannels = v
}

// GetMappings returns the Mappings field value if set, zero value otherwise.
func (o *ChannelMappingOptionsDto) GetMappings() []NameValuePair {
	if o == nil || IsNil(o.Mappings) {
		var ret []NameValuePair
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelMappingOptionsDto) GetMappingsOk() ([]NameValuePair, bool) {
	if o == nil || IsNil(o.Mappings) {
		return nil, false
	}
	return o.Mappings, true
}

// HasMappings returns a boolean if a field has been set.
func (o *ChannelMappingOptionsDto) HasMappings() bool {
	if o != nil && !IsNil(o.Mappings) {
		return true
	}

	return false
}

// SetMappings gets a reference to the given []NameValuePair and assigns it to the Mappings field.
func (o *ChannelMappingOptionsDto) SetMappings(v []NameValuePair) {
	o.Mappings = v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelMappingOptionsDto) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderName.Get()
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelMappingOptionsDto) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderName.Get(), o.ProviderName.IsSet()
}

// HasProviderName returns a boolean if a field has been set.
func (o *ChannelMappingOptionsDto) HasProviderName() bool {
	if o != nil && o.ProviderName.IsSet() {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given NullableString and assigns it to the ProviderName field.
func (o *ChannelMappingOptionsDto) SetProviderName(v string) {
	o.ProviderName.Set(&v)
}
// SetProviderNameNil sets the value for ProviderName to be an explicit nil
func (o *ChannelMappingOptionsDto) SetProviderNameNil() {
	o.ProviderName.Set(nil)
}

// UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
func (o *ChannelMappingOptionsDto) UnsetProviderName() {
	o.ProviderName.Unset()
}

func (o ChannelMappingOptionsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelMappingOptionsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TunerChannels) {
		toSerialize["TunerChannels"] = o.TunerChannels
	}
	if !IsNil(o.ProviderChannels) {
		toSerialize["ProviderChannels"] = o.ProviderChannels
	}
	if !IsNil(o.Mappings) {
		toSerialize["Mappings"] = o.Mappings
	}
	if o.ProviderName.IsSet() {
		toSerialize["ProviderName"] = o.ProviderName.Get()
	}
	return toSerialize, nil
}

type NullableChannelMappingOptionsDto struct {
	value *ChannelMappingOptionsDto
	isSet bool
}

func (v NullableChannelMappingOptionsDto) Get() *ChannelMappingOptionsDto {
	return v.value
}

func (v *NullableChannelMappingOptionsDto) Set(val *ChannelMappingOptionsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelMappingOptionsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelMappingOptionsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelMappingOptionsDto(val *ChannelMappingOptionsDto) *NullableChannelMappingOptionsDto {
	return &NullableChannelMappingOptionsDto{value: val, isSet: true}
}

func (v NullableChannelMappingOptionsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelMappingOptionsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


