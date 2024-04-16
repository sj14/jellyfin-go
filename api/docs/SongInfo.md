# SongInfo

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
**AlbumArtists** | Pointer to **[]string** |  | [optional] 
**Album** | Pointer to **NullableString** |  | [optional] 
**Artists** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSongInfo

`func NewSongInfo() *SongInfo`

NewSongInfo instantiates a new SongInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSongInfoWithDefaults

`func NewSongInfoWithDefaults() *SongInfo`

NewSongInfoWithDefaults instantiates a new SongInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SongInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SongInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SongInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SongInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SongInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SongInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOriginalTitle

`func (o *SongInfo) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *SongInfo) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *SongInfo) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *SongInfo) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### SetOriginalTitleNil

`func (o *SongInfo) SetOriginalTitleNil(b bool)`

 SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil

### UnsetOriginalTitle
`func (o *SongInfo) UnsetOriginalTitle()`

UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
### GetPath

`func (o *SongInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SongInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SongInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SongInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *SongInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *SongInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetMetadataLanguage

`func (o *SongInfo) GetMetadataLanguage() string`

GetMetadataLanguage returns the MetadataLanguage field if non-nil, zero value otherwise.

### GetMetadataLanguageOk

`func (o *SongInfo) GetMetadataLanguageOk() (*string, bool)`

GetMetadataLanguageOk returns a tuple with the MetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLanguage

`func (o *SongInfo) SetMetadataLanguage(v string)`

SetMetadataLanguage sets MetadataLanguage field to given value.

### HasMetadataLanguage

`func (o *SongInfo) HasMetadataLanguage() bool`

HasMetadataLanguage returns a boolean if a field has been set.

### SetMetadataLanguageNil

`func (o *SongInfo) SetMetadataLanguageNil(b bool)`

 SetMetadataLanguageNil sets the value for MetadataLanguage to be an explicit nil

### UnsetMetadataLanguage
`func (o *SongInfo) UnsetMetadataLanguage()`

UnsetMetadataLanguage ensures that no value is present for MetadataLanguage, not even an explicit nil
### GetMetadataCountryCode

`func (o *SongInfo) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *SongInfo) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *SongInfo) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *SongInfo) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### SetMetadataCountryCodeNil

`func (o *SongInfo) SetMetadataCountryCodeNil(b bool)`

 SetMetadataCountryCodeNil sets the value for MetadataCountryCode to be an explicit nil

### UnsetMetadataCountryCode
`func (o *SongInfo) UnsetMetadataCountryCode()`

UnsetMetadataCountryCode ensures that no value is present for MetadataCountryCode, not even an explicit nil
### GetProviderIds

`func (o *SongInfo) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *SongInfo) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *SongInfo) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *SongInfo) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### SetProviderIdsNil

`func (o *SongInfo) SetProviderIdsNil(b bool)`

 SetProviderIdsNil sets the value for ProviderIds to be an explicit nil

### UnsetProviderIds
`func (o *SongInfo) UnsetProviderIds()`

UnsetProviderIds ensures that no value is present for ProviderIds, not even an explicit nil
### GetYear

`func (o *SongInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *SongInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *SongInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *SongInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *SongInfo) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *SongInfo) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetIndexNumber

`func (o *SongInfo) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *SongInfo) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *SongInfo) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *SongInfo) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *SongInfo) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *SongInfo) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetParentIndexNumber

`func (o *SongInfo) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *SongInfo) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *SongInfo) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *SongInfo) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *SongInfo) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *SongInfo) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetPremiereDate

`func (o *SongInfo) GetPremiereDate() time.Time`

GetPremiereDate returns the PremiereDate field if non-nil, zero value otherwise.

### GetPremiereDateOk

`func (o *SongInfo) GetPremiereDateOk() (*time.Time, bool)`

GetPremiereDateOk returns a tuple with the PremiereDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiereDate

`func (o *SongInfo) SetPremiereDate(v time.Time)`

SetPremiereDate sets PremiereDate field to given value.

### HasPremiereDate

`func (o *SongInfo) HasPremiereDate() bool`

HasPremiereDate returns a boolean if a field has been set.

### SetPremiereDateNil

`func (o *SongInfo) SetPremiereDateNil(b bool)`

 SetPremiereDateNil sets the value for PremiereDate to be an explicit nil

### UnsetPremiereDate
`func (o *SongInfo) UnsetPremiereDate()`

UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
### GetIsAutomated

`func (o *SongInfo) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *SongInfo) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *SongInfo) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *SongInfo) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetAlbumArtists

`func (o *SongInfo) GetAlbumArtists() []string`

GetAlbumArtists returns the AlbumArtists field if non-nil, zero value otherwise.

### GetAlbumArtistsOk

`func (o *SongInfo) GetAlbumArtistsOk() (*[]string, bool)`

GetAlbumArtistsOk returns a tuple with the AlbumArtists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtists

`func (o *SongInfo) SetAlbumArtists(v []string)`

SetAlbumArtists sets AlbumArtists field to given value.

### HasAlbumArtists

`func (o *SongInfo) HasAlbumArtists() bool`

HasAlbumArtists returns a boolean if a field has been set.

### SetAlbumArtistsNil

`func (o *SongInfo) SetAlbumArtistsNil(b bool)`

 SetAlbumArtistsNil sets the value for AlbumArtists to be an explicit nil

### UnsetAlbumArtists
`func (o *SongInfo) UnsetAlbumArtists()`

UnsetAlbumArtists ensures that no value is present for AlbumArtists, not even an explicit nil
### GetAlbum

`func (o *SongInfo) GetAlbum() string`

GetAlbum returns the Album field if non-nil, zero value otherwise.

### GetAlbumOk

`func (o *SongInfo) GetAlbumOk() (*string, bool)`

GetAlbumOk returns a tuple with the Album field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbum

`func (o *SongInfo) SetAlbum(v string)`

SetAlbum sets Album field to given value.

### HasAlbum

`func (o *SongInfo) HasAlbum() bool`

HasAlbum returns a boolean if a field has been set.

### SetAlbumNil

`func (o *SongInfo) SetAlbumNil(b bool)`

 SetAlbumNil sets the value for Album to be an explicit nil

### UnsetAlbum
`func (o *SongInfo) UnsetAlbum()`

UnsetAlbum ensures that no value is present for Album, not even an explicit nil
### GetArtists

`func (o *SongInfo) GetArtists() []string`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *SongInfo) GetArtistsOk() (*[]string, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *SongInfo) SetArtists(v []string)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *SongInfo) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### SetArtistsNil

`func (o *SongInfo) SetArtistsNil(b bool)`

 SetArtistsNil sets the value for Artists to be an explicit nil

### UnsetArtists
`func (o *SongInfo) UnsetArtists()`

UnsetArtists ensures that no value is present for Artists, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


