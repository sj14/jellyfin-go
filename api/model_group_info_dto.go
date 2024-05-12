/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the GroupInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupInfoDto{}

// GroupInfoDto Class GroupInfoDto.
type GroupInfoDto struct {
	// Gets the group identifier.
	GroupId *string `json:"GroupId,omitempty"`
	// Gets the group name.
	GroupName *string `json:"GroupName,omitempty"`
	// Gets the group state.
	State *GroupStateType `json:"State,omitempty"`
	// Gets the participants.
	Participants []string `json:"Participants,omitempty"`
	// Gets the date when this DTO has been created.
	LastUpdatedAt *time.Time `json:"LastUpdatedAt,omitempty"`
}

// NewGroupInfoDto instantiates a new GroupInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupInfoDto() *GroupInfoDto {
	this := GroupInfoDto{}
	return &this
}

// NewGroupInfoDtoWithDefaults instantiates a new GroupInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupInfoDtoWithDefaults() *GroupInfoDto {
	this := GroupInfoDto{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *GroupInfoDto) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInfoDto) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *GroupInfoDto) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *GroupInfoDto) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *GroupInfoDto) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInfoDto) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *GroupInfoDto) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *GroupInfoDto) SetGroupName(v string) {
	o.GroupName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GroupInfoDto) GetState() GroupStateType {
	if o == nil || IsNil(o.State) {
		var ret GroupStateType
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInfoDto) GetStateOk() (*GroupStateType, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GroupInfoDto) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given GroupStateType and assigns it to the State field.
func (o *GroupInfoDto) SetState(v GroupStateType) {
	o.State = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *GroupInfoDto) GetParticipants() []string {
	if o == nil || IsNil(o.Participants) {
		var ret []string
		return ret
	}
	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInfoDto) GetParticipantsOk() ([]string, bool) {
	if o == nil || IsNil(o.Participants) {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *GroupInfoDto) HasParticipants() bool {
	if o != nil && !IsNil(o.Participants) {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []string and assigns it to the Participants field.
func (o *GroupInfoDto) SetParticipants(v []string) {
	o.Participants = v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value if set, zero value otherwise.
func (o *GroupInfoDto) GetLastUpdatedAt() time.Time {
	if o == nil || IsNil(o.LastUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupInfoDto) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedAt) {
		return nil, false
	}
	return o.LastUpdatedAt, true
}

// HasLastUpdatedAt returns a boolean if a field has been set.
func (o *GroupInfoDto) HasLastUpdatedAt() bool {
	if o != nil && !IsNil(o.LastUpdatedAt) {
		return true
	}

	return false
}

// SetLastUpdatedAt gets a reference to the given time.Time and assigns it to the LastUpdatedAt field.
func (o *GroupInfoDto) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = &v
}

func (o GroupInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["GroupId"] = o.GroupId
	}
	if !IsNil(o.GroupName) {
		toSerialize["GroupName"] = o.GroupName
	}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if !IsNil(o.Participants) {
		toSerialize["Participants"] = o.Participants
	}
	if !IsNil(o.LastUpdatedAt) {
		toSerialize["LastUpdatedAt"] = o.LastUpdatedAt
	}
	return toSerialize, nil
}

type NullableGroupInfoDto struct {
	value *GroupInfoDto
	isSet bool
}

func (v NullableGroupInfoDto) Get() *GroupInfoDto {
	return v.value
}

func (v *NullableGroupInfoDto) Set(val *GroupInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupInfoDto(val *GroupInfoDto) *NullableGroupInfoDto {
	return &NullableGroupInfoDto{value: val, isSet: true}
}

func (v NullableGroupInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


