# AddVirtualFolderDtoLibraryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnablePhotos** | Pointer to **bool** |  | [optional] 
**EnableRealtimeMonitor** | Pointer to **bool** |  | [optional] 
**EnableChapterImageExtraction** | Pointer to **bool** |  | [optional] 
**ExtractChapterImagesDuringLibraryScan** | Pointer to **bool** |  | [optional] 
**PathInfos** | Pointer to [**[]MediaPathInfo**](MediaPathInfo.md) |  | [optional] 
**SaveLocalMetadata** | Pointer to **bool** |  | [optional] 
**EnableInternetProviders** | Pointer to **bool** |  | [optional] 
**EnableAutomaticSeriesGrouping** | Pointer to **bool** |  | [optional] 
**EnableEmbeddedTitles** | Pointer to **bool** |  | [optional] 
**EnableEmbeddedEpisodeInfos** | Pointer to **bool** |  | [optional] 
**AutomaticRefreshIntervalDays** | Pointer to **int32** |  | [optional] 
**PreferredMetadataLanguage** | Pointer to **NullableString** | Gets or sets the preferred metadata language. | [optional] 
**MetadataCountryCode** | Pointer to **NullableString** | Gets or sets the metadata country code. | [optional] 
**SeasonZeroDisplayName** | Pointer to **string** |  | [optional] 
**MetadataSavers** | Pointer to **[]string** |  | [optional] 
**DisabledLocalMetadataReaders** | Pointer to **[]string** |  | [optional] 
**LocalMetadataReaderOrder** | Pointer to **[]string** |  | [optional] 
**DisabledSubtitleFetchers** | Pointer to **[]string** |  | [optional] 
**SubtitleFetcherOrder** | Pointer to **[]string** |  | [optional] 
**SkipSubtitlesIfEmbeddedSubtitlesPresent** | Pointer to **bool** |  | [optional] 
**SkipSubtitlesIfAudioTrackMatches** | Pointer to **bool** |  | [optional] 
**SubtitleDownloadLanguages** | Pointer to **[]string** |  | [optional] 
**RequirePerfectSubtitleMatch** | Pointer to **bool** |  | [optional] 
**SaveSubtitlesWithMedia** | Pointer to **bool** |  | [optional] 
**AutomaticallyAddToCollection** | Pointer to **bool** |  | [optional] 
**AllowEmbeddedSubtitles** | Pointer to [**EmbeddedSubtitleOptions**](EmbeddedSubtitleOptions.md) | An enum representing the options to disable embedded subs. | [optional] 
**TypeOptions** | Pointer to [**[]TypeOptions**](TypeOptions.md) |  | [optional] 

## Methods

### NewAddVirtualFolderDtoLibraryOptions

`func NewAddVirtualFolderDtoLibraryOptions() *AddVirtualFolderDtoLibraryOptions`

NewAddVirtualFolderDtoLibraryOptions instantiates a new AddVirtualFolderDtoLibraryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVirtualFolderDtoLibraryOptionsWithDefaults

`func NewAddVirtualFolderDtoLibraryOptionsWithDefaults() *AddVirtualFolderDtoLibraryOptions`

NewAddVirtualFolderDtoLibraryOptionsWithDefaults instantiates a new AddVirtualFolderDtoLibraryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnablePhotos

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnablePhotos() bool`

GetEnablePhotos returns the EnablePhotos field if non-nil, zero value otherwise.

### GetEnablePhotosOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnablePhotosOk() (*bool, bool)`

GetEnablePhotosOk returns a tuple with the EnablePhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePhotos

`func (o *AddVirtualFolderDtoLibraryOptions) SetEnablePhotos(v bool)`

SetEnablePhotos sets EnablePhotos field to given value.

### HasEnablePhotos

`func (o *AddVirtualFolderDtoLibraryOptions) HasEnablePhotos() bool`

HasEnablePhotos returns a boolean if a field has been set.

### GetEnableRealtimeMonitor

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnableRealtimeMonitor() bool`

GetEnableRealtimeMonitor returns the EnableRealtimeMonitor field if non-nil, zero value otherwise.

### GetEnableRealtimeMonitorOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnableRealtimeMonitorOk() (*bool, bool)`

GetEnableRealtimeMonitorOk returns a tuple with the EnableRealtimeMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRealtimeMonitor

`func (o *AddVirtualFolderDtoLibraryOptions) SetEnableRealtimeMonitor(v bool)`

SetEnableRealtimeMonitor sets EnableRealtimeMonitor field to given value.

### HasEnableRealtimeMonitor

`func (o *AddVirtualFolderDtoLibraryOptions) HasEnableRealtimeMonitor() bool`

HasEnableRealtimeMonitor returns a boolean if a field has been set.

### GetEnableChapterImageExtraction

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnableChapterImageExtraction() bool`

GetEnableChapterImageExtraction returns the EnableChapterImageExtraction field if non-nil, zero value otherwise.

### GetEnableChapterImageExtractionOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnableChapterImageExtractionOk() (*bool, bool)`

GetEnableChapterImageExtractionOk returns a tuple with the EnableChapterImageExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChapterImageExtraction

`func (o *AddVirtualFolderDtoLibraryOptions) SetEnableChapterImageExtraction(v bool)`

SetEnableChapterImageExtraction sets EnableChapterImageExtraction field to given value.

### HasEnableChapterImageExtraction

`func (o *AddVirtualFolderDtoLibraryOptions) HasEnableChapterImageExtraction() bool`

HasEnableChapterImageExtraction returns a boolean if a field has been set.

### GetExtractChapterImagesDuringLibraryScan

`func (o *AddVirtualFolderDtoLibraryOptions) GetExtractChapterImagesDuringLibraryScan() bool`

GetExtractChapterImagesDuringLibraryScan returns the ExtractChapterImagesDuringLibraryScan field if non-nil, zero value otherwise.

### GetExtractChapterImagesDuringLibraryScanOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetExtractChapterImagesDuringLibraryScanOk() (*bool, bool)`

GetExtractChapterImagesDuringLibraryScanOk returns a tuple with the ExtractChapterImagesDuringLibraryScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractChapterImagesDuringLibraryScan

`func (o *AddVirtualFolderDtoLibraryOptions) SetExtractChapterImagesDuringLibraryScan(v bool)`

SetExtractChapterImagesDuringLibraryScan sets ExtractChapterImagesDuringLibraryScan field to given value.

### HasExtractChapterImagesDuringLibraryScan

`func (o *AddVirtualFolderDtoLibraryOptions) HasExtractChapterImagesDuringLibraryScan() bool`

HasExtractChapterImagesDuringLibraryScan returns a boolean if a field has been set.

### GetPathInfos

`func (o *AddVirtualFolderDtoLibraryOptions) GetPathInfos() []MediaPathInfo`

GetPathInfos returns the PathInfos field if non-nil, zero value otherwise.

### GetPathInfosOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetPathInfosOk() (*[]MediaPathInfo, bool)`

GetPathInfosOk returns a tuple with the PathInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathInfos

`func (o *AddVirtualFolderDtoLibraryOptions) SetPathInfos(v []MediaPathInfo)`

SetPathInfos sets PathInfos field to given value.

### HasPathInfos

`func (o *AddVirtualFolderDtoLibraryOptions) HasPathInfos() bool`

HasPathInfos returns a boolean if a field has been set.

### GetSaveLocalMetadata

`func (o *AddVirtualFolderDtoLibraryOptions) GetSaveLocalMetadata() bool`

GetSaveLocalMetadata returns the SaveLocalMetadata field if non-nil, zero value otherwise.

### GetSaveLocalMetadataOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetSaveLocalMetadataOk() (*bool, bool)`

GetSaveLocalMetadataOk returns a tuple with the SaveLocalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLocalMetadata

`func (o *AddVirtualFolderDtoLibraryOptions) SetSaveLocalMetadata(v bool)`

SetSaveLocalMetadata sets SaveLocalMetadata field to given value.

### HasSaveLocalMetadata

`func (o *AddVirtualFolderDtoLibraryOptions) HasSaveLocalMetadata() bool`

HasSaveLocalMetadata returns a boolean if a field has been set.

### GetEnableInternetProviders

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnableInternetProviders() bool`

GetEnableInternetProviders returns the EnableInternetProviders field if non-nil, zero value otherwise.

### GetEnableInternetProvidersOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnableInternetProvidersOk() (*bool, bool)`

GetEnableInternetProvidersOk returns a tuple with the EnableInternetProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternetProviders

`func (o *AddVirtualFolderDtoLibraryOptions) SetEnableInternetProviders(v bool)`

SetEnableInternetProviders sets EnableInternetProviders field to given value.

### HasEnableInternetProviders

`func (o *AddVirtualFolderDtoLibraryOptions) HasEnableInternetProviders() bool`

HasEnableInternetProviders returns a boolean if a field has been set.

### GetEnableAutomaticSeriesGrouping

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnableAutomaticSeriesGrouping() bool`

GetEnableAutomaticSeriesGrouping returns the EnableAutomaticSeriesGrouping field if non-nil, zero value otherwise.

### GetEnableAutomaticSeriesGroupingOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnableAutomaticSeriesGroupingOk() (*bool, bool)`

GetEnableAutomaticSeriesGroupingOk returns a tuple with the EnableAutomaticSeriesGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticSeriesGrouping

`func (o *AddVirtualFolderDtoLibraryOptions) SetEnableAutomaticSeriesGrouping(v bool)`

SetEnableAutomaticSeriesGrouping sets EnableAutomaticSeriesGrouping field to given value.

### HasEnableAutomaticSeriesGrouping

`func (o *AddVirtualFolderDtoLibraryOptions) HasEnableAutomaticSeriesGrouping() bool`

HasEnableAutomaticSeriesGrouping returns a boolean if a field has been set.

### GetEnableEmbeddedTitles

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnableEmbeddedTitles() bool`

GetEnableEmbeddedTitles returns the EnableEmbeddedTitles field if non-nil, zero value otherwise.

### GetEnableEmbeddedTitlesOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnableEmbeddedTitlesOk() (*bool, bool)`

GetEnableEmbeddedTitlesOk returns a tuple with the EnableEmbeddedTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedTitles

`func (o *AddVirtualFolderDtoLibraryOptions) SetEnableEmbeddedTitles(v bool)`

SetEnableEmbeddedTitles sets EnableEmbeddedTitles field to given value.

### HasEnableEmbeddedTitles

`func (o *AddVirtualFolderDtoLibraryOptions) HasEnableEmbeddedTitles() bool`

HasEnableEmbeddedTitles returns a boolean if a field has been set.

### GetEnableEmbeddedEpisodeInfos

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnableEmbeddedEpisodeInfos() bool`

GetEnableEmbeddedEpisodeInfos returns the EnableEmbeddedEpisodeInfos field if non-nil, zero value otherwise.

### GetEnableEmbeddedEpisodeInfosOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetEnableEmbeddedEpisodeInfosOk() (*bool, bool)`

GetEnableEmbeddedEpisodeInfosOk returns a tuple with the EnableEmbeddedEpisodeInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedEpisodeInfos

`func (o *AddVirtualFolderDtoLibraryOptions) SetEnableEmbeddedEpisodeInfos(v bool)`

SetEnableEmbeddedEpisodeInfos sets EnableEmbeddedEpisodeInfos field to given value.

### HasEnableEmbeddedEpisodeInfos

`func (o *AddVirtualFolderDtoLibraryOptions) HasEnableEmbeddedEpisodeInfos() bool`

HasEnableEmbeddedEpisodeInfos returns a boolean if a field has been set.

### GetAutomaticRefreshIntervalDays

`func (o *AddVirtualFolderDtoLibraryOptions) GetAutomaticRefreshIntervalDays() int32`

GetAutomaticRefreshIntervalDays returns the AutomaticRefreshIntervalDays field if non-nil, zero value otherwise.

### GetAutomaticRefreshIntervalDaysOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetAutomaticRefreshIntervalDaysOk() (*int32, bool)`

GetAutomaticRefreshIntervalDaysOk returns a tuple with the AutomaticRefreshIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticRefreshIntervalDays

`func (o *AddVirtualFolderDtoLibraryOptions) SetAutomaticRefreshIntervalDays(v int32)`

SetAutomaticRefreshIntervalDays sets AutomaticRefreshIntervalDays field to given value.

### HasAutomaticRefreshIntervalDays

`func (o *AddVirtualFolderDtoLibraryOptions) HasAutomaticRefreshIntervalDays() bool`

HasAutomaticRefreshIntervalDays returns a boolean if a field has been set.

### GetPreferredMetadataLanguage

`func (o *AddVirtualFolderDtoLibraryOptions) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *AddVirtualFolderDtoLibraryOptions) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *AddVirtualFolderDtoLibraryOptions) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### SetPreferredMetadataLanguageNil

`func (o *AddVirtualFolderDtoLibraryOptions) SetPreferredMetadataLanguageNil(b bool)`

 SetPreferredMetadataLanguageNil sets the value for PreferredMetadataLanguage to be an explicit nil

### UnsetPreferredMetadataLanguage
`func (o *AddVirtualFolderDtoLibraryOptions) UnsetPreferredMetadataLanguage()`

UnsetPreferredMetadataLanguage ensures that no value is present for PreferredMetadataLanguage, not even an explicit nil
### GetMetadataCountryCode

`func (o *AddVirtualFolderDtoLibraryOptions) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *AddVirtualFolderDtoLibraryOptions) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *AddVirtualFolderDtoLibraryOptions) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### SetMetadataCountryCodeNil

`func (o *AddVirtualFolderDtoLibraryOptions) SetMetadataCountryCodeNil(b bool)`

 SetMetadataCountryCodeNil sets the value for MetadataCountryCode to be an explicit nil

### UnsetMetadataCountryCode
`func (o *AddVirtualFolderDtoLibraryOptions) UnsetMetadataCountryCode()`

UnsetMetadataCountryCode ensures that no value is present for MetadataCountryCode, not even an explicit nil
### GetSeasonZeroDisplayName

`func (o *AddVirtualFolderDtoLibraryOptions) GetSeasonZeroDisplayName() string`

GetSeasonZeroDisplayName returns the SeasonZeroDisplayName field if non-nil, zero value otherwise.

### GetSeasonZeroDisplayNameOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetSeasonZeroDisplayNameOk() (*string, bool)`

GetSeasonZeroDisplayNameOk returns a tuple with the SeasonZeroDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonZeroDisplayName

`func (o *AddVirtualFolderDtoLibraryOptions) SetSeasonZeroDisplayName(v string)`

SetSeasonZeroDisplayName sets SeasonZeroDisplayName field to given value.

### HasSeasonZeroDisplayName

`func (o *AddVirtualFolderDtoLibraryOptions) HasSeasonZeroDisplayName() bool`

HasSeasonZeroDisplayName returns a boolean if a field has been set.

### GetMetadataSavers

`func (o *AddVirtualFolderDtoLibraryOptions) GetMetadataSavers() []string`

GetMetadataSavers returns the MetadataSavers field if non-nil, zero value otherwise.

### GetMetadataSaversOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetMetadataSaversOk() (*[]string, bool)`

GetMetadataSaversOk returns a tuple with the MetadataSavers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSavers

`func (o *AddVirtualFolderDtoLibraryOptions) SetMetadataSavers(v []string)`

SetMetadataSavers sets MetadataSavers field to given value.

### HasMetadataSavers

`func (o *AddVirtualFolderDtoLibraryOptions) HasMetadataSavers() bool`

HasMetadataSavers returns a boolean if a field has been set.

### SetMetadataSaversNil

`func (o *AddVirtualFolderDtoLibraryOptions) SetMetadataSaversNil(b bool)`

 SetMetadataSaversNil sets the value for MetadataSavers to be an explicit nil

### UnsetMetadataSavers
`func (o *AddVirtualFolderDtoLibraryOptions) UnsetMetadataSavers()`

UnsetMetadataSavers ensures that no value is present for MetadataSavers, not even an explicit nil
### GetDisabledLocalMetadataReaders

`func (o *AddVirtualFolderDtoLibraryOptions) GetDisabledLocalMetadataReaders() []string`

GetDisabledLocalMetadataReaders returns the DisabledLocalMetadataReaders field if non-nil, zero value otherwise.

### GetDisabledLocalMetadataReadersOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetDisabledLocalMetadataReadersOk() (*[]string, bool)`

GetDisabledLocalMetadataReadersOk returns a tuple with the DisabledLocalMetadataReaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledLocalMetadataReaders

`func (o *AddVirtualFolderDtoLibraryOptions) SetDisabledLocalMetadataReaders(v []string)`

SetDisabledLocalMetadataReaders sets DisabledLocalMetadataReaders field to given value.

### HasDisabledLocalMetadataReaders

`func (o *AddVirtualFolderDtoLibraryOptions) HasDisabledLocalMetadataReaders() bool`

HasDisabledLocalMetadataReaders returns a boolean if a field has been set.

### GetLocalMetadataReaderOrder

`func (o *AddVirtualFolderDtoLibraryOptions) GetLocalMetadataReaderOrder() []string`

GetLocalMetadataReaderOrder returns the LocalMetadataReaderOrder field if non-nil, zero value otherwise.

### GetLocalMetadataReaderOrderOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetLocalMetadataReaderOrderOk() (*[]string, bool)`

GetLocalMetadataReaderOrderOk returns a tuple with the LocalMetadataReaderOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMetadataReaderOrder

`func (o *AddVirtualFolderDtoLibraryOptions) SetLocalMetadataReaderOrder(v []string)`

SetLocalMetadataReaderOrder sets LocalMetadataReaderOrder field to given value.

### HasLocalMetadataReaderOrder

`func (o *AddVirtualFolderDtoLibraryOptions) HasLocalMetadataReaderOrder() bool`

HasLocalMetadataReaderOrder returns a boolean if a field has been set.

### SetLocalMetadataReaderOrderNil

`func (o *AddVirtualFolderDtoLibraryOptions) SetLocalMetadataReaderOrderNil(b bool)`

 SetLocalMetadataReaderOrderNil sets the value for LocalMetadataReaderOrder to be an explicit nil

### UnsetLocalMetadataReaderOrder
`func (o *AddVirtualFolderDtoLibraryOptions) UnsetLocalMetadataReaderOrder()`

UnsetLocalMetadataReaderOrder ensures that no value is present for LocalMetadataReaderOrder, not even an explicit nil
### GetDisabledSubtitleFetchers

`func (o *AddVirtualFolderDtoLibraryOptions) GetDisabledSubtitleFetchers() []string`

GetDisabledSubtitleFetchers returns the DisabledSubtitleFetchers field if non-nil, zero value otherwise.

### GetDisabledSubtitleFetchersOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetDisabledSubtitleFetchersOk() (*[]string, bool)`

GetDisabledSubtitleFetchersOk returns a tuple with the DisabledSubtitleFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledSubtitleFetchers

`func (o *AddVirtualFolderDtoLibraryOptions) SetDisabledSubtitleFetchers(v []string)`

SetDisabledSubtitleFetchers sets DisabledSubtitleFetchers field to given value.

### HasDisabledSubtitleFetchers

`func (o *AddVirtualFolderDtoLibraryOptions) HasDisabledSubtitleFetchers() bool`

HasDisabledSubtitleFetchers returns a boolean if a field has been set.

### GetSubtitleFetcherOrder

`func (o *AddVirtualFolderDtoLibraryOptions) GetSubtitleFetcherOrder() []string`

GetSubtitleFetcherOrder returns the SubtitleFetcherOrder field if non-nil, zero value otherwise.

### GetSubtitleFetcherOrderOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetSubtitleFetcherOrderOk() (*[]string, bool)`

GetSubtitleFetcherOrderOk returns a tuple with the SubtitleFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleFetcherOrder

`func (o *AddVirtualFolderDtoLibraryOptions) SetSubtitleFetcherOrder(v []string)`

SetSubtitleFetcherOrder sets SubtitleFetcherOrder field to given value.

### HasSubtitleFetcherOrder

`func (o *AddVirtualFolderDtoLibraryOptions) HasSubtitleFetcherOrder() bool`

HasSubtitleFetcherOrder returns a boolean if a field has been set.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *AddVirtualFolderDtoLibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

GetSkipSubtitlesIfEmbeddedSubtitlesPresent returns the SkipSubtitlesIfEmbeddedSubtitlesPresent field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk() (*bool, bool)`

GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk returns a tuple with the SkipSubtitlesIfEmbeddedSubtitlesPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *AddVirtualFolderDtoLibraryOptions) SetSkipSubtitlesIfEmbeddedSubtitlesPresent(v bool)`

SetSkipSubtitlesIfEmbeddedSubtitlesPresent sets SkipSubtitlesIfEmbeddedSubtitlesPresent field to given value.

### HasSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *AddVirtualFolderDtoLibraryOptions) HasSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

HasSkipSubtitlesIfEmbeddedSubtitlesPresent returns a boolean if a field has been set.

### GetSkipSubtitlesIfAudioTrackMatches

`func (o *AddVirtualFolderDtoLibraryOptions) GetSkipSubtitlesIfAudioTrackMatches() bool`

GetSkipSubtitlesIfAudioTrackMatches returns the SkipSubtitlesIfAudioTrackMatches field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfAudioTrackMatchesOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetSkipSubtitlesIfAudioTrackMatchesOk() (*bool, bool)`

GetSkipSubtitlesIfAudioTrackMatchesOk returns a tuple with the SkipSubtitlesIfAudioTrackMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfAudioTrackMatches

`func (o *AddVirtualFolderDtoLibraryOptions) SetSkipSubtitlesIfAudioTrackMatches(v bool)`

SetSkipSubtitlesIfAudioTrackMatches sets SkipSubtitlesIfAudioTrackMatches field to given value.

### HasSkipSubtitlesIfAudioTrackMatches

`func (o *AddVirtualFolderDtoLibraryOptions) HasSkipSubtitlesIfAudioTrackMatches() bool`

HasSkipSubtitlesIfAudioTrackMatches returns a boolean if a field has been set.

### GetSubtitleDownloadLanguages

`func (o *AddVirtualFolderDtoLibraryOptions) GetSubtitleDownloadLanguages() []string`

GetSubtitleDownloadLanguages returns the SubtitleDownloadLanguages field if non-nil, zero value otherwise.

### GetSubtitleDownloadLanguagesOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetSubtitleDownloadLanguagesOk() (*[]string, bool)`

GetSubtitleDownloadLanguagesOk returns a tuple with the SubtitleDownloadLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleDownloadLanguages

`func (o *AddVirtualFolderDtoLibraryOptions) SetSubtitleDownloadLanguages(v []string)`

SetSubtitleDownloadLanguages sets SubtitleDownloadLanguages field to given value.

### HasSubtitleDownloadLanguages

`func (o *AddVirtualFolderDtoLibraryOptions) HasSubtitleDownloadLanguages() bool`

HasSubtitleDownloadLanguages returns a boolean if a field has been set.

### SetSubtitleDownloadLanguagesNil

`func (o *AddVirtualFolderDtoLibraryOptions) SetSubtitleDownloadLanguagesNil(b bool)`

 SetSubtitleDownloadLanguagesNil sets the value for SubtitleDownloadLanguages to be an explicit nil

### UnsetSubtitleDownloadLanguages
`func (o *AddVirtualFolderDtoLibraryOptions) UnsetSubtitleDownloadLanguages()`

UnsetSubtitleDownloadLanguages ensures that no value is present for SubtitleDownloadLanguages, not even an explicit nil
### GetRequirePerfectSubtitleMatch

`func (o *AddVirtualFolderDtoLibraryOptions) GetRequirePerfectSubtitleMatch() bool`

GetRequirePerfectSubtitleMatch returns the RequirePerfectSubtitleMatch field if non-nil, zero value otherwise.

### GetRequirePerfectSubtitleMatchOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetRequirePerfectSubtitleMatchOk() (*bool, bool)`

GetRequirePerfectSubtitleMatchOk returns a tuple with the RequirePerfectSubtitleMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePerfectSubtitleMatch

`func (o *AddVirtualFolderDtoLibraryOptions) SetRequirePerfectSubtitleMatch(v bool)`

SetRequirePerfectSubtitleMatch sets RequirePerfectSubtitleMatch field to given value.

### HasRequirePerfectSubtitleMatch

`func (o *AddVirtualFolderDtoLibraryOptions) HasRequirePerfectSubtitleMatch() bool`

HasRequirePerfectSubtitleMatch returns a boolean if a field has been set.

### GetSaveSubtitlesWithMedia

`func (o *AddVirtualFolderDtoLibraryOptions) GetSaveSubtitlesWithMedia() bool`

GetSaveSubtitlesWithMedia returns the SaveSubtitlesWithMedia field if non-nil, zero value otherwise.

### GetSaveSubtitlesWithMediaOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetSaveSubtitlesWithMediaOk() (*bool, bool)`

GetSaveSubtitlesWithMediaOk returns a tuple with the SaveSubtitlesWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveSubtitlesWithMedia

`func (o *AddVirtualFolderDtoLibraryOptions) SetSaveSubtitlesWithMedia(v bool)`

SetSaveSubtitlesWithMedia sets SaveSubtitlesWithMedia field to given value.

### HasSaveSubtitlesWithMedia

`func (o *AddVirtualFolderDtoLibraryOptions) HasSaveSubtitlesWithMedia() bool`

HasSaveSubtitlesWithMedia returns a boolean if a field has been set.

### GetAutomaticallyAddToCollection

`func (o *AddVirtualFolderDtoLibraryOptions) GetAutomaticallyAddToCollection() bool`

GetAutomaticallyAddToCollection returns the AutomaticallyAddToCollection field if non-nil, zero value otherwise.

### GetAutomaticallyAddToCollectionOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetAutomaticallyAddToCollectionOk() (*bool, bool)`

GetAutomaticallyAddToCollectionOk returns a tuple with the AutomaticallyAddToCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyAddToCollection

`func (o *AddVirtualFolderDtoLibraryOptions) SetAutomaticallyAddToCollection(v bool)`

SetAutomaticallyAddToCollection sets AutomaticallyAddToCollection field to given value.

### HasAutomaticallyAddToCollection

`func (o *AddVirtualFolderDtoLibraryOptions) HasAutomaticallyAddToCollection() bool`

HasAutomaticallyAddToCollection returns a boolean if a field has been set.

### GetAllowEmbeddedSubtitles

`func (o *AddVirtualFolderDtoLibraryOptions) GetAllowEmbeddedSubtitles() EmbeddedSubtitleOptions`

GetAllowEmbeddedSubtitles returns the AllowEmbeddedSubtitles field if non-nil, zero value otherwise.

### GetAllowEmbeddedSubtitlesOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetAllowEmbeddedSubtitlesOk() (*EmbeddedSubtitleOptions, bool)`

GetAllowEmbeddedSubtitlesOk returns a tuple with the AllowEmbeddedSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmbeddedSubtitles

`func (o *AddVirtualFolderDtoLibraryOptions) SetAllowEmbeddedSubtitles(v EmbeddedSubtitleOptions)`

SetAllowEmbeddedSubtitles sets AllowEmbeddedSubtitles field to given value.

### HasAllowEmbeddedSubtitles

`func (o *AddVirtualFolderDtoLibraryOptions) HasAllowEmbeddedSubtitles() bool`

HasAllowEmbeddedSubtitles returns a boolean if a field has been set.

### GetTypeOptions

`func (o *AddVirtualFolderDtoLibraryOptions) GetTypeOptions() []TypeOptions`

GetTypeOptions returns the TypeOptions field if non-nil, zero value otherwise.

### GetTypeOptionsOk

`func (o *AddVirtualFolderDtoLibraryOptions) GetTypeOptionsOk() (*[]TypeOptions, bool)`

GetTypeOptionsOk returns a tuple with the TypeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOptions

`func (o *AddVirtualFolderDtoLibraryOptions) SetTypeOptions(v []TypeOptions)`

SetTypeOptions sets TypeOptions field to given value.

### HasTypeOptions

`func (o *AddVirtualFolderDtoLibraryOptions) HasTypeOptions() bool`

HasTypeOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


