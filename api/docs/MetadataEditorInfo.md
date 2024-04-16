# MetadataEditorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentalRatingOptions** | Pointer to [**[]ParentalRating**](ParentalRating.md) |  | [optional] 
**Countries** | Pointer to [**[]CountryInfo**](CountryInfo.md) |  | [optional] 
**Cultures** | Pointer to [**[]CultureDto**](CultureDto.md) |  | [optional] 
**ExternalIdInfos** | Pointer to [**[]ExternalIdInfo**](ExternalIdInfo.md) |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**ContentTypeOptions** | Pointer to [**[]NameValuePair**](NameValuePair.md) |  | [optional] 

## Methods

### NewMetadataEditorInfo

`func NewMetadataEditorInfo() *MetadataEditorInfo`

NewMetadataEditorInfo instantiates a new MetadataEditorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataEditorInfoWithDefaults

`func NewMetadataEditorInfoWithDefaults() *MetadataEditorInfo`

NewMetadataEditorInfoWithDefaults instantiates a new MetadataEditorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentalRatingOptions

`func (o *MetadataEditorInfo) GetParentalRatingOptions() []ParentalRating`

GetParentalRatingOptions returns the ParentalRatingOptions field if non-nil, zero value otherwise.

### GetParentalRatingOptionsOk

`func (o *MetadataEditorInfo) GetParentalRatingOptionsOk() (*[]ParentalRating, bool)`

GetParentalRatingOptionsOk returns a tuple with the ParentalRatingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentalRatingOptions

`func (o *MetadataEditorInfo) SetParentalRatingOptions(v []ParentalRating)`

SetParentalRatingOptions sets ParentalRatingOptions field to given value.

### HasParentalRatingOptions

`func (o *MetadataEditorInfo) HasParentalRatingOptions() bool`

HasParentalRatingOptions returns a boolean if a field has been set.

### GetCountries

`func (o *MetadataEditorInfo) GetCountries() []CountryInfo`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *MetadataEditorInfo) GetCountriesOk() (*[]CountryInfo, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *MetadataEditorInfo) SetCountries(v []CountryInfo)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *MetadataEditorInfo) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetCultures

`func (o *MetadataEditorInfo) GetCultures() []CultureDto`

GetCultures returns the Cultures field if non-nil, zero value otherwise.

### GetCulturesOk

`func (o *MetadataEditorInfo) GetCulturesOk() (*[]CultureDto, bool)`

GetCulturesOk returns a tuple with the Cultures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultures

`func (o *MetadataEditorInfo) SetCultures(v []CultureDto)`

SetCultures sets Cultures field to given value.

### HasCultures

`func (o *MetadataEditorInfo) HasCultures() bool`

HasCultures returns a boolean if a field has been set.

### GetExternalIdInfos

`func (o *MetadataEditorInfo) GetExternalIdInfos() []ExternalIdInfo`

GetExternalIdInfos returns the ExternalIdInfos field if non-nil, zero value otherwise.

### GetExternalIdInfosOk

`func (o *MetadataEditorInfo) GetExternalIdInfosOk() (*[]ExternalIdInfo, bool)`

GetExternalIdInfosOk returns a tuple with the ExternalIdInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdInfos

`func (o *MetadataEditorInfo) SetExternalIdInfos(v []ExternalIdInfo)`

SetExternalIdInfos sets ExternalIdInfos field to given value.

### HasExternalIdInfos

`func (o *MetadataEditorInfo) HasExternalIdInfos() bool`

HasExternalIdInfos returns a boolean if a field has been set.

### GetContentType

`func (o *MetadataEditorInfo) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *MetadataEditorInfo) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *MetadataEditorInfo) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *MetadataEditorInfo) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *MetadataEditorInfo) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *MetadataEditorInfo) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetContentTypeOptions

`func (o *MetadataEditorInfo) GetContentTypeOptions() []NameValuePair`

GetContentTypeOptions returns the ContentTypeOptions field if non-nil, zero value otherwise.

### GetContentTypeOptionsOk

`func (o *MetadataEditorInfo) GetContentTypeOptionsOk() (*[]NameValuePair, bool)`

GetContentTypeOptionsOk returns a tuple with the ContentTypeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypeOptions

`func (o *MetadataEditorInfo) SetContentTypeOptions(v []NameValuePair)`

SetContentTypeOptions sets ContentTypeOptions field to given value.

### HasContentTypeOptions

`func (o *MetadataEditorInfo) HasContentTypeOptions() bool`

HasContentTypeOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


