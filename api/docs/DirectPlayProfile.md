# DirectPlayProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **NullableString** |  | [optional] 
**AudioCodec** | Pointer to **NullableString** |  | [optional] 
**VideoCodec** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**DlnaProfileType**](DlnaProfileType.md) |  | [optional] 

## Methods

### NewDirectPlayProfile

`func NewDirectPlayProfile() *DirectPlayProfile`

NewDirectPlayProfile instantiates a new DirectPlayProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectPlayProfileWithDefaults

`func NewDirectPlayProfileWithDefaults() *DirectPlayProfile`

NewDirectPlayProfileWithDefaults instantiates a new DirectPlayProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *DirectPlayProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *DirectPlayProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *DirectPlayProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *DirectPlayProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *DirectPlayProfile) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *DirectPlayProfile) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetAudioCodec

`func (o *DirectPlayProfile) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *DirectPlayProfile) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *DirectPlayProfile) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *DirectPlayProfile) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### SetAudioCodecNil

`func (o *DirectPlayProfile) SetAudioCodecNil(b bool)`

 SetAudioCodecNil sets the value for AudioCodec to be an explicit nil

### UnsetAudioCodec
`func (o *DirectPlayProfile) UnsetAudioCodec()`

UnsetAudioCodec ensures that no value is present for AudioCodec, not even an explicit nil
### GetVideoCodec

`func (o *DirectPlayProfile) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *DirectPlayProfile) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *DirectPlayProfile) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *DirectPlayProfile) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### SetVideoCodecNil

`func (o *DirectPlayProfile) SetVideoCodecNil(b bool)`

 SetVideoCodecNil sets the value for VideoCodec to be an explicit nil

### UnsetVideoCodec
`func (o *DirectPlayProfile) UnsetVideoCodec()`

UnsetVideoCodec ensures that no value is present for VideoCodec, not even an explicit nil
### GetType

`func (o *DirectPlayProfile) GetType() DlnaProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DirectPlayProfile) GetTypeOk() (*DlnaProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DirectPlayProfile) SetType(v DlnaProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *DirectPlayProfile) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


