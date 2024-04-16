# TrailerInfo

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

## Methods

### NewTrailerInfo

`func NewTrailerInfo() *TrailerInfo`

NewTrailerInfo instantiates a new TrailerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailerInfoWithDefaults

`func NewTrailerInfoWithDefaults() *TrailerInfo`

NewTrailerInfoWithDefaults instantiates a new TrailerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TrailerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrailerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrailerInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrailerInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TrailerInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TrailerInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOriginalTitle

`func (o *TrailerInfo) GetOriginalTitle() string`

GetOriginalTitle returns the OriginalTitle field if non-nil, zero value otherwise.

### GetOriginalTitleOk

`func (o *TrailerInfo) GetOriginalTitleOk() (*string, bool)`

GetOriginalTitleOk returns a tuple with the OriginalTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTitle

`func (o *TrailerInfo) SetOriginalTitle(v string)`

SetOriginalTitle sets OriginalTitle field to given value.

### HasOriginalTitle

`func (o *TrailerInfo) HasOriginalTitle() bool`

HasOriginalTitle returns a boolean if a field has been set.

### SetOriginalTitleNil

`func (o *TrailerInfo) SetOriginalTitleNil(b bool)`

 SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil

### UnsetOriginalTitle
`func (o *TrailerInfo) UnsetOriginalTitle()`

UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
### GetPath

`func (o *TrailerInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TrailerInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TrailerInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TrailerInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *TrailerInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *TrailerInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetMetadataLanguage

`func (o *TrailerInfo) GetMetadataLanguage() string`

GetMetadataLanguage returns the MetadataLanguage field if non-nil, zero value otherwise.

### GetMetadataLanguageOk

`func (o *TrailerInfo) GetMetadataLanguageOk() (*string, bool)`

GetMetadataLanguageOk returns a tuple with the MetadataLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLanguage

`func (o *TrailerInfo) SetMetadataLanguage(v string)`

SetMetadataLanguage sets MetadataLanguage field to given value.

### HasMetadataLanguage

`func (o *TrailerInfo) HasMetadataLanguage() bool`

HasMetadataLanguage returns a boolean if a field has been set.

### SetMetadataLanguageNil

`func (o *TrailerInfo) SetMetadataLanguageNil(b bool)`

 SetMetadataLanguageNil sets the value for MetadataLanguage to be an explicit nil

### UnsetMetadataLanguage
`func (o *TrailerInfo) UnsetMetadataLanguage()`

UnsetMetadataLanguage ensures that no value is present for MetadataLanguage, not even an explicit nil
### GetMetadataCountryCode

`func (o *TrailerInfo) GetMetadataCountryCode() string`

GetMetadataCountryCode returns the MetadataCountryCode field if non-nil, zero value otherwise.

### GetMetadataCountryCodeOk

`func (o *TrailerInfo) GetMetadataCountryCodeOk() (*string, bool)`

GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataCountryCode

`func (o *TrailerInfo) SetMetadataCountryCode(v string)`

SetMetadataCountryCode sets MetadataCountryCode field to given value.

### HasMetadataCountryCode

`func (o *TrailerInfo) HasMetadataCountryCode() bool`

HasMetadataCountryCode returns a boolean if a field has been set.

### SetMetadataCountryCodeNil

`func (o *TrailerInfo) SetMetadataCountryCodeNil(b bool)`

 SetMetadataCountryCodeNil sets the value for MetadataCountryCode to be an explicit nil

### UnsetMetadataCountryCode
`func (o *TrailerInfo) UnsetMetadataCountryCode()`

UnsetMetadataCountryCode ensures that no value is present for MetadataCountryCode, not even an explicit nil
### GetProviderIds

`func (o *TrailerInfo) GetProviderIds() map[string]string`

GetProviderIds returns the ProviderIds field if non-nil, zero value otherwise.

### GetProviderIdsOk

`func (o *TrailerInfo) GetProviderIdsOk() (*map[string]string, bool)`

GetProviderIdsOk returns a tuple with the ProviderIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderIds

`func (o *TrailerInfo) SetProviderIds(v map[string]string)`

SetProviderIds sets ProviderIds field to given value.

### HasProviderIds

`func (o *TrailerInfo) HasProviderIds() bool`

HasProviderIds returns a boolean if a field has been set.

### SetProviderIdsNil

`func (o *TrailerInfo) SetProviderIdsNil(b bool)`

 SetProviderIdsNil sets the value for ProviderIds to be an explicit nil

### UnsetProviderIds
`func (o *TrailerInfo) UnsetProviderIds()`

UnsetProviderIds ensures that no value is present for ProviderIds, not even an explicit nil
### GetYear

`func (o *TrailerInfo) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *TrailerInfo) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *TrailerInfo) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *TrailerInfo) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *TrailerInfo) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *TrailerInfo) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetIndexNumber

`func (o *TrailerInfo) GetIndexNumber() int32`

GetIndexNumber returns the IndexNumber field if non-nil, zero value otherwise.

### GetIndexNumberOk

`func (o *TrailerInfo) GetIndexNumberOk() (*int32, bool)`

GetIndexNumberOk returns a tuple with the IndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexNumber

`func (o *TrailerInfo) SetIndexNumber(v int32)`

SetIndexNumber sets IndexNumber field to given value.

### HasIndexNumber

`func (o *TrailerInfo) HasIndexNumber() bool`

HasIndexNumber returns a boolean if a field has been set.

### SetIndexNumberNil

`func (o *TrailerInfo) SetIndexNumberNil(b bool)`

 SetIndexNumberNil sets the value for IndexNumber to be an explicit nil

### UnsetIndexNumber
`func (o *TrailerInfo) UnsetIndexNumber()`

UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
### GetParentIndexNumber

`func (o *TrailerInfo) GetParentIndexNumber() int32`

GetParentIndexNumber returns the ParentIndexNumber field if non-nil, zero value otherwise.

### GetParentIndexNumberOk

`func (o *TrailerInfo) GetParentIndexNumberOk() (*int32, bool)`

GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIndexNumber

`func (o *TrailerInfo) SetParentIndexNumber(v int32)`

SetParentIndexNumber sets ParentIndexNumber field to given value.

### HasParentIndexNumber

`func (o *TrailerInfo) HasParentIndexNumber() bool`

HasParentIndexNumber returns a boolean if a field has been set.

### SetParentIndexNumberNil

`func (o *TrailerInfo) SetParentIndexNumberNil(b bool)`

 SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil

### UnsetParentIndexNumber
`func (o *TrailerInfo) UnsetParentIndexNumber()`

UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
### GetPremiereDate

`func (o *TrailerInfo) GetPremiereDate() time.Time`

GetPremiereDate returns the PremiereDate field if non-nil, zero value otherwise.

### GetPremiereDateOk

`func (o *TrailerInfo) GetPremiereDateOk() (*time.Time, bool)`

GetPremiereDateOk returns a tuple with the PremiereDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiereDate

`func (o *TrailerInfo) SetPremiereDate(v time.Time)`

SetPremiereDate sets PremiereDate field to given value.

### HasPremiereDate

`func (o *TrailerInfo) HasPremiereDate() bool`

HasPremiereDate returns a boolean if a field has been set.

### SetPremiereDateNil

`func (o *TrailerInfo) SetPremiereDateNil(b bool)`

 SetPremiereDateNil sets the value for PremiereDate to be an explicit nil

### UnsetPremiereDate
`func (o *TrailerInfo) UnsetPremiereDate()`

UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
### GetIsAutomated

`func (o *TrailerInfo) GetIsAutomated() bool`

GetIsAutomated returns the IsAutomated field if non-nil, zero value otherwise.

### GetIsAutomatedOk

`func (o *TrailerInfo) GetIsAutomatedOk() (*bool, bool)`

GetIsAutomatedOk returns a tuple with the IsAutomated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomated

`func (o *TrailerInfo) SetIsAutomated(v bool)`

SetIsAutomated sets IsAutomated field to given value.

### HasIsAutomated

`func (o *TrailerInfo) HasIsAutomated() bool`

HasIsAutomated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


