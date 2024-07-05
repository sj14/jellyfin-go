# TaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**State** | Pointer to [**TaskState**](TaskState.md) | Gets or sets the state of the task. | [optional] 
**CurrentProgressPercentage** | Pointer to **NullableFloat64** | Gets or sets the progress. | [optional] 
**Id** | Pointer to **NullableString** | Gets or sets the id. | [optional] 
**LastExecutionResult** | Pointer to [**NullableTaskResult**](TaskResult.md) | Gets or sets the last execution result. | [optional] 
**Triggers** | Pointer to [**[]TaskTriggerInfo**](TaskTriggerInfo.md) | Gets or sets the triggers. | [optional] 
**Description** | Pointer to **NullableString** | Gets or sets the description. | [optional] 
**Category** | Pointer to **NullableString** | Gets or sets the category. | [optional] 
**IsHidden** | Pointer to **bool** | Gets or sets a value indicating whether this instance is hidden. | [optional] 
**Key** | Pointer to **NullableString** | Gets or sets the key. | [optional] 

## Methods

### NewTaskInfo

`func NewTaskInfo() *TaskInfo`

NewTaskInfo instantiates a new TaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskInfoWithDefaults

`func NewTaskInfoWithDefaults() *TaskInfo`

NewTaskInfoWithDefaults instantiates a new TaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TaskInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaskInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaskInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaskInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TaskInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TaskInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetState

`func (o *TaskInfo) GetState() TaskState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TaskInfo) GetStateOk() (*TaskState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TaskInfo) SetState(v TaskState)`

SetState sets State field to given value.

### HasState

`func (o *TaskInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCurrentProgressPercentage

`func (o *TaskInfo) GetCurrentProgressPercentage() float64`

GetCurrentProgressPercentage returns the CurrentProgressPercentage field if non-nil, zero value otherwise.

### GetCurrentProgressPercentageOk

`func (o *TaskInfo) GetCurrentProgressPercentageOk() (*float64, bool)`

GetCurrentProgressPercentageOk returns a tuple with the CurrentProgressPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProgressPercentage

`func (o *TaskInfo) SetCurrentProgressPercentage(v float64)`

SetCurrentProgressPercentage sets CurrentProgressPercentage field to given value.

### HasCurrentProgressPercentage

`func (o *TaskInfo) HasCurrentProgressPercentage() bool`

HasCurrentProgressPercentage returns a boolean if a field has been set.

### SetCurrentProgressPercentageNil

`func (o *TaskInfo) SetCurrentProgressPercentageNil(b bool)`

 SetCurrentProgressPercentageNil sets the value for CurrentProgressPercentage to be an explicit nil

### UnsetCurrentProgressPercentage
`func (o *TaskInfo) UnsetCurrentProgressPercentage()`

UnsetCurrentProgressPercentage ensures that no value is present for CurrentProgressPercentage, not even an explicit nil
### GetId

`func (o *TaskInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TaskInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TaskInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLastExecutionResult

`func (o *TaskInfo) GetLastExecutionResult() TaskResult`

GetLastExecutionResult returns the LastExecutionResult field if non-nil, zero value otherwise.

### GetLastExecutionResultOk

`func (o *TaskInfo) GetLastExecutionResultOk() (*TaskResult, bool)`

GetLastExecutionResultOk returns a tuple with the LastExecutionResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionResult

`func (o *TaskInfo) SetLastExecutionResult(v TaskResult)`

SetLastExecutionResult sets LastExecutionResult field to given value.

### HasLastExecutionResult

`func (o *TaskInfo) HasLastExecutionResult() bool`

HasLastExecutionResult returns a boolean if a field has been set.

### SetLastExecutionResultNil

`func (o *TaskInfo) SetLastExecutionResultNil(b bool)`

 SetLastExecutionResultNil sets the value for LastExecutionResult to be an explicit nil

### UnsetLastExecutionResult
`func (o *TaskInfo) UnsetLastExecutionResult()`

UnsetLastExecutionResult ensures that no value is present for LastExecutionResult, not even an explicit nil
### GetTriggers

`func (o *TaskInfo) GetTriggers() []TaskTriggerInfo`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *TaskInfo) GetTriggersOk() (*[]TaskTriggerInfo, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *TaskInfo) SetTriggers(v []TaskTriggerInfo)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *TaskInfo) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggersNil

`func (o *TaskInfo) SetTriggersNil(b bool)`

 SetTriggersNil sets the value for Triggers to be an explicit nil

### UnsetTriggers
`func (o *TaskInfo) UnsetTriggers()`

UnsetTriggers ensures that no value is present for Triggers, not even an explicit nil
### GetDescription

`func (o *TaskInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaskInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaskInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaskInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TaskInfo) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TaskInfo) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCategory

`func (o *TaskInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TaskInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TaskInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TaskInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *TaskInfo) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *TaskInfo) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetIsHidden

`func (o *TaskInfo) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *TaskInfo) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *TaskInfo) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *TaskInfo) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetKey

`func (o *TaskInfo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TaskInfo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TaskInfo) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TaskInfo) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *TaskInfo) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *TaskInfo) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


