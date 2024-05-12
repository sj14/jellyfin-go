# GroupStateUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**GroupStateType**](GroupStateType.md) | Gets the state of the group. | [optional] 
**Reason** | Pointer to [**PlaybackRequestType**](PlaybackRequestType.md) | Gets the reason of the state change. | [optional] 

## Methods

### NewGroupStateUpdate

`func NewGroupStateUpdate() *GroupStateUpdate`

NewGroupStateUpdate instantiates a new GroupStateUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupStateUpdateWithDefaults

`func NewGroupStateUpdateWithDefaults() *GroupStateUpdate`

NewGroupStateUpdateWithDefaults instantiates a new GroupStateUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *GroupStateUpdate) GetState() GroupStateType`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GroupStateUpdate) GetStateOk() (*GroupStateType, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GroupStateUpdate) SetState(v GroupStateType)`

SetState sets State field to given value.

### HasState

`func (o *GroupStateUpdate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetReason

`func (o *GroupStateUpdate) GetReason() PlaybackRequestType`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GroupStateUpdate) GetReasonOk() (*PlaybackRequestType, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GroupStateUpdate) SetReason(v PlaybackRequestType)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GroupStateUpdate) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


