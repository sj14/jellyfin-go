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

// checks if the CultureDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CultureDto{}

// CultureDto Class CultureDto.
type CultureDto struct {
	// Gets the name.
	Name *string `json:"Name,omitempty"`
	// Gets the display name.
	DisplayName *string `json:"DisplayName,omitempty"`
	// Gets the name of the two letter ISO language.
	TwoLetterISOLanguageName *string `json:"TwoLetterISOLanguageName,omitempty"`
	// Gets the name of the three letter ISO language.
	ThreeLetterISOLanguageName NullableString `json:"ThreeLetterISOLanguageName,omitempty"`
	ThreeLetterISOLanguageNames []string `json:"ThreeLetterISOLanguageNames,omitempty"`
}

// NewCultureDto instantiates a new CultureDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCultureDto() *CultureDto {
	this := CultureDto{}
	return &this
}

// NewCultureDtoWithDefaults instantiates a new CultureDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCultureDtoWithDefaults() *CultureDto {
	this := CultureDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CultureDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CultureDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CultureDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CultureDto) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CultureDto) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CultureDto) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CultureDto) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CultureDto) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetTwoLetterISOLanguageName returns the TwoLetterISOLanguageName field value if set, zero value otherwise.
func (o *CultureDto) GetTwoLetterISOLanguageName() string {
	if o == nil || IsNil(o.TwoLetterISOLanguageName) {
		var ret string
		return ret
	}
	return *o.TwoLetterISOLanguageName
}

// GetTwoLetterISOLanguageNameOk returns a tuple with the TwoLetterISOLanguageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CultureDto) GetTwoLetterISOLanguageNameOk() (*string, bool) {
	if o == nil || IsNil(o.TwoLetterISOLanguageName) {
		return nil, false
	}
	return o.TwoLetterISOLanguageName, true
}

// HasTwoLetterISOLanguageName returns a boolean if a field has been set.
func (o *CultureDto) HasTwoLetterISOLanguageName() bool {
	if o != nil && !IsNil(o.TwoLetterISOLanguageName) {
		return true
	}

	return false
}

// SetTwoLetterISOLanguageName gets a reference to the given string and assigns it to the TwoLetterISOLanguageName field.
func (o *CultureDto) SetTwoLetterISOLanguageName(v string) {
	o.TwoLetterISOLanguageName = &v
}

// GetThreeLetterISOLanguageName returns the ThreeLetterISOLanguageName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CultureDto) GetThreeLetterISOLanguageName() string {
	if o == nil || IsNil(o.ThreeLetterISOLanguageName.Get()) {
		var ret string
		return ret
	}
	return *o.ThreeLetterISOLanguageName.Get()
}

// GetThreeLetterISOLanguageNameOk returns a tuple with the ThreeLetterISOLanguageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CultureDto) GetThreeLetterISOLanguageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreeLetterISOLanguageName.Get(), o.ThreeLetterISOLanguageName.IsSet()
}

// HasThreeLetterISOLanguageName returns a boolean if a field has been set.
func (o *CultureDto) HasThreeLetterISOLanguageName() bool {
	if o != nil && o.ThreeLetterISOLanguageName.IsSet() {
		return true
	}

	return false
}

// SetThreeLetterISOLanguageName gets a reference to the given NullableString and assigns it to the ThreeLetterISOLanguageName field.
func (o *CultureDto) SetThreeLetterISOLanguageName(v string) {
	o.ThreeLetterISOLanguageName.Set(&v)
}
// SetThreeLetterISOLanguageNameNil sets the value for ThreeLetterISOLanguageName to be an explicit nil
func (o *CultureDto) SetThreeLetterISOLanguageNameNil() {
	o.ThreeLetterISOLanguageName.Set(nil)
}

// UnsetThreeLetterISOLanguageName ensures that no value is present for ThreeLetterISOLanguageName, not even an explicit nil
func (o *CultureDto) UnsetThreeLetterISOLanguageName() {
	o.ThreeLetterISOLanguageName.Unset()
}

// GetThreeLetterISOLanguageNames returns the ThreeLetterISOLanguageNames field value if set, zero value otherwise.
func (o *CultureDto) GetThreeLetterISOLanguageNames() []string {
	if o == nil || IsNil(o.ThreeLetterISOLanguageNames) {
		var ret []string
		return ret
	}
	return o.ThreeLetterISOLanguageNames
}

// GetThreeLetterISOLanguageNamesOk returns a tuple with the ThreeLetterISOLanguageNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CultureDto) GetThreeLetterISOLanguageNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ThreeLetterISOLanguageNames) {
		return nil, false
	}
	return o.ThreeLetterISOLanguageNames, true
}

// HasThreeLetterISOLanguageNames returns a boolean if a field has been set.
func (o *CultureDto) HasThreeLetterISOLanguageNames() bool {
	if o != nil && !IsNil(o.ThreeLetterISOLanguageNames) {
		return true
	}

	return false
}

// SetThreeLetterISOLanguageNames gets a reference to the given []string and assigns it to the ThreeLetterISOLanguageNames field.
func (o *CultureDto) SetThreeLetterISOLanguageNames(v []string) {
	o.ThreeLetterISOLanguageNames = v
}

func (o CultureDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CultureDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.DisplayName) {
		toSerialize["DisplayName"] = o.DisplayName
	}
	if !IsNil(o.TwoLetterISOLanguageName) {
		toSerialize["TwoLetterISOLanguageName"] = o.TwoLetterISOLanguageName
	}
	if o.ThreeLetterISOLanguageName.IsSet() {
		toSerialize["ThreeLetterISOLanguageName"] = o.ThreeLetterISOLanguageName.Get()
	}
	if !IsNil(o.ThreeLetterISOLanguageNames) {
		toSerialize["ThreeLetterISOLanguageNames"] = o.ThreeLetterISOLanguageNames
	}
	return toSerialize, nil
}

type NullableCultureDto struct {
	value *CultureDto
	isSet bool
}

func (v NullableCultureDto) Get() *CultureDto {
	return v.value
}

func (v *NullableCultureDto) Set(val *CultureDto) {
	v.value = val
	v.isSet = true
}

func (v NullableCultureDto) IsSet() bool {
	return v.isSet
}

func (v *NullableCultureDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCultureDto(val *CultureDto) *NullableCultureDto {
	return &NullableCultureDto{value: val, isSet: true}
}

func (v NullableCultureDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCultureDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


