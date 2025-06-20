/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the UpdateUserItemDataDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserItemDataDto{}

// UpdateUserItemDataDto This is used by the api to get information about a item user data.
type UpdateUserItemDataDto struct {
	// Gets or sets the rating.
	Rating NullableFloat64 `json:"Rating,omitempty"`
	// Gets or sets the played percentage.
	PlayedPercentage NullableFloat64 `json:"PlayedPercentage,omitempty"`
	// Gets or sets the unplayed item count.
	UnplayedItemCount NullableInt32 `json:"UnplayedItemCount,omitempty"`
	// Gets or sets the playback position ticks.
	PlaybackPositionTicks NullableInt64 `json:"PlaybackPositionTicks,omitempty"`
	// Gets or sets the play count.
	PlayCount NullableInt32 `json:"PlayCount,omitempty"`
	// Gets or sets a value indicating whether this instance is favorite.
	IsFavorite NullableBool `json:"IsFavorite,omitempty"`
	// Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UpdateUserItemDataDto is likes.
	Likes NullableBool `json:"Likes,omitempty"`
	// Gets or sets the last played date.
	LastPlayedDate NullableTime `json:"LastPlayedDate,omitempty"`
	// Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UserItemDataDto is played.
	Played NullableBool `json:"Played,omitempty"`
	// Gets or sets the key.
	Key NullableString `json:"Key,omitempty"`
	// Gets or sets the item identifier.
	ItemId NullableString `json:"ItemId,omitempty"`
}

// NewUpdateUserItemDataDto instantiates a new UpdateUserItemDataDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserItemDataDto() *UpdateUserItemDataDto {
	this := UpdateUserItemDataDto{}
	return &this
}

// NewUpdateUserItemDataDtoWithDefaults instantiates a new UpdateUserItemDataDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserItemDataDtoWithDefaults() *UpdateUserItemDataDto {
	this := UpdateUserItemDataDto{}
	return &this
}

// GetRating returns the Rating field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserItemDataDto) GetRating() float64 {
	if o == nil || IsNil(o.Rating.Get()) {
		var ret float64
		return ret
	}
	return *o.Rating.Get()
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserItemDataDto) GetRatingOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rating.Get(), o.Rating.IsSet()
}

// HasRating returns a boolean if a field has been set.
func (o *UpdateUserItemDataDto) HasRating() bool {
	if o != nil && o.Rating.IsSet() {
		return true
	}

	return false
}

// SetRating gets a reference to the given NullableFloat64 and assigns it to the Rating field.
func (o *UpdateUserItemDataDto) SetRating(v float64) {
	o.Rating.Set(&v)
}
// SetRatingNil sets the value for Rating to be an explicit nil
func (o *UpdateUserItemDataDto) SetRatingNil() {
	o.Rating.Set(nil)
}

// UnsetRating ensures that no value is present for Rating, not even an explicit nil
func (o *UpdateUserItemDataDto) UnsetRating() {
	o.Rating.Unset()
}

// GetPlayedPercentage returns the PlayedPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserItemDataDto) GetPlayedPercentage() float64 {
	if o == nil || IsNil(o.PlayedPercentage.Get()) {
		var ret float64
		return ret
	}
	return *o.PlayedPercentage.Get()
}

// GetPlayedPercentageOk returns a tuple with the PlayedPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserItemDataDto) GetPlayedPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlayedPercentage.Get(), o.PlayedPercentage.IsSet()
}

// HasPlayedPercentage returns a boolean if a field has been set.
func (o *UpdateUserItemDataDto) HasPlayedPercentage() bool {
	if o != nil && o.PlayedPercentage.IsSet() {
		return true
	}

	return false
}

// SetPlayedPercentage gets a reference to the given NullableFloat64 and assigns it to the PlayedPercentage field.
func (o *UpdateUserItemDataDto) SetPlayedPercentage(v float64) {
	o.PlayedPercentage.Set(&v)
}
// SetPlayedPercentageNil sets the value for PlayedPercentage to be an explicit nil
func (o *UpdateUserItemDataDto) SetPlayedPercentageNil() {
	o.PlayedPercentage.Set(nil)
}

// UnsetPlayedPercentage ensures that no value is present for PlayedPercentage, not even an explicit nil
func (o *UpdateUserItemDataDto) UnsetPlayedPercentage() {
	o.PlayedPercentage.Unset()
}

// GetUnplayedItemCount returns the UnplayedItemCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserItemDataDto) GetUnplayedItemCount() int32 {
	if o == nil || IsNil(o.UnplayedItemCount.Get()) {
		var ret int32
		return ret
	}
	return *o.UnplayedItemCount.Get()
}

// GetUnplayedItemCountOk returns a tuple with the UnplayedItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserItemDataDto) GetUnplayedItemCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnplayedItemCount.Get(), o.UnplayedItemCount.IsSet()
}

// HasUnplayedItemCount returns a boolean if a field has been set.
func (o *UpdateUserItemDataDto) HasUnplayedItemCount() bool {
	if o != nil && o.UnplayedItemCount.IsSet() {
		return true
	}

	return false
}

// SetUnplayedItemCount gets a reference to the given NullableInt32 and assigns it to the UnplayedItemCount field.
func (o *UpdateUserItemDataDto) SetUnplayedItemCount(v int32) {
	o.UnplayedItemCount.Set(&v)
}
// SetUnplayedItemCountNil sets the value for UnplayedItemCount to be an explicit nil
func (o *UpdateUserItemDataDto) SetUnplayedItemCountNil() {
	o.UnplayedItemCount.Set(nil)
}

// UnsetUnplayedItemCount ensures that no value is present for UnplayedItemCount, not even an explicit nil
func (o *UpdateUserItemDataDto) UnsetUnplayedItemCount() {
	o.UnplayedItemCount.Unset()
}

// GetPlaybackPositionTicks returns the PlaybackPositionTicks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserItemDataDto) GetPlaybackPositionTicks() int64 {
	if o == nil || IsNil(o.PlaybackPositionTicks.Get()) {
		var ret int64
		return ret
	}
	return *o.PlaybackPositionTicks.Get()
}

// GetPlaybackPositionTicksOk returns a tuple with the PlaybackPositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserItemDataDto) GetPlaybackPositionTicksOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaybackPositionTicks.Get(), o.PlaybackPositionTicks.IsSet()
}

// HasPlaybackPositionTicks returns a boolean if a field has been set.
func (o *UpdateUserItemDataDto) HasPlaybackPositionTicks() bool {
	if o != nil && o.PlaybackPositionTicks.IsSet() {
		return true
	}

	return false
}

// SetPlaybackPositionTicks gets a reference to the given NullableInt64 and assigns it to the PlaybackPositionTicks field.
func (o *UpdateUserItemDataDto) SetPlaybackPositionTicks(v int64) {
	o.PlaybackPositionTicks.Set(&v)
}
// SetPlaybackPositionTicksNil sets the value for PlaybackPositionTicks to be an explicit nil
func (o *UpdateUserItemDataDto) SetPlaybackPositionTicksNil() {
	o.PlaybackPositionTicks.Set(nil)
}

// UnsetPlaybackPositionTicks ensures that no value is present for PlaybackPositionTicks, not even an explicit nil
func (o *UpdateUserItemDataDto) UnsetPlaybackPositionTicks() {
	o.PlaybackPositionTicks.Unset()
}

// GetPlayCount returns the PlayCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserItemDataDto) GetPlayCount() int32 {
	if o == nil || IsNil(o.PlayCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PlayCount.Get()
}

// GetPlayCountOk returns a tuple with the PlayCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserItemDataDto) GetPlayCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlayCount.Get(), o.PlayCount.IsSet()
}

// HasPlayCount returns a boolean if a field has been set.
func (o *UpdateUserItemDataDto) HasPlayCount() bool {
	if o != nil && o.PlayCount.IsSet() {
		return true
	}

	return false
}

// SetPlayCount gets a reference to the given NullableInt32 and assigns it to the PlayCount field.
func (o *UpdateUserItemDataDto) SetPlayCount(v int32) {
	o.PlayCount.Set(&v)
}
// SetPlayCountNil sets the value for PlayCount to be an explicit nil
func (o *UpdateUserItemDataDto) SetPlayCountNil() {
	o.PlayCount.Set(nil)
}

// UnsetPlayCount ensures that no value is present for PlayCount, not even an explicit nil
func (o *UpdateUserItemDataDto) UnsetPlayCount() {
	o.PlayCount.Unset()
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserItemDataDto) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite.Get()) {
		var ret bool
		return ret
	}
	return *o.IsFavorite.Get()
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserItemDataDto) GetIsFavoriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsFavorite.Get(), o.IsFavorite.IsSet()
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *UpdateUserItemDataDto) HasIsFavorite() bool {
	if o != nil && o.IsFavorite.IsSet() {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given NullableBool and assigns it to the IsFavorite field.
func (o *UpdateUserItemDataDto) SetIsFavorite(v bool) {
	o.IsFavorite.Set(&v)
}
// SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil
func (o *UpdateUserItemDataDto) SetIsFavoriteNil() {
	o.IsFavorite.Set(nil)
}

// UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
func (o *UpdateUserItemDataDto) UnsetIsFavorite() {
	o.IsFavorite.Unset()
}

// GetLikes returns the Likes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserItemDataDto) GetLikes() bool {
	if o == nil || IsNil(o.Likes.Get()) {
		var ret bool
		return ret
	}
	return *o.Likes.Get()
}

// GetLikesOk returns a tuple with the Likes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserItemDataDto) GetLikesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Likes.Get(), o.Likes.IsSet()
}

// HasLikes returns a boolean if a field has been set.
func (o *UpdateUserItemDataDto) HasLikes() bool {
	if o != nil && o.Likes.IsSet() {
		return true
	}

	return false
}

// SetLikes gets a reference to the given NullableBool and assigns it to the Likes field.
func (o *UpdateUserItemDataDto) SetLikes(v bool) {
	o.Likes.Set(&v)
}
// SetLikesNil sets the value for Likes to be an explicit nil
func (o *UpdateUserItemDataDto) SetLikesNil() {
	o.Likes.Set(nil)
}

// UnsetLikes ensures that no value is present for Likes, not even an explicit nil
func (o *UpdateUserItemDataDto) UnsetLikes() {
	o.Likes.Unset()
}

// GetLastPlayedDate returns the LastPlayedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserItemDataDto) GetLastPlayedDate() time.Time {
	if o == nil || IsNil(o.LastPlayedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastPlayedDate.Get()
}

// GetLastPlayedDateOk returns a tuple with the LastPlayedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserItemDataDto) GetLastPlayedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastPlayedDate.Get(), o.LastPlayedDate.IsSet()
}

// HasLastPlayedDate returns a boolean if a field has been set.
func (o *UpdateUserItemDataDto) HasLastPlayedDate() bool {
	if o != nil && o.LastPlayedDate.IsSet() {
		return true
	}

	return false
}

// SetLastPlayedDate gets a reference to the given NullableTime and assigns it to the LastPlayedDate field.
func (o *UpdateUserItemDataDto) SetLastPlayedDate(v time.Time) {
	o.LastPlayedDate.Set(&v)
}
// SetLastPlayedDateNil sets the value for LastPlayedDate to be an explicit nil
func (o *UpdateUserItemDataDto) SetLastPlayedDateNil() {
	o.LastPlayedDate.Set(nil)
}

// UnsetLastPlayedDate ensures that no value is present for LastPlayedDate, not even an explicit nil
func (o *UpdateUserItemDataDto) UnsetLastPlayedDate() {
	o.LastPlayedDate.Unset()
}

// GetPlayed returns the Played field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserItemDataDto) GetPlayed() bool {
	if o == nil || IsNil(o.Played.Get()) {
		var ret bool
		return ret
	}
	return *o.Played.Get()
}

// GetPlayedOk returns a tuple with the Played field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserItemDataDto) GetPlayedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Played.Get(), o.Played.IsSet()
}

// HasPlayed returns a boolean if a field has been set.
func (o *UpdateUserItemDataDto) HasPlayed() bool {
	if o != nil && o.Played.IsSet() {
		return true
	}

	return false
}

// SetPlayed gets a reference to the given NullableBool and assigns it to the Played field.
func (o *UpdateUserItemDataDto) SetPlayed(v bool) {
	o.Played.Set(&v)
}
// SetPlayedNil sets the value for Played to be an explicit nil
func (o *UpdateUserItemDataDto) SetPlayedNil() {
	o.Played.Set(nil)
}

// UnsetPlayed ensures that no value is present for Played, not even an explicit nil
func (o *UpdateUserItemDataDto) UnsetPlayed() {
	o.Played.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserItemDataDto) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserItemDataDto) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateUserItemDataDto) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *UpdateUserItemDataDto) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *UpdateUserItemDataDto) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *UpdateUserItemDataDto) UnsetKey() {
	o.Key.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserItemDataDto) GetItemId() string {
	if o == nil || IsNil(o.ItemId.Get()) {
		var ret string
		return ret
	}
	return *o.ItemId.Get()
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserItemDataDto) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemId.Get(), o.ItemId.IsSet()
}

// HasItemId returns a boolean if a field has been set.
func (o *UpdateUserItemDataDto) HasItemId() bool {
	if o != nil && o.ItemId.IsSet() {
		return true
	}

	return false
}

// SetItemId gets a reference to the given NullableString and assigns it to the ItemId field.
func (o *UpdateUserItemDataDto) SetItemId(v string) {
	o.ItemId.Set(&v)
}
// SetItemIdNil sets the value for ItemId to be an explicit nil
func (o *UpdateUserItemDataDto) SetItemIdNil() {
	o.ItemId.Set(nil)
}

// UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
func (o *UpdateUserItemDataDto) UnsetItemId() {
	o.ItemId.Unset()
}

func (o UpdateUserItemDataDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserItemDataDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Rating.IsSet() {
		toSerialize["Rating"] = o.Rating.Get()
	}
	if o.PlayedPercentage.IsSet() {
		toSerialize["PlayedPercentage"] = o.PlayedPercentage.Get()
	}
	if o.UnplayedItemCount.IsSet() {
		toSerialize["UnplayedItemCount"] = o.UnplayedItemCount.Get()
	}
	if o.PlaybackPositionTicks.IsSet() {
		toSerialize["PlaybackPositionTicks"] = o.PlaybackPositionTicks.Get()
	}
	if o.PlayCount.IsSet() {
		toSerialize["PlayCount"] = o.PlayCount.Get()
	}
	if o.IsFavorite.IsSet() {
		toSerialize["IsFavorite"] = o.IsFavorite.Get()
	}
	if o.Likes.IsSet() {
		toSerialize["Likes"] = o.Likes.Get()
	}
	if o.LastPlayedDate.IsSet() {
		toSerialize["LastPlayedDate"] = o.LastPlayedDate.Get()
	}
	if o.Played.IsSet() {
		toSerialize["Played"] = o.Played.Get()
	}
	if o.Key.IsSet() {
		toSerialize["Key"] = o.Key.Get()
	}
	if o.ItemId.IsSet() {
		toSerialize["ItemId"] = o.ItemId.Get()
	}
	return toSerialize, nil
}

type NullableUpdateUserItemDataDto struct {
	value *UpdateUserItemDataDto
	isSet bool
}

func (v NullableUpdateUserItemDataDto) Get() *UpdateUserItemDataDto {
	return v.value
}

func (v *NullableUpdateUserItemDataDto) Set(val *UpdateUserItemDataDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserItemDataDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserItemDataDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserItemDataDto(val *UpdateUserItemDataDto) *NullableUpdateUserItemDataDto {
	return &NullableUpdateUserItemDataDto{value: val, isSet: true}
}

func (v NullableUpdateUserItemDataDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserItemDataDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


