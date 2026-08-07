# AuthenticationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**NullableUserDto**](UserDto.md) | Gets or sets the user. | [optional] 
**SessionInfo** | Pointer to [**NullableSessionInfoDto**](SessionInfoDto.md) | Gets or sets the session info. | [optional] 
**AccessToken** | Pointer to **NullableString** | Gets or sets the access token. | [optional] 
**ServerId** | Pointer to **NullableString** | Gets or sets the server id. | [optional] 

## Methods

### NewAuthenticationResult

`func NewAuthenticationResult() *AuthenticationResult`

NewAuthenticationResult instantiates a new AuthenticationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationResultWithDefaults

`func NewAuthenticationResultWithDefaults() *AuthenticationResult`

NewAuthenticationResultWithDefaults instantiates a new AuthenticationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *AuthenticationResult) GetUser() UserDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AuthenticationResult) GetUserOk() (*UserDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AuthenticationResult) SetUser(v UserDto)`

SetUser sets User field to given value.

### HasUser

`func (o *AuthenticationResult) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *AuthenticationResult) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *AuthenticationResult) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetSessionInfo

`func (o *AuthenticationResult) GetSessionInfo() SessionInfoDto`

GetSessionInfo returns the SessionInfo field if non-nil, zero value otherwise.

### GetSessionInfoOk

`func (o *AuthenticationResult) GetSessionInfoOk() (*SessionInfoDto, bool)`

GetSessionInfoOk returns a tuple with the SessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionInfo

`func (o *AuthenticationResult) SetSessionInfo(v SessionInfoDto)`

SetSessionInfo sets SessionInfo field to given value.

### HasSessionInfo

`func (o *AuthenticationResult) HasSessionInfo() bool`

HasSessionInfo returns a boolean if a field has been set.

### SetSessionInfoNil

`func (o *AuthenticationResult) SetSessionInfoNil(b bool)`

 SetSessionInfoNil sets the value for SessionInfo to be an explicit nil

### UnsetSessionInfo
`func (o *AuthenticationResult) UnsetSessionInfo()`

UnsetSessionInfo ensures that no value is present for SessionInfo, not even an explicit nil
### GetAccessToken

`func (o *AuthenticationResult) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthenticationResult) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthenticationResult) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AuthenticationResult) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *AuthenticationResult) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *AuthenticationResult) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetServerId

`func (o *AuthenticationResult) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *AuthenticationResult) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *AuthenticationResult) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *AuthenticationResult) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *AuthenticationResult) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *AuthenticationResult) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


