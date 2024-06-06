/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the UserUpdatedMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUpdatedMessage{}

// UserUpdatedMessage User updated message.
type UserUpdatedMessage struct {
	Data NullableUserUpdatedMessageData `json:"Data,omitempty"`
	// Gets or sets the message id.
	MessageId *string `json:"MessageId,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewUserUpdatedMessage instantiates a new UserUpdatedMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdatedMessage() *UserUpdatedMessage {
	this := UserUpdatedMessage{}
	return &this
}

// NewUserUpdatedMessageWithDefaults instantiates a new UserUpdatedMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdatedMessageWithDefaults() *UserUpdatedMessage {
	this := UserUpdatedMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdatedMessage) GetData() UserUpdatedMessageData {
	if o == nil || IsNil(o.Data.Get()) {
		var ret UserUpdatedMessageData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdatedMessage) GetDataOk() (*UserUpdatedMessageData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *UserUpdatedMessage) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableUserUpdatedMessageData and assigns it to the Data field.
func (o *UserUpdatedMessage) SetData(v UserUpdatedMessageData) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *UserUpdatedMessage) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *UserUpdatedMessage) UnsetData() {
	o.Data.Unset()
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *UserUpdatedMessage) GetMessageId() string {
	if o == nil || IsNil(o.MessageId) {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdatedMessage) GetMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessageId) {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *UserUpdatedMessage) HasMessageId() bool {
	if o != nil && !IsNil(o.MessageId) {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *UserUpdatedMessage) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *UserUpdatedMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdatedMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *UserUpdatedMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *UserUpdatedMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o UserUpdatedMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUpdatedMessage) ToMap() (map[string]interface{}, error) {
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

type NullableUserUpdatedMessage struct {
	value *UserUpdatedMessage
	isSet bool
}

func (v NullableUserUpdatedMessage) Get() *UserUpdatedMessage {
	return v.value
}

func (v *NullableUserUpdatedMessage) Set(val *UserUpdatedMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdatedMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdatedMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdatedMessage(val *UserUpdatedMessage) *NullableUserUpdatedMessage {
	return &NullableUserUpdatedMessage{value: val, isSet: true}
}

func (v NullableUserUpdatedMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdatedMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


