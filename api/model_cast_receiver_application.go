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

// checks if the CastReceiverApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastReceiverApplication{}

// CastReceiverApplication The cast receiver application model.
type CastReceiverApplication struct {
	// Gets or sets the cast receiver application id.
	Id *string `json:"Id,omitempty"`
	// Gets or sets the cast receiver application name.
	Name *string `json:"Name,omitempty"`
}

// NewCastReceiverApplication instantiates a new CastReceiverApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastReceiverApplication() *CastReceiverApplication {
	this := CastReceiverApplication{}
	return &this
}

// NewCastReceiverApplicationWithDefaults instantiates a new CastReceiverApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastReceiverApplicationWithDefaults() *CastReceiverApplication {
	this := CastReceiverApplication{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CastReceiverApplication) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastReceiverApplication) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CastReceiverApplication) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CastReceiverApplication) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CastReceiverApplication) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastReceiverApplication) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CastReceiverApplication) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CastReceiverApplication) SetName(v string) {
	o.Name = &v
}

func (o CastReceiverApplication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastReceiverApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCastReceiverApplication struct {
	value *CastReceiverApplication
	isSet bool
}

func (v NullableCastReceiverApplication) Get() *CastReceiverApplication {
	return v.value
}

func (v *NullableCastReceiverApplication) Set(val *CastReceiverApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableCastReceiverApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableCastReceiverApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastReceiverApplication(val *CastReceiverApplication) *NullableCastReceiverApplication {
	return &NullableCastReceiverApplication{value: val, isSet: true}
}

func (v NullableCastReceiverApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastReceiverApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


