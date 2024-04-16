# ValidatePathDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidateWritable** | Pointer to **bool** | Gets or sets a value indicating whether validate if path is writable. | [optional] 
**Path** | Pointer to **NullableString** | Gets or sets the path. | [optional] 
**IsFile** | Pointer to **NullableBool** | Gets or sets is path file. | [optional] 

## Methods

### NewValidatePathDto

`func NewValidatePathDto() *ValidatePathDto`

NewValidatePathDto instantiates a new ValidatePathDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePathDtoWithDefaults

`func NewValidatePathDtoWithDefaults() *ValidatePathDto`

NewValidatePathDtoWithDefaults instantiates a new ValidatePathDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidateWritable

`func (o *ValidatePathDto) GetValidateWritable() bool`

GetValidateWritable returns the ValidateWritable field if non-nil, zero value otherwise.

### GetValidateWritableOk

`func (o *ValidatePathDto) GetValidateWritableOk() (*bool, bool)`

GetValidateWritableOk returns a tuple with the ValidateWritable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateWritable

`func (o *ValidatePathDto) SetValidateWritable(v bool)`

SetValidateWritable sets ValidateWritable field to given value.

### HasValidateWritable

`func (o *ValidatePathDto) HasValidateWritable() bool`

HasValidateWritable returns a boolean if a field has been set.

### GetPath

`func (o *ValidatePathDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ValidatePathDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ValidatePathDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ValidatePathDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ValidatePathDto) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ValidatePathDto) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetIsFile

`func (o *ValidatePathDto) GetIsFile() bool`

GetIsFile returns the IsFile field if non-nil, zero value otherwise.

### GetIsFileOk

`func (o *ValidatePathDto) GetIsFileOk() (*bool, bool)`

GetIsFileOk returns a tuple with the IsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFile

`func (o *ValidatePathDto) SetIsFile(v bool)`

SetIsFile sets IsFile field to given value.

### HasIsFile

`func (o *ValidatePathDto) HasIsFile() bool`

HasIsFile returns a boolean if a field has been set.

### SetIsFileNil

`func (o *ValidatePathDto) SetIsFileNil(b bool)`

 SetIsFileNil sets the value for IsFile to be an explicit nil

### UnsetIsFile
`func (o *ValidatePathDto) UnsetIsFile()`

UnsetIsFile ensures that no value is present for IsFile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


