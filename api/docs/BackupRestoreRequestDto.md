# BackupRestoreRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveFileName** | Pointer to **string** | Gets or Sets the name of the backup archive to restore from. Must be present in MediaBrowser.Common.Configuration.IApplicationPaths.BackupPath. | [optional] 

## Methods

### NewBackupRestoreRequestDto

`func NewBackupRestoreRequestDto() *BackupRestoreRequestDto`

NewBackupRestoreRequestDto instantiates a new BackupRestoreRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreRequestDtoWithDefaults

`func NewBackupRestoreRequestDtoWithDefaults() *BackupRestoreRequestDto`

NewBackupRestoreRequestDtoWithDefaults instantiates a new BackupRestoreRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveFileName

`func (o *BackupRestoreRequestDto) GetArchiveFileName() string`

GetArchiveFileName returns the ArchiveFileName field if non-nil, zero value otherwise.

### GetArchiveFileNameOk

`func (o *BackupRestoreRequestDto) GetArchiveFileNameOk() (*string, bool)`

GetArchiveFileNameOk returns a tuple with the ArchiveFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFileName

`func (o *BackupRestoreRequestDto) SetArchiveFileName(v string)`

SetArchiveFileName sets ArchiveFileName field to given value.

### HasArchiveFileName

`func (o *BackupRestoreRequestDto) HasArchiveFileName() bool`

HasArchiveFileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


