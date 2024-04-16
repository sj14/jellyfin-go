# ChannelFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Gets or sets the name. | [optional] 
**Id** | Pointer to **string** | Gets or sets the identifier. | [optional] 
**CanSearch** | Pointer to **bool** | Gets or sets a value indicating whether this instance can search. | [optional] 
**MediaTypes** | Pointer to [**[]ChannelMediaType**](ChannelMediaType.md) | Gets or sets the media types. | [optional] 
**ContentTypes** | Pointer to [**[]ChannelMediaContentType**](ChannelMediaContentType.md) | Gets or sets the content types. | [optional] 
**MaxPageSize** | Pointer to **NullableInt32** | Gets or sets the maximum number of records the channel allows retrieving at a time. | [optional] 
**AutoRefreshLevels** | Pointer to **NullableInt32** | Gets or sets the automatic refresh levels. | [optional] 
**DefaultSortFields** | Pointer to [**[]ChannelItemSortField**](ChannelItemSortField.md) | Gets or sets the default sort orders. | [optional] 
**SupportsSortOrderToggle** | Pointer to **bool** | Gets or sets a value indicating whether a sort ascending/descending toggle is supported. | [optional] 
**SupportsLatestMedia** | Pointer to **bool** | Gets or sets a value indicating whether [supports latest media]. | [optional] 
**CanFilter** | Pointer to **bool** | Gets or sets a value indicating whether this instance can filter. | [optional] 
**SupportsContentDownloading** | Pointer to **bool** | Gets or sets a value indicating whether [supports content downloading]. | [optional] 

## Methods

### NewChannelFeatures

`func NewChannelFeatures() *ChannelFeatures`

NewChannelFeatures instantiates a new ChannelFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelFeaturesWithDefaults

`func NewChannelFeaturesWithDefaults() *ChannelFeatures`

NewChannelFeaturesWithDefaults instantiates a new ChannelFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ChannelFeatures) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelFeatures) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelFeatures) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChannelFeatures) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ChannelFeatures) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelFeatures) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelFeatures) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelFeatures) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCanSearch

`func (o *ChannelFeatures) GetCanSearch() bool`

GetCanSearch returns the CanSearch field if non-nil, zero value otherwise.

### GetCanSearchOk

`func (o *ChannelFeatures) GetCanSearchOk() (*bool, bool)`

GetCanSearchOk returns a tuple with the CanSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanSearch

`func (o *ChannelFeatures) SetCanSearch(v bool)`

SetCanSearch sets CanSearch field to given value.

### HasCanSearch

`func (o *ChannelFeatures) HasCanSearch() bool`

HasCanSearch returns a boolean if a field has been set.

### GetMediaTypes

`func (o *ChannelFeatures) GetMediaTypes() []ChannelMediaType`

GetMediaTypes returns the MediaTypes field if non-nil, zero value otherwise.

### GetMediaTypesOk

`func (o *ChannelFeatures) GetMediaTypesOk() (*[]ChannelMediaType, bool)`

GetMediaTypesOk returns a tuple with the MediaTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaTypes

`func (o *ChannelFeatures) SetMediaTypes(v []ChannelMediaType)`

SetMediaTypes sets MediaTypes field to given value.

### HasMediaTypes

`func (o *ChannelFeatures) HasMediaTypes() bool`

HasMediaTypes returns a boolean if a field has been set.

### GetContentTypes

`func (o *ChannelFeatures) GetContentTypes() []ChannelMediaContentType`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *ChannelFeatures) GetContentTypesOk() (*[]ChannelMediaContentType, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *ChannelFeatures) SetContentTypes(v []ChannelMediaContentType)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *ChannelFeatures) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetMaxPageSize

`func (o *ChannelFeatures) GetMaxPageSize() int32`

GetMaxPageSize returns the MaxPageSize field if non-nil, zero value otherwise.

### GetMaxPageSizeOk

`func (o *ChannelFeatures) GetMaxPageSizeOk() (*int32, bool)`

GetMaxPageSizeOk returns a tuple with the MaxPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPageSize

`func (o *ChannelFeatures) SetMaxPageSize(v int32)`

SetMaxPageSize sets MaxPageSize field to given value.

### HasMaxPageSize

`func (o *ChannelFeatures) HasMaxPageSize() bool`

HasMaxPageSize returns a boolean if a field has been set.

### SetMaxPageSizeNil

`func (o *ChannelFeatures) SetMaxPageSizeNil(b bool)`

 SetMaxPageSizeNil sets the value for MaxPageSize to be an explicit nil

### UnsetMaxPageSize
`func (o *ChannelFeatures) UnsetMaxPageSize()`

UnsetMaxPageSize ensures that no value is present for MaxPageSize, not even an explicit nil
### GetAutoRefreshLevels

`func (o *ChannelFeatures) GetAutoRefreshLevels() int32`

GetAutoRefreshLevels returns the AutoRefreshLevels field if non-nil, zero value otherwise.

### GetAutoRefreshLevelsOk

`func (o *ChannelFeatures) GetAutoRefreshLevelsOk() (*int32, bool)`

GetAutoRefreshLevelsOk returns a tuple with the AutoRefreshLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRefreshLevels

`func (o *ChannelFeatures) SetAutoRefreshLevels(v int32)`

SetAutoRefreshLevels sets AutoRefreshLevels field to given value.

### HasAutoRefreshLevels

`func (o *ChannelFeatures) HasAutoRefreshLevels() bool`

HasAutoRefreshLevels returns a boolean if a field has been set.

### SetAutoRefreshLevelsNil

`func (o *ChannelFeatures) SetAutoRefreshLevelsNil(b bool)`

 SetAutoRefreshLevelsNil sets the value for AutoRefreshLevels to be an explicit nil

### UnsetAutoRefreshLevels
`func (o *ChannelFeatures) UnsetAutoRefreshLevels()`

UnsetAutoRefreshLevels ensures that no value is present for AutoRefreshLevels, not even an explicit nil
### GetDefaultSortFields

`func (o *ChannelFeatures) GetDefaultSortFields() []ChannelItemSortField`

GetDefaultSortFields returns the DefaultSortFields field if non-nil, zero value otherwise.

### GetDefaultSortFieldsOk

`func (o *ChannelFeatures) GetDefaultSortFieldsOk() (*[]ChannelItemSortField, bool)`

GetDefaultSortFieldsOk returns a tuple with the DefaultSortFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSortFields

`func (o *ChannelFeatures) SetDefaultSortFields(v []ChannelItemSortField)`

SetDefaultSortFields sets DefaultSortFields field to given value.

### HasDefaultSortFields

`func (o *ChannelFeatures) HasDefaultSortFields() bool`

HasDefaultSortFields returns a boolean if a field has been set.

### GetSupportsSortOrderToggle

`func (o *ChannelFeatures) GetSupportsSortOrderToggle() bool`

GetSupportsSortOrderToggle returns the SupportsSortOrderToggle field if non-nil, zero value otherwise.

### GetSupportsSortOrderToggleOk

`func (o *ChannelFeatures) GetSupportsSortOrderToggleOk() (*bool, bool)`

GetSupportsSortOrderToggleOk returns a tuple with the SupportsSortOrderToggle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsSortOrderToggle

`func (o *ChannelFeatures) SetSupportsSortOrderToggle(v bool)`

SetSupportsSortOrderToggle sets SupportsSortOrderToggle field to given value.

### HasSupportsSortOrderToggle

`func (o *ChannelFeatures) HasSupportsSortOrderToggle() bool`

HasSupportsSortOrderToggle returns a boolean if a field has been set.

### GetSupportsLatestMedia

`func (o *ChannelFeatures) GetSupportsLatestMedia() bool`

GetSupportsLatestMedia returns the SupportsLatestMedia field if non-nil, zero value otherwise.

### GetSupportsLatestMediaOk

`func (o *ChannelFeatures) GetSupportsLatestMediaOk() (*bool, bool)`

GetSupportsLatestMediaOk returns a tuple with the SupportsLatestMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsLatestMedia

`func (o *ChannelFeatures) SetSupportsLatestMedia(v bool)`

SetSupportsLatestMedia sets SupportsLatestMedia field to given value.

### HasSupportsLatestMedia

`func (o *ChannelFeatures) HasSupportsLatestMedia() bool`

HasSupportsLatestMedia returns a boolean if a field has been set.

### GetCanFilter

`func (o *ChannelFeatures) GetCanFilter() bool`

GetCanFilter returns the CanFilter field if non-nil, zero value otherwise.

### GetCanFilterOk

`func (o *ChannelFeatures) GetCanFilterOk() (*bool, bool)`

GetCanFilterOk returns a tuple with the CanFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanFilter

`func (o *ChannelFeatures) SetCanFilter(v bool)`

SetCanFilter sets CanFilter field to given value.

### HasCanFilter

`func (o *ChannelFeatures) HasCanFilter() bool`

HasCanFilter returns a boolean if a field has been set.

### GetSupportsContentDownloading

`func (o *ChannelFeatures) GetSupportsContentDownloading() bool`

GetSupportsContentDownloading returns the SupportsContentDownloading field if non-nil, zero value otherwise.

### GetSupportsContentDownloadingOk

`func (o *ChannelFeatures) GetSupportsContentDownloadingOk() (*bool, bool)`

GetSupportsContentDownloadingOk returns a tuple with the SupportsContentDownloading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsContentDownloading

`func (o *ChannelFeatures) SetSupportsContentDownloading(v bool)`

SetSupportsContentDownloading sets SupportsContentDownloading field to given value.

### HasSupportsContentDownloading

`func (o *ChannelFeatures) HasSupportsContentDownloading() bool`

HasSupportsContentDownloading returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


