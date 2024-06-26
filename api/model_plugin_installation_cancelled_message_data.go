/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PluginInstallationCancelledMessageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginInstallationCancelledMessageData{}

// PluginInstallationCancelledMessageData Gets or sets the data.
type PluginInstallationCancelledMessageData struct {
	// Gets or sets the Id.
	Guid *string `json:"Guid,omitempty"`
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the version.
	Version NullableString `json:"Version,omitempty"`
	// Gets or sets the changelog for this version.
	Changelog NullableString `json:"Changelog,omitempty"`
	// Gets or sets the source URL.
	SourceUrl NullableString `json:"SourceUrl,omitempty"`
	// Gets or sets a checksum for the binary.
	Checksum NullableString `json:"Checksum,omitempty"`
	PackageInfo NullableInstallationInfoPackageInfo `json:"PackageInfo,omitempty"`
}

// NewPluginInstallationCancelledMessageData instantiates a new PluginInstallationCancelledMessageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginInstallationCancelledMessageData() *PluginInstallationCancelledMessageData {
	this := PluginInstallationCancelledMessageData{}
	return &this
}

// NewPluginInstallationCancelledMessageDataWithDefaults instantiates a new PluginInstallationCancelledMessageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginInstallationCancelledMessageDataWithDefaults() *PluginInstallationCancelledMessageData {
	this := PluginInstallationCancelledMessageData{}
	return &this
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *PluginInstallationCancelledMessageData) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginInstallationCancelledMessageData) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *PluginInstallationCancelledMessageData) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *PluginInstallationCancelledMessageData) SetGuid(v string) {
	o.Guid = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PluginInstallationCancelledMessageData) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PluginInstallationCancelledMessageData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PluginInstallationCancelledMessageData) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PluginInstallationCancelledMessageData) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PluginInstallationCancelledMessageData) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PluginInstallationCancelledMessageData) UnsetName() {
	o.Name.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PluginInstallationCancelledMessageData) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PluginInstallationCancelledMessageData) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *PluginInstallationCancelledMessageData) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *PluginInstallationCancelledMessageData) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *PluginInstallationCancelledMessageData) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *PluginInstallationCancelledMessageData) UnsetVersion() {
	o.Version.Unset()
}

// GetChangelog returns the Changelog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PluginInstallationCancelledMessageData) GetChangelog() string {
	if o == nil || IsNil(o.Changelog.Get()) {
		var ret string
		return ret
	}
	return *o.Changelog.Get()
}

// GetChangelogOk returns a tuple with the Changelog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PluginInstallationCancelledMessageData) GetChangelogOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changelog.Get(), o.Changelog.IsSet()
}

// HasChangelog returns a boolean if a field has been set.
func (o *PluginInstallationCancelledMessageData) HasChangelog() bool {
	if o != nil && o.Changelog.IsSet() {
		return true
	}

	return false
}

// SetChangelog gets a reference to the given NullableString and assigns it to the Changelog field.
func (o *PluginInstallationCancelledMessageData) SetChangelog(v string) {
	o.Changelog.Set(&v)
}
// SetChangelogNil sets the value for Changelog to be an explicit nil
func (o *PluginInstallationCancelledMessageData) SetChangelogNil() {
	o.Changelog.Set(nil)
}

// UnsetChangelog ensures that no value is present for Changelog, not even an explicit nil
func (o *PluginInstallationCancelledMessageData) UnsetChangelog() {
	o.Changelog.Unset()
}

// GetSourceUrl returns the SourceUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PluginInstallationCancelledMessageData) GetSourceUrl() string {
	if o == nil || IsNil(o.SourceUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SourceUrl.Get()
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PluginInstallationCancelledMessageData) GetSourceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceUrl.Get(), o.SourceUrl.IsSet()
}

// HasSourceUrl returns a boolean if a field has been set.
func (o *PluginInstallationCancelledMessageData) HasSourceUrl() bool {
	if o != nil && o.SourceUrl.IsSet() {
		return true
	}

	return false
}

// SetSourceUrl gets a reference to the given NullableString and assigns it to the SourceUrl field.
func (o *PluginInstallationCancelledMessageData) SetSourceUrl(v string) {
	o.SourceUrl.Set(&v)
}
// SetSourceUrlNil sets the value for SourceUrl to be an explicit nil
func (o *PluginInstallationCancelledMessageData) SetSourceUrlNil() {
	o.SourceUrl.Set(nil)
}

// UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
func (o *PluginInstallationCancelledMessageData) UnsetSourceUrl() {
	o.SourceUrl.Unset()
}

// GetChecksum returns the Checksum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PluginInstallationCancelledMessageData) GetChecksum() string {
	if o == nil || IsNil(o.Checksum.Get()) {
		var ret string
		return ret
	}
	return *o.Checksum.Get()
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PluginInstallationCancelledMessageData) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Checksum.Get(), o.Checksum.IsSet()
}

// HasChecksum returns a boolean if a field has been set.
func (o *PluginInstallationCancelledMessageData) HasChecksum() bool {
	if o != nil && o.Checksum.IsSet() {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given NullableString and assigns it to the Checksum field.
func (o *PluginInstallationCancelledMessageData) SetChecksum(v string) {
	o.Checksum.Set(&v)
}
// SetChecksumNil sets the value for Checksum to be an explicit nil
func (o *PluginInstallationCancelledMessageData) SetChecksumNil() {
	o.Checksum.Set(nil)
}

// UnsetChecksum ensures that no value is present for Checksum, not even an explicit nil
func (o *PluginInstallationCancelledMessageData) UnsetChecksum() {
	o.Checksum.Unset()
}

// GetPackageInfo returns the PackageInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PluginInstallationCancelledMessageData) GetPackageInfo() InstallationInfoPackageInfo {
	if o == nil || IsNil(o.PackageInfo.Get()) {
		var ret InstallationInfoPackageInfo
		return ret
	}
	return *o.PackageInfo.Get()
}

// GetPackageInfoOk returns a tuple with the PackageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PluginInstallationCancelledMessageData) GetPackageInfoOk() (*InstallationInfoPackageInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageInfo.Get(), o.PackageInfo.IsSet()
}

// HasPackageInfo returns a boolean if a field has been set.
func (o *PluginInstallationCancelledMessageData) HasPackageInfo() bool {
	if o != nil && o.PackageInfo.IsSet() {
		return true
	}

	return false
}

// SetPackageInfo gets a reference to the given NullableInstallationInfoPackageInfo and assigns it to the PackageInfo field.
func (o *PluginInstallationCancelledMessageData) SetPackageInfo(v InstallationInfoPackageInfo) {
	o.PackageInfo.Set(&v)
}
// SetPackageInfoNil sets the value for PackageInfo to be an explicit nil
func (o *PluginInstallationCancelledMessageData) SetPackageInfoNil() {
	o.PackageInfo.Set(nil)
}

// UnsetPackageInfo ensures that no value is present for PackageInfo, not even an explicit nil
func (o *PluginInstallationCancelledMessageData) UnsetPackageInfo() {
	o.PackageInfo.Unset()
}

func (o PluginInstallationCancelledMessageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginInstallationCancelledMessageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Guid) {
		toSerialize["Guid"] = o.Guid
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Version.IsSet() {
		toSerialize["Version"] = o.Version.Get()
	}
	if o.Changelog.IsSet() {
		toSerialize["Changelog"] = o.Changelog.Get()
	}
	if o.SourceUrl.IsSet() {
		toSerialize["SourceUrl"] = o.SourceUrl.Get()
	}
	if o.Checksum.IsSet() {
		toSerialize["Checksum"] = o.Checksum.Get()
	}
	if o.PackageInfo.IsSet() {
		toSerialize["PackageInfo"] = o.PackageInfo.Get()
	}
	return toSerialize, nil
}

type NullablePluginInstallationCancelledMessageData struct {
	value *PluginInstallationCancelledMessageData
	isSet bool
}

func (v NullablePluginInstallationCancelledMessageData) Get() *PluginInstallationCancelledMessageData {
	return v.value
}

func (v *NullablePluginInstallationCancelledMessageData) Set(val *PluginInstallationCancelledMessageData) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginInstallationCancelledMessageData) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginInstallationCancelledMessageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginInstallationCancelledMessageData(val *PluginInstallationCancelledMessageData) *NullablePluginInstallationCancelledMessageData {
	return &NullablePluginInstallationCancelledMessageData{value: val, isSet: true}
}

func (v NullablePluginInstallationCancelledMessageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginInstallationCancelledMessageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


