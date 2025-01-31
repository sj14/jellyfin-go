/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the QuickConnectResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickConnectResult{}

// QuickConnectResult Stores the state of an quick connect request.
type QuickConnectResult struct {
	// Gets or sets a value indicating whether this request is authorized.
	Authenticated *bool `json:"Authenticated,omitempty"`
	// Gets the secret value used to uniquely identify this request. Can be used to retrieve authentication information.
	Secret *string `json:"Secret,omitempty"`
	// Gets the user facing code used so the user can quickly differentiate this request from others.
	Code *string `json:"Code,omitempty"`
	// Gets the requesting device id.
	DeviceId *string `json:"DeviceId,omitempty"`
	// Gets the requesting device name.
	DeviceName *string `json:"DeviceName,omitempty"`
	// Gets the requesting app name.
	AppName *string `json:"AppName,omitempty"`
	// Gets the requesting app version.
	AppVersion *string `json:"AppVersion,omitempty"`
	// Gets or sets the DateTime that this request was created.
	DateAdded *time.Time `json:"DateAdded,omitempty"`
}

// NewQuickConnectResult instantiates a new QuickConnectResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickConnectResult() *QuickConnectResult {
	this := QuickConnectResult{}
	return &this
}

// NewQuickConnectResultWithDefaults instantiates a new QuickConnectResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickConnectResultWithDefaults() *QuickConnectResult {
	this := QuickConnectResult{}
	return &this
}

// GetAuthenticated returns the Authenticated field value if set, zero value otherwise.
func (o *QuickConnectResult) GetAuthenticated() bool {
	if o == nil || IsNil(o.Authenticated) {
		var ret bool
		return ret
	}
	return *o.Authenticated
}

// GetAuthenticatedOk returns a tuple with the Authenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickConnectResult) GetAuthenticatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Authenticated) {
		return nil, false
	}
	return o.Authenticated, true
}

// HasAuthenticated returns a boolean if a field has been set.
func (o *QuickConnectResult) HasAuthenticated() bool {
	if o != nil && !IsNil(o.Authenticated) {
		return true
	}

	return false
}

// SetAuthenticated gets a reference to the given bool and assigns it to the Authenticated field.
func (o *QuickConnectResult) SetAuthenticated(v bool) {
	o.Authenticated = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *QuickConnectResult) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickConnectResult) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *QuickConnectResult) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *QuickConnectResult) SetSecret(v string) {
	o.Secret = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *QuickConnectResult) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickConnectResult) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *QuickConnectResult) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *QuickConnectResult) SetCode(v string) {
	o.Code = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *QuickConnectResult) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickConnectResult) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *QuickConnectResult) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *QuickConnectResult) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *QuickConnectResult) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickConnectResult) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *QuickConnectResult) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *QuickConnectResult) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *QuickConnectResult) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickConnectResult) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *QuickConnectResult) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *QuickConnectResult) SetAppName(v string) {
	o.AppName = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *QuickConnectResult) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion) {
		var ret string
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickConnectResult) GetAppVersionOk() (*string, bool) {
	if o == nil || IsNil(o.AppVersion) {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *QuickConnectResult) HasAppVersion() bool {
	if o != nil && !IsNil(o.AppVersion) {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given string and assigns it to the AppVersion field.
func (o *QuickConnectResult) SetAppVersion(v string) {
	o.AppVersion = &v
}

// GetDateAdded returns the DateAdded field value if set, zero value otherwise.
func (o *QuickConnectResult) GetDateAdded() time.Time {
	if o == nil || IsNil(o.DateAdded) {
		var ret time.Time
		return ret
	}
	return *o.DateAdded
}

// GetDateAddedOk returns a tuple with the DateAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickConnectResult) GetDateAddedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateAdded) {
		return nil, false
	}
	return o.DateAdded, true
}

// HasDateAdded returns a boolean if a field has been set.
func (o *QuickConnectResult) HasDateAdded() bool {
	if o != nil && !IsNil(o.DateAdded) {
		return true
	}

	return false
}

// SetDateAdded gets a reference to the given time.Time and assigns it to the DateAdded field.
func (o *QuickConnectResult) SetDateAdded(v time.Time) {
	o.DateAdded = &v
}

func (o QuickConnectResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickConnectResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authenticated) {
		toSerialize["Authenticated"] = o.Authenticated
	}
	if !IsNil(o.Secret) {
		toSerialize["Secret"] = o.Secret
	}
	if !IsNil(o.Code) {
		toSerialize["Code"] = o.Code
	}
	if !IsNil(o.DeviceId) {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if !IsNil(o.DeviceName) {
		toSerialize["DeviceName"] = o.DeviceName
	}
	if !IsNil(o.AppName) {
		toSerialize["AppName"] = o.AppName
	}
	if !IsNil(o.AppVersion) {
		toSerialize["AppVersion"] = o.AppVersion
	}
	if !IsNil(o.DateAdded) {
		toSerialize["DateAdded"] = o.DateAdded
	}
	return toSerialize, nil
}

type NullableQuickConnectResult struct {
	value *QuickConnectResult
	isSet bool
}

func (v NullableQuickConnectResult) Get() *QuickConnectResult {
	return v.value
}

func (v *NullableQuickConnectResult) Set(val *QuickConnectResult) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickConnectResult) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickConnectResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickConnectResult(val *QuickConnectResult) *NullableQuickConnectResult {
	return &NullableQuickConnectResult{value: val, isSet: true}
}

func (v NullableQuickConnectResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickConnectResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


