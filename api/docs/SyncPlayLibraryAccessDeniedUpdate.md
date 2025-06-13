# SyncPlayLibraryAccessDeniedUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Gets the group identifier. | [optional] [readonly] 
**Data** | Pointer to **string** | Gets the update data. | [optional] [readonly] 
**Type** | Pointer to [**GroupUpdateType**](GroupUpdateType.md) | Enum GroupUpdateType. | [optional] [readonly] [default to GROUPUPDATETYPE_LIBRARY_ACCESS_DENIED]

## Methods

### NewSyncPlayLibraryAccessDeniedUpdate

`func NewSyncPlayLibraryAccessDeniedUpdate() *SyncPlayLibraryAccessDeniedUpdate`

NewSyncPlayLibraryAccessDeniedUpdate instantiates a new SyncPlayLibraryAccessDeniedUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncPlayLibraryAccessDeniedUpdateWithDefaults

`func NewSyncPlayLibraryAccessDeniedUpdateWithDefaults() *SyncPlayLibraryAccessDeniedUpdate`

NewSyncPlayLibraryAccessDeniedUpdateWithDefaults instantiates a new SyncPlayLibraryAccessDeniedUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SyncPlayLibraryAccessDeniedUpdate) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SyncPlayLibraryAccessDeniedUpdate) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SyncPlayLibraryAccessDeniedUpdate) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *SyncPlayLibraryAccessDeniedUpdate) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetData

`func (o *SyncPlayLibraryAccessDeniedUpdate) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SyncPlayLibraryAccessDeniedUpdate) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SyncPlayLibraryAccessDeniedUpdate) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *SyncPlayLibraryAccessDeniedUpdate) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *SyncPlayLibraryAccessDeniedUpdate) GetType() GroupUpdateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyncPlayLibraryAccessDeniedUpdate) GetTypeOk() (*GroupUpdateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyncPlayLibraryAccessDeniedUpdate) SetType(v GroupUpdateType)`

SetType sets Type field to given value.

### HasType

`func (o *SyncPlayLibraryAccessDeniedUpdate) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


