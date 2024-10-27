# PlaylistDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenAccess** | Pointer to **bool** | Gets or sets a value indicating whether the playlist is publicly readable. | [optional] 
**Shares** | Pointer to [**[]PlaylistUserPermissions**](PlaylistUserPermissions.md) | Gets or sets the share permissions. | [optional] 
**ItemIds** | Pointer to **[]string** | Gets or sets the item ids. | [optional] 

## Methods

### NewPlaylistDto

`func NewPlaylistDto() *PlaylistDto`

NewPlaylistDto instantiates a new PlaylistDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaylistDtoWithDefaults

`func NewPlaylistDtoWithDefaults() *PlaylistDto`

NewPlaylistDtoWithDefaults instantiates a new PlaylistDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenAccess

`func (o *PlaylistDto) GetOpenAccess() bool`

GetOpenAccess returns the OpenAccess field if non-nil, zero value otherwise.

### GetOpenAccessOk

`func (o *PlaylistDto) GetOpenAccessOk() (*bool, bool)`

GetOpenAccessOk returns a tuple with the OpenAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAccess

`func (o *PlaylistDto) SetOpenAccess(v bool)`

SetOpenAccess sets OpenAccess field to given value.

### HasOpenAccess

`func (o *PlaylistDto) HasOpenAccess() bool`

HasOpenAccess returns a boolean if a field has been set.

### GetShares

`func (o *PlaylistDto) GetShares() []PlaylistUserPermissions`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *PlaylistDto) GetSharesOk() (*[]PlaylistUserPermissions, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *PlaylistDto) SetShares(v []PlaylistUserPermissions)`

SetShares sets Shares field to given value.

### HasShares

`func (o *PlaylistDto) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetItemIds

`func (o *PlaylistDto) GetItemIds() []string`

GetItemIds returns the ItemIds field if non-nil, zero value otherwise.

### GetItemIdsOk

`func (o *PlaylistDto) GetItemIdsOk() (*[]string, bool)`

GetItemIdsOk returns a tuple with the ItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemIds

`func (o *PlaylistDto) SetItemIds(v []string)`

SetItemIds sets ItemIds field to given value.

### HasItemIds

`func (o *PlaylistDto) HasItemIds() bool`

HasItemIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


