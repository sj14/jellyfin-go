# LibraryStorageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Gets or sets the Library Id. | [optional] 
**Name** | Pointer to **string** | Gets or sets the name of the library. | [optional] 
**Folders** | Pointer to [**[]FolderStorageDto**](FolderStorageDto.md) | Gets or sets the storage informations about the folders used in a library. | [optional] 

## Methods

### NewLibraryStorageDto

`func NewLibraryStorageDto() *LibraryStorageDto`

NewLibraryStorageDto instantiates a new LibraryStorageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryStorageDtoWithDefaults

`func NewLibraryStorageDtoWithDefaults() *LibraryStorageDto`

NewLibraryStorageDtoWithDefaults instantiates a new LibraryStorageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LibraryStorageDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LibraryStorageDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LibraryStorageDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LibraryStorageDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LibraryStorageDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LibraryStorageDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LibraryStorageDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LibraryStorageDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFolders

`func (o *LibraryStorageDto) GetFolders() []FolderStorageDto`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *LibraryStorageDto) GetFoldersOk() (*[]FolderStorageDto, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *LibraryStorageDto) SetFolders(v []FolderStorageDto)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *LibraryStorageDto) HasFolders() bool`

HasFolders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


