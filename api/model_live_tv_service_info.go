/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the LiveTvServiceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveTvServiceInfo{}

// LiveTvServiceInfo Class ServiceInfo.
type LiveTvServiceInfo struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the home page URL.
	HomePageUrl NullableString `json:"HomePageUrl,omitempty"`
	// Gets or sets the status.
	Status *LiveTvServiceStatus `json:"Status,omitempty"`
	// Gets or sets the status message.
	StatusMessage NullableString `json:"StatusMessage,omitempty"`
	// Gets or sets the version.
	Version NullableString `json:"Version,omitempty"`
	// Gets or sets a value indicating whether this instance has update available.
	HasUpdateAvailable *bool `json:"HasUpdateAvailable,omitempty"`
	// Gets or sets a value indicating whether this instance is visible.
	IsVisible *bool `json:"IsVisible,omitempty"`
	Tuners []string `json:"Tuners,omitempty"`
}

// NewLiveTvServiceInfo instantiates a new LiveTvServiceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveTvServiceInfo() *LiveTvServiceInfo {
	this := LiveTvServiceInfo{}
	return &this
}

// NewLiveTvServiceInfoWithDefaults instantiates a new LiveTvServiceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveTvServiceInfoWithDefaults() *LiveTvServiceInfo {
	this := LiveTvServiceInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvServiceInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvServiceInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *LiveTvServiceInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *LiveTvServiceInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *LiveTvServiceInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *LiveTvServiceInfo) UnsetName() {
	o.Name.Unset()
}

// GetHomePageUrl returns the HomePageUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvServiceInfo) GetHomePageUrl() string {
	if o == nil || IsNil(o.HomePageUrl.Get()) {
		var ret string
		return ret
	}
	return *o.HomePageUrl.Get()
}

// GetHomePageUrlOk returns a tuple with the HomePageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvServiceInfo) GetHomePageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HomePageUrl.Get(), o.HomePageUrl.IsSet()
}

// HasHomePageUrl returns a boolean if a field has been set.
func (o *LiveTvServiceInfo) HasHomePageUrl() bool {
	if o != nil && o.HomePageUrl.IsSet() {
		return true
	}

	return false
}

// SetHomePageUrl gets a reference to the given NullableString and assigns it to the HomePageUrl field.
func (o *LiveTvServiceInfo) SetHomePageUrl(v string) {
	o.HomePageUrl.Set(&v)
}
// SetHomePageUrlNil sets the value for HomePageUrl to be an explicit nil
func (o *LiveTvServiceInfo) SetHomePageUrlNil() {
	o.HomePageUrl.Set(nil)
}

// UnsetHomePageUrl ensures that no value is present for HomePageUrl, not even an explicit nil
func (o *LiveTvServiceInfo) UnsetHomePageUrl() {
	o.HomePageUrl.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LiveTvServiceInfo) GetStatus() LiveTvServiceStatus {
	if o == nil || IsNil(o.Status) {
		var ret LiveTvServiceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvServiceInfo) GetStatusOk() (*LiveTvServiceStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LiveTvServiceInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given LiveTvServiceStatus and assigns it to the Status field.
func (o *LiveTvServiceInfo) SetStatus(v LiveTvServiceStatus) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvServiceInfo) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage.Get()) {
		var ret string
		return ret
	}
	return *o.StatusMessage.Get()
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvServiceInfo) GetStatusMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusMessage.Get(), o.StatusMessage.IsSet()
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *LiveTvServiceInfo) HasStatusMessage() bool {
	if o != nil && o.StatusMessage.IsSet() {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given NullableString and assigns it to the StatusMessage field.
func (o *LiveTvServiceInfo) SetStatusMessage(v string) {
	o.StatusMessage.Set(&v)
}
// SetStatusMessageNil sets the value for StatusMessage to be an explicit nil
func (o *LiveTvServiceInfo) SetStatusMessageNil() {
	o.StatusMessage.Set(nil)
}

// UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
func (o *LiveTvServiceInfo) UnsetStatusMessage() {
	o.StatusMessage.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvServiceInfo) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvServiceInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *LiveTvServiceInfo) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *LiveTvServiceInfo) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *LiveTvServiceInfo) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *LiveTvServiceInfo) UnsetVersion() {
	o.Version.Unset()
}

// GetHasUpdateAvailable returns the HasUpdateAvailable field value if set, zero value otherwise.
func (o *LiveTvServiceInfo) GetHasUpdateAvailable() bool {
	if o == nil || IsNil(o.HasUpdateAvailable) {
		var ret bool
		return ret
	}
	return *o.HasUpdateAvailable
}

// GetHasUpdateAvailableOk returns a tuple with the HasUpdateAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvServiceInfo) GetHasUpdateAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.HasUpdateAvailable) {
		return nil, false
	}
	return o.HasUpdateAvailable, true
}

// HasHasUpdateAvailable returns a boolean if a field has been set.
func (o *LiveTvServiceInfo) HasHasUpdateAvailable() bool {
	if o != nil && !IsNil(o.HasUpdateAvailable) {
		return true
	}

	return false
}

// SetHasUpdateAvailable gets a reference to the given bool and assigns it to the HasUpdateAvailable field.
func (o *LiveTvServiceInfo) SetHasUpdateAvailable(v bool) {
	o.HasUpdateAvailable = &v
}

// GetIsVisible returns the IsVisible field value if set, zero value otherwise.
func (o *LiveTvServiceInfo) GetIsVisible() bool {
	if o == nil || IsNil(o.IsVisible) {
		var ret bool
		return ret
	}
	return *o.IsVisible
}

// GetIsVisibleOk returns a tuple with the IsVisible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvServiceInfo) GetIsVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVisible) {
		return nil, false
	}
	return o.IsVisible, true
}

// HasIsVisible returns a boolean if a field has been set.
func (o *LiveTvServiceInfo) HasIsVisible() bool {
	if o != nil && !IsNil(o.IsVisible) {
		return true
	}

	return false
}

// SetIsVisible gets a reference to the given bool and assigns it to the IsVisible field.
func (o *LiveTvServiceInfo) SetIsVisible(v bool) {
	o.IsVisible = &v
}

// GetTuners returns the Tuners field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvServiceInfo) GetTuners() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tuners
}

// GetTunersOk returns a tuple with the Tuners field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvServiceInfo) GetTunersOk() ([]string, bool) {
	if o == nil || IsNil(o.Tuners) {
		return nil, false
	}
	return o.Tuners, true
}

// HasTuners returns a boolean if a field has been set.
func (o *LiveTvServiceInfo) HasTuners() bool {
	if o != nil && !IsNil(o.Tuners) {
		return true
	}

	return false
}

// SetTuners gets a reference to the given []string and assigns it to the Tuners field.
func (o *LiveTvServiceInfo) SetTuners(v []string) {
	o.Tuners = v
}

func (o LiveTvServiceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveTvServiceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.HomePageUrl.IsSet() {
		toSerialize["HomePageUrl"] = o.HomePageUrl.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if o.StatusMessage.IsSet() {
		toSerialize["StatusMessage"] = o.StatusMessage.Get()
	}
	if o.Version.IsSet() {
		toSerialize["Version"] = o.Version.Get()
	}
	if !IsNil(o.HasUpdateAvailable) {
		toSerialize["HasUpdateAvailable"] = o.HasUpdateAvailable
	}
	if !IsNil(o.IsVisible) {
		toSerialize["IsVisible"] = o.IsVisible
	}
	if o.Tuners != nil {
		toSerialize["Tuners"] = o.Tuners
	}
	return toSerialize, nil
}

type NullableLiveTvServiceInfo struct {
	value *LiveTvServiceInfo
	isSet bool
}

func (v NullableLiveTvServiceInfo) Get() *LiveTvServiceInfo {
	return v.value
}

func (v *NullableLiveTvServiceInfo) Set(val *LiveTvServiceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveTvServiceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveTvServiceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveTvServiceInfo(val *LiveTvServiceInfo) *NullableLiveTvServiceInfo {
	return &NullableLiveTvServiceInfo{value: val, isSet: true}
}

func (v NullableLiveTvServiceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveTvServiceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


