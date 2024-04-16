# DeviceOptionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Gets or sets the id. | [optional] 
**DeviceId** | Pointer to **NullableString** | Gets or sets the device id. | [optional] 
**CustomName** | Pointer to **NullableString** | Gets or sets the custom name. | [optional] 

## Methods

### NewDeviceOptionsDto

`func NewDeviceOptionsDto() *DeviceOptionsDto`

NewDeviceOptionsDto instantiates a new DeviceOptionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceOptionsDtoWithDefaults

`func NewDeviceOptionsDtoWithDefaults() *DeviceOptionsDto`

NewDeviceOptionsDtoWithDefaults instantiates a new DeviceOptionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceOptionsDto) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceOptionsDto) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceOptionsDto) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceOptionsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeviceId

`func (o *DeviceOptionsDto) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceOptionsDto) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceOptionsDto) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceOptionsDto) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *DeviceOptionsDto) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *DeviceOptionsDto) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetCustomName

`func (o *DeviceOptionsDto) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *DeviceOptionsDto) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *DeviceOptionsDto) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.

### HasCustomName

`func (o *DeviceOptionsDto) HasCustomName() bool`

HasCustomName returns a boolean if a field has been set.

### SetCustomNameNil

`func (o *DeviceOptionsDto) SetCustomNameNil(b bool)`

 SetCustomNameNil sets the value for CustomName to be an explicit nil

### UnsetCustomName
`func (o *DeviceOptionsDto) UnsetCustomName()`

UnsetCustomName ensures that no value is present for CustomName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


