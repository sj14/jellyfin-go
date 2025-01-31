/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ContainerProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerProfile{}

// ContainerProfile Defines the MediaBrowser.Model.Dlna.ContainerProfile.
type ContainerProfile struct {
	// Gets or sets the MediaBrowser.Model.Dlna.DlnaProfileType which this container must meet.
	Type *DlnaProfileType `json:"Type,omitempty"`
	// Gets or sets the list of MediaBrowser.Model.Dlna.ProfileCondition which this container will be applied to.
	Conditions []ProfileCondition `json:"Conditions,omitempty"`
	// Gets or sets the container(s) which this container must meet.
	Container NullableString `json:"Container,omitempty"`
	// Gets or sets the sub container(s) which this container must meet.
	SubContainer NullableString `json:"SubContainer,omitempty"`
}

// NewContainerProfile instantiates a new ContainerProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerProfile() *ContainerProfile {
	this := ContainerProfile{}
	return &this
}

// NewContainerProfileWithDefaults instantiates a new ContainerProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerProfileWithDefaults() *ContainerProfile {
	this := ContainerProfile{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContainerProfile) GetType() DlnaProfileType {
	if o == nil || IsNil(o.Type) {
		var ret DlnaProfileType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerProfile) GetTypeOk() (*DlnaProfileType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContainerProfile) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DlnaProfileType and assigns it to the Type field.
func (o *ContainerProfile) SetType(v DlnaProfileType) {
	o.Type = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ContainerProfile) GetConditions() []ProfileCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []ProfileCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerProfile) GetConditionsOk() ([]ProfileCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ContainerProfile) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ProfileCondition and assigns it to the Conditions field.
func (o *ContainerProfile) SetConditions(v []ProfileCondition) {
	o.Conditions = v
}

// GetContainer returns the Container field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerProfile) GetContainer() string {
	if o == nil || IsNil(o.Container.Get()) {
		var ret string
		return ret
	}
	return *o.Container.Get()
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerProfile) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Container.Get(), o.Container.IsSet()
}

// HasContainer returns a boolean if a field has been set.
func (o *ContainerProfile) HasContainer() bool {
	if o != nil && o.Container.IsSet() {
		return true
	}

	return false
}

// SetContainer gets a reference to the given NullableString and assigns it to the Container field.
func (o *ContainerProfile) SetContainer(v string) {
	o.Container.Set(&v)
}
// SetContainerNil sets the value for Container to be an explicit nil
func (o *ContainerProfile) SetContainerNil() {
	o.Container.Set(nil)
}

// UnsetContainer ensures that no value is present for Container, not even an explicit nil
func (o *ContainerProfile) UnsetContainer() {
	o.Container.Unset()
}

// GetSubContainer returns the SubContainer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerProfile) GetSubContainer() string {
	if o == nil || IsNil(o.SubContainer.Get()) {
		var ret string
		return ret
	}
	return *o.SubContainer.Get()
}

// GetSubContainerOk returns a tuple with the SubContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerProfile) GetSubContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubContainer.Get(), o.SubContainer.IsSet()
}

// HasSubContainer returns a boolean if a field has been set.
func (o *ContainerProfile) HasSubContainer() bool {
	if o != nil && o.SubContainer.IsSet() {
		return true
	}

	return false
}

// SetSubContainer gets a reference to the given NullableString and assigns it to the SubContainer field.
func (o *ContainerProfile) SetSubContainer(v string) {
	o.SubContainer.Set(&v)
}
// SetSubContainerNil sets the value for SubContainer to be an explicit nil
func (o *ContainerProfile) SetSubContainerNil() {
	o.SubContainer.Set(nil)
}

// UnsetSubContainer ensures that no value is present for SubContainer, not even an explicit nil
func (o *ContainerProfile) UnsetSubContainer() {
	o.SubContainer.Unset()
}

func (o ContainerProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.Conditions) {
		toSerialize["Conditions"] = o.Conditions
	}
	if o.Container.IsSet() {
		toSerialize["Container"] = o.Container.Get()
	}
	if o.SubContainer.IsSet() {
		toSerialize["SubContainer"] = o.SubContainer.Get()
	}
	return toSerialize, nil
}

type NullableContainerProfile struct {
	value *ContainerProfile
	isSet bool
}

func (v NullableContainerProfile) Get() *ContainerProfile {
	return v.value
}

func (v *NullableContainerProfile) Set(val *ContainerProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerProfile(val *ContainerProfile) *NullableContainerProfile {
	return &NullableContainerProfile{value: val, isSet: true}
}

func (v NullableContainerProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


