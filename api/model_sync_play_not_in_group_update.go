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

// checks if the SyncPlayNotInGroupUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyncPlayNotInGroupUpdate{}

// SyncPlayNotInGroupUpdate struct for SyncPlayNotInGroupUpdate
type SyncPlayNotInGroupUpdate struct {
	// Gets the group identifier.
	GroupId *string `json:"GroupId,omitempty"`
	// Gets the update data.
	Data *string `json:"Data,omitempty"`
	// Enum GroupUpdateType.
	Type *GroupUpdateType `json:"Type,omitempty"`
}

// NewSyncPlayNotInGroupUpdate instantiates a new SyncPlayNotInGroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncPlayNotInGroupUpdate() *SyncPlayNotInGroupUpdate {
	this := SyncPlayNotInGroupUpdate{}
	return &this
}

// NewSyncPlayNotInGroupUpdateWithDefaults instantiates a new SyncPlayNotInGroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncPlayNotInGroupUpdateWithDefaults() *SyncPlayNotInGroupUpdate {
	this := SyncPlayNotInGroupUpdate{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *SyncPlayNotInGroupUpdate) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncPlayNotInGroupUpdate) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *SyncPlayNotInGroupUpdate) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *SyncPlayNotInGroupUpdate) SetGroupId(v string) {
	o.GroupId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SyncPlayNotInGroupUpdate) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncPlayNotInGroupUpdate) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SyncPlayNotInGroupUpdate) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *SyncPlayNotInGroupUpdate) SetData(v string) {
	o.Data = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyncPlayNotInGroupUpdate) GetType() GroupUpdateType {
	if o == nil || IsNil(o.Type) {
		var ret GroupUpdateType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncPlayNotInGroupUpdate) GetTypeOk() (*GroupUpdateType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyncPlayNotInGroupUpdate) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given GroupUpdateType and assigns it to the Type field.
func (o *SyncPlayNotInGroupUpdate) SetType(v GroupUpdateType) {
	o.Type = &v
}

func (o SyncPlayNotInGroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyncPlayNotInGroupUpdate) ToMap() (map[string]interface{}, error) {
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

type NullableSyncPlayNotInGroupUpdate struct {
	value *SyncPlayNotInGroupUpdate
	isSet bool
}

func (v NullableSyncPlayNotInGroupUpdate) Get() *SyncPlayNotInGroupUpdate {
	return v.value
}

func (v *NullableSyncPlayNotInGroupUpdate) Set(val *SyncPlayNotInGroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncPlayNotInGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncPlayNotInGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncPlayNotInGroupUpdate(val *SyncPlayNotInGroupUpdate) *NullableSyncPlayNotInGroupUpdate {
	return &NullableSyncPlayNotInGroupUpdate{value: val, isSet: true}
}

func (v NullableSyncPlayNotInGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncPlayNotInGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


