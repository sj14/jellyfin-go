# NotificationOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** |  | [optional] 
**DisabledMonitorUsers** | Pointer to **[]string** | Gets or sets user Ids to not monitor (it&#39;s opt out). | [optional] 
**SendToUsers** | Pointer to **[]string** | Gets or sets user Ids to send to (if SendToUserMode &#x3D;&#x3D; Custom). | [optional] 
**Enabled** | Pointer to **bool** | Gets or sets a value indicating whether this MediaBrowser.Model.Notifications.NotificationOption is enabled. | [optional] 
**DisabledServices** | Pointer to **[]string** | Gets or sets the disabled services. | [optional] 
**SendToUserMode** | Pointer to [**SendToUserType**](SendToUserType.md) | Gets or sets the send to user mode. | [optional] 

## Methods

### NewNotificationOption

`func NewNotificationOption() *NotificationOption`

NewNotificationOption instantiates a new NotificationOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationOptionWithDefaults

`func NewNotificationOptionWithDefaults() *NotificationOption`

NewNotificationOptionWithDefaults instantiates a new NotificationOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationOption) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *NotificationOption) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *NotificationOption) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDisabledMonitorUsers

`func (o *NotificationOption) GetDisabledMonitorUsers() []string`

GetDisabledMonitorUsers returns the DisabledMonitorUsers field if non-nil, zero value otherwise.

### GetDisabledMonitorUsersOk

`func (o *NotificationOption) GetDisabledMonitorUsersOk() (*[]string, bool)`

GetDisabledMonitorUsersOk returns a tuple with the DisabledMonitorUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledMonitorUsers

`func (o *NotificationOption) SetDisabledMonitorUsers(v []string)`

SetDisabledMonitorUsers sets DisabledMonitorUsers field to given value.

### HasDisabledMonitorUsers

`func (o *NotificationOption) HasDisabledMonitorUsers() bool`

HasDisabledMonitorUsers returns a boolean if a field has been set.

### GetSendToUsers

`func (o *NotificationOption) GetSendToUsers() []string`

GetSendToUsers returns the SendToUsers field if non-nil, zero value otherwise.

### GetSendToUsersOk

`func (o *NotificationOption) GetSendToUsersOk() (*[]string, bool)`

GetSendToUsersOk returns a tuple with the SendToUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendToUsers

`func (o *NotificationOption) SetSendToUsers(v []string)`

SetSendToUsers sets SendToUsers field to given value.

### HasSendToUsers

`func (o *NotificationOption) HasSendToUsers() bool`

HasSendToUsers returns a boolean if a field has been set.

### GetEnabled

`func (o *NotificationOption) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NotificationOption) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NotificationOption) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NotificationOption) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDisabledServices

`func (o *NotificationOption) GetDisabledServices() []string`

GetDisabledServices returns the DisabledServices field if non-nil, zero value otherwise.

### GetDisabledServicesOk

`func (o *NotificationOption) GetDisabledServicesOk() (*[]string, bool)`

GetDisabledServicesOk returns a tuple with the DisabledServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledServices

`func (o *NotificationOption) SetDisabledServices(v []string)`

SetDisabledServices sets DisabledServices field to given value.

### HasDisabledServices

`func (o *NotificationOption) HasDisabledServices() bool`

HasDisabledServices returns a boolean if a field has been set.

### GetSendToUserMode

`func (o *NotificationOption) GetSendToUserMode() SendToUserType`

GetSendToUserMode returns the SendToUserMode field if non-nil, zero value otherwise.

### GetSendToUserModeOk

`func (o *NotificationOption) GetSendToUserModeOk() (*SendToUserType, bool)`

GetSendToUserModeOk returns a tuple with the SendToUserMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendToUserMode

`func (o *NotificationOption) SetSendToUserMode(v SendToUserType)`

SetSendToUserMode sets SendToUserMode field to given value.

### HasSendToUserMode

`func (o *NotificationOption) HasSendToUserMode() bool`

HasSendToUserMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


