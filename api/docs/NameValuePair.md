# NameValuePair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**Value** | Pointer to **NullableString** | Gets or sets the value. | [optional] 

## Methods

### NewNameValuePair

`func NewNameValuePair() *NameValuePair`

NewNameValuePair instantiates a new NameValuePair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameValuePairWithDefaults

`func NewNameValuePairWithDefaults() *NameValuePair`

NewNameValuePairWithDefaults instantiates a new NameValuePair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NameValuePair) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NameValuePair) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NameValuePair) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NameValuePair) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *NameValuePair) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NameValuePair) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *NameValuePair) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NameValuePair) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NameValuePair) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *NameValuePair) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *NameValuePair) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *NameValuePair) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


