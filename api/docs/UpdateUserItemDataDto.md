# UpdateUserItemDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | Pointer to **NullableFloat64** | Gets or sets the rating. | [optional] 
**PlayedPercentage** | Pointer to **NullableFloat64** | Gets or sets the played percentage. | [optional] 
**UnplayedItemCount** | Pointer to **NullableInt32** | Gets or sets the unplayed item count. | [optional] 
**PlaybackPositionTicks** | Pointer to **NullableInt64** | Gets or sets the playback position ticks. | [optional] 
**PlayCount** | Pointer to **NullableInt32** | Gets or sets the play count. | [optional] 
**IsFavorite** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is favorite. | [optional] 
**Likes** | Pointer to **NullableBool** | Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UpdateUserItemDataDto is likes. | [optional] 
**LastPlayedDate** | Pointer to **NullableTime** | Gets or sets the last played date. | [optional] 
**Played** | Pointer to **NullableBool** | Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UserItemDataDto is played. | [optional] 
**Key** | Pointer to **NullableString** | Gets or sets the key. | [optional] 
**ItemId** | Pointer to **NullableString** | Gets or sets the item identifier. | [optional] 

## Methods

### NewUpdateUserItemDataDto

`func NewUpdateUserItemDataDto() *UpdateUserItemDataDto`

NewUpdateUserItemDataDto instantiates a new UpdateUserItemDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserItemDataDtoWithDefaults

`func NewUpdateUserItemDataDtoWithDefaults() *UpdateUserItemDataDto`

NewUpdateUserItemDataDtoWithDefaults instantiates a new UpdateUserItemDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *UpdateUserItemDataDto) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *UpdateUserItemDataDto) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *UpdateUserItemDataDto) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *UpdateUserItemDataDto) HasRating() bool`

HasRating returns a boolean if a field has been set.

### SetRatingNil

`func (o *UpdateUserItemDataDto) SetRatingNil(b bool)`

 SetRatingNil sets the value for Rating to be an explicit nil

### UnsetRating
`func (o *UpdateUserItemDataDto) UnsetRating()`

UnsetRating ensures that no value is present for Rating, not even an explicit nil
### GetPlayedPercentage

`func (o *UpdateUserItemDataDto) GetPlayedPercentage() float64`

GetPlayedPercentage returns the PlayedPercentage field if non-nil, zero value otherwise.

### GetPlayedPercentageOk

`func (o *UpdateUserItemDataDto) GetPlayedPercentageOk() (*float64, bool)`

GetPlayedPercentageOk returns a tuple with the PlayedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayedPercentage

`func (o *UpdateUserItemDataDto) SetPlayedPercentage(v float64)`

SetPlayedPercentage sets PlayedPercentage field to given value.

### HasPlayedPercentage

`func (o *UpdateUserItemDataDto) HasPlayedPercentage() bool`

HasPlayedPercentage returns a boolean if a field has been set.

### SetPlayedPercentageNil

`func (o *UpdateUserItemDataDto) SetPlayedPercentageNil(b bool)`

 SetPlayedPercentageNil sets the value for PlayedPercentage to be an explicit nil

### UnsetPlayedPercentage
`func (o *UpdateUserItemDataDto) UnsetPlayedPercentage()`

UnsetPlayedPercentage ensures that no value is present for PlayedPercentage, not even an explicit nil
### GetUnplayedItemCount

`func (o *UpdateUserItemDataDto) GetUnplayedItemCount() int32`

GetUnplayedItemCount returns the UnplayedItemCount field if non-nil, zero value otherwise.

### GetUnplayedItemCountOk

`func (o *UpdateUserItemDataDto) GetUnplayedItemCountOk() (*int32, bool)`

GetUnplayedItemCountOk returns a tuple with the UnplayedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnplayedItemCount

`func (o *UpdateUserItemDataDto) SetUnplayedItemCount(v int32)`

SetUnplayedItemCount sets UnplayedItemCount field to given value.

### HasUnplayedItemCount

`func (o *UpdateUserItemDataDto) HasUnplayedItemCount() bool`

HasUnplayedItemCount returns a boolean if a field has been set.

### SetUnplayedItemCountNil

`func (o *UpdateUserItemDataDto) SetUnplayedItemCountNil(b bool)`

 SetUnplayedItemCountNil sets the value for UnplayedItemCount to be an explicit nil

### UnsetUnplayedItemCount
`func (o *UpdateUserItemDataDto) UnsetUnplayedItemCount()`

UnsetUnplayedItemCount ensures that no value is present for UnplayedItemCount, not even an explicit nil
### GetPlaybackPositionTicks

`func (o *UpdateUserItemDataDto) GetPlaybackPositionTicks() int64`

GetPlaybackPositionTicks returns the PlaybackPositionTicks field if non-nil, zero value otherwise.

### GetPlaybackPositionTicksOk

`func (o *UpdateUserItemDataDto) GetPlaybackPositionTicksOk() (*int64, bool)`

GetPlaybackPositionTicksOk returns a tuple with the PlaybackPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackPositionTicks

`func (o *UpdateUserItemDataDto) SetPlaybackPositionTicks(v int64)`

SetPlaybackPositionTicks sets PlaybackPositionTicks field to given value.

### HasPlaybackPositionTicks

`func (o *UpdateUserItemDataDto) HasPlaybackPositionTicks() bool`

HasPlaybackPositionTicks returns a boolean if a field has been set.

### SetPlaybackPositionTicksNil

`func (o *UpdateUserItemDataDto) SetPlaybackPositionTicksNil(b bool)`

 SetPlaybackPositionTicksNil sets the value for PlaybackPositionTicks to be an explicit nil

### UnsetPlaybackPositionTicks
`func (o *UpdateUserItemDataDto) UnsetPlaybackPositionTicks()`

UnsetPlaybackPositionTicks ensures that no value is present for PlaybackPositionTicks, not even an explicit nil
### GetPlayCount

`func (o *UpdateUserItemDataDto) GetPlayCount() int32`

GetPlayCount returns the PlayCount field if non-nil, zero value otherwise.

### GetPlayCountOk

`func (o *UpdateUserItemDataDto) GetPlayCountOk() (*int32, bool)`

GetPlayCountOk returns a tuple with the PlayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayCount

`func (o *UpdateUserItemDataDto) SetPlayCount(v int32)`

SetPlayCount sets PlayCount field to given value.

### HasPlayCount

`func (o *UpdateUserItemDataDto) HasPlayCount() bool`

HasPlayCount returns a boolean if a field has been set.

### SetPlayCountNil

`func (o *UpdateUserItemDataDto) SetPlayCountNil(b bool)`

 SetPlayCountNil sets the value for PlayCount to be an explicit nil

### UnsetPlayCount
`func (o *UpdateUserItemDataDto) UnsetPlayCount()`

UnsetPlayCount ensures that no value is present for PlayCount, not even an explicit nil
### GetIsFavorite

`func (o *UpdateUserItemDataDto) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *UpdateUserItemDataDto) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *UpdateUserItemDataDto) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *UpdateUserItemDataDto) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### SetIsFavoriteNil

`func (o *UpdateUserItemDataDto) SetIsFavoriteNil(b bool)`

 SetIsFavoriteNil sets the value for IsFavorite to be an explicit nil

### UnsetIsFavorite
`func (o *UpdateUserItemDataDto) UnsetIsFavorite()`

UnsetIsFavorite ensures that no value is present for IsFavorite, not even an explicit nil
### GetLikes

`func (o *UpdateUserItemDataDto) GetLikes() bool`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *UpdateUserItemDataDto) GetLikesOk() (*bool, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *UpdateUserItemDataDto) SetLikes(v bool)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *UpdateUserItemDataDto) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### SetLikesNil

`func (o *UpdateUserItemDataDto) SetLikesNil(b bool)`

 SetLikesNil sets the value for Likes to be an explicit nil

### UnsetLikes
`func (o *UpdateUserItemDataDto) UnsetLikes()`

UnsetLikes ensures that no value is present for Likes, not even an explicit nil
### GetLastPlayedDate

`func (o *UpdateUserItemDataDto) GetLastPlayedDate() time.Time`

GetLastPlayedDate returns the LastPlayedDate field if non-nil, zero value otherwise.

### GetLastPlayedDateOk

`func (o *UpdateUserItemDataDto) GetLastPlayedDateOk() (*time.Time, bool)`

GetLastPlayedDateOk returns a tuple with the LastPlayedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlayedDate

`func (o *UpdateUserItemDataDto) SetLastPlayedDate(v time.Time)`

SetLastPlayedDate sets LastPlayedDate field to given value.

### HasLastPlayedDate

`func (o *UpdateUserItemDataDto) HasLastPlayedDate() bool`

HasLastPlayedDate returns a boolean if a field has been set.

### SetLastPlayedDateNil

`func (o *UpdateUserItemDataDto) SetLastPlayedDateNil(b bool)`

 SetLastPlayedDateNil sets the value for LastPlayedDate to be an explicit nil

### UnsetLastPlayedDate
`func (o *UpdateUserItemDataDto) UnsetLastPlayedDate()`

UnsetLastPlayedDate ensures that no value is present for LastPlayedDate, not even an explicit nil
### GetPlayed

`func (o *UpdateUserItemDataDto) GetPlayed() bool`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *UpdateUserItemDataDto) GetPlayedOk() (*bool, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *UpdateUserItemDataDto) SetPlayed(v bool)`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *UpdateUserItemDataDto) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### SetPlayedNil

`func (o *UpdateUserItemDataDto) SetPlayedNil(b bool)`

 SetPlayedNil sets the value for Played to be an explicit nil

### UnsetPlayed
`func (o *UpdateUserItemDataDto) UnsetPlayed()`

UnsetPlayed ensures that no value is present for Played, not even an explicit nil
### GetKey

`func (o *UpdateUserItemDataDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateUserItemDataDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateUserItemDataDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateUserItemDataDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *UpdateUserItemDataDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *UpdateUserItemDataDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetItemId

`func (o *UpdateUserItemDataDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *UpdateUserItemDataDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *UpdateUserItemDataDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *UpdateUserItemDataDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *UpdateUserItemDataDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *UpdateUserItemDataDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


