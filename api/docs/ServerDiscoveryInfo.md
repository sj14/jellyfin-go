# ServerDiscoveryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Gets the address. | [optional] 
**Id** | Pointer to **string** | Gets the server identifier. | [optional] 
**Name** | Pointer to **string** | Gets the name. | [optional] 
**EndpointAddress** | Pointer to **NullableString** | Gets the endpoint address. | [optional] 

## Methods

### NewServerDiscoveryInfo

`func NewServerDiscoveryInfo() *ServerDiscoveryInfo`

NewServerDiscoveryInfo instantiates a new ServerDiscoveryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerDiscoveryInfoWithDefaults

`func NewServerDiscoveryInfoWithDefaults() *ServerDiscoveryInfo`

NewServerDiscoveryInfoWithDefaults instantiates a new ServerDiscoveryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ServerDiscoveryInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ServerDiscoveryInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ServerDiscoveryInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ServerDiscoveryInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetId

`func (o *ServerDiscoveryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerDiscoveryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerDiscoveryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerDiscoveryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServerDiscoveryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerDiscoveryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerDiscoveryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerDiscoveryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEndpointAddress

`func (o *ServerDiscoveryInfo) GetEndpointAddress() string`

GetEndpointAddress returns the EndpointAddress field if non-nil, zero value otherwise.

### GetEndpointAddressOk

`func (o *ServerDiscoveryInfo) GetEndpointAddressOk() (*string, bool)`

GetEndpointAddressOk returns a tuple with the EndpointAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointAddress

`func (o *ServerDiscoveryInfo) SetEndpointAddress(v string)`

SetEndpointAddress sets EndpointAddress field to given value.

### HasEndpointAddress

`func (o *ServerDiscoveryInfo) HasEndpointAddress() bool`

HasEndpointAddress returns a boolean if a field has been set.

### SetEndpointAddressNil

`func (o *ServerDiscoveryInfo) SetEndpointAddressNil(b bool)`

 SetEndpointAddressNil sets the value for EndpointAddress to be an explicit nil

### UnsetEndpointAddress
`func (o *ServerDiscoveryInfo) UnsetEndpointAddress()`

UnsetEndpointAddress ensures that no value is present for EndpointAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


