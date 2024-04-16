# ListingsProviderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to **NullableString** |  | [optional] 
**ListingsId** | Pointer to **NullableString** |  | [optional] 
**ZipCode** | Pointer to **NullableString** |  | [optional] 
**Country** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** |  | [optional] 
**EnabledTuners** | Pointer to **[]string** |  | [optional] 
**EnableAllTuners** | Pointer to **bool** |  | [optional] 
**NewsCategories** | Pointer to **[]string** |  | [optional] 
**SportsCategories** | Pointer to **[]string** |  | [optional] 
**KidsCategories** | Pointer to **[]string** |  | [optional] 
**MovieCategories** | Pointer to **[]string** |  | [optional] 
**ChannelMappings** | Pointer to [**[]NameValuePair**](NameValuePair.md) |  | [optional] 
**MoviePrefix** | Pointer to **NullableString** |  | [optional] 
**PreferredLanguage** | Pointer to **NullableString** |  | [optional] 
**UserAgent** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListingsProviderInfo

`func NewListingsProviderInfo() *ListingsProviderInfo`

NewListingsProviderInfo instantiates a new ListingsProviderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListingsProviderInfoWithDefaults

`func NewListingsProviderInfoWithDefaults() *ListingsProviderInfo`

NewListingsProviderInfoWithDefaults instantiates a new ListingsProviderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListingsProviderInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListingsProviderInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListingsProviderInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListingsProviderInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ListingsProviderInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ListingsProviderInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *ListingsProviderInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListingsProviderInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListingsProviderInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListingsProviderInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ListingsProviderInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ListingsProviderInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUsername

`func (o *ListingsProviderInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListingsProviderInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListingsProviderInfo) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ListingsProviderInfo) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ListingsProviderInfo) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ListingsProviderInfo) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ListingsProviderInfo) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ListingsProviderInfo) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ListingsProviderInfo) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ListingsProviderInfo) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ListingsProviderInfo) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ListingsProviderInfo) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetListingsId

`func (o *ListingsProviderInfo) GetListingsId() string`

GetListingsId returns the ListingsId field if non-nil, zero value otherwise.

### GetListingsIdOk

`func (o *ListingsProviderInfo) GetListingsIdOk() (*string, bool)`

GetListingsIdOk returns a tuple with the ListingsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingsId

`func (o *ListingsProviderInfo) SetListingsId(v string)`

SetListingsId sets ListingsId field to given value.

### HasListingsId

`func (o *ListingsProviderInfo) HasListingsId() bool`

HasListingsId returns a boolean if a field has been set.

### SetListingsIdNil

`func (o *ListingsProviderInfo) SetListingsIdNil(b bool)`

 SetListingsIdNil sets the value for ListingsId to be an explicit nil

### UnsetListingsId
`func (o *ListingsProviderInfo) UnsetListingsId()`

UnsetListingsId ensures that no value is present for ListingsId, not even an explicit nil
### GetZipCode

`func (o *ListingsProviderInfo) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *ListingsProviderInfo) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *ListingsProviderInfo) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *ListingsProviderInfo) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *ListingsProviderInfo) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *ListingsProviderInfo) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
### GetCountry

`func (o *ListingsProviderInfo) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ListingsProviderInfo) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ListingsProviderInfo) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ListingsProviderInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ListingsProviderInfo) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ListingsProviderInfo) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetPath

`func (o *ListingsProviderInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ListingsProviderInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ListingsProviderInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ListingsProviderInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ListingsProviderInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ListingsProviderInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetEnabledTuners

`func (o *ListingsProviderInfo) GetEnabledTuners() []string`

GetEnabledTuners returns the EnabledTuners field if non-nil, zero value otherwise.

### GetEnabledTunersOk

`func (o *ListingsProviderInfo) GetEnabledTunersOk() (*[]string, bool)`

GetEnabledTunersOk returns a tuple with the EnabledTuners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledTuners

`func (o *ListingsProviderInfo) SetEnabledTuners(v []string)`

SetEnabledTuners sets EnabledTuners field to given value.

### HasEnabledTuners

`func (o *ListingsProviderInfo) HasEnabledTuners() bool`

HasEnabledTuners returns a boolean if a field has been set.

### SetEnabledTunersNil

`func (o *ListingsProviderInfo) SetEnabledTunersNil(b bool)`

 SetEnabledTunersNil sets the value for EnabledTuners to be an explicit nil

### UnsetEnabledTuners
`func (o *ListingsProviderInfo) UnsetEnabledTuners()`

UnsetEnabledTuners ensures that no value is present for EnabledTuners, not even an explicit nil
### GetEnableAllTuners

`func (o *ListingsProviderInfo) GetEnableAllTuners() bool`

GetEnableAllTuners returns the EnableAllTuners field if non-nil, zero value otherwise.

### GetEnableAllTunersOk

`func (o *ListingsProviderInfo) GetEnableAllTunersOk() (*bool, bool)`

GetEnableAllTunersOk returns a tuple with the EnableAllTuners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllTuners

`func (o *ListingsProviderInfo) SetEnableAllTuners(v bool)`

SetEnableAllTuners sets EnableAllTuners field to given value.

### HasEnableAllTuners

`func (o *ListingsProviderInfo) HasEnableAllTuners() bool`

HasEnableAllTuners returns a boolean if a field has been set.

### GetNewsCategories

`func (o *ListingsProviderInfo) GetNewsCategories() []string`

GetNewsCategories returns the NewsCategories field if non-nil, zero value otherwise.

### GetNewsCategoriesOk

`func (o *ListingsProviderInfo) GetNewsCategoriesOk() (*[]string, bool)`

GetNewsCategoriesOk returns a tuple with the NewsCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewsCategories

`func (o *ListingsProviderInfo) SetNewsCategories(v []string)`

SetNewsCategories sets NewsCategories field to given value.

### HasNewsCategories

`func (o *ListingsProviderInfo) HasNewsCategories() bool`

HasNewsCategories returns a boolean if a field has been set.

### SetNewsCategoriesNil

`func (o *ListingsProviderInfo) SetNewsCategoriesNil(b bool)`

 SetNewsCategoriesNil sets the value for NewsCategories to be an explicit nil

### UnsetNewsCategories
`func (o *ListingsProviderInfo) UnsetNewsCategories()`

UnsetNewsCategories ensures that no value is present for NewsCategories, not even an explicit nil
### GetSportsCategories

`func (o *ListingsProviderInfo) GetSportsCategories() []string`

GetSportsCategories returns the SportsCategories field if non-nil, zero value otherwise.

### GetSportsCategoriesOk

`func (o *ListingsProviderInfo) GetSportsCategoriesOk() (*[]string, bool)`

GetSportsCategoriesOk returns a tuple with the SportsCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSportsCategories

`func (o *ListingsProviderInfo) SetSportsCategories(v []string)`

SetSportsCategories sets SportsCategories field to given value.

### HasSportsCategories

`func (o *ListingsProviderInfo) HasSportsCategories() bool`

HasSportsCategories returns a boolean if a field has been set.

### SetSportsCategoriesNil

`func (o *ListingsProviderInfo) SetSportsCategoriesNil(b bool)`

 SetSportsCategoriesNil sets the value for SportsCategories to be an explicit nil

### UnsetSportsCategories
`func (o *ListingsProviderInfo) UnsetSportsCategories()`

UnsetSportsCategories ensures that no value is present for SportsCategories, not even an explicit nil
### GetKidsCategories

`func (o *ListingsProviderInfo) GetKidsCategories() []string`

GetKidsCategories returns the KidsCategories field if non-nil, zero value otherwise.

### GetKidsCategoriesOk

`func (o *ListingsProviderInfo) GetKidsCategoriesOk() (*[]string, bool)`

GetKidsCategoriesOk returns a tuple with the KidsCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKidsCategories

`func (o *ListingsProviderInfo) SetKidsCategories(v []string)`

SetKidsCategories sets KidsCategories field to given value.

### HasKidsCategories

`func (o *ListingsProviderInfo) HasKidsCategories() bool`

HasKidsCategories returns a boolean if a field has been set.

### SetKidsCategoriesNil

`func (o *ListingsProviderInfo) SetKidsCategoriesNil(b bool)`

 SetKidsCategoriesNil sets the value for KidsCategories to be an explicit nil

### UnsetKidsCategories
`func (o *ListingsProviderInfo) UnsetKidsCategories()`

UnsetKidsCategories ensures that no value is present for KidsCategories, not even an explicit nil
### GetMovieCategories

`func (o *ListingsProviderInfo) GetMovieCategories() []string`

GetMovieCategories returns the MovieCategories field if non-nil, zero value otherwise.

### GetMovieCategoriesOk

`func (o *ListingsProviderInfo) GetMovieCategoriesOk() (*[]string, bool)`

GetMovieCategoriesOk returns a tuple with the MovieCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieCategories

`func (o *ListingsProviderInfo) SetMovieCategories(v []string)`

SetMovieCategories sets MovieCategories field to given value.

### HasMovieCategories

`func (o *ListingsProviderInfo) HasMovieCategories() bool`

HasMovieCategories returns a boolean if a field has been set.

### SetMovieCategoriesNil

`func (o *ListingsProviderInfo) SetMovieCategoriesNil(b bool)`

 SetMovieCategoriesNil sets the value for MovieCategories to be an explicit nil

### UnsetMovieCategories
`func (o *ListingsProviderInfo) UnsetMovieCategories()`

UnsetMovieCategories ensures that no value is present for MovieCategories, not even an explicit nil
### GetChannelMappings

`func (o *ListingsProviderInfo) GetChannelMappings() []NameValuePair`

GetChannelMappings returns the ChannelMappings field if non-nil, zero value otherwise.

### GetChannelMappingsOk

`func (o *ListingsProviderInfo) GetChannelMappingsOk() (*[]NameValuePair, bool)`

GetChannelMappingsOk returns a tuple with the ChannelMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelMappings

`func (o *ListingsProviderInfo) SetChannelMappings(v []NameValuePair)`

SetChannelMappings sets ChannelMappings field to given value.

### HasChannelMappings

`func (o *ListingsProviderInfo) HasChannelMappings() bool`

HasChannelMappings returns a boolean if a field has been set.

### SetChannelMappingsNil

`func (o *ListingsProviderInfo) SetChannelMappingsNil(b bool)`

 SetChannelMappingsNil sets the value for ChannelMappings to be an explicit nil

### UnsetChannelMappings
`func (o *ListingsProviderInfo) UnsetChannelMappings()`

UnsetChannelMappings ensures that no value is present for ChannelMappings, not even an explicit nil
### GetMoviePrefix

`func (o *ListingsProviderInfo) GetMoviePrefix() string`

GetMoviePrefix returns the MoviePrefix field if non-nil, zero value otherwise.

### GetMoviePrefixOk

`func (o *ListingsProviderInfo) GetMoviePrefixOk() (*string, bool)`

GetMoviePrefixOk returns a tuple with the MoviePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoviePrefix

`func (o *ListingsProviderInfo) SetMoviePrefix(v string)`

SetMoviePrefix sets MoviePrefix field to given value.

### HasMoviePrefix

`func (o *ListingsProviderInfo) HasMoviePrefix() bool`

HasMoviePrefix returns a boolean if a field has been set.

### SetMoviePrefixNil

`func (o *ListingsProviderInfo) SetMoviePrefixNil(b bool)`

 SetMoviePrefixNil sets the value for MoviePrefix to be an explicit nil

### UnsetMoviePrefix
`func (o *ListingsProviderInfo) UnsetMoviePrefix()`

UnsetMoviePrefix ensures that no value is present for MoviePrefix, not even an explicit nil
### GetPreferredLanguage

`func (o *ListingsProviderInfo) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *ListingsProviderInfo) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *ListingsProviderInfo) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *ListingsProviderInfo) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguageNil

`func (o *ListingsProviderInfo) SetPreferredLanguageNil(b bool)`

 SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil

### UnsetPreferredLanguage
`func (o *ListingsProviderInfo) UnsetPreferredLanguage()`

UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
### GetUserAgent

`func (o *ListingsProviderInfo) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ListingsProviderInfo) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ListingsProviderInfo) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ListingsProviderInfo) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *ListingsProviderInfo) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *ListingsProviderInfo) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


