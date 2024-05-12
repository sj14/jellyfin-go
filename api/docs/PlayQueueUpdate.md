# PlayQueueUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to [**PlayQueueUpdateReason**](PlayQueueUpdateReason.md) | Gets the request type that originated this update. | [optional] 
**LastUpdate** | Pointer to **time.Time** | Gets the UTC time of the last change to the playing queue. | [optional] 
**Playlist** | Pointer to [**[]SyncPlayQueueItem**](SyncPlayQueueItem.md) | Gets the playlist. | [optional] 
**PlayingItemIndex** | Pointer to **int32** | Gets the playing item index in the playlist. | [optional] 
**StartPositionTicks** | Pointer to **int64** | Gets the start position ticks. | [optional] 
**IsPlaying** | Pointer to **bool** | Gets a value indicating whether the current item is playing. | [optional] 
**ShuffleMode** | Pointer to [**GroupShuffleMode**](GroupShuffleMode.md) | Gets the shuffle mode. | [optional] 
**RepeatMode** | Pointer to [**GroupRepeatMode**](GroupRepeatMode.md) | Gets the repeat mode. | [optional] 

## Methods

### NewPlayQueueUpdate

`func NewPlayQueueUpdate() *PlayQueueUpdate`

NewPlayQueueUpdate instantiates a new PlayQueueUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayQueueUpdateWithDefaults

`func NewPlayQueueUpdateWithDefaults() *PlayQueueUpdate`

NewPlayQueueUpdateWithDefaults instantiates a new PlayQueueUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *PlayQueueUpdate) GetReason() PlayQueueUpdateReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PlayQueueUpdate) GetReasonOk() (*PlayQueueUpdateReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PlayQueueUpdate) SetReason(v PlayQueueUpdateReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PlayQueueUpdate) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetLastUpdate

`func (o *PlayQueueUpdate) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *PlayQueueUpdate) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *PlayQueueUpdate) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *PlayQueueUpdate) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetPlaylist

`func (o *PlayQueueUpdate) GetPlaylist() []SyncPlayQueueItem`

GetPlaylist returns the Playlist field if non-nil, zero value otherwise.

### GetPlaylistOk

`func (o *PlayQueueUpdate) GetPlaylistOk() (*[]SyncPlayQueueItem, bool)`

GetPlaylistOk returns a tuple with the Playlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylist

`func (o *PlayQueueUpdate) SetPlaylist(v []SyncPlayQueueItem)`

SetPlaylist sets Playlist field to given value.

### HasPlaylist

`func (o *PlayQueueUpdate) HasPlaylist() bool`

HasPlaylist returns a boolean if a field has been set.

### GetPlayingItemIndex

`func (o *PlayQueueUpdate) GetPlayingItemIndex() int32`

GetPlayingItemIndex returns the PlayingItemIndex field if non-nil, zero value otherwise.

### GetPlayingItemIndexOk

`func (o *PlayQueueUpdate) GetPlayingItemIndexOk() (*int32, bool)`

GetPlayingItemIndexOk returns a tuple with the PlayingItemIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayingItemIndex

`func (o *PlayQueueUpdate) SetPlayingItemIndex(v int32)`

SetPlayingItemIndex sets PlayingItemIndex field to given value.

### HasPlayingItemIndex

`func (o *PlayQueueUpdate) HasPlayingItemIndex() bool`

HasPlayingItemIndex returns a boolean if a field has been set.

### GetStartPositionTicks

`func (o *PlayQueueUpdate) GetStartPositionTicks() int64`

GetStartPositionTicks returns the StartPositionTicks field if non-nil, zero value otherwise.

### GetStartPositionTicksOk

`func (o *PlayQueueUpdate) GetStartPositionTicksOk() (*int64, bool)`

GetStartPositionTicksOk returns a tuple with the StartPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPositionTicks

`func (o *PlayQueueUpdate) SetStartPositionTicks(v int64)`

SetStartPositionTicks sets StartPositionTicks field to given value.

### HasStartPositionTicks

`func (o *PlayQueueUpdate) HasStartPositionTicks() bool`

HasStartPositionTicks returns a boolean if a field has been set.

### GetIsPlaying

`func (o *PlayQueueUpdate) GetIsPlaying() bool`

GetIsPlaying returns the IsPlaying field if non-nil, zero value otherwise.

### GetIsPlayingOk

`func (o *PlayQueueUpdate) GetIsPlayingOk() (*bool, bool)`

GetIsPlayingOk returns a tuple with the IsPlaying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlaying

`func (o *PlayQueueUpdate) SetIsPlaying(v bool)`

SetIsPlaying sets IsPlaying field to given value.

### HasIsPlaying

`func (o *PlayQueueUpdate) HasIsPlaying() bool`

HasIsPlaying returns a boolean if a field has been set.

### GetShuffleMode

`func (o *PlayQueueUpdate) GetShuffleMode() GroupShuffleMode`

GetShuffleMode returns the ShuffleMode field if non-nil, zero value otherwise.

### GetShuffleModeOk

`func (o *PlayQueueUpdate) GetShuffleModeOk() (*GroupShuffleMode, bool)`

GetShuffleModeOk returns a tuple with the ShuffleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShuffleMode

`func (o *PlayQueueUpdate) SetShuffleMode(v GroupShuffleMode)`

SetShuffleMode sets ShuffleMode field to given value.

### HasShuffleMode

`func (o *PlayQueueUpdate) HasShuffleMode() bool`

HasShuffleMode returns a boolean if a field has been set.

### GetRepeatMode

`func (o *PlayQueueUpdate) GetRepeatMode() GroupRepeatMode`

GetRepeatMode returns the RepeatMode field if non-nil, zero value otherwise.

### GetRepeatModeOk

`func (o *PlayQueueUpdate) GetRepeatModeOk() (*GroupRepeatMode, bool)`

GetRepeatModeOk returns a tuple with the RepeatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatMode

`func (o *PlayQueueUpdate) SetRepeatMode(v GroupRepeatMode)`

SetRepeatMode sets RepeatMode field to given value.

### HasRepeatMode

`func (o *PlayQueueUpdate) HasRepeatMode() bool`

HasRepeatMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


