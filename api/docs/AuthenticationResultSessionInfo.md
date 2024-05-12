# AuthenticationResultSessionInfo

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
**NowPlayingItem** | Pointer to [**NullableSessionInfoNowPlayingItem**](SessionInfoNowPlayingItem.md) |  | [optional] 
**NowViewingItem** | Pointer to [**NullableSessionInfoNowPlayingItem**](SessionInfoNowPlayingItem.md) |  | [optional] 
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

### NewAuthenticationResultSessionInfo

`func NewAuthenticationResultSessionInfo() *AuthenticationResultSessionInfo`

NewAuthenticationResultSessionInfo instantiates a new AuthenticationResultSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationResultSessionInfoWithDefaults

`func NewAuthenticationResultSessionInfoWithDefaults() *AuthenticationResultSessionInfo`

NewAuthenticationResultSessionInfoWithDefaults instantiates a new AuthenticationResultSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayState

`func (o *AuthenticationResultSessionInfo) GetPlayState() PlayerStateInfo`

GetPlayState returns the PlayState field if non-nil, zero value otherwise.

### GetPlayStateOk

`func (o *AuthenticationResultSessionInfo) GetPlayStateOk() (*PlayerStateInfo, bool)`

GetPlayStateOk returns a tuple with the PlayState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayState

`func (o *AuthenticationResultSessionInfo) SetPlayState(v PlayerStateInfo)`

SetPlayState sets PlayState field to given value.

### HasPlayState

`func (o *AuthenticationResultSessionInfo) HasPlayState() bool`

HasPlayState returns a boolean if a field has been set.

### SetPlayStateNil

`func (o *AuthenticationResultSessionInfo) SetPlayStateNil(b bool)`

 SetPlayStateNil sets the value for PlayState to be an explicit nil

### UnsetPlayState
`func (o *AuthenticationResultSessionInfo) UnsetPlayState()`

UnsetPlayState ensures that no value is present for PlayState, not even an explicit nil
### GetAdditionalUsers

`func (o *AuthenticationResultSessionInfo) GetAdditionalUsers() []SessionUserInfo`

GetAdditionalUsers returns the AdditionalUsers field if non-nil, zero value otherwise.

### GetAdditionalUsersOk

`func (o *AuthenticationResultSessionInfo) GetAdditionalUsersOk() (*[]SessionUserInfo, bool)`

GetAdditionalUsersOk returns a tuple with the AdditionalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUsers

`func (o *AuthenticationResultSessionInfo) SetAdditionalUsers(v []SessionUserInfo)`

SetAdditionalUsers sets AdditionalUsers field to given value.

### HasAdditionalUsers

`func (o *AuthenticationResultSessionInfo) HasAdditionalUsers() bool`

HasAdditionalUsers returns a boolean if a field has been set.

### SetAdditionalUsersNil

`func (o *AuthenticationResultSessionInfo) SetAdditionalUsersNil(b bool)`

 SetAdditionalUsersNil sets the value for AdditionalUsers to be an explicit nil

### UnsetAdditionalUsers
`func (o *AuthenticationResultSessionInfo) UnsetAdditionalUsers()`

UnsetAdditionalUsers ensures that no value is present for AdditionalUsers, not even an explicit nil
### GetCapabilities

`func (o *AuthenticationResultSessionInfo) GetCapabilities() ClientCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AuthenticationResultSessionInfo) GetCapabilitiesOk() (*ClientCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AuthenticationResultSessionInfo) SetCapabilities(v ClientCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *AuthenticationResultSessionInfo) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *AuthenticationResultSessionInfo) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *AuthenticationResultSessionInfo) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetRemoteEndPoint

`func (o *AuthenticationResultSessionInfo) GetRemoteEndPoint() string`

GetRemoteEndPoint returns the RemoteEndPoint field if non-nil, zero value otherwise.

### GetRemoteEndPointOk

`func (o *AuthenticationResultSessionInfo) GetRemoteEndPointOk() (*string, bool)`

GetRemoteEndPointOk returns a tuple with the RemoteEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteEndPoint

`func (o *AuthenticationResultSessionInfo) SetRemoteEndPoint(v string)`

SetRemoteEndPoint sets RemoteEndPoint field to given value.

### HasRemoteEndPoint

`func (o *AuthenticationResultSessionInfo) HasRemoteEndPoint() bool`

HasRemoteEndPoint returns a boolean if a field has been set.

### SetRemoteEndPointNil

`func (o *AuthenticationResultSessionInfo) SetRemoteEndPointNil(b bool)`

 SetRemoteEndPointNil sets the value for RemoteEndPoint to be an explicit nil

### UnsetRemoteEndPoint
`func (o *AuthenticationResultSessionInfo) UnsetRemoteEndPoint()`

UnsetRemoteEndPoint ensures that no value is present for RemoteEndPoint, not even an explicit nil
### GetPlayableMediaTypes

`func (o *AuthenticationResultSessionInfo) GetPlayableMediaTypes() []MediaType`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *AuthenticationResultSessionInfo) GetPlayableMediaTypesOk() (*[]MediaType, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *AuthenticationResultSessionInfo) SetPlayableMediaTypes(v []MediaType)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *AuthenticationResultSessionInfo) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### SetPlayableMediaTypesNil

`func (o *AuthenticationResultSessionInfo) SetPlayableMediaTypesNil(b bool)`

 SetPlayableMediaTypesNil sets the value for PlayableMediaTypes to be an explicit nil

### UnsetPlayableMediaTypes
`func (o *AuthenticationResultSessionInfo) UnsetPlayableMediaTypes()`

UnsetPlayableMediaTypes ensures that no value is present for PlayableMediaTypes, not even an explicit nil
### GetId

`func (o *AuthenticationResultSessionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationResultSessionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationResultSessionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationResultSessionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AuthenticationResultSessionInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AuthenticationResultSessionInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetUserId

`func (o *AuthenticationResultSessionInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthenticationResultSessionInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthenticationResultSessionInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuthenticationResultSessionInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *AuthenticationResultSessionInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AuthenticationResultSessionInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AuthenticationResultSessionInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AuthenticationResultSessionInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *AuthenticationResultSessionInfo) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *AuthenticationResultSessionInfo) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetClient

`func (o *AuthenticationResultSessionInfo) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *AuthenticationResultSessionInfo) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *AuthenticationResultSessionInfo) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *AuthenticationResultSessionInfo) HasClient() bool`

HasClient returns a boolean if a field has been set.

### SetClientNil

`func (o *AuthenticationResultSessionInfo) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *AuthenticationResultSessionInfo) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil
### GetLastActivityDate

`func (o *AuthenticationResultSessionInfo) GetLastActivityDate() time.Time`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *AuthenticationResultSessionInfo) GetLastActivityDateOk() (*time.Time, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *AuthenticationResultSessionInfo) SetLastActivityDate(v time.Time)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *AuthenticationResultSessionInfo) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### GetLastPlaybackCheckIn

`func (o *AuthenticationResultSessionInfo) GetLastPlaybackCheckIn() time.Time`

GetLastPlaybackCheckIn returns the LastPlaybackCheckIn field if non-nil, zero value otherwise.

### GetLastPlaybackCheckInOk

`func (o *AuthenticationResultSessionInfo) GetLastPlaybackCheckInOk() (*time.Time, bool)`

GetLastPlaybackCheckInOk returns a tuple with the LastPlaybackCheckIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlaybackCheckIn

`func (o *AuthenticationResultSessionInfo) SetLastPlaybackCheckIn(v time.Time)`

SetLastPlaybackCheckIn sets LastPlaybackCheckIn field to given value.

### HasLastPlaybackCheckIn

`func (o *AuthenticationResultSessionInfo) HasLastPlaybackCheckIn() bool`

HasLastPlaybackCheckIn returns a boolean if a field has been set.

### GetLastPausedDate

`func (o *AuthenticationResultSessionInfo) GetLastPausedDate() time.Time`

GetLastPausedDate returns the LastPausedDate field if non-nil, zero value otherwise.

### GetLastPausedDateOk

`func (o *AuthenticationResultSessionInfo) GetLastPausedDateOk() (*time.Time, bool)`

GetLastPausedDateOk returns a tuple with the LastPausedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPausedDate

`func (o *AuthenticationResultSessionInfo) SetLastPausedDate(v time.Time)`

SetLastPausedDate sets LastPausedDate field to given value.

### HasLastPausedDate

`func (o *AuthenticationResultSessionInfo) HasLastPausedDate() bool`

HasLastPausedDate returns a boolean if a field has been set.

### SetLastPausedDateNil

`func (o *AuthenticationResultSessionInfo) SetLastPausedDateNil(b bool)`

 SetLastPausedDateNil sets the value for LastPausedDate to be an explicit nil

### UnsetLastPausedDate
`func (o *AuthenticationResultSessionInfo) UnsetLastPausedDate()`

UnsetLastPausedDate ensures that no value is present for LastPausedDate, not even an explicit nil
### GetDeviceName

`func (o *AuthenticationResultSessionInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *AuthenticationResultSessionInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *AuthenticationResultSessionInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *AuthenticationResultSessionInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *AuthenticationResultSessionInfo) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *AuthenticationResultSessionInfo) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetDeviceType

`func (o *AuthenticationResultSessionInfo) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *AuthenticationResultSessionInfo) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *AuthenticationResultSessionInfo) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *AuthenticationResultSessionInfo) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *AuthenticationResultSessionInfo) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *AuthenticationResultSessionInfo) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetNowPlayingItem

`func (o *AuthenticationResultSessionInfo) GetNowPlayingItem() SessionInfoNowPlayingItem`

GetNowPlayingItem returns the NowPlayingItem field if non-nil, zero value otherwise.

### GetNowPlayingItemOk

`func (o *AuthenticationResultSessionInfo) GetNowPlayingItemOk() (*SessionInfoNowPlayingItem, bool)`

GetNowPlayingItemOk returns a tuple with the NowPlayingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingItem

`func (o *AuthenticationResultSessionInfo) SetNowPlayingItem(v SessionInfoNowPlayingItem)`

SetNowPlayingItem sets NowPlayingItem field to given value.

### HasNowPlayingItem

`func (o *AuthenticationResultSessionInfo) HasNowPlayingItem() bool`

HasNowPlayingItem returns a boolean if a field has been set.

### SetNowPlayingItemNil

`func (o *AuthenticationResultSessionInfo) SetNowPlayingItemNil(b bool)`

 SetNowPlayingItemNil sets the value for NowPlayingItem to be an explicit nil

### UnsetNowPlayingItem
`func (o *AuthenticationResultSessionInfo) UnsetNowPlayingItem()`

UnsetNowPlayingItem ensures that no value is present for NowPlayingItem, not even an explicit nil
### GetNowViewingItem

`func (o *AuthenticationResultSessionInfo) GetNowViewingItem() SessionInfoNowPlayingItem`

GetNowViewingItem returns the NowViewingItem field if non-nil, zero value otherwise.

### GetNowViewingItemOk

`func (o *AuthenticationResultSessionInfo) GetNowViewingItemOk() (*SessionInfoNowPlayingItem, bool)`

GetNowViewingItemOk returns a tuple with the NowViewingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowViewingItem

`func (o *AuthenticationResultSessionInfo) SetNowViewingItem(v SessionInfoNowPlayingItem)`

SetNowViewingItem sets NowViewingItem field to given value.

### HasNowViewingItem

`func (o *AuthenticationResultSessionInfo) HasNowViewingItem() bool`

HasNowViewingItem returns a boolean if a field has been set.

### SetNowViewingItemNil

`func (o *AuthenticationResultSessionInfo) SetNowViewingItemNil(b bool)`

 SetNowViewingItemNil sets the value for NowViewingItem to be an explicit nil

### UnsetNowViewingItem
`func (o *AuthenticationResultSessionInfo) UnsetNowViewingItem()`

UnsetNowViewingItem ensures that no value is present for NowViewingItem, not even an explicit nil
### GetDeviceId

`func (o *AuthenticationResultSessionInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AuthenticationResultSessionInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AuthenticationResultSessionInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AuthenticationResultSessionInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *AuthenticationResultSessionInfo) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *AuthenticationResultSessionInfo) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetApplicationVersion

`func (o *AuthenticationResultSessionInfo) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *AuthenticationResultSessionInfo) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *AuthenticationResultSessionInfo) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *AuthenticationResultSessionInfo) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### SetApplicationVersionNil

`func (o *AuthenticationResultSessionInfo) SetApplicationVersionNil(b bool)`

 SetApplicationVersionNil sets the value for ApplicationVersion to be an explicit nil

### UnsetApplicationVersion
`func (o *AuthenticationResultSessionInfo) UnsetApplicationVersion()`

UnsetApplicationVersion ensures that no value is present for ApplicationVersion, not even an explicit nil
### GetTranscodingInfo

`func (o *AuthenticationResultSessionInfo) GetTranscodingInfo() TranscodingInfo`

GetTranscodingInfo returns the TranscodingInfo field if non-nil, zero value otherwise.

### GetTranscodingInfoOk

`func (o *AuthenticationResultSessionInfo) GetTranscodingInfoOk() (*TranscodingInfo, bool)`

GetTranscodingInfoOk returns a tuple with the TranscodingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingInfo

`func (o *AuthenticationResultSessionInfo) SetTranscodingInfo(v TranscodingInfo)`

SetTranscodingInfo sets TranscodingInfo field to given value.

### HasTranscodingInfo

`func (o *AuthenticationResultSessionInfo) HasTranscodingInfo() bool`

HasTranscodingInfo returns a boolean if a field has been set.

### SetTranscodingInfoNil

`func (o *AuthenticationResultSessionInfo) SetTranscodingInfoNil(b bool)`

 SetTranscodingInfoNil sets the value for TranscodingInfo to be an explicit nil

### UnsetTranscodingInfo
`func (o *AuthenticationResultSessionInfo) UnsetTranscodingInfo()`

UnsetTranscodingInfo ensures that no value is present for TranscodingInfo, not even an explicit nil
### GetIsActive

`func (o *AuthenticationResultSessionInfo) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AuthenticationResultSessionInfo) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AuthenticationResultSessionInfo) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AuthenticationResultSessionInfo) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetSupportsMediaControl

`func (o *AuthenticationResultSessionInfo) GetSupportsMediaControl() bool`

GetSupportsMediaControl returns the SupportsMediaControl field if non-nil, zero value otherwise.

### GetSupportsMediaControlOk

`func (o *AuthenticationResultSessionInfo) GetSupportsMediaControlOk() (*bool, bool)`

GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMediaControl

`func (o *AuthenticationResultSessionInfo) SetSupportsMediaControl(v bool)`

SetSupportsMediaControl sets SupportsMediaControl field to given value.

### HasSupportsMediaControl

`func (o *AuthenticationResultSessionInfo) HasSupportsMediaControl() bool`

HasSupportsMediaControl returns a boolean if a field has been set.

### GetSupportsRemoteControl

`func (o *AuthenticationResultSessionInfo) GetSupportsRemoteControl() bool`

GetSupportsRemoteControl returns the SupportsRemoteControl field if non-nil, zero value otherwise.

### GetSupportsRemoteControlOk

`func (o *AuthenticationResultSessionInfo) GetSupportsRemoteControlOk() (*bool, bool)`

GetSupportsRemoteControlOk returns a tuple with the SupportsRemoteControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsRemoteControl

`func (o *AuthenticationResultSessionInfo) SetSupportsRemoteControl(v bool)`

SetSupportsRemoteControl sets SupportsRemoteControl field to given value.

### HasSupportsRemoteControl

`func (o *AuthenticationResultSessionInfo) HasSupportsRemoteControl() bool`

HasSupportsRemoteControl returns a boolean if a field has been set.

### GetNowPlayingQueue

`func (o *AuthenticationResultSessionInfo) GetNowPlayingQueue() []QueueItem`

GetNowPlayingQueue returns the NowPlayingQueue field if non-nil, zero value otherwise.

### GetNowPlayingQueueOk

`func (o *AuthenticationResultSessionInfo) GetNowPlayingQueueOk() (*[]QueueItem, bool)`

GetNowPlayingQueueOk returns a tuple with the NowPlayingQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueue

`func (o *AuthenticationResultSessionInfo) SetNowPlayingQueue(v []QueueItem)`

SetNowPlayingQueue sets NowPlayingQueue field to given value.

### HasNowPlayingQueue

`func (o *AuthenticationResultSessionInfo) HasNowPlayingQueue() bool`

HasNowPlayingQueue returns a boolean if a field has been set.

### SetNowPlayingQueueNil

`func (o *AuthenticationResultSessionInfo) SetNowPlayingQueueNil(b bool)`

 SetNowPlayingQueueNil sets the value for NowPlayingQueue to be an explicit nil

### UnsetNowPlayingQueue
`func (o *AuthenticationResultSessionInfo) UnsetNowPlayingQueue()`

UnsetNowPlayingQueue ensures that no value is present for NowPlayingQueue, not even an explicit nil
### GetNowPlayingQueueFullItems

`func (o *AuthenticationResultSessionInfo) GetNowPlayingQueueFullItems() []BaseItemDto`

GetNowPlayingQueueFullItems returns the NowPlayingQueueFullItems field if non-nil, zero value otherwise.

### GetNowPlayingQueueFullItemsOk

`func (o *AuthenticationResultSessionInfo) GetNowPlayingQueueFullItemsOk() (*[]BaseItemDto, bool)`

GetNowPlayingQueueFullItemsOk returns a tuple with the NowPlayingQueueFullItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowPlayingQueueFullItems

`func (o *AuthenticationResultSessionInfo) SetNowPlayingQueueFullItems(v []BaseItemDto)`

SetNowPlayingQueueFullItems sets NowPlayingQueueFullItems field to given value.

### HasNowPlayingQueueFullItems

`func (o *AuthenticationResultSessionInfo) HasNowPlayingQueueFullItems() bool`

HasNowPlayingQueueFullItems returns a boolean if a field has been set.

### SetNowPlayingQueueFullItemsNil

`func (o *AuthenticationResultSessionInfo) SetNowPlayingQueueFullItemsNil(b bool)`

 SetNowPlayingQueueFullItemsNil sets the value for NowPlayingQueueFullItems to be an explicit nil

### UnsetNowPlayingQueueFullItems
`func (o *AuthenticationResultSessionInfo) UnsetNowPlayingQueueFullItems()`

UnsetNowPlayingQueueFullItems ensures that no value is present for NowPlayingQueueFullItems, not even an explicit nil
### GetHasCustomDeviceName

`func (o *AuthenticationResultSessionInfo) GetHasCustomDeviceName() bool`

GetHasCustomDeviceName returns the HasCustomDeviceName field if non-nil, zero value otherwise.

### GetHasCustomDeviceNameOk

`func (o *AuthenticationResultSessionInfo) GetHasCustomDeviceNameOk() (*bool, bool)`

GetHasCustomDeviceNameOk returns a tuple with the HasCustomDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCustomDeviceName

`func (o *AuthenticationResultSessionInfo) SetHasCustomDeviceName(v bool)`

SetHasCustomDeviceName sets HasCustomDeviceName field to given value.

### HasHasCustomDeviceName

`func (o *AuthenticationResultSessionInfo) HasHasCustomDeviceName() bool`

HasHasCustomDeviceName returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *AuthenticationResultSessionInfo) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *AuthenticationResultSessionInfo) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *AuthenticationResultSessionInfo) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *AuthenticationResultSessionInfo) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### SetPlaylistItemIdNil

`func (o *AuthenticationResultSessionInfo) SetPlaylistItemIdNil(b bool)`

 SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil

### UnsetPlaylistItemId
`func (o *AuthenticationResultSessionInfo) UnsetPlaylistItemId()`

UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil
### GetServerId

`func (o *AuthenticationResultSessionInfo) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *AuthenticationResultSessionInfo) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *AuthenticationResultSessionInfo) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *AuthenticationResultSessionInfo) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *AuthenticationResultSessionInfo) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *AuthenticationResultSessionInfo) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetUserPrimaryImageTag

`func (o *AuthenticationResultSessionInfo) GetUserPrimaryImageTag() string`

GetUserPrimaryImageTag returns the UserPrimaryImageTag field if non-nil, zero value otherwise.

### GetUserPrimaryImageTagOk

`func (o *AuthenticationResultSessionInfo) GetUserPrimaryImageTagOk() (*string, bool)`

GetUserPrimaryImageTagOk returns a tuple with the UserPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrimaryImageTag

`func (o *AuthenticationResultSessionInfo) SetUserPrimaryImageTag(v string)`

SetUserPrimaryImageTag sets UserPrimaryImageTag field to given value.

### HasUserPrimaryImageTag

`func (o *AuthenticationResultSessionInfo) HasUserPrimaryImageTag() bool`

HasUserPrimaryImageTag returns a boolean if a field has been set.

### SetUserPrimaryImageTagNil

`func (o *AuthenticationResultSessionInfo) SetUserPrimaryImageTagNil(b bool)`

 SetUserPrimaryImageTagNil sets the value for UserPrimaryImageTag to be an explicit nil

### UnsetUserPrimaryImageTag
`func (o *AuthenticationResultSessionInfo) UnsetUserPrimaryImageTag()`

UnsetUserPrimaryImageTag ensures that no value is present for UserPrimaryImageTag, not even an explicit nil
### GetSupportedCommands

`func (o *AuthenticationResultSessionInfo) GetSupportedCommands() []GeneralCommandType`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *AuthenticationResultSessionInfo) GetSupportedCommandsOk() (*[]GeneralCommandType, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *AuthenticationResultSessionInfo) SetSupportedCommands(v []GeneralCommandType)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *AuthenticationResultSessionInfo) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.

### SetSupportedCommandsNil

`func (o *AuthenticationResultSessionInfo) SetSupportedCommandsNil(b bool)`

 SetSupportedCommandsNil sets the value for SupportedCommands to be an explicit nil

### UnsetSupportedCommands
`func (o *AuthenticationResultSessionInfo) UnsetSupportedCommands()`

UnsetSupportedCommands ensures that no value is present for SupportedCommands, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


