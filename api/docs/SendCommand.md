# SendCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] 
**PlaylistItemId** | Pointer to **string** | Gets the playlist identifier of the playing item. | [optional] 
**When** | Pointer to **time.Time** | Gets or sets the UTC time when to execute the command. | [optional] 
**PositionTicks** | Pointer to **NullableInt64** | Gets the position ticks. | [optional] 
**Command** | Pointer to [**SendCommandType**](SendCommandType.md) | Gets the command. | [optional] 
**EmittedAt** | Pointer to **time.Time** | Gets the UTC time when this command has been emitted. | [optional] 

## Methods

### NewSendCommand

`func NewSendCommand() *SendCommand`

NewSendCommand instantiates a new SendCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendCommandWithDefaults

`func NewSendCommandWithDefaults() *SendCommand`

NewSendCommandWithDefaults instantiates a new SendCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SendCommand) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SendCommand) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SendCommand) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SendCommand) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *SendCommand) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *SendCommand) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *SendCommand) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *SendCommand) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### GetWhen

`func (o *SendCommand) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *SendCommand) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *SendCommand) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *SendCommand) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetPositionTicks

`func (o *SendCommand) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *SendCommand) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *SendCommand) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *SendCommand) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *SendCommand) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *SendCommand) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetCommand

`func (o *SendCommand) GetCommand() SendCommandType`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *SendCommand) GetCommandOk() (*SendCommandType, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *SendCommand) SetCommand(v SendCommandType)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *SendCommand) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetEmittedAt

`func (o *SendCommand) GetEmittedAt() time.Time`

GetEmittedAt returns the EmittedAt field if non-nil, zero value otherwise.

### GetEmittedAtOk

`func (o *SendCommand) GetEmittedAtOk() (*time.Time, bool)`

GetEmittedAtOk returns a tuple with the EmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmittedAt

`func (o *SendCommand) SetEmittedAt(v time.Time)`

SetEmittedAt sets EmittedAt field to given value.

### HasEmittedAt

`func (o *SendCommand) HasEmittedAt() bool`

HasEmittedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


