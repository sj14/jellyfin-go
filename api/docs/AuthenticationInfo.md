# AuthenticationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Gets or sets the identifier. | [optional] 
**AccessToken** | Pointer to **NullableString** | Gets or sets the access token. | [optional] 
**DeviceId** | Pointer to **NullableString** | Gets or sets the device identifier. | [optional] 
**AppName** | Pointer to **NullableString** | Gets or sets the name of the application. | [optional] 
**AppVersion** | Pointer to **NullableString** | Gets or sets the application version. | [optional] 
**DeviceName** | Pointer to **NullableString** | Gets or sets the name of the device. | [optional] 
**UserId** | Pointer to **string** | Gets or sets the user identifier. | [optional] 
**IsActive** | Pointer to **bool** | Gets or sets a value indicating whether this instance is active. | [optional] 
**DateCreated** | Pointer to **time.Time** | Gets or sets the date created. | [optional] 
**DateRevoked** | Pointer to **NullableTime** | Gets or sets the date revoked. | [optional] 
**DateLastActivity** | Pointer to **time.Time** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAuthenticationInfo

`func NewAuthenticationInfo() *AuthenticationInfo`

NewAuthenticationInfo instantiates a new AuthenticationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationInfoWithDefaults

`func NewAuthenticationInfoWithDefaults() *AuthenticationInfo`

NewAuthenticationInfoWithDefaults instantiates a new AuthenticationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthenticationInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthenticationInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthenticationInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AuthenticationInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccessToken

`func (o *AuthenticationInfo) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AuthenticationInfo) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AuthenticationInfo) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *AuthenticationInfo) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *AuthenticationInfo) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *AuthenticationInfo) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetDeviceId

`func (o *AuthenticationInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AuthenticationInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AuthenticationInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *AuthenticationInfo) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *AuthenticationInfo) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *AuthenticationInfo) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetAppName

`func (o *AuthenticationInfo) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *AuthenticationInfo) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *AuthenticationInfo) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *AuthenticationInfo) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *AuthenticationInfo) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *AuthenticationInfo) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil
### GetAppVersion

`func (o *AuthenticationInfo) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *AuthenticationInfo) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *AuthenticationInfo) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *AuthenticationInfo) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *AuthenticationInfo) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *AuthenticationInfo) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetDeviceName

`func (o *AuthenticationInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *AuthenticationInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *AuthenticationInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *AuthenticationInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### SetDeviceNameNil

`func (o *AuthenticationInfo) SetDeviceNameNil(b bool)`

 SetDeviceNameNil sets the value for DeviceName to be an explicit nil

### UnsetDeviceName
`func (o *AuthenticationInfo) UnsetDeviceName()`

UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
### GetUserId

`func (o *AuthenticationInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuthenticationInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuthenticationInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuthenticationInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetIsActive

`func (o *AuthenticationInfo) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AuthenticationInfo) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AuthenticationInfo) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *AuthenticationInfo) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDateCreated

`func (o *AuthenticationInfo) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AuthenticationInfo) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AuthenticationInfo) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AuthenticationInfo) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateRevoked

`func (o *AuthenticationInfo) GetDateRevoked() time.Time`

GetDateRevoked returns the DateRevoked field if non-nil, zero value otherwise.

### GetDateRevokedOk

`func (o *AuthenticationInfo) GetDateRevokedOk() (*time.Time, bool)`

GetDateRevokedOk returns a tuple with the DateRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRevoked

`func (o *AuthenticationInfo) SetDateRevoked(v time.Time)`

SetDateRevoked sets DateRevoked field to given value.

### HasDateRevoked

`func (o *AuthenticationInfo) HasDateRevoked() bool`

HasDateRevoked returns a boolean if a field has been set.

### SetDateRevokedNil

`func (o *AuthenticationInfo) SetDateRevokedNil(b bool)`

 SetDateRevokedNil sets the value for DateRevoked to be an explicit nil

### UnsetDateRevoked
`func (o *AuthenticationInfo) UnsetDateRevoked()`

UnsetDateRevoked ensures that no value is present for DateRevoked, not even an explicit nil
### GetDateLastActivity

`func (o *AuthenticationInfo) GetDateLastActivity() time.Time`

GetDateLastActivity returns the DateLastActivity field if non-nil, zero value otherwise.

### GetDateLastActivityOk

`func (o *AuthenticationInfo) GetDateLastActivityOk() (*time.Time, bool)`

GetDateLastActivityOk returns a tuple with the DateLastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastActivity

`func (o *AuthenticationInfo) SetDateLastActivity(v time.Time)`

SetDateLastActivity sets DateLastActivity field to given value.

### HasDateLastActivity

`func (o *AuthenticationInfo) HasDateLastActivity() bool`

HasDateLastActivity returns a boolean if a field has been set.

### GetUserName

`func (o *AuthenticationInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *AuthenticationInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *AuthenticationInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *AuthenticationInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *AuthenticationInfo) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *AuthenticationInfo) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


