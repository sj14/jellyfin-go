# UpdatePlaylistDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name of the new playlist. | [optional] 
**Ids** | Pointer to **[]string** | Gets or sets item ids of the playlist. | [optional] 
**Users** | Pointer to [**[]PlaylistUserPermissions**](PlaylistUserPermissions.md) | Gets or sets the playlist users. | [optional] 
**IsPublic** | Pointer to **NullableBool** | Gets or sets a value indicating whether the playlist is public. | [optional] 

## Methods

### NewUpdatePlaylistDto

`func NewUpdatePlaylistDto() *UpdatePlaylistDto`

NewUpdatePlaylistDto instantiates a new UpdatePlaylistDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePlaylistDtoWithDefaults

`func NewUpdatePlaylistDtoWithDefaults() *UpdatePlaylistDto`

NewUpdatePlaylistDtoWithDefaults instantiates a new UpdatePlaylistDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePlaylistDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePlaylistDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePlaylistDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePlaylistDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdatePlaylistDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdatePlaylistDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIds

`func (o *UpdatePlaylistDto) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *UpdatePlaylistDto) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *UpdatePlaylistDto) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *UpdatePlaylistDto) HasIds() bool`

HasIds returns a boolean if a field has been set.

### SetIdsNil

`func (o *UpdatePlaylistDto) SetIdsNil(b bool)`

 SetIdsNil sets the value for Ids to be an explicit nil

### UnsetIds
`func (o *UpdatePlaylistDto) UnsetIds()`

UnsetIds ensures that no value is present for Ids, not even an explicit nil
### GetUsers

`func (o *UpdatePlaylistDto) GetUsers() []PlaylistUserPermissions`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdatePlaylistDto) GetUsersOk() (*[]PlaylistUserPermissions, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdatePlaylistDto) SetUsers(v []PlaylistUserPermissions)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UpdatePlaylistDto) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *UpdatePlaylistDto) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *UpdatePlaylistDto) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil
### GetIsPublic

`func (o *UpdatePlaylistDto) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *UpdatePlaylistDto) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *UpdatePlaylistDto) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *UpdatePlaylistDto) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### SetIsPublicNil

`func (o *UpdatePlaylistDto) SetIsPublicNil(b bool)`

 SetIsPublicNil sets the value for IsPublic to be an explicit nil

### UnsetIsPublic
`func (o *UpdatePlaylistDto) UnsetIsPublic()`

UnsetIsPublic ensures that no value is present for IsPublic, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


