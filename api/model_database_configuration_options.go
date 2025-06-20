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

// checks if the DatabaseConfigurationOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseConfigurationOptions{}

// DatabaseConfigurationOptions Options to configure jellyfins managed database.
type DatabaseConfigurationOptions struct {
	// Gets or Sets the type of database jellyfin should use.
	DatabaseType *string `json:"DatabaseType,omitempty"`
	// Gets or sets the options required to use a custom database provider.
	CustomProviderOptions NullableCustomDatabaseOptions `json:"CustomProviderOptions,omitempty"`
	// Gets or Sets the kind of locking behavior jellyfin should perform. Possible options are \"NoLock\", \"Pessimistic\", \"Optimistic\".  Defaults to \"NoLock\".
	LockingBehavior *DatabaseLockingBehaviorTypes `json:"LockingBehavior,omitempty"`
}

// NewDatabaseConfigurationOptions instantiates a new DatabaseConfigurationOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseConfigurationOptions() *DatabaseConfigurationOptions {
	this := DatabaseConfigurationOptions{}
	return &this
}

// NewDatabaseConfigurationOptionsWithDefaults instantiates a new DatabaseConfigurationOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseConfigurationOptionsWithDefaults() *DatabaseConfigurationOptions {
	this := DatabaseConfigurationOptions{}
	return &this
}

// GetDatabaseType returns the DatabaseType field value if set, zero value otherwise.
func (o *DatabaseConfigurationOptions) GetDatabaseType() string {
	if o == nil || IsNil(o.DatabaseType) {
		var ret string
		return ret
	}
	return *o.DatabaseType
}

// GetDatabaseTypeOk returns a tuple with the DatabaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseConfigurationOptions) GetDatabaseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DatabaseType) {
		return nil, false
	}
	return o.DatabaseType, true
}

// HasDatabaseType returns a boolean if a field has been set.
func (o *DatabaseConfigurationOptions) HasDatabaseType() bool {
	if o != nil && !IsNil(o.DatabaseType) {
		return true
	}

	return false
}

// SetDatabaseType gets a reference to the given string and assigns it to the DatabaseType field.
func (o *DatabaseConfigurationOptions) SetDatabaseType(v string) {
	o.DatabaseType = &v
}

// GetCustomProviderOptions returns the CustomProviderOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatabaseConfigurationOptions) GetCustomProviderOptions() CustomDatabaseOptions {
	if o == nil || IsNil(o.CustomProviderOptions.Get()) {
		var ret CustomDatabaseOptions
		return ret
	}
	return *o.CustomProviderOptions.Get()
}

// GetCustomProviderOptionsOk returns a tuple with the CustomProviderOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatabaseConfigurationOptions) GetCustomProviderOptionsOk() (*CustomDatabaseOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomProviderOptions.Get(), o.CustomProviderOptions.IsSet()
}

// HasCustomProviderOptions returns a boolean if a field has been set.
func (o *DatabaseConfigurationOptions) HasCustomProviderOptions() bool {
	if o != nil && o.CustomProviderOptions.IsSet() {
		return true
	}

	return false
}

// SetCustomProviderOptions gets a reference to the given NullableCustomDatabaseOptions and assigns it to the CustomProviderOptions field.
func (o *DatabaseConfigurationOptions) SetCustomProviderOptions(v CustomDatabaseOptions) {
	o.CustomProviderOptions.Set(&v)
}
// SetCustomProviderOptionsNil sets the value for CustomProviderOptions to be an explicit nil
func (o *DatabaseConfigurationOptions) SetCustomProviderOptionsNil() {
	o.CustomProviderOptions.Set(nil)
}

// UnsetCustomProviderOptions ensures that no value is present for CustomProviderOptions, not even an explicit nil
func (o *DatabaseConfigurationOptions) UnsetCustomProviderOptions() {
	o.CustomProviderOptions.Unset()
}

// GetLockingBehavior returns the LockingBehavior field value if set, zero value otherwise.
func (o *DatabaseConfigurationOptions) GetLockingBehavior() DatabaseLockingBehaviorTypes {
	if o == nil || IsNil(o.LockingBehavior) {
		var ret DatabaseLockingBehaviorTypes
		return ret
	}
	return *o.LockingBehavior
}

// GetLockingBehaviorOk returns a tuple with the LockingBehavior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseConfigurationOptions) GetLockingBehaviorOk() (*DatabaseLockingBehaviorTypes, bool) {
	if o == nil || IsNil(o.LockingBehavior) {
		return nil, false
	}
	return o.LockingBehavior, true
}

// HasLockingBehavior returns a boolean if a field has been set.
func (o *DatabaseConfigurationOptions) HasLockingBehavior() bool {
	if o != nil && !IsNil(o.LockingBehavior) {
		return true
	}

	return false
}

// SetLockingBehavior gets a reference to the given DatabaseLockingBehaviorTypes and assigns it to the LockingBehavior field.
func (o *DatabaseConfigurationOptions) SetLockingBehavior(v DatabaseLockingBehaviorTypes) {
	o.LockingBehavior = &v
}

func (o DatabaseConfigurationOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseConfigurationOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatabaseType) {
		toSerialize["DatabaseType"] = o.DatabaseType
	}
	if o.CustomProviderOptions.IsSet() {
		toSerialize["CustomProviderOptions"] = o.CustomProviderOptions.Get()
	}
	if !IsNil(o.LockingBehavior) {
		toSerialize["LockingBehavior"] = o.LockingBehavior
	}
	return toSerialize, nil
}

type NullableDatabaseConfigurationOptions struct {
	value *DatabaseConfigurationOptions
	isSet bool
}

func (v NullableDatabaseConfigurationOptions) Get() *DatabaseConfigurationOptions {
	return v.value
}

func (v *NullableDatabaseConfigurationOptions) Set(val *DatabaseConfigurationOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseConfigurationOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseConfigurationOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseConfigurationOptions(val *DatabaseConfigurationOptions) *NullableDatabaseConfigurationOptions {
	return &NullableDatabaseConfigurationOptions{value: val, isSet: true}
}

func (v NullableDatabaseConfigurationOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseConfigurationOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


