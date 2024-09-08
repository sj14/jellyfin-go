/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SessionsStopMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionsStopMessage{}

// SessionsStopMessage Sessions stop message.
type SessionsStopMessage struct {
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewSessionsStopMessage instantiates a new SessionsStopMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionsStopMessage() *SessionsStopMessage {
	this := SessionsStopMessage{}
	return &this
}

// NewSessionsStopMessageWithDefaults instantiates a new SessionsStopMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionsStopMessageWithDefaults() *SessionsStopMessage {
	this := SessionsStopMessage{}
	return &this
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *SessionsStopMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsStopMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *SessionsStopMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *SessionsStopMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o SessionsStopMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionsStopMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageType) {
		toSerialize["MessageType"] = o.MessageType
	}
	return toSerialize, nil
}

type NullableSessionsStopMessage struct {
	value *SessionsStopMessage
	isSet bool
}

func (v NullableSessionsStopMessage) Get() *SessionsStopMessage {
	return v.value
}

func (v *NullableSessionsStopMessage) Set(val *SessionsStopMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionsStopMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionsStopMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionsStopMessage(val *SessionsStopMessage) *NullableSessionsStopMessage {
	return &NullableSessionsStopMessage{value: val, isSet: true}
}

func (v NullableSessionsStopMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionsStopMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


