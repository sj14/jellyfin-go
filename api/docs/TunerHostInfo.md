# TunerHostInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**DeviceId** | Pointer to **NullableString** |  | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**ImportFavoritesOnly** | Pointer to **bool** |  | [optional] 
**AllowHWTranscoding** | Pointer to **bool** |  | [optional] 
**AllowFmp4TranscodingContainer** | Pointer to **bool** |  | [optional] 
**AllowStreamSharing** | Pointer to **bool** |  | [optional] 
**FallbackMaxStreamingBitrate** | Pointer to **int32** |  | [optional] 
**EnableStreamLooping** | Pointer to **bool** |  | [optional] 
**Source** | Pointer to **NullableString** |  | [optional] 
**TunerCount** | Pointer to **int32** |  | [optional] 
**UserAgent** | Pointer to **NullableString** |  | [optional] 
**IgnoreDts** | Pointer to **bool** |  | [optional] 

## Methods

### NewTunerHostInfo

`func NewTunerHostInfo() *TunerHostInfo`

NewTunerHostInfo instantiates a new TunerHostInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTunerHostInfoWithDefaults

`func NewTunerHostInfoWithDefaults() *TunerHostInfo`

NewTunerHostInfoWithDefaults instantiates a new TunerHostInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TunerHostInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TunerHostInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TunerHostInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TunerHostInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TunerHostInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TunerHostInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUrl

`func (o *TunerHostInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TunerHostInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TunerHostInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TunerHostInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TunerHostInfo) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TunerHostInfo) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetType

`func (o *TunerHostInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TunerHostInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TunerHostInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TunerHostInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TunerHostInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TunerHostInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDeviceId

`func (o *TunerHostInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *TunerHostInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *TunerHostInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *TunerHostInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *TunerHostInfo) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *TunerHostInfo) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetFriendlyName

`func (o *TunerHostInfo) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *TunerHostInfo) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *TunerHostInfo) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *TunerHostInfo) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *TunerHostInfo) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *TunerHostInfo) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetImportFavoritesOnly

`func (o *TunerHostInfo) GetImportFavoritesOnly() bool`

GetImportFavoritesOnly returns the ImportFavoritesOnly field if non-nil, zero value otherwise.

### GetImportFavoritesOnlyOk

`func (o *TunerHostInfo) GetImportFavoritesOnlyOk() (*bool, bool)`

GetImportFavoritesOnlyOk returns a tuple with the ImportFavoritesOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportFavoritesOnly

`func (o *TunerHostInfo) SetImportFavoritesOnly(v bool)`

SetImportFavoritesOnly sets ImportFavoritesOnly field to given value.

### HasImportFavoritesOnly

`func (o *TunerHostInfo) HasImportFavoritesOnly() bool`

HasImportFavoritesOnly returns a boolean if a field has been set.

### GetAllowHWTranscoding

`func (o *TunerHostInfo) GetAllowHWTranscoding() bool`

GetAllowHWTranscoding returns the AllowHWTranscoding field if non-nil, zero value otherwise.

### GetAllowHWTranscodingOk

`func (o *TunerHostInfo) GetAllowHWTranscodingOk() (*bool, bool)`

GetAllowHWTranscodingOk returns a tuple with the AllowHWTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHWTranscoding

`func (o *TunerHostInfo) SetAllowHWTranscoding(v bool)`

SetAllowHWTranscoding sets AllowHWTranscoding field to given value.

### HasAllowHWTranscoding

`func (o *TunerHostInfo) HasAllowHWTranscoding() bool`

HasAllowHWTranscoding returns a boolean if a field has been set.

### GetAllowFmp4TranscodingContainer

`func (o *TunerHostInfo) GetAllowFmp4TranscodingContainer() bool`

GetAllowFmp4TranscodingContainer returns the AllowFmp4TranscodingContainer field if non-nil, zero value otherwise.

### GetAllowFmp4TranscodingContainerOk

`func (o *TunerHostInfo) GetAllowFmp4TranscodingContainerOk() (*bool, bool)`

GetAllowFmp4TranscodingContainerOk returns a tuple with the AllowFmp4TranscodingContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFmp4TranscodingContainer

`func (o *TunerHostInfo) SetAllowFmp4TranscodingContainer(v bool)`

SetAllowFmp4TranscodingContainer sets AllowFmp4TranscodingContainer field to given value.

### HasAllowFmp4TranscodingContainer

`func (o *TunerHostInfo) HasAllowFmp4TranscodingContainer() bool`

HasAllowFmp4TranscodingContainer returns a boolean if a field has been set.

### GetAllowStreamSharing

`func (o *TunerHostInfo) GetAllowStreamSharing() bool`

GetAllowStreamSharing returns the AllowStreamSharing field if non-nil, zero value otherwise.

### GetAllowStreamSharingOk

`func (o *TunerHostInfo) GetAllowStreamSharingOk() (*bool, bool)`

GetAllowStreamSharingOk returns a tuple with the AllowStreamSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStreamSharing

`func (o *TunerHostInfo) SetAllowStreamSharing(v bool)`

SetAllowStreamSharing sets AllowStreamSharing field to given value.

### HasAllowStreamSharing

`func (o *TunerHostInfo) HasAllowStreamSharing() bool`

HasAllowStreamSharing returns a boolean if a field has been set.

### GetFallbackMaxStreamingBitrate

`func (o *TunerHostInfo) GetFallbackMaxStreamingBitrate() int32`

GetFallbackMaxStreamingBitrate returns the FallbackMaxStreamingBitrate field if non-nil, zero value otherwise.

### GetFallbackMaxStreamingBitrateOk

`func (o *TunerHostInfo) GetFallbackMaxStreamingBitrateOk() (*int32, bool)`

GetFallbackMaxStreamingBitrateOk returns a tuple with the FallbackMaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackMaxStreamingBitrate

`func (o *TunerHostInfo) SetFallbackMaxStreamingBitrate(v int32)`

SetFallbackMaxStreamingBitrate sets FallbackMaxStreamingBitrate field to given value.

### HasFallbackMaxStreamingBitrate

`func (o *TunerHostInfo) HasFallbackMaxStreamingBitrate() bool`

HasFallbackMaxStreamingBitrate returns a boolean if a field has been set.

### GetEnableStreamLooping

`func (o *TunerHostInfo) GetEnableStreamLooping() bool`

GetEnableStreamLooping returns the EnableStreamLooping field if non-nil, zero value otherwise.

### GetEnableStreamLoopingOk

`func (o *TunerHostInfo) GetEnableStreamLoopingOk() (*bool, bool)`

GetEnableStreamLoopingOk returns a tuple with the EnableStreamLooping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStreamLooping

`func (o *TunerHostInfo) SetEnableStreamLooping(v bool)`

SetEnableStreamLooping sets EnableStreamLooping field to given value.

### HasEnableStreamLooping

`func (o *TunerHostInfo) HasEnableStreamLooping() bool`

HasEnableStreamLooping returns a boolean if a field has been set.

### GetSource

`func (o *TunerHostInfo) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TunerHostInfo) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TunerHostInfo) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TunerHostInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *TunerHostInfo) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *TunerHostInfo) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetTunerCount

`func (o *TunerHostInfo) GetTunerCount() int32`

GetTunerCount returns the TunerCount field if non-nil, zero value otherwise.

### GetTunerCountOk

`func (o *TunerHostInfo) GetTunerCountOk() (*int32, bool)`

GetTunerCountOk returns a tuple with the TunerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunerCount

`func (o *TunerHostInfo) SetTunerCount(v int32)`

SetTunerCount sets TunerCount field to given value.

### HasTunerCount

`func (o *TunerHostInfo) HasTunerCount() bool`

HasTunerCount returns a boolean if a field has been set.

### GetUserAgent

`func (o *TunerHostInfo) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *TunerHostInfo) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *TunerHostInfo) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *TunerHostInfo) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *TunerHostInfo) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *TunerHostInfo) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetIgnoreDts

`func (o *TunerHostInfo) GetIgnoreDts() bool`

GetIgnoreDts returns the IgnoreDts field if non-nil, zero value otherwise.

### GetIgnoreDtsOk

`func (o *TunerHostInfo) GetIgnoreDtsOk() (*bool, bool)`

GetIgnoreDtsOk returns a tuple with the IgnoreDts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDts

`func (o *TunerHostInfo) SetIgnoreDts(v bool)`

SetIgnoreDts sets IgnoreDts field to given value.

### HasIgnoreDts

`func (o *TunerHostInfo) HasIgnoreDts() bool`

HasIgnoreDts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


