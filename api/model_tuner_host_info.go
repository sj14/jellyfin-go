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

// checks if the TunerHostInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TunerHostInfo{}

// TunerHostInfo struct for TunerHostInfo
type TunerHostInfo struct {
	Id NullableString `json:"Id,omitempty"`
	Url NullableString `json:"Url,omitempty"`
	Type NullableString `json:"Type,omitempty"`
	DeviceId NullableString `json:"DeviceId,omitempty"`
	FriendlyName NullableString `json:"FriendlyName,omitempty"`
	ImportFavoritesOnly *bool `json:"ImportFavoritesOnly,omitempty"`
	AllowHWTranscoding *bool `json:"AllowHWTranscoding,omitempty"`
	EnableStreamLooping *bool `json:"EnableStreamLooping,omitempty"`
	Source NullableString `json:"Source,omitempty"`
	TunerCount *int32 `json:"TunerCount,omitempty"`
	UserAgent NullableString `json:"UserAgent,omitempty"`
	IgnoreDts *bool `json:"IgnoreDts,omitempty"`
}

// NewTunerHostInfo instantiates a new TunerHostInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunerHostInfo() *TunerHostInfo {
	this := TunerHostInfo{}
	return &this
}

// NewTunerHostInfoWithDefaults instantiates a new TunerHostInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunerHostInfoWithDefaults() *TunerHostInfo {
	this := TunerHostInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunerHostInfo) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunerHostInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *TunerHostInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *TunerHostInfo) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *TunerHostInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *TunerHostInfo) UnsetId() {
	o.Id.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunerHostInfo) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunerHostInfo) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *TunerHostInfo) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *TunerHostInfo) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *TunerHostInfo) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *TunerHostInfo) UnsetUrl() {
	o.Url.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunerHostInfo) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunerHostInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *TunerHostInfo) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *TunerHostInfo) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *TunerHostInfo) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *TunerHostInfo) UnsetType() {
	o.Type.Unset()
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunerHostInfo) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId.Get()) {
		var ret string
		return ret
	}
	return *o.DeviceId.Get()
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunerHostInfo) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceId.Get(), o.DeviceId.IsSet()
}

// HasDeviceId returns a boolean if a field has been set.
func (o *TunerHostInfo) HasDeviceId() bool {
	if o != nil && o.DeviceId.IsSet() {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given NullableString and assigns it to the DeviceId field.
func (o *TunerHostInfo) SetDeviceId(v string) {
	o.DeviceId.Set(&v)
}
// SetDeviceIdNil sets the value for DeviceId to be an explicit nil
func (o *TunerHostInfo) SetDeviceIdNil() {
	o.DeviceId.Set(nil)
}

// UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
func (o *TunerHostInfo) UnsetDeviceId() {
	o.DeviceId.Unset()
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunerHostInfo) GetFriendlyName() string {
	if o == nil || IsNil(o.FriendlyName.Get()) {
		var ret string
		return ret
	}
	return *o.FriendlyName.Get()
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunerHostInfo) GetFriendlyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FriendlyName.Get(), o.FriendlyName.IsSet()
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *TunerHostInfo) HasFriendlyName() bool {
	if o != nil && o.FriendlyName.IsSet() {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given NullableString and assigns it to the FriendlyName field.
func (o *TunerHostInfo) SetFriendlyName(v string) {
	o.FriendlyName.Set(&v)
}
// SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil
func (o *TunerHostInfo) SetFriendlyNameNil() {
	o.FriendlyName.Set(nil)
}

// UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
func (o *TunerHostInfo) UnsetFriendlyName() {
	o.FriendlyName.Unset()
}

// GetImportFavoritesOnly returns the ImportFavoritesOnly field value if set, zero value otherwise.
func (o *TunerHostInfo) GetImportFavoritesOnly() bool {
	if o == nil || IsNil(o.ImportFavoritesOnly) {
		var ret bool
		return ret
	}
	return *o.ImportFavoritesOnly
}

// GetImportFavoritesOnlyOk returns a tuple with the ImportFavoritesOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunerHostInfo) GetImportFavoritesOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ImportFavoritesOnly) {
		return nil, false
	}
	return o.ImportFavoritesOnly, true
}

// HasImportFavoritesOnly returns a boolean if a field has been set.
func (o *TunerHostInfo) HasImportFavoritesOnly() bool {
	if o != nil && !IsNil(o.ImportFavoritesOnly) {
		return true
	}

	return false
}

// SetImportFavoritesOnly gets a reference to the given bool and assigns it to the ImportFavoritesOnly field.
func (o *TunerHostInfo) SetImportFavoritesOnly(v bool) {
	o.ImportFavoritesOnly = &v
}

// GetAllowHWTranscoding returns the AllowHWTranscoding field value if set, zero value otherwise.
func (o *TunerHostInfo) GetAllowHWTranscoding() bool {
	if o == nil || IsNil(o.AllowHWTranscoding) {
		var ret bool
		return ret
	}
	return *o.AllowHWTranscoding
}

// GetAllowHWTranscodingOk returns a tuple with the AllowHWTranscoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunerHostInfo) GetAllowHWTranscodingOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowHWTranscoding) {
		return nil, false
	}
	return o.AllowHWTranscoding, true
}

// HasAllowHWTranscoding returns a boolean if a field has been set.
func (o *TunerHostInfo) HasAllowHWTranscoding() bool {
	if o != nil && !IsNil(o.AllowHWTranscoding) {
		return true
	}

	return false
}

// SetAllowHWTranscoding gets a reference to the given bool and assigns it to the AllowHWTranscoding field.
func (o *TunerHostInfo) SetAllowHWTranscoding(v bool) {
	o.AllowHWTranscoding = &v
}

// GetEnableStreamLooping returns the EnableStreamLooping field value if set, zero value otherwise.
func (o *TunerHostInfo) GetEnableStreamLooping() bool {
	if o == nil || IsNil(o.EnableStreamLooping) {
		var ret bool
		return ret
	}
	return *o.EnableStreamLooping
}

// GetEnableStreamLoopingOk returns a tuple with the EnableStreamLooping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunerHostInfo) GetEnableStreamLoopingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableStreamLooping) {
		return nil, false
	}
	return o.EnableStreamLooping, true
}

// HasEnableStreamLooping returns a boolean if a field has been set.
func (o *TunerHostInfo) HasEnableStreamLooping() bool {
	if o != nil && !IsNil(o.EnableStreamLooping) {
		return true
	}

	return false
}

// SetEnableStreamLooping gets a reference to the given bool and assigns it to the EnableStreamLooping field.
func (o *TunerHostInfo) SetEnableStreamLooping(v bool) {
	o.EnableStreamLooping = &v
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunerHostInfo) GetSource() string {
	if o == nil || IsNil(o.Source.Get()) {
		var ret string
		return ret
	}
	return *o.Source.Get()
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunerHostInfo) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Source.Get(), o.Source.IsSet()
}

// HasSource returns a boolean if a field has been set.
func (o *TunerHostInfo) HasSource() bool {
	if o != nil && o.Source.IsSet() {
		return true
	}

	return false
}

// SetSource gets a reference to the given NullableString and assigns it to the Source field.
func (o *TunerHostInfo) SetSource(v string) {
	o.Source.Set(&v)
}
// SetSourceNil sets the value for Source to be an explicit nil
func (o *TunerHostInfo) SetSourceNil() {
	o.Source.Set(nil)
}

// UnsetSource ensures that no value is present for Source, not even an explicit nil
func (o *TunerHostInfo) UnsetSource() {
	o.Source.Unset()
}

// GetTunerCount returns the TunerCount field value if set, zero value otherwise.
func (o *TunerHostInfo) GetTunerCount() int32 {
	if o == nil || IsNil(o.TunerCount) {
		var ret int32
		return ret
	}
	return *o.TunerCount
}

// GetTunerCountOk returns a tuple with the TunerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunerHostInfo) GetTunerCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TunerCount) {
		return nil, false
	}
	return o.TunerCount, true
}

// HasTunerCount returns a boolean if a field has been set.
func (o *TunerHostInfo) HasTunerCount() bool {
	if o != nil && !IsNil(o.TunerCount) {
		return true
	}

	return false
}

// SetTunerCount gets a reference to the given int32 and assigns it to the TunerCount field.
func (o *TunerHostInfo) SetTunerCount(v int32) {
	o.TunerCount = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TunerHostInfo) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent.Get()) {
		var ret string
		return ret
	}
	return *o.UserAgent.Get()
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TunerHostInfo) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserAgent.Get(), o.UserAgent.IsSet()
}

// HasUserAgent returns a boolean if a field has been set.
func (o *TunerHostInfo) HasUserAgent() bool {
	if o != nil && o.UserAgent.IsSet() {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given NullableString and assigns it to the UserAgent field.
func (o *TunerHostInfo) SetUserAgent(v string) {
	o.UserAgent.Set(&v)
}
// SetUserAgentNil sets the value for UserAgent to be an explicit nil
func (o *TunerHostInfo) SetUserAgentNil() {
	o.UserAgent.Set(nil)
}

// UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
func (o *TunerHostInfo) UnsetUserAgent() {
	o.UserAgent.Unset()
}

// GetIgnoreDts returns the IgnoreDts field value if set, zero value otherwise.
func (o *TunerHostInfo) GetIgnoreDts() bool {
	if o == nil || IsNil(o.IgnoreDts) {
		var ret bool
		return ret
	}
	return *o.IgnoreDts
}

// GetIgnoreDtsOk returns a tuple with the IgnoreDts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TunerHostInfo) GetIgnoreDtsOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreDts) {
		return nil, false
	}
	return o.IgnoreDts, true
}

// HasIgnoreDts returns a boolean if a field has been set.
func (o *TunerHostInfo) HasIgnoreDts() bool {
	if o != nil && !IsNil(o.IgnoreDts) {
		return true
	}

	return false
}

// SetIgnoreDts gets a reference to the given bool and assigns it to the IgnoreDts field.
func (o *TunerHostInfo) SetIgnoreDts(v bool) {
	o.IgnoreDts = &v
}

func (o TunerHostInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TunerHostInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.Url.IsSet() {
		toSerialize["Url"] = o.Url.Get()
	}
	if o.Type.IsSet() {
		toSerialize["Type"] = o.Type.Get()
	}
	if o.DeviceId.IsSet() {
		toSerialize["DeviceId"] = o.DeviceId.Get()
	}
	if o.FriendlyName.IsSet() {
		toSerialize["FriendlyName"] = o.FriendlyName.Get()
	}
	if !IsNil(o.ImportFavoritesOnly) {
		toSerialize["ImportFavoritesOnly"] = o.ImportFavoritesOnly
	}
	if !IsNil(o.AllowHWTranscoding) {
		toSerialize["AllowHWTranscoding"] = o.AllowHWTranscoding
	}
	if !IsNil(o.EnableStreamLooping) {
		toSerialize["EnableStreamLooping"] = o.EnableStreamLooping
	}
	if o.Source.IsSet() {
		toSerialize["Source"] = o.Source.Get()
	}
	if !IsNil(o.TunerCount) {
		toSerialize["TunerCount"] = o.TunerCount
	}
	if o.UserAgent.IsSet() {
		toSerialize["UserAgent"] = o.UserAgent.Get()
	}
	if !IsNil(o.IgnoreDts) {
		toSerialize["IgnoreDts"] = o.IgnoreDts
	}
	return toSerialize, nil
}

type NullableTunerHostInfo struct {
	value *TunerHostInfo
	isSet bool
}

func (v NullableTunerHostInfo) Get() *TunerHostInfo {
	return v.value
}

func (v *NullableTunerHostInfo) Set(val *TunerHostInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTunerHostInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTunerHostInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunerHostInfo(val *TunerHostInfo) *NullableTunerHostInfo {
	return &NullableTunerHostInfo{value: val, isSet: true}
}

func (v NullableTunerHostInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunerHostInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


