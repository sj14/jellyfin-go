# ForgotPasswordDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnteredUsername** | **string** | Gets or sets the entered username to have its password reset. | 

## Methods

### NewForgotPasswordDto

`func NewForgotPasswordDto(enteredUsername string, ) *ForgotPasswordDto`

NewForgotPasswordDto instantiates a new ForgotPasswordDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForgotPasswordDtoWithDefaults

`func NewForgotPasswordDtoWithDefaults() *ForgotPasswordDto`

NewForgotPasswordDtoWithDefaults instantiates a new ForgotPasswordDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnteredUsername

`func (o *ForgotPasswordDto) GetEnteredUsername() string`

GetEnteredUsername returns the EnteredUsername field if non-nil, zero value otherwise.

### GetEnteredUsernameOk

`func (o *ForgotPasswordDto) GetEnteredUsernameOk() (*string, bool)`

GetEnteredUsernameOk returns a tuple with the EnteredUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredUsername

`func (o *ForgotPasswordDto) SetEnteredUsername(v string)`

SetEnteredUsername sets EnteredUsername field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


