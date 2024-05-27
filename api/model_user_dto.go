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

// checks if the UserDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDto{}

// UserDto Class UserDto.
type UserDto struct {
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the server identifier.
	ServerId NullableString `json:"ServerId,omitempty"`
	// Gets or sets the name of the server.  This is not used by the server and is for client-side usage only.
	ServerName NullableString `json:"ServerName,omitempty"`
	// Gets or sets the id.
	Id *string `json:"Id,omitempty"`
	// Gets or sets the primary image tag.
	PrimaryImageTag NullableString `json:"PrimaryImageTag,omitempty"`
	// Gets or sets a value indicating whether this instance has password.
	HasPassword *bool `json:"HasPassword,omitempty"`
	// Gets or sets a value indicating whether this instance has configured password.
	HasConfiguredPassword *bool `json:"HasConfiguredPassword,omitempty"`
	// Gets or sets a value indicating whether this instance has configured easy password.
	// Deprecated
	HasConfiguredEasyPassword *bool `json:"HasConfiguredEasyPassword,omitempty"`
	// Gets or sets whether async login is enabled or not.
	EnableAutoLogin NullableBool `json:"EnableAutoLogin,omitempty"`
	// Gets or sets the last login date.
	LastLoginDate NullableTime `json:"LastLoginDate,omitempty"`
	// Gets or sets the last activity date.
	LastActivityDate NullableTime `json:"LastActivityDate,omitempty"`
	Configuration NullableUserDtoConfiguration `json:"Configuration,omitempty"`
	Policy NullableUserDtoPolicy `json:"Policy,omitempty"`
	// Gets or sets the primary image aspect ratio.
	PrimaryImageAspectRatio NullableFloat64 `json:"PrimaryImageAspectRatio,omitempty"`
}

// NewUserDto instantiates a new UserDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDto() *UserDto {
	this := UserDto{}
	return &this
}

// NewUserDtoWithDefaults instantiates a new UserDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDtoWithDefaults() *UserDto {
	this := UserDto{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDto) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UserDto) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UserDto) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UserDto) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UserDto) UnsetName() {
	o.Name.Unset()
}

// GetServerId returns the ServerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDto) GetServerId() string {
	if o == nil || IsNil(o.ServerId.Get()) {
		var ret string
		return ret
	}
	return *o.ServerId.Get()
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDto) GetServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerId.Get(), o.ServerId.IsSet()
}

// HasServerId returns a boolean if a field has been set.
func (o *UserDto) HasServerId() bool {
	if o != nil && o.ServerId.IsSet() {
		return true
	}

	return false
}

// SetServerId gets a reference to the given NullableString and assigns it to the ServerId field.
func (o *UserDto) SetServerId(v string) {
	o.ServerId.Set(&v)
}
// SetServerIdNil sets the value for ServerId to be an explicit nil
func (o *UserDto) SetServerIdNil() {
	o.ServerId.Set(nil)
}

// UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
func (o *UserDto) UnsetServerId() {
	o.ServerId.Unset()
}

// GetServerName returns the ServerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDto) GetServerName() string {
	if o == nil || IsNil(o.ServerName.Get()) {
		var ret string
		return ret
	}
	return *o.ServerName.Get()
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDto) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerName.Get(), o.ServerName.IsSet()
}

// HasServerName returns a boolean if a field has been set.
func (o *UserDto) HasServerName() bool {
	if o != nil && o.ServerName.IsSet() {
		return true
	}

	return false
}

// SetServerName gets a reference to the given NullableString and assigns it to the ServerName field.
func (o *UserDto) SetServerName(v string) {
	o.ServerName.Set(&v)
}
// SetServerNameNil sets the value for ServerName to be an explicit nil
func (o *UserDto) SetServerNameNil() {
	o.ServerName.Set(nil)
}

// UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
func (o *UserDto) UnsetServerName() {
	o.ServerName.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserDto) SetId(v string) {
	o.Id = &v
}

// GetPrimaryImageTag returns the PrimaryImageTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDto) GetPrimaryImageTag() string {
	if o == nil || IsNil(o.PrimaryImageTag.Get()) {
		var ret string
		return ret
	}
	return *o.PrimaryImageTag.Get()
}

// GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDto) GetPrimaryImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryImageTag.Get(), o.PrimaryImageTag.IsSet()
}

// HasPrimaryImageTag returns a boolean if a field has been set.
func (o *UserDto) HasPrimaryImageTag() bool {
	if o != nil && o.PrimaryImageTag.IsSet() {
		return true
	}

	return false
}

// SetPrimaryImageTag gets a reference to the given NullableString and assigns it to the PrimaryImageTag field.
func (o *UserDto) SetPrimaryImageTag(v string) {
	o.PrimaryImageTag.Set(&v)
}
// SetPrimaryImageTagNil sets the value for PrimaryImageTag to be an explicit nil
func (o *UserDto) SetPrimaryImageTagNil() {
	o.PrimaryImageTag.Set(nil)
}

// UnsetPrimaryImageTag ensures that no value is present for PrimaryImageTag, not even an explicit nil
func (o *UserDto) UnsetPrimaryImageTag() {
	o.PrimaryImageTag.Unset()
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *UserDto) GetHasPassword() bool {
	if o == nil || IsNil(o.HasPassword) {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDto) GetHasPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPassword) {
		return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *UserDto) HasHasPassword() bool {
	if o != nil && !IsNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *UserDto) SetHasPassword(v bool) {
	o.HasPassword = &v
}

// GetHasConfiguredPassword returns the HasConfiguredPassword field value if set, zero value otherwise.
func (o *UserDto) GetHasConfiguredPassword() bool {
	if o == nil || IsNil(o.HasConfiguredPassword) {
		var ret bool
		return ret
	}
	return *o.HasConfiguredPassword
}

// GetHasConfiguredPasswordOk returns a tuple with the HasConfiguredPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDto) GetHasConfiguredPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasConfiguredPassword) {
		return nil, false
	}
	return o.HasConfiguredPassword, true
}

// HasHasConfiguredPassword returns a boolean if a field has been set.
func (o *UserDto) HasHasConfiguredPassword() bool {
	if o != nil && !IsNil(o.HasConfiguredPassword) {
		return true
	}

	return false
}

// SetHasConfiguredPassword gets a reference to the given bool and assigns it to the HasConfiguredPassword field.
func (o *UserDto) SetHasConfiguredPassword(v bool) {
	o.HasConfiguredPassword = &v
}

// GetHasConfiguredEasyPassword returns the HasConfiguredEasyPassword field value if set, zero value otherwise.
// Deprecated
func (o *UserDto) GetHasConfiguredEasyPassword() bool {
	if o == nil || IsNil(o.HasConfiguredEasyPassword) {
		var ret bool
		return ret
	}
	return *o.HasConfiguredEasyPassword
}

// GetHasConfiguredEasyPasswordOk returns a tuple with the HasConfiguredEasyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UserDto) GetHasConfiguredEasyPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasConfiguredEasyPassword) {
		return nil, false
	}
	return o.HasConfiguredEasyPassword, true
}

// HasHasConfiguredEasyPassword returns a boolean if a field has been set.
func (o *UserDto) HasHasConfiguredEasyPassword() bool {
	if o != nil && !IsNil(o.HasConfiguredEasyPassword) {
		return true
	}

	return false
}

// SetHasConfiguredEasyPassword gets a reference to the given bool and assigns it to the HasConfiguredEasyPassword field.
// Deprecated
func (o *UserDto) SetHasConfiguredEasyPassword(v bool) {
	o.HasConfiguredEasyPassword = &v
}

// GetEnableAutoLogin returns the EnableAutoLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDto) GetEnableAutoLogin() bool {
	if o == nil || IsNil(o.EnableAutoLogin.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableAutoLogin.Get()
}

// GetEnableAutoLoginOk returns a tuple with the EnableAutoLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDto) GetEnableAutoLoginOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableAutoLogin.Get(), o.EnableAutoLogin.IsSet()
}

// HasEnableAutoLogin returns a boolean if a field has been set.
func (o *UserDto) HasEnableAutoLogin() bool {
	if o != nil && o.EnableAutoLogin.IsSet() {
		return true
	}

	return false
}

// SetEnableAutoLogin gets a reference to the given NullableBool and assigns it to the EnableAutoLogin field.
func (o *UserDto) SetEnableAutoLogin(v bool) {
	o.EnableAutoLogin.Set(&v)
}
// SetEnableAutoLoginNil sets the value for EnableAutoLogin to be an explicit nil
func (o *UserDto) SetEnableAutoLoginNil() {
	o.EnableAutoLogin.Set(nil)
}

// UnsetEnableAutoLogin ensures that no value is present for EnableAutoLogin, not even an explicit nil
func (o *UserDto) UnsetEnableAutoLogin() {
	o.EnableAutoLogin.Unset()
}

// GetLastLoginDate returns the LastLoginDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDto) GetLastLoginDate() time.Time {
	if o == nil || IsNil(o.LastLoginDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastLoginDate.Get()
}

// GetLastLoginDateOk returns a tuple with the LastLoginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDto) GetLastLoginDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLoginDate.Get(), o.LastLoginDate.IsSet()
}

// HasLastLoginDate returns a boolean if a field has been set.
func (o *UserDto) HasLastLoginDate() bool {
	if o != nil && o.LastLoginDate.IsSet() {
		return true
	}

	return false
}

// SetLastLoginDate gets a reference to the given NullableTime and assigns it to the LastLoginDate field.
func (o *UserDto) SetLastLoginDate(v time.Time) {
	o.LastLoginDate.Set(&v)
}
// SetLastLoginDateNil sets the value for LastLoginDate to be an explicit nil
func (o *UserDto) SetLastLoginDateNil() {
	o.LastLoginDate.Set(nil)
}

// UnsetLastLoginDate ensures that no value is present for LastLoginDate, not even an explicit nil
func (o *UserDto) UnsetLastLoginDate() {
	o.LastLoginDate.Unset()
}

// GetLastActivityDate returns the LastActivityDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDto) GetLastActivityDate() time.Time {
	if o == nil || IsNil(o.LastActivityDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastActivityDate.Get()
}

// GetLastActivityDateOk returns a tuple with the LastActivityDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDto) GetLastActivityDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastActivityDate.Get(), o.LastActivityDate.IsSet()
}

// HasLastActivityDate returns a boolean if a field has been set.
func (o *UserDto) HasLastActivityDate() bool {
	if o != nil && o.LastActivityDate.IsSet() {
		return true
	}

	return false
}

// SetLastActivityDate gets a reference to the given NullableTime and assigns it to the LastActivityDate field.
func (o *UserDto) SetLastActivityDate(v time.Time) {
	o.LastActivityDate.Set(&v)
}
// SetLastActivityDateNil sets the value for LastActivityDate to be an explicit nil
func (o *UserDto) SetLastActivityDateNil() {
	o.LastActivityDate.Set(nil)
}

// UnsetLastActivityDate ensures that no value is present for LastActivityDate, not even an explicit nil
func (o *UserDto) UnsetLastActivityDate() {
	o.LastActivityDate.Unset()
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDto) GetConfiguration() UserDtoConfiguration {
	if o == nil || IsNil(o.Configuration.Get()) {
		var ret UserDtoConfiguration
		return ret
	}
	return *o.Configuration.Get()
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDto) GetConfigurationOk() (*UserDtoConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Configuration.Get(), o.Configuration.IsSet()
}

// HasConfiguration returns a boolean if a field has been set.
func (o *UserDto) HasConfiguration() bool {
	if o != nil && o.Configuration.IsSet() {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given NullableUserDtoConfiguration and assigns it to the Configuration field.
func (o *UserDto) SetConfiguration(v UserDtoConfiguration) {
	o.Configuration.Set(&v)
}
// SetConfigurationNil sets the value for Configuration to be an explicit nil
func (o *UserDto) SetConfigurationNil() {
	o.Configuration.Set(nil)
}

// UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
func (o *UserDto) UnsetConfiguration() {
	o.Configuration.Unset()
}

// GetPolicy returns the Policy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDto) GetPolicy() UserDtoPolicy {
	if o == nil || IsNil(o.Policy.Get()) {
		var ret UserDtoPolicy
		return ret
	}
	return *o.Policy.Get()
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDto) GetPolicyOk() (*UserDtoPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policy.Get(), o.Policy.IsSet()
}

// HasPolicy returns a boolean if a field has been set.
func (o *UserDto) HasPolicy() bool {
	if o != nil && o.Policy.IsSet() {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given NullableUserDtoPolicy and assigns it to the Policy field.
func (o *UserDto) SetPolicy(v UserDtoPolicy) {
	o.Policy.Set(&v)
}
// SetPolicyNil sets the value for Policy to be an explicit nil
func (o *UserDto) SetPolicyNil() {
	o.Policy.Set(nil)
}

// UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
func (o *UserDto) UnsetPolicy() {
	o.Policy.Unset()
}

// GetPrimaryImageAspectRatio returns the PrimaryImageAspectRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDto) GetPrimaryImageAspectRatio() float64 {
	if o == nil || IsNil(o.PrimaryImageAspectRatio.Get()) {
		var ret float64
		return ret
	}
	return *o.PrimaryImageAspectRatio.Get()
}

// GetPrimaryImageAspectRatioOk returns a tuple with the PrimaryImageAspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDto) GetPrimaryImageAspectRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryImageAspectRatio.Get(), o.PrimaryImageAspectRatio.IsSet()
}

// HasPrimaryImageAspectRatio returns a boolean if a field has been set.
func (o *UserDto) HasPrimaryImageAspectRatio() bool {
	if o != nil && o.PrimaryImageAspectRatio.IsSet() {
		return true
	}

	return false
}

// SetPrimaryImageAspectRatio gets a reference to the given NullableFloat64 and assigns it to the PrimaryImageAspectRatio field.
func (o *UserDto) SetPrimaryImageAspectRatio(v float64) {
	o.PrimaryImageAspectRatio.Set(&v)
}
// SetPrimaryImageAspectRatioNil sets the value for PrimaryImageAspectRatio to be an explicit nil
func (o *UserDto) SetPrimaryImageAspectRatioNil() {
	o.PrimaryImageAspectRatio.Set(nil)
}

// UnsetPrimaryImageAspectRatio ensures that no value is present for PrimaryImageAspectRatio, not even an explicit nil
func (o *UserDto) UnsetPrimaryImageAspectRatio() {
	o.PrimaryImageAspectRatio.Unset()
}

func (o UserDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.ServerId.IsSet() {
		toSerialize["ServerId"] = o.ServerId.Get()
	}
	if o.ServerName.IsSet() {
		toSerialize["ServerName"] = o.ServerName.Get()
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if o.PrimaryImageTag.IsSet() {
		toSerialize["PrimaryImageTag"] = o.PrimaryImageTag.Get()
	}
	if !IsNil(o.HasPassword) {
		toSerialize["HasPassword"] = o.HasPassword
	}
	if !IsNil(o.HasConfiguredPassword) {
		toSerialize["HasConfiguredPassword"] = o.HasConfiguredPassword
	}
	if !IsNil(o.HasConfiguredEasyPassword) {
		toSerialize["HasConfiguredEasyPassword"] = o.HasConfiguredEasyPassword
	}
	if o.EnableAutoLogin.IsSet() {
		toSerialize["EnableAutoLogin"] = o.EnableAutoLogin.Get()
	}
	if o.LastLoginDate.IsSet() {
		toSerialize["LastLoginDate"] = o.LastLoginDate.Get()
	}
	if o.LastActivityDate.IsSet() {
		toSerialize["LastActivityDate"] = o.LastActivityDate.Get()
	}
	if o.Configuration.IsSet() {
		toSerialize["Configuration"] = o.Configuration.Get()
	}
	if o.Policy.IsSet() {
		toSerialize["Policy"] = o.Policy.Get()
	}
	if o.PrimaryImageAspectRatio.IsSet() {
		toSerialize["PrimaryImageAspectRatio"] = o.PrimaryImageAspectRatio.Get()
	}
	return toSerialize, nil
}

type NullableUserDto struct {
	value *UserDto
	isSet bool
}

func (v NullableUserDto) Get() *UserDto {
	return v.value
}

func (v *NullableUserDto) Set(val *UserDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDto(val *UserDto) *NullableUserDto {
	return &NullableUserDto{value: val, isSet: true}
}

func (v NullableUserDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


