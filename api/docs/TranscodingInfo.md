# TranscodingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioCodec** | Pointer to **NullableString** |  | [optional] 
**VideoCodec** | Pointer to **NullableString** |  | [optional] 
**Container** | Pointer to **NullableString** |  | [optional] 
**IsVideoDirect** | Pointer to **bool** |  | [optional] 
**IsAudioDirect** | Pointer to **bool** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**Framerate** | Pointer to **NullableFloat32** |  | [optional] 
**CompletionPercentage** | Pointer to **NullableFloat64** |  | [optional] 
**Width** | Pointer to **NullableInt32** |  | [optional] 
**Height** | Pointer to **NullableInt32** |  | [optional] 
**AudioChannels** | Pointer to **NullableInt32** |  | [optional] 
**HardwareAccelerationType** | Pointer to [**NullableHardwareEncodingType**](HardwareEncodingType.md) |  | [optional] 
**TranscodeReasons** | Pointer to [**[]TranscodeReason**](TranscodeReason.md) |  | [optional] 

## Methods

### NewTranscodingInfo

`func NewTranscodingInfo() *TranscodingInfo`

NewTranscodingInfo instantiates a new TranscodingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscodingInfoWithDefaults

`func NewTranscodingInfoWithDefaults() *TranscodingInfo`

NewTranscodingInfoWithDefaults instantiates a new TranscodingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioCodec

`func (o *TranscodingInfo) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *TranscodingInfo) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *TranscodingInfo) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *TranscodingInfo) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### SetAudioCodecNil

`func (o *TranscodingInfo) SetAudioCodecNil(b bool)`

 SetAudioCodecNil sets the value for AudioCodec to be an explicit nil

### UnsetAudioCodec
`func (o *TranscodingInfo) UnsetAudioCodec()`

UnsetAudioCodec ensures that no value is present for AudioCodec, not even an explicit nil
### GetVideoCodec

`func (o *TranscodingInfo) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *TranscodingInfo) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *TranscodingInfo) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *TranscodingInfo) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### SetVideoCodecNil

`func (o *TranscodingInfo) SetVideoCodecNil(b bool)`

 SetVideoCodecNil sets the value for VideoCodec to be an explicit nil

### UnsetVideoCodec
`func (o *TranscodingInfo) UnsetVideoCodec()`

UnsetVideoCodec ensures that no value is present for VideoCodec, not even an explicit nil
### GetContainer

`func (o *TranscodingInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *TranscodingInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *TranscodingInfo) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *TranscodingInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *TranscodingInfo) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *TranscodingInfo) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetIsVideoDirect

`func (o *TranscodingInfo) GetIsVideoDirect() bool`

GetIsVideoDirect returns the IsVideoDirect field if non-nil, zero value otherwise.

### GetIsVideoDirectOk

`func (o *TranscodingInfo) GetIsVideoDirectOk() (*bool, bool)`

GetIsVideoDirectOk returns a tuple with the IsVideoDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVideoDirect

`func (o *TranscodingInfo) SetIsVideoDirect(v bool)`

SetIsVideoDirect sets IsVideoDirect field to given value.

### HasIsVideoDirect

`func (o *TranscodingInfo) HasIsVideoDirect() bool`

HasIsVideoDirect returns a boolean if a field has been set.

### GetIsAudioDirect

`func (o *TranscodingInfo) GetIsAudioDirect() bool`

GetIsAudioDirect returns the IsAudioDirect field if non-nil, zero value otherwise.

### GetIsAudioDirectOk

`func (o *TranscodingInfo) GetIsAudioDirectOk() (*bool, bool)`

GetIsAudioDirectOk returns a tuple with the IsAudioDirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAudioDirect

`func (o *TranscodingInfo) SetIsAudioDirect(v bool)`

SetIsAudioDirect sets IsAudioDirect field to given value.

### HasIsAudioDirect

`func (o *TranscodingInfo) HasIsAudioDirect() bool`

HasIsAudioDirect returns a boolean if a field has been set.

### GetBitrate

`func (o *TranscodingInfo) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *TranscodingInfo) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *TranscodingInfo) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *TranscodingInfo) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *TranscodingInfo) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *TranscodingInfo) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetFramerate

`func (o *TranscodingInfo) GetFramerate() float32`

GetFramerate returns the Framerate field if non-nil, zero value otherwise.

### GetFramerateOk

`func (o *TranscodingInfo) GetFramerateOk() (*float32, bool)`

GetFramerateOk returns a tuple with the Framerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramerate

`func (o *TranscodingInfo) SetFramerate(v float32)`

SetFramerate sets Framerate field to given value.

### HasFramerate

`func (o *TranscodingInfo) HasFramerate() bool`

HasFramerate returns a boolean if a field has been set.

### SetFramerateNil

`func (o *TranscodingInfo) SetFramerateNil(b bool)`

 SetFramerateNil sets the value for Framerate to be an explicit nil

### UnsetFramerate
`func (o *TranscodingInfo) UnsetFramerate()`

UnsetFramerate ensures that no value is present for Framerate, not even an explicit nil
### GetCompletionPercentage

`func (o *TranscodingInfo) GetCompletionPercentage() float64`

GetCompletionPercentage returns the CompletionPercentage field if non-nil, zero value otherwise.

### GetCompletionPercentageOk

`func (o *TranscodingInfo) GetCompletionPercentageOk() (*float64, bool)`

GetCompletionPercentageOk returns a tuple with the CompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionPercentage

`func (o *TranscodingInfo) SetCompletionPercentage(v float64)`

SetCompletionPercentage sets CompletionPercentage field to given value.

### HasCompletionPercentage

`func (o *TranscodingInfo) HasCompletionPercentage() bool`

HasCompletionPercentage returns a boolean if a field has been set.

### SetCompletionPercentageNil

`func (o *TranscodingInfo) SetCompletionPercentageNil(b bool)`

 SetCompletionPercentageNil sets the value for CompletionPercentage to be an explicit nil

### UnsetCompletionPercentage
`func (o *TranscodingInfo) UnsetCompletionPercentage()`

UnsetCompletionPercentage ensures that no value is present for CompletionPercentage, not even an explicit nil
### GetWidth

`func (o *TranscodingInfo) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *TranscodingInfo) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *TranscodingInfo) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *TranscodingInfo) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *TranscodingInfo) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *TranscodingInfo) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetHeight

`func (o *TranscodingInfo) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *TranscodingInfo) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *TranscodingInfo) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *TranscodingInfo) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *TranscodingInfo) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *TranscodingInfo) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetAudioChannels

`func (o *TranscodingInfo) GetAudioChannels() int32`

GetAudioChannels returns the AudioChannels field if non-nil, zero value otherwise.

### GetAudioChannelsOk

`func (o *TranscodingInfo) GetAudioChannelsOk() (*int32, bool)`

GetAudioChannelsOk returns a tuple with the AudioChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioChannels

`func (o *TranscodingInfo) SetAudioChannels(v int32)`

SetAudioChannels sets AudioChannels field to given value.

### HasAudioChannels

`func (o *TranscodingInfo) HasAudioChannels() bool`

HasAudioChannels returns a boolean if a field has been set.

### SetAudioChannelsNil

`func (o *TranscodingInfo) SetAudioChannelsNil(b bool)`

 SetAudioChannelsNil sets the value for AudioChannels to be an explicit nil

### UnsetAudioChannels
`func (o *TranscodingInfo) UnsetAudioChannels()`

UnsetAudioChannels ensures that no value is present for AudioChannels, not even an explicit nil
### GetHardwareAccelerationType

`func (o *TranscodingInfo) GetHardwareAccelerationType() HardwareEncodingType`

GetHardwareAccelerationType returns the HardwareAccelerationType field if non-nil, zero value otherwise.

### GetHardwareAccelerationTypeOk

`func (o *TranscodingInfo) GetHardwareAccelerationTypeOk() (*HardwareEncodingType, bool)`

GetHardwareAccelerationTypeOk returns a tuple with the HardwareAccelerationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareAccelerationType

`func (o *TranscodingInfo) SetHardwareAccelerationType(v HardwareEncodingType)`

SetHardwareAccelerationType sets HardwareAccelerationType field to given value.

### HasHardwareAccelerationType

`func (o *TranscodingInfo) HasHardwareAccelerationType() bool`

HasHardwareAccelerationType returns a boolean if a field has been set.

### SetHardwareAccelerationTypeNil

`func (o *TranscodingInfo) SetHardwareAccelerationTypeNil(b bool)`

 SetHardwareAccelerationTypeNil sets the value for HardwareAccelerationType to be an explicit nil

### UnsetHardwareAccelerationType
`func (o *TranscodingInfo) UnsetHardwareAccelerationType()`

UnsetHardwareAccelerationType ensures that no value is present for HardwareAccelerationType, not even an explicit nil
### GetTranscodeReasons

`func (o *TranscodingInfo) GetTranscodeReasons() []TranscodeReason`

GetTranscodeReasons returns the TranscodeReasons field if non-nil, zero value otherwise.

### GetTranscodeReasonsOk

`func (o *TranscodingInfo) GetTranscodeReasonsOk() (*[]TranscodeReason, bool)`

GetTranscodeReasonsOk returns a tuple with the TranscodeReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodeReasons

`func (o *TranscodingInfo) SetTranscodeReasons(v []TranscodeReason)`

SetTranscodeReasons sets TranscodeReasons field to given value.

### HasTranscodeReasons

`func (o *TranscodingInfo) HasTranscodeReasons() bool`

HasTranscodeReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


