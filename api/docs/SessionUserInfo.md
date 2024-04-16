# SessionUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | Gets or sets the user identifier. | [optional] 
**UserName** | Pointer to **NullableString** | Gets or sets the name of the user. | [optional] 

## Methods

### NewSessionUserInfo

`func NewSessionUserInfo() *SessionUserInfo`

NewSessionUserInfo instantiates a new SessionUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionUserInfoWithDefaults

`func NewSessionUserInfoWithDefaults() *SessionUserInfo`

NewSessionUserInfoWithDefaults instantiates a new SessionUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SessionUserInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SessionUserInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SessionUserInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SessionUserInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *SessionUserInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SessionUserInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SessionUserInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SessionUserInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *SessionUserInfo) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *SessionUserInfo) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


