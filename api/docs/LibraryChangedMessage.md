# LibraryChangedMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**NullableLibraryUpdateInfo**](LibraryUpdateInfo.md) | Class LibraryUpdateInfo. | [optional] 
**MessageId** | Pointer to **string** | Gets or sets the message id. | [optional] 
**MessageType** | Pointer to [**SessionMessageType**](SessionMessageType.md) | The different kinds of messages that are used in the WebSocket api. | [optional] [readonly] [default to SESSIONMESSAGETYPE_LIBRARY_CHANGED]

## Methods

### NewLibraryChangedMessage

`func NewLibraryChangedMessage() *LibraryChangedMessage`

NewLibraryChangedMessage instantiates a new LibraryChangedMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryChangedMessageWithDefaults

`func NewLibraryChangedMessageWithDefaults() *LibraryChangedMessage`

NewLibraryChangedMessageWithDefaults instantiates a new LibraryChangedMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LibraryChangedMessage) GetData() LibraryUpdateInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LibraryChangedMessage) GetDataOk() (*LibraryUpdateInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LibraryChangedMessage) SetData(v LibraryUpdateInfo)`

SetData sets Data field to given value.

### HasData

`func (o *LibraryChangedMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *LibraryChangedMessage) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *LibraryChangedMessage) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetMessageId

`func (o *LibraryChangedMessage) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *LibraryChangedMessage) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *LibraryChangedMessage) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *LibraryChangedMessage) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetMessageType

`func (o *LibraryChangedMessage) GetMessageType() SessionMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *LibraryChangedMessage) GetMessageTypeOk() (*SessionMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *LibraryChangedMessage) SetMessageType(v SessionMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *LibraryChangedMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


