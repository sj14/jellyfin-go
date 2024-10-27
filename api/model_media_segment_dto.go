/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the MediaSegmentDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MediaSegmentDto{}

// MediaSegmentDto Api model for MediaSegment's.
type MediaSegmentDto struct {
	// Gets or sets the id of the media segment.
	Id *string `json:"Id,omitempty"`
	// Gets or sets the id of the associated item.
	ItemId *string `json:"ItemId,omitempty"`
	// Defines the types of content an individual Jellyfin.Data.Entities.MediaSegment represents.
	Type *MediaSegmentType `json:"Type,omitempty"`
	// Gets or sets the start of the segment.
	StartTicks *int64 `json:"StartTicks,omitempty"`
	// Gets or sets the end of the segment.
	EndTicks *int64 `json:"EndTicks,omitempty"`
}

// NewMediaSegmentDto instantiates a new MediaSegmentDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaSegmentDto() *MediaSegmentDto {
	this := MediaSegmentDto{}
	return &this
}

// NewMediaSegmentDtoWithDefaults instantiates a new MediaSegmentDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaSegmentDtoWithDefaults() *MediaSegmentDto {
	this := MediaSegmentDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MediaSegmentDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSegmentDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MediaSegmentDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MediaSegmentDto) SetId(v string) {
	o.Id = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *MediaSegmentDto) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSegmentDto) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *MediaSegmentDto) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *MediaSegmentDto) SetItemId(v string) {
	o.ItemId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MediaSegmentDto) GetType() MediaSegmentType {
	if o == nil || IsNil(o.Type) {
		var ret MediaSegmentType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSegmentDto) GetTypeOk() (*MediaSegmentType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MediaSegmentDto) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given MediaSegmentType and assigns it to the Type field.
func (o *MediaSegmentDto) SetType(v MediaSegmentType) {
	o.Type = &v
}

// GetStartTicks returns the StartTicks field value if set, zero value otherwise.
func (o *MediaSegmentDto) GetStartTicks() int64 {
	if o == nil || IsNil(o.StartTicks) {
		var ret int64
		return ret
	}
	return *o.StartTicks
}

// GetStartTicksOk returns a tuple with the StartTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSegmentDto) GetStartTicksOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTicks) {
		return nil, false
	}
	return o.StartTicks, true
}

// HasStartTicks returns a boolean if a field has been set.
func (o *MediaSegmentDto) HasStartTicks() bool {
	if o != nil && !IsNil(o.StartTicks) {
		return true
	}

	return false
}

// SetStartTicks gets a reference to the given int64 and assigns it to the StartTicks field.
func (o *MediaSegmentDto) SetStartTicks(v int64) {
	o.StartTicks = &v
}

// GetEndTicks returns the EndTicks field value if set, zero value otherwise.
func (o *MediaSegmentDto) GetEndTicks() int64 {
	if o == nil || IsNil(o.EndTicks) {
		var ret int64
		return ret
	}
	return *o.EndTicks
}

// GetEndTicksOk returns a tuple with the EndTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaSegmentDto) GetEndTicksOk() (*int64, bool) {
	if o == nil || IsNil(o.EndTicks) {
		return nil, false
	}
	return o.EndTicks, true
}

// HasEndTicks returns a boolean if a field has been set.
func (o *MediaSegmentDto) HasEndTicks() bool {
	if o != nil && !IsNil(o.EndTicks) {
		return true
	}

	return false
}

// SetEndTicks gets a reference to the given int64 and assigns it to the EndTicks field.
func (o *MediaSegmentDto) SetEndTicks(v int64) {
	o.EndTicks = &v
}

func (o MediaSegmentDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MediaSegmentDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
	}
	if !IsNil(o.Type) {
		toSerialize["Type"] = o.Type
	}
	if !IsNil(o.StartTicks) {
		toSerialize["StartTicks"] = o.StartTicks
	}
	if !IsNil(o.EndTicks) {
		toSerialize["EndTicks"] = o.EndTicks
	}
	return toSerialize, nil
}

type NullableMediaSegmentDto struct {
	value *MediaSegmentDto
	isSet bool
}

func (v NullableMediaSegmentDto) Get() *MediaSegmentDto {
	return v.value
}

func (v *NullableMediaSegmentDto) Set(val *MediaSegmentDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaSegmentDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaSegmentDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaSegmentDto(val *MediaSegmentDto) *NullableMediaSegmentDto {
	return &NullableMediaSegmentDto{value: val, isSet: true}
}

func (v NullableMediaSegmentDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaSegmentDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


