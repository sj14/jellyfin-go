/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SeriesInfoRemoteSearchQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesInfoRemoteSearchQuery{}

// SeriesInfoRemoteSearchQuery struct for SeriesInfoRemoteSearchQuery
type SeriesInfoRemoteSearchQuery struct {
	SearchInfo NullableSeriesInfo `json:"SearchInfo,omitempty"`
	ItemId *string `json:"ItemId,omitempty"`
	// Gets or sets the provider name to search within if set.
	SearchProviderName NullableString `json:"SearchProviderName,omitempty"`
	// Gets or sets a value indicating whether disabled providers should be included.
	IncludeDisabledProviders *bool `json:"IncludeDisabledProviders,omitempty"`
}

// NewSeriesInfoRemoteSearchQuery instantiates a new SeriesInfoRemoteSearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesInfoRemoteSearchQuery() *SeriesInfoRemoteSearchQuery {
	this := SeriesInfoRemoteSearchQuery{}
	return &this
}

// NewSeriesInfoRemoteSearchQueryWithDefaults instantiates a new SeriesInfoRemoteSearchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesInfoRemoteSearchQueryWithDefaults() *SeriesInfoRemoteSearchQuery {
	this := SeriesInfoRemoteSearchQuery{}
	return &this
}

// GetSearchInfo returns the SearchInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesInfoRemoteSearchQuery) GetSearchInfo() SeriesInfo {
	if o == nil || IsNil(o.SearchInfo.Get()) {
		var ret SeriesInfo
		return ret
	}
	return *o.SearchInfo.Get()
}

// GetSearchInfoOk returns a tuple with the SearchInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesInfoRemoteSearchQuery) GetSearchInfoOk() (*SeriesInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchInfo.Get(), o.SearchInfo.IsSet()
}

// HasSearchInfo returns a boolean if a field has been set.
func (o *SeriesInfoRemoteSearchQuery) HasSearchInfo() bool {
	if o != nil && o.SearchInfo.IsSet() {
		return true
	}

	return false
}

// SetSearchInfo gets a reference to the given NullableSeriesInfo and assigns it to the SearchInfo field.
func (o *SeriesInfoRemoteSearchQuery) SetSearchInfo(v SeriesInfo) {
	o.SearchInfo.Set(&v)
}
// SetSearchInfoNil sets the value for SearchInfo to be an explicit nil
func (o *SeriesInfoRemoteSearchQuery) SetSearchInfoNil() {
	o.SearchInfo.Set(nil)
}

// UnsetSearchInfo ensures that no value is present for SearchInfo, not even an explicit nil
func (o *SeriesInfoRemoteSearchQuery) UnsetSearchInfo() {
	o.SearchInfo.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *SeriesInfoRemoteSearchQuery) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesInfoRemoteSearchQuery) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *SeriesInfoRemoteSearchQuery) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *SeriesInfoRemoteSearchQuery) SetItemId(v string) {
	o.ItemId = &v
}

// GetSearchProviderName returns the SearchProviderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesInfoRemoteSearchQuery) GetSearchProviderName() string {
	if o == nil || IsNil(o.SearchProviderName.Get()) {
		var ret string
		return ret
	}
	return *o.SearchProviderName.Get()
}

// GetSearchProviderNameOk returns a tuple with the SearchProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesInfoRemoteSearchQuery) GetSearchProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchProviderName.Get(), o.SearchProviderName.IsSet()
}

// HasSearchProviderName returns a boolean if a field has been set.
func (o *SeriesInfoRemoteSearchQuery) HasSearchProviderName() bool {
	if o != nil && o.SearchProviderName.IsSet() {
		return true
	}

	return false
}

// SetSearchProviderName gets a reference to the given NullableString and assigns it to the SearchProviderName field.
func (o *SeriesInfoRemoteSearchQuery) SetSearchProviderName(v string) {
	o.SearchProviderName.Set(&v)
}
// SetSearchProviderNameNil sets the value for SearchProviderName to be an explicit nil
func (o *SeriesInfoRemoteSearchQuery) SetSearchProviderNameNil() {
	o.SearchProviderName.Set(nil)
}

// UnsetSearchProviderName ensures that no value is present for SearchProviderName, not even an explicit nil
func (o *SeriesInfoRemoteSearchQuery) UnsetSearchProviderName() {
	o.SearchProviderName.Unset()
}

// GetIncludeDisabledProviders returns the IncludeDisabledProviders field value if set, zero value otherwise.
func (o *SeriesInfoRemoteSearchQuery) GetIncludeDisabledProviders() bool {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		var ret bool
		return ret
	}
	return *o.IncludeDisabledProviders
}

// GetIncludeDisabledProvidersOk returns a tuple with the IncludeDisabledProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesInfoRemoteSearchQuery) GetIncludeDisabledProvidersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		return nil, false
	}
	return o.IncludeDisabledProviders, true
}

// HasIncludeDisabledProviders returns a boolean if a field has been set.
func (o *SeriesInfoRemoteSearchQuery) HasIncludeDisabledProviders() bool {
	if o != nil && !IsNil(o.IncludeDisabledProviders) {
		return true
	}

	return false
}

// SetIncludeDisabledProviders gets a reference to the given bool and assigns it to the IncludeDisabledProviders field.
func (o *SeriesInfoRemoteSearchQuery) SetIncludeDisabledProviders(v bool) {
	o.IncludeDisabledProviders = &v
}

func (o SeriesInfoRemoteSearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesInfoRemoteSearchQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SearchInfo.IsSet() {
		toSerialize["SearchInfo"] = o.SearchInfo.Get()
	}
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
	}
	if o.SearchProviderName.IsSet() {
		toSerialize["SearchProviderName"] = o.SearchProviderName.Get()
	}
	if !IsNil(o.IncludeDisabledProviders) {
		toSerialize["IncludeDisabledProviders"] = o.IncludeDisabledProviders
	}
	return toSerialize, nil
}

type NullableSeriesInfoRemoteSearchQuery struct {
	value *SeriesInfoRemoteSearchQuery
	isSet bool
}

func (v NullableSeriesInfoRemoteSearchQuery) Get() *SeriesInfoRemoteSearchQuery {
	return v.value
}

func (v *NullableSeriesInfoRemoteSearchQuery) Set(val *SeriesInfoRemoteSearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesInfoRemoteSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesInfoRemoteSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesInfoRemoteSearchQuery(val *SeriesInfoRemoteSearchQuery) *NullableSeriesInfoRemoteSearchQuery {
	return &NullableSeriesInfoRemoteSearchQuery{value: val, isSet: true}
}

func (v NullableSeriesInfoRemoteSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesInfoRemoteSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


