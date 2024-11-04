/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the LogFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogFile{}

// LogFile struct for LogFile
type LogFile struct {
	// Gets or sets the date created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// Gets or sets the date modified.
	DateModified *time.Time `json:"DateModified,omitempty"`
	// Gets or sets the size.
	Size *int64 `json:"Size,omitempty"`
	// Gets or sets the name.
	Name *string `json:"Name,omitempty"`
}

// NewLogFile instantiates a new LogFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogFile() *LogFile {
	this := LogFile{}
	return &this
}

// NewLogFileWithDefaults instantiates a new LogFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogFileWithDefaults() *LogFile {
	this := LogFile{}
	return &this
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *LogFile) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogFile) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *LogFile) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *LogFile) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetDateModified returns the DateModified field value if set, zero value otherwise.
func (o *LogFile) GetDateModified() time.Time {
	if o == nil || IsNil(o.DateModified) {
		var ret time.Time
		return ret
	}
	return *o.DateModified
}

// GetDateModifiedOk returns a tuple with the DateModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogFile) GetDateModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateModified) {
		return nil, false
	}
	return o.DateModified, true
}

// HasDateModified returns a boolean if a field has been set.
func (o *LogFile) HasDateModified() bool {
	if o != nil && !IsNil(o.DateModified) {
		return true
	}

	return false
}

// SetDateModified gets a reference to the given time.Time and assigns it to the DateModified field.
func (o *LogFile) SetDateModified(v time.Time) {
	o.DateModified = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *LogFile) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogFile) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *LogFile) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *LogFile) SetSize(v int64) {
	o.Size = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogFile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogFile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogFile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogFile) SetName(v string) {
	o.Name = &v
}

func (o LogFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateCreated) {
		toSerialize["DateCreated"] = o.DateCreated
	}
	if !IsNil(o.DateModified) {
		toSerialize["DateModified"] = o.DateModified
	}
	if !IsNil(o.Size) {
		toSerialize["Size"] = o.Size
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableLogFile struct {
	value *LogFile
	isSet bool
}

func (v NullableLogFile) Get() *LogFile {
	return v.value
}

func (v *NullableLogFile) Set(val *LogFile) {
	v.value = val
	v.isSet = true
}

func (v NullableLogFile) IsSet() bool {
	return v.isSet
}

func (v *NullableLogFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogFile(val *LogFile) *NullableLogFile {
	return &NullableLogFile{value: val, isSet: true}
}

func (v NullableLogFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


