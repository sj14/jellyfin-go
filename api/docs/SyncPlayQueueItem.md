# SyncPlayQueueItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **string** | Gets the item identifier. | [optional] 
**PlaylistItemId** | Pointer to **string** | Gets the playlist identifier of the item. | [optional] [readonly] 

## Methods

### NewSyncPlayQueueItem

`func NewSyncPlayQueueItem() *SyncPlayQueueItem`

NewSyncPlayQueueItem instantiates a new SyncPlayQueueItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyncPlayQueueItemWithDefaults

`func NewSyncPlayQueueItemWithDefaults() *SyncPlayQueueItem`

NewSyncPlayQueueItemWithDefaults instantiates a new SyncPlayQueueItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *SyncPlayQueueItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *SyncPlayQueueItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *SyncPlayQueueItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *SyncPlayQueueItem) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *SyncPlayQueueItem) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *SyncPlayQueueItem) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *SyncPlayQueueItem) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *SyncPlayQueueItem) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


