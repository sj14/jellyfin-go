/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ActivityLogEntryStartMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityLogEntryStartMessage{}

// ActivityLogEntryStartMessage Activity log entry start message.  Data is the timing data encoded as \"$initialDelay,$interval\" in ms.
type ActivityLogEntryStartMessage struct {
	// Gets or sets the data.
	Data NullableString `json:"Data,omitempty"`
	// The different kinds of messages that are used in the WebSocket api.
	MessageType *SessionMessageType `json:"MessageType,omitempty"`
}

// NewActivityLogEntryStartMessage instantiates a new ActivityLogEntryStartMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityLogEntryStartMessage() *ActivityLogEntryStartMessage {
	this := ActivityLogEntryStartMessage{}
	return &this
}

// NewActivityLogEntryStartMessageWithDefaults instantiates a new ActivityLogEntryStartMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityLogEntryStartMessageWithDefaults() *ActivityLogEntryStartMessage {
	this := ActivityLogEntryStartMessage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActivityLogEntryStartMessage) GetData() string {
	if o == nil || IsNil(o.Data.Get()) {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActivityLogEntryStartMessage) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *ActivityLogEntryStartMessage) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *ActivityLogEntryStartMessage) SetData(v string) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *ActivityLogEntryStartMessage) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *ActivityLogEntryStartMessage) UnsetData() {
	o.Data.Unset()
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *ActivityLogEntryStartMessage) GetMessageType() SessionMessageType {
	if o == nil || IsNil(o.MessageType) {
		var ret SessionMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLogEntryStartMessage) GetMessageTypeOk() (*SessionMessageType, bool) {
	if o == nil || IsNil(o.MessageType) {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *ActivityLogEntryStartMessage) HasMessageType() bool {
	if o != nil && !IsNil(o.MessageType) {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given SessionMessageType and assigns it to the MessageType field.
func (o *ActivityLogEntryStartMessage) SetMessageType(v SessionMessageType) {
	o.MessageType = &v
}

func (o ActivityLogEntryStartMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityLogEntryStartMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data.IsSet() {
		toSerialize["Data"] = o.Data.Get()
	}
	if !IsNil(o.MessageType) {
		toSerialize["MessageType"] = o.MessageType
	}
	return toSerialize, nil
}

type NullableActivityLogEntryStartMessage struct {
	value *ActivityLogEntryStartMessage
	isSet bool
}

func (v NullableActivityLogEntryStartMessage) Get() *ActivityLogEntryStartMessage {
	return v.value
}

func (v *NullableActivityLogEntryStartMessage) Set(val *ActivityLogEntryStartMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityLogEntryStartMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityLogEntryStartMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityLogEntryStartMessage(val *ActivityLogEntryStartMessage) *NullableActivityLogEntryStartMessage {
	return &NullableActivityLogEntryStartMessage{value: val, isSet: true}
}

func (v NullableActivityLogEntryStartMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityLogEntryStartMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


