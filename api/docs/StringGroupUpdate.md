# StringGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] [readonly] 
**Type** | Pointer to [**GroupUpdateType**](GroupUpdateType.md) | Gets the update type. | [optional] 
**Data** | Pointer to **string** | Gets the update data. | [optional] 

## Methods

### NewStringGroupUpdate

`func NewStringGroupUpdate() *StringGroupUpdate`

NewStringGroupUpdate instantiates a new StringGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringGroupUpdateWithDefaults

`func NewStringGroupUpdateWithDefaults() *StringGroupUpdate`

NewStringGroupUpdateWithDefaults instantiates a new StringGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *StringGroupUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *StringGroupUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *StringGroupUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *StringGroupUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetType

`func (o *StringGroupUpdate) GetType() GroupUpdateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StringGroupUpdate) GetTypeOk() (*GroupUpdateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StringGroupUpdate) SetType(v GroupUpdateType)`

SetType sets Type field to given value.

### HasType

`func (o *StringGroupUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *StringGroupUpdate) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StringGroupUpdate) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StringGroupUpdate) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *StringGroupUpdate) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


