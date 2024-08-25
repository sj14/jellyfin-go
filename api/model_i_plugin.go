/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the IPlugin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPlugin{}

// IPlugin Defines the MediaBrowser.Common.Plugins.IPlugin.
type IPlugin struct {
	// Gets the name of the plugin.
	Name NullableString `json:"Name,omitempty"`
	// Gets the Description.
	Description NullableString `json:"Description,omitempty"`
	// Gets the unique id.
	Id *string `json:"Id,omitempty"`
	// Gets the plugin version.
	Version NullableString `json:"Version,omitempty"`
	// Gets the path to the assembly file.
	AssemblyFilePath NullableString `json:"AssemblyFilePath,omitempty"`
	// Gets a value indicating whether the plugin can be uninstalled.
	CanUninstall *bool `json:"CanUninstall,omitempty"`
	// Gets the full path to the data folder, where the plugin can store any miscellaneous files needed.
	DataFolderPath NullableString `json:"DataFolderPath,omitempty"`
}

// NewIPlugin instantiates a new IPlugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPlugin() *IPlugin {
	this := IPlugin{}
	return &this
}

// NewIPluginWithDefaults instantiates a new IPlugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPluginWithDefaults() *IPlugin {
	this := IPlugin{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPlugin) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPlugin) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IPlugin) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IPlugin) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *IPlugin) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IPlugin) UnsetName() {
	o.Name.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPlugin) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPlugin) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IPlugin) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *IPlugin) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *IPlugin) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *IPlugin) UnsetDescription() {
	o.Description.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IPlugin) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPlugin) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IPlugin) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IPlugin) SetId(v string) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPlugin) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPlugin) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *IPlugin) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *IPlugin) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *IPlugin) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *IPlugin) UnsetVersion() {
	o.Version.Unset()
}

// GetAssemblyFilePath returns the AssemblyFilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPlugin) GetAssemblyFilePath() string {
	if o == nil || IsNil(o.AssemblyFilePath.Get()) {
		var ret string
		return ret
	}
	return *o.AssemblyFilePath.Get()
}

// GetAssemblyFilePathOk returns a tuple with the AssemblyFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPlugin) GetAssemblyFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssemblyFilePath.Get(), o.AssemblyFilePath.IsSet()
}

// HasAssemblyFilePath returns a boolean if a field has been set.
func (o *IPlugin) HasAssemblyFilePath() bool {
	if o != nil && o.AssemblyFilePath.IsSet() {
		return true
	}

	return false
}

// SetAssemblyFilePath gets a reference to the given NullableString and assigns it to the AssemblyFilePath field.
func (o *IPlugin) SetAssemblyFilePath(v string) {
	o.AssemblyFilePath.Set(&v)
}
// SetAssemblyFilePathNil sets the value for AssemblyFilePath to be an explicit nil
func (o *IPlugin) SetAssemblyFilePathNil() {
	o.AssemblyFilePath.Set(nil)
}

// UnsetAssemblyFilePath ensures that no value is present for AssemblyFilePath, not even an explicit nil
func (o *IPlugin) UnsetAssemblyFilePath() {
	o.AssemblyFilePath.Unset()
}

// GetCanUninstall returns the CanUninstall field value if set, zero value otherwise.
func (o *IPlugin) GetCanUninstall() bool {
	if o == nil || IsNil(o.CanUninstall) {
		var ret bool
		return ret
	}
	return *o.CanUninstall
}

// GetCanUninstallOk returns a tuple with the CanUninstall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPlugin) GetCanUninstallOk() (*bool, bool) {
	if o == nil || IsNil(o.CanUninstall) {
		return nil, false
	}
	return o.CanUninstall, true
}

// HasCanUninstall returns a boolean if a field has been set.
func (o *IPlugin) HasCanUninstall() bool {
	if o != nil && !IsNil(o.CanUninstall) {
		return true
	}

	return false
}

// SetCanUninstall gets a reference to the given bool and assigns it to the CanUninstall field.
func (o *IPlugin) SetCanUninstall(v bool) {
	o.CanUninstall = &v
}

// GetDataFolderPath returns the DataFolderPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IPlugin) GetDataFolderPath() string {
	if o == nil || IsNil(o.DataFolderPath.Get()) {
		var ret string
		return ret
	}
	return *o.DataFolderPath.Get()
}

// GetDataFolderPathOk returns a tuple with the DataFolderPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IPlugin) GetDataFolderPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataFolderPath.Get(), o.DataFolderPath.IsSet()
}

// HasDataFolderPath returns a boolean if a field has been set.
func (o *IPlugin) HasDataFolderPath() bool {
	if o != nil && o.DataFolderPath.IsSet() {
		return true
	}

	return false
}

// SetDataFolderPath gets a reference to the given NullableString and assigns it to the DataFolderPath field.
func (o *IPlugin) SetDataFolderPath(v string) {
	o.DataFolderPath.Set(&v)
}
// SetDataFolderPathNil sets the value for DataFolderPath to be an explicit nil
func (o *IPlugin) SetDataFolderPathNil() {
	o.DataFolderPath.Set(nil)
}

// UnsetDataFolderPath ensures that no value is present for DataFolderPath, not even an explicit nil
func (o *IPlugin) UnsetDataFolderPath() {
	o.DataFolderPath.Unset()
}

func (o IPlugin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPlugin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.Version.IsSet() {
		toSerialize["Version"] = o.Version.Get()
	}
	if o.AssemblyFilePath.IsSet() {
		toSerialize["AssemblyFilePath"] = o.AssemblyFilePath.Get()
	}
	if !IsNil(o.CanUninstall) {
		toSerialize["CanUninstall"] = o.CanUninstall
	}
	if o.DataFolderPath.IsSet() {
		toSerialize["DataFolderPath"] = o.DataFolderPath.Get()
	}
	return toSerialize, nil
}

type NullableIPlugin struct {
	value *IPlugin
	isSet bool
}

func (v NullableIPlugin) Get() *IPlugin {
	return v.value
}

func (v *NullableIPlugin) Set(val *IPlugin) {
	v.value = val
	v.isSet = true
}

func (v NullableIPlugin) IsSet() bool {
	return v.isSet
}

func (v *NullableIPlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPlugin(val *IPlugin) *NullableIPlugin {
	return &NullableIPlugin{value: val, isSet: true}
}

func (v NullableIPlugin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


