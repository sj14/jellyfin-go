# CreateUserByName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Gets or sets the username. | 
**Password** | Pointer to **NullableString** | Gets or sets the password. | [optional] 

## Methods

### NewCreateUserByName

`func NewCreateUserByName(name string, ) *CreateUserByName`

NewCreateUserByName instantiates a new CreateUserByName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserByNameWithDefaults

`func NewCreateUserByNameWithDefaults() *CreateUserByName`

NewCreateUserByNameWithDefaults instantiates a new CreateUserByName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUserByName) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserByName) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserByName) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateUserByName) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateUserByName) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateUserByName) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateUserByName) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CreateUserByName) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CreateUserByName) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


