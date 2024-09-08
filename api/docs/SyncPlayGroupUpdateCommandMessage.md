# SyncPlayGroupUpdateCommandMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableGroupUpdate**](GroupUpdate.md) | Group update without data. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**SessionMessageType**](SessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SESSIONMESSAGETYPE_SYNC_PLAY_GROUP_UPDATE]

## Methods

### NewSyncPlayGroupUpdateCommandMessage

`func NewSyncPlayGroupUpdateCommandMessage() *SyncPlayGroupUpdateCommandMessage`

NewSyncPlayGroupUpdateCommandMessage instantiates a new SyncPlayGroupUpdateCommandMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncPlayGroupUpdateCommandMessageWithDefaults

`func NewSyncPlayGroupUpdateCommandMessageWithDefaults() *SyncPlayGroupUpdateCommandMessage`

NewSyncPlayGroupUpdateCommandMessageWithDefaults instantiates a new SyncPlayGroupUpdateCommandMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SyncPlayGroupUpdateCommandMessage) GetData() GroupUpdate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SyncPlayGroupUpdateCommandMessage) GetDataOk() (*GroupUpdate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SyncPlayGroupUpdateCommandMessage) SetData(v GroupUpdate)`

SetData sets Data field to given value.

### HasData

`func (o *SyncPlayGroupUpdateCommandMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *SyncPlayGroupUpdateCommandMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *SyncPlayGroupUpdateCommandMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *SyncPlayGroupUpdateCommandMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SyncPlayGroupUpdateCommandMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SyncPlayGroupUpdateCommandMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *SyncPlayGroupUpdateCommandMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *SyncPlayGroupUpdateCommandMessage) GetMessageType() SessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *SyncPlayGroupUpdateCommandMessage) GetMessageTypeOk() (*SessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *SyncPlayGroupUpdateCommandMessage) SetMessageType(v SessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *SyncPlayGroupUpdateCommandMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


