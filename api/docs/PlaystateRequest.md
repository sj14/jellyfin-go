# PlaystateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to [**PlaystateCommand**](PlaystateCommand.md) | Enum PlaystateCommand. | [optional] 
**SeekPositionTicks** | Pointer to **NullableInt64** |  | [optional] 
**ControllingUserId** | Pointer to **NullableString** | Gets or sets the controlling user identifier. | [optional] 

## Methods

### NewPlaystateRequest

`func NewPlaystateRequest() *PlaystateRequest`

NewPlaystateRequest instantiates a new PlaystateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaystateRequestWithDefaults

`func NewPlaystateRequestWithDefaults() *PlaystateRequest`

NewPlaystateRequestWithDefaults instantiates a new PlaystateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *PlaystateRequest) GetCommand() PlaystateCommand`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *PlaystateRequest) GetCommandOk() (*PlaystateCommand, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *PlaystateRequest) SetCommand(v PlaystateCommand)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *PlaystateRequest) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetSeekPositionTicks

`func (o *PlaystateRequest) GetSeekPositionTicks() int64`

GetSeekPositionTicks returns the SeekPositionTicks field if non-nil, zero value otherwise.

### GetSeekPositionTicksOk

`func (o *PlaystateRequest) GetSeekPositionTicksOk() (*int64, bool)`

GetSeekPositionTicksOk returns a tuple with the SeekPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeekPositionTicks

`func (o *PlaystateRequest) SetSeekPositionTicks(v int64)`

SetSeekPositionTicks sets SeekPositionTicks field to given value.

### HasSeekPositionTicks

`func (o *PlaystateRequest) HasSeekPositionTicks() bool`

HasSeekPositionTicks returns a boolean if a field has been set.

### SetSeekPositionTicksNil

`func (o *PlaystateRequest) SetSeekPositionTicksNil(b bool)`

 SetSeekPositionTicksNil sets the value for SeekPositionTicks to be an explicit nil

### UnsetSeekPositionTicks
`func (o *PlaystateRequest) UnsetSeekPositionTicks()`

UnsetSeekPositionTicks ensures that no value is present for SeekPositionTicks, not even an explicit nil
### GetControllingUserId

`func (o *PlaystateRequest) GetControllingUserId() string`

GetControllingUserId returns the ControllingUserId field if non-nil, zero value otherwise.

### GetControllingUserIdOk

`func (o *PlaystateRequest) GetControllingUserIdOk() (*string, bool)`

GetControllingUserIdOk returns a tuple with the ControllingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllingUserId

`func (o *PlaystateRequest) SetControllingUserId(v string)`

SetControllingUserId sets ControllingUserId field to given value.

### HasControllingUserId

`func (o *PlaystateRequest) HasControllingUserId() bool`

HasControllingUserId returns a boolean if a field has been set.

### SetControllingUserIdNil

`func (o *PlaystateRequest) SetControllingUserIdNil(b bool)`

 SetControllingUserIdNil sets the value for ControllingUserId to be an explicit nil

### UnsetControllingUserId
`func (o *PlaystateRequest) UnsetControllingUserId()`

UnsetControllingUserId ensures that no value is present for ControllingUserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


