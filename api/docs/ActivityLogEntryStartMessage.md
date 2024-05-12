# ActivityLogEntryStartMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **NullableString** | Gets or sets the data. | [optional] 
**MessageType** | Pointer to [**SessionMessageType**](SessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SESSIONMESSAGETYPE_ACTIVITY_LOG_ENTRY_START]

## Methods

### NewActivityLogEntryStartMessage

`func NewActivityLogEntryStartMessage() *ActivityLogEntryStartMessage`

NewActivityLogEntryStartMessage instantiates a new ActivityLogEntryStartMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogEntryStartMessageWithDefaults

`func NewActivityLogEntryStartMessageWithDefaults() *ActivityLogEntryStartMessage`

NewActivityLogEntryStartMessageWithDefaults instantiates a new ActivityLogEntryStartMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ActivityLogEntryStartMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ActivityLogEntryStartMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ActivityLogEntryStartMessage) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *ActivityLogEntryStartMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ActivityLogEntryStartMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ActivityLogEntryStartMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageType

`func (o *ActivityLogEntryStartMessage) GetMessageType() SessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ActivityLogEntryStartMessage) GetMessageTypeOk() (*SessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ActivityLogEntryStartMessage) SetMessageType(v SessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ActivityLogEntryStartMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


