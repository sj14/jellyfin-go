# OutboundWebSocketMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GroupUpdate**](GroupUpdate.md) | Group update data | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**SessionMessageType**](SessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SESSIONMESSAGETYPE_SYNC_PLAY_GROUP_UPDATE]

## Methods

### NewOutboundWebSocketMessage

`func NewOutboundWebSocketMessage() *OutboundWebSocketMessage`

NewOutboundWebSocketMessage instantiates a new OutboundWebSocketMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundWebSocketMessageWithDefaults

`func NewOutboundWebSocketMessageWithDefaults() *OutboundWebSocketMessage`

NewOutboundWebSocketMessageWithDefaults instantiates a new OutboundWebSocketMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OutboundWebSocketMessage) GetData() GroupUpdate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OutboundWebSocketMessage) GetDataOk() (*GroupUpdate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OutboundWebSocketMessage) SetData(v GroupUpdate)`

SetData sets Data field to given value.

### HasData

`func (o *OutboundWebSocketMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessageId

`func (o *OutboundWebSocketMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *OutboundWebSocketMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *OutboundWebSocketMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *OutboundWebSocketMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *OutboundWebSocketMessage) GetMessageType() SessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *OutboundWebSocketMessage) GetMessageTypeOk() (*SessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *OutboundWebSocketMessage) SetMessageType(v SessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *OutboundWebSocketMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


