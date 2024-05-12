# GeneralCommandMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableGeneralCommandMessageData**](GeneralCommandMessageData.md) |  | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**SessionMessageType**](SessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SESSIONMESSAGETYPE_GENERAL_COMMAND]

## Methods

### NewGeneralCommandMessage

`func NewGeneralCommandMessage() *GeneralCommandMessage`

NewGeneralCommandMessage instantiates a new GeneralCommandMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralCommandMessageWithDefaults

`func NewGeneralCommandMessageWithDefaults() *GeneralCommandMessage`

NewGeneralCommandMessageWithDefaults instantiates a new GeneralCommandMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GeneralCommandMessage) GetData() GeneralCommandMessageData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GeneralCommandMessage) GetDataOk() (*GeneralCommandMessageData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GeneralCommandMessage) SetData(v GeneralCommandMessageData)`

SetData sets Data field to given value.

### HasData

`func (o *GeneralCommandMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *GeneralCommandMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *GeneralCommandMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *GeneralCommandMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *GeneralCommandMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *GeneralCommandMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *GeneralCommandMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *GeneralCommandMessage) GetMessageType() SessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *GeneralCommandMessage) GetMessageTypeOk() (*SessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *GeneralCommandMessage) SetMessageType(v SessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *GeneralCommandMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


