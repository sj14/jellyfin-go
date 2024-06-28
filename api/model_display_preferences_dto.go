/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the DisplayPreferencesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisplayPreferencesDto{}

// DisplayPreferencesDto Defines the display preferences for any item that supports them (usually Folders).
type DisplayPreferencesDto struct {
	// Gets or sets the user id.
	Id NullableString `json:"Id,omitempty"`
	// Gets or sets the type of the view.
	ViewType NullableString `json:"ViewType,omitempty"`
	// Gets or sets the sort by.
	SortBy NullableString `json:"SortBy,omitempty"`
	// Gets or sets the index by.
	IndexBy NullableString `json:"IndexBy,omitempty"`
	// Gets or sets a value indicating whether [remember indexing].
	RememberIndexing *bool `json:"RememberIndexing,omitempty"`
	// Gets or sets the height of the primary image.
	PrimaryImageHeight *int32 `json:"PrimaryImageHeight,omitempty"`
	// Gets or sets the width of the primary image.
	PrimaryImageWidth *int32 `json:"PrimaryImageWidth,omitempty"`
	// Gets or sets the custom prefs.
	CustomPrefs *map[string]string `json:"CustomPrefs,omitempty"`
	// Gets or sets the scroll direction.
	ScrollDirection *ScrollDirection `json:"ScrollDirection,omitempty"`
	// Gets or sets a value indicating whether to show backdrops on this item.
	ShowBackdrop *bool `json:"ShowBackdrop,omitempty"`
	// Gets or sets a value indicating whether [remember sorting].
	RememberSorting *bool `json:"RememberSorting,omitempty"`
	// Gets or sets the sort order.
	SortOrder *SortOrder `json:"SortOrder,omitempty"`
	// Gets or sets a value indicating whether [show sidebar].
	ShowSidebar *bool `json:"ShowSidebar,omitempty"`
	// Gets or sets the client.
	Client NullableString `json:"Client,omitempty"`
}

// NewDisplayPreferencesDto instantiates a new DisplayPreferencesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisplayPreferencesDto() *DisplayPreferencesDto {
	this := DisplayPreferencesDto{}
	return &this
}

// NewDisplayPreferencesDtoWithDefaults instantiates a new DisplayPreferencesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisplayPreferencesDtoWithDefaults() *DisplayPreferencesDto {
	this := DisplayPreferencesDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisplayPreferencesDto) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DisplayPreferencesDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DisplayPreferencesDto) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DisplayPreferencesDto) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DisplayPreferencesDto) UnsetId() {
	o.Id.Unset()
}

// GetViewType returns the ViewType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisplayPreferencesDto) GetViewType() string {
	if o == nil || IsNil(o.ViewType.Get()) {
		var ret string
		return ret
	}
	return *o.ViewType.Get()
}

// GetViewTypeOk returns a tuple with the ViewType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DisplayPreferencesDto) GetViewTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewType.Get(), o.ViewType.IsSet()
}

// HasViewType returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasViewType() bool {
	if o != nil && o.ViewType.IsSet() {
		return true
	}

	return false
}

// SetViewType gets a reference to the given NullableString and assigns it to the ViewType field.
func (o *DisplayPreferencesDto) SetViewType(v string) {
	o.ViewType.Set(&v)
}
// SetViewTypeNil sets the value for ViewType to be an explicit nil
func (o *DisplayPreferencesDto) SetViewTypeNil() {
	o.ViewType.Set(nil)
}

// UnsetViewType ensures that no value is present for ViewType, not even an explicit nil
func (o *DisplayPreferencesDto) UnsetViewType() {
	o.ViewType.Unset()
}

// GetSortBy returns the SortBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisplayPreferencesDto) GetSortBy() string {
	if o == nil || IsNil(o.SortBy.Get()) {
		var ret string
		return ret
	}
	return *o.SortBy.Get()
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DisplayPreferencesDto) GetSortByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SortBy.Get(), o.SortBy.IsSet()
}

// HasSortBy returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasSortBy() bool {
	if o != nil && o.SortBy.IsSet() {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given NullableString and assigns it to the SortBy field.
func (o *DisplayPreferencesDto) SetSortBy(v string) {
	o.SortBy.Set(&v)
}
// SetSortByNil sets the value for SortBy to be an explicit nil
func (o *DisplayPreferencesDto) SetSortByNil() {
	o.SortBy.Set(nil)
}

// UnsetSortBy ensures that no value is present for SortBy, not even an explicit nil
func (o *DisplayPreferencesDto) UnsetSortBy() {
	o.SortBy.Unset()
}

// GetIndexBy returns the IndexBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisplayPreferencesDto) GetIndexBy() string {
	if o == nil || IsNil(o.IndexBy.Get()) {
		var ret string
		return ret
	}
	return *o.IndexBy.Get()
}

// GetIndexByOk returns a tuple with the IndexBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DisplayPreferencesDto) GetIndexByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexBy.Get(), o.IndexBy.IsSet()
}

// HasIndexBy returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasIndexBy() bool {
	if o != nil && o.IndexBy.IsSet() {
		return true
	}

	return false
}

// SetIndexBy gets a reference to the given NullableString and assigns it to the IndexBy field.
func (o *DisplayPreferencesDto) SetIndexBy(v string) {
	o.IndexBy.Set(&v)
}
// SetIndexByNil sets the value for IndexBy to be an explicit nil
func (o *DisplayPreferencesDto) SetIndexByNil() {
	o.IndexBy.Set(nil)
}

// UnsetIndexBy ensures that no value is present for IndexBy, not even an explicit nil
func (o *DisplayPreferencesDto) UnsetIndexBy() {
	o.IndexBy.Unset()
}

// GetRememberIndexing returns the RememberIndexing field value if set, zero value otherwise.
func (o *DisplayPreferencesDto) GetRememberIndexing() bool {
	if o == nil || IsNil(o.RememberIndexing) {
		var ret bool
		return ret
	}
	return *o.RememberIndexing
}

// GetRememberIndexingOk returns a tuple with the RememberIndexing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisplayPreferencesDto) GetRememberIndexingOk() (*bool, bool) {
	if o == nil || IsNil(o.RememberIndexing) {
		return nil, false
	}
	return o.RememberIndexing, true
}

// HasRememberIndexing returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasRememberIndexing() bool {
	if o != nil && !IsNil(o.RememberIndexing) {
		return true
	}

	return false
}

// SetRememberIndexing gets a reference to the given bool and assigns it to the RememberIndexing field.
func (o *DisplayPreferencesDto) SetRememberIndexing(v bool) {
	o.RememberIndexing = &v
}

// GetPrimaryImageHeight returns the PrimaryImageHeight field value if set, zero value otherwise.
func (o *DisplayPreferencesDto) GetPrimaryImageHeight() int32 {
	if o == nil || IsNil(o.PrimaryImageHeight) {
		var ret int32
		return ret
	}
	return *o.PrimaryImageHeight
}

// GetPrimaryImageHeightOk returns a tuple with the PrimaryImageHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisplayPreferencesDto) GetPrimaryImageHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.PrimaryImageHeight) {
		return nil, false
	}
	return o.PrimaryImageHeight, true
}

// HasPrimaryImageHeight returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasPrimaryImageHeight() bool {
	if o != nil && !IsNil(o.PrimaryImageHeight) {
		return true
	}

	return false
}

// SetPrimaryImageHeight gets a reference to the given int32 and assigns it to the PrimaryImageHeight field.
func (o *DisplayPreferencesDto) SetPrimaryImageHeight(v int32) {
	o.PrimaryImageHeight = &v
}

// GetPrimaryImageWidth returns the PrimaryImageWidth field value if set, zero value otherwise.
func (o *DisplayPreferencesDto) GetPrimaryImageWidth() int32 {
	if o == nil || IsNil(o.PrimaryImageWidth) {
		var ret int32
		return ret
	}
	return *o.PrimaryImageWidth
}

// GetPrimaryImageWidthOk returns a tuple with the PrimaryImageWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisplayPreferencesDto) GetPrimaryImageWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.PrimaryImageWidth) {
		return nil, false
	}
	return o.PrimaryImageWidth, true
}

// HasPrimaryImageWidth returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasPrimaryImageWidth() bool {
	if o != nil && !IsNil(o.PrimaryImageWidth) {
		return true
	}

	return false
}

// SetPrimaryImageWidth gets a reference to the given int32 and assigns it to the PrimaryImageWidth field.
func (o *DisplayPreferencesDto) SetPrimaryImageWidth(v int32) {
	o.PrimaryImageWidth = &v
}

// GetCustomPrefs returns the CustomPrefs field value if set, zero value otherwise.
func (o *DisplayPreferencesDto) GetCustomPrefs() map[string]string {
	if o == nil || IsNil(o.CustomPrefs) {
		var ret map[string]string
		return ret
	}
	return *o.CustomPrefs
}

// GetCustomPrefsOk returns a tuple with the CustomPrefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisplayPreferencesDto) GetCustomPrefsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CustomPrefs) {
		return nil, false
	}
	return o.CustomPrefs, true
}

// HasCustomPrefs returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasCustomPrefs() bool {
	if o != nil && !IsNil(o.CustomPrefs) {
		return true
	}

	return false
}

// SetCustomPrefs gets a reference to the given map[string]string and assigns it to the CustomPrefs field.
func (o *DisplayPreferencesDto) SetCustomPrefs(v map[string]string) {
	o.CustomPrefs = &v
}

// GetScrollDirection returns the ScrollDirection field value if set, zero value otherwise.
func (o *DisplayPreferencesDto) GetScrollDirection() ScrollDirection {
	if o == nil || IsNil(o.ScrollDirection) {
		var ret ScrollDirection
		return ret
	}
	return *o.ScrollDirection
}

// GetScrollDirectionOk returns a tuple with the ScrollDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisplayPreferencesDto) GetScrollDirectionOk() (*ScrollDirection, bool) {
	if o == nil || IsNil(o.ScrollDirection) {
		return nil, false
	}
	return o.ScrollDirection, true
}

// HasScrollDirection returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasScrollDirection() bool {
	if o != nil && !IsNil(o.ScrollDirection) {
		return true
	}

	return false
}

// SetScrollDirection gets a reference to the given ScrollDirection and assigns it to the ScrollDirection field.
func (o *DisplayPreferencesDto) SetScrollDirection(v ScrollDirection) {
	o.ScrollDirection = &v
}

// GetShowBackdrop returns the ShowBackdrop field value if set, zero value otherwise.
func (o *DisplayPreferencesDto) GetShowBackdrop() bool {
	if o == nil || IsNil(o.ShowBackdrop) {
		var ret bool
		return ret
	}
	return *o.ShowBackdrop
}

// GetShowBackdropOk returns a tuple with the ShowBackdrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisplayPreferencesDto) GetShowBackdropOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowBackdrop) {
		return nil, false
	}
	return o.ShowBackdrop, true
}

// HasShowBackdrop returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasShowBackdrop() bool {
	if o != nil && !IsNil(o.ShowBackdrop) {
		return true
	}

	return false
}

// SetShowBackdrop gets a reference to the given bool and assigns it to the ShowBackdrop field.
func (o *DisplayPreferencesDto) SetShowBackdrop(v bool) {
	o.ShowBackdrop = &v
}

// GetRememberSorting returns the RememberSorting field value if set, zero value otherwise.
func (o *DisplayPreferencesDto) GetRememberSorting() bool {
	if o == nil || IsNil(o.RememberSorting) {
		var ret bool
		return ret
	}
	return *o.RememberSorting
}

// GetRememberSortingOk returns a tuple with the RememberSorting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisplayPreferencesDto) GetRememberSortingOk() (*bool, bool) {
	if o == nil || IsNil(o.RememberSorting) {
		return nil, false
	}
	return o.RememberSorting, true
}

// HasRememberSorting returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasRememberSorting() bool {
	if o != nil && !IsNil(o.RememberSorting) {
		return true
	}

	return false
}

// SetRememberSorting gets a reference to the given bool and assigns it to the RememberSorting field.
func (o *DisplayPreferencesDto) SetRememberSorting(v bool) {
	o.RememberSorting = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *DisplayPreferencesDto) GetSortOrder() SortOrder {
	if o == nil || IsNil(o.SortOrder) {
		var ret SortOrder
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisplayPreferencesDto) GetSortOrderOk() (*SortOrder, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given SortOrder and assigns it to the SortOrder field.
func (o *DisplayPreferencesDto) SetSortOrder(v SortOrder) {
	o.SortOrder = &v
}

// GetShowSidebar returns the ShowSidebar field value if set, zero value otherwise.
func (o *DisplayPreferencesDto) GetShowSidebar() bool {
	if o == nil || IsNil(o.ShowSidebar) {
		var ret bool
		return ret
	}
	return *o.ShowSidebar
}

// GetShowSidebarOk returns a tuple with the ShowSidebar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisplayPreferencesDto) GetShowSidebarOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowSidebar) {
		return nil, false
	}
	return o.ShowSidebar, true
}

// HasShowSidebar returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasShowSidebar() bool {
	if o != nil && !IsNil(o.ShowSidebar) {
		return true
	}

	return false
}

// SetShowSidebar gets a reference to the given bool and assigns it to the ShowSidebar field.
func (o *DisplayPreferencesDto) SetShowSidebar(v bool) {
	o.ShowSidebar = &v
}

// GetClient returns the Client field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DisplayPreferencesDto) GetClient() string {
	if o == nil || IsNil(o.Client.Get()) {
		var ret string
		return ret
	}
	return *o.Client.Get()
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DisplayPreferencesDto) GetClientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Client.Get(), o.Client.IsSet()
}

// HasClient returns a boolean if a field has been set.
func (o *DisplayPreferencesDto) HasClient() bool {
	if o != nil && o.Client.IsSet() {
		return true
	}

	return false
}

// SetClient gets a reference to the given NullableString and assigns it to the Client field.
func (o *DisplayPreferencesDto) SetClient(v string) {
	o.Client.Set(&v)
}
// SetClientNil sets the value for Client to be an explicit nil
func (o *DisplayPreferencesDto) SetClientNil() {
	o.Client.Set(nil)
}

// UnsetClient ensures that no value is present for Client, not even an explicit nil
func (o *DisplayPreferencesDto) UnsetClient() {
	o.Client.Unset()
}

func (o DisplayPreferencesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisplayPreferencesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.ViewType.IsSet() {
		toSerialize["ViewType"] = o.ViewType.Get()
	}
	if o.SortBy.IsSet() {
		toSerialize["SortBy"] = o.SortBy.Get()
	}
	if o.IndexBy.IsSet() {
		toSerialize["IndexBy"] = o.IndexBy.Get()
	}
	if !IsNil(o.RememberIndexing) {
		toSerialize["RememberIndexing"] = o.RememberIndexing
	}
	if !IsNil(o.PrimaryImageHeight) {
		toSerialize["PrimaryImageHeight"] = o.PrimaryImageHeight
	}
	if !IsNil(o.PrimaryImageWidth) {
		toSerialize["PrimaryImageWidth"] = o.PrimaryImageWidth
	}
	if !IsNil(o.CustomPrefs) {
		toSerialize["CustomPrefs"] = o.CustomPrefs
	}
	if !IsNil(o.ScrollDirection) {
		toSerialize["ScrollDirection"] = o.ScrollDirection
	}
	if !IsNil(o.ShowBackdrop) {
		toSerialize["ShowBackdrop"] = o.ShowBackdrop
	}
	if !IsNil(o.RememberSorting) {
		toSerialize["RememberSorting"] = o.RememberSorting
	}
	if !IsNil(o.SortOrder) {
		toSerialize["SortOrder"] = o.SortOrder
	}
	if !IsNil(o.ShowSidebar) {
		toSerialize["ShowSidebar"] = o.ShowSidebar
	}
	if o.Client.IsSet() {
		toSerialize["Client"] = o.Client.Get()
	}
	return toSerialize, nil
}

type NullableDisplayPreferencesDto struct {
	value *DisplayPreferencesDto
	isSet bool
}

func (v NullableDisplayPreferencesDto) Get() *DisplayPreferencesDto {
	return v.value
}

func (v *NullableDisplayPreferencesDto) Set(val *DisplayPreferencesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableDisplayPreferencesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableDisplayPreferencesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisplayPreferencesDto(val *DisplayPreferencesDto) *NullableDisplayPreferencesDto {
	return &NullableDisplayPreferencesDto{value: val, isSet: true}
}

func (v NullableDisplayPreferencesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisplayPreferencesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


