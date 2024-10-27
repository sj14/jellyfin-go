/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ServerRestartingMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerRestartingMessage{}

// ServerRestartingMessage Server restarting down message.
type ServerRestartingMessage struct {
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewServerRestartingMessage instantiates a new ServerRestartingMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerRestartingMessage() *ServerRestartingMessage {
	this := ServerRestartingMessage{}
	return &this
}

// NewServerRestartingMessageWithDefaults instantiates a new ServerRestartingMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerRestartingMessageWithDefaults() *ServerRestartingMessage {
	this := ServerRestartingMessage{}
	return &this
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *ServerRestartingMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerRestartingMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *ServerRestartingMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *ServerRestartingMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *ServerRestartingMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerRestartingMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *ServerRestartingMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *ServerRestartingMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o ServerRestartingMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerRestartingMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageId) {
		toSerialize["MessageId"] = o.MessageId
	}
	if !IsNil(o.MessageType) {
		toSerialize["MessageType"] = o.MessageType
	}
	return toSerialize, nil
}

type NullableServerRestartingMessage struct {
	value *ServerRestartingMessage
	isSet bool
}

func (v NullableServerRestartingMessage) Get() *ServerRestartingMessage {
	return v.value
}

func (v *NullableServerRestartingMessage) Set(val *ServerRestartingMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableServerRestartingMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableServerRestartingMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerRestartingMessage(val *ServerRestartingMessage) *NullableServerRestartingMessage {
	return &NullableServerRestartingMessage{value: val, isSet: true}
}

func (v NullableServerRestartingMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerRestartingMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


