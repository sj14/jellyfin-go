/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the PublicSystemInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicSystemInfo{}

// PublicSystemInfo struct for PublicSystemInfo
type PublicSystemInfo struct {
	// Gets or sets the local address.
	LocalAddress NullableString `json:"LocalAddress,omitempty"`
	// Gets or sets the name of the server.
	ServerName NullableString `json:"ServerName,omitempty"`
	// Gets or sets the server version.
	Version NullableString `json:"Version,omitempty"`
	// Gets or sets the product name. This is the AssemblyProduct name.
	ProductName NullableString `json:"ProductName,omitempty"`
	// Gets or sets the operating system.
	// Deprecated
	OperatingSystem NullableString `json:"OperatingSystem,omitempty"`
	// Gets or sets the id.
	Id NullableString `json:"Id,omitempty"`
	// Gets or sets a value indicating whether the startup wizard is completed.
	StartupWizardCompleted NullableBool `json:"StartupWizardCompleted,omitempty"`
}

// NewPublicSystemInfo instantiates a new PublicSystemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSystemInfo() *PublicSystemInfo {
	this := PublicSystemInfo{}
	return &this
}

// NewPublicSystemInfoWithDefaults instantiates a new PublicSystemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSystemInfoWithDefaults() *PublicSystemInfo {
	this := PublicSystemInfo{}
	return &this
}

// GetLocalAddress returns the LocalAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicSystemInfo) GetLocalAddress() string {
	if o == nil || IsNil(o.LocalAddress.Get()) {
		var ret string
		return ret
	}
	return *o.LocalAddress.Get()
}

// GetLocalAddressOk returns a tuple with the LocalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicSystemInfo) GetLocalAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalAddress.Get(), o.LocalAddress.IsSet()
}

// HasLocalAddress returns a boolean if a field has been set.
func (o *PublicSystemInfo) HasLocalAddress() bool {
	if o != nil && o.LocalAddress.IsSet() {
		return true
	}

	return false
}

// SetLocalAddress gets a reference to the given NullableString and assigns it to the LocalAddress field.
func (o *PublicSystemInfo) SetLocalAddress(v string) {
	o.LocalAddress.Set(&v)
}
// SetLocalAddressNil sets the value for LocalAddress to be an explicit nil
func (o *PublicSystemInfo) SetLocalAddressNil() {
	o.LocalAddress.Set(nil)
}

// UnsetLocalAddress ensures that no value is present for LocalAddress, not even an explicit nil
func (o *PublicSystemInfo) UnsetLocalAddress() {
	o.LocalAddress.Unset()
}

// GetServerName returns the ServerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicSystemInfo) GetServerName() string {
	if o == nil || IsNil(o.ServerName.Get()) {
		var ret string
		return ret
	}
	return *o.ServerName.Get()
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicSystemInfo) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerName.Get(), o.ServerName.IsSet()
}

// HasServerName returns a boolean if a field has been set.
func (o *PublicSystemInfo) HasServerName() bool {
	if o != nil && o.ServerName.IsSet() {
		return true
	}

	return false
}

// SetServerName gets a reference to the given NullableString and assigns it to the ServerName field.
func (o *PublicSystemInfo) SetServerName(v string) {
	o.ServerName.Set(&v)
}
// SetServerNameNil sets the value for ServerName to be an explicit nil
func (o *PublicSystemInfo) SetServerNameNil() {
	o.ServerName.Set(nil)
}

// UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
func (o *PublicSystemInfo) UnsetServerName() {
	o.ServerName.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicSystemInfo) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicSystemInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *PublicSystemInfo) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *PublicSystemInfo) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *PublicSystemInfo) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *PublicSystemInfo) UnsetVersion() {
	o.Version.Unset()
}

// GetProductName returns the ProductName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicSystemInfo) GetProductName() string {
	if o == nil || IsNil(o.ProductName.Get()) {
		var ret string
		return ret
	}
	return *o.ProductName.Get()
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicSystemInfo) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductName.Get(), o.ProductName.IsSet()
}

// HasProductName returns a boolean if a field has been set.
func (o *PublicSystemInfo) HasProductName() bool {
	if o != nil && o.ProductName.IsSet() {
		return true
	}

	return false
}

// SetProductName gets a reference to the given NullableString and assigns it to the ProductName field.
func (o *PublicSystemInfo) SetProductName(v string) {
	o.ProductName.Set(&v)
}
// SetProductNameNil sets the value for ProductName to be an explicit nil
func (o *PublicSystemInfo) SetProductNameNil() {
	o.ProductName.Set(nil)
}

// UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
func (o *PublicSystemInfo) UnsetProductName() {
	o.ProductName.Unset()
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *PublicSystemInfo) GetOperatingSystem() string {
	if o == nil || IsNil(o.OperatingSystem.Get()) {
		var ret string
		return ret
	}
	return *o.OperatingSystem.Get()
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *PublicSystemInfo) GetOperatingSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperatingSystem.Get(), o.OperatingSystem.IsSet()
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *PublicSystemInfo) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem.IsSet() {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given NullableString and assigns it to the OperatingSystem field.
// Deprecated
func (o *PublicSystemInfo) SetOperatingSystem(v string) {
	o.OperatingSystem.Set(&v)
}
// SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil
func (o *PublicSystemInfo) SetOperatingSystemNil() {
	o.OperatingSystem.Set(nil)
}

// UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
func (o *PublicSystemInfo) UnsetOperatingSystem() {
	o.OperatingSystem.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicSystemInfo) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicSystemInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *PublicSystemInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *PublicSystemInfo) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *PublicSystemInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *PublicSystemInfo) UnsetId() {
	o.Id.Unset()
}

// GetStartupWizardCompleted returns the StartupWizardCompleted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicSystemInfo) GetStartupWizardCompleted() bool {
	if o == nil || IsNil(o.StartupWizardCompleted.Get()) {
		var ret bool
		return ret
	}
	return *o.StartupWizardCompleted.Get()
}

// GetStartupWizardCompletedOk returns a tuple with the StartupWizardCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicSystemInfo) GetStartupWizardCompletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartupWizardCompleted.Get(), o.StartupWizardCompleted.IsSet()
}

// HasStartupWizardCompleted returns a boolean if a field has been set.
func (o *PublicSystemInfo) HasStartupWizardCompleted() bool {
	if o != nil && o.StartupWizardCompleted.IsSet() {
		return true
	}

	return false
}

// SetStartupWizardCompleted gets a reference to the given NullableBool and assigns it to the StartupWizardCompleted field.
func (o *PublicSystemInfo) SetStartupWizardCompleted(v bool) {
	o.StartupWizardCompleted.Set(&v)
}
// SetStartupWizardCompletedNil sets the value for StartupWizardCompleted to be an explicit nil
func (o *PublicSystemInfo) SetStartupWizardCompletedNil() {
	o.StartupWizardCompleted.Set(nil)
}

// UnsetStartupWizardCompleted ensures that no value is present for StartupWizardCompleted, not even an explicit nil
func (o *PublicSystemInfo) UnsetStartupWizardCompleted() {
	o.StartupWizardCompleted.Unset()
}

func (o PublicSystemInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicSystemInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LocalAddress.IsSet() {
		toSerialize["LocalAddress"] = o.LocalAddress.Get()
	}
	if o.ServerName.IsSet() {
		toSerialize["ServerName"] = o.ServerName.Get()
	}
	if o.Version.IsSet() {
		toSerialize["Version"] = o.Version.Get()
	}
	if o.ProductName.IsSet() {
		toSerialize["ProductName"] = o.ProductName.Get()
	}
	if o.OperatingSystem.IsSet() {
		toSerialize["OperatingSystem"] = o.OperatingSystem.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.StartupWizardCompleted.IsSet() {
		toSerialize["StartupWizardCompleted"] = o.StartupWizardCompleted.Get()
	}
	return toSerialize, nil
}

type NullablePublicSystemInfo struct {
	value *PublicSystemInfo
	isSet bool
}

func (v NullablePublicSystemInfo) Get() *PublicSystemInfo {
	return v.value
}

func (v *NullablePublicSystemInfo) Set(val *PublicSystemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSystemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSystemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSystemInfo(val *PublicSystemInfo) *NullablePublicSystemInfo {
	return &NullablePublicSystemInfo{value: val, isSet: true}
}

func (v NullablePublicSystemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSystemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


