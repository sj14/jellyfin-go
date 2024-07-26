/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PackageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageInfo{}

// PackageInfo Class PackageInfo.
type PackageInfo struct {
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

// NewPackageInfo instantiates a new PackageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageInfo() *PackageInfo {
	this := PackageInfo{}
	return &this
}

// NewPackageInfoWithDefaults instantiates a new PackageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageInfoWithDefaults() *PackageInfo {
	this := PackageInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PackageInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PackageInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PackageInfo) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PackageInfo) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PackageInfo) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PackageInfo) SetDescription(v string) {
	o.Description = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise.
func (o *PackageInfo) GetOverview() string {
	if o == nil || IsNil(o.Overview) {
		var ret string
		return ret
	}
	return *o.Overview
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageInfo) GetOverviewOk() (*string, bool) {
	if o == nil || IsNil(o.Overview) {
		return nil, false
	}
	return o.Overview, true
}

// HasOverview returns a boolean if a field has been set.
func (o *PackageInfo) HasOverview() bool {
	if o != nil && !IsNil(o.Overview) {
		return true
	}

	return false
}

// SetOverview gets a reference to the given string and assigns it to the Overview field.
func (o *PackageInfo) SetOverview(v string) {
	o.Overview = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *PackageInfo) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageInfo) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *PackageInfo) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *PackageInfo) SetOwner(v string) {
	o.Owner = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *PackageInfo) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageInfo) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *PackageInfo) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *PackageInfo) SetCategory(v string) {
	o.Category = &v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *PackageInfo) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageInfo) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *PackageInfo) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *PackageInfo) SetGuid(v string) {
	o.Guid = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *PackageInfo) GetVersions() []VersionInfo {
	if o == nil || IsNil(o.Versions) {
		var ret []VersionInfo
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageInfo) GetVersionsOk() ([]VersionInfo, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *PackageInfo) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []VersionInfo and assigns it to the Versions field.
func (o *PackageInfo) SetVersions(v []VersionInfo) {
	o.Versions = v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PackageInfo) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PackageInfo) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// HasImageUrl returns a boolean if a field has been set.
func (o *PackageInfo) HasImageUrl() bool {
	if o != nil && o.ImageUrl.IsSet() {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given NullableString and assigns it to the ImageUrl field.
func (o *PackageInfo) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}
// SetImageUrlNil sets the value for ImageUrl to be an explicit nil
func (o *PackageInfo) SetImageUrlNil() {
	o.ImageUrl.Set(nil)
}

// UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
func (o *PackageInfo) UnsetImageUrl() {
	o.ImageUrl.Unset()
}

func (o PackageInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageInfo) ToMap() (map[string]interface{}, error) {
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

type NullablePackageInfo struct {
	value *PackageInfo
	isSet bool
}

func (v NullablePackageInfo) Get() *PackageInfo {
	return v.value
}

func (v *NullablePackageInfo) Set(val *PackageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageInfo(val *PackageInfo) *NullablePackageInfo {
	return &NullablePackageInfo{value: val, isSet: true}
}

func (v NullablePackageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


