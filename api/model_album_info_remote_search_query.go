/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AlbumInfoRemoteSearchQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlbumInfoRemoteSearchQuery{}

// AlbumInfoRemoteSearchQuery struct for AlbumInfoRemoteSearchQuery
type AlbumInfoRemoteSearchQuery struct {
	SearchInfo NullableAlbumInfo `json:"SearchInfo,omitempty"`
	ItemId *string `json:"ItemId,omitempty"`
	// Gets or sets the provider name to search within if set.
	SearchProviderName NullableString `json:"SearchProviderName,omitempty"`
	// Gets or sets a value indicating whether disabled providers should be included.
	IncludeDisabledProviders *bool `json:"IncludeDisabledProviders,omitempty"`
}

// NewAlbumInfoRemoteSearchQuery instantiates a new AlbumInfoRemoteSearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlbumInfoRemoteSearchQuery() *AlbumInfoRemoteSearchQuery {
	this := AlbumInfoRemoteSearchQuery{}
	return &this
}

// NewAlbumInfoRemoteSearchQueryWithDefaults instantiates a new AlbumInfoRemoteSearchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlbumInfoRemoteSearchQueryWithDefaults() *AlbumInfoRemoteSearchQuery {
	this := AlbumInfoRemoteSearchQuery{}
	return &this
}

// GetSearchInfo returns the SearchInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumInfoRemoteSearchQuery) GetSearchInfo() AlbumInfo {
	if o == nil || IsNil(o.SearchInfo.Get()) {
		var ret AlbumInfo
		return ret
	}
	return *o.SearchInfo.Get()
}

// GetSearchInfoOk returns a tuple with the SearchInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumInfoRemoteSearchQuery) GetSearchInfoOk() (*AlbumInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchInfo.Get(), o.SearchInfo.IsSet()
}

// HasSearchInfo returns a boolean if a field has been set.
func (o *AlbumInfoRemoteSearchQuery) HasSearchInfo() bool {
	if o != nil && o.SearchInfo.IsSet() {
		return true
	}

	return false
}

// SetSearchInfo gets a reference to the given NullableAlbumInfo and assigns it to the SearchInfo field.
func (o *AlbumInfoRemoteSearchQuery) SetSearchInfo(v AlbumInfo) {
	o.SearchInfo.Set(&v)
}
// SetSearchInfoNil sets the value for SearchInfo to be an explicit nil
func (o *AlbumInfoRemoteSearchQuery) SetSearchInfoNil() {
	o.SearchInfo.Set(nil)
}

// UnsetSearchInfo ensures that no value is present for SearchInfo, not even an explicit nil
func (o *AlbumInfoRemoteSearchQuery) UnsetSearchInfo() {
	o.SearchInfo.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *AlbumInfoRemoteSearchQuery) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumInfoRemoteSearchQuery) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *AlbumInfoRemoteSearchQuery) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *AlbumInfoRemoteSearchQuery) SetItemId(v string) {
	o.ItemId = &v
}

// GetSearchProviderName returns the SearchProviderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AlbumInfoRemoteSearchQuery) GetSearchProviderName() string {
	if o == nil || IsNil(o.SearchProviderName.Get()) {
		var ret string
		return ret
	}
	return *o.SearchProviderName.Get()
}

// GetSearchProviderNameOk returns a tuple with the SearchProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AlbumInfoRemoteSearchQuery) GetSearchProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchProviderName.Get(), o.SearchProviderName.IsSet()
}

// HasSearchProviderName returns a boolean if a field has been set.
func (o *AlbumInfoRemoteSearchQuery) HasSearchProviderName() bool {
	if o != nil && o.SearchProviderName.IsSet() {
		return true
	}

	return false
}

// SetSearchProviderName gets a reference to the given NullableString and assigns it to the SearchProviderName field.
func (o *AlbumInfoRemoteSearchQuery) SetSearchProviderName(v string) {
	o.SearchProviderName.Set(&v)
}
// SetSearchProviderNameNil sets the value for SearchProviderName to be an explicit nil
func (o *AlbumInfoRemoteSearchQuery) SetSearchProviderNameNil() {
	o.SearchProviderName.Set(nil)
}

// UnsetSearchProviderName ensures that no value is present for SearchProviderName, not even an explicit nil
func (o *AlbumInfoRemoteSearchQuery) UnsetSearchProviderName() {
	o.SearchProviderName.Unset()
}

// GetIncludeDisabledProviders returns the IncludeDisabledProviders field value if set, zero value otherwise.
func (o *AlbumInfoRemoteSearchQuery) GetIncludeDisabledProviders() bool {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		var ret bool
		return ret
	}
	return *o.IncludeDisabledProviders
}

// GetIncludeDisabledProvidersOk returns a tuple with the IncludeDisabledProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlbumInfoRemoteSearchQuery) GetIncludeDisabledProvidersOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludeDisabledProviders) {
		return nil, false
	}
	return o.IncludeDisabledProviders, true
}

// HasIncludeDisabledProviders returns a boolean if a field has been set.
func (o *AlbumInfoRemoteSearchQuery) HasIncludeDisabledProviders() bool {
	if o != nil && !IsNil(o.IncludeDisabledProviders) {
		return true
	}

	return false
}

// SetIncludeDisabledProviders gets a reference to the given bool and assigns it to the IncludeDisabledProviders field.
func (o *AlbumInfoRemoteSearchQuery) SetIncludeDisabledProviders(v bool) {
	o.IncludeDisabledProviders = &v
}

func (o AlbumInfoRemoteSearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlbumInfoRemoteSearchQuery) ToMap() (map[string]interface{}, error) {
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

type NullableAlbumInfoRemoteSearchQuery struct {
	value *AlbumInfoRemoteSearchQuery
	isSet bool
}

func (v NullableAlbumInfoRemoteSearchQuery) Get() *AlbumInfoRemoteSearchQuery {
	return v.value
}

func (v *NullableAlbumInfoRemoteSearchQuery) Set(val *AlbumInfoRemoteSearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableAlbumInfoRemoteSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableAlbumInfoRemoteSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlbumInfoRemoteSearchQuery(val *AlbumInfoRemoteSearchQuery) *NullableAlbumInfoRemoteSearchQuery {
	return &NullableAlbumInfoRemoteSearchQuery{value: val, isSet: true}
}

func (v NullableAlbumInfoRemoteSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlbumInfoRemoteSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


