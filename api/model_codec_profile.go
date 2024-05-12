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

// checks if the CodecProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CodecProfile{}

// CodecProfile struct for CodecProfile
type CodecProfile struct {
	Type *CodecType `json:"Type,omitempty"`
	Conditions []ProfileCondition `json:"Conditions,omitempty"`
	ApplyConditions []ProfileCondition `json:"ApplyConditions,omitempty"`
	Codec NullableString `json:"Codec,omitempty"`
	Container NullableString `json:"Container,omitempty"`
}

// NewCodecProfile instantiates a new CodecProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodecProfile() *CodecProfile {
	this := CodecProfile{}
	return &this
}

// NewCodecProfileWithDefaults instantiates a new CodecProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodecProfileWithDefaults() *CodecProfile {
	this := CodecProfile{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CodecProfile) GetType() CodecType {
	if o == nil || IsNil(o.Type) {
		var ret CodecType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodecProfile) GetTypeOk() (*CodecType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CodecProfile) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CodecType and assigns it to the Type field.
func (o *CodecProfile) SetType(v CodecType) {
	o.Type = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodecProfile) GetConditions() []ProfileCondition {
	if o == nil {
		var ret []ProfileCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodecProfile) GetConditionsOk() ([]ProfileCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *CodecProfile) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ProfileCondition and assigns it to the Conditions field.
func (o *CodecProfile) SetConditions(v []ProfileCondition) {
	o.Conditions = v
}

// GetApplyConditions returns the ApplyConditions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodecProfile) GetApplyConditions() []ProfileCondition {
	if o == nil {
		var ret []ProfileCondition
		return ret
	}
	return o.ApplyConditions
}

// GetApplyConditionsOk returns a tuple with the ApplyConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodecProfile) GetApplyConditionsOk() ([]ProfileCondition, bool) {
	if o == nil || IsNil(o.ApplyConditions) {
		return nil, false
	}
	return o.ApplyConditions, true
}

// HasApplyConditions returns a boolean if a field has been set.
func (o *CodecProfile) HasApplyConditions() bool {
	if o != nil && !IsNil(o.ApplyConditions) {
		return true
	}

	return false
}

// SetApplyConditions gets a reference to the given []ProfileCondition and assigns it to the ApplyConditions field.
func (o *CodecProfile) SetApplyConditions(v []ProfileCondition) {
	o.ApplyConditions = v
}

// GetCodec returns the Codec field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodecProfile) GetCodec() string {
	if o == nil || IsNil(o.Codec.Get()) {
		var ret string
		return ret
	}
	return *o.Codec.Get()
}

// GetCodecOk returns a tuple with the Codec field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodecProfile) GetCodecOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Codec.Get(), o.Codec.IsSet()
}

// HasCodec returns a boolean if a field has been set.
func (o *CodecProfile) HasCodec() bool {
	if o != nil && o.Codec.IsSet() {
		return true
	}

	return false
}

// SetCodec gets a reference to the given NullableString and assigns it to the Codec field.
func (o *CodecProfile) SetCodec(v string) {
	o.Codec.Set(&v)
}
// SetCodecNil sets the value for Codec to be an explicit nil
func (o *CodecProfile) SetCodecNil() {
	o.Codec.Set(nil)
}

// UnsetCodec ensures that no value is present for Codec, not even an explicit nil
func (o *CodecProfile) UnsetCodec() {
	o.Codec.Unset()
}

// GetContainer returns the Container field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodecProfile) GetContainer() string {
	if o == nil || IsNil(o.Container.Get()) {
		var ret string
		return ret
	}
	return *o.Container.Get()
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodecProfile) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Container.Get(), o.Container.IsSet()
}

// HasContainer returns a boolean if a field has been set.
func (o *CodecProfile) HasContainer() bool {
	if o != nil && o.Container.IsSet() {
		return true
	}

	return false
}

// SetContainer gets a reference to the given NullableString and assigns it to the Container field.
func (o *CodecProfile) SetContainer(v string) {
	o.Container.Set(&v)
}
// SetContainerNil sets the value for Container to be an explicit nil
func (o *CodecProfile) SetContainerNil() {
	o.Container.Set(nil)
}

// UnsetContainer ensures that no value is present for Container, not even an explicit nil
func (o *CodecProfile) UnsetContainer() {
	o.Container.Unset()
}

func (o CodecProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodecProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.Conditions != nil {
		toSerialize["Conditions"] = o.Conditions
	}
	if o.ApplyConditions != nil {
		toSerialize["ApplyConditions"] = o.ApplyConditions
	}
	if o.Codec.IsSet() {
		toSerialize["Codec"] = o.Codec.Get()
	}
	if o.Container.IsSet() {
		toSerialize["Container"] = o.Container.Get()
	}
	return toSerialize, nil
}

type NullableCodecProfile struct {
	value *CodecProfile
	isSet bool
}

func (v NullableCodecProfile) Get() *CodecProfile {
	return v.value
}

func (v *NullableCodecProfile) Set(val *CodecProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableCodecProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableCodecProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodecProfile(val *CodecProfile) *NullableCodecProfile {
	return &NullableCodecProfile{value: val, isSet: true}
}

func (v NullableCodecProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodecProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


