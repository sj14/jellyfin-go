# ImageOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ImageType**](ImageType.md) | Gets or sets the type. | [optional] 
**Limit** | Pointer to **int32** | Gets or sets the limit. | [optional] 
**MinWidth** | Pointer to **int32** | Gets or sets the minimum width. | [optional] 

## Methods

### NewImageOption

`func NewImageOption() *ImageOption`

NewImageOption instantiates a new ImageOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageOptionWithDefaults

`func NewImageOptionWithDefaults() *ImageOption`

NewImageOptionWithDefaults instantiates a new ImageOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ImageOption) GetType() ImageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageOption) GetTypeOk() (*ImageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageOption) SetType(v ImageType)`

SetType sets Type field to given value.

### HasType

`func (o *ImageOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLimit

`func (o *ImageOption) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ImageOption) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ImageOption) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ImageOption) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMinWidth

`func (o *ImageOption) GetMinWidth() int32`

GetMinWidth returns the MinWidth field if non-nil, zero value otherwise.

### GetMinWidthOk

`func (o *ImageOption) GetMinWidthOk() (*int32, bool)`

GetMinWidthOk returns a tuple with the MinWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWidth

`func (o *ImageOption) SetMinWidth(v int32)`

SetMinWidth sets MinWidth field to given value.

### HasMinWidth

`func (o *ImageOption) HasMinWidth() bool`

HasMinWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


