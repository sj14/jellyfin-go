# DeviceOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Gets the id. | [optional] [readonly] 
**DeviceId** | Pointer to **string** | Gets the device id. | [optional] 
**CustomName** | Pointer to **NullableString** | Gets or sets the custom name. | [optional] 

## Methods

### NewDeviceOptions

`func NewDeviceOptions() *DeviceOptions`

NewDeviceOptions instantiates a new DeviceOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceOptionsWithDefaults

`func NewDeviceOptionsWithDefaults() *DeviceOptions`

NewDeviceOptionsWithDefaults instantiates a new DeviceOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceOptions) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceOptions) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceOptions) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceOptions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeviceId

`func (o *DeviceOptions) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceOptions) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceOptions) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceOptions) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetCustomName

`func (o *DeviceOptions) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *DeviceOptions) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *DeviceOptions) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *DeviceOptions) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *DeviceOptions) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *DeviceOptions) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


