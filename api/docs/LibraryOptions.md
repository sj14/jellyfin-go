# LibraryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**EnablePhotos** | Pointer to **bool** |  | [optional] 
**EnableRealtimeMonitor** | Pointer to **bool** |  | [optional] 
**EnableLUFSScan** | Pointer to **bool** |  | [optional] 
**EnableChapterImageExtraction** | Pointer to **bool** |  | [optional] 
**ExtractChapterImagesDuringLibraryScan** | Pointer to **bool** |  | [optional] 
**EnableTrickplayImageExtraction** | Pointer to **bool** |  | [optional] 
**ExtractTrickplayImagesDuringLibraryScan** | Pointer to **bool** |  | [optional] 
**PathInfos** | Pointer to [**[]MediaPathInfo**](MediaPathInfo.md) |  | [optional] 
**SaveLocalMetadata** | Pointer to **bool** |  | [optional] 
**EnableInternetProviders** | Pointer to **bool** |  | [optional] 
**EnableAutomaticSeriesGrouping** | Pointer to **bool** |  | [optional] 
**EnableEmbeddedTitles** | Pointer to **bool** |  | [optional] 
**EnableEmbeddedExtrasTitles** | Pointer to **bool** |  | [optional] 
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
**DisabledMediaSegmentProviders** | Pointer to **[]string** |  | [optional] 
**MediaSegmentProvideOrder** | Pointer to **[]string** |  | [optional] 
**SkipSubtitlesIfEmbeddedSubtitlesPresent** | Pointer to **bool** |  | [optional] 
**SkipSubtitlesIfAudioTrackMatches** | Pointer to **bool** |  | [optional] 
**SubtitleDownloadLanguages** | Pointer to **[]string** |  | [optional] 
**RequirePerfectSubtitleMatch** | Pointer to **bool** |  | [optional] 
**SaveSubtitlesWithMedia** | Pointer to **bool** |  | [optional] 
**SaveLyricsWithMedia** | Pointer to **bool** |  | [optional] [default to false]
**SaveTrickplayWithMedia** | Pointer to **bool** |  | [optional] [default to false]
**DisabledLyricFetchers** | Pointer to **[]string** |  | [optional] 
**LyricFetcherOrder** | Pointer to **[]string** |  | [optional] 
**PreferNonstandardArtistsTag** | Pointer to **bool** |  | [optional] [default to false]
**UseCustomTagDelimiters** | Pointer to **bool** |  | [optional] [default to false]
**CustomTagDelimiters** | Pointer to **[]string** |  | [optional] 
**DelimiterWhitelist** | Pointer to **[]string** |  | [optional] 
**AutomaticallyAddToCollection** | Pointer to **bool** |  | [optional] 
**AllowEmbeddedSubtitles** | Pointer to [**EmbeddedSubtitleOptions**](EmbeddedSubtitleOptions.md) | An enum representing the options to disable embedded subs. | [optional] 
**TypeOptions** | Pointer to [**[]TypeOptions**](TypeOptions.md) |  | [optional] 

## Methods

### NewLibraryOptions

`func NewLibraryOptions() *LibraryOptions`

NewLibraryOptions instantiates a new LibraryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryOptionsWithDefaults

`func NewLibraryOptionsWithDefaults() *LibraryOptions`

NewLibraryOptionsWithDefaults instantiates a new LibraryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *LibraryOptions) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LibraryOptions) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LibraryOptions) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LibraryOptions) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnablePhotos

`func (o *LibraryOptions) GetEnablePhotos() bool`

GetEnablePhotos returns the EnablePhotos field if non-nil, zero value otherwise.

### GetEnablePhotosOk

`func (o *LibraryOptions) GetEnablePhotosOk() (*bool, bool)`

GetEnablePhotosOk returns a tuple with the EnablePhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePhotos

`func (o *LibraryOptions) SetEnablePhotos(v bool)`

SetEnablePhotos sets EnablePhotos field to given value.

### HasEnablePhotos

`func (o *LibraryOptions) HasEnablePhotos() bool`

HasEnablePhotos returns a boolean if a field has been set.

### GetEnableRealtimeMonitor

`func (o *LibraryOptions) GetEnableRealtimeMonitor() bool`

GetEnableRealtimeMonitor returns the EnableRealtimeMonitor field if non-nil, zero value otherwise.

### GetEnableRealtimeMonitorOk

`func (o *LibraryOptions) GetEnableRealtimeMonitorOk() (*bool, bool)`

GetEnableRealtimeMonitorOk returns a tuple with the EnableRealtimeMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRealtimeMonitor

`func (o *LibraryOptions) SetEnableRealtimeMonitor(v bool)`

SetEnableRealtimeMonitor sets EnableRealtimeMonitor field to given value.

### HasEnableRealtimeMonitor

`func (o *LibraryOptions) HasEnableRealtimeMonitor() bool`

HasEnableRealtimeMonitor returns a boolean if a field has been set.

### GetEnableLUFSScan

`func (o *LibraryOptions) GetEnableLUFSScan() bool`

GetEnableLUFSScan returns the EnableLUFSScan field if non-nil, zero value otherwise.

### GetEnableLUFSScanOk

`func (o *LibraryOptions) GetEnableLUFSScanOk() (*bool, bool)`

GetEnableLUFSScanOk returns a tuple with the EnableLUFSScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLUFSScan

`func (o *LibraryOptions) SetEnableLUFSScan(v bool)`

SetEnableLUFSScan sets EnableLUFSScan field to given value.

### HasEnableLUFSScan

`func (o *LibraryOptions) HasEnableLUFSScan() bool`

HasEnableLUFSScan returns a boolean if a field has been set.

### GetEnableChapterImageExtraction

`func (o *LibraryOptions) GetEnableChapterImageExtraction() bool`

GetEnableChapterImageExtraction returns the EnableChapterImageExtraction field if non-nil, zero value otherwise.

### GetEnableChapterImageExtractionOk

`func (o *LibraryOptions) GetEnableChapterImageExtractionOk() (*bool, bool)`

GetEnableChapterImageExtractionOk returns a tuple with the EnableChapterImageExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChapterImageExtraction

`func (o *LibraryOptions) SetEnableChapterImageExtraction(v bool)`

SetEnableChapterImageExtraction sets EnableChapterImageExtraction field to given value.

### HasEnableChapterImageExtraction

`func (o *LibraryOptions) HasEnableChapterImageExtraction() bool`

HasEnableChapterImageExtraction returns a boolean if a field has been set.

### GetExtractChapterImagesDuringLibraryScan

`func (o *LibraryOptions) GetExtractChapterImagesDuringLibraryScan() bool`

GetExtractChapterImagesDuringLibraryScan returns the ExtractChapterImagesDuringLibraryScan field if non-nil, zero value otherwise.

### GetExtractChapterImagesDuringLibraryScanOk

`func (o *LibraryOptions) GetExtractChapterImagesDuringLibraryScanOk() (*bool, bool)`

GetExtractChapterImagesDuringLibraryScanOk returns a tuple with the ExtractChapterImagesDuringLibraryScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractChapterImagesDuringLibraryScan

`func (o *LibraryOptions) SetExtractChapterImagesDuringLibraryScan(v bool)`

SetExtractChapterImagesDuringLibraryScan sets ExtractChapterImagesDuringLibraryScan field to given value.

### HasExtractChapterImagesDuringLibraryScan

`func (o *LibraryOptions) HasExtractChapterImagesDuringLibraryScan() bool`

HasExtractChapterImagesDuringLibraryScan returns a boolean if a field has been set.

### GetEnableTrickplayImageExtraction

`func (o *LibraryOptions) GetEnableTrickplayImageExtraction() bool`

GetEnableTrickplayImageExtraction returns the EnableTrickplayImageExtraction field if non-nil, zero value otherwise.

### GetEnableTrickplayImageExtractionOk

`func (o *LibraryOptions) GetEnableTrickplayImageExtractionOk() (*bool, bool)`

GetEnableTrickplayImageExtractionOk returns a tuple with the EnableTrickplayImageExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTrickplayImageExtraction

`func (o *LibraryOptions) SetEnableTrickplayImageExtraction(v bool)`

SetEnableTrickplayImageExtraction sets EnableTrickplayImageExtraction field to given value.

### HasEnableTrickplayImageExtraction

`func (o *LibraryOptions) HasEnableTrickplayImageExtraction() bool`

HasEnableTrickplayImageExtraction returns a boolean if a field has been set.

### GetExtractTrickplayImagesDuringLibraryScan

`func (o *LibraryOptions) GetExtractTrickplayImagesDuringLibraryScan() bool`

GetExtractTrickplayImagesDuringLibraryScan returns the ExtractTrickplayImagesDuringLibraryScan field if non-nil, zero value otherwise.

### GetExtractTrickplayImagesDuringLibraryScanOk

`func (o *LibraryOptions) GetExtractTrickplayImagesDuringLibraryScanOk() (*bool, bool)`

GetExtractTrickplayImagesDuringLibraryScanOk returns a tuple with the ExtractTrickplayImagesDuringLibraryScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractTrickplayImagesDuringLibraryScan

`func (o *LibraryOptions) SetExtractTrickplayImagesDuringLibraryScan(v bool)`

SetExtractTrickplayImagesDuringLibraryScan sets ExtractTrickplayImagesDuringLibraryScan field to given value.

### HasExtractTrickplayImagesDuringLibraryScan

`func (o *LibraryOptions) HasExtractTrickplayImagesDuringLibraryScan() bool`

HasExtractTrickplayImagesDuringLibraryScan returns a boolean if a field has been set.

### GetPathInfos

`func (o *LibraryOptions) GetPathInfos() []MediaPathInfo`

GetPathInfos returns the PathInfos field if non-nil, zero value otherwise.

### GetPathInfosOk

`func (o *LibraryOptions) GetPathInfosOk() (*[]MediaPathInfo, bool)`

GetPathInfosOk returns a tuple with the PathInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathInfos

`func (o *LibraryOptions) SetPathInfos(v []MediaPathInfo)`

SetPathInfos sets PathInfos field to given value.

### HasPathInfos

`func (o *LibraryOptions) HasPathInfos() bool`

HasPathInfos returns a boolean if a field has been set.

### GetSaveLocalMetadata

`func (o *LibraryOptions) GetSaveLocalMetadata() bool`

GetSaveLocalMetadata returns the SaveLocalMetadata field if non-nil, zero value otherwise.

### GetSaveLocalMetadataOk

`func (o *LibraryOptions) GetSaveLocalMetadataOk() (*bool, bool)`

GetSaveLocalMetadataOk returns a tuple with the SaveLocalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLocalMetadata

`func (o *LibraryOptions) SetSaveLocalMetadata(v bool)`

SetSaveLocalMetadata sets SaveLocalMetadata field to given value.

### HasSaveLocalMetadata

`func (o *LibraryOptions) HasSaveLocalMetadata() bool`

HasSaveLocalMetadata returns a boolean if a field has been set.

### GetEnableInternetProviders

`func (o *LibraryOptions) GetEnableInternetProviders() bool`

GetEnableInternetProviders returns the EnableInternetProviders field if non-nil, zero value otherwise.

### GetEnableInternetProvidersOk

`func (o *LibraryOptions) GetEnableInternetProvidersOk() (*bool, bool)`

GetEnableInternetProvidersOk returns a tuple with the EnableInternetProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInternetProviders

`func (o *LibraryOptions) SetEnableInternetProviders(v bool)`

SetEnableInternetProviders sets EnableInternetProviders field to given value.

### HasEnableInternetProviders

`func (o *LibraryOptions) HasEnableInternetProviders() bool`

HasEnableInternetProviders returns a boolean if a field has been set.

### GetEnableAutomaticSeriesGrouping

`func (o *LibraryOptions) GetEnableAutomaticSeriesGrouping() bool`

GetEnableAutomaticSeriesGrouping returns the EnableAutomaticSeriesGrouping field if non-nil, zero value otherwise.

### GetEnableAutomaticSeriesGroupingOk

`func (o *LibraryOptions) GetEnableAutomaticSeriesGroupingOk() (*bool, bool)`

GetEnableAutomaticSeriesGroupingOk returns a tuple with the EnableAutomaticSeriesGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticSeriesGrouping

`func (o *LibraryOptions) SetEnableAutomaticSeriesGrouping(v bool)`

SetEnableAutomaticSeriesGrouping sets EnableAutomaticSeriesGrouping field to given value.

### HasEnableAutomaticSeriesGrouping

`func (o *LibraryOptions) HasEnableAutomaticSeriesGrouping() bool`

HasEnableAutomaticSeriesGrouping returns a boolean if a field has been set.

### GetEnableEmbeddedTitles

`func (o *LibraryOptions) GetEnableEmbeddedTitles() bool`

GetEnableEmbeddedTitles returns the EnableEmbeddedTitles field if non-nil, zero value otherwise.

### GetEnableEmbeddedTitlesOk

`func (o *LibraryOptions) GetEnableEmbeddedTitlesOk() (*bool, bool)`

GetEnableEmbeddedTitlesOk returns a tuple with the EnableEmbeddedTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedTitles

`func (o *LibraryOptions) SetEnableEmbeddedTitles(v bool)`

SetEnableEmbeddedTitles sets EnableEmbeddedTitles field to given value.

### HasEnableEmbeddedTitles

`func (o *LibraryOptions) HasEnableEmbeddedTitles() bool`

HasEnableEmbeddedTitles returns a boolean if a field has been set.

### GetEnableEmbeddedExtrasTitles

`func (o *LibraryOptions) GetEnableEmbeddedExtrasTitles() bool`

GetEnableEmbeddedExtrasTitles returns the EnableEmbeddedExtrasTitles field if non-nil, zero value otherwise.

### GetEnableEmbeddedExtrasTitlesOk

`func (o *LibraryOptions) GetEnableEmbeddedExtrasTitlesOk() (*bool, bool)`

GetEnableEmbeddedExtrasTitlesOk returns a tuple with the EnableEmbeddedExtrasTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedExtrasTitles

`func (o *LibraryOptions) SetEnableEmbeddedExtrasTitles(v bool)`

SetEnableEmbeddedExtrasTitles sets EnableEmbeddedExtrasTitles field to given value.

### HasEnableEmbeddedExtrasTitles

`func (o *LibraryOptions) HasEnableEmbeddedExtrasTitles() bool`

HasEnableEmbeddedExtrasTitles returns a boolean if a field has been set.

### GetEnableEmbeddedEpisodeInfos

`func (o *LibraryOptions) GetEnableEmbeddedEpisodeInfos() bool`

GetEnableEmbeddedEpisodeInfos returns the EnableEmbeddedEpisodeInfos field if non-nil, zero value otherwise.

### GetEnableEmbeddedEpisodeInfosOk

`func (o *LibraryOptions) GetEnableEmbeddedEpisodeInfosOk() (*bool, bool)`

GetEnableEmbeddedEpisodeInfosOk returns a tuple with the EnableEmbeddedEpisodeInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmbeddedEpisodeInfos

`func (o *LibraryOptions) SetEnableEmbeddedEpisodeInfos(v bool)`

SetEnableEmbeddedEpisodeInfos sets EnableEmbeddedEpisodeInfos field to given value.

### HasEnableEmbeddedEpisodeInfos

`func (o *LibraryOptions) HasEnableEmbeddedEpisodeInfos() bool`

HasEnableEmbeddedEpisodeInfos returns a boolean if a field has been set.

### GetAutomaticRefreshIntervalDays

`func (o *LibraryOptions) GetAutomaticRefreshIntervalDays() int32`

GetAutomaticRefreshIntervalDays returns the AutomaticRefreshIntervalDays field if non-nil, zero value otherwise.

### GetAutomaticRefreshIntervalDaysOk

`func (o *LibraryOptions) GetAutomaticRefreshIntervalDaysOk() (*int32, bool)`

GetAutomaticRefreshIntervalDaysOk returns a tuple with the AutomaticRefreshIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticRefreshIntervalDays

`func (o *LibraryOptions) SetAutomaticRefreshIntervalDays(v int32)`

SetAutomaticRefreshIntervalDays sets AutomaticRefreshIntervalDays field to given value.

### HasAutomaticRefreshIntervalDays

`func (o *LibraryOptions) HasAutomaticRefreshIntervalDays() bool`

HasAutomaticRefreshIntervalDays returns a boolean if a field has been set.

### GetPreferredMetadataLanguage

`func (o *LibraryOptions) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *LibraryOptions) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *LibraryOptions) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *LibraryOptions) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### SetPreferredMetadataLanguageNil

`func (o *LibraryOptions) SetPreferredMetadataLanguageNil(b bool)`

 SetPreferredMetadataLanguageNil sets the value for PreferredMetadataLanguage to be an explicit nil

### UnsetPreferredMetadataLanguage
`func (o *LibraryOptions) UnsetPreferredMetadataLanguage()`

UnsetPreferredMetadataLanguage ensures that no value is present for PreferredMetadataLanguage, not even an explicit nil
### GetMetadataCountryCode

`func (o *LibraryOptions) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *LibraryOptions) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *LibraryOptions) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *LibraryOptions) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### SetMetadataCountryCodeNil

`func (o *LibraryOptions) SetMetadataCountryCodeNil(b bool)`

 SetMetadataCountryCodeNil sets the value for MetadataCountryCode to be an explicit nil

### UnsetMetadataCountryCode
`func (o *LibraryOptions) UnsetMetadataCountryCode()`

UnsetMetadataCountryCode ensures that no value is present for MetadataCountryCode, not even an explicit nil
### GetSeasonZeroDisplayName

`func (o *LibraryOptions) GetSeasonZeroDisplayName() string`

GetSeasonZeroDisplayName returns the SeasonZeroDisplayName field if non-nil, zero value otherwise.

### GetSeasonZeroDisplayNameOk

`func (o *LibraryOptions) GetSeasonZeroDisplayNameOk() (*string, bool)`

GetSeasonZeroDisplayNameOk returns a tuple with the SeasonZeroDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonZeroDisplayName

`func (o *LibraryOptions) SetSeasonZeroDisplayName(v string)`

SetSeasonZeroDisplayName sets SeasonZeroDisplayName field to given value.

### HasSeasonZeroDisplayName

`func (o *LibraryOptions) HasSeasonZeroDisplayName() bool`

HasSeasonZeroDisplayName returns a boolean if a field has been set.

### GetMetadataSavers

`func (o *LibraryOptions) GetMetadataSavers() []string`

GetMetadataSavers returns the MetadataSavers field if non-nil, zero value otherwise.

### GetMetadataSaversOk

`func (o *LibraryOptions) GetMetadataSaversOk() (*[]string, bool)`

GetMetadataSaversOk returns a tuple with the MetadataSavers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataSavers

`func (o *LibraryOptions) SetMetadataSavers(v []string)`

SetMetadataSavers sets MetadataSavers field to given value.

### HasMetadataSavers

`func (o *LibraryOptions) HasMetadataSavers() bool`

HasMetadataSavers returns a boolean if a field has been set.

### SetMetadataSaversNil

`func (o *LibraryOptions) SetMetadataSaversNil(b bool)`

 SetMetadataSaversNil sets the value for MetadataSavers to be an explicit nil

### UnsetMetadataSavers
`func (o *LibraryOptions) UnsetMetadataSavers()`

UnsetMetadataSavers ensures that no value is present for MetadataSavers, not even an explicit nil
### GetDisabledLocalMetadataReaders

`func (o *LibraryOptions) GetDisabledLocalMetadataReaders() []string`

GetDisabledLocalMetadataReaders returns the DisabledLocalMetadataReaders field if non-nil, zero value otherwise.

### GetDisabledLocalMetadataReadersOk

`func (o *LibraryOptions) GetDisabledLocalMetadataReadersOk() (*[]string, bool)`

GetDisabledLocalMetadataReadersOk returns a tuple with the DisabledLocalMetadataReaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledLocalMetadataReaders

`func (o *LibraryOptions) SetDisabledLocalMetadataReaders(v []string)`

SetDisabledLocalMetadataReaders sets DisabledLocalMetadataReaders field to given value.

### HasDisabledLocalMetadataReaders

`func (o *LibraryOptions) HasDisabledLocalMetadataReaders() bool`

HasDisabledLocalMetadataReaders returns a boolean if a field has been set.

### GetLocalMetadataReaderOrder

`func (o *LibraryOptions) GetLocalMetadataReaderOrder() []string`

GetLocalMetadataReaderOrder returns the LocalMetadataReaderOrder field if non-nil, zero value otherwise.

### GetLocalMetadataReaderOrderOk

`func (o *LibraryOptions) GetLocalMetadataReaderOrderOk() (*[]string, bool)`

GetLocalMetadataReaderOrderOk returns a tuple with the LocalMetadataReaderOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMetadataReaderOrder

`func (o *LibraryOptions) SetLocalMetadataReaderOrder(v []string)`

SetLocalMetadataReaderOrder sets LocalMetadataReaderOrder field to given value.

### HasLocalMetadataReaderOrder

`func (o *LibraryOptions) HasLocalMetadataReaderOrder() bool`

HasLocalMetadataReaderOrder returns a boolean if a field has been set.

### SetLocalMetadataReaderOrderNil

`func (o *LibraryOptions) SetLocalMetadataReaderOrderNil(b bool)`

 SetLocalMetadataReaderOrderNil sets the value for LocalMetadataReaderOrder to be an explicit nil

### UnsetLocalMetadataReaderOrder
`func (o *LibraryOptions) UnsetLocalMetadataReaderOrder()`

UnsetLocalMetadataReaderOrder ensures that no value is present for LocalMetadataReaderOrder, not even an explicit nil
### GetDisabledSubtitleFetchers

`func (o *LibraryOptions) GetDisabledSubtitleFetchers() []string`

GetDisabledSubtitleFetchers returns the DisabledSubtitleFetchers field if non-nil, zero value otherwise.

### GetDisabledSubtitleFetchersOk

`func (o *LibraryOptions) GetDisabledSubtitleFetchersOk() (*[]string, bool)`

GetDisabledSubtitleFetchersOk returns a tuple with the DisabledSubtitleFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledSubtitleFetchers

`func (o *LibraryOptions) SetDisabledSubtitleFetchers(v []string)`

SetDisabledSubtitleFetchers sets DisabledSubtitleFetchers field to given value.

### HasDisabledSubtitleFetchers

`func (o *LibraryOptions) HasDisabledSubtitleFetchers() bool`

HasDisabledSubtitleFetchers returns a boolean if a field has been set.

### GetSubtitleFetcherOrder

`func (o *LibraryOptions) GetSubtitleFetcherOrder() []string`

GetSubtitleFetcherOrder returns the SubtitleFetcherOrder field if non-nil, zero value otherwise.

### GetSubtitleFetcherOrderOk

`func (o *LibraryOptions) GetSubtitleFetcherOrderOk() (*[]string, bool)`

GetSubtitleFetcherOrderOk returns a tuple with the SubtitleFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleFetcherOrder

`func (o *LibraryOptions) SetSubtitleFetcherOrder(v []string)`

SetSubtitleFetcherOrder sets SubtitleFetcherOrder field to given value.

### HasSubtitleFetcherOrder

`func (o *LibraryOptions) HasSubtitleFetcherOrder() bool`

HasSubtitleFetcherOrder returns a boolean if a field has been set.

### GetDisabledMediaSegmentProviders

`func (o *LibraryOptions) GetDisabledMediaSegmentProviders() []string`

GetDisabledMediaSegmentProviders returns the DisabledMediaSegmentProviders field if non-nil, zero value otherwise.

### GetDisabledMediaSegmentProvidersOk

`func (o *LibraryOptions) GetDisabledMediaSegmentProvidersOk() (*[]string, bool)`

GetDisabledMediaSegmentProvidersOk returns a tuple with the DisabledMediaSegmentProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledMediaSegmentProviders

`func (o *LibraryOptions) SetDisabledMediaSegmentProviders(v []string)`

SetDisabledMediaSegmentProviders sets DisabledMediaSegmentProviders field to given value.

### HasDisabledMediaSegmentProviders

`func (o *LibraryOptions) HasDisabledMediaSegmentProviders() bool`

HasDisabledMediaSegmentProviders returns a boolean if a field has been set.

### GetMediaSegmentProvideOrder

`func (o *LibraryOptions) GetMediaSegmentProvideOrder() []string`

GetMediaSegmentProvideOrder returns the MediaSegmentProvideOrder field if non-nil, zero value otherwise.

### GetMediaSegmentProvideOrderOk

`func (o *LibraryOptions) GetMediaSegmentProvideOrderOk() (*[]string, bool)`

GetMediaSegmentProvideOrderOk returns a tuple with the MediaSegmentProvideOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSegmentProvideOrder

`func (o *LibraryOptions) SetMediaSegmentProvideOrder(v []string)`

SetMediaSegmentProvideOrder sets MediaSegmentProvideOrder field to given value.

### HasMediaSegmentProvideOrder

`func (o *LibraryOptions) HasMediaSegmentProvideOrder() bool`

HasMediaSegmentProvideOrder returns a boolean if a field has been set.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *LibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

GetSkipSubtitlesIfEmbeddedSubtitlesPresent returns the SkipSubtitlesIfEmbeddedSubtitlesPresent field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk

`func (o *LibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk() (*bool, bool)`

GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk returns a tuple with the SkipSubtitlesIfEmbeddedSubtitlesPresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *LibraryOptions) SetSkipSubtitlesIfEmbeddedSubtitlesPresent(v bool)`

SetSkipSubtitlesIfEmbeddedSubtitlesPresent sets SkipSubtitlesIfEmbeddedSubtitlesPresent field to given value.

### HasSkipSubtitlesIfEmbeddedSubtitlesPresent

`func (o *LibraryOptions) HasSkipSubtitlesIfEmbeddedSubtitlesPresent() bool`

HasSkipSubtitlesIfEmbeddedSubtitlesPresent returns a boolean if a field has been set.

### GetSkipSubtitlesIfAudioTrackMatches

`func (o *LibraryOptions) GetSkipSubtitlesIfAudioTrackMatches() bool`

GetSkipSubtitlesIfAudioTrackMatches returns the SkipSubtitlesIfAudioTrackMatches field if non-nil, zero value otherwise.

### GetSkipSubtitlesIfAudioTrackMatchesOk

`func (o *LibraryOptions) GetSkipSubtitlesIfAudioTrackMatchesOk() (*bool, bool)`

GetSkipSubtitlesIfAudioTrackMatchesOk returns a tuple with the SkipSubtitlesIfAudioTrackMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSubtitlesIfAudioTrackMatches

`func (o *LibraryOptions) SetSkipSubtitlesIfAudioTrackMatches(v bool)`

SetSkipSubtitlesIfAudioTrackMatches sets SkipSubtitlesIfAudioTrackMatches field to given value.

### HasSkipSubtitlesIfAudioTrackMatches

`func (o *LibraryOptions) HasSkipSubtitlesIfAudioTrackMatches() bool`

HasSkipSubtitlesIfAudioTrackMatches returns a boolean if a field has been set.

### GetSubtitleDownloadLanguages

`func (o *LibraryOptions) GetSubtitleDownloadLanguages() []string`

GetSubtitleDownloadLanguages returns the SubtitleDownloadLanguages field if non-nil, zero value otherwise.

### GetSubtitleDownloadLanguagesOk

`func (o *LibraryOptions) GetSubtitleDownloadLanguagesOk() (*[]string, bool)`

GetSubtitleDownloadLanguagesOk returns a tuple with the SubtitleDownloadLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleDownloadLanguages

`func (o *LibraryOptions) SetSubtitleDownloadLanguages(v []string)`

SetSubtitleDownloadLanguages sets SubtitleDownloadLanguages field to given value.

### HasSubtitleDownloadLanguages

`func (o *LibraryOptions) HasSubtitleDownloadLanguages() bool`

HasSubtitleDownloadLanguages returns a boolean if a field has been set.

### SetSubtitleDownloadLanguagesNil

`func (o *LibraryOptions) SetSubtitleDownloadLanguagesNil(b bool)`

 SetSubtitleDownloadLanguagesNil sets the value for SubtitleDownloadLanguages to be an explicit nil

### UnsetSubtitleDownloadLanguages
`func (o *LibraryOptions) UnsetSubtitleDownloadLanguages()`

UnsetSubtitleDownloadLanguages ensures that no value is present for SubtitleDownloadLanguages, not even an explicit nil
### GetRequirePerfectSubtitleMatch

`func (o *LibraryOptions) GetRequirePerfectSubtitleMatch() bool`

GetRequirePerfectSubtitleMatch returns the RequirePerfectSubtitleMatch field if non-nil, zero value otherwise.

### GetRequirePerfectSubtitleMatchOk

`func (o *LibraryOptions) GetRequirePerfectSubtitleMatchOk() (*bool, bool)`

GetRequirePerfectSubtitleMatchOk returns a tuple with the RequirePerfectSubtitleMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirePerfectSubtitleMatch

`func (o *LibraryOptions) SetRequirePerfectSubtitleMatch(v bool)`

SetRequirePerfectSubtitleMatch sets RequirePerfectSubtitleMatch field to given value.

### HasRequirePerfectSubtitleMatch

`func (o *LibraryOptions) HasRequirePerfectSubtitleMatch() bool`

HasRequirePerfectSubtitleMatch returns a boolean if a field has been set.

### GetSaveSubtitlesWithMedia

`func (o *LibraryOptions) GetSaveSubtitlesWithMedia() bool`

GetSaveSubtitlesWithMedia returns the SaveSubtitlesWithMedia field if non-nil, zero value otherwise.

### GetSaveSubtitlesWithMediaOk

`func (o *LibraryOptions) GetSaveSubtitlesWithMediaOk() (*bool, bool)`

GetSaveSubtitlesWithMediaOk returns a tuple with the SaveSubtitlesWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveSubtitlesWithMedia

`func (o *LibraryOptions) SetSaveSubtitlesWithMedia(v bool)`

SetSaveSubtitlesWithMedia sets SaveSubtitlesWithMedia field to given value.

### HasSaveSubtitlesWithMedia

`func (o *LibraryOptions) HasSaveSubtitlesWithMedia() bool`

HasSaveSubtitlesWithMedia returns a boolean if a field has been set.

### GetSaveLyricsWithMedia

`func (o *LibraryOptions) GetSaveLyricsWithMedia() bool`

GetSaveLyricsWithMedia returns the SaveLyricsWithMedia field if non-nil, zero value otherwise.

### GetSaveLyricsWithMediaOk

`func (o *LibraryOptions) GetSaveLyricsWithMediaOk() (*bool, bool)`

GetSaveLyricsWithMediaOk returns a tuple with the SaveLyricsWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveLyricsWithMedia

`func (o *LibraryOptions) SetSaveLyricsWithMedia(v bool)`

SetSaveLyricsWithMedia sets SaveLyricsWithMedia field to given value.

### HasSaveLyricsWithMedia

`func (o *LibraryOptions) HasSaveLyricsWithMedia() bool`

HasSaveLyricsWithMedia returns a boolean if a field has been set.

### GetSaveTrickplayWithMedia

`func (o *LibraryOptions) GetSaveTrickplayWithMedia() bool`

GetSaveTrickplayWithMedia returns the SaveTrickplayWithMedia field if non-nil, zero value otherwise.

### GetSaveTrickplayWithMediaOk

`func (o *LibraryOptions) GetSaveTrickplayWithMediaOk() (*bool, bool)`

GetSaveTrickplayWithMediaOk returns a tuple with the SaveTrickplayWithMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveTrickplayWithMedia

`func (o *LibraryOptions) SetSaveTrickplayWithMedia(v bool)`

SetSaveTrickplayWithMedia sets SaveTrickplayWithMedia field to given value.

### HasSaveTrickplayWithMedia

`func (o *LibraryOptions) HasSaveTrickplayWithMedia() bool`

HasSaveTrickplayWithMedia returns a boolean if a field has been set.

### GetDisabledLyricFetchers

`func (o *LibraryOptions) GetDisabledLyricFetchers() []string`

GetDisabledLyricFetchers returns the DisabledLyricFetchers field if non-nil, zero value otherwise.

### GetDisabledLyricFetchersOk

`func (o *LibraryOptions) GetDisabledLyricFetchersOk() (*[]string, bool)`

GetDisabledLyricFetchersOk returns a tuple with the DisabledLyricFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledLyricFetchers

`func (o *LibraryOptions) SetDisabledLyricFetchers(v []string)`

SetDisabledLyricFetchers sets DisabledLyricFetchers field to given value.

### HasDisabledLyricFetchers

`func (o *LibraryOptions) HasDisabledLyricFetchers() bool`

HasDisabledLyricFetchers returns a boolean if a field has been set.

### GetLyricFetcherOrder

`func (o *LibraryOptions) GetLyricFetcherOrder() []string`

GetLyricFetcherOrder returns the LyricFetcherOrder field if non-nil, zero value otherwise.

### GetLyricFetcherOrderOk

`func (o *LibraryOptions) GetLyricFetcherOrderOk() (*[]string, bool)`

GetLyricFetcherOrderOk returns a tuple with the LyricFetcherOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyricFetcherOrder

`func (o *LibraryOptions) SetLyricFetcherOrder(v []string)`

SetLyricFetcherOrder sets LyricFetcherOrder field to given value.

### HasLyricFetcherOrder

`func (o *LibraryOptions) HasLyricFetcherOrder() bool`

HasLyricFetcherOrder returns a boolean if a field has been set.

### GetPreferNonstandardArtistsTag

`func (o *LibraryOptions) GetPreferNonstandardArtistsTag() bool`

GetPreferNonstandardArtistsTag returns the PreferNonstandardArtistsTag field if non-nil, zero value otherwise.

### GetPreferNonstandardArtistsTagOk

`func (o *LibraryOptions) GetPreferNonstandardArtistsTagOk() (*bool, bool)`

GetPreferNonstandardArtistsTagOk returns a tuple with the PreferNonstandardArtistsTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferNonstandardArtistsTag

`func (o *LibraryOptions) SetPreferNonstandardArtistsTag(v bool)`

SetPreferNonstandardArtistsTag sets PreferNonstandardArtistsTag field to given value.

### HasPreferNonstandardArtistsTag

`func (o *LibraryOptions) HasPreferNonstandardArtistsTag() bool`

HasPreferNonstandardArtistsTag returns a boolean if a field has been set.

### GetUseCustomTagDelimiters

`func (o *LibraryOptions) GetUseCustomTagDelimiters() bool`

GetUseCustomTagDelimiters returns the UseCustomTagDelimiters field if non-nil, zero value otherwise.

### GetUseCustomTagDelimitersOk

`func (o *LibraryOptions) GetUseCustomTagDelimitersOk() (*bool, bool)`

GetUseCustomTagDelimitersOk returns a tuple with the UseCustomTagDelimiters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomTagDelimiters

`func (o *LibraryOptions) SetUseCustomTagDelimiters(v bool)`

SetUseCustomTagDelimiters sets UseCustomTagDelimiters field to given value.

### HasUseCustomTagDelimiters

`func (o *LibraryOptions) HasUseCustomTagDelimiters() bool`

HasUseCustomTagDelimiters returns a boolean if a field has been set.

### GetCustomTagDelimiters

`func (o *LibraryOptions) GetCustomTagDelimiters() []string`

GetCustomTagDelimiters returns the CustomTagDelimiters field if non-nil, zero value otherwise.

### GetCustomTagDelimitersOk

`func (o *LibraryOptions) GetCustomTagDelimitersOk() (*[]string, bool)`

GetCustomTagDelimitersOk returns a tuple with the CustomTagDelimiters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTagDelimiters

`func (o *LibraryOptions) SetCustomTagDelimiters(v []string)`

SetCustomTagDelimiters sets CustomTagDelimiters field to given value.

### HasCustomTagDelimiters

`func (o *LibraryOptions) HasCustomTagDelimiters() bool`

HasCustomTagDelimiters returns a boolean if a field has been set.

### GetDelimiterWhitelist

`func (o *LibraryOptions) GetDelimiterWhitelist() []string`

GetDelimiterWhitelist returns the DelimiterWhitelist field if non-nil, zero value otherwise.

### GetDelimiterWhitelistOk

`func (o *LibraryOptions) GetDelimiterWhitelistOk() (*[]string, bool)`

GetDelimiterWhitelistOk returns a tuple with the DelimiterWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiterWhitelist

`func (o *LibraryOptions) SetDelimiterWhitelist(v []string)`

SetDelimiterWhitelist sets DelimiterWhitelist field to given value.

### HasDelimiterWhitelist

`func (o *LibraryOptions) HasDelimiterWhitelist() bool`

HasDelimiterWhitelist returns a boolean if a field has been set.

### GetAutomaticallyAddToCollection

`func (o *LibraryOptions) GetAutomaticallyAddToCollection() bool`

GetAutomaticallyAddToCollection returns the AutomaticallyAddToCollection field if non-nil, zero value otherwise.

### GetAutomaticallyAddToCollectionOk

`func (o *LibraryOptions) GetAutomaticallyAddToCollectionOk() (*bool, bool)`

GetAutomaticallyAddToCollectionOk returns a tuple with the AutomaticallyAddToCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticallyAddToCollection

`func (o *LibraryOptions) SetAutomaticallyAddToCollection(v bool)`

SetAutomaticallyAddToCollection sets AutomaticallyAddToCollection field to given value.

### HasAutomaticallyAddToCollection

`func (o *LibraryOptions) HasAutomaticallyAddToCollection() bool`

HasAutomaticallyAddToCollection returns a boolean if a field has been set.

### GetAllowEmbeddedSubtitles

`func (o *LibraryOptions) GetAllowEmbeddedSubtitles() EmbeddedSubtitleOptions`

GetAllowEmbeddedSubtitles returns the AllowEmbeddedSubtitles field if non-nil, zero value otherwise.

### GetAllowEmbeddedSubtitlesOk

`func (o *LibraryOptions) GetAllowEmbeddedSubtitlesOk() (*EmbeddedSubtitleOptions, bool)`

GetAllowEmbeddedSubtitlesOk returns a tuple with the AllowEmbeddedSubtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmbeddedSubtitles

`func (o *LibraryOptions) SetAllowEmbeddedSubtitles(v EmbeddedSubtitleOptions)`

SetAllowEmbeddedSubtitles sets AllowEmbeddedSubtitles field to given value.

### HasAllowEmbeddedSubtitles

`func (o *LibraryOptions) HasAllowEmbeddedSubtitles() bool`

HasAllowEmbeddedSubtitles returns a boolean if a field has been set.

### GetTypeOptions

`func (o *LibraryOptions) GetTypeOptions() []TypeOptions`

GetTypeOptions returns the TypeOptions field if non-nil, zero value otherwise.

### GetTypeOptionsOk

`func (o *LibraryOptions) GetTypeOptionsOk() (*[]TypeOptions, bool)`

GetTypeOptionsOk returns a tuple with the TypeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOptions

`func (o *LibraryOptions) SetTypeOptions(v []TypeOptions)`

SetTypeOptions sets TypeOptions field to given value.

### HasTypeOptions

`func (o *LibraryOptions) HasTypeOptions() bool`

HasTypeOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


