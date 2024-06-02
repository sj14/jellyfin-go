/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PlaystateMessageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaystateMessageData{}

// PlaystateMessageData Gets or sets the data.
type PlaystateMessageData struct {
	// Enum PlaystateCommand.
	Command *PlaystateCommand `json:"Command,omitempty"`
	SeekPositionTicks NullableInt64 `json:"SeekPositionTicks,omitempty"`
	// Gets or sets the controlling user identifier.
	ControllingUserId NullableString `json:"ControllingUserId,omitempty"`
}

// NewPlaystateMessageData instantiates a new PlaystateMessageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaystateMessageData() *PlaystateMessageData {
	this := PlaystateMessageData{}
	return &this
}

// NewPlaystateMessageDataWithDefaults instantiates a new PlaystateMessageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaystateMessageDataWithDefaults() *PlaystateMessageData {
	this := PlaystateMessageData{}
	return &this
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *PlaystateMessageData) GetCommand() PlaystateCommand {
	if o == nil || IsNil(o.Command) {
		var ret PlaystateCommand
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaystateMessageData) GetCommandOk() (*PlaystateCommand, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *PlaystateMessageData) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given PlaystateCommand and assigns it to the Command field.
func (o *PlaystateMessageData) SetCommand(v PlaystateCommand) {
	o.Command = &v
}

// GetSeekPositionTicks returns the SeekPositionTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaystateMessageData) GetSeekPositionTicks() int64 {
	if o == nil || IsNil(o.SeekPositionTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.SeekPositionTicks.Get()
}

// GetSeekPositionTicksOk returns a tuple with the SeekPositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaystateMessageData) GetSeekPositionTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeekPositionTicks.Get(), o.SeekPositionTicks.IsSet()
}

// HasSeekPositionTicks returns a boolean if a field has been set.
func (o *PlaystateMessageData) HasSeekPositionTicks() bool {
	if o != nil && o.SeekPositionTicks.IsSet() {
		return true
	}

	return false
}

// SetSeekPositionTicks gets a reference to the given NullableInt64 and assigns it to the SeekPositionTicks field.
func (o *PlaystateMessageData) SetSeekPositionTicks(v int64) {
	o.SeekPositionTicks.Set(&v)
}
// SetSeekPositionTicksNil sets the value for SeekPositionTicks to be an explicit nil
func (o *PlaystateMessageData) SetSeekPositionTicksNil() {
	o.SeekPositionTicks.Set(nil)
}

// UnsetSeekPositionTicks ensures that no value is present for SeekPositionTicks, not even an explicit nil
func (o *PlaystateMessageData) UnsetSeekPositionTicks() {
	o.SeekPositionTicks.Unset()
}

// GetControllingUserId returns the ControllingUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaystateMessageData) GetControllingUserId() string {
	if o == nil || IsNil(o.ControllingUserId.Get()) {
		var ret string
		return ret
	}
	return *o.ControllingUserId.Get()
}

// GetControllingUserIdOk returns a tuple with the ControllingUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaystateMessageData) GetControllingUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ControllingUserId.Get(), o.ControllingUserId.IsSet()
}

// HasControllingUserId returns a boolean if a field has been set.
func (o *PlaystateMessageData) HasControllingUserId() bool {
	if o != nil && o.ControllingUserId.IsSet() {
		return true
	}

	return false
}

// SetControllingUserId gets a reference to the given NullableString and assigns it to the ControllingUserId field.
func (o *PlaystateMessageData) SetControllingUserId(v string) {
	o.ControllingUserId.Set(&v)
}
// SetControllingUserIdNil sets the value for ControllingUserId to be an explicit nil
func (o *PlaystateMessageData) SetControllingUserIdNil() {
	o.ControllingUserId.Set(nil)
}

// UnsetControllingUserId ensures that no value is present for ControllingUserId, not even an explicit nil
func (o *PlaystateMessageData) UnsetControllingUserId() {
	o.ControllingUserId.Unset()
}

func (o PlaystateMessageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaystateMessageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Command) {
		toSerialize["Command"] = o.Command
	}
	if o.SeekPositionTicks.IsSet() {
		toSerialize["SeekPositionTicks"] = o.SeekPositionTicks.Get()
	}
	if o.ControllingUserId.IsSet() {
		toSerialize["ControllingUserId"] = o.ControllingUserId.Get()
	}
	return toSerialize, nil
}

type NullablePlaystateMessageData struct {
	value *PlaystateMessageData
	isSet bool
}

func (v NullablePlaystateMessageData) Get() *PlaystateMessageData {
	return v.value
}

func (v *NullablePlaystateMessageData) Set(val *PlaystateMessageData) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaystateMessageData) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaystateMessageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaystateMessageData(val *PlaystateMessageData) *NullablePlaystateMessageData {
	return &NullablePlaystateMessageData{value: val, isSet: true}
}

func (v NullablePlaystateMessageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaystateMessageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


