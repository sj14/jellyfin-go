# NotificationsSummaryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnreadCount** | Pointer to **int32** | Gets or sets the number of unread notifications. | [optional] 
**MaxUnreadNotificationLevel** | Pointer to [**NullableNotificationLevel**](NotificationLevel.md) | Gets or sets the maximum unread notification level. | [optional] 

## Methods

### NewNotificationsSummaryDto

`func NewNotificationsSummaryDto() *NotificationsSummaryDto`

NewNotificationsSummaryDto instantiates a new NotificationsSummaryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSummaryDtoWithDefaults

`func NewNotificationsSummaryDtoWithDefaults() *NotificationsSummaryDto`

NewNotificationsSummaryDtoWithDefaults instantiates a new NotificationsSummaryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnreadCount

`func (o *NotificationsSummaryDto) GetUnreadCount() int32`

GetUnreadCount returns the UnreadCount field if non-nil, zero value otherwise.

### GetUnreadCountOk

`func (o *NotificationsSummaryDto) GetUnreadCountOk() (*int32, bool)`

GetUnreadCountOk returns a tuple with the UnreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadCount

`func (o *NotificationsSummaryDto) SetUnreadCount(v int32)`

SetUnreadCount sets UnreadCount field to given value.

### HasUnreadCount

`func (o *NotificationsSummaryDto) HasUnreadCount() bool`

HasUnreadCount returns a boolean if a field has been set.

### GetMaxUnreadNotificationLevel

`func (o *NotificationsSummaryDto) GetMaxUnreadNotificationLevel() NotificationLevel`

GetMaxUnreadNotificationLevel returns the MaxUnreadNotificationLevel field if non-nil, zero value otherwise.

### GetMaxUnreadNotificationLevelOk

`func (o *NotificationsSummaryDto) GetMaxUnreadNotificationLevelOk() (*NotificationLevel, bool)`

GetMaxUnreadNotificationLevelOk returns a tuple with the MaxUnreadNotificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnreadNotificationLevel

`func (o *NotificationsSummaryDto) SetMaxUnreadNotificationLevel(v NotificationLevel)`

SetMaxUnreadNotificationLevel sets MaxUnreadNotificationLevel field to given value.

### HasMaxUnreadNotificationLevel

`func (o *NotificationsSummaryDto) HasMaxUnreadNotificationLevel() bool`

HasMaxUnreadNotificationLevel returns a boolean if a field has been set.

### SetMaxUnreadNotificationLevelNil

`func (o *NotificationsSummaryDto) SetMaxUnreadNotificationLevelNil(b bool)`

 SetMaxUnreadNotificationLevelNil sets the value for MaxUnreadNotificationLevel to be an explicit nil

### UnsetMaxUnreadNotificationLevel
`func (o *NotificationsSummaryDto) UnsetMaxUnreadNotificationLevel()`

UnsetMaxUnreadNotificationLevel ensures that no value is present for MaxUnreadNotificationLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


