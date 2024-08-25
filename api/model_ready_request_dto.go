/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the ReadyRequestDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReadyRequestDto{}

// ReadyRequestDto Class ReadyRequest.
type ReadyRequestDto struct {
	// Gets or sets when the request has been made by the client.
	When *time.Time `json:"When,omitempty"`
	// Gets or sets the position ticks.
	PositionTicks *int64 `json:"PositionTicks,omitempty"`
	// Gets or sets a value indicating whether the client playback is unpaused.
	IsPlaying *bool `json:"IsPlaying,omitempty"`
	// Gets or sets the playlist item identifier of the playing item.
	PlaylistItemId *string `json:"PlaylistItemId,omitempty"`
}

// NewReadyRequestDto instantiates a new ReadyRequestDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadyRequestDto() *ReadyRequestDto {
	this := ReadyRequestDto{}
	return &this
}

// NewReadyRequestDtoWithDefaults instantiates a new ReadyRequestDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadyRequestDtoWithDefaults() *ReadyRequestDto {
	this := ReadyRequestDto{}
	return &this
}

// GetWhen returns the When field value if set, zero value otherwise.
func (o *ReadyRequestDto) GetWhen() time.Time {
	if o == nil || IsNil(o.When) {
		var ret time.Time
		return ret
	}
	return *o.When
}

// GetWhenOk returns a tuple with the When field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadyRequestDto) GetWhenOk() (*time.Time, bool) {
	if o == nil || IsNil(o.When) {
		return nil, false
	}
	return o.When, true
}

// HasWhen returns a boolean if a field has been set.
func (o *ReadyRequestDto) HasWhen() bool {
	if o != nil && !IsNil(o.When) {
		return true
	}

	return false
}

// SetWhen gets a reference to the given time.Time and assigns it to the When field.
func (o *ReadyRequestDto) SetWhen(v time.Time) {
	o.When = &v
}

// GetPositionTicks returns the PositionTicks field value if set, zero value otherwise.
func (o *ReadyRequestDto) GetPositionTicks() int64 {
	if o == nil || IsNil(o.PositionTicks) {
		var ret int64
		return ret
	}
	return *o.PositionTicks
}

// GetPositionTicksOk returns a tuple with the PositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadyRequestDto) GetPositionTicksOk() (*int64, bool) {
	if o == nil || IsNil(o.PositionTicks) {
		return nil, false
	}
	return o.PositionTicks, true
}

// HasPositionTicks returns a boolean if a field has been set.
func (o *ReadyRequestDto) HasPositionTicks() bool {
	if o != nil && !IsNil(o.PositionTicks) {
		return true
	}

	return false
}

// SetPositionTicks gets a reference to the given int64 and assigns it to the PositionTicks field.
func (o *ReadyRequestDto) SetPositionTicks(v int64) {
	o.PositionTicks = &v
}

// GetIsPlaying returns the IsPlaying field value if set, zero value otherwise.
func (o *ReadyRequestDto) GetIsPlaying() bool {
	if o == nil || IsNil(o.IsPlaying) {
		var ret bool
		return ret
	}
	return *o.IsPlaying
}

// GetIsPlayingOk returns a tuple with the IsPlaying field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadyRequestDto) GetIsPlayingOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPlaying) {
		return nil, false
	}
	return o.IsPlaying, true
}

// HasIsPlaying returns a boolean if a field has been set.
func (o *ReadyRequestDto) HasIsPlaying() bool {
	if o != nil && !IsNil(o.IsPlaying) {
		return true
	}

	return false
}

// SetIsPlaying gets a reference to the given bool and assigns it to the IsPlaying field.
func (o *ReadyRequestDto) SetIsPlaying(v bool) {
	o.IsPlaying = &v
}

// GetPlaylistItemId returns the PlaylistItemId field value if set, zero value otherwise.
func (o *ReadyRequestDto) GetPlaylistItemId() string {
	if o == nil || IsNil(o.PlaylistItemId) {
		var ret string
		return ret
	}
	return *o.PlaylistItemId
}

// GetPlaylistItemIdOk returns a tuple with the PlaylistItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadyRequestDto) GetPlaylistItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlaylistItemId) {
		return nil, false
	}
	return o.PlaylistItemId, true
}

// HasPlaylistItemId returns a boolean if a field has been set.
func (o *ReadyRequestDto) HasPlaylistItemId() bool {
	if o != nil && !IsNil(o.PlaylistItemId) {
		return true
	}

	return false
}

// SetPlaylistItemId gets a reference to the given string and assigns it to the PlaylistItemId field.
func (o *ReadyRequestDto) SetPlaylistItemId(v string) {
	o.PlaylistItemId = &v
}

func (o ReadyRequestDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReadyRequestDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.When) {
		toSerialize["When"] = o.When
	}
	if !IsNil(o.PositionTicks) {
		toSerialize["PositionTicks"] = o.PositionTicks
	}
	if !IsNil(o.IsPlaying) {
		toSerialize["IsPlaying"] = o.IsPlaying
	}
	if !IsNil(o.PlaylistItemId) {
		toSerialize["PlaylistItemId"] = o.PlaylistItemId
	}
	return toSerialize, nil
}

type NullableReadyRequestDto struct {
	value *ReadyRequestDto
	isSet bool
}

func (v NullableReadyRequestDto) Get() *ReadyRequestDto {
	return v.value
}

func (v *NullableReadyRequestDto) Set(val *ReadyRequestDto) {
	v.value = val
	v.isSet = true
}

func (v NullableReadyRequestDto) IsSet() bool {
	return v.isSet
}

func (v *NullableReadyRequestDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadyRequestDto(val *ReadyRequestDto) *NullableReadyRequestDto {
	return &NullableReadyRequestDto{value: val, isSet: true}
}

func (v NullableReadyRequestDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadyRequestDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


