# ObjectGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] 
**Type** | Pointer to [**GroupUpdateType**](GroupUpdateType.md) | Gets the update type. | [optional] 
**Data** | Pointer to **interface{}** | Gets the update data. | [optional] 

## Methods

### NewObjectGroupUpdate

`func NewObjectGroupUpdate() *ObjectGroupUpdate`

NewObjectGroupUpdate instantiates a new ObjectGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectGroupUpdateWithDefaults

`func NewObjectGroupUpdateWithDefaults() *ObjectGroupUpdate`

NewObjectGroupUpdateWithDefaults instantiates a new ObjectGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ObjectGroupUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ObjectGroupUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ObjectGroupUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ObjectGroupUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetType

`func (o *ObjectGroupUpdate) GetType() GroupUpdateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObjectGroupUpdate) GetTypeOk() (*GroupUpdateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObjectGroupUpdate) SetType(v GroupUpdateType)`

SetType sets Type field to given value.

### HasType

`func (o *ObjectGroupUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetData

`func (o *ObjectGroupUpdate) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ObjectGroupUpdate) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ObjectGroupUpdate) SetData(v interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ObjectGroupUpdate) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *ObjectGroupUpdate) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ObjectGroupUpdate) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


