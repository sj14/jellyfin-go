/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ChannelFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelFeatures{}

// ChannelFeatures struct for ChannelFeatures
type ChannelFeatures struct {
	// Gets or sets the name.
	Name *string `json:"Name,omitempty"`
	// Gets or sets the identifier.
	Id *string `json:"Id,omitempty"`
	// Gets or sets a value indicating whether this instance can search.
	CanSearch *bool `json:"CanSearch,omitempty"`
	// Gets or sets the media types.
	MediaTypes []ChannelMediaType `json:"MediaTypes,omitempty"`
	// Gets or sets the content types.
	ContentTypes []ChannelMediaContentType `json:"ContentTypes,omitempty"`
	// Gets or sets the maximum number of records the channel allows retrieving at a time.
	MaxPageSize NullableInt32 `json:"MaxPageSize,omitempty"`
	// Gets or sets the automatic refresh levels.
	AutoRefreshLevels NullableInt32 `json:"AutoRefreshLevels,omitempty"`
	// Gets or sets the default sort orders.
	DefaultSortFields []ChannelItemSortField `json:"DefaultSortFields,omitempty"`
	// Gets or sets a value indicating whether a sort ascending/descending toggle is supported.
	SupportsSortOrderToggle *bool `json:"SupportsSortOrderToggle,omitempty"`
	// Gets or sets a value indicating whether [supports latest media].
	SupportsLatestMedia *bool `json:"SupportsLatestMedia,omitempty"`
	// Gets or sets a value indicating whether this instance can filter.
	CanFilter *bool `json:"CanFilter,omitempty"`
	// Gets or sets a value indicating whether [supports content downloading].
	SupportsContentDownloading *bool `json:"SupportsContentDownloading,omitempty"`
}

// NewChannelFeatures instantiates a new ChannelFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelFeatures() *ChannelFeatures {
	this := ChannelFeatures{}
	return &this
}

// NewChannelFeaturesWithDefaults instantiates a new ChannelFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelFeaturesWithDefaults() *ChannelFeatures {
	this := ChannelFeatures{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChannelFeatures) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelFeatures) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChannelFeatures) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChannelFeatures) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelFeatures) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelFeatures) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelFeatures) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChannelFeatures) SetId(v string) {
	o.Id = &v
}

// GetCanSearch returns the CanSearch field value if set, zero value otherwise.
func (o *ChannelFeatures) GetCanSearch() bool {
	if o == nil || IsNil(o.CanSearch) {
		var ret bool
		return ret
	}
	return *o.CanSearch
}

// GetCanSearchOk returns a tuple with the CanSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelFeatures) GetCanSearchOk() (*bool, bool) {
	if o == nil || IsNil(o.CanSearch) {
		return nil, false
	}
	return o.CanSearch, true
}

// HasCanSearch returns a boolean if a field has been set.
func (o *ChannelFeatures) HasCanSearch() bool {
	if o != nil && !IsNil(o.CanSearch) {
		return true
	}

	return false
}

// SetCanSearch gets a reference to the given bool and assigns it to the CanSearch field.
func (o *ChannelFeatures) SetCanSearch(v bool) {
	o.CanSearch = &v
}

// GetMediaTypes returns the MediaTypes field value if set, zero value otherwise.
func (o *ChannelFeatures) GetMediaTypes() []ChannelMediaType {
	if o == nil || IsNil(o.MediaTypes) {
		var ret []ChannelMediaType
		return ret
	}
	return o.MediaTypes
}

// GetMediaTypesOk returns a tuple with the MediaTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelFeatures) GetMediaTypesOk() ([]ChannelMediaType, bool) {
	if o == nil || IsNil(o.MediaTypes) {
		return nil, false
	}
	return o.MediaTypes, true
}

// HasMediaTypes returns a boolean if a field has been set.
func (o *ChannelFeatures) HasMediaTypes() bool {
	if o != nil && !IsNil(o.MediaTypes) {
		return true
	}

	return false
}

// SetMediaTypes gets a reference to the given []ChannelMediaType and assigns it to the MediaTypes field.
func (o *ChannelFeatures) SetMediaTypes(v []ChannelMediaType) {
	o.MediaTypes = v
}

// GetContentTypes returns the ContentTypes field value if set, zero value otherwise.
func (o *ChannelFeatures) GetContentTypes() []ChannelMediaContentType {
	if o == nil || IsNil(o.ContentTypes) {
		var ret []ChannelMediaContentType
		return ret
	}
	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelFeatures) GetContentTypesOk() ([]ChannelMediaContentType, bool) {
	if o == nil || IsNil(o.ContentTypes) {
		return nil, false
	}
	return o.ContentTypes, true
}

// HasContentTypes returns a boolean if a field has been set.
func (o *ChannelFeatures) HasContentTypes() bool {
	if o != nil && !IsNil(o.ContentTypes) {
		return true
	}

	return false
}

// SetContentTypes gets a reference to the given []ChannelMediaContentType and assigns it to the ContentTypes field.
func (o *ChannelFeatures) SetContentTypes(v []ChannelMediaContentType) {
	o.ContentTypes = v
}

// GetMaxPageSize returns the MaxPageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelFeatures) GetMaxPageSize() int32 {
	if o == nil || IsNil(o.MaxPageSize.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxPageSize.Get()
}

// GetMaxPageSizeOk returns a tuple with the MaxPageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelFeatures) GetMaxPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxPageSize.Get(), o.MaxPageSize.IsSet()
}

// HasMaxPageSize returns a boolean if a field has been set.
func (o *ChannelFeatures) HasMaxPageSize() bool {
	if o != nil && o.MaxPageSize.IsSet() {
		return true
	}

	return false
}

// SetMaxPageSize gets a reference to the given NullableInt32 and assigns it to the MaxPageSize field.
func (o *ChannelFeatures) SetMaxPageSize(v int32) {
	o.MaxPageSize.Set(&v)
}
// SetMaxPageSizeNil sets the value for MaxPageSize to be an explicit nil
func (o *ChannelFeatures) SetMaxPageSizeNil() {
	o.MaxPageSize.Set(nil)
}

// UnsetMaxPageSize ensures that no value is present for MaxPageSize, not even an explicit nil
func (o *ChannelFeatures) UnsetMaxPageSize() {
	o.MaxPageSize.Unset()
}

// GetAutoRefreshLevels returns the AutoRefreshLevels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChannelFeatures) GetAutoRefreshLevels() int32 {
	if o == nil || IsNil(o.AutoRefreshLevels.Get()) {
		var ret int32
		return ret
	}
	return *o.AutoRefreshLevels.Get()
}

// GetAutoRefreshLevelsOk returns a tuple with the AutoRefreshLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChannelFeatures) GetAutoRefreshLevelsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutoRefreshLevels.Get(), o.AutoRefreshLevels.IsSet()
}

// HasAutoRefreshLevels returns a boolean if a field has been set.
func (o *ChannelFeatures) HasAutoRefreshLevels() bool {
	if o != nil && o.AutoRefreshLevels.IsSet() {
		return true
	}

	return false
}

// SetAutoRefreshLevels gets a reference to the given NullableInt32 and assigns it to the AutoRefreshLevels field.
func (o *ChannelFeatures) SetAutoRefreshLevels(v int32) {
	o.AutoRefreshLevels.Set(&v)
}
// SetAutoRefreshLevelsNil sets the value for AutoRefreshLevels to be an explicit nil
func (o *ChannelFeatures) SetAutoRefreshLevelsNil() {
	o.AutoRefreshLevels.Set(nil)
}

// UnsetAutoRefreshLevels ensures that no value is present for AutoRefreshLevels, not even an explicit nil
func (o *ChannelFeatures) UnsetAutoRefreshLevels() {
	o.AutoRefreshLevels.Unset()
}

// GetDefaultSortFields returns the DefaultSortFields field value if set, zero value otherwise.
func (o *ChannelFeatures) GetDefaultSortFields() []ChannelItemSortField {
	if o == nil || IsNil(o.DefaultSortFields) {
		var ret []ChannelItemSortField
		return ret
	}
	return o.DefaultSortFields
}

// GetDefaultSortFieldsOk returns a tuple with the DefaultSortFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelFeatures) GetDefaultSortFieldsOk() ([]ChannelItemSortField, bool) {
	if o == nil || IsNil(o.DefaultSortFields) {
		return nil, false
	}
	return o.DefaultSortFields, true
}

// HasDefaultSortFields returns a boolean if a field has been set.
func (o *ChannelFeatures) HasDefaultSortFields() bool {
	if o != nil && !IsNil(o.DefaultSortFields) {
		return true
	}

	return false
}

// SetDefaultSortFields gets a reference to the given []ChannelItemSortField and assigns it to the DefaultSortFields field.
func (o *ChannelFeatures) SetDefaultSortFields(v []ChannelItemSortField) {
	o.DefaultSortFields = v
}

// GetSupportsSortOrderToggle returns the SupportsSortOrderToggle field value if set, zero value otherwise.
func (o *ChannelFeatures) GetSupportsSortOrderToggle() bool {
	if o == nil || IsNil(o.SupportsSortOrderToggle) {
		var ret bool
		return ret
	}
	return *o.SupportsSortOrderToggle
}

// GetSupportsSortOrderToggleOk returns a tuple with the SupportsSortOrderToggle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelFeatures) GetSupportsSortOrderToggleOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsSortOrderToggle) {
		return nil, false
	}
	return o.SupportsSortOrderToggle, true
}

// HasSupportsSortOrderToggle returns a boolean if a field has been set.
func (o *ChannelFeatures) HasSupportsSortOrderToggle() bool {
	if o != nil && !IsNil(o.SupportsSortOrderToggle) {
		return true
	}

	return false
}

// SetSupportsSortOrderToggle gets a reference to the given bool and assigns it to the SupportsSortOrderToggle field.
func (o *ChannelFeatures) SetSupportsSortOrderToggle(v bool) {
	o.SupportsSortOrderToggle = &v
}

// GetSupportsLatestMedia returns the SupportsLatestMedia field value if set, zero value otherwise.
func (o *ChannelFeatures) GetSupportsLatestMedia() bool {
	if o == nil || IsNil(o.SupportsLatestMedia) {
		var ret bool
		return ret
	}
	return *o.SupportsLatestMedia
}

// GetSupportsLatestMediaOk returns a tuple with the SupportsLatestMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelFeatures) GetSupportsLatestMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsLatestMedia) {
		return nil, false
	}
	return o.SupportsLatestMedia, true
}

// HasSupportsLatestMedia returns a boolean if a field has been set.
func (o *ChannelFeatures) HasSupportsLatestMedia() bool {
	if o != nil && !IsNil(o.SupportsLatestMedia) {
		return true
	}

	return false
}

// SetSupportsLatestMedia gets a reference to the given bool and assigns it to the SupportsLatestMedia field.
func (o *ChannelFeatures) SetSupportsLatestMedia(v bool) {
	o.SupportsLatestMedia = &v
}

// GetCanFilter returns the CanFilter field value if set, zero value otherwise.
func (o *ChannelFeatures) GetCanFilter() bool {
	if o == nil || IsNil(o.CanFilter) {
		var ret bool
		return ret
	}
	return *o.CanFilter
}

// GetCanFilterOk returns a tuple with the CanFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelFeatures) GetCanFilterOk() (*bool, bool) {
	if o == nil || IsNil(o.CanFilter) {
		return nil, false
	}
	return o.CanFilter, true
}

// HasCanFilter returns a boolean if a field has been set.
func (o *ChannelFeatures) HasCanFilter() bool {
	if o != nil && !IsNil(o.CanFilter) {
		return true
	}

	return false
}

// SetCanFilter gets a reference to the given bool and assigns it to the CanFilter field.
func (o *ChannelFeatures) SetCanFilter(v bool) {
	o.CanFilter = &v
}

// GetSupportsContentDownloading returns the SupportsContentDownloading field value if set, zero value otherwise.
func (o *ChannelFeatures) GetSupportsContentDownloading() bool {
	if o == nil || IsNil(o.SupportsContentDownloading) {
		var ret bool
		return ret
	}
	return *o.SupportsContentDownloading
}

// GetSupportsContentDownloadingOk returns a tuple with the SupportsContentDownloading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelFeatures) GetSupportsContentDownloadingOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportsContentDownloading) {
		return nil, false
	}
	return o.SupportsContentDownloading, true
}

// HasSupportsContentDownloading returns a boolean if a field has been set.
func (o *ChannelFeatures) HasSupportsContentDownloading() bool {
	if o != nil && !IsNil(o.SupportsContentDownloading) {
		return true
	}

	return false
}

// SetSupportsContentDownloading gets a reference to the given bool and assigns it to the SupportsContentDownloading field.
func (o *ChannelFeatures) SetSupportsContentDownloading(v bool) {
	o.SupportsContentDownloading = &v
}

func (o ChannelFeatures) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.CanSearch) {
		toSerialize["CanSearch"] = o.CanSearch
	}
	if !IsNil(o.MediaTypes) {
		toSerialize["MediaTypes"] = o.MediaTypes
	}
	if !IsNil(o.ContentTypes) {
		toSerialize["ContentTypes"] = o.ContentTypes
	}
	if o.MaxPageSize.IsSet() {
		toSerialize["MaxPageSize"] = o.MaxPageSize.Get()
	}
	if o.AutoRefreshLevels.IsSet() {
		toSerialize["AutoRefreshLevels"] = o.AutoRefreshLevels.Get()
	}
	if !IsNil(o.DefaultSortFields) {
		toSerialize["DefaultSortFields"] = o.DefaultSortFields
	}
	if !IsNil(o.SupportsSortOrderToggle) {
		toSerialize["SupportsSortOrderToggle"] = o.SupportsSortOrderToggle
	}
	if !IsNil(o.SupportsLatestMedia) {
		toSerialize["SupportsLatestMedia"] = o.SupportsLatestMedia
	}
	if !IsNil(o.CanFilter) {
		toSerialize["CanFilter"] = o.CanFilter
	}
	if !IsNil(o.SupportsContentDownloading) {
		toSerialize["SupportsContentDownloading"] = o.SupportsContentDownloading
	}
	return toSerialize, nil
}

type NullableChannelFeatures struct {
	value *ChannelFeatures
	isSet bool
}

func (v NullableChannelFeatures) Get() *ChannelFeatures {
	return v.value
}

func (v *NullableChannelFeatures) Set(val *ChannelFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelFeatures(val *ChannelFeatures) *NullableChannelFeatures {
	return &NullableChannelFeatures{value: val, isSet: true}
}

func (v NullableChannelFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


