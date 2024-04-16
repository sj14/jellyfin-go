# PlaybackStartInfo

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

### NewPlaybackStartInfo

`func NewPlaybackStartInfo() *PlaybackStartInfo`

NewPlaybackStartInfo instantiates a new PlaybackStartInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaybackStartInfoWithDefaults

`func NewPlaybackStartInfoWithDefaults() *PlaybackStartInfo`

NewPlaybackStartInfoWithDefaults instantiates a new PlaybackStartInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanSeek

`func (o *PlaybackStartInfo) GetCanSeek() bool`

GetCanSeek returns the CanSeek field if non-nil, zero value otherwise.

### GetCanSeekOk

`func (o *PlaybackStartInfo) GetCanSeekOk() (*bool, bool)`

GetCanSeekOk returns a tuple with the CanSeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSeek

`func (o *PlaybackStartInfo) SetCanSeek(v bool)`

SetCanSeek sets CanSeek field to given value.

### HasCanSeek

`func (o *PlaybackStartInfo) HasCanSeek() bool`

HasCanSeek returns a boolean if a field has been set.

### GetItem

`func (o *PlaybackStartInfo) GetItem() PlaybackProgressInfoItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *PlaybackStartInfo) GetItemOk() (*PlaybackProgressInfoItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *PlaybackStartInfo) SetItem(v PlaybackProgressInfoItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *PlaybackStartInfo) HasItem() bool`

HasItem returns a boolean if a field has been set.

### SetItemNil

`func (o *PlaybackStartInfo) SetItemNil(b bool)`

 SetItemNil sets the value for Item to be an explicit nil

### UnsetItem
`func (o *PlaybackStartInfo) UnsetItem()`

UnsetItem ensures that no value is present for Item, not even an explicit nil
### GetItemId

`func (o *PlaybackStartInfo) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *PlaybackStartInfo) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *PlaybackStartInfo) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *PlaybackStartInfo) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetSessionId

`func (o *PlaybackStartInfo) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *PlaybackStartInfo) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *PlaybackStartInfo) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *PlaybackStartInfo) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *PlaybackStartInfo) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *PlaybackStartInfo) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
### GetMediaSourceId

`func (o *PlaybackStartInfo) GetMediaSourceId() string`

GetMediaSourceId returns the MediaSourceId field if non-nil, zero value otherwise.

### GetMediaSourceIdOk

`func (o *PlaybackStartInfo) GetMediaSourceIdOk() (*string, bool)`

GetMediaSourceIdOk returns a tuple with the MediaSourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSourceId

`func (o *PlaybackStartInfo) SetMediaSourceId(v string)`

SetMediaSourceId sets MediaSourceId field to given value.

### HasMediaSourceId

`func (o *PlaybackStartInfo) HasMediaSourceId() bool`

HasMediaSourceId returns a boolean if a field has been set.

### SetMediaSourceIdNil

`func (o *PlaybackStartInfo) SetMediaSourceIdNil(b bool)`

 SetMediaSourceIdNil sets the value for MediaSourceId to be an explicit nil

### UnsetMediaSourceId
`func (o *PlaybackStartInfo) UnsetMediaSourceId()`

UnsetMediaSourceId ensures that no value is present for MediaSourceId, not even an explicit nil
### GetAudioStreamIndex

`func (o *PlaybackStartInfo) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *PlaybackStartInfo) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *PlaybackStartInfo) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *PlaybackStartInfo) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *PlaybackStartInfo) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *PlaybackStartInfo) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *PlaybackStartInfo) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *PlaybackStartInfo) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *PlaybackStartInfo) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *PlaybackStartInfo) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *PlaybackStartInfo) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *PlaybackStartInfo) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetIsPaused

`func (o *PlaybackStartInfo) GetIsPaused() bool`

GetIsPaused returns the IsPaused field if non-nil, zero value otherwise.

### GetIsPausedOk

`func (o *PlaybackStartInfo) GetIsPausedOk() (*bool, bool)`

GetIsPausedOk returns a tuple with the IsPaused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaused

`func (o *PlaybackStartInfo) SetIsPaused(v bool)`

SetIsPaused sets IsPaused field to given value.

### HasIsPaused

`func (o *PlaybackStartInfo) HasIsPaused() bool`

HasIsPaused returns a boolean if a field has been set.

### GetIsMuted

`func (o *PlaybackStartInfo) GetIsMuted() bool`

GetIsMuted returns the IsMuted field if non-nil, zero value otherwise.

### GetIsMutedOk

`func (o *PlaybackStartInfo) GetIsMutedOk() (*bool, bool)`

GetIsMutedOk returns a tuple with the IsMuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMuted

`func (o *PlaybackStartInfo) SetIsMuted(v bool)`

SetIsMuted sets IsMuted field to given value.

### HasIsMuted

`func (o *PlaybackStartInfo) HasIsMuted() bool`

HasIsMuted returns a boolean if a field has been set.

### GetPositionTicks

`func (o *PlaybackStartInfo) GetPositionTicks() int64`

GetPositionTicks returns the PositionTicks field if non-nil, zero value otherwise.

### GetPositionTicksOk

`func (o *PlaybackStartInfo) GetPositionTicksOk() (*int64, bool)`

GetPositionTicksOk returns a tuple with the PositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionTicks

`func (o *PlaybackStartInfo) SetPositionTicks(v int64)`

SetPositionTicks sets PositionTicks field to given value.

### HasPositionTicks

`func (o *PlaybackStartInfo) HasPositionTicks() bool`

HasPositionTicks returns a boolean if a field has been set.

### SetPositionTicksNil

`func (o *PlaybackStartInfo) SetPositionTicksNil(b bool)`

 SetPositionTicksNil sets the value for PositionTicks to be an explicit nil

### UnsetPositionTicks
`func (o *PlaybackStartInfo) UnsetPositionTicks()`

UnsetPositionTicks ensures that no value is present for PositionTicks, not even an explicit nil
### GetPlaybackStartTimeTicks

`func (o *PlaybackStartInfo) GetPlaybackStartTimeTicks() int64`

GetPlaybackStartTimeTicks returns the PlaybackStartTimeTicks field if non-nil, zero value otherwise.

### GetPlaybackStartTimeTicksOk

`func (o *PlaybackStartInfo) GetPlaybackStartTimeTicksOk() (*int64, bool)`

GetPlaybackStartTimeTicksOk returns a tuple with the PlaybackStartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackStartTimeTicks

`func (o *PlaybackStartInfo) SetPlaybackStartTimeTicks(v int64)`

SetPlaybackStartTimeTicks sets PlaybackStartTimeTicks field to given value.

### HasPlaybackStartTimeTicks

`func (o *PlaybackStartInfo) HasPlaybackStartTimeTicks() bool`

HasPlaybackStartTimeTicks returns a boolean if a field has been set.

### SetPlaybackStartTimeTicksNil

`func (o *PlaybackStartInfo) SetPlaybackStartTimeTicksNil(b bool)`

 SetPlaybackStartTimeTicksNil sets the value for PlaybackStartTimeTicks to be an explicit nil

### UnsetPlaybackStartTimeTicks
`func (o *PlaybackStartInfo) UnsetPlaybackStartTimeTicks()`

UnsetPlaybackStartTimeTicks ensures that no value is present for PlaybackStartTimeTicks, not even an explicit nil
### GetVolumeLevel

`func (o *PlaybackStartInfo) GetVolumeLevel() int32`

GetVolumeLevel returns the VolumeLevel field if non-nil, zero value otherwise.

### GetVolumeLevelOk

`func (o *PlaybackStartInfo) GetVolumeLevelOk() (*int32, bool)`

GetVolumeLevelOk returns a tuple with the VolumeLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeLevel

`func (o *PlaybackStartInfo) SetVolumeLevel(v int32)`

SetVolumeLevel sets VolumeLevel field to given value.

### HasVolumeLevel

`func (o *PlaybackStartInfo) HasVolumeLevel() bool`

HasVolumeLevel returns a boolean if a field has been set.

### SetVolumeLevelNil

`func (o *PlaybackStartInfo) SetVolumeLevelNil(b bool)`

 SetVolumeLevelNil sets the value for VolumeLevel to be an explicit nil

### UnsetVolumeLevel
`func (o *PlaybackStartInfo) UnsetVolumeLevel()`

UnsetVolumeLevel ensures that no value is present for VolumeLevel, not even an explicit nil
### GetBrightness

`func (o *PlaybackStartInfo) GetBrightness() int32`

GetBrightness returns the Brightness field if non-nil, zero value otherwise.

### GetBrightnessOk

`func (o *PlaybackStartInfo) GetBrightnessOk() (*int32, bool)`

GetBrightnessOk returns a tuple with the Brightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrightness

`func (o *PlaybackStartInfo) SetBrightness(v int32)`

SetBrightness sets Brightness field to given value.

### HasBrightness

`func (o *PlaybackStartInfo) HasBrightness() bool`

HasBrightness returns a boolean if a field has been set.

### SetBrightnessNil

`func (o *PlaybackStartInfo) SetBrightnessNil(b bool)`

 SetBrightnessNil sets the value for Brightness to be an explicit nil

### UnsetBrightness
`func (o *PlaybackStartInfo) UnsetBrightness()`

UnsetBrightness ensures that no value is present for Brightness, not even an explicit nil
### GetAspectRatio

`func (o *PlaybackStartInfo) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *PlaybackStartInfo) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *PlaybackStartInfo) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *PlaybackStartInfo) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### SetAspectRatioNil

`func (o *PlaybackStartInfo) SetAspectRatioNil(b bool)`

 SetAspectRatioNil sets the value for AspectRatio to be an explicit nil

### UnsetAspectRatio
`func (o *PlaybackStartInfo) UnsetAspectRatio()`

UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
### GetPlayMethod

`func (o *PlaybackStartInfo) GetPlayMethod() PlayMethod`

GetPlayMethod returns the PlayMethod field if non-nil, zero value otherwise.

### GetPlayMethodOk

`func (o *PlaybackStartInfo) GetPlayMethodOk() (*PlayMethod, bool)`

GetPlayMethodOk returns a tuple with the PlayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayMethod

`func (o *PlaybackStartInfo) SetPlayMethod(v PlayMethod)`

SetPlayMethod sets PlayMethod field to given value.

### HasPlayMethod

`func (o *PlaybackStartInfo) HasPlayMethod() bool`

HasPlayMethod returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *PlaybackStartInfo) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *PlaybackStartInfo) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *PlaybackStartInfo) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *PlaybackStartInfo) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### SetLiveStreamIdNil

`func (o *PlaybackStartInfo) SetLiveStreamIdNil(b bool)`

 SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil

### UnsetLiveStreamId
`func (o *PlaybackStartInfo) UnsetLiveStreamId()`

UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil
### GetPlaySessionId

`func (o *PlaybackStartInfo) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *PlaybackStartInfo) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *PlaybackStartInfo) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *PlaybackStartInfo) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### SetPlaySessionIdNil

`func (o *PlaybackStartInfo) SetPlaySessionIdNil(b bool)`

 SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil

### UnsetPlaySessionId
`func (o *PlaybackStartInfo) UnsetPlaySessionId()`

UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
### GetRepeatMode

`func (o *PlaybackStartInfo) GetRepeatMode() RepeatMode`

GetRepeatMode returns the RepeatMode field if non-nil, zero value otherwise.

### GetRepeatModeOk

`func (o *PlaybackStartInfo) GetRepeatModeOk() (*RepeatMode, bool)`

GetRepeatModeOk returns a tuple with the RepeatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatMode

`func (o *PlaybackStartInfo) SetRepeatMode(v RepeatMode)`

SetRepeatMode sets RepeatMode field to given value.

### HasRepeatMode

`func (o *PlaybackStartInfo) HasRepeatMode() bool`

HasRepeatMode returns a boolean if a field has been set.

### GetNowPlayingQueue

`func (o *PlaybackStartInfo) GetNowPlayingQueue() []QueueItem`

GetNowPlayingQueue returns the NowPlayingQueue field if non-nil, zero value otherwise.

### GetNowPlayingQueueOk

`func (o *PlaybackStartInfo) GetNowPlayingQueueOk() (*[]QueueItem, bool)`

GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueue

`func (o *PlaybackStartInfo) SetNowPlayingQueue(v []QueueItem)`

SetNowPlayingQueue sets NowPlayingQueue field to given value.

### HasNowPlayingQueue

`func (o *PlaybackStartInfo) HasNowPlayingQueue() bool`

HasNowPlayingQueue returns a boolean if a field has been set.

### SetNowPlayingQueueNil

`func (o *PlaybackStartInfo) SetNowPlayingQueueNil(b bool)`

 SetNowPlayingQueueNil sets the value for NowPlayingQueue to be an explicit nil

### UnsetNowPlayingQueue
`func (o *PlaybackStartInfo) UnsetNowPlayingQueue()`

UnsetNowPlayingQueue ensures that no value is present for NowPlayingQueue, not even an explicit nil
### GetPlaylistItemId

`func (o *PlaybackStartInfo) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *PlaybackStartInfo) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *PlaybackStartInfo) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *PlaybackStartInfo) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### SetPlaylistItemIdNil

`func (o *PlaybackStartInfo) SetPlaylistItemIdNil(b bool)`

 SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil

### UnsetPlaylistItemId
`func (o *PlaybackStartInfo) UnsetPlaylistItemId()`

UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


