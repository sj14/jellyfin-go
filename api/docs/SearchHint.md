# SearchHint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **string** | Gets or sets the item id. | [optional] 
**Id** | Pointer to **string** | Gets or sets the item id. | [optional] 
**Name** | Pointer to **string** | Gets or sets the name. | [optional] 
**MatchedTerm** | Pointer to **NullableString** | Gets or sets the matched term. | [optional] 
**IndexNumber** | Pointer to **NullableInt32** | Gets or sets the index number. | [optional] 
**ProductionYear** | Pointer to **NullableInt32** | Gets or sets the production year. | [optional] 
**ParentIndexNumber** | Pointer to **NullableInt32** | Gets or sets the parent index number. | [optional] 
**PrimaryImageTag** | Pointer to **NullableString** | Gets or sets the image tag. | [optional] 
**ThumbImageTag** | Pointer to **NullableString** | Gets or sets the thumb image tag. | [optional] 
**ThumbImageItemId** | Pointer to **NullableString** | Gets or sets the thumb image item identifier. | [optional] 
**BackdropImageTag** | Pointer to **NullableString** | Gets or sets the backdrop image tag. | [optional] 
**BackdropImageItemId** | Pointer to **NullableString** | Gets or sets the backdrop image item identifier. | [optional] 
**Type** | Pointer to [**BaseItemKind**](BaseItemKind.md) | Gets or sets the type. | [optional] 
**IsFolder** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is folder. | [optional] 
**RunTimeTicks** | Pointer to **NullableInt64** | Gets or sets the run time ticks. | [optional] 
**MediaType** | Pointer to [**MediaType**](MediaType.md) | Gets or sets the type of the media. | [optional] [default to MEDIATYPE_UNKNOWN]
**StartDate** | Pointer to **NullableTime** | Gets or sets the start date. | [optional] 
**EndDate** | Pointer to **NullableTime** | Gets or sets the end date. | [optional] 
**Series** | Pointer to **NullableString** | Gets or sets the series. | [optional] 
**Status** | Pointer to **NullableString** | Gets or sets the status. | [optional] 
**Album** | Pointer to **NullableString** | Gets or sets the album. | [optional] 
**AlbumId** | Pointer to **NullableString** | Gets or sets the album id. | [optional] 
**AlbumArtist** | Pointer to **NullableString** | Gets or sets the album artist. | [optional] 
**Artists** | Pointer to **[]string** | Gets or sets the artists. | [optional] 
**SongCount** | Pointer to **NullableInt32** | Gets or sets the song count. | [optional] 
**EpisodeCount** | Pointer to **NullableInt32** | Gets or sets the episode count. | [optional] 
**ChannelId** | Pointer to **NullableString** | Gets or sets the channel identifier. | [optional] 
**ChannelName** | Pointer to **NullableString** | Gets or sets the name of the channel. | [optional] 
**PrimaryImageAspectRatio** | Pointer to **NullableFloat64** | Gets or sets the primary image aspect ratio. | [optional] 

## Methods

### NewSearchHint

`func NewSearchHint() *SearchHint`

NewSearchHint instantiates a new SearchHint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchHintWithDefaults

`func NewSearchHintWithDefaults() *SearchHint`

NewSearchHintWithDefaults instantiates a new SearchHint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *SearchHint) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *SearchHint) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *SearchHint) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *SearchHint) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetId

`func (o *SearchHint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchHint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchHint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchHint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SearchHint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchHint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchHint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SearchHint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMatchedTerm

`func (o *SearchHint) GetMatchedTerm() string`

GetMatchedTerm returns the MatchedTerm field if non-nil, zero value otherwise.

### GetMatchedTermOk

`func (o *SearchHint) GetMatchedTermOk() (*string, bool)`

GetMatchedTermOk returns a tuple with the MatchedTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedTerm

`func (o *SearchHint) SetMatchedTerm(v string)`

SetMatchedTerm sets MatchedTerm field to given value.

### HasMatchedTerm

`func (o *SearchHint) HasMatchedTerm() bool`

HasMatchedTerm returns a boolean if a field has been set.

### SetMatchedTermNil

`func (o *SearchHint) SetMatchedTermNil(b bool)`

 SetMatchedTermNil sets the value for MatchedTerm to be an explicit nil

### UnsetMatchedTerm
`func (o *SearchHint) UnsetMatchedTerm()`

UnsetMatchedTerm ensures that no value is present for MatchedTerm, not even an explicit nil
### GetIndexNumber

`func (o *SearchHint) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *SearchHint) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *SearchHint) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *SearchHint) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *SearchHint) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *SearchHint) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetProductionYear

`func (o *SearchHint) GetProductionYear() int32`

GetProductionYear returns the ProductionYear field if non-nil, zero value otherwise.

### GetProductionYearOk

`func (o *SearchHint) GetProductionYearOk() (*int32, bool)`

GetProductionYearOk returns a tuple with the ProductionYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionYear

`func (o *SearchHint) SetProductionYear(v int32)`

SetProductionYear sets ProductionYear field to given value.

### HasProductionYear

`func (o *SearchHint) HasProductionYear() bool`

HasProductionYear returns a boolean if a field has been set.

### SetProductionYearNil

`func (o *SearchHint) SetProductionYearNil(b bool)`

 SetProductionYearNil sets the value for ProductionYear to be an explicit nil

### UnsetProductionYear
`func (o *SearchHint) UnsetProductionYear()`

UnsetProductionYear ensures that no value is present for ProductionYear, not even an explicit nil
### GetParentIndexNumber

`func (o *SearchHint) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *SearchHint) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *SearchHint) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *SearchHint) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *SearchHint) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *SearchHint) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetPrimaryImageTag

`func (o *SearchHint) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *SearchHint) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *SearchHint) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *SearchHint) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.

### SetPrimaryImageTagNil

`func (o *SearchHint) SetPrimaryImageTagNil(b bool)`

 SetPrimaryImageTagNil sets the value for PrimaryImageTag to be an explicit nil

### UnsetPrimaryImageTag
`func (o *SearchHint) UnsetPrimaryImageTag()`

UnsetPrimaryImageTag ensures that no value is present for PrimaryImageTag, not even an explicit nil
### GetThumbImageTag

`func (o *SearchHint) GetThumbImageTag() string`

GetThumbImageTag returns the ThumbImageTag field if non-nil, zero value otherwise.

### GetThumbImageTagOk

`func (o *SearchHint) GetThumbImageTagOk() (*string, bool)`

GetThumbImageTagOk returns a tuple with the ThumbImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbImageTag

`func (o *SearchHint) SetThumbImageTag(v string)`

SetThumbImageTag sets ThumbImageTag field to given value.

### HasThumbImageTag

`func (o *SearchHint) HasThumbImageTag() bool`

HasThumbImageTag returns a boolean if a field has been set.

### SetThumbImageTagNil

`func (o *SearchHint) SetThumbImageTagNil(b bool)`

 SetThumbImageTagNil sets the value for ThumbImageTag to be an explicit nil

### UnsetThumbImageTag
`func (o *SearchHint) UnsetThumbImageTag()`

UnsetThumbImageTag ensures that no value is present for ThumbImageTag, not even an explicit nil
### GetThumbImageItemId

`func (o *SearchHint) GetThumbImageItemId() string`

GetThumbImageItemId returns the ThumbImageItemId field if non-nil, zero value otherwise.

### GetThumbImageItemIdOk

`func (o *SearchHint) GetThumbImageItemIdOk() (*string, bool)`

GetThumbImageItemIdOk returns a tuple with the ThumbImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbImageItemId

`func (o *SearchHint) SetThumbImageItemId(v string)`

SetThumbImageItemId sets ThumbImageItemId field to given value.

### HasThumbImageItemId

`func (o *SearchHint) HasThumbImageItemId() bool`

HasThumbImageItemId returns a boolean if a field has been set.

### SetThumbImageItemIdNil

`func (o *SearchHint) SetThumbImageItemIdNil(b bool)`

 SetThumbImageItemIdNil sets the value for ThumbImageItemId to be an explicit nil

### UnsetThumbImageItemId
`func (o *SearchHint) UnsetThumbImageItemId()`

UnsetThumbImageItemId ensures that no value is present for ThumbImageItemId, not even an explicit nil
### GetBackdropImageTag

`func (o *SearchHint) GetBackdropImageTag() string`

GetBackdropImageTag returns the BackdropImageTag field if non-nil, zero value otherwise.

### GetBackdropImageTagOk

`func (o *SearchHint) GetBackdropImageTagOk() (*string, bool)`

GetBackdropImageTagOk returns a tuple with the BackdropImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdropImageTag

`func (o *SearchHint) SetBackdropImageTag(v string)`

SetBackdropImageTag sets BackdropImageTag field to given value.

### HasBackdropImageTag

`func (o *SearchHint) HasBackdropImageTag() bool`

HasBackdropImageTag returns a boolean if a field has been set.

### SetBackdropImageTagNil

`func (o *SearchHint) SetBackdropImageTagNil(b bool)`

 SetBackdropImageTagNil sets the value for BackdropImageTag to be an explicit nil

### UnsetBackdropImageTag
`func (o *SearchHint) UnsetBackdropImageTag()`

UnsetBackdropImageTag ensures that no value is present for BackdropImageTag, not even an explicit nil
### GetBackdropImageItemId

`func (o *SearchHint) GetBackdropImageItemId() string`

GetBackdropImageItemId returns the BackdropImageItemId field if non-nil, zero value otherwise.

### GetBackdropImageItemIdOk

`func (o *SearchHint) GetBackdropImageItemIdOk() (*string, bool)`

GetBackdropImageItemIdOk returns a tuple with the BackdropImageItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdropImageItemId

`func (o *SearchHint) SetBackdropImageItemId(v string)`

SetBackdropImageItemId sets BackdropImageItemId field to given value.

### HasBackdropImageItemId

`func (o *SearchHint) HasBackdropImageItemId() bool`

HasBackdropImageItemId returns a boolean if a field has been set.

### SetBackdropImageItemIdNil

`func (o *SearchHint) SetBackdropImageItemIdNil(b bool)`

 SetBackdropImageItemIdNil sets the value for BackdropImageItemId to be an explicit nil

### UnsetBackdropImageItemId
`func (o *SearchHint) UnsetBackdropImageItemId()`

UnsetBackdropImageItemId ensures that no value is present for BackdropImageItemId, not even an explicit nil
### GetType

`func (o *SearchHint) GetType() BaseItemKind`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchHint) GetTypeOk() (*BaseItemKind, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchHint) SetType(v BaseItemKind)`

SetType sets Type field to given value.

### HasType

`func (o *SearchHint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsFolder

`func (o *SearchHint) GetIsFolder() bool`

GetIsFolder returns the IsFolder field if non-nil, zero value otherwise.

### GetIsFolderOk

`func (o *SearchHint) GetIsFolderOk() (*bool, bool)`

GetIsFolderOk returns a tuple with the IsFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFolder

`func (o *SearchHint) SetIsFolder(v bool)`

SetIsFolder sets IsFolder field to given value.

### HasIsFolder

`func (o *SearchHint) HasIsFolder() bool`

HasIsFolder returns a boolean if a field has been set.

### SetIsFolderNil

`func (o *SearchHint) SetIsFolderNil(b bool)`

 SetIsFolderNil sets the value for IsFolder to be an explicit nil

### UnsetIsFolder
`func (o *SearchHint) UnsetIsFolder()`

UnsetIsFolder ensures that no value is present for IsFolder, not even an explicit nil
### GetRunTimeTicks

`func (o *SearchHint) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *SearchHint) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *SearchHint) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *SearchHint) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *SearchHint) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *SearchHint) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetMediaType

`func (o *SearchHint) GetMediaType() MediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *SearchHint) GetMediaTypeOk() (*MediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *SearchHint) SetMediaType(v MediaType)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *SearchHint) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetStartDate

`func (o *SearchHint) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SearchHint) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SearchHint) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SearchHint) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *SearchHint) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *SearchHint) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *SearchHint) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SearchHint) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SearchHint) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SearchHint) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *SearchHint) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *SearchHint) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetSeries

`func (o *SearchHint) GetSeries() string`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *SearchHint) GetSeriesOk() (*string, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *SearchHint) SetSeries(v string)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *SearchHint) HasSeries() bool`

HasSeries returns a boolean if a field has been set.

### SetSeriesNil

`func (o *SearchHint) SetSeriesNil(b bool)`

 SetSeriesNil sets the value for Series to be an explicit nil

### UnsetSeries
`func (o *SearchHint) UnsetSeries()`

UnsetSeries ensures that no value is present for Series, not even an explicit nil
### GetStatus

`func (o *SearchHint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchHint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchHint) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchHint) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *SearchHint) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *SearchHint) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAlbum

`func (o *SearchHint) GetAlbum() string`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *SearchHint) GetAlbumOk() (*string, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *SearchHint) SetAlbum(v string)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *SearchHint) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### SetAlbumNil

`func (o *SearchHint) SetAlbumNil(b bool)`

 SetAlbumNil sets the value for Album to be an explicit nil

### UnsetAlbum
`func (o *SearchHint) UnsetAlbum()`

UnsetAlbum ensures that no value is present for Album, not even an explicit nil
### GetAlbumId

`func (o *SearchHint) GetAlbumId() string`

GetAlbumId returns the AlbumId field if non-nil, zero value otherwise.

### GetAlbumIdOk

`func (o *SearchHint) GetAlbumIdOk() (*string, bool)`

GetAlbumIdOk returns a tuple with the AlbumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumId

`func (o *SearchHint) SetAlbumId(v string)`

SetAlbumId sets AlbumId field to given value.

### HasAlbumId

`func (o *SearchHint) HasAlbumId() bool`

HasAlbumId returns a boolean if a field has been set.

### SetAlbumIdNil

`func (o *SearchHint) SetAlbumIdNil(b bool)`

 SetAlbumIdNil sets the value for AlbumId to be an explicit nil

### UnsetAlbumId
`func (o *SearchHint) UnsetAlbumId()`

UnsetAlbumId ensures that no value is present for AlbumId, not even an explicit nil
### GetAlbumArtist

`func (o *SearchHint) GetAlbumArtist() string`

GetAlbumArtist returns the AlbumArtist field if non-nil, zero value otherwise.

### GetAlbumArtistOk

`func (o *SearchHint) GetAlbumArtistOk() (*string, bool)`

GetAlbumArtistOk returns a tuple with the AlbumArtist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtist

`func (o *SearchHint) SetAlbumArtist(v string)`

SetAlbumArtist sets AlbumArtist field to given value.

### HasAlbumArtist

`func (o *SearchHint) HasAlbumArtist() bool`

HasAlbumArtist returns a boolean if a field has been set.

### SetAlbumArtistNil

`func (o *SearchHint) SetAlbumArtistNil(b bool)`

 SetAlbumArtistNil sets the value for AlbumArtist to be an explicit nil

### UnsetAlbumArtist
`func (o *SearchHint) UnsetAlbumArtist()`

UnsetAlbumArtist ensures that no value is present for AlbumArtist, not even an explicit nil
### GetArtists

`func (o *SearchHint) GetArtists() []string`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *SearchHint) GetArtistsOk() (*[]string, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *SearchHint) SetArtists(v []string)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *SearchHint) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### GetSongCount

`func (o *SearchHint) GetSongCount() int32`

GetSongCount returns the SongCount field if non-nil, zero value otherwise.

### GetSongCountOk

`func (o *SearchHint) GetSongCountOk() (*int32, bool)`

GetSongCountOk returns a tuple with the SongCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongCount

`func (o *SearchHint) SetSongCount(v int32)`

SetSongCount sets SongCount field to given value.

### HasSongCount

`func (o *SearchHint) HasSongCount() bool`

HasSongCount returns a boolean if a field has been set.

### SetSongCountNil

`func (o *SearchHint) SetSongCountNil(b bool)`

 SetSongCountNil sets the value for SongCount to be an explicit nil

### UnsetSongCount
`func (o *SearchHint) UnsetSongCount()`

UnsetSongCount ensures that no value is present for SongCount, not even an explicit nil
### GetEpisodeCount

`func (o *SearchHint) GetEpisodeCount() int32`

GetEpisodeCount returns the EpisodeCount field if non-nil, zero value otherwise.

### GetEpisodeCountOk

`func (o *SearchHint) GetEpisodeCountOk() (*int32, bool)`

GetEpisodeCountOk returns a tuple with the EpisodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodeCount

`func (o *SearchHint) SetEpisodeCount(v int32)`

SetEpisodeCount sets EpisodeCount field to given value.

### HasEpisodeCount

`func (o *SearchHint) HasEpisodeCount() bool`

HasEpisodeCount returns a boolean if a field has been set.

### SetEpisodeCountNil

`func (o *SearchHint) SetEpisodeCountNil(b bool)`

 SetEpisodeCountNil sets the value for EpisodeCount to be an explicit nil

### UnsetEpisodeCount
`func (o *SearchHint) UnsetEpisodeCount()`

UnsetEpisodeCount ensures that no value is present for EpisodeCount, not even an explicit nil
### GetChannelId

`func (o *SearchHint) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SearchHint) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SearchHint) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *SearchHint) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *SearchHint) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *SearchHint) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetChannelName

`func (o *SearchHint) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *SearchHint) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *SearchHint) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *SearchHint) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### SetChannelNameNil

`func (o *SearchHint) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *SearchHint) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetPrimaryImageAspectRatio

`func (o *SearchHint) GetPrimaryImageAspectRatio() float64`

GetPrimaryImageAspectRatio returns the PrimaryImageAspectRatio field if non-nil, zero value otherwise.

### GetPrimaryImageAspectRatioOk

`func (o *SearchHint) GetPrimaryImageAspectRatioOk() (*float64, bool)`

GetPrimaryImageAspectRatioOk returns a tuple with the PrimaryImageAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageAspectRatio

`func (o *SearchHint) SetPrimaryImageAspectRatio(v float64)`

SetPrimaryImageAspectRatio sets PrimaryImageAspectRatio field to given value.

### HasPrimaryImageAspectRatio

`func (o *SearchHint) HasPrimaryImageAspectRatio() bool`

HasPrimaryImageAspectRatio returns a boolean if a field has been set.

### SetPrimaryImageAspectRatioNil

`func (o *SearchHint) SetPrimaryImageAspectRatioNil(b bool)`

 SetPrimaryImageAspectRatioNil sets the value for PrimaryImageAspectRatio to be an explicit nil

### UnsetPrimaryImageAspectRatio
`func (o *SearchHint) UnsetPrimaryImageAspectRatio()`

UnsetPrimaryImageAspectRatio ensures that no value is present for PrimaryImageAspectRatio, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


