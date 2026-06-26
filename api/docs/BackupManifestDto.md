# BackupManifestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerVersion** | **string** | Gets or sets the jellyfin version this backup was created with. | 
**BackupEngineVersion** | **string** | Gets or sets the backup engine version this backup was created with. | 
**DateCreated** | **time.Time** | Gets or sets the date this backup was created with. | 
**Path** | **string** | Gets or sets the path to the backup on the system. | 
**Options** | [**BackupOptionsDto**](BackupOptionsDto.md) | Gets or sets the contents of the backup archive. | 

## Methods

### NewBackupManifestDto

`func NewBackupManifestDto(serverVersion string, backupEngineVersion string, dateCreated time.Time, path string, options BackupOptionsDto, ) *BackupManifestDto`

NewBackupManifestDto instantiates a new BackupManifestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupManifestDtoWithDefaults

`func NewBackupManifestDtoWithDefaults() *BackupManifestDto`

NewBackupManifestDtoWithDefaults instantiates a new BackupManifestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerVersion

`func (o *BackupManifestDto) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *BackupManifestDto) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *BackupManifestDto) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.


### GetBackupEngineVersion

`func (o *BackupManifestDto) GetBackupEngineVersion() string`

GetBackupEngineVersion returns the BackupEngineVersion field if non-nil, zero value otherwise.

### GetBackupEngineVersionOk

`func (o *BackupManifestDto) GetBackupEngineVersionOk() (*string, bool)`

GetBackupEngineVersionOk returns a tuple with the BackupEngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEngineVersion

`func (o *BackupManifestDto) SetBackupEngineVersion(v string)`

SetBackupEngineVersion sets BackupEngineVersion field to given value.


### GetDateCreated

`func (o *BackupManifestDto) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *BackupManifestDto) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *BackupManifestDto) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.


### GetPath

`func (o *BackupManifestDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BackupManifestDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BackupManifestDto) SetPath(v string)`

SetPath sets Path field to given value.


### GetOptions

`func (o *BackupManifestDto) GetOptions() BackupOptionsDto`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BackupManifestDto) GetOptionsOk() (*BackupOptionsDto, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BackupManifestDto) SetOptions(v BackupOptionsDto)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


