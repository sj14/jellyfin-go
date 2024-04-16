# PlaybackInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaSources** | Pointer to [**[]MediaSourceInfo**](MediaSourceInfo.md) | Gets or sets the media sources. | [optional] 
**PlaySessionId** | Pointer to **NullableString** | Gets or sets the play session identifier. | [optional] 
**ErrorCode** | Pointer to [**NullablePlaybackErrorCode**](PlaybackErrorCode.md) | Gets or sets the error code. | [optional] 

## Methods

### NewPlaybackInfoResponse

`func NewPlaybackInfoResponse() *PlaybackInfoResponse`

NewPlaybackInfoResponse instantiates a new PlaybackInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaybackInfoResponseWithDefaults

`func NewPlaybackInfoResponseWithDefaults() *PlaybackInfoResponse`

NewPlaybackInfoResponseWithDefaults instantiates a new PlaybackInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaSources

`func (o *PlaybackInfoResponse) GetMediaSources() []MediaSourceInfo`

GetMediaSources returns the MediaSources field if non-nil, zero value otherwise.

### GetMediaSourcesOk

`func (o *PlaybackInfoResponse) GetMediaSourcesOk() (*[]MediaSourceInfo, bool)`

GetMediaSourcesOk returns a tuple with the MediaSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSources

`func (o *PlaybackInfoResponse) SetMediaSources(v []MediaSourceInfo)`

SetMediaSources sets MediaSources field to given value.

### HasMediaSources

`func (o *PlaybackInfoResponse) HasMediaSources() bool`

HasMediaSources returns a boolean if a field has been set.

### GetPlaySessionId

`func (o *PlaybackInfoResponse) GetPlaySessionId() string`

GetPlaySessionId returns the PlaySessionId field if non-nil, zero value otherwise.

### GetPlaySessionIdOk

`func (o *PlaybackInfoResponse) GetPlaySessionIdOk() (*string, bool)`

GetPlaySessionIdOk returns a tuple with the PlaySessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaySessionId

`func (o *PlaybackInfoResponse) SetPlaySessionId(v string)`

SetPlaySessionId sets PlaySessionId field to given value.

### HasPlaySessionId

`func (o *PlaybackInfoResponse) HasPlaySessionId() bool`

HasPlaySessionId returns a boolean if a field has been set.

### SetPlaySessionIdNil

`func (o *PlaybackInfoResponse) SetPlaySessionIdNil(b bool)`

 SetPlaySessionIdNil sets the value for PlaySessionId to be an explicit nil

### UnsetPlaySessionId
`func (o *PlaybackInfoResponse) UnsetPlaySessionId()`

UnsetPlaySessionId ensures that no value is present for PlaySessionId, not even an explicit nil
### GetErrorCode

`func (o *PlaybackInfoResponse) GetErrorCode() PlaybackErrorCode`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *PlaybackInfoResponse) GetErrorCodeOk() (*PlaybackErrorCode, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *PlaybackInfoResponse) SetErrorCode(v PlaybackErrorCode)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *PlaybackInfoResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *PlaybackInfoResponse) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *PlaybackInfoResponse) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


