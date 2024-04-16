# SeekRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PositionTicks** | Pointer to **int64** | Gets or sets the position ticks. | [optional] 

## Methods

### NewSeekRequestDto

`func NewSeekRequestDto() *SeekRequestDto`

NewSeekRequestDto instantiates a new SeekRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeekRequestDtoWithDefaults

`func NewSeekRequestDtoWithDefaults() *SeekRequestDto`

NewSeekRequestDtoWithDefaults instantiates a new SeekRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionTicks

`func (o *SeekRequestDto) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *SeekRequestDto) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *SeekRequestDto) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *SeekRequestDto) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


