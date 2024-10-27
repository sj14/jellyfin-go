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

// checks if the PlayRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlayRequestDto{}

// PlayRequestDto Class PlayRequestDto.
type PlayRequestDto struct {
	// Gets or sets the playing queue.
	PlayingQueue []string `json:"PlayingQueue,omitempty"`
	// Gets or sets the position of the playing item in the queue.
	PlayingItemPosition *int32 `json:"PlayingItemPosition,omitempty"`
	// Gets or sets the start position ticks.
	StartPositionTicks *int64 `json:"StartPositionTicks,omitempty"`
}

// NewPlayRequestDto instantiates a new PlayRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayRequestDto() *PlayRequestDto {
	this := PlayRequestDto{}
	return &this
}

// NewPlayRequestDtoWithDefaults instantiates a new PlayRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayRequestDtoWithDefaults() *PlayRequestDto {
	this := PlayRequestDto{}
	return &this
}

// GetPlayingQueue returns the PlayingQueue field value if set, zero value otherwise.
func (o *PlayRequestDto) GetPlayingQueue() []string {
	if o == nil || IsNil(o.PlayingQueue) {
		var ret []string
		return ret
	}
	return o.PlayingQueue
}

// GetPlayingQueueOk returns a tuple with the PlayingQueue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayRequestDto) GetPlayingQueueOk() ([]string, bool) {
	if o == nil || IsNil(o.PlayingQueue) {
		return nil, false
	}
	return o.PlayingQueue, true
}

// HasPlayingQueue returns a boolean if a field has been set.
func (o *PlayRequestDto) HasPlayingQueue() bool {
	if o != nil && !IsNil(o.PlayingQueue) {
		return true
	}

	return false
}

// SetPlayingQueue gets a reference to the given []string and assigns it to the PlayingQueue field.
func (o *PlayRequestDto) SetPlayingQueue(v []string) {
	o.PlayingQueue = v
}

// GetPlayingItemPosition returns the PlayingItemPosition field value if set, zero value otherwise.
func (o *PlayRequestDto) GetPlayingItemPosition() int32 {
	if o == nil || IsNil(o.PlayingItemPosition) {
		var ret int32
		return ret
	}
	return *o.PlayingItemPosition
}

// GetPlayingItemPositionOk returns a tuple with the PlayingItemPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayRequestDto) GetPlayingItemPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.PlayingItemPosition) {
		return nil, false
	}
	return o.PlayingItemPosition, true
}

// HasPlayingItemPosition returns a boolean if a field has been set.
func (o *PlayRequestDto) HasPlayingItemPosition() bool {
	if o != nil && !IsNil(o.PlayingItemPosition) {
		return true
	}

	return false
}

// SetPlayingItemPosition gets a reference to the given int32 and assigns it to the PlayingItemPosition field.
func (o *PlayRequestDto) SetPlayingItemPosition(v int32) {
	o.PlayingItemPosition = &v
}

// GetStartPositionTicks returns the StartPositionTicks field value if set, zero value otherwise.
func (o *PlayRequestDto) GetStartPositionTicks() int64 {
	if o == nil || IsNil(o.StartPositionTicks) {
		var ret int64
		return ret
	}
	return *o.StartPositionTicks
}

// GetStartPositionTicksOk returns a tuple with the StartPositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlayRequestDto) GetStartPositionTicksOk() (*int64, bool) {
	if o == nil || IsNil(o.StartPositionTicks) {
		return nil, false
	}
	return o.StartPositionTicks, true
}

// HasStartPositionTicks returns a boolean if a field has been set.
func (o *PlayRequestDto) HasStartPositionTicks() bool {
	if o != nil && !IsNil(o.StartPositionTicks) {
		return true
	}

	return false
}

// SetStartPositionTicks gets a reference to the given int64 and assigns it to the StartPositionTicks field.
func (o *PlayRequestDto) SetStartPositionTicks(v int64) {
	o.StartPositionTicks = &v
}

func (o PlayRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlayRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlayingQueue) {
		toSerialize["PlayingQueue"] = o.PlayingQueue
	}
	if !IsNil(o.PlayingItemPosition) {
		toSerialize["PlayingItemPosition"] = o.PlayingItemPosition
	}
	if !IsNil(o.StartPositionTicks) {
		toSerialize["StartPositionTicks"] = o.StartPositionTicks
	}
	return toSerialize, nil
}

type NullablePlayRequestDto struct {
	value *PlayRequestDto
	isSet bool
}

func (v NullablePlayRequestDto) Get() *PlayRequestDto {
	return v.value
}

func (v *NullablePlayRequestDto) Set(val *PlayRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayRequestDto(val *PlayRequestDto) *NullablePlayRequestDto {
	return &NullablePlayRequestDto{value: val, isSet: true}
}

func (v NullablePlayRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


