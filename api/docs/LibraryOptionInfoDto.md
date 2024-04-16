# LibraryOptionInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets name. | [optional] 
**DefaultEnabled** | Pointer to **bool** | Gets or sets a value indicating whether default enabled. | [optional] 

## Methods

### NewLibraryOptionInfoDto

`func NewLibraryOptionInfoDto() *LibraryOptionInfoDto`

NewLibraryOptionInfoDto instantiates a new LibraryOptionInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryOptionInfoDtoWithDefaults

`func NewLibraryOptionInfoDtoWithDefaults() *LibraryOptionInfoDto`

NewLibraryOptionInfoDtoWithDefaults instantiates a new LibraryOptionInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LibraryOptionInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LibraryOptionInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LibraryOptionInfoDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LibraryOptionInfoDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LibraryOptionInfoDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LibraryOptionInfoDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDefaultEnabled

`func (o *LibraryOptionInfoDto) GetDefaultEnabled() bool`

GetDefaultEnabled returns the DefaultEnabled field if non-nil, zero value otherwise.

### GetDefaultEnabledOk

`func (o *LibraryOptionInfoDto) GetDefaultEnabledOk() (*bool, bool)`

GetDefaultEnabledOk returns a tuple with the DefaultEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEnabled

`func (o *LibraryOptionInfoDto) SetDefaultEnabled(v bool)`

SetDefaultEnabled sets DefaultEnabled field to given value.

### HasDefaultEnabled

`func (o *LibraryOptionInfoDto) HasDefaultEnabled() bool`

HasDefaultEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


