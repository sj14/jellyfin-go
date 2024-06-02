/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AllThemeMediaResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllThemeMediaResult{}

// AllThemeMediaResult struct for AllThemeMediaResult
type AllThemeMediaResult struct {
	ThemeVideosResult NullableAllThemeMediaResultThemeVideosResult `json:"ThemeVideosResult,omitempty"`
	ThemeSongsResult NullableAllThemeMediaResultThemeVideosResult `json:"ThemeSongsResult,omitempty"`
	SoundtrackSongsResult NullableAllThemeMediaResultThemeVideosResult `json:"SoundtrackSongsResult,omitempty"`
}

// NewAllThemeMediaResult instantiates a new AllThemeMediaResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllThemeMediaResult() *AllThemeMediaResult {
	this := AllThemeMediaResult{}
	return &this
}

// NewAllThemeMediaResultWithDefaults instantiates a new AllThemeMediaResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllThemeMediaResultWithDefaults() *AllThemeMediaResult {
	this := AllThemeMediaResult{}
	return &this
}

// GetThemeVideosResult returns the ThemeVideosResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllThemeMediaResult) GetThemeVideosResult() AllThemeMediaResultThemeVideosResult {
	if o == nil || IsNil(o.ThemeVideosResult.Get()) {
		var ret AllThemeMediaResultThemeVideosResult
		return ret
	}
	return *o.ThemeVideosResult.Get()
}

// GetThemeVideosResultOk returns a tuple with the ThemeVideosResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AllThemeMediaResult) GetThemeVideosResultOk() (*AllThemeMediaResultThemeVideosResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThemeVideosResult.Get(), o.ThemeVideosResult.IsSet()
}

// HasThemeVideosResult returns a boolean if a field has been set.
func (o *AllThemeMediaResult) HasThemeVideosResult() bool {
	if o != nil && o.ThemeVideosResult.IsSet() {
		return true
	}

	return false
}

// SetThemeVideosResult gets a reference to the given NullableAllThemeMediaResultThemeVideosResult and assigns it to the ThemeVideosResult field.
func (o *AllThemeMediaResult) SetThemeVideosResult(v AllThemeMediaResultThemeVideosResult) {
	o.ThemeVideosResult.Set(&v)
}
// SetThemeVideosResultNil sets the value for ThemeVideosResult to be an explicit nil
func (o *AllThemeMediaResult) SetThemeVideosResultNil() {
	o.ThemeVideosResult.Set(nil)
}

// UnsetThemeVideosResult ensures that no value is present for ThemeVideosResult, not even an explicit nil
func (o *AllThemeMediaResult) UnsetThemeVideosResult() {
	o.ThemeVideosResult.Unset()
}

// GetThemeSongsResult returns the ThemeSongsResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllThemeMediaResult) GetThemeSongsResult() AllThemeMediaResultThemeVideosResult {
	if o == nil || IsNil(o.ThemeSongsResult.Get()) {
		var ret AllThemeMediaResultThemeVideosResult
		return ret
	}
	return *o.ThemeSongsResult.Get()
}

// GetThemeSongsResultOk returns a tuple with the ThemeSongsResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AllThemeMediaResult) GetThemeSongsResultOk() (*AllThemeMediaResultThemeVideosResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThemeSongsResult.Get(), o.ThemeSongsResult.IsSet()
}

// HasThemeSongsResult returns a boolean if a field has been set.
func (o *AllThemeMediaResult) HasThemeSongsResult() bool {
	if o != nil && o.ThemeSongsResult.IsSet() {
		return true
	}

	return false
}

// SetThemeSongsResult gets a reference to the given NullableAllThemeMediaResultThemeVideosResult and assigns it to the ThemeSongsResult field.
func (o *AllThemeMediaResult) SetThemeSongsResult(v AllThemeMediaResultThemeVideosResult) {
	o.ThemeSongsResult.Set(&v)
}
// SetThemeSongsResultNil sets the value for ThemeSongsResult to be an explicit nil
func (o *AllThemeMediaResult) SetThemeSongsResultNil() {
	o.ThemeSongsResult.Set(nil)
}

// UnsetThemeSongsResult ensures that no value is present for ThemeSongsResult, not even an explicit nil
func (o *AllThemeMediaResult) UnsetThemeSongsResult() {
	o.ThemeSongsResult.Unset()
}

// GetSoundtrackSongsResult returns the SoundtrackSongsResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AllThemeMediaResult) GetSoundtrackSongsResult() AllThemeMediaResultThemeVideosResult {
	if o == nil || IsNil(o.SoundtrackSongsResult.Get()) {
		var ret AllThemeMediaResultThemeVideosResult
		return ret
	}
	return *o.SoundtrackSongsResult.Get()
}

// GetSoundtrackSongsResultOk returns a tuple with the SoundtrackSongsResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AllThemeMediaResult) GetSoundtrackSongsResultOk() (*AllThemeMediaResultThemeVideosResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.SoundtrackSongsResult.Get(), o.SoundtrackSongsResult.IsSet()
}

// HasSoundtrackSongsResult returns a boolean if a field has been set.
func (o *AllThemeMediaResult) HasSoundtrackSongsResult() bool {
	if o != nil && o.SoundtrackSongsResult.IsSet() {
		return true
	}

	return false
}

// SetSoundtrackSongsResult gets a reference to the given NullableAllThemeMediaResultThemeVideosResult and assigns it to the SoundtrackSongsResult field.
func (o *AllThemeMediaResult) SetSoundtrackSongsResult(v AllThemeMediaResultThemeVideosResult) {
	o.SoundtrackSongsResult.Set(&v)
}
// SetSoundtrackSongsResultNil sets the value for SoundtrackSongsResult to be an explicit nil
func (o *AllThemeMediaResult) SetSoundtrackSongsResultNil() {
	o.SoundtrackSongsResult.Set(nil)
}

// UnsetSoundtrackSongsResult ensures that no value is present for SoundtrackSongsResult, not even an explicit nil
func (o *AllThemeMediaResult) UnsetSoundtrackSongsResult() {
	o.SoundtrackSongsResult.Unset()
}

func (o AllThemeMediaResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllThemeMediaResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ThemeVideosResult.IsSet() {
		toSerialize["ThemeVideosResult"] = o.ThemeVideosResult.Get()
	}
	if o.ThemeSongsResult.IsSet() {
		toSerialize["ThemeSongsResult"] = o.ThemeSongsResult.Get()
	}
	if o.SoundtrackSongsResult.IsSet() {
		toSerialize["SoundtrackSongsResult"] = o.SoundtrackSongsResult.Get()
	}
	return toSerialize, nil
}

type NullableAllThemeMediaResult struct {
	value *AllThemeMediaResult
	isSet bool
}

func (v NullableAllThemeMediaResult) Get() *AllThemeMediaResult {
	return v.value
}

func (v *NullableAllThemeMediaResult) Set(val *AllThemeMediaResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAllThemeMediaResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAllThemeMediaResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllThemeMediaResult(val *AllThemeMediaResult) *NullableAllThemeMediaResult {
	return &NullableAllThemeMediaResult{value: val, isSet: true}
}

func (v NullableAllThemeMediaResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllThemeMediaResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


