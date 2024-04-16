# CreatePlaylistDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name of the new playlist. | [optional] 
**Ids** | Pointer to **[]string** | Gets or sets item ids to add to the playlist. | [optional] 
**UserId** | Pointer to **NullableString** | Gets or sets the user id. | [optional] 
**MediaType** | Pointer to **NullableString** | Gets or sets the media type. | [optional] 

## Methods

### NewCreatePlaylistDto

`func NewCreatePlaylistDto() *CreatePlaylistDto`

NewCreatePlaylistDto instantiates a new CreatePlaylistDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePlaylistDtoWithDefaults

`func NewCreatePlaylistDtoWithDefaults() *CreatePlaylistDto`

NewCreatePlaylistDtoWithDefaults instantiates a new CreatePlaylistDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePlaylistDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePlaylistDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePlaylistDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatePlaylistDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreatePlaylistDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreatePlaylistDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIds

`func (o *CreatePlaylistDto) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *CreatePlaylistDto) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *CreatePlaylistDto) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *CreatePlaylistDto) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetUserId

`func (o *CreatePlaylistDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreatePlaylistDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreatePlaylistDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreatePlaylistDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *CreatePlaylistDto) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *CreatePlaylistDto) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetMediaType

`func (o *CreatePlaylistDto) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *CreatePlaylistDto) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *CreatePlaylistDto) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *CreatePlaylistDto) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *CreatePlaylistDto) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *CreatePlaylistDto) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


