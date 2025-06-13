# LyricLineCue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | Pointer to **int32** | Gets the character index of the lyric. | [optional] 
**Start** | Pointer to **int64** | Gets the timestamp the lyric is synced to in ticks. | [optional] 
**End** | Pointer to **NullableInt64** | Gets the end timestamp the lyric is synced to in ticks. | [optional] 

## Methods

### NewLyricLineCue

`func NewLyricLineCue() *LyricLineCue`

NewLyricLineCue instantiates a new LyricLineCue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLyricLineCueWithDefaults

`func NewLyricLineCueWithDefaults() *LyricLineCue`

NewLyricLineCueWithDefaults instantiates a new LyricLineCue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *LyricLineCue) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *LyricLineCue) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *LyricLineCue) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *LyricLineCue) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetStart

`func (o *LyricLineCue) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *LyricLineCue) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *LyricLineCue) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *LyricLineCue) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *LyricLineCue) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *LyricLineCue) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *LyricLineCue) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *LyricLineCue) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *LyricLineCue) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *LyricLineCue) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


