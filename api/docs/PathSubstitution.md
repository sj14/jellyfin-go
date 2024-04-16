# PathSubstitution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Gets or sets the value to substitute. | [optional] 
**To** | Pointer to **string** | Gets or sets the value to substitution with. | [optional] 

## Methods

### NewPathSubstitution

`func NewPathSubstitution() *PathSubstitution`

NewPathSubstitution instantiates a new PathSubstitution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathSubstitutionWithDefaults

`func NewPathSubstitutionWithDefaults() *PathSubstitution`

NewPathSubstitutionWithDefaults instantiates a new PathSubstitution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *PathSubstitution) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PathSubstitution) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PathSubstitution) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PathSubstitution) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *PathSubstitution) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *PathSubstitution) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *PathSubstitution) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *PathSubstitution) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


