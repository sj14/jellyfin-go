/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the VersionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionInfo{}

// VersionInfo Defines the MediaBrowser.Model.Updates.VersionInfo class.
type VersionInfo struct {
	// Gets or sets the version.
	Version *string `json:"version,omitempty"`
	// Gets the version as a System.Version.
	VersionNumber *string `json:"VersionNumber,omitempty"`
	// Gets or sets the changelog for this version.
	Changelog NullableString `json:"changelog,omitempty"`
	// Gets or sets the ABI that this version was built against.
	TargetAbi NullableString `json:"targetAbi,omitempty"`
	// Gets or sets the source URL.
	SourceUrl NullableString `json:"sourceUrl,omitempty"`
	// Gets or sets a checksum for the binary.
	Checksum NullableString `json:"checksum,omitempty"`
	// Gets or sets a timestamp of when the binary was built.
	Timestamp NullableString `json:"timestamp,omitempty"`
	// Gets or sets the repository name.
	RepositoryName *string `json:"repositoryName,omitempty"`
	// Gets or sets the repository url.
	RepositoryUrl *string `json:"repositoryUrl,omitempty"`
}

// NewVersionInfo instantiates a new VersionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionInfo() *VersionInfo {
	this := VersionInfo{}
	return &this
}

// NewVersionInfoWithDefaults instantiates a new VersionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionInfoWithDefaults() *VersionInfo {
	this := VersionInfo{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VersionInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VersionInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VersionInfo) SetVersion(v string) {
	o.Version = &v
}

// GetVersionNumber returns the VersionNumber field value if set, zero value otherwise.
func (o *VersionInfo) GetVersionNumber() string {
	if o == nil || IsNil(o.VersionNumber) {
		var ret string
		return ret
	}
	return *o.VersionNumber
}

// GetVersionNumberOk returns a tuple with the VersionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetVersionNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VersionNumber) {
		return nil, false
	}
	return o.VersionNumber, true
}

// HasVersionNumber returns a boolean if a field has been set.
func (o *VersionInfo) HasVersionNumber() bool {
	if o != nil && !IsNil(o.VersionNumber) {
		return true
	}

	return false
}

// SetVersionNumber gets a reference to the given string and assigns it to the VersionNumber field.
func (o *VersionInfo) SetVersionNumber(v string) {
	o.VersionNumber = &v
}

// GetChangelog returns the Changelog field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionInfo) GetChangelog() string {
	if o == nil || IsNil(o.Changelog.Get()) {
		var ret string
		return ret
	}
	return *o.Changelog.Get()
}

// GetChangelogOk returns a tuple with the Changelog field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionInfo) GetChangelogOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changelog.Get(), o.Changelog.IsSet()
}

// HasChangelog returns a boolean if a field has been set.
func (o *VersionInfo) HasChangelog() bool {
	if o != nil && o.Changelog.IsSet() {
		return true
	}

	return false
}

// SetChangelog gets a reference to the given NullableString and assigns it to the Changelog field.
func (o *VersionInfo) SetChangelog(v string) {
	o.Changelog.Set(&v)
}
// SetChangelogNil sets the value for Changelog to be an explicit nil
func (o *VersionInfo) SetChangelogNil() {
	o.Changelog.Set(nil)
}

// UnsetChangelog ensures that no value is present for Changelog, not even an explicit nil
func (o *VersionInfo) UnsetChangelog() {
	o.Changelog.Unset()
}

// GetTargetAbi returns the TargetAbi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionInfo) GetTargetAbi() string {
	if o == nil || IsNil(o.TargetAbi.Get()) {
		var ret string
		return ret
	}
	return *o.TargetAbi.Get()
}

// GetTargetAbiOk returns a tuple with the TargetAbi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionInfo) GetTargetAbiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetAbi.Get(), o.TargetAbi.IsSet()
}

// HasTargetAbi returns a boolean if a field has been set.
func (o *VersionInfo) HasTargetAbi() bool {
	if o != nil && o.TargetAbi.IsSet() {
		return true
	}

	return false
}

// SetTargetAbi gets a reference to the given NullableString and assigns it to the TargetAbi field.
func (o *VersionInfo) SetTargetAbi(v string) {
	o.TargetAbi.Set(&v)
}
// SetTargetAbiNil sets the value for TargetAbi to be an explicit nil
func (o *VersionInfo) SetTargetAbiNil() {
	o.TargetAbi.Set(nil)
}

// UnsetTargetAbi ensures that no value is present for TargetAbi, not even an explicit nil
func (o *VersionInfo) UnsetTargetAbi() {
	o.TargetAbi.Unset()
}

// GetSourceUrl returns the SourceUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionInfo) GetSourceUrl() string {
	if o == nil || IsNil(o.SourceUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SourceUrl.Get()
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionInfo) GetSourceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceUrl.Get(), o.SourceUrl.IsSet()
}

// HasSourceUrl returns a boolean if a field has been set.
func (o *VersionInfo) HasSourceUrl() bool {
	if o != nil && o.SourceUrl.IsSet() {
		return true
	}

	return false
}

// SetSourceUrl gets a reference to the given NullableString and assigns it to the SourceUrl field.
func (o *VersionInfo) SetSourceUrl(v string) {
	o.SourceUrl.Set(&v)
}
// SetSourceUrlNil sets the value for SourceUrl to be an explicit nil
func (o *VersionInfo) SetSourceUrlNil() {
	o.SourceUrl.Set(nil)
}

// UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
func (o *VersionInfo) UnsetSourceUrl() {
	o.SourceUrl.Unset()
}

// GetChecksum returns the Checksum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionInfo) GetChecksum() string {
	if o == nil || IsNil(o.Checksum.Get()) {
		var ret string
		return ret
	}
	return *o.Checksum.Get()
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionInfo) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Checksum.Get(), o.Checksum.IsSet()
}

// HasChecksum returns a boolean if a field has been set.
func (o *VersionInfo) HasChecksum() bool {
	if o != nil && o.Checksum.IsSet() {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given NullableString and assigns it to the Checksum field.
func (o *VersionInfo) SetChecksum(v string) {
	o.Checksum.Set(&v)
}
// SetChecksumNil sets the value for Checksum to be an explicit nil
func (o *VersionInfo) SetChecksumNil() {
	o.Checksum.Set(nil)
}

// UnsetChecksum ensures that no value is present for Checksum, not even an explicit nil
func (o *VersionInfo) UnsetChecksum() {
	o.Checksum.Unset()
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VersionInfo) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp.Get()) {
		var ret string
		return ret
	}
	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VersionInfo) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// HasTimestamp returns a boolean if a field has been set.
func (o *VersionInfo) HasTimestamp() bool {
	if o != nil && o.Timestamp.IsSet() {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given NullableString and assigns it to the Timestamp field.
func (o *VersionInfo) SetTimestamp(v string) {
	o.Timestamp.Set(&v)
}
// SetTimestampNil sets the value for Timestamp to be an explicit nil
func (o *VersionInfo) SetTimestampNil() {
	o.Timestamp.Set(nil)
}

// UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
func (o *VersionInfo) UnsetTimestamp() {
	o.Timestamp.Unset()
}

// GetRepositoryName returns the RepositoryName field value if set, zero value otherwise.
func (o *VersionInfo) GetRepositoryName() string {
	if o == nil || IsNil(o.RepositoryName) {
		var ret string
		return ret
	}
	return *o.RepositoryName
}

// GetRepositoryNameOk returns a tuple with the RepositoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetRepositoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryName) {
		return nil, false
	}
	return o.RepositoryName, true
}

// HasRepositoryName returns a boolean if a field has been set.
func (o *VersionInfo) HasRepositoryName() bool {
	if o != nil && !IsNil(o.RepositoryName) {
		return true
	}

	return false
}

// SetRepositoryName gets a reference to the given string and assigns it to the RepositoryName field.
func (o *VersionInfo) SetRepositoryName(v string) {
	o.RepositoryName = &v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *VersionInfo) GetRepositoryUrl() string {
	if o == nil || IsNil(o.RepositoryUrl) {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryUrl) {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *VersionInfo) HasRepositoryUrl() bool {
	if o != nil && !IsNil(o.RepositoryUrl) {
		return true
	}

	return false
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *VersionInfo) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

func (o VersionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.VersionNumber) {
		toSerialize["VersionNumber"] = o.VersionNumber
	}
	if o.Changelog.IsSet() {
		toSerialize["changelog"] = o.Changelog.Get()
	}
	if o.TargetAbi.IsSet() {
		toSerialize["targetAbi"] = o.TargetAbi.Get()
	}
	if o.SourceUrl.IsSet() {
		toSerialize["sourceUrl"] = o.SourceUrl.Get()
	}
	if o.Checksum.IsSet() {
		toSerialize["checksum"] = o.Checksum.Get()
	}
	if o.Timestamp.IsSet() {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if !IsNil(o.RepositoryName) {
		toSerialize["repositoryName"] = o.RepositoryName
	}
	if !IsNil(o.RepositoryUrl) {
		toSerialize["repositoryUrl"] = o.RepositoryUrl
	}
	return toSerialize, nil
}

type NullableVersionInfo struct {
	value *VersionInfo
	isSet bool
}

func (v NullableVersionInfo) Get() *VersionInfo {
	return v.value
}

func (v *NullableVersionInfo) Set(val *VersionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionInfo(val *VersionInfo) *NullableVersionInfo {
	return &NullableVersionInfo{value: val, isSet: true}
}

func (v NullableVersionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


