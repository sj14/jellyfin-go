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

// checks if the DeviceOptionsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceOptionsDto{}

// DeviceOptionsDto A dto representing custom options for a device.
type DeviceOptionsDto struct {
	// Gets or sets the id.
	Id *int32 `json:"Id,omitempty"`
	// Gets or sets the device id.
	DeviceId NullableString `json:"DeviceId,omitempty"`
	// Gets or sets the custom name.
	CustomName NullableString `json:"CustomName,omitempty"`
}

// NewDeviceOptionsDto instantiates a new DeviceOptionsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceOptionsDto() *DeviceOptionsDto {
	this := DeviceOptionsDto{}
	return &this
}

// NewDeviceOptionsDtoWithDefaults instantiates a new DeviceOptionsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceOptionsDtoWithDefaults() *DeviceOptionsDto {
	this := DeviceOptionsDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceOptionsDto) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceOptionsDto) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceOptionsDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DeviceOptionsDto) SetId(v int32) {
	o.Id = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceOptionsDto) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId.Get()) {
		var ret string
		return ret
	}
	return *o.DeviceId.Get()
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceOptionsDto) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceId.Get(), o.DeviceId.IsSet()
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DeviceOptionsDto) HasDeviceId() bool {
	if o != nil && o.DeviceId.IsSet() {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given NullableString and assigns it to the DeviceId field.
func (o *DeviceOptionsDto) SetDeviceId(v string) {
	o.DeviceId.Set(&v)
}
// SetDeviceIdNil sets the value for DeviceId to be an explicit nil
func (o *DeviceOptionsDto) SetDeviceIdNil() {
	o.DeviceId.Set(nil)
}

// UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
func (o *DeviceOptionsDto) UnsetDeviceId() {
	o.DeviceId.Unset()
}

// GetCustomName returns the CustomName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeviceOptionsDto) GetCustomName() string {
	if o == nil || IsNil(o.CustomName.Get()) {
		var ret string
		return ret
	}
	return *o.CustomName.Get()
}

// GetCustomNameOk returns a tuple with the CustomName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeviceOptionsDto) GetCustomNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomName.Get(), o.CustomName.IsSet()
}

// HasCustomName returns a boolean if a field has been set.
func (o *DeviceOptionsDto) HasCustomName() bool {
	if o != nil && o.CustomName.IsSet() {
		return true
	}

	return false
}

// SetCustomName gets a reference to the given NullableString and assigns it to the CustomName field.
func (o *DeviceOptionsDto) SetCustomName(v string) {
	o.CustomName.Set(&v)
}
// SetCustomNameNil sets the value for CustomName to be an explicit nil
func (o *DeviceOptionsDto) SetCustomNameNil() {
	o.CustomName.Set(nil)
}

// UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil
func (o *DeviceOptionsDto) UnsetCustomName() {
	o.CustomName.Unset()
}

func (o DeviceOptionsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceOptionsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.DeviceId.IsSet() {
		toSerialize["DeviceId"] = o.DeviceId.Get()
	}
	if o.CustomName.IsSet() {
		toSerialize["CustomName"] = o.CustomName.Get()
	}
	return toSerialize, nil
}

type NullableDeviceOptionsDto struct {
	value *DeviceOptionsDto
	isSet bool
}

func (v NullableDeviceOptionsDto) Get() *DeviceOptionsDto {
	return v.value
}

func (v *NullableDeviceOptionsDto) Set(val *DeviceOptionsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceOptionsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceOptionsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceOptionsDto(val *DeviceOptionsDto) *NullableDeviceOptionsDto {
	return &NullableDeviceOptionsDto{value: val, isSet: true}
}

func (v NullableDeviceOptionsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceOptionsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


