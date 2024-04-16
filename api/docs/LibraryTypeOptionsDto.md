# LibraryTypeOptionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | Gets or sets the type. | [optional] 
**MetadataFetchers** | Pointer to [**[]LibraryOptionInfoDto**](LibraryOptionInfoDto.md) | Gets or sets the metadata fetchers. | [optional] 
**ImageFetchers** | Pointer to [**[]LibraryOptionInfoDto**](LibraryOptionInfoDto.md) | Gets or sets the image fetchers. | [optional] 
**SupportedImageTypes** | Pointer to [**[]ImageType**](ImageType.md) | Gets or sets the supported image types. | [optional] 
**DefaultImageOptions** | Pointer to [**[]ImageOption**](ImageOption.md) | Gets or sets the default image options. | [optional] 

## Methods

### NewLibraryTypeOptionsDto

`func NewLibraryTypeOptionsDto() *LibraryTypeOptionsDto`

NewLibraryTypeOptionsDto instantiates a new LibraryTypeOptionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryTypeOptionsDtoWithDefaults

`func NewLibraryTypeOptionsDtoWithDefaults() *LibraryTypeOptionsDto`

NewLibraryTypeOptionsDtoWithDefaults instantiates a new LibraryTypeOptionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LibraryTypeOptionsDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LibraryTypeOptionsDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LibraryTypeOptionsDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LibraryTypeOptionsDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *LibraryTypeOptionsDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *LibraryTypeOptionsDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetMetadataFetchers

`func (o *LibraryTypeOptionsDto) GetMetadataFetchers() []LibraryOptionInfoDto`

GetMetadataFetchers returns the MetadataFetchers field if non-nil, zero value otherwise.

### GetMetadataFetchersOk

`func (o *LibraryTypeOptionsDto) GetMetadataFetchersOk() (*[]LibraryOptionInfoDto, bool)`

GetMetadataFetchersOk returns a tuple with the MetadataFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataFetchers

`func (o *LibraryTypeOptionsDto) SetMetadataFetchers(v []LibraryOptionInfoDto)`

SetMetadataFetchers sets MetadataFetchers field to given value.

### HasMetadataFetchers

`func (o *LibraryTypeOptionsDto) HasMetadataFetchers() bool`

HasMetadataFetchers returns a boolean if a field has been set.

### GetImageFetchers

`func (o *LibraryTypeOptionsDto) GetImageFetchers() []LibraryOptionInfoDto`

GetImageFetchers returns the ImageFetchers field if non-nil, zero value otherwise.

### GetImageFetchersOk

`func (o *LibraryTypeOptionsDto) GetImageFetchersOk() (*[]LibraryOptionInfoDto, bool)`

GetImageFetchersOk returns a tuple with the ImageFetchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFetchers

`func (o *LibraryTypeOptionsDto) SetImageFetchers(v []LibraryOptionInfoDto)`

SetImageFetchers sets ImageFetchers field to given value.

### HasImageFetchers

`func (o *LibraryTypeOptionsDto) HasImageFetchers() bool`

HasImageFetchers returns a boolean if a field has been set.

### GetSupportedImageTypes

`func (o *LibraryTypeOptionsDto) GetSupportedImageTypes() []ImageType`

GetSupportedImageTypes returns the SupportedImageTypes field if non-nil, zero value otherwise.

### GetSupportedImageTypesOk

`func (o *LibraryTypeOptionsDto) GetSupportedImageTypesOk() (*[]ImageType, bool)`

GetSupportedImageTypesOk returns a tuple with the SupportedImageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedImageTypes

`func (o *LibraryTypeOptionsDto) SetSupportedImageTypes(v []ImageType)`

SetSupportedImageTypes sets SupportedImageTypes field to given value.

### HasSupportedImageTypes

`func (o *LibraryTypeOptionsDto) HasSupportedImageTypes() bool`

HasSupportedImageTypes returns a boolean if a field has been set.

### GetDefaultImageOptions

`func (o *LibraryTypeOptionsDto) GetDefaultImageOptions() []ImageOption`

GetDefaultImageOptions returns the DefaultImageOptions field if non-nil, zero value otherwise.

### GetDefaultImageOptionsOk

`func (o *LibraryTypeOptionsDto) GetDefaultImageOptionsOk() (*[]ImageOption, bool)`

GetDefaultImageOptionsOk returns a tuple with the DefaultImageOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageOptions

`func (o *LibraryTypeOptionsDto) SetDefaultImageOptions(v []ImageOption)`

SetDefaultImageOptions sets DefaultImageOptions field to given value.

### HasDefaultImageOptions

`func (o *LibraryTypeOptionsDto) HasDefaultImageOptions() bool`

HasDefaultImageOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


