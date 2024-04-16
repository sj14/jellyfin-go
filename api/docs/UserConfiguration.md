# UserConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioLanguagePreference** | Pointer to **NullableString** | Gets or sets the audio language preference. | [optional] 
**PlayDefaultAudioTrack** | Pointer to **bool** | Gets or sets a value indicating whether [play default audio track]. | [optional] 
**SubtitleLanguagePreference** | Pointer to **NullableString** | Gets or sets the subtitle language preference. | [optional] 
**DisplayMissingEpisodes** | Pointer to **bool** |  | [optional] 
**GroupedFolders** | Pointer to **[]string** |  | [optional] 
**SubtitleMode** | Pointer to [**SubtitlePlaybackMode**](SubtitlePlaybackMode.md) | An enum representing a subtitle playback mode. | [optional] 
**DisplayCollectionsView** | Pointer to **bool** |  | [optional] 
**EnableLocalPassword** | Pointer to **bool** |  | [optional] 
**OrderedViews** | Pointer to **[]string** |  | [optional] 
**LatestItemsExcludes** | Pointer to **[]string** |  | [optional] 
**MyMediaExcludes** | Pointer to **[]string** |  | [optional] 
**HidePlayedInLatest** | Pointer to **bool** |  | [optional] 
**RememberAudioSelections** | Pointer to **bool** |  | [optional] 
**RememberSubtitleSelections** | Pointer to **bool** |  | [optional] 
**EnableNextEpisodeAutoPlay** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserConfiguration

`func NewUserConfiguration() *UserConfiguration`

NewUserConfiguration instantiates a new UserConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConfigurationWithDefaults

`func NewUserConfigurationWithDefaults() *UserConfiguration`

NewUserConfigurationWithDefaults instantiates a new UserConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioLanguagePreference

`func (o *UserConfiguration) GetAudioLanguagePreference() string`

GetAudioLanguagePreference returns the AudioLanguagePreference field if non-nil, zero value otherwise.

### GetAudioLanguagePreferenceOk

`func (o *UserConfiguration) GetAudioLanguagePreferenceOk() (*string, bool)`

GetAudioLanguagePreferenceOk returns a tuple with the AudioLanguagePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioLanguagePreference

`func (o *UserConfiguration) SetAudioLanguagePreference(v string)`

SetAudioLanguagePreference sets AudioLanguagePreference field to given value.

### HasAudioLanguagePreference

`func (o *UserConfiguration) HasAudioLanguagePreference() bool`

HasAudioLanguagePreference returns a boolean if a field has been set.

### SetAudioLanguagePreferenceNil

`func (o *UserConfiguration) SetAudioLanguagePreferenceNil(b bool)`

 SetAudioLanguagePreferenceNil sets the value for AudioLanguagePreference to be an explicit nil

### UnsetAudioLanguagePreference
`func (o *UserConfiguration) UnsetAudioLanguagePreference()`

UnsetAudioLanguagePreference ensures that no value is present for AudioLanguagePreference, not even an explicit nil
### GetPlayDefaultAudioTrack

`func (o *UserConfiguration) GetPlayDefaultAudioTrack() bool`

GetPlayDefaultAudioTrack returns the PlayDefaultAudioTrack field if non-nil, zero value otherwise.

### GetPlayDefaultAudioTrackOk

`func (o *UserConfiguration) GetPlayDefaultAudioTrackOk() (*bool, bool)`

GetPlayDefaultAudioTrackOk returns a tuple with the PlayDefaultAudioTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayDefaultAudioTrack

`func (o *UserConfiguration) SetPlayDefaultAudioTrack(v bool)`

SetPlayDefaultAudioTrack sets PlayDefaultAudioTrack field to given value.

### HasPlayDefaultAudioTrack

`func (o *UserConfiguration) HasPlayDefaultAudioTrack() bool`

HasPlayDefaultAudioTrack returns a boolean if a field has been set.

### GetSubtitleLanguagePreference

`func (o *UserConfiguration) GetSubtitleLanguagePreference() string`

GetSubtitleLanguagePreference returns the SubtitleLanguagePreference field if non-nil, zero value otherwise.

### GetSubtitleLanguagePreferenceOk

`func (o *UserConfiguration) GetSubtitleLanguagePreferenceOk() (*string, bool)`

GetSubtitleLanguagePreferenceOk returns a tuple with the SubtitleLanguagePreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleLanguagePreference

`func (o *UserConfiguration) SetSubtitleLanguagePreference(v string)`

SetSubtitleLanguagePreference sets SubtitleLanguagePreference field to given value.

### HasSubtitleLanguagePreference

`func (o *UserConfiguration) HasSubtitleLanguagePreference() bool`

HasSubtitleLanguagePreference returns a boolean if a field has been set.

### SetSubtitleLanguagePreferenceNil

`func (o *UserConfiguration) SetSubtitleLanguagePreferenceNil(b bool)`

 SetSubtitleLanguagePreferenceNil sets the value for SubtitleLanguagePreference to be an explicit nil

### UnsetSubtitleLanguagePreference
`func (o *UserConfiguration) UnsetSubtitleLanguagePreference()`

UnsetSubtitleLanguagePreference ensures that no value is present for SubtitleLanguagePreference, not even an explicit nil
### GetDisplayMissingEpisodes

`func (o *UserConfiguration) GetDisplayMissingEpisodes() bool`

GetDisplayMissingEpisodes returns the DisplayMissingEpisodes field if non-nil, zero value otherwise.

### GetDisplayMissingEpisodesOk

`func (o *UserConfiguration) GetDisplayMissingEpisodesOk() (*bool, bool)`

GetDisplayMissingEpisodesOk returns a tuple with the DisplayMissingEpisodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMissingEpisodes

`func (o *UserConfiguration) SetDisplayMissingEpisodes(v bool)`

SetDisplayMissingEpisodes sets DisplayMissingEpisodes field to given value.

### HasDisplayMissingEpisodes

`func (o *UserConfiguration) HasDisplayMissingEpisodes() bool`

HasDisplayMissingEpisodes returns a boolean if a field has been set.

### GetGroupedFolders

`func (o *UserConfiguration) GetGroupedFolders() []string`

GetGroupedFolders returns the GroupedFolders field if non-nil, zero value otherwise.

### GetGroupedFoldersOk

`func (o *UserConfiguration) GetGroupedFoldersOk() (*[]string, bool)`

GetGroupedFoldersOk returns a tuple with the GroupedFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedFolders

`func (o *UserConfiguration) SetGroupedFolders(v []string)`

SetGroupedFolders sets GroupedFolders field to given value.

### HasGroupedFolders

`func (o *UserConfiguration) HasGroupedFolders() bool`

HasGroupedFolders returns a boolean if a field has been set.

### GetSubtitleMode

`func (o *UserConfiguration) GetSubtitleMode() SubtitlePlaybackMode`

GetSubtitleMode returns the SubtitleMode field if non-nil, zero value otherwise.

### GetSubtitleModeOk

`func (o *UserConfiguration) GetSubtitleModeOk() (*SubtitlePlaybackMode, bool)`

GetSubtitleModeOk returns a tuple with the SubtitleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleMode

`func (o *UserConfiguration) SetSubtitleMode(v SubtitlePlaybackMode)`

SetSubtitleMode sets SubtitleMode field to given value.

### HasSubtitleMode

`func (o *UserConfiguration) HasSubtitleMode() bool`

HasSubtitleMode returns a boolean if a field has been set.

### GetDisplayCollectionsView

`func (o *UserConfiguration) GetDisplayCollectionsView() bool`

GetDisplayCollectionsView returns the DisplayCollectionsView field if non-nil, zero value otherwise.

### GetDisplayCollectionsViewOk

`func (o *UserConfiguration) GetDisplayCollectionsViewOk() (*bool, bool)`

GetDisplayCollectionsViewOk returns a tuple with the DisplayCollectionsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayCollectionsView

`func (o *UserConfiguration) SetDisplayCollectionsView(v bool)`

SetDisplayCollectionsView sets DisplayCollectionsView field to given value.

### HasDisplayCollectionsView

`func (o *UserConfiguration) HasDisplayCollectionsView() bool`

HasDisplayCollectionsView returns a boolean if a field has been set.

### GetEnableLocalPassword

`func (o *UserConfiguration) GetEnableLocalPassword() bool`

GetEnableLocalPassword returns the EnableLocalPassword field if non-nil, zero value otherwise.

### GetEnableLocalPasswordOk

`func (o *UserConfiguration) GetEnableLocalPasswordOk() (*bool, bool)`

GetEnableLocalPasswordOk returns a tuple with the EnableLocalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLocalPassword

`func (o *UserConfiguration) SetEnableLocalPassword(v bool)`

SetEnableLocalPassword sets EnableLocalPassword field to given value.

### HasEnableLocalPassword

`func (o *UserConfiguration) HasEnableLocalPassword() bool`

HasEnableLocalPassword returns a boolean if a field has been set.

### GetOrderedViews

`func (o *UserConfiguration) GetOrderedViews() []string`

GetOrderedViews returns the OrderedViews field if non-nil, zero value otherwise.

### GetOrderedViewsOk

`func (o *UserConfiguration) GetOrderedViewsOk() (*[]string, bool)`

GetOrderedViewsOk returns a tuple with the OrderedViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderedViews

`func (o *UserConfiguration) SetOrderedViews(v []string)`

SetOrderedViews sets OrderedViews field to given value.

### HasOrderedViews

`func (o *UserConfiguration) HasOrderedViews() bool`

HasOrderedViews returns a boolean if a field has been set.

### GetLatestItemsExcludes

`func (o *UserConfiguration) GetLatestItemsExcludes() []string`

GetLatestItemsExcludes returns the LatestItemsExcludes field if non-nil, zero value otherwise.

### GetLatestItemsExcludesOk

`func (o *UserConfiguration) GetLatestItemsExcludesOk() (*[]string, bool)`

GetLatestItemsExcludesOk returns a tuple with the LatestItemsExcludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestItemsExcludes

`func (o *UserConfiguration) SetLatestItemsExcludes(v []string)`

SetLatestItemsExcludes sets LatestItemsExcludes field to given value.

### HasLatestItemsExcludes

`func (o *UserConfiguration) HasLatestItemsExcludes() bool`

HasLatestItemsExcludes returns a boolean if a field has been set.

### GetMyMediaExcludes

`func (o *UserConfiguration) GetMyMediaExcludes() []string`

GetMyMediaExcludes returns the MyMediaExcludes field if non-nil, zero value otherwise.

### GetMyMediaExcludesOk

`func (o *UserConfiguration) GetMyMediaExcludesOk() (*[]string, bool)`

GetMyMediaExcludesOk returns a tuple with the MyMediaExcludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyMediaExcludes

`func (o *UserConfiguration) SetMyMediaExcludes(v []string)`

SetMyMediaExcludes sets MyMediaExcludes field to given value.

### HasMyMediaExcludes

`func (o *UserConfiguration) HasMyMediaExcludes() bool`

HasMyMediaExcludes returns a boolean if a field has been set.

### GetHidePlayedInLatest

`func (o *UserConfiguration) GetHidePlayedInLatest() bool`

GetHidePlayedInLatest returns the HidePlayedInLatest field if non-nil, zero value otherwise.

### GetHidePlayedInLatestOk

`func (o *UserConfiguration) GetHidePlayedInLatestOk() (*bool, bool)`

GetHidePlayedInLatestOk returns a tuple with the HidePlayedInLatest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePlayedInLatest

`func (o *UserConfiguration) SetHidePlayedInLatest(v bool)`

SetHidePlayedInLatest sets HidePlayedInLatest field to given value.

### HasHidePlayedInLatest

`func (o *UserConfiguration) HasHidePlayedInLatest() bool`

HasHidePlayedInLatest returns a boolean if a field has been set.

### GetRememberAudioSelections

`func (o *UserConfiguration) GetRememberAudioSelections() bool`

GetRememberAudioSelections returns the RememberAudioSelections field if non-nil, zero value otherwise.

### GetRememberAudioSelectionsOk

`func (o *UserConfiguration) GetRememberAudioSelectionsOk() (*bool, bool)`

GetRememberAudioSelectionsOk returns a tuple with the RememberAudioSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberAudioSelections

`func (o *UserConfiguration) SetRememberAudioSelections(v bool)`

SetRememberAudioSelections sets RememberAudioSelections field to given value.

### HasRememberAudioSelections

`func (o *UserConfiguration) HasRememberAudioSelections() bool`

HasRememberAudioSelections returns a boolean if a field has been set.

### GetRememberSubtitleSelections

`func (o *UserConfiguration) GetRememberSubtitleSelections() bool`

GetRememberSubtitleSelections returns the RememberSubtitleSelections field if non-nil, zero value otherwise.

### GetRememberSubtitleSelectionsOk

`func (o *UserConfiguration) GetRememberSubtitleSelectionsOk() (*bool, bool)`

GetRememberSubtitleSelectionsOk returns a tuple with the RememberSubtitleSelections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberSubtitleSelections

`func (o *UserConfiguration) SetRememberSubtitleSelections(v bool)`

SetRememberSubtitleSelections sets RememberSubtitleSelections field to given value.

### HasRememberSubtitleSelections

`func (o *UserConfiguration) HasRememberSubtitleSelections() bool`

HasRememberSubtitleSelections returns a boolean if a field has been set.

### GetEnableNextEpisodeAutoPlay

`func (o *UserConfiguration) GetEnableNextEpisodeAutoPlay() bool`

GetEnableNextEpisodeAutoPlay returns the EnableNextEpisodeAutoPlay field if non-nil, zero value otherwise.

### GetEnableNextEpisodeAutoPlayOk

`func (o *UserConfiguration) GetEnableNextEpisodeAutoPlayOk() (*bool, bool)`

GetEnableNextEpisodeAutoPlayOk returns a tuple with the EnableNextEpisodeAutoPlay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNextEpisodeAutoPlay

`func (o *UserConfiguration) SetEnableNextEpisodeAutoPlay(v bool)`

SetEnableNextEpisodeAutoPlay sets EnableNextEpisodeAutoPlay field to given value.

### HasEnableNextEpisodeAutoPlay

`func (o *UserConfiguration) HasEnableNextEpisodeAutoPlay() bool`

HasEnableNextEpisodeAutoPlay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


