# BrandingOptionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginDisclaimer** | Pointer to **NullableString** | Gets or sets the login disclaimer. | [optional] 
**CustomCss** | Pointer to **NullableString** | Gets or sets the custom CSS. | [optional] 
**SplashscreenEnabled** | Pointer to **bool** | Gets or sets a value indicating whether to enable the splashscreen. | [optional] 

## Methods

### NewBrandingOptionsDto

`func NewBrandingOptionsDto() *BrandingOptionsDto`

NewBrandingOptionsDto instantiates a new BrandingOptionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingOptionsDtoWithDefaults

`func NewBrandingOptionsDtoWithDefaults() *BrandingOptionsDto`

NewBrandingOptionsDtoWithDefaults instantiates a new BrandingOptionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginDisclaimer

`func (o *BrandingOptionsDto) GetLoginDisclaimer() string`

GetLoginDisclaimer returns the LoginDisclaimer field if non-nil, zero value otherwise.

### GetLoginDisclaimerOk

`func (o *BrandingOptionsDto) GetLoginDisclaimerOk() (*string, bool)`

GetLoginDisclaimerOk returns a tuple with the LoginDisclaimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDisclaimer

`func (o *BrandingOptionsDto) SetLoginDisclaimer(v string)`

SetLoginDisclaimer sets LoginDisclaimer field to given value.

### HasLoginDisclaimer

`func (o *BrandingOptionsDto) HasLoginDisclaimer() bool`

HasLoginDisclaimer returns a boolean if a field has been set.

### SetLoginDisclaimerNil

`func (o *BrandingOptionsDto) SetLoginDisclaimerNil(b bool)`

 SetLoginDisclaimerNil sets the value for LoginDisclaimer to be an explicit nil

### UnsetLoginDisclaimer
`func (o *BrandingOptionsDto) UnsetLoginDisclaimer()`

UnsetLoginDisclaimer ensures that no value is present for LoginDisclaimer, not even an explicit nil
### GetCustomCss

`func (o *BrandingOptionsDto) GetCustomCss() string`

GetCustomCss returns the CustomCss field if non-nil, zero value otherwise.

### GetCustomCssOk

`func (o *BrandingOptionsDto) GetCustomCssOk() (*string, bool)`

GetCustomCssOk returns a tuple with the CustomCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCss

`func (o *BrandingOptionsDto) SetCustomCss(v string)`

SetCustomCss sets CustomCss field to given value.

### HasCustomCss

`func (o *BrandingOptionsDto) HasCustomCss() bool`

HasCustomCss returns a boolean if a field has been set.

### SetCustomCssNil

`func (o *BrandingOptionsDto) SetCustomCssNil(b bool)`

 SetCustomCssNil sets the value for CustomCss to be an explicit nil

### UnsetCustomCss
`func (o *BrandingOptionsDto) UnsetCustomCss()`

UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
### GetSplashscreenEnabled

`func (o *BrandingOptionsDto) GetSplashscreenEnabled() bool`

GetSplashscreenEnabled returns the SplashscreenEnabled field if non-nil, zero value otherwise.

### GetSplashscreenEnabledOk

`func (o *BrandingOptionsDto) GetSplashscreenEnabledOk() (*bool, bool)`

GetSplashscreenEnabledOk returns a tuple with the SplashscreenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplashscreenEnabled

`func (o *BrandingOptionsDto) SetSplashscreenEnabled(v bool)`

SetSplashscreenEnabled sets SplashscreenEnabled field to given value.

### HasSplashscreenEnabled

`func (o *BrandingOptionsDto) HasSplashscreenEnabled() bool`

HasSplashscreenEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


