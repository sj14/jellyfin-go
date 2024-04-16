# NotificationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Gets or sets the notification ID. Defaults to an empty string. | [optional] 
**UserId** | Pointer to **string** | Gets or sets the notification&#39;s user ID. Defaults to an empty string. | [optional] 
**Date** | Pointer to **time.Time** | Gets or sets the notification date. | [optional] 
**IsRead** | Pointer to **bool** | Gets or sets a value indicating whether the notification has been read. Defaults to false. | [optional] 
**Name** | Pointer to **string** | Gets or sets the notification&#39;s name. Defaults to an empty string. | [optional] 
**Description** | Pointer to **string** | Gets or sets the notification&#39;s description. Defaults to an empty string. | [optional] 
**Url** | Pointer to **string** | Gets or sets the notification&#39;s URL. Defaults to an empty string. | [optional] 
**Level** | Pointer to [**NotificationLevel**](NotificationLevel.md) | Gets or sets the notification level. | [optional] 

## Methods

### NewNotificationDto

`func NewNotificationDto() *NotificationDto`

NewNotificationDto instantiates a new NotificationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationDtoWithDefaults

`func NewNotificationDtoWithDefaults() *NotificationDto`

NewNotificationDtoWithDefaults instantiates a new NotificationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *NotificationDto) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *NotificationDto) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *NotificationDto) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *NotificationDto) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDate

`func (o *NotificationDto) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *NotificationDto) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *NotificationDto) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *NotificationDto) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetIsRead

`func (o *NotificationDto) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *NotificationDto) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *NotificationDto) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.

### HasIsRead

`func (o *NotificationDto) HasIsRead() bool`

HasIsRead returns a boolean if a field has been set.

### GetName

`func (o *NotificationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NotificationDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NotificationDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NotificationDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NotificationDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *NotificationDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NotificationDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NotificationDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NotificationDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLevel

`func (o *NotificationDto) GetLevel() NotificationLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *NotificationDto) GetLevelOk() (*NotificationLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *NotificationDto) SetLevel(v NotificationLevel)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *NotificationDto) HasLevel() bool`

HasLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


