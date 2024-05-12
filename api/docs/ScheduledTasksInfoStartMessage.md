# ScheduledTasksInfoStartMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **NullableString** | Gets or sets the data. | [optional] 
**MessageType** | Pointer to [**SessionMessageType**](SessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SESSIONMESSAGETYPE_SCHEDULED_TASKS_INFO_START]

## Methods

### NewScheduledTasksInfoStartMessage

`func NewScheduledTasksInfoStartMessage() *ScheduledTasksInfoStartMessage`

NewScheduledTasksInfoStartMessage instantiates a new ScheduledTasksInfoStartMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledTasksInfoStartMessageWithDefaults

`func NewScheduledTasksInfoStartMessageWithDefaults() *ScheduledTasksInfoStartMessage`

NewScheduledTasksInfoStartMessageWithDefaults instantiates a new ScheduledTasksInfoStartMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ScheduledTasksInfoStartMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ScheduledTasksInfoStartMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ScheduledTasksInfoStartMessage) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ScheduledTasksInfoStartMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ScheduledTasksInfoStartMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ScheduledTasksInfoStartMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageType

`func (o *ScheduledTasksInfoStartMessage) GetMessageType() SessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ScheduledTasksInfoStartMessage) GetMessageTypeOk() (*SessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ScheduledTasksInfoStartMessage) SetMessageType(v SessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ScheduledTasksInfoStartMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


