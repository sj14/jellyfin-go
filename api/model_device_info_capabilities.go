/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeviceInfoCapabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceInfoCapabilities{}

// DeviceInfoCapabilities Gets or sets the capabilities.
type DeviceInfoCapabilities struct {
	PlayableMediaTypes []MediaType `json:"PlayableMediaTypes,omitempty"`
	SupportedCommands []GeneralCommandType `json:"SupportedCommands,omitempty"`
	SupportsMediaControl *bool `json:"SupportsMediaControl,omitempty"`
	SupportsPersistentIdentifier *bool `json:"SupportsPersistentIdentifier,omitempty"`
	DeviceProfile NullableClientCapabilitiesDeviceProfile `json:"DeviceProfile,omitempty"`
	AppStoreUrl NullableString `json:"AppStoreUrl,omitempty"`
	IconUrl NullableString `json:"IconUrl,omitempty"`
	// Deprecated
	SupportsContentUploading NullableBool `json:"SupportsContentUploading,omitempty"`
	// Deprecated
	SupportsSync NullableBool `json:"SupportsSync,omitempty"`
}

// NewDeviceInfoCapabilities instantiates a new DeviceInfoCapabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceInfoCapabilities() *DeviceInfoCapabilities {
	this := DeviceInfoCapabilities{}
	var supportsContentUploading bool = false
	this.SupportsContentUploading = *NewNullableBool(&supportsContentUploading)
	var supportsSync bool = false
	this.SupportsSync = *NewNullableBool(&supportsSync)
	return &this
}

// NewDeviceInfoCapabilitiesWithDefaults instantiates a new DeviceInfoCapabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceInfoCapabilitiesWithDefaults() *DeviceInfoCapabilities {
	this := DeviceInfoCapabilities{}
	var supportsContentUploading bool = false
	this.SupportsContentUploading = *NewNullableBool(&supportsContentUploading)
	var supportsSync bool = false
	this.SupportsSync = *NewNullableBool(&supportsSync)
	return &this
}

// GetPlayableMediaTypes returns the PlayableMediaTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoCapabilities) GetPlayableMediaTypes() []MediaType {
	if o == nil {
		var ret []MediaType
		return ret
	}
	return o.PlayableMediaTypes
}

// GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoCapabilities) GetPlayableMediaTypesOk() ([]MediaType, bool) {
	if o == nil || IsNil(o.PlayableMediaTypes) {
		return nil, false
	}
	return o.PlayableMediaTypes, true
}

// HasPlayableMediaTypes returns a boolean if a field has been set.
func (o *DeviceInfoCapabilities) HasPlayableMediaTypes() bool {
	if o != nil && !IsNil(o.PlayableMediaTypes) {
		return true
	}

	return false
}

// SetPlayableMediaTypes gets a reference to the given []MediaType and assigns it to the PlayableMediaTypes field.
func (o *DeviceInfoCapabilities) SetPlayableMediaTypes(v []MediaType) {
	o.PlayableMediaTypes = v
}

// GetSupportedCommands returns the SupportedCommands field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoCapabilities) GetSupportedCommands() []GeneralCommandType {
	if o == nil {
		var ret []GeneralCommandType
		return ret
	}
	return o.SupportedCommands
}

// GetSupportedCommandsOk returns a tuple with the SupportedCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoCapabilities) GetSupportedCommandsOk() ([]GeneralCommandType, bool) {
	if o == nil || IsNil(o.SupportedCommands) {
		return nil, false
	}
	return o.SupportedCommands, true
}

// HasSupportedCommands returns a boolean if a field has been set.
func (o *DeviceInfoCapabilities) HasSupportedCommands() bool {
	if o != nil && !IsNil(o.SupportedCommands) {
		return true
	}

	return false
}

// SetSupportedCommands gets a reference to the given []GeneralCommandType and assigns it to the SupportedCommands field.
func (o *DeviceInfoCapabilities) SetSupportedCommands(v []GeneralCommandType) {
	o.SupportedCommands = v
}

// GetSupportsMediaControl returns the SupportsMediaControl field value if set, zero value otherwise.
func (o *DeviceInfoCapabilities) GetSupportsMediaControl() bool {
	if o == nil || IsNil(o.SupportsMediaControl) {
		var ret bool
		return ret
	}
	return *o.SupportsMediaControl
}

// GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfoCapabilities) GetSupportsMediaControlOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsMediaControl) {
		return nil, false
	}
	return o.SupportsMediaControl, true
}

// HasSupportsMediaControl returns a boolean if a field has been set.
func (o *DeviceInfoCapabilities) HasSupportsMediaControl() bool {
	if o != nil && !IsNil(o.SupportsMediaControl) {
		return true
	}

	return false
}

// SetSupportsMediaControl gets a reference to the given bool and assigns it to the SupportsMediaControl field.
func (o *DeviceInfoCapabilities) SetSupportsMediaControl(v bool) {
	o.SupportsMediaControl = &v
}

// GetSupportsPersistentIdentifier returns the SupportsPersistentIdentifier field value if set, zero value otherwise.
func (o *DeviceInfoCapabilities) GetSupportsPersistentIdentifier() bool {
	if o == nil || IsNil(o.SupportsPersistentIdentifier) {
		var ret bool
		return ret
	}
	return *o.SupportsPersistentIdentifier
}

// GetSupportsPersistentIdentifierOk returns a tuple with the SupportsPersistentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInfoCapabilities) GetSupportsPersistentIdentifierOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsPersistentIdentifier) {
		return nil, false
	}
	return o.SupportsPersistentIdentifier, true
}

// HasSupportsPersistentIdentifier returns a boolean if a field has been set.
func (o *DeviceInfoCapabilities) HasSupportsPersistentIdentifier() bool {
	if o != nil && !IsNil(o.SupportsPersistentIdentifier) {
		return true
	}

	return false
}

// SetSupportsPersistentIdentifier gets a reference to the given bool and assigns it to the SupportsPersistentIdentifier field.
func (o *DeviceInfoCapabilities) SetSupportsPersistentIdentifier(v bool) {
	o.SupportsPersistentIdentifier = &v
}

// GetDeviceProfile returns the DeviceProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoCapabilities) GetDeviceProfile() ClientCapabilitiesDeviceProfile {
	if o == nil || IsNil(o.DeviceProfile.Get()) {
		var ret ClientCapabilitiesDeviceProfile
		return ret
	}
	return *o.DeviceProfile.Get()
}

// GetDeviceProfileOk returns a tuple with the DeviceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoCapabilities) GetDeviceProfileOk() (*ClientCapabilitiesDeviceProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceProfile.Get(), o.DeviceProfile.IsSet()
}

// HasDeviceProfile returns a boolean if a field has been set.
func (o *DeviceInfoCapabilities) HasDeviceProfile() bool {
	if o != nil && o.DeviceProfile.IsSet() {
		return true
	}

	return false
}

// SetDeviceProfile gets a reference to the given NullableClientCapabilitiesDeviceProfile and assigns it to the DeviceProfile field.
func (o *DeviceInfoCapabilities) SetDeviceProfile(v ClientCapabilitiesDeviceProfile) {
	o.DeviceProfile.Set(&v)
}
// SetDeviceProfileNil sets the value for DeviceProfile to be an explicit nil
func (o *DeviceInfoCapabilities) SetDeviceProfileNil() {
	o.DeviceProfile.Set(nil)
}

// UnsetDeviceProfile ensures that no value is present for DeviceProfile, not even an explicit nil
func (o *DeviceInfoCapabilities) UnsetDeviceProfile() {
	o.DeviceProfile.Unset()
}

// GetAppStoreUrl returns the AppStoreUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoCapabilities) GetAppStoreUrl() string {
	if o == nil || IsNil(o.AppStoreUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AppStoreUrl.Get()
}

// GetAppStoreUrlOk returns a tuple with the AppStoreUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoCapabilities) GetAppStoreUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppStoreUrl.Get(), o.AppStoreUrl.IsSet()
}

// HasAppStoreUrl returns a boolean if a field has been set.
func (o *DeviceInfoCapabilities) HasAppStoreUrl() bool {
	if o != nil && o.AppStoreUrl.IsSet() {
		return true
	}

	return false
}

// SetAppStoreUrl gets a reference to the given NullableString and assigns it to the AppStoreUrl field.
func (o *DeviceInfoCapabilities) SetAppStoreUrl(v string) {
	o.AppStoreUrl.Set(&v)
}
// SetAppStoreUrlNil sets the value for AppStoreUrl to be an explicit nil
func (o *DeviceInfoCapabilities) SetAppStoreUrlNil() {
	o.AppStoreUrl.Set(nil)
}

// UnsetAppStoreUrl ensures that no value is present for AppStoreUrl, not even an explicit nil
func (o *DeviceInfoCapabilities) UnsetAppStoreUrl() {
	o.AppStoreUrl.Unset()
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceInfoCapabilities) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceInfoCapabilities) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *DeviceInfoCapabilities) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *DeviceInfoCapabilities) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}
// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *DeviceInfoCapabilities) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *DeviceInfoCapabilities) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetSupportsContentUploading returns the SupportsContentUploading field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *DeviceInfoCapabilities) GetSupportsContentUploading() bool {
	if o == nil || IsNil(o.SupportsContentUploading.Get()) {
		var ret bool
		return ret
	}
	return *o.SupportsContentUploading.Get()
}

// GetSupportsContentUploadingOk returns a tuple with the SupportsContentUploading field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *DeviceInfoCapabilities) GetSupportsContentUploadingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportsContentUploading.Get(), o.SupportsContentUploading.IsSet()
}

// HasSupportsContentUploading returns a boolean if a field has been set.
func (o *DeviceInfoCapabilities) HasSupportsContentUploading() bool {
	if o != nil && o.SupportsContentUploading.IsSet() {
		return true
	}

	return false
}

// SetSupportsContentUploading gets a reference to the given NullableBool and assigns it to the SupportsContentUploading field.
// Deprecated
func (o *DeviceInfoCapabilities) SetSupportsContentUploading(v bool) {
	o.SupportsContentUploading.Set(&v)
}
// SetSupportsContentUploadingNil sets the value for SupportsContentUploading to be an explicit nil
func (o *DeviceInfoCapabilities) SetSupportsContentUploadingNil() {
	o.SupportsContentUploading.Set(nil)
}

// UnsetSupportsContentUploading ensures that no value is present for SupportsContentUploading, not even an explicit nil
func (o *DeviceInfoCapabilities) UnsetSupportsContentUploading() {
	o.SupportsContentUploading.Unset()
}

// GetSupportsSync returns the SupportsSync field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *DeviceInfoCapabilities) GetSupportsSync() bool {
	if o == nil || IsNil(o.SupportsSync.Get()) {
		var ret bool
		return ret
	}
	return *o.SupportsSync.Get()
}

// GetSupportsSyncOk returns a tuple with the SupportsSync field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *DeviceInfoCapabilities) GetSupportsSyncOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportsSync.Get(), o.SupportsSync.IsSet()
}

// HasSupportsSync returns a boolean if a field has been set.
func (o *DeviceInfoCapabilities) HasSupportsSync() bool {
	if o != nil && o.SupportsSync.IsSet() {
		return true
	}

	return false
}

// SetSupportsSync gets a reference to the given NullableBool and assigns it to the SupportsSync field.
// Deprecated
func (o *DeviceInfoCapabilities) SetSupportsSync(v bool) {
	o.SupportsSync.Set(&v)
}
// SetSupportsSyncNil sets the value for SupportsSync to be an explicit nil
func (o *DeviceInfoCapabilities) SetSupportsSyncNil() {
	o.SupportsSync.Set(nil)
}

// UnsetSupportsSync ensures that no value is present for SupportsSync, not even an explicit nil
func (o *DeviceInfoCapabilities) UnsetSupportsSync() {
	o.SupportsSync.Unset()
}

func (o DeviceInfoCapabilities) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceInfoCapabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PlayableMediaTypes != nil {
		toSerialize["PlayableMediaTypes"] = o.PlayableMediaTypes
	}
	if o.SupportedCommands != nil {
		toSerialize["SupportedCommands"] = o.SupportedCommands
	}
	if !IsNil(o.SupportsMediaControl) {
		toSerialize["SupportsMediaControl"] = o.SupportsMediaControl
	}
	if !IsNil(o.SupportsPersistentIdentifier) {
		toSerialize["SupportsPersistentIdentifier"] = o.SupportsPersistentIdentifier
	}
	if o.DeviceProfile.IsSet() {
		toSerialize["DeviceProfile"] = o.DeviceProfile.Get()
	}
	if o.AppStoreUrl.IsSet() {
		toSerialize["AppStoreUrl"] = o.AppStoreUrl.Get()
	}
	if o.IconUrl.IsSet() {
		toSerialize["IconUrl"] = o.IconUrl.Get()
	}
	if o.SupportsContentUploading.IsSet() {
		toSerialize["SupportsContentUploading"] = o.SupportsContentUploading.Get()
	}
	if o.SupportsSync.IsSet() {
		toSerialize["SupportsSync"] = o.SupportsSync.Get()
	}
	return toSerialize, nil
}

type NullableDeviceInfoCapabilities struct {
	value *DeviceInfoCapabilities
	isSet bool
}

func (v NullableDeviceInfoCapabilities) Get() *DeviceInfoCapabilities {
	return v.value
}

func (v *NullableDeviceInfoCapabilities) Set(val *DeviceInfoCapabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceInfoCapabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceInfoCapabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceInfoCapabilities(val *DeviceInfoCapabilities) *NullableDeviceInfoCapabilities {
	return &NullableDeviceInfoCapabilities{value: val, isSet: true}
}

func (v NullableDeviceInfoCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceInfoCapabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


