# MediaUpdateInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updates** | Pointer to [**[]MediaUpdateInfoPathDto**](MediaUpdateInfoPathDto.md) | Gets or sets the list of updates. | [optional] 

## Methods

### NewMediaUpdateInfoDto

`func NewMediaUpdateInfoDto() *MediaUpdateInfoDto`

NewMediaUpdateInfoDto instantiates a new MediaUpdateInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaUpdateInfoDtoWithDefaults

`func NewMediaUpdateInfoDtoWithDefaults() *MediaUpdateInfoDto`

NewMediaUpdateInfoDtoWithDefaults instantiates a new MediaUpdateInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdates

`func (o *MediaUpdateInfoDto) GetUpdates() []MediaUpdateInfoPathDto`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *MediaUpdateInfoDto) GetUpdatesOk() (*[]MediaUpdateInfoPathDto, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *MediaUpdateInfoDto) SetUpdates(v []MediaUpdateInfoPathDto)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *MediaUpdateInfoDto) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


