# PlaystateMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullablePlaystateMessageData**](PlaystateMessageData.md) |  | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**SessionMessageType**](SessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SESSIONMESSAGETYPE_PLAYSTATE]

## Methods

### NewPlaystateMessage

`func NewPlaystateMessage() *PlaystateMessage`

NewPlaystateMessage instantiates a new PlaystateMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaystateMessageWithDefaults

`func NewPlaystateMessageWithDefaults() *PlaystateMessage`

NewPlaystateMessageWithDefaults instantiates a new PlaystateMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PlaystateMessage) GetData() PlaystateMessageData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PlaystateMessage) GetDataOk() (*PlaystateMessageData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PlaystateMessage) SetData(v PlaystateMessageData)`

SetData sets Data field to given value.

### HasData

`func (o *PlaystateMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *PlaystateMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *PlaystateMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *PlaystateMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *PlaystateMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *PlaystateMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *PlaystateMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *PlaystateMessage) GetMessageType() SessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *PlaystateMessage) GetMessageTypeOk() (*SessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *PlaystateMessage) SetMessageType(v SessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *PlaystateMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


