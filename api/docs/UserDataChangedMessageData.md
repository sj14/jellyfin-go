# UserDataChangedMessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **NullableString** | Gets or sets the user id. | [optional] 
**UserDataList** | Pointer to [**[]UserItemDataDto**](UserItemDataDto.md) | Gets or sets the user data list. | [optional] 

## Methods

### NewUserDataChangedMessageData

`func NewUserDataChangedMessageData() *UserDataChangedMessageData`

NewUserDataChangedMessageData instantiates a new UserDataChangedMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataChangedMessageDataWithDefaults

`func NewUserDataChangedMessageDataWithDefaults() *UserDataChangedMessageData`

NewUserDataChangedMessageDataWithDefaults instantiates a new UserDataChangedMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserDataChangedMessageData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserDataChangedMessageData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserDataChangedMessageData) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserDataChangedMessageData) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *UserDataChangedMessageData) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UserDataChangedMessageData) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetUserDataList

`func (o *UserDataChangedMessageData) GetUserDataList() []UserItemDataDto`

GetUserDataList returns the UserDataList field if non-nil, zero value otherwise.

### GetUserDataListOk

`func (o *UserDataChangedMessageData) GetUserDataListOk() (*[]UserItemDataDto, bool)`

GetUserDataListOk returns a tuple with the UserDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataList

`func (o *UserDataChangedMessageData) SetUserDataList(v []UserItemDataDto)`

SetUserDataList sets UserDataList field to given value.

### HasUserDataList

`func (o *UserDataChangedMessageData) HasUserDataList() bool`

HasUserDataList returns a boolean if a field has been set.

### SetUserDataListNil

`func (o *UserDataChangedMessageData) SetUserDataListNil(b bool)`

 SetUserDataListNil sets the value for UserDataList to be an explicit nil

### UnsetUserDataList
`func (o *UserDataChangedMessageData) UnsetUserDataList()`

UnsetUserDataList ensures that no value is present for UserDataList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


