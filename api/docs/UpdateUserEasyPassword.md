# UpdateUserEasyPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPassword** | Pointer to **NullableString** | Gets or sets the new sha1-hashed password. | [optional] 
**NewPw** | Pointer to **NullableString** | Gets or sets the new password. | [optional] 
**ResetPassword** | Pointer to **bool** | Gets or sets a value indicating whether to reset the password. | [optional] 

## Methods

### NewUpdateUserEasyPassword

`func NewUpdateUserEasyPassword() *UpdateUserEasyPassword`

NewUpdateUserEasyPassword instantiates a new UpdateUserEasyPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserEasyPasswordWithDefaults

`func NewUpdateUserEasyPasswordWithDefaults() *UpdateUserEasyPassword`

NewUpdateUserEasyPasswordWithDefaults instantiates a new UpdateUserEasyPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPassword

`func (o *UpdateUserEasyPassword) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UpdateUserEasyPassword) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UpdateUserEasyPassword) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *UpdateUserEasyPassword) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### SetNewPasswordNil

`func (o *UpdateUserEasyPassword) SetNewPasswordNil(b bool)`

 SetNewPasswordNil sets the value for NewPassword to be an explicit nil

### UnsetNewPassword
`func (o *UpdateUserEasyPassword) UnsetNewPassword()`

UnsetNewPassword ensures that no value is present for NewPassword, not even an explicit nil
### GetNewPw

`func (o *UpdateUserEasyPassword) GetNewPw() string`

GetNewPw returns the NewPw field if non-nil, zero value otherwise.

### GetNewPwOk

`func (o *UpdateUserEasyPassword) GetNewPwOk() (*string, bool)`

GetNewPwOk returns a tuple with the NewPw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPw

`func (o *UpdateUserEasyPassword) SetNewPw(v string)`

SetNewPw sets NewPw field to given value.

### HasNewPw

`func (o *UpdateUserEasyPassword) HasNewPw() bool`

HasNewPw returns a boolean if a field has been set.

### SetNewPwNil

`func (o *UpdateUserEasyPassword) SetNewPwNil(b bool)`

 SetNewPwNil sets the value for NewPw to be an explicit nil

### UnsetNewPw
`func (o *UpdateUserEasyPassword) UnsetNewPw()`

UnsetNewPw ensures that no value is present for NewPw, not even an explicit nil
### GetResetPassword

`func (o *UpdateUserEasyPassword) GetResetPassword() bool`

GetResetPassword returns the ResetPassword field if non-nil, zero value otherwise.

### GetResetPasswordOk

`func (o *UpdateUserEasyPassword) GetResetPasswordOk() (*bool, bool)`

GetResetPasswordOk returns a tuple with the ResetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPassword

`func (o *UpdateUserEasyPassword) SetResetPassword(v bool)`

SetResetPassword sets ResetPassword field to given value.

### HasResetPassword

`func (o *UpdateUserEasyPassword) HasResetPassword() bool`

HasResetPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


