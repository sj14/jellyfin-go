# AccessSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Gets the id of this instance. | [optional] [readonly] 
**UserId** | Pointer to **string** | Gets the id of the associated user. | [optional] 
**DayOfWeek** | Pointer to [**DynamicDayOfWeek**](DynamicDayOfWeek.md) | Gets or sets the day of week. | [optional] 
**StartHour** | Pointer to **float64** | Gets or sets the start hour. | [optional] 
**EndHour** | Pointer to **float64** | Gets or sets the end hour. | [optional] 

## Methods

### NewAccessSchedule

`func NewAccessSchedule() *AccessSchedule`

NewAccessSchedule instantiates a new AccessSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessScheduleWithDefaults

`func NewAccessScheduleWithDefaults() *AccessSchedule`

NewAccessScheduleWithDefaults instantiates a new AccessSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessSchedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessSchedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessSchedule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AccessSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *AccessSchedule) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AccessSchedule) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AccessSchedule) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AccessSchedule) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *AccessSchedule) GetDayOfWeek() DynamicDayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *AccessSchedule) GetDayOfWeekOk() (*DynamicDayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *AccessSchedule) SetDayOfWeek(v DynamicDayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *AccessSchedule) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetStartHour

`func (o *AccessSchedule) GetStartHour() float64`

GetStartHour returns the StartHour field if non-nil, zero value otherwise.

### GetStartHourOk

`func (o *AccessSchedule) GetStartHourOk() (*float64, bool)`

GetStartHourOk returns a tuple with the StartHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartHour

`func (o *AccessSchedule) SetStartHour(v float64)`

SetStartHour sets StartHour field to given value.

### HasStartHour

`func (o *AccessSchedule) HasStartHour() bool`

HasStartHour returns a boolean if a field has been set.

### GetEndHour

`func (o *AccessSchedule) GetEndHour() float64`

GetEndHour returns the EndHour field if non-nil, zero value otherwise.

### GetEndHourOk

`func (o *AccessSchedule) GetEndHourOk() (*float64, bool)`

GetEndHourOk returns a tuple with the EndHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHour

`func (o *AccessSchedule) SetEndHour(v float64)`

SetEndHour sets EndHour field to given value.

### HasEndHour

`func (o *AccessSchedule) HasEndHour() bool`

HasEndHour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


