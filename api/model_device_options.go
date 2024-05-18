/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DeviceOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceOptions{}

// DeviceOptions An entity representing custom options for a device.
type DeviceOptions struct {
	// Gets the id.
	Id *int32 `json:"Id,omitempty"`
	// Gets the device id.
	DeviceId *string `json:"DeviceId,omitempty"`
	// Gets or sets the custom name.
	CustomName NullableString `json:"CustomName,omitempty"`
}

// NewDeviceOptions instantiates a new DeviceOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceOptions() *DeviceOptions {
	this := DeviceOptions{}
	return &this
}

// NewDeviceOptionsWithDefaults instantiates a new DeviceOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceOptionsWithDefaults() *DeviceOptions {
	this := DeviceOptions{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceOptions) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceOptions) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceOptions) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DeviceOptions) SetId(v int32) {
	o.Id = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *DeviceOptions) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceOptions) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DeviceOptions) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *DeviceOptions) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetCustomName returns the CustomName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceOptions) GetCustomName() string {
	if o == nil || IsNil(o.CustomName.Get()) {
		var ret string
		return ret
	}
	return *o.CustomName.Get()
}

// GetCustomNameOk returns a tuple with the CustomName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceOptions) GetCustomNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomName.Get(), o.CustomName.IsSet()
}

// HasCustomName returns a boolean if a field has been set.
func (o *DeviceOptions) HasCustomName() bool {
	if o != nil && o.CustomName.IsSet() {
		return true
	}

	return false
}

// SetCustomName gets a reference to the given NullableString and assigns it to the CustomName field.
func (o *DeviceOptions) SetCustomName(v string) {
	o.CustomName.Set(&v)
}
// SetCustomNameNil sets the value for CustomName to be an explicit nil
func (o *DeviceOptions) SetCustomNameNil() {
	o.CustomName.Set(nil)
}

// UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
func (o *DeviceOptions) UnsetCustomName() {
	o.CustomName.Unset()
}

func (o DeviceOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.DeviceId) {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.CustomName.IsSet() {
		toSerialize["CustomName"] = o.CustomName.Get()
	}
	return toSerialize, nil
}

type NullableDeviceOptions struct {
	value *DeviceOptions
	isSet bool
}

func (v NullableDeviceOptions) Get() *DeviceOptions {
	return v.value
}

func (v *NullableDeviceOptions) Set(val *DeviceOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceOptions(val *DeviceOptions) *NullableDeviceOptions {
	return &NullableDeviceOptions{value: val, isSet: true}
}

func (v NullableDeviceOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


