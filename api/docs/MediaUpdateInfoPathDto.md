# MediaUpdateInfoPathDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **NullableString** | Gets or sets media path. | [optional] 
**UpdateType** | Pointer to **NullableString** | Gets or sets media update type.  Created, Modified, Deleted. | [optional] 

## Methods

### NewMediaUpdateInfoPathDto

`func NewMediaUpdateInfoPathDto() *MediaUpdateInfoPathDto`

NewMediaUpdateInfoPathDto instantiates a new MediaUpdateInfoPathDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaUpdateInfoPathDtoWithDefaults

`func NewMediaUpdateInfoPathDtoWithDefaults() *MediaUpdateInfoPathDto`

NewMediaUpdateInfoPathDtoWithDefaults instantiates a new MediaUpdateInfoPathDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *MediaUpdateInfoPathDto) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MediaUpdateInfoPathDto) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MediaUpdateInfoPathDto) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MediaUpdateInfoPathDto) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *MediaUpdateInfoPathDto) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *MediaUpdateInfoPathDto) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetUpdateType

`func (o *MediaUpdateInfoPathDto) GetUpdateType() string`

GetUpdateType returns the UpdateType field if non-nil, zero value otherwise.

### GetUpdateTypeOk

`func (o *MediaUpdateInfoPathDto) GetUpdateTypeOk() (*string, bool)`

GetUpdateTypeOk returns a tuple with the UpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateType

`func (o *MediaUpdateInfoPathDto) SetUpdateType(v string)`

SetUpdateType sets UpdateType field to given value.

### HasUpdateType

`func (o *MediaUpdateInfoPathDto) HasUpdateType() bool`

HasUpdateType returns a boolean if a field has been set.

### SetUpdateTypeNil

`func (o *MediaUpdateInfoPathDto) SetUpdateTypeNil(b bool)`

 SetUpdateTypeNil sets the value for UpdateType to be an explicit nil

### UnsetUpdateType
`func (o *MediaUpdateInfoPathDto) UnsetUpdateType()`

UnsetUpdateType ensures that no value is present for UpdateType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


