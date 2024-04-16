# FileSystemEntryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Gets the name. | [optional] 
**Path** | Pointer to **string** | Gets the path. | [optional] 
**Type** | Pointer to [**FileSystemEntryType**](FileSystemEntryType.md) | Gets the type. | [optional] 

## Methods

### NewFileSystemEntryInfo

`func NewFileSystemEntryInfo() *FileSystemEntryInfo`

NewFileSystemEntryInfo instantiates a new FileSystemEntryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSystemEntryInfoWithDefaults

`func NewFileSystemEntryInfoWithDefaults() *FileSystemEntryInfo`

NewFileSystemEntryInfoWithDefaults instantiates a new FileSystemEntryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FileSystemEntryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FileSystemEntryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FileSystemEntryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FileSystemEntryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *FileSystemEntryInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FileSystemEntryInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FileSystemEntryInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *FileSystemEntryInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *FileSystemEntryInfo) GetType() FileSystemEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FileSystemEntryInfo) GetTypeOk() (*FileSystemEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FileSystemEntryInfo) SetType(v FileSystemEntryType)`

SetType sets Type field to given value.

### HasType

`func (o *FileSystemEntryInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


