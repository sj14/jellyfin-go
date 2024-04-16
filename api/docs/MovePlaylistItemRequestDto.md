# MovePlaylistItemRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaylistItemId** | Pointer to **string** | Gets or sets the playlist identifier of the item. | [optional] 
**NewIndex** | Pointer to **int32** | Gets or sets the new position. | [optional] 

## Methods

### NewMovePlaylistItemRequestDto

`func NewMovePlaylistItemRequestDto() *MovePlaylistItemRequestDto`

NewMovePlaylistItemRequestDto instantiates a new MovePlaylistItemRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovePlaylistItemRequestDtoWithDefaults

`func NewMovePlaylistItemRequestDtoWithDefaults() *MovePlaylistItemRequestDto`

NewMovePlaylistItemRequestDtoWithDefaults instantiates a new MovePlaylistItemRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaylistItemId

`func (o *MovePlaylistItemRequestDto) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *MovePlaylistItemRequestDto) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *MovePlaylistItemRequestDto) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *MovePlaylistItemRequestDto) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.

### GetNewIndex

`func (o *MovePlaylistItemRequestDto) GetNewIndex() int32`

GetNewIndex returns the NewIndex field if non-nil, zero value otherwise.

### GetNewIndexOk

`func (o *MovePlaylistItemRequestDto) GetNewIndexOk() (*int32, bool)`

GetNewIndexOk returns a tuple with the NewIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewIndex

`func (o *MovePlaylistItemRequestDto) SetNewIndex(v int32)`

SetNewIndex sets NewIndex field to given value.

### HasNewIndex

`func (o *MovePlaylistItemRequestDto) HasNewIndex() bool`

HasNewIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


