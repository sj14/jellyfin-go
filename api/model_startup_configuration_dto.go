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

// checks if the StartupConfigurationDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartupConfigurationDto{}

// StartupConfigurationDto The startup configuration DTO.
type StartupConfigurationDto struct {
	// Gets or sets the server name.
	ServerName NullableString `json:"ServerName,omitempty"`
	// Gets or sets UI language culture.
	UICulture NullableString `json:"UICulture,omitempty"`
	// Gets or sets the metadata country code.
	MetadataCountryCode NullableString `json:"MetadataCountryCode,omitempty"`
	// Gets or sets the preferred language for the metadata.
	PreferredMetadataLanguage NullableString `json:"PreferredMetadataLanguage,omitempty"`
}

// NewStartupConfigurationDto instantiates a new StartupConfigurationDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartupConfigurationDto() *StartupConfigurationDto {
	this := StartupConfigurationDto{}
	return &this
}

// NewStartupConfigurationDtoWithDefaults instantiates a new StartupConfigurationDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartupConfigurationDtoWithDefaults() *StartupConfigurationDto {
	this := StartupConfigurationDto{}
	return &this
}

// GetServerName returns the ServerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartupConfigurationDto) GetServerName() string {
	if o == nil || IsNil(o.ServerName.Get()) {
		var ret string
		return ret
	}
	return *o.ServerName.Get()
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartupConfigurationDto) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerName.Get(), o.ServerName.IsSet()
}

// HasServerName returns a boolean if a field has been set.
func (o *StartupConfigurationDto) HasServerName() bool {
	if o != nil && o.ServerName.IsSet() {
		return true
	}

	return false
}

// SetServerName gets a reference to the given NullableString and assigns it to the ServerName field.
func (o *StartupConfigurationDto) SetServerName(v string) {
	o.ServerName.Set(&v)
}
// SetServerNameNil sets the value for ServerName to be an explicit nil
func (o *StartupConfigurationDto) SetServerNameNil() {
	o.ServerName.Set(nil)
}

// UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
func (o *StartupConfigurationDto) UnsetServerName() {
	o.ServerName.Unset()
}

// GetUICulture returns the UICulture field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartupConfigurationDto) GetUICulture() string {
	if o == nil || IsNil(o.UICulture.Get()) {
		var ret string
		return ret
	}
	return *o.UICulture.Get()
}

// GetUICultureOk returns a tuple with the UICulture field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartupConfigurationDto) GetUICultureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UICulture.Get(), o.UICulture.IsSet()
}

// HasUICulture returns a boolean if a field has been set.
func (o *StartupConfigurationDto) HasUICulture() bool {
	if o != nil && o.UICulture.IsSet() {
		return true
	}

	return false
}

// SetUICulture gets a reference to the given NullableString and assigns it to the UICulture field.
func (o *StartupConfigurationDto) SetUICulture(v string) {
	o.UICulture.Set(&v)
}
// SetUICultureNil sets the value for UICulture to be an explicit nil
func (o *StartupConfigurationDto) SetUICultureNil() {
	o.UICulture.Set(nil)
}

// UnsetUICulture ensures that no value is present for UICulture, not even an explicit nil
func (o *StartupConfigurationDto) UnsetUICulture() {
	o.UICulture.Unset()
}

// GetMetadataCountryCode returns the MetadataCountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartupConfigurationDto) GetMetadataCountryCode() string {
	if o == nil || IsNil(o.MetadataCountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataCountryCode.Get()
}

// GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartupConfigurationDto) GetMetadataCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataCountryCode.Get(), o.MetadataCountryCode.IsSet()
}

// HasMetadataCountryCode returns a boolean if a field has been set.
func (o *StartupConfigurationDto) HasMetadataCountryCode() bool {
	if o != nil && o.MetadataCountryCode.IsSet() {
		return true
	}

	return false
}

// SetMetadataCountryCode gets a reference to the given NullableString and assigns it to the MetadataCountryCode field.
func (o *StartupConfigurationDto) SetMetadataCountryCode(v string) {
	o.MetadataCountryCode.Set(&v)
}
// SetMetadataCountryCodeNil sets the value for MetadataCountryCode to be an explicit nil
func (o *StartupConfigurationDto) SetMetadataCountryCodeNil() {
	o.MetadataCountryCode.Set(nil)
}

// UnsetMetadataCountryCode ensures that no value is present for MetadataCountryCode, not even an explicit nil
func (o *StartupConfigurationDto) UnsetMetadataCountryCode() {
	o.MetadataCountryCode.Unset()
}

// GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StartupConfigurationDto) GetPreferredMetadataLanguage() string {
	if o == nil || IsNil(o.PreferredMetadataLanguage.Get()) {
		var ret string
		return ret
	}
	return *o.PreferredMetadataLanguage.Get()
}

// GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StartupConfigurationDto) GetPreferredMetadataLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreferredMetadataLanguage.Get(), o.PreferredMetadataLanguage.IsSet()
}

// HasPreferredMetadataLanguage returns a boolean if a field has been set.
func (o *StartupConfigurationDto) HasPreferredMetadataLanguage() bool {
	if o != nil && o.PreferredMetadataLanguage.IsSet() {
		return true
	}

	return false
}

// SetPreferredMetadataLanguage gets a reference to the given NullableString and assigns it to the PreferredMetadataLanguage field.
func (o *StartupConfigurationDto) SetPreferredMetadataLanguage(v string) {
	o.PreferredMetadataLanguage.Set(&v)
}
// SetPreferredMetadataLanguageNil sets the value for PreferredMetadataLanguage to be an explicit nil
func (o *StartupConfigurationDto) SetPreferredMetadataLanguageNil() {
	o.PreferredMetadataLanguage.Set(nil)
}

// UnsetPreferredMetadataLanguage ensures that no value is present for PreferredMetadataLanguage, not even an explicit nil
func (o *StartupConfigurationDto) UnsetPreferredMetadataLanguage() {
	o.PreferredMetadataLanguage.Unset()
}

func (o StartupConfigurationDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartupConfigurationDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ServerName.IsSet() {
		toSerialize["ServerName"] = o.ServerName.Get()
	}
	if o.UICulture.IsSet() {
		toSerialize["UICulture"] = o.UICulture.Get()
	}
	if o.MetadataCountryCode.IsSet() {
		toSerialize["MetadataCountryCode"] = o.MetadataCountryCode.Get()
	}
	if o.PreferredMetadataLanguage.IsSet() {
		toSerialize["PreferredMetadataLanguage"] = o.PreferredMetadataLanguage.Get()
	}
	return toSerialize, nil
}

type NullableStartupConfigurationDto struct {
	value *StartupConfigurationDto
	isSet bool
}

func (v NullableStartupConfigurationDto) Get() *StartupConfigurationDto {
	return v.value
}

func (v *NullableStartupConfigurationDto) Set(val *StartupConfigurationDto) {
	v.value = val
	v.isSet = true
}

func (v NullableStartupConfigurationDto) IsSet() bool {
	return v.isSet
}

func (v *NullableStartupConfigurationDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartupConfigurationDto(val *StartupConfigurationDto) *NullableStartupConfigurationDto {
	return &NullableStartupConfigurationDto{value: val, isSet: true}
}

func (v NullableStartupConfigurationDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartupConfigurationDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


