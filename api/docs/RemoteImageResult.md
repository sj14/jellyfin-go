# RemoteImageResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to [**[]RemoteImageInfo**](RemoteImageInfo.md) | Gets or sets the images. | [optional] 
**TotalRecordCount** | Pointer to **int32** | Gets or sets the total record count. | [optional] 
**Providers** | Pointer to **[]string** | Gets or sets the providers. | [optional] 

## Methods

### NewRemoteImageResult

`func NewRemoteImageResult() *RemoteImageResult`

NewRemoteImageResult instantiates a new RemoteImageResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteImageResultWithDefaults

`func NewRemoteImageResultWithDefaults() *RemoteImageResult`

NewRemoteImageResultWithDefaults instantiates a new RemoteImageResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *RemoteImageResult) GetImages() []RemoteImageInfo`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *RemoteImageResult) GetImagesOk() (*[]RemoteImageInfo, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *RemoteImageResult) SetImages(v []RemoteImageInfo)`

SetImages sets Images field to given value.

### HasImages

`func (o *RemoteImageResult) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *RemoteImageResult) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *RemoteImageResult) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetTotalRecordCount

`func (o *RemoteImageResult) GetTotalRecordCount() int32`

GetTotalRecordCount returns the TotalRecordCount field if non-nil, zero value otherwise.

### GetTotalRecordCountOk

`func (o *RemoteImageResult) GetTotalRecordCountOk() (*int32, bool)`

GetTotalRecordCountOk returns a tuple with the TotalRecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRecordCount

`func (o *RemoteImageResult) SetTotalRecordCount(v int32)`

SetTotalRecordCount sets TotalRecordCount field to given value.

### HasTotalRecordCount

`func (o *RemoteImageResult) HasTotalRecordCount() bool`

HasTotalRecordCount returns a boolean if a field has been set.

### GetProviders

`func (o *RemoteImageResult) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *RemoteImageResult) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *RemoteImageResult) SetProviders(v []string)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *RemoteImageResult) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### SetProvidersNil

`func (o *RemoteImageResult) SetProvidersNil(b bool)`

 SetProvidersNil sets the value for Providers to be an explicit nil

### UnsetProviders
`func (o *RemoteImageResult) UnsetProviders()`

UnsetProviders ensures that no value is present for Providers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


