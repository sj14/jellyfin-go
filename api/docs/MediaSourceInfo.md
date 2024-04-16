# MediaSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | Pointer to [**MediaProtocol**](MediaProtocol.md) |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**EncoderPath** | Pointer to **NullableString** |  | [optional] 
**EncoderProtocol** | Pointer to [**NullableMediaProtocol**](MediaProtocol.md) |  | [optional] 
**Type** | Pointer to [**MediaSourceType**](MediaSourceType.md) |  | [optional] 
**Container** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableInt64** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**IsRemote** | Pointer to **bool** | Gets or sets a value indicating whether the media is remote.  Differentiate internet url vs local network. | [optional] 
**ETag** | Pointer to **NullableString** |  | [optional] 
**RunTimeTicks** | Pointer to **NullableInt64** |  | [optional] 
**ReadAtNativeFramerate** | Pointer to **bool** |  | [optional] 
**IgnoreDts** | Pointer to **bool** |  | [optional] 
**IgnoreIndex** | Pointer to **bool** |  | [optional] 
**GenPtsInput** | Pointer to **bool** |  | [optional] 
**SupportsTranscoding** | Pointer to **bool** |  | [optional] 
**SupportsDirectStream** | Pointer to **bool** |  | [optional] 
**SupportsDirectPlay** | Pointer to **bool** |  | [optional] 
**IsInfiniteStream** | Pointer to **bool** |  | [optional] 
**RequiresOpening** | Pointer to **bool** |  | [optional] 
**OpenToken** | Pointer to **NullableString** |  | [optional] 
**RequiresClosing** | Pointer to **bool** |  | [optional] 
**LiveStreamId** | Pointer to **NullableString** |  | [optional] 
**BufferMs** | Pointer to **NullableInt32** |  | [optional] 
**RequiresLooping** | Pointer to **bool** |  | [optional] 
**SupportsProbing** | Pointer to **bool** |  | [optional] 
**VideoType** | Pointer to [**NullableVideoType**](VideoType.md) |  | [optional] 
**IsoType** | Pointer to [**NullableIsoType**](IsoType.md) |  | [optional] 
**Video3DFormat** | Pointer to [**NullableVideo3DFormat**](Video3DFormat.md) |  | [optional] 
**MediaStreams** | Pointer to [**[]MediaStream**](MediaStream.md) |  | [optional] 
**MediaAttachments** | Pointer to [**[]MediaAttachment**](MediaAttachment.md) |  | [optional] 
**Formats** | Pointer to **[]string** |  | [optional] 
**Bitrate** | Pointer to **NullableInt32** |  | [optional] 
**Timestamp** | Pointer to [**NullableTransportStreamTimestamp**](TransportStreamTimestamp.md) |  | [optional] 
**RequiredHttpHeaders** | Pointer to **map[string]string** |  | [optional] 
**TranscodingUrl** | Pointer to **NullableString** |  | [optional] 
**TranscodingSubProtocol** | Pointer to **NullableString** |  | [optional] 
**TranscodingContainer** | Pointer to **NullableString** |  | [optional] 
**AnalyzeDurationMs** | Pointer to **NullableInt32** |  | [optional] 
**DefaultAudioStreamIndex** | Pointer to **NullableInt32** |  | [optional] 
**DefaultSubtitleStreamIndex** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewMediaSourceInfo

`func NewMediaSourceInfo() *MediaSourceInfo`

NewMediaSourceInfo instantiates a new MediaSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaSourceInfoWithDefaults

`func NewMediaSourceInfoWithDefaults() *MediaSourceInfo`

NewMediaSourceInfoWithDefaults instantiates a new MediaSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *MediaSourceInfo) GetProtocol() MediaProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *MediaSourceInfo) GetProtocolOk() (*MediaProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *MediaSourceInfo) SetProtocol(v MediaProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *MediaSourceInfo) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetId

`func (o *MediaSourceInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaSourceInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaSourceInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MediaSourceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MediaSourceInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MediaSourceInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPath

`func (o *MediaSourceInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MediaSourceInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MediaSourceInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MediaSourceInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *MediaSourceInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *MediaSourceInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetEncoderPath

`func (o *MediaSourceInfo) GetEncoderPath() string`

GetEncoderPath returns the EncoderPath field if non-nil, zero value otherwise.

### GetEncoderPathOk

`func (o *MediaSourceInfo) GetEncoderPathOk() (*string, bool)`

GetEncoderPathOk returns a tuple with the EncoderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderPath

`func (o *MediaSourceInfo) SetEncoderPath(v string)`

SetEncoderPath sets EncoderPath field to given value.

### HasEncoderPath

`func (o *MediaSourceInfo) HasEncoderPath() bool`

HasEncoderPath returns a boolean if a field has been set.

### SetEncoderPathNil

`func (o *MediaSourceInfo) SetEncoderPathNil(b bool)`

 SetEncoderPathNil sets the value for EncoderPath to be an explicit nil

### UnsetEncoderPath
`func (o *MediaSourceInfo) UnsetEncoderPath()`

UnsetEncoderPath ensures that no value is present for EncoderPath, not even an explicit nil
### GetEncoderProtocol

`func (o *MediaSourceInfo) GetEncoderProtocol() MediaProtocol`

GetEncoderProtocol returns the EncoderProtocol field if non-nil, zero value otherwise.

### GetEncoderProtocolOk

`func (o *MediaSourceInfo) GetEncoderProtocolOk() (*MediaProtocol, bool)`

GetEncoderProtocolOk returns a tuple with the EncoderProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderProtocol

`func (o *MediaSourceInfo) SetEncoderProtocol(v MediaProtocol)`

SetEncoderProtocol sets EncoderProtocol field to given value.

### HasEncoderProtocol

`func (o *MediaSourceInfo) HasEncoderProtocol() bool`

HasEncoderProtocol returns a boolean if a field has been set.

### SetEncoderProtocolNil

`func (o *MediaSourceInfo) SetEncoderProtocolNil(b bool)`

 SetEncoderProtocolNil sets the value for EncoderProtocol to be an explicit nil

### UnsetEncoderProtocol
`func (o *MediaSourceInfo) UnsetEncoderProtocol()`

UnsetEncoderProtocol ensures that no value is present for EncoderProtocol, not even an explicit nil
### GetType

`func (o *MediaSourceInfo) GetType() MediaSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MediaSourceInfo) GetTypeOk() (*MediaSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MediaSourceInfo) SetType(v MediaSourceType)`

SetType sets Type field to given value.

### HasType

`func (o *MediaSourceInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContainer

`func (o *MediaSourceInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *MediaSourceInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *MediaSourceInfo) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *MediaSourceInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *MediaSourceInfo) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *MediaSourceInfo) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetSize

`func (o *MediaSourceInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MediaSourceInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MediaSourceInfo) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *MediaSourceInfo) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *MediaSourceInfo) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *MediaSourceInfo) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetName

`func (o *MediaSourceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MediaSourceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MediaSourceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MediaSourceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MediaSourceInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MediaSourceInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsRemote

`func (o *MediaSourceInfo) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *MediaSourceInfo) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *MediaSourceInfo) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *MediaSourceInfo) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetETag

`func (o *MediaSourceInfo) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MediaSourceInfo) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MediaSourceInfo) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MediaSourceInfo) HasETag() bool`

HasETag returns a boolean if a field has been set.

### SetETagNil

`func (o *MediaSourceInfo) SetETagNil(b bool)`

 SetETagNil sets the value for ETag to be an explicit nil

### UnsetETag
`func (o *MediaSourceInfo) UnsetETag()`

UnsetETag ensures that no value is present for ETag, not even an explicit nil
### GetRunTimeTicks

`func (o *MediaSourceInfo) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *MediaSourceInfo) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *MediaSourceInfo) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *MediaSourceInfo) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *MediaSourceInfo) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *MediaSourceInfo) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetReadAtNativeFramerate

`func (o *MediaSourceInfo) GetReadAtNativeFramerate() bool`

GetReadAtNativeFramerate returns the ReadAtNativeFramerate field if non-nil, zero value otherwise.

### GetReadAtNativeFramerateOk

`func (o *MediaSourceInfo) GetReadAtNativeFramerateOk() (*bool, bool)`

GetReadAtNativeFramerateOk returns a tuple with the ReadAtNativeFramerate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAtNativeFramerate

`func (o *MediaSourceInfo) SetReadAtNativeFramerate(v bool)`

SetReadAtNativeFramerate sets ReadAtNativeFramerate field to given value.

### HasReadAtNativeFramerate

`func (o *MediaSourceInfo) HasReadAtNativeFramerate() bool`

HasReadAtNativeFramerate returns a boolean if a field has been set.

### GetIgnoreDts

`func (o *MediaSourceInfo) GetIgnoreDts() bool`

GetIgnoreDts returns the IgnoreDts field if non-nil, zero value otherwise.

### GetIgnoreDtsOk

`func (o *MediaSourceInfo) GetIgnoreDtsOk() (*bool, bool)`

GetIgnoreDtsOk returns a tuple with the IgnoreDts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDts

`func (o *MediaSourceInfo) SetIgnoreDts(v bool)`

SetIgnoreDts sets IgnoreDts field to given value.

### HasIgnoreDts

`func (o *MediaSourceInfo) HasIgnoreDts() bool`

HasIgnoreDts returns a boolean if a field has been set.

### GetIgnoreIndex

`func (o *MediaSourceInfo) GetIgnoreIndex() bool`

GetIgnoreIndex returns the IgnoreIndex field if non-nil, zero value otherwise.

### GetIgnoreIndexOk

`func (o *MediaSourceInfo) GetIgnoreIndexOk() (*bool, bool)`

GetIgnoreIndexOk returns a tuple with the IgnoreIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreIndex

`func (o *MediaSourceInfo) SetIgnoreIndex(v bool)`

SetIgnoreIndex sets IgnoreIndex field to given value.

### HasIgnoreIndex

`func (o *MediaSourceInfo) HasIgnoreIndex() bool`

HasIgnoreIndex returns a boolean if a field has been set.

### GetGenPtsInput

`func (o *MediaSourceInfo) GetGenPtsInput() bool`

GetGenPtsInput returns the GenPtsInput field if non-nil, zero value otherwise.

### GetGenPtsInputOk

`func (o *MediaSourceInfo) GetGenPtsInputOk() (*bool, bool)`

GetGenPtsInputOk returns a tuple with the GenPtsInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenPtsInput

`func (o *MediaSourceInfo) SetGenPtsInput(v bool)`

SetGenPtsInput sets GenPtsInput field to given value.

### HasGenPtsInput

`func (o *MediaSourceInfo) HasGenPtsInput() bool`

HasGenPtsInput returns a boolean if a field has been set.

### GetSupportsTranscoding

`func (o *MediaSourceInfo) GetSupportsTranscoding() bool`

GetSupportsTranscoding returns the SupportsTranscoding field if non-nil, zero value otherwise.

### GetSupportsTranscodingOk

`func (o *MediaSourceInfo) GetSupportsTranscodingOk() (*bool, bool)`

GetSupportsTranscodingOk returns a tuple with the SupportsTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsTranscoding

`func (o *MediaSourceInfo) SetSupportsTranscoding(v bool)`

SetSupportsTranscoding sets SupportsTranscoding field to given value.

### HasSupportsTranscoding

`func (o *MediaSourceInfo) HasSupportsTranscoding() bool`

HasSupportsTranscoding returns a boolean if a field has been set.

### GetSupportsDirectStream

`func (o *MediaSourceInfo) GetSupportsDirectStream() bool`

GetSupportsDirectStream returns the SupportsDirectStream field if non-nil, zero value otherwise.

### GetSupportsDirectStreamOk

`func (o *MediaSourceInfo) GetSupportsDirectStreamOk() (*bool, bool)`

GetSupportsDirectStreamOk returns a tuple with the SupportsDirectStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDirectStream

`func (o *MediaSourceInfo) SetSupportsDirectStream(v bool)`

SetSupportsDirectStream sets SupportsDirectStream field to given value.

### HasSupportsDirectStream

`func (o *MediaSourceInfo) HasSupportsDirectStream() bool`

HasSupportsDirectStream returns a boolean if a field has been set.

### GetSupportsDirectPlay

`func (o *MediaSourceInfo) GetSupportsDirectPlay() bool`

GetSupportsDirectPlay returns the SupportsDirectPlay field if non-nil, zero value otherwise.

### GetSupportsDirectPlayOk

`func (o *MediaSourceInfo) GetSupportsDirectPlayOk() (*bool, bool)`

GetSupportsDirectPlayOk returns a tuple with the SupportsDirectPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDirectPlay

`func (o *MediaSourceInfo) SetSupportsDirectPlay(v bool)`

SetSupportsDirectPlay sets SupportsDirectPlay field to given value.

### HasSupportsDirectPlay

`func (o *MediaSourceInfo) HasSupportsDirectPlay() bool`

HasSupportsDirectPlay returns a boolean if a field has been set.

### GetIsInfiniteStream

`func (o *MediaSourceInfo) GetIsInfiniteStream() bool`

GetIsInfiniteStream returns the IsInfiniteStream field if non-nil, zero value otherwise.

### GetIsInfiniteStreamOk

`func (o *MediaSourceInfo) GetIsInfiniteStreamOk() (*bool, bool)`

GetIsInfiniteStreamOk returns a tuple with the IsInfiniteStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInfiniteStream

`func (o *MediaSourceInfo) SetIsInfiniteStream(v bool)`

SetIsInfiniteStream sets IsInfiniteStream field to given value.

### HasIsInfiniteStream

`func (o *MediaSourceInfo) HasIsInfiniteStream() bool`

HasIsInfiniteStream returns a boolean if a field has been set.

### GetRequiresOpening

`func (o *MediaSourceInfo) GetRequiresOpening() bool`

GetRequiresOpening returns the RequiresOpening field if non-nil, zero value otherwise.

### GetRequiresOpeningOk

`func (o *MediaSourceInfo) GetRequiresOpeningOk() (*bool, bool)`

GetRequiresOpeningOk returns a tuple with the RequiresOpening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresOpening

`func (o *MediaSourceInfo) SetRequiresOpening(v bool)`

SetRequiresOpening sets RequiresOpening field to given value.

### HasRequiresOpening

`func (o *MediaSourceInfo) HasRequiresOpening() bool`

HasRequiresOpening returns a boolean if a field has been set.

### GetOpenToken

`func (o *MediaSourceInfo) GetOpenToken() string`

GetOpenToken returns the OpenToken field if non-nil, zero value otherwise.

### GetOpenTokenOk

`func (o *MediaSourceInfo) GetOpenTokenOk() (*string, bool)`

GetOpenTokenOk returns a tuple with the OpenToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenToken

`func (o *MediaSourceInfo) SetOpenToken(v string)`

SetOpenToken sets OpenToken field to given value.

### HasOpenToken

`func (o *MediaSourceInfo) HasOpenToken() bool`

HasOpenToken returns a boolean if a field has been set.

### SetOpenTokenNil

`func (o *MediaSourceInfo) SetOpenTokenNil(b bool)`

 SetOpenTokenNil sets the value for OpenToken to be an explicit nil

### UnsetOpenToken
`func (o *MediaSourceInfo) UnsetOpenToken()`

UnsetOpenToken ensures that no value is present for OpenToken, not even an explicit nil
### GetRequiresClosing

`func (o *MediaSourceInfo) GetRequiresClosing() bool`

GetRequiresClosing returns the RequiresClosing field if non-nil, zero value otherwise.

### GetRequiresClosingOk

`func (o *MediaSourceInfo) GetRequiresClosingOk() (*bool, bool)`

GetRequiresClosingOk returns a tuple with the RequiresClosing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresClosing

`func (o *MediaSourceInfo) SetRequiresClosing(v bool)`

SetRequiresClosing sets RequiresClosing field to given value.

### HasRequiresClosing

`func (o *MediaSourceInfo) HasRequiresClosing() bool`

HasRequiresClosing returns a boolean if a field has been set.

### GetLiveStreamId

`func (o *MediaSourceInfo) GetLiveStreamId() string`

GetLiveStreamId returns the LiveStreamId field if non-nil, zero value otherwise.

### GetLiveStreamIdOk

`func (o *MediaSourceInfo) GetLiveStreamIdOk() (*string, bool)`

GetLiveStreamIdOk returns a tuple with the LiveStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveStreamId

`func (o *MediaSourceInfo) SetLiveStreamId(v string)`

SetLiveStreamId sets LiveStreamId field to given value.

### HasLiveStreamId

`func (o *MediaSourceInfo) HasLiveStreamId() bool`

HasLiveStreamId returns a boolean if a field has been set.

### SetLiveStreamIdNil

`func (o *MediaSourceInfo) SetLiveStreamIdNil(b bool)`

 SetLiveStreamIdNil sets the value for LiveStreamId to be an explicit nil

### UnsetLiveStreamId
`func (o *MediaSourceInfo) UnsetLiveStreamId()`

UnsetLiveStreamId ensures that no value is present for LiveStreamId, not even an explicit nil
### GetBufferMs

`func (o *MediaSourceInfo) GetBufferMs() int32`

GetBufferMs returns the BufferMs field if non-nil, zero value otherwise.

### GetBufferMsOk

`func (o *MediaSourceInfo) GetBufferMsOk() (*int32, bool)`

GetBufferMsOk returns a tuple with the BufferMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferMs

`func (o *MediaSourceInfo) SetBufferMs(v int32)`

SetBufferMs sets BufferMs field to given value.

### HasBufferMs

`func (o *MediaSourceInfo) HasBufferMs() bool`

HasBufferMs returns a boolean if a field has been set.

### SetBufferMsNil

`func (o *MediaSourceInfo) SetBufferMsNil(b bool)`

 SetBufferMsNil sets the value for BufferMs to be an explicit nil

### UnsetBufferMs
`func (o *MediaSourceInfo) UnsetBufferMs()`

UnsetBufferMs ensures that no value is present for BufferMs, not even an explicit nil
### GetRequiresLooping

`func (o *MediaSourceInfo) GetRequiresLooping() bool`

GetRequiresLooping returns the RequiresLooping field if non-nil, zero value otherwise.

### GetRequiresLoopingOk

`func (o *MediaSourceInfo) GetRequiresLoopingOk() (*bool, bool)`

GetRequiresLoopingOk returns a tuple with the RequiresLooping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresLooping

`func (o *MediaSourceInfo) SetRequiresLooping(v bool)`

SetRequiresLooping sets RequiresLooping field to given value.

### HasRequiresLooping

`func (o *MediaSourceInfo) HasRequiresLooping() bool`

HasRequiresLooping returns a boolean if a field has been set.

### GetSupportsProbing

`func (o *MediaSourceInfo) GetSupportsProbing() bool`

GetSupportsProbing returns the SupportsProbing field if non-nil, zero value otherwise.

### GetSupportsProbingOk

`func (o *MediaSourceInfo) GetSupportsProbingOk() (*bool, bool)`

GetSupportsProbingOk returns a tuple with the SupportsProbing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsProbing

`func (o *MediaSourceInfo) SetSupportsProbing(v bool)`

SetSupportsProbing sets SupportsProbing field to given value.

### HasSupportsProbing

`func (o *MediaSourceInfo) HasSupportsProbing() bool`

HasSupportsProbing returns a boolean if a field has been set.

### GetVideoType

`func (o *MediaSourceInfo) GetVideoType() VideoType`

GetVideoType returns the VideoType field if non-nil, zero value otherwise.

### GetVideoTypeOk

`func (o *MediaSourceInfo) GetVideoTypeOk() (*VideoType, bool)`

GetVideoTypeOk returns a tuple with the VideoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoType

`func (o *MediaSourceInfo) SetVideoType(v VideoType)`

SetVideoType sets VideoType field to given value.

### HasVideoType

`func (o *MediaSourceInfo) HasVideoType() bool`

HasVideoType returns a boolean if a field has been set.

### SetVideoTypeNil

`func (o *MediaSourceInfo) SetVideoTypeNil(b bool)`

 SetVideoTypeNil sets the value for VideoType to be an explicit nil

### UnsetVideoType
`func (o *MediaSourceInfo) UnsetVideoType()`

UnsetVideoType ensures that no value is present for VideoType, not even an explicit nil
### GetIsoType

`func (o *MediaSourceInfo) GetIsoType() IsoType`

GetIsoType returns the IsoType field if non-nil, zero value otherwise.

### GetIsoTypeOk

`func (o *MediaSourceInfo) GetIsoTypeOk() (*IsoType, bool)`

GetIsoTypeOk returns a tuple with the IsoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoType

`func (o *MediaSourceInfo) SetIsoType(v IsoType)`

SetIsoType sets IsoType field to given value.

### HasIsoType

`func (o *MediaSourceInfo) HasIsoType() bool`

HasIsoType returns a boolean if a field has been set.

### SetIsoTypeNil

`func (o *MediaSourceInfo) SetIsoTypeNil(b bool)`

 SetIsoTypeNil sets the value for IsoType to be an explicit nil

### UnsetIsoType
`func (o *MediaSourceInfo) UnsetIsoType()`

UnsetIsoType ensures that no value is present for IsoType, not even an explicit nil
### GetVideo3DFormat

`func (o *MediaSourceInfo) GetVideo3DFormat() Video3DFormat`

GetVideo3DFormat returns the Video3DFormat field if non-nil, zero value otherwise.

### GetVideo3DFormatOk

`func (o *MediaSourceInfo) GetVideo3DFormatOk() (*Video3DFormat, bool)`

GetVideo3DFormatOk returns a tuple with the Video3DFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo3DFormat

`func (o *MediaSourceInfo) SetVideo3DFormat(v Video3DFormat)`

SetVideo3DFormat sets Video3DFormat field to given value.

### HasVideo3DFormat

`func (o *MediaSourceInfo) HasVideo3DFormat() bool`

HasVideo3DFormat returns a boolean if a field has been set.

### SetVideo3DFormatNil

`func (o *MediaSourceInfo) SetVideo3DFormatNil(b bool)`

 SetVideo3DFormatNil sets the value for Video3DFormat to be an explicit nil

### UnsetVideo3DFormat
`func (o *MediaSourceInfo) UnsetVideo3DFormat()`

UnsetVideo3DFormat ensures that no value is present for Video3DFormat, not even an explicit nil
### GetMediaStreams

`func (o *MediaSourceInfo) GetMediaStreams() []MediaStream`

GetMediaStreams returns the MediaStreams field if non-nil, zero value otherwise.

### GetMediaStreamsOk

`func (o *MediaSourceInfo) GetMediaStreamsOk() (*[]MediaStream, bool)`

GetMediaStreamsOk returns a tuple with the MediaStreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaStreams

`func (o *MediaSourceInfo) SetMediaStreams(v []MediaStream)`

SetMediaStreams sets MediaStreams field to given value.

### HasMediaStreams

`func (o *MediaSourceInfo) HasMediaStreams() bool`

HasMediaStreams returns a boolean if a field has been set.

### SetMediaStreamsNil

`func (o *MediaSourceInfo) SetMediaStreamsNil(b bool)`

 SetMediaStreamsNil sets the value for MediaStreams to be an explicit nil

### UnsetMediaStreams
`func (o *MediaSourceInfo) UnsetMediaStreams()`

UnsetMediaStreams ensures that no value is present for MediaStreams, not even an explicit nil
### GetMediaAttachments

`func (o *MediaSourceInfo) GetMediaAttachments() []MediaAttachment`

GetMediaAttachments returns the MediaAttachments field if non-nil, zero value otherwise.

### GetMediaAttachmentsOk

`func (o *MediaSourceInfo) GetMediaAttachmentsOk() (*[]MediaAttachment, bool)`

GetMediaAttachmentsOk returns a tuple with the MediaAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaAttachments

`func (o *MediaSourceInfo) SetMediaAttachments(v []MediaAttachment)`

SetMediaAttachments sets MediaAttachments field to given value.

### HasMediaAttachments

`func (o *MediaSourceInfo) HasMediaAttachments() bool`

HasMediaAttachments returns a boolean if a field has been set.

### SetMediaAttachmentsNil

`func (o *MediaSourceInfo) SetMediaAttachmentsNil(b bool)`

 SetMediaAttachmentsNil sets the value for MediaAttachments to be an explicit nil

### UnsetMediaAttachments
`func (o *MediaSourceInfo) UnsetMediaAttachments()`

UnsetMediaAttachments ensures that no value is present for MediaAttachments, not even an explicit nil
### GetFormats

`func (o *MediaSourceInfo) GetFormats() []string`

GetFormats returns the Formats field if non-nil, zero value otherwise.

### GetFormatsOk

`func (o *MediaSourceInfo) GetFormatsOk() (*[]string, bool)`

GetFormatsOk returns a tuple with the Formats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormats

`func (o *MediaSourceInfo) SetFormats(v []string)`

SetFormats sets Formats field to given value.

### HasFormats

`func (o *MediaSourceInfo) HasFormats() bool`

HasFormats returns a boolean if a field has been set.

### SetFormatsNil

`func (o *MediaSourceInfo) SetFormatsNil(b bool)`

 SetFormatsNil sets the value for Formats to be an explicit nil

### UnsetFormats
`func (o *MediaSourceInfo) UnsetFormats()`

UnsetFormats ensures that no value is present for Formats, not even an explicit nil
### GetBitrate

`func (o *MediaSourceInfo) GetBitrate() int32`

GetBitrate returns the Bitrate field if non-nil, zero value otherwise.

### GetBitrateOk

`func (o *MediaSourceInfo) GetBitrateOk() (*int32, bool)`

GetBitrateOk returns a tuple with the Bitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrate

`func (o *MediaSourceInfo) SetBitrate(v int32)`

SetBitrate sets Bitrate field to given value.

### HasBitrate

`func (o *MediaSourceInfo) HasBitrate() bool`

HasBitrate returns a boolean if a field has been set.

### SetBitrateNil

`func (o *MediaSourceInfo) SetBitrateNil(b bool)`

 SetBitrateNil sets the value for Bitrate to be an explicit nil

### UnsetBitrate
`func (o *MediaSourceInfo) UnsetBitrate()`

UnsetBitrate ensures that no value is present for Bitrate, not even an explicit nil
### GetTimestamp

`func (o *MediaSourceInfo) GetTimestamp() TransportStreamTimestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MediaSourceInfo) GetTimestampOk() (*TransportStreamTimestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MediaSourceInfo) SetTimestamp(v TransportStreamTimestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MediaSourceInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *MediaSourceInfo) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *MediaSourceInfo) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetRequiredHttpHeaders

`func (o *MediaSourceInfo) GetRequiredHttpHeaders() map[string]string`

GetRequiredHttpHeaders returns the RequiredHttpHeaders field if non-nil, zero value otherwise.

### GetRequiredHttpHeadersOk

`func (o *MediaSourceInfo) GetRequiredHttpHeadersOk() (*map[string]string, bool)`

GetRequiredHttpHeadersOk returns a tuple with the RequiredHttpHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredHttpHeaders

`func (o *MediaSourceInfo) SetRequiredHttpHeaders(v map[string]string)`

SetRequiredHttpHeaders sets RequiredHttpHeaders field to given value.

### HasRequiredHttpHeaders

`func (o *MediaSourceInfo) HasRequiredHttpHeaders() bool`

HasRequiredHttpHeaders returns a boolean if a field has been set.

### SetRequiredHttpHeadersNil

`func (o *MediaSourceInfo) SetRequiredHttpHeadersNil(b bool)`

 SetRequiredHttpHeadersNil sets the value for RequiredHttpHeaders to be an explicit nil

### UnsetRequiredHttpHeaders
`func (o *MediaSourceInfo) UnsetRequiredHttpHeaders()`

UnsetRequiredHttpHeaders ensures that no value is present for RequiredHttpHeaders, not even an explicit nil
### GetTranscodingUrl

`func (o *MediaSourceInfo) GetTranscodingUrl() string`

GetTranscodingUrl returns the TranscodingUrl field if non-nil, zero value otherwise.

### GetTranscodingUrlOk

`func (o *MediaSourceInfo) GetTranscodingUrlOk() (*string, bool)`

GetTranscodingUrlOk returns a tuple with the TranscodingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingUrl

`func (o *MediaSourceInfo) SetTranscodingUrl(v string)`

SetTranscodingUrl sets TranscodingUrl field to given value.

### HasTranscodingUrl

`func (o *MediaSourceInfo) HasTranscodingUrl() bool`

HasTranscodingUrl returns a boolean if a field has been set.

### SetTranscodingUrlNil

`func (o *MediaSourceInfo) SetTranscodingUrlNil(b bool)`

 SetTranscodingUrlNil sets the value for TranscodingUrl to be an explicit nil

### UnsetTranscodingUrl
`func (o *MediaSourceInfo) UnsetTranscodingUrl()`

UnsetTranscodingUrl ensures that no value is present for TranscodingUrl, not even an explicit nil
### GetTranscodingSubProtocol

`func (o *MediaSourceInfo) GetTranscodingSubProtocol() string`

GetTranscodingSubProtocol returns the TranscodingSubProtocol field if non-nil, zero value otherwise.

### GetTranscodingSubProtocolOk

`func (o *MediaSourceInfo) GetTranscodingSubProtocolOk() (*string, bool)`

GetTranscodingSubProtocolOk returns a tuple with the TranscodingSubProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingSubProtocol

`func (o *MediaSourceInfo) SetTranscodingSubProtocol(v string)`

SetTranscodingSubProtocol sets TranscodingSubProtocol field to given value.

### HasTranscodingSubProtocol

`func (o *MediaSourceInfo) HasTranscodingSubProtocol() bool`

HasTranscodingSubProtocol returns a boolean if a field has been set.

### SetTranscodingSubProtocolNil

`func (o *MediaSourceInfo) SetTranscodingSubProtocolNil(b bool)`

 SetTranscodingSubProtocolNil sets the value for TranscodingSubProtocol to be an explicit nil

### UnsetTranscodingSubProtocol
`func (o *MediaSourceInfo) UnsetTranscodingSubProtocol()`

UnsetTranscodingSubProtocol ensures that no value is present for TranscodingSubProtocol, not even an explicit nil
### GetTranscodingContainer

`func (o *MediaSourceInfo) GetTranscodingContainer() string`

GetTranscodingContainer returns the TranscodingContainer field if non-nil, zero value otherwise.

### GetTranscodingContainerOk

`func (o *MediaSourceInfo) GetTranscodingContainerOk() (*string, bool)`

GetTranscodingContainerOk returns a tuple with the TranscodingContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingContainer

`func (o *MediaSourceInfo) SetTranscodingContainer(v string)`

SetTranscodingContainer sets TranscodingContainer field to given value.

### HasTranscodingContainer

`func (o *MediaSourceInfo) HasTranscodingContainer() bool`

HasTranscodingContainer returns a boolean if a field has been set.

### SetTranscodingContainerNil

`func (o *MediaSourceInfo) SetTranscodingContainerNil(b bool)`

 SetTranscodingContainerNil sets the value for TranscodingContainer to be an explicit nil

### UnsetTranscodingContainer
`func (o *MediaSourceInfo) UnsetTranscodingContainer()`

UnsetTranscodingContainer ensures that no value is present for TranscodingContainer, not even an explicit nil
### GetAnalyzeDurationMs

`func (o *MediaSourceInfo) GetAnalyzeDurationMs() int32`

GetAnalyzeDurationMs returns the AnalyzeDurationMs field if non-nil, zero value otherwise.

### GetAnalyzeDurationMsOk

`func (o *MediaSourceInfo) GetAnalyzeDurationMsOk() (*int32, bool)`

GetAnalyzeDurationMsOk returns a tuple with the AnalyzeDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzeDurationMs

`func (o *MediaSourceInfo) SetAnalyzeDurationMs(v int32)`

SetAnalyzeDurationMs sets AnalyzeDurationMs field to given value.

### HasAnalyzeDurationMs

`func (o *MediaSourceInfo) HasAnalyzeDurationMs() bool`

HasAnalyzeDurationMs returns a boolean if a field has been set.

### SetAnalyzeDurationMsNil

`func (o *MediaSourceInfo) SetAnalyzeDurationMsNil(b bool)`

 SetAnalyzeDurationMsNil sets the value for AnalyzeDurationMs to be an explicit nil

### UnsetAnalyzeDurationMs
`func (o *MediaSourceInfo) UnsetAnalyzeDurationMs()`

UnsetAnalyzeDurationMs ensures that no value is present for AnalyzeDurationMs, not even an explicit nil
### GetDefaultAudioStreamIndex

`func (o *MediaSourceInfo) GetDefaultAudioStreamIndex() int32`

GetDefaultAudioStreamIndex returns the DefaultAudioStreamIndex field if non-nil, zero value otherwise.

### GetDefaultAudioStreamIndexOk

`func (o *MediaSourceInfo) GetDefaultAudioStreamIndexOk() (*int32, bool)`

GetDefaultAudioStreamIndexOk returns a tuple with the DefaultAudioStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAudioStreamIndex

`func (o *MediaSourceInfo) SetDefaultAudioStreamIndex(v int32)`

SetDefaultAudioStreamIndex sets DefaultAudioStreamIndex field to given value.

### HasDefaultAudioStreamIndex

`func (o *MediaSourceInfo) HasDefaultAudioStreamIndex() bool`

HasDefaultAudioStreamIndex returns a boolean if a field has been set.

### SetDefaultAudioStreamIndexNil

`func (o *MediaSourceInfo) SetDefaultAudioStreamIndexNil(b bool)`

 SetDefaultAudioStreamIndexNil sets the value for DefaultAudioStreamIndex to be an explicit nil

### UnsetDefaultAudioStreamIndex
`func (o *MediaSourceInfo) UnsetDefaultAudioStreamIndex()`

UnsetDefaultAudioStreamIndex ensures that no value is present for DefaultAudioStreamIndex, not even an explicit nil
### GetDefaultSubtitleStreamIndex

`func (o *MediaSourceInfo) GetDefaultSubtitleStreamIndex() int32`

GetDefaultSubtitleStreamIndex returns the DefaultSubtitleStreamIndex field if non-nil, zero value otherwise.

### GetDefaultSubtitleStreamIndexOk

`func (o *MediaSourceInfo) GetDefaultSubtitleStreamIndexOk() (*int32, bool)`

GetDefaultSubtitleStreamIndexOk returns a tuple with the DefaultSubtitleStreamIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubtitleStreamIndex

`func (o *MediaSourceInfo) SetDefaultSubtitleStreamIndex(v int32)`

SetDefaultSubtitleStreamIndex sets DefaultSubtitleStreamIndex field to given value.

### HasDefaultSubtitleStreamIndex

`func (o *MediaSourceInfo) HasDefaultSubtitleStreamIndex() bool`

HasDefaultSubtitleStreamIndex returns a boolean if a field has been set.

### SetDefaultSubtitleStreamIndexNil

`func (o *MediaSourceInfo) SetDefaultSubtitleStreamIndexNil(b bool)`

 SetDefaultSubtitleStreamIndexNil sets the value for DefaultSubtitleStreamIndex to be an explicit nil

### UnsetDefaultSubtitleStreamIndex
`func (o *MediaSourceInfo) UnsetDefaultSubtitleStreamIndex()`

UnsetDefaultSubtitleStreamIndex ensures that no value is present for DefaultSubtitleStreamIndex, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


