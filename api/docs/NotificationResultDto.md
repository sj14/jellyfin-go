# NotificationResultDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notifications** | Pointer to [**[]NotificationDto**](NotificationDto.md) | Gets or sets the current page of notifications. | [optional] 
**TotalRecordCount** | Pointer to **int32** | Gets or sets the total number of notifications. | [optional] 

## Methods

### NewNotificationResultDto

`func NewNotificationResultDto() *NotificationResultDto`

NewNotificationResultDto instantiates a new NotificationResultDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationResultDtoWithDefaults

`func NewNotificationResultDtoWithDefaults() *NotificationResultDto`

NewNotificationResultDtoWithDefaults instantiates a new NotificationResultDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifications

`func (o *NotificationResultDto) GetNotifications() []NotificationDto`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *NotificationResultDto) GetNotificationsOk() (*[]NotificationDto, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *NotificationResultDto) SetNotifications(v []NotificationDto)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *NotificationResultDto) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetTotalRecordCount

`func (o *NotificationResultDto) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *NotificationResultDto) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *NotificationResultDto) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *NotificationResultDto) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


