# ClientCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayableMediaTypes** | Pointer to [**[]MediaType**](MediaType.md) |  | [optional] 
**SupportedCommands** | Pointer to [**[]GeneralCommandType**](GeneralCommandType.md) |  | [optional] 
**SupportsMediaControl** | Pointer to **bool** |  | [optional] 
**SupportsPersistentIdentifier** | Pointer to **bool** |  | [optional] 
**DeviceProfile** | Pointer to [**NullableDeviceProfile**](DeviceProfile.md) | A MediaBrowser.Model.Dlna.DeviceProfile represents a set of metadata which determines which content a certain device is able to play.  &lt;br /&gt;  Specifically, it defines the supported &lt;see cref&#x3D;\&quot;P:MediaBrowser.Model.Dlna.DeviceProfile.ContainerProfiles\&quot;&gt;containers&lt;/see&gt; and  &lt;see cref&#x3D;\&quot;P:MediaBrowser.Model.Dlna.DeviceProfile.CodecProfiles\&quot;&gt;codecs&lt;/see&gt; (video and/or audio, including codec profiles and levels)  the device is able to direct play (without transcoding or remuxing),  as well as which &lt;see cref&#x3D;\&quot;P:MediaBrowser.Model.Dlna.DeviceProfile.TranscodingProfiles\&quot;&gt;containers/codecs to transcode to&lt;/see&gt; in case it isn&#39;t. | [optional] 
**AppStoreUrl** | Pointer to **NullableString** |  | [optional] 
**IconUrl** | Pointer to **NullableString** |  | [optional] 
**SupportsContentUploading** | Pointer to **NullableBool** |  | [optional] [default to false]
**SupportsSync** | Pointer to **NullableBool** |  | [optional] [default to false]

## Methods

### NewClientCapabilities

`func NewClientCapabilities() *ClientCapabilities`

NewClientCapabilities instantiates a new ClientCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCapabilitiesWithDefaults

`func NewClientCapabilitiesWithDefaults() *ClientCapabilities`

NewClientCapabilitiesWithDefaults instantiates a new ClientCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayableMediaTypes

`func (o *ClientCapabilities) GetPlayableMediaTypes() []MediaType`

GetPlayableMediaTypes returns the PlayableMediaTypes field if non-nil, zero value otherwise.

### GetPlayableMediaTypesOk

`func (o *ClientCapabilities) GetPlayableMediaTypesOk() (*[]MediaType, bool)`

GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayableMediaTypes

`func (o *ClientCapabilities) SetPlayableMediaTypes(v []MediaType)`

SetPlayableMediaTypes sets PlayableMediaTypes field to given value.

### HasPlayableMediaTypes

`func (o *ClientCapabilities) HasPlayableMediaTypes() bool`

HasPlayableMediaTypes returns a boolean if a field has been set.

### SetPlayableMediaTypesNil

`func (o *ClientCapabilities) SetPlayableMediaTypesNil(b bool)`

 SetPlayableMediaTypesNil sets the value for PlayableMediaTypes to be an explicit nil

### UnsetPlayableMediaTypes
`func (o *ClientCapabilities) UnsetPlayableMediaTypes()`

UnsetPlayableMediaTypes ensures that no value is present for PlayableMediaTypes, not even an explicit nil
### GetSupportedCommands

`func (o *ClientCapabilities) GetSupportedCommands() []GeneralCommandType`

GetSupportedCommands returns the SupportedCommands field if non-nil, zero value otherwise.

### GetSupportedCommandsOk

`func (o *ClientCapabilities) GetSupportedCommandsOk() (*[]GeneralCommandType, bool)`

GetSupportedCommandsOk returns a tuple with the SupportedCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCommands

`func (o *ClientCapabilities) SetSupportedCommands(v []GeneralCommandType)`

SetSupportedCommands sets SupportedCommands field to given value.

### HasSupportedCommands

`func (o *ClientCapabilities) HasSupportedCommands() bool`

HasSupportedCommands returns a boolean if a field has been set.

### SetSupportedCommandsNil

`func (o *ClientCapabilities) SetSupportedCommandsNil(b bool)`

 SetSupportedCommandsNil sets the value for SupportedCommands to be an explicit nil

### UnsetSupportedCommands
`func (o *ClientCapabilities) UnsetSupportedCommands()`

UnsetSupportedCommands ensures that no value is present for SupportedCommands, not even an explicit nil
### GetSupportsMediaControl

`func (o *ClientCapabilities) GetSupportsMediaControl() bool`

GetSupportsMediaControl returns the SupportsMediaControl field if non-nil, zero value otherwise.

### GetSupportsMediaControlOk

`func (o *ClientCapabilities) GetSupportsMediaControlOk() (*bool, bool)`

GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMediaControl

`func (o *ClientCapabilities) SetSupportsMediaControl(v bool)`

SetSupportsMediaControl sets SupportsMediaControl field to given value.

### HasSupportsMediaControl

`func (o *ClientCapabilities) HasSupportsMediaControl() bool`

HasSupportsMediaControl returns a boolean if a field has been set.

### GetSupportsPersistentIdentifier

`func (o *ClientCapabilities) GetSupportsPersistentIdentifier() bool`

GetSupportsPersistentIdentifier returns the SupportsPersistentIdentifier field if non-nil, zero value otherwise.

### GetSupportsPersistentIdentifierOk

`func (o *ClientCapabilities) GetSupportsPersistentIdentifierOk() (*bool, bool)`

GetSupportsPersistentIdentifierOk returns a tuple with the SupportsPersistentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPersistentIdentifier

`func (o *ClientCapabilities) SetSupportsPersistentIdentifier(v bool)`

SetSupportsPersistentIdentifier sets SupportsPersistentIdentifier field to given value.

### HasSupportsPersistentIdentifier

`func (o *ClientCapabilities) HasSupportsPersistentIdentifier() bool`

HasSupportsPersistentIdentifier returns a boolean if a field has been set.

### GetDeviceProfile

`func (o *ClientCapabilities) GetDeviceProfile() DeviceProfile`

GetDeviceProfile returns the DeviceProfile field if non-nil, zero value otherwise.

### GetDeviceProfileOk

`func (o *ClientCapabilities) GetDeviceProfileOk() (*DeviceProfile, bool)`

GetDeviceProfileOk returns a tuple with the DeviceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceProfile

`func (o *ClientCapabilities) SetDeviceProfile(v DeviceProfile)`

SetDeviceProfile sets DeviceProfile field to given value.

### HasDeviceProfile

`func (o *ClientCapabilities) HasDeviceProfile() bool`

HasDeviceProfile returns a boolean if a field has been set.

### SetDeviceProfileNil

`func (o *ClientCapabilities) SetDeviceProfileNil(b bool)`

 SetDeviceProfileNil sets the value for DeviceProfile to be an explicit nil

### UnsetDeviceProfile
`func (o *ClientCapabilities) UnsetDeviceProfile()`

UnsetDeviceProfile ensures that no value is present for DeviceProfile, not even an explicit nil
### GetAppStoreUrl

`func (o *ClientCapabilities) GetAppStoreUrl() string`

GetAppStoreUrl returns the AppStoreUrl field if non-nil, zero value otherwise.

### GetAppStoreUrlOk

`func (o *ClientCapabilities) GetAppStoreUrlOk() (*string, bool)`

GetAppStoreUrlOk returns a tuple with the AppStoreUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppStoreUrl

`func (o *ClientCapabilities) SetAppStoreUrl(v string)`

SetAppStoreUrl sets AppStoreUrl field to given value.

### HasAppStoreUrl

`func (o *ClientCapabilities) HasAppStoreUrl() bool`

HasAppStoreUrl returns a boolean if a field has been set.

### SetAppStoreUrlNil

`func (o *ClientCapabilities) SetAppStoreUrlNil(b bool)`

 SetAppStoreUrlNil sets the value for AppStoreUrl to be an explicit nil

### UnsetAppStoreUrl
`func (o *ClientCapabilities) UnsetAppStoreUrl()`

UnsetAppStoreUrl ensures that no value is present for AppStoreUrl, not even an explicit nil
### GetIconUrl

`func (o *ClientCapabilities) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ClientCapabilities) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ClientCapabilities) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ClientCapabilities) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *ClientCapabilities) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *ClientCapabilities) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetSupportsContentUploading

`func (o *ClientCapabilities) GetSupportsContentUploading() bool`

GetSupportsContentUploading returns the SupportsContentUploading field if non-nil, zero value otherwise.

### GetSupportsContentUploadingOk

`func (o *ClientCapabilities) GetSupportsContentUploadingOk() (*bool, bool)`

GetSupportsContentUploadingOk returns a tuple with the SupportsContentUploading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsContentUploading

`func (o *ClientCapabilities) SetSupportsContentUploading(v bool)`

SetSupportsContentUploading sets SupportsContentUploading field to given value.

### HasSupportsContentUploading

`func (o *ClientCapabilities) HasSupportsContentUploading() bool`

HasSupportsContentUploading returns a boolean if a field has been set.

### SetSupportsContentUploadingNil

`func (o *ClientCapabilities) SetSupportsContentUploadingNil(b bool)`

 SetSupportsContentUploadingNil sets the value for SupportsContentUploading to be an explicit nil

### UnsetSupportsContentUploading
`func (o *ClientCapabilities) UnsetSupportsContentUploading()`

UnsetSupportsContentUploading ensures that no value is present for SupportsContentUploading, not even an explicit nil
### GetSupportsSync

`func (o *ClientCapabilities) GetSupportsSync() bool`

GetSupportsSync returns the SupportsSync field if non-nil, zero value otherwise.

### GetSupportsSyncOk

`func (o *ClientCapabilities) GetSupportsSyncOk() (*bool, bool)`

GetSupportsSyncOk returns a tuple with the SupportsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSync

`func (o *ClientCapabilities) SetSupportsSync(v bool)`

SetSupportsSync sets SupportsSync field to given value.

### HasSupportsSync

`func (o *ClientCapabilities) HasSupportsSync() bool`

HasSupportsSync returns a boolean if a field has been set.

### SetSupportsSyncNil

`func (o *ClientCapabilities) SetSupportsSyncNil(b bool)`

 SetSupportsSyncNil sets the value for SupportsSync to be an explicit nil

### UnsetSupportsSync
`func (o *ClientCapabilities) UnsetSupportsSync()`

UnsetSupportsSync ensures that no value is present for SupportsSync, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


