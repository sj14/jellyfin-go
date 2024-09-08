# SessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayState** | Pointer to [**NullablePlayerStateInfo**](PlayerStateInfo.md) |  | [optional] 
**AdditionalUsers** | Pointer to [**[]SessionUserInfo**](SessionUserInfo.md) |  | [optional] 
**Capabilities** | Pointer to [**NullableClientCapabilities**](ClientCapabilities.md) |  | [optional] 
**RemoteEndPoint** | Pointer to **NullableString** | Gets or sets the remote end point. | [optional] 
**PlayableMediaTypes** | Pointer to [**[]MediaType**](MediaType.md) | Gets the playable media types. | [optional] [readonly] 
**Id** | Pointer to **NullableString** | Gets or sets the id. | [optional] 
**UserId** | Pointer to **string** | Gets or sets the user id. | [optional] 
**UserName** | Pointer to **NullableString** | Gets or sets the username. | [optional] 
**Client** | Pointer to **NullableString** | Gets or sets the type of the client. | [optional] 
**LastActivityDate** | Pointer to **time.Time** | Gets or sets the last activity date. | [optional] 
**LastPlaybackCheckIn** | Pointer to **time.Time** | Gets or sets the last playback check in. | [optional] 
**LastPausedDate** | Pointer to **NullableTime** | Gets or sets the last paused date. | [optional] 
**DeviceName** | Pointer to **NullableString** | Gets or sets the name of the device. | [optional] 
**DeviceType** | Pointer to **NullableString** | Gets or sets the type of the device. | [optional] 
**NowPlayingItem** | Pointer to [**NullableBaseItemDto**](BaseItemDto.md) | This is strictly used as a data transfer object from the api layer.  This holds information about a BaseItem in a format that is convenient for the client. | [optional] 
**NowViewingItem** | Pointer to [**NullableBaseItemDto**](BaseItemDto.md) | This is strictly used as a data transfer object from the api layer.  This holds information about a BaseItem in a format that is convenient for the client. | [optional] 
**DeviceId** | Pointer to **NullableString** | Gets or sets the device id. | [optional] 
**ApplicationVersion** | Pointer to **NullableString** | Gets or sets the application version. | [optional] 
**TranscodingInfo** | Pointer to [**NullableTranscodingInfo**](TranscodingInfo.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Gets a value indicating whether this instance is active. | [optional] [readonly] 
**SupportsMediaControl** | Pointer to **bool** |  | [optional] [readonly] 
**SupportsRemoteControl** | Pointer to **bool** |  | [optional] [readonly] 
**NowPlayingQueue** | Pointer to [**[]QueueItem**](QueueItem.md) |  | [optional] 
**NowPlayingQueueFullItems** | Pointer to [**[]BaseItemDto**](BaseItemDto.md) |  | [optional] 
**HasCustomDeviceName** | Pointer to **bool** |  | [optional] 
**PlaylistItemId** | Pointer to **NullableString** |  | [optional] 
**ServerId** | Pointer to **NullableString** |  | [optional] 
**UserPrimaryImageTag** | Pointer to **NullableString** |  | [optional] 
**SupportedCommands** | Pointer to [**[]GeneralCommandType**](GeneralCommandType.md) | Gets the supported commands. | [optional] [readonly] 

## Methods

### NewSessionInfo

`func NewSessionInfo() *SessionInfo`

NewSessionInfo instantiates a new SessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionInfoWithDefaults

`func NewSessionInfoWithDefaults() *SessionInfo`

NewSessionInfoWithDefaults instantiates a new SessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayState

`func (o *SessionInfo) GetPlayState() PlayerStateInfo`

GetPlayState returns the PlayState field if non-nil, zero value otherwise.

### GetPlayStateOk

`func (o *SessionInfo) GetPlayStateOk() (*PlayerStateInfo, bool)`

GetPlayStateOk returns a tuple with the PlayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayState

`func (o *SessionInfo) SetPlayState(v PlayerStateInfo)`

SetPlayState sets PlayState field to given value.

### HasPlayState

`func (o *SessionInfo) HasPlayState() bool`

HasPlayState returns a boolean if a field has been set.

### SetPlayStateNil

`func (o *SessionInfo) SetPlayStateNil(b bool)`

 SetPlayStateNil sets the value for PlayState to be an explicit nil

### UnsetPlayState
`func (o *SessionInfo) UnsetPlayState()`

UnsetPlayState ensures that no value is present for PlayState, not even an explicit nil
### GetAdditionalUsers

`func (o *SessionInfo) GetAdditionalUsers() []SessionUserInfo`

GetAdditionalUsers returns the AdditionalUsers field if non-nil, zero value otherwise.

### GetAdditionalUsersOk

`func (o *SessionInfo) GetAdditionalUsersOk() (*[]SessionUserInfo, bool)`

GetAdditionalUsersOk returns a tuple with the AdditionalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUsers

`func (o *SessionInfo) SetAdditionalUsers(v []SessionUserInfo)`

SetAdditionalUsers sets AdditionalUsers field to given value.

### HasAdditionalUsers

`func (o *SessionInfo) HasAdditionalUsers() bool`

HasAdditionalUsers returns a boolean if a field has been set.

### SetAdditionalUsersNil

`func (o *SessionInfo) SetAdditionalUsersNil(b bool)`

 SetAdditionalUsersNil sets the value for AdditionalUsers to be an explicit nil

### UnsetAdditionalUsers
`func (o *SessionInfo) UnsetAdditionalUsers()`

UnsetAdditionalUsers ensures that no value is present for AdditionalUsers, not even an explicit nil
### GetCapabilities

`func (o *SessionInfo) GetCapabilities() ClientCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *SessionInfo) GetCapabilitiesOk() (*ClientCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *SessionInfo) SetCapabilities(v ClientCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *SessionInfo) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *SessionInfo) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *SessionInfo) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetRemoteEndPoint

`func (o *SessionInfo) GetRemoteEndPoint() string`

GetRemoteEndPoint returns the RemoteEndPoint field if non-nil, zero value otherwise.

### GetRemoteEndPointOk

`func (o *SessionInfo) GetRemoteEndPointOk() (*string, bool)`

GetRemoteEndPointOk returns a tuple with the RemoteEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndPoint

`func (o *SessionInfo) SetRemoteEndPoint(v string)`

SetRemoteEndPoint sets RemoteEndPoint field to given value.

### HasRemoteEndPoint

`func (o *SessionInfo) HasRemoteEndPoint() bool`

HasRemoteEndPoint returns a boolean if a field has been set.

### SetRemoteEndPointNil

`func (o *SessionInfo) SetRemoteEndPointNil(b bool)`

 SetRemoteEndPointNil sets the value for RemoteEndPoint to be an explicit nil

### UnsetRemoteEndPoint
`func (o *SessionInfo) UnsetRemoteEndPoint()`

UnsetRemoteEndPoint ensures that no value is present for RemoteEndPoint, not even an explicit nil
### GetPlayableMediaTypes

`func (o *SessionInfo) GetPlayableMediaTypes() []MediaType`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *SessionInfo) GetPlayableMediaTypesOk() (*[]MediaType, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *SessionInfo) SetPlayableMediaTypes(v []MediaType)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *SessionInfo) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### SetPlayableMediaTypesNil

`func (o *SessionInfo) SetPlayableMediaTypesNil(b bool)`

 SetPlayableMediaTypesNil sets the value for PlayableMediaTypes to be an explicit nil

### UnsetPlayableMediaTypes
`func (o *SessionInfo) UnsetPlayableMediaTypes()`

UnsetPlayableMediaTypes ensures that no value is present for PlayableMediaTypes, not even an explicit nil
### GetId

`func (o *SessionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SessionInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SessionInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUserId

`func (o *SessionInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SessionInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SessionInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SessionInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *SessionInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SessionInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SessionInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SessionInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *SessionInfo) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *SessionInfo) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetClient

`func (o *SessionInfo) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *SessionInfo) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *SessionInfo) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *SessionInfo) HasClient() bool`

HasClient returns a boolean if a field has been set.

### SetClientNil

`func (o *SessionInfo) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *SessionInfo) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil
### GetLastActivityDate

`func (o *SessionInfo) GetLastActivityDate() time.Time`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *SessionInfo) GetLastActivityDateOk() (*time.Time, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *SessionInfo) SetLastActivityDate(v time.Time)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *SessionInfo) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### GetLastPlaybackCheckIn

`func (o *SessionInfo) GetLastPlaybackCheckIn() time.Time`

GetLastPlaybackCheckIn returns the LastPlaybackCheckIn field if non-nil, zero value otherwise.

### GetLastPlaybackCheckInOk

`func (o *SessionInfo) GetLastPlaybackCheckInOk() (*time.Time, bool)`

GetLastPlaybackCheckInOk returns a tuple with the LastPlaybackCheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlaybackCheckIn

`func (o *SessionInfo) SetLastPlaybackCheckIn(v time.Time)`

SetLastPlaybackCheckIn sets LastPlaybackCheckIn field to given value.

### HasLastPlaybackCheckIn

`func (o *SessionInfo) HasLastPlaybackCheckIn() bool`

HasLastPlaybackCheckIn returns a boolean if a field has been set.

### GetLastPausedDate

`func (o *SessionInfo) GetLastPausedDate() time.Time`

GetLastPausedDate returns the LastPausedDate field if non-nil, zero value otherwise.

### GetLastPausedDateOk

`func (o *SessionInfo) GetLastPausedDateOk() (*time.Time, bool)`

GetLastPausedDateOk returns a tuple with the LastPausedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPausedDate

`func (o *SessionInfo) SetLastPausedDate(v time.Time)`

SetLastPausedDate sets LastPausedDate field to given value.

### HasLastPausedDate

`func (o *SessionInfo) HasLastPausedDate() bool`

HasLastPausedDate returns a boolean if a field has been set.

### SetLastPausedDateNil

`func (o *SessionInfo) SetLastPausedDateNil(b bool)`

 SetLastPausedDateNil sets the value for LastPausedDate to be an explicit nil

### UnsetLastPausedDate
`func (o *SessionInfo) UnsetLastPausedDate()`

UnsetLastPausedDate ensures that no value is present for LastPausedDate, not even an explicit nil
### GetDeviceName

`func (o *SessionInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *SessionInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *SessionInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *SessionInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *SessionInfo) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *SessionInfo) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetDeviceType

`func (o *SessionInfo) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *SessionInfo) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *SessionInfo) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *SessionInfo) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *SessionInfo) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *SessionInfo) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetNowPlayingItem

`func (o *SessionInfo) GetNowPlayingItem() BaseItemDto`

GetNowPlayingItem returns the NowPlayingItem field if non-nil, zero value otherwise.

### GetNowPlayingItemOk

`func (o *SessionInfo) GetNowPlayingItemOk() (*BaseItemDto, bool)`

GetNowPlayingItemOk returns a tuple with the NowPlayingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingItem

`func (o *SessionInfo) SetNowPlayingItem(v BaseItemDto)`

SetNowPlayingItem sets NowPlayingItem field to given value.

### HasNowPlayingItem

`func (o *SessionInfo) HasNowPlayingItem() bool`

HasNowPlayingItem returns a boolean if a field has been set.

### SetNowPlayingItemNil

`func (o *SessionInfo) SetNowPlayingItemNil(b bool)`

 SetNowPlayingItemNil sets the value for NowPlayingItem to be an explicit nil

### UnsetNowPlayingItem
`func (o *SessionInfo) UnsetNowPlayingItem()`

UnsetNowPlayingItem ensures that no value is present for NowPlayingItem, not even an explicit nil
### GetNowViewingItem

`func (o *SessionInfo) GetNowViewingItem() BaseItemDto`

GetNowViewingItem returns the NowViewingItem field if non-nil, zero value otherwise.

### GetNowViewingItemOk

`func (o *SessionInfo) GetNowViewingItemOk() (*BaseItemDto, bool)`

GetNowViewingItemOk returns a tuple with the NowViewingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowViewingItem

`func (o *SessionInfo) SetNowViewingItem(v BaseItemDto)`

SetNowViewingItem sets NowViewingItem field to given value.

### HasNowViewingItem

`func (o *SessionInfo) HasNowViewingItem() bool`

HasNowViewingItem returns a boolean if a field has been set.

### SetNowViewingItemNil

`func (o *SessionInfo) SetNowViewingItemNil(b bool)`

 SetNowViewingItemNil sets the value for NowViewingItem to be an explicit nil

### UnsetNowViewingItem
`func (o *SessionInfo) UnsetNowViewingItem()`

UnsetNowViewingItem ensures that no value is present for NowViewingItem, not even an explicit nil
### GetDeviceId

`func (o *SessionInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SessionInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SessionInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *SessionInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *SessionInfo) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *SessionInfo) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetApplicationVersion

`func (o *SessionInfo) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *SessionInfo) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *SessionInfo) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *SessionInfo) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### SetApplicationVersionNil

`func (o *SessionInfo) SetApplicationVersionNil(b bool)`

 SetApplicationVersionNil sets the value for ApplicationVersion to be an explicit nil

### UnsetApplicationVersion
`func (o *SessionInfo) UnsetApplicationVersion()`

UnsetApplicationVersion ensures that no value is present for ApplicationVersion, not even an explicit nil
### GetTranscodingInfo

`func (o *SessionInfo) GetTranscodingInfo() TranscodingInfo`

GetTranscodingInfo returns the TranscodingInfo field if non-nil, zero value otherwise.

### GetTranscodingInfoOk

`func (o *SessionInfo) GetTranscodingInfoOk() (*TranscodingInfo, bool)`

GetTranscodingInfoOk returns a tuple with the TranscodingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingInfo

`func (o *SessionInfo) SetTranscodingInfo(v TranscodingInfo)`

SetTranscodingInfo sets TranscodingInfo field to given value.

### HasTranscodingInfo

`func (o *SessionInfo) HasTranscodingInfo() bool`

HasTranscodingInfo returns a boolean if a field has been set.

### SetTranscodingInfoNil

`func (o *SessionInfo) SetTranscodingInfoNil(b bool)`

 SetTranscodingInfoNil sets the value for TranscodingInfo to be an explicit nil

### UnsetTranscodingInfo
`func (o *SessionInfo) UnsetTranscodingInfo()`

UnsetTranscodingInfo ensures that no value is present for TranscodingInfo, not even an explicit nil
### GetIsActive

`func (o *SessionInfo) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SessionInfo) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SessionInfo) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *SessionInfo) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetSupportsMediaControl

`func (o *SessionInfo) GetSupportsMediaControl() bool`

GetSupportsMediaControl returns the SupportsMediaControl field if non-nil, zero value otherwise.

### GetSupportsMediaControlOk

`func (o *SessionInfo) GetSupportsMediaControlOk() (*bool, bool)`

GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMediaControl

`func (o *SessionInfo) SetSupportsMediaControl(v bool)`

SetSupportsMediaControl sets SupportsMediaControl field to given value.

### HasSupportsMediaControl

`func (o *SessionInfo) HasSupportsMediaControl() bool`

HasSupportsMediaControl returns a boolean if a field has been set.

### GetSupportsRemoteControl

`func (o *SessionInfo) GetSupportsRemoteControl() bool`

GetSupportsRemoteControl returns the SupportsRemoteControl field if non-nil, zero value otherwise.

### GetSupportsRemoteControlOk

`func (o *SessionInfo) GetSupportsRemoteControlOk() (*bool, bool)`

GetSupportsRemoteControlOk returns a tuple with the SupportsRemoteControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRemoteControl

`func (o *SessionInfo) SetSupportsRemoteControl(v bool)`

SetSupportsRemoteControl sets SupportsRemoteControl field to given value.

### HasSupportsRemoteControl

`func (o *SessionInfo) HasSupportsRemoteControl() bool`

HasSupportsRemoteControl returns a boolean if a field has been set.

### GetNowPlayingQueue

`func (o *SessionInfo) GetNowPlayingQueue() []QueueItem`

GetNowPlayingQueue returns the NowPlayingQueue field if non-nil, zero value otherwise.

### GetNowPlayingQueueOk

`func (o *SessionInfo) GetNowPlayingQueueOk() (*[]QueueItem, bool)`

GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueue

`func (o *SessionInfo) SetNowPlayingQueue(v []QueueItem)`

SetNowPlayingQueue sets NowPlayingQueue field to given value.

### HasNowPlayingQueue

`func (o *SessionInfo) HasNowPlayingQueue() bool`

HasNowPlayingQueue returns a boolean if a field has been set.

### SetNowPlayingQueueNil

`func (o *SessionInfo) SetNowPlayingQueueNil(b bool)`

 SetNowPlayingQueueNil sets the value for NowPlayingQueue to be an explicit nil

### UnsetNowPlayingQueue
`func (o *SessionInfo) UnsetNowPlayingQueue()`

UnsetNowPlayingQueue ensures that no value is present for NowPlayingQueue, not even an explicit nil
### GetNowPlayingQueueFullItems

`func (o *SessionInfo) GetNowPlayingQueueFullItems() []BaseItemDto`

GetNowPlayingQueueFullItems returns the NowPlayingQueueFullItems field if non-nil, zero value otherwise.

### GetNowPlayingQueueFullItemsOk

`func (o *SessionInfo) GetNowPlayingQueueFullItemsOk() (*[]BaseItemDto, bool)`

GetNowPlayingQueueFullItemsOk returns a tuple with the NowPlayingQueueFullItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueueFullItems

`func (o *SessionInfo) SetNowPlayingQueueFullItems(v []BaseItemDto)`

SetNowPlayingQueueFullItems sets NowPlayingQueueFullItems field to given value.

### HasNowPlayingQueueFullItems

`func (o *SessionInfo) HasNowPlayingQueueFullItems() bool`

HasNowPlayingQueueFullItems returns a boolean if a field has been set.

### SetNowPlayingQueueFullItemsNil

`func (o *SessionInfo) SetNowPlayingQueueFullItemsNil(b bool)`

 SetNowPlayingQueueFullItemsNil sets the value for NowPlayingQueueFullItems to be an explicit nil

### UnsetNowPlayingQueueFullItems
`func (o *SessionInfo) UnsetNowPlayingQueueFullItems()`

UnsetNowPlayingQueueFullItems ensures that no value is present for NowPlayingQueueFullItems, not even an explicit nil
### GetHasCustomDeviceName

`func (o *SessionInfo) GetHasCustomDeviceName() bool`

GetHasCustomDeviceName returns the HasCustomDeviceName field if non-nil, zero value otherwise.

### GetHasCustomDeviceNameOk

`func (o *SessionInfo) GetHasCustomDeviceNameOk() (*bool, bool)`

GetHasCustomDeviceNameOk returns a tuple with the HasCustomDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCustomDeviceName

`func (o *SessionInfo) SetHasCustomDeviceName(v bool)`

SetHasCustomDeviceName sets HasCustomDeviceName field to given value.

### HasHasCustomDeviceName

`func (o *SessionInfo) HasHasCustomDeviceName() bool`

HasHasCustomDeviceName returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *SessionInfo) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *SessionInfo) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *SessionInfo) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *SessionInfo) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### SetPlaylistItemIdNil

`func (o *SessionInfo) SetPlaylistItemIdNil(b bool)`

 SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil

### UnsetPlaylistItemId
`func (o *SessionInfo) UnsetPlaylistItemId()`

UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil
### GetServerId

`func (o *SessionInfo) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *SessionInfo) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *SessionInfo) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *SessionInfo) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *SessionInfo) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *SessionInfo) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetUserPrimaryImageTag

`func (o *SessionInfo) GetUserPrimaryImageTag() string`

GetUserPrimaryImageTag returns the UserPrimaryImageTag field if non-nil, zero value otherwise.

### GetUserPrimaryImageTagOk

`func (o *SessionInfo) GetUserPrimaryImageTagOk() (*string, bool)`

GetUserPrimaryImageTagOk returns a tuple with the UserPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrimaryImageTag

`func (o *SessionInfo) SetUserPrimaryImageTag(v string)`

SetUserPrimaryImageTag sets UserPrimaryImageTag field to given value.

### HasUserPrimaryImageTag

`func (o *SessionInfo) HasUserPrimaryImageTag() bool`

HasUserPrimaryImageTag returns a boolean if a field has been set.

### SetUserPrimaryImageTagNil

`func (o *SessionInfo) SetUserPrimaryImageTagNil(b bool)`

 SetUserPrimaryImageTagNil sets the value for UserPrimaryImageTag to be an explicit nil

### UnsetUserPrimaryImageTag
`func (o *SessionInfo) UnsetUserPrimaryImageTag()`

UnsetUserPrimaryImageTag ensures that no value is present for UserPrimaryImageTag, not even an explicit nil
### GetSupportedCommands

`func (o *SessionInfo) GetSupportedCommands() []GeneralCommandType`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *SessionInfo) GetSupportedCommandsOk() (*[]GeneralCommandType, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *SessionInfo) SetSupportedCommands(v []GeneralCommandType)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *SessionInfo) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.

### SetSupportedCommandsNil

`func (o *SessionInfo) SetSupportedCommandsNil(b bool)`

 SetSupportedCommandsNil sets the value for SupportedCommands to be an explicit nil

### UnsetSupportedCommands
`func (o *SessionInfo) UnsetSupportedCommands()`

UnsetSupportedCommands ensures that no value is present for SupportedCommands, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


