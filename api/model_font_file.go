/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the FontFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FontFile{}

// FontFile Class FontFile.
type FontFile struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the size.
	Size *int64 `json:"Size,omitempty"`
	// Gets or sets the date created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// Gets or sets the date modified.
	DateModified *time.Time `json:"DateModified,omitempty"`
}

// NewFontFile instantiates a new FontFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFontFile() *FontFile {
	this := FontFile{}
	return &this
}

// NewFontFileWithDefaults instantiates a new FontFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFontFileWithDefaults() *FontFile {
	this := FontFile{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FontFile) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FontFile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *FontFile) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *FontFile) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *FontFile) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *FontFile) UnsetName() {
	o.Name.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FontFile) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FontFile) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FontFile) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *FontFile) SetSize(v int64) {
	o.Size = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *FontFile) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FontFile) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *FontFile) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *FontFile) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetDateModified returns the DateModified field value if set, zero value otherwise.
func (o *FontFile) GetDateModified() time.Time {
	if o == nil || IsNil(o.DateModified) {
		var ret time.Time
		return ret
	}
	return *o.DateModified
}

// GetDateModifiedOk returns a tuple with the DateModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FontFile) GetDateModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateModified) {
		return nil, false
	}
	return o.DateModified, true
}

// HasDateModified returns a boolean if a field has been set.
func (o *FontFile) HasDateModified() bool {
	if o != nil && !IsNil(o.DateModified) {
		return true
	}

	return false
}

// SetDateModified gets a reference to the given time.Time and assigns it to the DateModified field.
func (o *FontFile) SetDateModified(v time.Time) {
	o.DateModified = &v
}

func (o FontFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FontFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	if !IsNil(o.DateCreated) {
		toSerialize["DateCreated"] = o.DateCreated
	}
	if !IsNil(o.DateModified) {
		toSerialize["DateModified"] = o.DateModified
	}
	return toSerialize, nil
}

type NullableFontFile struct {
	value *FontFile
	isSet bool
}

func (v NullableFontFile) Get() *FontFile {
	return v.value
}

func (v *NullableFontFile) Set(val *FontFile) {
	v.value = val
	v.isSet = true
}

func (v NullableFontFile) IsSet() bool {
	return v.isSet
}

func (v *NullableFontFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFontFile(val *FontFile) *NullableFontFile {
	return &NullableFontFile{value: val, isSet: true}
}

func (v NullableFontFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFontFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


