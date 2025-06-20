/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// WebSocketMessage - Represents the possible websocket types
type WebSocketMessage struct {
	InboundWebSocketMessage *InboundWebSocketMessage
	OutboundWebSocketMessage *OutboundWebSocketMessage
}

// InboundWebSocketMessageAsWebSocketMessage is a convenience function that returns InboundWebSocketMessage wrapped in WebSocketMessage
func InboundWebSocketMessageAsWebSocketMessage(v *InboundWebSocketMessage) WebSocketMessage {
	return WebSocketMessage{
		InboundWebSocketMessage: v,
	}
}

// OutboundWebSocketMessageAsWebSocketMessage is a convenience function that returns OutboundWebSocketMessage wrapped in WebSocketMessage
func OutboundWebSocketMessageAsWebSocketMessage(v *OutboundWebSocketMessage) WebSocketMessage {
	return WebSocketMessage{
		OutboundWebSocketMessage: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *WebSocketMessage) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InboundWebSocketMessage
	err = newStrictDecoder(data).Decode(&dst.InboundWebSocketMessage)
	if err == nil {
		jsonInboundWebSocketMessage, _ := json.Marshal(dst.InboundWebSocketMessage)
		if string(jsonInboundWebSocketMessage) == "{}" { // empty struct
			dst.InboundWebSocketMessage = nil
		} else {
			if err = validator.Validate(dst.InboundWebSocketMessage); err != nil {
				dst.InboundWebSocketMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.InboundWebSocketMessage = nil
	}

	// try to unmarshal data into OutboundWebSocketMessage
	err = newStrictDecoder(data).Decode(&dst.OutboundWebSocketMessage)
	if err == nil {
		jsonOutboundWebSocketMessage, _ := json.Marshal(dst.OutboundWebSocketMessage)
		if string(jsonOutboundWebSocketMessage) == "{}" { // empty struct
			dst.OutboundWebSocketMessage = nil
		} else {
			if err = validator.Validate(dst.OutboundWebSocketMessage); err != nil {
				dst.OutboundWebSocketMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.OutboundWebSocketMessage = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InboundWebSocketMessage = nil
		dst.OutboundWebSocketMessage = nil

		return fmt.Errorf("data matches more than one schema in oneOf(WebSocketMessage)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(WebSocketMessage)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src WebSocketMessage) MarshalJSON() ([]byte, error) {
	if src.InboundWebSocketMessage != nil {
		return json.Marshal(&src.InboundWebSocketMessage)
	}

	if src.OutboundWebSocketMessage != nil {
		return json.Marshal(&src.OutboundWebSocketMessage)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *WebSocketMessage) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InboundWebSocketMessage != nil {
		return obj.InboundWebSocketMessage
	}

	if obj.OutboundWebSocketMessage != nil {
		return obj.OutboundWebSocketMessage
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj WebSocketMessage) GetActualInstanceValue() (interface{}) {
	if obj.InboundWebSocketMessage != nil {
		return *obj.InboundWebSocketMessage
	}

	if obj.OutboundWebSocketMessage != nil {
		return *obj.OutboundWebSocketMessage
	}

	// all schemas are nil
	return nil
}

type NullableWebSocketMessage struct {
	value *WebSocketMessage
	isSet bool
}

func (v NullableWebSocketMessage) Get() *WebSocketMessage {
	return v.value
}

func (v *NullableWebSocketMessage) Set(val *WebSocketMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableWebSocketMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableWebSocketMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebSocketMessage(val *WebSocketMessage) *NullableWebSocketMessage {
	return &NullableWebSocketMessage{value: val, isSet: true}
}

func (v NullableWebSocketMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebSocketMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


