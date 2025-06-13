# SyncPlayGroupUpdateMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GroupUpdate**](GroupUpdate.md) | Group update data | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**SessionMessageType**](SessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SESSIONMESSAGETYPE_SYNC_PLAY_GROUP_UPDATE]

## Methods

### NewSyncPlayGroupUpdateMessage

`func NewSyncPlayGroupUpdateMessage() *SyncPlayGroupUpdateMessage`

NewSyncPlayGroupUpdateMessage instantiates a new SyncPlayGroupUpdateMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncPlayGroupUpdateMessageWithDefaults

`func NewSyncPlayGroupUpdateMessageWithDefaults() *SyncPlayGroupUpdateMessage`

NewSyncPlayGroupUpdateMessageWithDefaults instantiates a new SyncPlayGroupUpdateMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SyncPlayGroupUpdateMessage) GetData() GroupUpdate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SyncPlayGroupUpdateMessage) GetDataOk() (*GroupUpdate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SyncPlayGroupUpdateMessage) SetData(v GroupUpdate)`

SetData sets Data field to given value.

### HasData

`func (o *SyncPlayGroupUpdateMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessageId

`func (o *SyncPlayGroupUpdateMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SyncPlayGroupUpdateMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SyncPlayGroupUpdateMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *SyncPlayGroupUpdateMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *SyncPlayGroupUpdateMessage) GetMessageType() SessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *SyncPlayGroupUpdateMessage) GetMessageTypeOk() (*SessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *SyncPlayGroupUpdateMessage) SetMessageType(v SessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *SyncPlayGroupUpdateMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


