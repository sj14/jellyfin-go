# AlbumInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**OriginalTitle** | Pointer to **NullableString** | Gets or sets the original title. | [optional] 
**Path** | Pointer to **NullableString** | Gets or sets the path. | [optional] 
**MetadataLanguage** | Pointer to **NullableString** | Gets or sets the metadata language. | [optional] 
**MetadataCountryCode** | Pointer to **NullableString** | Gets or sets the metadata country code. | [optional] 
**ProviderIds** | Pointer to **map[string]string** | Gets or sets the provider ids. | [optional] 
**Year** | Pointer to **NullableInt32** | Gets or sets the year. | [optional] 
**IndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**ParentIndexNumber** | Pointer to **NullableInt32** |  | [optional] 
**PremiereDate** | Pointer to **NullableTime** |  | [optional] 
**IsAutomated** | Pointer to **bool** |  | [optional] 
**AlbumArtists** | Pointer to **[]string** | Gets or sets the album artist. | [optional] 
**ArtistProviderIds** | Pointer to **map[string]string** | Gets or sets the artist provider ids. | [optional] 
**SongInfos** | Pointer to [**[]SongInfo**](SongInfo.md) |  | [optional] 

## Methods

### NewAlbumInfo

`func NewAlbumInfo() *AlbumInfo`

NewAlbumInfo instantiates a new AlbumInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlbumInfoWithDefaults

`func NewAlbumInfoWithDefaults() *AlbumInfo`

NewAlbumInfoWithDefaults instantiates a new AlbumInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlbumInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlbumInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlbumInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlbumInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AlbumInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AlbumInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOriginalTitle

`func (o *AlbumInfo) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *AlbumInfo) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *AlbumInfo) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *AlbumInfo) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### SetOriginalTitleNil

`func (o *AlbumInfo) SetOriginalTitleNil(b bool)`

 SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil

### UnsetOriginalTitle
`func (o *AlbumInfo) UnsetOriginalTitle()`

UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
### GetPath

`func (o *AlbumInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AlbumInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AlbumInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AlbumInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *AlbumInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *AlbumInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetMetadataLanguage

`func (o *AlbumInfo) GetMetadataLanguage() string`

GetMetadataLanguage returns the MetadataLanguage field if non-nil, zero value otherwise.

### GetMetadataLanguageOk

`func (o *AlbumInfo) GetMetadataLanguageOk() (*string, bool)`

GetMetadataLanguageOk returns a tuple with the MetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLanguage

`func (o *AlbumInfo) SetMetadataLanguage(v string)`

SetMetadataLanguage sets MetadataLanguage field to given value.

### HasMetadataLanguage

`func (o *AlbumInfo) HasMetadataLanguage() bool`

HasMetadataLanguage returns a boolean if a field has been set.

### SetMetadataLanguageNil

`func (o *AlbumInfo) SetMetadataLanguageNil(b bool)`

 SetMetadataLanguageNil sets the value for MetadataLanguage to be an explicit nil

### UnsetMetadataLanguage
`func (o *AlbumInfo) UnsetMetadataLanguage()`

UnsetMetadataLanguage ensures that no value is present for MetadataLanguage, not even an explicit nil
### GetMetadataCountryCode

`func (o *AlbumInfo) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *AlbumInfo) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *AlbumInfo) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *AlbumInfo) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### SetMetadataCountryCodeNil

`func (o *AlbumInfo) SetMetadataCountryCodeNil(b bool)`

 SetMetadataCountryCodeNil sets the value for MetadataCountryCode to be an explicit nil

### UnsetMetadataCountryCode
`func (o *AlbumInfo) UnsetMetadataCountryCode()`

UnsetMetadataCountryCode ensures that no value is present for MetadataCountryCode, not even an explicit nil
### GetProviderIds

`func (o *AlbumInfo) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *AlbumInfo) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *AlbumInfo) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *AlbumInfo) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### SetProviderIdsNil

`func (o *AlbumInfo) SetProviderIdsNil(b bool)`

 SetProviderIdsNil sets the value for ProviderIds to be an explicit nil

### UnsetProviderIds
`func (o *AlbumInfo) UnsetProviderIds()`

UnsetProviderIds ensures that no value is present for ProviderIds, not even an explicit nil
### GetYear

`func (o *AlbumInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *AlbumInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *AlbumInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *AlbumInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *AlbumInfo) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *AlbumInfo) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetIndexNumber

`func (o *AlbumInfo) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *AlbumInfo) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *AlbumInfo) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *AlbumInfo) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *AlbumInfo) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *AlbumInfo) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetParentIndexNumber

`func (o *AlbumInfo) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *AlbumInfo) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *AlbumInfo) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *AlbumInfo) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *AlbumInfo) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *AlbumInfo) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetPremiereDate

`func (o *AlbumInfo) GetPremiereDate() time.Time`

GetPremiereDate returns the PremiereDate field if non-nil, zero value otherwise.

### GetPremiereDateOk

`func (o *AlbumInfo) GetPremiereDateOk() (*time.Time, bool)`

GetPremiereDateOk returns a tuple with the PremiereDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiereDate

`func (o *AlbumInfo) SetPremiereDate(v time.Time)`

SetPremiereDate sets PremiereDate field to given value.

### HasPremiereDate

`func (o *AlbumInfo) HasPremiereDate() bool`

HasPremiereDate returns a boolean if a field has been set.

### SetPremiereDateNil

`func (o *AlbumInfo) SetPremiereDateNil(b bool)`

 SetPremiereDateNil sets the value for PremiereDate to be an explicit nil

### UnsetPremiereDate
`func (o *AlbumInfo) UnsetPremiereDate()`

UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
### GetIsAutomated

`func (o *AlbumInfo) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *AlbumInfo) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *AlbumInfo) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *AlbumInfo) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetAlbumArtists

`func (o *AlbumInfo) GetAlbumArtists() []string`

GetAlbumArtists returns the AlbumArtists field if non-nil, zero value otherwise.

### GetAlbumArtistsOk

`func (o *AlbumInfo) GetAlbumArtistsOk() (*[]string, bool)`

GetAlbumArtistsOk returns a tuple with the AlbumArtists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtists

`func (o *AlbumInfo) SetAlbumArtists(v []string)`

SetAlbumArtists sets AlbumArtists field to given value.

### HasAlbumArtists

`func (o *AlbumInfo) HasAlbumArtists() bool`

HasAlbumArtists returns a boolean if a field has been set.

### GetArtistProviderIds

`func (o *AlbumInfo) GetArtistProviderIds() map[string]string`

GetArtistProviderIds returns the ArtistProviderIds field if non-nil, zero value otherwise.

### GetArtistProviderIdsOk

`func (o *AlbumInfo) GetArtistProviderIdsOk() (*map[string]string, bool)`

GetArtistProviderIdsOk returns a tuple with the ArtistProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistProviderIds

`func (o *AlbumInfo) SetArtistProviderIds(v map[string]string)`

SetArtistProviderIds sets ArtistProviderIds field to given value.

### HasArtistProviderIds

`func (o *AlbumInfo) HasArtistProviderIds() bool`

HasArtistProviderIds returns a boolean if a field has been set.

### GetSongInfos

`func (o *AlbumInfo) GetSongInfos() []SongInfo`

GetSongInfos returns the SongInfos field if non-nil, zero value otherwise.

### GetSongInfosOk

`func (o *AlbumInfo) GetSongInfosOk() (*[]SongInfo, bool)`

GetSongInfosOk returns a tuple with the SongInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongInfos

`func (o *AlbumInfo) SetSongInfos(v []SongInfo)`

SetSongInfos sets SongInfos field to given value.

### HasSongInfos

`func (o *AlbumInfo) HasSongInfos() bool`

HasSongInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


