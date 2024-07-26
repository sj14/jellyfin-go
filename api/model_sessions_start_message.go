/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SessionsStartMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionsStartMessage{}

// SessionsStartMessage Sessions start message.  Data is the timing data encoded as \"$initialDelay,$interval\" in ms.
type SessionsStartMessage struct {
	// Gets or sets the data.
	Data NullableString `json:"Data,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewSessionsStartMessage instantiates a new SessionsStartMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionsStartMessage() *SessionsStartMessage {
	this := SessionsStartMessage{}
	return &this
}

// NewSessionsStartMessageWithDefaults instantiates a new SessionsStartMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionsStartMessageWithDefaults() *SessionsStartMessage {
	this := SessionsStartMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionsStartMessage) GetData() string {
	if o == nil || IsNil(o.Data.Get()) {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionsStartMessage) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *SessionsStartMessage) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *SessionsStartMessage) SetData(v string) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *SessionsStartMessage) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *SessionsStartMessage) UnsetData() {
	o.Data.Unset()
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *SessionsStartMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionsStartMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *SessionsStartMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *SessionsStartMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o SessionsStartMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionsStartMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["Data"] = o.Data.Get()
	}
	if !IsNil(o.MessageType) {
		toSerialize["MessageType"] = o.MessageType
	}
	return toSerialize, nil
}

type NullableSessionsStartMessage struct {
	value *SessionsStartMessage
	isSet bool
}

func (v NullableSessionsStartMessage) Get() *SessionsStartMessage {
	return v.value
}

func (v *NullableSessionsStartMessage) Set(val *SessionsStartMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionsStartMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionsStartMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionsStartMessage(val *SessionsStartMessage) *NullableSessionsStartMessage {
	return &NullableSessionsStartMessage{value: val, isSet: true}
}

func (v NullableSessionsStartMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionsStartMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


