# SetPlaylistItemRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaylistItemId** | Pointer to **string** | Gets or sets the playlist identifier of the playing item. | [optional] 

## Methods

### NewSetPlaylistItemRequestDto

`func NewSetPlaylistItemRequestDto() *SetPlaylistItemRequestDto`

NewSetPlaylistItemRequestDto instantiates a new SetPlaylistItemRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetPlaylistItemRequestDtoWithDefaults

`func NewSetPlaylistItemRequestDtoWithDefaults() *SetPlaylistItemRequestDto`

NewSetPlaylistItemRequestDtoWithDefaults instantiates a new SetPlaylistItemRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaylistItemId

`func (o *SetPlaylistItemRequestDto) GetPlaylistItemId() string`

GetPlaylistItemId returns the PlaylistItemId field if non-nil, zero value otherwise.

### GetPlaylistItemIdOk

`func (o *SetPlaylistItemRequestDto) GetPlaylistItemIdOk() (*string, bool)`

GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistItemId

`func (o *SetPlaylistItemRequestDto) SetPlaylistItemId(v string)`

SetPlaylistItemId sets PlaylistItemId field to given value.

### HasPlaylistItemId

`func (o *SetPlaylistItemRequestDto) HasPlaylistItemId() bool`

HasPlaylistItemId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


