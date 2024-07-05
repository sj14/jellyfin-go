# UpdateLibraryOptionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Gets or sets the library item id. | [optional] 
**LibraryOptions** | Pointer to [**NullableLibraryOptions**](LibraryOptions.md) | Gets or sets library options. | [optional] 

## Methods

### NewUpdateLibraryOptionsDto

`func NewUpdateLibraryOptionsDto() *UpdateLibraryOptionsDto`

NewUpdateLibraryOptionsDto instantiates a new UpdateLibraryOptionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLibraryOptionsDtoWithDefaults

`func NewUpdateLibraryOptionsDtoWithDefaults() *UpdateLibraryOptionsDto`

NewUpdateLibraryOptionsDtoWithDefaults instantiates a new UpdateLibraryOptionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateLibraryOptionsDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateLibraryOptionsDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateLibraryOptionsDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateLibraryOptionsDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLibraryOptions

`func (o *UpdateLibraryOptionsDto) GetLibraryOptions() LibraryOptions`

GetLibraryOptions returns the LibraryOptions field if non-nil, zero value otherwise.

### GetLibraryOptionsOk

`func (o *UpdateLibraryOptionsDto) GetLibraryOptionsOk() (*LibraryOptions, bool)`

GetLibraryOptionsOk returns a tuple with the LibraryOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryOptions

`func (o *UpdateLibraryOptionsDto) SetLibraryOptions(v LibraryOptions)`

SetLibraryOptions sets LibraryOptions field to given value.

### HasLibraryOptions

`func (o *UpdateLibraryOptionsDto) HasLibraryOptions() bool`

HasLibraryOptions returns a boolean if a field has been set.

### SetLibraryOptionsNil

`func (o *UpdateLibraryOptionsDto) SetLibraryOptionsNil(b bool)`

 SetLibraryOptionsNil sets the value for LibraryOptions to be an explicit nil

### UnsetLibraryOptions
`func (o *UpdateLibraryOptionsDto) UnsetLibraryOptions()`

UnsetLibraryOptions ensures that no value is present for LibraryOptions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


