/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ClientCapabilitiesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientCapabilitiesDto{}

// ClientCapabilitiesDto Client capabilities dto.
type ClientCapabilitiesDto struct {
	// Gets or sets the list of playable media types.
	PlayableMediaTypes []MediaType `json:"PlayableMediaTypes,omitempty"`
	// Gets or sets the list of supported commands.
	SupportedCommands []GeneralCommandType `json:"SupportedCommands,omitempty"`
	// Gets or sets a value indicating whether session supports media control.
	SupportsMediaControl *bool `json:"SupportsMediaControl,omitempty"`
	// Gets or sets a value indicating whether session supports a persistent identifier.
	SupportsPersistentIdentifier *bool `json:"SupportsPersistentIdentifier,omitempty"`
	// Gets or sets the device profile.
	DeviceProfile NullableDeviceProfile `json:"DeviceProfile,omitempty"`
	// Gets or sets the app store url.
	AppStoreUrl NullableString `json:"AppStoreUrl,omitempty"`
	// Gets or sets the icon url.
	IconUrl NullableString `json:"IconUrl,omitempty"`
}

// NewClientCapabilitiesDto instantiates a new ClientCapabilitiesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientCapabilitiesDto() *ClientCapabilitiesDto {
	this := ClientCapabilitiesDto{}
	return &this
}

// NewClientCapabilitiesDtoWithDefaults instantiates a new ClientCapabilitiesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientCapabilitiesDtoWithDefaults() *ClientCapabilitiesDto {
	this := ClientCapabilitiesDto{}
	return &this
}

// GetPlayableMediaTypes returns the PlayableMediaTypes field value if set, zero value otherwise.
func (o *ClientCapabilitiesDto) GetPlayableMediaTypes() []MediaType {
	if o == nil || IsNil(o.PlayableMediaTypes) {
		var ret []MediaType
		return ret
	}
	return o.PlayableMediaTypes
}

// GetPlayableMediaTypesOk returns a tuple with the PlayableMediaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCapabilitiesDto) GetPlayableMediaTypesOk() ([]MediaType, bool) {
	if o == nil || IsNil(o.PlayableMediaTypes) {
		return nil, false
	}
	return o.PlayableMediaTypes, true
}

// HasPlayableMediaTypes returns a boolean if a field has been set.
func (o *ClientCapabilitiesDto) HasPlayableMediaTypes() bool {
	if o != nil && !IsNil(o.PlayableMediaTypes) {
		return true
	}

	return false
}

// SetPlayableMediaTypes gets a reference to the given []MediaType and assigns it to the PlayableMediaTypes field.
func (o *ClientCapabilitiesDto) SetPlayableMediaTypes(v []MediaType) {
	o.PlayableMediaTypes = v
}

// GetSupportedCommands returns the SupportedCommands field value if set, zero value otherwise.
func (o *ClientCapabilitiesDto) GetSupportedCommands() []GeneralCommandType {
	if o == nil || IsNil(o.SupportedCommands) {
		var ret []GeneralCommandType
		return ret
	}
	return o.SupportedCommands
}

// GetSupportedCommandsOk returns a tuple with the SupportedCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCapabilitiesDto) GetSupportedCommandsOk() ([]GeneralCommandType, bool) {
	if o == nil || IsNil(o.SupportedCommands) {
		return nil, false
	}
	return o.SupportedCommands, true
}

// HasSupportedCommands returns a boolean if a field has been set.
func (o *ClientCapabilitiesDto) HasSupportedCommands() bool {
	if o != nil && !IsNil(o.SupportedCommands) {
		return true
	}

	return false
}

// SetSupportedCommands gets a reference to the given []GeneralCommandType and assigns it to the SupportedCommands field.
func (o *ClientCapabilitiesDto) SetSupportedCommands(v []GeneralCommandType) {
	o.SupportedCommands = v
}

// GetSupportsMediaControl returns the SupportsMediaControl field value if set, zero value otherwise.
func (o *ClientCapabilitiesDto) GetSupportsMediaControl() bool {
	if o == nil || IsNil(o.SupportsMediaControl) {
		var ret bool
		return ret
	}
	return *o.SupportsMediaControl
}

// GetSupportsMediaControlOk returns a tuple with the SupportsMediaControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCapabilitiesDto) GetSupportsMediaControlOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsMediaControl) {
		return nil, false
	}
	return o.SupportsMediaControl, true
}

// HasSupportsMediaControl returns a boolean if a field has been set.
func (o *ClientCapabilitiesDto) HasSupportsMediaControl() bool {
	if o != nil && !IsNil(o.SupportsMediaControl) {
		return true
	}

	return false
}

// SetSupportsMediaControl gets a reference to the given bool and assigns it to the SupportsMediaControl field.
func (o *ClientCapabilitiesDto) SetSupportsMediaControl(v bool) {
	o.SupportsMediaControl = &v
}

// GetSupportsPersistentIdentifier returns the SupportsPersistentIdentifier field value if set, zero value otherwise.
func (o *ClientCapabilitiesDto) GetSupportsPersistentIdentifier() bool {
	if o == nil || IsNil(o.SupportsPersistentIdentifier) {
		var ret bool
		return ret
	}
	return *o.SupportsPersistentIdentifier
}

// GetSupportsPersistentIdentifierOk returns a tuple with the SupportsPersistentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientCapabilitiesDto) GetSupportsPersistentIdentifierOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsPersistentIdentifier) {
		return nil, false
	}
	return o.SupportsPersistentIdentifier, true
}

// HasSupportsPersistentIdentifier returns a boolean if a field has been set.
func (o *ClientCapabilitiesDto) HasSupportsPersistentIdentifier() bool {
	if o != nil && !IsNil(o.SupportsPersistentIdentifier) {
		return true
	}

	return false
}

// SetSupportsPersistentIdentifier gets a reference to the given bool and assigns it to the SupportsPersistentIdentifier field.
func (o *ClientCapabilitiesDto) SetSupportsPersistentIdentifier(v bool) {
	o.SupportsPersistentIdentifier = &v
}

// GetDeviceProfile returns the DeviceProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClientCapabilitiesDto) GetDeviceProfile() DeviceProfile {
	if o == nil || IsNil(o.DeviceProfile.Get()) {
		var ret DeviceProfile
		return ret
	}
	return *o.DeviceProfile.Get()
}

// GetDeviceProfileOk returns a tuple with the DeviceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientCapabilitiesDto) GetDeviceProfileOk() (*DeviceProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceProfile.Get(), o.DeviceProfile.IsSet()
}

// HasDeviceProfile returns a boolean if a field has been set.
func (o *ClientCapabilitiesDto) HasDeviceProfile() bool {
	if o != nil && o.DeviceProfile.IsSet() {
		return true
	}

	return false
}

// SetDeviceProfile gets a reference to the given NullableDeviceProfile and assigns it to the DeviceProfile field.
func (o *ClientCapabilitiesDto) SetDeviceProfile(v DeviceProfile) {
	o.DeviceProfile.Set(&v)
}
// SetDeviceProfileNil sets the value for DeviceProfile to be an explicit nil
func (o *ClientCapabilitiesDto) SetDeviceProfileNil() {
	o.DeviceProfile.Set(nil)
}

// UnsetDeviceProfile ensures that no value is present for DeviceProfile, not even an explicit nil
func (o *ClientCapabilitiesDto) UnsetDeviceProfile() {
	o.DeviceProfile.Unset()
}

// GetAppStoreUrl returns the AppStoreUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClientCapabilitiesDto) GetAppStoreUrl() string {
	if o == nil || IsNil(o.AppStoreUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AppStoreUrl.Get()
}

// GetAppStoreUrlOk returns a tuple with the AppStoreUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientCapabilitiesDto) GetAppStoreUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppStoreUrl.Get(), o.AppStoreUrl.IsSet()
}

// HasAppStoreUrl returns a boolean if a field has been set.
func (o *ClientCapabilitiesDto) HasAppStoreUrl() bool {
	if o != nil && o.AppStoreUrl.IsSet() {
		return true
	}

	return false
}

// SetAppStoreUrl gets a reference to the given NullableString and assigns it to the AppStoreUrl field.
func (o *ClientCapabilitiesDto) SetAppStoreUrl(v string) {
	o.AppStoreUrl.Set(&v)
}
// SetAppStoreUrlNil sets the value for AppStoreUrl to be an explicit nil
func (o *ClientCapabilitiesDto) SetAppStoreUrlNil() {
	o.AppStoreUrl.Set(nil)
}

// UnsetAppStoreUrl ensures that no value is present for AppStoreUrl, not even an explicit nil
func (o *ClientCapabilitiesDto) UnsetAppStoreUrl() {
	o.AppStoreUrl.Unset()
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClientCapabilitiesDto) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientCapabilitiesDto) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *ClientCapabilitiesDto) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *ClientCapabilitiesDto) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}
// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *ClientCapabilitiesDto) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *ClientCapabilitiesDto) UnsetIconUrl() {
	o.IconUrl.Unset()
}

func (o ClientCapabilitiesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientCapabilitiesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlayableMediaTypes) {
		toSerialize["PlayableMediaTypes"] = o.PlayableMediaTypes
	}
	if !IsNil(o.SupportedCommands) {
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
	return toSerialize, nil
}

type NullableClientCapabilitiesDto struct {
	value *ClientCapabilitiesDto
	isSet bool
}

func (v NullableClientCapabilitiesDto) Get() *ClientCapabilitiesDto {
	return v.value
}

func (v *NullableClientCapabilitiesDto) Set(val *ClientCapabilitiesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableClientCapabilitiesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableClientCapabilitiesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientCapabilitiesDto(val *ClientCapabilitiesDto) *NullableClientCapabilitiesDto {
	return &NullableClientCapabilitiesDto{value: val, isSet: true}
}

func (v NullableClientCapabilitiesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientCapabilitiesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


