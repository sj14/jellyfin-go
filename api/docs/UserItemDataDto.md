# UserItemDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | Pointer to **NullableFloat64** | Gets or sets the rating. | [optional] 
**PlayedPercentage** | Pointer to **NullableFloat64** | Gets or sets the played percentage. | [optional] 
**UnplayedItemCount** | Pointer to **NullableInt32** | Gets or sets the unplayed item count. | [optional] 
**PlaybackPositionTicks** | Pointer to **int64** | Gets or sets the playback position ticks. | [optional] 
**PlayCount** | Pointer to **int32** | Gets or sets the play count. | [optional] 
**IsFavorite** | Pointer to **bool** | Gets or sets a value indicating whether this instance is favorite. | [optional] 
**Likes** | Pointer to **NullableBool** | Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UserItemDataDto is likes. | [optional] 
**LastPlayedDate** | Pointer to **NullableTime** | Gets or sets the last played date. | [optional] 
**Played** | Pointer to **bool** | Gets or sets a value indicating whether this MediaBrowser.Model.Dto.UserItemDataDto is played. | [optional] 
**Key** | Pointer to **NullableString** | Gets or sets the key. | [optional] 
**ItemId** | Pointer to **NullableString** | Gets or sets the item identifier. | [optional] 

## Methods

### NewUserItemDataDto

`func NewUserItemDataDto() *UserItemDataDto`

NewUserItemDataDto instantiates a new UserItemDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserItemDataDtoWithDefaults

`func NewUserItemDataDtoWithDefaults() *UserItemDataDto`

NewUserItemDataDtoWithDefaults instantiates a new UserItemDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *UserItemDataDto) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *UserItemDataDto) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *UserItemDataDto) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *UserItemDataDto) HasRating() bool`

HasRating returns a boolean if a field has been set.

### SetRatingNil

`func (o *UserItemDataDto) SetRatingNil(b bool)`

 SetRatingNil sets the value for Rating to be an explicit nil

### UnsetRating
`func (o *UserItemDataDto) UnsetRating()`

UnsetRating ensures that no value is present for Rating, not even an explicit nil
### GetPlayedPercentage

`func (o *UserItemDataDto) GetPlayedPercentage() float64`

GetPlayedPercentage returns the PlayedPercentage field if non-nil, zero value otherwise.

### GetPlayedPercentageOk

`func (o *UserItemDataDto) GetPlayedPercentageOk() (*float64, bool)`

GetPlayedPercentageOk returns a tuple with the PlayedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayedPercentage

`func (o *UserItemDataDto) SetPlayedPercentage(v float64)`

SetPlayedPercentage sets PlayedPercentage field to given value.

### HasPlayedPercentage

`func (o *UserItemDataDto) HasPlayedPercentage() bool`

HasPlayedPercentage returns a boolean if a field has been set.

### SetPlayedPercentageNil

`func (o *UserItemDataDto) SetPlayedPercentageNil(b bool)`

 SetPlayedPercentageNil sets the value for PlayedPercentage to be an explicit nil

### UnsetPlayedPercentage
`func (o *UserItemDataDto) UnsetPlayedPercentage()`

UnsetPlayedPercentage ensures that no value is present for PlayedPercentage, not even an explicit nil
### GetUnplayedItemCount

`func (o *UserItemDataDto) GetUnplayedItemCount() int32`

GetUnplayedItemCount returns the UnplayedItemCount field if non-nil, zero value otherwise.

### GetUnplayedItemCountOk

`func (o *UserItemDataDto) GetUnplayedItemCountOk() (*int32, bool)`

GetUnplayedItemCountOk returns a tuple with the UnplayedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnplayedItemCount

`func (o *UserItemDataDto) SetUnplayedItemCount(v int32)`

SetUnplayedItemCount sets UnplayedItemCount field to given value.

### HasUnplayedItemCount

`func (o *UserItemDataDto) HasUnplayedItemCount() bool`

HasUnplayedItemCount returns a boolean if a field has been set.

### SetUnplayedItemCountNil

`func (o *UserItemDataDto) SetUnplayedItemCountNil(b bool)`

 SetUnplayedItemCountNil sets the value for UnplayedItemCount to be an explicit nil

### UnsetUnplayedItemCount
`func (o *UserItemDataDto) UnsetUnplayedItemCount()`

UnsetUnplayedItemCount ensures that no value is present for UnplayedItemCount, not even an explicit nil
### GetPlaybackPositionTicks

`func (o *UserItemDataDto) GetPlaybackPositionTicks() int64`

GetPlaybackPositionTicks returns the PlaybackPositionTicks field if non-nil, zero value otherwise.

### GetPlaybackPositionTicksOk

`func (o *UserItemDataDto) GetPlaybackPositionTicksOk() (*int64, bool)`

GetPlaybackPositionTicksOk returns a tuple with the PlaybackPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackPositionTicks

`func (o *UserItemDataDto) SetPlaybackPositionTicks(v int64)`

SetPlaybackPositionTicks sets PlaybackPositionTicks field to given value.

### HasPlaybackPositionTicks

`func (o *UserItemDataDto) HasPlaybackPositionTicks() bool`

HasPlaybackPositionTicks returns a boolean if a field has been set.

### GetPlayCount

`func (o *UserItemDataDto) GetPlayCount() int32`

GetPlayCount returns the PlayCount field if non-nil, zero value otherwise.

### GetPlayCountOk

`func (o *UserItemDataDto) GetPlayCountOk() (*int32, bool)`

GetPlayCountOk returns a tuple with the PlayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayCount

`func (o *UserItemDataDto) SetPlayCount(v int32)`

SetPlayCount sets PlayCount field to given value.

### HasPlayCount

`func (o *UserItemDataDto) HasPlayCount() bool`

HasPlayCount returns a boolean if a field has been set.

### GetIsFavorite

`func (o *UserItemDataDto) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *UserItemDataDto) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *UserItemDataDto) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *UserItemDataDto) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetLikes

`func (o *UserItemDataDto) GetLikes() bool`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *UserItemDataDto) GetLikesOk() (*bool, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *UserItemDataDto) SetLikes(v bool)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *UserItemDataDto) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### SetLikesNil

`func (o *UserItemDataDto) SetLikesNil(b bool)`

 SetLikesNil sets the value for Likes to be an explicit nil

### UnsetLikes
`func (o *UserItemDataDto) UnsetLikes()`

UnsetLikes ensures that no value is present for Likes, not even an explicit nil
### GetLastPlayedDate

`func (o *UserItemDataDto) GetLastPlayedDate() time.Time`

GetLastPlayedDate returns the LastPlayedDate field if non-nil, zero value otherwise.

### GetLastPlayedDateOk

`func (o *UserItemDataDto) GetLastPlayedDateOk() (*time.Time, bool)`

GetLastPlayedDateOk returns a tuple with the LastPlayedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlayedDate

`func (o *UserItemDataDto) SetLastPlayedDate(v time.Time)`

SetLastPlayedDate sets LastPlayedDate field to given value.

### HasLastPlayedDate

`func (o *UserItemDataDto) HasLastPlayedDate() bool`

HasLastPlayedDate returns a boolean if a field has been set.

### SetLastPlayedDateNil

`func (o *UserItemDataDto) SetLastPlayedDateNil(b bool)`

 SetLastPlayedDateNil sets the value for LastPlayedDate to be an explicit nil

### UnsetLastPlayedDate
`func (o *UserItemDataDto) UnsetLastPlayedDate()`

UnsetLastPlayedDate ensures that no value is present for LastPlayedDate, not even an explicit nil
### GetPlayed

`func (o *UserItemDataDto) GetPlayed() bool`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *UserItemDataDto) GetPlayedOk() (*bool, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *UserItemDataDto) SetPlayed(v bool)`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *UserItemDataDto) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### GetKey

`func (o *UserItemDataDto) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserItemDataDto) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserItemDataDto) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UserItemDataDto) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *UserItemDataDto) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *UserItemDataDto) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetItemId

`func (o *UserItemDataDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *UserItemDataDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *UserItemDataDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *UserItemDataDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *UserItemDataDto) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *UserItemDataDto) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


