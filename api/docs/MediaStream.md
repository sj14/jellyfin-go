# MediaStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codec** | Pointer to **NullableString** | Gets or sets the codec. | [optional] 
**CodecTag** | Pointer to **NullableString** | Gets or sets the codec tag. | [optional] 
**Language** | Pointer to **NullableString** | Gets or sets the language. | [optional] 
**ColorRange** | Pointer to **NullableString** | Gets or sets the color range. | [optional] 
**ColorSpace** | Pointer to **NullableString** | Gets or sets the color space. | [optional] 
**ColorTransfer** | Pointer to **NullableString** | Gets or sets the color transfer. | [optional] 
**ColorPrimaries** | Pointer to **NullableString** | Gets or sets the color primaries. | [optional] 
**DvVersionMajor** | Pointer to **NullableInt32** | Gets or sets the Dolby Vision version major. | [optional] 
**DvVersionMinor** | Pointer to **NullableInt32** | Gets or sets the Dolby Vision version minor. | [optional] 
**DvProfile** | Pointer to **NullableInt32** | Gets or sets the Dolby Vision profile. | [optional] 
**DvLevel** | Pointer to **NullableInt32** | Gets or sets the Dolby Vision level. | [optional] 
**RpuPresentFlag** | Pointer to **NullableInt32** | Gets or sets the Dolby Vision rpu present flag. | [optional] 
**ElPresentFlag** | Pointer to **NullableInt32** | Gets or sets the Dolby Vision el present flag. | [optional] 
**BlPresentFlag** | Pointer to **NullableInt32** | Gets or sets the Dolby Vision bl present flag. | [optional] 
**DvBlSignalCompatibilityId** | Pointer to **NullableInt32** | Gets or sets the Dolby Vision bl signal compatibility id. | [optional] 
**Rotation** | Pointer to **NullableInt32** | Gets or sets the Rotation in degrees. | [optional] 
**Comment** | Pointer to **NullableString** | Gets or sets the comment. | [optional] 
**TimeBase** | Pointer to **NullableString** | Gets or sets the time base. | [optional] 
**CodecTimeBase** | Pointer to **NullableString** | Gets or sets the codec time base. | [optional] 
**Title** | Pointer to **NullableString** | Gets or sets the title. | [optional] 
**Hdr10PlusPresentFlag** | Pointer to **NullableBool** |  | [optional] 
**VideoRange** | Pointer to [**VideoRange**](VideoRange.md) | Gets the video range. | [optional] [readonly] [default to VIDEORANGE_UNKNOWN]
**VideoRangeType** | Pointer to [**VideoRangeType**](VideoRangeType.md) | Gets the video range type. | [optional] [readonly] [default to VIDEORANGETYPE_UNKNOWN]
**VideoDoViTitle** | Pointer to **NullableString** | Gets the video dovi title. | [optional] [readonly] 
**AudioSpatialFormat** | Pointer to [**AudioSpatialFormat**](AudioSpatialFormat.md) | Gets the audio spatial format. | [optional] [readonly] [default to AUDIOSPATIALFORMAT_NONE]
**LocalizedUndefined** | Pointer to **NullableString** |  | [optional] 
**LocalizedDefault** | Pointer to **NullableString** |  | [optional] 
**LocalizedForced** | Pointer to **NullableString** |  | [optional] 
**LocalizedExternal** | Pointer to **NullableString** |  | [optional] 
**LocalizedHearingImpaired** | Pointer to **NullableString** |  | [optional] 
**DisplayTitle** | Pointer to **NullableString** |  | [optional] [readonly] 
**NalLengthSize** | Pointer to **NullableString** |  | [optional] 
**IsInterlaced** | Pointer to **bool** | Gets or sets a value indicating whether this instance is interlaced. | [optional] 
**IsAVC** | Pointer to **NullableBool** |  | [optional] 
**ChannelLayout** | Pointer to **NullableString** | Gets or sets the channel layout. | [optional] 
**BitRate** | Pointer to **NullableInt32** | Gets or sets the bit rate. | [optional] 
**BitDepth** | Pointer to **NullableInt32** | Gets or sets the bit depth. | [optional] 
**RefFrames** | Pointer to **NullableInt32** | Gets or sets the reference frames. | [optional] 
**PacketLength** | Pointer to **NullableInt32** | Gets or sets the length of the packet. | [optional] 
**Channels** | Pointer to **NullableInt32** | Gets or sets the channels. | [optional] 
**SampleRate** | Pointer to **NullableInt32** | Gets or sets the sample rate. | [optional] 
**IsDefault** | Pointer to **bool** | Gets or sets a value indicating whether this instance is default. | [optional] 
**IsForced** | Pointer to **bool** | Gets or sets a value indicating whether this instance is forced. | [optional] 
**IsHearingImpaired** | Pointer to **bool** | Gets or sets a value indicating whether this instance is for the hearing impaired. | [optional] 
**Height** | Pointer to **NullableInt32** | Gets or sets the height. | [optional] 
**Width** | Pointer to **NullableInt32** | Gets or sets the width. | [optional] 
**AverageFrameRate** | Pointer to **NullableFloat32** | Gets or sets the average frame rate. | [optional] 
**RealFrameRate** | Pointer to **NullableFloat32** | Gets or sets the real frame rate. | [optional] 
**ReferenceFrameRate** | Pointer to **NullableFloat32** | Gets the framerate used as reference.  Prefer AverageFrameRate, if that is null or an unrealistic value  then fallback to RealFrameRate. | [optional] [readonly] 
**Profile** | Pointer to **NullableString** | Gets or sets the profile. | [optional] 
**Type** | Pointer to [**MediaStreamType**](MediaStreamType.md) | Gets or sets the type. | [optional] 
**AspectRatio** | Pointer to **NullableString** | Gets or sets the aspect ratio. | [optional] 
**Index** | Pointer to **int32** | Gets or sets the index. | [optional] 
**Score** | Pointer to **NullableInt32** | Gets or sets the score. | [optional] 
**IsExternal** | Pointer to **bool** | Gets or sets a value indicating whether this instance is external. | [optional] 
**DeliveryMethod** | Pointer to [**NullableSubtitleDeliveryMethod**](SubtitleDeliveryMethod.md) | Gets or sets the method. | [optional] 
**DeliveryUrl** | Pointer to **NullableString** | Gets or sets the delivery URL. | [optional] 
**IsExternalUrl** | Pointer to **NullableBool** | Gets or sets a value indicating whether this instance is external URL. | [optional] 
**IsTextSubtitleStream** | Pointer to **bool** |  | [optional] [readonly] 
**SupportsExternalStream** | Pointer to **bool** | Gets or sets a value indicating whether [supports external stream]. | [optional] 
**Path** | Pointer to **NullableString** | Gets or sets the filename. | [optional] 
**PixelFormat** | Pointer to **NullableString** | Gets or sets the pixel format. | [optional] 
**Level** | Pointer to **NullableFloat64** | Gets or sets the level. | [optional] 
**IsAnamorphic** | Pointer to **NullableBool** | Gets or sets whether this instance is anamorphic. | [optional] 

## Methods

### NewMediaStream

`func NewMediaStream() *MediaStream`

NewMediaStream instantiates a new MediaStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaStreamWithDefaults

`func NewMediaStreamWithDefaults() *MediaStream`

NewMediaStreamWithDefaults instantiates a new MediaStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodec

`func (o *MediaStream) GetCodec() string`

GetCodec returns the Codec field if non-nil, zero value otherwise.

### GetCodecOk

`func (o *MediaStream) GetCodecOk() (*string, bool)`

GetCodecOk returns a tuple with the Codec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodec

`func (o *MediaStream) SetCodec(v string)`

SetCodec sets Codec field to given value.

### HasCodec

`func (o *MediaStream) HasCodec() bool`

HasCodec returns a boolean if a field has been set.

### SetCodecNil

`func (o *MediaStream) SetCodecNil(b bool)`

 SetCodecNil sets the value for Codec to be an explicit nil

### UnsetCodec
`func (o *MediaStream) UnsetCodec()`

UnsetCodec ensures that no value is present for Codec, not even an explicit nil
### GetCodecTag

`func (o *MediaStream) GetCodecTag() string`

GetCodecTag returns the CodecTag field if non-nil, zero value otherwise.

### GetCodecTagOk

`func (o *MediaStream) GetCodecTagOk() (*string, bool)`

GetCodecTagOk returns a tuple with the CodecTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecTag

`func (o *MediaStream) SetCodecTag(v string)`

SetCodecTag sets CodecTag field to given value.

### HasCodecTag

`func (o *MediaStream) HasCodecTag() bool`

HasCodecTag returns a boolean if a field has been set.

### SetCodecTagNil

`func (o *MediaStream) SetCodecTagNil(b bool)`

 SetCodecTagNil sets the value for CodecTag to be an explicit nil

### UnsetCodecTag
`func (o *MediaStream) UnsetCodecTag()`

UnsetCodecTag ensures that no value is present for CodecTag, not even an explicit nil
### GetLanguage

`func (o *MediaStream) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MediaStream) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MediaStream) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MediaStream) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *MediaStream) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *MediaStream) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetColorRange

`func (o *MediaStream) GetColorRange() string`

GetColorRange returns the ColorRange field if non-nil, zero value otherwise.

### GetColorRangeOk

`func (o *MediaStream) GetColorRangeOk() (*string, bool)`

GetColorRangeOk returns a tuple with the ColorRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorRange

`func (o *MediaStream) SetColorRange(v string)`

SetColorRange sets ColorRange field to given value.

### HasColorRange

`func (o *MediaStream) HasColorRange() bool`

HasColorRange returns a boolean if a field has been set.

### SetColorRangeNil

`func (o *MediaStream) SetColorRangeNil(b bool)`

 SetColorRangeNil sets the value for ColorRange to be an explicit nil

### UnsetColorRange
`func (o *MediaStream) UnsetColorRange()`

UnsetColorRange ensures that no value is present for ColorRange, not even an explicit nil
### GetColorSpace

`func (o *MediaStream) GetColorSpace() string`

GetColorSpace returns the ColorSpace field if non-nil, zero value otherwise.

### GetColorSpaceOk

`func (o *MediaStream) GetColorSpaceOk() (*string, bool)`

GetColorSpaceOk returns a tuple with the ColorSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorSpace

`func (o *MediaStream) SetColorSpace(v string)`

SetColorSpace sets ColorSpace field to given value.

### HasColorSpace

`func (o *MediaStream) HasColorSpace() bool`

HasColorSpace returns a boolean if a field has been set.

### SetColorSpaceNil

`func (o *MediaStream) SetColorSpaceNil(b bool)`

 SetColorSpaceNil sets the value for ColorSpace to be an explicit nil

### UnsetColorSpace
`func (o *MediaStream) UnsetColorSpace()`

UnsetColorSpace ensures that no value is present for ColorSpace, not even an explicit nil
### GetColorTransfer

`func (o *MediaStream) GetColorTransfer() string`

GetColorTransfer returns the ColorTransfer field if non-nil, zero value otherwise.

### GetColorTransferOk

`func (o *MediaStream) GetColorTransferOk() (*string, bool)`

GetColorTransferOk returns a tuple with the ColorTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorTransfer

`func (o *MediaStream) SetColorTransfer(v string)`

SetColorTransfer sets ColorTransfer field to given value.

### HasColorTransfer

`func (o *MediaStream) HasColorTransfer() bool`

HasColorTransfer returns a boolean if a field has been set.

### SetColorTransferNil

`func (o *MediaStream) SetColorTransferNil(b bool)`

 SetColorTransferNil sets the value for ColorTransfer to be an explicit nil

### UnsetColorTransfer
`func (o *MediaStream) UnsetColorTransfer()`

UnsetColorTransfer ensures that no value is present for ColorTransfer, not even an explicit nil
### GetColorPrimaries

`func (o *MediaStream) GetColorPrimaries() string`

GetColorPrimaries returns the ColorPrimaries field if non-nil, zero value otherwise.

### GetColorPrimariesOk

`func (o *MediaStream) GetColorPrimariesOk() (*string, bool)`

GetColorPrimariesOk returns a tuple with the ColorPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorPrimaries

`func (o *MediaStream) SetColorPrimaries(v string)`

SetColorPrimaries sets ColorPrimaries field to given value.

### HasColorPrimaries

`func (o *MediaStream) HasColorPrimaries() bool`

HasColorPrimaries returns a boolean if a field has been set.

### SetColorPrimariesNil

`func (o *MediaStream) SetColorPrimariesNil(b bool)`

 SetColorPrimariesNil sets the value for ColorPrimaries to be an explicit nil

### UnsetColorPrimaries
`func (o *MediaStream) UnsetColorPrimaries()`

UnsetColorPrimaries ensures that no value is present for ColorPrimaries, not even an explicit nil
### GetDvVersionMajor

`func (o *MediaStream) GetDvVersionMajor() int32`

GetDvVersionMajor returns the DvVersionMajor field if non-nil, zero value otherwise.

### GetDvVersionMajorOk

`func (o *MediaStream) GetDvVersionMajorOk() (*int32, bool)`

GetDvVersionMajorOk returns a tuple with the DvVersionMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvVersionMajor

`func (o *MediaStream) SetDvVersionMajor(v int32)`

SetDvVersionMajor sets DvVersionMajor field to given value.

### HasDvVersionMajor

`func (o *MediaStream) HasDvVersionMajor() bool`

HasDvVersionMajor returns a boolean if a field has been set.

### SetDvVersionMajorNil

`func (o *MediaStream) SetDvVersionMajorNil(b bool)`

 SetDvVersionMajorNil sets the value for DvVersionMajor to be an explicit nil

### UnsetDvVersionMajor
`func (o *MediaStream) UnsetDvVersionMajor()`

UnsetDvVersionMajor ensures that no value is present for DvVersionMajor, not even an explicit nil
### GetDvVersionMinor

`func (o *MediaStream) GetDvVersionMinor() int32`

GetDvVersionMinor returns the DvVersionMinor field if non-nil, zero value otherwise.

### GetDvVersionMinorOk

`func (o *MediaStream) GetDvVersionMinorOk() (*int32, bool)`

GetDvVersionMinorOk returns a tuple with the DvVersionMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvVersionMinor

`func (o *MediaStream) SetDvVersionMinor(v int32)`

SetDvVersionMinor sets DvVersionMinor field to given value.

### HasDvVersionMinor

`func (o *MediaStream) HasDvVersionMinor() bool`

HasDvVersionMinor returns a boolean if a field has been set.

### SetDvVersionMinorNil

`func (o *MediaStream) SetDvVersionMinorNil(b bool)`

 SetDvVersionMinorNil sets the value for DvVersionMinor to be an explicit nil

### UnsetDvVersionMinor
`func (o *MediaStream) UnsetDvVersionMinor()`

UnsetDvVersionMinor ensures that no value is present for DvVersionMinor, not even an explicit nil
### GetDvProfile

`func (o *MediaStream) GetDvProfile() int32`

GetDvProfile returns the DvProfile field if non-nil, zero value otherwise.

### GetDvProfileOk

`func (o *MediaStream) GetDvProfileOk() (*int32, bool)`

GetDvProfileOk returns a tuple with the DvProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvProfile

`func (o *MediaStream) SetDvProfile(v int32)`

SetDvProfile sets DvProfile field to given value.

### HasDvProfile

`func (o *MediaStream) HasDvProfile() bool`

HasDvProfile returns a boolean if a field has been set.

### SetDvProfileNil

`func (o *MediaStream) SetDvProfileNil(b bool)`

 SetDvProfileNil sets the value for DvProfile to be an explicit nil

### UnsetDvProfile
`func (o *MediaStream) UnsetDvProfile()`

UnsetDvProfile ensures that no value is present for DvProfile, not even an explicit nil
### GetDvLevel

`func (o *MediaStream) GetDvLevel() int32`

GetDvLevel returns the DvLevel field if non-nil, zero value otherwise.

### GetDvLevelOk

`func (o *MediaStream) GetDvLevelOk() (*int32, bool)`

GetDvLevelOk returns a tuple with the DvLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvLevel

`func (o *MediaStream) SetDvLevel(v int32)`

SetDvLevel sets DvLevel field to given value.

### HasDvLevel

`func (o *MediaStream) HasDvLevel() bool`

HasDvLevel returns a boolean if a field has been set.

### SetDvLevelNil

`func (o *MediaStream) SetDvLevelNil(b bool)`

 SetDvLevelNil sets the value for DvLevel to be an explicit nil

### UnsetDvLevel
`func (o *MediaStream) UnsetDvLevel()`

UnsetDvLevel ensures that no value is present for DvLevel, not even an explicit nil
### GetRpuPresentFlag

`func (o *MediaStream) GetRpuPresentFlag() int32`

GetRpuPresentFlag returns the RpuPresentFlag field if non-nil, zero value otherwise.

### GetRpuPresentFlagOk

`func (o *MediaStream) GetRpuPresentFlagOk() (*int32, bool)`

GetRpuPresentFlagOk returns a tuple with the RpuPresentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpuPresentFlag

`func (o *MediaStream) SetRpuPresentFlag(v int32)`

SetRpuPresentFlag sets RpuPresentFlag field to given value.

### HasRpuPresentFlag

`func (o *MediaStream) HasRpuPresentFlag() bool`

HasRpuPresentFlag returns a boolean if a field has been set.

### SetRpuPresentFlagNil

`func (o *MediaStream) SetRpuPresentFlagNil(b bool)`

 SetRpuPresentFlagNil sets the value for RpuPresentFlag to be an explicit nil

### UnsetRpuPresentFlag
`func (o *MediaStream) UnsetRpuPresentFlag()`

UnsetRpuPresentFlag ensures that no value is present for RpuPresentFlag, not even an explicit nil
### GetElPresentFlag

`func (o *MediaStream) GetElPresentFlag() int32`

GetElPresentFlag returns the ElPresentFlag field if non-nil, zero value otherwise.

### GetElPresentFlagOk

`func (o *MediaStream) GetElPresentFlagOk() (*int32, bool)`

GetElPresentFlagOk returns a tuple with the ElPresentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElPresentFlag

`func (o *MediaStream) SetElPresentFlag(v int32)`

SetElPresentFlag sets ElPresentFlag field to given value.

### HasElPresentFlag

`func (o *MediaStream) HasElPresentFlag() bool`

HasElPresentFlag returns a boolean if a field has been set.

### SetElPresentFlagNil

`func (o *MediaStream) SetElPresentFlagNil(b bool)`

 SetElPresentFlagNil sets the value for ElPresentFlag to be an explicit nil

### UnsetElPresentFlag
`func (o *MediaStream) UnsetElPresentFlag()`

UnsetElPresentFlag ensures that no value is present for ElPresentFlag, not even an explicit nil
### GetBlPresentFlag

`func (o *MediaStream) GetBlPresentFlag() int32`

GetBlPresentFlag returns the BlPresentFlag field if non-nil, zero value otherwise.

### GetBlPresentFlagOk

`func (o *MediaStream) GetBlPresentFlagOk() (*int32, bool)`

GetBlPresentFlagOk returns a tuple with the BlPresentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlPresentFlag

`func (o *MediaStream) SetBlPresentFlag(v int32)`

SetBlPresentFlag sets BlPresentFlag field to given value.

### HasBlPresentFlag

`func (o *MediaStream) HasBlPresentFlag() bool`

HasBlPresentFlag returns a boolean if a field has been set.

### SetBlPresentFlagNil

`func (o *MediaStream) SetBlPresentFlagNil(b bool)`

 SetBlPresentFlagNil sets the value for BlPresentFlag to be an explicit nil

### UnsetBlPresentFlag
`func (o *MediaStream) UnsetBlPresentFlag()`

UnsetBlPresentFlag ensures that no value is present for BlPresentFlag, not even an explicit nil
### GetDvBlSignalCompatibilityId

`func (o *MediaStream) GetDvBlSignalCompatibilityId() int32`

GetDvBlSignalCompatibilityId returns the DvBlSignalCompatibilityId field if non-nil, zero value otherwise.

### GetDvBlSignalCompatibilityIdOk

`func (o *MediaStream) GetDvBlSignalCompatibilityIdOk() (*int32, bool)`

GetDvBlSignalCompatibilityIdOk returns a tuple with the DvBlSignalCompatibilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvBlSignalCompatibilityId

`func (o *MediaStream) SetDvBlSignalCompatibilityId(v int32)`

SetDvBlSignalCompatibilityId sets DvBlSignalCompatibilityId field to given value.

### HasDvBlSignalCompatibilityId

`func (o *MediaStream) HasDvBlSignalCompatibilityId() bool`

HasDvBlSignalCompatibilityId returns a boolean if a field has been set.

### SetDvBlSignalCompatibilityIdNil

`func (o *MediaStream) SetDvBlSignalCompatibilityIdNil(b bool)`

 SetDvBlSignalCompatibilityIdNil sets the value for DvBlSignalCompatibilityId to be an explicit nil

### UnsetDvBlSignalCompatibilityId
`func (o *MediaStream) UnsetDvBlSignalCompatibilityId()`

UnsetDvBlSignalCompatibilityId ensures that no value is present for DvBlSignalCompatibilityId, not even an explicit nil
### GetRotation

`func (o *MediaStream) GetRotation() int32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *MediaStream) GetRotationOk() (*int32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *MediaStream) SetRotation(v int32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *MediaStream) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### SetRotationNil

`func (o *MediaStream) SetRotationNil(b bool)`

 SetRotationNil sets the value for Rotation to be an explicit nil

### UnsetRotation
`func (o *MediaStream) UnsetRotation()`

UnsetRotation ensures that no value is present for Rotation, not even an explicit nil
### GetComment

`func (o *MediaStream) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MediaStream) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MediaStream) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MediaStream) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *MediaStream) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *MediaStream) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetTimeBase

`func (o *MediaStream) GetTimeBase() string`

GetTimeBase returns the TimeBase field if non-nil, zero value otherwise.

### GetTimeBaseOk

`func (o *MediaStream) GetTimeBaseOk() (*string, bool)`

GetTimeBaseOk returns a tuple with the TimeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBase

`func (o *MediaStream) SetTimeBase(v string)`

SetTimeBase sets TimeBase field to given value.

### HasTimeBase

`func (o *MediaStream) HasTimeBase() bool`

HasTimeBase returns a boolean if a field has been set.

### SetTimeBaseNil

`func (o *MediaStream) SetTimeBaseNil(b bool)`

 SetTimeBaseNil sets the value for TimeBase to be an explicit nil

### UnsetTimeBase
`func (o *MediaStream) UnsetTimeBase()`

UnsetTimeBase ensures that no value is present for TimeBase, not even an explicit nil
### GetCodecTimeBase

`func (o *MediaStream) GetCodecTimeBase() string`

GetCodecTimeBase returns the CodecTimeBase field if non-nil, zero value otherwise.

### GetCodecTimeBaseOk

`func (o *MediaStream) GetCodecTimeBaseOk() (*string, bool)`

GetCodecTimeBaseOk returns a tuple with the CodecTimeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodecTimeBase

`func (o *MediaStream) SetCodecTimeBase(v string)`

SetCodecTimeBase sets CodecTimeBase field to given value.

### HasCodecTimeBase

`func (o *MediaStream) HasCodecTimeBase() bool`

HasCodecTimeBase returns a boolean if a field has been set.

### SetCodecTimeBaseNil

`func (o *MediaStream) SetCodecTimeBaseNil(b bool)`

 SetCodecTimeBaseNil sets the value for CodecTimeBase to be an explicit nil

### UnsetCodecTimeBase
`func (o *MediaStream) UnsetCodecTimeBase()`

UnsetCodecTimeBase ensures that no value is present for CodecTimeBase, not even an explicit nil
### GetTitle

`func (o *MediaStream) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MediaStream) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MediaStream) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MediaStream) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MediaStream) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MediaStream) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetHdr10PlusPresentFlag

`func (o *MediaStream) GetHdr10PlusPresentFlag() bool`

GetHdr10PlusPresentFlag returns the Hdr10PlusPresentFlag field if non-nil, zero value otherwise.

### GetHdr10PlusPresentFlagOk

`func (o *MediaStream) GetHdr10PlusPresentFlagOk() (*bool, bool)`

GetHdr10PlusPresentFlagOk returns a tuple with the Hdr10PlusPresentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdr10PlusPresentFlag

`func (o *MediaStream) SetHdr10PlusPresentFlag(v bool)`

SetHdr10PlusPresentFlag sets Hdr10PlusPresentFlag field to given value.

### HasHdr10PlusPresentFlag

`func (o *MediaStream) HasHdr10PlusPresentFlag() bool`

HasHdr10PlusPresentFlag returns a boolean if a field has been set.

### SetHdr10PlusPresentFlagNil

`func (o *MediaStream) SetHdr10PlusPresentFlagNil(b bool)`

 SetHdr10PlusPresentFlagNil sets the value for Hdr10PlusPresentFlag to be an explicit nil

### UnsetHdr10PlusPresentFlag
`func (o *MediaStream) UnsetHdr10PlusPresentFlag()`

UnsetHdr10PlusPresentFlag ensures that no value is present for Hdr10PlusPresentFlag, not even an explicit nil
### GetVideoRange

`func (o *MediaStream) GetVideoRange() VideoRange`

GetVideoRange returns the VideoRange field if non-nil, zero value otherwise.

### GetVideoRangeOk

`func (o *MediaStream) GetVideoRangeOk() (*VideoRange, bool)`

GetVideoRangeOk returns a tuple with the VideoRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoRange

`func (o *MediaStream) SetVideoRange(v VideoRange)`

SetVideoRange sets VideoRange field to given value.

### HasVideoRange

`func (o *MediaStream) HasVideoRange() bool`

HasVideoRange returns a boolean if a field has been set.

### GetVideoRangeType

`func (o *MediaStream) GetVideoRangeType() VideoRangeType`

GetVideoRangeType returns the VideoRangeType field if non-nil, zero value otherwise.

### GetVideoRangeTypeOk

`func (o *MediaStream) GetVideoRangeTypeOk() (*VideoRangeType, bool)`

GetVideoRangeTypeOk returns a tuple with the VideoRangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoRangeType

`func (o *MediaStream) SetVideoRangeType(v VideoRangeType)`

SetVideoRangeType sets VideoRangeType field to given value.

### HasVideoRangeType

`func (o *MediaStream) HasVideoRangeType() bool`

HasVideoRangeType returns a boolean if a field has been set.

### GetVideoDoViTitle

`func (o *MediaStream) GetVideoDoViTitle() string`

GetVideoDoViTitle returns the VideoDoViTitle field if non-nil, zero value otherwise.

### GetVideoDoViTitleOk

`func (o *MediaStream) GetVideoDoViTitleOk() (*string, bool)`

GetVideoDoViTitleOk returns a tuple with the VideoDoViTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoDoViTitle

`func (o *MediaStream) SetVideoDoViTitle(v string)`

SetVideoDoViTitle sets VideoDoViTitle field to given value.

### HasVideoDoViTitle

`func (o *MediaStream) HasVideoDoViTitle() bool`

HasVideoDoViTitle returns a boolean if a field has been set.

### SetVideoDoViTitleNil

`func (o *MediaStream) SetVideoDoViTitleNil(b bool)`

 SetVideoDoViTitleNil sets the value for VideoDoViTitle to be an explicit nil

### UnsetVideoDoViTitle
`func (o *MediaStream) UnsetVideoDoViTitle()`

UnsetVideoDoViTitle ensures that no value is present for VideoDoViTitle, not even an explicit nil
### GetAudioSpatialFormat

`func (o *MediaStream) GetAudioSpatialFormat() AudioSpatialFormat`

GetAudioSpatialFormat returns the AudioSpatialFormat field if non-nil, zero value otherwise.

### GetAudioSpatialFormatOk

`func (o *MediaStream) GetAudioSpatialFormatOk() (*AudioSpatialFormat, bool)`

GetAudioSpatialFormatOk returns a tuple with the AudioSpatialFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioSpatialFormat

`func (o *MediaStream) SetAudioSpatialFormat(v AudioSpatialFormat)`

SetAudioSpatialFormat sets AudioSpatialFormat field to given value.

### HasAudioSpatialFormat

`func (o *MediaStream) HasAudioSpatialFormat() bool`

HasAudioSpatialFormat returns a boolean if a field has been set.

### GetLocalizedUndefined

`func (o *MediaStream) GetLocalizedUndefined() string`

GetLocalizedUndefined returns the LocalizedUndefined field if non-nil, zero value otherwise.

### GetLocalizedUndefinedOk

`func (o *MediaStream) GetLocalizedUndefinedOk() (*string, bool)`

GetLocalizedUndefinedOk returns a tuple with the LocalizedUndefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedUndefined

`func (o *MediaStream) SetLocalizedUndefined(v string)`

SetLocalizedUndefined sets LocalizedUndefined field to given value.

### HasLocalizedUndefined

`func (o *MediaStream) HasLocalizedUndefined() bool`

HasLocalizedUndefined returns a boolean if a field has been set.

### SetLocalizedUndefinedNil

`func (o *MediaStream) SetLocalizedUndefinedNil(b bool)`

 SetLocalizedUndefinedNil sets the value for LocalizedUndefined to be an explicit nil

### UnsetLocalizedUndefined
`func (o *MediaStream) UnsetLocalizedUndefined()`

UnsetLocalizedUndefined ensures that no value is present for LocalizedUndefined, not even an explicit nil
### GetLocalizedDefault

`func (o *MediaStream) GetLocalizedDefault() string`

GetLocalizedDefault returns the LocalizedDefault field if non-nil, zero value otherwise.

### GetLocalizedDefaultOk

`func (o *MediaStream) GetLocalizedDefaultOk() (*string, bool)`

GetLocalizedDefaultOk returns a tuple with the LocalizedDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedDefault

`func (o *MediaStream) SetLocalizedDefault(v string)`

SetLocalizedDefault sets LocalizedDefault field to given value.

### HasLocalizedDefault

`func (o *MediaStream) HasLocalizedDefault() bool`

HasLocalizedDefault returns a boolean if a field has been set.

### SetLocalizedDefaultNil

`func (o *MediaStream) SetLocalizedDefaultNil(b bool)`

 SetLocalizedDefaultNil sets the value for LocalizedDefault to be an explicit nil

### UnsetLocalizedDefault
`func (o *MediaStream) UnsetLocalizedDefault()`

UnsetLocalizedDefault ensures that no value is present for LocalizedDefault, not even an explicit nil
### GetLocalizedForced

`func (o *MediaStream) GetLocalizedForced() string`

GetLocalizedForced returns the LocalizedForced field if non-nil, zero value otherwise.

### GetLocalizedForcedOk

`func (o *MediaStream) GetLocalizedForcedOk() (*string, bool)`

GetLocalizedForcedOk returns a tuple with the LocalizedForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedForced

`func (o *MediaStream) SetLocalizedForced(v string)`

SetLocalizedForced sets LocalizedForced field to given value.

### HasLocalizedForced

`func (o *MediaStream) HasLocalizedForced() bool`

HasLocalizedForced returns a boolean if a field has been set.

### SetLocalizedForcedNil

`func (o *MediaStream) SetLocalizedForcedNil(b bool)`

 SetLocalizedForcedNil sets the value for LocalizedForced to be an explicit nil

### UnsetLocalizedForced
`func (o *MediaStream) UnsetLocalizedForced()`

UnsetLocalizedForced ensures that no value is present for LocalizedForced, not even an explicit nil
### GetLocalizedExternal

`func (o *MediaStream) GetLocalizedExternal() string`

GetLocalizedExternal returns the LocalizedExternal field if non-nil, zero value otherwise.

### GetLocalizedExternalOk

`func (o *MediaStream) GetLocalizedExternalOk() (*string, bool)`

GetLocalizedExternalOk returns a tuple with the LocalizedExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedExternal

`func (o *MediaStream) SetLocalizedExternal(v string)`

SetLocalizedExternal sets LocalizedExternal field to given value.

### HasLocalizedExternal

`func (o *MediaStream) HasLocalizedExternal() bool`

HasLocalizedExternal returns a boolean if a field has been set.

### SetLocalizedExternalNil

`func (o *MediaStream) SetLocalizedExternalNil(b bool)`

 SetLocalizedExternalNil sets the value for LocalizedExternal to be an explicit nil

### UnsetLocalizedExternal
`func (o *MediaStream) UnsetLocalizedExternal()`

UnsetLocalizedExternal ensures that no value is present for LocalizedExternal, not even an explicit nil
### GetLocalizedHearingImpaired

`func (o *MediaStream) GetLocalizedHearingImpaired() string`

GetLocalizedHearingImpaired returns the LocalizedHearingImpaired field if non-nil, zero value otherwise.

### GetLocalizedHearingImpairedOk

`func (o *MediaStream) GetLocalizedHearingImpairedOk() (*string, bool)`

GetLocalizedHearingImpairedOk returns a tuple with the LocalizedHearingImpaired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedHearingImpaired

`func (o *MediaStream) SetLocalizedHearingImpaired(v string)`

SetLocalizedHearingImpaired sets LocalizedHearingImpaired field to given value.

### HasLocalizedHearingImpaired

`func (o *MediaStream) HasLocalizedHearingImpaired() bool`

HasLocalizedHearingImpaired returns a boolean if a field has been set.

### SetLocalizedHearingImpairedNil

`func (o *MediaStream) SetLocalizedHearingImpairedNil(b bool)`

 SetLocalizedHearingImpairedNil sets the value for LocalizedHearingImpaired to be an explicit nil

### UnsetLocalizedHearingImpaired
`func (o *MediaStream) UnsetLocalizedHearingImpaired()`

UnsetLocalizedHearingImpaired ensures that no value is present for LocalizedHearingImpaired, not even an explicit nil
### GetDisplayTitle

`func (o *MediaStream) GetDisplayTitle() string`

GetDisplayTitle returns the DisplayTitle field if non-nil, zero value otherwise.

### GetDisplayTitleOk

`func (o *MediaStream) GetDisplayTitleOk() (*string, bool)`

GetDisplayTitleOk returns a tuple with the DisplayTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTitle

`func (o *MediaStream) SetDisplayTitle(v string)`

SetDisplayTitle sets DisplayTitle field to given value.

### HasDisplayTitle

`func (o *MediaStream) HasDisplayTitle() bool`

HasDisplayTitle returns a boolean if a field has been set.

### SetDisplayTitleNil

`func (o *MediaStream) SetDisplayTitleNil(b bool)`

 SetDisplayTitleNil sets the value for DisplayTitle to be an explicit nil

### UnsetDisplayTitle
`func (o *MediaStream) UnsetDisplayTitle()`

UnsetDisplayTitle ensures that no value is present for DisplayTitle, not even an explicit nil
### GetNalLengthSize

`func (o *MediaStream) GetNalLengthSize() string`

GetNalLengthSize returns the NalLengthSize field if non-nil, zero value otherwise.

### GetNalLengthSizeOk

`func (o *MediaStream) GetNalLengthSizeOk() (*string, bool)`

GetNalLengthSizeOk returns a tuple with the NalLengthSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNalLengthSize

`func (o *MediaStream) SetNalLengthSize(v string)`

SetNalLengthSize sets NalLengthSize field to given value.

### HasNalLengthSize

`func (o *MediaStream) HasNalLengthSize() bool`

HasNalLengthSize returns a boolean if a field has been set.

### SetNalLengthSizeNil

`func (o *MediaStream) SetNalLengthSizeNil(b bool)`

 SetNalLengthSizeNil sets the value for NalLengthSize to be an explicit nil

### UnsetNalLengthSize
`func (o *MediaStream) UnsetNalLengthSize()`

UnsetNalLengthSize ensures that no value is present for NalLengthSize, not even an explicit nil
### GetIsInterlaced

`func (o *MediaStream) GetIsInterlaced() bool`

GetIsInterlaced returns the IsInterlaced field if non-nil, zero value otherwise.

### GetIsInterlacedOk

`func (o *MediaStream) GetIsInterlacedOk() (*bool, bool)`

GetIsInterlacedOk returns a tuple with the IsInterlaced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInterlaced

`func (o *MediaStream) SetIsInterlaced(v bool)`

SetIsInterlaced sets IsInterlaced field to given value.

### HasIsInterlaced

`func (o *MediaStream) HasIsInterlaced() bool`

HasIsInterlaced returns a boolean if a field has been set.

### GetIsAVC

`func (o *MediaStream) GetIsAVC() bool`

GetIsAVC returns the IsAVC field if non-nil, zero value otherwise.

### GetIsAVCOk

`func (o *MediaStream) GetIsAVCOk() (*bool, bool)`

GetIsAVCOk returns a tuple with the IsAVC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAVC

`func (o *MediaStream) SetIsAVC(v bool)`

SetIsAVC sets IsAVC field to given value.

### HasIsAVC

`func (o *MediaStream) HasIsAVC() bool`

HasIsAVC returns a boolean if a field has been set.

### SetIsAVCNil

`func (o *MediaStream) SetIsAVCNil(b bool)`

 SetIsAVCNil sets the value for IsAVC to be an explicit nil

### UnsetIsAVC
`func (o *MediaStream) UnsetIsAVC()`

UnsetIsAVC ensures that no value is present for IsAVC, not even an explicit nil
### GetChannelLayout

`func (o *MediaStream) GetChannelLayout() string`

GetChannelLayout returns the ChannelLayout field if non-nil, zero value otherwise.

### GetChannelLayoutOk

`func (o *MediaStream) GetChannelLayoutOk() (*string, bool)`

GetChannelLayoutOk returns a tuple with the ChannelLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLayout

`func (o *MediaStream) SetChannelLayout(v string)`

SetChannelLayout sets ChannelLayout field to given value.

### HasChannelLayout

`func (o *MediaStream) HasChannelLayout() bool`

HasChannelLayout returns a boolean if a field has been set.

### SetChannelLayoutNil

`func (o *MediaStream) SetChannelLayoutNil(b bool)`

 SetChannelLayoutNil sets the value for ChannelLayout to be an explicit nil

### UnsetChannelLayout
`func (o *MediaStream) UnsetChannelLayout()`

UnsetChannelLayout ensures that no value is present for ChannelLayout, not even an explicit nil
### GetBitRate

`func (o *MediaStream) GetBitRate() int32`

GetBitRate returns the BitRate field if non-nil, zero value otherwise.

### GetBitRateOk

`func (o *MediaStream) GetBitRateOk() (*int32, bool)`

GetBitRateOk returns a tuple with the BitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitRate

`func (o *MediaStream) SetBitRate(v int32)`

SetBitRate sets BitRate field to given value.

### HasBitRate

`func (o *MediaStream) HasBitRate() bool`

HasBitRate returns a boolean if a field has been set.

### SetBitRateNil

`func (o *MediaStream) SetBitRateNil(b bool)`

 SetBitRateNil sets the value for BitRate to be an explicit nil

### UnsetBitRate
`func (o *MediaStream) UnsetBitRate()`

UnsetBitRate ensures that no value is present for BitRate, not even an explicit nil
### GetBitDepth

`func (o *MediaStream) GetBitDepth() int32`

GetBitDepth returns the BitDepth field if non-nil, zero value otherwise.

### GetBitDepthOk

`func (o *MediaStream) GetBitDepthOk() (*int32, bool)`

GetBitDepthOk returns a tuple with the BitDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitDepth

`func (o *MediaStream) SetBitDepth(v int32)`

SetBitDepth sets BitDepth field to given value.

### HasBitDepth

`func (o *MediaStream) HasBitDepth() bool`

HasBitDepth returns a boolean if a field has been set.

### SetBitDepthNil

`func (o *MediaStream) SetBitDepthNil(b bool)`

 SetBitDepthNil sets the value for BitDepth to be an explicit nil

### UnsetBitDepth
`func (o *MediaStream) UnsetBitDepth()`

UnsetBitDepth ensures that no value is present for BitDepth, not even an explicit nil
### GetRefFrames

`func (o *MediaStream) GetRefFrames() int32`

GetRefFrames returns the RefFrames field if non-nil, zero value otherwise.

### GetRefFramesOk

`func (o *MediaStream) GetRefFramesOk() (*int32, bool)`

GetRefFramesOk returns a tuple with the RefFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefFrames

`func (o *MediaStream) SetRefFrames(v int32)`

SetRefFrames sets RefFrames field to given value.

### HasRefFrames

`func (o *MediaStream) HasRefFrames() bool`

HasRefFrames returns a boolean if a field has been set.

### SetRefFramesNil

`func (o *MediaStream) SetRefFramesNil(b bool)`

 SetRefFramesNil sets the value for RefFrames to be an explicit nil

### UnsetRefFrames
`func (o *MediaStream) UnsetRefFrames()`

UnsetRefFrames ensures that no value is present for RefFrames, not even an explicit nil
### GetPacketLength

`func (o *MediaStream) GetPacketLength() int32`

GetPacketLength returns the PacketLength field if non-nil, zero value otherwise.

### GetPacketLengthOk

`func (o *MediaStream) GetPacketLengthOk() (*int32, bool)`

GetPacketLengthOk returns a tuple with the PacketLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketLength

`func (o *MediaStream) SetPacketLength(v int32)`

SetPacketLength sets PacketLength field to given value.

### HasPacketLength

`func (o *MediaStream) HasPacketLength() bool`

HasPacketLength returns a boolean if a field has been set.

### SetPacketLengthNil

`func (o *MediaStream) SetPacketLengthNil(b bool)`

 SetPacketLengthNil sets the value for PacketLength to be an explicit nil

### UnsetPacketLength
`func (o *MediaStream) UnsetPacketLength()`

UnsetPacketLength ensures that no value is present for PacketLength, not even an explicit nil
### GetChannels

`func (o *MediaStream) GetChannels() int32`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *MediaStream) GetChannelsOk() (*int32, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *MediaStream) SetChannels(v int32)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *MediaStream) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *MediaStream) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *MediaStream) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil
### GetSampleRate

`func (o *MediaStream) GetSampleRate() int32`

GetSampleRate returns the SampleRate field if non-nil, zero value otherwise.

### GetSampleRateOk

`func (o *MediaStream) GetSampleRateOk() (*int32, bool)`

GetSampleRateOk returns a tuple with the SampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRate

`func (o *MediaStream) SetSampleRate(v int32)`

SetSampleRate sets SampleRate field to given value.

### HasSampleRate

`func (o *MediaStream) HasSampleRate() bool`

HasSampleRate returns a boolean if a field has been set.

### SetSampleRateNil

`func (o *MediaStream) SetSampleRateNil(b bool)`

 SetSampleRateNil sets the value for SampleRate to be an explicit nil

### UnsetSampleRate
`func (o *MediaStream) UnsetSampleRate()`

UnsetSampleRate ensures that no value is present for SampleRate, not even an explicit nil
### GetIsDefault

`func (o *MediaStream) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MediaStream) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MediaStream) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *MediaStream) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsForced

`func (o *MediaStream) GetIsForced() bool`

GetIsForced returns the IsForced field if non-nil, zero value otherwise.

### GetIsForcedOk

`func (o *MediaStream) GetIsForcedOk() (*bool, bool)`

GetIsForcedOk returns a tuple with the IsForced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForced

`func (o *MediaStream) SetIsForced(v bool)`

SetIsForced sets IsForced field to given value.

### HasIsForced

`func (o *MediaStream) HasIsForced() bool`

HasIsForced returns a boolean if a field has been set.

### GetIsHearingImpaired

`func (o *MediaStream) GetIsHearingImpaired() bool`

GetIsHearingImpaired returns the IsHearingImpaired field if non-nil, zero value otherwise.

### GetIsHearingImpairedOk

`func (o *MediaStream) GetIsHearingImpairedOk() (*bool, bool)`

GetIsHearingImpairedOk returns a tuple with the IsHearingImpaired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHearingImpaired

`func (o *MediaStream) SetIsHearingImpaired(v bool)`

SetIsHearingImpaired sets IsHearingImpaired field to given value.

### HasIsHearingImpaired

`func (o *MediaStream) HasIsHearingImpaired() bool`

HasIsHearingImpaired returns a boolean if a field has been set.

### GetHeight

`func (o *MediaStream) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *MediaStream) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *MediaStream) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *MediaStream) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *MediaStream) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *MediaStream) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetWidth

`func (o *MediaStream) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *MediaStream) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *MediaStream) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *MediaStream) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *MediaStream) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *MediaStream) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetAverageFrameRate

`func (o *MediaStream) GetAverageFrameRate() float32`

GetAverageFrameRate returns the AverageFrameRate field if non-nil, zero value otherwise.

### GetAverageFrameRateOk

`func (o *MediaStream) GetAverageFrameRateOk() (*float32, bool)`

GetAverageFrameRateOk returns a tuple with the AverageFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageFrameRate

`func (o *MediaStream) SetAverageFrameRate(v float32)`

SetAverageFrameRate sets AverageFrameRate field to given value.

### HasAverageFrameRate

`func (o *MediaStream) HasAverageFrameRate() bool`

HasAverageFrameRate returns a boolean if a field has been set.

### SetAverageFrameRateNil

`func (o *MediaStream) SetAverageFrameRateNil(b bool)`

 SetAverageFrameRateNil sets the value for AverageFrameRate to be an explicit nil

### UnsetAverageFrameRate
`func (o *MediaStream) UnsetAverageFrameRate()`

UnsetAverageFrameRate ensures that no value is present for AverageFrameRate, not even an explicit nil
### GetRealFrameRate

`func (o *MediaStream) GetRealFrameRate() float32`

GetRealFrameRate returns the RealFrameRate field if non-nil, zero value otherwise.

### GetRealFrameRateOk

`func (o *MediaStream) GetRealFrameRateOk() (*float32, bool)`

GetRealFrameRateOk returns a tuple with the RealFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealFrameRate

`func (o *MediaStream) SetRealFrameRate(v float32)`

SetRealFrameRate sets RealFrameRate field to given value.

### HasRealFrameRate

`func (o *MediaStream) HasRealFrameRate() bool`

HasRealFrameRate returns a boolean if a field has been set.

### SetRealFrameRateNil

`func (o *MediaStream) SetRealFrameRateNil(b bool)`

 SetRealFrameRateNil sets the value for RealFrameRate to be an explicit nil

### UnsetRealFrameRate
`func (o *MediaStream) UnsetRealFrameRate()`

UnsetRealFrameRate ensures that no value is present for RealFrameRate, not even an explicit nil
### GetReferenceFrameRate

`func (o *MediaStream) GetReferenceFrameRate() float32`

GetReferenceFrameRate returns the ReferenceFrameRate field if non-nil, zero value otherwise.

### GetReferenceFrameRateOk

`func (o *MediaStream) GetReferenceFrameRateOk() (*float32, bool)`

GetReferenceFrameRateOk returns a tuple with the ReferenceFrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceFrameRate

`func (o *MediaStream) SetReferenceFrameRate(v float32)`

SetReferenceFrameRate sets ReferenceFrameRate field to given value.

### HasReferenceFrameRate

`func (o *MediaStream) HasReferenceFrameRate() bool`

HasReferenceFrameRate returns a boolean if a field has been set.

### SetReferenceFrameRateNil

`func (o *MediaStream) SetReferenceFrameRateNil(b bool)`

 SetReferenceFrameRateNil sets the value for ReferenceFrameRate to be an explicit nil

### UnsetReferenceFrameRate
`func (o *MediaStream) UnsetReferenceFrameRate()`

UnsetReferenceFrameRate ensures that no value is present for ReferenceFrameRate, not even an explicit nil
### GetProfile

`func (o *MediaStream) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *MediaStream) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *MediaStream) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *MediaStream) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *MediaStream) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *MediaStream) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetType

`func (o *MediaStream) GetType() MediaStreamType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MediaStream) GetTypeOk() (*MediaStreamType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MediaStream) SetType(v MediaStreamType)`

SetType sets Type field to given value.

### HasType

`func (o *MediaStream) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAspectRatio

`func (o *MediaStream) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *MediaStream) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *MediaStream) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *MediaStream) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### SetAspectRatioNil

`func (o *MediaStream) SetAspectRatioNil(b bool)`

 SetAspectRatioNil sets the value for AspectRatio to be an explicit nil

### UnsetAspectRatio
`func (o *MediaStream) UnsetAspectRatio()`

UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
### GetIndex

`func (o *MediaStream) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *MediaStream) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *MediaStream) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *MediaStream) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetScore

`func (o *MediaStream) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *MediaStream) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *MediaStream) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *MediaStream) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *MediaStream) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *MediaStream) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetIsExternal

`func (o *MediaStream) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *MediaStream) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *MediaStream) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *MediaStream) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetDeliveryMethod

`func (o *MediaStream) GetDeliveryMethod() SubtitleDeliveryMethod`

GetDeliveryMethod returns the DeliveryMethod field if non-nil, zero value otherwise.

### GetDeliveryMethodOk

`func (o *MediaStream) GetDeliveryMethodOk() (*SubtitleDeliveryMethod, bool)`

GetDeliveryMethodOk returns a tuple with the DeliveryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethod

`func (o *MediaStream) SetDeliveryMethod(v SubtitleDeliveryMethod)`

SetDeliveryMethod sets DeliveryMethod field to given value.

### HasDeliveryMethod

`func (o *MediaStream) HasDeliveryMethod() bool`

HasDeliveryMethod returns a boolean if a field has been set.

### SetDeliveryMethodNil

`func (o *MediaStream) SetDeliveryMethodNil(b bool)`

 SetDeliveryMethodNil sets the value for DeliveryMethod to be an explicit nil

### UnsetDeliveryMethod
`func (o *MediaStream) UnsetDeliveryMethod()`

UnsetDeliveryMethod ensures that no value is present for DeliveryMethod, not even an explicit nil
### GetDeliveryUrl

`func (o *MediaStream) GetDeliveryUrl() string`

GetDeliveryUrl returns the DeliveryUrl field if non-nil, zero value otherwise.

### GetDeliveryUrlOk

`func (o *MediaStream) GetDeliveryUrlOk() (*string, bool)`

GetDeliveryUrlOk returns a tuple with the DeliveryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryUrl

`func (o *MediaStream) SetDeliveryUrl(v string)`

SetDeliveryUrl sets DeliveryUrl field to given value.

### HasDeliveryUrl

`func (o *MediaStream) HasDeliveryUrl() bool`

HasDeliveryUrl returns a boolean if a field has been set.

### SetDeliveryUrlNil

`func (o *MediaStream) SetDeliveryUrlNil(b bool)`

 SetDeliveryUrlNil sets the value for DeliveryUrl to be an explicit nil

### UnsetDeliveryUrl
`func (o *MediaStream) UnsetDeliveryUrl()`

UnsetDeliveryUrl ensures that no value is present for DeliveryUrl, not even an explicit nil
### GetIsExternalUrl

`func (o *MediaStream) GetIsExternalUrl() bool`

GetIsExternalUrl returns the IsExternalUrl field if non-nil, zero value otherwise.

### GetIsExternalUrlOk

`func (o *MediaStream) GetIsExternalUrlOk() (*bool, bool)`

GetIsExternalUrlOk returns a tuple with the IsExternalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternalUrl

`func (o *MediaStream) SetIsExternalUrl(v bool)`

SetIsExternalUrl sets IsExternalUrl field to given value.

### HasIsExternalUrl

`func (o *MediaStream) HasIsExternalUrl() bool`

HasIsExternalUrl returns a boolean if a field has been set.

### SetIsExternalUrlNil

`func (o *MediaStream) SetIsExternalUrlNil(b bool)`

 SetIsExternalUrlNil sets the value for IsExternalUrl to be an explicit nil

### UnsetIsExternalUrl
`func (o *MediaStream) UnsetIsExternalUrl()`

UnsetIsExternalUrl ensures that no value is present for IsExternalUrl, not even an explicit nil
### GetIsTextSubtitleStream

`func (o *MediaStream) GetIsTextSubtitleStream() bool`

GetIsTextSubtitleStream returns the IsTextSubtitleStream field if non-nil, zero value otherwise.

### GetIsTextSubtitleStreamOk

`func (o *MediaStream) GetIsTextSubtitleStreamOk() (*bool, bool)`

GetIsTextSubtitleStreamOk returns a tuple with the IsTextSubtitleStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTextSubtitleStream

`func (o *MediaStream) SetIsTextSubtitleStream(v bool)`

SetIsTextSubtitleStream sets IsTextSubtitleStream field to given value.

### HasIsTextSubtitleStream

`func (o *MediaStream) HasIsTextSubtitleStream() bool`

HasIsTextSubtitleStream returns a boolean if a field has been set.

### GetSupportsExternalStream

`func (o *MediaStream) GetSupportsExternalStream() bool`

GetSupportsExternalStream returns the SupportsExternalStream field if non-nil, zero value otherwise.

### GetSupportsExternalStreamOk

`func (o *MediaStream) GetSupportsExternalStreamOk() (*bool, bool)`

GetSupportsExternalStreamOk returns a tuple with the SupportsExternalStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsExternalStream

`func (o *MediaStream) SetSupportsExternalStream(v bool)`

SetSupportsExternalStream sets SupportsExternalStream field to given value.

### HasSupportsExternalStream

`func (o *MediaStream) HasSupportsExternalStream() bool`

HasSupportsExternalStream returns a boolean if a field has been set.

### GetPath

`func (o *MediaStream) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MediaStream) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MediaStream) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MediaStream) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *MediaStream) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *MediaStream) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetPixelFormat

`func (o *MediaStream) GetPixelFormat() string`

GetPixelFormat returns the PixelFormat field if non-nil, zero value otherwise.

### GetPixelFormatOk

`func (o *MediaStream) GetPixelFormatOk() (*string, bool)`

GetPixelFormatOk returns a tuple with the PixelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixelFormat

`func (o *MediaStream) SetPixelFormat(v string)`

SetPixelFormat sets PixelFormat field to given value.

### HasPixelFormat

`func (o *MediaStream) HasPixelFormat() bool`

HasPixelFormat returns a boolean if a field has been set.

### SetPixelFormatNil

`func (o *MediaStream) SetPixelFormatNil(b bool)`

 SetPixelFormatNil sets the value for PixelFormat to be an explicit nil

### UnsetPixelFormat
`func (o *MediaStream) UnsetPixelFormat()`

UnsetPixelFormat ensures that no value is present for PixelFormat, not even an explicit nil
### GetLevel

`func (o *MediaStream) GetLevel() float64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *MediaStream) GetLevelOk() (*float64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *MediaStream) SetLevel(v float64)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *MediaStream) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *MediaStream) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *MediaStream) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetIsAnamorphic

`func (o *MediaStream) GetIsAnamorphic() bool`

GetIsAnamorphic returns the IsAnamorphic field if non-nil, zero value otherwise.

### GetIsAnamorphicOk

`func (o *MediaStream) GetIsAnamorphicOk() (*bool, bool)`

GetIsAnamorphicOk returns a tuple with the IsAnamorphic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnamorphic

`func (o *MediaStream) SetIsAnamorphic(v bool)`

SetIsAnamorphic sets IsAnamorphic field to given value.

### HasIsAnamorphic

`func (o *MediaStream) HasIsAnamorphic() bool`

HasIsAnamorphic returns a boolean if a field has been set.

### SetIsAnamorphicNil

`func (o *MediaStream) SetIsAnamorphicNil(b bool)`

 SetIsAnamorphicNil sets the value for IsAnamorphic to be an explicit nil

### UnsetIsAnamorphic
`func (o *MediaStream) UnsetIsAnamorphic()`

UnsetIsAnamorphic ensures that no value is present for IsAnamorphic, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


