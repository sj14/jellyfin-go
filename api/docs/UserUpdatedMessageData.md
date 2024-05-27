# UserUpdatedMessageData

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

### NewUserUpdatedMessageData

`func NewUserUpdatedMessageData() *UserUpdatedMessageData`

NewUserUpdatedMessageData instantiates a new UserUpdatedMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdatedMessageDataWithDefaults

`func NewUserUpdatedMessageDataWithDefaults() *UserUpdatedMessageData`

NewUserUpdatedMessageDataWithDefaults instantiates a new UserUpdatedMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserUpdatedMessageData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserUpdatedMessageData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserUpdatedMessageData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserUpdatedMessageData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UserUpdatedMessageData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UserUpdatedMessageData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetServerId

`func (o *UserUpdatedMessageData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *UserUpdatedMessageData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *UserUpdatedMessageData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *UserUpdatedMessageData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *UserUpdatedMessageData) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *UserUpdatedMessageData) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetServerName

`func (o *UserUpdatedMessageData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *UserUpdatedMessageData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *UserUpdatedMessageData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *UserUpdatedMessageData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *UserUpdatedMessageData) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *UserUpdatedMessageData) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetId

`func (o *UserUpdatedMessageData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserUpdatedMessageData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserUpdatedMessageData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserUpdatedMessageData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPrimaryImageTag

`func (o *UserUpdatedMessageData) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *UserUpdatedMessageData) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *UserUpdatedMessageData) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *UserUpdatedMessageData) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.

### SetPrimaryImageTagNil

`func (o *UserUpdatedMessageData) SetPrimaryImageTagNil(b bool)`

 SetPrimaryImageTagNil sets the value for PrimaryImageTag to be an explicit nil

### UnsetPrimaryImageTag
`func (o *UserUpdatedMessageData) UnsetPrimaryImageTag()`

UnsetPrimaryImageTag ensures that no value is present for PrimaryImageTag, not even an explicit nil
### GetHasPassword

`func (o *UserUpdatedMessageData) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *UserUpdatedMessageData) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *UserUpdatedMessageData) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *UserUpdatedMessageData) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetHasConfiguredPassword

`func (o *UserUpdatedMessageData) GetHasConfiguredPassword() bool`

GetHasConfiguredPassword returns the HasConfiguredPassword field if non-nil, zero value otherwise.

### GetHasConfiguredPasswordOk

`func (o *UserUpdatedMessageData) GetHasConfiguredPasswordOk() (*bool, bool)`

GetHasConfiguredPasswordOk returns a tuple with the HasConfiguredPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfiguredPassword

`func (o *UserUpdatedMessageData) SetHasConfiguredPassword(v bool)`

SetHasConfiguredPassword sets HasConfiguredPassword field to given value.

### HasHasConfiguredPassword

`func (o *UserUpdatedMessageData) HasHasConfiguredPassword() bool`

HasHasConfiguredPassword returns a boolean if a field has been set.

### GetHasConfiguredEasyPassword

`func (o *UserUpdatedMessageData) GetHasConfiguredEasyPassword() bool`

GetHasConfiguredEasyPassword returns the HasConfiguredEasyPassword field if non-nil, zero value otherwise.

### GetHasConfiguredEasyPasswordOk

`func (o *UserUpdatedMessageData) GetHasConfiguredEasyPasswordOk() (*bool, bool)`

GetHasConfiguredEasyPasswordOk returns a tuple with the HasConfiguredEasyPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfiguredEasyPassword

`func (o *UserUpdatedMessageData) SetHasConfiguredEasyPassword(v bool)`

SetHasConfiguredEasyPassword sets HasConfiguredEasyPassword field to given value.

### HasHasConfiguredEasyPassword

`func (o *UserUpdatedMessageData) HasHasConfiguredEasyPassword() bool`

HasHasConfiguredEasyPassword returns a boolean if a field has been set.

### GetEnableAutoLogin

`func (o *UserUpdatedMessageData) GetEnableAutoLogin() bool`

GetEnableAutoLogin returns the EnableAutoLogin field if non-nil, zero value otherwise.

### GetEnableAutoLoginOk

`func (o *UserUpdatedMessageData) GetEnableAutoLoginOk() (*bool, bool)`

GetEnableAutoLoginOk returns a tuple with the EnableAutoLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoLogin

`func (o *UserUpdatedMessageData) SetEnableAutoLogin(v bool)`

SetEnableAutoLogin sets EnableAutoLogin field to given value.

### HasEnableAutoLogin

`func (o *UserUpdatedMessageData) HasEnableAutoLogin() bool`

HasEnableAutoLogin returns a boolean if a field has been set.

### SetEnableAutoLoginNil

`func (o *UserUpdatedMessageData) SetEnableAutoLoginNil(b bool)`

 SetEnableAutoLoginNil sets the value for EnableAutoLogin to be an explicit nil

### UnsetEnableAutoLogin
`func (o *UserUpdatedMessageData) UnsetEnableAutoLogin()`

UnsetEnableAutoLogin ensures that no value is present for EnableAutoLogin, not even an explicit nil
### GetLastLoginDate

`func (o *UserUpdatedMessageData) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *UserUpdatedMessageData) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *UserUpdatedMessageData) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *UserUpdatedMessageData) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### SetLastLoginDateNil

`func (o *UserUpdatedMessageData) SetLastLoginDateNil(b bool)`

 SetLastLoginDateNil sets the value for LastLoginDate to be an explicit nil

### UnsetLastLoginDate
`func (o *UserUpdatedMessageData) UnsetLastLoginDate()`

UnsetLastLoginDate ensures that no value is present for LastLoginDate, not even an explicit nil
### GetLastActivityDate

`func (o *UserUpdatedMessageData) GetLastActivityDate() time.Time`

GetLastActivityDate returns the LastActivityDate field if non-nil, zero value otherwise.

### GetLastActivityDateOk

`func (o *UserUpdatedMessageData) GetLastActivityDateOk() (*time.Time, bool)`

GetLastActivityDateOk returns a tuple with the LastActivityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivityDate

`func (o *UserUpdatedMessageData) SetLastActivityDate(v time.Time)`

SetLastActivityDate sets LastActivityDate field to given value.

### HasLastActivityDate

`func (o *UserUpdatedMessageData) HasLastActivityDate() bool`

HasLastActivityDate returns a boolean if a field has been set.

### SetLastActivityDateNil

`func (o *UserUpdatedMessageData) SetLastActivityDateNil(b bool)`

 SetLastActivityDateNil sets the value for LastActivityDate to be an explicit nil

### UnsetLastActivityDate
`func (o *UserUpdatedMessageData) UnsetLastActivityDate()`

UnsetLastActivityDate ensures that no value is present for LastActivityDate, not even an explicit nil
### GetConfiguration

`func (o *UserUpdatedMessageData) GetConfiguration() UserDtoConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *UserUpdatedMessageData) GetConfigurationOk() (*UserDtoConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *UserUpdatedMessageData) SetConfiguration(v UserDtoConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *UserUpdatedMessageData) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *UserUpdatedMessageData) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *UserUpdatedMessageData) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetPolicy

`func (o *UserUpdatedMessageData) GetPolicy() UserDtoPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UserUpdatedMessageData) GetPolicyOk() (*UserDtoPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *UserUpdatedMessageData) SetPolicy(v UserDtoPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *UserUpdatedMessageData) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *UserUpdatedMessageData) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *UserUpdatedMessageData) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetPrimaryImageAspectRatio

`func (o *UserUpdatedMessageData) GetPrimaryImageAspectRatio() float64`

GetPrimaryImageAspectRatio returns the PrimaryImageAspectRatio field if non-nil, zero value otherwise.

### GetPrimaryImageAspectRatioOk

`func (o *UserUpdatedMessageData) GetPrimaryImageAspectRatioOk() (*float64, bool)`

GetPrimaryImageAspectRatioOk returns a tuple with the PrimaryImageAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageAspectRatio

`func (o *UserUpdatedMessageData) SetPrimaryImageAspectRatio(v float64)`

SetPrimaryImageAspectRatio sets PrimaryImageAspectRatio field to given value.

### HasPrimaryImageAspectRatio

`func (o *UserUpdatedMessageData) HasPrimaryImageAspectRatio() bool`

HasPrimaryImageAspectRatio returns a boolean if a field has been set.

### SetPrimaryImageAspectRatioNil

`func (o *UserUpdatedMessageData) SetPrimaryImageAspectRatioNil(b bool)`

 SetPrimaryImageAspectRatioNil sets the value for PrimaryImageAspectRatio to be an explicit nil

### UnsetPrimaryImageAspectRatio
`func (o *UserUpdatedMessageData) UnsetPrimaryImageAspectRatio()`

UnsetPrimaryImageAspectRatio ensures that no value is present for PrimaryImageAspectRatio, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


