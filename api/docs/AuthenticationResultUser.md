# AuthenticationResultUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**ServerId** | Pointer to **NullableString** | Gets or sets the server identifier. | [optional] 
**ServerName** | Pointer to **NullableString** | Gets or sets the name of the server.  This is not used by the server and is for client-side usage only. | [optional] 
**Id** | Pointer to **string** | Gets or sets the id. | [optional] 
**PrimaryImageTag** | Pointer to **NullableString** | Gets or sets the primary image tag. | [optional] 
**HasPassword** | Pointer to **bool** | Gets or sets a value indicating whether this instance has password. | [optional] 
**HasConfiguredPassword** | Pointer to **bool** | Gets or sets a value indicating whether this instance has configured password. | [optional] 
**HasConfiguredEasyPassword** | Pointer to **bool** | Gets or sets a value indicating whether this instance has configured easy password. | [optional] 
**EnableAutoLogin** | Pointer to **NullableBool** | Gets or sets whether async login is enabled or not. | [optional] 
**LastLoginDate** | Pointer to **NullableTime** | Gets or sets the last login date. | [optional] 
**LastActivityDate** | Pointer to **NullableTime** | Gets or sets the last activity date. | [optional] 
**Configuration** | Pointer to [**NullableUserDtoConfiguration**](UserDtoConfiguration.md) |  | [optional] 
**Policy** | Pointer to [**NullableUserDtoPolicy**](UserDtoPolicy.md) |  | [optional] 
**PrimaryImageAspectRatio** | Pointer to **NullableFloat64** | Gets or sets the primary image aspect ratio. | [optional] 

## Methods

### NewAuthenticationResultUser

`func NewAuthenticationResultUser() *AuthenticationResultUser`

NewAuthenticationResultUser instantiates a new AuthenticationResultUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationResultUserWithDefaults

`func NewAuthenticationResultUserWithDefaults() *AuthenticationResultUser`

NewAuthenticationResultUserWithDefaults instantiates a new AuthenticationResultUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthenticationResultUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthenticationResultUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthenticationResultUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthenticationResultUser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AuthenticationResultUser) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AuthenticationResultUser) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetServerId

`func (o *AuthenticationResultUser) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *AuthenticationResultUser) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *AuthenticationResultUser) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *AuthenticationResultUser) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *AuthenticationResultUser) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *AuthenticationResultUser) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetServerName

`func (o *AuthenticationResultUser) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AuthenticationResultUser) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AuthenticationResultUser) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *AuthenticationResultUser) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *AuthenticationResultUser) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *AuthenticationResultUser) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetId

`func (o *AuthenticationResultUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationResultUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationResultUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationResultUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrimaryImageTag

`func (o *AuthenticationResultUser) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *AuthenticationResultUser) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *AuthenticationResultUser) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *AuthenticationResultUser) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.

### SetPrimaryImageTagNil

`func (o *AuthenticationResultUser) SetPrimaryImageTagNil(b bool)`

 SetPrimaryImageTagNil sets the value for PrimaryImageTag to be an explicit nil

### UnsetPrimaryImageTag
`func (o *AuthenticationResultUser) UnsetPrimaryImageTag()`

UnsetPrimaryImageTag ensures that no value is present for PrimaryImageTag, not even an explicit nil
### GetHasPassword

`func (o *AuthenticationResultUser) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *AuthenticationResultUser) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *AuthenticationResultUser) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *AuthenticationResultUser) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetHasConfiguredPassword

`func (o *AuthenticationResultUser) GetHasConfiguredPassword() bool`

GetHasConfiguredPassword returns the HasConfiguredPassword field if non-nil, zero value otherwise.

### GetHasConfiguredPasswordOk

`func (o *AuthenticationResultUser) GetHasConfiguredPasswordOk() (*bool, bool)`

GetHasConfiguredPasswordOk returns a tuple with the HasConfiguredPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfiguredPassword

`func (o *AuthenticationResultUser) SetHasConfiguredPassword(v bool)`

SetHasConfiguredPassword sets HasConfiguredPassword field to given value.

### HasHasConfiguredPassword

`func (o *AuthenticationResultUser) HasHasConfiguredPassword() bool`

HasHasConfiguredPassword returns a boolean if a field has been set.

### GetHasConfiguredEasyPassword

`func (o *AuthenticationResultUser) GetHasConfiguredEasyPassword() bool`

GetHasConfiguredEasyPassword returns the HasConfiguredEasyPassword field if non-nil, zero value otherwise.

### GetHasConfiguredEasyPasswordOk

`func (o *AuthenticationResultUser) GetHasConfiguredEasyPasswordOk() (*bool, bool)`

GetHasConfiguredEasyPasswordOk returns a tuple with the HasConfiguredEasyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfiguredEasyPassword

`func (o *AuthenticationResultUser) SetHasConfiguredEasyPassword(v bool)`

SetHasConfiguredEasyPassword sets HasConfiguredEasyPassword field to given value.

### HasHasConfiguredEasyPassword

`func (o *AuthenticationResultUser) HasHasConfiguredEasyPassword() bool`

HasHasConfiguredEasyPassword returns a boolean if a field has been set.

### GetEnableAutoLogin

`func (o *AuthenticationResultUser) GetEnableAutoLogin() bool`

GetEnableAutoLogin returns the EnableAutoLogin field if non-nil, zero value otherwise.

### GetEnableAutoLoginOk

`func (o *AuthenticationResultUser) GetEnableAutoLoginOk() (*bool, bool)`

GetEnableAutoLoginOk returns a tuple with the EnableAutoLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoLogin

`func (o *AuthenticationResultUser) SetEnableAutoLogin(v bool)`

SetEnableAutoLogin sets EnableAutoLogin field to given value.

### HasEnableAutoLogin

`func (o *AuthenticationResultUser) HasEnableAutoLogin() bool`

HasEnableAutoLogin returns a boolean if a field has been set.

### SetEnableAutoLoginNil

`func (o *AuthenticationResultUser) SetEnableAutoLoginNil(b bool)`

 SetEnableAutoLoginNil sets the value for EnableAutoLogin to be an explicit nil

### UnsetEnableAutoLogin
`func (o *AuthenticationResultUser) UnsetEnableAutoLogin()`

UnsetEnableAutoLogin ensures that no value is present for EnableAutoLogin, not even an explicit nil
### GetLastLoginDate

`func (o *AuthenticationResultUser) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *AuthenticationResultUser) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *AuthenticationResultUser) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *AuthenticationResultUser) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### SetLastLoginDateNil

`func (o *AuthenticationResultUser) SetLastLoginDateNil(b bool)`

 SetLastLoginDateNil sets the value for LastLoginDate to be an explicit nil

### UnsetLastLoginDate
`func (o *AuthenticationResultUser) UnsetLastLoginDate()`

UnsetLastLoginDate ensures that no value is present for LastLoginDate, not even an explicit nil
### GetLastActivityDate

`func (o *AuthenticationResultUser) GetLastActivityDate() time.Time`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *AuthenticationResultUser) GetLastActivityDateOk() (*time.Time, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *AuthenticationResultUser) SetLastActivityDate(v time.Time)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *AuthenticationResultUser) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### SetLastActivityDateNil

`func (o *AuthenticationResultUser) SetLastActivityDateNil(b bool)`

 SetLastActivityDateNil sets the value for LastActivityDate to be an explicit nil

### UnsetLastActivityDate
`func (o *AuthenticationResultUser) UnsetLastActivityDate()`

UnsetLastActivityDate ensures that no value is present for LastActivityDate, not even an explicit nil
### GetConfiguration

`func (o *AuthenticationResultUser) GetConfiguration() UserDtoConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AuthenticationResultUser) GetConfigurationOk() (*UserDtoConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AuthenticationResultUser) SetConfiguration(v UserDtoConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *AuthenticationResultUser) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *AuthenticationResultUser) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *AuthenticationResultUser) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetPolicy

`func (o *AuthenticationResultUser) GetPolicy() UserDtoPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *AuthenticationResultUser) GetPolicyOk() (*UserDtoPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *AuthenticationResultUser) SetPolicy(v UserDtoPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *AuthenticationResultUser) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *AuthenticationResultUser) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *AuthenticationResultUser) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetPrimaryImageAspectRatio

`func (o *AuthenticationResultUser) GetPrimaryImageAspectRatio() float64`

GetPrimaryImageAspectRatio returns the PrimaryImageAspectRatio field if non-nil, zero value otherwise.

### GetPrimaryImageAspectRatioOk

`func (o *AuthenticationResultUser) GetPrimaryImageAspectRatioOk() (*float64, bool)`

GetPrimaryImageAspectRatioOk returns a tuple with the PrimaryImageAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageAspectRatio

`func (o *AuthenticationResultUser) SetPrimaryImageAspectRatio(v float64)`

SetPrimaryImageAspectRatio sets PrimaryImageAspectRatio field to given value.

### HasPrimaryImageAspectRatio

`func (o *AuthenticationResultUser) HasPrimaryImageAspectRatio() bool`

HasPrimaryImageAspectRatio returns a boolean if a field has been set.

### SetPrimaryImageAspectRatioNil

`func (o *AuthenticationResultUser) SetPrimaryImageAspectRatioNil(b bool)`

 SetPrimaryImageAspectRatioNil sets the value for PrimaryImageAspectRatio to be an explicit nil

### UnsetPrimaryImageAspectRatio
`func (o *AuthenticationResultUser) UnsetPrimaryImageAspectRatio()`

UnsetPrimaryImageAspectRatio ensures that no value is present for PrimaryImageAspectRatio, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


