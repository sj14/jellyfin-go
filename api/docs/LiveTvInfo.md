# LiveTvInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to [**[]LiveTvServiceInfo**](LiveTvServiceInfo.md) | Gets or sets the services. | [optional] 
**IsEnabled** | Pointer to **bool** | Gets or sets a value indicating whether this instance is enabled. | [optional] 
**EnabledUsers** | Pointer to **[]string** | Gets or sets the enabled users. | [optional] 

## Methods

### NewLiveTvInfo

`func NewLiveTvInfo() *LiveTvInfo`

NewLiveTvInfo instantiates a new LiveTvInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveTvInfoWithDefaults

`func NewLiveTvInfoWithDefaults() *LiveTvInfo`

NewLiveTvInfoWithDefaults instantiates a new LiveTvInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *LiveTvInfo) GetServices() []LiveTvServiceInfo`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *LiveTvInfo) GetServicesOk() (*[]LiveTvServiceInfo, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *LiveTvInfo) SetServices(v []LiveTvServiceInfo)`

SetServices sets Services field to given value.

### HasServices

`func (o *LiveTvInfo) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetIsEnabled

`func (o *LiveTvInfo) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *LiveTvInfo) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *LiveTvInfo) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *LiveTvInfo) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetEnabledUsers

`func (o *LiveTvInfo) GetEnabledUsers() []string`

GetEnabledUsers returns the EnabledUsers field if non-nil, zero value otherwise.

### GetEnabledUsersOk

`func (o *LiveTvInfo) GetEnabledUsersOk() (*[]string, bool)`

GetEnabledUsersOk returns a tuple with the EnabledUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledUsers

`func (o *LiveTvInfo) SetEnabledUsers(v []string)`

SetEnabledUsers sets EnabledUsers field to given value.

### HasEnabledUsers

`func (o *LiveTvInfo) HasEnabledUsers() bool`

HasEnabledUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


