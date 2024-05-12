# InboundWebSocketMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **NullableString** | Gets or sets the data. | [optional] 
**MessageType** | Pointer to [**SessionMessageType**](SessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SESSIONMESSAGETYPE_SESSIONS_STOP]

## Methods

### NewInboundWebSocketMessage

`func NewInboundWebSocketMessage() *InboundWebSocketMessage`

NewInboundWebSocketMessage instantiates a new InboundWebSocketMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundWebSocketMessageWithDefaults

`func NewInboundWebSocketMessageWithDefaults() *InboundWebSocketMessage`

NewInboundWebSocketMessageWithDefaults instantiates a new InboundWebSocketMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *InboundWebSocketMessage) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InboundWebSocketMessage) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InboundWebSocketMessage) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *InboundWebSocketMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *InboundWebSocketMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *InboundWebSocketMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageType

`func (o *InboundWebSocketMessage) GetMessageType() SessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *InboundWebSocketMessage) GetMessageTypeOk() (*SessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *InboundWebSocketMessage) SetMessageType(v SessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *InboundWebSocketMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


