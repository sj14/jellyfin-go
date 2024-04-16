# PlayRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayingQueue** | Pointer to **[]string** | Gets or sets the playing queue. | [optional] 
**PlayingItemPosition** | Pointer to **int32** | Gets or sets the position of the playing item in the queue. | [optional] 
**StartPositionTicks** | Pointer to **int64** | Gets or sets the start position ticks. | [optional] 

## Methods

### NewPlayRequestDto

`func NewPlayRequestDto() *PlayRequestDto`

NewPlayRequestDto instantiates a new PlayRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayRequestDtoWithDefaults

`func NewPlayRequestDtoWithDefaults() *PlayRequestDto`

NewPlayRequestDtoWithDefaults instantiates a new PlayRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayingQueue

`func (o *PlayRequestDto) GetPlayingQueue() []string`

GetPlayingQueue returns the PlayingQueue field if non-nil, zero value otherwise.

### GetPlayingQueueOk

`func (o *PlayRequestDto) GetPlayingQueueOk() (*[]string, bool)`

GetPlayingQueueOk returns a tuple with the PlayingQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayingQueue

`func (o *PlayRequestDto) SetPlayingQueue(v []string)`

SetPlayingQueue sets PlayingQueue field to given value.

### HasPlayingQueue

`func (o *PlayRequestDto) HasPlayingQueue() bool`

HasPlayingQueue returns a boolean if a field has been set.

### GetPlayingItemPosition

`func (o *PlayRequestDto) GetPlayingItemPosition() int32`

GetPlayingItemPosition returns the PlayingItemPosition field if non-nil, zero value otherwise.

### GetPlayingItemPositionOk

`func (o *PlayRequestDto) GetPlayingItemPositionOk() (*int32, bool)`

GetPlayingItemPositionOk returns a tuple with the PlayingItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayingItemPosition

`func (o *PlayRequestDto) SetPlayingItemPosition(v int32)`

SetPlayingItemPosition sets PlayingItemPosition field to given value.

### HasPlayingItemPosition

`func (o *PlayRequestDto) HasPlayingItemPosition() bool`

HasPlayingItemPosition returns a boolean if a field has been set.

### GetStartPositionTicks

`func (o *PlayRequestDto) GetStartPositionTicks() int64`

GetStartPositionTicks returns the StartPositionTicks field if non-nil, zero value otherwise.

### GetStartPositionTicksOk

`func (o *PlayRequestDto) GetStartPositionTicksOk() (*int64, bool)`

GetStartPositionTicksOk returns a tuple with the StartPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPositionTicks

`func (o *PlayRequestDto) SetStartPositionTicks(v int64)`

SetStartPositionTicks sets StartPositionTicks field to given value.

### HasStartPositionTicks

`func (o *PlayRequestDto) HasStartPositionTicks() bool`

HasStartPositionTicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


