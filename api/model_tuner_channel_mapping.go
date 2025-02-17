/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TunerChannelMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TunerChannelMapping{}

// TunerChannelMapping struct for TunerChannelMapping
type TunerChannelMapping struct {
	Name NullableString `json:"Name,omitempty"`
	ProviderChannelName NullableString `json:"ProviderChannelName,omitempty"`
	ProviderChannelId NullableString `json:"ProviderChannelId,omitempty"`
	Id NullableString `json:"Id,omitempty"`
}

// NewTunerChannelMapping instantiates a new TunerChannelMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunerChannelMapping() *TunerChannelMapping {
	this := TunerChannelMapping{}
	return &this
}

// NewTunerChannelMappingWithDefaults instantiates a new TunerChannelMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunerChannelMappingWithDefaults() *TunerChannelMapping {
	this := TunerChannelMapping{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunerChannelMapping) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunerChannelMapping) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TunerChannelMapping) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TunerChannelMapping) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TunerChannelMapping) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TunerChannelMapping) UnsetName() {
	o.Name.Unset()
}

// GetProviderChannelName returns the ProviderChannelName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunerChannelMapping) GetProviderChannelName() string {
	if o == nil || IsNil(o.ProviderChannelName.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderChannelName.Get()
}

// GetProviderChannelNameOk returns a tuple with the ProviderChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunerChannelMapping) GetProviderChannelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderChannelName.Get(), o.ProviderChannelName.IsSet()
}

// HasProviderChannelName returns a boolean if a field has been set.
func (o *TunerChannelMapping) HasProviderChannelName() bool {
	if o != nil && o.ProviderChannelName.IsSet() {
		return true
	}

	return false
}

// SetProviderChannelName gets a reference to the given NullableString and assigns it to the ProviderChannelName field.
func (o *TunerChannelMapping) SetProviderChannelName(v string) {
	o.ProviderChannelName.Set(&v)
}
// SetProviderChannelNameNil sets the value for ProviderChannelName to be an explicit nil
func (o *TunerChannelMapping) SetProviderChannelNameNil() {
	o.ProviderChannelName.Set(nil)
}

// UnsetProviderChannelName ensures that no value is present for ProviderChannelName, not even an explicit nil
func (o *TunerChannelMapping) UnsetProviderChannelName() {
	o.ProviderChannelName.Unset()
}

// GetProviderChannelId returns the ProviderChannelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunerChannelMapping) GetProviderChannelId() string {
	if o == nil || IsNil(o.ProviderChannelId.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderChannelId.Get()
}

// GetProviderChannelIdOk returns a tuple with the ProviderChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunerChannelMapping) GetProviderChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderChannelId.Get(), o.ProviderChannelId.IsSet()
}

// HasProviderChannelId returns a boolean if a field has been set.
func (o *TunerChannelMapping) HasProviderChannelId() bool {
	if o != nil && o.ProviderChannelId.IsSet() {
		return true
	}

	return false
}

// SetProviderChannelId gets a reference to the given NullableString and assigns it to the ProviderChannelId field.
func (o *TunerChannelMapping) SetProviderChannelId(v string) {
	o.ProviderChannelId.Set(&v)
}
// SetProviderChannelIdNil sets the value for ProviderChannelId to be an explicit nil
func (o *TunerChannelMapping) SetProviderChannelIdNil() {
	o.ProviderChannelId.Set(nil)
}

// UnsetProviderChannelId ensures that no value is present for ProviderChannelId, not even an explicit nil
func (o *TunerChannelMapping) UnsetProviderChannelId() {
	o.ProviderChannelId.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunerChannelMapping) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunerChannelMapping) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TunerChannelMapping) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *TunerChannelMapping) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *TunerChannelMapping) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TunerChannelMapping) UnsetId() {
	o.Id.Unset()
}

func (o TunerChannelMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TunerChannelMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.ProviderChannelName.IsSet() {
		toSerialize["ProviderChannelName"] = o.ProviderChannelName.Get()
	}
	if o.ProviderChannelId.IsSet() {
		toSerialize["ProviderChannelId"] = o.ProviderChannelId.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	return toSerialize, nil
}

type NullableTunerChannelMapping struct {
	value *TunerChannelMapping
	isSet bool
}

func (v NullableTunerChannelMapping) Get() *TunerChannelMapping {
	return v.value
}

func (v *NullableTunerChannelMapping) Set(val *TunerChannelMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableTunerChannelMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableTunerChannelMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunerChannelMapping(val *TunerChannelMapping) *NullableTunerChannelMapping {
	return &NullableTunerChannelMapping{value: val, isSet: true}
}

func (v NullableTunerChannelMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunerChannelMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


