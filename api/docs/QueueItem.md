# QueueItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PlaylistItemId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewQueueItem

`func NewQueueItem() *QueueItem`

NewQueueItem instantiates a new QueueItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueItemWithDefaults

`func NewQueueItemWithDefaults() *QueueItem`

NewQueueItemWithDefaults instantiates a new QueueItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueueItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueueItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueueItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QueueItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlaylistItemId

`func (o *QueueItem) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *QueueItem) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *QueueItem) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *QueueItem) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### SetPlaylistItemIdNil

`func (o *QueueItem) SetPlaylistItemIdNil(b bool)`

 SetPlaylistItemIdNil sets the value for PlaylistItemId to be an explicit nil

### UnsetPlaylistItemId
`func (o *QueueItem) UnsetPlaylistItemId()`

UnsetPlaylistItemId ensures that no value is present for PlaylistItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


