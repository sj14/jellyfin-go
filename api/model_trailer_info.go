/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the TrailerInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrailerInfo{}

// TrailerInfo struct for TrailerInfo
type TrailerInfo struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the original title.
	OriginalTitle NullableString `json:"OriginalTitle,omitempty"`
	// Gets or sets the path.
	Path NullableString `json:"Path,omitempty"`
	// Gets or sets the metadata language.
	MetadataLanguage NullableString `json:"MetadataLanguage,omitempty"`
	// Gets or sets the metadata country code.
	MetadataCountryCode NullableString `json:"MetadataCountryCode,omitempty"`
	// Gets or sets the provider ids.
	ProviderIds map[string]string `json:"ProviderIds,omitempty"`
	// Gets or sets the year.
	Year NullableInt32 `json:"Year,omitempty"`
	IndexNumber NullableInt32 `json:"IndexNumber,omitempty"`
	ParentIndexNumber NullableInt32 `json:"ParentIndexNumber,omitempty"`
	PremiereDate NullableTime `json:"PremiereDate,omitempty"`
	IsAutomated *bool `json:"IsAutomated,omitempty"`
}

// NewTrailerInfo instantiates a new TrailerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrailerInfo() *TrailerInfo {
	this := TrailerInfo{}
	return &this
}

// NewTrailerInfoWithDefaults instantiates a new TrailerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrailerInfoWithDefaults() *TrailerInfo {
	this := TrailerInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *TrailerInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *TrailerInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *TrailerInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *TrailerInfo) UnsetName() {
	o.Name.Unset()
}

// GetOriginalTitle returns the OriginalTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerInfo) GetOriginalTitle() string {
	if o == nil || IsNil(o.OriginalTitle.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalTitle.Get()
}

// GetOriginalTitleOk returns a tuple with the OriginalTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerInfo) GetOriginalTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalTitle.Get(), o.OriginalTitle.IsSet()
}

// HasOriginalTitle returns a boolean if a field has been set.
func (o *TrailerInfo) HasOriginalTitle() bool {
	if o != nil && o.OriginalTitle.IsSet() {
		return true
	}

	return false
}

// SetOriginalTitle gets a reference to the given NullableString and assigns it to the OriginalTitle field.
func (o *TrailerInfo) SetOriginalTitle(v string) {
	o.OriginalTitle.Set(&v)
}
// SetOriginalTitleNil sets the value for OriginalTitle to be an explicit nil
func (o *TrailerInfo) SetOriginalTitleNil() {
	o.OriginalTitle.Set(nil)
}

// UnsetOriginalTitle ensures that no value is present for OriginalTitle, not even an explicit nil
func (o *TrailerInfo) UnsetOriginalTitle() {
	o.OriginalTitle.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerInfo) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerInfo) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *TrailerInfo) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *TrailerInfo) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *TrailerInfo) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *TrailerInfo) UnsetPath() {
	o.Path.Unset()
}

// GetMetadataLanguage returns the MetadataLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerInfo) GetMetadataLanguage() string {
	if o == nil || IsNil(o.MetadataLanguage.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataLanguage.Get()
}

// GetMetadataLanguageOk returns a tuple with the MetadataLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerInfo) GetMetadataLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataLanguage.Get(), o.MetadataLanguage.IsSet()
}

// HasMetadataLanguage returns a boolean if a field has been set.
func (o *TrailerInfo) HasMetadataLanguage() bool {
	if o != nil && o.MetadataLanguage.IsSet() {
		return true
	}

	return false
}

// SetMetadataLanguage gets a reference to the given NullableString and assigns it to the MetadataLanguage field.
func (o *TrailerInfo) SetMetadataLanguage(v string) {
	o.MetadataLanguage.Set(&v)
}
// SetMetadataLanguageNil sets the value for MetadataLanguage to be an explicit nil
func (o *TrailerInfo) SetMetadataLanguageNil() {
	o.MetadataLanguage.Set(nil)
}

// UnsetMetadataLanguage ensures that no value is present for MetadataLanguage, not even an explicit nil
func (o *TrailerInfo) UnsetMetadataLanguage() {
	o.MetadataLanguage.Unset()
}

// GetMetadataCountryCode returns the MetadataCountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerInfo) GetMetadataCountryCode() string {
	if o == nil || IsNil(o.MetadataCountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataCountryCode.Get()
}

// GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerInfo) GetMetadataCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataCountryCode.Get(), o.MetadataCountryCode.IsSet()
}

// HasMetadataCountryCode returns a boolean if a field has been set.
func (o *TrailerInfo) HasMetadataCountryCode() bool {
	if o != nil && o.MetadataCountryCode.IsSet() {
		return true
	}

	return false
}

// SetMetadataCountryCode gets a reference to the given NullableString and assigns it to the MetadataCountryCode field.
func (o *TrailerInfo) SetMetadataCountryCode(v string) {
	o.MetadataCountryCode.Set(&v)
}
// SetMetadataCountryCodeNil sets the value for MetadataCountryCode to be an explicit nil
func (o *TrailerInfo) SetMetadataCountryCodeNil() {
	o.MetadataCountryCode.Set(nil)
}

// UnsetMetadataCountryCode ensures that no value is present for MetadataCountryCode, not even an explicit nil
func (o *TrailerInfo) UnsetMetadataCountryCode() {
	o.MetadataCountryCode.Unset()
}

// GetProviderIds returns the ProviderIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerInfo) GetProviderIds() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.ProviderIds
}

// GetProviderIdsOk returns a tuple with the ProviderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerInfo) GetProviderIdsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ProviderIds) {
		return nil, false
	}
	return &o.ProviderIds, true
}

// HasProviderIds returns a boolean if a field has been set.
func (o *TrailerInfo) HasProviderIds() bool {
	if o != nil && !IsNil(o.ProviderIds) {
		return true
	}

	return false
}

// SetProviderIds gets a reference to the given map[string]string and assigns it to the ProviderIds field.
func (o *TrailerInfo) SetProviderIds(v map[string]string) {
	o.ProviderIds = v
}

// GetYear returns the Year field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerInfo) GetYear() int32 {
	if o == nil || IsNil(o.Year.Get()) {
		var ret int32
		return ret
	}
	return *o.Year.Get()
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerInfo) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Year.Get(), o.Year.IsSet()
}

// HasYear returns a boolean if a field has been set.
func (o *TrailerInfo) HasYear() bool {
	if o != nil && o.Year.IsSet() {
		return true
	}

	return false
}

// SetYear gets a reference to the given NullableInt32 and assigns it to the Year field.
func (o *TrailerInfo) SetYear(v int32) {
	o.Year.Set(&v)
}
// SetYearNil sets the value for Year to be an explicit nil
func (o *TrailerInfo) SetYearNil() {
	o.Year.Set(nil)
}

// UnsetYear ensures that no value is present for Year, not even an explicit nil
func (o *TrailerInfo) UnsetYear() {
	o.Year.Unset()
}

// GetIndexNumber returns the IndexNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerInfo) GetIndexNumber() int32 {
	if o == nil || IsNil(o.IndexNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.IndexNumber.Get()
}

// GetIndexNumberOk returns a tuple with the IndexNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerInfo) GetIndexNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexNumber.Get(), o.IndexNumber.IsSet()
}

// HasIndexNumber returns a boolean if a field has been set.
func (o *TrailerInfo) HasIndexNumber() bool {
	if o != nil && o.IndexNumber.IsSet() {
		return true
	}

	return false
}

// SetIndexNumber gets a reference to the given NullableInt32 and assigns it to the IndexNumber field.
func (o *TrailerInfo) SetIndexNumber(v int32) {
	o.IndexNumber.Set(&v)
}
// SetIndexNumberNil sets the value for IndexNumber to be an explicit nil
func (o *TrailerInfo) SetIndexNumberNil() {
	o.IndexNumber.Set(nil)
}

// UnsetIndexNumber ensures that no value is present for IndexNumber, not even an explicit nil
func (o *TrailerInfo) UnsetIndexNumber() {
	o.IndexNumber.Unset()
}

// GetParentIndexNumber returns the ParentIndexNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerInfo) GetParentIndexNumber() int32 {
	if o == nil || IsNil(o.ParentIndexNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.ParentIndexNumber.Get()
}

// GetParentIndexNumberOk returns a tuple with the ParentIndexNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerInfo) GetParentIndexNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentIndexNumber.Get(), o.ParentIndexNumber.IsSet()
}

// HasParentIndexNumber returns a boolean if a field has been set.
func (o *TrailerInfo) HasParentIndexNumber() bool {
	if o != nil && o.ParentIndexNumber.IsSet() {
		return true
	}

	return false
}

// SetParentIndexNumber gets a reference to the given NullableInt32 and assigns it to the ParentIndexNumber field.
func (o *TrailerInfo) SetParentIndexNumber(v int32) {
	o.ParentIndexNumber.Set(&v)
}
// SetParentIndexNumberNil sets the value for ParentIndexNumber to be an explicit nil
func (o *TrailerInfo) SetParentIndexNumberNil() {
	o.ParentIndexNumber.Set(nil)
}

// UnsetParentIndexNumber ensures that no value is present for ParentIndexNumber, not even an explicit nil
func (o *TrailerInfo) UnsetParentIndexNumber() {
	o.ParentIndexNumber.Unset()
}

// GetPremiereDate returns the PremiereDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrailerInfo) GetPremiereDate() time.Time {
	if o == nil || IsNil(o.PremiereDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.PremiereDate.Get()
}

// GetPremiereDateOk returns a tuple with the PremiereDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrailerInfo) GetPremiereDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PremiereDate.Get(), o.PremiereDate.IsSet()
}

// HasPremiereDate returns a boolean if a field has been set.
func (o *TrailerInfo) HasPremiereDate() bool {
	if o != nil && o.PremiereDate.IsSet() {
		return true
	}

	return false
}

// SetPremiereDate gets a reference to the given NullableTime and assigns it to the PremiereDate field.
func (o *TrailerInfo) SetPremiereDate(v time.Time) {
	o.PremiereDate.Set(&v)
}
// SetPremiereDateNil sets the value for PremiereDate to be an explicit nil
func (o *TrailerInfo) SetPremiereDateNil() {
	o.PremiereDate.Set(nil)
}

// UnsetPremiereDate ensures that no value is present for PremiereDate, not even an explicit nil
func (o *TrailerInfo) UnsetPremiereDate() {
	o.PremiereDate.Unset()
}

// GetIsAutomated returns the IsAutomated field value if set, zero value otherwise.
func (o *TrailerInfo) GetIsAutomated() bool {
	if o == nil || IsNil(o.IsAutomated) {
		var ret bool
		return ret
	}
	return *o.IsAutomated
}

// GetIsAutomatedOk returns a tuple with the IsAutomated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerInfo) GetIsAutomatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutomated) {
		return nil, false
	}
	return o.IsAutomated, true
}

// HasIsAutomated returns a boolean if a field has been set.
func (o *TrailerInfo) HasIsAutomated() bool {
	if o != nil && !IsNil(o.IsAutomated) {
		return true
	}

	return false
}

// SetIsAutomated gets a reference to the given bool and assigns it to the IsAutomated field.
func (o *TrailerInfo) SetIsAutomated(v bool) {
	o.IsAutomated = &v
}

func (o TrailerInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrailerInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.OriginalTitle.IsSet() {
		toSerialize["OriginalTitle"] = o.OriginalTitle.Get()
	}
	if o.Path.IsSet() {
		toSerialize["Path"] = o.Path.Get()
	}
	if o.MetadataLanguage.IsSet() {
		toSerialize["MetadataLanguage"] = o.MetadataLanguage.Get()
	}
	if o.MetadataCountryCode.IsSet() {
		toSerialize["MetadataCountryCode"] = o.MetadataCountryCode.Get()
	}
	if o.ProviderIds != nil {
		toSerialize["ProviderIds"] = o.ProviderIds
	}
	if o.Year.IsSet() {
		toSerialize["Year"] = o.Year.Get()
	}
	if o.IndexNumber.IsSet() {
		toSerialize["IndexNumber"] = o.IndexNumber.Get()
	}
	if o.ParentIndexNumber.IsSet() {
		toSerialize["ParentIndexNumber"] = o.ParentIndexNumber.Get()
	}
	if o.PremiereDate.IsSet() {
		toSerialize["PremiereDate"] = o.PremiereDate.Get()
	}
	if !IsNil(o.IsAutomated) {
		toSerialize["IsAutomated"] = o.IsAutomated
	}
	return toSerialize, nil
}

type NullableTrailerInfo struct {
	value *TrailerInfo
	isSet bool
}

func (v NullableTrailerInfo) Get() *TrailerInfo {
	return v.value
}

func (v *NullableTrailerInfo) Set(val *TrailerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTrailerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTrailerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrailerInfo(val *TrailerInfo) *NullableTrailerInfo {
	return &NullableTrailerInfo{value: val, isSet: true}
}

func (v NullableTrailerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrailerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


