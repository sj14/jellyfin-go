/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the DeviceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceInfo{}

// DeviceInfo struct for DeviceInfo
type DeviceInfo struct {
	Name NullableString `json:"Name,omitempty"`
	CustomName NullableString `json:"CustomName,omitempty"`
	// Gets or sets the access token.
	AccessToken NullableString `json:"AccessToken,omitempty"`
	// Gets or sets the identifier.
	Id NullableString `json:"Id,omitempty"`
	// Gets or sets the last name of the user.
	LastUserName NullableString `json:"LastUserName,omitempty"`
	// Gets or sets the name of the application.
	AppName NullableString `json:"AppName,omitempty"`
	// Gets or sets the application version.
	AppVersion NullableString `json:"AppVersion,omitempty"`
	// Gets or sets the last user identifier.
	LastUserId *string `json:"LastUserId,omitempty"`
	// Gets or sets the date last modified.
	DateLastActivity *time.Time `json:"DateLastActivity,omitempty"`
	Capabilities NullableDeviceInfoCapabilities `json:"Capabilities,omitempty"`
	IconUrl NullableString `json:"IconUrl,omitempty"`
}

// NewDeviceInfo instantiates a new DeviceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceInfo() *DeviceInfo {
	this := DeviceInfo{}
	return &this
}

// NewDeviceInfoWithDefaults instantiates a new DeviceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceInfoWithDefaults() *DeviceInfo {
	this := DeviceInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DeviceInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DeviceInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DeviceInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DeviceInfo) UnsetName() {
	o.Name.Unset()
}

// GetCustomName returns the CustomName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfo) GetCustomName() string {
	if o == nil || IsNil(o.CustomName.Get()) {
		var ret string
		return ret
	}
	return *o.CustomName.Get()
}

// GetCustomNameOk returns a tuple with the CustomName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfo) GetCustomNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomName.Get(), o.CustomName.IsSet()
}

// HasCustomName returns a boolean if a field has been set.
func (o *DeviceInfo) HasCustomName() bool {
	if o != nil && o.CustomName.IsSet() {
		return true
	}

	return false
}

// SetCustomName gets a reference to the given NullableString and assigns it to the CustomName field.
func (o *DeviceInfo) SetCustomName(v string) {
	o.CustomName.Set(&v)
}
// SetCustomNameNil sets the value for CustomName to be an explicit nil
func (o *DeviceInfo) SetCustomNameNil() {
	o.CustomName.Set(nil)
}

// UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
func (o *DeviceInfo) UnsetCustomName() {
	o.CustomName.Unset()
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfo) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfo) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *DeviceInfo) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *DeviceInfo) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *DeviceInfo) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *DeviceInfo) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfo) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DeviceInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DeviceInfo) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DeviceInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DeviceInfo) UnsetId() {
	o.Id.Unset()
}

// GetLastUserName returns the LastUserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfo) GetLastUserName() string {
	if o == nil || IsNil(o.LastUserName.Get()) {
		var ret string
		return ret
	}
	return *o.LastUserName.Get()
}

// GetLastUserNameOk returns a tuple with the LastUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfo) GetLastUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUserName.Get(), o.LastUserName.IsSet()
}

// HasLastUserName returns a boolean if a field has been set.
func (o *DeviceInfo) HasLastUserName() bool {
	if o != nil && o.LastUserName.IsSet() {
		return true
	}

	return false
}

// SetLastUserName gets a reference to the given NullableString and assigns it to the LastUserName field.
func (o *DeviceInfo) SetLastUserName(v string) {
	o.LastUserName.Set(&v)
}
// SetLastUserNameNil sets the value for LastUserName to be an explicit nil
func (o *DeviceInfo) SetLastUserNameNil() {
	o.LastUserName.Set(nil)
}

// UnsetLastUserName ensures that no value is present for LastUserName, not even an explicit nil
func (o *DeviceInfo) UnsetLastUserName() {
	o.LastUserName.Unset()
}

// GetAppName returns the AppName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfo) GetAppName() string {
	if o == nil || IsNil(o.AppName.Get()) {
		var ret string
		return ret
	}
	return *o.AppName.Get()
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfo) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppName.Get(), o.AppName.IsSet()
}

// HasAppName returns a boolean if a field has been set.
func (o *DeviceInfo) HasAppName() bool {
	if o != nil && o.AppName.IsSet() {
		return true
	}

	return false
}

// SetAppName gets a reference to the given NullableString and assigns it to the AppName field.
func (o *DeviceInfo) SetAppName(v string) {
	o.AppName.Set(&v)
}
// SetAppNameNil sets the value for AppName to be an explicit nil
func (o *DeviceInfo) SetAppNameNil() {
	o.AppName.Set(nil)
}

// UnsetAppName ensures that no value is present for AppName, not even an explicit nil
func (o *DeviceInfo) UnsetAppName() {
	o.AppName.Unset()
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfo) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion.Get()) {
		var ret string
		return ret
	}
	return *o.AppVersion.Get()
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfo) GetAppVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppVersion.Get(), o.AppVersion.IsSet()
}

// HasAppVersion returns a boolean if a field has been set.
func (o *DeviceInfo) HasAppVersion() bool {
	if o != nil && o.AppVersion.IsSet() {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given NullableString and assigns it to the AppVersion field.
func (o *DeviceInfo) SetAppVersion(v string) {
	o.AppVersion.Set(&v)
}
// SetAppVersionNil sets the value for AppVersion to be an explicit nil
func (o *DeviceInfo) SetAppVersionNil() {
	o.AppVersion.Set(nil)
}

// UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
func (o *DeviceInfo) UnsetAppVersion() {
	o.AppVersion.Unset()
}

// GetLastUserId returns the LastUserId field value if set, zero value otherwise.
func (o *DeviceInfo) GetLastUserId() string {
	if o == nil || IsNil(o.LastUserId) {
		var ret string
		return ret
	}
	return *o.LastUserId
}

// GetLastUserIdOk returns a tuple with the LastUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetLastUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastUserId) {
		return nil, false
	}
	return o.LastUserId, true
}

// HasLastUserId returns a boolean if a field has been set.
func (o *DeviceInfo) HasLastUserId() bool {
	if o != nil && !IsNil(o.LastUserId) {
		return true
	}

	return false
}

// SetLastUserId gets a reference to the given string and assigns it to the LastUserId field.
func (o *DeviceInfo) SetLastUserId(v string) {
	o.LastUserId = &v
}

// GetDateLastActivity returns the DateLastActivity field value if set, zero value otherwise.
func (o *DeviceInfo) GetDateLastActivity() time.Time {
	if o == nil || IsNil(o.DateLastActivity) {
		var ret time.Time
		return ret
	}
	return *o.DateLastActivity
}

// GetDateLastActivityOk returns a tuple with the DateLastActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfo) GetDateLastActivityOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateLastActivity) {
		return nil, false
	}
	return o.DateLastActivity, true
}

// HasDateLastActivity returns a boolean if a field has been set.
func (o *DeviceInfo) HasDateLastActivity() bool {
	if o != nil && !IsNil(o.DateLastActivity) {
		return true
	}

	return false
}

// SetDateLastActivity gets a reference to the given time.Time and assigns it to the DateLastActivity field.
func (o *DeviceInfo) SetDateLastActivity(v time.Time) {
	o.DateLastActivity = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfo) GetCapabilities() DeviceInfoCapabilities {
	if o == nil || IsNil(o.Capabilities.Get()) {
		var ret DeviceInfoCapabilities
		return ret
	}
	return *o.Capabilities.Get()
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfo) GetCapabilitiesOk() (*DeviceInfoCapabilities, bool) {
	if o == nil {
		return nil, false
	}
	return o.Capabilities.Get(), o.Capabilities.IsSet()
}

// HasCapabilities returns a boolean if a field has been set.
func (o *DeviceInfo) HasCapabilities() bool {
	if o != nil && o.Capabilities.IsSet() {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given NullableDeviceInfoCapabilities and assigns it to the Capabilities field.
func (o *DeviceInfo) SetCapabilities(v DeviceInfoCapabilities) {
	o.Capabilities.Set(&v)
}
// SetCapabilitiesNil sets the value for Capabilities to be an explicit nil
func (o *DeviceInfo) SetCapabilitiesNil() {
	o.Capabilities.Set(nil)
}

// UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
func (o *DeviceInfo) UnsetCapabilities() {
	o.Capabilities.Unset()
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfo) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfo) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *DeviceInfo) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *DeviceInfo) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}
// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *DeviceInfo) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *DeviceInfo) UnsetIconUrl() {
	o.IconUrl.Unset()
}

func (o DeviceInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.CustomName.IsSet() {
		toSerialize["CustomName"] = o.CustomName.Get()
	}
	if o.AccessToken.IsSet() {
		toSerialize["AccessToken"] = o.AccessToken.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.LastUserName.IsSet() {
		toSerialize["LastUserName"] = o.LastUserName.Get()
	}
	if o.AppName.IsSet() {
		toSerialize["AppName"] = o.AppName.Get()
	}
	if o.AppVersion.IsSet() {
		toSerialize["AppVersion"] = o.AppVersion.Get()
	}
	if !IsNil(o.LastUserId) {
		toSerialize["LastUserId"] = o.LastUserId
	}
	if !IsNil(o.DateLastActivity) {
		toSerialize["DateLastActivity"] = o.DateLastActivity
	}
	if o.Capabilities.IsSet() {
		toSerialize["Capabilities"] = o.Capabilities.Get()
	}
	if o.IconUrl.IsSet() {
		toSerialize["IconUrl"] = o.IconUrl.Get()
	}
	return toSerialize, nil
}

type NullableDeviceInfo struct {
	value *DeviceInfo
	isSet bool
}

func (v NullableDeviceInfo) Get() *DeviceInfo {
	return v.value
}

func (v *NullableDeviceInfo) Set(val *DeviceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceInfo(val *DeviceInfo) *NullableDeviceInfo {
	return &NullableDeviceInfo{value: val, isSet: true}
}

func (v NullableDeviceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


