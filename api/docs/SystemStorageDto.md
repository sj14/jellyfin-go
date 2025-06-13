# SystemStorageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramDataFolder** | Pointer to [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the program data folder. | [optional] 
**WebFolder** | Pointer to [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the web UI resources folder. | [optional] 
**ImageCacheFolder** | Pointer to [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the folder where images are cached. | [optional] 
**CacheFolder** | Pointer to [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the cache folder. | [optional] 
**LogFolder** | Pointer to [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the folder where logfiles are saved to. | [optional] 
**InternalMetadataFolder** | Pointer to [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the folder where metadata is stored. | [optional] 
**TranscodingTempFolder** | Pointer to [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the transcoding cache. | [optional] 
**Libraries** | Pointer to [**[]LibraryStorageDto**](LibraryStorageDto.md) | Gets or sets the storage informations of all libraries. | [optional] 

## Methods

### NewSystemStorageDto

`func NewSystemStorageDto() *SystemStorageDto`

NewSystemStorageDto instantiates a new SystemStorageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemStorageDtoWithDefaults

`func NewSystemStorageDtoWithDefaults() *SystemStorageDto`

NewSystemStorageDtoWithDefaults instantiates a new SystemStorageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramDataFolder

`func (o *SystemStorageDto) GetProgramDataFolder() FolderStorageDto`

GetProgramDataFolder returns the ProgramDataFolder field if non-nil, zero value otherwise.

### GetProgramDataFolderOk

`func (o *SystemStorageDto) GetProgramDataFolderOk() (*FolderStorageDto, bool)`

GetProgramDataFolderOk returns a tuple with the ProgramDataFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramDataFolder

`func (o *SystemStorageDto) SetProgramDataFolder(v FolderStorageDto)`

SetProgramDataFolder sets ProgramDataFolder field to given value.

### HasProgramDataFolder

`func (o *SystemStorageDto) HasProgramDataFolder() bool`

HasProgramDataFolder returns a boolean if a field has been set.

### GetWebFolder

`func (o *SystemStorageDto) GetWebFolder() FolderStorageDto`

GetWebFolder returns the WebFolder field if non-nil, zero value otherwise.

### GetWebFolderOk

`func (o *SystemStorageDto) GetWebFolderOk() (*FolderStorageDto, bool)`

GetWebFolderOk returns a tuple with the WebFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebFolder

`func (o *SystemStorageDto) SetWebFolder(v FolderStorageDto)`

SetWebFolder sets WebFolder field to given value.

### HasWebFolder

`func (o *SystemStorageDto) HasWebFolder() bool`

HasWebFolder returns a boolean if a field has been set.

### GetImageCacheFolder

`func (o *SystemStorageDto) GetImageCacheFolder() FolderStorageDto`

GetImageCacheFolder returns the ImageCacheFolder field if non-nil, zero value otherwise.

### GetImageCacheFolderOk

`func (o *SystemStorageDto) GetImageCacheFolderOk() (*FolderStorageDto, bool)`

GetImageCacheFolderOk returns a tuple with the ImageCacheFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCacheFolder

`func (o *SystemStorageDto) SetImageCacheFolder(v FolderStorageDto)`

SetImageCacheFolder sets ImageCacheFolder field to given value.

### HasImageCacheFolder

`func (o *SystemStorageDto) HasImageCacheFolder() bool`

HasImageCacheFolder returns a boolean if a field has been set.

### GetCacheFolder

`func (o *SystemStorageDto) GetCacheFolder() FolderStorageDto`

GetCacheFolder returns the CacheFolder field if non-nil, zero value otherwise.

### GetCacheFolderOk

`func (o *SystemStorageDto) GetCacheFolderOk() (*FolderStorageDto, bool)`

GetCacheFolderOk returns a tuple with the CacheFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheFolder

`func (o *SystemStorageDto) SetCacheFolder(v FolderStorageDto)`

SetCacheFolder sets CacheFolder field to given value.

### HasCacheFolder

`func (o *SystemStorageDto) HasCacheFolder() bool`

HasCacheFolder returns a boolean if a field has been set.

### GetLogFolder

`func (o *SystemStorageDto) GetLogFolder() FolderStorageDto`

GetLogFolder returns the LogFolder field if non-nil, zero value otherwise.

### GetLogFolderOk

`func (o *SystemStorageDto) GetLogFolderOk() (*FolderStorageDto, bool)`

GetLogFolderOk returns a tuple with the LogFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogFolder

`func (o *SystemStorageDto) SetLogFolder(v FolderStorageDto)`

SetLogFolder sets LogFolder field to given value.

### HasLogFolder

`func (o *SystemStorageDto) HasLogFolder() bool`

HasLogFolder returns a boolean if a field has been set.

### GetInternalMetadataFolder

`func (o *SystemStorageDto) GetInternalMetadataFolder() FolderStorageDto`

GetInternalMetadataFolder returns the InternalMetadataFolder field if non-nil, zero value otherwise.

### GetInternalMetadataFolderOk

`func (o *SystemStorageDto) GetInternalMetadataFolderOk() (*FolderStorageDto, bool)`

GetInternalMetadataFolderOk returns a tuple with the InternalMetadataFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalMetadataFolder

`func (o *SystemStorageDto) SetInternalMetadataFolder(v FolderStorageDto)`

SetInternalMetadataFolder sets InternalMetadataFolder field to given value.

### HasInternalMetadataFolder

`func (o *SystemStorageDto) HasInternalMetadataFolder() bool`

HasInternalMetadataFolder returns a boolean if a field has been set.

### GetTranscodingTempFolder

`func (o *SystemStorageDto) GetTranscodingTempFolder() FolderStorageDto`

GetTranscodingTempFolder returns the TranscodingTempFolder field if non-nil, zero value otherwise.

### GetTranscodingTempFolderOk

`func (o *SystemStorageDto) GetTranscodingTempFolderOk() (*FolderStorageDto, bool)`

GetTranscodingTempFolderOk returns a tuple with the TranscodingTempFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingTempFolder

`func (o *SystemStorageDto) SetTranscodingTempFolder(v FolderStorageDto)`

SetTranscodingTempFolder sets TranscodingTempFolder field to given value.

### HasTranscodingTempFolder

`func (o *SystemStorageDto) HasTranscodingTempFolder() bool`

HasTranscodingTempFolder returns a boolean if a field has been set.

### GetLibraries

`func (o *SystemStorageDto) GetLibraries() []LibraryStorageDto`

GetLibraries returns the Libraries field if non-nil, zero value otherwise.

### GetLibrariesOk

`func (o *SystemStorageDto) GetLibrariesOk() (*[]LibraryStorageDto, bool)`

GetLibrariesOk returns a tuple with the Libraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraries

`func (o *SystemStorageDto) SetLibraries(v []LibraryStorageDto)`

SetLibraries sets Libraries field to given value.

### HasLibraries

`func (o *SystemStorageDto) HasLibraries() bool`

HasLibraries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


