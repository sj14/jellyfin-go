# UserDataChangeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | Gets or sets the user id. | [optional] 
**UserDataList** | Pointer to [**[]UserItemDataDto**](UserItemDataDto.md) | Gets or sets the user data list. | [optional] 

## Methods

### NewUserDataChangeInfo

`func NewUserDataChangeInfo() *UserDataChangeInfo`

NewUserDataChangeInfo instantiates a new UserDataChangeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataChangeInfoWithDefaults

`func NewUserDataChangeInfoWithDefaults() *UserDataChangeInfo`

NewUserDataChangeInfoWithDefaults instantiates a new UserDataChangeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserDataChangeInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserDataChangeInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserDataChangeInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserDataChangeInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserDataList

`func (o *UserDataChangeInfo) GetUserDataList() []UserItemDataDto`

GetUserDataList returns the UserDataList field if non-nil, zero value otherwise.

### GetUserDataListOk

`func (o *UserDataChangeInfo) GetUserDataListOk() (*[]UserItemDataDto, bool)`

GetUserDataListOk returns a tuple with the UserDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataList

`func (o *UserDataChangeInfo) SetUserDataList(v []UserItemDataDto)`

SetUserDataList sets UserDataList field to given value.

### HasUserDataList

`func (o *UserDataChangeInfo) HasUserDataList() bool`

HasUserDataList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


