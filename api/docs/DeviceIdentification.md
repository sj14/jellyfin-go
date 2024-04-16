# DeviceIdentification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FriendlyName** | Pointer to **string** | Gets or sets the name of the friendly. | [optional] 
**ModelNumber** | Pointer to **string** | Gets or sets the model number. | [optional] 
**SerialNumber** | Pointer to **string** | Gets or sets the serial number. | [optional] 
**ModelName** | Pointer to **string** | Gets or sets the name of the model. | [optional] 
**ModelDescription** | Pointer to **string** | Gets or sets the model description. | [optional] 
**ModelUrl** | Pointer to **string** | Gets or sets the model URL. | [optional] 
**Manufacturer** | Pointer to **string** | Gets or sets the manufacturer. | [optional] 
**ManufacturerUrl** | Pointer to **string** | Gets or sets the manufacturer URL. | [optional] 
**Headers** | Pointer to [**[]HttpHeaderInfo**](HttpHeaderInfo.md) | Gets or sets the headers. | [optional] 

## Methods

### NewDeviceIdentification

`func NewDeviceIdentification() *DeviceIdentification`

NewDeviceIdentification instantiates a new DeviceIdentification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceIdentificationWithDefaults

`func NewDeviceIdentificationWithDefaults() *DeviceIdentification`

NewDeviceIdentificationWithDefaults instantiates a new DeviceIdentification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFriendlyName

`func (o *DeviceIdentification) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *DeviceIdentification) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *DeviceIdentification) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *DeviceIdentification) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetModelNumber

`func (o *DeviceIdentification) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *DeviceIdentification) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *DeviceIdentification) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *DeviceIdentification) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetSerialNumber

`func (o *DeviceIdentification) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *DeviceIdentification) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *DeviceIdentification) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *DeviceIdentification) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetModelName

`func (o *DeviceIdentification) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *DeviceIdentification) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *DeviceIdentification) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *DeviceIdentification) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetModelDescription

`func (o *DeviceIdentification) GetModelDescription() string`

GetModelDescription returns the ModelDescription field if non-nil, zero value otherwise.

### GetModelDescriptionOk

`func (o *DeviceIdentification) GetModelDescriptionOk() (*string, bool)`

GetModelDescriptionOk returns a tuple with the ModelDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelDescription

`func (o *DeviceIdentification) SetModelDescription(v string)`

SetModelDescription sets ModelDescription field to given value.

### HasModelDescription

`func (o *DeviceIdentification) HasModelDescription() bool`

HasModelDescription returns a boolean if a field has been set.

### GetModelUrl

`func (o *DeviceIdentification) GetModelUrl() string`

GetModelUrl returns the ModelUrl field if non-nil, zero value otherwise.

### GetModelUrlOk

`func (o *DeviceIdentification) GetModelUrlOk() (*string, bool)`

GetModelUrlOk returns a tuple with the ModelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelUrl

`func (o *DeviceIdentification) SetModelUrl(v string)`

SetModelUrl sets ModelUrl field to given value.

### HasModelUrl

`func (o *DeviceIdentification) HasModelUrl() bool`

HasModelUrl returns a boolean if a field has been set.

### GetManufacturer

`func (o *DeviceIdentification) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *DeviceIdentification) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *DeviceIdentification) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *DeviceIdentification) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetManufacturerUrl

`func (o *DeviceIdentification) GetManufacturerUrl() string`

GetManufacturerUrl returns the ManufacturerUrl field if non-nil, zero value otherwise.

### GetManufacturerUrlOk

`func (o *DeviceIdentification) GetManufacturerUrlOk() (*string, bool)`

GetManufacturerUrlOk returns a tuple with the ManufacturerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerUrl

`func (o *DeviceIdentification) SetManufacturerUrl(v string)`

SetManufacturerUrl sets ManufacturerUrl field to given value.

### HasManufacturerUrl

`func (o *DeviceIdentification) HasManufacturerUrl() bool`

HasManufacturerUrl returns a boolean if a field has been set.

### GetHeaders

`func (o *DeviceIdentification) GetHeaders() []HttpHeaderInfo`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *DeviceIdentification) GetHeadersOk() (*[]HttpHeaderInfo, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *DeviceIdentification) SetHeaders(v []HttpHeaderInfo)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *DeviceIdentification) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


