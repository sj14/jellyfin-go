# ClientCapabilitiesDeviceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name of this device profile. | [optional] 
**Id** | Pointer to **NullableString** | Gets or sets the Id. | [optional] 
**Identification** | Pointer to [**NullableDeviceProfileIdentification**](DeviceProfileIdentification.md) |  | [optional] 
**FriendlyName** | Pointer to **NullableString** | Gets or sets the friendly name of the device profile, which can be shown to users. | [optional] 
**Manufacturer** | Pointer to **NullableString** | Gets or sets the manufacturer of the device which this profile represents. | [optional] 
**ManufacturerUrl** | Pointer to **NullableString** | Gets or sets an url for the manufacturer of the device which this profile represents. | [optional] 
**ModelName** | Pointer to **NullableString** | Gets or sets the model name of the device which this profile represents. | [optional] 
**ModelDescription** | Pointer to **NullableString** | Gets or sets the model description of the device which this profile represents. | [optional] 
**ModelNumber** | Pointer to **NullableString** | Gets or sets the model number of the device which this profile represents. | [optional] 
**ModelUrl** | Pointer to **NullableString** | Gets or sets the ModelUrl. | [optional] 
**SerialNumber** | Pointer to **NullableString** | Gets or sets the serial number of the device which this profile represents. | [optional] 
**EnableAlbumArtInDidl** | Pointer to **bool** | Gets or sets a value indicating whether EnableAlbumArtInDidl. | [optional] [default to false]
**EnableSingleAlbumArtLimit** | Pointer to **bool** | Gets or sets a value indicating whether EnableSingleAlbumArtLimit. | [optional] [default to false]
**EnableSingleSubtitleLimit** | Pointer to **bool** | Gets or sets a value indicating whether EnableSingleSubtitleLimit. | [optional] [default to false]
**SupportedMediaTypes** | Pointer to **string** | Gets or sets the SupportedMediaTypes. | [optional] 
**UserId** | Pointer to **NullableString** | Gets or sets the UserId. | [optional] 
**AlbumArtPn** | Pointer to **NullableString** | Gets or sets the AlbumArtPn. | [optional] 
**MaxAlbumArtWidth** | Pointer to **NullableInt32** | Gets or sets the MaxAlbumArtWidth. | [optional] 
**MaxAlbumArtHeight** | Pointer to **NullableInt32** | Gets or sets the MaxAlbumArtHeight. | [optional] 
**MaxIconWidth** | Pointer to **NullableInt32** | Gets or sets the maximum allowed width of embedded icons. | [optional] 
**MaxIconHeight** | Pointer to **NullableInt32** | Gets or sets the maximum allowed height of embedded icons. | [optional] 
**MaxStreamingBitrate** | Pointer to **NullableInt32** | Gets or sets the maximum allowed bitrate for all streamed content. | [optional] 
**MaxStaticBitrate** | Pointer to **NullableInt32** | Gets or sets the maximum allowed bitrate for statically streamed content (&#x3D; direct played files). | [optional] 
**MusicStreamingTranscodingBitrate** | Pointer to **NullableInt32** | Gets or sets the maximum allowed bitrate for transcoded music streams. | [optional] 
**MaxStaticMusicBitrate** | Pointer to **NullableInt32** | Gets or sets the maximum allowed bitrate for statically streamed (&#x3D; direct played) music files. | [optional] 
**SonyAggregationFlags** | Pointer to **NullableString** | Gets or sets the content of the aggregationFlags element in the urn:schemas-sonycom:av namespace. | [optional] 
**ProtocolInfo** | Pointer to **NullableString** | Gets or sets the ProtocolInfo. | [optional] 
**TimelineOffsetSeconds** | Pointer to **int32** | Gets or sets the TimelineOffsetSeconds. | [optional] [default to 0]
**RequiresPlainVideoItems** | Pointer to **bool** | Gets or sets a value indicating whether RequiresPlainVideoItems. | [optional] [default to false]
**RequiresPlainFolders** | Pointer to **bool** | Gets or sets a value indicating whether RequiresPlainFolders. | [optional] [default to false]
**EnableMSMediaReceiverRegistrar** | Pointer to **bool** | Gets or sets a value indicating whether EnableMSMediaReceiverRegistrar. | [optional] [default to false]
**IgnoreTranscodeByteRangeRequests** | Pointer to **bool** | Gets or sets a value indicating whether IgnoreTranscodeByteRangeRequests. | [optional] [default to false]
**XmlRootAttributes** | Pointer to [**[]XmlAttribute**](XmlAttribute.md) | Gets or sets the XmlRootAttributes. | [optional] 
**DirectPlayProfiles** | Pointer to [**[]DirectPlayProfile**](DirectPlayProfile.md) | Gets or sets the direct play profiles. | [optional] 
**TranscodingProfiles** | Pointer to [**[]TranscodingProfile**](TranscodingProfile.md) | Gets or sets the transcoding profiles. | [optional] 
**ContainerProfiles** | Pointer to [**[]ContainerProfile**](ContainerProfile.md) | Gets or sets the container profiles. | [optional] 
**CodecProfiles** | Pointer to [**[]CodecProfile**](CodecProfile.md) | Gets or sets the codec profiles. | [optional] 
**ResponseProfiles** | Pointer to [**[]ResponseProfile**](ResponseProfile.md) | Gets or sets the ResponseProfiles. | [optional] 
**SubtitleProfiles** | Pointer to [**[]SubtitleProfile**](SubtitleProfile.md) | Gets or sets the subtitle profiles. | [optional] 

## Methods

### NewClientCapabilitiesDeviceProfile

`func NewClientCapabilitiesDeviceProfile() *ClientCapabilitiesDeviceProfile`

NewClientCapabilitiesDeviceProfile instantiates a new ClientCapabilitiesDeviceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCapabilitiesDeviceProfileWithDefaults

`func NewClientCapabilitiesDeviceProfileWithDefaults() *ClientCapabilitiesDeviceProfile`

NewClientCapabilitiesDeviceProfileWithDefaults instantiates a new ClientCapabilitiesDeviceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientCapabilitiesDeviceProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientCapabilitiesDeviceProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientCapabilitiesDeviceProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientCapabilitiesDeviceProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ClientCapabilitiesDeviceProfile) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ClientCapabilitiesDeviceProfile) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *ClientCapabilitiesDeviceProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientCapabilitiesDeviceProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientCapabilitiesDeviceProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientCapabilitiesDeviceProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ClientCapabilitiesDeviceProfile) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ClientCapabilitiesDeviceProfile) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentification

`func (o *ClientCapabilitiesDeviceProfile) GetIdentification() DeviceProfileIdentification`

GetIdentification returns the Identification field if non-nil, zero value otherwise.

### GetIdentificationOk

`func (o *ClientCapabilitiesDeviceProfile) GetIdentificationOk() (*DeviceProfileIdentification, bool)`

GetIdentificationOk returns a tuple with the Identification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentification

`func (o *ClientCapabilitiesDeviceProfile) SetIdentification(v DeviceProfileIdentification)`

SetIdentification sets Identification field to given value.

### HasIdentification

`func (o *ClientCapabilitiesDeviceProfile) HasIdentification() bool`

HasIdentification returns a boolean if a field has been set.

### SetIdentificationNil

`func (o *ClientCapabilitiesDeviceProfile) SetIdentificationNil(b bool)`

 SetIdentificationNil sets the value for Identification to be an explicit nil

### UnsetIdentification
`func (o *ClientCapabilitiesDeviceProfile) UnsetIdentification()`

UnsetIdentification ensures that no value is present for Identification, not even an explicit nil
### GetFriendlyName

`func (o *ClientCapabilitiesDeviceProfile) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *ClientCapabilitiesDeviceProfile) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *ClientCapabilitiesDeviceProfile) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *ClientCapabilitiesDeviceProfile) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *ClientCapabilitiesDeviceProfile) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *ClientCapabilitiesDeviceProfile) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetManufacturer

`func (o *ClientCapabilitiesDeviceProfile) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *ClientCapabilitiesDeviceProfile) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *ClientCapabilitiesDeviceProfile) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *ClientCapabilitiesDeviceProfile) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *ClientCapabilitiesDeviceProfile) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *ClientCapabilitiesDeviceProfile) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetManufacturerUrl

`func (o *ClientCapabilitiesDeviceProfile) GetManufacturerUrl() string`

GetManufacturerUrl returns the ManufacturerUrl field if non-nil, zero value otherwise.

### GetManufacturerUrlOk

`func (o *ClientCapabilitiesDeviceProfile) GetManufacturerUrlOk() (*string, bool)`

GetManufacturerUrlOk returns a tuple with the ManufacturerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerUrl

`func (o *ClientCapabilitiesDeviceProfile) SetManufacturerUrl(v string)`

SetManufacturerUrl sets ManufacturerUrl field to given value.

### HasManufacturerUrl

`func (o *ClientCapabilitiesDeviceProfile) HasManufacturerUrl() bool`

HasManufacturerUrl returns a boolean if a field has been set.

### SetManufacturerUrlNil

`func (o *ClientCapabilitiesDeviceProfile) SetManufacturerUrlNil(b bool)`

 SetManufacturerUrlNil sets the value for ManufacturerUrl to be an explicit nil

### UnsetManufacturerUrl
`func (o *ClientCapabilitiesDeviceProfile) UnsetManufacturerUrl()`

UnsetManufacturerUrl ensures that no value is present for ManufacturerUrl, not even an explicit nil
### GetModelName

`func (o *ClientCapabilitiesDeviceProfile) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ClientCapabilitiesDeviceProfile) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ClientCapabilitiesDeviceProfile) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *ClientCapabilitiesDeviceProfile) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### SetModelNameNil

`func (o *ClientCapabilitiesDeviceProfile) SetModelNameNil(b bool)`

 SetModelNameNil sets the value for ModelName to be an explicit nil

### UnsetModelName
`func (o *ClientCapabilitiesDeviceProfile) UnsetModelName()`

UnsetModelName ensures that no value is present for ModelName, not even an explicit nil
### GetModelDescription

`func (o *ClientCapabilitiesDeviceProfile) GetModelDescription() string`

GetModelDescription returns the ModelDescription field if non-nil, zero value otherwise.

### GetModelDescriptionOk

`func (o *ClientCapabilitiesDeviceProfile) GetModelDescriptionOk() (*string, bool)`

GetModelDescriptionOk returns a tuple with the ModelDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelDescription

`func (o *ClientCapabilitiesDeviceProfile) SetModelDescription(v string)`

SetModelDescription sets ModelDescription field to given value.

### HasModelDescription

`func (o *ClientCapabilitiesDeviceProfile) HasModelDescription() bool`

HasModelDescription returns a boolean if a field has been set.

### SetModelDescriptionNil

`func (o *ClientCapabilitiesDeviceProfile) SetModelDescriptionNil(b bool)`

 SetModelDescriptionNil sets the value for ModelDescription to be an explicit nil

### UnsetModelDescription
`func (o *ClientCapabilitiesDeviceProfile) UnsetModelDescription()`

UnsetModelDescription ensures that no value is present for ModelDescription, not even an explicit nil
### GetModelNumber

`func (o *ClientCapabilitiesDeviceProfile) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *ClientCapabilitiesDeviceProfile) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *ClientCapabilitiesDeviceProfile) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *ClientCapabilitiesDeviceProfile) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### SetModelNumberNil

`func (o *ClientCapabilitiesDeviceProfile) SetModelNumberNil(b bool)`

 SetModelNumberNil sets the value for ModelNumber to be an explicit nil

### UnsetModelNumber
`func (o *ClientCapabilitiesDeviceProfile) UnsetModelNumber()`

UnsetModelNumber ensures that no value is present for ModelNumber, not even an explicit nil
### GetModelUrl

`func (o *ClientCapabilitiesDeviceProfile) GetModelUrl() string`

GetModelUrl returns the ModelUrl field if non-nil, zero value otherwise.

### GetModelUrlOk

`func (o *ClientCapabilitiesDeviceProfile) GetModelUrlOk() (*string, bool)`

GetModelUrlOk returns a tuple with the ModelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelUrl

`func (o *ClientCapabilitiesDeviceProfile) SetModelUrl(v string)`

SetModelUrl sets ModelUrl field to given value.

### HasModelUrl

`func (o *ClientCapabilitiesDeviceProfile) HasModelUrl() bool`

HasModelUrl returns a boolean if a field has been set.

### SetModelUrlNil

`func (o *ClientCapabilitiesDeviceProfile) SetModelUrlNil(b bool)`

 SetModelUrlNil sets the value for ModelUrl to be an explicit nil

### UnsetModelUrl
`func (o *ClientCapabilitiesDeviceProfile) UnsetModelUrl()`

UnsetModelUrl ensures that no value is present for ModelUrl, not even an explicit nil
### GetSerialNumber

`func (o *ClientCapabilitiesDeviceProfile) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ClientCapabilitiesDeviceProfile) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ClientCapabilitiesDeviceProfile) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ClientCapabilitiesDeviceProfile) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *ClientCapabilitiesDeviceProfile) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *ClientCapabilitiesDeviceProfile) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetEnableAlbumArtInDidl

`func (o *ClientCapabilitiesDeviceProfile) GetEnableAlbumArtInDidl() bool`

GetEnableAlbumArtInDidl returns the EnableAlbumArtInDidl field if non-nil, zero value otherwise.

### GetEnableAlbumArtInDidlOk

`func (o *ClientCapabilitiesDeviceProfile) GetEnableAlbumArtInDidlOk() (*bool, bool)`

GetEnableAlbumArtInDidlOk returns a tuple with the EnableAlbumArtInDidl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAlbumArtInDidl

`func (o *ClientCapabilitiesDeviceProfile) SetEnableAlbumArtInDidl(v bool)`

SetEnableAlbumArtInDidl sets EnableAlbumArtInDidl field to given value.

### HasEnableAlbumArtInDidl

`func (o *ClientCapabilitiesDeviceProfile) HasEnableAlbumArtInDidl() bool`

HasEnableAlbumArtInDidl returns a boolean if a field has been set.

### GetEnableSingleAlbumArtLimit

`func (o *ClientCapabilitiesDeviceProfile) GetEnableSingleAlbumArtLimit() bool`

GetEnableSingleAlbumArtLimit returns the EnableSingleAlbumArtLimit field if non-nil, zero value otherwise.

### GetEnableSingleAlbumArtLimitOk

`func (o *ClientCapabilitiesDeviceProfile) GetEnableSingleAlbumArtLimitOk() (*bool, bool)`

GetEnableSingleAlbumArtLimitOk returns a tuple with the EnableSingleAlbumArtLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSingleAlbumArtLimit

`func (o *ClientCapabilitiesDeviceProfile) SetEnableSingleAlbumArtLimit(v bool)`

SetEnableSingleAlbumArtLimit sets EnableSingleAlbumArtLimit field to given value.

### HasEnableSingleAlbumArtLimit

`func (o *ClientCapabilitiesDeviceProfile) HasEnableSingleAlbumArtLimit() bool`

HasEnableSingleAlbumArtLimit returns a boolean if a field has been set.

### GetEnableSingleSubtitleLimit

`func (o *ClientCapabilitiesDeviceProfile) GetEnableSingleSubtitleLimit() bool`

GetEnableSingleSubtitleLimit returns the EnableSingleSubtitleLimit field if non-nil, zero value otherwise.

### GetEnableSingleSubtitleLimitOk

`func (o *ClientCapabilitiesDeviceProfile) GetEnableSingleSubtitleLimitOk() (*bool, bool)`

GetEnableSingleSubtitleLimitOk returns a tuple with the EnableSingleSubtitleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSingleSubtitleLimit

`func (o *ClientCapabilitiesDeviceProfile) SetEnableSingleSubtitleLimit(v bool)`

SetEnableSingleSubtitleLimit sets EnableSingleSubtitleLimit field to given value.

### HasEnableSingleSubtitleLimit

`func (o *ClientCapabilitiesDeviceProfile) HasEnableSingleSubtitleLimit() bool`

HasEnableSingleSubtitleLimit returns a boolean if a field has been set.

### GetSupportedMediaTypes

`func (o *ClientCapabilitiesDeviceProfile) GetSupportedMediaTypes() string`

GetSupportedMediaTypes returns the SupportedMediaTypes field if non-nil, zero value otherwise.

### GetSupportedMediaTypesOk

`func (o *ClientCapabilitiesDeviceProfile) GetSupportedMediaTypesOk() (*string, bool)`

GetSupportedMediaTypesOk returns a tuple with the SupportedMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMediaTypes

`func (o *ClientCapabilitiesDeviceProfile) SetSupportedMediaTypes(v string)`

SetSupportedMediaTypes sets SupportedMediaTypes field to given value.

### HasSupportedMediaTypes

`func (o *ClientCapabilitiesDeviceProfile) HasSupportedMediaTypes() bool`

HasSupportedMediaTypes returns a boolean if a field has been set.

### GetUserId

`func (o *ClientCapabilitiesDeviceProfile) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ClientCapabilitiesDeviceProfile) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ClientCapabilitiesDeviceProfile) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ClientCapabilitiesDeviceProfile) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ClientCapabilitiesDeviceProfile) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ClientCapabilitiesDeviceProfile) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetAlbumArtPn

`func (o *ClientCapabilitiesDeviceProfile) GetAlbumArtPn() string`

GetAlbumArtPn returns the AlbumArtPn field if non-nil, zero value otherwise.

### GetAlbumArtPnOk

`func (o *ClientCapabilitiesDeviceProfile) GetAlbumArtPnOk() (*string, bool)`

GetAlbumArtPnOk returns a tuple with the AlbumArtPn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtPn

`func (o *ClientCapabilitiesDeviceProfile) SetAlbumArtPn(v string)`

SetAlbumArtPn sets AlbumArtPn field to given value.

### HasAlbumArtPn

`func (o *ClientCapabilitiesDeviceProfile) HasAlbumArtPn() bool`

HasAlbumArtPn returns a boolean if a field has been set.

### SetAlbumArtPnNil

`func (o *ClientCapabilitiesDeviceProfile) SetAlbumArtPnNil(b bool)`

 SetAlbumArtPnNil sets the value for AlbumArtPn to be an explicit nil

### UnsetAlbumArtPn
`func (o *ClientCapabilitiesDeviceProfile) UnsetAlbumArtPn()`

UnsetAlbumArtPn ensures that no value is present for AlbumArtPn, not even an explicit nil
### GetMaxAlbumArtWidth

`func (o *ClientCapabilitiesDeviceProfile) GetMaxAlbumArtWidth() int32`

GetMaxAlbumArtWidth returns the MaxAlbumArtWidth field if non-nil, zero value otherwise.

### GetMaxAlbumArtWidthOk

`func (o *ClientCapabilitiesDeviceProfile) GetMaxAlbumArtWidthOk() (*int32, bool)`

GetMaxAlbumArtWidthOk returns a tuple with the MaxAlbumArtWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlbumArtWidth

`func (o *ClientCapabilitiesDeviceProfile) SetMaxAlbumArtWidth(v int32)`

SetMaxAlbumArtWidth sets MaxAlbumArtWidth field to given value.

### HasMaxAlbumArtWidth

`func (o *ClientCapabilitiesDeviceProfile) HasMaxAlbumArtWidth() bool`

HasMaxAlbumArtWidth returns a boolean if a field has been set.

### SetMaxAlbumArtWidthNil

`func (o *ClientCapabilitiesDeviceProfile) SetMaxAlbumArtWidthNil(b bool)`

 SetMaxAlbumArtWidthNil sets the value for MaxAlbumArtWidth to be an explicit nil

### UnsetMaxAlbumArtWidth
`func (o *ClientCapabilitiesDeviceProfile) UnsetMaxAlbumArtWidth()`

UnsetMaxAlbumArtWidth ensures that no value is present for MaxAlbumArtWidth, not even an explicit nil
### GetMaxAlbumArtHeight

`func (o *ClientCapabilitiesDeviceProfile) GetMaxAlbumArtHeight() int32`

GetMaxAlbumArtHeight returns the MaxAlbumArtHeight field if non-nil, zero value otherwise.

### GetMaxAlbumArtHeightOk

`func (o *ClientCapabilitiesDeviceProfile) GetMaxAlbumArtHeightOk() (*int32, bool)`

GetMaxAlbumArtHeightOk returns a tuple with the MaxAlbumArtHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlbumArtHeight

`func (o *ClientCapabilitiesDeviceProfile) SetMaxAlbumArtHeight(v int32)`

SetMaxAlbumArtHeight sets MaxAlbumArtHeight field to given value.

### HasMaxAlbumArtHeight

`func (o *ClientCapabilitiesDeviceProfile) HasMaxAlbumArtHeight() bool`

HasMaxAlbumArtHeight returns a boolean if a field has been set.

### SetMaxAlbumArtHeightNil

`func (o *ClientCapabilitiesDeviceProfile) SetMaxAlbumArtHeightNil(b bool)`

 SetMaxAlbumArtHeightNil sets the value for MaxAlbumArtHeight to be an explicit nil

### UnsetMaxAlbumArtHeight
`func (o *ClientCapabilitiesDeviceProfile) UnsetMaxAlbumArtHeight()`

UnsetMaxAlbumArtHeight ensures that no value is present for MaxAlbumArtHeight, not even an explicit nil
### GetMaxIconWidth

`func (o *ClientCapabilitiesDeviceProfile) GetMaxIconWidth() int32`

GetMaxIconWidth returns the MaxIconWidth field if non-nil, zero value otherwise.

### GetMaxIconWidthOk

`func (o *ClientCapabilitiesDeviceProfile) GetMaxIconWidthOk() (*int32, bool)`

GetMaxIconWidthOk returns a tuple with the MaxIconWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIconWidth

`func (o *ClientCapabilitiesDeviceProfile) SetMaxIconWidth(v int32)`

SetMaxIconWidth sets MaxIconWidth field to given value.

### HasMaxIconWidth

`func (o *ClientCapabilitiesDeviceProfile) HasMaxIconWidth() bool`

HasMaxIconWidth returns a boolean if a field has been set.

### SetMaxIconWidthNil

`func (o *ClientCapabilitiesDeviceProfile) SetMaxIconWidthNil(b bool)`

 SetMaxIconWidthNil sets the value for MaxIconWidth to be an explicit nil

### UnsetMaxIconWidth
`func (o *ClientCapabilitiesDeviceProfile) UnsetMaxIconWidth()`

UnsetMaxIconWidth ensures that no value is present for MaxIconWidth, not even an explicit nil
### GetMaxIconHeight

`func (o *ClientCapabilitiesDeviceProfile) GetMaxIconHeight() int32`

GetMaxIconHeight returns the MaxIconHeight field if non-nil, zero value otherwise.

### GetMaxIconHeightOk

`func (o *ClientCapabilitiesDeviceProfile) GetMaxIconHeightOk() (*int32, bool)`

GetMaxIconHeightOk returns a tuple with the MaxIconHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIconHeight

`func (o *ClientCapabilitiesDeviceProfile) SetMaxIconHeight(v int32)`

SetMaxIconHeight sets MaxIconHeight field to given value.

### HasMaxIconHeight

`func (o *ClientCapabilitiesDeviceProfile) HasMaxIconHeight() bool`

HasMaxIconHeight returns a boolean if a field has been set.

### SetMaxIconHeightNil

`func (o *ClientCapabilitiesDeviceProfile) SetMaxIconHeightNil(b bool)`

 SetMaxIconHeightNil sets the value for MaxIconHeight to be an explicit nil

### UnsetMaxIconHeight
`func (o *ClientCapabilitiesDeviceProfile) UnsetMaxIconHeight()`

UnsetMaxIconHeight ensures that no value is present for MaxIconHeight, not even an explicit nil
### GetMaxStreamingBitrate

`func (o *ClientCapabilitiesDeviceProfile) GetMaxStreamingBitrate() int32`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *ClientCapabilitiesDeviceProfile) GetMaxStreamingBitrateOk() (*int32, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *ClientCapabilitiesDeviceProfile) SetMaxStreamingBitrate(v int32)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *ClientCapabilitiesDeviceProfile) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *ClientCapabilitiesDeviceProfile) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *ClientCapabilitiesDeviceProfile) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetMaxStaticBitrate

`func (o *ClientCapabilitiesDeviceProfile) GetMaxStaticBitrate() int32`

GetMaxStaticBitrate returns the MaxStaticBitrate field if non-nil, zero value otherwise.

### GetMaxStaticBitrateOk

`func (o *ClientCapabilitiesDeviceProfile) GetMaxStaticBitrateOk() (*int32, bool)`

GetMaxStaticBitrateOk returns a tuple with the MaxStaticBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaticBitrate

`func (o *ClientCapabilitiesDeviceProfile) SetMaxStaticBitrate(v int32)`

SetMaxStaticBitrate sets MaxStaticBitrate field to given value.

### HasMaxStaticBitrate

`func (o *ClientCapabilitiesDeviceProfile) HasMaxStaticBitrate() bool`

HasMaxStaticBitrate returns a boolean if a field has been set.

### SetMaxStaticBitrateNil

`func (o *ClientCapabilitiesDeviceProfile) SetMaxStaticBitrateNil(b bool)`

 SetMaxStaticBitrateNil sets the value for MaxStaticBitrate to be an explicit nil

### UnsetMaxStaticBitrate
`func (o *ClientCapabilitiesDeviceProfile) UnsetMaxStaticBitrate()`

UnsetMaxStaticBitrate ensures that no value is present for MaxStaticBitrate, not even an explicit nil
### GetMusicStreamingTranscodingBitrate

`func (o *ClientCapabilitiesDeviceProfile) GetMusicStreamingTranscodingBitrate() int32`

GetMusicStreamingTranscodingBitrate returns the MusicStreamingTranscodingBitrate field if non-nil, zero value otherwise.

### GetMusicStreamingTranscodingBitrateOk

`func (o *ClientCapabilitiesDeviceProfile) GetMusicStreamingTranscodingBitrateOk() (*int32, bool)`

GetMusicStreamingTranscodingBitrateOk returns a tuple with the MusicStreamingTranscodingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicStreamingTranscodingBitrate

`func (o *ClientCapabilitiesDeviceProfile) SetMusicStreamingTranscodingBitrate(v int32)`

SetMusicStreamingTranscodingBitrate sets MusicStreamingTranscodingBitrate field to given value.

### HasMusicStreamingTranscodingBitrate

`func (o *ClientCapabilitiesDeviceProfile) HasMusicStreamingTranscodingBitrate() bool`

HasMusicStreamingTranscodingBitrate returns a boolean if a field has been set.

### SetMusicStreamingTranscodingBitrateNil

`func (o *ClientCapabilitiesDeviceProfile) SetMusicStreamingTranscodingBitrateNil(b bool)`

 SetMusicStreamingTranscodingBitrateNil sets the value for MusicStreamingTranscodingBitrate to be an explicit nil

### UnsetMusicStreamingTranscodingBitrate
`func (o *ClientCapabilitiesDeviceProfile) UnsetMusicStreamingTranscodingBitrate()`

UnsetMusicStreamingTranscodingBitrate ensures that no value is present for MusicStreamingTranscodingBitrate, not even an explicit nil
### GetMaxStaticMusicBitrate

`func (o *ClientCapabilitiesDeviceProfile) GetMaxStaticMusicBitrate() int32`

GetMaxStaticMusicBitrate returns the MaxStaticMusicBitrate field if non-nil, zero value otherwise.

### GetMaxStaticMusicBitrateOk

`func (o *ClientCapabilitiesDeviceProfile) GetMaxStaticMusicBitrateOk() (*int32, bool)`

GetMaxStaticMusicBitrateOk returns a tuple with the MaxStaticMusicBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaticMusicBitrate

`func (o *ClientCapabilitiesDeviceProfile) SetMaxStaticMusicBitrate(v int32)`

SetMaxStaticMusicBitrate sets MaxStaticMusicBitrate field to given value.

### HasMaxStaticMusicBitrate

`func (o *ClientCapabilitiesDeviceProfile) HasMaxStaticMusicBitrate() bool`

HasMaxStaticMusicBitrate returns a boolean if a field has been set.

### SetMaxStaticMusicBitrateNil

`func (o *ClientCapabilitiesDeviceProfile) SetMaxStaticMusicBitrateNil(b bool)`

 SetMaxStaticMusicBitrateNil sets the value for MaxStaticMusicBitrate to be an explicit nil

### UnsetMaxStaticMusicBitrate
`func (o *ClientCapabilitiesDeviceProfile) UnsetMaxStaticMusicBitrate()`

UnsetMaxStaticMusicBitrate ensures that no value is present for MaxStaticMusicBitrate, not even an explicit nil
### GetSonyAggregationFlags

`func (o *ClientCapabilitiesDeviceProfile) GetSonyAggregationFlags() string`

GetSonyAggregationFlags returns the SonyAggregationFlags field if non-nil, zero value otherwise.

### GetSonyAggregationFlagsOk

`func (o *ClientCapabilitiesDeviceProfile) GetSonyAggregationFlagsOk() (*string, bool)`

GetSonyAggregationFlagsOk returns a tuple with the SonyAggregationFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSonyAggregationFlags

`func (o *ClientCapabilitiesDeviceProfile) SetSonyAggregationFlags(v string)`

SetSonyAggregationFlags sets SonyAggregationFlags field to given value.

### HasSonyAggregationFlags

`func (o *ClientCapabilitiesDeviceProfile) HasSonyAggregationFlags() bool`

HasSonyAggregationFlags returns a boolean if a field has been set.

### SetSonyAggregationFlagsNil

`func (o *ClientCapabilitiesDeviceProfile) SetSonyAggregationFlagsNil(b bool)`

 SetSonyAggregationFlagsNil sets the value for SonyAggregationFlags to be an explicit nil

### UnsetSonyAggregationFlags
`func (o *ClientCapabilitiesDeviceProfile) UnsetSonyAggregationFlags()`

UnsetSonyAggregationFlags ensures that no value is present for SonyAggregationFlags, not even an explicit nil
### GetProtocolInfo

`func (o *ClientCapabilitiesDeviceProfile) GetProtocolInfo() string`

GetProtocolInfo returns the ProtocolInfo field if non-nil, zero value otherwise.

### GetProtocolInfoOk

`func (o *ClientCapabilitiesDeviceProfile) GetProtocolInfoOk() (*string, bool)`

GetProtocolInfoOk returns a tuple with the ProtocolInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolInfo

`func (o *ClientCapabilitiesDeviceProfile) SetProtocolInfo(v string)`

SetProtocolInfo sets ProtocolInfo field to given value.

### HasProtocolInfo

`func (o *ClientCapabilitiesDeviceProfile) HasProtocolInfo() bool`

HasProtocolInfo returns a boolean if a field has been set.

### SetProtocolInfoNil

`func (o *ClientCapabilitiesDeviceProfile) SetProtocolInfoNil(b bool)`

 SetProtocolInfoNil sets the value for ProtocolInfo to be an explicit nil

### UnsetProtocolInfo
`func (o *ClientCapabilitiesDeviceProfile) UnsetProtocolInfo()`

UnsetProtocolInfo ensures that no value is present for ProtocolInfo, not even an explicit nil
### GetTimelineOffsetSeconds

`func (o *ClientCapabilitiesDeviceProfile) GetTimelineOffsetSeconds() int32`

GetTimelineOffsetSeconds returns the TimelineOffsetSeconds field if non-nil, zero value otherwise.

### GetTimelineOffsetSecondsOk

`func (o *ClientCapabilitiesDeviceProfile) GetTimelineOffsetSecondsOk() (*int32, bool)`

GetTimelineOffsetSecondsOk returns a tuple with the TimelineOffsetSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineOffsetSeconds

`func (o *ClientCapabilitiesDeviceProfile) SetTimelineOffsetSeconds(v int32)`

SetTimelineOffsetSeconds sets TimelineOffsetSeconds field to given value.

### HasTimelineOffsetSeconds

`func (o *ClientCapabilitiesDeviceProfile) HasTimelineOffsetSeconds() bool`

HasTimelineOffsetSeconds returns a boolean if a field has been set.

### GetRequiresPlainVideoItems

`func (o *ClientCapabilitiesDeviceProfile) GetRequiresPlainVideoItems() bool`

GetRequiresPlainVideoItems returns the RequiresPlainVideoItems field if non-nil, zero value otherwise.

### GetRequiresPlainVideoItemsOk

`func (o *ClientCapabilitiesDeviceProfile) GetRequiresPlainVideoItemsOk() (*bool, bool)`

GetRequiresPlainVideoItemsOk returns a tuple with the RequiresPlainVideoItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPlainVideoItems

`func (o *ClientCapabilitiesDeviceProfile) SetRequiresPlainVideoItems(v bool)`

SetRequiresPlainVideoItems sets RequiresPlainVideoItems field to given value.

### HasRequiresPlainVideoItems

`func (o *ClientCapabilitiesDeviceProfile) HasRequiresPlainVideoItems() bool`

HasRequiresPlainVideoItems returns a boolean if a field has been set.

### GetRequiresPlainFolders

`func (o *ClientCapabilitiesDeviceProfile) GetRequiresPlainFolders() bool`

GetRequiresPlainFolders returns the RequiresPlainFolders field if non-nil, zero value otherwise.

### GetRequiresPlainFoldersOk

`func (o *ClientCapabilitiesDeviceProfile) GetRequiresPlainFoldersOk() (*bool, bool)`

GetRequiresPlainFoldersOk returns a tuple with the RequiresPlainFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPlainFolders

`func (o *ClientCapabilitiesDeviceProfile) SetRequiresPlainFolders(v bool)`

SetRequiresPlainFolders sets RequiresPlainFolders field to given value.

### HasRequiresPlainFolders

`func (o *ClientCapabilitiesDeviceProfile) HasRequiresPlainFolders() bool`

HasRequiresPlainFolders returns a boolean if a field has been set.

### GetEnableMSMediaReceiverRegistrar

`func (o *ClientCapabilitiesDeviceProfile) GetEnableMSMediaReceiverRegistrar() bool`

GetEnableMSMediaReceiverRegistrar returns the EnableMSMediaReceiverRegistrar field if non-nil, zero value otherwise.

### GetEnableMSMediaReceiverRegistrarOk

`func (o *ClientCapabilitiesDeviceProfile) GetEnableMSMediaReceiverRegistrarOk() (*bool, bool)`

GetEnableMSMediaReceiverRegistrarOk returns a tuple with the EnableMSMediaReceiverRegistrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMSMediaReceiverRegistrar

`func (o *ClientCapabilitiesDeviceProfile) SetEnableMSMediaReceiverRegistrar(v bool)`

SetEnableMSMediaReceiverRegistrar sets EnableMSMediaReceiverRegistrar field to given value.

### HasEnableMSMediaReceiverRegistrar

`func (o *ClientCapabilitiesDeviceProfile) HasEnableMSMediaReceiverRegistrar() bool`

HasEnableMSMediaReceiverRegistrar returns a boolean if a field has been set.

### GetIgnoreTranscodeByteRangeRequests

`func (o *ClientCapabilitiesDeviceProfile) GetIgnoreTranscodeByteRangeRequests() bool`

GetIgnoreTranscodeByteRangeRequests returns the IgnoreTranscodeByteRangeRequests field if non-nil, zero value otherwise.

### GetIgnoreTranscodeByteRangeRequestsOk

`func (o *ClientCapabilitiesDeviceProfile) GetIgnoreTranscodeByteRangeRequestsOk() (*bool, bool)`

GetIgnoreTranscodeByteRangeRequestsOk returns a tuple with the IgnoreTranscodeByteRangeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreTranscodeByteRangeRequests

`func (o *ClientCapabilitiesDeviceProfile) SetIgnoreTranscodeByteRangeRequests(v bool)`

SetIgnoreTranscodeByteRangeRequests sets IgnoreTranscodeByteRangeRequests field to given value.

### HasIgnoreTranscodeByteRangeRequests

`func (o *ClientCapabilitiesDeviceProfile) HasIgnoreTranscodeByteRangeRequests() bool`

HasIgnoreTranscodeByteRangeRequests returns a boolean if a field has been set.

### GetXmlRootAttributes

`func (o *ClientCapabilitiesDeviceProfile) GetXmlRootAttributes() []XmlAttribute`

GetXmlRootAttributes returns the XmlRootAttributes field if non-nil, zero value otherwise.

### GetXmlRootAttributesOk

`func (o *ClientCapabilitiesDeviceProfile) GetXmlRootAttributesOk() (*[]XmlAttribute, bool)`

GetXmlRootAttributesOk returns a tuple with the XmlRootAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlRootAttributes

`func (o *ClientCapabilitiesDeviceProfile) SetXmlRootAttributes(v []XmlAttribute)`

SetXmlRootAttributes sets XmlRootAttributes field to given value.

### HasXmlRootAttributes

`func (o *ClientCapabilitiesDeviceProfile) HasXmlRootAttributes() bool`

HasXmlRootAttributes returns a boolean if a field has been set.

### GetDirectPlayProfiles

`func (o *ClientCapabilitiesDeviceProfile) GetDirectPlayProfiles() []DirectPlayProfile`

GetDirectPlayProfiles returns the DirectPlayProfiles field if non-nil, zero value otherwise.

### GetDirectPlayProfilesOk

`func (o *ClientCapabilitiesDeviceProfile) GetDirectPlayProfilesOk() (*[]DirectPlayProfile, bool)`

GetDirectPlayProfilesOk returns a tuple with the DirectPlayProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPlayProfiles

`func (o *ClientCapabilitiesDeviceProfile) SetDirectPlayProfiles(v []DirectPlayProfile)`

SetDirectPlayProfiles sets DirectPlayProfiles field to given value.

### HasDirectPlayProfiles

`func (o *ClientCapabilitiesDeviceProfile) HasDirectPlayProfiles() bool`

HasDirectPlayProfiles returns a boolean if a field has been set.

### GetTranscodingProfiles

`func (o *ClientCapabilitiesDeviceProfile) GetTranscodingProfiles() []TranscodingProfile`

GetTranscodingProfiles returns the TranscodingProfiles field if non-nil, zero value otherwise.

### GetTranscodingProfilesOk

`func (o *ClientCapabilitiesDeviceProfile) GetTranscodingProfilesOk() (*[]TranscodingProfile, bool)`

GetTranscodingProfilesOk returns a tuple with the TranscodingProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingProfiles

`func (o *ClientCapabilitiesDeviceProfile) SetTranscodingProfiles(v []TranscodingProfile)`

SetTranscodingProfiles sets TranscodingProfiles field to given value.

### HasTranscodingProfiles

`func (o *ClientCapabilitiesDeviceProfile) HasTranscodingProfiles() bool`

HasTranscodingProfiles returns a boolean if a field has been set.

### GetContainerProfiles

`func (o *ClientCapabilitiesDeviceProfile) GetContainerProfiles() []ContainerProfile`

GetContainerProfiles returns the ContainerProfiles field if non-nil, zero value otherwise.

### GetContainerProfilesOk

`func (o *ClientCapabilitiesDeviceProfile) GetContainerProfilesOk() (*[]ContainerProfile, bool)`

GetContainerProfilesOk returns a tuple with the ContainerProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerProfiles

`func (o *ClientCapabilitiesDeviceProfile) SetContainerProfiles(v []ContainerProfile)`

SetContainerProfiles sets ContainerProfiles field to given value.

### HasContainerProfiles

`func (o *ClientCapabilitiesDeviceProfile) HasContainerProfiles() bool`

HasContainerProfiles returns a boolean if a field has been set.

### GetCodecProfiles

`func (o *ClientCapabilitiesDeviceProfile) GetCodecProfiles() []CodecProfile`

GetCodecProfiles returns the CodecProfiles field if non-nil, zero value otherwise.

### GetCodecProfilesOk

`func (o *ClientCapabilitiesDeviceProfile) GetCodecProfilesOk() (*[]CodecProfile, bool)`

GetCodecProfilesOk returns a tuple with the CodecProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecProfiles

`func (o *ClientCapabilitiesDeviceProfile) SetCodecProfiles(v []CodecProfile)`

SetCodecProfiles sets CodecProfiles field to given value.

### HasCodecProfiles

`func (o *ClientCapabilitiesDeviceProfile) HasCodecProfiles() bool`

HasCodecProfiles returns a boolean if a field has been set.

### GetResponseProfiles

`func (o *ClientCapabilitiesDeviceProfile) GetResponseProfiles() []ResponseProfile`

GetResponseProfiles returns the ResponseProfiles field if non-nil, zero value otherwise.

### GetResponseProfilesOk

`func (o *ClientCapabilitiesDeviceProfile) GetResponseProfilesOk() (*[]ResponseProfile, bool)`

GetResponseProfilesOk returns a tuple with the ResponseProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseProfiles

`func (o *ClientCapabilitiesDeviceProfile) SetResponseProfiles(v []ResponseProfile)`

SetResponseProfiles sets ResponseProfiles field to given value.

### HasResponseProfiles

`func (o *ClientCapabilitiesDeviceProfile) HasResponseProfiles() bool`

HasResponseProfiles returns a boolean if a field has been set.

### GetSubtitleProfiles

`func (o *ClientCapabilitiesDeviceProfile) GetSubtitleProfiles() []SubtitleProfile`

GetSubtitleProfiles returns the SubtitleProfiles field if non-nil, zero value otherwise.

### GetSubtitleProfilesOk

`func (o *ClientCapabilitiesDeviceProfile) GetSubtitleProfilesOk() (*[]SubtitleProfile, bool)`

GetSubtitleProfilesOk returns a tuple with the SubtitleProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleProfiles

`func (o *ClientCapabilitiesDeviceProfile) SetSubtitleProfiles(v []SubtitleProfile)`

SetSubtitleProfiles sets SubtitleProfiles field to given value.

### HasSubtitleProfiles

`func (o *ClientCapabilitiesDeviceProfile) HasSubtitleProfiles() bool`

HasSubtitleProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


