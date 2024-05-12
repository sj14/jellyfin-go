# ScheduledTasksInfoStopMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageType** | Pointer to [**SessionMessageType**](SessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SESSIONMESSAGETYPE_SCHEDULED_TASKS_INFO_STOP]

## Methods

### NewScheduledTasksInfoStopMessage

`func NewScheduledTasksInfoStopMessage() *ScheduledTasksInfoStopMessage`

NewScheduledTasksInfoStopMessage instantiates a new ScheduledTasksInfoStopMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledTasksInfoStopMessageWithDefaults

`func NewScheduledTasksInfoStopMessageWithDefaults() *ScheduledTasksInfoStopMessage`

NewScheduledTasksInfoStopMessageWithDefaults instantiates a new ScheduledTasksInfoStopMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageType

`func (o *ScheduledTasksInfoStopMessage) GetMessageType() SessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ScheduledTasksInfoStopMessage) GetMessageTypeOk() (*SessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ScheduledTasksInfoStopMessage) SetMessageType(v SessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ScheduledTasksInfoStopMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


