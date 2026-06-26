# QueryFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Genres** | Pointer to [**[]NameGuidPair**](NameGuidPair.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**AudioLanguages** | Pointer to [**[]NameValuePair**](NameValuePair.md) |  | [optional] 
**SubtitleLanguages** | Pointer to [**[]NameValuePair**](NameValuePair.md) |  | [optional] 

## Methods

### NewQueryFilters

`func NewQueryFilters() *QueryFilters`

NewQueryFilters instantiates a new QueryFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryFiltersWithDefaults

`func NewQueryFiltersWithDefaults() *QueryFilters`

NewQueryFiltersWithDefaults instantiates a new QueryFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenres

`func (o *QueryFilters) GetGenres() []NameGuidPair`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *QueryFilters) GetGenresOk() (*[]NameGuidPair, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *QueryFilters) SetGenres(v []NameGuidPair)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *QueryFilters) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### SetGenresNil

`func (o *QueryFilters) SetGenresNil(b bool)`

 SetGenresNil sets the value for Genres to be an explicit nil

### UnsetGenres
`func (o *QueryFilters) UnsetGenres()`

UnsetGenres ensures that no value is present for Genres, not even an explicit nil
### GetTags

`func (o *QueryFilters) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *QueryFilters) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *QueryFilters) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *QueryFilters) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *QueryFilters) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *QueryFilters) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetAudioLanguages

`func (o *QueryFilters) GetAudioLanguages() []NameValuePair`

GetAudioLanguages returns the AudioLanguages field if non-nil, zero value otherwise.

### GetAudioLanguagesOk

`func (o *QueryFilters) GetAudioLanguagesOk() (*[]NameValuePair, bool)`

GetAudioLanguagesOk returns a tuple with the AudioLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioLanguages

`func (o *QueryFilters) SetAudioLanguages(v []NameValuePair)`

SetAudioLanguages sets AudioLanguages field to given value.

### HasAudioLanguages

`func (o *QueryFilters) HasAudioLanguages() bool`

HasAudioLanguages returns a boolean if a field has been set.

### SetAudioLanguagesNil

`func (o *QueryFilters) SetAudioLanguagesNil(b bool)`

 SetAudioLanguagesNil sets the value for AudioLanguages to be an explicit nil

### UnsetAudioLanguages
`func (o *QueryFilters) UnsetAudioLanguages()`

UnsetAudioLanguages ensures that no value is present for AudioLanguages, not even an explicit nil
### GetSubtitleLanguages

`func (o *QueryFilters) GetSubtitleLanguages() []NameValuePair`

GetSubtitleLanguages returns the SubtitleLanguages field if non-nil, zero value otherwise.

### GetSubtitleLanguagesOk

`func (o *QueryFilters) GetSubtitleLanguagesOk() (*[]NameValuePair, bool)`

GetSubtitleLanguagesOk returns a tuple with the SubtitleLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleLanguages

`func (o *QueryFilters) SetSubtitleLanguages(v []NameValuePair)`

SetSubtitleLanguages sets SubtitleLanguages field to given value.

### HasSubtitleLanguages

`func (o *QueryFilters) HasSubtitleLanguages() bool`

HasSubtitleLanguages returns a boolean if a field has been set.

### SetSubtitleLanguagesNil

`func (o *QueryFilters) SetSubtitleLanguagesNil(b bool)`

 SetSubtitleLanguagesNil sets the value for SubtitleLanguages to be an explicit nil

### UnsetSubtitleLanguages
`func (o *QueryFilters) UnsetSubtitleLanguages()`

UnsetSubtitleLanguages ensures that no value is present for SubtitleLanguages, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


