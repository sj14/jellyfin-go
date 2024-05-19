/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the BoxSetInfoRemoteSearchQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoxSetInfoRemoteSearchQuery{}

// BoxSetInfoRemoteSearchQuery struct for BoxSetInfoRemoteSearchQuery
type BoxSetInfoRemoteSearchQuery struct {
	SearchInfo NullableBoxSetInfo `json:"SearchInfo,omitempty"`
	ItemId *string `json:"ItemId,omitempty"`
	// Gets or sets the provider name to search within if set.
	SearchProviderName NullableString `json:"SearchProviderName,omitempty"`
	// Gets or sets a value indicating whether disabled providers should be included.
	IncludeDisabledProviders *bool `json:"IncludeDisabledProviders,omitempty"`
}

// NewBoxSetInfoRemoteSearchQuery instantiates a new BoxSetInfoRemoteSearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoxSetInfoRemoteSearchQuery() *BoxSetInfoRemoteSearchQuery {
	this := BoxSetInfoRemoteSearchQuery{}
	return &this
}

// NewBoxSetInfoRemoteSearchQueryWithDefaults instantiates a new BoxSetInfoRemoteSearchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoxSetInfoRemoteSearchQueryWithDefaults() *BoxSetInfoRemoteSearchQuery {
	this := BoxSetInfoRemoteSearchQuery{}
	return &this
}

// GetSearchInfo returns the SearchInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoxSetInfoRemoteSearchQuery) GetSearchInfo() BoxSetInfo {
	if o == nil || IsNil(o.SearchInfo.Get()) {
		var ret BoxSetInfo
		return ret
	}
	return *o.SearchInfo.Get()
}

// GetSearchInfoOk returns a tuple with the SearchInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoxSetInfoRemoteSearchQuery) GetSearchInfoOk() (*BoxSetInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchInfo.Get(), o.SearchInfo.IsSet()
}

// HasSearchInfo returns a boolean if a field has been set.
func (o *BoxSetInfoRemoteSearchQuery) HasSearchInfo() bool {
	if o != nil && o.SearchInfo.IsSet() {
		return true
	}

	return false
}

// SetSearchInfo gets a reference to the given NullableBoxSetInfo and assigns it to the SearchInfo field.
func (o *BoxSetInfoRemoteSearchQuery) SetSearchInfo(v BoxSetInfo) {
	o.SearchInfo.Set(&v)
}
// SetSearchInfoNil sets the value for SearchInfo to be an explicit nil
func (o *BoxSetInfoRemoteSearchQuery) SetSearchInfoNil() {
	o.SearchInfo.Set(nil)
}

// UnsetSearchInfo ensures that no value is present for SearchInfo, not even an explicit nil
func (o *BoxSetInfoRemoteSearchQuery) UnsetSearchInfo() {
	o.SearchInfo.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *BoxSetInfoRemoteSearchQuery) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoxSetInfoRemoteSearchQuery) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *BoxSetInfoRemoteSearchQuery) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *BoxSetInfoRemoteSearchQuery) SetItemId(v string) {
	o.ItemId = &v
}

// GetSearchProviderName returns the SearchProviderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BoxSetInfoRemoteSearchQuery) GetSearchProviderName() string {
	if o == nil || IsNil(o.SearchProviderName.Get()) {
		var ret string
		return ret
	}
	return *o.SearchProviderName.Get()
}

// GetSearchProviderNameOk returns a tuple with the SearchProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BoxSetInfoRemoteSearchQuery) GetSearchProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchProviderName.Get(), o.SearchProviderName.IsSet()
}

// HasSearchProviderName returns a boolean if a field has been set.
func (o *BoxSetInfoRemoteSearchQuery) HasSearchProviderName() bool {
	if o != nil && o.SearchProviderName.IsSet() {
		return true
	}

	return false
}

// SetSearchProviderName gets a reference to the given NullableString and assigns it to the SearchProviderName field.
func (o *BoxSetInfoRemoteSearchQuery) SetSearchProviderName(v string) {
	o.SearchProviderName.Set(&v)
}
// SetSearchProviderNameNil sets the value for SearchProviderName to be an explicit nil
func (o *BoxSetInfoRemoteSearchQuery) SetSearchProviderNameNil() {
	o.SearchProviderName.Set(nil)
}

// UnsetSearchProviderName ensures that no value is present for SearchProviderName, not even an explicit nil
func (o *BoxSetInfoRemoteSearchQuery) UnsetSearchProviderName() {
	o.SearchProviderName.Unset()
}

// GetIncludeDisabledProviders returns the IncludeDisabledProviders field value if set, zero value otherwise.
func (o *BoxSetInfoRemoteSearchQuery) GetIncludeDisabledProviders() bool {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		var ret bool
		return ret
	}
	return *o.IncludeDisabledProviders
}

// GetIncludeDisabledProvidersOk returns a tuple with the IncludeDisabledProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoxSetInfoRemoteSearchQuery) GetIncludeDisabledProvidersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		return nil, false
	}
	return o.IncludeDisabledProviders, true
}

// HasIncludeDisabledProviders returns a boolean if a field has been set.
func (o *BoxSetInfoRemoteSearchQuery) HasIncludeDisabledProviders() bool {
	if o != nil && !IsNil(o.IncludeDisabledProviders) {
		return true
	}

	return false
}

// SetIncludeDisabledProviders gets a reference to the given bool and assigns it to the IncludeDisabledProviders field.
func (o *BoxSetInfoRemoteSearchQuery) SetIncludeDisabledProviders(v bool) {
	o.IncludeDisabledProviders = &v
}

func (o BoxSetInfoRemoteSearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoxSetInfoRemoteSearchQuery) ToMap() (map[string]interface{}, error) {
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

type NullableBoxSetInfoRemoteSearchQuery struct {
	value *BoxSetInfoRemoteSearchQuery
	isSet bool
}

func (v NullableBoxSetInfoRemoteSearchQuery) Get() *BoxSetInfoRemoteSearchQuery {
	return v.value
}

func (v *NullableBoxSetInfoRemoteSearchQuery) Set(val *BoxSetInfoRemoteSearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableBoxSetInfoRemoteSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableBoxSetInfoRemoteSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoxSetInfoRemoteSearchQuery(val *BoxSetInfoRemoteSearchQuery) *NullableBoxSetInfoRemoteSearchQuery {
	return &NullableBoxSetInfoRemoteSearchQuery{value: val, isSet: true}
}

func (v NullableBoxSetInfoRemoteSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoxSetInfoRemoteSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


