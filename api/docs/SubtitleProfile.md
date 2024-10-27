# SubtitleProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **NullableString** | Gets or sets the format. | [optional] 
**Method** | Pointer to [**SubtitleDeliveryMethod**](SubtitleDeliveryMethod.md) | Gets or sets the delivery method. | [optional] 
**DidlMode** | Pointer to **NullableString** | Gets or sets the DIDL mode. | [optional] 
**Language** | Pointer to **NullableString** | Gets or sets the language. | [optional] 
**Container** | Pointer to **NullableString** | Gets or sets the container. | [optional] 

## Methods

### NewSubtitleProfile

`func NewSubtitleProfile() *SubtitleProfile`

NewSubtitleProfile instantiates a new SubtitleProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubtitleProfileWithDefaults

`func NewSubtitleProfileWithDefaults() *SubtitleProfile`

NewSubtitleProfileWithDefaults instantiates a new SubtitleProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *SubtitleProfile) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SubtitleProfile) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SubtitleProfile) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SubtitleProfile) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *SubtitleProfile) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *SubtitleProfile) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetMethod

`func (o *SubtitleProfile) GetMethod() SubtitleDeliveryMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SubtitleProfile) GetMethodOk() (*SubtitleDeliveryMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SubtitleProfile) SetMethod(v SubtitleDeliveryMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *SubtitleProfile) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetDidlMode

`func (o *SubtitleProfile) GetDidlMode() string`

GetDidlMode returns the DidlMode field if non-nil, zero value otherwise.

### GetDidlModeOk

`func (o *SubtitleProfile) GetDidlModeOk() (*string, bool)`

GetDidlModeOk returns a tuple with the DidlMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDidlMode

`func (o *SubtitleProfile) SetDidlMode(v string)`

SetDidlMode sets DidlMode field to given value.

### HasDidlMode

`func (o *SubtitleProfile) HasDidlMode() bool`

HasDidlMode returns a boolean if a field has been set.

### SetDidlModeNil

`func (o *SubtitleProfile) SetDidlModeNil(b bool)`

 SetDidlModeNil sets the value for DidlMode to be an explicit nil

### UnsetDidlMode
`func (o *SubtitleProfile) UnsetDidlMode()`

UnsetDidlMode ensures that no value is present for DidlMode, not even an explicit nil
### GetLanguage

`func (o *SubtitleProfile) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *SubtitleProfile) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *SubtitleProfile) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *SubtitleProfile) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *SubtitleProfile) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *SubtitleProfile) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetContainer

`func (o *SubtitleProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *SubtitleProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *SubtitleProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *SubtitleProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *SubtitleProfile) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *SubtitleProfile) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


