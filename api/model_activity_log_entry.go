/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the ActivityLogEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityLogEntry{}

// ActivityLogEntry An activity log entry.
type ActivityLogEntry struct {
	// Gets or sets the identifier.
	Id *int64 `json:"Id,omitempty"`
	// Gets or sets the name.
	Name *string `json:"Name,omitempty"`
	// Gets or sets the overview.
	Overview NullableString `json:"Overview,omitempty"`
	// Gets or sets the short overview.
	ShortOverview NullableString `json:"ShortOverview,omitempty"`
	// Gets or sets the type.
	Type *string `json:"Type,omitempty"`
	// Gets or sets the item identifier.
	ItemId NullableString `json:"ItemId,omitempty"`
	// Gets or sets the date.
	Date *time.Time `json:"Date,omitempty"`
	// Gets or sets the user identifier.
	UserId *string `json:"UserId,omitempty"`
	// Gets or sets the user primary image tag.
	// Deprecated
	UserPrimaryImageTag NullableString `json:"UserPrimaryImageTag,omitempty"`
	// Gets or sets the log severity.
	Severity *LogLevel `json:"Severity,omitempty"`
}

// NewActivityLogEntry instantiates a new ActivityLogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityLogEntry() *ActivityLogEntry {
	this := ActivityLogEntry{}
	return &this
}

// NewActivityLogEntryWithDefaults instantiates a new ActivityLogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityLogEntryWithDefaults() *ActivityLogEntry {
	this := ActivityLogEntry{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ActivityLogEntry) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLogEntry) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ActivityLogEntry) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ActivityLogEntry) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ActivityLogEntry) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLogEntry) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ActivityLogEntry) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ActivityLogEntry) SetName(v string) {
	o.Name = &v
}

// GetOverview returns the Overview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActivityLogEntry) GetOverview() string {
	if o == nil || IsNil(o.Overview.Get()) {
		var ret string
		return ret
	}
	return *o.Overview.Get()
}

// GetOverviewOk returns a tuple with the Overview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActivityLogEntry) GetOverviewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Overview.Get(), o.Overview.IsSet()
}

// HasOverview returns a boolean if a field has been set.
func (o *ActivityLogEntry) HasOverview() bool {
	if o != nil && o.Overview.IsSet() {
		return true
	}

	return false
}

// SetOverview gets a reference to the given NullableString and assigns it to the Overview field.
func (o *ActivityLogEntry) SetOverview(v string) {
	o.Overview.Set(&v)
}
// SetOverviewNil sets the value for Overview to be an explicit nil
func (o *ActivityLogEntry) SetOverviewNil() {
	o.Overview.Set(nil)
}

// UnsetOverview ensures that no value is present for Overview, not even an explicit nil
func (o *ActivityLogEntry) UnsetOverview() {
	o.Overview.Unset()
}

// GetShortOverview returns the ShortOverview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActivityLogEntry) GetShortOverview() string {
	if o == nil || IsNil(o.ShortOverview.Get()) {
		var ret string
		return ret
	}
	return *o.ShortOverview.Get()
}

// GetShortOverviewOk returns a tuple with the ShortOverview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActivityLogEntry) GetShortOverviewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortOverview.Get(), o.ShortOverview.IsSet()
}

// HasShortOverview returns a boolean if a field has been set.
func (o *ActivityLogEntry) HasShortOverview() bool {
	if o != nil && o.ShortOverview.IsSet() {
		return true
	}

	return false
}

// SetShortOverview gets a reference to the given NullableString and assigns it to the ShortOverview field.
func (o *ActivityLogEntry) SetShortOverview(v string) {
	o.ShortOverview.Set(&v)
}
// SetShortOverviewNil sets the value for ShortOverview to be an explicit nil
func (o *ActivityLogEntry) SetShortOverviewNil() {
	o.ShortOverview.Set(nil)
}

// UnsetShortOverview ensures that no value is present for ShortOverview, not even an explicit nil
func (o *ActivityLogEntry) UnsetShortOverview() {
	o.ShortOverview.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ActivityLogEntry) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLogEntry) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ActivityLogEntry) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ActivityLogEntry) SetType(v string) {
	o.Type = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActivityLogEntry) GetItemId() string {
	if o == nil || IsNil(o.ItemId.Get()) {
		var ret string
		return ret
	}
	return *o.ItemId.Get()
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActivityLogEntry) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemId.Get(), o.ItemId.IsSet()
}

// HasItemId returns a boolean if a field has been set.
func (o *ActivityLogEntry) HasItemId() bool {
	if o != nil && o.ItemId.IsSet() {
		return true
	}

	return false
}

// SetItemId gets a reference to the given NullableString and assigns it to the ItemId field.
func (o *ActivityLogEntry) SetItemId(v string) {
	o.ItemId.Set(&v)
}
// SetItemIdNil sets the value for ItemId to be an explicit nil
func (o *ActivityLogEntry) SetItemIdNil() {
	o.ItemId.Set(nil)
}

// UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
func (o *ActivityLogEntry) UnsetItemId() {
	o.ItemId.Unset()
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *ActivityLogEntry) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLogEntry) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *ActivityLogEntry) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *ActivityLogEntry) SetDate(v time.Time) {
	o.Date = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ActivityLogEntry) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLogEntry) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ActivityLogEntry) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ActivityLogEntry) SetUserId(v string) {
	o.UserId = &v
}

// GetUserPrimaryImageTag returns the UserPrimaryImageTag field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *ActivityLogEntry) GetUserPrimaryImageTag() string {
	if o == nil || IsNil(o.UserPrimaryImageTag.Get()) {
		var ret string
		return ret
	}
	return *o.UserPrimaryImageTag.Get()
}

// GetUserPrimaryImageTagOk returns a tuple with the UserPrimaryImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *ActivityLogEntry) GetUserPrimaryImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserPrimaryImageTag.Get(), o.UserPrimaryImageTag.IsSet()
}

// HasUserPrimaryImageTag returns a boolean if a field has been set.
func (o *ActivityLogEntry) HasUserPrimaryImageTag() bool {
	if o != nil && o.UserPrimaryImageTag.IsSet() {
		return true
	}

	return false
}

// SetUserPrimaryImageTag gets a reference to the given NullableString and assigns it to the UserPrimaryImageTag field.
// Deprecated
func (o *ActivityLogEntry) SetUserPrimaryImageTag(v string) {
	o.UserPrimaryImageTag.Set(&v)
}
// SetUserPrimaryImageTagNil sets the value for UserPrimaryImageTag to be an explicit nil
func (o *ActivityLogEntry) SetUserPrimaryImageTagNil() {
	o.UserPrimaryImageTag.Set(nil)
}

// UnsetUserPrimaryImageTag ensures that no value is present for UserPrimaryImageTag, not even an explicit nil
func (o *ActivityLogEntry) UnsetUserPrimaryImageTag() {
	o.UserPrimaryImageTag.Unset()
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ActivityLogEntry) GetSeverity() LogLevel {
	if o == nil || IsNil(o.Severity) {
		var ret LogLevel
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityLogEntry) GetSeverityOk() (*LogLevel, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ActivityLogEntry) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given LogLevel and assigns it to the Severity field.
func (o *ActivityLogEntry) SetSeverity(v LogLevel) {
	o.Severity = &v
}

func (o ActivityLogEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityLogEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if o.Overview.IsSet() {
		toSerialize["Overview"] = o.Overview.Get()
	}
	if o.ShortOverview.IsSet() {
		toSerialize["ShortOverview"] = o.ShortOverview.Get()
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if o.ItemId.IsSet() {
		toSerialize["ItemId"] = o.ItemId.Get()
	}
	if !IsNil(o.Date) {
		toSerialize["Date"] = o.Date
	}
	if !IsNil(o.UserId) {
		toSerialize["UserId"] = o.UserId
	}
	if o.UserPrimaryImageTag.IsSet() {
		toSerialize["UserPrimaryImageTag"] = o.UserPrimaryImageTag.Get()
	}
	if !IsNil(o.Severity) {
		toSerialize["Severity"] = o.Severity
	}
	return toSerialize, nil
}

type NullableActivityLogEntry struct {
	value *ActivityLogEntry
	isSet bool
}

func (v NullableActivityLogEntry) Get() *ActivityLogEntry {
	return v.value
}

func (v *NullableActivityLogEntry) Set(val *ActivityLogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityLogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityLogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityLogEntry(val *ActivityLogEntry) *NullableActivityLogEntry {
	return &NullableActivityLogEntry{value: val, isSet: true}
}

func (v NullableActivityLogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityLogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


