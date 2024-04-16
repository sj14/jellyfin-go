# RecommendationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]BaseItemDto**](BaseItemDto.md) |  | [optional] 
**RecommendationType** | Pointer to [**RecommendationType**](RecommendationType.md) |  | [optional] 
**BaselineItemName** | Pointer to **NullableString** |  | [optional] 
**CategoryId** | Pointer to **string** |  | [optional] 

## Methods

### NewRecommendationDto

`func NewRecommendationDto() *RecommendationDto`

NewRecommendationDto instantiates a new RecommendationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendationDtoWithDefaults

`func NewRecommendationDtoWithDefaults() *RecommendationDto`

NewRecommendationDtoWithDefaults instantiates a new RecommendationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *RecommendationDto) GetItems() []BaseItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RecommendationDto) GetItemsOk() (*[]BaseItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RecommendationDto) SetItems(v []BaseItemDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *RecommendationDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *RecommendationDto) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *RecommendationDto) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetRecommendationType

`func (o *RecommendationDto) GetRecommendationType() RecommendationType`

GetRecommendationType returns the RecommendationType field if non-nil, zero value otherwise.

### GetRecommendationTypeOk

`func (o *RecommendationDto) GetRecommendationTypeOk() (*RecommendationType, bool)`

GetRecommendationTypeOk returns a tuple with the RecommendationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationType

`func (o *RecommendationDto) SetRecommendationType(v RecommendationType)`

SetRecommendationType sets RecommendationType field to given value.

### HasRecommendationType

`func (o *RecommendationDto) HasRecommendationType() bool`

HasRecommendationType returns a boolean if a field has been set.

### GetBaselineItemName

`func (o *RecommendationDto) GetBaselineItemName() string`

GetBaselineItemName returns the BaselineItemName field if non-nil, zero value otherwise.

### GetBaselineItemNameOk

`func (o *RecommendationDto) GetBaselineItemNameOk() (*string, bool)`

GetBaselineItemNameOk returns a tuple with the BaselineItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineItemName

`func (o *RecommendationDto) SetBaselineItemName(v string)`

SetBaselineItemName sets BaselineItemName field to given value.

### HasBaselineItemName

`func (o *RecommendationDto) HasBaselineItemName() bool`

HasBaselineItemName returns a boolean if a field has been set.

### SetBaselineItemNameNil

`func (o *RecommendationDto) SetBaselineItemNameNil(b bool)`

 SetBaselineItemNameNil sets the value for BaselineItemName to be an explicit nil

### UnsetBaselineItemName
`func (o *RecommendationDto) UnsetBaselineItemName()`

UnsetBaselineItemName ensures that no value is present for BaselineItemName, not even an explicit nil
### GetCategoryId

`func (o *RecommendationDto) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *RecommendationDto) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *RecommendationDto) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *RecommendationDto) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


