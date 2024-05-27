# ClientCapabilitiesDtoDeviceProfile

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

### NewClientCapabilitiesDtoDeviceProfile

`func NewClientCapabilitiesDtoDeviceProfile() *ClientCapabilitiesDtoDeviceProfile`

NewClientCapabilitiesDtoDeviceProfile instantiates a new ClientCapabilitiesDtoDeviceProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCapabilitiesDtoDeviceProfileWithDefaults

`func NewClientCapabilitiesDtoDeviceProfileWithDefaults() *ClientCapabilitiesDtoDeviceProfile`

NewClientCapabilitiesDtoDeviceProfileWithDefaults instantiates a new ClientCapabilitiesDtoDeviceProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClientCapabilitiesDtoDeviceProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientCapabilitiesDtoDeviceProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientCapabilitiesDtoDeviceProfile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientCapabilitiesDtoDeviceProfile) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ClientCapabilitiesDtoDeviceProfile) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ClientCapabilitiesDtoDeviceProfile) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *ClientCapabilitiesDtoDeviceProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientCapabilitiesDtoDeviceProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientCapabilitiesDtoDeviceProfile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientCapabilitiesDtoDeviceProfile) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ClientCapabilitiesDtoDeviceProfile) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ClientCapabilitiesDtoDeviceProfile) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetMaxStreamingBitrate

`func (o *ClientCapabilitiesDtoDeviceProfile) GetMaxStreamingBitrate() int32`

GetMaxStreamingBitrate returns the MaxStreamingBitrate field if non-nil, zero value otherwise.

### GetMaxStreamingBitrateOk

`func (o *ClientCapabilitiesDtoDeviceProfile) GetMaxStreamingBitrateOk() (*int32, bool)`

GetMaxStreamingBitrateOk returns a tuple with the MaxStreamingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStreamingBitrate

`func (o *ClientCapabilitiesDtoDeviceProfile) SetMaxStreamingBitrate(v int32)`

SetMaxStreamingBitrate sets MaxStreamingBitrate field to given value.

### HasMaxStreamingBitrate

`func (o *ClientCapabilitiesDtoDeviceProfile) HasMaxStreamingBitrate() bool`

HasMaxStreamingBitrate returns a boolean if a field has been set.

### SetMaxStreamingBitrateNil

`func (o *ClientCapabilitiesDtoDeviceProfile) SetMaxStreamingBitrateNil(b bool)`

 SetMaxStreamingBitrateNil sets the value for MaxStreamingBitrate to be an explicit nil

### UnsetMaxStreamingBitrate
`func (o *ClientCapabilitiesDtoDeviceProfile) UnsetMaxStreamingBitrate()`

UnsetMaxStreamingBitrate ensures that no value is present for MaxStreamingBitrate, not even an explicit nil
### GetMaxStaticBitrate

`func (o *ClientCapabilitiesDtoDeviceProfile) GetMaxStaticBitrate() int32`

GetMaxStaticBitrate returns the MaxStaticBitrate field if non-nil, zero value otherwise.

### GetMaxStaticBitrateOk

`func (o *ClientCapabilitiesDtoDeviceProfile) GetMaxStaticBitrateOk() (*int32, bool)`

GetMaxStaticBitrateOk returns a tuple with the MaxStaticBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaticBitrate

`func (o *ClientCapabilitiesDtoDeviceProfile) SetMaxStaticBitrate(v int32)`

SetMaxStaticBitrate sets MaxStaticBitrate field to given value.

### HasMaxStaticBitrate

`func (o *ClientCapabilitiesDtoDeviceProfile) HasMaxStaticBitrate() bool`

HasMaxStaticBitrate returns a boolean if a field has been set.

### SetMaxStaticBitrateNil

`func (o *ClientCapabilitiesDtoDeviceProfile) SetMaxStaticBitrateNil(b bool)`

 SetMaxStaticBitrateNil sets the value for MaxStaticBitrate to be an explicit nil

### UnsetMaxStaticBitrate
`func (o *ClientCapabilitiesDtoDeviceProfile) UnsetMaxStaticBitrate()`

UnsetMaxStaticBitrate ensures that no value is present for MaxStaticBitrate, not even an explicit nil
### GetMusicStreamingTranscodingBitrate

`func (o *ClientCapabilitiesDtoDeviceProfile) GetMusicStreamingTranscodingBitrate() int32`

GetMusicStreamingTranscodingBitrate returns the MusicStreamingTranscodingBitrate field if non-nil, zero value otherwise.

### GetMusicStreamingTranscodingBitrateOk

`func (o *ClientCapabilitiesDtoDeviceProfile) GetMusicStreamingTranscodingBitrateOk() (*int32, bool)`

GetMusicStreamingTranscodingBitrateOk returns a tuple with the MusicStreamingTranscodingBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicStreamingTranscodingBitrate

`func (o *ClientCapabilitiesDtoDeviceProfile) SetMusicStreamingTranscodingBitrate(v int32)`

SetMusicStreamingTranscodingBitrate sets MusicStreamingTranscodingBitrate field to given value.

### HasMusicStreamingTranscodingBitrate

`func (o *ClientCapabilitiesDtoDeviceProfile) HasMusicStreamingTranscodingBitrate() bool`

HasMusicStreamingTranscodingBitrate returns a boolean if a field has been set.

### SetMusicStreamingTranscodingBitrateNil

`func (o *ClientCapabilitiesDtoDeviceProfile) SetMusicStreamingTranscodingBitrateNil(b bool)`

 SetMusicStreamingTranscodingBitrateNil sets the value for MusicStreamingTranscodingBitrate to be an explicit nil

### UnsetMusicStreamingTranscodingBitrate
`func (o *ClientCapabilitiesDtoDeviceProfile) UnsetMusicStreamingTranscodingBitrate()`

UnsetMusicStreamingTranscodingBitrate ensures that no value is present for MusicStreamingTranscodingBitrate, not even an explicit nil
### GetMaxStaticMusicBitrate

`func (o *ClientCapabilitiesDtoDeviceProfile) GetMaxStaticMusicBitrate() int32`

GetMaxStaticMusicBitrate returns the MaxStaticMusicBitrate field if non-nil, zero value otherwise.

### GetMaxStaticMusicBitrateOk

`func (o *ClientCapabilitiesDtoDeviceProfile) GetMaxStaticMusicBitrateOk() (*int32, bool)`

GetMaxStaticMusicBitrateOk returns a tuple with the MaxStaticMusicBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStaticMusicBitrate

`func (o *ClientCapabilitiesDtoDeviceProfile) SetMaxStaticMusicBitrate(v int32)`

SetMaxStaticMusicBitrate sets MaxStaticMusicBitrate field to given value.

### HasMaxStaticMusicBitrate

`func (o *ClientCapabilitiesDtoDeviceProfile) HasMaxStaticMusicBitrate() bool`

HasMaxStaticMusicBitrate returns a boolean if a field has been set.

### SetMaxStaticMusicBitrateNil

`func (o *ClientCapabilitiesDtoDeviceProfile) SetMaxStaticMusicBitrateNil(b bool)`

 SetMaxStaticMusicBitrateNil sets the value for MaxStaticMusicBitrate to be an explicit nil

### UnsetMaxStaticMusicBitrate
`func (o *ClientCapabilitiesDtoDeviceProfile) UnsetMaxStaticMusicBitrate()`

UnsetMaxStaticMusicBitrate ensures that no value is present for MaxStaticMusicBitrate, not even an explicit nil
### GetDirectPlayProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) GetDirectPlayProfiles() []DirectPlayProfile`

GetDirectPlayProfiles returns the DirectPlayProfiles field if non-nil, zero value otherwise.

### GetDirectPlayProfilesOk

`func (o *ClientCapabilitiesDtoDeviceProfile) GetDirectPlayProfilesOk() (*[]DirectPlayProfile, bool)`

GetDirectPlayProfilesOk returns a tuple with the DirectPlayProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectPlayProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) SetDirectPlayProfiles(v []DirectPlayProfile)`

SetDirectPlayProfiles sets DirectPlayProfiles field to given value.

### HasDirectPlayProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) HasDirectPlayProfiles() bool`

HasDirectPlayProfiles returns a boolean if a field has been set.

### GetTranscodingProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) GetTranscodingProfiles() []TranscodingProfile`

GetTranscodingProfiles returns the TranscodingProfiles field if non-nil, zero value otherwise.

### GetTranscodingProfilesOk

`func (o *ClientCapabilitiesDtoDeviceProfile) GetTranscodingProfilesOk() (*[]TranscodingProfile, bool)`

GetTranscodingProfilesOk returns a tuple with the TranscodingProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) SetTranscodingProfiles(v []TranscodingProfile)`

SetTranscodingProfiles sets TranscodingProfiles field to given value.

### HasTranscodingProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) HasTranscodingProfiles() bool`

HasTranscodingProfiles returns a boolean if a field has been set.

### GetContainerProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) GetContainerProfiles() []ContainerProfile`

GetContainerProfiles returns the ContainerProfiles field if non-nil, zero value otherwise.

### GetContainerProfilesOk

`func (o *ClientCapabilitiesDtoDeviceProfile) GetContainerProfilesOk() (*[]ContainerProfile, bool)`

GetContainerProfilesOk returns a tuple with the ContainerProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) SetContainerProfiles(v []ContainerProfile)`

SetContainerProfiles sets ContainerProfiles field to given value.

### HasContainerProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) HasContainerProfiles() bool`

HasContainerProfiles returns a boolean if a field has been set.

### GetCodecProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) GetCodecProfiles() []CodecProfile`

GetCodecProfiles returns the CodecProfiles field if non-nil, zero value otherwise.

### GetCodecProfilesOk

`func (o *ClientCapabilitiesDtoDeviceProfile) GetCodecProfilesOk() (*[]CodecProfile, bool)`

GetCodecProfilesOk returns a tuple with the CodecProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) SetCodecProfiles(v []CodecProfile)`

SetCodecProfiles sets CodecProfiles field to given value.

### HasCodecProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) HasCodecProfiles() bool`

HasCodecProfiles returns a boolean if a field has been set.

### GetSubtitleProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) GetSubtitleProfiles() []SubtitleProfile`

GetSubtitleProfiles returns the SubtitleProfiles field if non-nil, zero value otherwise.

### GetSubtitleProfilesOk

`func (o *ClientCapabilitiesDtoDeviceProfile) GetSubtitleProfilesOk() (*[]SubtitleProfile, bool)`

GetSubtitleProfilesOk returns a tuple with the SubtitleProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitleProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) SetSubtitleProfiles(v []SubtitleProfile)`

SetSubtitleProfiles sets SubtitleProfiles field to given value.

### HasSubtitleProfiles

`func (o *ClientCapabilitiesDtoDeviceProfile) HasSubtitleProfiles() bool`

HasSubtitleProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


