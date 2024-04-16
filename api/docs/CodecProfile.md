# CodecProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CodecType**](CodecType.md) |  | [optional] 
**Conditions** | Pointer to [**[]ProfileCondition**](ProfileCondition.md) |  | [optional] 
**ApplyConditions** | Pointer to [**[]ProfileCondition**](ProfileCondition.md) |  | [optional] 
**Codec** | Pointer to **NullableString** |  | [optional] 
**Container** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCodecProfile

`func NewCodecProfile() *CodecProfile`

NewCodecProfile instantiates a new CodecProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodecProfileWithDefaults

`func NewCodecProfileWithDefaults() *CodecProfile`

NewCodecProfileWithDefaults instantiates a new CodecProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CodecProfile) GetType() CodecType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CodecProfile) GetTypeOk() (*CodecType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CodecProfile) SetType(v CodecType)`

SetType sets Type field to given value.

### HasType

`func (o *CodecProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConditions

`func (o *CodecProfile) GetConditions() []ProfileCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CodecProfile) GetConditionsOk() (*[]ProfileCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CodecProfile) SetConditions(v []ProfileCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CodecProfile) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *CodecProfile) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *CodecProfile) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetApplyConditions

`func (o *CodecProfile) GetApplyConditions() []ProfileCondition`

GetApplyConditions returns the ApplyConditions field if non-nil, zero value otherwise.

### GetApplyConditionsOk

`func (o *CodecProfile) GetApplyConditionsOk() (*[]ProfileCondition, bool)`

GetApplyConditionsOk returns a tuple with the ApplyConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyConditions

`func (o *CodecProfile) SetApplyConditions(v []ProfileCondition)`

SetApplyConditions sets ApplyConditions field to given value.

### HasApplyConditions

`func (o *CodecProfile) HasApplyConditions() bool`

HasApplyConditions returns a boolean if a field has been set.

### SetApplyConditionsNil

`func (o *CodecProfile) SetApplyConditionsNil(b bool)`

 SetApplyConditionsNil sets the value for ApplyConditions to be an explicit nil

### UnsetApplyConditions
`func (o *CodecProfile) UnsetApplyConditions()`

UnsetApplyConditions ensures that no value is present for ApplyConditions, not even an explicit nil
### GetCodec

`func (o *CodecProfile) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *CodecProfile) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *CodecProfile) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *CodecProfile) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### SetCodecNil

`func (o *CodecProfile) SetCodecNil(b bool)`

 SetCodecNil sets the value for Codec to be an explicit nil

### UnsetCodec
`func (o *CodecProfile) UnsetCodec()`

UnsetCodec ensures that no value is present for Codec, not even an explicit nil
### GetContainer

`func (o *CodecProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *CodecProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *CodecProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *CodecProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *CodecProfile) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *CodecProfile) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


