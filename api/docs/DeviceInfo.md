# DeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**AccessToken** | Pointer to **NullableString** | Gets or sets the access token. | [optional] 
**Id** | Pointer to **NullableString** | Gets or sets the identifier. | [optional] 
**LastUserName** | Pointer to **NullableString** | Gets or sets the last name of the user. | [optional] 
**AppName** | Pointer to **NullableString** | Gets or sets the name of the application. | [optional] 
**AppVersion** | Pointer to **NullableString** | Gets or sets the application version. | [optional] 
**LastUserId** | Pointer to **string** | Gets or sets the last user identifier. | [optional] 
**DateLastActivity** | Pointer to **time.Time** | Gets or sets the date last modified. | [optional] 
**Capabilities** | Pointer to [**NullableDeviceInfoCapabilities**](DeviceInfoCapabilities.md) |  | [optional] 
**IconUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDeviceInfo

`func NewDeviceInfo() *DeviceInfo`

NewDeviceInfo instantiates a new DeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInfoWithDefaults

`func NewDeviceInfoWithDefaults() *DeviceInfo`

NewDeviceInfoWithDefaults instantiates a new DeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeviceInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeviceInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAccessToken

`func (o *DeviceInfo) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *DeviceInfo) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *DeviceInfo) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *DeviceInfo) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *DeviceInfo) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *DeviceInfo) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetId

`func (o *DeviceInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DeviceInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeviceInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLastUserName

`func (o *DeviceInfo) GetLastUserName() string`

GetLastUserName returns the LastUserName field if non-nil, zero value otherwise.

### GetLastUserNameOk

`func (o *DeviceInfo) GetLastUserNameOk() (*string, bool)`

GetLastUserNameOk returns a tuple with the LastUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUserName

`func (o *DeviceInfo) SetLastUserName(v string)`

SetLastUserName sets LastUserName field to given value.

### HasLastUserName

`func (o *DeviceInfo) HasLastUserName() bool`

HasLastUserName returns a boolean if a field has been set.

### SetLastUserNameNil

`func (o *DeviceInfo) SetLastUserNameNil(b bool)`

 SetLastUserNameNil sets the value for LastUserName to be an explicit nil

### UnsetLastUserName
`func (o *DeviceInfo) UnsetLastUserName()`

UnsetLastUserName ensures that no value is present for LastUserName, not even an explicit nil
### GetAppName

`func (o *DeviceInfo) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *DeviceInfo) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *DeviceInfo) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *DeviceInfo) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### SetAppNameNil

`func (o *DeviceInfo) SetAppNameNil(b bool)`

 SetAppNameNil sets the value for AppName to be an explicit nil

### UnsetAppName
`func (o *DeviceInfo) UnsetAppName()`

UnsetAppName ensures that no value is present for AppName, not even an explicit nil
### GetAppVersion

`func (o *DeviceInfo) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *DeviceInfo) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *DeviceInfo) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.

### HasAppVersion

`func (o *DeviceInfo) HasAppVersion() bool`

HasAppVersion returns a boolean if a field has been set.

### SetAppVersionNil

`func (o *DeviceInfo) SetAppVersionNil(b bool)`

 SetAppVersionNil sets the value for AppVersion to be an explicit nil

### UnsetAppVersion
`func (o *DeviceInfo) UnsetAppVersion()`

UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
### GetLastUserId

`func (o *DeviceInfo) GetLastUserId() string`

GetLastUserId returns the LastUserId field if non-nil, zero value otherwise.

### GetLastUserIdOk

`func (o *DeviceInfo) GetLastUserIdOk() (*string, bool)`

GetLastUserIdOk returns a tuple with the LastUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUserId

`func (o *DeviceInfo) SetLastUserId(v string)`

SetLastUserId sets LastUserId field to given value.

### HasLastUserId

`func (o *DeviceInfo) HasLastUserId() bool`

HasLastUserId returns a boolean if a field has been set.

### GetDateLastActivity

`func (o *DeviceInfo) GetDateLastActivity() time.Time`

GetDateLastActivity returns the DateLastActivity field if non-nil, zero value otherwise.

### GetDateLastActivityOk

`func (o *DeviceInfo) GetDateLastActivityOk() (*time.Time, bool)`

GetDateLastActivityOk returns a tuple with the DateLastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastActivity

`func (o *DeviceInfo) SetDateLastActivity(v time.Time)`

SetDateLastActivity sets DateLastActivity field to given value.

### HasDateLastActivity

`func (o *DeviceInfo) HasDateLastActivity() bool`

HasDateLastActivity returns a boolean if a field has been set.

### GetCapabilities

`func (o *DeviceInfo) GetCapabilities() DeviceInfoCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *DeviceInfo) GetCapabilitiesOk() (*DeviceInfoCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *DeviceInfo) SetCapabilities(v DeviceInfoCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *DeviceInfo) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *DeviceInfo) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *DeviceInfo) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetIconUrl

`func (o *DeviceInfo) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *DeviceInfo) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *DeviceInfo) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *DeviceInfo) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *DeviceInfo) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *DeviceInfo) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


