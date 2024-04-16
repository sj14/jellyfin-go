# ForgotPasswordResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**ForgotPasswordAction**](ForgotPasswordAction.md) | Gets or sets the action. | [optional] 
**PinFile** | Pointer to **NullableString** | Gets or sets the pin file. | [optional] 
**PinExpirationDate** | Pointer to **NullableTime** | Gets or sets the pin expiration date. | [optional] 

## Methods

### NewForgotPasswordResult

`func NewForgotPasswordResult() *ForgotPasswordResult`

NewForgotPasswordResult instantiates a new ForgotPasswordResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForgotPasswordResultWithDefaults

`func NewForgotPasswordResultWithDefaults() *ForgotPasswordResult`

NewForgotPasswordResultWithDefaults instantiates a new ForgotPasswordResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ForgotPasswordResult) GetAction() ForgotPasswordAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ForgotPasswordResult) GetActionOk() (*ForgotPasswordAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ForgotPasswordResult) SetAction(v ForgotPasswordAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ForgotPasswordResult) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPinFile

`func (o *ForgotPasswordResult) GetPinFile() string`

GetPinFile returns the PinFile field if non-nil, zero value otherwise.

### GetPinFileOk

`func (o *ForgotPasswordResult) GetPinFileOk() (*string, bool)`

GetPinFileOk returns a tuple with the PinFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinFile

`func (o *ForgotPasswordResult) SetPinFile(v string)`

SetPinFile sets PinFile field to given value.

### HasPinFile

`func (o *ForgotPasswordResult) HasPinFile() bool`

HasPinFile returns a boolean if a field has been set.

### SetPinFileNil

`func (o *ForgotPasswordResult) SetPinFileNil(b bool)`

 SetPinFileNil sets the value for PinFile to be an explicit nil

### UnsetPinFile
`func (o *ForgotPasswordResult) UnsetPinFile()`

UnsetPinFile ensures that no value is present for PinFile, not even an explicit nil
### GetPinExpirationDate

`func (o *ForgotPasswordResult) GetPinExpirationDate() time.Time`

GetPinExpirationDate returns the PinExpirationDate field if non-nil, zero value otherwise.

### GetPinExpirationDateOk

`func (o *ForgotPasswordResult) GetPinExpirationDateOk() (*time.Time, bool)`

GetPinExpirationDateOk returns a tuple with the PinExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinExpirationDate

`func (o *ForgotPasswordResult) SetPinExpirationDate(v time.Time)`

SetPinExpirationDate sets PinExpirationDate field to given value.

### HasPinExpirationDate

`func (o *ForgotPasswordResult) HasPinExpirationDate() bool`

HasPinExpirationDate returns a boolean if a field has been set.

### SetPinExpirationDateNil

`func (o *ForgotPasswordResult) SetPinExpirationDateNil(b bool)`

 SetPinExpirationDateNil sets the value for PinExpirationDate to be an explicit nil

### UnsetPinExpirationDate
`func (o *ForgotPasswordResult) UnsetPinExpirationDate()`

UnsetPinExpirationDate ensures that no value is present for PinExpirationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


