/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the TimerCancelledMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimerCancelledMessage{}

// TimerCancelledMessage Timer cancelled message.
type TimerCancelledMessage struct {
	// Gets or sets the data.
	Data NullableTimerEventInfo `json:"Data,omitempty"`
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewTimerCancelledMessage instantiates a new TimerCancelledMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimerCancelledMessage() *TimerCancelledMessage {
	this := TimerCancelledMessage{}
	return &this
}

// NewTimerCancelledMessageWithDefaults instantiates a new TimerCancelledMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimerCancelledMessageWithDefaults() *TimerCancelledMessage {
	this := TimerCancelledMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TimerCancelledMessage) GetData() TimerEventInfo {
	if o == nil || IsNil(o.Data.Get()) {
		var ret TimerEventInfo
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TimerCancelledMessage) GetDataOk() (*TimerEventInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *TimerCancelledMessage) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableTimerEventInfo and assigns it to the Data field.
func (o *TimerCancelledMessage) SetData(v TimerEventInfo) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *TimerCancelledMessage) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *TimerCancelledMessage) UnsetData() {
	o.Data.Unset()
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *TimerCancelledMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerCancelledMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *TimerCancelledMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *TimerCancelledMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *TimerCancelledMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimerCancelledMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *TimerCancelledMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *TimerCancelledMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o TimerCancelledMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimerCancelledMessage) ToMap() (map[string]interface{}, error) {
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

type NullableTimerCancelledMessage struct {
	value *TimerCancelledMessage
	isSet bool
}

func (v NullableTimerCancelledMessage) Get() *TimerCancelledMessage {
	return v.value
}

func (v *NullableTimerCancelledMessage) Set(val *TimerCancelledMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableTimerCancelledMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableTimerCancelledMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimerCancelledMessage(val *TimerCancelledMessage) *NullableTimerCancelledMessage {
	return &NullableTimerCancelledMessage{value: val, isSet: true}
}

func (v NullableTimerCancelledMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimerCancelledMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


