/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MetadataEditorInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataEditorInfo{}

// MetadataEditorInfo struct for MetadataEditorInfo
type MetadataEditorInfo struct {
	ParentalRatingOptions []ParentalRating `json:"ParentalRatingOptions,omitempty"`
	Countries []CountryInfo `json:"Countries,omitempty"`
	Cultures []CultureDto `json:"Cultures,omitempty"`
	ExternalIdInfos []ExternalIdInfo `json:"ExternalIdInfos,omitempty"`
	ContentType NullableCollectionType `json:"ContentType,omitempty"`
	ContentTypeOptions []NameValuePair `json:"ContentTypeOptions,omitempty"`
}

// NewMetadataEditorInfo instantiates a new MetadataEditorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataEditorInfo() *MetadataEditorInfo {
	this := MetadataEditorInfo{}
	return &this
}

// NewMetadataEditorInfoWithDefaults instantiates a new MetadataEditorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataEditorInfoWithDefaults() *MetadataEditorInfo {
	this := MetadataEditorInfo{}
	return &this
}

// GetParentalRatingOptions returns the ParentalRatingOptions field value if set, zero value otherwise.
func (o *MetadataEditorInfo) GetParentalRatingOptions() []ParentalRating {
	if o == nil || IsNil(o.ParentalRatingOptions) {
		var ret []ParentalRating
		return ret
	}
	return o.ParentalRatingOptions
}

// GetParentalRatingOptionsOk returns a tuple with the ParentalRatingOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataEditorInfo) GetParentalRatingOptionsOk() ([]ParentalRating, bool) {
	if o == nil || IsNil(o.ParentalRatingOptions) {
		return nil, false
	}
	return o.ParentalRatingOptions, true
}

// HasParentalRatingOptions returns a boolean if a field has been set.
func (o *MetadataEditorInfo) HasParentalRatingOptions() bool {
	if o != nil && !IsNil(o.ParentalRatingOptions) {
		return true
	}

	return false
}

// SetParentalRatingOptions gets a reference to the given []ParentalRating and assigns it to the ParentalRatingOptions field.
func (o *MetadataEditorInfo) SetParentalRatingOptions(v []ParentalRating) {
	o.ParentalRatingOptions = v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *MetadataEditorInfo) GetCountries() []CountryInfo {
	if o == nil || IsNil(o.Countries) {
		var ret []CountryInfo
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataEditorInfo) GetCountriesOk() ([]CountryInfo, bool) {
	if o == nil || IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *MetadataEditorInfo) HasCountries() bool {
	if o != nil && !IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []CountryInfo and assigns it to the Countries field.
func (o *MetadataEditorInfo) SetCountries(v []CountryInfo) {
	o.Countries = v
}

// GetCultures returns the Cultures field value if set, zero value otherwise.
func (o *MetadataEditorInfo) GetCultures() []CultureDto {
	if o == nil || IsNil(o.Cultures) {
		var ret []CultureDto
		return ret
	}
	return o.Cultures
}

// GetCulturesOk returns a tuple with the Cultures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataEditorInfo) GetCulturesOk() ([]CultureDto, bool) {
	if o == nil || IsNil(o.Cultures) {
		return nil, false
	}
	return o.Cultures, true
}

// HasCultures returns a boolean if a field has been set.
func (o *MetadataEditorInfo) HasCultures() bool {
	if o != nil && !IsNil(o.Cultures) {
		return true
	}

	return false
}

// SetCultures gets a reference to the given []CultureDto and assigns it to the Cultures field.
func (o *MetadataEditorInfo) SetCultures(v []CultureDto) {
	o.Cultures = v
}

// GetExternalIdInfos returns the ExternalIdInfos field value if set, zero value otherwise.
func (o *MetadataEditorInfo) GetExternalIdInfos() []ExternalIdInfo {
	if o == nil || IsNil(o.ExternalIdInfos) {
		var ret []ExternalIdInfo
		return ret
	}
	return o.ExternalIdInfos
}

// GetExternalIdInfosOk returns a tuple with the ExternalIdInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataEditorInfo) GetExternalIdInfosOk() ([]ExternalIdInfo, bool) {
	if o == nil || IsNil(o.ExternalIdInfos) {
		return nil, false
	}
	return o.ExternalIdInfos, true
}

// HasExternalIdInfos returns a boolean if a field has been set.
func (o *MetadataEditorInfo) HasExternalIdInfos() bool {
	if o != nil && !IsNil(o.ExternalIdInfos) {
		return true
	}

	return false
}

// SetExternalIdInfos gets a reference to the given []ExternalIdInfo and assigns it to the ExternalIdInfos field.
func (o *MetadataEditorInfo) SetExternalIdInfos(v []ExternalIdInfo) {
	o.ExternalIdInfos = v
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataEditorInfo) GetContentType() CollectionType {
	if o == nil || IsNil(o.ContentType.Get()) {
		var ret CollectionType
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataEditorInfo) GetContentTypeOk() (*CollectionType, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *MetadataEditorInfo) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableCollectionType and assigns it to the ContentType field.
func (o *MetadataEditorInfo) SetContentType(v CollectionType) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *MetadataEditorInfo) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *MetadataEditorInfo) UnsetContentType() {
	o.ContentType.Unset()
}

// GetContentTypeOptions returns the ContentTypeOptions field value if set, zero value otherwise.
func (o *MetadataEditorInfo) GetContentTypeOptions() []NameValuePair {
	if o == nil || IsNil(o.ContentTypeOptions) {
		var ret []NameValuePair
		return ret
	}
	return o.ContentTypeOptions
}

// GetContentTypeOptionsOk returns a tuple with the ContentTypeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataEditorInfo) GetContentTypeOptionsOk() ([]NameValuePair, bool) {
	if o == nil || IsNil(o.ContentTypeOptions) {
		return nil, false
	}
	return o.ContentTypeOptions, true
}

// HasContentTypeOptions returns a boolean if a field has been set.
func (o *MetadataEditorInfo) HasContentTypeOptions() bool {
	if o != nil && !IsNil(o.ContentTypeOptions) {
		return true
	}

	return false
}

// SetContentTypeOptions gets a reference to the given []NameValuePair and assigns it to the ContentTypeOptions field.
func (o *MetadataEditorInfo) SetContentTypeOptions(v []NameValuePair) {
	o.ContentTypeOptions = v
}

func (o MetadataEditorInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataEditorInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParentalRatingOptions) {
		toSerialize["ParentalRatingOptions"] = o.ParentalRatingOptions
	}
	if !IsNil(o.Countries) {
		toSerialize["Countries"] = o.Countries
	}
	if !IsNil(o.Cultures) {
		toSerialize["Cultures"] = o.Cultures
	}
	if !IsNil(o.ExternalIdInfos) {
		toSerialize["ExternalIdInfos"] = o.ExternalIdInfos
	}
	if o.ContentType.IsSet() {
		toSerialize["ContentType"] = o.ContentType.Get()
	}
	if !IsNil(o.ContentTypeOptions) {
		toSerialize["ContentTypeOptions"] = o.ContentTypeOptions
	}
	return toSerialize, nil
}

type NullableMetadataEditorInfo struct {
	value *MetadataEditorInfo
	isSet bool
}

func (v NullableMetadataEditorInfo) Get() *MetadataEditorInfo {
	return v.value
}

func (v *NullableMetadataEditorInfo) Set(val *MetadataEditorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataEditorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataEditorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataEditorInfo(val *MetadataEditorInfo) *NullableMetadataEditorInfo {
	return &NullableMetadataEditorInfo{value: val, isSet: true}
}

func (v NullableMetadataEditorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataEditorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


