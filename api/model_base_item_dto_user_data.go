/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the BaseItemDtoUserData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseItemDtoUserData{}

// BaseItemDtoUserData Gets or sets the user data for this item based on the user it's being requested for.
type BaseItemDtoUserData struct {
	// Gets or sets the rating.
	Rating NullableFloat64 `json:"Rating,omitempty"`
	// Gets or sets the played percentage.
	PlayedPercentage NullableFloat64 `json:"PlayedPercentage,omitempty"`
	// Gets or sets the unplayed item count.
	UnplayedItemCount NullableInt32 `json:"UnplayedItemCount,omitempty"`
	// Gets or sets the playback position ticks.
	PlaybackPositionTicks *int64 `json:"PlaybackPositionTicks,omitempty"`
	// Gets or sets the play count.
	PlayCount *int32 `json:"PlayCount,omitempty"`
	// Gets or sets a value indicating whether this instance is favorite.
	IsFavorite *bool `json:"IsFavorite,omitempty"`
	// Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UserItemDataDto is likes.
	Likes NullableBool `json:"Likes,omitempty"`
	// Gets or sets the last played date.
	LastPlayedDate NullableTime `json:"LastPlayedDate,omitempty"`
	// Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UserItemDataDto is played.
	Played *bool `json:"Played,omitempty"`
	// Gets or sets the key.
	Key NullableString `json:"Key,omitempty"`
	// Gets or sets the item identifier.
	ItemId NullableString `json:"ItemId,omitempty"`
}

// NewBaseItemDtoUserData instantiates a new BaseItemDtoUserData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseItemDtoUserData() *BaseItemDtoUserData {
	this := BaseItemDtoUserData{}
	return &this
}

// NewBaseItemDtoUserDataWithDefaults instantiates a new BaseItemDtoUserData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseItemDtoUserDataWithDefaults() *BaseItemDtoUserData {
	this := BaseItemDtoUserData{}
	return &this
}

// GetRating returns the Rating field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemDtoUserData) GetRating() float64 {
	if o == nil || IsNil(o.Rating.Get()) {
		var ret float64
		return ret
	}
	return *o.Rating.Get()
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemDtoUserData) GetRatingOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rating.Get(), o.Rating.IsSet()
}

// HasRating returns a boolean if a field has been set.
func (o *BaseItemDtoUserData) HasRating() bool {
	if o != nil && o.Rating.IsSet() {
		return true
	}

	return false
}

// SetRating gets a reference to the given NullableFloat64 and assigns it to the Rating field.
func (o *BaseItemDtoUserData) SetRating(v float64) {
	o.Rating.Set(&v)
}
// SetRatingNil sets the value for Rating to be an explicit nil
func (o *BaseItemDtoUserData) SetRatingNil() {
	o.Rating.Set(nil)
}

// UnsetRating ensures that no value is present for Rating, not even an explicit nil
func (o *BaseItemDtoUserData) UnsetRating() {
	o.Rating.Unset()
}

// GetPlayedPercentage returns the PlayedPercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemDtoUserData) GetPlayedPercentage() float64 {
	if o == nil || IsNil(o.PlayedPercentage.Get()) {
		var ret float64
		return ret
	}
	return *o.PlayedPercentage.Get()
}

// GetPlayedPercentageOk returns a tuple with the PlayedPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemDtoUserData) GetPlayedPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlayedPercentage.Get(), o.PlayedPercentage.IsSet()
}

// HasPlayedPercentage returns a boolean if a field has been set.
func (o *BaseItemDtoUserData) HasPlayedPercentage() bool {
	if o != nil && o.PlayedPercentage.IsSet() {
		return true
	}

	return false
}

// SetPlayedPercentage gets a reference to the given NullableFloat64 and assigns it to the PlayedPercentage field.
func (o *BaseItemDtoUserData) SetPlayedPercentage(v float64) {
	o.PlayedPercentage.Set(&v)
}
// SetPlayedPercentageNil sets the value for PlayedPercentage to be an explicit nil
func (o *BaseItemDtoUserData) SetPlayedPercentageNil() {
	o.PlayedPercentage.Set(nil)
}

// UnsetPlayedPercentage ensures that no value is present for PlayedPercentage, not even an explicit nil
func (o *BaseItemDtoUserData) UnsetPlayedPercentage() {
	o.PlayedPercentage.Unset()
}

// GetUnplayedItemCount returns the UnplayedItemCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemDtoUserData) GetUnplayedItemCount() int32 {
	if o == nil || IsNil(o.UnplayedItemCount.Get()) {
		var ret int32
		return ret
	}
	return *o.UnplayedItemCount.Get()
}

// GetUnplayedItemCountOk returns a tuple with the UnplayedItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemDtoUserData) GetUnplayedItemCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnplayedItemCount.Get(), o.UnplayedItemCount.IsSet()
}

// HasUnplayedItemCount returns a boolean if a field has been set.
func (o *BaseItemDtoUserData) HasUnplayedItemCount() bool {
	if o != nil && o.UnplayedItemCount.IsSet() {
		return true
	}

	return false
}

// SetUnplayedItemCount gets a reference to the given NullableInt32 and assigns it to the UnplayedItemCount field.
func (o *BaseItemDtoUserData) SetUnplayedItemCount(v int32) {
	o.UnplayedItemCount.Set(&v)
}
// SetUnplayedItemCountNil sets the value for UnplayedItemCount to be an explicit nil
func (o *BaseItemDtoUserData) SetUnplayedItemCountNil() {
	o.UnplayedItemCount.Set(nil)
}

// UnsetUnplayedItemCount ensures that no value is present for UnplayedItemCount, not even an explicit nil
func (o *BaseItemDtoUserData) UnsetUnplayedItemCount() {
	o.UnplayedItemCount.Unset()
}

// GetPlaybackPositionTicks returns the PlaybackPositionTicks field value if set, zero value otherwise.
func (o *BaseItemDtoUserData) GetPlaybackPositionTicks() int64 {
	if o == nil || IsNil(o.PlaybackPositionTicks) {
		var ret int64
		return ret
	}
	return *o.PlaybackPositionTicks
}

// GetPlaybackPositionTicksOk returns a tuple with the PlaybackPositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoUserData) GetPlaybackPositionTicksOk() (*int64, bool) {
	if o == nil || IsNil(o.PlaybackPositionTicks) {
		return nil, false
	}
	return o.PlaybackPositionTicks, true
}

// HasPlaybackPositionTicks returns a boolean if a field has been set.
func (o *BaseItemDtoUserData) HasPlaybackPositionTicks() bool {
	if o != nil && !IsNil(o.PlaybackPositionTicks) {
		return true
	}

	return false
}

// SetPlaybackPositionTicks gets a reference to the given int64 and assigns it to the PlaybackPositionTicks field.
func (o *BaseItemDtoUserData) SetPlaybackPositionTicks(v int64) {
	o.PlaybackPositionTicks = &v
}

// GetPlayCount returns the PlayCount field value if set, zero value otherwise.
func (o *BaseItemDtoUserData) GetPlayCount() int32 {
	if o == nil || IsNil(o.PlayCount) {
		var ret int32
		return ret
	}
	return *o.PlayCount
}

// GetPlayCountOk returns a tuple with the PlayCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoUserData) GetPlayCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PlayCount) {
		return nil, false
	}
	return o.PlayCount, true
}

// HasPlayCount returns a boolean if a field has been set.
func (o *BaseItemDtoUserData) HasPlayCount() bool {
	if o != nil && !IsNil(o.PlayCount) {
		return true
	}

	return false
}

// SetPlayCount gets a reference to the given int32 and assigns it to the PlayCount field.
func (o *BaseItemDtoUserData) SetPlayCount(v int32) {
	o.PlayCount = &v
}

// GetIsFavorite returns the IsFavorite field value if set, zero value otherwise.
func (o *BaseItemDtoUserData) GetIsFavorite() bool {
	if o == nil || IsNil(o.IsFavorite) {
		var ret bool
		return ret
	}
	return *o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoUserData) GetIsFavoriteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFavorite) {
		return nil, false
	}
	return o.IsFavorite, true
}

// HasIsFavorite returns a boolean if a field has been set.
func (o *BaseItemDtoUserData) HasIsFavorite() bool {
	if o != nil && !IsNil(o.IsFavorite) {
		return true
	}

	return false
}

// SetIsFavorite gets a reference to the given bool and assigns it to the IsFavorite field.
func (o *BaseItemDtoUserData) SetIsFavorite(v bool) {
	o.IsFavorite = &v
}

// GetLikes returns the Likes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemDtoUserData) GetLikes() bool {
	if o == nil || IsNil(o.Likes.Get()) {
		var ret bool
		return ret
	}
	return *o.Likes.Get()
}

// GetLikesOk returns a tuple with the Likes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemDtoUserData) GetLikesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Likes.Get(), o.Likes.IsSet()
}

// HasLikes returns a boolean if a field has been set.
func (o *BaseItemDtoUserData) HasLikes() bool {
	if o != nil && o.Likes.IsSet() {
		return true
	}

	return false
}

// SetLikes gets a reference to the given NullableBool and assigns it to the Likes field.
func (o *BaseItemDtoUserData) SetLikes(v bool) {
	o.Likes.Set(&v)
}
// SetLikesNil sets the value for Likes to be an explicit nil
func (o *BaseItemDtoUserData) SetLikesNil() {
	o.Likes.Set(nil)
}

// UnsetLikes ensures that no value is present for Likes, not even an explicit nil
func (o *BaseItemDtoUserData) UnsetLikes() {
	o.Likes.Unset()
}

// GetLastPlayedDate returns the LastPlayedDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemDtoUserData) GetLastPlayedDate() time.Time {
	if o == nil || IsNil(o.LastPlayedDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastPlayedDate.Get()
}

// GetLastPlayedDateOk returns a tuple with the LastPlayedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemDtoUserData) GetLastPlayedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastPlayedDate.Get(), o.LastPlayedDate.IsSet()
}

// HasLastPlayedDate returns a boolean if a field has been set.
func (o *BaseItemDtoUserData) HasLastPlayedDate() bool {
	if o != nil && o.LastPlayedDate.IsSet() {
		return true
	}

	return false
}

// SetLastPlayedDate gets a reference to the given NullableTime and assigns it to the LastPlayedDate field.
func (o *BaseItemDtoUserData) SetLastPlayedDate(v time.Time) {
	o.LastPlayedDate.Set(&v)
}
// SetLastPlayedDateNil sets the value for LastPlayedDate to be an explicit nil
func (o *BaseItemDtoUserData) SetLastPlayedDateNil() {
	o.LastPlayedDate.Set(nil)
}

// UnsetLastPlayedDate ensures that no value is present for LastPlayedDate, not even an explicit nil
func (o *BaseItemDtoUserData) UnsetLastPlayedDate() {
	o.LastPlayedDate.Unset()
}

// GetPlayed returns the Played field value if set, zero value otherwise.
func (o *BaseItemDtoUserData) GetPlayed() bool {
	if o == nil || IsNil(o.Played) {
		var ret bool
		return ret
	}
	return *o.Played
}

// GetPlayedOk returns a tuple with the Played field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoUserData) GetPlayedOk() (*bool, bool) {
	if o == nil || IsNil(o.Played) {
		return nil, false
	}
	return o.Played, true
}

// HasPlayed returns a boolean if a field has been set.
func (o *BaseItemDtoUserData) HasPlayed() bool {
	if o != nil && !IsNil(o.Played) {
		return true
	}

	return false
}

// SetPlayed gets a reference to the given bool and assigns it to the Played field.
func (o *BaseItemDtoUserData) SetPlayed(v bool) {
	o.Played = &v
}

// GetKey returns the Key field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemDtoUserData) GetKey() string {
	if o == nil || IsNil(o.Key.Get()) {
		var ret string
		return ret
	}
	return *o.Key.Get()
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemDtoUserData) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key.Get(), o.Key.IsSet()
}

// HasKey returns a boolean if a field has been set.
func (o *BaseItemDtoUserData) HasKey() bool {
	if o != nil && o.Key.IsSet() {
		return true
	}

	return false
}

// SetKey gets a reference to the given NullableString and assigns it to the Key field.
func (o *BaseItemDtoUserData) SetKey(v string) {
	o.Key.Set(&v)
}
// SetKeyNil sets the value for Key to be an explicit nil
func (o *BaseItemDtoUserData) SetKeyNil() {
	o.Key.Set(nil)
}

// UnsetKey ensures that no value is present for Key, not even an explicit nil
func (o *BaseItemDtoUserData) UnsetKey() {
	o.Key.Unset()
}

// GetItemId returns the ItemId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BaseItemDtoUserData) GetItemId() string {
	if o == nil || IsNil(o.ItemId.Get()) {
		var ret string
		return ret
	}
	return *o.ItemId.Get()
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BaseItemDtoUserData) GetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemId.Get(), o.ItemId.IsSet()
}

// HasItemId returns a boolean if a field has been set.
func (o *BaseItemDtoUserData) HasItemId() bool {
	if o != nil && o.ItemId.IsSet() {
		return true
	}

	return false
}

// SetItemId gets a reference to the given NullableString and assigns it to the ItemId field.
func (o *BaseItemDtoUserData) SetItemId(v string) {
	o.ItemId.Set(&v)
}
// SetItemIdNil sets the value for ItemId to be an explicit nil
func (o *BaseItemDtoUserData) SetItemIdNil() {
	o.ItemId.Set(nil)
}

// UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
func (o *BaseItemDtoUserData) UnsetItemId() {
	o.ItemId.Unset()
}

func (o BaseItemDtoUserData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseItemDtoUserData) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PlaybackPositionTicks) {
		toSerialize["PlaybackPositionTicks"] = o.PlaybackPositionTicks
	}
	if !IsNil(o.PlayCount) {
		toSerialize["PlayCount"] = o.PlayCount
	}
	if !IsNil(o.IsFavorite) {
		toSerialize["IsFavorite"] = o.IsFavorite
	}
	if o.Likes.IsSet() {
		toSerialize["Likes"] = o.Likes.Get()
	}
	if o.LastPlayedDate.IsSet() {
		toSerialize["LastPlayedDate"] = o.LastPlayedDate.Get()
	}
	if !IsNil(o.Played) {
		toSerialize["Played"] = o.Played
	}
	if o.Key.IsSet() {
		toSerialize["Key"] = o.Key.Get()
	}
	if o.ItemId.IsSet() {
		toSerialize["ItemId"] = o.ItemId.Get()
	}
	return toSerialize, nil
}

type NullableBaseItemDtoUserData struct {
	value *BaseItemDtoUserData
	isSet bool
}

func (v NullableBaseItemDtoUserData) Get() *BaseItemDtoUserData {
	return v.value
}

func (v *NullableBaseItemDtoUserData) Set(val *BaseItemDtoUserData) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseItemDtoUserData) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseItemDtoUserData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseItemDtoUserData(val *BaseItemDtoUserData) *NullableBaseItemDtoUserData {
	return &NullableBaseItemDtoUserData{value: val, isSet: true}
}

func (v NullableBaseItemDtoUserData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseItemDtoUserData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


