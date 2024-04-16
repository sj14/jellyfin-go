# BaseItemDtoQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]BaseItemDto**](BaseItemDto.md) | Gets or sets the items. | [optional] 
**TotalRecordCount** | Pointer to **int32** | Gets or sets the total number of records available. | [optional] 
**StartIndex** | Pointer to **int32** | Gets or sets the index of the first record in Items. | [optional] 

## Methods

### NewBaseItemDtoQueryResult

`func NewBaseItemDtoQueryResult() *BaseItemDtoQueryResult`

NewBaseItemDtoQueryResult instantiates a new BaseItemDtoQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseItemDtoQueryResultWithDefaults

`func NewBaseItemDtoQueryResultWithDefaults() *BaseItemDtoQueryResult`

NewBaseItemDtoQueryResultWithDefaults instantiates a new BaseItemDtoQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *BaseItemDtoQueryResult) GetItems() []BaseItemDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BaseItemDtoQueryResult) GetItemsOk() (*[]BaseItemDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BaseItemDtoQueryResult) SetItems(v []BaseItemDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *BaseItemDtoQueryResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *BaseItemDtoQueryResult) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *BaseItemDtoQueryResult) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetTotalRecordCount

`func (o *BaseItemDtoQueryResult) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *BaseItemDtoQueryResult) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *BaseItemDtoQueryResult) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *BaseItemDtoQueryResult) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.

### GetStartIndex

`func (o *BaseItemDtoQueryResult) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *BaseItemDtoQueryResult) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *BaseItemDtoQueryResult) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *BaseItemDtoQueryResult) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


