/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AuthenticationInfoQueryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationInfoQueryResult{}

// AuthenticationInfoQueryResult Query result container.
type AuthenticationInfoQueryResult struct {
	// Gets or sets the items.
	Items []AuthenticationInfo `json:"Items,omitempty"`
	// Gets or sets the total number of records available.
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty"`
	// Gets or sets the index of the first record in Items.
	StartIndex *int32 `json:"StartIndex,omitempty"`
}

// NewAuthenticationInfoQueryResult instantiates a new AuthenticationInfoQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationInfoQueryResult() *AuthenticationInfoQueryResult {
	this := AuthenticationInfoQueryResult{}
	return &this
}

// NewAuthenticationInfoQueryResultWithDefaults instantiates a new AuthenticationInfoQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationInfoQueryResultWithDefaults() *AuthenticationInfoQueryResult {
	this := AuthenticationInfoQueryResult{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AuthenticationInfoQueryResult) GetItems() []AuthenticationInfo {
	if o == nil || IsNil(o.Items) {
		var ret []AuthenticationInfo
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoQueryResult) GetItemsOk() ([]AuthenticationInfo, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AuthenticationInfoQueryResult) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AuthenticationInfo and assigns it to the Items field.
func (o *AuthenticationInfoQueryResult) SetItems(v []AuthenticationInfo) {
	o.Items = v
}

// GetTotalRecordCount returns the TotalRecordCount field value if set, zero value otherwise.
func (o *AuthenticationInfoQueryResult) GetTotalRecordCount() int32 {
	if o == nil || IsNil(o.TotalRecordCount) {
		var ret int32
		return ret
	}
	return *o.TotalRecordCount
}

// GetTotalRecordCountOk returns a tuple with the TotalRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoQueryResult) GetTotalRecordCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalRecordCount) {
		return nil, false
	}
	return o.TotalRecordCount, true
}

// HasTotalRecordCount returns a boolean if a field has been set.
func (o *AuthenticationInfoQueryResult) HasTotalRecordCount() bool {
	if o != nil && !IsNil(o.TotalRecordCount) {
		return true
	}

	return false
}

// SetTotalRecordCount gets a reference to the given int32 and assigns it to the TotalRecordCount field.
func (o *AuthenticationInfoQueryResult) SetTotalRecordCount(v int32) {
	o.TotalRecordCount = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *AuthenticationInfoQueryResult) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex) {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoQueryResult) GetStartIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.StartIndex) {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *AuthenticationInfoQueryResult) HasStartIndex() bool {
	if o != nil && !IsNil(o.StartIndex) {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *AuthenticationInfoQueryResult) SetStartIndex(v int32) {
	o.StartIndex = &v
}

func (o AuthenticationInfoQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationInfoQueryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["Items"] = o.Items
	}
	if !IsNil(o.TotalRecordCount) {
		toSerialize["TotalRecordCount"] = o.TotalRecordCount
	}
	if !IsNil(o.StartIndex) {
		toSerialize["StartIndex"] = o.StartIndex
	}
	return toSerialize, nil
}

type NullableAuthenticationInfoQueryResult struct {
	value *AuthenticationInfoQueryResult
	isSet bool
}

func (v NullableAuthenticationInfoQueryResult) Get() *AuthenticationInfoQueryResult {
	return v.value
}

func (v *NullableAuthenticationInfoQueryResult) Set(val *AuthenticationInfoQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationInfoQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationInfoQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationInfoQueryResult(val *AuthenticationInfoQueryResult) *NullableAuthenticationInfoQueryResult {
	return &NullableAuthenticationInfoQueryResult{value: val, isSet: true}
}

func (v NullableAuthenticationInfoQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationInfoQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


