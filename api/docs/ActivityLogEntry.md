# ActivityLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Gets or sets the identifier. | [optional] 
**Name** | Pointer to **string** | Gets or sets the name. | [optional] 
**Overview** | Pointer to **NullableString** | Gets or sets the overview. | [optional] 
**ShortOverview** | Pointer to **NullableString** | Gets or sets the short overview. | [optional] 
**Type** | Pointer to **string** | Gets or sets the type. | [optional] 
**ItemId** | Pointer to **NullableString** | Gets or sets the item identifier. | [optional] 
**Date** | Pointer to **time.Time** | Gets or sets the date. | [optional] 
**UserId** | Pointer to **string** | Gets or sets the user identifier. | [optional] 
**UserPrimaryImageTag** | Pointer to **NullableString** | Gets or sets the user primary image tag. | [optional] 
**Severity** | Pointer to [**LogLevel**](LogLevel.md) | Gets or sets the log severity. | [optional] 

## Methods

### NewActivityLogEntry

`func NewActivityLogEntry() *ActivityLogEntry`

NewActivityLogEntry instantiates a new ActivityLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogEntryWithDefaults

`func NewActivityLogEntryWithDefaults() *ActivityLogEntry`

NewActivityLogEntryWithDefaults instantiates a new ActivityLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityLogEntry) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityLogEntry) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityLogEntry) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityLogEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ActivityLogEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivityLogEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivityLogEntry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActivityLogEntry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverview

`func (o *ActivityLogEntry) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *ActivityLogEntry) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *ActivityLogEntry) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *ActivityLogEntry) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *ActivityLogEntry) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *ActivityLogEntry) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetShortOverview

`func (o *ActivityLogEntry) GetShortOverview() string`

GetShortOverview returns the ShortOverview field if non-nil, zero value otherwise.

### GetShortOverviewOk

`func (o *ActivityLogEntry) GetShortOverviewOk() (*string, bool)`

GetShortOverviewOk returns a tuple with the ShortOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortOverview

`func (o *ActivityLogEntry) SetShortOverview(v string)`

SetShortOverview sets ShortOverview field to given value.

### HasShortOverview

`func (o *ActivityLogEntry) HasShortOverview() bool`

HasShortOverview returns a boolean if a field has been set.

### SetShortOverviewNil

`func (o *ActivityLogEntry) SetShortOverviewNil(b bool)`

 SetShortOverviewNil sets the value for ShortOverview to be an explicit nil

### UnsetShortOverview
`func (o *ActivityLogEntry) UnsetShortOverview()`

UnsetShortOverview ensures that no value is present for ShortOverview, not even an explicit nil
### GetType

`func (o *ActivityLogEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityLogEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityLogEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActivityLogEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetItemId

`func (o *ActivityLogEntry) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ActivityLogEntry) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ActivityLogEntry) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ActivityLogEntry) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### SetItemIdNil

`func (o *ActivityLogEntry) SetItemIdNil(b bool)`

 SetItemIdNil sets the value for ItemId to be an explicit nil

### UnsetItemId
`func (o *ActivityLogEntry) UnsetItemId()`

UnsetItemId ensures that no value is present for ItemId, not even an explicit nil
### GetDate

`func (o *ActivityLogEntry) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ActivityLogEntry) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ActivityLogEntry) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *ActivityLogEntry) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetUserId

`func (o *ActivityLogEntry) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ActivityLogEntry) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ActivityLogEntry) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ActivityLogEntry) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserPrimaryImageTag

`func (o *ActivityLogEntry) GetUserPrimaryImageTag() string`

GetUserPrimaryImageTag returns the UserPrimaryImageTag field if non-nil, zero value otherwise.

### GetUserPrimaryImageTagOk

`func (o *ActivityLogEntry) GetUserPrimaryImageTagOk() (*string, bool)`

GetUserPrimaryImageTagOk returns a tuple with the UserPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrimaryImageTag

`func (o *ActivityLogEntry) SetUserPrimaryImageTag(v string)`

SetUserPrimaryImageTag sets UserPrimaryImageTag field to given value.

### HasUserPrimaryImageTag

`func (o *ActivityLogEntry) HasUserPrimaryImageTag() bool`

HasUserPrimaryImageTag returns a boolean if a field has been set.

### SetUserPrimaryImageTagNil

`func (o *ActivityLogEntry) SetUserPrimaryImageTagNil(b bool)`

 SetUserPrimaryImageTagNil sets the value for UserPrimaryImageTag to be an explicit nil

### UnsetUserPrimaryImageTag
`func (o *ActivityLogEntry) UnsetUserPrimaryImageTag()`

UnsetUserPrimaryImageTag ensures that no value is present for UserPrimaryImageTag, not even an explicit nil
### GetSeverity

`func (o *ActivityLogEntry) GetSeverity() LogLevel`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ActivityLogEntry) GetSeverityOk() (*LogLevel, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ActivityLogEntry) SetSeverity(v LogLevel)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ActivityLogEntry) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


