# ServerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogFileRetentionDays** | Pointer to **int32** | Gets or sets the number of days we should retain log files. | [optional] 
**IsStartupWizardCompleted** | Pointer to **bool** | Gets or sets a value indicating whether this instance is first run. | [optional] 
**CachePath** | Pointer to **NullableString** | Gets or sets the cache path. | [optional] 
**PreviousVersion** | Pointer to **NullableString** | Gets or sets the last known version that was ran using the configuration. | [optional] 
**PreviousVersionStr** | Pointer to **NullableString** | Gets or sets the stringified PreviousVersion to be stored/loaded,  because System.Version itself isn&#39;t xml-serializable. | [optional] 
**EnableMetrics** | Pointer to **bool** | Gets or sets a value indicating whether to enable prometheus metrics exporting. | [optional] 
**EnableNormalizedItemByNameIds** | Pointer to **bool** |  | [optional] 
**IsPortAuthorized** | Pointer to **bool** | Gets or sets a value indicating whether this instance is port authorized. | [optional] 
**QuickConnectAvailable** | Pointer to **bool** | Gets or sets a value indicating whether quick connect is available for use on this server. | [optional] 
**EnableCaseSensitiveItemIds** | Pointer to **bool** | Gets or sets a value indicating whether [enable case sensitive item ids]. | [optional] 
**DisableLiveTvChannelUserDataName** | Pointer to **bool** |  | [optional] 
**MetadataPath** | Pointer to **string** | Gets or sets the metadata path. | [optional] 
**MetadataNetworkPath** | Pointer to **string** |  | [optional] 
**PreferredMetadataLanguage** | Pointer to **string** | Gets or sets the preferred metadata language. | [optional] 
**MetadataCountryCode** | Pointer to **string** | Gets or sets the metadata country code. | [optional] 
**SortReplaceCharacters** | Pointer to **[]string** | Gets or sets characters to be replaced with a &#39; &#39; in strings to create a sort name. | [optional] 
**SortRemoveCharacters** | Pointer to **[]string** | Gets or sets characters to be removed from strings to create a sort name. | [optional] 
**SortRemoveWords** | Pointer to **[]string** | Gets or sets words to be removed from strings to create a sort name. | [optional] 
**MinResumePct** | Pointer to **int32** | Gets or sets the minimum percentage of an item that must be played in order for playstate to be updated. | [optional] 
**MaxResumePct** | Pointer to **int32** | Gets or sets the maximum percentage of an item that can be played while still saving playstate. If this percentage is crossed playstate will be reset to the beginning and the item will be marked watched. | [optional] 
**MinResumeDurationSeconds** | Pointer to **int32** | Gets or sets the minimum duration that an item must have in order to be eligible for playstate updates.. | [optional] 
**MinAudiobookResume** | Pointer to **int32** | Gets or sets the minimum minutes of a book that must be played in order for playstate to be updated. | [optional] 
**MaxAudiobookResume** | Pointer to **int32** | Gets or sets the remaining minutes of a book that can be played while still saving playstate. If this percentage is crossed playstate will be reset to the beginning and the item will be marked watched. | [optional] 
**LibraryMonitorDelay** | Pointer to **int32** | Gets or sets the delay in seconds that we will wait after a file system change to try and discover what has been added/removed  Some delay is necessary with some items because their creation is not atomic.  It involves the creation of several  different directories and files. | [optional] 
**ImageSavingConvention** | Pointer to [**ImageSavingConvention**](ImageSavingConvention.md) | Gets or sets the image saving convention. | [optional] 
**MetadataOptions** | Pointer to [**[]MetadataOptions**](MetadataOptions.md) |  | [optional] 
**SkipDeserializationForBasicTypes** | Pointer to **bool** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**UICulture** | Pointer to **string** |  | [optional] 
**SaveMetadataHidden** | Pointer to **bool** |  | [optional] 
**ContentTypes** | Pointer to [**[]NameValuePair**](NameValuePair.md) |  | [optional] 
**RemoteClientBitrateLimit** | Pointer to **int32** |  | [optional] 
**EnableFolderView** | Pointer to **bool** |  | [optional] 
**EnableGroupingIntoCollections** | Pointer to **bool** |  | [optional] 
**DisplaySpecialsWithinSeasons** | Pointer to **bool** |  | [optional] 
**CodecsUsed** | Pointer to **[]string** |  | [optional] 
**PluginRepositories** | Pointer to [**[]RepositoryInfo**](RepositoryInfo.md) |  | [optional] 
**EnableExternalContentInSuggestions** | Pointer to **bool** |  | [optional] 
**ImageExtractionTimeoutMs** | Pointer to **int32** |  | [optional] 
**PathSubstitutions** | Pointer to [**[]PathSubstitution**](PathSubstitution.md) |  | [optional] 
**EnableSlowResponseWarning** | Pointer to **bool** | Gets or sets a value indicating whether slow server responses should be logged as a warning. | [optional] 
**SlowResponseThresholdMs** | Pointer to **int64** | Gets or sets the threshold for the slow response time warning in ms. | [optional] 
**CorsHosts** | Pointer to **[]string** | Gets or sets the cors hosts. | [optional] 
**ActivityLogRetentionDays** | Pointer to **NullableInt32** | Gets or sets the number of days we should retain activity logs. | [optional] 
**LibraryScanFanoutConcurrency** | Pointer to **int32** | Gets or sets the how the library scan fans out. | [optional] 
**LibraryMetadataRefreshConcurrency** | Pointer to **int32** | Gets or sets the how many metadata refreshes can run concurrently. | [optional] 
**RemoveOldPlugins** | Pointer to **bool** | Gets or sets a value indicating whether older plugins should automatically be deleted from the plugin folder. | [optional] 
**AllowClientLogUpload** | Pointer to **bool** | Gets or sets a value indicating whether clients should be allowed to upload logs. | [optional] 

## Methods

### NewServerConfiguration

`func NewServerConfiguration() *ServerConfiguration`

NewServerConfiguration instantiates a new ServerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigurationWithDefaults

`func NewServerConfigurationWithDefaults() *ServerConfiguration`

NewServerConfigurationWithDefaults instantiates a new ServerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogFileRetentionDays

`func (o *ServerConfiguration) GetLogFileRetentionDays() int32`

GetLogFileRetentionDays returns the LogFileRetentionDays field if non-nil, zero value otherwise.

### GetLogFileRetentionDaysOk

`func (o *ServerConfiguration) GetLogFileRetentionDaysOk() (*int32, bool)`

GetLogFileRetentionDaysOk returns a tuple with the LogFileRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFileRetentionDays

`func (o *ServerConfiguration) SetLogFileRetentionDays(v int32)`

SetLogFileRetentionDays sets LogFileRetentionDays field to given value.

### HasLogFileRetentionDays

`func (o *ServerConfiguration) HasLogFileRetentionDays() bool`

HasLogFileRetentionDays returns a boolean if a field has been set.

### GetIsStartupWizardCompleted

`func (o *ServerConfiguration) GetIsStartupWizardCompleted() bool`

GetIsStartupWizardCompleted returns the IsStartupWizardCompleted field if non-nil, zero value otherwise.

### GetIsStartupWizardCompletedOk

`func (o *ServerConfiguration) GetIsStartupWizardCompletedOk() (*bool, bool)`

GetIsStartupWizardCompletedOk returns a tuple with the IsStartupWizardCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStartupWizardCompleted

`func (o *ServerConfiguration) SetIsStartupWizardCompleted(v bool)`

SetIsStartupWizardCompleted sets IsStartupWizardCompleted field to given value.

### HasIsStartupWizardCompleted

`func (o *ServerConfiguration) HasIsStartupWizardCompleted() bool`

HasIsStartupWizardCompleted returns a boolean if a field has been set.

### GetCachePath

`func (o *ServerConfiguration) GetCachePath() string`

GetCachePath returns the CachePath field if non-nil, zero value otherwise.

### GetCachePathOk

`func (o *ServerConfiguration) GetCachePathOk() (*string, bool)`

GetCachePathOk returns a tuple with the CachePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachePath

`func (o *ServerConfiguration) SetCachePath(v string)`

SetCachePath sets CachePath field to given value.

### HasCachePath

`func (o *ServerConfiguration) HasCachePath() bool`

HasCachePath returns a boolean if a field has been set.

### SetCachePathNil

`func (o *ServerConfiguration) SetCachePathNil(b bool)`

 SetCachePathNil sets the value for CachePath to be an explicit nil

### UnsetCachePath
`func (o *ServerConfiguration) UnsetCachePath()`

UnsetCachePath ensures that no value is present for CachePath, not even an explicit nil
### GetPreviousVersion

`func (o *ServerConfiguration) GetPreviousVersion() string`

GetPreviousVersion returns the PreviousVersion field if non-nil, zero value otherwise.

### GetPreviousVersionOk

`func (o *ServerConfiguration) GetPreviousVersionOk() (*string, bool)`

GetPreviousVersionOk returns a tuple with the PreviousVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersion

`func (o *ServerConfiguration) SetPreviousVersion(v string)`

SetPreviousVersion sets PreviousVersion field to given value.

### HasPreviousVersion

`func (o *ServerConfiguration) HasPreviousVersion() bool`

HasPreviousVersion returns a boolean if a field has been set.

### SetPreviousVersionNil

`func (o *ServerConfiguration) SetPreviousVersionNil(b bool)`

 SetPreviousVersionNil sets the value for PreviousVersion to be an explicit nil

### UnsetPreviousVersion
`func (o *ServerConfiguration) UnsetPreviousVersion()`

UnsetPreviousVersion ensures that no value is present for PreviousVersion, not even an explicit nil
### GetPreviousVersionStr

`func (o *ServerConfiguration) GetPreviousVersionStr() string`

GetPreviousVersionStr returns the PreviousVersionStr field if non-nil, zero value otherwise.

### GetPreviousVersionStrOk

`func (o *ServerConfiguration) GetPreviousVersionStrOk() (*string, bool)`

GetPreviousVersionStrOk returns a tuple with the PreviousVersionStr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVersionStr

`func (o *ServerConfiguration) SetPreviousVersionStr(v string)`

SetPreviousVersionStr sets PreviousVersionStr field to given value.

### HasPreviousVersionStr

`func (o *ServerConfiguration) HasPreviousVersionStr() bool`

HasPreviousVersionStr returns a boolean if a field has been set.

### SetPreviousVersionStrNil

`func (o *ServerConfiguration) SetPreviousVersionStrNil(b bool)`

 SetPreviousVersionStrNil sets the value for PreviousVersionStr to be an explicit nil

### UnsetPreviousVersionStr
`func (o *ServerConfiguration) UnsetPreviousVersionStr()`

UnsetPreviousVersionStr ensures that no value is present for PreviousVersionStr, not even an explicit nil
### GetEnableMetrics

`func (o *ServerConfiguration) GetEnableMetrics() bool`

GetEnableMetrics returns the EnableMetrics field if non-nil, zero value otherwise.

### GetEnableMetricsOk

`func (o *ServerConfiguration) GetEnableMetricsOk() (*bool, bool)`

GetEnableMetricsOk returns a tuple with the EnableMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMetrics

`func (o *ServerConfiguration) SetEnableMetrics(v bool)`

SetEnableMetrics sets EnableMetrics field to given value.

### HasEnableMetrics

`func (o *ServerConfiguration) HasEnableMetrics() bool`

HasEnableMetrics returns a boolean if a field has been set.

### GetEnableNormalizedItemByNameIds

`func (o *ServerConfiguration) GetEnableNormalizedItemByNameIds() bool`

GetEnableNormalizedItemByNameIds returns the EnableNormalizedItemByNameIds field if non-nil, zero value otherwise.

### GetEnableNormalizedItemByNameIdsOk

`func (o *ServerConfiguration) GetEnableNormalizedItemByNameIdsOk() (*bool, bool)`

GetEnableNormalizedItemByNameIdsOk returns a tuple with the EnableNormalizedItemByNameIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNormalizedItemByNameIds

`func (o *ServerConfiguration) SetEnableNormalizedItemByNameIds(v bool)`

SetEnableNormalizedItemByNameIds sets EnableNormalizedItemByNameIds field to given value.

### HasEnableNormalizedItemByNameIds

`func (o *ServerConfiguration) HasEnableNormalizedItemByNameIds() bool`

HasEnableNormalizedItemByNameIds returns a boolean if a field has been set.

### GetIsPortAuthorized

`func (o *ServerConfiguration) GetIsPortAuthorized() bool`

GetIsPortAuthorized returns the IsPortAuthorized field if non-nil, zero value otherwise.

### GetIsPortAuthorizedOk

`func (o *ServerConfiguration) GetIsPortAuthorizedOk() (*bool, bool)`

GetIsPortAuthorizedOk returns a tuple with the IsPortAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPortAuthorized

`func (o *ServerConfiguration) SetIsPortAuthorized(v bool)`

SetIsPortAuthorized sets IsPortAuthorized field to given value.

### HasIsPortAuthorized

`func (o *ServerConfiguration) HasIsPortAuthorized() bool`

HasIsPortAuthorized returns a boolean if a field has been set.

### GetQuickConnectAvailable

`func (o *ServerConfiguration) GetQuickConnectAvailable() bool`

GetQuickConnectAvailable returns the QuickConnectAvailable field if non-nil, zero value otherwise.

### GetQuickConnectAvailableOk

`func (o *ServerConfiguration) GetQuickConnectAvailableOk() (*bool, bool)`

GetQuickConnectAvailableOk returns a tuple with the QuickConnectAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickConnectAvailable

`func (o *ServerConfiguration) SetQuickConnectAvailable(v bool)`

SetQuickConnectAvailable sets QuickConnectAvailable field to given value.

### HasQuickConnectAvailable

`func (o *ServerConfiguration) HasQuickConnectAvailable() bool`

HasQuickConnectAvailable returns a boolean if a field has been set.

### GetEnableCaseSensitiveItemIds

`func (o *ServerConfiguration) GetEnableCaseSensitiveItemIds() bool`

GetEnableCaseSensitiveItemIds returns the EnableCaseSensitiveItemIds field if non-nil, zero value otherwise.

### GetEnableCaseSensitiveItemIdsOk

`func (o *ServerConfiguration) GetEnableCaseSensitiveItemIdsOk() (*bool, bool)`

GetEnableCaseSensitiveItemIdsOk returns a tuple with the EnableCaseSensitiveItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCaseSensitiveItemIds

`func (o *ServerConfiguration) SetEnableCaseSensitiveItemIds(v bool)`

SetEnableCaseSensitiveItemIds sets EnableCaseSensitiveItemIds field to given value.

### HasEnableCaseSensitiveItemIds

`func (o *ServerConfiguration) HasEnableCaseSensitiveItemIds() bool`

HasEnableCaseSensitiveItemIds returns a boolean if a field has been set.

### GetDisableLiveTvChannelUserDataName

`func (o *ServerConfiguration) GetDisableLiveTvChannelUserDataName() bool`

GetDisableLiveTvChannelUserDataName returns the DisableLiveTvChannelUserDataName field if non-nil, zero value otherwise.

### GetDisableLiveTvChannelUserDataNameOk

`func (o *ServerConfiguration) GetDisableLiveTvChannelUserDataNameOk() (*bool, bool)`

GetDisableLiveTvChannelUserDataNameOk returns a tuple with the DisableLiveTvChannelUserDataName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLiveTvChannelUserDataName

`func (o *ServerConfiguration) SetDisableLiveTvChannelUserDataName(v bool)`

SetDisableLiveTvChannelUserDataName sets DisableLiveTvChannelUserDataName field to given value.

### HasDisableLiveTvChannelUserDataName

`func (o *ServerConfiguration) HasDisableLiveTvChannelUserDataName() bool`

HasDisableLiveTvChannelUserDataName returns a boolean if a field has been set.

### GetMetadataPath

`func (o *ServerConfiguration) GetMetadataPath() string`

GetMetadataPath returns the MetadataPath field if non-nil, zero value otherwise.

### GetMetadataPathOk

`func (o *ServerConfiguration) GetMetadataPathOk() (*string, bool)`

GetMetadataPathOk returns a tuple with the MetadataPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataPath

`func (o *ServerConfiguration) SetMetadataPath(v string)`

SetMetadataPath sets MetadataPath field to given value.

### HasMetadataPath

`func (o *ServerConfiguration) HasMetadataPath() bool`

HasMetadataPath returns a boolean if a field has been set.

### GetMetadataNetworkPath

`func (o *ServerConfiguration) GetMetadataNetworkPath() string`

GetMetadataNetworkPath returns the MetadataNetworkPath field if non-nil, zero value otherwise.

### GetMetadataNetworkPathOk

`func (o *ServerConfiguration) GetMetadataNetworkPathOk() (*string, bool)`

GetMetadataNetworkPathOk returns a tuple with the MetadataNetworkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataNetworkPath

`func (o *ServerConfiguration) SetMetadataNetworkPath(v string)`

SetMetadataNetworkPath sets MetadataNetworkPath field to given value.

### HasMetadataNetworkPath

`func (o *ServerConfiguration) HasMetadataNetworkPath() bool`

HasMetadataNetworkPath returns a boolean if a field has been set.

### GetPreferredMetadataLanguage

`func (o *ServerConfiguration) GetPreferredMetadataLanguage() string`

GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field if non-nil, zero value otherwise.

### GetPreferredMetadataLanguageOk

`func (o *ServerConfiguration) GetPreferredMetadataLanguageOk() (*string, bool)`

GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredMetadataLanguage

`func (o *ServerConfiguration) SetPreferredMetadataLanguage(v string)`

SetPreferredMetadataLanguage sets PreferredMetadataLanguage field to given value.

### HasPreferredMetadataLanguage

`func (o *ServerConfiguration) HasPreferredMetadataLanguage() bool`

HasPreferredMetadataLanguage returns a boolean if a field has been set.

### GetMetadataCountryCode

`func (o *ServerConfiguration) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *ServerConfiguration) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *ServerConfiguration) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *ServerConfiguration) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### GetSortReplaceCharacters

`func (o *ServerConfiguration) GetSortReplaceCharacters() []string`

GetSortReplaceCharacters returns the SortReplaceCharacters field if non-nil, zero value otherwise.

### GetSortReplaceCharactersOk

`func (o *ServerConfiguration) GetSortReplaceCharactersOk() (*[]string, bool)`

GetSortReplaceCharactersOk returns a tuple with the SortReplaceCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortReplaceCharacters

`func (o *ServerConfiguration) SetSortReplaceCharacters(v []string)`

SetSortReplaceCharacters sets SortReplaceCharacters field to given value.

### HasSortReplaceCharacters

`func (o *ServerConfiguration) HasSortReplaceCharacters() bool`

HasSortReplaceCharacters returns a boolean if a field has been set.

### GetSortRemoveCharacters

`func (o *ServerConfiguration) GetSortRemoveCharacters() []string`

GetSortRemoveCharacters returns the SortRemoveCharacters field if non-nil, zero value otherwise.

### GetSortRemoveCharactersOk

`func (o *ServerConfiguration) GetSortRemoveCharactersOk() (*[]string, bool)`

GetSortRemoveCharactersOk returns a tuple with the SortRemoveCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortRemoveCharacters

`func (o *ServerConfiguration) SetSortRemoveCharacters(v []string)`

SetSortRemoveCharacters sets SortRemoveCharacters field to given value.

### HasSortRemoveCharacters

`func (o *ServerConfiguration) HasSortRemoveCharacters() bool`

HasSortRemoveCharacters returns a boolean if a field has been set.

### GetSortRemoveWords

`func (o *ServerConfiguration) GetSortRemoveWords() []string`

GetSortRemoveWords returns the SortRemoveWords field if non-nil, zero value otherwise.

### GetSortRemoveWordsOk

`func (o *ServerConfiguration) GetSortRemoveWordsOk() (*[]string, bool)`

GetSortRemoveWordsOk returns a tuple with the SortRemoveWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortRemoveWords

`func (o *ServerConfiguration) SetSortRemoveWords(v []string)`

SetSortRemoveWords sets SortRemoveWords field to given value.

### HasSortRemoveWords

`func (o *ServerConfiguration) HasSortRemoveWords() bool`

HasSortRemoveWords returns a boolean if a field has been set.

### GetMinResumePct

`func (o *ServerConfiguration) GetMinResumePct() int32`

GetMinResumePct returns the MinResumePct field if non-nil, zero value otherwise.

### GetMinResumePctOk

`func (o *ServerConfiguration) GetMinResumePctOk() (*int32, bool)`

GetMinResumePctOk returns a tuple with the MinResumePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinResumePct

`func (o *ServerConfiguration) SetMinResumePct(v int32)`

SetMinResumePct sets MinResumePct field to given value.

### HasMinResumePct

`func (o *ServerConfiguration) HasMinResumePct() bool`

HasMinResumePct returns a boolean if a field has been set.

### GetMaxResumePct

`func (o *ServerConfiguration) GetMaxResumePct() int32`

GetMaxResumePct returns the MaxResumePct field if non-nil, zero value otherwise.

### GetMaxResumePctOk

`func (o *ServerConfiguration) GetMaxResumePctOk() (*int32, bool)`

GetMaxResumePctOk returns a tuple with the MaxResumePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResumePct

`func (o *ServerConfiguration) SetMaxResumePct(v int32)`

SetMaxResumePct sets MaxResumePct field to given value.

### HasMaxResumePct

`func (o *ServerConfiguration) HasMaxResumePct() bool`

HasMaxResumePct returns a boolean if a field has been set.

### GetMinResumeDurationSeconds

`func (o *ServerConfiguration) GetMinResumeDurationSeconds() int32`

GetMinResumeDurationSeconds returns the MinResumeDurationSeconds field if non-nil, zero value otherwise.

### GetMinResumeDurationSecondsOk

`func (o *ServerConfiguration) GetMinResumeDurationSecondsOk() (*int32, bool)`

GetMinResumeDurationSecondsOk returns a tuple with the MinResumeDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinResumeDurationSeconds

`func (o *ServerConfiguration) SetMinResumeDurationSeconds(v int32)`

SetMinResumeDurationSeconds sets MinResumeDurationSeconds field to given value.

### HasMinResumeDurationSeconds

`func (o *ServerConfiguration) HasMinResumeDurationSeconds() bool`

HasMinResumeDurationSeconds returns a boolean if a field has been set.

### GetMinAudiobookResume

`func (o *ServerConfiguration) GetMinAudiobookResume() int32`

GetMinAudiobookResume returns the MinAudiobookResume field if non-nil, zero value otherwise.

### GetMinAudiobookResumeOk

`func (o *ServerConfiguration) GetMinAudiobookResumeOk() (*int32, bool)`

GetMinAudiobookResumeOk returns a tuple with the MinAudiobookResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAudiobookResume

`func (o *ServerConfiguration) SetMinAudiobookResume(v int32)`

SetMinAudiobookResume sets MinAudiobookResume field to given value.

### HasMinAudiobookResume

`func (o *ServerConfiguration) HasMinAudiobookResume() bool`

HasMinAudiobookResume returns a boolean if a field has been set.

### GetMaxAudiobookResume

`func (o *ServerConfiguration) GetMaxAudiobookResume() int32`

GetMaxAudiobookResume returns the MaxAudiobookResume field if non-nil, zero value otherwise.

### GetMaxAudiobookResumeOk

`func (o *ServerConfiguration) GetMaxAudiobookResumeOk() (*int32, bool)`

GetMaxAudiobookResumeOk returns a tuple with the MaxAudiobookResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAudiobookResume

`func (o *ServerConfiguration) SetMaxAudiobookResume(v int32)`

SetMaxAudiobookResume sets MaxAudiobookResume field to given value.

### HasMaxAudiobookResume

`func (o *ServerConfiguration) HasMaxAudiobookResume() bool`

HasMaxAudiobookResume returns a boolean if a field has been set.

### GetLibraryMonitorDelay

`func (o *ServerConfiguration) GetLibraryMonitorDelay() int32`

GetLibraryMonitorDelay returns the LibraryMonitorDelay field if non-nil, zero value otherwise.

### GetLibraryMonitorDelayOk

`func (o *ServerConfiguration) GetLibraryMonitorDelayOk() (*int32, bool)`

GetLibraryMonitorDelayOk returns a tuple with the LibraryMonitorDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryMonitorDelay

`func (o *ServerConfiguration) SetLibraryMonitorDelay(v int32)`

SetLibraryMonitorDelay sets LibraryMonitorDelay field to given value.

### HasLibraryMonitorDelay

`func (o *ServerConfiguration) HasLibraryMonitorDelay() bool`

HasLibraryMonitorDelay returns a boolean if a field has been set.

### GetImageSavingConvention

`func (o *ServerConfiguration) GetImageSavingConvention() ImageSavingConvention`

GetImageSavingConvention returns the ImageSavingConvention field if non-nil, zero value otherwise.

### GetImageSavingConventionOk

`func (o *ServerConfiguration) GetImageSavingConventionOk() (*ImageSavingConvention, bool)`

GetImageSavingConventionOk returns a tuple with the ImageSavingConvention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSavingConvention

`func (o *ServerConfiguration) SetImageSavingConvention(v ImageSavingConvention)`

SetImageSavingConvention sets ImageSavingConvention field to given value.

### HasImageSavingConvention

`func (o *ServerConfiguration) HasImageSavingConvention() bool`

HasImageSavingConvention returns a boolean if a field has been set.

### GetMetadataOptions

`func (o *ServerConfiguration) GetMetadataOptions() []MetadataOptions`

GetMetadataOptions returns the MetadataOptions field if non-nil, zero value otherwise.

### GetMetadataOptionsOk

`func (o *ServerConfiguration) GetMetadataOptionsOk() (*[]MetadataOptions, bool)`

GetMetadataOptionsOk returns a tuple with the MetadataOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataOptions

`func (o *ServerConfiguration) SetMetadataOptions(v []MetadataOptions)`

SetMetadataOptions sets MetadataOptions field to given value.

### HasMetadataOptions

`func (o *ServerConfiguration) HasMetadataOptions() bool`

HasMetadataOptions returns a boolean if a field has been set.

### GetSkipDeserializationForBasicTypes

`func (o *ServerConfiguration) GetSkipDeserializationForBasicTypes() bool`

GetSkipDeserializationForBasicTypes returns the SkipDeserializationForBasicTypes field if non-nil, zero value otherwise.

### GetSkipDeserializationForBasicTypesOk

`func (o *ServerConfiguration) GetSkipDeserializationForBasicTypesOk() (*bool, bool)`

GetSkipDeserializationForBasicTypesOk returns a tuple with the SkipDeserializationForBasicTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDeserializationForBasicTypes

`func (o *ServerConfiguration) SetSkipDeserializationForBasicTypes(v bool)`

SetSkipDeserializationForBasicTypes sets SkipDeserializationForBasicTypes field to given value.

### HasSkipDeserializationForBasicTypes

`func (o *ServerConfiguration) HasSkipDeserializationForBasicTypes() bool`

HasSkipDeserializationForBasicTypes returns a boolean if a field has been set.

### GetServerName

`func (o *ServerConfiguration) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ServerConfiguration) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ServerConfiguration) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ServerConfiguration) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetUICulture

`func (o *ServerConfiguration) GetUICulture() string`

GetUICulture returns the UICulture field if non-nil, zero value otherwise.

### GetUICultureOk

`func (o *ServerConfiguration) GetUICultureOk() (*string, bool)`

GetUICultureOk returns a tuple with the UICulture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUICulture

`func (o *ServerConfiguration) SetUICulture(v string)`

SetUICulture sets UICulture field to given value.

### HasUICulture

`func (o *ServerConfiguration) HasUICulture() bool`

HasUICulture returns a boolean if a field has been set.

### GetSaveMetadataHidden

`func (o *ServerConfiguration) GetSaveMetadataHidden() bool`

GetSaveMetadataHidden returns the SaveMetadataHidden field if non-nil, zero value otherwise.

### GetSaveMetadataHiddenOk

`func (o *ServerConfiguration) GetSaveMetadataHiddenOk() (*bool, bool)`

GetSaveMetadataHiddenOk returns a tuple with the SaveMetadataHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveMetadataHidden

`func (o *ServerConfiguration) SetSaveMetadataHidden(v bool)`

SetSaveMetadataHidden sets SaveMetadataHidden field to given value.

### HasSaveMetadataHidden

`func (o *ServerConfiguration) HasSaveMetadataHidden() bool`

HasSaveMetadataHidden returns a boolean if a field has been set.

### GetContentTypes

`func (o *ServerConfiguration) GetContentTypes() []NameValuePair`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *ServerConfiguration) GetContentTypesOk() (*[]NameValuePair, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *ServerConfiguration) SetContentTypes(v []NameValuePair)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *ServerConfiguration) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetRemoteClientBitrateLimit

`func (o *ServerConfiguration) GetRemoteClientBitrateLimit() int32`

GetRemoteClientBitrateLimit returns the RemoteClientBitrateLimit field if non-nil, zero value otherwise.

### GetRemoteClientBitrateLimitOk

`func (o *ServerConfiguration) GetRemoteClientBitrateLimitOk() (*int32, bool)`

GetRemoteClientBitrateLimitOk returns a tuple with the RemoteClientBitrateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteClientBitrateLimit

`func (o *ServerConfiguration) SetRemoteClientBitrateLimit(v int32)`

SetRemoteClientBitrateLimit sets RemoteClientBitrateLimit field to given value.

### HasRemoteClientBitrateLimit

`func (o *ServerConfiguration) HasRemoteClientBitrateLimit() bool`

HasRemoteClientBitrateLimit returns a boolean if a field has been set.

### GetEnableFolderView

`func (o *ServerConfiguration) GetEnableFolderView() bool`

GetEnableFolderView returns the EnableFolderView field if non-nil, zero value otherwise.

### GetEnableFolderViewOk

`func (o *ServerConfiguration) GetEnableFolderViewOk() (*bool, bool)`

GetEnableFolderViewOk returns a tuple with the EnableFolderView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFolderView

`func (o *ServerConfiguration) SetEnableFolderView(v bool)`

SetEnableFolderView sets EnableFolderView field to given value.

### HasEnableFolderView

`func (o *ServerConfiguration) HasEnableFolderView() bool`

HasEnableFolderView returns a boolean if a field has been set.

### GetEnableGroupingIntoCollections

`func (o *ServerConfiguration) GetEnableGroupingIntoCollections() bool`

GetEnableGroupingIntoCollections returns the EnableGroupingIntoCollections field if non-nil, zero value otherwise.

### GetEnableGroupingIntoCollectionsOk

`func (o *ServerConfiguration) GetEnableGroupingIntoCollectionsOk() (*bool, bool)`

GetEnableGroupingIntoCollectionsOk returns a tuple with the EnableGroupingIntoCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGroupingIntoCollections

`func (o *ServerConfiguration) SetEnableGroupingIntoCollections(v bool)`

SetEnableGroupingIntoCollections sets EnableGroupingIntoCollections field to given value.

### HasEnableGroupingIntoCollections

`func (o *ServerConfiguration) HasEnableGroupingIntoCollections() bool`

HasEnableGroupingIntoCollections returns a boolean if a field has been set.

### GetDisplaySpecialsWithinSeasons

`func (o *ServerConfiguration) GetDisplaySpecialsWithinSeasons() bool`

GetDisplaySpecialsWithinSeasons returns the DisplaySpecialsWithinSeasons field if non-nil, zero value otherwise.

### GetDisplaySpecialsWithinSeasonsOk

`func (o *ServerConfiguration) GetDisplaySpecialsWithinSeasonsOk() (*bool, bool)`

GetDisplaySpecialsWithinSeasonsOk returns a tuple with the DisplaySpecialsWithinSeasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySpecialsWithinSeasons

`func (o *ServerConfiguration) SetDisplaySpecialsWithinSeasons(v bool)`

SetDisplaySpecialsWithinSeasons sets DisplaySpecialsWithinSeasons field to given value.

### HasDisplaySpecialsWithinSeasons

`func (o *ServerConfiguration) HasDisplaySpecialsWithinSeasons() bool`

HasDisplaySpecialsWithinSeasons returns a boolean if a field has been set.

### GetCodecsUsed

`func (o *ServerConfiguration) GetCodecsUsed() []string`

GetCodecsUsed returns the CodecsUsed field if non-nil, zero value otherwise.

### GetCodecsUsedOk

`func (o *ServerConfiguration) GetCodecsUsedOk() (*[]string, bool)`

GetCodecsUsedOk returns a tuple with the CodecsUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecsUsed

`func (o *ServerConfiguration) SetCodecsUsed(v []string)`

SetCodecsUsed sets CodecsUsed field to given value.

### HasCodecsUsed

`func (o *ServerConfiguration) HasCodecsUsed() bool`

HasCodecsUsed returns a boolean if a field has been set.

### GetPluginRepositories

`func (o *ServerConfiguration) GetPluginRepositories() []RepositoryInfo`

GetPluginRepositories returns the PluginRepositories field if non-nil, zero value otherwise.

### GetPluginRepositoriesOk

`func (o *ServerConfiguration) GetPluginRepositoriesOk() (*[]RepositoryInfo, bool)`

GetPluginRepositoriesOk returns a tuple with the PluginRepositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginRepositories

`func (o *ServerConfiguration) SetPluginRepositories(v []RepositoryInfo)`

SetPluginRepositories sets PluginRepositories field to given value.

### HasPluginRepositories

`func (o *ServerConfiguration) HasPluginRepositories() bool`

HasPluginRepositories returns a boolean if a field has been set.

### GetEnableExternalContentInSuggestions

`func (o *ServerConfiguration) GetEnableExternalContentInSuggestions() bool`

GetEnableExternalContentInSuggestions returns the EnableExternalContentInSuggestions field if non-nil, zero value otherwise.

### GetEnableExternalContentInSuggestionsOk

`func (o *ServerConfiguration) GetEnableExternalContentInSuggestionsOk() (*bool, bool)`

GetEnableExternalContentInSuggestionsOk returns a tuple with the EnableExternalContentInSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalContentInSuggestions

`func (o *ServerConfiguration) SetEnableExternalContentInSuggestions(v bool)`

SetEnableExternalContentInSuggestions sets EnableExternalContentInSuggestions field to given value.

### HasEnableExternalContentInSuggestions

`func (o *ServerConfiguration) HasEnableExternalContentInSuggestions() bool`

HasEnableExternalContentInSuggestions returns a boolean if a field has been set.

### GetImageExtractionTimeoutMs

`func (o *ServerConfiguration) GetImageExtractionTimeoutMs() int32`

GetImageExtractionTimeoutMs returns the ImageExtractionTimeoutMs field if non-nil, zero value otherwise.

### GetImageExtractionTimeoutMsOk

`func (o *ServerConfiguration) GetImageExtractionTimeoutMsOk() (*int32, bool)`

GetImageExtractionTimeoutMsOk returns a tuple with the ImageExtractionTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageExtractionTimeoutMs

`func (o *ServerConfiguration) SetImageExtractionTimeoutMs(v int32)`

SetImageExtractionTimeoutMs sets ImageExtractionTimeoutMs field to given value.

### HasImageExtractionTimeoutMs

`func (o *ServerConfiguration) HasImageExtractionTimeoutMs() bool`

HasImageExtractionTimeoutMs returns a boolean if a field has been set.

### GetPathSubstitutions

`func (o *ServerConfiguration) GetPathSubstitutions() []PathSubstitution`

GetPathSubstitutions returns the PathSubstitutions field if non-nil, zero value otherwise.

### GetPathSubstitutionsOk

`func (o *ServerConfiguration) GetPathSubstitutionsOk() (*[]PathSubstitution, bool)`

GetPathSubstitutionsOk returns a tuple with the PathSubstitutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathSubstitutions

`func (o *ServerConfiguration) SetPathSubstitutions(v []PathSubstitution)`

SetPathSubstitutions sets PathSubstitutions field to given value.

### HasPathSubstitutions

`func (o *ServerConfiguration) HasPathSubstitutions() bool`

HasPathSubstitutions returns a boolean if a field has been set.

### GetEnableSlowResponseWarning

`func (o *ServerConfiguration) GetEnableSlowResponseWarning() bool`

GetEnableSlowResponseWarning returns the EnableSlowResponseWarning field if non-nil, zero value otherwise.

### GetEnableSlowResponseWarningOk

`func (o *ServerConfiguration) GetEnableSlowResponseWarningOk() (*bool, bool)`

GetEnableSlowResponseWarningOk returns a tuple with the EnableSlowResponseWarning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSlowResponseWarning

`func (o *ServerConfiguration) SetEnableSlowResponseWarning(v bool)`

SetEnableSlowResponseWarning sets EnableSlowResponseWarning field to given value.

### HasEnableSlowResponseWarning

`func (o *ServerConfiguration) HasEnableSlowResponseWarning() bool`

HasEnableSlowResponseWarning returns a boolean if a field has been set.

### GetSlowResponseThresholdMs

`func (o *ServerConfiguration) GetSlowResponseThresholdMs() int64`

GetSlowResponseThresholdMs returns the SlowResponseThresholdMs field if non-nil, zero value otherwise.

### GetSlowResponseThresholdMsOk

`func (o *ServerConfiguration) GetSlowResponseThresholdMsOk() (*int64, bool)`

GetSlowResponseThresholdMsOk returns a tuple with the SlowResponseThresholdMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowResponseThresholdMs

`func (o *ServerConfiguration) SetSlowResponseThresholdMs(v int64)`

SetSlowResponseThresholdMs sets SlowResponseThresholdMs field to given value.

### HasSlowResponseThresholdMs

`func (o *ServerConfiguration) HasSlowResponseThresholdMs() bool`

HasSlowResponseThresholdMs returns a boolean if a field has been set.

### GetCorsHosts

`func (o *ServerConfiguration) GetCorsHosts() []string`

GetCorsHosts returns the CorsHosts field if non-nil, zero value otherwise.

### GetCorsHostsOk

`func (o *ServerConfiguration) GetCorsHostsOk() (*[]string, bool)`

GetCorsHostsOk returns a tuple with the CorsHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsHosts

`func (o *ServerConfiguration) SetCorsHosts(v []string)`

SetCorsHosts sets CorsHosts field to given value.

### HasCorsHosts

`func (o *ServerConfiguration) HasCorsHosts() bool`

HasCorsHosts returns a boolean if a field has been set.

### GetActivityLogRetentionDays

`func (o *ServerConfiguration) GetActivityLogRetentionDays() int32`

GetActivityLogRetentionDays returns the ActivityLogRetentionDays field if non-nil, zero value otherwise.

### GetActivityLogRetentionDaysOk

`func (o *ServerConfiguration) GetActivityLogRetentionDaysOk() (*int32, bool)`

GetActivityLogRetentionDaysOk returns a tuple with the ActivityLogRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityLogRetentionDays

`func (o *ServerConfiguration) SetActivityLogRetentionDays(v int32)`

SetActivityLogRetentionDays sets ActivityLogRetentionDays field to given value.

### HasActivityLogRetentionDays

`func (o *ServerConfiguration) HasActivityLogRetentionDays() bool`

HasActivityLogRetentionDays returns a boolean if a field has been set.

### SetActivityLogRetentionDaysNil

`func (o *ServerConfiguration) SetActivityLogRetentionDaysNil(b bool)`

 SetActivityLogRetentionDaysNil sets the value for ActivityLogRetentionDays to be an explicit nil

### UnsetActivityLogRetentionDays
`func (o *ServerConfiguration) UnsetActivityLogRetentionDays()`

UnsetActivityLogRetentionDays ensures that no value is present for ActivityLogRetentionDays, not even an explicit nil
### GetLibraryScanFanoutConcurrency

`func (o *ServerConfiguration) GetLibraryScanFanoutConcurrency() int32`

GetLibraryScanFanoutConcurrency returns the LibraryScanFanoutConcurrency field if non-nil, zero value otherwise.

### GetLibraryScanFanoutConcurrencyOk

`func (o *ServerConfiguration) GetLibraryScanFanoutConcurrencyOk() (*int32, bool)`

GetLibraryScanFanoutConcurrencyOk returns a tuple with the LibraryScanFanoutConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryScanFanoutConcurrency

`func (o *ServerConfiguration) SetLibraryScanFanoutConcurrency(v int32)`

SetLibraryScanFanoutConcurrency sets LibraryScanFanoutConcurrency field to given value.

### HasLibraryScanFanoutConcurrency

`func (o *ServerConfiguration) HasLibraryScanFanoutConcurrency() bool`

HasLibraryScanFanoutConcurrency returns a boolean if a field has been set.

### GetLibraryMetadataRefreshConcurrency

`func (o *ServerConfiguration) GetLibraryMetadataRefreshConcurrency() int32`

GetLibraryMetadataRefreshConcurrency returns the LibraryMetadataRefreshConcurrency field if non-nil, zero value otherwise.

### GetLibraryMetadataRefreshConcurrencyOk

`func (o *ServerConfiguration) GetLibraryMetadataRefreshConcurrencyOk() (*int32, bool)`

GetLibraryMetadataRefreshConcurrencyOk returns a tuple with the LibraryMetadataRefreshConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryMetadataRefreshConcurrency

`func (o *ServerConfiguration) SetLibraryMetadataRefreshConcurrency(v int32)`

SetLibraryMetadataRefreshConcurrency sets LibraryMetadataRefreshConcurrency field to given value.

### HasLibraryMetadataRefreshConcurrency

`func (o *ServerConfiguration) HasLibraryMetadataRefreshConcurrency() bool`

HasLibraryMetadataRefreshConcurrency returns a boolean if a field has been set.

### GetRemoveOldPlugins

`func (o *ServerConfiguration) GetRemoveOldPlugins() bool`

GetRemoveOldPlugins returns the RemoveOldPlugins field if non-nil, zero value otherwise.

### GetRemoveOldPluginsOk

`func (o *ServerConfiguration) GetRemoveOldPluginsOk() (*bool, bool)`

GetRemoveOldPluginsOk returns a tuple with the RemoveOldPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveOldPlugins

`func (o *ServerConfiguration) SetRemoveOldPlugins(v bool)`

SetRemoveOldPlugins sets RemoveOldPlugins field to given value.

### HasRemoveOldPlugins

`func (o *ServerConfiguration) HasRemoveOldPlugins() bool`

HasRemoveOldPlugins returns a boolean if a field has been set.

### GetAllowClientLogUpload

`func (o *ServerConfiguration) GetAllowClientLogUpload() bool`

GetAllowClientLogUpload returns the AllowClientLogUpload field if non-nil, zero value otherwise.

### GetAllowClientLogUploadOk

`func (o *ServerConfiguration) GetAllowClientLogUploadOk() (*bool, bool)`

GetAllowClientLogUploadOk returns a tuple with the AllowClientLogUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowClientLogUpload

`func (o *ServerConfiguration) SetAllowClientLogUpload(v bool)`

SetAllowClientLogUpload sets AllowClientLogUpload field to given value.

### HasAllowClientLogUpload

`func (o *ServerConfiguration) HasAllowClientLogUpload() bool`

HasAllowClientLogUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


