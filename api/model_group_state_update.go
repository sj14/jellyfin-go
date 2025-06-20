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

// checks if the GroupStateUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupStateUpdate{}

// GroupStateUpdate Class GroupStateUpdate.
type GroupStateUpdate struct {
	// Gets the state of the group.
	State *GroupStateType `json:"State,omitempty"`
	// Gets the reason of the state change.
	Reason *PlaybackRequestType `json:"Reason,omitempty"`
}

// NewGroupStateUpdate instantiates a new GroupStateUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupStateUpdate() *GroupStateUpdate {
	this := GroupStateUpdate{}
	return &this
}

// NewGroupStateUpdateWithDefaults instantiates a new GroupStateUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupStateUpdateWithDefaults() *GroupStateUpdate {
	this := GroupStateUpdate{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GroupStateUpdate) GetState() GroupStateType {
	if o == nil || IsNil(o.State) {
		var ret GroupStateType
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStateUpdate) GetStateOk() (*GroupStateType, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GroupStateUpdate) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given GroupStateType and assigns it to the State field.
func (o *GroupStateUpdate) SetState(v GroupStateType) {
	o.State = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *GroupStateUpdate) GetReason() PlaybackRequestType {
	if o == nil || IsNil(o.Reason) {
		var ret PlaybackRequestType
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupStateUpdate) GetReasonOk() (*PlaybackRequestType, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *GroupStateUpdate) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given PlaybackRequestType and assigns it to the Reason field.
func (o *GroupStateUpdate) SetReason(v PlaybackRequestType) {
	o.Reason = &v
}

func (o GroupStateUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupStateUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if !IsNil(o.Reason) {
		toSerialize["Reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableGroupStateUpdate struct {
	value *GroupStateUpdate
	isSet bool
}

func (v NullableGroupStateUpdate) Get() *GroupStateUpdate {
	return v.value
}

func (v *NullableGroupStateUpdate) Set(val *GroupStateUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupStateUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupStateUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupStateUpdate(val *GroupStateUpdate) *NullableGroupStateUpdate {
	return &NullableGroupStateUpdate{value: val, isSet: true}
}

func (v NullableGroupStateUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupStateUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


