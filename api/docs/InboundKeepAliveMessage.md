# InboundKeepAliveMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageType** | Pointer to [**SessionMessageType**](SessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SESSIONMESSAGETYPE_KEEP_ALIVE]

## Methods

### NewInboundKeepAliveMessage

`func NewInboundKeepAliveMessage() *InboundKeepAliveMessage`

NewInboundKeepAliveMessage instantiates a new InboundKeepAliveMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboundKeepAliveMessageWithDefaults

`func NewInboundKeepAliveMessageWithDefaults() *InboundKeepAliveMessage`

NewInboundKeepAliveMessageWithDefaults instantiates a new InboundKeepAliveMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageType

`func (o *InboundKeepAliveMessage) GetMessageType() SessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *InboundKeepAliveMessage) GetMessageTypeOk() (*SessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *InboundKeepAliveMessage) SetMessageType(v SessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *InboundKeepAliveMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


