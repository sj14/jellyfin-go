/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PluginInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginInfo{}

// PluginInfo This is a serializable stub class that is used by the api to provide information about installed plugins.
type PluginInfo struct {
	// Gets or sets the name.
	Name *string `json:"Name,omitempty"`
	// Gets or sets the version.
	Version *string `json:"Version,omitempty"`
	// Gets or sets the name of the configuration file.
	ConfigurationFileName NullableString `json:"ConfigurationFileName,omitempty"`
	// Gets or sets the description.
	Description *string `json:"Description,omitempty"`
	// Gets or sets the unique id.
	Id *string `json:"Id,omitempty"`
	// Gets or sets a value indicating whether the plugin can be uninstalled.
	CanUninstall *bool `json:"CanUninstall,omitempty"`
	// Gets or sets a value indicating whether this plugin has a valid image.
	HasImage *bool `json:"HasImage,omitempty"`
	// Gets or sets a value indicating the status of the plugin.
	Status *PluginStatus `json:"Status,omitempty"`
}

// NewPluginInfo instantiates a new PluginInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginInfo() *PluginInfo {
	this := PluginInfo{}
	return &this
}

// NewPluginInfoWithDefaults instantiates a new PluginInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginInfoWithDefaults() *PluginInfo {
	this := PluginInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PluginInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PluginInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PluginInfo) SetName(v string) {
	o.Name = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PluginInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PluginInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PluginInfo) SetVersion(v string) {
	o.Version = &v
}

// GetConfigurationFileName returns the ConfigurationFileName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PluginInfo) GetConfigurationFileName() string {
	if o == nil || IsNil(o.ConfigurationFileName.Get()) {
		var ret string
		return ret
	}
	return *o.ConfigurationFileName.Get()
}

// GetConfigurationFileNameOk returns a tuple with the ConfigurationFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PluginInfo) GetConfigurationFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigurationFileName.Get(), o.ConfigurationFileName.IsSet()
}

// HasConfigurationFileName returns a boolean if a field has been set.
func (o *PluginInfo) HasConfigurationFileName() bool {
	if o != nil && o.ConfigurationFileName.IsSet() {
		return true
	}

	return false
}

// SetConfigurationFileName gets a reference to the given NullableString and assigns it to the ConfigurationFileName field.
func (o *PluginInfo) SetConfigurationFileName(v string) {
	o.ConfigurationFileName.Set(&v)
}
// SetConfigurationFileNameNil sets the value for ConfigurationFileName to be an explicit nil
func (o *PluginInfo) SetConfigurationFileNameNil() {
	o.ConfigurationFileName.Set(nil)
}

// UnsetConfigurationFileName ensures that no value is present for ConfigurationFileName, not even an explicit nil
func (o *PluginInfo) UnsetConfigurationFileName() {
	o.ConfigurationFileName.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PluginInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PluginInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PluginInfo) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PluginInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PluginInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PluginInfo) SetId(v string) {
	o.Id = &v
}

// GetCanUninstall returns the CanUninstall field value if set, zero value otherwise.
func (o *PluginInfo) GetCanUninstall() bool {
	if o == nil || IsNil(o.CanUninstall) {
		var ret bool
		return ret
	}
	return *o.CanUninstall
}

// GetCanUninstallOk returns a tuple with the CanUninstall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInfo) GetCanUninstallOk() (*bool, bool) {
	if o == nil || IsNil(o.CanUninstall) {
		return nil, false
	}
	return o.CanUninstall, true
}

// HasCanUninstall returns a boolean if a field has been set.
func (o *PluginInfo) HasCanUninstall() bool {
	if o != nil && !IsNil(o.CanUninstall) {
		return true
	}

	return false
}

// SetCanUninstall gets a reference to the given bool and assigns it to the CanUninstall field.
func (o *PluginInfo) SetCanUninstall(v bool) {
	o.CanUninstall = &v
}

// GetHasImage returns the HasImage field value if set, zero value otherwise.
func (o *PluginInfo) GetHasImage() bool {
	if o == nil || IsNil(o.HasImage) {
		var ret bool
		return ret
	}
	return *o.HasImage
}

// GetHasImageOk returns a tuple with the HasImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInfo) GetHasImageOk() (*bool, bool) {
	if o == nil || IsNil(o.HasImage) {
		return nil, false
	}
	return o.HasImage, true
}

// HasHasImage returns a boolean if a field has been set.
func (o *PluginInfo) HasHasImage() bool {
	if o != nil && !IsNil(o.HasImage) {
		return true
	}

	return false
}

// SetHasImage gets a reference to the given bool and assigns it to the HasImage field.
func (o *PluginInfo) SetHasImage(v bool) {
	o.HasImage = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PluginInfo) GetStatus() PluginStatus {
	if o == nil || IsNil(o.Status) {
		var ret PluginStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInfo) GetStatusOk() (*PluginStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PluginInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given PluginStatus and assigns it to the Status field.
func (o *PluginInfo) SetStatus(v PluginStatus) {
	o.Status = &v
}

func (o PluginInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Version) {
		toSerialize["Version"] = o.Version
	}
	if o.ConfigurationFileName.IsSet() {
		toSerialize["ConfigurationFileName"] = o.ConfigurationFileName.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.CanUninstall) {
		toSerialize["CanUninstall"] = o.CanUninstall
	}
	if !IsNil(o.HasImage) {
		toSerialize["HasImage"] = o.HasImage
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	return toSerialize, nil
}

type NullablePluginInfo struct {
	value *PluginInfo
	isSet bool
}

func (v NullablePluginInfo) Get() *PluginInfo {
	return v.value
}

func (v *NullablePluginInfo) Set(val *PluginInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginInfo(val *PluginInfo) *NullablePluginInfo {
	return &NullablePluginInfo{value: val, isSet: true}
}

func (v NullablePluginInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


