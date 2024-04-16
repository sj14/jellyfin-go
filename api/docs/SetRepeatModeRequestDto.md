# SetRepeatModeRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to [**GroupRepeatMode**](GroupRepeatMode.md) | Enum GroupRepeatMode. | [optional] 

## Methods

### NewSetRepeatModeRequestDto

`func NewSetRepeatModeRequestDto() *SetRepeatModeRequestDto`

NewSetRepeatModeRequestDto instantiates a new SetRepeatModeRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetRepeatModeRequestDtoWithDefaults

`func NewSetRepeatModeRequestDtoWithDefaults() *SetRepeatModeRequestDto`

NewSetRepeatModeRequestDtoWithDefaults instantiates a new SetRepeatModeRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *SetRepeatModeRequestDto) GetMode() GroupRepeatMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SetRepeatModeRequestDto) GetModeOk() (*GroupRepeatMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SetRepeatModeRequestDto) SetMode(v GroupRepeatMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SetRepeatModeRequestDto) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


