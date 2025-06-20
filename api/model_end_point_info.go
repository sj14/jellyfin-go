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

// checks if the EndPointInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndPointInfo{}

// EndPointInfo struct for EndPointInfo
type EndPointInfo struct {
	IsLocal *bool `json:"IsLocal,omitempty"`
	IsInNetwork *bool `json:"IsInNetwork,omitempty"`
}

// NewEndPointInfo instantiates a new EndPointInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndPointInfo() *EndPointInfo {
	this := EndPointInfo{}
	return &this
}

// NewEndPointInfoWithDefaults instantiates a new EndPointInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndPointInfoWithDefaults() *EndPointInfo {
	this := EndPointInfo{}
	return &this
}

// GetIsLocal returns the IsLocal field value if set, zero value otherwise.
func (o *EndPointInfo) GetIsLocal() bool {
	if o == nil || IsNil(o.IsLocal) {
		var ret bool
		return ret
	}
	return *o.IsLocal
}

// GetIsLocalOk returns a tuple with the IsLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndPointInfo) GetIsLocalOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLocal) {
		return nil, false
	}
	return o.IsLocal, true
}

// HasIsLocal returns a boolean if a field has been set.
func (o *EndPointInfo) HasIsLocal() bool {
	if o != nil && !IsNil(o.IsLocal) {
		return true
	}

	return false
}

// SetIsLocal gets a reference to the given bool and assigns it to the IsLocal field.
func (o *EndPointInfo) SetIsLocal(v bool) {
	o.IsLocal = &v
}

// GetIsInNetwork returns the IsInNetwork field value if set, zero value otherwise.
func (o *EndPointInfo) GetIsInNetwork() bool {
	if o == nil || IsNil(o.IsInNetwork) {
		var ret bool
		return ret
	}
	return *o.IsInNetwork
}

// GetIsInNetworkOk returns a tuple with the IsInNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndPointInfo) GetIsInNetworkOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInNetwork) {
		return nil, false
	}
	return o.IsInNetwork, true
}

// HasIsInNetwork returns a boolean if a field has been set.
func (o *EndPointInfo) HasIsInNetwork() bool {
	if o != nil && !IsNil(o.IsInNetwork) {
		return true
	}

	return false
}

// SetIsInNetwork gets a reference to the given bool and assigns it to the IsInNetwork field.
func (o *EndPointInfo) SetIsInNetwork(v bool) {
	o.IsInNetwork = &v
}

func (o EndPointInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndPointInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsLocal) {
		toSerialize["IsLocal"] = o.IsLocal
	}
	if !IsNil(o.IsInNetwork) {
		toSerialize["IsInNetwork"] = o.IsInNetwork
	}
	return toSerialize, nil
}

type NullableEndPointInfo struct {
	value *EndPointInfo
	isSet bool
}

func (v NullableEndPointInfo) Get() *EndPointInfo {
	return v.value
}

func (v *NullableEndPointInfo) Set(val *EndPointInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEndPointInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEndPointInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndPointInfo(val *EndPointInfo) *NullableEndPointInfo {
	return &NullableEndPointInfo{value: val, isSet: true}
}

func (v NullableEndPointInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndPointInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


