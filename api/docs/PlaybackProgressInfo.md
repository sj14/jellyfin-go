# PlaybackProgressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanSeek** | Pointer to **bool** | Gets or sets a value indicating whether this instance can seek. | [optional] 
**Item** | Pointer to [**NullablePlaybackProgressInfoItem**](PlaybackProgressInfoItem.md) |  | [optional] 
**ItemId** | Pointer to **string** | Gets or sets the item identifier. | [optional] 
**SessionId** | Pointer to **NullableString** | Gets or sets the session id. | [optional] 
**MediaSourceId** | Pointer to **NullableString** | Gets or sets the media version identifier. | [optional] 
**AudioStreamIndex** | Pointer to **NullableInt32** | Gets or sets the index of the audio stream. | [optional] 
**SubtitleStreamIndex** | Pointer to **NullableInt32** | Gets or sets the index of the subtitle stream. | [optional] 
**IsPaused** | Pointer to **bool** | Gets or sets a value indicating whether this instance is paused. | [optional] 
**IsMuted** | Pointer to **bool** | Gets or sets a value indicating whether this instance is muted. | [optional] 
**PositionTicks** | Pointer to **NullableInt64** | Gets or sets the position ticks. | [optional] 
**PlaybackStartTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**VolumeLevel** | Pointer to **NullableInt32** | Gets or sets the volume level. | [optional] 
**Brightness** | Pointer to **NullableInt32** |  | [optional] 
**AspectRatio** | Pointer to **NullableString** |  | [optional] 
**PlayMethod** | Pointer to [**PlayMethod**](PlayMethod.md) | Gets or sets the play method. | [optional] 
**LiveStreamId** | Pointer to **NullableString** | Gets or sets the live stream identifier. | [optional] 
**PlaySessionId** | Pointer to **NullableString** | Gets or sets the play session identifier. | [optional] 
**RepeatMode** | Pointer to [**RepeatMode**](RepeatMode.md) | Gets or sets the repeat mode. | [optional] 
**NowPlayingQueue** | Pointer to [**[]QueueItem**](QueueItem.md) |  | [optional] 
**PlaylistItemId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPlaybackProgressInfo

`func NewPlaybackProgressInfo() *PlaybackProgressInfo`

NewPlaybackProgressInfo instantiates a new PlaybackProgressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaybackProgressInfoWithDefaults

`func NewPlaybackProgressInfoWithDefaults() *PlaybackProgressInfo`

NewPlaybackProgressInfoWithDefaults instantiates a new PlaybackProgressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanSeek

`func (o *PlaybackProgressInfo) GetCanSeek() bool`

GetCanSeek returns the CanSeek field if non-nil, zero value otherwise.

### GetCanSeekOk

`func (o *PlaybackProgressInfo) GetCanSeekOk() (*bool, bool)`

GetCanSeekOk returns a tuple with the CanSeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeek

`func (o *PlaybackProgressInfo) SetCanSeek(v bool)`

SetCanSeek sets CanSeek field to given value.

### HasCanSeek

`func (o *PlaybackProgressInfo) HasCanSeek() bool`

HasCanSeek returns a boolean if a field has been set.

### GetItem

`func (o *PlaybackProgressInfo) GetItem() PlaybackProgressInfoItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *PlaybackProgressInfo) GetItemOk() (*PlaybackProgressInfoItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *PlaybackProgressInfo) SetItem(v PlaybackProgressInfoItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *PlaybackProgressInfo) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItemNil

`func (o *PlaybackProgressInfo) SetItemNil(b bool)`

 SetItemNil sets the value for Item to be an explicit nil

### UnsetItem
`func (o *PlaybackProgressInfo) UnsetItem()`

UnsetItem ensures that no value is present for Item, not even an explicit nil
### GetItemId

`func (o *PlaybackProgressInfo) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *PlaybackProgressInfo) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *PlaybackProgressInfo) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *PlaybackProgressInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSessionId

`func (o *PlaybackProgressInfo) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *PlaybackProgressInfo) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *PlaybackProgressInfo) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *PlaybackProgressInfo) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *PlaybackProgressInfo) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *PlaybackProgressInfo) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetMediaSourceId

`func (o *PlaybackProgressInfo) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *PlaybackProgressInfo) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *PlaybackProgressInfo) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *PlaybackProgressInfo) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### SetMediaSourceIdNil

`func (o *PlaybackProgressInfo) SetMediaSourceIdNil(b bool)`

 SetMediaSourceIdNil sets the value for MediaSourceId to be an explicit nil

### UnsetMediaSourceId
`func (o *PlaybackProgressInfo) UnsetMediaSourceId()`

UnsetMediaSourceId ensures that no value is present for MediaSourceId, not even an explicit nil
### GetAudioStreamIndex

`func (o *PlaybackProgressInfo) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *PlaybackProgressInfo) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *PlaybackProgressInfo) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *PlaybackProgressInfo) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *PlaybackProgressInfo) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *PlaybackProgressInfo) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *PlaybackProgressInfo) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *PlaybackProgressInfo) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *PlaybackProgressInfo) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *PlaybackProgressInfo) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *PlaybackProgressInfo) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *PlaybackProgressInfo) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetIsPaused

`func (o *PlaybackProgressInfo) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *PlaybackProgressInfo) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *PlaybackProgressInfo) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *PlaybackProgressInfo) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### GetIsMuted

`func (o *PlaybackProgressInfo) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *PlaybackProgressInfo) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *PlaybackProgressInfo) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *PlaybackProgressInfo) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetPositionTicks

`func (o *PlaybackProgressInfo) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *PlaybackProgressInfo) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *PlaybackProgressInfo) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *PlaybackProgressInfo) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *PlaybackProgressInfo) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *PlaybackProgressInfo) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetPlaybackStartTimeTicks

`func (o *PlaybackProgressInfo) GetPlaybackStartTimeTicks() int64`

GetPlaybackStartTimeTicks returns the PlaybackStartTimeTicks field if non-nil, zero value otherwise.

### GetPlaybackStartTimeTicksOk

`func (o *PlaybackProgressInfo) GetPlaybackStartTimeTicksOk() (*int64, bool)`

GetPlaybackStartTimeTicksOk returns a tuple with the PlaybackStartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackStartTimeTicks

`func (o *PlaybackProgressInfo) SetPlaybackStartTimeTicks(v int64)`

SetPlaybackStartTimeTicks sets PlaybackStartTimeTicks field to given value.

### HasPlaybackStartTimeTicks

`func (o *PlaybackProgressInfo) HasPlaybackStartTimeTicks() bool`

HasPlaybackStartTimeTicks returns a boolean if a field has been set.

### SetPlaybackStartTimeTicksNil

`func (o *PlaybackProgressInfo) SetPlaybackStartTimeTicksNil(b bool)`

 SetPlaybackStartTimeTicksNil sets the value for PlaybackStartTimeTicks to be an explicit nil

### UnsetPlaybackStartTimeTicks
`func (o *PlaybackProgressInfo) UnsetPlaybackStartTimeTicks()`

UnsetPlaybackStartTimeTicks ensures that no value is present for PlaybackStartTimeTicks, not even an explicit nil
### GetVolumeLevel

`func (o *PlaybackProgressInfo) GetVolumeLevel() int32`

GetVolumeLevel returns the VolumeLevel field if non-nil, zero value otherwise.

### GetVolumeLevelOk

`func (o *PlaybackProgressInfo) GetVolumeLevelOk() (*int32, bool)`

GetVolumeLevelOk returns a tuple with the VolumeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLevel

`func (o *PlaybackProgressInfo) SetVolumeLevel(v int32)`

SetVolumeLevel sets VolumeLevel field to given value.

### HasVolumeLevel

`func (o *PlaybackProgressInfo) HasVolumeLevel() bool`

HasVolumeLevel returns a boolean if a field has been set.

### SetVolumeLevelNil

`func (o *PlaybackProgressInfo) SetVolumeLevelNil(b bool)`

 SetVolumeLevelNil sets the value for VolumeLevel to be an explicit nil

### UnsetVolumeLevel
`func (o *PlaybackProgressInfo) UnsetVolumeLevel()`

UnsetVolumeLevel ensures that no value is present for VolumeLevel, not even an explicit nil
### GetBrightness

`func (o *PlaybackProgressInfo) GetBrightness() int32`

GetBrightness returns the Brightness field if non-nil, zero value otherwise.

### GetBrightnessOk

`func (o *PlaybackProgressInfo) GetBrightnessOk() (*int32, bool)`

GetBrightnessOk returns a tuple with the Brightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrightness

`func (o *PlaybackProgressInfo) SetBrightness(v int32)`

SetBrightness sets Brightness field to given value.

### HasBrightness

`func (o *PlaybackProgressInfo) HasBrightness() bool`

HasBrightness returns a boolean if a field has been set.

### SetBrightnessNil

`func (o *PlaybackProgressInfo) SetBrightnessNil(b bool)`

 SetBrightnessNil sets the value for Brightness to be an explicit nil

### UnsetBrightness
`func (o *PlaybackProgressInfo) UnsetBrightness()`

UnsetBrightness ensures that no value is present for Brightness, not even an explicit nil
### GetAspectRatio

`func (o *PlaybackProgressInfo) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *PlaybackProgressInfo) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *PlaybackProgressInfo) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *PlaybackProgressInfo) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### SetAspectRatioNil

`func (o *PlaybackProgressInfo) SetAspectRatioNil(b bool)`

 SetAspectRatioNil sets the value for AspectRatio to be an explicit nil

### UnsetAspectRatio
`func (o *PlaybackProgressInfo) UnsetAspectRatio()`

UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
### GetPlayMethod

`func (o *PlaybackProgressInfo) GetPlayMethod() PlayMethod`

GetPlayMethod returns the PlayMethod field if non-nil, zero value otherwise.

### GetPlayMethodOk

`func (o *PlaybackProgressInfo) GetPlayMethodOk() (*PlayMethod, bool)`

GetPlayMethodOk returns a tuple with the PlayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayMethod

`func (o *PlaybackProgressInfo) SetPlayMethod(v PlayMethod)`

SetPlayMethod sets PlayMethod field to given value.

### HasPlayMethod

`func (o *PlaybackProgressInfo) HasPlayMethod() bool`

HasPlayMethod returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *PlaybackProgressInfo) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *PlaybackProgressInfo) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *PlaybackProgressInfo) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *PlaybackProgressInfo) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### SetLiveStreamIdNil

`func (o *PlaybackProgressInfo) SetLiveStreamIdNil(b bool)`

 SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil

### UnsetLiveStreamId
`func (o *PlaybackProgressInfo) UnsetLiveStreamId()`

UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil
### GetPlaySessionId

`func (o *PlaybackProgressInfo) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *PlaybackProgressInfo) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *PlaybackProgressInfo) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *PlaybackProgressInfo) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### SetPlaySessionIdNil

`func (o *PlaybackProgressInfo) SetPlaySessionIdNil(b bool)`

 SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil

### UnsetPlaySessionId
`func (o *PlaybackProgressInfo) UnsetPlaySessionId()`

UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
### GetRepeatMode

`func (o *PlaybackProgressInfo) GetRepeatMode() RepeatMode`

GetRepeatMode returns the RepeatMode field if non-nil, zero value otherwise.

### GetRepeatModeOk

`func (o *PlaybackProgressInfo) GetRepeatModeOk() (*RepeatMode, bool)`

GetRepeatModeOk returns a tuple with the RepeatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatMode

`func (o *PlaybackProgressInfo) SetRepeatMode(v RepeatMode)`

SetRepeatMode sets RepeatMode field to given value.

### HasRepeatMode

`func (o *PlaybackProgressInfo) HasRepeatMode() bool`

HasRepeatMode returns a boolean if a field has been set.

### GetNowPlayingQueue

`func (o *PlaybackProgressInfo) GetNowPlayingQueue() []QueueItem`

GetNowPlayingQueue returns the NowPlayingQueue field if non-nil, zero value otherwise.

### GetNowPlayingQueueOk

`func (o *PlaybackProgressInfo) GetNowPlayingQueueOk() (*[]QueueItem, bool)`

GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueue

`func (o *PlaybackProgressInfo) SetNowPlayingQueue(v []QueueItem)`

SetNowPlayingQueue sets NowPlayingQueue field to given value.

### HasNowPlayingQueue

`func (o *PlaybackProgressInfo) HasNowPlayingQueue() bool`

HasNowPlayingQueue returns a boolean if a field has been set.

### SetNowPlayingQueueNil

`func (o *PlaybackProgressInfo) SetNowPlayingQueueNil(b bool)`

 SetNowPlayingQueueNil sets the value for NowPlayingQueue to be an explicit nil

### UnsetNowPlayingQueue
`func (o *PlaybackProgressInfo) UnsetNowPlayingQueue()`

UnsetNowPlayingQueue ensures that no value is present for NowPlayingQueue, not even an explicit nil
### GetPlaylistItemId

`func (o *PlaybackProgressInfo) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *PlaybackProgressInfo) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *PlaybackProgressInfo) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *PlaybackProgressInfo) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### SetPlaylistItemIdNil

`func (o *PlaybackProgressInfo) SetPlaylistItemIdNil(b bool)`

 SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil

### UnsetPlaylistItemId
`func (o *PlaybackProgressInfo) UnsetPlaylistItemId()`

UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


