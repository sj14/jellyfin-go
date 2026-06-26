# SystemStorageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramDataFolder** | [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the program data folder. | 
**WebFolder** | [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the web UI resources folder. | 
**ImageCacheFolder** | [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the folder where images are cached. | 
**CacheFolder** | [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the cache folder. | 
**LogFolder** | [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the folder where logfiles are saved to. | 
**InternalMetadataFolder** | [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the folder where metadata is stored. | 
**TranscodingTempFolder** | [**FolderStorageDto**](FolderStorageDto.md) | Gets or sets the Storage information of the transcoding cache. | 
**Libraries** | [**[]LibraryStorageDto**](LibraryStorageDto.md) | Gets or sets the storage informations of all libraries. | 

## Methods

### NewSystemStorageDto

`func NewSystemStorageDto(programDataFolder FolderStorageDto, webFolder FolderStorageDto, imageCacheFolder FolderStorageDto, cacheFolder FolderStorageDto, logFolder FolderStorageDto, internalMetadataFolder FolderStorageDto, transcodingTempFolder FolderStorageDto, libraries []LibraryStorageDto, ) *SystemStorageDto`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


