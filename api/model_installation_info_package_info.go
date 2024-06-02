/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the InstallationInfoPackageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstallationInfoPackageInfo{}

// InstallationInfoPackageInfo Gets or sets package information for the installation.
type InstallationInfoPackageInfo struct {
	// Gets or sets the name.
	Name *string `json:"name,omitempty"`
	// Gets or sets a long description of the plugin containing features or helpful explanations.
	Description *string `json:"description,omitempty"`
	// Gets or sets a short overview of what the plugin does.
	Overview *string `json:"overview,omitempty"`
	// Gets or sets the owner.
	Owner *string `json:"owner,omitempty"`
	// Gets or sets the category.
	Category *string `json:"category,omitempty"`
	// Gets or sets the guid of the assembly associated with this plugin.  This is used to identify the proper item for automatic updates.
	Guid *string `json:"guid,omitempty"`
	// Gets or sets the versions.
	Versions []VersionInfo `json:"versions,omitempty"`
	// Gets or sets the image url for the package.
	ImageUrl NullableString `json:"imageUrl,omitempty"`
}

// NewInstallationInfoPackageInfo instantiates a new InstallationInfoPackageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallationInfoPackageInfo() *InstallationInfoPackageInfo {
	this := InstallationInfoPackageInfo{}
	return &this
}

// NewInstallationInfoPackageInfoWithDefaults instantiates a new InstallationInfoPackageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallationInfoPackageInfoWithDefaults() *InstallationInfoPackageInfo {
	this := InstallationInfoPackageInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstallationInfoPackageInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationInfoPackageInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstallationInfoPackageInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstallationInfoPackageInfo) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InstallationInfoPackageInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationInfoPackageInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InstallationInfoPackageInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InstallationInfoPackageInfo) SetDescription(v string) {
	o.Description = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *InstallationInfoPackageInfo) GetOverview() string {
	if o == nil || IsNil(o.Overview) {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationInfoPackageInfo) GetOverviewOk() (*string, bool) {
	if o == nil || IsNil(o.Overview) {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *InstallationInfoPackageInfo) HasOverview() bool {
	if o != nil && !IsNil(o.Overview) {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *InstallationInfoPackageInfo) SetOverview(v string) {
	o.Overview = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *InstallationInfoPackageInfo) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationInfoPackageInfo) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *InstallationInfoPackageInfo) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *InstallationInfoPackageInfo) SetOwner(v string) {
	o.Owner = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InstallationInfoPackageInfo) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationInfoPackageInfo) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InstallationInfoPackageInfo) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InstallationInfoPackageInfo) SetCategory(v string) {
	o.Category = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *InstallationInfoPackageInfo) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationInfoPackageInfo) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *InstallationInfoPackageInfo) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *InstallationInfoPackageInfo) SetGuid(v string) {
	o.Guid = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *InstallationInfoPackageInfo) GetVersions() []VersionInfo {
	if o == nil || IsNil(o.Versions) {
		var ret []VersionInfo
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstallationInfoPackageInfo) GetVersionsOk() ([]VersionInfo, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *InstallationInfoPackageInfo) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []VersionInfo and assigns it to the Versions field.
func (o *InstallationInfoPackageInfo) SetVersions(v []VersionInfo) {
	o.Versions = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstallationInfoPackageInfo) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstallationInfoPackageInfo) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *InstallationInfoPackageInfo) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *InstallationInfoPackageInfo) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *InstallationInfoPackageInfo) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *InstallationInfoPackageInfo) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

func (o InstallationInfoPackageInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallationInfoPackageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Overview) {
		toSerialize["overview"] = o.Overview
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if o.ImageUrl.IsSet() {
		toSerialize["imageUrl"] = o.ImageUrl.Get()
	}
	return toSerialize, nil
}

type NullableInstallationInfoPackageInfo struct {
	value *InstallationInfoPackageInfo
	isSet bool
}

func (v NullableInstallationInfoPackageInfo) Get() *InstallationInfoPackageInfo {
	return v.value
}

func (v *NullableInstallationInfoPackageInfo) Set(val *InstallationInfoPackageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallationInfoPackageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallationInfoPackageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallationInfoPackageInfo(val *InstallationInfoPackageInfo) *NullableInstallationInfoPackageInfo {
	return &NullableInstallationInfoPackageInfo{value: val, isSet: true}
}

func (v NullableInstallationInfoPackageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallationInfoPackageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


