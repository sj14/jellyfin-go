# BackupOptionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **bool** | Gets or sets a value indicating whether the archive contains the Metadata contents. | [optional] 
**Trickplay** | Pointer to **bool** | Gets or sets a value indicating whether the archive contains the Trickplay contents. | [optional] 
**Subtitles** | Pointer to **bool** | Gets or sets a value indicating whether the archive contains the Subtitle contents. | [optional] 
**Database** | Pointer to **bool** | Gets or sets a value indicating whether the archive contains the Database contents. | [optional] 

## Methods

### NewBackupOptionsDto

`func NewBackupOptionsDto() *BackupOptionsDto`

NewBackupOptionsDto instantiates a new BackupOptionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupOptionsDtoWithDefaults

`func NewBackupOptionsDtoWithDefaults() *BackupOptionsDto`

NewBackupOptionsDtoWithDefaults instantiates a new BackupOptionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *BackupOptionsDto) GetMetadata() bool`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BackupOptionsDto) GetMetadataOk() (*bool, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BackupOptionsDto) SetMetadata(v bool)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BackupOptionsDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTrickplay

`func (o *BackupOptionsDto) GetTrickplay() bool`

GetTrickplay returns the Trickplay field if non-nil, zero value otherwise.

### GetTrickplayOk

`func (o *BackupOptionsDto) GetTrickplayOk() (*bool, bool)`

GetTrickplayOk returns a tuple with the Trickplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrickplay

`func (o *BackupOptionsDto) SetTrickplay(v bool)`

SetTrickplay sets Trickplay field to given value.

### HasTrickplay

`func (o *BackupOptionsDto) HasTrickplay() bool`

HasTrickplay returns a boolean if a field has been set.

### GetSubtitles

`func (o *BackupOptionsDto) GetSubtitles() bool`

GetSubtitles returns the Subtitles field if non-nil, zero value otherwise.

### GetSubtitlesOk

`func (o *BackupOptionsDto) GetSubtitlesOk() (*bool, bool)`

GetSubtitlesOk returns a tuple with the Subtitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitles

`func (o *BackupOptionsDto) SetSubtitles(v bool)`

SetSubtitles sets Subtitles field to given value.

### HasSubtitles

`func (o *BackupOptionsDto) HasSubtitles() bool`

HasSubtitles returns a boolean if a field has been set.

### GetDatabase

`func (o *BackupOptionsDto) GetDatabase() bool`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *BackupOptionsDto) GetDatabaseOk() (*bool, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *BackupOptionsDto) SetDatabase(v bool)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *BackupOptionsDto) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


