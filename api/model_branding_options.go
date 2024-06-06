/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the BrandingOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingOptions{}

// BrandingOptions The branding options.
type BrandingOptions struct {
	// Gets or sets the login disclaimer.
	LoginDisclaimer NullableString `json:"LoginDisclaimer,omitempty"`
	// Gets or sets the custom CSS.
	CustomCss NullableString `json:"CustomCss,omitempty"`
	// Gets or sets a value indicating whether to enable the splashscreen.
	SplashscreenEnabled *bool `json:"SplashscreenEnabled,omitempty"`
}

// NewBrandingOptions instantiates a new BrandingOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingOptions() *BrandingOptions {
	this := BrandingOptions{}
	return &this
}

// NewBrandingOptionsWithDefaults instantiates a new BrandingOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingOptionsWithDefaults() *BrandingOptions {
	this := BrandingOptions{}
	return &this
}

// GetLoginDisclaimer returns the LoginDisclaimer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrandingOptions) GetLoginDisclaimer() string {
	if o == nil || IsNil(o.LoginDisclaimer.Get()) {
		var ret string
		return ret
	}
	return *o.LoginDisclaimer.Get()
}

// GetLoginDisclaimerOk returns a tuple with the LoginDisclaimer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrandingOptions) GetLoginDisclaimerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoginDisclaimer.Get(), o.LoginDisclaimer.IsSet()
}

// HasLoginDisclaimer returns a boolean if a field has been set.
func (o *BrandingOptions) HasLoginDisclaimer() bool {
	if o != nil && o.LoginDisclaimer.IsSet() {
		return true
	}

	return false
}

// SetLoginDisclaimer gets a reference to the given NullableString and assigns it to the LoginDisclaimer field.
func (o *BrandingOptions) SetLoginDisclaimer(v string) {
	o.LoginDisclaimer.Set(&v)
}
// SetLoginDisclaimerNil sets the value for LoginDisclaimer to be an explicit nil
func (o *BrandingOptions) SetLoginDisclaimerNil() {
	o.LoginDisclaimer.Set(nil)
}

// UnsetLoginDisclaimer ensures that no value is present for LoginDisclaimer, not even an explicit nil
func (o *BrandingOptions) UnsetLoginDisclaimer() {
	o.LoginDisclaimer.Unset()
}

// GetCustomCss returns the CustomCss field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BrandingOptions) GetCustomCss() string {
	if o == nil || IsNil(o.CustomCss.Get()) {
		var ret string
		return ret
	}
	return *o.CustomCss.Get()
}

// GetCustomCssOk returns a tuple with the CustomCss field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BrandingOptions) GetCustomCssOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomCss.Get(), o.CustomCss.IsSet()
}

// HasCustomCss returns a boolean if a field has been set.
func (o *BrandingOptions) HasCustomCss() bool {
	if o != nil && o.CustomCss.IsSet() {
		return true
	}

	return false
}

// SetCustomCss gets a reference to the given NullableString and assigns it to the CustomCss field.
func (o *BrandingOptions) SetCustomCss(v string) {
	o.CustomCss.Set(&v)
}
// SetCustomCssNil sets the value for CustomCss to be an explicit nil
func (o *BrandingOptions) SetCustomCssNil() {
	o.CustomCss.Set(nil)
}

// UnsetCustomCss ensures that no value is present for CustomCss, not even an explicit nil
func (o *BrandingOptions) UnsetCustomCss() {
	o.CustomCss.Unset()
}

// GetSplashscreenEnabled returns the SplashscreenEnabled field value if set, zero value otherwise.
func (o *BrandingOptions) GetSplashscreenEnabled() bool {
	if o == nil || IsNil(o.SplashscreenEnabled) {
		var ret bool
		return ret
	}
	return *o.SplashscreenEnabled
}

// GetSplashscreenEnabledOk returns a tuple with the SplashscreenEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingOptions) GetSplashscreenEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SplashscreenEnabled) {
		return nil, false
	}
	return o.SplashscreenEnabled, true
}

// HasSplashscreenEnabled returns a boolean if a field has been set.
func (o *BrandingOptions) HasSplashscreenEnabled() bool {
	if o != nil && !IsNil(o.SplashscreenEnabled) {
		return true
	}

	return false
}

// SetSplashscreenEnabled gets a reference to the given bool and assigns it to the SplashscreenEnabled field.
func (o *BrandingOptions) SetSplashscreenEnabled(v bool) {
	o.SplashscreenEnabled = &v
}

func (o BrandingOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LoginDisclaimer.IsSet() {
		toSerialize["LoginDisclaimer"] = o.LoginDisclaimer.Get()
	}
	if o.CustomCss.IsSet() {
		toSerialize["CustomCss"] = o.CustomCss.Get()
	}
	if !IsNil(o.SplashscreenEnabled) {
		toSerialize["SplashscreenEnabled"] = o.SplashscreenEnabled
	}
	return toSerialize, nil
}

type NullableBrandingOptions struct {
	value *BrandingOptions
	isSet bool
}

func (v NullableBrandingOptions) Get() *BrandingOptions {
	return v.value
}

func (v *NullableBrandingOptions) Set(val *BrandingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingOptions(val *BrandingOptions) *NullableBrandingOptions {
	return &NullableBrandingOptions{value: val, isSet: true}
}

func (v NullableBrandingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


