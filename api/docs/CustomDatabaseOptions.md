# CustomDatabaseOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginName** | Pointer to **string** | Gets or sets the Plugin name to search for database providers. | [optional] 
**PluginAssembly** | Pointer to **string** | Gets or sets the plugin assembly to search for providers. | [optional] 
**ConnectionString** | Pointer to **string** | Gets or sets the connection string for the custom database provider. | [optional] 
**Options** | Pointer to [**[]CustomDatabaseOption**](CustomDatabaseOption.md) | Gets or sets the list of extra options for the custom provider. | [optional] 

## Methods

### NewCustomDatabaseOptions

`func NewCustomDatabaseOptions() *CustomDatabaseOptions`

NewCustomDatabaseOptions instantiates a new CustomDatabaseOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDatabaseOptionsWithDefaults

`func NewCustomDatabaseOptionsWithDefaults() *CustomDatabaseOptions`

NewCustomDatabaseOptionsWithDefaults instantiates a new CustomDatabaseOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPluginName

`func (o *CustomDatabaseOptions) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *CustomDatabaseOptions) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *CustomDatabaseOptions) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.

### HasPluginName

`func (o *CustomDatabaseOptions) HasPluginName() bool`

HasPluginName returns a boolean if a field has been set.

### GetPluginAssembly

`func (o *CustomDatabaseOptions) GetPluginAssembly() string`

GetPluginAssembly returns the PluginAssembly field if non-nil, zero value otherwise.

### GetPluginAssemblyOk

`func (o *CustomDatabaseOptions) GetPluginAssemblyOk() (*string, bool)`

GetPluginAssemblyOk returns a tuple with the PluginAssembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginAssembly

`func (o *CustomDatabaseOptions) SetPluginAssembly(v string)`

SetPluginAssembly sets PluginAssembly field to given value.

### HasPluginAssembly

`func (o *CustomDatabaseOptions) HasPluginAssembly() bool`

HasPluginAssembly returns a boolean if a field has been set.

### GetConnectionString

`func (o *CustomDatabaseOptions) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *CustomDatabaseOptions) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *CustomDatabaseOptions) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### HasConnectionString

`func (o *CustomDatabaseOptions) HasConnectionString() bool`

HasConnectionString returns a boolean if a field has been set.

### GetOptions

`func (o *CustomDatabaseOptions) GetOptions() []CustomDatabaseOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CustomDatabaseOptions) GetOptionsOk() (*[]CustomDatabaseOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CustomDatabaseOptions) SetOptions(v []CustomDatabaseOption)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CustomDatabaseOptions) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


