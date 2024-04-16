# PackageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Gets or sets the name. | [optional] 
**Description** | Pointer to **string** | Gets or sets a long description of the plugin containing features or helpful explanations. | [optional] 
**Overview** | Pointer to **string** | Gets or sets a short overview of what the plugin does. | [optional] 
**Owner** | Pointer to **string** | Gets or sets the owner. | [optional] 
**Category** | Pointer to **string** | Gets or sets the category. | [optional] 
**Guid** | Pointer to **string** | Gets or sets the guid of the assembly associated with this plugin.  This is used to identify the proper item for automatic updates. | [optional] 
**Versions** | Pointer to [**[]VersionInfo**](VersionInfo.md) | Gets or sets the versions. | [optional] 
**ImageUrl** | Pointer to **NullableString** | Gets or sets the image url for the package. | [optional] 

## Methods

### NewPackageInfo

`func NewPackageInfo() *PackageInfo`

NewPackageInfo instantiates a new PackageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageInfoWithDefaults

`func NewPackageInfoWithDefaults() *PackageInfo`

NewPackageInfoWithDefaults instantiates a new PackageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PackageInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PackageInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PackageInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PackageInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PackageInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackageInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackageInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PackageInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOverview

`func (o *PackageInfo) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *PackageInfo) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *PackageInfo) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *PackageInfo) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetOwner

`func (o *PackageInfo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PackageInfo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PackageInfo) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PackageInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCategory

`func (o *PackageInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PackageInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PackageInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PackageInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetGuid

`func (o *PackageInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *PackageInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *PackageInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *PackageInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetVersions

`func (o *PackageInfo) GetVersions() []VersionInfo`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *PackageInfo) GetVersionsOk() (*[]VersionInfo, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *PackageInfo) SetVersions(v []VersionInfo)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *PackageInfo) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetImageUrl

`func (o *PackageInfo) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *PackageInfo) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *PackageInfo) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *PackageInfo) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *PackageInfo) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *PackageInfo) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


