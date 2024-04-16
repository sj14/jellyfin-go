# EndPointInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLocal** | Pointer to **bool** |  | [optional] 
**IsInNetwork** | Pointer to **bool** |  | [optional] 

## Methods

### NewEndPointInfo

`func NewEndPointInfo() *EndPointInfo`

NewEndPointInfo instantiates a new EndPointInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndPointInfoWithDefaults

`func NewEndPointInfoWithDefaults() *EndPointInfo`

NewEndPointInfoWithDefaults instantiates a new EndPointInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLocal

`func (o *EndPointInfo) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *EndPointInfo) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *EndPointInfo) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *EndPointInfo) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetIsInNetwork

`func (o *EndPointInfo) GetIsInNetwork() bool`

GetIsInNetwork returns the IsInNetwork field if non-nil, zero value otherwise.

### GetIsInNetworkOk

`func (o *EndPointInfo) GetIsInNetworkOk() (*bool, bool)`

GetIsInNetworkOk returns a tuple with the IsInNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInNetwork

`func (o *EndPointInfo) SetIsInNetwork(v bool)`

SetIsInNetwork sets IsInNetwork field to given value.

### HasIsInNetwork

`func (o *EndPointInfo) HasIsInNetwork() bool`

HasIsInNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


