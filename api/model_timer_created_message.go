/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TimerCreatedMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimerCreatedMessage{}

// TimerCreatedMessage Timer created message.
type TimerCreatedMessage struct {
	Data NullableSeriesTimerCancelledMessageData `json:"Data,omitempty"`
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewTimerCreatedMessage instantiates a new TimerCreatedMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimerCreatedMessage() *TimerCreatedMessage {
	this := TimerCreatedMessage{}
	return &this
}

// NewTimerCreatedMessageWithDefaults instantiates a new TimerCreatedMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimerCreatedMessageWithDefaults() *TimerCreatedMessage {
	this := TimerCreatedMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimerCreatedMessage) GetData() SeriesTimerCancelledMessageData {
	if o == nil || IsNil(o.Data.Get()) {
		var ret SeriesTimerCancelledMessageData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimerCreatedMessage) GetDataOk() (*SeriesTimerCancelledMessageData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *TimerCreatedMessage) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableSeriesTimerCancelledMessageData and assigns it to the Data field.
func (o *TimerCreatedMessage) SetData(v SeriesTimerCancelledMessageData) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *TimerCreatedMessage) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *TimerCreatedMessage) UnsetData() {
	o.Data.Unset()
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *TimerCreatedMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerCreatedMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *TimerCreatedMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *TimerCreatedMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *TimerCreatedMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerCreatedMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *TimerCreatedMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *TimerCreatedMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o TimerCreatedMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimerCreatedMessage) ToMap() (map[string]interface{}, error) {
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

type NullableTimerCreatedMessage struct {
	value *TimerCreatedMessage
	isSet bool
}

func (v NullableTimerCreatedMessage) Get() *TimerCreatedMessage {
	return v.value
}

func (v *NullableTimerCreatedMessage) Set(val *TimerCreatedMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableTimerCreatedMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableTimerCreatedMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimerCreatedMessage(val *TimerCreatedMessage) *NullableTimerCreatedMessage {
	return &NullableTimerCreatedMessage{value: val, isSet: true}
}

func (v NullableTimerCreatedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimerCreatedMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


