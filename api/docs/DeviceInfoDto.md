# DeviceInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**CustomName** | Pointer to **NullableString** | Gets or sets the custom name. | [optional] 
**AccessToken** | Pointer to **NullableString** | Gets or sets the access token. | [optional] 
**Id** | Pointer to **NullableString** | Gets or sets the identifier. | [optional] 
**LastUserName** | Pointer to **NullableString** | Gets or sets the last name of the user. | [optional] 
**AppName** | Pointer to **NullableString** | Gets or sets the name of the application. | [optional] 
**AppVersion** | Pointer to **NullableString** | Gets or sets the application version. | [optional] 
**LastUserId** | Pointer to **NullableString** | Gets or sets the last user identifier. | [optional] 
**DateLastActivity** | Pointer to **NullableTime** | Gets or sets the date last modified. | [optional] 
**Capabilities** | Pointer to [**ClientCapabilitiesDto**](ClientCapabilitiesDto.md) | Gets or sets the capabilities. | [optional] 
**IconUrl** | Pointer to **NullableString** | Gets or sets the icon URL. | [optional] 

## Methods

### NewDeviceInfoDto

`func NewDeviceInfoDto() *DeviceInfoDto`

NewDeviceInfoDto instantiates a new DeviceInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInfoDtoWithDefaults

`func NewDeviceInfoDtoWithDefaults() *DeviceInfoDto`

NewDeviceInfoDtoWithDefaults instantiates a new DeviceInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceInfoDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceInfoDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeviceInfoDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeviceInfoDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCustomName

`func (o *DeviceInfoDto) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *DeviceInfoDto) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *DeviceInfoDto) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *DeviceInfoDto) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *DeviceInfoDto) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *DeviceInfoDto) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
### GetAccessToken

`func (o *DeviceInfoDto) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *DeviceInfoDto) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *DeviceInfoDto) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *DeviceInfoDto) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *DeviceInfoDto) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *DeviceInfoDto) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetId

`func (o *DeviceInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DeviceInfoDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeviceInfoDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLastUserName

`func (o *DeviceInfoDto) GetLastUserName() string`

GetLastUserName returns the LastUserName field if non-nil, zero value otherwise.

### GetLastUserNameOk

`func (o *DeviceInfoDto) GetLastUserNameOk() (*string, bool)`

GetLastUserNameOk returns a tuple with the LastUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUserName

`func (o *DeviceInfoDto) SetLastUserName(v string)`

SetLastUserName sets LastUserName field to given value.

### HasLastUserName

`func (o *DeviceInfoDto) HasLastUserName() bool`

HasLastUserName returns a boolean if a field has been set.

### SetLastUserNameNil

`func (o *DeviceInfoDto) SetLastUserNameNil(b bool)`

 SetLastUserNameNil sets the value for LastUserName to be an explicit nil

### UnsetLastUserName
`func (o *DeviceInfoDto) UnsetLastUserName()`

UnsetLastUserName ensures that no value is present for LastUserName, not even an explicit nil
### GetAppName

`func (o *DeviceInfoDto) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *DeviceInfoDto) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *DeviceInfoDto) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *DeviceInfoDto) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *DeviceInfoDto) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *DeviceInfoDto) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil
### GetAppVersion

`func (o *DeviceInfoDto) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *DeviceInfoDto) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *DeviceInfoDto) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *DeviceInfoDto) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *DeviceInfoDto) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *DeviceInfoDto) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetLastUserId

`func (o *DeviceInfoDto) GetLastUserId() string`

GetLastUserId returns the LastUserId field if non-nil, zero value otherwise.

### GetLastUserIdOk

`func (o *DeviceInfoDto) GetLastUserIdOk() (*string, bool)`

GetLastUserIdOk returns a tuple with the LastUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUserId

`func (o *DeviceInfoDto) SetLastUserId(v string)`

SetLastUserId sets LastUserId field to given value.

### HasLastUserId

`func (o *DeviceInfoDto) HasLastUserId() bool`

HasLastUserId returns a boolean if a field has been set.

### SetLastUserIdNil

`func (o *DeviceInfoDto) SetLastUserIdNil(b bool)`

 SetLastUserIdNil sets the value for LastUserId to be an explicit nil

### UnsetLastUserId
`func (o *DeviceInfoDto) UnsetLastUserId()`

UnsetLastUserId ensures that no value is present for LastUserId, not even an explicit nil
### GetDateLastActivity

`func (o *DeviceInfoDto) GetDateLastActivity() time.Time`

GetDateLastActivity returns the DateLastActivity field if non-nil, zero value otherwise.

### GetDateLastActivityOk

`func (o *DeviceInfoDto) GetDateLastActivityOk() (*time.Time, bool)`

GetDateLastActivityOk returns a tuple with the DateLastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastActivity

`func (o *DeviceInfoDto) SetDateLastActivity(v time.Time)`

SetDateLastActivity sets DateLastActivity field to given value.

### HasDateLastActivity

`func (o *DeviceInfoDto) HasDateLastActivity() bool`

HasDateLastActivity returns a boolean if a field has been set.

### SetDateLastActivityNil

`func (o *DeviceInfoDto) SetDateLastActivityNil(b bool)`

 SetDateLastActivityNil sets the value for DateLastActivity to be an explicit nil

### UnsetDateLastActivity
`func (o *DeviceInfoDto) UnsetDateLastActivity()`

UnsetDateLastActivity ensures that no value is present for DateLastActivity, not even an explicit nil
### GetCapabilities

`func (o *DeviceInfoDto) GetCapabilities() ClientCapabilitiesDto`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *DeviceInfoDto) GetCapabilitiesOk() (*ClientCapabilitiesDto, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *DeviceInfoDto) SetCapabilities(v ClientCapabilitiesDto)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *DeviceInfoDto) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetIconUrl

`func (o *DeviceInfoDto) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *DeviceInfoDto) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *DeviceInfoDto) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *DeviceInfoDto) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *DeviceInfoDto) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *DeviceInfoDto) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


