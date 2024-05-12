# ContainerProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**DlnaProfileType**](DlnaProfileType.md) |  | [optional] 
**Conditions** | Pointer to [**[]ProfileCondition**](ProfileCondition.md) |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 

## Methods

### NewContainerProfile

`func NewContainerProfile() *ContainerProfile`

NewContainerProfile instantiates a new ContainerProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerProfileWithDefaults

`func NewContainerProfileWithDefaults() *ContainerProfile`

NewContainerProfileWithDefaults instantiates a new ContainerProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContainerProfile) GetType() DlnaProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContainerProfile) GetTypeOk() (*DlnaProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContainerProfile) SetType(v DlnaProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *ContainerProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConditions

`func (o *ContainerProfile) GetConditions() []ProfileCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ContainerProfile) GetConditionsOk() (*[]ProfileCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ContainerProfile) SetConditions(v []ProfileCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ContainerProfile) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetContainer

`func (o *ContainerProfile) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ContainerProfile) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ContainerProfile) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ContainerProfile) HasContainer() bool`

HasContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


