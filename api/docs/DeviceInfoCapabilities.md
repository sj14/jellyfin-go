# DeviceInfoCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayableMediaTypes** | Pointer to **[]string** |  | [optional] 
**SupportedCommands** | Pointer to [**[]GeneralCommandType**](GeneralCommandType.md) |  | [optional] 
**SupportsMediaControl** | Pointer to **bool** |  | [optional] 
**SupportsContentUploading** | Pointer to **bool** |  | [optional] 
**MessageCallbackUrl** | Pointer to **NullableString** |  | [optional] 
**SupportsPersistentIdentifier** | Pointer to **bool** |  | [optional] 
**SupportsSync** | Pointer to **bool** |  | [optional] 
**DeviceProfile** | Pointer to [**NullableClientCapabilitiesDeviceProfile**](ClientCapabilitiesDeviceProfile.md) |  | [optional] 
**AppStoreUrl** | Pointer to **NullableString** |  | [optional] 
**IconUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDeviceInfoCapabilities

`func NewDeviceInfoCapabilities() *DeviceInfoCapabilities`

NewDeviceInfoCapabilities instantiates a new DeviceInfoCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInfoCapabilitiesWithDefaults

`func NewDeviceInfoCapabilitiesWithDefaults() *DeviceInfoCapabilities`

NewDeviceInfoCapabilitiesWithDefaults instantiates a new DeviceInfoCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayableMediaTypes

`func (o *DeviceInfoCapabilities) GetPlayableMediaTypes() []string`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *DeviceInfoCapabilities) GetPlayableMediaTypesOk() (*[]string, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *DeviceInfoCapabilities) SetPlayableMediaTypes(v []string)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *DeviceInfoCapabilities) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### SetPlayableMediaTypesNil

`func (o *DeviceInfoCapabilities) SetPlayableMediaTypesNil(b bool)`

 SetPlayableMediaTypesNil sets the value for PlayableMediaTypes to be an explicit nil

### UnsetPlayableMediaTypes
`func (o *DeviceInfoCapabilities) UnsetPlayableMediaTypes()`

UnsetPlayableMediaTypes ensures that no value is present for PlayableMediaTypes, not even an explicit nil
### GetSupportedCommands

`func (o *DeviceInfoCapabilities) GetSupportedCommands() []GeneralCommandType`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *DeviceInfoCapabilities) GetSupportedCommandsOk() (*[]GeneralCommandType, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *DeviceInfoCapabilities) SetSupportedCommands(v []GeneralCommandType)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *DeviceInfoCapabilities) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.

### SetSupportedCommandsNil

`func (o *DeviceInfoCapabilities) SetSupportedCommandsNil(b bool)`

 SetSupportedCommandsNil sets the value for SupportedCommands to be an explicit nil

### UnsetSupportedCommands
`func (o *DeviceInfoCapabilities) UnsetSupportedCommands()`

UnsetSupportedCommands ensures that no value is present for SupportedCommands, not even an explicit nil
### GetSupportsMediaControl

`func (o *DeviceInfoCapabilities) GetSupportsMediaControl() bool`

GetSupportsMediaControl returns the SupportsMediaControl field if non-nil, zero value otherwise.

### GetSupportsMediaControlOk

`func (o *DeviceInfoCapabilities) GetSupportsMediaControlOk() (*bool, bool)`

GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMediaControl

`func (o *DeviceInfoCapabilities) SetSupportsMediaControl(v bool)`

SetSupportsMediaControl sets SupportsMediaControl field to given value.

### HasSupportsMediaControl

`func (o *DeviceInfoCapabilities) HasSupportsMediaControl() bool`

HasSupportsMediaControl returns a boolean if a field has been set.

### GetSupportsContentUploading

`func (o *DeviceInfoCapabilities) GetSupportsContentUploading() bool`

GetSupportsContentUploading returns the SupportsContentUploading field if non-nil, zero value otherwise.

### GetSupportsContentUploadingOk

`func (o *DeviceInfoCapabilities) GetSupportsContentUploadingOk() (*bool, bool)`

GetSupportsContentUploadingOk returns a tuple with the SupportsContentUploading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsContentUploading

`func (o *DeviceInfoCapabilities) SetSupportsContentUploading(v bool)`

SetSupportsContentUploading sets SupportsContentUploading field to given value.

### HasSupportsContentUploading

`func (o *DeviceInfoCapabilities) HasSupportsContentUploading() bool`

HasSupportsContentUploading returns a boolean if a field has been set.

### GetMessageCallbackUrl

`func (o *DeviceInfoCapabilities) GetMessageCallbackUrl() string`

GetMessageCallbackUrl returns the MessageCallbackUrl field if non-nil, zero value otherwise.

### GetMessageCallbackUrlOk

`func (o *DeviceInfoCapabilities) GetMessageCallbackUrlOk() (*string, bool)`

GetMessageCallbackUrlOk returns a tuple with the MessageCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCallbackUrl

`func (o *DeviceInfoCapabilities) SetMessageCallbackUrl(v string)`

SetMessageCallbackUrl sets MessageCallbackUrl field to given value.

### HasMessageCallbackUrl

`func (o *DeviceInfoCapabilities) HasMessageCallbackUrl() bool`

HasMessageCallbackUrl returns a boolean if a field has been set.

### SetMessageCallbackUrlNil

`func (o *DeviceInfoCapabilities) SetMessageCallbackUrlNil(b bool)`

 SetMessageCallbackUrlNil sets the value for MessageCallbackUrl to be an explicit nil

### UnsetMessageCallbackUrl
`func (o *DeviceInfoCapabilities) UnsetMessageCallbackUrl()`

UnsetMessageCallbackUrl ensures that no value is present for MessageCallbackUrl, not even an explicit nil
### GetSupportsPersistentIdentifier

`func (o *DeviceInfoCapabilities) GetSupportsPersistentIdentifier() bool`

GetSupportsPersistentIdentifier returns the SupportsPersistentIdentifier field if non-nil, zero value otherwise.

### GetSupportsPersistentIdentifierOk

`func (o *DeviceInfoCapabilities) GetSupportsPersistentIdentifierOk() (*bool, bool)`

GetSupportsPersistentIdentifierOk returns a tuple with the SupportsPersistentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPersistentIdentifier

`func (o *DeviceInfoCapabilities) SetSupportsPersistentIdentifier(v bool)`

SetSupportsPersistentIdentifier sets SupportsPersistentIdentifier field to given value.

### HasSupportsPersistentIdentifier

`func (o *DeviceInfoCapabilities) HasSupportsPersistentIdentifier() bool`

HasSupportsPersistentIdentifier returns a boolean if a field has been set.

### GetSupportsSync

`func (o *DeviceInfoCapabilities) GetSupportsSync() bool`

GetSupportsSync returns the SupportsSync field if non-nil, zero value otherwise.

### GetSupportsSyncOk

`func (o *DeviceInfoCapabilities) GetSupportsSyncOk() (*bool, bool)`

GetSupportsSyncOk returns a tuple with the SupportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSync

`func (o *DeviceInfoCapabilities) SetSupportsSync(v bool)`

SetSupportsSync sets SupportsSync field to given value.

### HasSupportsSync

`func (o *DeviceInfoCapabilities) HasSupportsSync() bool`

HasSupportsSync returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *DeviceInfoCapabilities) GetDeviceProfile() ClientCapabilitiesDeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *DeviceInfoCapabilities) GetDeviceProfileOk() (*ClientCapabilitiesDeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *DeviceInfoCapabilities) SetDeviceProfile(v ClientCapabilitiesDeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *DeviceInfoCapabilities) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### SetDeviceProfileNil

`func (o *DeviceInfoCapabilities) SetDeviceProfileNil(b bool)`

 SetDeviceProfileNil sets the value for DeviceProfile to be an explicit nil

### UnsetDeviceProfile
`func (o *DeviceInfoCapabilities) UnsetDeviceProfile()`

UnsetDeviceProfile ensures that no value is present for DeviceProfile, not even an explicit nil
### GetAppStoreUrl

`func (o *DeviceInfoCapabilities) GetAppStoreUrl() string`

GetAppStoreUrl returns the AppStoreUrl field if non-nil, zero value otherwise.

### GetAppStoreUrlOk

`func (o *DeviceInfoCapabilities) GetAppStoreUrlOk() (*string, bool)`

GetAppStoreUrlOk returns a tuple with the AppStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreUrl

`func (o *DeviceInfoCapabilities) SetAppStoreUrl(v string)`

SetAppStoreUrl sets AppStoreUrl field to given value.

### HasAppStoreUrl

`func (o *DeviceInfoCapabilities) HasAppStoreUrl() bool`

HasAppStoreUrl returns a boolean if a field has been set.

### SetAppStoreUrlNil

`func (o *DeviceInfoCapabilities) SetAppStoreUrlNil(b bool)`

 SetAppStoreUrlNil sets the value for AppStoreUrl to be an explicit nil

### UnsetAppStoreUrl
`func (o *DeviceInfoCapabilities) UnsetAppStoreUrl()`

UnsetAppStoreUrl ensures that no value is present for AppStoreUrl, not even an explicit nil
### GetIconUrl

`func (o *DeviceInfoCapabilities) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *DeviceInfoCapabilities) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *DeviceInfoCapabilities) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *DeviceInfoCapabilities) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *DeviceInfoCapabilities) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *DeviceInfoCapabilities) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


