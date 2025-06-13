# ParentalRating

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Gets or sets the name. | [optional] 
**Value** | Pointer to **NullableInt32** | Gets or sets the value. | [optional] 
**RatingScore** | Pointer to [**NullableParentalRatingScore**](ParentalRatingScore.md) | Gets or sets the rating score. | [optional] 

## Methods

### NewParentalRating

`func NewParentalRating() *ParentalRating`

NewParentalRating instantiates a new ParentalRating object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParentalRatingWithDefaults

`func NewParentalRatingWithDefaults() *ParentalRating`

NewParentalRatingWithDefaults instantiates a new ParentalRating object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ParentalRating) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParentalRating) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParentalRating) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParentalRating) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ParentalRating) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ParentalRating) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ParentalRating) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ParentalRating) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ParentalRating) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ParentalRating) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetRatingScore

`func (o *ParentalRating) GetRatingScore() ParentalRatingScore`

GetRatingScore returns the RatingScore field if non-nil, zero value otherwise.

### GetRatingScoreOk

`func (o *ParentalRating) GetRatingScoreOk() (*ParentalRatingScore, bool)`

GetRatingScoreOk returns a tuple with the RatingScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingScore

`func (o *ParentalRating) SetRatingScore(v ParentalRatingScore)`

SetRatingScore sets RatingScore field to given value.

### HasRatingScore

`func (o *ParentalRating) HasRatingScore() bool`

HasRatingScore returns a boolean if a field has been set.

### SetRatingScoreNil

`func (o *ParentalRating) SetRatingScoreNil(b bool)`

 SetRatingScoreNil sets the value for RatingScore to be an explicit nil

### UnsetRatingScore
`func (o *ParentalRating) UnsetRatingScore()`

UnsetRatingScore ensures that no value is present for RatingScore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


