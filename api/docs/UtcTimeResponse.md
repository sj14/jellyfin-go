# UtcTimeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestReceptionTime** | Pointer to **time.Time** | Gets the UTC time when request has been received. | [optional] 
**ResponseTransmissionTime** | Pointer to **time.Time** | Gets the UTC time when response has been sent. | [optional] 

## Methods

### NewUtcTimeResponse

`func NewUtcTimeResponse() *UtcTimeResponse`

NewUtcTimeResponse instantiates a new UtcTimeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtcTimeResponseWithDefaults

`func NewUtcTimeResponseWithDefaults() *UtcTimeResponse`

NewUtcTimeResponseWithDefaults instantiates a new UtcTimeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestReceptionTime

`func (o *UtcTimeResponse) GetRequestReceptionTime() time.Time`

GetRequestReceptionTime returns the RequestReceptionTime field if non-nil, zero value otherwise.

### GetRequestReceptionTimeOk

`func (o *UtcTimeResponse) GetRequestReceptionTimeOk() (*time.Time, bool)`

GetRequestReceptionTimeOk returns a tuple with the RequestReceptionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestReceptionTime

`func (o *UtcTimeResponse) SetRequestReceptionTime(v time.Time)`

SetRequestReceptionTime sets RequestReceptionTime field to given value.

### HasRequestReceptionTime

`func (o *UtcTimeResponse) HasRequestReceptionTime() bool`

HasRequestReceptionTime returns a boolean if a field has been set.

### GetResponseTransmissionTime

`func (o *UtcTimeResponse) GetResponseTransmissionTime() time.Time`

GetResponseTransmissionTime returns the ResponseTransmissionTime field if non-nil, zero value otherwise.

### GetResponseTransmissionTimeOk

`func (o *UtcTimeResponse) GetResponseTransmissionTimeOk() (*time.Time, bool)`

GetResponseTransmissionTimeOk returns a tuple with the ResponseTransmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTransmissionTime

`func (o *UtcTimeResponse) SetResponseTransmissionTime(v time.Time)`

SetResponseTransmissionTime sets ResponseTransmissionTime field to given value.

### HasResponseTransmissionTime

`func (o *UtcTimeResponse) HasResponseTransmissionTime() bool`

HasResponseTransmissionTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


