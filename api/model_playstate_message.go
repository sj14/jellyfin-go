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

// checks if the PlaystateMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaystateMessage{}

// PlaystateMessage Playstate message.
type PlaystateMessage struct {
	// Gets or sets the data.
	Data NullablePlaystateRequest `json:"Data,omitempty"`
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewPlaystateMessage instantiates a new PlaystateMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaystateMessage() *PlaystateMessage {
	this := PlaystateMessage{}
	return &this
}

// NewPlaystateMessageWithDefaults instantiates a new PlaystateMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaystateMessageWithDefaults() *PlaystateMessage {
	this := PlaystateMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlaystateMessage) GetData() PlaystateRequest {
	if o == nil || IsNil(o.Data.Get()) {
		var ret PlaystateRequest
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlaystateMessage) GetDataOk() (*PlaystateRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *PlaystateMessage) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullablePlaystateRequest and assigns it to the Data field.
func (o *PlaystateMessage) SetData(v PlaystateRequest) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *PlaystateMessage) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *PlaystateMessage) UnsetData() {
	o.Data.Unset()
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *PlaystateMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaystateMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *PlaystateMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *PlaystateMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *PlaystateMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaystateMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *PlaystateMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *PlaystateMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o PlaystateMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaystateMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["Data"] = o.Data.Get()
	}
	if !IsNil(o.MessageId) {
		toSerialize["MessageId"] = o.MessageId
	}
	if !IsNil(o.MessageType) {
		toSerialize["MessageType"] = o.MessageType
	}
	return toSerialize, nil
}

type NullablePlaystateMessage struct {
	value *PlaystateMessage
	isSet bool
}

func (v NullablePlaystateMessage) Get() *PlaystateMessage {
	return v.value
}

func (v *NullablePlaystateMessage) Set(val *PlaystateMessage) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaystateMessage) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaystateMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaystateMessage(val *PlaystateMessage) *NullablePlaystateMessage {
	return &NullablePlaystateMessage{value: val, isSet: true}
}

func (v NullablePlaystateMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaystateMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


