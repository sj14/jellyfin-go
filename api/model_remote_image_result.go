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

// checks if the RemoteImageResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteImageResult{}

// RemoteImageResult Class RemoteImageResult.
type RemoteImageResult struct {
	// Gets or sets the images.
	Images []RemoteImageInfo `json:"Images,omitempty"`
	// Gets or sets the total record count.
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty"`
	// Gets or sets the providers.
	Providers []string `json:"Providers,omitempty"`
}

// NewRemoteImageResult instantiates a new RemoteImageResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteImageResult() *RemoteImageResult {
	this := RemoteImageResult{}
	return &this
}

// NewRemoteImageResultWithDefaults instantiates a new RemoteImageResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteImageResultWithDefaults() *RemoteImageResult {
	this := RemoteImageResult{}
	return &this
}

// GetImages returns the Images field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteImageResult) GetImages() []RemoteImageInfo {
	if o == nil {
		var ret []RemoteImageInfo
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteImageResult) GetImagesOk() ([]RemoteImageInfo, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *RemoteImageResult) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given []RemoteImageInfo and assigns it to the Images field.
func (o *RemoteImageResult) SetImages(v []RemoteImageInfo) {
	o.Images = v
}

// GetTotalRecordCount returns the TotalRecordCount field value if set, zero value otherwise.
func (o *RemoteImageResult) GetTotalRecordCount() int32 {
	if o == nil || IsNil(o.TotalRecordCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordCount
}

// GetTotalRecordCountOk returns a tuple with the TotalRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteImageResult) GetTotalRecordCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordCount) {
		return nil, false
	}
	return o.TotalRecordCount, true
}

// HasTotalRecordCount returns a boolean if a field has been set.
func (o *RemoteImageResult) HasTotalRecordCount() bool {
	if o != nil && !IsNil(o.TotalRecordCount) {
		return true
	}

	return false
}

// SetTotalRecordCount gets a reference to the given int32 and assigns it to the TotalRecordCount field.
func (o *RemoteImageResult) SetTotalRecordCount(v int32) {
	o.TotalRecordCount = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteImageResult) GetProviders() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteImageResult) GetProvidersOk() ([]string, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *RemoteImageResult) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []string and assigns it to the Providers field.
func (o *RemoteImageResult) SetProviders(v []string) {
	o.Providers = v
}

func (o RemoteImageResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteImageResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Images != nil {
		toSerialize["Images"] = o.Images
	}
	if !IsNil(o.TotalRecordCount) {
		toSerialize["TotalRecordCount"] = o.TotalRecordCount
	}
	if o.Providers != nil {
		toSerialize["Providers"] = o.Providers
	}
	return toSerialize, nil
}

type NullableRemoteImageResult struct {
	value *RemoteImageResult
	isSet bool
}

func (v NullableRemoteImageResult) Get() *RemoteImageResult {
	return v.value
}

func (v *NullableRemoteImageResult) Set(val *RemoteImageResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteImageResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteImageResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteImageResult(val *RemoteImageResult) *NullableRemoteImageResult {
	return &NullableRemoteImageResult{value: val, isSet: true}
}

func (v NullableRemoteImageResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteImageResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


