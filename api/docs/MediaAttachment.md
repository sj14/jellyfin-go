# MediaAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codec** | Pointer to **NullableString** | Gets or sets the codec. | [optional] 
**CodecTag** | Pointer to **NullableString** | Gets or sets the codec tag. | [optional] 
**Comment** | Pointer to **NullableString** | Gets or sets the comment. | [optional] 
**Index** | Pointer to **int32** | Gets or sets the index. | [optional] 
**FileName** | Pointer to **NullableString** | Gets or sets the filename. | [optional] 
**MimeType** | Pointer to **NullableString** | Gets or sets the MIME type. | [optional] 
**DeliveryUrl** | Pointer to **NullableString** | Gets or sets the delivery URL. | [optional] 

## Methods

### NewMediaAttachment

`func NewMediaAttachment() *MediaAttachment`

NewMediaAttachment instantiates a new MediaAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaAttachmentWithDefaults

`func NewMediaAttachmentWithDefaults() *MediaAttachment`

NewMediaAttachmentWithDefaults instantiates a new MediaAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodec

`func (o *MediaAttachment) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *MediaAttachment) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *MediaAttachment) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *MediaAttachment) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### SetCodecNil

`func (o *MediaAttachment) SetCodecNil(b bool)`

 SetCodecNil sets the value for Codec to be an explicit nil

### UnsetCodec
`func (o *MediaAttachment) UnsetCodec()`

UnsetCodec ensures that no value is present for Codec, not even an explicit nil
### GetCodecTag

`func (o *MediaAttachment) GetCodecTag() string`

GetCodecTag returns the CodecTag field if non-nil, zero value otherwise.

### GetCodecTagOk

`func (o *MediaAttachment) GetCodecTagOk() (*string, bool)`

GetCodecTagOk returns a tuple with the CodecTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecTag

`func (o *MediaAttachment) SetCodecTag(v string)`

SetCodecTag sets CodecTag field to given value.

### HasCodecTag

`func (o *MediaAttachment) HasCodecTag() bool`

HasCodecTag returns a boolean if a field has been set.

### SetCodecTagNil

`func (o *MediaAttachment) SetCodecTagNil(b bool)`

 SetCodecTagNil sets the value for CodecTag to be an explicit nil

### UnsetCodecTag
`func (o *MediaAttachment) UnsetCodecTag()`

UnsetCodecTag ensures that no value is present for CodecTag, not even an explicit nil
### GetComment

`func (o *MediaAttachment) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MediaAttachment) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MediaAttachment) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MediaAttachment) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *MediaAttachment) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *MediaAttachment) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetIndex

`func (o *MediaAttachment) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MediaAttachment) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MediaAttachment) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *MediaAttachment) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetFileName

`func (o *MediaAttachment) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MediaAttachment) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *MediaAttachment) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *MediaAttachment) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### SetFileNameNil

`func (o *MediaAttachment) SetFileNameNil(b bool)`

 SetFileNameNil sets the value for FileName to be an explicit nil

### UnsetFileName
`func (o *MediaAttachment) UnsetFileName()`

UnsetFileName ensures that no value is present for FileName, not even an explicit nil
### GetMimeType

`func (o *MediaAttachment) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *MediaAttachment) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *MediaAttachment) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *MediaAttachment) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### SetMimeTypeNil

`func (o *MediaAttachment) SetMimeTypeNil(b bool)`

 SetMimeTypeNil sets the value for MimeType to be an explicit nil

### UnsetMimeType
`func (o *MediaAttachment) UnsetMimeType()`

UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil
### GetDeliveryUrl

`func (o *MediaAttachment) GetDeliveryUrl() string`

GetDeliveryUrl returns the DeliveryUrl field if non-nil, zero value otherwise.

### GetDeliveryUrlOk

`func (o *MediaAttachment) GetDeliveryUrlOk() (*string, bool)`

GetDeliveryUrlOk returns a tuple with the DeliveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryUrl

`func (o *MediaAttachment) SetDeliveryUrl(v string)`

SetDeliveryUrl sets DeliveryUrl field to given value.

### HasDeliveryUrl

`func (o *MediaAttachment) HasDeliveryUrl() bool`

HasDeliveryUrl returns a boolean if a field has been set.

### SetDeliveryUrlNil

`func (o *MediaAttachment) SetDeliveryUrlNil(b bool)`

 SetDeliveryUrlNil sets the value for DeliveryUrl to be an explicit nil

### UnsetDeliveryUrl
`func (o *MediaAttachment) UnsetDeliveryUrl()`

UnsetDeliveryUrl ensures that no value is present for DeliveryUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


