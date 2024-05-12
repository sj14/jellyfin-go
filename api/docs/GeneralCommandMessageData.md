# GeneralCommandMessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**GeneralCommandType**](GeneralCommandType.md) | This exists simply to identify a set of known commands. | [optional] 
**ControllingUserId** | Pointer to **string** |  | [optional] 
**Arguments** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewGeneralCommandMessageData

`func NewGeneralCommandMessageData() *GeneralCommandMessageData`

NewGeneralCommandMessageData instantiates a new GeneralCommandMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralCommandMessageDataWithDefaults

`func NewGeneralCommandMessageDataWithDefaults() *GeneralCommandMessageData`

NewGeneralCommandMessageDataWithDefaults instantiates a new GeneralCommandMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GeneralCommandMessageData) GetName() GeneralCommandType`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeneralCommandMessageData) GetNameOk() (*GeneralCommandType, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeneralCommandMessageData) SetName(v GeneralCommandType)`

SetName sets Name field to given value.

### HasName

`func (o *GeneralCommandMessageData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetControllingUserId

`func (o *GeneralCommandMessageData) GetControllingUserId() string`

GetControllingUserId returns the ControllingUserId field if non-nil, zero value otherwise.

### GetControllingUserIdOk

`func (o *GeneralCommandMessageData) GetControllingUserIdOk() (*string, bool)`

GetControllingUserIdOk returns a tuple with the ControllingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllingUserId

`func (o *GeneralCommandMessageData) SetControllingUserId(v string)`

SetControllingUserId sets ControllingUserId field to given value.

### HasControllingUserId

`func (o *GeneralCommandMessageData) HasControllingUserId() bool`

HasControllingUserId returns a boolean if a field has been set.

### GetArguments

`func (o *GeneralCommandMessageData) GetArguments() map[string]string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *GeneralCommandMessageData) GetArgumentsOk() (*map[string]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *GeneralCommandMessageData) SetArguments(v map[string]string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *GeneralCommandMessageData) HasArguments() bool`

HasArguments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


