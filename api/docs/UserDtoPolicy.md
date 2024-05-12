# UserDtoPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAdministrator** | Pointer to **bool** | Gets or sets a value indicating whether this instance is administrator. | [optional] 
**IsHidden** | Pointer to **bool** | Gets or sets a value indicating whether this instance is hidden. | [optional] 
**EnableCollectionManagement** | Pointer to **bool** | Gets or sets a value indicating whether this instance can manage collections. | [optional] [default to false]
**EnableSubtitleManagement** | Pointer to **bool** | Gets or sets a value indicating whether this instance can manage subtitles. | [optional] [default to false]
**EnableLyricManagement** | Pointer to **bool** | Gets or sets a value indicating whether this user can manage lyrics. | [optional] [default to false]
**IsDisabled** | Pointer to **bool** | Gets or sets a value indicating whether this instance is disabled. | [optional] 
**MaxParentalRating** | Pointer to **NullableInt32** | Gets or sets the max parental rating. | [optional] 
**BlockedTags** | Pointer to **[]string** |  | [optional] 
**AllowedTags** | Pointer to **[]string** |  | [optional] 
**EnableUserPreferenceAccess** | Pointer to **bool** |  | [optional] 
**AccessSchedules** | Pointer to [**[]AccessSchedule**](AccessSchedule.md) |  | [optional] 
**BlockUnratedItems** | Pointer to [**[]UnratedItem**](UnratedItem.md) |  | [optional] 
**EnableRemoteControlOfOtherUsers** | Pointer to **bool** |  | [optional] 
**EnableSharedDeviceControl** | Pointer to **bool** |  | [optional] 
**EnableRemoteAccess** | Pointer to **bool** |  | [optional] 
**EnableLiveTvManagement** | Pointer to **bool** |  | [optional] 
**EnableLiveTvAccess** | Pointer to **bool** |  | [optional] 
**EnableMediaPlayback** | Pointer to **bool** |  | [optional] 
**EnableAudioPlaybackTranscoding** | Pointer to **bool** |  | [optional] 
**EnableVideoPlaybackTranscoding** | Pointer to **bool** |  | [optional] 
**EnablePlaybackRemuxing** | Pointer to **bool** |  | [optional] 
**ForceRemoteSourceTranscoding** | Pointer to **bool** |  | [optional] 
**EnableContentDeletion** | Pointer to **bool** |  | [optional] 
**EnableContentDeletionFromFolders** | Pointer to **[]string** |  | [optional] 
**EnableContentDownloading** | Pointer to **bool** |  | [optional] 
**EnableSyncTranscoding** | Pointer to **bool** | Gets or sets a value indicating whether [enable synchronize]. | [optional] 
**EnableMediaConversion** | Pointer to **bool** |  | [optional] 
**EnabledDevices** | Pointer to **[]string** |  | [optional] 
**EnableAllDevices** | Pointer to **bool** |  | [optional] 
**EnabledChannels** | Pointer to **[]string** |  | [optional] 
**EnableAllChannels** | Pointer to **bool** |  | [optional] 
**EnabledFolders** | Pointer to **[]string** |  | [optional] 
**EnableAllFolders** | Pointer to **bool** |  | [optional] 
**InvalidLoginAttemptCount** | Pointer to **int32** |  | [optional] 
**LoginAttemptsBeforeLockout** | Pointer to **int32** |  | [optional] 
**MaxActiveSessions** | Pointer to **int32** |  | [optional] 
**EnablePublicSharing** | Pointer to **bool** |  | [optional] 
**BlockedMediaFolders** | Pointer to **[]string** |  | [optional] 
**BlockedChannels** | Pointer to **[]string** |  | [optional] 
**RemoteClientBitrateLimit** | Pointer to **int32** |  | [optional] 
**AuthenticationProviderId** | **string** |  | 
**PasswordResetProviderId** | **string** |  | 
**SyncPlayAccess** | Pointer to [**SyncPlayUserAccessType**](SyncPlayUserAccessType.md) | Gets or sets a value indicating what SyncPlay features the user can access. | [optional] 

## Methods

### NewUserDtoPolicy

`func NewUserDtoPolicy(authenticationProviderId string, passwordResetProviderId string, ) *UserDtoPolicy`

NewUserDtoPolicy instantiates a new UserDtoPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDtoPolicyWithDefaults

`func NewUserDtoPolicyWithDefaults() *UserDtoPolicy`

NewUserDtoPolicyWithDefaults instantiates a new UserDtoPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAdministrator

`func (o *UserDtoPolicy) GetIsAdministrator() bool`

GetIsAdministrator returns the IsAdministrator field if non-nil, zero value otherwise.

### GetIsAdministratorOk

`func (o *UserDtoPolicy) GetIsAdministratorOk() (*bool, bool)`

GetIsAdministratorOk returns a tuple with the IsAdministrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdministrator

`func (o *UserDtoPolicy) SetIsAdministrator(v bool)`

SetIsAdministrator sets IsAdministrator field to given value.

### HasIsAdministrator

`func (o *UserDtoPolicy) HasIsAdministrator() bool`

HasIsAdministrator returns a boolean if a field has been set.

### GetIsHidden

`func (o *UserDtoPolicy) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *UserDtoPolicy) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *UserDtoPolicy) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *UserDtoPolicy) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetEnableCollectionManagement

`func (o *UserDtoPolicy) GetEnableCollectionManagement() bool`

GetEnableCollectionManagement returns the EnableCollectionManagement field if non-nil, zero value otherwise.

### GetEnableCollectionManagementOk

`func (o *UserDtoPolicy) GetEnableCollectionManagementOk() (*bool, bool)`

GetEnableCollectionManagementOk returns a tuple with the EnableCollectionManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCollectionManagement

`func (o *UserDtoPolicy) SetEnableCollectionManagement(v bool)`

SetEnableCollectionManagement sets EnableCollectionManagement field to given value.

### HasEnableCollectionManagement

`func (o *UserDtoPolicy) HasEnableCollectionManagement() bool`

HasEnableCollectionManagement returns a boolean if a field has been set.

### GetEnableSubtitleManagement

`func (o *UserDtoPolicy) GetEnableSubtitleManagement() bool`

GetEnableSubtitleManagement returns the EnableSubtitleManagement field if non-nil, zero value otherwise.

### GetEnableSubtitleManagementOk

`func (o *UserDtoPolicy) GetEnableSubtitleManagementOk() (*bool, bool)`

GetEnableSubtitleManagementOk returns a tuple with the EnableSubtitleManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubtitleManagement

`func (o *UserDtoPolicy) SetEnableSubtitleManagement(v bool)`

SetEnableSubtitleManagement sets EnableSubtitleManagement field to given value.

### HasEnableSubtitleManagement

`func (o *UserDtoPolicy) HasEnableSubtitleManagement() bool`

HasEnableSubtitleManagement returns a boolean if a field has been set.

### GetEnableLyricManagement

`func (o *UserDtoPolicy) GetEnableLyricManagement() bool`

GetEnableLyricManagement returns the EnableLyricManagement field if non-nil, zero value otherwise.

### GetEnableLyricManagementOk

`func (o *UserDtoPolicy) GetEnableLyricManagementOk() (*bool, bool)`

GetEnableLyricManagementOk returns a tuple with the EnableLyricManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLyricManagement

`func (o *UserDtoPolicy) SetEnableLyricManagement(v bool)`

SetEnableLyricManagement sets EnableLyricManagement field to given value.

### HasEnableLyricManagement

`func (o *UserDtoPolicy) HasEnableLyricManagement() bool`

HasEnableLyricManagement returns a boolean if a field has been set.

### GetIsDisabled

`func (o *UserDtoPolicy) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *UserDtoPolicy) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *UserDtoPolicy) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *UserDtoPolicy) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetMaxParentalRating

`func (o *UserDtoPolicy) GetMaxParentalRating() int32`

GetMaxParentalRating returns the MaxParentalRating field if non-nil, zero value otherwise.

### GetMaxParentalRatingOk

`func (o *UserDtoPolicy) GetMaxParentalRatingOk() (*int32, bool)`

GetMaxParentalRatingOk returns a tuple with the MaxParentalRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParentalRating

`func (o *UserDtoPolicy) SetMaxParentalRating(v int32)`

SetMaxParentalRating sets MaxParentalRating field to given value.

### HasMaxParentalRating

`func (o *UserDtoPolicy) HasMaxParentalRating() bool`

HasMaxParentalRating returns a boolean if a field has been set.

### SetMaxParentalRatingNil

`func (o *UserDtoPolicy) SetMaxParentalRatingNil(b bool)`

 SetMaxParentalRatingNil sets the value for MaxParentalRating to be an explicit nil

### UnsetMaxParentalRating
`func (o *UserDtoPolicy) UnsetMaxParentalRating()`

UnsetMaxParentalRating ensures that no value is present for MaxParentalRating, not even an explicit nil
### GetBlockedTags

`func (o *UserDtoPolicy) GetBlockedTags() []string`

GetBlockedTags returns the BlockedTags field if non-nil, zero value otherwise.

### GetBlockedTagsOk

`func (o *UserDtoPolicy) GetBlockedTagsOk() (*[]string, bool)`

GetBlockedTagsOk returns a tuple with the BlockedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedTags

`func (o *UserDtoPolicy) SetBlockedTags(v []string)`

SetBlockedTags sets BlockedTags field to given value.

### HasBlockedTags

`func (o *UserDtoPolicy) HasBlockedTags() bool`

HasBlockedTags returns a boolean if a field has been set.

### SetBlockedTagsNil

`func (o *UserDtoPolicy) SetBlockedTagsNil(b bool)`

 SetBlockedTagsNil sets the value for BlockedTags to be an explicit nil

### UnsetBlockedTags
`func (o *UserDtoPolicy) UnsetBlockedTags()`

UnsetBlockedTags ensures that no value is present for BlockedTags, not even an explicit nil
### GetAllowedTags

`func (o *UserDtoPolicy) GetAllowedTags() []string`

GetAllowedTags returns the AllowedTags field if non-nil, zero value otherwise.

### GetAllowedTagsOk

`func (o *UserDtoPolicy) GetAllowedTagsOk() (*[]string, bool)`

GetAllowedTagsOk returns a tuple with the AllowedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTags

`func (o *UserDtoPolicy) SetAllowedTags(v []string)`

SetAllowedTags sets AllowedTags field to given value.

### HasAllowedTags

`func (o *UserDtoPolicy) HasAllowedTags() bool`

HasAllowedTags returns a boolean if a field has been set.

### SetAllowedTagsNil

`func (o *UserDtoPolicy) SetAllowedTagsNil(b bool)`

 SetAllowedTagsNil sets the value for AllowedTags to be an explicit nil

### UnsetAllowedTags
`func (o *UserDtoPolicy) UnsetAllowedTags()`

UnsetAllowedTags ensures that no value is present for AllowedTags, not even an explicit nil
### GetEnableUserPreferenceAccess

`func (o *UserDtoPolicy) GetEnableUserPreferenceAccess() bool`

GetEnableUserPreferenceAccess returns the EnableUserPreferenceAccess field if non-nil, zero value otherwise.

### GetEnableUserPreferenceAccessOk

`func (o *UserDtoPolicy) GetEnableUserPreferenceAccessOk() (*bool, bool)`

GetEnableUserPreferenceAccessOk returns a tuple with the EnableUserPreferenceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserPreferenceAccess

`func (o *UserDtoPolicy) SetEnableUserPreferenceAccess(v bool)`

SetEnableUserPreferenceAccess sets EnableUserPreferenceAccess field to given value.

### HasEnableUserPreferenceAccess

`func (o *UserDtoPolicy) HasEnableUserPreferenceAccess() bool`

HasEnableUserPreferenceAccess returns a boolean if a field has been set.

### GetAccessSchedules

`func (o *UserDtoPolicy) GetAccessSchedules() []AccessSchedule`

GetAccessSchedules returns the AccessSchedules field if non-nil, zero value otherwise.

### GetAccessSchedulesOk

`func (o *UserDtoPolicy) GetAccessSchedulesOk() (*[]AccessSchedule, bool)`

GetAccessSchedulesOk returns a tuple with the AccessSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSchedules

`func (o *UserDtoPolicy) SetAccessSchedules(v []AccessSchedule)`

SetAccessSchedules sets AccessSchedules field to given value.

### HasAccessSchedules

`func (o *UserDtoPolicy) HasAccessSchedules() bool`

HasAccessSchedules returns a boolean if a field has been set.

### SetAccessSchedulesNil

`func (o *UserDtoPolicy) SetAccessSchedulesNil(b bool)`

 SetAccessSchedulesNil sets the value for AccessSchedules to be an explicit nil

### UnsetAccessSchedules
`func (o *UserDtoPolicy) UnsetAccessSchedules()`

UnsetAccessSchedules ensures that no value is present for AccessSchedules, not even an explicit nil
### GetBlockUnratedItems

`func (o *UserDtoPolicy) GetBlockUnratedItems() []UnratedItem`

GetBlockUnratedItems returns the BlockUnratedItems field if non-nil, zero value otherwise.

### GetBlockUnratedItemsOk

`func (o *UserDtoPolicy) GetBlockUnratedItemsOk() (*[]UnratedItem, bool)`

GetBlockUnratedItemsOk returns a tuple with the BlockUnratedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockUnratedItems

`func (o *UserDtoPolicy) SetBlockUnratedItems(v []UnratedItem)`

SetBlockUnratedItems sets BlockUnratedItems field to given value.

### HasBlockUnratedItems

`func (o *UserDtoPolicy) HasBlockUnratedItems() bool`

HasBlockUnratedItems returns a boolean if a field has been set.

### SetBlockUnratedItemsNil

`func (o *UserDtoPolicy) SetBlockUnratedItemsNil(b bool)`

 SetBlockUnratedItemsNil sets the value for BlockUnratedItems to be an explicit nil

### UnsetBlockUnratedItems
`func (o *UserDtoPolicy) UnsetBlockUnratedItems()`

UnsetBlockUnratedItems ensures that no value is present for BlockUnratedItems, not even an explicit nil
### GetEnableRemoteControlOfOtherUsers

`func (o *UserDtoPolicy) GetEnableRemoteControlOfOtherUsers() bool`

GetEnableRemoteControlOfOtherUsers returns the EnableRemoteControlOfOtherUsers field if non-nil, zero value otherwise.

### GetEnableRemoteControlOfOtherUsersOk

`func (o *UserDtoPolicy) GetEnableRemoteControlOfOtherUsersOk() (*bool, bool)`

GetEnableRemoteControlOfOtherUsersOk returns a tuple with the EnableRemoteControlOfOtherUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteControlOfOtherUsers

`func (o *UserDtoPolicy) SetEnableRemoteControlOfOtherUsers(v bool)`

SetEnableRemoteControlOfOtherUsers sets EnableRemoteControlOfOtherUsers field to given value.

### HasEnableRemoteControlOfOtherUsers

`func (o *UserDtoPolicy) HasEnableRemoteControlOfOtherUsers() bool`

HasEnableRemoteControlOfOtherUsers returns a boolean if a field has been set.

### GetEnableSharedDeviceControl

`func (o *UserDtoPolicy) GetEnableSharedDeviceControl() bool`

GetEnableSharedDeviceControl returns the EnableSharedDeviceControl field if non-nil, zero value otherwise.

### GetEnableSharedDeviceControlOk

`func (o *UserDtoPolicy) GetEnableSharedDeviceControlOk() (*bool, bool)`

GetEnableSharedDeviceControlOk returns a tuple with the EnableSharedDeviceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSharedDeviceControl

`func (o *UserDtoPolicy) SetEnableSharedDeviceControl(v bool)`

SetEnableSharedDeviceControl sets EnableSharedDeviceControl field to given value.

### HasEnableSharedDeviceControl

`func (o *UserDtoPolicy) HasEnableSharedDeviceControl() bool`

HasEnableSharedDeviceControl returns a boolean if a field has been set.

### GetEnableRemoteAccess

`func (o *UserDtoPolicy) GetEnableRemoteAccess() bool`

GetEnableRemoteAccess returns the EnableRemoteAccess field if non-nil, zero value otherwise.

### GetEnableRemoteAccessOk

`func (o *UserDtoPolicy) GetEnableRemoteAccessOk() (*bool, bool)`

GetEnableRemoteAccessOk returns a tuple with the EnableRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteAccess

`func (o *UserDtoPolicy) SetEnableRemoteAccess(v bool)`

SetEnableRemoteAccess sets EnableRemoteAccess field to given value.

### HasEnableRemoteAccess

`func (o *UserDtoPolicy) HasEnableRemoteAccess() bool`

HasEnableRemoteAccess returns a boolean if a field has been set.

### GetEnableLiveTvManagement

`func (o *UserDtoPolicy) GetEnableLiveTvManagement() bool`

GetEnableLiveTvManagement returns the EnableLiveTvManagement field if non-nil, zero value otherwise.

### GetEnableLiveTvManagementOk

`func (o *UserDtoPolicy) GetEnableLiveTvManagementOk() (*bool, bool)`

GetEnableLiveTvManagementOk returns a tuple with the EnableLiveTvManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLiveTvManagement

`func (o *UserDtoPolicy) SetEnableLiveTvManagement(v bool)`

SetEnableLiveTvManagement sets EnableLiveTvManagement field to given value.

### HasEnableLiveTvManagement

`func (o *UserDtoPolicy) HasEnableLiveTvManagement() bool`

HasEnableLiveTvManagement returns a boolean if a field has been set.

### GetEnableLiveTvAccess

`func (o *UserDtoPolicy) GetEnableLiveTvAccess() bool`

GetEnableLiveTvAccess returns the EnableLiveTvAccess field if non-nil, zero value otherwise.

### GetEnableLiveTvAccessOk

`func (o *UserDtoPolicy) GetEnableLiveTvAccessOk() (*bool, bool)`

GetEnableLiveTvAccessOk returns a tuple with the EnableLiveTvAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLiveTvAccess

`func (o *UserDtoPolicy) SetEnableLiveTvAccess(v bool)`

SetEnableLiveTvAccess sets EnableLiveTvAccess field to given value.

### HasEnableLiveTvAccess

`func (o *UserDtoPolicy) HasEnableLiveTvAccess() bool`

HasEnableLiveTvAccess returns a boolean if a field has been set.

### GetEnableMediaPlayback

`func (o *UserDtoPolicy) GetEnableMediaPlayback() bool`

GetEnableMediaPlayback returns the EnableMediaPlayback field if non-nil, zero value otherwise.

### GetEnableMediaPlaybackOk

`func (o *UserDtoPolicy) GetEnableMediaPlaybackOk() (*bool, bool)`

GetEnableMediaPlaybackOk returns a tuple with the EnableMediaPlayback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMediaPlayback

`func (o *UserDtoPolicy) SetEnableMediaPlayback(v bool)`

SetEnableMediaPlayback sets EnableMediaPlayback field to given value.

### HasEnableMediaPlayback

`func (o *UserDtoPolicy) HasEnableMediaPlayback() bool`

HasEnableMediaPlayback returns a boolean if a field has been set.

### GetEnableAudioPlaybackTranscoding

`func (o *UserDtoPolicy) GetEnableAudioPlaybackTranscoding() bool`

GetEnableAudioPlaybackTranscoding returns the EnableAudioPlaybackTranscoding field if non-nil, zero value otherwise.

### GetEnableAudioPlaybackTranscodingOk

`func (o *UserDtoPolicy) GetEnableAudioPlaybackTranscodingOk() (*bool, bool)`

GetEnableAudioPlaybackTranscodingOk returns a tuple with the EnableAudioPlaybackTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAudioPlaybackTranscoding

`func (o *UserDtoPolicy) SetEnableAudioPlaybackTranscoding(v bool)`

SetEnableAudioPlaybackTranscoding sets EnableAudioPlaybackTranscoding field to given value.

### HasEnableAudioPlaybackTranscoding

`func (o *UserDtoPolicy) HasEnableAudioPlaybackTranscoding() bool`

HasEnableAudioPlaybackTranscoding returns a boolean if a field has been set.

### GetEnableVideoPlaybackTranscoding

`func (o *UserDtoPolicy) GetEnableVideoPlaybackTranscoding() bool`

GetEnableVideoPlaybackTranscoding returns the EnableVideoPlaybackTranscoding field if non-nil, zero value otherwise.

### GetEnableVideoPlaybackTranscodingOk

`func (o *UserDtoPolicy) GetEnableVideoPlaybackTranscodingOk() (*bool, bool)`

GetEnableVideoPlaybackTranscodingOk returns a tuple with the EnableVideoPlaybackTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVideoPlaybackTranscoding

`func (o *UserDtoPolicy) SetEnableVideoPlaybackTranscoding(v bool)`

SetEnableVideoPlaybackTranscoding sets EnableVideoPlaybackTranscoding field to given value.

### HasEnableVideoPlaybackTranscoding

`func (o *UserDtoPolicy) HasEnableVideoPlaybackTranscoding() bool`

HasEnableVideoPlaybackTranscoding returns a boolean if a field has been set.

### GetEnablePlaybackRemuxing

`func (o *UserDtoPolicy) GetEnablePlaybackRemuxing() bool`

GetEnablePlaybackRemuxing returns the EnablePlaybackRemuxing field if non-nil, zero value otherwise.

### GetEnablePlaybackRemuxingOk

`func (o *UserDtoPolicy) GetEnablePlaybackRemuxingOk() (*bool, bool)`

GetEnablePlaybackRemuxingOk returns a tuple with the EnablePlaybackRemuxing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePlaybackRemuxing

`func (o *UserDtoPolicy) SetEnablePlaybackRemuxing(v bool)`

SetEnablePlaybackRemuxing sets EnablePlaybackRemuxing field to given value.

### HasEnablePlaybackRemuxing

`func (o *UserDtoPolicy) HasEnablePlaybackRemuxing() bool`

HasEnablePlaybackRemuxing returns a boolean if a field has been set.

### GetForceRemoteSourceTranscoding

`func (o *UserDtoPolicy) GetForceRemoteSourceTranscoding() bool`

GetForceRemoteSourceTranscoding returns the ForceRemoteSourceTranscoding field if non-nil, zero value otherwise.

### GetForceRemoteSourceTranscodingOk

`func (o *UserDtoPolicy) GetForceRemoteSourceTranscodingOk() (*bool, bool)`

GetForceRemoteSourceTranscodingOk returns a tuple with the ForceRemoteSourceTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRemoteSourceTranscoding

`func (o *UserDtoPolicy) SetForceRemoteSourceTranscoding(v bool)`

SetForceRemoteSourceTranscoding sets ForceRemoteSourceTranscoding field to given value.

### HasForceRemoteSourceTranscoding

`func (o *UserDtoPolicy) HasForceRemoteSourceTranscoding() bool`

HasForceRemoteSourceTranscoding returns a boolean if a field has been set.

### GetEnableContentDeletion

`func (o *UserDtoPolicy) GetEnableContentDeletion() bool`

GetEnableContentDeletion returns the EnableContentDeletion field if non-nil, zero value otherwise.

### GetEnableContentDeletionOk

`func (o *UserDtoPolicy) GetEnableContentDeletionOk() (*bool, bool)`

GetEnableContentDeletionOk returns a tuple with the EnableContentDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContentDeletion

`func (o *UserDtoPolicy) SetEnableContentDeletion(v bool)`

SetEnableContentDeletion sets EnableContentDeletion field to given value.

### HasEnableContentDeletion

`func (o *UserDtoPolicy) HasEnableContentDeletion() bool`

HasEnableContentDeletion returns a boolean if a field has been set.

### GetEnableContentDeletionFromFolders

`func (o *UserDtoPolicy) GetEnableContentDeletionFromFolders() []string`

GetEnableContentDeletionFromFolders returns the EnableContentDeletionFromFolders field if non-nil, zero value otherwise.

### GetEnableContentDeletionFromFoldersOk

`func (o *UserDtoPolicy) GetEnableContentDeletionFromFoldersOk() (*[]string, bool)`

GetEnableContentDeletionFromFoldersOk returns a tuple with the EnableContentDeletionFromFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContentDeletionFromFolders

`func (o *UserDtoPolicy) SetEnableContentDeletionFromFolders(v []string)`

SetEnableContentDeletionFromFolders sets EnableContentDeletionFromFolders field to given value.

### HasEnableContentDeletionFromFolders

`func (o *UserDtoPolicy) HasEnableContentDeletionFromFolders() bool`

HasEnableContentDeletionFromFolders returns a boolean if a field has been set.

### SetEnableContentDeletionFromFoldersNil

`func (o *UserDtoPolicy) SetEnableContentDeletionFromFoldersNil(b bool)`

 SetEnableContentDeletionFromFoldersNil sets the value for EnableContentDeletionFromFolders to be an explicit nil

### UnsetEnableContentDeletionFromFolders
`func (o *UserDtoPolicy) UnsetEnableContentDeletionFromFolders()`

UnsetEnableContentDeletionFromFolders ensures that no value is present for EnableContentDeletionFromFolders, not even an explicit nil
### GetEnableContentDownloading

`func (o *UserDtoPolicy) GetEnableContentDownloading() bool`

GetEnableContentDownloading returns the EnableContentDownloading field if non-nil, zero value otherwise.

### GetEnableContentDownloadingOk

`func (o *UserDtoPolicy) GetEnableContentDownloadingOk() (*bool, bool)`

GetEnableContentDownloadingOk returns a tuple with the EnableContentDownloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContentDownloading

`func (o *UserDtoPolicy) SetEnableContentDownloading(v bool)`

SetEnableContentDownloading sets EnableContentDownloading field to given value.

### HasEnableContentDownloading

`func (o *UserDtoPolicy) HasEnableContentDownloading() bool`

HasEnableContentDownloading returns a boolean if a field has been set.

### GetEnableSyncTranscoding

`func (o *UserDtoPolicy) GetEnableSyncTranscoding() bool`

GetEnableSyncTranscoding returns the EnableSyncTranscoding field if non-nil, zero value otherwise.

### GetEnableSyncTranscodingOk

`func (o *UserDtoPolicy) GetEnableSyncTranscodingOk() (*bool, bool)`

GetEnableSyncTranscodingOk returns a tuple with the EnableSyncTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSyncTranscoding

`func (o *UserDtoPolicy) SetEnableSyncTranscoding(v bool)`

SetEnableSyncTranscoding sets EnableSyncTranscoding field to given value.

### HasEnableSyncTranscoding

`func (o *UserDtoPolicy) HasEnableSyncTranscoding() bool`

HasEnableSyncTranscoding returns a boolean if a field has been set.

### GetEnableMediaConversion

`func (o *UserDtoPolicy) GetEnableMediaConversion() bool`

GetEnableMediaConversion returns the EnableMediaConversion field if non-nil, zero value otherwise.

### GetEnableMediaConversionOk

`func (o *UserDtoPolicy) GetEnableMediaConversionOk() (*bool, bool)`

GetEnableMediaConversionOk returns a tuple with the EnableMediaConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMediaConversion

`func (o *UserDtoPolicy) SetEnableMediaConversion(v bool)`

SetEnableMediaConversion sets EnableMediaConversion field to given value.

### HasEnableMediaConversion

`func (o *UserDtoPolicy) HasEnableMediaConversion() bool`

HasEnableMediaConversion returns a boolean if a field has been set.

### GetEnabledDevices

`func (o *UserDtoPolicy) GetEnabledDevices() []string`

GetEnabledDevices returns the EnabledDevices field if non-nil, zero value otherwise.

### GetEnabledDevicesOk

`func (o *UserDtoPolicy) GetEnabledDevicesOk() (*[]string, bool)`

GetEnabledDevicesOk returns a tuple with the EnabledDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledDevices

`func (o *UserDtoPolicy) SetEnabledDevices(v []string)`

SetEnabledDevices sets EnabledDevices field to given value.

### HasEnabledDevices

`func (o *UserDtoPolicy) HasEnabledDevices() bool`

HasEnabledDevices returns a boolean if a field has been set.

### SetEnabledDevicesNil

`func (o *UserDtoPolicy) SetEnabledDevicesNil(b bool)`

 SetEnabledDevicesNil sets the value for EnabledDevices to be an explicit nil

### UnsetEnabledDevices
`func (o *UserDtoPolicy) UnsetEnabledDevices()`

UnsetEnabledDevices ensures that no value is present for EnabledDevices, not even an explicit nil
### GetEnableAllDevices

`func (o *UserDtoPolicy) GetEnableAllDevices() bool`

GetEnableAllDevices returns the EnableAllDevices field if non-nil, zero value otherwise.

### GetEnableAllDevicesOk

`func (o *UserDtoPolicy) GetEnableAllDevicesOk() (*bool, bool)`

GetEnableAllDevicesOk returns a tuple with the EnableAllDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllDevices

`func (o *UserDtoPolicy) SetEnableAllDevices(v bool)`

SetEnableAllDevices sets EnableAllDevices field to given value.

### HasEnableAllDevices

`func (o *UserDtoPolicy) HasEnableAllDevices() bool`

HasEnableAllDevices returns a boolean if a field has been set.

### GetEnabledChannels

`func (o *UserDtoPolicy) GetEnabledChannels() []string`

GetEnabledChannels returns the EnabledChannels field if non-nil, zero value otherwise.

### GetEnabledChannelsOk

`func (o *UserDtoPolicy) GetEnabledChannelsOk() (*[]string, bool)`

GetEnabledChannelsOk returns a tuple with the EnabledChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledChannels

`func (o *UserDtoPolicy) SetEnabledChannels(v []string)`

SetEnabledChannels sets EnabledChannels field to given value.

### HasEnabledChannels

`func (o *UserDtoPolicy) HasEnabledChannels() bool`

HasEnabledChannels returns a boolean if a field has been set.

### SetEnabledChannelsNil

`func (o *UserDtoPolicy) SetEnabledChannelsNil(b bool)`

 SetEnabledChannelsNil sets the value for EnabledChannels to be an explicit nil

### UnsetEnabledChannels
`func (o *UserDtoPolicy) UnsetEnabledChannels()`

UnsetEnabledChannels ensures that no value is present for EnabledChannels, not even an explicit nil
### GetEnableAllChannels

`func (o *UserDtoPolicy) GetEnableAllChannels() bool`

GetEnableAllChannels returns the EnableAllChannels field if non-nil, zero value otherwise.

### GetEnableAllChannelsOk

`func (o *UserDtoPolicy) GetEnableAllChannelsOk() (*bool, bool)`

GetEnableAllChannelsOk returns a tuple with the EnableAllChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllChannels

`func (o *UserDtoPolicy) SetEnableAllChannels(v bool)`

SetEnableAllChannels sets EnableAllChannels field to given value.

### HasEnableAllChannels

`func (o *UserDtoPolicy) HasEnableAllChannels() bool`

HasEnableAllChannels returns a boolean if a field has been set.

### GetEnabledFolders

`func (o *UserDtoPolicy) GetEnabledFolders() []string`

GetEnabledFolders returns the EnabledFolders field if non-nil, zero value otherwise.

### GetEnabledFoldersOk

`func (o *UserDtoPolicy) GetEnabledFoldersOk() (*[]string, bool)`

GetEnabledFoldersOk returns a tuple with the EnabledFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFolders

`func (o *UserDtoPolicy) SetEnabledFolders(v []string)`

SetEnabledFolders sets EnabledFolders field to given value.

### HasEnabledFolders

`func (o *UserDtoPolicy) HasEnabledFolders() bool`

HasEnabledFolders returns a boolean if a field has been set.

### SetEnabledFoldersNil

`func (o *UserDtoPolicy) SetEnabledFoldersNil(b bool)`

 SetEnabledFoldersNil sets the value for EnabledFolders to be an explicit nil

### UnsetEnabledFolders
`func (o *UserDtoPolicy) UnsetEnabledFolders()`

UnsetEnabledFolders ensures that no value is present for EnabledFolders, not even an explicit nil
### GetEnableAllFolders

`func (o *UserDtoPolicy) GetEnableAllFolders() bool`

GetEnableAllFolders returns the EnableAllFolders field if non-nil, zero value otherwise.

### GetEnableAllFoldersOk

`func (o *UserDtoPolicy) GetEnableAllFoldersOk() (*bool, bool)`

GetEnableAllFoldersOk returns a tuple with the EnableAllFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllFolders

`func (o *UserDtoPolicy) SetEnableAllFolders(v bool)`

SetEnableAllFolders sets EnableAllFolders field to given value.

### HasEnableAllFolders

`func (o *UserDtoPolicy) HasEnableAllFolders() bool`

HasEnableAllFolders returns a boolean if a field has been set.

### GetInvalidLoginAttemptCount

`func (o *UserDtoPolicy) GetInvalidLoginAttemptCount() int32`

GetInvalidLoginAttemptCount returns the InvalidLoginAttemptCount field if non-nil, zero value otherwise.

### GetInvalidLoginAttemptCountOk

`func (o *UserDtoPolicy) GetInvalidLoginAttemptCountOk() (*int32, bool)`

GetInvalidLoginAttemptCountOk returns a tuple with the InvalidLoginAttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidLoginAttemptCount

`func (o *UserDtoPolicy) SetInvalidLoginAttemptCount(v int32)`

SetInvalidLoginAttemptCount sets InvalidLoginAttemptCount field to given value.

### HasInvalidLoginAttemptCount

`func (o *UserDtoPolicy) HasInvalidLoginAttemptCount() bool`

HasInvalidLoginAttemptCount returns a boolean if a field has been set.

### GetLoginAttemptsBeforeLockout

`func (o *UserDtoPolicy) GetLoginAttemptsBeforeLockout() int32`

GetLoginAttemptsBeforeLockout returns the LoginAttemptsBeforeLockout field if non-nil, zero value otherwise.

### GetLoginAttemptsBeforeLockoutOk

`func (o *UserDtoPolicy) GetLoginAttemptsBeforeLockoutOk() (*int32, bool)`

GetLoginAttemptsBeforeLockoutOk returns a tuple with the LoginAttemptsBeforeLockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAttemptsBeforeLockout

`func (o *UserDtoPolicy) SetLoginAttemptsBeforeLockout(v int32)`

SetLoginAttemptsBeforeLockout sets LoginAttemptsBeforeLockout field to given value.

### HasLoginAttemptsBeforeLockout

`func (o *UserDtoPolicy) HasLoginAttemptsBeforeLockout() bool`

HasLoginAttemptsBeforeLockout returns a boolean if a field has been set.

### GetMaxActiveSessions

`func (o *UserDtoPolicy) GetMaxActiveSessions() int32`

GetMaxActiveSessions returns the MaxActiveSessions field if non-nil, zero value otherwise.

### GetMaxActiveSessionsOk

`func (o *UserDtoPolicy) GetMaxActiveSessionsOk() (*int32, bool)`

GetMaxActiveSessionsOk returns a tuple with the MaxActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActiveSessions

`func (o *UserDtoPolicy) SetMaxActiveSessions(v int32)`

SetMaxActiveSessions sets MaxActiveSessions field to given value.

### HasMaxActiveSessions

`func (o *UserDtoPolicy) HasMaxActiveSessions() bool`

HasMaxActiveSessions returns a boolean if a field has been set.

### GetEnablePublicSharing

`func (o *UserDtoPolicy) GetEnablePublicSharing() bool`

GetEnablePublicSharing returns the EnablePublicSharing field if non-nil, zero value otherwise.

### GetEnablePublicSharingOk

`func (o *UserDtoPolicy) GetEnablePublicSharingOk() (*bool, bool)`

GetEnablePublicSharingOk returns a tuple with the EnablePublicSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublicSharing

`func (o *UserDtoPolicy) SetEnablePublicSharing(v bool)`

SetEnablePublicSharing sets EnablePublicSharing field to given value.

### HasEnablePublicSharing

`func (o *UserDtoPolicy) HasEnablePublicSharing() bool`

HasEnablePublicSharing returns a boolean if a field has been set.

### GetBlockedMediaFolders

`func (o *UserDtoPolicy) GetBlockedMediaFolders() []string`

GetBlockedMediaFolders returns the BlockedMediaFolders field if non-nil, zero value otherwise.

### GetBlockedMediaFoldersOk

`func (o *UserDtoPolicy) GetBlockedMediaFoldersOk() (*[]string, bool)`

GetBlockedMediaFoldersOk returns a tuple with the BlockedMediaFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedMediaFolders

`func (o *UserDtoPolicy) SetBlockedMediaFolders(v []string)`

SetBlockedMediaFolders sets BlockedMediaFolders field to given value.

### HasBlockedMediaFolders

`func (o *UserDtoPolicy) HasBlockedMediaFolders() bool`

HasBlockedMediaFolders returns a boolean if a field has been set.

### SetBlockedMediaFoldersNil

`func (o *UserDtoPolicy) SetBlockedMediaFoldersNil(b bool)`

 SetBlockedMediaFoldersNil sets the value for BlockedMediaFolders to be an explicit nil

### UnsetBlockedMediaFolders
`func (o *UserDtoPolicy) UnsetBlockedMediaFolders()`

UnsetBlockedMediaFolders ensures that no value is present for BlockedMediaFolders, not even an explicit nil
### GetBlockedChannels

`func (o *UserDtoPolicy) GetBlockedChannels() []string`

GetBlockedChannels returns the BlockedChannels field if non-nil, zero value otherwise.

### GetBlockedChannelsOk

`func (o *UserDtoPolicy) GetBlockedChannelsOk() (*[]string, bool)`

GetBlockedChannelsOk returns a tuple with the BlockedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedChannels

`func (o *UserDtoPolicy) SetBlockedChannels(v []string)`

SetBlockedChannels sets BlockedChannels field to given value.

### HasBlockedChannels

`func (o *UserDtoPolicy) HasBlockedChannels() bool`

HasBlockedChannels returns a boolean if a field has been set.

### SetBlockedChannelsNil

`func (o *UserDtoPolicy) SetBlockedChannelsNil(b bool)`

 SetBlockedChannelsNil sets the value for BlockedChannels to be an explicit nil

### UnsetBlockedChannels
`func (o *UserDtoPolicy) UnsetBlockedChannels()`

UnsetBlockedChannels ensures that no value is present for BlockedChannels, not even an explicit nil
### GetRemoteClientBitrateLimit

`func (o *UserDtoPolicy) GetRemoteClientBitrateLimit() int32`

GetRemoteClientBitrateLimit returns the RemoteClientBitrateLimit field if non-nil, zero value otherwise.

### GetRemoteClientBitrateLimitOk

`func (o *UserDtoPolicy) GetRemoteClientBitrateLimitOk() (*int32, bool)`

GetRemoteClientBitrateLimitOk returns a tuple with the RemoteClientBitrateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClientBitrateLimit

`func (o *UserDtoPolicy) SetRemoteClientBitrateLimit(v int32)`

SetRemoteClientBitrateLimit sets RemoteClientBitrateLimit field to given value.

### HasRemoteClientBitrateLimit

`func (o *UserDtoPolicy) HasRemoteClientBitrateLimit() bool`

HasRemoteClientBitrateLimit returns a boolean if a field has been set.

### GetAuthenticationProviderId

`func (o *UserDtoPolicy) GetAuthenticationProviderId() string`

GetAuthenticationProviderId returns the AuthenticationProviderId field if non-nil, zero value otherwise.

### GetAuthenticationProviderIdOk

`func (o *UserDtoPolicy) GetAuthenticationProviderIdOk() (*string, bool)`

GetAuthenticationProviderIdOk returns a tuple with the AuthenticationProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProviderId

`func (o *UserDtoPolicy) SetAuthenticationProviderId(v string)`

SetAuthenticationProviderId sets AuthenticationProviderId field to given value.


### GetPasswordResetProviderId

`func (o *UserDtoPolicy) GetPasswordResetProviderId() string`

GetPasswordResetProviderId returns the PasswordResetProviderId field if non-nil, zero value otherwise.

### GetPasswordResetProviderIdOk

`func (o *UserDtoPolicy) GetPasswordResetProviderIdOk() (*string, bool)`

GetPasswordResetProviderIdOk returns a tuple with the PasswordResetProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetProviderId

`func (o *UserDtoPolicy) SetPasswordResetProviderId(v string)`

SetPasswordResetProviderId sets PasswordResetProviderId field to given value.


### GetSyncPlayAccess

`func (o *UserDtoPolicy) GetSyncPlayAccess() SyncPlayUserAccessType`

GetSyncPlayAccess returns the SyncPlayAccess field if non-nil, zero value otherwise.

### GetSyncPlayAccessOk

`func (o *UserDtoPolicy) GetSyncPlayAccessOk() (*SyncPlayUserAccessType, bool)`

GetSyncPlayAccessOk returns a tuple with the SyncPlayAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPlayAccess

`func (o *UserDtoPolicy) SetSyncPlayAccess(v SyncPlayUserAccessType)`

SetSyncPlayAccess sets SyncPlayAccess field to given value.

### HasSyncPlayAccess

`func (o *UserDtoPolicy) HasSyncPlayAccess() bool`

HasSyncPlayAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


