# ClientCapabilitiesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayableMediaTypes** | Pointer to [**[]MediaType**](MediaType.md) | Gets or sets the list of playable media types. | [optional] 
**SupportedCommands** | Pointer to [**[]GeneralCommandType**](GeneralCommandType.md) | Gets or sets the list of supported commands. | [optional] 
**SupportsMediaControl** | Pointer to **bool** | Gets or sets a value indicating whether session supports media control. | [optional] 
**SupportsPersistentIdentifier** | Pointer to **bool** | Gets or sets a value indicating whether session supports a persistent identifier. | [optional] 
**DeviceProfile** | Pointer to [**NullableClientCapabilitiesDtoDeviceProfile**](ClientCapabilitiesDtoDeviceProfile.md) |  | [optional] 
**AppStoreUrl** | Pointer to **NullableString** | Gets or sets the app store url. | [optional] 
**IconUrl** | Pointer to **NullableString** | Gets or sets the icon url. | [optional] 
**SupportsContentUploading** | Pointer to **NullableBool** |  | [optional] [default to false]
**SupportsSync** | Pointer to **NullableBool** |  | [optional] [default to false]

## Methods

### NewClientCapabilitiesDto

`func NewClientCapabilitiesDto() *ClientCapabilitiesDto`

NewClientCapabilitiesDto instantiates a new ClientCapabilitiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCapabilitiesDtoWithDefaults

`func NewClientCapabilitiesDtoWithDefaults() *ClientCapabilitiesDto`

NewClientCapabilitiesDtoWithDefaults instantiates a new ClientCapabilitiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayableMediaTypes

`func (o *ClientCapabilitiesDto) GetPlayableMediaTypes() []MediaType`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *ClientCapabilitiesDto) GetPlayableMediaTypesOk() (*[]MediaType, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *ClientCapabilitiesDto) SetPlayableMediaTypes(v []MediaType)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *ClientCapabilitiesDto) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### GetSupportedCommands

`func (o *ClientCapabilitiesDto) GetSupportedCommands() []GeneralCommandType`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *ClientCapabilitiesDto) GetSupportedCommandsOk() (*[]GeneralCommandType, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *ClientCapabilitiesDto) SetSupportedCommands(v []GeneralCommandType)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *ClientCapabilitiesDto) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.

### GetSupportsMediaControl

`func (o *ClientCapabilitiesDto) GetSupportsMediaControl() bool`

GetSupportsMediaControl returns the SupportsMediaControl field if non-nil, zero value otherwise.

### GetSupportsMediaControlOk

`func (o *ClientCapabilitiesDto) GetSupportsMediaControlOk() (*bool, bool)`

GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMediaControl

`func (o *ClientCapabilitiesDto) SetSupportsMediaControl(v bool)`

SetSupportsMediaControl sets SupportsMediaControl field to given value.

### HasSupportsMediaControl

`func (o *ClientCapabilitiesDto) HasSupportsMediaControl() bool`

HasSupportsMediaControl returns a boolean if a field has been set.

### GetSupportsPersistentIdentifier

`func (o *ClientCapabilitiesDto) GetSupportsPersistentIdentifier() bool`

GetSupportsPersistentIdentifier returns the SupportsPersistentIdentifier field if non-nil, zero value otherwise.

### GetSupportsPersistentIdentifierOk

`func (o *ClientCapabilitiesDto) GetSupportsPersistentIdentifierOk() (*bool, bool)`

GetSupportsPersistentIdentifierOk returns a tuple with the SupportsPersistentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPersistentIdentifier

`func (o *ClientCapabilitiesDto) SetSupportsPersistentIdentifier(v bool)`

SetSupportsPersistentIdentifier sets SupportsPersistentIdentifier field to given value.

### HasSupportsPersistentIdentifier

`func (o *ClientCapabilitiesDto) HasSupportsPersistentIdentifier() bool`

HasSupportsPersistentIdentifier returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *ClientCapabilitiesDto) GetDeviceProfile() ClientCapabilitiesDtoDeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *ClientCapabilitiesDto) GetDeviceProfileOk() (*ClientCapabilitiesDtoDeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *ClientCapabilitiesDto) SetDeviceProfile(v ClientCapabilitiesDtoDeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *ClientCapabilitiesDto) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### SetDeviceProfileNil

`func (o *ClientCapabilitiesDto) SetDeviceProfileNil(b bool)`

 SetDeviceProfileNil sets the value for DeviceProfile to be an explicit nil

### UnsetDeviceProfile
`func (o *ClientCapabilitiesDto) UnsetDeviceProfile()`

UnsetDeviceProfile ensures that no value is present for DeviceProfile, not even an explicit nil
### GetAppStoreUrl

`func (o *ClientCapabilitiesDto) GetAppStoreUrl() string`

GetAppStoreUrl returns the AppStoreUrl field if non-nil, zero value otherwise.

### GetAppStoreUrlOk

`func (o *ClientCapabilitiesDto) GetAppStoreUrlOk() (*string, bool)`

GetAppStoreUrlOk returns a tuple with the AppStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreUrl

`func (o *ClientCapabilitiesDto) SetAppStoreUrl(v string)`

SetAppStoreUrl sets AppStoreUrl field to given value.

### HasAppStoreUrl

`func (o *ClientCapabilitiesDto) HasAppStoreUrl() bool`

HasAppStoreUrl returns a boolean if a field has been set.

### SetAppStoreUrlNil

`func (o *ClientCapabilitiesDto) SetAppStoreUrlNil(b bool)`

 SetAppStoreUrlNil sets the value for AppStoreUrl to be an explicit nil

### UnsetAppStoreUrl
`func (o *ClientCapabilitiesDto) UnsetAppStoreUrl()`

UnsetAppStoreUrl ensures that no value is present for AppStoreUrl, not even an explicit nil
### GetIconUrl

`func (o *ClientCapabilitiesDto) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ClientCapabilitiesDto) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ClientCapabilitiesDto) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ClientCapabilitiesDto) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *ClientCapabilitiesDto) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *ClientCapabilitiesDto) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetSupportsContentUploading

`func (o *ClientCapabilitiesDto) GetSupportsContentUploading() bool`

GetSupportsContentUploading returns the SupportsContentUploading field if non-nil, zero value otherwise.

### GetSupportsContentUploadingOk

`func (o *ClientCapabilitiesDto) GetSupportsContentUploadingOk() (*bool, bool)`

GetSupportsContentUploadingOk returns a tuple with the SupportsContentUploading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsContentUploading

`func (o *ClientCapabilitiesDto) SetSupportsContentUploading(v bool)`

SetSupportsContentUploading sets SupportsContentUploading field to given value.

### HasSupportsContentUploading

`func (o *ClientCapabilitiesDto) HasSupportsContentUploading() bool`

HasSupportsContentUploading returns a boolean if a field has been set.

### SetSupportsContentUploadingNil

`func (o *ClientCapabilitiesDto) SetSupportsContentUploadingNil(b bool)`

 SetSupportsContentUploadingNil sets the value for SupportsContentUploading to be an explicit nil

### UnsetSupportsContentUploading
`func (o *ClientCapabilitiesDto) UnsetSupportsContentUploading()`

UnsetSupportsContentUploading ensures that no value is present for SupportsContentUploading, not even an explicit nil
### GetSupportsSync

`func (o *ClientCapabilitiesDto) GetSupportsSync() bool`

GetSupportsSync returns the SupportsSync field if non-nil, zero value otherwise.

### GetSupportsSyncOk

`func (o *ClientCapabilitiesDto) GetSupportsSyncOk() (*bool, bool)`

GetSupportsSyncOk returns a tuple with the SupportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSync

`func (o *ClientCapabilitiesDto) SetSupportsSync(v bool)`

SetSupportsSync sets SupportsSync field to given value.

### HasSupportsSync

`func (o *ClientCapabilitiesDto) HasSupportsSync() bool`

HasSupportsSync returns a boolean if a field has been set.

### SetSupportsSyncNil

`func (o *ClientCapabilitiesDto) SetSupportsSyncNil(b bool)`

 SetSupportsSyncNil sets the value for SupportsSync to be an explicit nil

### UnsetSupportsSync
`func (o *ClientCapabilitiesDto) UnsetSupportsSync()`

UnsetSupportsSync ensures that no value is present for SupportsSync, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


