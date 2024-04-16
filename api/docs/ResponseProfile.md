# ResponseProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **NullableString** |  | [optional] 
**AudioCodec** | Pointer to **NullableString** |  | [optional] 
**VideoCodec** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**DlnaProfileType**](DlnaProfileType.md) |  | [optional] 
**OrgPn** | Pointer to **NullableString** |  | [optional] 
**MimeType** | Pointer to **NullableString** |  | [optional] 
**Conditions** | Pointer to [**[]ProfileCondition**](ProfileCondition.md) |  | [optional] 

## Methods

### NewResponseProfile

`func NewResponseProfile() *ResponseProfile`

NewResponseProfile instantiates a new ResponseProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseProfileWithDefaults

`func NewResponseProfileWithDefaults() *ResponseProfile`

NewResponseProfileWithDefaults instantiates a new ResponseProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *ResponseProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ResponseProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ResponseProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ResponseProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *ResponseProfile) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *ResponseProfile) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetAudioCodec

`func (o *ResponseProfile) GetAudioCodec() string`

GetAudioCodec returns the AudioCodec field if non-nil, zero value otherwise.

### GetAudioCodecOk

`func (o *ResponseProfile) GetAudioCodecOk() (*string, bool)`

GetAudioCodecOk returns a tuple with the AudioCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioCodec

`func (o *ResponseProfile) SetAudioCodec(v string)`

SetAudioCodec sets AudioCodec field to given value.

### HasAudioCodec

`func (o *ResponseProfile) HasAudioCodec() bool`

HasAudioCodec returns a boolean if a field has been set.

### SetAudioCodecNil

`func (o *ResponseProfile) SetAudioCodecNil(b bool)`

 SetAudioCodecNil sets the value for AudioCodec to be an explicit nil

### UnsetAudioCodec
`func (o *ResponseProfile) UnsetAudioCodec()`

UnsetAudioCodec ensures that no value is present for AudioCodec, not even an explicit nil
### GetVideoCodec

`func (o *ResponseProfile) GetVideoCodec() string`

GetVideoCodec returns the VideoCodec field if non-nil, zero value otherwise.

### GetVideoCodecOk

`func (o *ResponseProfile) GetVideoCodecOk() (*string, bool)`

GetVideoCodecOk returns a tuple with the VideoCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCodec

`func (o *ResponseProfile) SetVideoCodec(v string)`

SetVideoCodec sets VideoCodec field to given value.

### HasVideoCodec

`func (o *ResponseProfile) HasVideoCodec() bool`

HasVideoCodec returns a boolean if a field has been set.

### SetVideoCodecNil

`func (o *ResponseProfile) SetVideoCodecNil(b bool)`

 SetVideoCodecNil sets the value for VideoCodec to be an explicit nil

### UnsetVideoCodec
`func (o *ResponseProfile) UnsetVideoCodec()`

UnsetVideoCodec ensures that no value is present for VideoCodec, not even an explicit nil
### GetType

`func (o *ResponseProfile) GetType() DlnaProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponseProfile) GetTypeOk() (*DlnaProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponseProfile) SetType(v DlnaProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *ResponseProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrgPn

`func (o *ResponseProfile) GetOrgPn() string`

GetOrgPn returns the OrgPn field if non-nil, zero value otherwise.

### GetOrgPnOk

`func (o *ResponseProfile) GetOrgPnOk() (*string, bool)`

GetOrgPnOk returns a tuple with the OrgPn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgPn

`func (o *ResponseProfile) SetOrgPn(v string)`

SetOrgPn sets OrgPn field to given value.

### HasOrgPn

`func (o *ResponseProfile) HasOrgPn() bool`

HasOrgPn returns a boolean if a field has been set.

### SetOrgPnNil

`func (o *ResponseProfile) SetOrgPnNil(b bool)`

 SetOrgPnNil sets the value for OrgPn to be an explicit nil

### UnsetOrgPn
`func (o *ResponseProfile) UnsetOrgPn()`

UnsetOrgPn ensures that no value is present for OrgPn, not even an explicit nil
### GetMimeType

`func (o *ResponseProfile) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *ResponseProfile) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *ResponseProfile) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *ResponseProfile) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### SetMimeTypeNil

`func (o *ResponseProfile) SetMimeTypeNil(b bool)`

 SetMimeTypeNil sets the value for MimeType to be an explicit nil

### UnsetMimeType
`func (o *ResponseProfile) UnsetMimeType()`

UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil
### GetConditions

`func (o *ResponseProfile) GetConditions() []ProfileCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ResponseProfile) GetConditionsOk() (*[]ProfileCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ResponseProfile) SetConditions(v []ProfileCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ResponseProfile) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *ResponseProfile) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *ResponseProfile) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


