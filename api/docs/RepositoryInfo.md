# RepositoryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**Url** | Pointer to **NullableString** | Gets or sets the URL. | [optional] 
**Enabled** | Pointer to **bool** | Gets or sets a value indicating whether the repository is enabled. | [optional] 

## Methods

### NewRepositoryInfo

`func NewRepositoryInfo() *RepositoryInfo`

NewRepositoryInfo instantiates a new RepositoryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryInfoWithDefaults

`func NewRepositoryInfoWithDefaults() *RepositoryInfo`

NewRepositoryInfoWithDefaults instantiates a new RepositoryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RepositoryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RepositoryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RepositoryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RepositoryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RepositoryInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RepositoryInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUrl

`func (o *RepositoryInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RepositoryInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RepositoryInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RepositoryInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *RepositoryInfo) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *RepositoryInfo) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetEnabled

`func (o *RepositoryInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RepositoryInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RepositoryInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RepositoryInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


