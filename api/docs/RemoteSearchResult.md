# RemoteSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**ProviderIds** | Pointer to **map[string]string** | Gets or sets the provider ids. | [optional] 
**ProductionYear** | Pointer to **NullableInt32** | Gets or sets the year. | [optional] 
**IndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**IndexNumberEnd** | Pointer to **NullableInt32** |  | [optional] 
**ParentIndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**PremiereDate** | Pointer to **NullableTime** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**SearchProviderName** | Pointer to **NullableString** |  | [optional] 
**Overview** | Pointer to **NullableString** |  | [optional] 
**AlbumArtist** | Pointer to [**NullableRemoteSearchResult**](RemoteSearchResult.md) |  | [optional] 
**Artists** | Pointer to [**[]RemoteSearchResult**](RemoteSearchResult.md) |  | [optional] 

## Methods

### NewRemoteSearchResult

`func NewRemoteSearchResult() *RemoteSearchResult`

NewRemoteSearchResult instantiates a new RemoteSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSearchResultWithDefaults

`func NewRemoteSearchResultWithDefaults() *RemoteSearchResult`

NewRemoteSearchResultWithDefaults instantiates a new RemoteSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RemoteSearchResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteSearchResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteSearchResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemoteSearchResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RemoteSearchResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RemoteSearchResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProviderIds

`func (o *RemoteSearchResult) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *RemoteSearchResult) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *RemoteSearchResult) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *RemoteSearchResult) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### SetProviderIdsNil

`func (o *RemoteSearchResult) SetProviderIdsNil(b bool)`

 SetProviderIdsNil sets the value for ProviderIds to be an explicit nil

### UnsetProviderIds
`func (o *RemoteSearchResult) UnsetProviderIds()`

UnsetProviderIds ensures that no value is present for ProviderIds, not even an explicit nil
### GetProductionYear

`func (o *RemoteSearchResult) GetProductionYear() int32`

GetProductionYear returns the ProductionYear field if non-nil, zero value otherwise.

### GetProductionYearOk

`func (o *RemoteSearchResult) GetProductionYearOk() (*int32, bool)`

GetProductionYearOk returns a tuple with the ProductionYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionYear

`func (o *RemoteSearchResult) SetProductionYear(v int32)`

SetProductionYear sets ProductionYear field to given value.

### HasProductionYear

`func (o *RemoteSearchResult) HasProductionYear() bool`

HasProductionYear returns a boolean if a field has been set.

### SetProductionYearNil

`func (o *RemoteSearchResult) SetProductionYearNil(b bool)`

 SetProductionYearNil sets the value for ProductionYear to be an explicit nil

### UnsetProductionYear
`func (o *RemoteSearchResult) UnsetProductionYear()`

UnsetProductionYear ensures that no value is present for ProductionYear, not even an explicit nil
### GetIndexNumber

`func (o *RemoteSearchResult) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *RemoteSearchResult) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *RemoteSearchResult) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *RemoteSearchResult) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *RemoteSearchResult) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *RemoteSearchResult) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetIndexNumberEnd

`func (o *RemoteSearchResult) GetIndexNumberEnd() int32`

GetIndexNumberEnd returns the IndexNumberEnd field if non-nil, zero value otherwise.

### GetIndexNumberEndOk

`func (o *RemoteSearchResult) GetIndexNumberEndOk() (*int32, bool)`

GetIndexNumberEndOk returns a tuple with the IndexNumberEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumberEnd

`func (o *RemoteSearchResult) SetIndexNumberEnd(v int32)`

SetIndexNumberEnd sets IndexNumberEnd field to given value.

### HasIndexNumberEnd

`func (o *RemoteSearchResult) HasIndexNumberEnd() bool`

HasIndexNumberEnd returns a boolean if a field has been set.

### SetIndexNumberEndNil

`func (o *RemoteSearchResult) SetIndexNumberEndNil(b bool)`

 SetIndexNumberEndNil sets the value for IndexNumberEnd to be an explicit nil

### UnsetIndexNumberEnd
`func (o *RemoteSearchResult) UnsetIndexNumberEnd()`

UnsetIndexNumberEnd ensures that no value is present for IndexNumberEnd, not even an explicit nil
### GetParentIndexNumber

`func (o *RemoteSearchResult) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *RemoteSearchResult) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *RemoteSearchResult) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *RemoteSearchResult) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *RemoteSearchResult) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *RemoteSearchResult) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetPremiereDate

`func (o *RemoteSearchResult) GetPremiereDate() time.Time`

GetPremiereDate returns the PremiereDate field if non-nil, zero value otherwise.

### GetPremiereDateOk

`func (o *RemoteSearchResult) GetPremiereDateOk() (*time.Time, bool)`

GetPremiereDateOk returns a tuple with the PremiereDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiereDate

`func (o *RemoteSearchResult) SetPremiereDate(v time.Time)`

SetPremiereDate sets PremiereDate field to given value.

### HasPremiereDate

`func (o *RemoteSearchResult) HasPremiereDate() bool`

HasPremiereDate returns a boolean if a field has been set.

### SetPremiereDateNil

`func (o *RemoteSearchResult) SetPremiereDateNil(b bool)`

 SetPremiereDateNil sets the value for PremiereDate to be an explicit nil

### UnsetPremiereDate
`func (o *RemoteSearchResult) UnsetPremiereDate()`

UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
### GetImageUrl

`func (o *RemoteSearchResult) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *RemoteSearchResult) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *RemoteSearchResult) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *RemoteSearchResult) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *RemoteSearchResult) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *RemoteSearchResult) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetSearchProviderName

`func (o *RemoteSearchResult) GetSearchProviderName() string`

GetSearchProviderName returns the SearchProviderName field if non-nil, zero value otherwise.

### GetSearchProviderNameOk

`func (o *RemoteSearchResult) GetSearchProviderNameOk() (*string, bool)`

GetSearchProviderNameOk returns a tuple with the SearchProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchProviderName

`func (o *RemoteSearchResult) SetSearchProviderName(v string)`

SetSearchProviderName sets SearchProviderName field to given value.

### HasSearchProviderName

`func (o *RemoteSearchResult) HasSearchProviderName() bool`

HasSearchProviderName returns a boolean if a field has been set.

### SetSearchProviderNameNil

`func (o *RemoteSearchResult) SetSearchProviderNameNil(b bool)`

 SetSearchProviderNameNil sets the value for SearchProviderName to be an explicit nil

### UnsetSearchProviderName
`func (o *RemoteSearchResult) UnsetSearchProviderName()`

UnsetSearchProviderName ensures that no value is present for SearchProviderName, not even an explicit nil
### GetOverview

`func (o *RemoteSearchResult) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *RemoteSearchResult) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *RemoteSearchResult) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *RemoteSearchResult) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *RemoteSearchResult) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *RemoteSearchResult) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetAlbumArtist

`func (o *RemoteSearchResult) GetAlbumArtist() RemoteSearchResult`

GetAlbumArtist returns the AlbumArtist field if non-nil, zero value otherwise.

### GetAlbumArtistOk

`func (o *RemoteSearchResult) GetAlbumArtistOk() (*RemoteSearchResult, bool)`

GetAlbumArtistOk returns a tuple with the AlbumArtist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtist

`func (o *RemoteSearchResult) SetAlbumArtist(v RemoteSearchResult)`

SetAlbumArtist sets AlbumArtist field to given value.

### HasAlbumArtist

`func (o *RemoteSearchResult) HasAlbumArtist() bool`

HasAlbumArtist returns a boolean if a field has been set.

### SetAlbumArtistNil

`func (o *RemoteSearchResult) SetAlbumArtistNil(b bool)`

 SetAlbumArtistNil sets the value for AlbumArtist to be an explicit nil

### UnsetAlbumArtist
`func (o *RemoteSearchResult) UnsetAlbumArtist()`

UnsetAlbumArtist ensures that no value is present for AlbumArtist, not even an explicit nil
### GetArtists

`func (o *RemoteSearchResult) GetArtists() []RemoteSearchResult`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *RemoteSearchResult) GetArtistsOk() (*[]RemoteSearchResult, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *RemoteSearchResult) SetArtists(v []RemoteSearchResult)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *RemoteSearchResult) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### SetArtistsNil

`func (o *RemoteSearchResult) SetArtistsNil(b bool)`

 SetArtistsNil sets the value for Artists to be an explicit nil

### UnsetArtists
`func (o *RemoteSearchResult) UnsetArtists()`

UnsetArtists ensures that no value is present for Artists, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


