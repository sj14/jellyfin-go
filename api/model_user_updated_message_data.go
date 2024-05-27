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

// checks if the UserUpdatedMessageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUpdatedMessageData{}

// UserUpdatedMessageData Gets or sets the data.
type UserUpdatedMessageData struct {
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

// NewUserUpdatedMessageData instantiates a new UserUpdatedMessageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdatedMessageData() *UserUpdatedMessageData {
	this := UserUpdatedMessageData{}
	return &this
}

// NewUserUpdatedMessageDataWithDefaults instantiates a new UserUpdatedMessageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdatedMessageDataWithDefaults() *UserUpdatedMessageData {
	this := UserUpdatedMessageData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdatedMessageData) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdatedMessageData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *UserUpdatedMessageData) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *UserUpdatedMessageData) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *UserUpdatedMessageData) UnsetName() {
	o.Name.Unset()
}

// GetServerId returns the ServerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdatedMessageData) GetServerId() string {
	if o == nil || IsNil(o.ServerId.Get()) {
		var ret string
		return ret
	}
	return *o.ServerId.Get()
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdatedMessageData) GetServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerId.Get(), o.ServerId.IsSet()
}

// HasServerId returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasServerId() bool {
	if o != nil && o.ServerId.IsSet() {
		return true
	}

	return false
}

// SetServerId gets a reference to the given NullableString and assigns it to the ServerId field.
func (o *UserUpdatedMessageData) SetServerId(v string) {
	o.ServerId.Set(&v)
}
// SetServerIdNil sets the value for ServerId to be an explicit nil
func (o *UserUpdatedMessageData) SetServerIdNil() {
	o.ServerId.Set(nil)
}

// UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
func (o *UserUpdatedMessageData) UnsetServerId() {
	o.ServerId.Unset()
}

// GetServerName returns the ServerName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdatedMessageData) GetServerName() string {
	if o == nil || IsNil(o.ServerName.Get()) {
		var ret string
		return ret
	}
	return *o.ServerName.Get()
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdatedMessageData) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerName.Get(), o.ServerName.IsSet()
}

// HasServerName returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasServerName() bool {
	if o != nil && o.ServerName.IsSet() {
		return true
	}

	return false
}

// SetServerName gets a reference to the given NullableString and assigns it to the ServerName field.
func (o *UserUpdatedMessageData) SetServerName(v string) {
	o.ServerName.Set(&v)
}
// SetServerNameNil sets the value for ServerName to be an explicit nil
func (o *UserUpdatedMessageData) SetServerNameNil() {
	o.ServerName.Set(nil)
}

// UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
func (o *UserUpdatedMessageData) UnsetServerName() {
	o.ServerName.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserUpdatedMessageData) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdatedMessageData) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserUpdatedMessageData) SetId(v string) {
	o.Id = &v
}

// GetPrimaryImageTag returns the PrimaryImageTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdatedMessageData) GetPrimaryImageTag() string {
	if o == nil || IsNil(o.PrimaryImageTag.Get()) {
		var ret string
		return ret
	}
	return *o.PrimaryImageTag.Get()
}

// GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdatedMessageData) GetPrimaryImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryImageTag.Get(), o.PrimaryImageTag.IsSet()
}

// HasPrimaryImageTag returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasPrimaryImageTag() bool {
	if o != nil && o.PrimaryImageTag.IsSet() {
		return true
	}

	return false
}

// SetPrimaryImageTag gets a reference to the given NullableString and assigns it to the PrimaryImageTag field.
func (o *UserUpdatedMessageData) SetPrimaryImageTag(v string) {
	o.PrimaryImageTag.Set(&v)
}
// SetPrimaryImageTagNil sets the value for PrimaryImageTag to be an explicit nil
func (o *UserUpdatedMessageData) SetPrimaryImageTagNil() {
	o.PrimaryImageTag.Set(nil)
}

// UnsetPrimaryImageTag ensures that no value is present for PrimaryImageTag, not even an explicit nil
func (o *UserUpdatedMessageData) UnsetPrimaryImageTag() {
	o.PrimaryImageTag.Unset()
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *UserUpdatedMessageData) GetHasPassword() bool {
	if o == nil || IsNil(o.HasPassword) {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdatedMessageData) GetHasPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPassword) {
		return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasHasPassword() bool {
	if o != nil && !IsNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *UserUpdatedMessageData) SetHasPassword(v bool) {
	o.HasPassword = &v
}

// GetHasConfiguredPassword returns the HasConfiguredPassword field value if set, zero value otherwise.
func (o *UserUpdatedMessageData) GetHasConfiguredPassword() bool {
	if o == nil || IsNil(o.HasConfiguredPassword) {
		var ret bool
		return ret
	}
	return *o.HasConfiguredPassword
}

// GetHasConfiguredPasswordOk returns a tuple with the HasConfiguredPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdatedMessageData) GetHasConfiguredPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasConfiguredPassword) {
		return nil, false
	}
	return o.HasConfiguredPassword, true
}

// HasHasConfiguredPassword returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasHasConfiguredPassword() bool {
	if o != nil && !IsNil(o.HasConfiguredPassword) {
		return true
	}

	return false
}

// SetHasConfiguredPassword gets a reference to the given bool and assigns it to the HasConfiguredPassword field.
func (o *UserUpdatedMessageData) SetHasConfiguredPassword(v bool) {
	o.HasConfiguredPassword = &v
}

// GetHasConfiguredEasyPassword returns the HasConfiguredEasyPassword field value if set, zero value otherwise.
// Deprecated
func (o *UserUpdatedMessageData) GetHasConfiguredEasyPassword() bool {
	if o == nil || IsNil(o.HasConfiguredEasyPassword) {
		var ret bool
		return ret
	}
	return *o.HasConfiguredEasyPassword
}

// GetHasConfiguredEasyPasswordOk returns a tuple with the HasConfiguredEasyPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UserUpdatedMessageData) GetHasConfiguredEasyPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasConfiguredEasyPassword) {
		return nil, false
	}
	return o.HasConfiguredEasyPassword, true
}

// HasHasConfiguredEasyPassword returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasHasConfiguredEasyPassword() bool {
	if o != nil && !IsNil(o.HasConfiguredEasyPassword) {
		return true
	}

	return false
}

// SetHasConfiguredEasyPassword gets a reference to the given bool and assigns it to the HasConfiguredEasyPassword field.
// Deprecated
func (o *UserUpdatedMessageData) SetHasConfiguredEasyPassword(v bool) {
	o.HasConfiguredEasyPassword = &v
}

// GetEnableAutoLogin returns the EnableAutoLogin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdatedMessageData) GetEnableAutoLogin() bool {
	if o == nil || IsNil(o.EnableAutoLogin.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableAutoLogin.Get()
}

// GetEnableAutoLoginOk returns a tuple with the EnableAutoLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdatedMessageData) GetEnableAutoLoginOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableAutoLogin.Get(), o.EnableAutoLogin.IsSet()
}

// HasEnableAutoLogin returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasEnableAutoLogin() bool {
	if o != nil && o.EnableAutoLogin.IsSet() {
		return true
	}

	return false
}

// SetEnableAutoLogin gets a reference to the given NullableBool and assigns it to the EnableAutoLogin field.
func (o *UserUpdatedMessageData) SetEnableAutoLogin(v bool) {
	o.EnableAutoLogin.Set(&v)
}
// SetEnableAutoLoginNil sets the value for EnableAutoLogin to be an explicit nil
func (o *UserUpdatedMessageData) SetEnableAutoLoginNil() {
	o.EnableAutoLogin.Set(nil)
}

// UnsetEnableAutoLogin ensures that no value is present for EnableAutoLogin, not even an explicit nil
func (o *UserUpdatedMessageData) UnsetEnableAutoLogin() {
	o.EnableAutoLogin.Unset()
}

// GetLastLoginDate returns the LastLoginDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdatedMessageData) GetLastLoginDate() time.Time {
	if o == nil || IsNil(o.LastLoginDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastLoginDate.Get()
}

// GetLastLoginDateOk returns a tuple with the LastLoginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdatedMessageData) GetLastLoginDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastLoginDate.Get(), o.LastLoginDate.IsSet()
}

// HasLastLoginDate returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasLastLoginDate() bool {
	if o != nil && o.LastLoginDate.IsSet() {
		return true
	}

	return false
}

// SetLastLoginDate gets a reference to the given NullableTime and assigns it to the LastLoginDate field.
func (o *UserUpdatedMessageData) SetLastLoginDate(v time.Time) {
	o.LastLoginDate.Set(&v)
}
// SetLastLoginDateNil sets the value for LastLoginDate to be an explicit nil
func (o *UserUpdatedMessageData) SetLastLoginDateNil() {
	o.LastLoginDate.Set(nil)
}

// UnsetLastLoginDate ensures that no value is present for LastLoginDate, not even an explicit nil
func (o *UserUpdatedMessageData) UnsetLastLoginDate() {
	o.LastLoginDate.Unset()
}

// GetLastActivityDate returns the LastActivityDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdatedMessageData) GetLastActivityDate() time.Time {
	if o == nil || IsNil(o.LastActivityDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastActivityDate.Get()
}

// GetLastActivityDateOk returns a tuple with the LastActivityDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdatedMessageData) GetLastActivityDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastActivityDate.Get(), o.LastActivityDate.IsSet()
}

// HasLastActivityDate returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasLastActivityDate() bool {
	if o != nil && o.LastActivityDate.IsSet() {
		return true
	}

	return false
}

// SetLastActivityDate gets a reference to the given NullableTime and assigns it to the LastActivityDate field.
func (o *UserUpdatedMessageData) SetLastActivityDate(v time.Time) {
	o.LastActivityDate.Set(&v)
}
// SetLastActivityDateNil sets the value for LastActivityDate to be an explicit nil
func (o *UserUpdatedMessageData) SetLastActivityDateNil() {
	o.LastActivityDate.Set(nil)
}

// UnsetLastActivityDate ensures that no value is present for LastActivityDate, not even an explicit nil
func (o *UserUpdatedMessageData) UnsetLastActivityDate() {
	o.LastActivityDate.Unset()
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdatedMessageData) GetConfiguration() UserDtoConfiguration {
	if o == nil || IsNil(o.Configuration.Get()) {
		var ret UserDtoConfiguration
		return ret
	}
	return *o.Configuration.Get()
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdatedMessageData) GetConfigurationOk() (*UserDtoConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Configuration.Get(), o.Configuration.IsSet()
}

// HasConfiguration returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasConfiguration() bool {
	if o != nil && o.Configuration.IsSet() {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given NullableUserDtoConfiguration and assigns it to the Configuration field.
func (o *UserUpdatedMessageData) SetConfiguration(v UserDtoConfiguration) {
	o.Configuration.Set(&v)
}
// SetConfigurationNil sets the value for Configuration to be an explicit nil
func (o *UserUpdatedMessageData) SetConfigurationNil() {
	o.Configuration.Set(nil)
}

// UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
func (o *UserUpdatedMessageData) UnsetConfiguration() {
	o.Configuration.Unset()
}

// GetPolicy returns the Policy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdatedMessageData) GetPolicy() UserDtoPolicy {
	if o == nil || IsNil(o.Policy.Get()) {
		var ret UserDtoPolicy
		return ret
	}
	return *o.Policy.Get()
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdatedMessageData) GetPolicyOk() (*UserDtoPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policy.Get(), o.Policy.IsSet()
}

// HasPolicy returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasPolicy() bool {
	if o != nil && o.Policy.IsSet() {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given NullableUserDtoPolicy and assigns it to the Policy field.
func (o *UserUpdatedMessageData) SetPolicy(v UserDtoPolicy) {
	o.Policy.Set(&v)
}
// SetPolicyNil sets the value for Policy to be an explicit nil
func (o *UserUpdatedMessageData) SetPolicyNil() {
	o.Policy.Set(nil)
}

// UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
func (o *UserUpdatedMessageData) UnsetPolicy() {
	o.Policy.Unset()
}

// GetPrimaryImageAspectRatio returns the PrimaryImageAspectRatio field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserUpdatedMessageData) GetPrimaryImageAspectRatio() float64 {
	if o == nil || IsNil(o.PrimaryImageAspectRatio.Get()) {
		var ret float64
		return ret
	}
	return *o.PrimaryImageAspectRatio.Get()
}

// GetPrimaryImageAspectRatioOk returns a tuple with the PrimaryImageAspectRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserUpdatedMessageData) GetPrimaryImageAspectRatioOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryImageAspectRatio.Get(), o.PrimaryImageAspectRatio.IsSet()
}

// HasPrimaryImageAspectRatio returns a boolean if a field has been set.
func (o *UserUpdatedMessageData) HasPrimaryImageAspectRatio() bool {
	if o != nil && o.PrimaryImageAspectRatio.IsSet() {
		return true
	}

	return false
}

// SetPrimaryImageAspectRatio gets a reference to the given NullableFloat64 and assigns it to the PrimaryImageAspectRatio field.
func (o *UserUpdatedMessageData) SetPrimaryImageAspectRatio(v float64) {
	o.PrimaryImageAspectRatio.Set(&v)
}
// SetPrimaryImageAspectRatioNil sets the value for PrimaryImageAspectRatio to be an explicit nil
func (o *UserUpdatedMessageData) SetPrimaryImageAspectRatioNil() {
	o.PrimaryImageAspectRatio.Set(nil)
}

// UnsetPrimaryImageAspectRatio ensures that no value is present for PrimaryImageAspectRatio, not even an explicit nil
func (o *UserUpdatedMessageData) UnsetPrimaryImageAspectRatio() {
	o.PrimaryImageAspectRatio.Unset()
}

func (o UserUpdatedMessageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUpdatedMessageData) ToMap() (map[string]interface{}, error) {
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

type NullableUserUpdatedMessageData struct {
	value *UserUpdatedMessageData
	isSet bool
}

func (v NullableUserUpdatedMessageData) Get() *UserUpdatedMessageData {
	return v.value
}

func (v *NullableUserUpdatedMessageData) Set(val *UserUpdatedMessageData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdatedMessageData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdatedMessageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdatedMessageData(val *UserUpdatedMessageData) *NullableUserUpdatedMessageData {
	return &NullableUserUpdatedMessageData{value: val, isSet: true}
}

func (v NullableUserUpdatedMessageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdatedMessageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

