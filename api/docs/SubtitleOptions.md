# SubtitleOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipIfEmbeddedSubtitlesPresent** | Pointer to **bool** |  | [optional] 
**SkipIfAudioTrackMatches** | Pointer to **bool** |  | [optional] 
**DownloadLanguages** | Pointer to **[]string** |  | [optional] 
**DownloadMovieSubtitles** | Pointer to **bool** |  | [optional] 
**DownloadEpisodeSubtitles** | Pointer to **bool** |  | [optional] 
**OpenSubtitlesUsername** | Pointer to **NullableString** |  | [optional] 
**OpenSubtitlesPasswordHash** | Pointer to **NullableString** |  | [optional] 
**IsOpenSubtitleVipAccount** | Pointer to **bool** |  | [optional] 
**RequirePerfectMatch** | Pointer to **bool** |  | [optional] 

## Methods

### NewSubtitleOptions

`func NewSubtitleOptions() *SubtitleOptions`

NewSubtitleOptions instantiates a new SubtitleOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubtitleOptionsWithDefaults

`func NewSubtitleOptionsWithDefaults() *SubtitleOptions`

NewSubtitleOptionsWithDefaults instantiates a new SubtitleOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipIfEmbeddedSubtitlesPresent

`func (o *SubtitleOptions) GetSkipIfEmbeddedSubtitlesPresent() bool`

GetSkipIfEmbeddedSubtitlesPresent returns the SkipIfEmbeddedSubtitlesPresent field if non-nil, zero value otherwise.

### GetSkipIfEmbeddedSubtitlesPresentOk

`func (o *SubtitleOptions) GetSkipIfEmbeddedSubtitlesPresentOk() (*bool, bool)`

GetSkipIfEmbeddedSubtitlesPresentOk returns a tuple with the SkipIfEmbeddedSubtitlesPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIfEmbeddedSubtitlesPresent

`func (o *SubtitleOptions) SetSkipIfEmbeddedSubtitlesPresent(v bool)`

SetSkipIfEmbeddedSubtitlesPresent sets SkipIfEmbeddedSubtitlesPresent field to given value.

### HasSkipIfEmbeddedSubtitlesPresent

`func (o *SubtitleOptions) HasSkipIfEmbeddedSubtitlesPresent() bool`

HasSkipIfEmbeddedSubtitlesPresent returns a boolean if a field has been set.

### GetSkipIfAudioTrackMatches

`func (o *SubtitleOptions) GetSkipIfAudioTrackMatches() bool`

GetSkipIfAudioTrackMatches returns the SkipIfAudioTrackMatches field if non-nil, zero value otherwise.

### GetSkipIfAudioTrackMatchesOk

`func (o *SubtitleOptions) GetSkipIfAudioTrackMatchesOk() (*bool, bool)`

GetSkipIfAudioTrackMatchesOk returns a tuple with the SkipIfAudioTrackMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipIfAudioTrackMatches

`func (o *SubtitleOptions) SetSkipIfAudioTrackMatches(v bool)`

SetSkipIfAudioTrackMatches sets SkipIfAudioTrackMatches field to given value.

### HasSkipIfAudioTrackMatches

`func (o *SubtitleOptions) HasSkipIfAudioTrackMatches() bool`

HasSkipIfAudioTrackMatches returns a boolean if a field has been set.

### GetDownloadLanguages

`func (o *SubtitleOptions) GetDownloadLanguages() []string`

GetDownloadLanguages returns the DownloadLanguages field if non-nil, zero value otherwise.

### GetDownloadLanguagesOk

`func (o *SubtitleOptions) GetDownloadLanguagesOk() (*[]string, bool)`

GetDownloadLanguagesOk returns a tuple with the DownloadLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadLanguages

`func (o *SubtitleOptions) SetDownloadLanguages(v []string)`

SetDownloadLanguages sets DownloadLanguages field to given value.

### HasDownloadLanguages

`func (o *SubtitleOptions) HasDownloadLanguages() bool`

HasDownloadLanguages returns a boolean if a field has been set.

### SetDownloadLanguagesNil

`func (o *SubtitleOptions) SetDownloadLanguagesNil(b bool)`

 SetDownloadLanguagesNil sets the value for DownloadLanguages to be an explicit nil

### UnsetDownloadLanguages
`func (o *SubtitleOptions) UnsetDownloadLanguages()`

UnsetDownloadLanguages ensures that no value is present for DownloadLanguages, not even an explicit nil
### GetDownloadMovieSubtitles

`func (o *SubtitleOptions) GetDownloadMovieSubtitles() bool`

GetDownloadMovieSubtitles returns the DownloadMovieSubtitles field if non-nil, zero value otherwise.

### GetDownloadMovieSubtitlesOk

`func (o *SubtitleOptions) GetDownloadMovieSubtitlesOk() (*bool, bool)`

GetDownloadMovieSubtitlesOk returns a tuple with the DownloadMovieSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadMovieSubtitles

`func (o *SubtitleOptions) SetDownloadMovieSubtitles(v bool)`

SetDownloadMovieSubtitles sets DownloadMovieSubtitles field to given value.

### HasDownloadMovieSubtitles

`func (o *SubtitleOptions) HasDownloadMovieSubtitles() bool`

HasDownloadMovieSubtitles returns a boolean if a field has been set.

### GetDownloadEpisodeSubtitles

`func (o *SubtitleOptions) GetDownloadEpisodeSubtitles() bool`

GetDownloadEpisodeSubtitles returns the DownloadEpisodeSubtitles field if non-nil, zero value otherwise.

### GetDownloadEpisodeSubtitlesOk

`func (o *SubtitleOptions) GetDownloadEpisodeSubtitlesOk() (*bool, bool)`

GetDownloadEpisodeSubtitlesOk returns a tuple with the DownloadEpisodeSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadEpisodeSubtitles

`func (o *SubtitleOptions) SetDownloadEpisodeSubtitles(v bool)`

SetDownloadEpisodeSubtitles sets DownloadEpisodeSubtitles field to given value.

### HasDownloadEpisodeSubtitles

`func (o *SubtitleOptions) HasDownloadEpisodeSubtitles() bool`

HasDownloadEpisodeSubtitles returns a boolean if a field has been set.

### GetOpenSubtitlesUsername

`func (o *SubtitleOptions) GetOpenSubtitlesUsername() string`

GetOpenSubtitlesUsername returns the OpenSubtitlesUsername field if non-nil, zero value otherwise.

### GetOpenSubtitlesUsernameOk

`func (o *SubtitleOptions) GetOpenSubtitlesUsernameOk() (*string, bool)`

GetOpenSubtitlesUsernameOk returns a tuple with the OpenSubtitlesUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenSubtitlesUsername

`func (o *SubtitleOptions) SetOpenSubtitlesUsername(v string)`

SetOpenSubtitlesUsername sets OpenSubtitlesUsername field to given value.

### HasOpenSubtitlesUsername

`func (o *SubtitleOptions) HasOpenSubtitlesUsername() bool`

HasOpenSubtitlesUsername returns a boolean if a field has been set.

### SetOpenSubtitlesUsernameNil

`func (o *SubtitleOptions) SetOpenSubtitlesUsernameNil(b bool)`

 SetOpenSubtitlesUsernameNil sets the value for OpenSubtitlesUsername to be an explicit nil

### UnsetOpenSubtitlesUsername
`func (o *SubtitleOptions) UnsetOpenSubtitlesUsername()`

UnsetOpenSubtitlesUsername ensures that no value is present for OpenSubtitlesUsername, not even an explicit nil
### GetOpenSubtitlesPasswordHash

`func (o *SubtitleOptions) GetOpenSubtitlesPasswordHash() string`

GetOpenSubtitlesPasswordHash returns the OpenSubtitlesPasswordHash field if non-nil, zero value otherwise.

### GetOpenSubtitlesPasswordHashOk

`func (o *SubtitleOptions) GetOpenSubtitlesPasswordHashOk() (*string, bool)`

GetOpenSubtitlesPasswordHashOk returns a tuple with the OpenSubtitlesPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenSubtitlesPasswordHash

`func (o *SubtitleOptions) SetOpenSubtitlesPasswordHash(v string)`

SetOpenSubtitlesPasswordHash sets OpenSubtitlesPasswordHash field to given value.

### HasOpenSubtitlesPasswordHash

`func (o *SubtitleOptions) HasOpenSubtitlesPasswordHash() bool`

HasOpenSubtitlesPasswordHash returns a boolean if a field has been set.

### SetOpenSubtitlesPasswordHashNil

`func (o *SubtitleOptions) SetOpenSubtitlesPasswordHashNil(b bool)`

 SetOpenSubtitlesPasswordHashNil sets the value for OpenSubtitlesPasswordHash to be an explicit nil

### UnsetOpenSubtitlesPasswordHash
`func (o *SubtitleOptions) UnsetOpenSubtitlesPasswordHash()`

UnsetOpenSubtitlesPasswordHash ensures that no value is present for OpenSubtitlesPasswordHash, not even an explicit nil
### GetIsOpenSubtitleVipAccount

`func (o *SubtitleOptions) GetIsOpenSubtitleVipAccount() bool`

GetIsOpenSubtitleVipAccount returns the IsOpenSubtitleVipAccount field if non-nil, zero value otherwise.

### GetIsOpenSubtitleVipAccountOk

`func (o *SubtitleOptions) GetIsOpenSubtitleVipAccountOk() (*bool, bool)`

GetIsOpenSubtitleVipAccountOk returns a tuple with the IsOpenSubtitleVipAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpenSubtitleVipAccount

`func (o *SubtitleOptions) SetIsOpenSubtitleVipAccount(v bool)`

SetIsOpenSubtitleVipAccount sets IsOpenSubtitleVipAccount field to given value.

### HasIsOpenSubtitleVipAccount

`func (o *SubtitleOptions) HasIsOpenSubtitleVipAccount() bool`

HasIsOpenSubtitleVipAccount returns a boolean if a field has been set.

### GetRequirePerfectMatch

`func (o *SubtitleOptions) GetRequirePerfectMatch() bool`

GetRequirePerfectMatch returns the RequirePerfectMatch field if non-nil, zero value otherwise.

### GetRequirePerfectMatchOk

`func (o *SubtitleOptions) GetRequirePerfectMatchOk() (*bool, bool)`

GetRequirePerfectMatchOk returns a tuple with the RequirePerfectMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePerfectMatch

`func (o *SubtitleOptions) SetRequirePerfectMatch(v bool)`

SetRequirePerfectMatch sets RequirePerfectMatch field to given value.

### HasRequirePerfectMatch

`func (o *SubtitleOptions) HasRequirePerfectMatch() bool`

HasRequirePerfectMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


