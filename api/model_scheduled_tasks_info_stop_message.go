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

// checks if the ScheduledTasksInfoStopMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduledTasksInfoStopMessage{}

// ScheduledTasksInfoStopMessage Scheduled tasks info stop message.
type ScheduledTasksInfoStopMessage struct {
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewScheduledTasksInfoStopMessage instantiates a new ScheduledTasksInfoStopMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledTasksInfoStopMessage() *ScheduledTasksInfoStopMessage {
	this := ScheduledTasksInfoStopMessage{}
	return &this
}

// NewScheduledTasksInfoStopMessageWithDefaults instantiates a new ScheduledTasksInfoStopMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledTasksInfoStopMessageWithDefaults() *ScheduledTasksInfoStopMessage {
	this := ScheduledTasksInfoStopMessage{}
	return &this
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *ScheduledTasksInfoStopMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTasksInfoStopMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *ScheduledTasksInfoStopMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *ScheduledTasksInfoStopMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o ScheduledTasksInfoStopMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduledTasksInfoStopMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageType) {
		toSerialize["MessageType"] = o.MessageType
	}
	return toSerialize, nil
}

type NullableScheduledTasksInfoStopMessage struct {
	value *ScheduledTasksInfoStopMessage
	isSet bool
}

func (v NullableScheduledTasksInfoStopMessage) Get() *ScheduledTasksInfoStopMessage {
	return v.value
}

func (v *NullableScheduledTasksInfoStopMessage) Set(val *ScheduledTasksInfoStopMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledTasksInfoStopMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledTasksInfoStopMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledTasksInfoStopMessage(val *ScheduledTasksInfoStopMessage) *NullableScheduledTasksInfoStopMessage {
	return &NullableScheduledTasksInfoStopMessage{value: val, isSet: true}
}

func (v NullableScheduledTasksInfoStopMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledTasksInfoStopMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


