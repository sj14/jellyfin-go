/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PathSubstitution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PathSubstitution{}

// PathSubstitution Defines the MediaBrowser.Model.Configuration.PathSubstitution.
type PathSubstitution struct {
	// Gets or sets the value to substitute.
	From *string `json:"From,omitempty"`
	// Gets or sets the value to substitution with.
	To *string `json:"To,omitempty"`
}

// NewPathSubstitution instantiates a new PathSubstitution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathSubstitution() *PathSubstitution {
	this := PathSubstitution{}
	return &this
}

// NewPathSubstitutionWithDefaults instantiates a new PathSubstitution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathSubstitutionWithDefaults() *PathSubstitution {
	this := PathSubstitution{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *PathSubstitution) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathSubstitution) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *PathSubstitution) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *PathSubstitution) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *PathSubstitution) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathSubstitution) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *PathSubstitution) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *PathSubstitution) SetTo(v string) {
	o.To = &v
}

func (o PathSubstitution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PathSubstitution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["From"] = o.From
	}
	if !IsNil(o.To) {
		toSerialize["To"] = o.To
	}
	return toSerialize, nil
}

type NullablePathSubstitution struct {
	value *PathSubstitution
	isSet bool
}

func (v NullablePathSubstitution) Get() *PathSubstitution {
	return v.value
}

func (v *NullablePathSubstitution) Set(val *PathSubstitution) {
	v.value = val
	v.isSet = true
}

func (v NullablePathSubstitution) IsSet() bool {
	return v.isSet
}

func (v *NullablePathSubstitution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathSubstitution(val *PathSubstitution) *NullablePathSubstitution {
	return &NullablePathSubstitution{value: val, isSet: true}
}

func (v NullablePathSubstitution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathSubstitution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


