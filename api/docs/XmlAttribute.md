# XmlAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name of the attribute. | [optional] 
**Value** | Pointer to **NullableString** | Gets or sets the value of the attribute. | [optional] 

## Methods

### NewXmlAttribute

`func NewXmlAttribute() *XmlAttribute`

NewXmlAttribute instantiates a new XmlAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXmlAttributeWithDefaults

`func NewXmlAttributeWithDefaults() *XmlAttribute`

NewXmlAttributeWithDefaults instantiates a new XmlAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *XmlAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *XmlAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *XmlAttribute) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *XmlAttribute) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *XmlAttribute) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *XmlAttribute) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetValue

`func (o *XmlAttribute) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *XmlAttribute) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *XmlAttribute) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *XmlAttribute) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *XmlAttribute) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *XmlAttribute) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


