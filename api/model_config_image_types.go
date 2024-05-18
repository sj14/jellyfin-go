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

// checks if the ConfigImageTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigImageTypes{}

// ConfigImageTypes struct for ConfigImageTypes
type ConfigImageTypes struct {
	BackdropSizes []string `json:"BackdropSizes,omitempty"`
	BaseUrl NullableString `json:"BaseUrl,omitempty"`
	LogoSizes []string `json:"LogoSizes,omitempty"`
	PosterSizes []string `json:"PosterSizes,omitempty"`
	ProfileSizes []string `json:"ProfileSizes,omitempty"`
	SecureBaseUrl NullableString `json:"SecureBaseUrl,omitempty"`
	StillSizes []string `json:"StillSizes,omitempty"`
}

// NewConfigImageTypes instantiates a new ConfigImageTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigImageTypes() *ConfigImageTypes {
	this := ConfigImageTypes{}
	return &this
}

// NewConfigImageTypesWithDefaults instantiates a new ConfigImageTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigImageTypesWithDefaults() *ConfigImageTypes {
	this := ConfigImageTypes{}
	return &this
}

// GetBackdropSizes returns the BackdropSizes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigImageTypes) GetBackdropSizes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BackdropSizes
}

// GetBackdropSizesOk returns a tuple with the BackdropSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigImageTypes) GetBackdropSizesOk() ([]string, bool) {
	if o == nil || IsNil(o.BackdropSizes) {
		return nil, false
	}
	return o.BackdropSizes, true
}

// HasBackdropSizes returns a boolean if a field has been set.
func (o *ConfigImageTypes) HasBackdropSizes() bool {
	if o != nil && !IsNil(o.BackdropSizes) {
		return true
	}

	return false
}

// SetBackdropSizes gets a reference to the given []string and assigns it to the BackdropSizes field.
func (o *ConfigImageTypes) SetBackdropSizes(v []string) {
	o.BackdropSizes = v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigImageTypes) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.BaseUrl.Get()
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigImageTypes) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseUrl.Get(), o.BaseUrl.IsSet()
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *ConfigImageTypes) HasBaseUrl() bool {
	if o != nil && o.BaseUrl.IsSet() {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given NullableString and assigns it to the BaseUrl field.
func (o *ConfigImageTypes) SetBaseUrl(v string) {
	o.BaseUrl.Set(&v)
}
// SetBaseUrlNil sets the value for BaseUrl to be an explicit nil
func (o *ConfigImageTypes) SetBaseUrlNil() {
	o.BaseUrl.Set(nil)
}

// UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
func (o *ConfigImageTypes) UnsetBaseUrl() {
	o.BaseUrl.Unset()
}

// GetLogoSizes returns the LogoSizes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigImageTypes) GetLogoSizes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.LogoSizes
}

// GetLogoSizesOk returns a tuple with the LogoSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigImageTypes) GetLogoSizesOk() ([]string, bool) {
	if o == nil || IsNil(o.LogoSizes) {
		return nil, false
	}
	return o.LogoSizes, true
}

// HasLogoSizes returns a boolean if a field has been set.
func (o *ConfigImageTypes) HasLogoSizes() bool {
	if o != nil && !IsNil(o.LogoSizes) {
		return true
	}

	return false
}

// SetLogoSizes gets a reference to the given []string and assigns it to the LogoSizes field.
func (o *ConfigImageTypes) SetLogoSizes(v []string) {
	o.LogoSizes = v
}

// GetPosterSizes returns the PosterSizes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigImageTypes) GetPosterSizes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PosterSizes
}

// GetPosterSizesOk returns a tuple with the PosterSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigImageTypes) GetPosterSizesOk() ([]string, bool) {
	if o == nil || IsNil(o.PosterSizes) {
		return nil, false
	}
	return o.PosterSizes, true
}

// HasPosterSizes returns a boolean if a field has been set.
func (o *ConfigImageTypes) HasPosterSizes() bool {
	if o != nil && !IsNil(o.PosterSizes) {
		return true
	}

	return false
}

// SetPosterSizes gets a reference to the given []string and assigns it to the PosterSizes field.
func (o *ConfigImageTypes) SetPosterSizes(v []string) {
	o.PosterSizes = v
}

// GetProfileSizes returns the ProfileSizes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigImageTypes) GetProfileSizes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ProfileSizes
}

// GetProfileSizesOk returns a tuple with the ProfileSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigImageTypes) GetProfileSizesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProfileSizes) {
		return nil, false
	}
	return o.ProfileSizes, true
}

// HasProfileSizes returns a boolean if a field has been set.
func (o *ConfigImageTypes) HasProfileSizes() bool {
	if o != nil && !IsNil(o.ProfileSizes) {
		return true
	}

	return false
}

// SetProfileSizes gets a reference to the given []string and assigns it to the ProfileSizes field.
func (o *ConfigImageTypes) SetProfileSizes(v []string) {
	o.ProfileSizes = v
}

// GetSecureBaseUrl returns the SecureBaseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigImageTypes) GetSecureBaseUrl() string {
	if o == nil || IsNil(o.SecureBaseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.SecureBaseUrl.Get()
}

// GetSecureBaseUrlOk returns a tuple with the SecureBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigImageTypes) GetSecureBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecureBaseUrl.Get(), o.SecureBaseUrl.IsSet()
}

// HasSecureBaseUrl returns a boolean if a field has been set.
func (o *ConfigImageTypes) HasSecureBaseUrl() bool {
	if o != nil && o.SecureBaseUrl.IsSet() {
		return true
	}

	return false
}

// SetSecureBaseUrl gets a reference to the given NullableString and assigns it to the SecureBaseUrl field.
func (o *ConfigImageTypes) SetSecureBaseUrl(v string) {
	o.SecureBaseUrl.Set(&v)
}
// SetSecureBaseUrlNil sets the value for SecureBaseUrl to be an explicit nil
func (o *ConfigImageTypes) SetSecureBaseUrlNil() {
	o.SecureBaseUrl.Set(nil)
}

// UnsetSecureBaseUrl ensures that no value is present for SecureBaseUrl, not even an explicit nil
func (o *ConfigImageTypes) UnsetSecureBaseUrl() {
	o.SecureBaseUrl.Unset()
}

// GetStillSizes returns the StillSizes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigImageTypes) GetStillSizes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.StillSizes
}

// GetStillSizesOk returns a tuple with the StillSizes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigImageTypes) GetStillSizesOk() ([]string, bool) {
	if o == nil || IsNil(o.StillSizes) {
		return nil, false
	}
	return o.StillSizes, true
}

// HasStillSizes returns a boolean if a field has been set.
func (o *ConfigImageTypes) HasStillSizes() bool {
	if o != nil && !IsNil(o.StillSizes) {
		return true
	}

	return false
}

// SetStillSizes gets a reference to the given []string and assigns it to the StillSizes field.
func (o *ConfigImageTypes) SetStillSizes(v []string) {
	o.StillSizes = v
}

func (o ConfigImageTypes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigImageTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BackdropSizes != nil {
		toSerialize["BackdropSizes"] = o.BackdropSizes
	}
	if o.BaseUrl.IsSet() {
		toSerialize["BaseUrl"] = o.BaseUrl.Get()
	}
	if o.LogoSizes != nil {
		toSerialize["LogoSizes"] = o.LogoSizes
	}
	if o.PosterSizes != nil {
		toSerialize["PosterSizes"] = o.PosterSizes
	}
	if o.ProfileSizes != nil {
		toSerialize["ProfileSizes"] = o.ProfileSizes
	}
	if o.SecureBaseUrl.IsSet() {
		toSerialize["SecureBaseUrl"] = o.SecureBaseUrl.Get()
	}
	if o.StillSizes != nil {
		toSerialize["StillSizes"] = o.StillSizes
	}
	return toSerialize, nil
}

type NullableConfigImageTypes struct {
	value *ConfigImageTypes
	isSet bool
}

func (v NullableConfigImageTypes) Get() *ConfigImageTypes {
	return v.value
}

func (v *NullableConfigImageTypes) Set(val *ConfigImageTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigImageTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigImageTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigImageTypes(val *ConfigImageTypes) *NullableConfigImageTypes {
	return &NullableConfigImageTypes{value: val, isSet: true}
}

func (v NullableConfigImageTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigImageTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


