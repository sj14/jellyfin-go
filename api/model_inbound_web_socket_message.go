/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// InboundWebSocketMessage - Represents the list of possible inbound websocket types
type InboundWebSocketMessage struct {
	ActivityLogEntryStartMessage *ActivityLogEntryStartMessage
	ActivityLogEntryStopMessage *ActivityLogEntryStopMessage
	InboundKeepAliveMessage *InboundKeepAliveMessage
	ScheduledTasksInfoStartMessage *ScheduledTasksInfoStartMessage
	ScheduledTasksInfoStopMessage *ScheduledTasksInfoStopMessage
	SessionsStartMessage *SessionsStartMessage
	SessionsStopMessage *SessionsStopMessage
}

// ActivityLogEntryStartMessageAsInboundWebSocketMessage is a convenience function that returns ActivityLogEntryStartMessage wrapped in InboundWebSocketMessage
func ActivityLogEntryStartMessageAsInboundWebSocketMessage(v *ActivityLogEntryStartMessage) InboundWebSocketMessage {
	return InboundWebSocketMessage{
		ActivityLogEntryStartMessage: v,
	}
}

// ActivityLogEntryStopMessageAsInboundWebSocketMessage is a convenience function that returns ActivityLogEntryStopMessage wrapped in InboundWebSocketMessage
func ActivityLogEntryStopMessageAsInboundWebSocketMessage(v *ActivityLogEntryStopMessage) InboundWebSocketMessage {
	return InboundWebSocketMessage{
		ActivityLogEntryStopMessage: v,
	}
}

// InboundKeepAliveMessageAsInboundWebSocketMessage is a convenience function that returns InboundKeepAliveMessage wrapped in InboundWebSocketMessage
func InboundKeepAliveMessageAsInboundWebSocketMessage(v *InboundKeepAliveMessage) InboundWebSocketMessage {
	return InboundWebSocketMessage{
		InboundKeepAliveMessage: v,
	}
}

// ScheduledTasksInfoStartMessageAsInboundWebSocketMessage is a convenience function that returns ScheduledTasksInfoStartMessage wrapped in InboundWebSocketMessage
func ScheduledTasksInfoStartMessageAsInboundWebSocketMessage(v *ScheduledTasksInfoStartMessage) InboundWebSocketMessage {
	return InboundWebSocketMessage{
		ScheduledTasksInfoStartMessage: v,
	}
}

// ScheduledTasksInfoStopMessageAsInboundWebSocketMessage is a convenience function that returns ScheduledTasksInfoStopMessage wrapped in InboundWebSocketMessage
func ScheduledTasksInfoStopMessageAsInboundWebSocketMessage(v *ScheduledTasksInfoStopMessage) InboundWebSocketMessage {
	return InboundWebSocketMessage{
		ScheduledTasksInfoStopMessage: v,
	}
}

// SessionsStartMessageAsInboundWebSocketMessage is a convenience function that returns SessionsStartMessage wrapped in InboundWebSocketMessage
func SessionsStartMessageAsInboundWebSocketMessage(v *SessionsStartMessage) InboundWebSocketMessage {
	return InboundWebSocketMessage{
		SessionsStartMessage: v,
	}
}

// SessionsStopMessageAsInboundWebSocketMessage is a convenience function that returns SessionsStopMessage wrapped in InboundWebSocketMessage
func SessionsStopMessageAsInboundWebSocketMessage(v *SessionsStopMessage) InboundWebSocketMessage {
	return InboundWebSocketMessage{
		SessionsStopMessage: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InboundWebSocketMessage) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ActivityLogEntryStartMessage
	err = newStrictDecoder(data).Decode(&dst.ActivityLogEntryStartMessage)
	if err == nil {
		jsonActivityLogEntryStartMessage, _ := json.Marshal(dst.ActivityLogEntryStartMessage)
		if string(jsonActivityLogEntryStartMessage) == "{}" { // empty struct
			dst.ActivityLogEntryStartMessage = nil
		} else {
			if err = validator.Validate(dst.ActivityLogEntryStartMessage); err != nil {
				dst.ActivityLogEntryStartMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActivityLogEntryStartMessage = nil
	}

	// try to unmarshal data into ActivityLogEntryStopMessage
	err = newStrictDecoder(data).Decode(&dst.ActivityLogEntryStopMessage)
	if err == nil {
		jsonActivityLogEntryStopMessage, _ := json.Marshal(dst.ActivityLogEntryStopMessage)
		if string(jsonActivityLogEntryStopMessage) == "{}" { // empty struct
			dst.ActivityLogEntryStopMessage = nil
		} else {
			if err = validator.Validate(dst.ActivityLogEntryStopMessage); err != nil {
				dst.ActivityLogEntryStopMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActivityLogEntryStopMessage = nil
	}

	// try to unmarshal data into InboundKeepAliveMessage
	err = newStrictDecoder(data).Decode(&dst.InboundKeepAliveMessage)
	if err == nil {
		jsonInboundKeepAliveMessage, _ := json.Marshal(dst.InboundKeepAliveMessage)
		if string(jsonInboundKeepAliveMessage) == "{}" { // empty struct
			dst.InboundKeepAliveMessage = nil
		} else {
			if err = validator.Validate(dst.InboundKeepAliveMessage); err != nil {
				dst.InboundKeepAliveMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.InboundKeepAliveMessage = nil
	}

	// try to unmarshal data into ScheduledTasksInfoStartMessage
	err = newStrictDecoder(data).Decode(&dst.ScheduledTasksInfoStartMessage)
	if err == nil {
		jsonScheduledTasksInfoStartMessage, _ := json.Marshal(dst.ScheduledTasksInfoStartMessage)
		if string(jsonScheduledTasksInfoStartMessage) == "{}" { // empty struct
			dst.ScheduledTasksInfoStartMessage = nil
		} else {
			if err = validator.Validate(dst.ScheduledTasksInfoStartMessage); err != nil {
				dst.ScheduledTasksInfoStartMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.ScheduledTasksInfoStartMessage = nil
	}

	// try to unmarshal data into ScheduledTasksInfoStopMessage
	err = newStrictDecoder(data).Decode(&dst.ScheduledTasksInfoStopMessage)
	if err == nil {
		jsonScheduledTasksInfoStopMessage, _ := json.Marshal(dst.ScheduledTasksInfoStopMessage)
		if string(jsonScheduledTasksInfoStopMessage) == "{}" { // empty struct
			dst.ScheduledTasksInfoStopMessage = nil
		} else {
			if err = validator.Validate(dst.ScheduledTasksInfoStopMessage); err != nil {
				dst.ScheduledTasksInfoStopMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.ScheduledTasksInfoStopMessage = nil
	}

	// try to unmarshal data into SessionsStartMessage
	err = newStrictDecoder(data).Decode(&dst.SessionsStartMessage)
	if err == nil {
		jsonSessionsStartMessage, _ := json.Marshal(dst.SessionsStartMessage)
		if string(jsonSessionsStartMessage) == "{}" { // empty struct
			dst.SessionsStartMessage = nil
		} else {
			if err = validator.Validate(dst.SessionsStartMessage); err != nil {
				dst.SessionsStartMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.SessionsStartMessage = nil
	}

	// try to unmarshal data into SessionsStopMessage
	err = newStrictDecoder(data).Decode(&dst.SessionsStopMessage)
	if err == nil {
		jsonSessionsStopMessage, _ := json.Marshal(dst.SessionsStopMessage)
		if string(jsonSessionsStopMessage) == "{}" { // empty struct
			dst.SessionsStopMessage = nil
		} else {
			if err = validator.Validate(dst.SessionsStopMessage); err != nil {
				dst.SessionsStopMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.SessionsStopMessage = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ActivityLogEntryStartMessage = nil
		dst.ActivityLogEntryStopMessage = nil
		dst.InboundKeepAliveMessage = nil
		dst.ScheduledTasksInfoStartMessage = nil
		dst.ScheduledTasksInfoStopMessage = nil
		dst.SessionsStartMessage = nil
		dst.SessionsStopMessage = nil

		return fmt.Errorf("data matches more than one schema in oneOf(InboundWebSocketMessage)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InboundWebSocketMessage)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InboundWebSocketMessage) MarshalJSON() ([]byte, error) {
	if src.ActivityLogEntryStartMessage != nil {
		return json.Marshal(&src.ActivityLogEntryStartMessage)
	}

	if src.ActivityLogEntryStopMessage != nil {
		return json.Marshal(&src.ActivityLogEntryStopMessage)
	}

	if src.InboundKeepAliveMessage != nil {
		return json.Marshal(&src.InboundKeepAliveMessage)
	}

	if src.ScheduledTasksInfoStartMessage != nil {
		return json.Marshal(&src.ScheduledTasksInfoStartMessage)
	}

	if src.ScheduledTasksInfoStopMessage != nil {
		return json.Marshal(&src.ScheduledTasksInfoStopMessage)
	}

	if src.SessionsStartMessage != nil {
		return json.Marshal(&src.SessionsStartMessage)
	}

	if src.SessionsStopMessage != nil {
		return json.Marshal(&src.SessionsStopMessage)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InboundWebSocketMessage) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ActivityLogEntryStartMessage != nil {
		return obj.ActivityLogEntryStartMessage
	}

	if obj.ActivityLogEntryStopMessage != nil {
		return obj.ActivityLogEntryStopMessage
	}

	if obj.InboundKeepAliveMessage != nil {
		return obj.InboundKeepAliveMessage
	}

	if obj.ScheduledTasksInfoStartMessage != nil {
		return obj.ScheduledTasksInfoStartMessage
	}

	if obj.ScheduledTasksInfoStopMessage != nil {
		return obj.ScheduledTasksInfoStopMessage
	}

	if obj.SessionsStartMessage != nil {
		return obj.SessionsStartMessage
	}

	if obj.SessionsStopMessage != nil {
		return obj.SessionsStopMessage
	}

	// all schemas are nil
	return nil
}

type NullableInboundWebSocketMessage struct {
	value *InboundWebSocketMessage
	isSet bool
}

func (v NullableInboundWebSocketMessage) Get() *InboundWebSocketMessage {
	return v.value
}

func (v *NullableInboundWebSocketMessage) Set(val *InboundWebSocketMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundWebSocketMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundWebSocketMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundWebSocketMessage(val *InboundWebSocketMessage) *NullableInboundWebSocketMessage {
	return &NullableInboundWebSocketMessage{value: val, isSet: true}
}

func (v NullableInboundWebSocketMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInboundWebSocketMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


