# BookInfo

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
**SeriesName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBookInfo

`func NewBookInfo() *BookInfo`

NewBookInfo instantiates a new BookInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBookInfoWithDefaults

`func NewBookInfoWithDefaults() *BookInfo`

NewBookInfoWithDefaults instantiates a new BookInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BookInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BookInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BookInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BookInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BookInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BookInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOriginalTitle

`func (o *BookInfo) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *BookInfo) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *BookInfo) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *BookInfo) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### SetOriginalTitleNil

`func (o *BookInfo) SetOriginalTitleNil(b bool)`

 SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil

### UnsetOriginalTitle
`func (o *BookInfo) UnsetOriginalTitle()`

UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
### GetPath

`func (o *BookInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BookInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BookInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BookInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *BookInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *BookInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetMetadataLanguage

`func (o *BookInfo) GetMetadataLanguage() string`

GetMetadataLanguage returns the MetadataLanguage field if non-nil, zero value otherwise.

### GetMetadataLanguageOk

`func (o *BookInfo) GetMetadataLanguageOk() (*string, bool)`

GetMetadataLanguageOk returns a tuple with the MetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLanguage

`func (o *BookInfo) SetMetadataLanguage(v string)`

SetMetadataLanguage sets MetadataLanguage field to given value.

### HasMetadataLanguage

`func (o *BookInfo) HasMetadataLanguage() bool`

HasMetadataLanguage returns a boolean if a field has been set.

### SetMetadataLanguageNil

`func (o *BookInfo) SetMetadataLanguageNil(b bool)`

 SetMetadataLanguageNil sets the value for MetadataLanguage to be an explicit nil

### UnsetMetadataLanguage
`func (o *BookInfo) UnsetMetadataLanguage()`

UnsetMetadataLanguage ensures that no value is present for MetadataLanguage, not even an explicit nil
### GetMetadataCountryCode

`func (o *BookInfo) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *BookInfo) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *BookInfo) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *BookInfo) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### SetMetadataCountryCodeNil

`func (o *BookInfo) SetMetadataCountryCodeNil(b bool)`

 SetMetadataCountryCodeNil sets the value for MetadataCountryCode to be an explicit nil

### UnsetMetadataCountryCode
`func (o *BookInfo) UnsetMetadataCountryCode()`

UnsetMetadataCountryCode ensures that no value is present for MetadataCountryCode, not even an explicit nil
### GetProviderIds

`func (o *BookInfo) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *BookInfo) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *BookInfo) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *BookInfo) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### SetProviderIdsNil

`func (o *BookInfo) SetProviderIdsNil(b bool)`

 SetProviderIdsNil sets the value for ProviderIds to be an explicit nil

### UnsetProviderIds
`func (o *BookInfo) UnsetProviderIds()`

UnsetProviderIds ensures that no value is present for ProviderIds, not even an explicit nil
### GetYear

`func (o *BookInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *BookInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *BookInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *BookInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *BookInfo) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *BookInfo) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetIndexNumber

`func (o *BookInfo) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *BookInfo) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *BookInfo) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *BookInfo) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *BookInfo) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *BookInfo) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetParentIndexNumber

`func (o *BookInfo) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *BookInfo) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *BookInfo) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *BookInfo) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *BookInfo) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *BookInfo) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetPremiereDate

`func (o *BookInfo) GetPremiereDate() time.Time`

GetPremiereDate returns the PremiereDate field if non-nil, zero value otherwise.

### GetPremiereDateOk

`func (o *BookInfo) GetPremiereDateOk() (*time.Time, bool)`

GetPremiereDateOk returns a tuple with the PremiereDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiereDate

`func (o *BookInfo) SetPremiereDate(v time.Time)`

SetPremiereDate sets PremiereDate field to given value.

### HasPremiereDate

`func (o *BookInfo) HasPremiereDate() bool`

HasPremiereDate returns a boolean if a field has been set.

### SetPremiereDateNil

`func (o *BookInfo) SetPremiereDateNil(b bool)`

 SetPremiereDateNil sets the value for PremiereDate to be an explicit nil

### UnsetPremiereDate
`func (o *BookInfo) UnsetPremiereDate()`

UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
### GetIsAutomated

`func (o *BookInfo) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *BookInfo) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *BookInfo) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *BookInfo) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.

### GetSeriesName

`func (o *BookInfo) GetSeriesName() string`

GetSeriesName returns the SeriesName field if non-nil, zero value otherwise.

### GetSeriesNameOk

`func (o *BookInfo) GetSeriesNameOk() (*string, bool)`

GetSeriesNameOk returns a tuple with the SeriesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesName

`func (o *BookInfo) SetSeriesName(v string)`

SetSeriesName sets SeriesName field to given value.

### HasSeriesName

`func (o *BookInfo) HasSeriesName() bool`

HasSeriesName returns a boolean if a field has been set.

### SetSeriesNameNil

`func (o *BookInfo) SetSeriesNameNil(b bool)`

 SetSeriesNameNil sets the value for SeriesName to be an explicit nil

### UnsetSeriesName
`func (o *BookInfo) UnsetSeriesName()`

UnsetSeriesName ensures that no value is present for SeriesName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


