# PlaystateMessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to [**PlaystateCommand**](PlaystateCommand.md) | Enum PlaystateCommand. | [optional] 
**SeekPositionTicks** | Pointer to **NullableInt64** |  | [optional] 
**ControllingUserId** | Pointer to **NullableString** | Gets or sets the controlling user identifier. | [optional] 

## Methods

### NewPlaystateMessageData

`func NewPlaystateMessageData() *PlaystateMessageData`

NewPlaystateMessageData instantiates a new PlaystateMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaystateMessageDataWithDefaults

`func NewPlaystateMessageDataWithDefaults() *PlaystateMessageData`

NewPlaystateMessageDataWithDefaults instantiates a new PlaystateMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *PlaystateMessageData) GetCommand() PlaystateCommand`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *PlaystateMessageData) GetCommandOk() (*PlaystateCommand, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *PlaystateMessageData) SetCommand(v PlaystateCommand)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *PlaystateMessageData) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetSeekPositionTicks

`func (o *PlaystateMessageData) GetSeekPositionTicks() int64`

GetSeekPositionTicks returns the SeekPositionTicks field if non-nil, zero value otherwise.

### GetSeekPositionTicksOk

`func (o *PlaystateMessageData) GetSeekPositionTicksOk() (*int64, bool)`

GetSeekPositionTicksOk returns a tuple with the SeekPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeekPositionTicks

`func (o *PlaystateMessageData) SetSeekPositionTicks(v int64)`

SetSeekPositionTicks sets SeekPositionTicks field to given value.

### HasSeekPositionTicks

`func (o *PlaystateMessageData) HasSeekPositionTicks() bool`

HasSeekPositionTicks returns a boolean if a field has been set.

### SetSeekPositionTicksNil

`func (o *PlaystateMessageData) SetSeekPositionTicksNil(b bool)`

 SetSeekPositionTicksNil sets the value for SeekPositionTicks to be an explicit nil

### UnsetSeekPositionTicks
`func (o *PlaystateMessageData) UnsetSeekPositionTicks()`

UnsetSeekPositionTicks ensures that no value is present for SeekPositionTicks, not even an explicit nil
### GetControllingUserId

`func (o *PlaystateMessageData) GetControllingUserId() string`

GetControllingUserId returns the ControllingUserId field if non-nil, zero value otherwise.

### GetControllingUserIdOk

`func (o *PlaystateMessageData) GetControllingUserIdOk() (*string, bool)`

GetControllingUserIdOk returns a tuple with the ControllingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllingUserId

`func (o *PlaystateMessageData) SetControllingUserId(v string)`

SetControllingUserId sets ControllingUserId field to given value.

### HasControllingUserId

`func (o *PlaystateMessageData) HasControllingUserId() bool`

HasControllingUserId returns a boolean if a field has been set.

### SetControllingUserIdNil

`func (o *PlaystateMessageData) SetControllingUserIdNil(b bool)`

 SetControllingUserIdNil sets the value for ControllingUserId to be an explicit nil

### UnsetControllingUserId
`func (o *PlaystateMessageData) UnsetControllingUserId()`

UnsetControllingUserId ensures that no value is present for ControllingUserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


