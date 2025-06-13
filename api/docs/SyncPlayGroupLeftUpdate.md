# SyncPlayGroupLeftUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] [readonly] 
**Data** | Pointer to **string** | Gets the update data. | [optional] [readonly] 
**Type** | Pointer to [**GroupUpdateType**](GroupUpdateType.md) | Enum GroupUpdateType. | [optional] [readonly] [default to GROUPUPDATETYPE_GROUP_LEFT]

## Methods

### NewSyncPlayGroupLeftUpdate

`func NewSyncPlayGroupLeftUpdate() *SyncPlayGroupLeftUpdate`

NewSyncPlayGroupLeftUpdate instantiates a new SyncPlayGroupLeftUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncPlayGroupLeftUpdateWithDefaults

`func NewSyncPlayGroupLeftUpdateWithDefaults() *SyncPlayGroupLeftUpdate`

NewSyncPlayGroupLeftUpdateWithDefaults instantiates a new SyncPlayGroupLeftUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SyncPlayGroupLeftUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SyncPlayGroupLeftUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SyncPlayGroupLeftUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SyncPlayGroupLeftUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetData

`func (o *SyncPlayGroupLeftUpdate) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SyncPlayGroupLeftUpdate) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SyncPlayGroupLeftUpdate) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SyncPlayGroupLeftUpdate) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *SyncPlayGroupLeftUpdate) GetType() GroupUpdateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyncPlayGroupLeftUpdate) GetTypeOk() (*GroupUpdateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyncPlayGroupLeftUpdate) SetType(v GroupUpdateType)`

SetType sets Type field to given value.

### HasType

`func (o *SyncPlayGroupLeftUpdate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


