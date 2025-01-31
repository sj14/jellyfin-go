/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the LiveTvInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveTvInfo{}

// LiveTvInfo struct for LiveTvInfo
type LiveTvInfo struct {
	// Gets or sets the services.
	Services []LiveTvServiceInfo `json:"Services,omitempty"`
	// Gets or sets a value indicating whether this instance is enabled.
	IsEnabled *bool `json:"IsEnabled,omitempty"`
	// Gets or sets the enabled users.
	EnabledUsers []string `json:"EnabledUsers,omitempty"`
}

// NewLiveTvInfo instantiates a new LiveTvInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveTvInfo() *LiveTvInfo {
	this := LiveTvInfo{}
	return &this
}

// NewLiveTvInfoWithDefaults instantiates a new LiveTvInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveTvInfoWithDefaults() *LiveTvInfo {
	this := LiveTvInfo{}
	return &this
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *LiveTvInfo) GetServices() []LiveTvServiceInfo {
	if o == nil || IsNil(o.Services) {
		var ret []LiveTvServiceInfo
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvInfo) GetServicesOk() ([]LiveTvServiceInfo, bool) {
	if o == nil || IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *LiveTvInfo) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given []LiveTvServiceInfo and assigns it to the Services field.
func (o *LiveTvInfo) SetServices(v []LiveTvServiceInfo) {
	o.Services = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *LiveTvInfo) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvInfo) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *LiveTvInfo) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *LiveTvInfo) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetEnabledUsers returns the EnabledUsers field value if set, zero value otherwise.
func (o *LiveTvInfo) GetEnabledUsers() []string {
	if o == nil || IsNil(o.EnabledUsers) {
		var ret []string
		return ret
	}
	return o.EnabledUsers
}

// GetEnabledUsersOk returns a tuple with the EnabledUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvInfo) GetEnabledUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.EnabledUsers) {
		return nil, false
	}
	return o.EnabledUsers, true
}

// HasEnabledUsers returns a boolean if a field has been set.
func (o *LiveTvInfo) HasEnabledUsers() bool {
	if o != nil && !IsNil(o.EnabledUsers) {
		return true
	}

	return false
}

// SetEnabledUsers gets a reference to the given []string and assigns it to the EnabledUsers field.
func (o *LiveTvInfo) SetEnabledUsers(v []string) {
	o.EnabledUsers = v
}

func (o LiveTvInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveTvInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Services) {
		toSerialize["Services"] = o.Services
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["IsEnabled"] = o.IsEnabled
	}
	if !IsNil(o.EnabledUsers) {
		toSerialize["EnabledUsers"] = o.EnabledUsers
	}
	return toSerialize, nil
}

type NullableLiveTvInfo struct {
	value *LiveTvInfo
	isSet bool
}

func (v NullableLiveTvInfo) Get() *LiveTvInfo {
	return v.value
}

func (v *NullableLiveTvInfo) Set(val *LiveTvInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveTvInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveTvInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveTvInfo(val *LiveTvInfo) *NullableLiveTvInfo {
	return &NullableLiveTvInfo{value: val, isSet: true}
}

func (v NullableLiveTvInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveTvInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


