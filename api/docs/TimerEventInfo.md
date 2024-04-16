# TimerEventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ProgramId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTimerEventInfo

`func NewTimerEventInfo() *TimerEventInfo`

NewTimerEventInfo instantiates a new TimerEventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimerEventInfoWithDefaults

`func NewTimerEventInfoWithDefaults() *TimerEventInfo`

NewTimerEventInfoWithDefaults instantiates a new TimerEventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimerEventInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimerEventInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimerEventInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TimerEventInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProgramId

`func (o *TimerEventInfo) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *TimerEventInfo) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *TimerEventInfo) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *TimerEventInfo) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramIdNil

`func (o *TimerEventInfo) SetProgramIdNil(b bool)`

 SetProgramIdNil sets the value for ProgramId to be an explicit nil

### UnsetProgramId
`func (o *TimerEventInfo) UnsetProgramId()`

UnsetProgramId ensures that no value is present for ProgramId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


