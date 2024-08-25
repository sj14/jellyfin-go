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

// checks if the ConfigurationPageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigurationPageInfo{}

// ConfigurationPageInfo The configuration page info.
type ConfigurationPageInfo struct {
	// Gets or sets the name.
	Name *string `json:"Name,omitempty"`
	// Gets or sets a value indicating whether the configurations page is enabled in the main menu.
	EnableInMainMenu *bool `json:"EnableInMainMenu,omitempty"`
	// Gets or sets the menu section.
	MenuSection NullableString `json:"MenuSection,omitempty"`
	// Gets or sets the menu icon.
	MenuIcon NullableString `json:"MenuIcon,omitempty"`
	// Gets or sets the display name.
	DisplayName NullableString `json:"DisplayName,omitempty"`
	// Gets or sets the plugin id.
	PluginId NullableString `json:"PluginId,omitempty"`
}

// NewConfigurationPageInfo instantiates a new ConfigurationPageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigurationPageInfo() *ConfigurationPageInfo {
	this := ConfigurationPageInfo{}
	return &this
}

// NewConfigurationPageInfoWithDefaults instantiates a new ConfigurationPageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigurationPageInfoWithDefaults() *ConfigurationPageInfo {
	this := ConfigurationPageInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigurationPageInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPageInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigurationPageInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConfigurationPageInfo) SetName(v string) {
	o.Name = &v
}

// GetEnableInMainMenu returns the EnableInMainMenu field value if set, zero value otherwise.
func (o *ConfigurationPageInfo) GetEnableInMainMenu() bool {
	if o == nil || IsNil(o.EnableInMainMenu) {
		var ret bool
		return ret
	}
	return *o.EnableInMainMenu
}

// GetEnableInMainMenuOk returns a tuple with the EnableInMainMenu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigurationPageInfo) GetEnableInMainMenuOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableInMainMenu) {
		return nil, false
	}
	return o.EnableInMainMenu, true
}

// HasEnableInMainMenu returns a boolean if a field has been set.
func (o *ConfigurationPageInfo) HasEnableInMainMenu() bool {
	if o != nil && !IsNil(o.EnableInMainMenu) {
		return true
	}

	return false
}

// SetEnableInMainMenu gets a reference to the given bool and assigns it to the EnableInMainMenu field.
func (o *ConfigurationPageInfo) SetEnableInMainMenu(v bool) {
	o.EnableInMainMenu = &v
}

// GetMenuSection returns the MenuSection field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationPageInfo) GetMenuSection() string {
	if o == nil || IsNil(o.MenuSection.Get()) {
		var ret string
		return ret
	}
	return *o.MenuSection.Get()
}

// GetMenuSectionOk returns a tuple with the MenuSection field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationPageInfo) GetMenuSectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MenuSection.Get(), o.MenuSection.IsSet()
}

// HasMenuSection returns a boolean if a field has been set.
func (o *ConfigurationPageInfo) HasMenuSection() bool {
	if o != nil && o.MenuSection.IsSet() {
		return true
	}

	return false
}

// SetMenuSection gets a reference to the given NullableString and assigns it to the MenuSection field.
func (o *ConfigurationPageInfo) SetMenuSection(v string) {
	o.MenuSection.Set(&v)
}
// SetMenuSectionNil sets the value for MenuSection to be an explicit nil
func (o *ConfigurationPageInfo) SetMenuSectionNil() {
	o.MenuSection.Set(nil)
}

// UnsetMenuSection ensures that no value is present for MenuSection, not even an explicit nil
func (o *ConfigurationPageInfo) UnsetMenuSection() {
	o.MenuSection.Unset()
}

// GetMenuIcon returns the MenuIcon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationPageInfo) GetMenuIcon() string {
	if o == nil || IsNil(o.MenuIcon.Get()) {
		var ret string
		return ret
	}
	return *o.MenuIcon.Get()
}

// GetMenuIconOk returns a tuple with the MenuIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationPageInfo) GetMenuIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MenuIcon.Get(), o.MenuIcon.IsSet()
}

// HasMenuIcon returns a boolean if a field has been set.
func (o *ConfigurationPageInfo) HasMenuIcon() bool {
	if o != nil && o.MenuIcon.IsSet() {
		return true
	}

	return false
}

// SetMenuIcon gets a reference to the given NullableString and assigns it to the MenuIcon field.
func (o *ConfigurationPageInfo) SetMenuIcon(v string) {
	o.MenuIcon.Set(&v)
}
// SetMenuIconNil sets the value for MenuIcon to be an explicit nil
func (o *ConfigurationPageInfo) SetMenuIconNil() {
	o.MenuIcon.Set(nil)
}

// UnsetMenuIcon ensures that no value is present for MenuIcon, not even an explicit nil
func (o *ConfigurationPageInfo) UnsetMenuIcon() {
	o.MenuIcon.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationPageInfo) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationPageInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ConfigurationPageInfo) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *ConfigurationPageInfo) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *ConfigurationPageInfo) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *ConfigurationPageInfo) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetPluginId returns the PluginId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigurationPageInfo) GetPluginId() string {
	if o == nil || IsNil(o.PluginId.Get()) {
		var ret string
		return ret
	}
	return *o.PluginId.Get()
}

// GetPluginIdOk returns a tuple with the PluginId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigurationPageInfo) GetPluginIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginId.Get(), o.PluginId.IsSet()
}

// HasPluginId returns a boolean if a field has been set.
func (o *ConfigurationPageInfo) HasPluginId() bool {
	if o != nil && o.PluginId.IsSet() {
		return true
	}

	return false
}

// SetPluginId gets a reference to the given NullableString and assigns it to the PluginId field.
func (o *ConfigurationPageInfo) SetPluginId(v string) {
	o.PluginId.Set(&v)
}
// SetPluginIdNil sets the value for PluginId to be an explicit nil
func (o *ConfigurationPageInfo) SetPluginIdNil() {
	o.PluginId.Set(nil)
}

// UnsetPluginId ensures that no value is present for PluginId, not even an explicit nil
func (o *ConfigurationPageInfo) UnsetPluginId() {
	o.PluginId.Unset()
}

func (o ConfigurationPageInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigurationPageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.EnableInMainMenu) {
		toSerialize["EnableInMainMenu"] = o.EnableInMainMenu
	}
	if o.MenuSection.IsSet() {
		toSerialize["MenuSection"] = o.MenuSection.Get()
	}
	if o.MenuIcon.IsSet() {
		toSerialize["MenuIcon"] = o.MenuIcon.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["DisplayName"] = o.DisplayName.Get()
	}
	if o.PluginId.IsSet() {
		toSerialize["PluginId"] = o.PluginId.Get()
	}
	return toSerialize, nil
}

type NullableConfigurationPageInfo struct {
	value *ConfigurationPageInfo
	isSet bool
}

func (v NullableConfigurationPageInfo) Get() *ConfigurationPageInfo {
	return v.value
}

func (v *NullableConfigurationPageInfo) Set(val *ConfigurationPageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigurationPageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigurationPageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigurationPageInfo(val *ConfigurationPageInfo) *NullableConfigurationPageInfo {
	return &NullableConfigurationPageInfo{value: val, isSet: true}
}

func (v NullableConfigurationPageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigurationPageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


