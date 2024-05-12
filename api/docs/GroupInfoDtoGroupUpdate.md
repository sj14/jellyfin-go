# GroupInfoDtoGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] [readonly] 
**Type** | Pointer to [**GroupUpdateType**](GroupUpdateType.md) | Gets the update type. | [optional] 
**Data** | Pointer to [**GroupInfoDto**](GroupInfoDto.md) | Gets the update data. | [optional] 

## Methods

### NewGroupInfoDtoGroupUpdate

`func NewGroupInfoDtoGroupUpdate() *GroupInfoDtoGroupUpdate`

NewGroupInfoDtoGroupUpdate instantiates a new GroupInfoDtoGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupInfoDtoGroupUpdateWithDefaults

`func NewGroupInfoDtoGroupUpdateWithDefaults() *GroupInfoDtoGroupUpdate`

NewGroupInfoDtoGroupUpdateWithDefaults instantiates a new GroupInfoDtoGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupInfoDtoGroupUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupInfoDtoGroupUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupInfoDtoGroupUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupInfoDtoGroupUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetType

`func (o *GroupInfoDtoGroupUpdate) GetType() GroupUpdateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupInfoDtoGroupUpdate) GetTypeOk() (*GroupUpdateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupInfoDtoGroupUpdate) SetType(v GroupUpdateType)`

SetType sets Type field to given value.

### HasType

`func (o *GroupInfoDtoGroupUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *GroupInfoDtoGroupUpdate) GetData() GroupInfoDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GroupInfoDtoGroupUpdate) GetDataOk() (*GroupInfoDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GroupInfoDtoGroupUpdate) SetData(v GroupInfoDto)`

SetData sets Data field to given value.

### HasData

`func (o *GroupInfoDtoGroupUpdate) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


