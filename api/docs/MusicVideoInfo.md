# MusicVideoInfo

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
**Artists** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMusicVideoInfo

`func NewMusicVideoInfo() *MusicVideoInfo`

NewMusicVideoInfo instantiates a new MusicVideoInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMusicVideoInfoWithDefaults

`func NewMusicVideoInfoWithDefaults() *MusicVideoInfo`

NewMusicVideoInfoWithDefaults instantiates a new MusicVideoInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MusicVideoInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MusicVideoInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MusicVideoInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MusicVideoInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MusicVideoInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MusicVideoInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOriginalTitle

`func (o *MusicVideoInfo) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *MusicVideoInfo) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *MusicVideoInfo) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *MusicVideoInfo) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### SetOriginalTitleNil

`func (o *MusicVideoInfo) SetOriginalTitleNil(b bool)`

 SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil

### UnsetOriginalTitle
`func (o *MusicVideoInfo) UnsetOriginalTitle()`

UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
### GetPath

`func (o *MusicVideoInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MusicVideoInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MusicVideoInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MusicVideoInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *MusicVideoInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *MusicVideoInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetMetadataLanguage

`func (o *MusicVideoInfo) GetMetadataLanguage() string`

GetMetadataLanguage returns the MetadataLanguage field if non-nil, zero value otherwise.

### GetMetadataLanguageOk

`func (o *MusicVideoInfo) GetMetadataLanguageOk() (*string, bool)`

GetMetadataLanguageOk returns a tuple with the MetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLanguage

`func (o *MusicVideoInfo) SetMetadataLanguage(v string)`

SetMetadataLanguage sets MetadataLanguage field to given value.

### HasMetadataLanguage

`func (o *MusicVideoInfo) HasMetadataLanguage() bool`

HasMetadataLanguage returns a boolean if a field has been set.

### SetMetadataLanguageNil

`func (o *MusicVideoInfo) SetMetadataLanguageNil(b bool)`

 SetMetadataLanguageNil sets the value for MetadataLanguage to be an explicit nil

### UnsetMetadataLanguage
`func (o *MusicVideoInfo) UnsetMetadataLanguage()`

UnsetMetadataLanguage ensures that no value is present for MetadataLanguage, not even an explicit nil
### GetMetadataCountryCode

`func (o *MusicVideoInfo) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *MusicVideoInfo) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *MusicVideoInfo) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *MusicVideoInfo) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### SetMetadataCountryCodeNil

`func (o *MusicVideoInfo) SetMetadataCountryCodeNil(b bool)`

 SetMetadataCountryCodeNil sets the value for MetadataCountryCode to be an explicit nil

### UnsetMetadataCountryCode
`func (o *MusicVideoInfo) UnsetMetadataCountryCode()`

UnsetMetadataCountryCode ensures that no value is present for MetadataCountryCode, not even an explicit nil
### GetProviderIds

`func (o *MusicVideoInfo) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *MusicVideoInfo) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *MusicVideoInfo) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *MusicVideoInfo) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### SetProviderIdsNil

`func (o *MusicVideoInfo) SetProviderIdsNil(b bool)`

 SetProviderIdsNil sets the value for ProviderIds to be an explicit nil

### UnsetProviderIds
`func (o *MusicVideoInfo) UnsetProviderIds()`

UnsetProviderIds ensures that no value is present for ProviderIds, not even an explicit nil
### GetYear

`func (o *MusicVideoInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *MusicVideoInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *MusicVideoInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *MusicVideoInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *MusicVideoInfo) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *MusicVideoInfo) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetIndexNumber

`func (o *MusicVideoInfo) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *MusicVideoInfo) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *MusicVideoInfo) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *MusicVideoInfo) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *MusicVideoInfo) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *MusicVideoInfo) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetParentIndexNumber

`func (o *MusicVideoInfo) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *MusicVideoInfo) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *MusicVideoInfo) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *MusicVideoInfo) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *MusicVideoInfo) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *MusicVideoInfo) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetPremiereDate

`func (o *MusicVideoInfo) GetPremiereDate() time.Time`

GetPremiereDate returns the PremiereDate field if non-nil, zero value otherwise.

### GetPremiereDateOk

`func (o *MusicVideoInfo) GetPremiereDateOk() (*time.Time, bool)`

GetPremiereDateOk returns a tuple with the PremiereDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiereDate

`func (o *MusicVideoInfo) SetPremiereDate(v time.Time)`

SetPremiereDate sets PremiereDate field to given value.

### HasPremiereDate

`func (o *MusicVideoInfo) HasPremiereDate() bool`

HasPremiereDate returns a boolean if a field has been set.

### SetPremiereDateNil

`func (o *MusicVideoInfo) SetPremiereDateNil(b bool)`

 SetPremiereDateNil sets the value for PremiereDate to be an explicit nil

### UnsetPremiereDate
`func (o *MusicVideoInfo) UnsetPremiereDate()`

UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
### GetIsAutomated

`func (o *MusicVideoInfo) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *MusicVideoInfo) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *MusicVideoInfo) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *MusicVideoInfo) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetArtists

`func (o *MusicVideoInfo) GetArtists() []string`

GetArtists returns the Artists field if non-nil, zero value otherwise.

### GetArtistsOk

`func (o *MusicVideoInfo) GetArtistsOk() (*[]string, bool)`

GetArtistsOk returns a tuple with the Artists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtists

`func (o *MusicVideoInfo) SetArtists(v []string)`

SetArtists sets Artists field to given value.

### HasArtists

`func (o *MusicVideoInfo) HasArtists() bool`

HasArtists returns a boolean if a field has been set.

### SetArtistsNil

`func (o *MusicVideoInfo) SetArtistsNil(b bool)`

 SetArtistsNil sets the value for Artists to be an explicit nil

### UnsetArtists
`func (o *MusicVideoInfo) UnsetArtists()`

UnsetArtists ensures that no value is present for Artists, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


