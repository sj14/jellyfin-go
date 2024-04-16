# TaskInfoLastExecutionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTimeUtc** | Pointer to **time.Time** | Gets or sets the start time UTC. | [optional] 
**EndTimeUtc** | Pointer to **time.Time** | Gets or sets the end time UTC. | [optional] 
**Status** | Pointer to [**TaskCompletionStatus**](TaskCompletionStatus.md) | Gets or sets the status. | [optional] 
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**Key** | Pointer to **NullableString** | Gets or sets the key. | [optional] 
**Id** | Pointer to **NullableString** | Gets or sets the id. | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Gets or sets the error message. | [optional] 
**LongErrorMessage** | Pointer to **NullableString** | Gets or sets the long error message. | [optional] 

## Methods

### NewTaskInfoLastExecutionResult

`func NewTaskInfoLastExecutionResult() *TaskInfoLastExecutionResult`

NewTaskInfoLastExecutionResult instantiates a new TaskInfoLastExecutionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskInfoLastExecutionResultWithDefaults

`func NewTaskInfoLastExecutionResultWithDefaults() *TaskInfoLastExecutionResult`

NewTaskInfoLastExecutionResultWithDefaults instantiates a new TaskInfoLastExecutionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimeUtc

`func (o *TaskInfoLastExecutionResult) GetStartTimeUtc() time.Time`

GetStartTimeUtc returns the StartTimeUtc field if non-nil, zero value otherwise.

### GetStartTimeUtcOk

`func (o *TaskInfoLastExecutionResult) GetStartTimeUtcOk() (*time.Time, bool)`

GetStartTimeUtcOk returns a tuple with the StartTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUtc

`func (o *TaskInfoLastExecutionResult) SetStartTimeUtc(v time.Time)`

SetStartTimeUtc sets StartTimeUtc field to given value.

### HasStartTimeUtc

`func (o *TaskInfoLastExecutionResult) HasStartTimeUtc() bool`

HasStartTimeUtc returns a boolean if a field has been set.

### GetEndTimeUtc

`func (o *TaskInfoLastExecutionResult) GetEndTimeUtc() time.Time`

GetEndTimeUtc returns the EndTimeUtc field if non-nil, zero value otherwise.

### GetEndTimeUtcOk

`func (o *TaskInfoLastExecutionResult) GetEndTimeUtcOk() (*time.Time, bool)`

GetEndTimeUtcOk returns a tuple with the EndTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUtc

`func (o *TaskInfoLastExecutionResult) SetEndTimeUtc(v time.Time)`

SetEndTimeUtc sets EndTimeUtc field to given value.

### HasEndTimeUtc

`func (o *TaskInfoLastExecutionResult) HasEndTimeUtc() bool`

HasEndTimeUtc returns a boolean if a field has been set.

### GetStatus

`func (o *TaskInfoLastExecutionResult) GetStatus() TaskCompletionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskInfoLastExecutionResult) GetStatusOk() (*TaskCompletionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskInfoLastExecutionResult) SetStatus(v TaskCompletionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskInfoLastExecutionResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetName

`func (o *TaskInfoLastExecutionResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskInfoLastExecutionResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskInfoLastExecutionResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskInfoLastExecutionResult) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TaskInfoLastExecutionResult) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TaskInfoLastExecutionResult) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKey

`func (o *TaskInfoLastExecutionResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TaskInfoLastExecutionResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TaskInfoLastExecutionResult) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TaskInfoLastExecutionResult) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *TaskInfoLastExecutionResult) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *TaskInfoLastExecutionResult) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetId

`func (o *TaskInfoLastExecutionResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskInfoLastExecutionResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskInfoLastExecutionResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskInfoLastExecutionResult) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TaskInfoLastExecutionResult) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaskInfoLastExecutionResult) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetErrorMessage

`func (o *TaskInfoLastExecutionResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *TaskInfoLastExecutionResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *TaskInfoLastExecutionResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *TaskInfoLastExecutionResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *TaskInfoLastExecutionResult) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *TaskInfoLastExecutionResult) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetLongErrorMessage

`func (o *TaskInfoLastExecutionResult) GetLongErrorMessage() string`

GetLongErrorMessage returns the LongErrorMessage field if non-nil, zero value otherwise.

### GetLongErrorMessageOk

`func (o *TaskInfoLastExecutionResult) GetLongErrorMessageOk() (*string, bool)`

GetLongErrorMessageOk returns a tuple with the LongErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongErrorMessage

`func (o *TaskInfoLastExecutionResult) SetLongErrorMessage(v string)`

SetLongErrorMessage sets LongErrorMessage field to given value.

### HasLongErrorMessage

`func (o *TaskInfoLastExecutionResult) HasLongErrorMessage() bool`

HasLongErrorMessage returns a boolean if a field has been set.

### SetLongErrorMessageNil

`func (o *TaskInfoLastExecutionResult) SetLongErrorMessageNil(b bool)`

 SetLongErrorMessageNil sets the value for LongErrorMessage to be an explicit nil

### UnsetLongErrorMessage
`func (o *TaskInfoLastExecutionResult) UnsetLongErrorMessage()`

UnsetLongErrorMessage ensures that no value is present for LongErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


