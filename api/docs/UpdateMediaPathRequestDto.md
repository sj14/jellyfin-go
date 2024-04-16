# UpdateMediaPathRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Gets or sets the library name. | 
**PathInfo** | [**MediaPathInfo**](MediaPathInfo.md) | Gets or sets library folder path information. | 

## Methods

### NewUpdateMediaPathRequestDto

`func NewUpdateMediaPathRequestDto(name string, pathInfo MediaPathInfo, ) *UpdateMediaPathRequestDto`

NewUpdateMediaPathRequestDto instantiates a new UpdateMediaPathRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMediaPathRequestDtoWithDefaults

`func NewUpdateMediaPathRequestDtoWithDefaults() *UpdateMediaPathRequestDto`

NewUpdateMediaPathRequestDtoWithDefaults instantiates a new UpdateMediaPathRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateMediaPathRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMediaPathRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMediaPathRequestDto) SetName(v string)`

SetName sets Name field to given value.


### GetPathInfo

`func (o *UpdateMediaPathRequestDto) GetPathInfo() MediaPathInfo`

GetPathInfo returns the PathInfo field if non-nil, zero value otherwise.

### GetPathInfoOk

`func (o *UpdateMediaPathRequestDto) GetPathInfoOk() (*MediaPathInfo, bool)`

GetPathInfoOk returns a tuple with the PathInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathInfo

`func (o *UpdateMediaPathRequestDto) SetPathInfo(v MediaPathInfo)`

SetPathInfo sets PathInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


