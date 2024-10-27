# SessionInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayState** | Pointer to [**NullablePlayerStateInfo**](PlayerStateInfo.md) | Gets or sets the play state. | [optional] 
**AdditionalUsers** | Pointer to [**[]SessionUserInfo**](SessionUserInfo.md) | Gets or sets the additional users. | [optional] 
**Capabilities** | Pointer to [**NullableClientCapabilitiesDto**](ClientCapabilitiesDto.md) | Gets or sets the client capabilities. | [optional] 
**RemoteEndPoint** | Pointer to **NullableString** | Gets or sets the remote end point. | [optional] 
**PlayableMediaTypes** | Pointer to [**[]MediaType**](MediaType.md) | Gets or sets the playable media types. | [optional] 
**Id** | Pointer to **NullableString** | Gets or sets the id. | [optional] 
**UserId** | Pointer to **string** | Gets or sets the user id. | [optional] 
**UserName** | Pointer to **NullableString** | Gets or sets the username. | [optional] 
**Client** | Pointer to **NullableString** | Gets or sets the type of the client. | [optional] 
**LastActivityDate** | Pointer to **time.Time** | Gets or sets the last activity date. | [optional] 
**LastPlaybackCheckIn** | Pointer to **time.Time** | Gets or sets the last playback check in. | [optional] 
**LastPausedDate** | Pointer to **NullableTime** | Gets or sets the last paused date. | [optional] 
**DeviceName** | Pointer to **NullableString** | Gets or sets the name of the device. | [optional] 
**DeviceType** | Pointer to **NullableString** | Gets or sets the type of the device. | [optional] 
**NowPlayingItem** | Pointer to [**NullableBaseItemDto**](BaseItemDto.md) | Gets or sets the now playing item. | [optional] 
**NowViewingItem** | Pointer to [**NullableBaseItemDto**](BaseItemDto.md) | Gets or sets the now viewing item. | [optional] 
**DeviceId** | Pointer to **NullableString** | Gets or sets the device id. | [optional] 
**ApplicationVersion** | Pointer to **NullableString** | Gets or sets the application version. | [optional] 
**TranscodingInfo** | Pointer to [**NullableTranscodingInfo**](TranscodingInfo.md) | Gets or sets the transcoding info. | [optional] 
**IsActive** | Pointer to **bool** | Gets or sets a value indicating whether this session is active. | [optional] 
**SupportsMediaControl** | Pointer to **bool** | Gets or sets a value indicating whether the session supports media control. | [optional] 
**SupportsRemoteControl** | Pointer to **bool** | Gets or sets a value indicating whether the session supports remote control. | [optional] 
**NowPlayingQueue** | Pointer to [**[]QueueItem**](QueueItem.md) | Gets or sets the now playing queue. | [optional] 
**NowPlayingQueueFullItems** | Pointer to [**[]BaseItemDto**](BaseItemDto.md) | Gets or sets the now playing queue full items. | [optional] 
**HasCustomDeviceName** | Pointer to **bool** | Gets or sets a value indicating whether the session has a custom device name. | [optional] 
**PlaylistItemId** | Pointer to **NullableString** | Gets or sets the playlist item id. | [optional] 
**ServerId** | Pointer to **NullableString** | Gets or sets the server id. | [optional] 
**UserPrimaryImageTag** | Pointer to **NullableString** | Gets or sets the user primary image tag. | [optional] 
**SupportedCommands** | Pointer to [**[]GeneralCommandType**](GeneralCommandType.md) | Gets or sets the supported commands. | [optional] 

## Methods

### NewSessionInfoDto

`func NewSessionInfoDto() *SessionInfoDto`

NewSessionInfoDto instantiates a new SessionInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionInfoDtoWithDefaults

`func NewSessionInfoDtoWithDefaults() *SessionInfoDto`

NewSessionInfoDtoWithDefaults instantiates a new SessionInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayState

`func (o *SessionInfoDto) GetPlayState() PlayerStateInfo`

GetPlayState returns the PlayState field if non-nil, zero value otherwise.

### GetPlayStateOk

`func (o *SessionInfoDto) GetPlayStateOk() (*PlayerStateInfo, bool)`

GetPlayStateOk returns a tuple with the PlayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayState

`func (o *SessionInfoDto) SetPlayState(v PlayerStateInfo)`

SetPlayState sets PlayState field to given value.

### HasPlayState

`func (o *SessionInfoDto) HasPlayState() bool`

HasPlayState returns a boolean if a field has been set.

### SetPlayStateNil

`func (o *SessionInfoDto) SetPlayStateNil(b bool)`

 SetPlayStateNil sets the value for PlayState to be an explicit nil

### UnsetPlayState
`func (o *SessionInfoDto) UnsetPlayState()`

UnsetPlayState ensures that no value is present for PlayState, not even an explicit nil
### GetAdditionalUsers

`func (o *SessionInfoDto) GetAdditionalUsers() []SessionUserInfo`

GetAdditionalUsers returns the AdditionalUsers field if non-nil, zero value otherwise.

### GetAdditionalUsersOk

`func (o *SessionInfoDto) GetAdditionalUsersOk() (*[]SessionUserInfo, bool)`

GetAdditionalUsersOk returns a tuple with the AdditionalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUsers

`func (o *SessionInfoDto) SetAdditionalUsers(v []SessionUserInfo)`

SetAdditionalUsers sets AdditionalUsers field to given value.

### HasAdditionalUsers

`func (o *SessionInfoDto) HasAdditionalUsers() bool`

HasAdditionalUsers returns a boolean if a field has been set.

### SetAdditionalUsersNil

`func (o *SessionInfoDto) SetAdditionalUsersNil(b bool)`

 SetAdditionalUsersNil sets the value for AdditionalUsers to be an explicit nil

### UnsetAdditionalUsers
`func (o *SessionInfoDto) UnsetAdditionalUsers()`

UnsetAdditionalUsers ensures that no value is present for AdditionalUsers, not even an explicit nil
### GetCapabilities

`func (o *SessionInfoDto) GetCapabilities() ClientCapabilitiesDto`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *SessionInfoDto) GetCapabilitiesOk() (*ClientCapabilitiesDto, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *SessionInfoDto) SetCapabilities(v ClientCapabilitiesDto)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *SessionInfoDto) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *SessionInfoDto) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *SessionInfoDto) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetRemoteEndPoint

`func (o *SessionInfoDto) GetRemoteEndPoint() string`

GetRemoteEndPoint returns the RemoteEndPoint field if non-nil, zero value otherwise.

### GetRemoteEndPointOk

`func (o *SessionInfoDto) GetRemoteEndPointOk() (*string, bool)`

GetRemoteEndPointOk returns a tuple with the RemoteEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndPoint

`func (o *SessionInfoDto) SetRemoteEndPoint(v string)`

SetRemoteEndPoint sets RemoteEndPoint field to given value.

### HasRemoteEndPoint

`func (o *SessionInfoDto) HasRemoteEndPoint() bool`

HasRemoteEndPoint returns a boolean if a field has been set.

### SetRemoteEndPointNil

`func (o *SessionInfoDto) SetRemoteEndPointNil(b bool)`

 SetRemoteEndPointNil sets the value for RemoteEndPoint to be an explicit nil

### UnsetRemoteEndPoint
`func (o *SessionInfoDto) UnsetRemoteEndPoint()`

UnsetRemoteEndPoint ensures that no value is present for RemoteEndPoint, not even an explicit nil
### GetPlayableMediaTypes

`func (o *SessionInfoDto) GetPlayableMediaTypes() []MediaType`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *SessionInfoDto) GetPlayableMediaTypesOk() (*[]MediaType, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *SessionInfoDto) SetPlayableMediaTypes(v []MediaType)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *SessionInfoDto) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### GetId

`func (o *SessionInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SessionInfoDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SessionInfoDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUserId

`func (o *SessionInfoDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SessionInfoDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SessionInfoDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SessionInfoDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *SessionInfoDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SessionInfoDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SessionInfoDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SessionInfoDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *SessionInfoDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *SessionInfoDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetClient

`func (o *SessionInfoDto) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *SessionInfoDto) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *SessionInfoDto) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *SessionInfoDto) HasClient() bool`

HasClient returns a boolean if a field has been set.

### SetClientNil

`func (o *SessionInfoDto) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *SessionInfoDto) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil
### GetLastActivityDate

`func (o *SessionInfoDto) GetLastActivityDate() time.Time`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *SessionInfoDto) GetLastActivityDateOk() (*time.Time, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *SessionInfoDto) SetLastActivityDate(v time.Time)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *SessionInfoDto) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### GetLastPlaybackCheckIn

`func (o *SessionInfoDto) GetLastPlaybackCheckIn() time.Time`

GetLastPlaybackCheckIn returns the LastPlaybackCheckIn field if non-nil, zero value otherwise.

### GetLastPlaybackCheckInOk

`func (o *SessionInfoDto) GetLastPlaybackCheckInOk() (*time.Time, bool)`

GetLastPlaybackCheckInOk returns a tuple with the LastPlaybackCheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlaybackCheckIn

`func (o *SessionInfoDto) SetLastPlaybackCheckIn(v time.Time)`

SetLastPlaybackCheckIn sets LastPlaybackCheckIn field to given value.

### HasLastPlaybackCheckIn

`func (o *SessionInfoDto) HasLastPlaybackCheckIn() bool`

HasLastPlaybackCheckIn returns a boolean if a field has been set.

### GetLastPausedDate

`func (o *SessionInfoDto) GetLastPausedDate() time.Time`

GetLastPausedDate returns the LastPausedDate field if non-nil, zero value otherwise.

### GetLastPausedDateOk

`func (o *SessionInfoDto) GetLastPausedDateOk() (*time.Time, bool)`

GetLastPausedDateOk returns a tuple with the LastPausedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPausedDate

`func (o *SessionInfoDto) SetLastPausedDate(v time.Time)`

SetLastPausedDate sets LastPausedDate field to given value.

### HasLastPausedDate

`func (o *SessionInfoDto) HasLastPausedDate() bool`

HasLastPausedDate returns a boolean if a field has been set.

### SetLastPausedDateNil

`func (o *SessionInfoDto) SetLastPausedDateNil(b bool)`

 SetLastPausedDateNil sets the value for LastPausedDate to be an explicit nil

### UnsetLastPausedDate
`func (o *SessionInfoDto) UnsetLastPausedDate()`

UnsetLastPausedDate ensures that no value is present for LastPausedDate, not even an explicit nil
### GetDeviceName

`func (o *SessionInfoDto) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SessionInfoDto) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SessionInfoDto) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *SessionInfoDto) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *SessionInfoDto) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *SessionInfoDto) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetDeviceType

`func (o *SessionInfoDto) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *SessionInfoDto) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *SessionInfoDto) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *SessionInfoDto) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *SessionInfoDto) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *SessionInfoDto) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetNowPlayingItem

`func (o *SessionInfoDto) GetNowPlayingItem() BaseItemDto`

GetNowPlayingItem returns the NowPlayingItem field if non-nil, zero value otherwise.

### GetNowPlayingItemOk

`func (o *SessionInfoDto) GetNowPlayingItemOk() (*BaseItemDto, bool)`

GetNowPlayingItemOk returns a tuple with the NowPlayingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingItem

`func (o *SessionInfoDto) SetNowPlayingItem(v BaseItemDto)`

SetNowPlayingItem sets NowPlayingItem field to given value.

### HasNowPlayingItem

`func (o *SessionInfoDto) HasNowPlayingItem() bool`

HasNowPlayingItem returns a boolean if a field has been set.

### SetNowPlayingItemNil

`func (o *SessionInfoDto) SetNowPlayingItemNil(b bool)`

 SetNowPlayingItemNil sets the value for NowPlayingItem to be an explicit nil

### UnsetNowPlayingItem
`func (o *SessionInfoDto) UnsetNowPlayingItem()`

UnsetNowPlayingItem ensures that no value is present for NowPlayingItem, not even an explicit nil
### GetNowViewingItem

`func (o *SessionInfoDto) GetNowViewingItem() BaseItemDto`

GetNowViewingItem returns the NowViewingItem field if non-nil, zero value otherwise.

### GetNowViewingItemOk

`func (o *SessionInfoDto) GetNowViewingItemOk() (*BaseItemDto, bool)`

GetNowViewingItemOk returns a tuple with the NowViewingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowViewingItem

`func (o *SessionInfoDto) SetNowViewingItem(v BaseItemDto)`

SetNowViewingItem sets NowViewingItem field to given value.

### HasNowViewingItem

`func (o *SessionInfoDto) HasNowViewingItem() bool`

HasNowViewingItem returns a boolean if a field has been set.

### SetNowViewingItemNil

`func (o *SessionInfoDto) SetNowViewingItemNil(b bool)`

 SetNowViewingItemNil sets the value for NowViewingItem to be an explicit nil

### UnsetNowViewingItem
`func (o *SessionInfoDto) UnsetNowViewingItem()`

UnsetNowViewingItem ensures that no value is present for NowViewingItem, not even an explicit nil
### GetDeviceId

`func (o *SessionInfoDto) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SessionInfoDto) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SessionInfoDto) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SessionInfoDto) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *SessionInfoDto) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *SessionInfoDto) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetApplicationVersion

`func (o *SessionInfoDto) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *SessionInfoDto) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *SessionInfoDto) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *SessionInfoDto) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### SetApplicationVersionNil

`func (o *SessionInfoDto) SetApplicationVersionNil(b bool)`

 SetApplicationVersionNil sets the value for ApplicationVersion to be an explicit nil

### UnsetApplicationVersion
`func (o *SessionInfoDto) UnsetApplicationVersion()`

UnsetApplicationVersion ensures that no value is present for ApplicationVersion, not even an explicit nil
### GetTranscodingInfo

`func (o *SessionInfoDto) GetTranscodingInfo() TranscodingInfo`

GetTranscodingInfo returns the TranscodingInfo field if non-nil, zero value otherwise.

### GetTranscodingInfoOk

`func (o *SessionInfoDto) GetTranscodingInfoOk() (*TranscodingInfo, bool)`

GetTranscodingInfoOk returns a tuple with the TranscodingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingInfo

`func (o *SessionInfoDto) SetTranscodingInfo(v TranscodingInfo)`

SetTranscodingInfo sets TranscodingInfo field to given value.

### HasTranscodingInfo

`func (o *SessionInfoDto) HasTranscodingInfo() bool`

HasTranscodingInfo returns a boolean if a field has been set.

### SetTranscodingInfoNil

`func (o *SessionInfoDto) SetTranscodingInfoNil(b bool)`

 SetTranscodingInfoNil sets the value for TranscodingInfo to be an explicit nil

### UnsetTranscodingInfo
`func (o *SessionInfoDto) UnsetTranscodingInfo()`

UnsetTranscodingInfo ensures that no value is present for TranscodingInfo, not even an explicit nil
### GetIsActive

`func (o *SessionInfoDto) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SessionInfoDto) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SessionInfoDto) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *SessionInfoDto) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetSupportsMediaControl

`func (o *SessionInfoDto) GetSupportsMediaControl() bool`

GetSupportsMediaControl returns the SupportsMediaControl field if non-nil, zero value otherwise.

### GetSupportsMediaControlOk

`func (o *SessionInfoDto) GetSupportsMediaControlOk() (*bool, bool)`

GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMediaControl

`func (o *SessionInfoDto) SetSupportsMediaControl(v bool)`

SetSupportsMediaControl sets SupportsMediaControl field to given value.

### HasSupportsMediaControl

`func (o *SessionInfoDto) HasSupportsMediaControl() bool`

HasSupportsMediaControl returns a boolean if a field has been set.

### GetSupportsRemoteControl

`func (o *SessionInfoDto) GetSupportsRemoteControl() bool`

GetSupportsRemoteControl returns the SupportsRemoteControl field if non-nil, zero value otherwise.

### GetSupportsRemoteControlOk

`func (o *SessionInfoDto) GetSupportsRemoteControlOk() (*bool, bool)`

GetSupportsRemoteControlOk returns a tuple with the SupportsRemoteControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRemoteControl

`func (o *SessionInfoDto) SetSupportsRemoteControl(v bool)`

SetSupportsRemoteControl sets SupportsRemoteControl field to given value.

### HasSupportsRemoteControl

`func (o *SessionInfoDto) HasSupportsRemoteControl() bool`

HasSupportsRemoteControl returns a boolean if a field has been set.

### GetNowPlayingQueue

`func (o *SessionInfoDto) GetNowPlayingQueue() []QueueItem`

GetNowPlayingQueue returns the NowPlayingQueue field if non-nil, zero value otherwise.

### GetNowPlayingQueueOk

`func (o *SessionInfoDto) GetNowPlayingQueueOk() (*[]QueueItem, bool)`

GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueue

`func (o *SessionInfoDto) SetNowPlayingQueue(v []QueueItem)`

SetNowPlayingQueue sets NowPlayingQueue field to given value.

### HasNowPlayingQueue

`func (o *SessionInfoDto) HasNowPlayingQueue() bool`

HasNowPlayingQueue returns a boolean if a field has been set.

### SetNowPlayingQueueNil

`func (o *SessionInfoDto) SetNowPlayingQueueNil(b bool)`

 SetNowPlayingQueueNil sets the value for NowPlayingQueue to be an explicit nil

### UnsetNowPlayingQueue
`func (o *SessionInfoDto) UnsetNowPlayingQueue()`

UnsetNowPlayingQueue ensures that no value is present for NowPlayingQueue, not even an explicit nil
### GetNowPlayingQueueFullItems

`func (o *SessionInfoDto) GetNowPlayingQueueFullItems() []BaseItemDto`

GetNowPlayingQueueFullItems returns the NowPlayingQueueFullItems field if non-nil, zero value otherwise.

### GetNowPlayingQueueFullItemsOk

`func (o *SessionInfoDto) GetNowPlayingQueueFullItemsOk() (*[]BaseItemDto, bool)`

GetNowPlayingQueueFullItemsOk returns a tuple with the NowPlayingQueueFullItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueueFullItems

`func (o *SessionInfoDto) SetNowPlayingQueueFullItems(v []BaseItemDto)`

SetNowPlayingQueueFullItems sets NowPlayingQueueFullItems field to given value.

### HasNowPlayingQueueFullItems

`func (o *SessionInfoDto) HasNowPlayingQueueFullItems() bool`

HasNowPlayingQueueFullItems returns a boolean if a field has been set.

### SetNowPlayingQueueFullItemsNil

`func (o *SessionInfoDto) SetNowPlayingQueueFullItemsNil(b bool)`

 SetNowPlayingQueueFullItemsNil sets the value for NowPlayingQueueFullItems to be an explicit nil

### UnsetNowPlayingQueueFullItems
`func (o *SessionInfoDto) UnsetNowPlayingQueueFullItems()`

UnsetNowPlayingQueueFullItems ensures that no value is present for NowPlayingQueueFullItems, not even an explicit nil
### GetHasCustomDeviceName

`func (o *SessionInfoDto) GetHasCustomDeviceName() bool`

GetHasCustomDeviceName returns the HasCustomDeviceName field if non-nil, zero value otherwise.

### GetHasCustomDeviceNameOk

`func (o *SessionInfoDto) GetHasCustomDeviceNameOk() (*bool, bool)`

GetHasCustomDeviceNameOk returns a tuple with the HasCustomDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCustomDeviceName

`func (o *SessionInfoDto) SetHasCustomDeviceName(v bool)`

SetHasCustomDeviceName sets HasCustomDeviceName field to given value.

### HasHasCustomDeviceName

`func (o *SessionInfoDto) HasHasCustomDeviceName() bool`

HasHasCustomDeviceName returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *SessionInfoDto) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *SessionInfoDto) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *SessionInfoDto) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *SessionInfoDto) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### SetPlaylistItemIdNil

`func (o *SessionInfoDto) SetPlaylistItemIdNil(b bool)`

 SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil

### UnsetPlaylistItemId
`func (o *SessionInfoDto) UnsetPlaylistItemId()`

UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil
### GetServerId

`func (o *SessionInfoDto) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *SessionInfoDto) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *SessionInfoDto) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *SessionInfoDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *SessionInfoDto) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *SessionInfoDto) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetUserPrimaryImageTag

`func (o *SessionInfoDto) GetUserPrimaryImageTag() string`

GetUserPrimaryImageTag returns the UserPrimaryImageTag field if non-nil, zero value otherwise.

### GetUserPrimaryImageTagOk

`func (o *SessionInfoDto) GetUserPrimaryImageTagOk() (*string, bool)`

GetUserPrimaryImageTagOk returns a tuple with the UserPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrimaryImageTag

`func (o *SessionInfoDto) SetUserPrimaryImageTag(v string)`

SetUserPrimaryImageTag sets UserPrimaryImageTag field to given value.

### HasUserPrimaryImageTag

`func (o *SessionInfoDto) HasUserPrimaryImageTag() bool`

HasUserPrimaryImageTag returns a boolean if a field has been set.

### SetUserPrimaryImageTagNil

`func (o *SessionInfoDto) SetUserPrimaryImageTagNil(b bool)`

 SetUserPrimaryImageTagNil sets the value for UserPrimaryImageTag to be an explicit nil

### UnsetUserPrimaryImageTag
`func (o *SessionInfoDto) UnsetUserPrimaryImageTag()`

UnsetUserPrimaryImageTag ensures that no value is present for UserPrimaryImageTag, not even an explicit nil
### GetSupportedCommands

`func (o *SessionInfoDto) GetSupportedCommands() []GeneralCommandType`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *SessionInfoDto) GetSupportedCommandsOk() (*[]GeneralCommandType, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *SessionInfoDto) SetSupportedCommands(v []GeneralCommandType)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *SessionInfoDto) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


