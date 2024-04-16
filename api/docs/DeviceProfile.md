# DeviceProfile

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

### NewDeviceProfile

`func NewDeviceProfile() *DeviceProfile`

NewDeviceProfile instantiates a new DeviceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceProfileWithDefaults

`func NewDeviceProfileWithDefaults() *DeviceProfile`

NewDeviceProfileWithDefaults instantiates a new DeviceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeviceProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeviceProfile) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeviceProfile) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *DeviceProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DeviceProfile) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeviceProfile) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentification

`func (o *DeviceProfile) GetIdentification() DeviceProfileIdentification`

GetIdentification returns the Identification field if non-nil, zero value otherwise.

### GetIdentificationOk

`func (o *DeviceProfile) GetIdentificationOk() (*DeviceProfileIdentification, bool)`

GetIdentificationOk returns a tuple with the Identification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentification

`func (o *DeviceProfile) SetIdentification(v DeviceProfileIdentification)`

SetIdentification sets Identification field to given value.

### HasIdentification

`func (o *DeviceProfile) HasIdentification() bool`

HasIdentification returns a boolean if a field has been set.

### SetIdentificationNil

`func (o *DeviceProfile) SetIdentificationNil(b bool)`

 SetIdentificationNil sets the value for Identification to be an explicit nil

### UnsetIdentification
`func (o *DeviceProfile) UnsetIdentification()`

UnsetIdentification ensures that no value is present for Identification, not even an explicit nil
### GetFriendlyName

`func (o *DeviceProfile) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *DeviceProfile) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *DeviceProfile) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *DeviceProfile) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *DeviceProfile) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *DeviceProfile) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetManufacturer

`func (o *DeviceProfile) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *DeviceProfile) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *DeviceProfile) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *DeviceProfile) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *DeviceProfile) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *DeviceProfile) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetManufacturerUrl

`func (o *DeviceProfile) GetManufacturerUrl() string`

GetManufacturerUrl returns the ManufacturerUrl field if non-nil, zero value otherwise.

### GetManufacturerUrlOk

`func (o *DeviceProfile) GetManufacturerUrlOk() (*string, bool)`

GetManufacturerUrlOk returns a tuple with the ManufacturerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerUrl

`func (o *DeviceProfile) SetManufacturerUrl(v string)`

SetManufacturerUrl sets ManufacturerUrl field to given value.

### HasManufacturerUrl

`func (o *DeviceProfile) HasManufacturerUrl() bool`

HasManufacturerUrl returns a boolean if a field has been set.

### SetManufacturerUrlNil

`func (o *DeviceProfile) SetManufacturerUrlNil(b bool)`

 SetManufacturerUrlNil sets the value for ManufacturerUrl to be an explicit nil

### UnsetManufacturerUrl
`func (o *DeviceProfile) UnsetManufacturerUrl()`

UnsetManufacturerUrl ensures that no value is present for ManufacturerUrl, not even an explicit nil
### GetModelName

`func (o *DeviceProfile) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *DeviceProfile) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *DeviceProfile) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *DeviceProfile) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### SetModelNameNil

`func (o *DeviceProfile) SetModelNameNil(b bool)`

 SetModelNameNil sets the value for ModelName to be an explicit nil

### UnsetModelName
`func (o *DeviceProfile) UnsetModelName()`

UnsetModelName ensures that no value is present for ModelName, not even an explicit nil
### GetModelDescription

`func (o *DeviceProfile) GetModelDescription() string`

GetModelDescription returns the ModelDescription field if non-nil, zero value otherwise.

### GetModelDescriptionOk

`func (o *DeviceProfile) GetModelDescriptionOk() (*string, bool)`

GetModelDescriptionOk returns a tuple with the ModelDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelDescription

`func (o *DeviceProfile) SetModelDescription(v string)`

SetModelDescription sets ModelDescription field to given value.

### HasModelDescription

`func (o *DeviceProfile) HasModelDescription() bool`

HasModelDescription returns a boolean if a field has been set.

### SetModelDescriptionNil

`func (o *DeviceProfile) SetModelDescriptionNil(b bool)`

 SetModelDescriptionNil sets the value for ModelDescription to be an explicit nil

### UnsetModelDescription
`func (o *DeviceProfile) UnsetModelDescription()`

UnsetModelDescription ensures that no value is present for ModelDescription, not even an explicit nil
### GetModelNumber

`func (o *DeviceProfile) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *DeviceProfile) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *DeviceProfile) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *DeviceProfile) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### SetModelNumberNil

`func (o *DeviceProfile) SetModelNumberNil(b bool)`

 SetModelNumberNil sets the value for ModelNumber to be an explicit nil

### UnsetModelNumber
`func (o *DeviceProfile) UnsetModelNumber()`

UnsetModelNumber ensures that no value is present for ModelNumber, not even an explicit nil
### GetModelUrl

`func (o *DeviceProfile) GetModelUrl() string`

GetModelUrl returns the ModelUrl field if non-nil, zero value otherwise.

### GetModelUrlOk

`func (o *DeviceProfile) GetModelUrlOk() (*string, bool)`

GetModelUrlOk returns a tuple with the ModelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelUrl

`func (o *DeviceProfile) SetModelUrl(v string)`

SetModelUrl sets ModelUrl field to given value.

### HasModelUrl

`func (o *DeviceProfile) HasModelUrl() bool`

HasModelUrl returns a boolean if a field has been set.

### SetModelUrlNil

`func (o *DeviceProfile) SetModelUrlNil(b bool)`

 SetModelUrlNil sets the value for ModelUrl to be an explicit nil

### UnsetModelUrl
`func (o *DeviceProfile) UnsetModelUrl()`

UnsetModelUrl ensures that no value is present for ModelUrl, not even an explicit nil
### GetSerialNumber

`func (o *DeviceProfile) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DeviceProfile) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DeviceProfile) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DeviceProfile) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *DeviceProfile) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *DeviceProfile) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetEnableAlbumArtInDidl

`func (o *DeviceProfile) GetEnableAlbumArtInDidl() bool`

GetEnableAlbumArtInDidl returns the EnableAlbumArtInDidl field if non-nil, zero value otherwise.

### GetEnableAlbumArtInDidlOk

`func (o *DeviceProfile) GetEnableAlbumArtInDidlOk() (*bool, bool)`

GetEnableAlbumArtInDidlOk returns a tuple with the EnableAlbumArtInDidl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAlbumArtInDidl

`func (o *DeviceProfile) SetEnableAlbumArtInDidl(v bool)`

SetEnableAlbumArtInDidl sets EnableAlbumArtInDidl field to given value.

### HasEnableAlbumArtInDidl

`func (o *DeviceProfile) HasEnableAlbumArtInDidl() bool`

HasEnableAlbumArtInDidl returns a boolean if a field has been set.

### GetEnableSingleAlbumArtLimit

`func (o *DeviceProfile) GetEnableSingleAlbumArtLimit() bool`

GetEnableSingleAlbumArtLimit returns the EnableSingleAlbumArtLimit field if non-nil, zero value otherwise.

### GetEnableSingleAlbumArtLimitOk

`func (o *DeviceProfile) GetEnableSingleAlbumArtLimitOk() (*bool, bool)`

GetEnableSingleAlbumArtLimitOk returns a tuple with the EnableSingleAlbumArtLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSingleAlbumArtLimit

`func (o *DeviceProfile) SetEnableSingleAlbumArtLimit(v bool)`

SetEnableSingleAlbumArtLimit sets EnableSingleAlbumArtLimit field to given value.

### HasEnableSingleAlbumArtLimit

`func (o *DeviceProfile) HasEnableSingleAlbumArtLimit() bool`

HasEnableSingleAlbumArtLimit returns a boolean if a field has been set.

### GetEnableSingleSubtitleLimit

`func (o *DeviceProfile) GetEnableSingleSubtitleLimit() bool`

GetEnableSingleSubtitleLimit returns the EnableSingleSubtitleLimit field if non-nil, zero value otherwise.

### GetEnableSingleSubtitleLimitOk

`func (o *DeviceProfile) GetEnableSingleSubtitleLimitOk() (*bool, bool)`

GetEnableSingleSubtitleLimitOk returns a tuple with the EnableSingleSubtitleLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSingleSubtitleLimit

`func (o *DeviceProfile) SetEnableSingleSubtitleLimit(v bool)`

SetEnableSingleSubtitleLimit sets EnableSingleSubtitleLimit field to given value.

### HasEnableSingleSubtitleLimit

`func (o *DeviceProfile) HasEnableSingleSubtitleLimit() bool`

HasEnableSingleSubtitleLimit returns a boolean if a field has been set.

### GetSupportedMediaTypes

`func (o *DeviceProfile) GetSupportedMediaTypes() string`

GetSupportedMediaTypes returns the SupportedMediaTypes field if non-nil, zero value otherwise.

### GetSupportedMediaTypesOk

`func (o *DeviceProfile) GetSupportedMediaTypesOk() (*string, bool)`

GetSupportedMediaTypesOk returns a tuple with the SupportedMediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMediaTypes

`func (o *DeviceProfile) SetSupportedMediaTypes(v string)`

SetSupportedMediaTypes sets SupportedMediaTypes field to given value.

### HasSupportedMediaTypes

`func (o *DeviceProfile) HasSupportedMediaTypes() bool`

HasSupportedMediaTypes returns a boolean if a field has been set.

### GetUserId

`func (o *DeviceProfile) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DeviceProfile) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DeviceProfile) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *DeviceProfile) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *DeviceProfile) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *DeviceProfile) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetAlbumArtPn

`func (o *DeviceProfile) GetAlbumArtPn() string`

GetAlbumArtPn returns the AlbumArtPn field if non-nil, zero value otherwise.

### GetAlbumArtPnOk

`func (o *DeviceProfile) GetAlbumArtPnOk() (*string, bool)`

GetAlbumArtPnOk returns a tuple with the AlbumArtPn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlbumArtPn

`func (o *DeviceProfile) SetAlbumArtPn(v string)`

SetAlbumArtPn sets AlbumArtPn field to given value.

### HasAlbumArtPn

`func (o *DeviceProfile) HasAlbumArtPn() bool`

HasAlbumArtPn returns a boolean if a field has been set.

### SetAlbumArtPnNil

`func (o *DeviceProfile) SetAlbumArtPnNil(b bool)`

 SetAlbumArtPnNil sets the value for AlbumArtPn to be an explicit nil

### UnsetAlbumArtPn
`func (o *DeviceProfile) UnsetAlbumArtPn()`

UnsetAlbumArtPn ensures that no value is present for AlbumArtPn, not even an explicit nil
### GetMaxAlbumArtWidth

`func (o *DeviceProfile) GetMaxAlbumArtWidth() int32`

GetMaxAlbumArtWidth returns the MaxAlbumArtWidth field if non-nil, zero value otherwise.

### GetMaxAlbumArtWidthOk

`func (o *DeviceProfile) GetMaxAlbumArtWidthOk() (*int32, bool)`

GetMaxAlbumArtWidthOk returns a tuple with the MaxAlbumArtWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlbumArtWidth

`func (o *DeviceProfile) SetMaxAlbumArtWidth(v int32)`

SetMaxAlbumArtWidth sets MaxAlbumArtWidth field to given value.

### HasMaxAlbumArtWidth

`func (o *DeviceProfile) HasMaxAlbumArtWidth() bool`

HasMaxAlbumArtWidth returns a boolean if a field has been set.

### SetMaxAlbumArtWidthNil

`func (o *DeviceProfile) SetMaxAlbumArtWidthNil(b bool)`

 SetMaxAlbumArtWidthNil sets the value for MaxAlbumArtWidth to be an explicit nil

### UnsetMaxAlbumArtWidth
`func (o *DeviceProfile) UnsetMaxAlbumArtWidth()`

UnsetMaxAlbumArtWidth ensures that no value is present for MaxAlbumArtWidth, not even an explicit nil
### GetMaxAlbumArtHeight

`func (o *DeviceProfile) GetMaxAlbumArtHeight() int32`

GetMaxAlbumArtHeight returns the MaxAlbumArtHeight field if non-nil, zero value otherwise.

### GetMaxAlbumArtHeightOk

`func (o *DeviceProfile) GetMaxAlbumArtHeightOk() (*int32, bool)`

GetMaxAlbumArtHeightOk returns a tuple with the MaxAlbumArtHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAlbumArtHeight

`func (o *DeviceProfile) SetMaxAlbumArtHeight(v int32)`

SetMaxAlbumArtHeight sets MaxAlbumArtHeight field to given value.

### HasMaxAlbumArtHeight

`func (o *DeviceProfile) HasMaxAlbumArtHeight() bool`

HasMaxAlbumArtHeight returns a boolean if a field has been set.

### SetMaxAlbumArtHeightNil

`func (o *DeviceProfile) SetMaxAlbumArtHeightNil(b bool)`

 SetMaxAlbumArtHeightNil sets the value for MaxAlbumArtHeight to be an explicit nil

### UnsetMaxAlbumArtHeight
`func (o *DeviceProfile) UnsetMaxAlbumArtHeight()`

UnsetMaxAlbumArtHeight ensures that no value is present for MaxAlbumArtHeight, not even an explicit nil
### GetMaxIconWidth

`func (o *DeviceProfile) GetMaxIconWidth() int32`

GetMaxIconWidth returns the MaxIconWidth field if non-nil, zero value otherwise.

### GetMaxIconWidthOk

`func (o *DeviceProfile) GetMaxIconWidthOk() (*int32, bool)`

GetMaxIconWidthOk returns a tuple with the MaxIconWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIconWidth

`func (o *DeviceProfile) SetMaxIconWidth(v int32)`

SetMaxIconWidth sets MaxIconWidth field to given value.

### HasMaxIconWidth

`func (o *DeviceProfile) HasMaxIconWidth() bool`

HasMaxIconWidth returns a boolean if a field has been set.

### SetMaxIconWidthNil

`func (o *DeviceProfile) SetMaxIconWidthNil(b bool)`

 SetMaxIconWidthNil sets the value for MaxIconWidth to be an explicit nil

### UnsetMaxIconWidth
`func (o *DeviceProfile) UnsetMaxIconWidth()`

UnsetMaxIconWidth ensures that no value is present for MaxIconWidth, not even an explicit nil
### GetMaxIconHeight

`func (o *DeviceProfile) GetMaxIconHeight() int32`

GetMaxIconHeight returns the MaxIconHeight field if non-nil, zero value otherwise.

### GetMaxIconHeightOk

`func (o *DeviceProfile) GetMaxIconHeightOk() (*int32, bool)`

GetMaxIconHeightOk returns a tuple with the MaxIconHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIconHeight

`func (o *DeviceProfile) SetMaxIconHeight(v int32)`

SetMaxIconHeight sets MaxIconHeight field to given value.

### HasMaxIconHeight

`func (o *DeviceProfile) HasMaxIconHeight() bool`

HasMaxIconHeight returns a boolean if a field has been set.

### SetMaxIconHeightNil

`func (o *DeviceProfile) SetMaxIconHeightNil(b bool)`

 SetMaxIconHeightNil sets the value for MaxIconHeight to be an explicit nil

### UnsetMaxIconHeight
`func (o *DeviceProfile) UnsetMaxIconHeight()`

UnsetMaxIconHeight ensures that no value is present for MaxIconHeight, not even an explicit nil
### GetMaxStreamingBitrate

`func (o *DeviceProfile) GetMaxStreamingBitrate() int32`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *DeviceProfile) GetMaxStreamingBitrateOk() (*int32, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *DeviceProfile) SetMaxStreamingBitrate(v int32)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *DeviceProfile) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *DeviceProfile) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *DeviceProfile) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetMaxStaticBitrate

`func (o *DeviceProfile) GetMaxStaticBitrate() int32`

GetMaxStaticBitrate returns the MaxStaticBitrate field if non-nil, zero value otherwise.

### GetMaxStaticBitrateOk

`func (o *DeviceProfile) GetMaxStaticBitrateOk() (*int32, bool)`

GetMaxStaticBitrateOk returns a tuple with the MaxStaticBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaticBitrate

`func (o *DeviceProfile) SetMaxStaticBitrate(v int32)`

SetMaxStaticBitrate sets MaxStaticBitrate field to given value.

### HasMaxStaticBitrate

`func (o *DeviceProfile) HasMaxStaticBitrate() bool`

HasMaxStaticBitrate returns a boolean if a field has been set.

### SetMaxStaticBitrateNil

`func (o *DeviceProfile) SetMaxStaticBitrateNil(b bool)`

 SetMaxStaticBitrateNil sets the value for MaxStaticBitrate to be an explicit nil

### UnsetMaxStaticBitrate
`func (o *DeviceProfile) UnsetMaxStaticBitrate()`

UnsetMaxStaticBitrate ensures that no value is present for MaxStaticBitrate, not even an explicit nil
### GetMusicStreamingTranscodingBitrate

`func (o *DeviceProfile) GetMusicStreamingTranscodingBitrate() int32`

GetMusicStreamingTranscodingBitrate returns the MusicStreamingTranscodingBitrate field if non-nil, zero value otherwise.

### GetMusicStreamingTranscodingBitrateOk

`func (o *DeviceProfile) GetMusicStreamingTranscodingBitrateOk() (*int32, bool)`

GetMusicStreamingTranscodingBitrateOk returns a tuple with the MusicStreamingTranscodingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicStreamingTranscodingBitrate

`func (o *DeviceProfile) SetMusicStreamingTranscodingBitrate(v int32)`

SetMusicStreamingTranscodingBitrate sets MusicStreamingTranscodingBitrate field to given value.

### HasMusicStreamingTranscodingBitrate

`func (o *DeviceProfile) HasMusicStreamingTranscodingBitrate() bool`

HasMusicStreamingTranscodingBitrate returns a boolean if a field has been set.

### SetMusicStreamingTranscodingBitrateNil

`func (o *DeviceProfile) SetMusicStreamingTranscodingBitrateNil(b bool)`

 SetMusicStreamingTranscodingBitrateNil sets the value for MusicStreamingTranscodingBitrate to be an explicit nil

### UnsetMusicStreamingTranscodingBitrate
`func (o *DeviceProfile) UnsetMusicStreamingTranscodingBitrate()`

UnsetMusicStreamingTranscodingBitrate ensures that no value is present for MusicStreamingTranscodingBitrate, not even an explicit nil
### GetMaxStaticMusicBitrate

`func (o *DeviceProfile) GetMaxStaticMusicBitrate() int32`

GetMaxStaticMusicBitrate returns the MaxStaticMusicBitrate field if non-nil, zero value otherwise.

### GetMaxStaticMusicBitrateOk

`func (o *DeviceProfile) GetMaxStaticMusicBitrateOk() (*int32, bool)`

GetMaxStaticMusicBitrateOk returns a tuple with the MaxStaticMusicBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaticMusicBitrate

`func (o *DeviceProfile) SetMaxStaticMusicBitrate(v int32)`

SetMaxStaticMusicBitrate sets MaxStaticMusicBitrate field to given value.

### HasMaxStaticMusicBitrate

`func (o *DeviceProfile) HasMaxStaticMusicBitrate() bool`

HasMaxStaticMusicBitrate returns a boolean if a field has been set.

### SetMaxStaticMusicBitrateNil

`func (o *DeviceProfile) SetMaxStaticMusicBitrateNil(b bool)`

 SetMaxStaticMusicBitrateNil sets the value for MaxStaticMusicBitrate to be an explicit nil

### UnsetMaxStaticMusicBitrate
`func (o *DeviceProfile) UnsetMaxStaticMusicBitrate()`

UnsetMaxStaticMusicBitrate ensures that no value is present for MaxStaticMusicBitrate, not even an explicit nil
### GetSonyAggregationFlags

`func (o *DeviceProfile) GetSonyAggregationFlags() string`

GetSonyAggregationFlags returns the SonyAggregationFlags field if non-nil, zero value otherwise.

### GetSonyAggregationFlagsOk

`func (o *DeviceProfile) GetSonyAggregationFlagsOk() (*string, bool)`

GetSonyAggregationFlagsOk returns a tuple with the SonyAggregationFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSonyAggregationFlags

`func (o *DeviceProfile) SetSonyAggregationFlags(v string)`

SetSonyAggregationFlags sets SonyAggregationFlags field to given value.

### HasSonyAggregationFlags

`func (o *DeviceProfile) HasSonyAggregationFlags() bool`

HasSonyAggregationFlags returns a boolean if a field has been set.

### SetSonyAggregationFlagsNil

`func (o *DeviceProfile) SetSonyAggregationFlagsNil(b bool)`

 SetSonyAggregationFlagsNil sets the value for SonyAggregationFlags to be an explicit nil

### UnsetSonyAggregationFlags
`func (o *DeviceProfile) UnsetSonyAggregationFlags()`

UnsetSonyAggregationFlags ensures that no value is present for SonyAggregationFlags, not even an explicit nil
### GetProtocolInfo

`func (o *DeviceProfile) GetProtocolInfo() string`

GetProtocolInfo returns the ProtocolInfo field if non-nil, zero value otherwise.

### GetProtocolInfoOk

`func (o *DeviceProfile) GetProtocolInfoOk() (*string, bool)`

GetProtocolInfoOk returns a tuple with the ProtocolInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolInfo

`func (o *DeviceProfile) SetProtocolInfo(v string)`

SetProtocolInfo sets ProtocolInfo field to given value.

### HasProtocolInfo

`func (o *DeviceProfile) HasProtocolInfo() bool`

HasProtocolInfo returns a boolean if a field has been set.

### SetProtocolInfoNil

`func (o *DeviceProfile) SetProtocolInfoNil(b bool)`

 SetProtocolInfoNil sets the value for ProtocolInfo to be an explicit nil

### UnsetProtocolInfo
`func (o *DeviceProfile) UnsetProtocolInfo()`

UnsetProtocolInfo ensures that no value is present for ProtocolInfo, not even an explicit nil
### GetTimelineOffsetSeconds

`func (o *DeviceProfile) GetTimelineOffsetSeconds() int32`

GetTimelineOffsetSeconds returns the TimelineOffsetSeconds field if non-nil, zero value otherwise.

### GetTimelineOffsetSecondsOk

`func (o *DeviceProfile) GetTimelineOffsetSecondsOk() (*int32, bool)`

GetTimelineOffsetSecondsOk returns a tuple with the TimelineOffsetSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimelineOffsetSeconds

`func (o *DeviceProfile) SetTimelineOffsetSeconds(v int32)`

SetTimelineOffsetSeconds sets TimelineOffsetSeconds field to given value.

### HasTimelineOffsetSeconds

`func (o *DeviceProfile) HasTimelineOffsetSeconds() bool`

HasTimelineOffsetSeconds returns a boolean if a field has been set.

### GetRequiresPlainVideoItems

`func (o *DeviceProfile) GetRequiresPlainVideoItems() bool`

GetRequiresPlainVideoItems returns the RequiresPlainVideoItems field if non-nil, zero value otherwise.

### GetRequiresPlainVideoItemsOk

`func (o *DeviceProfile) GetRequiresPlainVideoItemsOk() (*bool, bool)`

GetRequiresPlainVideoItemsOk returns a tuple with the RequiresPlainVideoItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPlainVideoItems

`func (o *DeviceProfile) SetRequiresPlainVideoItems(v bool)`

SetRequiresPlainVideoItems sets RequiresPlainVideoItems field to given value.

### HasRequiresPlainVideoItems

`func (o *DeviceProfile) HasRequiresPlainVideoItems() bool`

HasRequiresPlainVideoItems returns a boolean if a field has been set.

### GetRequiresPlainFolders

`func (o *DeviceProfile) GetRequiresPlainFolders() bool`

GetRequiresPlainFolders returns the RequiresPlainFolders field if non-nil, zero value otherwise.

### GetRequiresPlainFoldersOk

`func (o *DeviceProfile) GetRequiresPlainFoldersOk() (*bool, bool)`

GetRequiresPlainFoldersOk returns a tuple with the RequiresPlainFolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPlainFolders

`func (o *DeviceProfile) SetRequiresPlainFolders(v bool)`

SetRequiresPlainFolders sets RequiresPlainFolders field to given value.

### HasRequiresPlainFolders

`func (o *DeviceProfile) HasRequiresPlainFolders() bool`

HasRequiresPlainFolders returns a boolean if a field has been set.

### GetEnableMSMediaReceiverRegistrar

`func (o *DeviceProfile) GetEnableMSMediaReceiverRegistrar() bool`

GetEnableMSMediaReceiverRegistrar returns the EnableMSMediaReceiverRegistrar field if non-nil, zero value otherwise.

### GetEnableMSMediaReceiverRegistrarOk

`func (o *DeviceProfile) GetEnableMSMediaReceiverRegistrarOk() (*bool, bool)`

GetEnableMSMediaReceiverRegistrarOk returns a tuple with the EnableMSMediaReceiverRegistrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMSMediaReceiverRegistrar

`func (o *DeviceProfile) SetEnableMSMediaReceiverRegistrar(v bool)`

SetEnableMSMediaReceiverRegistrar sets EnableMSMediaReceiverRegistrar field to given value.

### HasEnableMSMediaReceiverRegistrar

`func (o *DeviceProfile) HasEnableMSMediaReceiverRegistrar() bool`

HasEnableMSMediaReceiverRegistrar returns a boolean if a field has been set.

### GetIgnoreTranscodeByteRangeRequests

`func (o *DeviceProfile) GetIgnoreTranscodeByteRangeRequests() bool`

GetIgnoreTranscodeByteRangeRequests returns the IgnoreTranscodeByteRangeRequests field if non-nil, zero value otherwise.

### GetIgnoreTranscodeByteRangeRequestsOk

`func (o *DeviceProfile) GetIgnoreTranscodeByteRangeRequestsOk() (*bool, bool)`

GetIgnoreTranscodeByteRangeRequestsOk returns a tuple with the IgnoreTranscodeByteRangeRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreTranscodeByteRangeRequests

`func (o *DeviceProfile) SetIgnoreTranscodeByteRangeRequests(v bool)`

SetIgnoreTranscodeByteRangeRequests sets IgnoreTranscodeByteRangeRequests field to given value.

### HasIgnoreTranscodeByteRangeRequests

`func (o *DeviceProfile) HasIgnoreTranscodeByteRangeRequests() bool`

HasIgnoreTranscodeByteRangeRequests returns a boolean if a field has been set.

### GetXmlRootAttributes

`func (o *DeviceProfile) GetXmlRootAttributes() []XmlAttribute`

GetXmlRootAttributes returns the XmlRootAttributes field if non-nil, zero value otherwise.

### GetXmlRootAttributesOk

`func (o *DeviceProfile) GetXmlRootAttributesOk() (*[]XmlAttribute, bool)`

GetXmlRootAttributesOk returns a tuple with the XmlRootAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlRootAttributes

`func (o *DeviceProfile) SetXmlRootAttributes(v []XmlAttribute)`

SetXmlRootAttributes sets XmlRootAttributes field to given value.

### HasXmlRootAttributes

`func (o *DeviceProfile) HasXmlRootAttributes() bool`

HasXmlRootAttributes returns a boolean if a field has been set.

### GetDirectPlayProfiles

`func (o *DeviceProfile) GetDirectPlayProfiles() []DirectPlayProfile`

GetDirectPlayProfiles returns the DirectPlayProfiles field if non-nil, zero value otherwise.

### GetDirectPlayProfilesOk

`func (o *DeviceProfile) GetDirectPlayProfilesOk() (*[]DirectPlayProfile, bool)`

GetDirectPlayProfilesOk returns a tuple with the DirectPlayProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPlayProfiles

`func (o *DeviceProfile) SetDirectPlayProfiles(v []DirectPlayProfile)`

SetDirectPlayProfiles sets DirectPlayProfiles field to given value.

### HasDirectPlayProfiles

`func (o *DeviceProfile) HasDirectPlayProfiles() bool`

HasDirectPlayProfiles returns a boolean if a field has been set.

### GetTranscodingProfiles

`func (o *DeviceProfile) GetTranscodingProfiles() []TranscodingProfile`

GetTranscodingProfiles returns the TranscodingProfiles field if non-nil, zero value otherwise.

### GetTranscodingProfilesOk

`func (o *DeviceProfile) GetTranscodingProfilesOk() (*[]TranscodingProfile, bool)`

GetTranscodingProfilesOk returns a tuple with the TranscodingProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingProfiles

`func (o *DeviceProfile) SetTranscodingProfiles(v []TranscodingProfile)`

SetTranscodingProfiles sets TranscodingProfiles field to given value.

### HasTranscodingProfiles

`func (o *DeviceProfile) HasTranscodingProfiles() bool`

HasTranscodingProfiles returns a boolean if a field has been set.

### GetContainerProfiles

`func (o *DeviceProfile) GetContainerProfiles() []ContainerProfile`

GetContainerProfiles returns the ContainerProfiles field if non-nil, zero value otherwise.

### GetContainerProfilesOk

`func (o *DeviceProfile) GetContainerProfilesOk() (*[]ContainerProfile, bool)`

GetContainerProfilesOk returns a tuple with the ContainerProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerProfiles

`func (o *DeviceProfile) SetContainerProfiles(v []ContainerProfile)`

SetContainerProfiles sets ContainerProfiles field to given value.

### HasContainerProfiles

`func (o *DeviceProfile) HasContainerProfiles() bool`

HasContainerProfiles returns a boolean if a field has been set.

### GetCodecProfiles

`func (o *DeviceProfile) GetCodecProfiles() []CodecProfile`

GetCodecProfiles returns the CodecProfiles field if non-nil, zero value otherwise.

### GetCodecProfilesOk

`func (o *DeviceProfile) GetCodecProfilesOk() (*[]CodecProfile, bool)`

GetCodecProfilesOk returns a tuple with the CodecProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecProfiles

`func (o *DeviceProfile) SetCodecProfiles(v []CodecProfile)`

SetCodecProfiles sets CodecProfiles field to given value.

### HasCodecProfiles

`func (o *DeviceProfile) HasCodecProfiles() bool`

HasCodecProfiles returns a boolean if a field has been set.

### GetResponseProfiles

`func (o *DeviceProfile) GetResponseProfiles() []ResponseProfile`

GetResponseProfiles returns the ResponseProfiles field if non-nil, zero value otherwise.

### GetResponseProfilesOk

`func (o *DeviceProfile) GetResponseProfilesOk() (*[]ResponseProfile, bool)`

GetResponseProfilesOk returns a tuple with the ResponseProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseProfiles

`func (o *DeviceProfile) SetResponseProfiles(v []ResponseProfile)`

SetResponseProfiles sets ResponseProfiles field to given value.

### HasResponseProfiles

`func (o *DeviceProfile) HasResponseProfiles() bool`

HasResponseProfiles returns a boolean if a field has been set.

### GetSubtitleProfiles

`func (o *DeviceProfile) GetSubtitleProfiles() []SubtitleProfile`

GetSubtitleProfiles returns the SubtitleProfiles field if non-nil, zero value otherwise.

### GetSubtitleProfilesOk

`func (o *DeviceProfile) GetSubtitleProfilesOk() (*[]SubtitleProfile, bool)`

GetSubtitleProfilesOk returns a tuple with the SubtitleProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleProfiles

`func (o *DeviceProfile) SetSubtitleProfiles(v []SubtitleProfile)`

SetSubtitleProfiles sets SubtitleProfiles field to given value.

### HasSubtitleProfiles

`func (o *DeviceProfile) HasSubtitleProfiles() bool`

HasSubtitleProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


