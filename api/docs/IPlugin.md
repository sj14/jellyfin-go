# IPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets the name of the plugin. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | Gets the Description. | [optional] [readonly] 
**Id** | Pointer to **string** | Gets the unique id. | [optional] [readonly] 
**Version** | Pointer to **NullableString** | Gets the plugin version. | [optional] [readonly] 
**AssemblyFilePath** | Pointer to **NullableString** | Gets the path to the assembly file. | [optional] [readonly] 
**CanUninstall** | Pointer to **bool** | Gets a value indicating whether the plugin can be uninstalled. | [optional] [readonly] 
**DataFolderPath** | Pointer to **NullableString** | Gets the full path to the data folder, where the plugin can store any miscellaneous files needed. | [optional] [readonly] 

## Methods

### NewIPlugin

`func NewIPlugin() *IPlugin`

NewIPlugin instantiates a new IPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPluginWithDefaults

`func NewIPluginWithDefaults() *IPlugin`

NewIPluginWithDefaults instantiates a new IPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IPlugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPlugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPlugin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IPlugin) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *IPlugin) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IPlugin) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *IPlugin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPlugin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPlugin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPlugin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *IPlugin) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *IPlugin) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *IPlugin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPlugin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPlugin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IPlugin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *IPlugin) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IPlugin) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IPlugin) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IPlugin) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *IPlugin) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *IPlugin) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetAssemblyFilePath

`func (o *IPlugin) GetAssemblyFilePath() string`

GetAssemblyFilePath returns the AssemblyFilePath field if non-nil, zero value otherwise.

### GetAssemblyFilePathOk

`func (o *IPlugin) GetAssemblyFilePathOk() (*string, bool)`

GetAssemblyFilePathOk returns a tuple with the AssemblyFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyFilePath

`func (o *IPlugin) SetAssemblyFilePath(v string)`

SetAssemblyFilePath sets AssemblyFilePath field to given value.

### HasAssemblyFilePath

`func (o *IPlugin) HasAssemblyFilePath() bool`

HasAssemblyFilePath returns a boolean if a field has been set.

### SetAssemblyFilePathNil

`func (o *IPlugin) SetAssemblyFilePathNil(b bool)`

 SetAssemblyFilePathNil sets the value for AssemblyFilePath to be an explicit nil

### UnsetAssemblyFilePath
`func (o *IPlugin) UnsetAssemblyFilePath()`

UnsetAssemblyFilePath ensures that no value is present for AssemblyFilePath, not even an explicit nil
### GetCanUninstall

`func (o *IPlugin) GetCanUninstall() bool`

GetCanUninstall returns the CanUninstall field if non-nil, zero value otherwise.

### GetCanUninstallOk

`func (o *IPlugin) GetCanUninstallOk() (*bool, bool)`

GetCanUninstallOk returns a tuple with the CanUninstall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUninstall

`func (o *IPlugin) SetCanUninstall(v bool)`

SetCanUninstall sets CanUninstall field to given value.

### HasCanUninstall

`func (o *IPlugin) HasCanUninstall() bool`

HasCanUninstall returns a boolean if a field has been set.

### GetDataFolderPath

`func (o *IPlugin) GetDataFolderPath() string`

GetDataFolderPath returns the DataFolderPath field if non-nil, zero value otherwise.

### GetDataFolderPathOk

`func (o *IPlugin) GetDataFolderPathOk() (*string, bool)`

GetDataFolderPathOk returns a tuple with the DataFolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataFolderPath

`func (o *IPlugin) SetDataFolderPath(v string)`

SetDataFolderPath sets DataFolderPath field to given value.

### HasDataFolderPath

`func (o *IPlugin) HasDataFolderPath() bool`

HasDataFolderPath returns a boolean if a field has been set.

### SetDataFolderPathNil

`func (o *IPlugin) SetDataFolderPathNil(b bool)`

 SetDataFolderPathNil sets the value for DataFolderPath to be an explicit nil

### UnsetDataFolderPath
`func (o *IPlugin) UnsetDataFolderPath()`

UnsetDataFolderPath ensures that no value is present for DataFolderPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


