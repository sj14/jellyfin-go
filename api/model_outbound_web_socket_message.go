/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// OutboundWebSocketMessage - Represents the list of possible outbound websocket types
type OutboundWebSocketMessage struct {
	ActivityLogEntryMessage *ActivityLogEntryMessage
	ForceKeepAliveMessage *ForceKeepAliveMessage
	GeneralCommandMessage *GeneralCommandMessage
	LibraryChangedMessage *LibraryChangedMessage
	OutboundKeepAliveMessage *OutboundKeepAliveMessage
	PlayMessage *PlayMessage
	PlaystateMessage *PlaystateMessage
	PluginInstallationCancelledMessage *PluginInstallationCancelledMessage
	PluginInstallationCompletedMessage *PluginInstallationCompletedMessage
	PluginInstallationFailedMessage *PluginInstallationFailedMessage
	PluginInstallingMessage *PluginInstallingMessage
	PluginUninstalledMessage *PluginUninstalledMessage
	RefreshProgressMessage *RefreshProgressMessage
	RestartRequiredMessage *RestartRequiredMessage
	ScheduledTaskEndedMessage *ScheduledTaskEndedMessage
	ScheduledTasksInfoMessage *ScheduledTasksInfoMessage
	SeriesTimerCancelledMessage *SeriesTimerCancelledMessage
	SeriesTimerCreatedMessage *SeriesTimerCreatedMessage
	ServerRestartingMessage *ServerRestartingMessage
	ServerShuttingDownMessage *ServerShuttingDownMessage
	SessionsMessage *SessionsMessage
	SyncPlayCommandMessage *SyncPlayCommandMessage
	SyncPlayGroupUpdateCommandMessage *SyncPlayGroupUpdateCommandMessage
	TimerCancelledMessage *TimerCancelledMessage
	TimerCreatedMessage *TimerCreatedMessage
	UserDataChangedMessage *UserDataChangedMessage
	UserDeletedMessage *UserDeletedMessage
	UserUpdatedMessage *UserUpdatedMessage
}

// ActivityLogEntryMessageAsOutboundWebSocketMessage is a convenience function that returns ActivityLogEntryMessage wrapped in OutboundWebSocketMessage
func ActivityLogEntryMessageAsOutboundWebSocketMessage(v *ActivityLogEntryMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		ActivityLogEntryMessage: v,
	}
}

// ForceKeepAliveMessageAsOutboundWebSocketMessage is a convenience function that returns ForceKeepAliveMessage wrapped in OutboundWebSocketMessage
func ForceKeepAliveMessageAsOutboundWebSocketMessage(v *ForceKeepAliveMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		ForceKeepAliveMessage: v,
	}
}

// GeneralCommandMessageAsOutboundWebSocketMessage is a convenience function that returns GeneralCommandMessage wrapped in OutboundWebSocketMessage
func GeneralCommandMessageAsOutboundWebSocketMessage(v *GeneralCommandMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		GeneralCommandMessage: v,
	}
}

// LibraryChangedMessageAsOutboundWebSocketMessage is a convenience function that returns LibraryChangedMessage wrapped in OutboundWebSocketMessage
func LibraryChangedMessageAsOutboundWebSocketMessage(v *LibraryChangedMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		LibraryChangedMessage: v,
	}
}

// OutboundKeepAliveMessageAsOutboundWebSocketMessage is a convenience function that returns OutboundKeepAliveMessage wrapped in OutboundWebSocketMessage
func OutboundKeepAliveMessageAsOutboundWebSocketMessage(v *OutboundKeepAliveMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		OutboundKeepAliveMessage: v,
	}
}

// PlayMessageAsOutboundWebSocketMessage is a convenience function that returns PlayMessage wrapped in OutboundWebSocketMessage
func PlayMessageAsOutboundWebSocketMessage(v *PlayMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		PlayMessage: v,
	}
}

// PlaystateMessageAsOutboundWebSocketMessage is a convenience function that returns PlaystateMessage wrapped in OutboundWebSocketMessage
func PlaystateMessageAsOutboundWebSocketMessage(v *PlaystateMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		PlaystateMessage: v,
	}
}

// PluginInstallationCancelledMessageAsOutboundWebSocketMessage is a convenience function that returns PluginInstallationCancelledMessage wrapped in OutboundWebSocketMessage
func PluginInstallationCancelledMessageAsOutboundWebSocketMessage(v *PluginInstallationCancelledMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		PluginInstallationCancelledMessage: v,
	}
}

// PluginInstallationCompletedMessageAsOutboundWebSocketMessage is a convenience function that returns PluginInstallationCompletedMessage wrapped in OutboundWebSocketMessage
func PluginInstallationCompletedMessageAsOutboundWebSocketMessage(v *PluginInstallationCompletedMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		PluginInstallationCompletedMessage: v,
	}
}

// PluginInstallationFailedMessageAsOutboundWebSocketMessage is a convenience function that returns PluginInstallationFailedMessage wrapped in OutboundWebSocketMessage
func PluginInstallationFailedMessageAsOutboundWebSocketMessage(v *PluginInstallationFailedMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		PluginInstallationFailedMessage: v,
	}
}

// PluginInstallingMessageAsOutboundWebSocketMessage is a convenience function that returns PluginInstallingMessage wrapped in OutboundWebSocketMessage
func PluginInstallingMessageAsOutboundWebSocketMessage(v *PluginInstallingMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		PluginInstallingMessage: v,
	}
}

// PluginUninstalledMessageAsOutboundWebSocketMessage is a convenience function that returns PluginUninstalledMessage wrapped in OutboundWebSocketMessage
func PluginUninstalledMessageAsOutboundWebSocketMessage(v *PluginUninstalledMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		PluginUninstalledMessage: v,
	}
}

// RefreshProgressMessageAsOutboundWebSocketMessage is a convenience function that returns RefreshProgressMessage wrapped in OutboundWebSocketMessage
func RefreshProgressMessageAsOutboundWebSocketMessage(v *RefreshProgressMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		RefreshProgressMessage: v,
	}
}

// RestartRequiredMessageAsOutboundWebSocketMessage is a convenience function that returns RestartRequiredMessage wrapped in OutboundWebSocketMessage
func RestartRequiredMessageAsOutboundWebSocketMessage(v *RestartRequiredMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		RestartRequiredMessage: v,
	}
}

// ScheduledTaskEndedMessageAsOutboundWebSocketMessage is a convenience function that returns ScheduledTaskEndedMessage wrapped in OutboundWebSocketMessage
func ScheduledTaskEndedMessageAsOutboundWebSocketMessage(v *ScheduledTaskEndedMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		ScheduledTaskEndedMessage: v,
	}
}

// ScheduledTasksInfoMessageAsOutboundWebSocketMessage is a convenience function that returns ScheduledTasksInfoMessage wrapped in OutboundWebSocketMessage
func ScheduledTasksInfoMessageAsOutboundWebSocketMessage(v *ScheduledTasksInfoMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		ScheduledTasksInfoMessage: v,
	}
}

// SeriesTimerCancelledMessageAsOutboundWebSocketMessage is a convenience function that returns SeriesTimerCancelledMessage wrapped in OutboundWebSocketMessage
func SeriesTimerCancelledMessageAsOutboundWebSocketMessage(v *SeriesTimerCancelledMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		SeriesTimerCancelledMessage: v,
	}
}

// SeriesTimerCreatedMessageAsOutboundWebSocketMessage is a convenience function that returns SeriesTimerCreatedMessage wrapped in OutboundWebSocketMessage
func SeriesTimerCreatedMessageAsOutboundWebSocketMessage(v *SeriesTimerCreatedMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		SeriesTimerCreatedMessage: v,
	}
}

// ServerRestartingMessageAsOutboundWebSocketMessage is a convenience function that returns ServerRestartingMessage wrapped in OutboundWebSocketMessage
func ServerRestartingMessageAsOutboundWebSocketMessage(v *ServerRestartingMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		ServerRestartingMessage: v,
	}
}

// ServerShuttingDownMessageAsOutboundWebSocketMessage is a convenience function that returns ServerShuttingDownMessage wrapped in OutboundWebSocketMessage
func ServerShuttingDownMessageAsOutboundWebSocketMessage(v *ServerShuttingDownMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		ServerShuttingDownMessage: v,
	}
}

// SessionsMessageAsOutboundWebSocketMessage is a convenience function that returns SessionsMessage wrapped in OutboundWebSocketMessage
func SessionsMessageAsOutboundWebSocketMessage(v *SessionsMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		SessionsMessage: v,
	}
}

// SyncPlayCommandMessageAsOutboundWebSocketMessage is a convenience function that returns SyncPlayCommandMessage wrapped in OutboundWebSocketMessage
func SyncPlayCommandMessageAsOutboundWebSocketMessage(v *SyncPlayCommandMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		SyncPlayCommandMessage: v,
	}
}

// SyncPlayGroupUpdateCommandMessageAsOutboundWebSocketMessage is a convenience function that returns SyncPlayGroupUpdateCommandMessage wrapped in OutboundWebSocketMessage
func SyncPlayGroupUpdateCommandMessageAsOutboundWebSocketMessage(v *SyncPlayGroupUpdateCommandMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		SyncPlayGroupUpdateCommandMessage: v,
	}
}

// TimerCancelledMessageAsOutboundWebSocketMessage is a convenience function that returns TimerCancelledMessage wrapped in OutboundWebSocketMessage
func TimerCancelledMessageAsOutboundWebSocketMessage(v *TimerCancelledMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		TimerCancelledMessage: v,
	}
}

// TimerCreatedMessageAsOutboundWebSocketMessage is a convenience function that returns TimerCreatedMessage wrapped in OutboundWebSocketMessage
func TimerCreatedMessageAsOutboundWebSocketMessage(v *TimerCreatedMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		TimerCreatedMessage: v,
	}
}

// UserDataChangedMessageAsOutboundWebSocketMessage is a convenience function that returns UserDataChangedMessage wrapped in OutboundWebSocketMessage
func UserDataChangedMessageAsOutboundWebSocketMessage(v *UserDataChangedMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		UserDataChangedMessage: v,
	}
}

// UserDeletedMessageAsOutboundWebSocketMessage is a convenience function that returns UserDeletedMessage wrapped in OutboundWebSocketMessage
func UserDeletedMessageAsOutboundWebSocketMessage(v *UserDeletedMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		UserDeletedMessage: v,
	}
}

// UserUpdatedMessageAsOutboundWebSocketMessage is a convenience function that returns UserUpdatedMessage wrapped in OutboundWebSocketMessage
func UserUpdatedMessageAsOutboundWebSocketMessage(v *UserUpdatedMessage) OutboundWebSocketMessage {
	return OutboundWebSocketMessage{
		UserUpdatedMessage: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *OutboundWebSocketMessage) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ActivityLogEntryMessage
	err = newStrictDecoder(data).Decode(&dst.ActivityLogEntryMessage)
	if err == nil {
		jsonActivityLogEntryMessage, _ := json.Marshal(dst.ActivityLogEntryMessage)
		if string(jsonActivityLogEntryMessage) == "{}" { // empty struct
			dst.ActivityLogEntryMessage = nil
		} else {
			if err = validator.Validate(dst.ActivityLogEntryMessage); err != nil {
				dst.ActivityLogEntryMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActivityLogEntryMessage = nil
	}

	// try to unmarshal data into ForceKeepAliveMessage
	err = newStrictDecoder(data).Decode(&dst.ForceKeepAliveMessage)
	if err == nil {
		jsonForceKeepAliveMessage, _ := json.Marshal(dst.ForceKeepAliveMessage)
		if string(jsonForceKeepAliveMessage) == "{}" { // empty struct
			dst.ForceKeepAliveMessage = nil
		} else {
			if err = validator.Validate(dst.ForceKeepAliveMessage); err != nil {
				dst.ForceKeepAliveMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.ForceKeepAliveMessage = nil
	}

	// try to unmarshal data into GeneralCommandMessage
	err = newStrictDecoder(data).Decode(&dst.GeneralCommandMessage)
	if err == nil {
		jsonGeneralCommandMessage, _ := json.Marshal(dst.GeneralCommandMessage)
		if string(jsonGeneralCommandMessage) == "{}" { // empty struct
			dst.GeneralCommandMessage = nil
		} else {
			if err = validator.Validate(dst.GeneralCommandMessage); err != nil {
				dst.GeneralCommandMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.GeneralCommandMessage = nil
	}

	// try to unmarshal data into LibraryChangedMessage
	err = newStrictDecoder(data).Decode(&dst.LibraryChangedMessage)
	if err == nil {
		jsonLibraryChangedMessage, _ := json.Marshal(dst.LibraryChangedMessage)
		if string(jsonLibraryChangedMessage) == "{}" { // empty struct
			dst.LibraryChangedMessage = nil
		} else {
			if err = validator.Validate(dst.LibraryChangedMessage); err != nil {
				dst.LibraryChangedMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.LibraryChangedMessage = nil
	}

	// try to unmarshal data into OutboundKeepAliveMessage
	err = newStrictDecoder(data).Decode(&dst.OutboundKeepAliveMessage)
	if err == nil {
		jsonOutboundKeepAliveMessage, _ := json.Marshal(dst.OutboundKeepAliveMessage)
		if string(jsonOutboundKeepAliveMessage) == "{}" { // empty struct
			dst.OutboundKeepAliveMessage = nil
		} else {
			if err = validator.Validate(dst.OutboundKeepAliveMessage); err != nil {
				dst.OutboundKeepAliveMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.OutboundKeepAliveMessage = nil
	}

	// try to unmarshal data into PlayMessage
	err = newStrictDecoder(data).Decode(&dst.PlayMessage)
	if err == nil {
		jsonPlayMessage, _ := json.Marshal(dst.PlayMessage)
		if string(jsonPlayMessage) == "{}" { // empty struct
			dst.PlayMessage = nil
		} else {
			if err = validator.Validate(dst.PlayMessage); err != nil {
				dst.PlayMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.PlayMessage = nil
	}

	// try to unmarshal data into PlaystateMessage
	err = newStrictDecoder(data).Decode(&dst.PlaystateMessage)
	if err == nil {
		jsonPlaystateMessage, _ := json.Marshal(dst.PlaystateMessage)
		if string(jsonPlaystateMessage) == "{}" { // empty struct
			dst.PlaystateMessage = nil
		} else {
			if err = validator.Validate(dst.PlaystateMessage); err != nil {
				dst.PlaystateMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.PlaystateMessage = nil
	}

	// try to unmarshal data into PluginInstallationCancelledMessage
	err = newStrictDecoder(data).Decode(&dst.PluginInstallationCancelledMessage)
	if err == nil {
		jsonPluginInstallationCancelledMessage, _ := json.Marshal(dst.PluginInstallationCancelledMessage)
		if string(jsonPluginInstallationCancelledMessage) == "{}" { // empty struct
			dst.PluginInstallationCancelledMessage = nil
		} else {
			if err = validator.Validate(dst.PluginInstallationCancelledMessage); err != nil {
				dst.PluginInstallationCancelledMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.PluginInstallationCancelledMessage = nil
	}

	// try to unmarshal data into PluginInstallationCompletedMessage
	err = newStrictDecoder(data).Decode(&dst.PluginInstallationCompletedMessage)
	if err == nil {
		jsonPluginInstallationCompletedMessage, _ := json.Marshal(dst.PluginInstallationCompletedMessage)
		if string(jsonPluginInstallationCompletedMessage) == "{}" { // empty struct
			dst.PluginInstallationCompletedMessage = nil
		} else {
			if err = validator.Validate(dst.PluginInstallationCompletedMessage); err != nil {
				dst.PluginInstallationCompletedMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.PluginInstallationCompletedMessage = nil
	}

	// try to unmarshal data into PluginInstallationFailedMessage
	err = newStrictDecoder(data).Decode(&dst.PluginInstallationFailedMessage)
	if err == nil {
		jsonPluginInstallationFailedMessage, _ := json.Marshal(dst.PluginInstallationFailedMessage)
		if string(jsonPluginInstallationFailedMessage) == "{}" { // empty struct
			dst.PluginInstallationFailedMessage = nil
		} else {
			if err = validator.Validate(dst.PluginInstallationFailedMessage); err != nil {
				dst.PluginInstallationFailedMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.PluginInstallationFailedMessage = nil
	}

	// try to unmarshal data into PluginInstallingMessage
	err = newStrictDecoder(data).Decode(&dst.PluginInstallingMessage)
	if err == nil {
		jsonPluginInstallingMessage, _ := json.Marshal(dst.PluginInstallingMessage)
		if string(jsonPluginInstallingMessage) == "{}" { // empty struct
			dst.PluginInstallingMessage = nil
		} else {
			if err = validator.Validate(dst.PluginInstallingMessage); err != nil {
				dst.PluginInstallingMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.PluginInstallingMessage = nil
	}

	// try to unmarshal data into PluginUninstalledMessage
	err = newStrictDecoder(data).Decode(&dst.PluginUninstalledMessage)
	if err == nil {
		jsonPluginUninstalledMessage, _ := json.Marshal(dst.PluginUninstalledMessage)
		if string(jsonPluginUninstalledMessage) == "{}" { // empty struct
			dst.PluginUninstalledMessage = nil
		} else {
			if err = validator.Validate(dst.PluginUninstalledMessage); err != nil {
				dst.PluginUninstalledMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.PluginUninstalledMessage = nil
	}

	// try to unmarshal data into RefreshProgressMessage
	err = newStrictDecoder(data).Decode(&dst.RefreshProgressMessage)
	if err == nil {
		jsonRefreshProgressMessage, _ := json.Marshal(dst.RefreshProgressMessage)
		if string(jsonRefreshProgressMessage) == "{}" { // empty struct
			dst.RefreshProgressMessage = nil
		} else {
			if err = validator.Validate(dst.RefreshProgressMessage); err != nil {
				dst.RefreshProgressMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.RefreshProgressMessage = nil
	}

	// try to unmarshal data into RestartRequiredMessage
	err = newStrictDecoder(data).Decode(&dst.RestartRequiredMessage)
	if err == nil {
		jsonRestartRequiredMessage, _ := json.Marshal(dst.RestartRequiredMessage)
		if string(jsonRestartRequiredMessage) == "{}" { // empty struct
			dst.RestartRequiredMessage = nil
		} else {
			if err = validator.Validate(dst.RestartRequiredMessage); err != nil {
				dst.RestartRequiredMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.RestartRequiredMessage = nil
	}

	// try to unmarshal data into ScheduledTaskEndedMessage
	err = newStrictDecoder(data).Decode(&dst.ScheduledTaskEndedMessage)
	if err == nil {
		jsonScheduledTaskEndedMessage, _ := json.Marshal(dst.ScheduledTaskEndedMessage)
		if string(jsonScheduledTaskEndedMessage) == "{}" { // empty struct
			dst.ScheduledTaskEndedMessage = nil
		} else {
			if err = validator.Validate(dst.ScheduledTaskEndedMessage); err != nil {
				dst.ScheduledTaskEndedMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.ScheduledTaskEndedMessage = nil
	}

	// try to unmarshal data into ScheduledTasksInfoMessage
	err = newStrictDecoder(data).Decode(&dst.ScheduledTasksInfoMessage)
	if err == nil {
		jsonScheduledTasksInfoMessage, _ := json.Marshal(dst.ScheduledTasksInfoMessage)
		if string(jsonScheduledTasksInfoMessage) == "{}" { // empty struct
			dst.ScheduledTasksInfoMessage = nil
		} else {
			if err = validator.Validate(dst.ScheduledTasksInfoMessage); err != nil {
				dst.ScheduledTasksInfoMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.ScheduledTasksInfoMessage = nil
	}

	// try to unmarshal data into SeriesTimerCancelledMessage
	err = newStrictDecoder(data).Decode(&dst.SeriesTimerCancelledMessage)
	if err == nil {
		jsonSeriesTimerCancelledMessage, _ := json.Marshal(dst.SeriesTimerCancelledMessage)
		if string(jsonSeriesTimerCancelledMessage) == "{}" { // empty struct
			dst.SeriesTimerCancelledMessage = nil
		} else {
			if err = validator.Validate(dst.SeriesTimerCancelledMessage); err != nil {
				dst.SeriesTimerCancelledMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.SeriesTimerCancelledMessage = nil
	}

	// try to unmarshal data into SeriesTimerCreatedMessage
	err = newStrictDecoder(data).Decode(&dst.SeriesTimerCreatedMessage)
	if err == nil {
		jsonSeriesTimerCreatedMessage, _ := json.Marshal(dst.SeriesTimerCreatedMessage)
		if string(jsonSeriesTimerCreatedMessage) == "{}" { // empty struct
			dst.SeriesTimerCreatedMessage = nil
		} else {
			if err = validator.Validate(dst.SeriesTimerCreatedMessage); err != nil {
				dst.SeriesTimerCreatedMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.SeriesTimerCreatedMessage = nil
	}

	// try to unmarshal data into ServerRestartingMessage
	err = newStrictDecoder(data).Decode(&dst.ServerRestartingMessage)
	if err == nil {
		jsonServerRestartingMessage, _ := json.Marshal(dst.ServerRestartingMessage)
		if string(jsonServerRestartingMessage) == "{}" { // empty struct
			dst.ServerRestartingMessage = nil
		} else {
			if err = validator.Validate(dst.ServerRestartingMessage); err != nil {
				dst.ServerRestartingMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.ServerRestartingMessage = nil
	}

	// try to unmarshal data into ServerShuttingDownMessage
	err = newStrictDecoder(data).Decode(&dst.ServerShuttingDownMessage)
	if err == nil {
		jsonServerShuttingDownMessage, _ := json.Marshal(dst.ServerShuttingDownMessage)
		if string(jsonServerShuttingDownMessage) == "{}" { // empty struct
			dst.ServerShuttingDownMessage = nil
		} else {
			if err = validator.Validate(dst.ServerShuttingDownMessage); err != nil {
				dst.ServerShuttingDownMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.ServerShuttingDownMessage = nil
	}

	// try to unmarshal data into SessionsMessage
	err = newStrictDecoder(data).Decode(&dst.SessionsMessage)
	if err == nil {
		jsonSessionsMessage, _ := json.Marshal(dst.SessionsMessage)
		if string(jsonSessionsMessage) == "{}" { // empty struct
			dst.SessionsMessage = nil
		} else {
			if err = validator.Validate(dst.SessionsMessage); err != nil {
				dst.SessionsMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.SessionsMessage = nil
	}

	// try to unmarshal data into SyncPlayCommandMessage
	err = newStrictDecoder(data).Decode(&dst.SyncPlayCommandMessage)
	if err == nil {
		jsonSyncPlayCommandMessage, _ := json.Marshal(dst.SyncPlayCommandMessage)
		if string(jsonSyncPlayCommandMessage) == "{}" { // empty struct
			dst.SyncPlayCommandMessage = nil
		} else {
			if err = validator.Validate(dst.SyncPlayCommandMessage); err != nil {
				dst.SyncPlayCommandMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.SyncPlayCommandMessage = nil
	}

	// try to unmarshal data into SyncPlayGroupUpdateCommandMessage
	err = newStrictDecoder(data).Decode(&dst.SyncPlayGroupUpdateCommandMessage)
	if err == nil {
		jsonSyncPlayGroupUpdateCommandMessage, _ := json.Marshal(dst.SyncPlayGroupUpdateCommandMessage)
		if string(jsonSyncPlayGroupUpdateCommandMessage) == "{}" { // empty struct
			dst.SyncPlayGroupUpdateCommandMessage = nil
		} else {
			if err = validator.Validate(dst.SyncPlayGroupUpdateCommandMessage); err != nil {
				dst.SyncPlayGroupUpdateCommandMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.SyncPlayGroupUpdateCommandMessage = nil
	}

	// try to unmarshal data into TimerCancelledMessage
	err = newStrictDecoder(data).Decode(&dst.TimerCancelledMessage)
	if err == nil {
		jsonTimerCancelledMessage, _ := json.Marshal(dst.TimerCancelledMessage)
		if string(jsonTimerCancelledMessage) == "{}" { // empty struct
			dst.TimerCancelledMessage = nil
		} else {
			if err = validator.Validate(dst.TimerCancelledMessage); err != nil {
				dst.TimerCancelledMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.TimerCancelledMessage = nil
	}

	// try to unmarshal data into TimerCreatedMessage
	err = newStrictDecoder(data).Decode(&dst.TimerCreatedMessage)
	if err == nil {
		jsonTimerCreatedMessage, _ := json.Marshal(dst.TimerCreatedMessage)
		if string(jsonTimerCreatedMessage) == "{}" { // empty struct
			dst.TimerCreatedMessage = nil
		} else {
			if err = validator.Validate(dst.TimerCreatedMessage); err != nil {
				dst.TimerCreatedMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.TimerCreatedMessage = nil
	}

	// try to unmarshal data into UserDataChangedMessage
	err = newStrictDecoder(data).Decode(&dst.UserDataChangedMessage)
	if err == nil {
		jsonUserDataChangedMessage, _ := json.Marshal(dst.UserDataChangedMessage)
		if string(jsonUserDataChangedMessage) == "{}" { // empty struct
			dst.UserDataChangedMessage = nil
		} else {
			if err = validator.Validate(dst.UserDataChangedMessage); err != nil {
				dst.UserDataChangedMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserDataChangedMessage = nil
	}

	// try to unmarshal data into UserDeletedMessage
	err = newStrictDecoder(data).Decode(&dst.UserDeletedMessage)
	if err == nil {
		jsonUserDeletedMessage, _ := json.Marshal(dst.UserDeletedMessage)
		if string(jsonUserDeletedMessage) == "{}" { // empty struct
			dst.UserDeletedMessage = nil
		} else {
			if err = validator.Validate(dst.UserDeletedMessage); err != nil {
				dst.UserDeletedMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserDeletedMessage = nil
	}

	// try to unmarshal data into UserUpdatedMessage
	err = newStrictDecoder(data).Decode(&dst.UserUpdatedMessage)
	if err == nil {
		jsonUserUpdatedMessage, _ := json.Marshal(dst.UserUpdatedMessage)
		if string(jsonUserUpdatedMessage) == "{}" { // empty struct
			dst.UserUpdatedMessage = nil
		} else {
			if err = validator.Validate(dst.UserUpdatedMessage); err != nil {
				dst.UserUpdatedMessage = nil
			} else {
				match++
			}
		}
	} else {
		dst.UserUpdatedMessage = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ActivityLogEntryMessage = nil
		dst.ForceKeepAliveMessage = nil
		dst.GeneralCommandMessage = nil
		dst.LibraryChangedMessage = nil
		dst.OutboundKeepAliveMessage = nil
		dst.PlayMessage = nil
		dst.PlaystateMessage = nil
		dst.PluginInstallationCancelledMessage = nil
		dst.PluginInstallationCompletedMessage = nil
		dst.PluginInstallationFailedMessage = nil
		dst.PluginInstallingMessage = nil
		dst.PluginUninstalledMessage = nil
		dst.RefreshProgressMessage = nil
		dst.RestartRequiredMessage = nil
		dst.ScheduledTaskEndedMessage = nil
		dst.ScheduledTasksInfoMessage = nil
		dst.SeriesTimerCancelledMessage = nil
		dst.SeriesTimerCreatedMessage = nil
		dst.ServerRestartingMessage = nil
		dst.ServerShuttingDownMessage = nil
		dst.SessionsMessage = nil
		dst.SyncPlayCommandMessage = nil
		dst.SyncPlayGroupUpdateCommandMessage = nil
		dst.TimerCancelledMessage = nil
		dst.TimerCreatedMessage = nil
		dst.UserDataChangedMessage = nil
		dst.UserDeletedMessage = nil
		dst.UserUpdatedMessage = nil

		return fmt.Errorf("data matches more than one schema in oneOf(OutboundWebSocketMessage)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(OutboundWebSocketMessage)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src OutboundWebSocketMessage) MarshalJSON() ([]byte, error) {
	if src.ActivityLogEntryMessage != nil {
		return json.Marshal(&src.ActivityLogEntryMessage)
	}

	if src.ForceKeepAliveMessage != nil {
		return json.Marshal(&src.ForceKeepAliveMessage)
	}

	if src.GeneralCommandMessage != nil {
		return json.Marshal(&src.GeneralCommandMessage)
	}

	if src.LibraryChangedMessage != nil {
		return json.Marshal(&src.LibraryChangedMessage)
	}

	if src.OutboundKeepAliveMessage != nil {
		return json.Marshal(&src.OutboundKeepAliveMessage)
	}

	if src.PlayMessage != nil {
		return json.Marshal(&src.PlayMessage)
	}

	if src.PlaystateMessage != nil {
		return json.Marshal(&src.PlaystateMessage)
	}

	if src.PluginInstallationCancelledMessage != nil {
		return json.Marshal(&src.PluginInstallationCancelledMessage)
	}

	if src.PluginInstallationCompletedMessage != nil {
		return json.Marshal(&src.PluginInstallationCompletedMessage)
	}

	if src.PluginInstallationFailedMessage != nil {
		return json.Marshal(&src.PluginInstallationFailedMessage)
	}

	if src.PluginInstallingMessage != nil {
		return json.Marshal(&src.PluginInstallingMessage)
	}

	if src.PluginUninstalledMessage != nil {
		return json.Marshal(&src.PluginUninstalledMessage)
	}

	if src.RefreshProgressMessage != nil {
		return json.Marshal(&src.RefreshProgressMessage)
	}

	if src.RestartRequiredMessage != nil {
		return json.Marshal(&src.RestartRequiredMessage)
	}

	if src.ScheduledTaskEndedMessage != nil {
		return json.Marshal(&src.ScheduledTaskEndedMessage)
	}

	if src.ScheduledTasksInfoMessage != nil {
		return json.Marshal(&src.ScheduledTasksInfoMessage)
	}

	if src.SeriesTimerCancelledMessage != nil {
		return json.Marshal(&src.SeriesTimerCancelledMessage)
	}

	if src.SeriesTimerCreatedMessage != nil {
		return json.Marshal(&src.SeriesTimerCreatedMessage)
	}

	if src.ServerRestartingMessage != nil {
		return json.Marshal(&src.ServerRestartingMessage)
	}

	if src.ServerShuttingDownMessage != nil {
		return json.Marshal(&src.ServerShuttingDownMessage)
	}

	if src.SessionsMessage != nil {
		return json.Marshal(&src.SessionsMessage)
	}

	if src.SyncPlayCommandMessage != nil {
		return json.Marshal(&src.SyncPlayCommandMessage)
	}

	if src.SyncPlayGroupUpdateCommandMessage != nil {
		return json.Marshal(&src.SyncPlayGroupUpdateCommandMessage)
	}

	if src.TimerCancelledMessage != nil {
		return json.Marshal(&src.TimerCancelledMessage)
	}

	if src.TimerCreatedMessage != nil {
		return json.Marshal(&src.TimerCreatedMessage)
	}

	if src.UserDataChangedMessage != nil {
		return json.Marshal(&src.UserDataChangedMessage)
	}

	if src.UserDeletedMessage != nil {
		return json.Marshal(&src.UserDeletedMessage)
	}

	if src.UserUpdatedMessage != nil {
		return json.Marshal(&src.UserUpdatedMessage)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *OutboundWebSocketMessage) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ActivityLogEntryMessage != nil {
		return obj.ActivityLogEntryMessage
	}

	if obj.ForceKeepAliveMessage != nil {
		return obj.ForceKeepAliveMessage
	}

	if obj.GeneralCommandMessage != nil {
		return obj.GeneralCommandMessage
	}

	if obj.LibraryChangedMessage != nil {
		return obj.LibraryChangedMessage
	}

	if obj.OutboundKeepAliveMessage != nil {
		return obj.OutboundKeepAliveMessage
	}

	if obj.PlayMessage != nil {
		return obj.PlayMessage
	}

	if obj.PlaystateMessage != nil {
		return obj.PlaystateMessage
	}

	if obj.PluginInstallationCancelledMessage != nil {
		return obj.PluginInstallationCancelledMessage
	}

	if obj.PluginInstallationCompletedMessage != nil {
		return obj.PluginInstallationCompletedMessage
	}

	if obj.PluginInstallationFailedMessage != nil {
		return obj.PluginInstallationFailedMessage
	}

	if obj.PluginInstallingMessage != nil {
		return obj.PluginInstallingMessage
	}

	if obj.PluginUninstalledMessage != nil {
		return obj.PluginUninstalledMessage
	}

	if obj.RefreshProgressMessage != nil {
		return obj.RefreshProgressMessage
	}

	if obj.RestartRequiredMessage != nil {
		return obj.RestartRequiredMessage
	}

	if obj.ScheduledTaskEndedMessage != nil {
		return obj.ScheduledTaskEndedMessage
	}

	if obj.ScheduledTasksInfoMessage != nil {
		return obj.ScheduledTasksInfoMessage
	}

	if obj.SeriesTimerCancelledMessage != nil {
		return obj.SeriesTimerCancelledMessage
	}

	if obj.SeriesTimerCreatedMessage != nil {
		return obj.SeriesTimerCreatedMessage
	}

	if obj.ServerRestartingMessage != nil {
		return obj.ServerRestartingMessage
	}

	if obj.ServerShuttingDownMessage != nil {
		return obj.ServerShuttingDownMessage
	}

	if obj.SessionsMessage != nil {
		return obj.SessionsMessage
	}

	if obj.SyncPlayCommandMessage != nil {
		return obj.SyncPlayCommandMessage
	}

	if obj.SyncPlayGroupUpdateCommandMessage != nil {
		return obj.SyncPlayGroupUpdateCommandMessage
	}

	if obj.TimerCancelledMessage != nil {
		return obj.TimerCancelledMessage
	}

	if obj.TimerCreatedMessage != nil {
		return obj.TimerCreatedMessage
	}

	if obj.UserDataChangedMessage != nil {
		return obj.UserDataChangedMessage
	}

	if obj.UserDeletedMessage != nil {
		return obj.UserDeletedMessage
	}

	if obj.UserUpdatedMessage != nil {
		return obj.UserUpdatedMessage
	}

	// all schemas are nil
	return nil
}

type NullableOutboundWebSocketMessage struct {
	value *OutboundWebSocketMessage
	isSet bool
}

func (v NullableOutboundWebSocketMessage) Get() *OutboundWebSocketMessage {
	return v.value
}

func (v *NullableOutboundWebSocketMessage) Set(val *OutboundWebSocketMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableOutboundWebSocketMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableOutboundWebSocketMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutboundWebSocketMessage(val *OutboundWebSocketMessage) *NullableOutboundWebSocketMessage {
	return &NullableOutboundWebSocketMessage{value: val, isSet: true}
}

func (v NullableOutboundWebSocketMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutboundWebSocketMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


