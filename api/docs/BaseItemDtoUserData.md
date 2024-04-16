# BaseItemDtoUserData

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

### NewBaseItemDtoUserData

`func NewBaseItemDtoUserData() *BaseItemDtoUserData`

NewBaseItemDtoUserData instantiates a new BaseItemDtoUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseItemDtoUserDataWithDefaults

`func NewBaseItemDtoUserDataWithDefaults() *BaseItemDtoUserData`

NewBaseItemDtoUserDataWithDefaults instantiates a new BaseItemDtoUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *BaseItemDtoUserData) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *BaseItemDtoUserData) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *BaseItemDtoUserData) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *BaseItemDtoUserData) HasRating() bool`

HasRating returns a boolean if a field has been set.

### SetRatingNil

`func (o *BaseItemDtoUserData) SetRatingNil(b bool)`

 SetRatingNil sets the value for Rating to be an explicit nil

### UnsetRating
`func (o *BaseItemDtoUserData) UnsetRating()`

UnsetRating ensures that no value is present for Rating, not even an explicit nil
### GetPlayedPercentage

`func (o *BaseItemDtoUserData) GetPlayedPercentage() float64`

GetPlayedPercentage returns the PlayedPercentage field if non-nil, zero value otherwise.

### GetPlayedPercentageOk

`func (o *BaseItemDtoUserData) GetPlayedPercentageOk() (*float64, bool)`

GetPlayedPercentageOk returns a tuple with the PlayedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayedPercentage

`func (o *BaseItemDtoUserData) SetPlayedPercentage(v float64)`

SetPlayedPercentage sets PlayedPercentage field to given value.

### HasPlayedPercentage

`func (o *BaseItemDtoUserData) HasPlayedPercentage() bool`

HasPlayedPercentage returns a boolean if a field has been set.

### SetPlayedPercentageNil

`func (o *BaseItemDtoUserData) SetPlayedPercentageNil(b bool)`

 SetPlayedPercentageNil sets the value for PlayedPercentage to be an explicit nil

### UnsetPlayedPercentage
`func (o *BaseItemDtoUserData) UnsetPlayedPercentage()`

UnsetPlayedPercentage ensures that no value is present for PlayedPercentage, not even an explicit nil
### GetUnplayedItemCount

`func (o *BaseItemDtoUserData) GetUnplayedItemCount() int32`

GetUnplayedItemCount returns the UnplayedItemCount field if non-nil, zero value otherwise.

### GetUnplayedItemCountOk

`func (o *BaseItemDtoUserData) GetUnplayedItemCountOk() (*int32, bool)`

GetUnplayedItemCountOk returns a tuple with the UnplayedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnplayedItemCount

`func (o *BaseItemDtoUserData) SetUnplayedItemCount(v int32)`

SetUnplayedItemCount sets UnplayedItemCount field to given value.

### HasUnplayedItemCount

`func (o *BaseItemDtoUserData) HasUnplayedItemCount() bool`

HasUnplayedItemCount returns a boolean if a field has been set.

### SetUnplayedItemCountNil

`func (o *BaseItemDtoUserData) SetUnplayedItemCountNil(b bool)`

 SetUnplayedItemCountNil sets the value for UnplayedItemCount to be an explicit nil

### UnsetUnplayedItemCount
`func (o *BaseItemDtoUserData) UnsetUnplayedItemCount()`

UnsetUnplayedItemCount ensures that no value is present for UnplayedItemCount, not even an explicit nil
### GetPlaybackPositionTicks

`func (o *BaseItemDtoUserData) GetPlaybackPositionTicks() int64`

GetPlaybackPositionTicks returns the PlaybackPositionTicks field if non-nil, zero value otherwise.

### GetPlaybackPositionTicksOk

`func (o *BaseItemDtoUserData) GetPlaybackPositionTicksOk() (*int64, bool)`

GetPlaybackPositionTicksOk returns a tuple with the PlaybackPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackPositionTicks

`func (o *BaseItemDtoUserData) SetPlaybackPositionTicks(v int64)`

SetPlaybackPositionTicks sets PlaybackPositionTicks field to given value.

### HasPlaybackPositionTicks

`func (o *BaseItemDtoUserData) HasPlaybackPositionTicks() bool`

HasPlaybackPositionTicks returns a boolean if a field has been set.

### GetPlayCount

`func (o *BaseItemDtoUserData) GetPlayCount() int32`

GetPlayCount returns the PlayCount field if non-nil, zero value otherwise.

### GetPlayCountOk

`func (o *BaseItemDtoUserData) GetPlayCountOk() (*int32, bool)`

GetPlayCountOk returns a tuple with the PlayCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayCount

`func (o *BaseItemDtoUserData) SetPlayCount(v int32)`

SetPlayCount sets PlayCount field to given value.

### HasPlayCount

`func (o *BaseItemDtoUserData) HasPlayCount() bool`

HasPlayCount returns a boolean if a field has been set.

### GetIsFavorite

`func (o *BaseItemDtoUserData) GetIsFavorite() bool`

GetIsFavorite returns the IsFavorite field if non-nil, zero value otherwise.

### GetIsFavoriteOk

`func (o *BaseItemDtoUserData) GetIsFavoriteOk() (*bool, bool)`

GetIsFavoriteOk returns a tuple with the IsFavorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorite

`func (o *BaseItemDtoUserData) SetIsFavorite(v bool)`

SetIsFavorite sets IsFavorite field to given value.

### HasIsFavorite

`func (o *BaseItemDtoUserData) HasIsFavorite() bool`

HasIsFavorite returns a boolean if a field has been set.

### GetLikes

`func (o *BaseItemDtoUserData) GetLikes() bool`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *BaseItemDtoUserData) GetLikesOk() (*bool, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *BaseItemDtoUserData) SetLikes(v bool)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *BaseItemDtoUserData) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### SetLikesNil

`func (o *BaseItemDtoUserData) SetLikesNil(b bool)`

 SetLikesNil sets the value for Likes to be an explicit nil

### UnsetLikes
`func (o *BaseItemDtoUserData) UnsetLikes()`

UnsetLikes ensures that no value is present for Likes, not even an explicit nil
### GetLastPlayedDate

`func (o *BaseItemDtoUserData) GetLastPlayedDate() time.Time`

GetLastPlayedDate returns the LastPlayedDate field if non-nil, zero value otherwise.

### GetLastPlayedDateOk

`func (o *BaseItemDtoUserData) GetLastPlayedDateOk() (*time.Time, bool)`

GetLastPlayedDateOk returns a tuple with the LastPlayedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPlayedDate

`func (o *BaseItemDtoUserData) SetLastPlayedDate(v time.Time)`

SetLastPlayedDate sets LastPlayedDate field to given value.

### HasLastPlayedDate

`func (o *BaseItemDtoUserData) HasLastPlayedDate() bool`

HasLastPlayedDate returns a boolean if a field has been set.

### SetLastPlayedDateNil

`func (o *BaseItemDtoUserData) SetLastPlayedDateNil(b bool)`

 SetLastPlayedDateNil sets the value for LastPlayedDate to be an explicit nil

### UnsetLastPlayedDate
`func (o *BaseItemDtoUserData) UnsetLastPlayedDate()`

UnsetLastPlayedDate ensures that no value is present for LastPlayedDate, not even an explicit nil
### GetPlayed

`func (o *BaseItemDtoUserData) GetPlayed() bool`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *BaseItemDtoUserData) GetPlayedOk() (*bool, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *BaseItemDtoUserData) SetPlayed(v bool)`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *BaseItemDtoUserData) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### GetKey

`func (o *BaseItemDtoUserData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BaseItemDtoUserData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BaseItemDtoUserData) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *BaseItemDtoUserData) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *BaseItemDtoUserData) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *BaseItemDtoUserData) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetItemId

`func (o *BaseItemDtoUserData) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *BaseItemDtoUserData) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *BaseItemDtoUserData) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *BaseItemDtoUserData) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *BaseItemDtoUserData) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *BaseItemDtoUserData) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


