# RemoveFromPlaylistRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaylistItemIds** | Pointer to **[]string** | Gets or sets the playlist identifiers ot the items. Ignored when clearing the playlist. | [optional] 
**ClearPlaylist** | Pointer to **bool** | Gets or sets a value indicating whether the entire playlist should be cleared. | [optional] 
**ClearPlayingItem** | Pointer to **bool** | Gets or sets a value indicating whether the playing item should be removed as well. Used only when clearing the playlist. | [optional] 

## Methods

### NewRemoveFromPlaylistRequestDto

`func NewRemoveFromPlaylistRequestDto() *RemoveFromPlaylistRequestDto`

NewRemoveFromPlaylistRequestDto instantiates a new RemoveFromPlaylistRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveFromPlaylistRequestDtoWithDefaults

`func NewRemoveFromPlaylistRequestDtoWithDefaults() *RemoveFromPlaylistRequestDto`

NewRemoveFromPlaylistRequestDtoWithDefaults instantiates a new RemoveFromPlaylistRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaylistItemIds

`func (o *RemoveFromPlaylistRequestDto) GetPlaylistItemIds() []string`

GetPlaylistItemIds returns the PlaylistItemIds field if non-nil, zero value otherwise.

### GetPlaylistItemIdsOk

`func (o *RemoveFromPlaylistRequestDto) GetPlaylistItemIdsOk() (*[]string, bool)`

GetPlaylistItemIdsOk returns a tuple with the PlaylistItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemIds

`func (o *RemoveFromPlaylistRequestDto) SetPlaylistItemIds(v []string)`

SetPlaylistItemIds sets PlaylistItemIds field to given value.

### HasPlaylistItemIds

`func (o *RemoveFromPlaylistRequestDto) HasPlaylistItemIds() bool`

HasPlaylistItemIds returns a boolean if a field has been set.

### GetClearPlaylist

`func (o *RemoveFromPlaylistRequestDto) GetClearPlaylist() bool`

GetClearPlaylist returns the ClearPlaylist field if non-nil, zero value otherwise.

### GetClearPlaylistOk

`func (o *RemoveFromPlaylistRequestDto) GetClearPlaylistOk() (*bool, bool)`

GetClearPlaylistOk returns a tuple with the ClearPlaylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearPlaylist

`func (o *RemoveFromPlaylistRequestDto) SetClearPlaylist(v bool)`

SetClearPlaylist sets ClearPlaylist field to given value.

### HasClearPlaylist

`func (o *RemoveFromPlaylistRequestDto) HasClearPlaylist() bool`

HasClearPlaylist returns a boolean if a field has been set.

### GetClearPlayingItem

`func (o *RemoveFromPlaylistRequestDto) GetClearPlayingItem() bool`

GetClearPlayingItem returns the ClearPlayingItem field if non-nil, zero value otherwise.

### GetClearPlayingItemOk

`func (o *RemoveFromPlaylistRequestDto) GetClearPlayingItemOk() (*bool, bool)`

GetClearPlayingItemOk returns a tuple with the ClearPlayingItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearPlayingItem

`func (o *RemoveFromPlaylistRequestDto) SetClearPlayingItem(v bool)`

SetClearPlayingItem sets ClearPlayingItem field to given value.

### HasClearPlayingItem

`func (o *RemoveFromPlaylistRequestDto) HasClearPlayingItem() bool`

HasClearPlayingItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


