/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MetadataOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataOptions{}

// MetadataOptions Class MetadataOptions.
type MetadataOptions struct {
	ItemType NullableString `json:"ItemType,omitempty"`
	DisabledMetadataSavers []string `json:"DisabledMetadataSavers,omitempty"`
	LocalMetadataReaderOrder []string `json:"LocalMetadataReaderOrder,omitempty"`
	DisabledMetadataFetchers []string `json:"DisabledMetadataFetchers,omitempty"`
	MetadataFetcherOrder []string `json:"MetadataFetcherOrder,omitempty"`
	DisabledImageFetchers []string `json:"DisabledImageFetchers,omitempty"`
	ImageFetcherOrder []string `json:"ImageFetcherOrder,omitempty"`
}

// NewMetadataOptions instantiates a new MetadataOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataOptions() *MetadataOptions {
	this := MetadataOptions{}
	return &this
}

// NewMetadataOptionsWithDefaults instantiates a new MetadataOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataOptionsWithDefaults() *MetadataOptions {
	this := MetadataOptions{}
	return &this
}

// GetItemType returns the ItemType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataOptions) GetItemType() string {
	if o == nil || IsNil(o.ItemType.Get()) {
		var ret string
		return ret
	}
	return *o.ItemType.Get()
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataOptions) GetItemTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemType.Get(), o.ItemType.IsSet()
}

// HasItemType returns a boolean if a field has been set.
func (o *MetadataOptions) HasItemType() bool {
	if o != nil && o.ItemType.IsSet() {
		return true
	}

	return false
}

// SetItemType gets a reference to the given NullableString and assigns it to the ItemType field.
func (o *MetadataOptions) SetItemType(v string) {
	o.ItemType.Set(&v)
}
// SetItemTypeNil sets the value for ItemType to be an explicit nil
func (o *MetadataOptions) SetItemTypeNil() {
	o.ItemType.Set(nil)
}

// UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
func (o *MetadataOptions) UnsetItemType() {
	o.ItemType.Unset()
}

// GetDisabledMetadataSavers returns the DisabledMetadataSavers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataOptions) GetDisabledMetadataSavers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DisabledMetadataSavers
}

// GetDisabledMetadataSaversOk returns a tuple with the DisabledMetadataSavers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataOptions) GetDisabledMetadataSaversOk() ([]string, bool) {
	if o == nil || IsNil(o.DisabledMetadataSavers) {
		return nil, false
	}
	return o.DisabledMetadataSavers, true
}

// HasDisabledMetadataSavers returns a boolean if a field has been set.
func (o *MetadataOptions) HasDisabledMetadataSavers() bool {
	if o != nil && !IsNil(o.DisabledMetadataSavers) {
		return true
	}

	return false
}

// SetDisabledMetadataSavers gets a reference to the given []string and assigns it to the DisabledMetadataSavers field.
func (o *MetadataOptions) SetDisabledMetadataSavers(v []string) {
	o.DisabledMetadataSavers = v
}

// GetLocalMetadataReaderOrder returns the LocalMetadataReaderOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataOptions) GetLocalMetadataReaderOrder() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.LocalMetadataReaderOrder
}

// GetLocalMetadataReaderOrderOk returns a tuple with the LocalMetadataReaderOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataOptions) GetLocalMetadataReaderOrderOk() ([]string, bool) {
	if o == nil || IsNil(o.LocalMetadataReaderOrder) {
		return nil, false
	}
	return o.LocalMetadataReaderOrder, true
}

// HasLocalMetadataReaderOrder returns a boolean if a field has been set.
func (o *MetadataOptions) HasLocalMetadataReaderOrder() bool {
	if o != nil && !IsNil(o.LocalMetadataReaderOrder) {
		return true
	}

	return false
}

// SetLocalMetadataReaderOrder gets a reference to the given []string and assigns it to the LocalMetadataReaderOrder field.
func (o *MetadataOptions) SetLocalMetadataReaderOrder(v []string) {
	o.LocalMetadataReaderOrder = v
}

// GetDisabledMetadataFetchers returns the DisabledMetadataFetchers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataOptions) GetDisabledMetadataFetchers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DisabledMetadataFetchers
}

// GetDisabledMetadataFetchersOk returns a tuple with the DisabledMetadataFetchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataOptions) GetDisabledMetadataFetchersOk() ([]string, bool) {
	if o == nil || IsNil(o.DisabledMetadataFetchers) {
		return nil, false
	}
	return o.DisabledMetadataFetchers, true
}

// HasDisabledMetadataFetchers returns a boolean if a field has been set.
func (o *MetadataOptions) HasDisabledMetadataFetchers() bool {
	if o != nil && !IsNil(o.DisabledMetadataFetchers) {
		return true
	}

	return false
}

// SetDisabledMetadataFetchers gets a reference to the given []string and assigns it to the DisabledMetadataFetchers field.
func (o *MetadataOptions) SetDisabledMetadataFetchers(v []string) {
	o.DisabledMetadataFetchers = v
}

// GetMetadataFetcherOrder returns the MetadataFetcherOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataOptions) GetMetadataFetcherOrder() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MetadataFetcherOrder
}

// GetMetadataFetcherOrderOk returns a tuple with the MetadataFetcherOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataOptions) GetMetadataFetcherOrderOk() ([]string, bool) {
	if o == nil || IsNil(o.MetadataFetcherOrder) {
		return nil, false
	}
	return o.MetadataFetcherOrder, true
}

// HasMetadataFetcherOrder returns a boolean if a field has been set.
func (o *MetadataOptions) HasMetadataFetcherOrder() bool {
	if o != nil && !IsNil(o.MetadataFetcherOrder) {
		return true
	}

	return false
}

// SetMetadataFetcherOrder gets a reference to the given []string and assigns it to the MetadataFetcherOrder field.
func (o *MetadataOptions) SetMetadataFetcherOrder(v []string) {
	o.MetadataFetcherOrder = v
}

// GetDisabledImageFetchers returns the DisabledImageFetchers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataOptions) GetDisabledImageFetchers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DisabledImageFetchers
}

// GetDisabledImageFetchersOk returns a tuple with the DisabledImageFetchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataOptions) GetDisabledImageFetchersOk() ([]string, bool) {
	if o == nil || IsNil(o.DisabledImageFetchers) {
		return nil, false
	}
	return o.DisabledImageFetchers, true
}

// HasDisabledImageFetchers returns a boolean if a field has been set.
func (o *MetadataOptions) HasDisabledImageFetchers() bool {
	if o != nil && !IsNil(o.DisabledImageFetchers) {
		return true
	}

	return false
}

// SetDisabledImageFetchers gets a reference to the given []string and assigns it to the DisabledImageFetchers field.
func (o *MetadataOptions) SetDisabledImageFetchers(v []string) {
	o.DisabledImageFetchers = v
}

// GetImageFetcherOrder returns the ImageFetcherOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MetadataOptions) GetImageFetcherOrder() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ImageFetcherOrder
}

// GetImageFetcherOrderOk returns a tuple with the ImageFetcherOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataOptions) GetImageFetcherOrderOk() ([]string, bool) {
	if o == nil || IsNil(o.ImageFetcherOrder) {
		return nil, false
	}
	return o.ImageFetcherOrder, true
}

// HasImageFetcherOrder returns a boolean if a field has been set.
func (o *MetadataOptions) HasImageFetcherOrder() bool {
	if o != nil && !IsNil(o.ImageFetcherOrder) {
		return true
	}

	return false
}

// SetImageFetcherOrder gets a reference to the given []string and assigns it to the ImageFetcherOrder field.
func (o *MetadataOptions) SetImageFetcherOrder(v []string) {
	o.ImageFetcherOrder = v
}

func (o MetadataOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ItemType.IsSet() {
		toSerialize["ItemType"] = o.ItemType.Get()
	}
	if o.DisabledMetadataSavers != nil {
		toSerialize["DisabledMetadataSavers"] = o.DisabledMetadataSavers
	}
	if o.LocalMetadataReaderOrder != nil {
		toSerialize["LocalMetadataReaderOrder"] = o.LocalMetadataReaderOrder
	}
	if o.DisabledMetadataFetchers != nil {
		toSerialize["DisabledMetadataFetchers"] = o.DisabledMetadataFetchers
	}
	if o.MetadataFetcherOrder != nil {
		toSerialize["MetadataFetcherOrder"] = o.MetadataFetcherOrder
	}
	if o.DisabledImageFetchers != nil {
		toSerialize["DisabledImageFetchers"] = o.DisabledImageFetchers
	}
	if o.ImageFetcherOrder != nil {
		toSerialize["ImageFetcherOrder"] = o.ImageFetcherOrder
	}
	return toSerialize, nil
}

type NullableMetadataOptions struct {
	value *MetadataOptions
	isSet bool
}

func (v NullableMetadataOptions) Get() *MetadataOptions {
	return v.value
}

func (v *NullableMetadataOptions) Set(val *MetadataOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataOptions(val *MetadataOptions) *NullableMetadataOptions {
	return &NullableMetadataOptions{value: val, isSet: true}
}

func (v NullableMetadataOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


