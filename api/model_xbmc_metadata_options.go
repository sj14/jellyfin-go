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

// checks if the XbmcMetadataOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XbmcMetadataOptions{}

// XbmcMetadataOptions struct for XbmcMetadataOptions
type XbmcMetadataOptions struct {
	UserId NullableString `json:"UserId,omitempty"`
	ReleaseDateFormat *string `json:"ReleaseDateFormat,omitempty"`
	SaveImagePathsInNfo *bool `json:"SaveImagePathsInNfo,omitempty"`
	EnablePathSubstitution *bool `json:"EnablePathSubstitution,omitempty"`
	EnableExtraThumbsDuplication *bool `json:"EnableExtraThumbsDuplication,omitempty"`
}

// NewXbmcMetadataOptions instantiates a new XbmcMetadataOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXbmcMetadataOptions() *XbmcMetadataOptions {
	this := XbmcMetadataOptions{}
	return &this
}

// NewXbmcMetadataOptionsWithDefaults instantiates a new XbmcMetadataOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXbmcMetadataOptionsWithDefaults() *XbmcMetadataOptions {
	this := XbmcMetadataOptions{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *XbmcMetadataOptions) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *XbmcMetadataOptions) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *XbmcMetadataOptions) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *XbmcMetadataOptions) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *XbmcMetadataOptions) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *XbmcMetadataOptions) UnsetUserId() {
	o.UserId.Unset()
}

// GetReleaseDateFormat returns the ReleaseDateFormat field value if set, zero value otherwise.
func (o *XbmcMetadataOptions) GetReleaseDateFormat() string {
	if o == nil || IsNil(o.ReleaseDateFormat) {
		var ret string
		return ret
	}
	return *o.ReleaseDateFormat
}

// GetReleaseDateFormatOk returns a tuple with the ReleaseDateFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XbmcMetadataOptions) GetReleaseDateFormatOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseDateFormat) {
		return nil, false
	}
	return o.ReleaseDateFormat, true
}

// HasReleaseDateFormat returns a boolean if a field has been set.
func (o *XbmcMetadataOptions) HasReleaseDateFormat() bool {
	if o != nil && !IsNil(o.ReleaseDateFormat) {
		return true
	}

	return false
}

// SetReleaseDateFormat gets a reference to the given string and assigns it to the ReleaseDateFormat field.
func (o *XbmcMetadataOptions) SetReleaseDateFormat(v string) {
	o.ReleaseDateFormat = &v
}

// GetSaveImagePathsInNfo returns the SaveImagePathsInNfo field value if set, zero value otherwise.
func (o *XbmcMetadataOptions) GetSaveImagePathsInNfo() bool {
	if o == nil || IsNil(o.SaveImagePathsInNfo) {
		var ret bool
		return ret
	}
	return *o.SaveImagePathsInNfo
}

// GetSaveImagePathsInNfoOk returns a tuple with the SaveImagePathsInNfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XbmcMetadataOptions) GetSaveImagePathsInNfoOk() (*bool, bool) {
	if o == nil || IsNil(o.SaveImagePathsInNfo) {
		return nil, false
	}
	return o.SaveImagePathsInNfo, true
}

// HasSaveImagePathsInNfo returns a boolean if a field has been set.
func (o *XbmcMetadataOptions) HasSaveImagePathsInNfo() bool {
	if o != nil && !IsNil(o.SaveImagePathsInNfo) {
		return true
	}

	return false
}

// SetSaveImagePathsInNfo gets a reference to the given bool and assigns it to the SaveImagePathsInNfo field.
func (o *XbmcMetadataOptions) SetSaveImagePathsInNfo(v bool) {
	o.SaveImagePathsInNfo = &v
}

// GetEnablePathSubstitution returns the EnablePathSubstitution field value if set, zero value otherwise.
func (o *XbmcMetadataOptions) GetEnablePathSubstitution() bool {
	if o == nil || IsNil(o.EnablePathSubstitution) {
		var ret bool
		return ret
	}
	return *o.EnablePathSubstitution
}

// GetEnablePathSubstitutionOk returns a tuple with the EnablePathSubstitution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XbmcMetadataOptions) GetEnablePathSubstitutionOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePathSubstitution) {
		return nil, false
	}
	return o.EnablePathSubstitution, true
}

// HasEnablePathSubstitution returns a boolean if a field has been set.
func (o *XbmcMetadataOptions) HasEnablePathSubstitution() bool {
	if o != nil && !IsNil(o.EnablePathSubstitution) {
		return true
	}

	return false
}

// SetEnablePathSubstitution gets a reference to the given bool and assigns it to the EnablePathSubstitution field.
func (o *XbmcMetadataOptions) SetEnablePathSubstitution(v bool) {
	o.EnablePathSubstitution = &v
}

// GetEnableExtraThumbsDuplication returns the EnableExtraThumbsDuplication field value if set, zero value otherwise.
func (o *XbmcMetadataOptions) GetEnableExtraThumbsDuplication() bool {
	if o == nil || IsNil(o.EnableExtraThumbsDuplication) {
		var ret bool
		return ret
	}
	return *o.EnableExtraThumbsDuplication
}

// GetEnableExtraThumbsDuplicationOk returns a tuple with the EnableExtraThumbsDuplication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XbmcMetadataOptions) GetEnableExtraThumbsDuplicationOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableExtraThumbsDuplication) {
		return nil, false
	}
	return o.EnableExtraThumbsDuplication, true
}

// HasEnableExtraThumbsDuplication returns a boolean if a field has been set.
func (o *XbmcMetadataOptions) HasEnableExtraThumbsDuplication() bool {
	if o != nil && !IsNil(o.EnableExtraThumbsDuplication) {
		return true
	}

	return false
}

// SetEnableExtraThumbsDuplication gets a reference to the given bool and assigns it to the EnableExtraThumbsDuplication field.
func (o *XbmcMetadataOptions) SetEnableExtraThumbsDuplication(v bool) {
	o.EnableExtraThumbsDuplication = &v
}

func (o XbmcMetadataOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XbmcMetadataOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.UserId.IsSet() {
		toSerialize["UserId"] = o.UserId.Get()
	}
	if !IsNil(o.ReleaseDateFormat) {
		toSerialize["ReleaseDateFormat"] = o.ReleaseDateFormat
	}
	if !IsNil(o.SaveImagePathsInNfo) {
		toSerialize["SaveImagePathsInNfo"] = o.SaveImagePathsInNfo
	}
	if !IsNil(o.EnablePathSubstitution) {
		toSerialize["EnablePathSubstitution"] = o.EnablePathSubstitution
	}
	if !IsNil(o.EnableExtraThumbsDuplication) {
		toSerialize["EnableExtraThumbsDuplication"] = o.EnableExtraThumbsDuplication
	}
	return toSerialize, nil
}

type NullableXbmcMetadataOptions struct {
	value *XbmcMetadataOptions
	isSet bool
}

func (v NullableXbmcMetadataOptions) Get() *XbmcMetadataOptions {
	return v.value
}

func (v *NullableXbmcMetadataOptions) Set(val *XbmcMetadataOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableXbmcMetadataOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableXbmcMetadataOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXbmcMetadataOptions(val *XbmcMetadataOptions) *NullableXbmcMetadataOptions {
	return &NullableXbmcMetadataOptions{value: val, isSet: true}
}

func (v NullableXbmcMetadataOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXbmcMetadataOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


