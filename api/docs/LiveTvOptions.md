# LiveTvOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuideDays** | Pointer to **NullableInt32** |  | [optional] 
**RecordingPath** | Pointer to **NullableString** |  | [optional] 
**MovieRecordingPath** | Pointer to **NullableString** |  | [optional] 
**SeriesRecordingPath** | Pointer to **NullableString** |  | [optional] 
**EnableRecordingSubfolders** | Pointer to **bool** |  | [optional] 
**EnableOriginalAudioWithEncodedRecordings** | Pointer to **bool** |  | [optional] 
**TunerHosts** | Pointer to [**[]TunerHostInfo**](TunerHostInfo.md) |  | [optional] 
**ListingProviders** | Pointer to [**[]ListingsProviderInfo**](ListingsProviderInfo.md) |  | [optional] 
**PrePaddingSeconds** | Pointer to **int32** |  | [optional] 
**PostPaddingSeconds** | Pointer to **int32** |  | [optional] 
**MediaLocationsCreated** | Pointer to **[]string** |  | [optional] 
**RecordingPostProcessor** | Pointer to **NullableString** |  | [optional] 
**RecordingPostProcessorArguments** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLiveTvOptions

`func NewLiveTvOptions() *LiveTvOptions`

NewLiveTvOptions instantiates a new LiveTvOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveTvOptionsWithDefaults

`func NewLiveTvOptionsWithDefaults() *LiveTvOptions`

NewLiveTvOptionsWithDefaults instantiates a new LiveTvOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuideDays

`func (o *LiveTvOptions) GetGuideDays() int32`

GetGuideDays returns the GuideDays field if non-nil, zero value otherwise.

### GetGuideDaysOk

`func (o *LiveTvOptions) GetGuideDaysOk() (*int32, bool)`

GetGuideDaysOk returns a tuple with the GuideDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuideDays

`func (o *LiveTvOptions) SetGuideDays(v int32)`

SetGuideDays sets GuideDays field to given value.

### HasGuideDays

`func (o *LiveTvOptions) HasGuideDays() bool`

HasGuideDays returns a boolean if a field has been set.

### SetGuideDaysNil

`func (o *LiveTvOptions) SetGuideDaysNil(b bool)`

 SetGuideDaysNil sets the value for GuideDays to be an explicit nil

### UnsetGuideDays
`func (o *LiveTvOptions) UnsetGuideDays()`

UnsetGuideDays ensures that no value is present for GuideDays, not even an explicit nil
### GetRecordingPath

`func (o *LiveTvOptions) GetRecordingPath() string`

GetRecordingPath returns the RecordingPath field if non-nil, zero value otherwise.

### GetRecordingPathOk

`func (o *LiveTvOptions) GetRecordingPathOk() (*string, bool)`

GetRecordingPathOk returns a tuple with the RecordingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingPath

`func (o *LiveTvOptions) SetRecordingPath(v string)`

SetRecordingPath sets RecordingPath field to given value.

### HasRecordingPath

`func (o *LiveTvOptions) HasRecordingPath() bool`

HasRecordingPath returns a boolean if a field has been set.

### SetRecordingPathNil

`func (o *LiveTvOptions) SetRecordingPathNil(b bool)`

 SetRecordingPathNil sets the value for RecordingPath to be an explicit nil

### UnsetRecordingPath
`func (o *LiveTvOptions) UnsetRecordingPath()`

UnsetRecordingPath ensures that no value is present for RecordingPath, not even an explicit nil
### GetMovieRecordingPath

`func (o *LiveTvOptions) GetMovieRecordingPath() string`

GetMovieRecordingPath returns the MovieRecordingPath field if non-nil, zero value otherwise.

### GetMovieRecordingPathOk

`func (o *LiveTvOptions) GetMovieRecordingPathOk() (*string, bool)`

GetMovieRecordingPathOk returns a tuple with the MovieRecordingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieRecordingPath

`func (o *LiveTvOptions) SetMovieRecordingPath(v string)`

SetMovieRecordingPath sets MovieRecordingPath field to given value.

### HasMovieRecordingPath

`func (o *LiveTvOptions) HasMovieRecordingPath() bool`

HasMovieRecordingPath returns a boolean if a field has been set.

### SetMovieRecordingPathNil

`func (o *LiveTvOptions) SetMovieRecordingPathNil(b bool)`

 SetMovieRecordingPathNil sets the value for MovieRecordingPath to be an explicit nil

### UnsetMovieRecordingPath
`func (o *LiveTvOptions) UnsetMovieRecordingPath()`

UnsetMovieRecordingPath ensures that no value is present for MovieRecordingPath, not even an explicit nil
### GetSeriesRecordingPath

`func (o *LiveTvOptions) GetSeriesRecordingPath() string`

GetSeriesRecordingPath returns the SeriesRecordingPath field if non-nil, zero value otherwise.

### GetSeriesRecordingPathOk

`func (o *LiveTvOptions) GetSeriesRecordingPathOk() (*string, bool)`

GetSeriesRecordingPathOk returns a tuple with the SeriesRecordingPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesRecordingPath

`func (o *LiveTvOptions) SetSeriesRecordingPath(v string)`

SetSeriesRecordingPath sets SeriesRecordingPath field to given value.

### HasSeriesRecordingPath

`func (o *LiveTvOptions) HasSeriesRecordingPath() bool`

HasSeriesRecordingPath returns a boolean if a field has been set.

### SetSeriesRecordingPathNil

`func (o *LiveTvOptions) SetSeriesRecordingPathNil(b bool)`

 SetSeriesRecordingPathNil sets the value for SeriesRecordingPath to be an explicit nil

### UnsetSeriesRecordingPath
`func (o *LiveTvOptions) UnsetSeriesRecordingPath()`

UnsetSeriesRecordingPath ensures that no value is present for SeriesRecordingPath, not even an explicit nil
### GetEnableRecordingSubfolders

`func (o *LiveTvOptions) GetEnableRecordingSubfolders() bool`

GetEnableRecordingSubfolders returns the EnableRecordingSubfolders field if non-nil, zero value otherwise.

### GetEnableRecordingSubfoldersOk

`func (o *LiveTvOptions) GetEnableRecordingSubfoldersOk() (*bool, bool)`

GetEnableRecordingSubfoldersOk returns a tuple with the EnableRecordingSubfolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecordingSubfolders

`func (o *LiveTvOptions) SetEnableRecordingSubfolders(v bool)`

SetEnableRecordingSubfolders sets EnableRecordingSubfolders field to given value.

### HasEnableRecordingSubfolders

`func (o *LiveTvOptions) HasEnableRecordingSubfolders() bool`

HasEnableRecordingSubfolders returns a boolean if a field has been set.

### GetEnableOriginalAudioWithEncodedRecordings

`func (o *LiveTvOptions) GetEnableOriginalAudioWithEncodedRecordings() bool`

GetEnableOriginalAudioWithEncodedRecordings returns the EnableOriginalAudioWithEncodedRecordings field if non-nil, zero value otherwise.

### GetEnableOriginalAudioWithEncodedRecordingsOk

`func (o *LiveTvOptions) GetEnableOriginalAudioWithEncodedRecordingsOk() (*bool, bool)`

GetEnableOriginalAudioWithEncodedRecordingsOk returns a tuple with the EnableOriginalAudioWithEncodedRecordings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOriginalAudioWithEncodedRecordings

`func (o *LiveTvOptions) SetEnableOriginalAudioWithEncodedRecordings(v bool)`

SetEnableOriginalAudioWithEncodedRecordings sets EnableOriginalAudioWithEncodedRecordings field to given value.

### HasEnableOriginalAudioWithEncodedRecordings

`func (o *LiveTvOptions) HasEnableOriginalAudioWithEncodedRecordings() bool`

HasEnableOriginalAudioWithEncodedRecordings returns a boolean if a field has been set.

### GetTunerHosts

`func (o *LiveTvOptions) GetTunerHosts() []TunerHostInfo`

GetTunerHosts returns the TunerHosts field if non-nil, zero value otherwise.

### GetTunerHostsOk

`func (o *LiveTvOptions) GetTunerHostsOk() (*[]TunerHostInfo, bool)`

GetTunerHostsOk returns a tuple with the TunerHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunerHosts

`func (o *LiveTvOptions) SetTunerHosts(v []TunerHostInfo)`

SetTunerHosts sets TunerHosts field to given value.

### HasTunerHosts

`func (o *LiveTvOptions) HasTunerHosts() bool`

HasTunerHosts returns a boolean if a field has been set.

### SetTunerHostsNil

`func (o *LiveTvOptions) SetTunerHostsNil(b bool)`

 SetTunerHostsNil sets the value for TunerHosts to be an explicit nil

### UnsetTunerHosts
`func (o *LiveTvOptions) UnsetTunerHosts()`

UnsetTunerHosts ensures that no value is present for TunerHosts, not even an explicit nil
### GetListingProviders

`func (o *LiveTvOptions) GetListingProviders() []ListingsProviderInfo`

GetListingProviders returns the ListingProviders field if non-nil, zero value otherwise.

### GetListingProvidersOk

`func (o *LiveTvOptions) GetListingProvidersOk() (*[]ListingsProviderInfo, bool)`

GetListingProvidersOk returns a tuple with the ListingProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingProviders

`func (o *LiveTvOptions) SetListingProviders(v []ListingsProviderInfo)`

SetListingProviders sets ListingProviders field to given value.

### HasListingProviders

`func (o *LiveTvOptions) HasListingProviders() bool`

HasListingProviders returns a boolean if a field has been set.

### SetListingProvidersNil

`func (o *LiveTvOptions) SetListingProvidersNil(b bool)`

 SetListingProvidersNil sets the value for ListingProviders to be an explicit nil

### UnsetListingProviders
`func (o *LiveTvOptions) UnsetListingProviders()`

UnsetListingProviders ensures that no value is present for ListingProviders, not even an explicit nil
### GetPrePaddingSeconds

`func (o *LiveTvOptions) GetPrePaddingSeconds() int32`

GetPrePaddingSeconds returns the PrePaddingSeconds field if non-nil, zero value otherwise.

### GetPrePaddingSecondsOk

`func (o *LiveTvOptions) GetPrePaddingSecondsOk() (*int32, bool)`

GetPrePaddingSecondsOk returns a tuple with the PrePaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaddingSeconds

`func (o *LiveTvOptions) SetPrePaddingSeconds(v int32)`

SetPrePaddingSeconds sets PrePaddingSeconds field to given value.

### HasPrePaddingSeconds

`func (o *LiveTvOptions) HasPrePaddingSeconds() bool`

HasPrePaddingSeconds returns a boolean if a field has been set.

### GetPostPaddingSeconds

`func (o *LiveTvOptions) GetPostPaddingSeconds() int32`

GetPostPaddingSeconds returns the PostPaddingSeconds field if non-nil, zero value otherwise.

### GetPostPaddingSecondsOk

`func (o *LiveTvOptions) GetPostPaddingSecondsOk() (*int32, bool)`

GetPostPaddingSecondsOk returns a tuple with the PostPaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPaddingSeconds

`func (o *LiveTvOptions) SetPostPaddingSeconds(v int32)`

SetPostPaddingSeconds sets PostPaddingSeconds field to given value.

### HasPostPaddingSeconds

`func (o *LiveTvOptions) HasPostPaddingSeconds() bool`

HasPostPaddingSeconds returns a boolean if a field has been set.

### GetMediaLocationsCreated

`func (o *LiveTvOptions) GetMediaLocationsCreated() []string`

GetMediaLocationsCreated returns the MediaLocationsCreated field if non-nil, zero value otherwise.

### GetMediaLocationsCreatedOk

`func (o *LiveTvOptions) GetMediaLocationsCreatedOk() (*[]string, bool)`

GetMediaLocationsCreatedOk returns a tuple with the MediaLocationsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaLocationsCreated

`func (o *LiveTvOptions) SetMediaLocationsCreated(v []string)`

SetMediaLocationsCreated sets MediaLocationsCreated field to given value.

### HasMediaLocationsCreated

`func (o *LiveTvOptions) HasMediaLocationsCreated() bool`

HasMediaLocationsCreated returns a boolean if a field has been set.

### SetMediaLocationsCreatedNil

`func (o *LiveTvOptions) SetMediaLocationsCreatedNil(b bool)`

 SetMediaLocationsCreatedNil sets the value for MediaLocationsCreated to be an explicit nil

### UnsetMediaLocationsCreated
`func (o *LiveTvOptions) UnsetMediaLocationsCreated()`

UnsetMediaLocationsCreated ensures that no value is present for MediaLocationsCreated, not even an explicit nil
### GetRecordingPostProcessor

`func (o *LiveTvOptions) GetRecordingPostProcessor() string`

GetRecordingPostProcessor returns the RecordingPostProcessor field if non-nil, zero value otherwise.

### GetRecordingPostProcessorOk

`func (o *LiveTvOptions) GetRecordingPostProcessorOk() (*string, bool)`

GetRecordingPostProcessorOk returns a tuple with the RecordingPostProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingPostProcessor

`func (o *LiveTvOptions) SetRecordingPostProcessor(v string)`

SetRecordingPostProcessor sets RecordingPostProcessor field to given value.

### HasRecordingPostProcessor

`func (o *LiveTvOptions) HasRecordingPostProcessor() bool`

HasRecordingPostProcessor returns a boolean if a field has been set.

### SetRecordingPostProcessorNil

`func (o *LiveTvOptions) SetRecordingPostProcessorNil(b bool)`

 SetRecordingPostProcessorNil sets the value for RecordingPostProcessor to be an explicit nil

### UnsetRecordingPostProcessor
`func (o *LiveTvOptions) UnsetRecordingPostProcessor()`

UnsetRecordingPostProcessor ensures that no value is present for RecordingPostProcessor, not even an explicit nil
### GetRecordingPostProcessorArguments

`func (o *LiveTvOptions) GetRecordingPostProcessorArguments() string`

GetRecordingPostProcessorArguments returns the RecordingPostProcessorArguments field if non-nil, zero value otherwise.

### GetRecordingPostProcessorArgumentsOk

`func (o *LiveTvOptions) GetRecordingPostProcessorArgumentsOk() (*string, bool)`

GetRecordingPostProcessorArgumentsOk returns a tuple with the RecordingPostProcessorArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingPostProcessorArguments

`func (o *LiveTvOptions) SetRecordingPostProcessorArguments(v string)`

SetRecordingPostProcessorArguments sets RecordingPostProcessorArguments field to given value.

### HasRecordingPostProcessorArguments

`func (o *LiveTvOptions) HasRecordingPostProcessorArguments() bool`

HasRecordingPostProcessorArguments returns a boolean if a field has been set.

### SetRecordingPostProcessorArgumentsNil

`func (o *LiveTvOptions) SetRecordingPostProcessorArgumentsNil(b bool)`

 SetRecordingPostProcessorArgumentsNil sets the value for RecordingPostProcessorArguments to be an explicit nil

### UnsetRecordingPostProcessorArguments
`func (o *LiveTvOptions) UnsetRecordingPostProcessorArguments()`

UnsetRecordingPostProcessorArguments ensures that no value is present for RecordingPostProcessorArguments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


