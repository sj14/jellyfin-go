# PingRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ping** | Pointer to **int64** | Gets or sets the ping time. | [optional] 

## Methods

### NewPingRequestDto

`func NewPingRequestDto() *PingRequestDto`

NewPingRequestDto instantiates a new PingRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingRequestDtoWithDefaults

`func NewPingRequestDtoWithDefaults() *PingRequestDto`

NewPingRequestDtoWithDefaults instantiates a new PingRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPing

`func (o *PingRequestDto) GetPing() int64`

GetPing returns the Ping field if non-nil, zero value otherwise.

### GetPingOk

`func (o *PingRequestDto) GetPingOk() (*int64, bool)`

GetPingOk returns a tuple with the Ping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPing

`func (o *PingRequestDto) SetPing(v int64)`

SetPing sets Ping field to given value.

### HasPing

`func (o *PingRequestDto) HasPing() bool`

HasPing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


