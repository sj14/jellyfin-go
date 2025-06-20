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

// checks if the CountryInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryInfo{}

// CountryInfo Class CountryInfo.
type CountryInfo struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the display name.
	DisplayName NullableString `json:"DisplayName,omitempty"`
	// Gets or sets the name of the two letter ISO region.
	TwoLetterISORegionName NullableString `json:"TwoLetterISORegionName,omitempty"`
	// Gets or sets the name of the three letter ISO region.
	ThreeLetterISORegionName NullableString `json:"ThreeLetterISORegionName,omitempty"`
}

// NewCountryInfo instantiates a new CountryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryInfo() *CountryInfo {
	this := CountryInfo{}
	return &this
}

// NewCountryInfoWithDefaults instantiates a new CountryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryInfoWithDefaults() *CountryInfo {
	this := CountryInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CountryInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CountryInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CountryInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CountryInfo) UnsetName() {
	o.Name.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryInfo) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CountryInfo) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *CountryInfo) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *CountryInfo) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *CountryInfo) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetTwoLetterISORegionName returns the TwoLetterISORegionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryInfo) GetTwoLetterISORegionName() string {
	if o == nil || IsNil(o.TwoLetterISORegionName.Get()) {
		var ret string
		return ret
	}
	return *o.TwoLetterISORegionName.Get()
}

// GetTwoLetterISORegionNameOk returns a tuple with the TwoLetterISORegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryInfo) GetTwoLetterISORegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwoLetterISORegionName.Get(), o.TwoLetterISORegionName.IsSet()
}

// HasTwoLetterISORegionName returns a boolean if a field has been set.
func (o *CountryInfo) HasTwoLetterISORegionName() bool {
	if o != nil && o.TwoLetterISORegionName.IsSet() {
		return true
	}

	return false
}

// SetTwoLetterISORegionName gets a reference to the given NullableString and assigns it to the TwoLetterISORegionName field.
func (o *CountryInfo) SetTwoLetterISORegionName(v string) {
	o.TwoLetterISORegionName.Set(&v)
}
// SetTwoLetterISORegionNameNil sets the value for TwoLetterISORegionName to be an explicit nil
func (o *CountryInfo) SetTwoLetterISORegionNameNil() {
	o.TwoLetterISORegionName.Set(nil)
}

// UnsetTwoLetterISORegionName ensures that no value is present for TwoLetterISORegionName, not even an explicit nil
func (o *CountryInfo) UnsetTwoLetterISORegionName() {
	o.TwoLetterISORegionName.Unset()
}

// GetThreeLetterISORegionName returns the ThreeLetterISORegionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CountryInfo) GetThreeLetterISORegionName() string {
	if o == nil || IsNil(o.ThreeLetterISORegionName.Get()) {
		var ret string
		return ret
	}
	return *o.ThreeLetterISORegionName.Get()
}

// GetThreeLetterISORegionNameOk returns a tuple with the ThreeLetterISORegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CountryInfo) GetThreeLetterISORegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreeLetterISORegionName.Get(), o.ThreeLetterISORegionName.IsSet()
}

// HasThreeLetterISORegionName returns a boolean if a field has been set.
func (o *CountryInfo) HasThreeLetterISORegionName() bool {
	if o != nil && o.ThreeLetterISORegionName.IsSet() {
		return true
	}

	return false
}

// SetThreeLetterISORegionName gets a reference to the given NullableString and assigns it to the ThreeLetterISORegionName field.
func (o *CountryInfo) SetThreeLetterISORegionName(v string) {
	o.ThreeLetterISORegionName.Set(&v)
}
// SetThreeLetterISORegionNameNil sets the value for ThreeLetterISORegionName to be an explicit nil
func (o *CountryInfo) SetThreeLetterISORegionNameNil() {
	o.ThreeLetterISORegionName.Set(nil)
}

// UnsetThreeLetterISORegionName ensures that no value is present for ThreeLetterISORegionName, not even an explicit nil
func (o *CountryInfo) UnsetThreeLetterISORegionName() {
	o.ThreeLetterISORegionName.Unset()
}

func (o CountryInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountryInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["DisplayName"] = o.DisplayName.Get()
	}
	if o.TwoLetterISORegionName.IsSet() {
		toSerialize["TwoLetterISORegionName"] = o.TwoLetterISORegionName.Get()
	}
	if o.ThreeLetterISORegionName.IsSet() {
		toSerialize["ThreeLetterISORegionName"] = o.ThreeLetterISORegionName.Get()
	}
	return toSerialize, nil
}

type NullableCountryInfo struct {
	value *CountryInfo
	isSet bool
}

func (v NullableCountryInfo) Get() *CountryInfo {
	return v.value
}

func (v *NullableCountryInfo) Set(val *CountryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryInfo(val *CountryInfo) *NullableCountryInfo {
	return &NullableCountryInfo{value: val, isSet: true}
}

func (v NullableCountryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


