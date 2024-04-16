# ForgotPasswordPinDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pin** | **string** | Gets or sets the entered pin to have the password reset. | 

## Methods

### NewForgotPasswordPinDto

`func NewForgotPasswordPinDto(pin string, ) *ForgotPasswordPinDto`

NewForgotPasswordPinDto instantiates a new ForgotPasswordPinDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForgotPasswordPinDtoWithDefaults

`func NewForgotPasswordPinDtoWithDefaults() *ForgotPasswordPinDto`

NewForgotPasswordPinDtoWithDefaults instantiates a new ForgotPasswordPinDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPin

`func (o *ForgotPasswordPinDto) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *ForgotPasswordPinDto) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *ForgotPasswordPinDto) SetPin(v string)`

SetPin sets Pin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


