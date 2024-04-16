# BufferRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**When** | Pointer to **time.Time** | Gets or sets when the request has been made by the client. | [optional] 
**PositionTicks** | Pointer to **int64** | Gets or sets the position ticks. | [optional] 
**IsPlaying** | Pointer to **bool** | Gets or sets a value indicating whether the client playback is unpaused. | [optional] 
**PlaylistItemId** | Pointer to **string** | Gets or sets the playlist item identifier of the playing item. | [optional] 

## Methods

### NewBufferRequestDto

`func NewBufferRequestDto() *BufferRequestDto`

NewBufferRequestDto instantiates a new BufferRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBufferRequestDtoWithDefaults

`func NewBufferRequestDtoWithDefaults() *BufferRequestDto`

NewBufferRequestDtoWithDefaults instantiates a new BufferRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhen

`func (o *BufferRequestDto) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *BufferRequestDto) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *BufferRequestDto) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *BufferRequestDto) HasWhen() bool`

HasWhen returns a boolean if a field has been set.

### GetPositionTicks

`func (o *BufferRequestDto) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *BufferRequestDto) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *BufferRequestDto) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *BufferRequestDto) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### GetIsPlaying

`func (o *BufferRequestDto) GetIsPlaying() bool`

GetIsPlaying returns the IsPlaying field if non-nil, zero value otherwise.

### GetIsPlayingOk

`func (o *BufferRequestDto) GetIsPlayingOk() (*bool, bool)`

GetIsPlayingOk returns a tuple with the IsPlaying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlaying

`func (o *BufferRequestDto) SetIsPlaying(v bool)`

SetIsPlaying sets IsPlaying field to given value.

### HasIsPlaying

`func (o *BufferRequestDto) HasIsPlaying() bool`

HasIsPlaying returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *BufferRequestDto) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *BufferRequestDto) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *BufferRequestDto) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *BufferRequestDto) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


