# InstallationInfoPackageInfo

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

### NewInstallationInfoPackageInfo

`func NewInstallationInfoPackageInfo() *InstallationInfoPackageInfo`

NewInstallationInfoPackageInfo instantiates a new InstallationInfoPackageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallationInfoPackageInfoWithDefaults

`func NewInstallationInfoPackageInfoWithDefaults() *InstallationInfoPackageInfo`

NewInstallationInfoPackageInfoWithDefaults instantiates a new InstallationInfoPackageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstallationInfoPackageInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstallationInfoPackageInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstallationInfoPackageInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstallationInfoPackageInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InstallationInfoPackageInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstallationInfoPackageInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstallationInfoPackageInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstallationInfoPackageInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOverview

`func (o *InstallationInfoPackageInfo) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *InstallationInfoPackageInfo) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *InstallationInfoPackageInfo) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *InstallationInfoPackageInfo) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetOwner

`func (o *InstallationInfoPackageInfo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *InstallationInfoPackageInfo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *InstallationInfoPackageInfo) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *InstallationInfoPackageInfo) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCategory

`func (o *InstallationInfoPackageInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InstallationInfoPackageInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InstallationInfoPackageInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InstallationInfoPackageInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetGuid

`func (o *InstallationInfoPackageInfo) GetGuid() string`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *InstallationInfoPackageInfo) GetGuidOk() (*string, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *InstallationInfoPackageInfo) SetGuid(v string)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *InstallationInfoPackageInfo) HasGuid() bool`

HasGuid returns a boolean if a field has been set.

### GetVersions

`func (o *InstallationInfoPackageInfo) GetVersions() []VersionInfo`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *InstallationInfoPackageInfo) GetVersionsOk() (*[]VersionInfo, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *InstallationInfoPackageInfo) SetVersions(v []VersionInfo)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *InstallationInfoPackageInfo) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetImageUrl

`func (o *InstallationInfoPackageInfo) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *InstallationInfoPackageInfo) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *InstallationInfoPackageInfo) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *InstallationInfoPackageInfo) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *InstallationInfoPackageInfo) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *InstallationInfoPackageInfo) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


