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

// checks if the SyncPlayLibraryAccessDeniedUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncPlayLibraryAccessDeniedUpdate{}

// SyncPlayLibraryAccessDeniedUpdate struct for SyncPlayLibraryAccessDeniedUpdate
type SyncPlayLibraryAccessDeniedUpdate struct {
	// Gets the group identifier.
	GroupId *string `json:"GroupId,omitempty"`
	// Gets the update data.
	Data *string `json:"Data,omitempty"`
	// Enum GroupUpdateType.
	Type *GroupUpdateType `json:"Type,omitempty"`
}

// NewSyncPlayLibraryAccessDeniedUpdate instantiates a new SyncPlayLibraryAccessDeniedUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncPlayLibraryAccessDeniedUpdate() *SyncPlayLibraryAccessDeniedUpdate {
	this := SyncPlayLibraryAccessDeniedUpdate{}
	return &this
}

// NewSyncPlayLibraryAccessDeniedUpdateWithDefaults instantiates a new SyncPlayLibraryAccessDeniedUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncPlayLibraryAccessDeniedUpdateWithDefaults() *SyncPlayLibraryAccessDeniedUpdate {
	this := SyncPlayLibraryAccessDeniedUpdate{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *SyncPlayLibraryAccessDeniedUpdate) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncPlayLibraryAccessDeniedUpdate) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *SyncPlayLibraryAccessDeniedUpdate) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *SyncPlayLibraryAccessDeniedUpdate) SetGroupId(v string) {
	o.GroupId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SyncPlayLibraryAccessDeniedUpdate) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncPlayLibraryAccessDeniedUpdate) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SyncPlayLibraryAccessDeniedUpdate) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *SyncPlayLibraryAccessDeniedUpdate) SetData(v string) {
	o.Data = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyncPlayLibraryAccessDeniedUpdate) GetType() GroupUpdateType {
	if o == nil || IsNil(o.Type) {
		var ret GroupUpdateType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncPlayLibraryAccessDeniedUpdate) GetTypeOk() (*GroupUpdateType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyncPlayLibraryAccessDeniedUpdate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given GroupUpdateType and assigns it to the Type field.
func (o *SyncPlayLibraryAccessDeniedUpdate) SetType(v GroupUpdateType) {
	o.Type = &v
}

func (o SyncPlayLibraryAccessDeniedUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncPlayLibraryAccessDeniedUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["GroupId"] = o.GroupId
	}
	if !IsNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSyncPlayLibraryAccessDeniedUpdate struct {
	value *SyncPlayLibraryAccessDeniedUpdate
	isSet bool
}

func (v NullableSyncPlayLibraryAccessDeniedUpdate) Get() *SyncPlayLibraryAccessDeniedUpdate {
	return v.value
}

func (v *NullableSyncPlayLibraryAccessDeniedUpdate) Set(val *SyncPlayLibraryAccessDeniedUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncPlayLibraryAccessDeniedUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncPlayLibraryAccessDeniedUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncPlayLibraryAccessDeniedUpdate(val *SyncPlayLibraryAccessDeniedUpdate) *NullableSyncPlayLibraryAccessDeniedUpdate {
	return &NullableSyncPlayLibraryAccessDeniedUpdate{value: val, isSet: true}
}

func (v NullableSyncPlayLibraryAccessDeniedUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncPlayLibraryAccessDeniedUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


