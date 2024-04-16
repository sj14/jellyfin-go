# ControlResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to **map[string]string** |  | [optional] [readonly] 
**Xml** | Pointer to **string** |  | [optional] 
**IsSuccessful** | Pointer to **bool** |  | [optional] 

## Methods

### NewControlResponse

`func NewControlResponse() *ControlResponse`

NewControlResponse instantiates a new ControlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControlResponseWithDefaults

`func NewControlResponseWithDefaults() *ControlResponse`

NewControlResponseWithDefaults instantiates a new ControlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *ControlResponse) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ControlResponse) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ControlResponse) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ControlResponse) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetXml

`func (o *ControlResponse) GetXml() string`

GetXml returns the Xml field if non-nil, zero value otherwise.

### GetXmlOk

`func (o *ControlResponse) GetXmlOk() (*string, bool)`

GetXmlOk returns a tuple with the Xml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXml

`func (o *ControlResponse) SetXml(v string)`

SetXml sets Xml field to given value.

### HasXml

`func (o *ControlResponse) HasXml() bool`

HasXml returns a boolean if a field has been set.

### GetIsSuccessful

`func (o *ControlResponse) GetIsSuccessful() bool`

GetIsSuccessful returns the IsSuccessful field if non-nil, zero value otherwise.

### GetIsSuccessfulOk

`func (o *ControlResponse) GetIsSuccessfulOk() (*bool, bool)`

GetIsSuccessfulOk returns a tuple with the IsSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccessful

`func (o *ControlResponse) SetIsSuccessful(v bool)`

SetIsSuccessful sets IsSuccessful field to given value.

### HasIsSuccessful

`func (o *ControlResponse) HasIsSuccessful() bool`

HasIsSuccessful returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


