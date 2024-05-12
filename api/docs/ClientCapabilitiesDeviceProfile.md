# ClientCapabilitiesDeviceProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name of this device profile. | [optional] 
**Id** | Pointer to **NullableString** | Gets or sets the Id. | [optional] 
**MaxStreamingBitrate** | Pointer to **NullableInt32** | Gets or sets the maximum allowed bitrate for all streamed content. | [optional] 
**MaxStaticBitrate** | Pointer to **NullableInt32** | Gets or sets the maximum allowed bitrate for statically streamed content (&#x3D; direct played files). | [optional] 
**MusicStreamingTranscodingBitrate** | Pointer to **NullableInt32** | Gets or sets the maximum allowed bitrate for transcoded music streams. | [optional] 
**MaxStaticMusicBitrate** | Pointer to **NullableInt32** | Gets or sets the maximum allowed bitrate for statically streamed (&#x3D; direct played) music files. | [optional] 
**DirectPlayProfiles** | Pointer to [**[]DirectPlayProfile**](DirectPlayProfile.md) | Gets or sets the direct play profiles. | [optional] 
**TranscodingProfiles** | Pointer to [**[]TranscodingProfile**](TranscodingProfile.md) | Gets or sets the transcoding profiles. | [optional] 
**ContainerProfiles** | Pointer to [**[]ContainerProfile**](ContainerProfile.md) | Gets or sets the container profiles. | [optional] 
**CodecProfiles** | Pointer to [**[]CodecProfile**](CodecProfile.md) | Gets or sets the codec profiles. | [optional] 
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


