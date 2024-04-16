# ProfileCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to [**ProfileConditionType**](ProfileConditionType.md) |  | [optional] 
**Property** | Pointer to [**ProfileConditionValue**](ProfileConditionValue.md) |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsRequired** | Pointer to **bool** |  | [optional] 

## Methods

### NewProfileCondition

`func NewProfileCondition() *ProfileCondition`

NewProfileCondition instantiates a new ProfileCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileConditionWithDefaults

`func NewProfileConditionWithDefaults() *ProfileCondition`

NewProfileConditionWithDefaults instantiates a new ProfileCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *ProfileCondition) GetCondition() ProfileConditionType`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ProfileCondition) GetConditionOk() (*ProfileConditionType, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ProfileCondition) SetCondition(v ProfileConditionType)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ProfileCondition) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetProperty

`func (o *ProfileCondition) GetProperty() ProfileConditionValue`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ProfileCondition) GetPropertyOk() (*ProfileConditionValue, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ProfileCondition) SetProperty(v ProfileConditionValue)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ProfileCondition) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetValue

`func (o *ProfileCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ProfileCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ProfileCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ProfileCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ProfileCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ProfileCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsRequired

`func (o *ProfileCondition) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ProfileCondition) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ProfileCondition) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *ProfileCondition) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


