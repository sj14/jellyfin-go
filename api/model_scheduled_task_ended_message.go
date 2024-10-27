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

// checks if the ScheduledTaskEndedMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduledTaskEndedMessage{}

// ScheduledTaskEndedMessage Scheduled task ended message.
type ScheduledTaskEndedMessage struct {
	// Class TaskExecutionInfo.
	Data NullableTaskResult `json:"Data,omitempty"`
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewScheduledTaskEndedMessage instantiates a new ScheduledTaskEndedMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledTaskEndedMessage() *ScheduledTaskEndedMessage {
	this := ScheduledTaskEndedMessage{}
	return &this
}

// NewScheduledTaskEndedMessageWithDefaults instantiates a new ScheduledTaskEndedMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledTaskEndedMessageWithDefaults() *ScheduledTaskEndedMessage {
	this := ScheduledTaskEndedMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ScheduledTaskEndedMessage) GetData() TaskResult {
	if o == nil || IsNil(o.Data.Get()) {
		var ret TaskResult
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ScheduledTaskEndedMessage) GetDataOk() (*TaskResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *ScheduledTaskEndedMessage) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableTaskResult and assigns it to the Data field.
func (o *ScheduledTaskEndedMessage) SetData(v TaskResult) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *ScheduledTaskEndedMessage) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *ScheduledTaskEndedMessage) UnsetData() {
	o.Data.Unset()
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *ScheduledTaskEndedMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskEndedMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *ScheduledTaskEndedMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *ScheduledTaskEndedMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *ScheduledTaskEndedMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskEndedMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *ScheduledTaskEndedMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *ScheduledTaskEndedMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o ScheduledTaskEndedMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduledTaskEndedMessage) ToMap() (map[string]interface{}, error) {
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

type NullableScheduledTaskEndedMessage struct {
	value *ScheduledTaskEndedMessage
	isSet bool
}

func (v NullableScheduledTaskEndedMessage) Get() *ScheduledTaskEndedMessage {
	return v.value
}

func (v *NullableScheduledTaskEndedMessage) Set(val *ScheduledTaskEndedMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledTaskEndedMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledTaskEndedMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledTaskEndedMessage(val *ScheduledTaskEndedMessage) *NullableScheduledTaskEndedMessage {
	return &NullableScheduledTaskEndedMessage{value: val, isSet: true}
}

func (v NullableScheduledTaskEndedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledTaskEndedMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


