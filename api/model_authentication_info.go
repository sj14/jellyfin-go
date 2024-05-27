/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the AuthenticationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationInfo{}

// AuthenticationInfo struct for AuthenticationInfo
type AuthenticationInfo struct {
	// Gets or sets the identifier.
	Id *int64 `json:"Id,omitempty"`
	// Gets or sets the access token.
	AccessToken NullableString `json:"AccessToken,omitempty"`
	// Gets or sets the device identifier.
	DeviceId NullableString `json:"DeviceId,omitempty"`
	// Gets or sets the name of the application.
	AppName NullableString `json:"AppName,omitempty"`
	// Gets or sets the application version.
	AppVersion NullableString `json:"AppVersion,omitempty"`
	// Gets or sets the name of the device.
	DeviceName NullableString `json:"DeviceName,omitempty"`
	// Gets or sets the user identifier.
	UserId *string `json:"UserId,omitempty"`
	// Gets or sets a value indicating whether this instance is active.
	IsActive *bool `json:"IsActive,omitempty"`
	// Gets or sets the date created.
	DateCreated *time.Time `json:"DateCreated,omitempty"`
	// Gets or sets the date revoked.
	DateRevoked NullableTime `json:"DateRevoked,omitempty"`
	DateLastActivity *time.Time `json:"DateLastActivity,omitempty"`
	UserName NullableString `json:"UserName,omitempty"`
}

// NewAuthenticationInfo instantiates a new AuthenticationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationInfo() *AuthenticationInfo {
	this := AuthenticationInfo{}
	return &this
}

// NewAuthenticationInfoWithDefaults instantiates a new AuthenticationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationInfoWithDefaults() *AuthenticationInfo {
	this := AuthenticationInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AuthenticationInfo) SetId(v int64) {
	o.Id = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationInfo) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken.Get()) {
		var ret string
		return ret
	}
	return *o.AccessToken.Get()
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationInfo) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccessToken.Get(), o.AccessToken.IsSet()
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasAccessToken() bool {
	if o != nil && o.AccessToken.IsSet() {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given NullableString and assigns it to the AccessToken field.
func (o *AuthenticationInfo) SetAccessToken(v string) {
	o.AccessToken.Set(&v)
}
// SetAccessTokenNil sets the value for AccessToken to be an explicit nil
func (o *AuthenticationInfo) SetAccessTokenNil() {
	o.AccessToken.Set(nil)
}

// UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
func (o *AuthenticationInfo) UnsetAccessToken() {
	o.AccessToken.Unset()
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationInfo) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId.Get()) {
		var ret string
		return ret
	}
	return *o.DeviceId.Get()
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationInfo) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceId.Get(), o.DeviceId.IsSet()
}

// HasDeviceId returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasDeviceId() bool {
	if o != nil && o.DeviceId.IsSet() {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given NullableString and assigns it to the DeviceId field.
func (o *AuthenticationInfo) SetDeviceId(v string) {
	o.DeviceId.Set(&v)
}
// SetDeviceIdNil sets the value for DeviceId to be an explicit nil
func (o *AuthenticationInfo) SetDeviceIdNil() {
	o.DeviceId.Set(nil)
}

// UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
func (o *AuthenticationInfo) UnsetDeviceId() {
	o.DeviceId.Unset()
}

// GetAppName returns the AppName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationInfo) GetAppName() string {
	if o == nil || IsNil(o.AppName.Get()) {
		var ret string
		return ret
	}
	return *o.AppName.Get()
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationInfo) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppName.Get(), o.AppName.IsSet()
}

// HasAppName returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasAppName() bool {
	if o != nil && o.AppName.IsSet() {
		return true
	}

	return false
}

// SetAppName gets a reference to the given NullableString and assigns it to the AppName field.
func (o *AuthenticationInfo) SetAppName(v string) {
	o.AppName.Set(&v)
}
// SetAppNameNil sets the value for AppName to be an explicit nil
func (o *AuthenticationInfo) SetAppNameNil() {
	o.AppName.Set(nil)
}

// UnsetAppName ensures that no value is present for AppName, not even an explicit nil
func (o *AuthenticationInfo) UnsetAppName() {
	o.AppName.Unset()
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationInfo) GetAppVersion() string {
	if o == nil || IsNil(o.AppVersion.Get()) {
		var ret string
		return ret
	}
	return *o.AppVersion.Get()
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationInfo) GetAppVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppVersion.Get(), o.AppVersion.IsSet()
}

// HasAppVersion returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasAppVersion() bool {
	if o != nil && o.AppVersion.IsSet() {
		return true
	}

	return false
}

// SetAppVersion gets a reference to the given NullableString and assigns it to the AppVersion field.
func (o *AuthenticationInfo) SetAppVersion(v string) {
	o.AppVersion.Set(&v)
}
// SetAppVersionNil sets the value for AppVersion to be an explicit nil
func (o *AuthenticationInfo) SetAppVersionNil() {
	o.AppVersion.Set(nil)
}

// UnsetAppVersion ensures that no value is present for AppVersion, not even an explicit nil
func (o *AuthenticationInfo) UnsetAppVersion() {
	o.AppVersion.Unset()
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationInfo) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName.Get()) {
		var ret string
		return ret
	}
	return *o.DeviceName.Get()
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationInfo) GetDeviceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceName.Get(), o.DeviceName.IsSet()
}

// HasDeviceName returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasDeviceName() bool {
	if o != nil && o.DeviceName.IsSet() {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given NullableString and assigns it to the DeviceName field.
func (o *AuthenticationInfo) SetDeviceName(v string) {
	o.DeviceName.Set(&v)
}
// SetDeviceNameNil sets the value for DeviceName to be an explicit nil
func (o *AuthenticationInfo) SetDeviceNameNil() {
	o.DeviceName.Set(nil)
}

// UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
func (o *AuthenticationInfo) UnsetDeviceName() {
	o.DeviceName.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AuthenticationInfo) SetUserId(v string) {
	o.UserId = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *AuthenticationInfo) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *AuthenticationInfo) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetDateRevoked returns the DateRevoked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationInfo) GetDateRevoked() time.Time {
	if o == nil || IsNil(o.DateRevoked.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DateRevoked.Get()
}

// GetDateRevokedOk returns a tuple with the DateRevoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationInfo) GetDateRevokedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateRevoked.Get(), o.DateRevoked.IsSet()
}

// HasDateRevoked returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasDateRevoked() bool {
	if o != nil && o.DateRevoked.IsSet() {
		return true
	}

	return false
}

// SetDateRevoked gets a reference to the given NullableTime and assigns it to the DateRevoked field.
func (o *AuthenticationInfo) SetDateRevoked(v time.Time) {
	o.DateRevoked.Set(&v)
}
// SetDateRevokedNil sets the value for DateRevoked to be an explicit nil
func (o *AuthenticationInfo) SetDateRevokedNil() {
	o.DateRevoked.Set(nil)
}

// UnsetDateRevoked ensures that no value is present for DateRevoked, not even an explicit nil
func (o *AuthenticationInfo) UnsetDateRevoked() {
	o.DateRevoked.Unset()
}

// GetDateLastActivity returns the DateLastActivity field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetDateLastActivity() time.Time {
	if o == nil || IsNil(o.DateLastActivity) {
		var ret time.Time
		return ret
	}
	return *o.DateLastActivity
}

// GetDateLastActivityOk returns a tuple with the DateLastActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetDateLastActivityOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateLastActivity) {
		return nil, false
	}
	return o.DateLastActivity, true
}

// HasDateLastActivity returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasDateLastActivity() bool {
	if o != nil && !IsNil(o.DateLastActivity) {
		return true
	}

	return false
}

// SetDateLastActivity gets a reference to the given time.Time and assigns it to the DateLastActivity field.
func (o *AuthenticationInfo) SetDateLastActivity(v time.Time) {
	o.DateLastActivity = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationInfo) GetUserName() string {
	if o == nil || IsNil(o.UserName.Get()) {
		var ret string
		return ret
	}
	return *o.UserName.Get()
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationInfo) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserName.Get(), o.UserName.IsSet()
}

// HasUserName returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasUserName() bool {
	if o != nil && o.UserName.IsSet() {
		return true
	}

	return false
}

// SetUserName gets a reference to the given NullableString and assigns it to the UserName field.
func (o *AuthenticationInfo) SetUserName(v string) {
	o.UserName.Set(&v)
}
// SetUserNameNil sets the value for UserName to be an explicit nil
func (o *AuthenticationInfo) SetUserNameNil() {
	o.UserName.Set(nil)
}

// UnsetUserName ensures that no value is present for UserName, not even an explicit nil
func (o *AuthenticationInfo) UnsetUserName() {
	o.UserName.Unset()
}

func (o AuthenticationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.AccessToken.IsSet() {
		toSerialize["AccessToken"] = o.AccessToken.Get()
	}
	if o.DeviceId.IsSet() {
		toSerialize["DeviceId"] = o.DeviceId.Get()
	}
	if o.AppName.IsSet() {
		toSerialize["AppName"] = o.AppName.Get()
	}
	if o.AppVersion.IsSet() {
		toSerialize["AppVersion"] = o.AppVersion.Get()
	}
	if o.DeviceName.IsSet() {
		toSerialize["DeviceName"] = o.DeviceName.Get()
	}
	if !IsNil(o.UserId) {
		toSerialize["UserId"] = o.UserId
	}
	if !IsNil(o.IsActive) {
		toSerialize["IsActive"] = o.IsActive
	}
	if !IsNil(o.DateCreated) {
		toSerialize["DateCreated"] = o.DateCreated
	}
	if o.DateRevoked.IsSet() {
		toSerialize["DateRevoked"] = o.DateRevoked.Get()
	}
	if !IsNil(o.DateLastActivity) {
		toSerialize["DateLastActivity"] = o.DateLastActivity
	}
	if o.UserName.IsSet() {
		toSerialize["UserName"] = o.UserName.Get()
	}
	return toSerialize, nil
}

type NullableAuthenticationInfo struct {
	value *AuthenticationInfo
	isSet bool
}

func (v NullableAuthenticationInfo) Get() *AuthenticationInfo {
	return v.value
}

func (v *NullableAuthenticationInfo) Set(val *AuthenticationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationInfo(val *AuthenticationInfo) *NullableAuthenticationInfo {
	return &NullableAuthenticationInfo{value: val, isSet: true}
}

func (v NullableAuthenticationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


