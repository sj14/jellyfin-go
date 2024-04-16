# DeviceInfoQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]DeviceInfo**](DeviceInfo.md) | Gets or sets the items. | [optional] 
**TotalRecordCount** | Pointer to **int32** | Gets or sets the total number of records available. | [optional] 
**StartIndex** | Pointer to **int32** | Gets or sets the index of the first record in Items. | [optional] 

## Methods

### NewDeviceInfoQueryResult

`func NewDeviceInfoQueryResult() *DeviceInfoQueryResult`

NewDeviceInfoQueryResult instantiates a new DeviceInfoQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInfoQueryResultWithDefaults

`func NewDeviceInfoQueryResultWithDefaults() *DeviceInfoQueryResult`

NewDeviceInfoQueryResultWithDefaults instantiates a new DeviceInfoQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DeviceInfoQueryResult) GetItems() []DeviceInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DeviceInfoQueryResult) GetItemsOk() (*[]DeviceInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DeviceInfoQueryResult) SetItems(v []DeviceInfo)`

SetItems sets Items field to given value.

### HasItems

`func (o *DeviceInfoQueryResult) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *DeviceInfoQueryResult) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *DeviceInfoQueryResult) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetTotalRecordCount

`func (o *DeviceInfoQueryResult) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *DeviceInfoQueryResult) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *DeviceInfoQueryResult) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *DeviceInfoQueryResult) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.

### GetStartIndex

`func (o *DeviceInfoQueryResult) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *DeviceInfoQueryResult) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *DeviceInfoQueryResult) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *DeviceInfoQueryResult) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


