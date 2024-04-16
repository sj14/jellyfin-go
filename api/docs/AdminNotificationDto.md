# AdminNotificationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the notification name. | [optional] 
**Description** | Pointer to **NullableString** | Gets or sets the notification description. | [optional] 
**NotificationLevel** | Pointer to [**NullableNotificationLevel**](NotificationLevel.md) | Gets or sets the notification level. | [optional] 
**Url** | Pointer to **NullableString** | Gets or sets the notification url. | [optional] 

## Methods

### NewAdminNotificationDto

`func NewAdminNotificationDto() *AdminNotificationDto`

NewAdminNotificationDto instantiates a new AdminNotificationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminNotificationDtoWithDefaults

`func NewAdminNotificationDtoWithDefaults() *AdminNotificationDto`

NewAdminNotificationDtoWithDefaults instantiates a new AdminNotificationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdminNotificationDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminNotificationDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminNotificationDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminNotificationDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdminNotificationDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdminNotificationDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AdminNotificationDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminNotificationDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminNotificationDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdminNotificationDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AdminNotificationDto) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AdminNotificationDto) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetNotificationLevel

`func (o *AdminNotificationDto) GetNotificationLevel() NotificationLevel`

GetNotificationLevel returns the NotificationLevel field if non-nil, zero value otherwise.

### GetNotificationLevelOk

`func (o *AdminNotificationDto) GetNotificationLevelOk() (*NotificationLevel, bool)`

GetNotificationLevelOk returns a tuple with the NotificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationLevel

`func (o *AdminNotificationDto) SetNotificationLevel(v NotificationLevel)`

SetNotificationLevel sets NotificationLevel field to given value.

### HasNotificationLevel

`func (o *AdminNotificationDto) HasNotificationLevel() bool`

HasNotificationLevel returns a boolean if a field has been set.

### SetNotificationLevelNil

`func (o *AdminNotificationDto) SetNotificationLevelNil(b bool)`

 SetNotificationLevelNil sets the value for NotificationLevel to be an explicit nil

### UnsetNotificationLevel
`func (o *AdminNotificationDto) UnsetNotificationLevel()`

UnsetNotificationLevel ensures that no value is present for NotificationLevel, not even an explicit nil
### GetUrl

`func (o *AdminNotificationDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdminNotificationDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdminNotificationDto) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdminNotificationDto) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *AdminNotificationDto) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *AdminNotificationDto) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


