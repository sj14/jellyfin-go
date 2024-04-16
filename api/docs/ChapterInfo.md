# ChapterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartPositionTicks** | Pointer to **int64** | Gets or sets the start position ticks. | [optional] 
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**ImagePath** | Pointer to **NullableString** | Gets or sets the image path. | [optional] 
**ImageDateModified** | Pointer to **time.Time** |  | [optional] 
**ImageTag** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewChapterInfo

`func NewChapterInfo() *ChapterInfo`

NewChapterInfo instantiates a new ChapterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChapterInfoWithDefaults

`func NewChapterInfoWithDefaults() *ChapterInfo`

NewChapterInfoWithDefaults instantiates a new ChapterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartPositionTicks

`func (o *ChapterInfo) GetStartPositionTicks() int64`

GetStartPositionTicks returns the StartPositionTicks field if non-nil, zero value otherwise.

### GetStartPositionTicksOk

`func (o *ChapterInfo) GetStartPositionTicksOk() (*int64, bool)`

GetStartPositionTicksOk returns a tuple with the StartPositionTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPositionTicks

`func (o *ChapterInfo) SetStartPositionTicks(v int64)`

SetStartPositionTicks sets StartPositionTicks field to given value.

### HasStartPositionTicks

`func (o *ChapterInfo) HasStartPositionTicks() bool`

HasStartPositionTicks returns a boolean if a field has been set.

### GetName

`func (o *ChapterInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChapterInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChapterInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChapterInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ChapterInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ChapterInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetImagePath

`func (o *ChapterInfo) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *ChapterInfo) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *ChapterInfo) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *ChapterInfo) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### SetImagePathNil

`func (o *ChapterInfo) SetImagePathNil(b bool)`

 SetImagePathNil sets the value for ImagePath to be an explicit nil

### UnsetImagePath
`func (o *ChapterInfo) UnsetImagePath()`

UnsetImagePath ensures that no value is present for ImagePath, not even an explicit nil
### GetImageDateModified

`func (o *ChapterInfo) GetImageDateModified() time.Time`

GetImageDateModified returns the ImageDateModified field if non-nil, zero value otherwise.

### GetImageDateModifiedOk

`func (o *ChapterInfo) GetImageDateModifiedOk() (*time.Time, bool)`

GetImageDateModifiedOk returns a tuple with the ImageDateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDateModified

`func (o *ChapterInfo) SetImageDateModified(v time.Time)`

SetImageDateModified sets ImageDateModified field to given value.

### HasImageDateModified

`func (o *ChapterInfo) HasImageDateModified() bool`

HasImageDateModified returns a boolean if a field has been set.

### GetImageTag

`func (o *ChapterInfo) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *ChapterInfo) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *ChapterInfo) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *ChapterInfo) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### SetImageTagNil

`func (o *ChapterInfo) SetImageTagNil(b bool)`

 SetImageTagNil sets the value for ImageTag to be an explicit nil

### UnsetImageTag
`func (o *ChapterInfo) UnsetImageTag()`

UnsetImageTag ensures that no value is present for ImageTag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


