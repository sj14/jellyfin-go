/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.8.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the BaseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseItem{}

// BaseItem Class BaseItem.
type BaseItem struct {
	Size NullableInt64 `json:"Size,omitempty"`
	Container NullableString `json:"Container,omitempty"`
	IsHD *bool `json:"IsHD,omitempty"`
	IsShortcut *bool `json:"IsShortcut,omitempty"`
	ShortcutPath NullableString `json:"ShortcutPath,omitempty"`
	Width *int32 `json:"Width,omitempty"`
	Height *int32 `json:"Height,omitempty"`
	ExtraIds []string `json:"ExtraIds,omitempty"`
	DateLastSaved *time.Time `json:"DateLastSaved,omitempty"`
	// Gets or sets the remote trailers.
	RemoteTrailers []MediaUrl `json:"RemoteTrailers,omitempty"`
	SupportsExternalTransfer *bool `json:"SupportsExternalTransfer,omitempty"`
}

// NewBaseItem instantiates a new BaseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseItem() *BaseItem {
	this := BaseItem{}
	return &this
}

// NewBaseItemWithDefaults instantiates a new BaseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseItemWithDefaults() *BaseItem {
	this := BaseItem{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItem) GetSize() int64 {
	if o == nil || IsNil(o.Size.Get()) {
		var ret int64
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItem) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *BaseItem) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt64 and assigns it to the Size field.
func (o *BaseItem) SetSize(v int64) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *BaseItem) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *BaseItem) UnsetSize() {
	o.Size.Unset()
}

// GetContainer returns the Container field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItem) GetContainer() string {
	if o == nil || IsNil(o.Container.Get()) {
		var ret string
		return ret
	}
	return *o.Container.Get()
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItem) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Container.Get(), o.Container.IsSet()
}

// HasContainer returns a boolean if a field has been set.
func (o *BaseItem) HasContainer() bool {
	if o != nil && o.Container.IsSet() {
		return true
	}

	return false
}

// SetContainer gets a reference to the given NullableString and assigns it to the Container field.
func (o *BaseItem) SetContainer(v string) {
	o.Container.Set(&v)
}
// SetContainerNil sets the value for Container to be an explicit nil
func (o *BaseItem) SetContainerNil() {
	o.Container.Set(nil)
}

// UnsetContainer ensures that no value is present for Container, not even an explicit nil
func (o *BaseItem) UnsetContainer() {
	o.Container.Unset()
}

// GetIsHD returns the IsHD field value if set, zero value otherwise.
func (o *BaseItem) GetIsHD() bool {
	if o == nil || IsNil(o.IsHD) {
		var ret bool
		return ret
	}
	return *o.IsHD
}

// GetIsHDOk returns a tuple with the IsHD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetIsHDOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHD) {
		return nil, false
	}
	return o.IsHD, true
}

// HasIsHD returns a boolean if a field has been set.
func (o *BaseItem) HasIsHD() bool {
	if o != nil && !IsNil(o.IsHD) {
		return true
	}

	return false
}

// SetIsHD gets a reference to the given bool and assigns it to the IsHD field.
func (o *BaseItem) SetIsHD(v bool) {
	o.IsHD = &v
}

// GetIsShortcut returns the IsShortcut field value if set, zero value otherwise.
func (o *BaseItem) GetIsShortcut() bool {
	if o == nil || IsNil(o.IsShortcut) {
		var ret bool
		return ret
	}
	return *o.IsShortcut
}

// GetIsShortcutOk returns a tuple with the IsShortcut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetIsShortcutOk() (*bool, bool) {
	if o == nil || IsNil(o.IsShortcut) {
		return nil, false
	}
	return o.IsShortcut, true
}

// HasIsShortcut returns a boolean if a field has been set.
func (o *BaseItem) HasIsShortcut() bool {
	if o != nil && !IsNil(o.IsShortcut) {
		return true
	}

	return false
}

// SetIsShortcut gets a reference to the given bool and assigns it to the IsShortcut field.
func (o *BaseItem) SetIsShortcut(v bool) {
	o.IsShortcut = &v
}

// GetShortcutPath returns the ShortcutPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItem) GetShortcutPath() string {
	if o == nil || IsNil(o.ShortcutPath.Get()) {
		var ret string
		return ret
	}
	return *o.ShortcutPath.Get()
}

// GetShortcutPathOk returns a tuple with the ShortcutPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItem) GetShortcutPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortcutPath.Get(), o.ShortcutPath.IsSet()
}

// HasShortcutPath returns a boolean if a field has been set.
func (o *BaseItem) HasShortcutPath() bool {
	if o != nil && o.ShortcutPath.IsSet() {
		return true
	}

	return false
}

// SetShortcutPath gets a reference to the given NullableString and assigns it to the ShortcutPath field.
func (o *BaseItem) SetShortcutPath(v string) {
	o.ShortcutPath.Set(&v)
}
// SetShortcutPathNil sets the value for ShortcutPath to be an explicit nil
func (o *BaseItem) SetShortcutPathNil() {
	o.ShortcutPath.Set(nil)
}

// UnsetShortcutPath ensures that no value is present for ShortcutPath, not even an explicit nil
func (o *BaseItem) UnsetShortcutPath() {
	o.ShortcutPath.Unset()
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *BaseItem) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *BaseItem) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *BaseItem) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *BaseItem) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *BaseItem) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *BaseItem) SetHeight(v int32) {
	o.Height = &v
}

// GetExtraIds returns the ExtraIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItem) GetExtraIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ExtraIds
}

// GetExtraIdsOk returns a tuple with the ExtraIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItem) GetExtraIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ExtraIds) {
		return nil, false
	}
	return o.ExtraIds, true
}

// HasExtraIds returns a boolean if a field has been set.
func (o *BaseItem) HasExtraIds() bool {
	if o != nil && !IsNil(o.ExtraIds) {
		return true
	}

	return false
}

// SetExtraIds gets a reference to the given []string and assigns it to the ExtraIds field.
func (o *BaseItem) SetExtraIds(v []string) {
	o.ExtraIds = v
}

// GetDateLastSaved returns the DateLastSaved field value if set, zero value otherwise.
func (o *BaseItem) GetDateLastSaved() time.Time {
	if o == nil || IsNil(o.DateLastSaved) {
		var ret time.Time
		return ret
	}
	return *o.DateLastSaved
}

// GetDateLastSavedOk returns a tuple with the DateLastSaved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetDateLastSavedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateLastSaved) {
		return nil, false
	}
	return o.DateLastSaved, true
}

// HasDateLastSaved returns a boolean if a field has been set.
func (o *BaseItem) HasDateLastSaved() bool {
	if o != nil && !IsNil(o.DateLastSaved) {
		return true
	}

	return false
}

// SetDateLastSaved gets a reference to the given time.Time and assigns it to the DateLastSaved field.
func (o *BaseItem) SetDateLastSaved(v time.Time) {
	o.DateLastSaved = &v
}

// GetRemoteTrailers returns the RemoteTrailers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItem) GetRemoteTrailers() []MediaUrl {
	if o == nil {
		var ret []MediaUrl
		return ret
	}
	return o.RemoteTrailers
}

// GetRemoteTrailersOk returns a tuple with the RemoteTrailers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItem) GetRemoteTrailersOk() ([]MediaUrl, bool) {
	if o == nil || IsNil(o.RemoteTrailers) {
		return nil, false
	}
	return o.RemoteTrailers, true
}

// HasRemoteTrailers returns a boolean if a field has been set.
func (o *BaseItem) HasRemoteTrailers() bool {
	if o != nil && !IsNil(o.RemoteTrailers) {
		return true
	}

	return false
}

// SetRemoteTrailers gets a reference to the given []MediaUrl and assigns it to the RemoteTrailers field.
func (o *BaseItem) SetRemoteTrailers(v []MediaUrl) {
	o.RemoteTrailers = v
}

// GetSupportsExternalTransfer returns the SupportsExternalTransfer field value if set, zero value otherwise.
func (o *BaseItem) GetSupportsExternalTransfer() bool {
	if o == nil || IsNil(o.SupportsExternalTransfer) {
		var ret bool
		return ret
	}
	return *o.SupportsExternalTransfer
}

// GetSupportsExternalTransferOk returns a tuple with the SupportsExternalTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItem) GetSupportsExternalTransferOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsExternalTransfer) {
		return nil, false
	}
	return o.SupportsExternalTransfer, true
}

// HasSupportsExternalTransfer returns a boolean if a field has been set.
func (o *BaseItem) HasSupportsExternalTransfer() bool {
	if o != nil && !IsNil(o.SupportsExternalTransfer) {
		return true
	}

	return false
}

// SetSupportsExternalTransfer gets a reference to the given bool and assigns it to the SupportsExternalTransfer field.
func (o *BaseItem) SetSupportsExternalTransfer(v bool) {
	o.SupportsExternalTransfer = &v
}

func (o BaseItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Size.IsSet() {
		toSerialize["Size"] = o.Size.Get()
	}
	if o.Container.IsSet() {
		toSerialize["Container"] = o.Container.Get()
	}
	if !IsNil(o.IsHD) {
		toSerialize["IsHD"] = o.IsHD
	}
	if !IsNil(o.IsShortcut) {
		toSerialize["IsShortcut"] = o.IsShortcut
	}
	if o.ShortcutPath.IsSet() {
		toSerialize["ShortcutPath"] = o.ShortcutPath.Get()
	}
	if !IsNil(o.Width) {
		toSerialize["Width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["Height"] = o.Height
	}
	if o.ExtraIds != nil {
		toSerialize["ExtraIds"] = o.ExtraIds
	}
	if !IsNil(o.DateLastSaved) {
		toSerialize["DateLastSaved"] = o.DateLastSaved
	}
	if o.RemoteTrailers != nil {
		toSerialize["RemoteTrailers"] = o.RemoteTrailers
	}
	if !IsNil(o.SupportsExternalTransfer) {
		toSerialize["SupportsExternalTransfer"] = o.SupportsExternalTransfer
	}
	return toSerialize, nil
}

type NullableBaseItem struct {
	value *BaseItem
	isSet bool
}

func (v NullableBaseItem) Get() *BaseItem {
	return v.value
}

func (v *NullableBaseItem) Set(val *BaseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseItem(val *BaseItem) *NullableBaseItem {
	return &NullableBaseItem{value: val, isSet: true}
}

func (v NullableBaseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


