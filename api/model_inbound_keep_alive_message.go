/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the InboundKeepAliveMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundKeepAliveMessage{}

// InboundKeepAliveMessage Keep alive websocket messages.
type InboundKeepAliveMessage struct {
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewInboundKeepAliveMessage instantiates a new InboundKeepAliveMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundKeepAliveMessage() *InboundKeepAliveMessage {
	this := InboundKeepAliveMessage{}
	return &this
}

// NewInboundKeepAliveMessageWithDefaults instantiates a new InboundKeepAliveMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundKeepAliveMessageWithDefaults() *InboundKeepAliveMessage {
	this := InboundKeepAliveMessage{}
	return &this
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *InboundKeepAliveMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InboundKeepAliveMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *InboundKeepAliveMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *InboundKeepAliveMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o InboundKeepAliveMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InboundKeepAliveMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageType) {
		toSerialize["MessageType"] = o.MessageType
	}
	return toSerialize, nil
}

type NullableInboundKeepAliveMessage struct {
	value *InboundKeepAliveMessage
	isSet bool
}

func (v NullableInboundKeepAliveMessage) Get() *InboundKeepAliveMessage {
	return v.value
}

func (v *NullableInboundKeepAliveMessage) Set(val *InboundKeepAliveMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundKeepAliveMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundKeepAliveMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundKeepAliveMessage(val *InboundKeepAliveMessage) *NullableInboundKeepAliveMessage {
	return &NullableInboundKeepAliveMessage{value: val, isSet: true}
}

func (v NullableInboundKeepAliveMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundKeepAliveMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


