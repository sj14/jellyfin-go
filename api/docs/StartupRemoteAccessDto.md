# StartupRemoteAccessDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableRemoteAccess** | **bool** | Gets or sets a value indicating whether enable remote access. | 
**EnableAutomaticPortMapping** | **bool** | Gets or sets a value indicating whether enable automatic port mapping. | 

## Methods

### NewStartupRemoteAccessDto

`func NewStartupRemoteAccessDto(enableRemoteAccess bool, enableAutomaticPortMapping bool, ) *StartupRemoteAccessDto`

NewStartupRemoteAccessDto instantiates a new StartupRemoteAccessDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartupRemoteAccessDtoWithDefaults

`func NewStartupRemoteAccessDtoWithDefaults() *StartupRemoteAccessDto`

NewStartupRemoteAccessDtoWithDefaults instantiates a new StartupRemoteAccessDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableRemoteAccess

`func (o *StartupRemoteAccessDto) GetEnableRemoteAccess() bool`

GetEnableRemoteAccess returns the EnableRemoteAccess field if non-nil, zero value otherwise.

### GetEnableRemoteAccessOk

`func (o *StartupRemoteAccessDto) GetEnableRemoteAccessOk() (*bool, bool)`

GetEnableRemoteAccessOk returns a tuple with the EnableRemoteAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRemoteAccess

`func (o *StartupRemoteAccessDto) SetEnableRemoteAccess(v bool)`

SetEnableRemoteAccess sets EnableRemoteAccess field to given value.


### GetEnableAutomaticPortMapping

`func (o *StartupRemoteAccessDto) GetEnableAutomaticPortMapping() bool`

GetEnableAutomaticPortMapping returns the EnableAutomaticPortMapping field if non-nil, zero value otherwise.

### GetEnableAutomaticPortMappingOk

`func (o *StartupRemoteAccessDto) GetEnableAutomaticPortMappingOk() (*bool, bool)`

GetEnableAutomaticPortMappingOk returns a tuple with the EnableAutomaticPortMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutomaticPortMapping

`func (o *StartupRemoteAccessDto) SetEnableAutomaticPortMapping(v bool)`

SetEnableAutomaticPortMapping sets EnableAutomaticPortMapping field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


