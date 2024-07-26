# UserPolicy

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
**SyncPlayAccess** | Pointer to [**SyncPlayUserAccessType**](SyncPlayUserAccessType.md) | Enum SyncPlayUserAccessType. | [optional] 

## Methods

### NewUserPolicy

`func NewUserPolicy(authenticationProviderId string, passwordResetProviderId string, ) *UserPolicy`

NewUserPolicy instantiates a new UserPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPolicyWithDefaults

`func NewUserPolicyWithDefaults() *UserPolicy`

NewUserPolicyWithDefaults instantiates a new UserPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAdministrator

`func (o *UserPolicy) GetIsAdministrator() bool`

GetIsAdministrator returns the IsAdministrator field if non-nil, zero value otherwise.

### GetIsAdministratorOk

`func (o *UserPolicy) GetIsAdministratorOk() (*bool, bool)`

GetIsAdministratorOk returns a tuple with the IsAdministrator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdministrator

`func (o *UserPolicy) SetIsAdministrator(v bool)`

SetIsAdministrator sets IsAdministrator field to given value.

### HasIsAdministrator

`func (o *UserPolicy) HasIsAdministrator() bool`

HasIsAdministrator returns a boolean if a field has been set.

### GetIsHidden

`func (o *UserPolicy) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *UserPolicy) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *UserPolicy) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *UserPolicy) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetEnableCollectionManagement

`func (o *UserPolicy) GetEnableCollectionManagement() bool`

GetEnableCollectionManagement returns the EnableCollectionManagement field if non-nil, zero value otherwise.

### GetEnableCollectionManagementOk

`func (o *UserPolicy) GetEnableCollectionManagementOk() (*bool, bool)`

GetEnableCollectionManagementOk returns a tuple with the EnableCollectionManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCollectionManagement

`func (o *UserPolicy) SetEnableCollectionManagement(v bool)`

SetEnableCollectionManagement sets EnableCollectionManagement field to given value.

### HasEnableCollectionManagement

`func (o *UserPolicy) HasEnableCollectionManagement() bool`

HasEnableCollectionManagement returns a boolean if a field has been set.

### GetEnableSubtitleManagement

`func (o *UserPolicy) GetEnableSubtitleManagement() bool`

GetEnableSubtitleManagement returns the EnableSubtitleManagement field if non-nil, zero value otherwise.

### GetEnableSubtitleManagementOk

`func (o *UserPolicy) GetEnableSubtitleManagementOk() (*bool, bool)`

GetEnableSubtitleManagementOk returns a tuple with the EnableSubtitleManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubtitleManagement

`func (o *UserPolicy) SetEnableSubtitleManagement(v bool)`

SetEnableSubtitleManagement sets EnableSubtitleManagement field to given value.

### HasEnableSubtitleManagement

`func (o *UserPolicy) HasEnableSubtitleManagement() bool`

HasEnableSubtitleManagement returns a boolean if a field has been set.

### GetEnableLyricManagement

`func (o *UserPolicy) GetEnableLyricManagement() bool`

GetEnableLyricManagement returns the EnableLyricManagement field if non-nil, zero value otherwise.

### GetEnableLyricManagementOk

`func (o *UserPolicy) GetEnableLyricManagementOk() (*bool, bool)`

GetEnableLyricManagementOk returns a tuple with the EnableLyricManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLyricManagement

`func (o *UserPolicy) SetEnableLyricManagement(v bool)`

SetEnableLyricManagement sets EnableLyricManagement field to given value.

### HasEnableLyricManagement

`func (o *UserPolicy) HasEnableLyricManagement() bool`

HasEnableLyricManagement returns a boolean if a field has been set.

### GetIsDisabled

`func (o *UserPolicy) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *UserPolicy) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *UserPolicy) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *UserPolicy) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetMaxParentalRating

`func (o *UserPolicy) GetMaxParentalRating() int32`

GetMaxParentalRating returns the MaxParentalRating field if non-nil, zero value otherwise.

### GetMaxParentalRatingOk

`func (o *UserPolicy) GetMaxParentalRatingOk() (*int32, bool)`

GetMaxParentalRatingOk returns a tuple with the MaxParentalRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParentalRating

`func (o *UserPolicy) SetMaxParentalRating(v int32)`

SetMaxParentalRating sets MaxParentalRating field to given value.

### HasMaxParentalRating

`func (o *UserPolicy) HasMaxParentalRating() bool`

HasMaxParentalRating returns a boolean if a field has been set.

### SetMaxParentalRatingNil

`func (o *UserPolicy) SetMaxParentalRatingNil(b bool)`

 SetMaxParentalRatingNil sets the value for MaxParentalRating to be an explicit nil

### UnsetMaxParentalRating
`func (o *UserPolicy) UnsetMaxParentalRating()`

UnsetMaxParentalRating ensures that no value is present for MaxParentalRating, not even an explicit nil
### GetBlockedTags

`func (o *UserPolicy) GetBlockedTags() []string`

GetBlockedTags returns the BlockedTags field if non-nil, zero value otherwise.

### GetBlockedTagsOk

`func (o *UserPolicy) GetBlockedTagsOk() (*[]string, bool)`

GetBlockedTagsOk returns a tuple with the BlockedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedTags

`func (o *UserPolicy) SetBlockedTags(v []string)`

SetBlockedTags sets BlockedTags field to given value.

### HasBlockedTags

`func (o *UserPolicy) HasBlockedTags() bool`

HasBlockedTags returns a boolean if a field has been set.

### SetBlockedTagsNil

`func (o *UserPolicy) SetBlockedTagsNil(b bool)`

 SetBlockedTagsNil sets the value for BlockedTags to be an explicit nil

### UnsetBlockedTags
`func (o *UserPolicy) UnsetBlockedTags()`

UnsetBlockedTags ensures that no value is present for BlockedTags, not even an explicit nil
### GetAllowedTags

`func (o *UserPolicy) GetAllowedTags() []string`

GetAllowedTags returns the AllowedTags field if non-nil, zero value otherwise.

### GetAllowedTagsOk

`func (o *UserPolicy) GetAllowedTagsOk() (*[]string, bool)`

GetAllowedTagsOk returns a tuple with the AllowedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTags

`func (o *UserPolicy) SetAllowedTags(v []string)`

SetAllowedTags sets AllowedTags field to given value.

### HasAllowedTags

`func (o *UserPolicy) HasAllowedTags() bool`

HasAllowedTags returns a boolean if a field has been set.

### SetAllowedTagsNil

`func (o *UserPolicy) SetAllowedTagsNil(b bool)`

 SetAllowedTagsNil sets the value for AllowedTags to be an explicit nil

### UnsetAllowedTags
`func (o *UserPolicy) UnsetAllowedTags()`

UnsetAllowedTags ensures that no value is present for AllowedTags, not even an explicit nil
### GetEnableUserPreferenceAccess

`func (o *UserPolicy) GetEnableUserPreferenceAccess() bool`

GetEnableUserPreferenceAccess returns the EnableUserPreferenceAccess field if non-nil, zero value otherwise.

### GetEnableUserPreferenceAccessOk

`func (o *UserPolicy) GetEnableUserPreferenceAccessOk() (*bool, bool)`

GetEnableUserPreferenceAccessOk returns a tuple with the EnableUserPreferenceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserPreferenceAccess

`func (o *UserPolicy) SetEnableUserPreferenceAccess(v bool)`

SetEnableUserPreferenceAccess sets EnableUserPreferenceAccess field to given value.

### HasEnableUserPreferenceAccess

`func (o *UserPolicy) HasEnableUserPreferenceAccess() bool`

HasEnableUserPreferenceAccess returns a boolean if a field has been set.

### GetAccessSchedules

`func (o *UserPolicy) GetAccessSchedules() []AccessSchedule`

GetAccessSchedules returns the AccessSchedules field if non-nil, zero value otherwise.

### GetAccessSchedulesOk

`func (o *UserPolicy) GetAccessSchedulesOk() (*[]AccessSchedule, bool)`

GetAccessSchedulesOk returns a tuple with the AccessSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessSchedules

`func (o *UserPolicy) SetAccessSchedules(v []AccessSchedule)`

SetAccessSchedules sets AccessSchedules field to given value.

### HasAccessSchedules

`func (o *UserPolicy) HasAccessSchedules() bool`

HasAccessSchedules returns a boolean if a field has been set.

### SetAccessSchedulesNil

`func (o *UserPolicy) SetAccessSchedulesNil(b bool)`

 SetAccessSchedulesNil sets the value for AccessSchedules to be an explicit nil

### UnsetAccessSchedules
`func (o *UserPolicy) UnsetAccessSchedules()`

UnsetAccessSchedules ensures that no value is present for AccessSchedules, not even an explicit nil
### GetBlockUnratedItems

`func (o *UserPolicy) GetBlockUnratedItems() []UnratedItem`

GetBlockUnratedItems returns the BlockUnratedItems field if non-nil, zero value otherwise.

### GetBlockUnratedItemsOk

`func (o *UserPolicy) GetBlockUnratedItemsOk() (*[]UnratedItem, bool)`

GetBlockUnratedItemsOk returns a tuple with the BlockUnratedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockUnratedItems

`func (o *UserPolicy) SetBlockUnratedItems(v []UnratedItem)`

SetBlockUnratedItems sets BlockUnratedItems field to given value.

### HasBlockUnratedItems

`func (o *UserPolicy) HasBlockUnratedItems() bool`

HasBlockUnratedItems returns a boolean if a field has been set.

### SetBlockUnratedItemsNil

`func (o *UserPolicy) SetBlockUnratedItemsNil(b bool)`

 SetBlockUnratedItemsNil sets the value for BlockUnratedItems to be an explicit nil

### UnsetBlockUnratedItems
`func (o *UserPolicy) UnsetBlockUnratedItems()`

UnsetBlockUnratedItems ensures that no value is present for BlockUnratedItems, not even an explicit nil
### GetEnableRemoteControlOfOtherUsers

`func (o *UserPolicy) GetEnableRemoteControlOfOtherUsers() bool`

GetEnableRemoteControlOfOtherUsers returns the EnableRemoteControlOfOtherUsers field if non-nil, zero value otherwise.

### GetEnableRemoteControlOfOtherUsersOk

`func (o *UserPolicy) GetEnableRemoteControlOfOtherUsersOk() (*bool, bool)`

GetEnableRemoteControlOfOtherUsersOk returns a tuple with the EnableRemoteControlOfOtherUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteControlOfOtherUsers

`func (o *UserPolicy) SetEnableRemoteControlOfOtherUsers(v bool)`

SetEnableRemoteControlOfOtherUsers sets EnableRemoteControlOfOtherUsers field to given value.

### HasEnableRemoteControlOfOtherUsers

`func (o *UserPolicy) HasEnableRemoteControlOfOtherUsers() bool`

HasEnableRemoteControlOfOtherUsers returns a boolean if a field has been set.

### GetEnableSharedDeviceControl

`func (o *UserPolicy) GetEnableSharedDeviceControl() bool`

GetEnableSharedDeviceControl returns the EnableSharedDeviceControl field if non-nil, zero value otherwise.

### GetEnableSharedDeviceControlOk

`func (o *UserPolicy) GetEnableSharedDeviceControlOk() (*bool, bool)`

GetEnableSharedDeviceControlOk returns a tuple with the EnableSharedDeviceControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSharedDeviceControl

`func (o *UserPolicy) SetEnableSharedDeviceControl(v bool)`

SetEnableSharedDeviceControl sets EnableSharedDeviceControl field to given value.

### HasEnableSharedDeviceControl

`func (o *UserPolicy) HasEnableSharedDeviceControl() bool`

HasEnableSharedDeviceControl returns a boolean if a field has been set.

### GetEnableRemoteAccess

`func (o *UserPolicy) GetEnableRemoteAccess() bool`

GetEnableRemoteAccess returns the EnableRemoteAccess field if non-nil, zero value otherwise.

### GetEnableRemoteAccessOk

`func (o *UserPolicy) GetEnableRemoteAccessOk() (*bool, bool)`

GetEnableRemoteAccessOk returns a tuple with the EnableRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteAccess

`func (o *UserPolicy) SetEnableRemoteAccess(v bool)`

SetEnableRemoteAccess sets EnableRemoteAccess field to given value.

### HasEnableRemoteAccess

`func (o *UserPolicy) HasEnableRemoteAccess() bool`

HasEnableRemoteAccess returns a boolean if a field has been set.

### GetEnableLiveTvManagement

`func (o *UserPolicy) GetEnableLiveTvManagement() bool`

GetEnableLiveTvManagement returns the EnableLiveTvManagement field if non-nil, zero value otherwise.

### GetEnableLiveTvManagementOk

`func (o *UserPolicy) GetEnableLiveTvManagementOk() (*bool, bool)`

GetEnableLiveTvManagementOk returns a tuple with the EnableLiveTvManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLiveTvManagement

`func (o *UserPolicy) SetEnableLiveTvManagement(v bool)`

SetEnableLiveTvManagement sets EnableLiveTvManagement field to given value.

### HasEnableLiveTvManagement

`func (o *UserPolicy) HasEnableLiveTvManagement() bool`

HasEnableLiveTvManagement returns a boolean if a field has been set.

### GetEnableLiveTvAccess

`func (o *UserPolicy) GetEnableLiveTvAccess() bool`

GetEnableLiveTvAccess returns the EnableLiveTvAccess field if non-nil, zero value otherwise.

### GetEnableLiveTvAccessOk

`func (o *UserPolicy) GetEnableLiveTvAccessOk() (*bool, bool)`

GetEnableLiveTvAccessOk returns a tuple with the EnableLiveTvAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLiveTvAccess

`func (o *UserPolicy) SetEnableLiveTvAccess(v bool)`

SetEnableLiveTvAccess sets EnableLiveTvAccess field to given value.

### HasEnableLiveTvAccess

`func (o *UserPolicy) HasEnableLiveTvAccess() bool`

HasEnableLiveTvAccess returns a boolean if a field has been set.

### GetEnableMediaPlayback

`func (o *UserPolicy) GetEnableMediaPlayback() bool`

GetEnableMediaPlayback returns the EnableMediaPlayback field if non-nil, zero value otherwise.

### GetEnableMediaPlaybackOk

`func (o *UserPolicy) GetEnableMediaPlaybackOk() (*bool, bool)`

GetEnableMediaPlaybackOk returns a tuple with the EnableMediaPlayback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMediaPlayback

`func (o *UserPolicy) SetEnableMediaPlayback(v bool)`

SetEnableMediaPlayback sets EnableMediaPlayback field to given value.

### HasEnableMediaPlayback

`func (o *UserPolicy) HasEnableMediaPlayback() bool`

HasEnableMediaPlayback returns a boolean if a field has been set.

### GetEnableAudioPlaybackTranscoding

`func (o *UserPolicy) GetEnableAudioPlaybackTranscoding() bool`

GetEnableAudioPlaybackTranscoding returns the EnableAudioPlaybackTranscoding field if non-nil, zero value otherwise.

### GetEnableAudioPlaybackTranscodingOk

`func (o *UserPolicy) GetEnableAudioPlaybackTranscodingOk() (*bool, bool)`

GetEnableAudioPlaybackTranscodingOk returns a tuple with the EnableAudioPlaybackTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAudioPlaybackTranscoding

`func (o *UserPolicy) SetEnableAudioPlaybackTranscoding(v bool)`

SetEnableAudioPlaybackTranscoding sets EnableAudioPlaybackTranscoding field to given value.

### HasEnableAudioPlaybackTranscoding

`func (o *UserPolicy) HasEnableAudioPlaybackTranscoding() bool`

HasEnableAudioPlaybackTranscoding returns a boolean if a field has been set.

### GetEnableVideoPlaybackTranscoding

`func (o *UserPolicy) GetEnableVideoPlaybackTranscoding() bool`

GetEnableVideoPlaybackTranscoding returns the EnableVideoPlaybackTranscoding field if non-nil, zero value otherwise.

### GetEnableVideoPlaybackTranscodingOk

`func (o *UserPolicy) GetEnableVideoPlaybackTranscodingOk() (*bool, bool)`

GetEnableVideoPlaybackTranscodingOk returns a tuple with the EnableVideoPlaybackTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVideoPlaybackTranscoding

`func (o *UserPolicy) SetEnableVideoPlaybackTranscoding(v bool)`

SetEnableVideoPlaybackTranscoding sets EnableVideoPlaybackTranscoding field to given value.

### HasEnableVideoPlaybackTranscoding

`func (o *UserPolicy) HasEnableVideoPlaybackTranscoding() bool`

HasEnableVideoPlaybackTranscoding returns a boolean if a field has been set.

### GetEnablePlaybackRemuxing

`func (o *UserPolicy) GetEnablePlaybackRemuxing() bool`

GetEnablePlaybackRemuxing returns the EnablePlaybackRemuxing field if non-nil, zero value otherwise.

### GetEnablePlaybackRemuxingOk

`func (o *UserPolicy) GetEnablePlaybackRemuxingOk() (*bool, bool)`

GetEnablePlaybackRemuxingOk returns a tuple with the EnablePlaybackRemuxing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePlaybackRemuxing

`func (o *UserPolicy) SetEnablePlaybackRemuxing(v bool)`

SetEnablePlaybackRemuxing sets EnablePlaybackRemuxing field to given value.

### HasEnablePlaybackRemuxing

`func (o *UserPolicy) HasEnablePlaybackRemuxing() bool`

HasEnablePlaybackRemuxing returns a boolean if a field has been set.

### GetForceRemoteSourceTranscoding

`func (o *UserPolicy) GetForceRemoteSourceTranscoding() bool`

GetForceRemoteSourceTranscoding returns the ForceRemoteSourceTranscoding field if non-nil, zero value otherwise.

### GetForceRemoteSourceTranscodingOk

`func (o *UserPolicy) GetForceRemoteSourceTranscodingOk() (*bool, bool)`

GetForceRemoteSourceTranscodingOk returns a tuple with the ForceRemoteSourceTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRemoteSourceTranscoding

`func (o *UserPolicy) SetForceRemoteSourceTranscoding(v bool)`

SetForceRemoteSourceTranscoding sets ForceRemoteSourceTranscoding field to given value.

### HasForceRemoteSourceTranscoding

`func (o *UserPolicy) HasForceRemoteSourceTranscoding() bool`

HasForceRemoteSourceTranscoding returns a boolean if a field has been set.

### GetEnableContentDeletion

`func (o *UserPolicy) GetEnableContentDeletion() bool`

GetEnableContentDeletion returns the EnableContentDeletion field if non-nil, zero value otherwise.

### GetEnableContentDeletionOk

`func (o *UserPolicy) GetEnableContentDeletionOk() (*bool, bool)`

GetEnableContentDeletionOk returns a tuple with the EnableContentDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContentDeletion

`func (o *UserPolicy) SetEnableContentDeletion(v bool)`

SetEnableContentDeletion sets EnableContentDeletion field to given value.

### HasEnableContentDeletion

`func (o *UserPolicy) HasEnableContentDeletion() bool`

HasEnableContentDeletion returns a boolean if a field has been set.

### GetEnableContentDeletionFromFolders

`func (o *UserPolicy) GetEnableContentDeletionFromFolders() []string`

GetEnableContentDeletionFromFolders returns the EnableContentDeletionFromFolders field if non-nil, zero value otherwise.

### GetEnableContentDeletionFromFoldersOk

`func (o *UserPolicy) GetEnableContentDeletionFromFoldersOk() (*[]string, bool)`

GetEnableContentDeletionFromFoldersOk returns a tuple with the EnableContentDeletionFromFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContentDeletionFromFolders

`func (o *UserPolicy) SetEnableContentDeletionFromFolders(v []string)`

SetEnableContentDeletionFromFolders sets EnableContentDeletionFromFolders field to given value.

### HasEnableContentDeletionFromFolders

`func (o *UserPolicy) HasEnableContentDeletionFromFolders() bool`

HasEnableContentDeletionFromFolders returns a boolean if a field has been set.

### SetEnableContentDeletionFromFoldersNil

`func (o *UserPolicy) SetEnableContentDeletionFromFoldersNil(b bool)`

 SetEnableContentDeletionFromFoldersNil sets the value for EnableContentDeletionFromFolders to be an explicit nil

### UnsetEnableContentDeletionFromFolders
`func (o *UserPolicy) UnsetEnableContentDeletionFromFolders()`

UnsetEnableContentDeletionFromFolders ensures that no value is present for EnableContentDeletionFromFolders, not even an explicit nil
### GetEnableContentDownloading

`func (o *UserPolicy) GetEnableContentDownloading() bool`

GetEnableContentDownloading returns the EnableContentDownloading field if non-nil, zero value otherwise.

### GetEnableContentDownloadingOk

`func (o *UserPolicy) GetEnableContentDownloadingOk() (*bool, bool)`

GetEnableContentDownloadingOk returns a tuple with the EnableContentDownloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableContentDownloading

`func (o *UserPolicy) SetEnableContentDownloading(v bool)`

SetEnableContentDownloading sets EnableContentDownloading field to given value.

### HasEnableContentDownloading

`func (o *UserPolicy) HasEnableContentDownloading() bool`

HasEnableContentDownloading returns a boolean if a field has been set.

### GetEnableSyncTranscoding

`func (o *UserPolicy) GetEnableSyncTranscoding() bool`

GetEnableSyncTranscoding returns the EnableSyncTranscoding field if non-nil, zero value otherwise.

### GetEnableSyncTranscodingOk

`func (o *UserPolicy) GetEnableSyncTranscodingOk() (*bool, bool)`

GetEnableSyncTranscodingOk returns a tuple with the EnableSyncTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSyncTranscoding

`func (o *UserPolicy) SetEnableSyncTranscoding(v bool)`

SetEnableSyncTranscoding sets EnableSyncTranscoding field to given value.

### HasEnableSyncTranscoding

`func (o *UserPolicy) HasEnableSyncTranscoding() bool`

HasEnableSyncTranscoding returns a boolean if a field has been set.

### GetEnableMediaConversion

`func (o *UserPolicy) GetEnableMediaConversion() bool`

GetEnableMediaConversion returns the EnableMediaConversion field if non-nil, zero value otherwise.

### GetEnableMediaConversionOk

`func (o *UserPolicy) GetEnableMediaConversionOk() (*bool, bool)`

GetEnableMediaConversionOk returns a tuple with the EnableMediaConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMediaConversion

`func (o *UserPolicy) SetEnableMediaConversion(v bool)`

SetEnableMediaConversion sets EnableMediaConversion field to given value.

### HasEnableMediaConversion

`func (o *UserPolicy) HasEnableMediaConversion() bool`

HasEnableMediaConversion returns a boolean if a field has been set.

### GetEnabledDevices

`func (o *UserPolicy) GetEnabledDevices() []string`

GetEnabledDevices returns the EnabledDevices field if non-nil, zero value otherwise.

### GetEnabledDevicesOk

`func (o *UserPolicy) GetEnabledDevicesOk() (*[]string, bool)`

GetEnabledDevicesOk returns a tuple with the EnabledDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledDevices

`func (o *UserPolicy) SetEnabledDevices(v []string)`

SetEnabledDevices sets EnabledDevices field to given value.

### HasEnabledDevices

`func (o *UserPolicy) HasEnabledDevices() bool`

HasEnabledDevices returns a boolean if a field has been set.

### SetEnabledDevicesNil

`func (o *UserPolicy) SetEnabledDevicesNil(b bool)`

 SetEnabledDevicesNil sets the value for EnabledDevices to be an explicit nil

### UnsetEnabledDevices
`func (o *UserPolicy) UnsetEnabledDevices()`

UnsetEnabledDevices ensures that no value is present for EnabledDevices, not even an explicit nil
### GetEnableAllDevices

`func (o *UserPolicy) GetEnableAllDevices() bool`

GetEnableAllDevices returns the EnableAllDevices field if non-nil, zero value otherwise.

### GetEnableAllDevicesOk

`func (o *UserPolicy) GetEnableAllDevicesOk() (*bool, bool)`

GetEnableAllDevicesOk returns a tuple with the EnableAllDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllDevices

`func (o *UserPolicy) SetEnableAllDevices(v bool)`

SetEnableAllDevices sets EnableAllDevices field to given value.

### HasEnableAllDevices

`func (o *UserPolicy) HasEnableAllDevices() bool`

HasEnableAllDevices returns a boolean if a field has been set.

### GetEnabledChannels

`func (o *UserPolicy) GetEnabledChannels() []string`

GetEnabledChannels returns the EnabledChannels field if non-nil, zero value otherwise.

### GetEnabledChannelsOk

`func (o *UserPolicy) GetEnabledChannelsOk() (*[]string, bool)`

GetEnabledChannelsOk returns a tuple with the EnabledChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledChannels

`func (o *UserPolicy) SetEnabledChannels(v []string)`

SetEnabledChannels sets EnabledChannels field to given value.

### HasEnabledChannels

`func (o *UserPolicy) HasEnabledChannels() bool`

HasEnabledChannels returns a boolean if a field has been set.

### SetEnabledChannelsNil

`func (o *UserPolicy) SetEnabledChannelsNil(b bool)`

 SetEnabledChannelsNil sets the value for EnabledChannels to be an explicit nil

### UnsetEnabledChannels
`func (o *UserPolicy) UnsetEnabledChannels()`

UnsetEnabledChannels ensures that no value is present for EnabledChannels, not even an explicit nil
### GetEnableAllChannels

`func (o *UserPolicy) GetEnableAllChannels() bool`

GetEnableAllChannels returns the EnableAllChannels field if non-nil, zero value otherwise.

### GetEnableAllChannelsOk

`func (o *UserPolicy) GetEnableAllChannelsOk() (*bool, bool)`

GetEnableAllChannelsOk returns a tuple with the EnableAllChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllChannels

`func (o *UserPolicy) SetEnableAllChannels(v bool)`

SetEnableAllChannels sets EnableAllChannels field to given value.

### HasEnableAllChannels

`func (o *UserPolicy) HasEnableAllChannels() bool`

HasEnableAllChannels returns a boolean if a field has been set.

### GetEnabledFolders

`func (o *UserPolicy) GetEnabledFolders() []string`

GetEnabledFolders returns the EnabledFolders field if non-nil, zero value otherwise.

### GetEnabledFoldersOk

`func (o *UserPolicy) GetEnabledFoldersOk() (*[]string, bool)`

GetEnabledFoldersOk returns a tuple with the EnabledFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFolders

`func (o *UserPolicy) SetEnabledFolders(v []string)`

SetEnabledFolders sets EnabledFolders field to given value.

### HasEnabledFolders

`func (o *UserPolicy) HasEnabledFolders() bool`

HasEnabledFolders returns a boolean if a field has been set.

### SetEnabledFoldersNil

`func (o *UserPolicy) SetEnabledFoldersNil(b bool)`

 SetEnabledFoldersNil sets the value for EnabledFolders to be an explicit nil

### UnsetEnabledFolders
`func (o *UserPolicy) UnsetEnabledFolders()`

UnsetEnabledFolders ensures that no value is present for EnabledFolders, not even an explicit nil
### GetEnableAllFolders

`func (o *UserPolicy) GetEnableAllFolders() bool`

GetEnableAllFolders returns the EnableAllFolders field if non-nil, zero value otherwise.

### GetEnableAllFoldersOk

`func (o *UserPolicy) GetEnableAllFoldersOk() (*bool, bool)`

GetEnableAllFoldersOk returns a tuple with the EnableAllFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllFolders

`func (o *UserPolicy) SetEnableAllFolders(v bool)`

SetEnableAllFolders sets EnableAllFolders field to given value.

### HasEnableAllFolders

`func (o *UserPolicy) HasEnableAllFolders() bool`

HasEnableAllFolders returns a boolean if a field has been set.

### GetInvalidLoginAttemptCount

`func (o *UserPolicy) GetInvalidLoginAttemptCount() int32`

GetInvalidLoginAttemptCount returns the InvalidLoginAttemptCount field if non-nil, zero value otherwise.

### GetInvalidLoginAttemptCountOk

`func (o *UserPolicy) GetInvalidLoginAttemptCountOk() (*int32, bool)`

GetInvalidLoginAttemptCountOk returns a tuple with the InvalidLoginAttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidLoginAttemptCount

`func (o *UserPolicy) SetInvalidLoginAttemptCount(v int32)`

SetInvalidLoginAttemptCount sets InvalidLoginAttemptCount field to given value.

### HasInvalidLoginAttemptCount

`func (o *UserPolicy) HasInvalidLoginAttemptCount() bool`

HasInvalidLoginAttemptCount returns a boolean if a field has been set.

### GetLoginAttemptsBeforeLockout

`func (o *UserPolicy) GetLoginAttemptsBeforeLockout() int32`

GetLoginAttemptsBeforeLockout returns the LoginAttemptsBeforeLockout field if non-nil, zero value otherwise.

### GetLoginAttemptsBeforeLockoutOk

`func (o *UserPolicy) GetLoginAttemptsBeforeLockoutOk() (*int32, bool)`

GetLoginAttemptsBeforeLockoutOk returns a tuple with the LoginAttemptsBeforeLockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginAttemptsBeforeLockout

`func (o *UserPolicy) SetLoginAttemptsBeforeLockout(v int32)`

SetLoginAttemptsBeforeLockout sets LoginAttemptsBeforeLockout field to given value.

### HasLoginAttemptsBeforeLockout

`func (o *UserPolicy) HasLoginAttemptsBeforeLockout() bool`

HasLoginAttemptsBeforeLockout returns a boolean if a field has been set.

### GetMaxActiveSessions

`func (o *UserPolicy) GetMaxActiveSessions() int32`

GetMaxActiveSessions returns the MaxActiveSessions field if non-nil, zero value otherwise.

### GetMaxActiveSessionsOk

`func (o *UserPolicy) GetMaxActiveSessionsOk() (*int32, bool)`

GetMaxActiveSessionsOk returns a tuple with the MaxActiveSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxActiveSessions

`func (o *UserPolicy) SetMaxActiveSessions(v int32)`

SetMaxActiveSessions sets MaxActiveSessions field to given value.

### HasMaxActiveSessions

`func (o *UserPolicy) HasMaxActiveSessions() bool`

HasMaxActiveSessions returns a boolean if a field has been set.

### GetEnablePublicSharing

`func (o *UserPolicy) GetEnablePublicSharing() bool`

GetEnablePublicSharing returns the EnablePublicSharing field if non-nil, zero value otherwise.

### GetEnablePublicSharingOk

`func (o *UserPolicy) GetEnablePublicSharingOk() (*bool, bool)`

GetEnablePublicSharingOk returns a tuple with the EnablePublicSharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePublicSharing

`func (o *UserPolicy) SetEnablePublicSharing(v bool)`

SetEnablePublicSharing sets EnablePublicSharing field to given value.

### HasEnablePublicSharing

`func (o *UserPolicy) HasEnablePublicSharing() bool`

HasEnablePublicSharing returns a boolean if a field has been set.

### GetBlockedMediaFolders

`func (o *UserPolicy) GetBlockedMediaFolders() []string`

GetBlockedMediaFolders returns the BlockedMediaFolders field if non-nil, zero value otherwise.

### GetBlockedMediaFoldersOk

`func (o *UserPolicy) GetBlockedMediaFoldersOk() (*[]string, bool)`

GetBlockedMediaFoldersOk returns a tuple with the BlockedMediaFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedMediaFolders

`func (o *UserPolicy) SetBlockedMediaFolders(v []string)`

SetBlockedMediaFolders sets BlockedMediaFolders field to given value.

### HasBlockedMediaFolders

`func (o *UserPolicy) HasBlockedMediaFolders() bool`

HasBlockedMediaFolders returns a boolean if a field has been set.

### SetBlockedMediaFoldersNil

`func (o *UserPolicy) SetBlockedMediaFoldersNil(b bool)`

 SetBlockedMediaFoldersNil sets the value for BlockedMediaFolders to be an explicit nil

### UnsetBlockedMediaFolders
`func (o *UserPolicy) UnsetBlockedMediaFolders()`

UnsetBlockedMediaFolders ensures that no value is present for BlockedMediaFolders, not even an explicit nil
### GetBlockedChannels

`func (o *UserPolicy) GetBlockedChannels() []string`

GetBlockedChannels returns the BlockedChannels field if non-nil, zero value otherwise.

### GetBlockedChannelsOk

`func (o *UserPolicy) GetBlockedChannelsOk() (*[]string, bool)`

GetBlockedChannelsOk returns a tuple with the BlockedChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedChannels

`func (o *UserPolicy) SetBlockedChannels(v []string)`

SetBlockedChannels sets BlockedChannels field to given value.

### HasBlockedChannels

`func (o *UserPolicy) HasBlockedChannels() bool`

HasBlockedChannels returns a boolean if a field has been set.

### SetBlockedChannelsNil

`func (o *UserPolicy) SetBlockedChannelsNil(b bool)`

 SetBlockedChannelsNil sets the value for BlockedChannels to be an explicit nil

### UnsetBlockedChannels
`func (o *UserPolicy) UnsetBlockedChannels()`

UnsetBlockedChannels ensures that no value is present for BlockedChannels, not even an explicit nil
### GetRemoteClientBitrateLimit

`func (o *UserPolicy) GetRemoteClientBitrateLimit() int32`

GetRemoteClientBitrateLimit returns the RemoteClientBitrateLimit field if non-nil, zero value otherwise.

### GetRemoteClientBitrateLimitOk

`func (o *UserPolicy) GetRemoteClientBitrateLimitOk() (*int32, bool)`

GetRemoteClientBitrateLimitOk returns a tuple with the RemoteClientBitrateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClientBitrateLimit

`func (o *UserPolicy) SetRemoteClientBitrateLimit(v int32)`

SetRemoteClientBitrateLimit sets RemoteClientBitrateLimit field to given value.

### HasRemoteClientBitrateLimit

`func (o *UserPolicy) HasRemoteClientBitrateLimit() bool`

HasRemoteClientBitrateLimit returns a boolean if a field has been set.

### GetAuthenticationProviderId

`func (o *UserPolicy) GetAuthenticationProviderId() string`

GetAuthenticationProviderId returns the AuthenticationProviderId field if non-nil, zero value otherwise.

### GetAuthenticationProviderIdOk

`func (o *UserPolicy) GetAuthenticationProviderIdOk() (*string, bool)`

GetAuthenticationProviderIdOk returns a tuple with the AuthenticationProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProviderId

`func (o *UserPolicy) SetAuthenticationProviderId(v string)`

SetAuthenticationProviderId sets AuthenticationProviderId field to given value.


### GetPasswordResetProviderId

`func (o *UserPolicy) GetPasswordResetProviderId() string`

GetPasswordResetProviderId returns the PasswordResetProviderId field if non-nil, zero value otherwise.

### GetPasswordResetProviderIdOk

`func (o *UserPolicy) GetPasswordResetProviderIdOk() (*string, bool)`

GetPasswordResetProviderIdOk returns a tuple with the PasswordResetProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordResetProviderId

`func (o *UserPolicy) SetPasswordResetProviderId(v string)`

SetPasswordResetProviderId sets PasswordResetProviderId field to given value.


### GetSyncPlayAccess

`func (o *UserPolicy) GetSyncPlayAccess() SyncPlayUserAccessType`

GetSyncPlayAccess returns the SyncPlayAccess field if non-nil, zero value otherwise.

### GetSyncPlayAccessOk

`func (o *UserPolicy) GetSyncPlayAccessOk() (*SyncPlayUserAccessType, bool)`

GetSyncPlayAccessOk returns a tuple with the SyncPlayAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncPlayAccess

`func (o *UserPolicy) SetSyncPlayAccess(v SyncPlayUserAccessType)`

SetSyncPlayAccess sets SyncPlayAccess field to given value.

### HasSyncPlayAccess

`func (o *UserPolicy) HasSyncPlayAccess() bool`

HasSyncPlayAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


