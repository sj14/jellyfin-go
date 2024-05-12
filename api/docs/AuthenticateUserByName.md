# AuthenticateUserByName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **NullableString** | Gets or sets the username. | [optional] 
**Pw** | Pointer to **NullableString** | Gets or sets the plain text password. | [optional] 

## Methods

### NewAuthenticateUserByName

`func NewAuthenticateUserByName() *AuthenticateUserByName`

NewAuthenticateUserByName instantiates a new AuthenticateUserByName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticateUserByNameWithDefaults

`func NewAuthenticateUserByNameWithDefaults() *AuthenticateUserByName`

NewAuthenticateUserByNameWithDefaults instantiates a new AuthenticateUserByName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AuthenticateUserByName) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AuthenticateUserByName) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AuthenticateUserByName) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AuthenticateUserByName) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *AuthenticateUserByName) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *AuthenticateUserByName) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPw

`func (o *AuthenticateUserByName) GetPw() string`

GetPw returns the Pw field if non-nil, zero value otherwise.

### GetPwOk

`func (o *AuthenticateUserByName) GetPwOk() (*string, bool)`

GetPwOk returns a tuple with the Pw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPw

`func (o *AuthenticateUserByName) SetPw(v string)`

SetPw sets Pw field to given value.

### HasPw

`func (o *AuthenticateUserByName) HasPw() bool`

HasPw returns a boolean if a field has been set.

### SetPwNil

`func (o *AuthenticateUserByName) SetPwNil(b bool)`

 SetPwNil sets the value for Pw to be an explicit nil

### UnsetPw
`func (o *AuthenticateUserByName) UnsetPw()`

UnsetPw ensures that no value is present for Pw, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


