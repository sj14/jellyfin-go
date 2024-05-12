# ScheduledTaskEndedMessageData

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

### NewScheduledTaskEndedMessageData

`func NewScheduledTaskEndedMessageData() *ScheduledTaskEndedMessageData`

NewScheduledTaskEndedMessageData instantiates a new ScheduledTaskEndedMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledTaskEndedMessageDataWithDefaults

`func NewScheduledTaskEndedMessageDataWithDefaults() *ScheduledTaskEndedMessageData`

NewScheduledTaskEndedMessageDataWithDefaults instantiates a new ScheduledTaskEndedMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTimeUtc

`func (o *ScheduledTaskEndedMessageData) GetStartTimeUtc() time.Time`

GetStartTimeUtc returns the StartTimeUtc field if non-nil, zero value otherwise.

### GetStartTimeUtcOk

`func (o *ScheduledTaskEndedMessageData) GetStartTimeUtcOk() (*time.Time, bool)`

GetStartTimeUtcOk returns a tuple with the StartTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUtc

`func (o *ScheduledTaskEndedMessageData) SetStartTimeUtc(v time.Time)`

SetStartTimeUtc sets StartTimeUtc field to given value.

### HasStartTimeUtc

`func (o *ScheduledTaskEndedMessageData) HasStartTimeUtc() bool`

HasStartTimeUtc returns a boolean if a field has been set.

### GetEndTimeUtc

`func (o *ScheduledTaskEndedMessageData) GetEndTimeUtc() time.Time`

GetEndTimeUtc returns the EndTimeUtc field if non-nil, zero value otherwise.

### GetEndTimeUtcOk

`func (o *ScheduledTaskEndedMessageData) GetEndTimeUtcOk() (*time.Time, bool)`

GetEndTimeUtcOk returns a tuple with the EndTimeUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUtc

`func (o *ScheduledTaskEndedMessageData) SetEndTimeUtc(v time.Time)`

SetEndTimeUtc sets EndTimeUtc field to given value.

### HasEndTimeUtc

`func (o *ScheduledTaskEndedMessageData) HasEndTimeUtc() bool`

HasEndTimeUtc returns a boolean if a field has been set.

### GetStatus

`func (o *ScheduledTaskEndedMessageData) GetStatus() TaskCompletionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduledTaskEndedMessageData) GetStatusOk() (*TaskCompletionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduledTaskEndedMessageData) SetStatus(v TaskCompletionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScheduledTaskEndedMessageData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetName

`func (o *ScheduledTaskEndedMessageData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduledTaskEndedMessageData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduledTaskEndedMessageData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScheduledTaskEndedMessageData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ScheduledTaskEndedMessageData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ScheduledTaskEndedMessageData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetKey

`func (o *ScheduledTaskEndedMessageData) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ScheduledTaskEndedMessageData) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ScheduledTaskEndedMessageData) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ScheduledTaskEndedMessageData) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *ScheduledTaskEndedMessageData) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *ScheduledTaskEndedMessageData) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetId

`func (o *ScheduledTaskEndedMessageData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduledTaskEndedMessageData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduledTaskEndedMessageData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduledTaskEndedMessageData) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ScheduledTaskEndedMessageData) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ScheduledTaskEndedMessageData) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetErrorMessage

`func (o *ScheduledTaskEndedMessageData) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ScheduledTaskEndedMessageData) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ScheduledTaskEndedMessageData) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ScheduledTaskEndedMessageData) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ScheduledTaskEndedMessageData) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ScheduledTaskEndedMessageData) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetLongErrorMessage

`func (o *ScheduledTaskEndedMessageData) GetLongErrorMessage() string`

GetLongErrorMessage returns the LongErrorMessage field if non-nil, zero value otherwise.

### GetLongErrorMessageOk

`func (o *ScheduledTaskEndedMessageData) GetLongErrorMessageOk() (*string, bool)`

GetLongErrorMessageOk returns a tuple with the LongErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongErrorMessage

`func (o *ScheduledTaskEndedMessageData) SetLongErrorMessage(v string)`

SetLongErrorMessage sets LongErrorMessage field to given value.

### HasLongErrorMessage

`func (o *ScheduledTaskEndedMessageData) HasLongErrorMessage() bool`

HasLongErrorMessage returns a boolean if a field has been set.

### SetLongErrorMessageNil

`func (o *ScheduledTaskEndedMessageData) SetLongErrorMessageNil(b bool)`

 SetLongErrorMessageNil sets the value for LongErrorMessage to be an explicit nil

### UnsetLongErrorMessage
`func (o *ScheduledTaskEndedMessageData) UnsetLongErrorMessage()`

UnsetLongErrorMessage ensures that no value is present for LongErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


