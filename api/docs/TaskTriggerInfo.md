# TaskTriggerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TaskTriggerInfoType**](TaskTriggerInfoType.md) | Gets or sets the type. | [optional] 
**TimeOfDayTicks** | Pointer to **NullableInt64** | Gets or sets the time of day. | [optional] 
**IntervalTicks** | Pointer to **NullableInt64** | Gets or sets the interval. | [optional] 
**DayOfWeek** | Pointer to [**NullableDayOfWeek**](DayOfWeek.md) | Gets or sets the day of week. | [optional] 
**MaxRuntimeTicks** | Pointer to **NullableInt64** | Gets or sets the maximum runtime ticks. | [optional] 

## Methods

### NewTaskTriggerInfo

`func NewTaskTriggerInfo() *TaskTriggerInfo`

NewTaskTriggerInfo instantiates a new TaskTriggerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskTriggerInfoWithDefaults

`func NewTaskTriggerInfoWithDefaults() *TaskTriggerInfo`

NewTaskTriggerInfoWithDefaults instantiates a new TaskTriggerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TaskTriggerInfo) GetType() TaskTriggerInfoType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskTriggerInfo) GetTypeOk() (*TaskTriggerInfoType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskTriggerInfo) SetType(v TaskTriggerInfoType)`

SetType sets Type field to given value.

### HasType

`func (o *TaskTriggerInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTimeOfDayTicks

`func (o *TaskTriggerInfo) GetTimeOfDayTicks() int64`

GetTimeOfDayTicks returns the TimeOfDayTicks field if non-nil, zero value otherwise.

### GetTimeOfDayTicksOk

`func (o *TaskTriggerInfo) GetTimeOfDayTicksOk() (*int64, bool)`

GetTimeOfDayTicksOk returns a tuple with the TimeOfDayTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfDayTicks

`func (o *TaskTriggerInfo) SetTimeOfDayTicks(v int64)`

SetTimeOfDayTicks sets TimeOfDayTicks field to given value.

### HasTimeOfDayTicks

`func (o *TaskTriggerInfo) HasTimeOfDayTicks() bool`

HasTimeOfDayTicks returns a boolean if a field has been set.

### SetTimeOfDayTicksNil

`func (o *TaskTriggerInfo) SetTimeOfDayTicksNil(b bool)`

 SetTimeOfDayTicksNil sets the value for TimeOfDayTicks to be an explicit nil

### UnsetTimeOfDayTicks
`func (o *TaskTriggerInfo) UnsetTimeOfDayTicks()`

UnsetTimeOfDayTicks ensures that no value is present for TimeOfDayTicks, not even an explicit nil
### GetIntervalTicks

`func (o *TaskTriggerInfo) GetIntervalTicks() int64`

GetIntervalTicks returns the IntervalTicks field if non-nil, zero value otherwise.

### GetIntervalTicksOk

`func (o *TaskTriggerInfo) GetIntervalTicksOk() (*int64, bool)`

GetIntervalTicksOk returns a tuple with the IntervalTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalTicks

`func (o *TaskTriggerInfo) SetIntervalTicks(v int64)`

SetIntervalTicks sets IntervalTicks field to given value.

### HasIntervalTicks

`func (o *TaskTriggerInfo) HasIntervalTicks() bool`

HasIntervalTicks returns a boolean if a field has been set.

### SetIntervalTicksNil

`func (o *TaskTriggerInfo) SetIntervalTicksNil(b bool)`

 SetIntervalTicksNil sets the value for IntervalTicks to be an explicit nil

### UnsetIntervalTicks
`func (o *TaskTriggerInfo) UnsetIntervalTicks()`

UnsetIntervalTicks ensures that no value is present for IntervalTicks, not even an explicit nil
### GetDayOfWeek

`func (o *TaskTriggerInfo) GetDayOfWeek() DayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *TaskTriggerInfo) GetDayOfWeekOk() (*DayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *TaskTriggerInfo) SetDayOfWeek(v DayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *TaskTriggerInfo) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### SetDayOfWeekNil

`func (o *TaskTriggerInfo) SetDayOfWeekNil(b bool)`

 SetDayOfWeekNil sets the value for DayOfWeek to be an explicit nil

### UnsetDayOfWeek
`func (o *TaskTriggerInfo) UnsetDayOfWeek()`

UnsetDayOfWeek ensures that no value is present for DayOfWeek, not even an explicit nil
### GetMaxRuntimeTicks

`func (o *TaskTriggerInfo) GetMaxRuntimeTicks() int64`

GetMaxRuntimeTicks returns the MaxRuntimeTicks field if non-nil, zero value otherwise.

### GetMaxRuntimeTicksOk

`func (o *TaskTriggerInfo) GetMaxRuntimeTicksOk() (*int64, bool)`

GetMaxRuntimeTicksOk returns a tuple with the MaxRuntimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRuntimeTicks

`func (o *TaskTriggerInfo) SetMaxRuntimeTicks(v int64)`

SetMaxRuntimeTicks sets MaxRuntimeTicks field to given value.

### HasMaxRuntimeTicks

`func (o *TaskTriggerInfo) HasMaxRuntimeTicks() bool`

HasMaxRuntimeTicks returns a boolean if a field has been set.

### SetMaxRuntimeTicksNil

`func (o *TaskTriggerInfo) SetMaxRuntimeTicksNil(b bool)`

 SetMaxRuntimeTicksNil sets the value for MaxRuntimeTicks to be an explicit nil

### UnsetMaxRuntimeTicks
`func (o *TaskTriggerInfo) UnsetMaxRuntimeTicks()`

UnsetMaxRuntimeTicks ensures that no value is present for MaxRuntimeTicks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


