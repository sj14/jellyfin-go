# DeviceProfileInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Gets or sets the identifier. | [optional] 
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**Type** | Pointer to [**DeviceProfileType**](DeviceProfileType.md) | Gets or sets the type. | [optional] 

## Methods

### NewDeviceProfileInfo

`func NewDeviceProfileInfo() *DeviceProfileInfo`

NewDeviceProfileInfo instantiates a new DeviceProfileInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceProfileInfoWithDefaults

`func NewDeviceProfileInfoWithDefaults() *DeviceProfileInfo`

NewDeviceProfileInfoWithDefaults instantiates a new DeviceProfileInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceProfileInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceProfileInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceProfileInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceProfileInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DeviceProfileInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DeviceProfileInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *DeviceProfileInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceProfileInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceProfileInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceProfileInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *DeviceProfileInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *DeviceProfileInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *DeviceProfileInfo) GetType() DeviceProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceProfileInfo) GetTypeOk() (*DeviceProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceProfileInfo) SetType(v DeviceProfileType)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceProfileInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


