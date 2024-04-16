# GetProgramsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelIds** | Pointer to **[]string** | Gets or sets the channels to return guide information for. | [optional] 
**UserId** | Pointer to **string** | Gets or sets optional. Filter by user id. | [optional] 
**MinStartDate** | Pointer to **NullableTime** | Gets or sets the minimum premiere start date.  Optional. | [optional] 
**HasAired** | Pointer to **NullableBool** | Gets or sets filter by programs that have completed airing, or not.  Optional. | [optional] 
**IsAiring** | Pointer to **NullableBool** | Gets or sets filter by programs that are currently airing, or not.  Optional. | [optional] 
**MaxStartDate** | Pointer to **NullableTime** | Gets or sets the maximum premiere start date.  Optional. | [optional] 
**MinEndDate** | Pointer to **NullableTime** | Gets or sets the minimum premiere end date.  Optional. | [optional] 
**MaxEndDate** | Pointer to **NullableTime** | Gets or sets the maximum premiere end date.  Optional. | [optional] 
**IsMovie** | Pointer to **NullableBool** | Gets or sets filter for movies.  Optional. | [optional] 
**IsSeries** | Pointer to **NullableBool** | Gets or sets filter for series.  Optional. | [optional] 
**IsNews** | Pointer to **NullableBool** | Gets or sets filter for news.  Optional. | [optional] 
**IsKids** | Pointer to **NullableBool** | Gets or sets filter for kids.  Optional. | [optional] 
**IsSports** | Pointer to **NullableBool** | Gets or sets filter for sports.  Optional. | [optional] 
**StartIndex** | Pointer to **NullableInt32** | Gets or sets the record index to start at. All items with a lower index will be dropped from the results.  Optional. | [optional] 
**Limit** | Pointer to **NullableInt32** | Gets or sets the maximum number of records to return.  Optional. | [optional] 
**SortBy** | Pointer to **[]string** | Gets or sets specify one or more sort orders, comma delimited. Options: Name, StartDate.  Optional. | [optional] 
**SortOrder** | Pointer to [**[]SortOrder**](SortOrder.md) | Gets or sets sort Order - Ascending,Descending. | [optional] 
**Genres** | Pointer to **[]string** | Gets or sets the genres to return guide information for. | [optional] 
**GenreIds** | Pointer to **[]string** | Gets or sets the genre ids to return guide information for. | [optional] 
**EnableImages** | Pointer to **NullableBool** | Gets or sets include image information in output.  Optional. | [optional] 
**EnableTotalRecordCount** | Pointer to **bool** | Gets or sets a value indicating whether retrieve total record count. | [optional] 
**ImageTypeLimit** | Pointer to **NullableInt32** | Gets or sets the max number of images to return, per image type.  Optional. | [optional] 
**EnableImageTypes** | Pointer to [**[]ImageType**](ImageType.md) | Gets or sets the image types to include in the output.  Optional. | [optional] 
**EnableUserData** | Pointer to **NullableBool** | Gets or sets include user data.  Optional. | [optional] 
**SeriesTimerId** | Pointer to **NullableString** | Gets or sets filter by series timer id.  Optional. | [optional] 
**LibrarySeriesId** | Pointer to **string** | Gets or sets filter by library series id.  Optional. | [optional] 
**Fields** | Pointer to [**[]ItemFields**](ItemFields.md) | Gets or sets specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines.  Optional. | [optional] 

## Methods

### NewGetProgramsDto

`func NewGetProgramsDto() *GetProgramsDto`

NewGetProgramsDto instantiates a new GetProgramsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProgramsDtoWithDefaults

`func NewGetProgramsDtoWithDefaults() *GetProgramsDto`

NewGetProgramsDtoWithDefaults instantiates a new GetProgramsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelIds

`func (o *GetProgramsDto) GetChannelIds() []string`

GetChannelIds returns the ChannelIds field if non-nil, zero value otherwise.

### GetChannelIdsOk

`func (o *GetProgramsDto) GetChannelIdsOk() (*[]string, bool)`

GetChannelIdsOk returns a tuple with the ChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIds

`func (o *GetProgramsDto) SetChannelIds(v []string)`

SetChannelIds sets ChannelIds field to given value.

### HasChannelIds

`func (o *GetProgramsDto) HasChannelIds() bool`

HasChannelIds returns a boolean if a field has been set.

### GetUserId

`func (o *GetProgramsDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetProgramsDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetProgramsDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetProgramsDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetMinStartDate

`func (o *GetProgramsDto) GetMinStartDate() time.Time`

GetMinStartDate returns the MinStartDate field if non-nil, zero value otherwise.

### GetMinStartDateOk

`func (o *GetProgramsDto) GetMinStartDateOk() (*time.Time, bool)`

GetMinStartDateOk returns a tuple with the MinStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStartDate

`func (o *GetProgramsDto) SetMinStartDate(v time.Time)`

SetMinStartDate sets MinStartDate field to given value.

### HasMinStartDate

`func (o *GetProgramsDto) HasMinStartDate() bool`

HasMinStartDate returns a boolean if a field has been set.

### SetMinStartDateNil

`func (o *GetProgramsDto) SetMinStartDateNil(b bool)`

 SetMinStartDateNil sets the value for MinStartDate to be an explicit nil

### UnsetMinStartDate
`func (o *GetProgramsDto) UnsetMinStartDate()`

UnsetMinStartDate ensures that no value is present for MinStartDate, not even an explicit nil
### GetHasAired

`func (o *GetProgramsDto) GetHasAired() bool`

GetHasAired returns the HasAired field if non-nil, zero value otherwise.

### GetHasAiredOk

`func (o *GetProgramsDto) GetHasAiredOk() (*bool, bool)`

GetHasAiredOk returns a tuple with the HasAired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAired

`func (o *GetProgramsDto) SetHasAired(v bool)`

SetHasAired sets HasAired field to given value.

### HasHasAired

`func (o *GetProgramsDto) HasHasAired() bool`

HasHasAired returns a boolean if a field has been set.

### SetHasAiredNil

`func (o *GetProgramsDto) SetHasAiredNil(b bool)`

 SetHasAiredNil sets the value for HasAired to be an explicit nil

### UnsetHasAired
`func (o *GetProgramsDto) UnsetHasAired()`

UnsetHasAired ensures that no value is present for HasAired, not even an explicit nil
### GetIsAiring

`func (o *GetProgramsDto) GetIsAiring() bool`

GetIsAiring returns the IsAiring field if non-nil, zero value otherwise.

### GetIsAiringOk

`func (o *GetProgramsDto) GetIsAiringOk() (*bool, bool)`

GetIsAiringOk returns a tuple with the IsAiring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAiring

`func (o *GetProgramsDto) SetIsAiring(v bool)`

SetIsAiring sets IsAiring field to given value.

### HasIsAiring

`func (o *GetProgramsDto) HasIsAiring() bool`

HasIsAiring returns a boolean if a field has been set.

### SetIsAiringNil

`func (o *GetProgramsDto) SetIsAiringNil(b bool)`

 SetIsAiringNil sets the value for IsAiring to be an explicit nil

### UnsetIsAiring
`func (o *GetProgramsDto) UnsetIsAiring()`

UnsetIsAiring ensures that no value is present for IsAiring, not even an explicit nil
### GetMaxStartDate

`func (o *GetProgramsDto) GetMaxStartDate() time.Time`

GetMaxStartDate returns the MaxStartDate field if non-nil, zero value otherwise.

### GetMaxStartDateOk

`func (o *GetProgramsDto) GetMaxStartDateOk() (*time.Time, bool)`

GetMaxStartDateOk returns a tuple with the MaxStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStartDate

`func (o *GetProgramsDto) SetMaxStartDate(v time.Time)`

SetMaxStartDate sets MaxStartDate field to given value.

### HasMaxStartDate

`func (o *GetProgramsDto) HasMaxStartDate() bool`

HasMaxStartDate returns a boolean if a field has been set.

### SetMaxStartDateNil

`func (o *GetProgramsDto) SetMaxStartDateNil(b bool)`

 SetMaxStartDateNil sets the value for MaxStartDate to be an explicit nil

### UnsetMaxStartDate
`func (o *GetProgramsDto) UnsetMaxStartDate()`

UnsetMaxStartDate ensures that no value is present for MaxStartDate, not even an explicit nil
### GetMinEndDate

`func (o *GetProgramsDto) GetMinEndDate() time.Time`

GetMinEndDate returns the MinEndDate field if non-nil, zero value otherwise.

### GetMinEndDateOk

`func (o *GetProgramsDto) GetMinEndDateOk() (*time.Time, bool)`

GetMinEndDateOk returns a tuple with the MinEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEndDate

`func (o *GetProgramsDto) SetMinEndDate(v time.Time)`

SetMinEndDate sets MinEndDate field to given value.

### HasMinEndDate

`func (o *GetProgramsDto) HasMinEndDate() bool`

HasMinEndDate returns a boolean if a field has been set.

### SetMinEndDateNil

`func (o *GetProgramsDto) SetMinEndDateNil(b bool)`

 SetMinEndDateNil sets the value for MinEndDate to be an explicit nil

### UnsetMinEndDate
`func (o *GetProgramsDto) UnsetMinEndDate()`

UnsetMinEndDate ensures that no value is present for MinEndDate, not even an explicit nil
### GetMaxEndDate

`func (o *GetProgramsDto) GetMaxEndDate() time.Time`

GetMaxEndDate returns the MaxEndDate field if non-nil, zero value otherwise.

### GetMaxEndDateOk

`func (o *GetProgramsDto) GetMaxEndDateOk() (*time.Time, bool)`

GetMaxEndDateOk returns a tuple with the MaxEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEndDate

`func (o *GetProgramsDto) SetMaxEndDate(v time.Time)`

SetMaxEndDate sets MaxEndDate field to given value.

### HasMaxEndDate

`func (o *GetProgramsDto) HasMaxEndDate() bool`

HasMaxEndDate returns a boolean if a field has been set.

### SetMaxEndDateNil

`func (o *GetProgramsDto) SetMaxEndDateNil(b bool)`

 SetMaxEndDateNil sets the value for MaxEndDate to be an explicit nil

### UnsetMaxEndDate
`func (o *GetProgramsDto) UnsetMaxEndDate()`

UnsetMaxEndDate ensures that no value is present for MaxEndDate, not even an explicit nil
### GetIsMovie

`func (o *GetProgramsDto) GetIsMovie() bool`

GetIsMovie returns the IsMovie field if non-nil, zero value otherwise.

### GetIsMovieOk

`func (o *GetProgramsDto) GetIsMovieOk() (*bool, bool)`

GetIsMovieOk returns a tuple with the IsMovie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMovie

`func (o *GetProgramsDto) SetIsMovie(v bool)`

SetIsMovie sets IsMovie field to given value.

### HasIsMovie

`func (o *GetProgramsDto) HasIsMovie() bool`

HasIsMovie returns a boolean if a field has been set.

### SetIsMovieNil

`func (o *GetProgramsDto) SetIsMovieNil(b bool)`

 SetIsMovieNil sets the value for IsMovie to be an explicit nil

### UnsetIsMovie
`func (o *GetProgramsDto) UnsetIsMovie()`

UnsetIsMovie ensures that no value is present for IsMovie, not even an explicit nil
### GetIsSeries

`func (o *GetProgramsDto) GetIsSeries() bool`

GetIsSeries returns the IsSeries field if non-nil, zero value otherwise.

### GetIsSeriesOk

`func (o *GetProgramsDto) GetIsSeriesOk() (*bool, bool)`

GetIsSeriesOk returns a tuple with the IsSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeries

`func (o *GetProgramsDto) SetIsSeries(v bool)`

SetIsSeries sets IsSeries field to given value.

### HasIsSeries

`func (o *GetProgramsDto) HasIsSeries() bool`

HasIsSeries returns a boolean if a field has been set.

### SetIsSeriesNil

`func (o *GetProgramsDto) SetIsSeriesNil(b bool)`

 SetIsSeriesNil sets the value for IsSeries to be an explicit nil

### UnsetIsSeries
`func (o *GetProgramsDto) UnsetIsSeries()`

UnsetIsSeries ensures that no value is present for IsSeries, not even an explicit nil
### GetIsNews

`func (o *GetProgramsDto) GetIsNews() bool`

GetIsNews returns the IsNews field if non-nil, zero value otherwise.

### GetIsNewsOk

`func (o *GetProgramsDto) GetIsNewsOk() (*bool, bool)`

GetIsNewsOk returns a tuple with the IsNews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNews

`func (o *GetProgramsDto) SetIsNews(v bool)`

SetIsNews sets IsNews field to given value.

### HasIsNews

`func (o *GetProgramsDto) HasIsNews() bool`

HasIsNews returns a boolean if a field has been set.

### SetIsNewsNil

`func (o *GetProgramsDto) SetIsNewsNil(b bool)`

 SetIsNewsNil sets the value for IsNews to be an explicit nil

### UnsetIsNews
`func (o *GetProgramsDto) UnsetIsNews()`

UnsetIsNews ensures that no value is present for IsNews, not even an explicit nil
### GetIsKids

`func (o *GetProgramsDto) GetIsKids() bool`

GetIsKids returns the IsKids field if non-nil, zero value otherwise.

### GetIsKidsOk

`func (o *GetProgramsDto) GetIsKidsOk() (*bool, bool)`

GetIsKidsOk returns a tuple with the IsKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKids

`func (o *GetProgramsDto) SetIsKids(v bool)`

SetIsKids sets IsKids field to given value.

### HasIsKids

`func (o *GetProgramsDto) HasIsKids() bool`

HasIsKids returns a boolean if a field has been set.

### SetIsKidsNil

`func (o *GetProgramsDto) SetIsKidsNil(b bool)`

 SetIsKidsNil sets the value for IsKids to be an explicit nil

### UnsetIsKids
`func (o *GetProgramsDto) UnsetIsKids()`

UnsetIsKids ensures that no value is present for IsKids, not even an explicit nil
### GetIsSports

`func (o *GetProgramsDto) GetIsSports() bool`

GetIsSports returns the IsSports field if non-nil, zero value otherwise.

### GetIsSportsOk

`func (o *GetProgramsDto) GetIsSportsOk() (*bool, bool)`

GetIsSportsOk returns a tuple with the IsSports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSports

`func (o *GetProgramsDto) SetIsSports(v bool)`

SetIsSports sets IsSports field to given value.

### HasIsSports

`func (o *GetProgramsDto) HasIsSports() bool`

HasIsSports returns a boolean if a field has been set.

### SetIsSportsNil

`func (o *GetProgramsDto) SetIsSportsNil(b bool)`

 SetIsSportsNil sets the value for IsSports to be an explicit nil

### UnsetIsSports
`func (o *GetProgramsDto) UnsetIsSports()`

UnsetIsSports ensures that no value is present for IsSports, not even an explicit nil
### GetStartIndex

`func (o *GetProgramsDto) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *GetProgramsDto) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *GetProgramsDto) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *GetProgramsDto) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.

### SetStartIndexNil

`func (o *GetProgramsDto) SetStartIndexNil(b bool)`

 SetStartIndexNil sets the value for StartIndex to be an explicit nil

### UnsetStartIndex
`func (o *GetProgramsDto) UnsetStartIndex()`

UnsetStartIndex ensures that no value is present for StartIndex, not even an explicit nil
### GetLimit

`func (o *GetProgramsDto) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetProgramsDto) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetProgramsDto) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetProgramsDto) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *GetProgramsDto) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *GetProgramsDto) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetSortBy

`func (o *GetProgramsDto) GetSortBy() []string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *GetProgramsDto) GetSortByOk() (*[]string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *GetProgramsDto) SetSortBy(v []string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *GetProgramsDto) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetSortOrder

`func (o *GetProgramsDto) GetSortOrder() []SortOrder`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *GetProgramsDto) GetSortOrderOk() (*[]SortOrder, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *GetProgramsDto) SetSortOrder(v []SortOrder)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *GetProgramsDto) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetGenres

`func (o *GetProgramsDto) GetGenres() []string`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *GetProgramsDto) GetGenresOk() (*[]string, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *GetProgramsDto) SetGenres(v []string)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *GetProgramsDto) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetGenreIds

`func (o *GetProgramsDto) GetGenreIds() []string`

GetGenreIds returns the GenreIds field if non-nil, zero value otherwise.

### GetGenreIdsOk

`func (o *GetProgramsDto) GetGenreIdsOk() (*[]string, bool)`

GetGenreIdsOk returns a tuple with the GenreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenreIds

`func (o *GetProgramsDto) SetGenreIds(v []string)`

SetGenreIds sets GenreIds field to given value.

### HasGenreIds

`func (o *GetProgramsDto) HasGenreIds() bool`

HasGenreIds returns a boolean if a field has been set.

### GetEnableImages

`func (o *GetProgramsDto) GetEnableImages() bool`

GetEnableImages returns the EnableImages field if non-nil, zero value otherwise.

### GetEnableImagesOk

`func (o *GetProgramsDto) GetEnableImagesOk() (*bool, bool)`

GetEnableImagesOk returns a tuple with the EnableImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImages

`func (o *GetProgramsDto) SetEnableImages(v bool)`

SetEnableImages sets EnableImages field to given value.

### HasEnableImages

`func (o *GetProgramsDto) HasEnableImages() bool`

HasEnableImages returns a boolean if a field has been set.

### SetEnableImagesNil

`func (o *GetProgramsDto) SetEnableImagesNil(b bool)`

 SetEnableImagesNil sets the value for EnableImages to be an explicit nil

### UnsetEnableImages
`func (o *GetProgramsDto) UnsetEnableImages()`

UnsetEnableImages ensures that no value is present for EnableImages, not even an explicit nil
### GetEnableTotalRecordCount

`func (o *GetProgramsDto) GetEnableTotalRecordCount() bool`

GetEnableTotalRecordCount returns the EnableTotalRecordCount field if non-nil, zero value otherwise.

### GetEnableTotalRecordCountOk

`func (o *GetProgramsDto) GetEnableTotalRecordCountOk() (*bool, bool)`

GetEnableTotalRecordCountOk returns a tuple with the EnableTotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTotalRecordCount

`func (o *GetProgramsDto) SetEnableTotalRecordCount(v bool)`

SetEnableTotalRecordCount sets EnableTotalRecordCount field to given value.

### HasEnableTotalRecordCount

`func (o *GetProgramsDto) HasEnableTotalRecordCount() bool`

HasEnableTotalRecordCount returns a boolean if a field has been set.

### GetImageTypeLimit

`func (o *GetProgramsDto) GetImageTypeLimit() int32`

GetImageTypeLimit returns the ImageTypeLimit field if non-nil, zero value otherwise.

### GetImageTypeLimitOk

`func (o *GetProgramsDto) GetImageTypeLimitOk() (*int32, bool)`

GetImageTypeLimitOk returns a tuple with the ImageTypeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTypeLimit

`func (o *GetProgramsDto) SetImageTypeLimit(v int32)`

SetImageTypeLimit sets ImageTypeLimit field to given value.

### HasImageTypeLimit

`func (o *GetProgramsDto) HasImageTypeLimit() bool`

HasImageTypeLimit returns a boolean if a field has been set.

### SetImageTypeLimitNil

`func (o *GetProgramsDto) SetImageTypeLimitNil(b bool)`

 SetImageTypeLimitNil sets the value for ImageTypeLimit to be an explicit nil

### UnsetImageTypeLimit
`func (o *GetProgramsDto) UnsetImageTypeLimit()`

UnsetImageTypeLimit ensures that no value is present for ImageTypeLimit, not even an explicit nil
### GetEnableImageTypes

`func (o *GetProgramsDto) GetEnableImageTypes() []ImageType`

GetEnableImageTypes returns the EnableImageTypes field if non-nil, zero value otherwise.

### GetEnableImageTypesOk

`func (o *GetProgramsDto) GetEnableImageTypesOk() (*[]ImageType, bool)`

GetEnableImageTypesOk returns a tuple with the EnableImageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImageTypes

`func (o *GetProgramsDto) SetEnableImageTypes(v []ImageType)`

SetEnableImageTypes sets EnableImageTypes field to given value.

### HasEnableImageTypes

`func (o *GetProgramsDto) HasEnableImageTypes() bool`

HasEnableImageTypes returns a boolean if a field has been set.

### GetEnableUserData

`func (o *GetProgramsDto) GetEnableUserData() bool`

GetEnableUserData returns the EnableUserData field if non-nil, zero value otherwise.

### GetEnableUserDataOk

`func (o *GetProgramsDto) GetEnableUserDataOk() (*bool, bool)`

GetEnableUserDataOk returns a tuple with the EnableUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserData

`func (o *GetProgramsDto) SetEnableUserData(v bool)`

SetEnableUserData sets EnableUserData field to given value.

### HasEnableUserData

`func (o *GetProgramsDto) HasEnableUserData() bool`

HasEnableUserData returns a boolean if a field has been set.

### SetEnableUserDataNil

`func (o *GetProgramsDto) SetEnableUserDataNil(b bool)`

 SetEnableUserDataNil sets the value for EnableUserData to be an explicit nil

### UnsetEnableUserData
`func (o *GetProgramsDto) UnsetEnableUserData()`

UnsetEnableUserData ensures that no value is present for EnableUserData, not even an explicit nil
### GetSeriesTimerId

`func (o *GetProgramsDto) GetSeriesTimerId() string`

GetSeriesTimerId returns the SeriesTimerId field if non-nil, zero value otherwise.

### GetSeriesTimerIdOk

`func (o *GetProgramsDto) GetSeriesTimerIdOk() (*string, bool)`

GetSeriesTimerIdOk returns a tuple with the SeriesTimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTimerId

`func (o *GetProgramsDto) SetSeriesTimerId(v string)`

SetSeriesTimerId sets SeriesTimerId field to given value.

### HasSeriesTimerId

`func (o *GetProgramsDto) HasSeriesTimerId() bool`

HasSeriesTimerId returns a boolean if a field has been set.

### SetSeriesTimerIdNil

`func (o *GetProgramsDto) SetSeriesTimerIdNil(b bool)`

 SetSeriesTimerIdNil sets the value for SeriesTimerId to be an explicit nil

### UnsetSeriesTimerId
`func (o *GetProgramsDto) UnsetSeriesTimerId()`

UnsetSeriesTimerId ensures that no value is present for SeriesTimerId, not even an explicit nil
### GetLibrarySeriesId

`func (o *GetProgramsDto) GetLibrarySeriesId() string`

GetLibrarySeriesId returns the LibrarySeriesId field if non-nil, zero value otherwise.

### GetLibrarySeriesIdOk

`func (o *GetProgramsDto) GetLibrarySeriesIdOk() (*string, bool)`

GetLibrarySeriesIdOk returns a tuple with the LibrarySeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrarySeriesId

`func (o *GetProgramsDto) SetLibrarySeriesId(v string)`

SetLibrarySeriesId sets LibrarySeriesId field to given value.

### HasLibrarySeriesId

`func (o *GetProgramsDto) HasLibrarySeriesId() bool`

HasLibrarySeriesId returns a boolean if a field has been set.

### GetFields

`func (o *GetProgramsDto) GetFields() []ItemFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GetProgramsDto) GetFieldsOk() (*[]ItemFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GetProgramsDto) SetFields(v []ItemFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *GetProgramsDto) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


