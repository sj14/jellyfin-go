# DatabaseConfigurationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseType** | Pointer to **string** | Gets or Sets the type of database jellyfin should use. | [optional] 
**CustomProviderOptions** | Pointer to [**NullableCustomDatabaseOptions**](CustomDatabaseOptions.md) | Gets or sets the options required to use a custom database provider. | [optional] 
**LockingBehavior** | Pointer to [**DatabaseLockingBehaviorTypes**](DatabaseLockingBehaviorTypes.md) | Gets or Sets the kind of locking behavior jellyfin should perform. Possible options are \&quot;NoLock\&quot;, \&quot;Pessimistic\&quot;, \&quot;Optimistic\&quot;.  Defaults to \&quot;NoLock\&quot;. | [optional] 

## Methods

### NewDatabaseConfigurationOptions

`func NewDatabaseConfigurationOptions() *DatabaseConfigurationOptions`

NewDatabaseConfigurationOptions instantiates a new DatabaseConfigurationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseConfigurationOptionsWithDefaults

`func NewDatabaseConfigurationOptionsWithDefaults() *DatabaseConfigurationOptions`

NewDatabaseConfigurationOptionsWithDefaults instantiates a new DatabaseConfigurationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseType

`func (o *DatabaseConfigurationOptions) GetDatabaseType() string`

GetDatabaseType returns the DatabaseType field if non-nil, zero value otherwise.

### GetDatabaseTypeOk

`func (o *DatabaseConfigurationOptions) GetDatabaseTypeOk() (*string, bool)`

GetDatabaseTypeOk returns a tuple with the DatabaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseType

`func (o *DatabaseConfigurationOptions) SetDatabaseType(v string)`

SetDatabaseType sets DatabaseType field to given value.

### HasDatabaseType

`func (o *DatabaseConfigurationOptions) HasDatabaseType() bool`

HasDatabaseType returns a boolean if a field has been set.

### GetCustomProviderOptions

`func (o *DatabaseConfigurationOptions) GetCustomProviderOptions() CustomDatabaseOptions`

GetCustomProviderOptions returns the CustomProviderOptions field if non-nil, zero value otherwise.

### GetCustomProviderOptionsOk

`func (o *DatabaseConfigurationOptions) GetCustomProviderOptionsOk() (*CustomDatabaseOptions, bool)`

GetCustomProviderOptionsOk returns a tuple with the CustomProviderOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProviderOptions

`func (o *DatabaseConfigurationOptions) SetCustomProviderOptions(v CustomDatabaseOptions)`

SetCustomProviderOptions sets CustomProviderOptions field to given value.

### HasCustomProviderOptions

`func (o *DatabaseConfigurationOptions) HasCustomProviderOptions() bool`

HasCustomProviderOptions returns a boolean if a field has been set.

### SetCustomProviderOptionsNil

`func (o *DatabaseConfigurationOptions) SetCustomProviderOptionsNil(b bool)`

 SetCustomProviderOptionsNil sets the value for CustomProviderOptions to be an explicit nil

### UnsetCustomProviderOptions
`func (o *DatabaseConfigurationOptions) UnsetCustomProviderOptions()`

UnsetCustomProviderOptions ensures that no value is present for CustomProviderOptions, not even an explicit nil
### GetLockingBehavior

`func (o *DatabaseConfigurationOptions) GetLockingBehavior() DatabaseLockingBehaviorTypes`

GetLockingBehavior returns the LockingBehavior field if non-nil, zero value otherwise.

### GetLockingBehaviorOk

`func (o *DatabaseConfigurationOptions) GetLockingBehaviorOk() (*DatabaseLockingBehaviorTypes, bool)`

GetLockingBehaviorOk returns a tuple with the LockingBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockingBehavior

`func (o *DatabaseConfigurationOptions) SetLockingBehavior(v DatabaseLockingBehaviorTypes)`

SetLockingBehavior sets LockingBehavior field to given value.

### HasLockingBehavior

`func (o *DatabaseConfigurationOptions) HasLockingBehavior() bool`

HasLockingBehavior returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


