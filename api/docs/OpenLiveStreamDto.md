# OpenLiveStreamDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenToken** | Pointer to **NullableString** | Gets or sets the open token. | [optional] 
**UserId** | Pointer to **NullableString** | Gets or sets the user id. | [optional] 
**PlaySessionId** | Pointer to **NullableString** | Gets or sets the play session id. | [optional] 
**MaxStreamingBitrate** | Pointer to **NullableInt32** | Gets or sets the max streaming bitrate. | [optional] 
**StartTimeTicks** | Pointer to **NullableInt64** | Gets or sets the start time in ticks. | [optional] 
**AudioStreamIndex** | Pointer to **NullableInt32** | Gets or sets the audio stream index. | [optional] 
**SubtitleStreamIndex** | Pointer to **NullableInt32** | Gets or sets the subtitle stream index. | [optional] 
**MaxAudioChannels** | Pointer to **NullableInt32** | Gets or sets the max audio channels. | [optional] 
**ItemId** | Pointer to **NullableString** | Gets or sets the item id. | [optional] 
**EnableDirectPlay** | Pointer to **NullableBool** | Gets or sets a value indicating whether to enable direct play. | [optional] 
**EnableDirectStream** | Pointer to **NullableBool** | Gets or sets a value indicating whether to enale direct stream. | [optional] 
**DeviceProfile** | Pointer to [**NullableClientCapabilitiesDeviceProfile**](ClientCapabilitiesDeviceProfile.md) |  | [optional] 
**DirectPlayProtocols** | Pointer to [**[]MediaProtocol**](MediaProtocol.md) | Gets or sets the device play protocols. | [optional] 

## Methods

### NewOpenLiveStreamDto

`func NewOpenLiveStreamDto() *OpenLiveStreamDto`

NewOpenLiveStreamDto instantiates a new OpenLiveStreamDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenLiveStreamDtoWithDefaults

`func NewOpenLiveStreamDtoWithDefaults() *OpenLiveStreamDto`

NewOpenLiveStreamDtoWithDefaults instantiates a new OpenLiveStreamDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenToken

`func (o *OpenLiveStreamDto) GetOpenToken() string`

GetOpenToken returns the OpenToken field if non-nil, zero value otherwise.

### GetOpenTokenOk

`func (o *OpenLiveStreamDto) GetOpenTokenOk() (*string, bool)`

GetOpenTokenOk returns a tuple with the OpenToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenToken

`func (o *OpenLiveStreamDto) SetOpenToken(v string)`

SetOpenToken sets OpenToken field to given value.

### HasOpenToken

`func (o *OpenLiveStreamDto) HasOpenToken() bool`

HasOpenToken returns a boolean if a field has been set.

### SetOpenTokenNil

`func (o *OpenLiveStreamDto) SetOpenTokenNil(b bool)`

 SetOpenTokenNil sets the value for OpenToken to be an explicit nil

### UnsetOpenToken
`func (o *OpenLiveStreamDto) UnsetOpenToken()`

UnsetOpenToken ensures that no value is present for OpenToken, not even an explicit nil
### GetUserId

`func (o *OpenLiveStreamDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OpenLiveStreamDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OpenLiveStreamDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OpenLiveStreamDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *OpenLiveStreamDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *OpenLiveStreamDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetPlaySessionId

`func (o *OpenLiveStreamDto) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *OpenLiveStreamDto) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *OpenLiveStreamDto) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *OpenLiveStreamDto) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### SetPlaySessionIdNil

`func (o *OpenLiveStreamDto) SetPlaySessionIdNil(b bool)`

 SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil

### UnsetPlaySessionId
`func (o *OpenLiveStreamDto) UnsetPlaySessionId()`

UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
### GetMaxStreamingBitrate

`func (o *OpenLiveStreamDto) GetMaxStreamingBitrate() int32`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *OpenLiveStreamDto) GetMaxStreamingBitrateOk() (*int32, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *OpenLiveStreamDto) SetMaxStreamingBitrate(v int32)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *OpenLiveStreamDto) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *OpenLiveStreamDto) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *OpenLiveStreamDto) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetStartTimeTicks

`func (o *OpenLiveStreamDto) GetStartTimeTicks() int64`

GetStartTimeTicks returns the StartTimeTicks field if non-nil, zero value otherwise.

### GetStartTimeTicksOk

`func (o *OpenLiveStreamDto) GetStartTimeTicksOk() (*int64, bool)`

GetStartTimeTicksOk returns a tuple with the StartTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeTicks

`func (o *OpenLiveStreamDto) SetStartTimeTicks(v int64)`

SetStartTimeTicks sets StartTimeTicks field to given value.

### HasStartTimeTicks

`func (o *OpenLiveStreamDto) HasStartTimeTicks() bool`

HasStartTimeTicks returns a boolean if a field has been set.

### SetStartTimeTicksNil

`func (o *OpenLiveStreamDto) SetStartTimeTicksNil(b bool)`

 SetStartTimeTicksNil sets the value for StartTimeTicks to be an explicit nil

### UnsetStartTimeTicks
`func (o *OpenLiveStreamDto) UnsetStartTimeTicks()`

UnsetStartTimeTicks ensures that no value is present for StartTimeTicks, not even an explicit nil
### GetAudioStreamIndex

`func (o *OpenLiveStreamDto) GetAudioStreamIndex() int32`

GetAudioStreamIndex returns the AudioStreamIndex field if non-nil, zero value otherwise.

### GetAudioStreamIndexOk

`func (o *OpenLiveStreamDto) GetAudioStreamIndexOk() (*int32, bool)`

GetAudioStreamIndexOk returns a tuple with the AudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioStreamIndex

`func (o *OpenLiveStreamDto) SetAudioStreamIndex(v int32)`

SetAudioStreamIndex sets AudioStreamIndex field to given value.

### HasAudioStreamIndex

`func (o *OpenLiveStreamDto) HasAudioStreamIndex() bool`

HasAudioStreamIndex returns a boolean if a field has been set.

### SetAudioStreamIndexNil

`func (o *OpenLiveStreamDto) SetAudioStreamIndexNil(b bool)`

 SetAudioStreamIndexNil sets the value for AudioStreamIndex to be an explicit nil

### UnsetAudioStreamIndex
`func (o *OpenLiveStreamDto) UnsetAudioStreamIndex()`

UnsetAudioStreamIndex ensures that no value is present for AudioStreamIndex, not even an explicit nil
### GetSubtitleStreamIndex

`func (o *OpenLiveStreamDto) GetSubtitleStreamIndex() int32`

GetSubtitleStreamIndex returns the SubtitleStreamIndex field if non-nil, zero value otherwise.

### GetSubtitleStreamIndexOk

`func (o *OpenLiveStreamDto) GetSubtitleStreamIndexOk() (*int32, bool)`

GetSubtitleStreamIndexOk returns a tuple with the SubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleStreamIndex

`func (o *OpenLiveStreamDto) SetSubtitleStreamIndex(v int32)`

SetSubtitleStreamIndex sets SubtitleStreamIndex field to given value.

### HasSubtitleStreamIndex

`func (o *OpenLiveStreamDto) HasSubtitleStreamIndex() bool`

HasSubtitleStreamIndex returns a boolean if a field has been set.

### SetSubtitleStreamIndexNil

`func (o *OpenLiveStreamDto) SetSubtitleStreamIndexNil(b bool)`

 SetSubtitleStreamIndexNil sets the value for SubtitleStreamIndex to be an explicit nil

### UnsetSubtitleStreamIndex
`func (o *OpenLiveStreamDto) UnsetSubtitleStreamIndex()`

UnsetSubtitleStreamIndex ensures that no value is present for SubtitleStreamIndex, not even an explicit nil
### GetMaxAudioChannels

`func (o *OpenLiveStreamDto) GetMaxAudioChannels() int32`

GetMaxAudioChannels returns the MaxAudioChannels field if non-nil, zero value otherwise.

### GetMaxAudioChannelsOk

`func (o *OpenLiveStreamDto) GetMaxAudioChannelsOk() (*int32, bool)`

GetMaxAudioChannelsOk returns a tuple with the MaxAudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudioChannels

`func (o *OpenLiveStreamDto) SetMaxAudioChannels(v int32)`

SetMaxAudioChannels sets MaxAudioChannels field to given value.

### HasMaxAudioChannels

`func (o *OpenLiveStreamDto) HasMaxAudioChannels() bool`

HasMaxAudioChannels returns a boolean if a field has been set.

### SetMaxAudioChannelsNil

`func (o *OpenLiveStreamDto) SetMaxAudioChannelsNil(b bool)`

 SetMaxAudioChannelsNil sets the value for MaxAudioChannels to be an explicit nil

### UnsetMaxAudioChannels
`func (o *OpenLiveStreamDto) UnsetMaxAudioChannels()`

UnsetMaxAudioChannels ensures that no value is present for MaxAudioChannels, not even an explicit nil
### GetItemId

`func (o *OpenLiveStreamDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *OpenLiveStreamDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *OpenLiveStreamDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *OpenLiveStreamDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *OpenLiveStreamDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *OpenLiveStreamDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetEnableDirectPlay

`func (o *OpenLiveStreamDto) GetEnableDirectPlay() bool`

GetEnableDirectPlay returns the EnableDirectPlay field if non-nil, zero value otherwise.

### GetEnableDirectPlayOk

`func (o *OpenLiveStreamDto) GetEnableDirectPlayOk() (*bool, bool)`

GetEnableDirectPlayOk returns a tuple with the EnableDirectPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectPlay

`func (o *OpenLiveStreamDto) SetEnableDirectPlay(v bool)`

SetEnableDirectPlay sets EnableDirectPlay field to given value.

### HasEnableDirectPlay

`func (o *OpenLiveStreamDto) HasEnableDirectPlay() bool`

HasEnableDirectPlay returns a boolean if a field has been set.

### SetEnableDirectPlayNil

`func (o *OpenLiveStreamDto) SetEnableDirectPlayNil(b bool)`

 SetEnableDirectPlayNil sets the value for EnableDirectPlay to be an explicit nil

### UnsetEnableDirectPlay
`func (o *OpenLiveStreamDto) UnsetEnableDirectPlay()`

UnsetEnableDirectPlay ensures that no value is present for EnableDirectPlay, not even an explicit nil
### GetEnableDirectStream

`func (o *OpenLiveStreamDto) GetEnableDirectStream() bool`

GetEnableDirectStream returns the EnableDirectStream field if non-nil, zero value otherwise.

### GetEnableDirectStreamOk

`func (o *OpenLiveStreamDto) GetEnableDirectStreamOk() (*bool, bool)`

GetEnableDirectStreamOk returns a tuple with the EnableDirectStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDirectStream

`func (o *OpenLiveStreamDto) SetEnableDirectStream(v bool)`

SetEnableDirectStream sets EnableDirectStream field to given value.

### HasEnableDirectStream

`func (o *OpenLiveStreamDto) HasEnableDirectStream() bool`

HasEnableDirectStream returns a boolean if a field has been set.

### SetEnableDirectStreamNil

`func (o *OpenLiveStreamDto) SetEnableDirectStreamNil(b bool)`

 SetEnableDirectStreamNil sets the value for EnableDirectStream to be an explicit nil

### UnsetEnableDirectStream
`func (o *OpenLiveStreamDto) UnsetEnableDirectStream()`

UnsetEnableDirectStream ensures that no value is present for EnableDirectStream, not even an explicit nil
### GetDeviceProfile

`func (o *OpenLiveStreamDto) GetDeviceProfile() ClientCapabilitiesDeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *OpenLiveStreamDto) GetDeviceProfileOk() (*ClientCapabilitiesDeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *OpenLiveStreamDto) SetDeviceProfile(v ClientCapabilitiesDeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *OpenLiveStreamDto) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### SetDeviceProfileNil

`func (o *OpenLiveStreamDto) SetDeviceProfileNil(b bool)`

 SetDeviceProfileNil sets the value for DeviceProfile to be an explicit nil

### UnsetDeviceProfile
`func (o *OpenLiveStreamDto) UnsetDeviceProfile()`

UnsetDeviceProfile ensures that no value is present for DeviceProfile, not even an explicit nil
### GetDirectPlayProtocols

`func (o *OpenLiveStreamDto) GetDirectPlayProtocols() []MediaProtocol`

GetDirectPlayProtocols returns the DirectPlayProtocols field if non-nil, zero value otherwise.

### GetDirectPlayProtocolsOk

`func (o *OpenLiveStreamDto) GetDirectPlayProtocolsOk() (*[]MediaProtocol, bool)`

GetDirectPlayProtocolsOk returns a tuple with the DirectPlayProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPlayProtocols

`func (o *OpenLiveStreamDto) SetDirectPlayProtocols(v []MediaProtocol)`

SetDirectPlayProtocols sets DirectPlayProtocols field to given value.

### HasDirectPlayProtocols

`func (o *OpenLiveStreamDto) HasDirectPlayProtocols() bool`

HasDirectPlayProtocols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


