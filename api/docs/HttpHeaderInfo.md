# HttpHeaderInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**Match** | Pointer to [**HeaderMatchType**](HeaderMatchType.md) |  | [optional] 

## Methods

### NewHttpHeaderInfo

`func NewHttpHeaderInfo() *HttpHeaderInfo`

NewHttpHeaderInfo instantiates a new HttpHeaderInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpHeaderInfoWithDefaults

`func NewHttpHeaderInfoWithDefaults() *HttpHeaderInfo`

NewHttpHeaderInfoWithDefaults instantiates a new HttpHeaderInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HttpHeaderInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HttpHeaderInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HttpHeaderInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HttpHeaderInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *HttpHeaderInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *HttpHeaderInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *HttpHeaderInfo) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HttpHeaderInfo) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HttpHeaderInfo) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *HttpHeaderInfo) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *HttpHeaderInfo) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *HttpHeaderInfo) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMatch

`func (o *HttpHeaderInfo) GetMatch() HeaderMatchType`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *HttpHeaderInfo) GetMatchOk() (*HeaderMatchType, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *HttpHeaderInfo) SetMatch(v HeaderMatchType)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *HttpHeaderInfo) HasMatch() bool`

HasMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


