# LiveTvServiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**HomePageUrl** | Pointer to **NullableString** | Gets or sets the home page URL. | [optional] 
**Status** | Pointer to [**LiveTvServiceStatus**](LiveTvServiceStatus.md) | Gets or sets the status. | [optional] 
**StatusMessage** | Pointer to **NullableString** | Gets or sets the status message. | [optional] 
**Version** | Pointer to **NullableString** | Gets or sets the version. | [optional] 
**HasUpdateAvailable** | Pointer to **bool** | Gets or sets a value indicating whether this instance has update available. | [optional] 
**IsVisible** | Pointer to **bool** | Gets or sets a value indicating whether this instance is visible. | [optional] 
**Tuners** | Pointer to **[]string** |  | [optional] 

## Methods

### NewLiveTvServiceInfo

`func NewLiveTvServiceInfo() *LiveTvServiceInfo`

NewLiveTvServiceInfo instantiates a new LiveTvServiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveTvServiceInfoWithDefaults

`func NewLiveTvServiceInfoWithDefaults() *LiveTvServiceInfo`

NewLiveTvServiceInfoWithDefaults instantiates a new LiveTvServiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LiveTvServiceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LiveTvServiceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LiveTvServiceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LiveTvServiceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LiveTvServiceInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LiveTvServiceInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetHomePageUrl

`func (o *LiveTvServiceInfo) GetHomePageUrl() string`

GetHomePageUrl returns the HomePageUrl field if non-nil, zero value otherwise.

### GetHomePageUrlOk

`func (o *LiveTvServiceInfo) GetHomePageUrlOk() (*string, bool)`

GetHomePageUrlOk returns a tuple with the HomePageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomePageUrl

`func (o *LiveTvServiceInfo) SetHomePageUrl(v string)`

SetHomePageUrl sets HomePageUrl field to given value.

### HasHomePageUrl

`func (o *LiveTvServiceInfo) HasHomePageUrl() bool`

HasHomePageUrl returns a boolean if a field has been set.

### SetHomePageUrlNil

`func (o *LiveTvServiceInfo) SetHomePageUrlNil(b bool)`

 SetHomePageUrlNil sets the value for HomePageUrl to be an explicit nil

### UnsetHomePageUrl
`func (o *LiveTvServiceInfo) UnsetHomePageUrl()`

UnsetHomePageUrl ensures that no value is present for HomePageUrl, not even an explicit nil
### GetStatus

`func (o *LiveTvServiceInfo) GetStatus() LiveTvServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LiveTvServiceInfo) GetStatusOk() (*LiveTvServiceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LiveTvServiceInfo) SetStatus(v LiveTvServiceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LiveTvServiceInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *LiveTvServiceInfo) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *LiveTvServiceInfo) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *LiveTvServiceInfo) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *LiveTvServiceInfo) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *LiveTvServiceInfo) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *LiveTvServiceInfo) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetVersion

`func (o *LiveTvServiceInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LiveTvServiceInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LiveTvServiceInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LiveTvServiceInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *LiveTvServiceInfo) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *LiveTvServiceInfo) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetHasUpdateAvailable

`func (o *LiveTvServiceInfo) GetHasUpdateAvailable() bool`

GetHasUpdateAvailable returns the HasUpdateAvailable field if non-nil, zero value otherwise.

### GetHasUpdateAvailableOk

`func (o *LiveTvServiceInfo) GetHasUpdateAvailableOk() (*bool, bool)`

GetHasUpdateAvailableOk returns a tuple with the HasUpdateAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdateAvailable

`func (o *LiveTvServiceInfo) SetHasUpdateAvailable(v bool)`

SetHasUpdateAvailable sets HasUpdateAvailable field to given value.

### HasHasUpdateAvailable

`func (o *LiveTvServiceInfo) HasHasUpdateAvailable() bool`

HasHasUpdateAvailable returns a boolean if a field has been set.

### GetIsVisible

`func (o *LiveTvServiceInfo) GetIsVisible() bool`

GetIsVisible returns the IsVisible field if non-nil, zero value otherwise.

### GetIsVisibleOk

`func (o *LiveTvServiceInfo) GetIsVisibleOk() (*bool, bool)`

GetIsVisibleOk returns a tuple with the IsVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisible

`func (o *LiveTvServiceInfo) SetIsVisible(v bool)`

SetIsVisible sets IsVisible field to given value.

### HasIsVisible

`func (o *LiveTvServiceInfo) HasIsVisible() bool`

HasIsVisible returns a boolean if a field has been set.

### GetTuners

`func (o *LiveTvServiceInfo) GetTuners() []string`

GetTuners returns the Tuners field if non-nil, zero value otherwise.

### GetTunersOk

`func (o *LiveTvServiceInfo) GetTunersOk() (*[]string, bool)`

GetTunersOk returns a tuple with the Tuners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuners

`func (o *LiveTvServiceInfo) SetTuners(v []string)`

SetTuners sets Tuners field to given value.

### HasTuners

`func (o *LiveTvServiceInfo) HasTuners() bool`

HasTuners returns a boolean if a field has been set.

### SetTunersNil

`func (o *LiveTvServiceInfo) SetTunersNil(b bool)`

 SetTunersNil sets the value for Tuners to be an explicit nil

### UnsetTuners
`func (o *LiveTvServiceInfo) UnsetTuners()`

UnsetTuners ensures that no value is present for Tuners, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


