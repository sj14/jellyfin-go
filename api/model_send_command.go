/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the SendCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendCommand{}

// SendCommand Class SendCommand.
type SendCommand struct {
	// Gets the group identifier.
	GroupId *string `json:"GroupId,omitempty"`
	// Gets the playlist identifier of the playing item.
	PlaylistItemId *string `json:"PlaylistItemId,omitempty"`
	// Gets or sets the UTC time when to execute the command.
	When *time.Time `json:"When,omitempty"`
	// Gets the position ticks.
	PositionTicks NullableInt64 `json:"PositionTicks,omitempty"`
	// Gets the command.
	Command *SendCommandType `json:"Command,omitempty"`
	// Gets the UTC time when this command has been emitted.
	EmittedAt *time.Time `json:"EmittedAt,omitempty"`
}

// NewSendCommand instantiates a new SendCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendCommand() *SendCommand {
	this := SendCommand{}
	return &this
}

// NewSendCommandWithDefaults instantiates a new SendCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendCommandWithDefaults() *SendCommand {
	this := SendCommand{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *SendCommand) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendCommand) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *SendCommand) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *SendCommand) SetGroupId(v string) {
	o.GroupId = &v
}

// GetPlaylistItemId returns the PlaylistItemId field value if set, zero value otherwise.
func (o *SendCommand) GetPlaylistItemId() string {
	if o == nil || IsNil(o.PlaylistItemId) {
		var ret string
		return ret
	}
	return *o.PlaylistItemId
}

// GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendCommand) GetPlaylistItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaylistItemId) {
		return nil, false
	}
	return o.PlaylistItemId, true
}

// HasPlaylistItemId returns a boolean if a field has been set.
func (o *SendCommand) HasPlaylistItemId() bool {
	if o != nil && !IsNil(o.PlaylistItemId) {
		return true
	}

	return false
}

// SetPlaylistItemId gets a reference to the given string and assigns it to the PlaylistItemId field.
func (o *SendCommand) SetPlaylistItemId(v string) {
	o.PlaylistItemId = &v
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *SendCommand) GetWhen() time.Time {
	if o == nil || IsNil(o.When) {
		var ret time.Time
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendCommand) GetWhenOk() (*time.Time, bool) {
	if o == nil || IsNil(o.When) {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *SendCommand) HasWhen() bool {
	if o != nil && !IsNil(o.When) {
		return true
	}

	return false
}

// SetWhen gets a reference to the given time.Time and assigns it to the When field.
func (o *SendCommand) SetWhen(v time.Time) {
	o.When = &v
}

// GetPositionTicks returns the PositionTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SendCommand) GetPositionTicks() int64 {
	if o == nil || IsNil(o.PositionTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.PositionTicks.Get()
}

// GetPositionTicksOk returns a tuple with the PositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SendCommand) GetPositionTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PositionTicks.Get(), o.PositionTicks.IsSet()
}

// HasPositionTicks returns a boolean if a field has been set.
func (o *SendCommand) HasPositionTicks() bool {
	if o != nil && o.PositionTicks.IsSet() {
		return true
	}

	return false
}

// SetPositionTicks gets a reference to the given NullableInt64 and assigns it to the PositionTicks field.
func (o *SendCommand) SetPositionTicks(v int64) {
	o.PositionTicks.Set(&v)
}
// SetPositionTicksNil sets the value for PositionTicks to be an explicit nil
func (o *SendCommand) SetPositionTicksNil() {
	o.PositionTicks.Set(nil)
}

// UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
func (o *SendCommand) UnsetPositionTicks() {
	o.PositionTicks.Unset()
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *SendCommand) GetCommand() SendCommandType {
	if o == nil || IsNil(o.Command) {
		var ret SendCommandType
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendCommand) GetCommandOk() (*SendCommandType, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *SendCommand) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given SendCommandType and assigns it to the Command field.
func (o *SendCommand) SetCommand(v SendCommandType) {
	o.Command = &v
}

// GetEmittedAt returns the EmittedAt field value if set, zero value otherwise.
func (o *SendCommand) GetEmittedAt() time.Time {
	if o == nil || IsNil(o.EmittedAt) {
		var ret time.Time
		return ret
	}
	return *o.EmittedAt
}

// GetEmittedAtOk returns a tuple with the EmittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendCommand) GetEmittedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EmittedAt) {
		return nil, false
	}
	return o.EmittedAt, true
}

// HasEmittedAt returns a boolean if a field has been set.
func (o *SendCommand) HasEmittedAt() bool {
	if o != nil && !IsNil(o.EmittedAt) {
		return true
	}

	return false
}

// SetEmittedAt gets a reference to the given time.Time and assigns it to the EmittedAt field.
func (o *SendCommand) SetEmittedAt(v time.Time) {
	o.EmittedAt = &v
}

func (o SendCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["GroupId"] = o.GroupId
	}
	if !IsNil(o.PlaylistItemId) {
		toSerialize["PlaylistItemId"] = o.PlaylistItemId
	}
	if !IsNil(o.When) {
		toSerialize["When"] = o.When
	}
	if o.PositionTicks.IsSet() {
		toSerialize["PositionTicks"] = o.PositionTicks.Get()
	}
	if !IsNil(o.Command) {
		toSerialize["Command"] = o.Command
	}
	if !IsNil(o.EmittedAt) {
		toSerialize["EmittedAt"] = o.EmittedAt
	}
	return toSerialize, nil
}

type NullableSendCommand struct {
	value *SendCommand
	isSet bool
}

func (v NullableSendCommand) Get() *SendCommand {
	return v.value
}

func (v *NullableSendCommand) Set(val *SendCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableSendCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableSendCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendCommand(val *SendCommand) *NullableSendCommand {
	return &NullableSendCommand{value: val, isSet: true}
}

func (v NullableSendCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


