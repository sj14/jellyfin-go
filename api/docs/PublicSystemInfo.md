# PublicSystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalAddress** | Pointer to **NullableString** | Gets or sets the local address. | [optional] 
**ServerName** | Pointer to **NullableString** | Gets or sets the name of the server. | [optional] 
**Version** | Pointer to **NullableString** | Gets or sets the server version. | [optional] 
**ProductName** | Pointer to **NullableString** | Gets or sets the product name. This is the AssemblyProduct name. | [optional] 
**OperatingSystem** | Pointer to **NullableString** | Gets or sets the operating system. | [optional] 
**Id** | Pointer to **NullableString** | Gets or sets the id. | [optional] 
**StartupWizardCompleted** | Pointer to **NullableBool** | Gets or sets a value indicating whether the startup wizard is completed. | [optional] 

## Methods

### NewPublicSystemInfo

`func NewPublicSystemInfo() *PublicSystemInfo`

NewPublicSystemInfo instantiates a new PublicSystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSystemInfoWithDefaults

`func NewPublicSystemInfoWithDefaults() *PublicSystemInfo`

NewPublicSystemInfoWithDefaults instantiates a new PublicSystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalAddress

`func (o *PublicSystemInfo) GetLocalAddress() string`

GetLocalAddress returns the LocalAddress field if non-nil, zero value otherwise.

### GetLocalAddressOk

`func (o *PublicSystemInfo) GetLocalAddressOk() (*string, bool)`

GetLocalAddressOk returns a tuple with the LocalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAddress

`func (o *PublicSystemInfo) SetLocalAddress(v string)`

SetLocalAddress sets LocalAddress field to given value.

### HasLocalAddress

`func (o *PublicSystemInfo) HasLocalAddress() bool`

HasLocalAddress returns a boolean if a field has been set.

### SetLocalAddressNil

`func (o *PublicSystemInfo) SetLocalAddressNil(b bool)`

 SetLocalAddressNil sets the value for LocalAddress to be an explicit nil

### UnsetLocalAddress
`func (o *PublicSystemInfo) UnsetLocalAddress()`

UnsetLocalAddress ensures that no value is present for LocalAddress, not even an explicit nil
### GetServerName

`func (o *PublicSystemInfo) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *PublicSystemInfo) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *PublicSystemInfo) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *PublicSystemInfo) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *PublicSystemInfo) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *PublicSystemInfo) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetVersion

`func (o *PublicSystemInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PublicSystemInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PublicSystemInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PublicSystemInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *PublicSystemInfo) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *PublicSystemInfo) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetProductName

`func (o *PublicSystemInfo) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *PublicSystemInfo) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *PublicSystemInfo) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *PublicSystemInfo) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *PublicSystemInfo) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *PublicSystemInfo) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetOperatingSystem

`func (o *PublicSystemInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *PublicSystemInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *PublicSystemInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *PublicSystemInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### SetOperatingSystemNil

`func (o *PublicSystemInfo) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *PublicSystemInfo) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetId

`func (o *PublicSystemInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicSystemInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicSystemInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicSystemInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PublicSystemInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PublicSystemInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetStartupWizardCompleted

`func (o *PublicSystemInfo) GetStartupWizardCompleted() bool`

GetStartupWizardCompleted returns the StartupWizardCompleted field if non-nil, zero value otherwise.

### GetStartupWizardCompletedOk

`func (o *PublicSystemInfo) GetStartupWizardCompletedOk() (*bool, bool)`

GetStartupWizardCompletedOk returns a tuple with the StartupWizardCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupWizardCompleted

`func (o *PublicSystemInfo) SetStartupWizardCompleted(v bool)`

SetStartupWizardCompleted sets StartupWizardCompleted field to given value.

### HasStartupWizardCompleted

`func (o *PublicSystemInfo) HasStartupWizardCompleted() bool`

HasStartupWizardCompleted returns a boolean if a field has been set.

### SetStartupWizardCompletedNil

`func (o *PublicSystemInfo) SetStartupWizardCompletedNil(b bool)`

 SetStartupWizardCompletedNil sets the value for StartupWizardCompleted to be an explicit nil

### UnsetStartupWizardCompleted
`func (o *PublicSystemInfo) UnsetStartupWizardCompleted()`

UnsetStartupWizardCompleted ensures that no value is present for StartupWizardCompleted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


