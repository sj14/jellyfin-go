/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ExternalIdInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalIdInfo{}

// ExternalIdInfo Represents the external id information for serialization to the client.
type ExternalIdInfo struct {
	// Gets or sets the display name of the external id provider (IE: IMDB, MusicBrainz, etc).
	Name *string `json:"Name,omitempty"`
	// Gets or sets the unique key for this id. This key should be unique across all providers.
	Key *string `json:"Key,omitempty"`
	// Gets or sets the specific media type for this id. This is used to distinguish between the different  external id types for providers with multiple ids.  A null value indicates there is no specific media type associated with the external id, or this is the  default id for the external provider so there is no need to specify a type.
	Type NullableExternalIdMediaType `json:"Type,omitempty"`
	// Gets or sets the URL format string.
	UrlFormatString NullableString `json:"UrlFormatString,omitempty"`
}

// NewExternalIdInfo instantiates a new ExternalIdInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalIdInfo() *ExternalIdInfo {
	this := ExternalIdInfo{}
	return &this
}

// NewExternalIdInfoWithDefaults instantiates a new ExternalIdInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalIdInfoWithDefaults() *ExternalIdInfo {
	this := ExternalIdInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExternalIdInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalIdInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExternalIdInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExternalIdInfo) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ExternalIdInfo) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalIdInfo) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ExternalIdInfo) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ExternalIdInfo) SetKey(v string) {
	o.Key = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalIdInfo) GetType() ExternalIdMediaType {
	if o == nil || IsNil(o.Type.Get()) {
		var ret ExternalIdMediaType
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalIdInfo) GetTypeOk() (*ExternalIdMediaType, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ExternalIdInfo) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableExternalIdMediaType and assigns it to the Type field.
func (o *ExternalIdInfo) SetType(v ExternalIdMediaType) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ExternalIdInfo) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ExternalIdInfo) UnsetType() {
	o.Type.Unset()
}

// GetUrlFormatString returns the UrlFormatString field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalIdInfo) GetUrlFormatString() string {
	if o == nil || IsNil(o.UrlFormatString.Get()) {
		var ret string
		return ret
	}
	return *o.UrlFormatString.Get()
}

// GetUrlFormatStringOk returns a tuple with the UrlFormatString field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalIdInfo) GetUrlFormatStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UrlFormatString.Get(), o.UrlFormatString.IsSet()
}

// HasUrlFormatString returns a boolean if a field has been set.
func (o *ExternalIdInfo) HasUrlFormatString() bool {
	if o != nil && o.UrlFormatString.IsSet() {
		return true
	}

	return false
}

// SetUrlFormatString gets a reference to the given NullableString and assigns it to the UrlFormatString field.
func (o *ExternalIdInfo) SetUrlFormatString(v string) {
	o.UrlFormatString.Set(&v)
}
// SetUrlFormatStringNil sets the value for UrlFormatString to be an explicit nil
func (o *ExternalIdInfo) SetUrlFormatStringNil() {
	o.UrlFormatString.Set(nil)
}

// UnsetUrlFormatString ensures that no value is present for UrlFormatString, not even an explicit nil
func (o *ExternalIdInfo) UnsetUrlFormatString() {
	o.UrlFormatString.Unset()
}

func (o ExternalIdInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalIdInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Key) {
		toSerialize["Key"] = o.Key
	}
	if o.Type.IsSet() {
		toSerialize["Type"] = o.Type.Get()
	}
	if o.UrlFormatString.IsSet() {
		toSerialize["UrlFormatString"] = o.UrlFormatString.Get()
	}
	return toSerialize, nil
}

type NullableExternalIdInfo struct {
	value *ExternalIdInfo
	isSet bool
}

func (v NullableExternalIdInfo) Get() *ExternalIdInfo {
	return v.value
}

func (v *NullableExternalIdInfo) Set(val *ExternalIdInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalIdInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalIdInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalIdInfo(val *ExternalIdInfo) *NullableExternalIdInfo {
	return &NullableExternalIdInfo{value: val, isSet: true}
}

func (v NullableExternalIdInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalIdInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


