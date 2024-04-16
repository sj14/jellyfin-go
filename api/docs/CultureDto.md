# CultureDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Gets the name. | [optional] 
**DisplayName** | Pointer to **string** | Gets the display name. | [optional] 
**TwoLetterISOLanguageName** | Pointer to **string** | Gets the name of the two letter ISO language. | [optional] 
**ThreeLetterISOLanguageName** | Pointer to **NullableString** | Gets the name of the three letter ISO language. | [optional] [readonly] 
**ThreeLetterISOLanguageNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCultureDto

`func NewCultureDto() *CultureDto`

NewCultureDto instantiates a new CultureDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCultureDtoWithDefaults

`func NewCultureDtoWithDefaults() *CultureDto`

NewCultureDtoWithDefaults instantiates a new CultureDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CultureDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CultureDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CultureDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CultureDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *CultureDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CultureDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CultureDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CultureDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetTwoLetterISOLanguageName

`func (o *CultureDto) GetTwoLetterISOLanguageName() string`

GetTwoLetterISOLanguageName returns the TwoLetterISOLanguageName field if non-nil, zero value otherwise.

### GetTwoLetterISOLanguageNameOk

`func (o *CultureDto) GetTwoLetterISOLanguageNameOk() (*string, bool)`

GetTwoLetterISOLanguageNameOk returns a tuple with the TwoLetterISOLanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoLetterISOLanguageName

`func (o *CultureDto) SetTwoLetterISOLanguageName(v string)`

SetTwoLetterISOLanguageName sets TwoLetterISOLanguageName field to given value.

### HasTwoLetterISOLanguageName

`func (o *CultureDto) HasTwoLetterISOLanguageName() bool`

HasTwoLetterISOLanguageName returns a boolean if a field has been set.

### GetThreeLetterISOLanguageName

`func (o *CultureDto) GetThreeLetterISOLanguageName() string`

GetThreeLetterISOLanguageName returns the ThreeLetterISOLanguageName field if non-nil, zero value otherwise.

### GetThreeLetterISOLanguageNameOk

`func (o *CultureDto) GetThreeLetterISOLanguageNameOk() (*string, bool)`

GetThreeLetterISOLanguageNameOk returns a tuple with the ThreeLetterISOLanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeLetterISOLanguageName

`func (o *CultureDto) SetThreeLetterISOLanguageName(v string)`

SetThreeLetterISOLanguageName sets ThreeLetterISOLanguageName field to given value.

### HasThreeLetterISOLanguageName

`func (o *CultureDto) HasThreeLetterISOLanguageName() bool`

HasThreeLetterISOLanguageName returns a boolean if a field has been set.

### SetThreeLetterISOLanguageNameNil

`func (o *CultureDto) SetThreeLetterISOLanguageNameNil(b bool)`

 SetThreeLetterISOLanguageNameNil sets the value for ThreeLetterISOLanguageName to be an explicit nil

### UnsetThreeLetterISOLanguageName
`func (o *CultureDto) UnsetThreeLetterISOLanguageName()`

UnsetThreeLetterISOLanguageName ensures that no value is present for ThreeLetterISOLanguageName, not even an explicit nil
### GetThreeLetterISOLanguageNames

`func (o *CultureDto) GetThreeLetterISOLanguageNames() []string`

GetThreeLetterISOLanguageNames returns the ThreeLetterISOLanguageNames field if non-nil, zero value otherwise.

### GetThreeLetterISOLanguageNamesOk

`func (o *CultureDto) GetThreeLetterISOLanguageNamesOk() (*[]string, bool)`

GetThreeLetterISOLanguageNamesOk returns a tuple with the ThreeLetterISOLanguageNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeLetterISOLanguageNames

`func (o *CultureDto) SetThreeLetterISOLanguageNames(v []string)`

SetThreeLetterISOLanguageNames sets ThreeLetterISOLanguageNames field to given value.

### HasThreeLetterISOLanguageNames

`func (o *CultureDto) HasThreeLetterISOLanguageNames() bool`

HasThreeLetterISOLanguageNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


