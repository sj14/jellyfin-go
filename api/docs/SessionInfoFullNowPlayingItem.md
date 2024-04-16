# SessionInfoFullNowPlayingItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **NullableInt64** |  | [optional] 
**Container** | Pointer to **NullableString** |  | [optional] 
**IsHD** | Pointer to **bool** |  | [optional] [readonly] 
**IsShortcut** | Pointer to **bool** |  | [optional] 
**ShortcutPath** | Pointer to **NullableString** |  | [optional] 
**Width** | Pointer to **int32** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**ExtraIds** | Pointer to **[]string** |  | [optional] 
**DateLastSaved** | Pointer to **time.Time** |  | [optional] 
**RemoteTrailers** | Pointer to [**[]MediaUrl**](MediaUrl.md) | Gets or sets the remote trailers. | [optional] 
**SupportsExternalTransfer** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewSessionInfoFullNowPlayingItem

`func NewSessionInfoFullNowPlayingItem() *SessionInfoFullNowPlayingItem`

NewSessionInfoFullNowPlayingItem instantiates a new SessionInfoFullNowPlayingItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionInfoFullNowPlayingItemWithDefaults

`func NewSessionInfoFullNowPlayingItemWithDefaults() *SessionInfoFullNowPlayingItem`

NewSessionInfoFullNowPlayingItemWithDefaults instantiates a new SessionInfoFullNowPlayingItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *SessionInfoFullNowPlayingItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SessionInfoFullNowPlayingItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SessionInfoFullNowPlayingItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *SessionInfoFullNowPlayingItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *SessionInfoFullNowPlayingItem) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *SessionInfoFullNowPlayingItem) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetContainer

`func (o *SessionInfoFullNowPlayingItem) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *SessionInfoFullNowPlayingItem) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *SessionInfoFullNowPlayingItem) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *SessionInfoFullNowPlayingItem) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *SessionInfoFullNowPlayingItem) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *SessionInfoFullNowPlayingItem) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetIsHD

`func (o *SessionInfoFullNowPlayingItem) GetIsHD() bool`

GetIsHD returns the IsHD field if non-nil, zero value otherwise.

### GetIsHDOk

`func (o *SessionInfoFullNowPlayingItem) GetIsHDOk() (*bool, bool)`

GetIsHDOk returns a tuple with the IsHD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHD

`func (o *SessionInfoFullNowPlayingItem) SetIsHD(v bool)`

SetIsHD sets IsHD field to given value.

### HasIsHD

`func (o *SessionInfoFullNowPlayingItem) HasIsHD() bool`

HasIsHD returns a boolean if a field has been set.

### GetIsShortcut

`func (o *SessionInfoFullNowPlayingItem) GetIsShortcut() bool`

GetIsShortcut returns the IsShortcut field if non-nil, zero value otherwise.

### GetIsShortcutOk

`func (o *SessionInfoFullNowPlayingItem) GetIsShortcutOk() (*bool, bool)`

GetIsShortcutOk returns a tuple with the IsShortcut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShortcut

`func (o *SessionInfoFullNowPlayingItem) SetIsShortcut(v bool)`

SetIsShortcut sets IsShortcut field to given value.

### HasIsShortcut

`func (o *SessionInfoFullNowPlayingItem) HasIsShortcut() bool`

HasIsShortcut returns a boolean if a field has been set.

### GetShortcutPath

`func (o *SessionInfoFullNowPlayingItem) GetShortcutPath() string`

GetShortcutPath returns the ShortcutPath field if non-nil, zero value otherwise.

### GetShortcutPathOk

`func (o *SessionInfoFullNowPlayingItem) GetShortcutPathOk() (*string, bool)`

GetShortcutPathOk returns a tuple with the ShortcutPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutPath

`func (o *SessionInfoFullNowPlayingItem) SetShortcutPath(v string)`

SetShortcutPath sets ShortcutPath field to given value.

### HasShortcutPath

`func (o *SessionInfoFullNowPlayingItem) HasShortcutPath() bool`

HasShortcutPath returns a boolean if a field has been set.

### SetShortcutPathNil

`func (o *SessionInfoFullNowPlayingItem) SetShortcutPathNil(b bool)`

 SetShortcutPathNil sets the value for ShortcutPath to be an explicit nil

### UnsetShortcutPath
`func (o *SessionInfoFullNowPlayingItem) UnsetShortcutPath()`

UnsetShortcutPath ensures that no value is present for ShortcutPath, not even an explicit nil
### GetWidth

`func (o *SessionInfoFullNowPlayingItem) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *SessionInfoFullNowPlayingItem) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *SessionInfoFullNowPlayingItem) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *SessionInfoFullNowPlayingItem) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *SessionInfoFullNowPlayingItem) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *SessionInfoFullNowPlayingItem) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *SessionInfoFullNowPlayingItem) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *SessionInfoFullNowPlayingItem) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetExtraIds

`func (o *SessionInfoFullNowPlayingItem) GetExtraIds() []string`

GetExtraIds returns the ExtraIds field if non-nil, zero value otherwise.

### GetExtraIdsOk

`func (o *SessionInfoFullNowPlayingItem) GetExtraIdsOk() (*[]string, bool)`

GetExtraIdsOk returns a tuple with the ExtraIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraIds

`func (o *SessionInfoFullNowPlayingItem) SetExtraIds(v []string)`

SetExtraIds sets ExtraIds field to given value.

### HasExtraIds

`func (o *SessionInfoFullNowPlayingItem) HasExtraIds() bool`

HasExtraIds returns a boolean if a field has been set.

### SetExtraIdsNil

`func (o *SessionInfoFullNowPlayingItem) SetExtraIdsNil(b bool)`

 SetExtraIdsNil sets the value for ExtraIds to be an explicit nil

### UnsetExtraIds
`func (o *SessionInfoFullNowPlayingItem) UnsetExtraIds()`

UnsetExtraIds ensures that no value is present for ExtraIds, not even an explicit nil
### GetDateLastSaved

`func (o *SessionInfoFullNowPlayingItem) GetDateLastSaved() time.Time`

GetDateLastSaved returns the DateLastSaved field if non-nil, zero value otherwise.

### GetDateLastSavedOk

`func (o *SessionInfoFullNowPlayingItem) GetDateLastSavedOk() (*time.Time, bool)`

GetDateLastSavedOk returns a tuple with the DateLastSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastSaved

`func (o *SessionInfoFullNowPlayingItem) SetDateLastSaved(v time.Time)`

SetDateLastSaved sets DateLastSaved field to given value.

### HasDateLastSaved

`func (o *SessionInfoFullNowPlayingItem) HasDateLastSaved() bool`

HasDateLastSaved returns a boolean if a field has been set.

### GetRemoteTrailers

`func (o *SessionInfoFullNowPlayingItem) GetRemoteTrailers() []MediaUrl`

GetRemoteTrailers returns the RemoteTrailers field if non-nil, zero value otherwise.

### GetRemoteTrailersOk

`func (o *SessionInfoFullNowPlayingItem) GetRemoteTrailersOk() (*[]MediaUrl, bool)`

GetRemoteTrailersOk returns a tuple with the RemoteTrailers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTrailers

`func (o *SessionInfoFullNowPlayingItem) SetRemoteTrailers(v []MediaUrl)`

SetRemoteTrailers sets RemoteTrailers field to given value.

### HasRemoteTrailers

`func (o *SessionInfoFullNowPlayingItem) HasRemoteTrailers() bool`

HasRemoteTrailers returns a boolean if a field has been set.

### SetRemoteTrailersNil

`func (o *SessionInfoFullNowPlayingItem) SetRemoteTrailersNil(b bool)`

 SetRemoteTrailersNil sets the value for RemoteTrailers to be an explicit nil

### UnsetRemoteTrailers
`func (o *SessionInfoFullNowPlayingItem) UnsetRemoteTrailers()`

UnsetRemoteTrailers ensures that no value is present for RemoteTrailers, not even an explicit nil
### GetSupportsExternalTransfer

`func (o *SessionInfoFullNowPlayingItem) GetSupportsExternalTransfer() bool`

GetSupportsExternalTransfer returns the SupportsExternalTransfer field if non-nil, zero value otherwise.

### GetSupportsExternalTransferOk

`func (o *SessionInfoFullNowPlayingItem) GetSupportsExternalTransferOk() (*bool, bool)`

GetSupportsExternalTransferOk returns a tuple with the SupportsExternalTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsExternalTransfer

`func (o *SessionInfoFullNowPlayingItem) SetSupportsExternalTransfer(v bool)`

SetSupportsExternalTransfer sets SupportsExternalTransfer field to given value.

### HasSupportsExternalTransfer

`func (o *SessionInfoFullNowPlayingItem) HasSupportsExternalTransfer() bool`

HasSupportsExternalTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


