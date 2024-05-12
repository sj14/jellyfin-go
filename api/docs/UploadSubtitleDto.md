# UploadSubtitleDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Language** | **string** | Gets or sets the subtitle language. | 
**Format** | **string** | Gets or sets the subtitle format. | 
**IsForced** | **bool** | Gets or sets a value indicating whether the subtitle is forced. | 
**IsHearingImpaired** | **bool** | Gets or sets a value indicating whether the subtitle is for hearing impaired. | 
**Data** | **string** | Gets or sets the subtitle data. | 

## Methods

### NewUploadSubtitleDto

`func NewUploadSubtitleDto(language string, format string, isForced bool, isHearingImpaired bool, data string, ) *UploadSubtitleDto`

NewUploadSubtitleDto instantiates a new UploadSubtitleDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadSubtitleDtoWithDefaults

`func NewUploadSubtitleDtoWithDefaults() *UploadSubtitleDto`

NewUploadSubtitleDtoWithDefaults instantiates a new UploadSubtitleDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguage

`func (o *UploadSubtitleDto) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UploadSubtitleDto) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UploadSubtitleDto) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetFormat

`func (o *UploadSubtitleDto) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UploadSubtitleDto) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UploadSubtitleDto) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetIsForced

`func (o *UploadSubtitleDto) GetIsForced() bool`

GetIsForced returns the IsForced field if non-nil, zero value otherwise.

### GetIsForcedOk

`func (o *UploadSubtitleDto) GetIsForcedOk() (*bool, bool)`

GetIsForcedOk returns a tuple with the IsForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForced

`func (o *UploadSubtitleDto) SetIsForced(v bool)`

SetIsForced sets IsForced field to given value.


### GetIsHearingImpaired

`func (o *UploadSubtitleDto) GetIsHearingImpaired() bool`

GetIsHearingImpaired returns the IsHearingImpaired field if non-nil, zero value otherwise.

### GetIsHearingImpairedOk

`func (o *UploadSubtitleDto) GetIsHearingImpairedOk() (*bool, bool)`

GetIsHearingImpairedOk returns a tuple with the IsHearingImpaired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHearingImpaired

`func (o *UploadSubtitleDto) SetIsHearingImpaired(v bool)`

SetIsHearingImpaired sets IsHearingImpaired field to given value.


### GetData

`func (o *UploadSubtitleDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UploadSubtitleDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UploadSubtitleDto) SetData(v string)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


