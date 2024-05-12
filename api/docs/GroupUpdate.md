# GroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] [readonly] 
**Type** | Pointer to [**GroupUpdateType**](GroupUpdateType.md) | Gets the update type. | [optional] 
**Data** | Pointer to [**PlayQueueUpdate**](PlayQueueUpdate.md) | Gets the update data. | [optional] 

## Methods

### NewGroupUpdate

`func NewGroupUpdate() *GroupUpdate`

NewGroupUpdate instantiates a new GroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUpdateWithDefaults

`func NewGroupUpdateWithDefaults() *GroupUpdate`

NewGroupUpdateWithDefaults instantiates a new GroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetType

`func (o *GroupUpdate) GetType() GroupUpdateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupUpdate) GetTypeOk() (*GroupUpdateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupUpdate) SetType(v GroupUpdateType)`

SetType sets Type field to given value.

### HasType

`func (o *GroupUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *GroupUpdate) GetData() PlayQueueUpdate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GroupUpdate) GetDataOk() (*PlayQueueUpdate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GroupUpdate) SetData(v PlayQueueUpdate)`

SetData sets Data field to given value.

### HasData

`func (o *GroupUpdate) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


