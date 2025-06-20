# SyncPlayStateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] [readonly] 
**Data** | Pointer to [**GroupStateUpdate**](GroupStateUpdate.md) | Gets the update data. | [optional] [readonly] 
**Type** | Pointer to [**GroupUpdateType**](GroupUpdateType.md) | Enum GroupUpdateType. | [optional] [readonly] [default to GROUPUPDATETYPE_STATE_UPDATE]

## Methods

### NewSyncPlayStateUpdate

`func NewSyncPlayStateUpdate() *SyncPlayStateUpdate`

NewSyncPlayStateUpdate instantiates a new SyncPlayStateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncPlayStateUpdateWithDefaults

`func NewSyncPlayStateUpdateWithDefaults() *SyncPlayStateUpdate`

NewSyncPlayStateUpdateWithDefaults instantiates a new SyncPlayStateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SyncPlayStateUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SyncPlayStateUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SyncPlayStateUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SyncPlayStateUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetData

`func (o *SyncPlayStateUpdate) GetData() GroupStateUpdate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SyncPlayStateUpdate) GetDataOk() (*GroupStateUpdate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SyncPlayStateUpdate) SetData(v GroupStateUpdate)`

SetData sets Data field to given value.

### HasData

`func (o *SyncPlayStateUpdate) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *SyncPlayStateUpdate) GetType() GroupUpdateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyncPlayStateUpdate) GetTypeOk() (*GroupUpdateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyncPlayStateUpdate) SetType(v GroupUpdateType)`

SetType sets Type field to given value.

### HasType

`func (o *SyncPlayStateUpdate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


