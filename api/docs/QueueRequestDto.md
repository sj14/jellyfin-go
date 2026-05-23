# QueueRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemIds** | Pointer to **[]string** | Gets or sets the items to enqueue. | [optional] 
**Mode** | Pointer to [**GroupQueueMode**](GroupQueueMode.md) | Gets or sets the mode in which to add the new items. | [optional] 

## Methods

### NewQueueRequestDto

`func NewQueueRequestDto() *QueueRequestDto`

NewQueueRequestDto instantiates a new QueueRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueRequestDtoWithDefaults

`func NewQueueRequestDtoWithDefaults() *QueueRequestDto`

NewQueueRequestDtoWithDefaults instantiates a new QueueRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemIds

`func (o *QueueRequestDto) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *QueueRequestDto) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *QueueRequestDto) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *QueueRequestDto) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.

### GetMode

`func (o *QueueRequestDto) GetMode() GroupQueueMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *QueueRequestDto) GetModeOk() (*GroupQueueMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *QueueRequestDto) SetMode(v GroupQueueMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *QueueRequestDto) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


