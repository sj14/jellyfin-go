# EncodingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncodingThreadCount** | Pointer to **int32** | Gets or sets the thread count used for encoding. | [optional] 
**TranscodingTempPath** | Pointer to **NullableString** | Gets or sets the temporary transcoding path. | [optional] 
**FallbackFontPath** | Pointer to **NullableString** | Gets or sets the path to the fallback font. | [optional] 
**EnableFallbackFont** | Pointer to **bool** | Gets or sets a value indicating whether to use the fallback font. | [optional] 
**EnableAudioVbr** | Pointer to **bool** | Gets or sets a value indicating whether audio VBR is enabled. | [optional] 
**DownMixAudioBoost** | Pointer to **float64** | Gets or sets the audio boost applied when downmixing audio. | [optional] 
**DownMixStereoAlgorithm** | Pointer to [**DownMixStereoAlgorithms**](DownMixStereoAlgorithms.md) | Gets or sets the algorithm used for downmixing audio to stereo. | [optional] 
**MaxMuxingQueueSize** | Pointer to **int32** | Gets or sets the maximum size of the muxing queue. | [optional] 
**EnableThrottling** | Pointer to **bool** | Gets or sets a value indicating whether throttling is enabled. | [optional] 
**ThrottleDelaySeconds** | Pointer to **int32** | Gets or sets the delay after which throttling happens. | [optional] 
**EnableSegmentDeletion** | Pointer to **bool** | Gets or sets a value indicating whether segment deletion is enabled. | [optional] 
**SegmentKeepSeconds** | Pointer to **int32** | Gets or sets seconds for which segments should be kept before being deleted. | [optional] 
**HardwareAccelerationType** | Pointer to **NullableString** | Gets or sets the hardware acceleration type. | [optional] 
**EncoderAppPath** | Pointer to **NullableString** | Gets or sets the FFmpeg path as set by the user via the UI. | [optional] 
**EncoderAppPathDisplay** | Pointer to **NullableString** | Gets or sets the current FFmpeg path being used by the system and displayed on the transcode page. | [optional] 
**VaapiDevice** | Pointer to **NullableString** | Gets or sets the VA-API device. | [optional] 
**EnableTonemapping** | Pointer to **bool** | Gets or sets a value indicating whether tonemapping is enabled. | [optional] 
**EnableVppTonemapping** | Pointer to **bool** | Gets or sets a value indicating whether VPP tonemapping is enabled. | [optional] 
**EnableVideoToolboxTonemapping** | Pointer to **bool** | Gets or sets a value indicating whether videotoolbox tonemapping is enabled. | [optional] 
**TonemappingAlgorithm** | Pointer to **NullableString** | Gets or sets the tone-mapping algorithm. | [optional] 
**TonemappingMode** | Pointer to **NullableString** | Gets or sets the tone-mapping mode. | [optional] 
**TonemappingRange** | Pointer to **NullableString** | Gets or sets the tone-mapping range. | [optional] 
**TonemappingDesat** | Pointer to **float64** | Gets or sets the tone-mapping desaturation. | [optional] 
**TonemappingPeak** | Pointer to **float64** | Gets or sets the tone-mapping peak. | [optional] 
**TonemappingParam** | Pointer to **float64** | Gets or sets the tone-mapping parameters. | [optional] 
**VppTonemappingBrightness** | Pointer to **float64** | Gets or sets the VPP tone-mapping brightness. | [optional] 
**VppTonemappingContrast** | Pointer to **float64** | Gets or sets the VPP tone-mapping contrast. | [optional] 
**H264Crf** | Pointer to **int32** | Gets or sets the H264 CRF. | [optional] 
**H265Crf** | Pointer to **int32** | Gets or sets the H265 CRF. | [optional] 
**EncoderPreset** | Pointer to **NullableString** | Gets or sets the encoder preset. | [optional] 
**DeinterlaceDoubleRate** | Pointer to **bool** | Gets or sets a value indicating whether the framerate is doubled when deinterlacing. | [optional] 
**DeinterlaceMethod** | Pointer to **NullableString** | Gets or sets the deinterlace method. | [optional] 
**EnableDecodingColorDepth10Hevc** | Pointer to **bool** | Gets or sets a value indicating whether 10bit HEVC decoding is enabled. | [optional] 
**EnableDecodingColorDepth10Vp9** | Pointer to **bool** | Gets or sets a value indicating whether 10bit VP9 decoding is enabled. | [optional] 
**EnableEnhancedNvdecDecoder** | Pointer to **bool** | Gets or sets a value indicating whether the enhanced NVDEC is enabled. | [optional] 
**PreferSystemNativeHwDecoder** | Pointer to **bool** | Gets or sets a value indicating whether the system native hardware decoder should be used. | [optional] 
**EnableIntelLowPowerH264HwEncoder** | Pointer to **bool** | Gets or sets a value indicating whether the Intel H264 low-power hardware encoder should be used. | [optional] 
**EnableIntelLowPowerHevcHwEncoder** | Pointer to **bool** | Gets or sets a value indicating whether the Intel HEVC low-power hardware encoder should be used. | [optional] 
**EnableHardwareEncoding** | Pointer to **bool** | Gets or sets a value indicating whether hardware encoding is enabled. | [optional] 
**AllowHevcEncoding** | Pointer to **bool** | Gets or sets a value indicating whether HEVC encoding is enabled. | [optional] 
**AllowAv1Encoding** | Pointer to **bool** | Gets or sets a value indicating whether AV1 encoding is enabled. | [optional] 
**EnableSubtitleExtraction** | Pointer to **bool** | Gets or sets a value indicating whether subtitle extraction is enabled. | [optional] 
**HardwareDecodingCodecs** | Pointer to **[]string** | Gets or sets the codecs hardware encoding is used for. | [optional] 
**AllowOnDemandMetadataBasedKeyframeExtractionForExtensions** | Pointer to **[]string** | Gets or sets the file extensions on-demand metadata based keyframe extraction is enabled for. | [optional] 

## Methods

### NewEncodingOptions

`func NewEncodingOptions() *EncodingOptions`

NewEncodingOptions instantiates a new EncodingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncodingOptionsWithDefaults

`func NewEncodingOptionsWithDefaults() *EncodingOptions`

NewEncodingOptionsWithDefaults instantiates a new EncodingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncodingThreadCount

`func (o *EncodingOptions) GetEncodingThreadCount() int32`

GetEncodingThreadCount returns the EncodingThreadCount field if non-nil, zero value otherwise.

### GetEncodingThreadCountOk

`func (o *EncodingOptions) GetEncodingThreadCountOk() (*int32, bool)`

GetEncodingThreadCountOk returns a tuple with the EncodingThreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingThreadCount

`func (o *EncodingOptions) SetEncodingThreadCount(v int32)`

SetEncodingThreadCount sets EncodingThreadCount field to given value.

### HasEncodingThreadCount

`func (o *EncodingOptions) HasEncodingThreadCount() bool`

HasEncodingThreadCount returns a boolean if a field has been set.

### GetTranscodingTempPath

`func (o *EncodingOptions) GetTranscodingTempPath() string`

GetTranscodingTempPath returns the TranscodingTempPath field if non-nil, zero value otherwise.

### GetTranscodingTempPathOk

`func (o *EncodingOptions) GetTranscodingTempPathOk() (*string, bool)`

GetTranscodingTempPathOk returns a tuple with the TranscodingTempPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscodingTempPath

`func (o *EncodingOptions) SetTranscodingTempPath(v string)`

SetTranscodingTempPath sets TranscodingTempPath field to given value.

### HasTranscodingTempPath

`func (o *EncodingOptions) HasTranscodingTempPath() bool`

HasTranscodingTempPath returns a boolean if a field has been set.

### SetTranscodingTempPathNil

`func (o *EncodingOptions) SetTranscodingTempPathNil(b bool)`

 SetTranscodingTempPathNil sets the value for TranscodingTempPath to be an explicit nil

### UnsetTranscodingTempPath
`func (o *EncodingOptions) UnsetTranscodingTempPath()`

UnsetTranscodingTempPath ensures that no value is present for TranscodingTempPath, not even an explicit nil
### GetFallbackFontPath

`func (o *EncodingOptions) GetFallbackFontPath() string`

GetFallbackFontPath returns the FallbackFontPath field if non-nil, zero value otherwise.

### GetFallbackFontPathOk

`func (o *EncodingOptions) GetFallbackFontPathOk() (*string, bool)`

GetFallbackFontPathOk returns a tuple with the FallbackFontPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackFontPath

`func (o *EncodingOptions) SetFallbackFontPath(v string)`

SetFallbackFontPath sets FallbackFontPath field to given value.

### HasFallbackFontPath

`func (o *EncodingOptions) HasFallbackFontPath() bool`

HasFallbackFontPath returns a boolean if a field has been set.

### SetFallbackFontPathNil

`func (o *EncodingOptions) SetFallbackFontPathNil(b bool)`

 SetFallbackFontPathNil sets the value for FallbackFontPath to be an explicit nil

### UnsetFallbackFontPath
`func (o *EncodingOptions) UnsetFallbackFontPath()`

UnsetFallbackFontPath ensures that no value is present for FallbackFontPath, not even an explicit nil
### GetEnableFallbackFont

`func (o *EncodingOptions) GetEnableFallbackFont() bool`

GetEnableFallbackFont returns the EnableFallbackFont field if non-nil, zero value otherwise.

### GetEnableFallbackFontOk

`func (o *EncodingOptions) GetEnableFallbackFontOk() (*bool, bool)`

GetEnableFallbackFontOk returns a tuple with the EnableFallbackFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFallbackFont

`func (o *EncodingOptions) SetEnableFallbackFont(v bool)`

SetEnableFallbackFont sets EnableFallbackFont field to given value.

### HasEnableFallbackFont

`func (o *EncodingOptions) HasEnableFallbackFont() bool`

HasEnableFallbackFont returns a boolean if a field has been set.

### GetEnableAudioVbr

`func (o *EncodingOptions) GetEnableAudioVbr() bool`

GetEnableAudioVbr returns the EnableAudioVbr field if non-nil, zero value otherwise.

### GetEnableAudioVbrOk

`func (o *EncodingOptions) GetEnableAudioVbrOk() (*bool, bool)`

GetEnableAudioVbrOk returns a tuple with the EnableAudioVbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAudioVbr

`func (o *EncodingOptions) SetEnableAudioVbr(v bool)`

SetEnableAudioVbr sets EnableAudioVbr field to given value.

### HasEnableAudioVbr

`func (o *EncodingOptions) HasEnableAudioVbr() bool`

HasEnableAudioVbr returns a boolean if a field has been set.

### GetDownMixAudioBoost

`func (o *EncodingOptions) GetDownMixAudioBoost() float64`

GetDownMixAudioBoost returns the DownMixAudioBoost field if non-nil, zero value otherwise.

### GetDownMixAudioBoostOk

`func (o *EncodingOptions) GetDownMixAudioBoostOk() (*float64, bool)`

GetDownMixAudioBoostOk returns a tuple with the DownMixAudioBoost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownMixAudioBoost

`func (o *EncodingOptions) SetDownMixAudioBoost(v float64)`

SetDownMixAudioBoost sets DownMixAudioBoost field to given value.

### HasDownMixAudioBoost

`func (o *EncodingOptions) HasDownMixAudioBoost() bool`

HasDownMixAudioBoost returns a boolean if a field has been set.

### GetDownMixStereoAlgorithm

`func (o *EncodingOptions) GetDownMixStereoAlgorithm() DownMixStereoAlgorithms`

GetDownMixStereoAlgorithm returns the DownMixStereoAlgorithm field if non-nil, zero value otherwise.

### GetDownMixStereoAlgorithmOk

`func (o *EncodingOptions) GetDownMixStereoAlgorithmOk() (*DownMixStereoAlgorithms, bool)`

GetDownMixStereoAlgorithmOk returns a tuple with the DownMixStereoAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownMixStereoAlgorithm

`func (o *EncodingOptions) SetDownMixStereoAlgorithm(v DownMixStereoAlgorithms)`

SetDownMixStereoAlgorithm sets DownMixStereoAlgorithm field to given value.

### HasDownMixStereoAlgorithm

`func (o *EncodingOptions) HasDownMixStereoAlgorithm() bool`

HasDownMixStereoAlgorithm returns a boolean if a field has been set.

### GetMaxMuxingQueueSize

`func (o *EncodingOptions) GetMaxMuxingQueueSize() int32`

GetMaxMuxingQueueSize returns the MaxMuxingQueueSize field if non-nil, zero value otherwise.

### GetMaxMuxingQueueSizeOk

`func (o *EncodingOptions) GetMaxMuxingQueueSizeOk() (*int32, bool)`

GetMaxMuxingQueueSizeOk returns a tuple with the MaxMuxingQueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMuxingQueueSize

`func (o *EncodingOptions) SetMaxMuxingQueueSize(v int32)`

SetMaxMuxingQueueSize sets MaxMuxingQueueSize field to given value.

### HasMaxMuxingQueueSize

`func (o *EncodingOptions) HasMaxMuxingQueueSize() bool`

HasMaxMuxingQueueSize returns a boolean if a field has been set.

### GetEnableThrottling

`func (o *EncodingOptions) GetEnableThrottling() bool`

GetEnableThrottling returns the EnableThrottling field if non-nil, zero value otherwise.

### GetEnableThrottlingOk

`func (o *EncodingOptions) GetEnableThrottlingOk() (*bool, bool)`

GetEnableThrottlingOk returns a tuple with the EnableThrottling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableThrottling

`func (o *EncodingOptions) SetEnableThrottling(v bool)`

SetEnableThrottling sets EnableThrottling field to given value.

### HasEnableThrottling

`func (o *EncodingOptions) HasEnableThrottling() bool`

HasEnableThrottling returns a boolean if a field has been set.

### GetThrottleDelaySeconds

`func (o *EncodingOptions) GetThrottleDelaySeconds() int32`

GetThrottleDelaySeconds returns the ThrottleDelaySeconds field if non-nil, zero value otherwise.

### GetThrottleDelaySecondsOk

`func (o *EncodingOptions) GetThrottleDelaySecondsOk() (*int32, bool)`

GetThrottleDelaySecondsOk returns a tuple with the ThrottleDelaySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottleDelaySeconds

`func (o *EncodingOptions) SetThrottleDelaySeconds(v int32)`

SetThrottleDelaySeconds sets ThrottleDelaySeconds field to given value.

### HasThrottleDelaySeconds

`func (o *EncodingOptions) HasThrottleDelaySeconds() bool`

HasThrottleDelaySeconds returns a boolean if a field has been set.

### GetEnableSegmentDeletion

`func (o *EncodingOptions) GetEnableSegmentDeletion() bool`

GetEnableSegmentDeletion returns the EnableSegmentDeletion field if non-nil, zero value otherwise.

### GetEnableSegmentDeletionOk

`func (o *EncodingOptions) GetEnableSegmentDeletionOk() (*bool, bool)`

GetEnableSegmentDeletionOk returns a tuple with the EnableSegmentDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSegmentDeletion

`func (o *EncodingOptions) SetEnableSegmentDeletion(v bool)`

SetEnableSegmentDeletion sets EnableSegmentDeletion field to given value.

### HasEnableSegmentDeletion

`func (o *EncodingOptions) HasEnableSegmentDeletion() bool`

HasEnableSegmentDeletion returns a boolean if a field has been set.

### GetSegmentKeepSeconds

`func (o *EncodingOptions) GetSegmentKeepSeconds() int32`

GetSegmentKeepSeconds returns the SegmentKeepSeconds field if non-nil, zero value otherwise.

### GetSegmentKeepSecondsOk

`func (o *EncodingOptions) GetSegmentKeepSecondsOk() (*int32, bool)`

GetSegmentKeepSecondsOk returns a tuple with the SegmentKeepSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentKeepSeconds

`func (o *EncodingOptions) SetSegmentKeepSeconds(v int32)`

SetSegmentKeepSeconds sets SegmentKeepSeconds field to given value.

### HasSegmentKeepSeconds

`func (o *EncodingOptions) HasSegmentKeepSeconds() bool`

HasSegmentKeepSeconds returns a boolean if a field has been set.

### GetHardwareAccelerationType

`func (o *EncodingOptions) GetHardwareAccelerationType() string`

GetHardwareAccelerationType returns the HardwareAccelerationType field if non-nil, zero value otherwise.

### GetHardwareAccelerationTypeOk

`func (o *EncodingOptions) GetHardwareAccelerationTypeOk() (*string, bool)`

GetHardwareAccelerationTypeOk returns a tuple with the HardwareAccelerationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareAccelerationType

`func (o *EncodingOptions) SetHardwareAccelerationType(v string)`

SetHardwareAccelerationType sets HardwareAccelerationType field to given value.

### HasHardwareAccelerationType

`func (o *EncodingOptions) HasHardwareAccelerationType() bool`

HasHardwareAccelerationType returns a boolean if a field has been set.

### SetHardwareAccelerationTypeNil

`func (o *EncodingOptions) SetHardwareAccelerationTypeNil(b bool)`

 SetHardwareAccelerationTypeNil sets the value for HardwareAccelerationType to be an explicit nil

### UnsetHardwareAccelerationType
`func (o *EncodingOptions) UnsetHardwareAccelerationType()`

UnsetHardwareAccelerationType ensures that no value is present for HardwareAccelerationType, not even an explicit nil
### GetEncoderAppPath

`func (o *EncodingOptions) GetEncoderAppPath() string`

GetEncoderAppPath returns the EncoderAppPath field if non-nil, zero value otherwise.

### GetEncoderAppPathOk

`func (o *EncodingOptions) GetEncoderAppPathOk() (*string, bool)`

GetEncoderAppPathOk returns a tuple with the EncoderAppPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderAppPath

`func (o *EncodingOptions) SetEncoderAppPath(v string)`

SetEncoderAppPath sets EncoderAppPath field to given value.

### HasEncoderAppPath

`func (o *EncodingOptions) HasEncoderAppPath() bool`

HasEncoderAppPath returns a boolean if a field has been set.

### SetEncoderAppPathNil

`func (o *EncodingOptions) SetEncoderAppPathNil(b bool)`

 SetEncoderAppPathNil sets the value for EncoderAppPath to be an explicit nil

### UnsetEncoderAppPath
`func (o *EncodingOptions) UnsetEncoderAppPath()`

UnsetEncoderAppPath ensures that no value is present for EncoderAppPath, not even an explicit nil
### GetEncoderAppPathDisplay

`func (o *EncodingOptions) GetEncoderAppPathDisplay() string`

GetEncoderAppPathDisplay returns the EncoderAppPathDisplay field if non-nil, zero value otherwise.

### GetEncoderAppPathDisplayOk

`func (o *EncodingOptions) GetEncoderAppPathDisplayOk() (*string, bool)`

GetEncoderAppPathDisplayOk returns a tuple with the EncoderAppPathDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderAppPathDisplay

`func (o *EncodingOptions) SetEncoderAppPathDisplay(v string)`

SetEncoderAppPathDisplay sets EncoderAppPathDisplay field to given value.

### HasEncoderAppPathDisplay

`func (o *EncodingOptions) HasEncoderAppPathDisplay() bool`

HasEncoderAppPathDisplay returns a boolean if a field has been set.

### SetEncoderAppPathDisplayNil

`func (o *EncodingOptions) SetEncoderAppPathDisplayNil(b bool)`

 SetEncoderAppPathDisplayNil sets the value for EncoderAppPathDisplay to be an explicit nil

### UnsetEncoderAppPathDisplay
`func (o *EncodingOptions) UnsetEncoderAppPathDisplay()`

UnsetEncoderAppPathDisplay ensures that no value is present for EncoderAppPathDisplay, not even an explicit nil
### GetVaapiDevice

`func (o *EncodingOptions) GetVaapiDevice() string`

GetVaapiDevice returns the VaapiDevice field if non-nil, zero value otherwise.

### GetVaapiDeviceOk

`func (o *EncodingOptions) GetVaapiDeviceOk() (*string, bool)`

GetVaapiDeviceOk returns a tuple with the VaapiDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaapiDevice

`func (o *EncodingOptions) SetVaapiDevice(v string)`

SetVaapiDevice sets VaapiDevice field to given value.

### HasVaapiDevice

`func (o *EncodingOptions) HasVaapiDevice() bool`

HasVaapiDevice returns a boolean if a field has been set.

### SetVaapiDeviceNil

`func (o *EncodingOptions) SetVaapiDeviceNil(b bool)`

 SetVaapiDeviceNil sets the value for VaapiDevice to be an explicit nil

### UnsetVaapiDevice
`func (o *EncodingOptions) UnsetVaapiDevice()`

UnsetVaapiDevice ensures that no value is present for VaapiDevice, not even an explicit nil
### GetEnableTonemapping

`func (o *EncodingOptions) GetEnableTonemapping() bool`

GetEnableTonemapping returns the EnableTonemapping field if non-nil, zero value otherwise.

### GetEnableTonemappingOk

`func (o *EncodingOptions) GetEnableTonemappingOk() (*bool, bool)`

GetEnableTonemappingOk returns a tuple with the EnableTonemapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTonemapping

`func (o *EncodingOptions) SetEnableTonemapping(v bool)`

SetEnableTonemapping sets EnableTonemapping field to given value.

### HasEnableTonemapping

`func (o *EncodingOptions) HasEnableTonemapping() bool`

HasEnableTonemapping returns a boolean if a field has been set.

### GetEnableVppTonemapping

`func (o *EncodingOptions) GetEnableVppTonemapping() bool`

GetEnableVppTonemapping returns the EnableVppTonemapping field if non-nil, zero value otherwise.

### GetEnableVppTonemappingOk

`func (o *EncodingOptions) GetEnableVppTonemappingOk() (*bool, bool)`

GetEnableVppTonemappingOk returns a tuple with the EnableVppTonemapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVppTonemapping

`func (o *EncodingOptions) SetEnableVppTonemapping(v bool)`

SetEnableVppTonemapping sets EnableVppTonemapping field to given value.

### HasEnableVppTonemapping

`func (o *EncodingOptions) HasEnableVppTonemapping() bool`

HasEnableVppTonemapping returns a boolean if a field has been set.

### GetEnableVideoToolboxTonemapping

`func (o *EncodingOptions) GetEnableVideoToolboxTonemapping() bool`

GetEnableVideoToolboxTonemapping returns the EnableVideoToolboxTonemapping field if non-nil, zero value otherwise.

### GetEnableVideoToolboxTonemappingOk

`func (o *EncodingOptions) GetEnableVideoToolboxTonemappingOk() (*bool, bool)`

GetEnableVideoToolboxTonemappingOk returns a tuple with the EnableVideoToolboxTonemapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableVideoToolboxTonemapping

`func (o *EncodingOptions) SetEnableVideoToolboxTonemapping(v bool)`

SetEnableVideoToolboxTonemapping sets EnableVideoToolboxTonemapping field to given value.

### HasEnableVideoToolboxTonemapping

`func (o *EncodingOptions) HasEnableVideoToolboxTonemapping() bool`

HasEnableVideoToolboxTonemapping returns a boolean if a field has been set.

### GetTonemappingAlgorithm

`func (o *EncodingOptions) GetTonemappingAlgorithm() string`

GetTonemappingAlgorithm returns the TonemappingAlgorithm field if non-nil, zero value otherwise.

### GetTonemappingAlgorithmOk

`func (o *EncodingOptions) GetTonemappingAlgorithmOk() (*string, bool)`

GetTonemappingAlgorithmOk returns a tuple with the TonemappingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingAlgorithm

`func (o *EncodingOptions) SetTonemappingAlgorithm(v string)`

SetTonemappingAlgorithm sets TonemappingAlgorithm field to given value.

### HasTonemappingAlgorithm

`func (o *EncodingOptions) HasTonemappingAlgorithm() bool`

HasTonemappingAlgorithm returns a boolean if a field has been set.

### SetTonemappingAlgorithmNil

`func (o *EncodingOptions) SetTonemappingAlgorithmNil(b bool)`

 SetTonemappingAlgorithmNil sets the value for TonemappingAlgorithm to be an explicit nil

### UnsetTonemappingAlgorithm
`func (o *EncodingOptions) UnsetTonemappingAlgorithm()`

UnsetTonemappingAlgorithm ensures that no value is present for TonemappingAlgorithm, not even an explicit nil
### GetTonemappingMode

`func (o *EncodingOptions) GetTonemappingMode() string`

GetTonemappingMode returns the TonemappingMode field if non-nil, zero value otherwise.

### GetTonemappingModeOk

`func (o *EncodingOptions) GetTonemappingModeOk() (*string, bool)`

GetTonemappingModeOk returns a tuple with the TonemappingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingMode

`func (o *EncodingOptions) SetTonemappingMode(v string)`

SetTonemappingMode sets TonemappingMode field to given value.

### HasTonemappingMode

`func (o *EncodingOptions) HasTonemappingMode() bool`

HasTonemappingMode returns a boolean if a field has been set.

### SetTonemappingModeNil

`func (o *EncodingOptions) SetTonemappingModeNil(b bool)`

 SetTonemappingModeNil sets the value for TonemappingMode to be an explicit nil

### UnsetTonemappingMode
`func (o *EncodingOptions) UnsetTonemappingMode()`

UnsetTonemappingMode ensures that no value is present for TonemappingMode, not even an explicit nil
### GetTonemappingRange

`func (o *EncodingOptions) GetTonemappingRange() string`

GetTonemappingRange returns the TonemappingRange field if non-nil, zero value otherwise.

### GetTonemappingRangeOk

`func (o *EncodingOptions) GetTonemappingRangeOk() (*string, bool)`

GetTonemappingRangeOk returns a tuple with the TonemappingRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingRange

`func (o *EncodingOptions) SetTonemappingRange(v string)`

SetTonemappingRange sets TonemappingRange field to given value.

### HasTonemappingRange

`func (o *EncodingOptions) HasTonemappingRange() bool`

HasTonemappingRange returns a boolean if a field has been set.

### SetTonemappingRangeNil

`func (o *EncodingOptions) SetTonemappingRangeNil(b bool)`

 SetTonemappingRangeNil sets the value for TonemappingRange to be an explicit nil

### UnsetTonemappingRange
`func (o *EncodingOptions) UnsetTonemappingRange()`

UnsetTonemappingRange ensures that no value is present for TonemappingRange, not even an explicit nil
### GetTonemappingDesat

`func (o *EncodingOptions) GetTonemappingDesat() float64`

GetTonemappingDesat returns the TonemappingDesat field if non-nil, zero value otherwise.

### GetTonemappingDesatOk

`func (o *EncodingOptions) GetTonemappingDesatOk() (*float64, bool)`

GetTonemappingDesatOk returns a tuple with the TonemappingDesat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingDesat

`func (o *EncodingOptions) SetTonemappingDesat(v float64)`

SetTonemappingDesat sets TonemappingDesat field to given value.

### HasTonemappingDesat

`func (o *EncodingOptions) HasTonemappingDesat() bool`

HasTonemappingDesat returns a boolean if a field has been set.

### GetTonemappingPeak

`func (o *EncodingOptions) GetTonemappingPeak() float64`

GetTonemappingPeak returns the TonemappingPeak field if non-nil, zero value otherwise.

### GetTonemappingPeakOk

`func (o *EncodingOptions) GetTonemappingPeakOk() (*float64, bool)`

GetTonemappingPeakOk returns a tuple with the TonemappingPeak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingPeak

`func (o *EncodingOptions) SetTonemappingPeak(v float64)`

SetTonemappingPeak sets TonemappingPeak field to given value.

### HasTonemappingPeak

`func (o *EncodingOptions) HasTonemappingPeak() bool`

HasTonemappingPeak returns a boolean if a field has been set.

### GetTonemappingParam

`func (o *EncodingOptions) GetTonemappingParam() float64`

GetTonemappingParam returns the TonemappingParam field if non-nil, zero value otherwise.

### GetTonemappingParamOk

`func (o *EncodingOptions) GetTonemappingParamOk() (*float64, bool)`

GetTonemappingParamOk returns a tuple with the TonemappingParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonemappingParam

`func (o *EncodingOptions) SetTonemappingParam(v float64)`

SetTonemappingParam sets TonemappingParam field to given value.

### HasTonemappingParam

`func (o *EncodingOptions) HasTonemappingParam() bool`

HasTonemappingParam returns a boolean if a field has been set.

### GetVppTonemappingBrightness

`func (o *EncodingOptions) GetVppTonemappingBrightness() float64`

GetVppTonemappingBrightness returns the VppTonemappingBrightness field if non-nil, zero value otherwise.

### GetVppTonemappingBrightnessOk

`func (o *EncodingOptions) GetVppTonemappingBrightnessOk() (*float64, bool)`

GetVppTonemappingBrightnessOk returns a tuple with the VppTonemappingBrightness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppTonemappingBrightness

`func (o *EncodingOptions) SetVppTonemappingBrightness(v float64)`

SetVppTonemappingBrightness sets VppTonemappingBrightness field to given value.

### HasVppTonemappingBrightness

`func (o *EncodingOptions) HasVppTonemappingBrightness() bool`

HasVppTonemappingBrightness returns a boolean if a field has been set.

### GetVppTonemappingContrast

`func (o *EncodingOptions) GetVppTonemappingContrast() float64`

GetVppTonemappingContrast returns the VppTonemappingContrast field if non-nil, zero value otherwise.

### GetVppTonemappingContrastOk

`func (o *EncodingOptions) GetVppTonemappingContrastOk() (*float64, bool)`

GetVppTonemappingContrastOk returns a tuple with the VppTonemappingContrast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVppTonemappingContrast

`func (o *EncodingOptions) SetVppTonemappingContrast(v float64)`

SetVppTonemappingContrast sets VppTonemappingContrast field to given value.

### HasVppTonemappingContrast

`func (o *EncodingOptions) HasVppTonemappingContrast() bool`

HasVppTonemappingContrast returns a boolean if a field has been set.

### GetH264Crf

`func (o *EncodingOptions) GetH264Crf() int32`

GetH264Crf returns the H264Crf field if non-nil, zero value otherwise.

### GetH264CrfOk

`func (o *EncodingOptions) GetH264CrfOk() (*int32, bool)`

GetH264CrfOk returns a tuple with the H264Crf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH264Crf

`func (o *EncodingOptions) SetH264Crf(v int32)`

SetH264Crf sets H264Crf field to given value.

### HasH264Crf

`func (o *EncodingOptions) HasH264Crf() bool`

HasH264Crf returns a boolean if a field has been set.

### GetH265Crf

`func (o *EncodingOptions) GetH265Crf() int32`

GetH265Crf returns the H265Crf field if non-nil, zero value otherwise.

### GetH265CrfOk

`func (o *EncodingOptions) GetH265CrfOk() (*int32, bool)`

GetH265CrfOk returns a tuple with the H265Crf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH265Crf

`func (o *EncodingOptions) SetH265Crf(v int32)`

SetH265Crf sets H265Crf field to given value.

### HasH265Crf

`func (o *EncodingOptions) HasH265Crf() bool`

HasH265Crf returns a boolean if a field has been set.

### GetEncoderPreset

`func (o *EncodingOptions) GetEncoderPreset() string`

GetEncoderPreset returns the EncoderPreset field if non-nil, zero value otherwise.

### GetEncoderPresetOk

`func (o *EncodingOptions) GetEncoderPresetOk() (*string, bool)`

GetEncoderPresetOk returns a tuple with the EncoderPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoderPreset

`func (o *EncodingOptions) SetEncoderPreset(v string)`

SetEncoderPreset sets EncoderPreset field to given value.

### HasEncoderPreset

`func (o *EncodingOptions) HasEncoderPreset() bool`

HasEncoderPreset returns a boolean if a field has been set.

### SetEncoderPresetNil

`func (o *EncodingOptions) SetEncoderPresetNil(b bool)`

 SetEncoderPresetNil sets the value for EncoderPreset to be an explicit nil

### UnsetEncoderPreset
`func (o *EncodingOptions) UnsetEncoderPreset()`

UnsetEncoderPreset ensures that no value is present for EncoderPreset, not even an explicit nil
### GetDeinterlaceDoubleRate

`func (o *EncodingOptions) GetDeinterlaceDoubleRate() bool`

GetDeinterlaceDoubleRate returns the DeinterlaceDoubleRate field if non-nil, zero value otherwise.

### GetDeinterlaceDoubleRateOk

`func (o *EncodingOptions) GetDeinterlaceDoubleRateOk() (*bool, bool)`

GetDeinterlaceDoubleRateOk returns a tuple with the DeinterlaceDoubleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeinterlaceDoubleRate

`func (o *EncodingOptions) SetDeinterlaceDoubleRate(v bool)`

SetDeinterlaceDoubleRate sets DeinterlaceDoubleRate field to given value.

### HasDeinterlaceDoubleRate

`func (o *EncodingOptions) HasDeinterlaceDoubleRate() bool`

HasDeinterlaceDoubleRate returns a boolean if a field has been set.

### GetDeinterlaceMethod

`func (o *EncodingOptions) GetDeinterlaceMethod() string`

GetDeinterlaceMethod returns the DeinterlaceMethod field if non-nil, zero value otherwise.

### GetDeinterlaceMethodOk

`func (o *EncodingOptions) GetDeinterlaceMethodOk() (*string, bool)`

GetDeinterlaceMethodOk returns a tuple with the DeinterlaceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeinterlaceMethod

`func (o *EncodingOptions) SetDeinterlaceMethod(v string)`

SetDeinterlaceMethod sets DeinterlaceMethod field to given value.

### HasDeinterlaceMethod

`func (o *EncodingOptions) HasDeinterlaceMethod() bool`

HasDeinterlaceMethod returns a boolean if a field has been set.

### SetDeinterlaceMethodNil

`func (o *EncodingOptions) SetDeinterlaceMethodNil(b bool)`

 SetDeinterlaceMethodNil sets the value for DeinterlaceMethod to be an explicit nil

### UnsetDeinterlaceMethod
`func (o *EncodingOptions) UnsetDeinterlaceMethod()`

UnsetDeinterlaceMethod ensures that no value is present for DeinterlaceMethod, not even an explicit nil
### GetEnableDecodingColorDepth10Hevc

`func (o *EncodingOptions) GetEnableDecodingColorDepth10Hevc() bool`

GetEnableDecodingColorDepth10Hevc returns the EnableDecodingColorDepth10Hevc field if non-nil, zero value otherwise.

### GetEnableDecodingColorDepth10HevcOk

`func (o *EncodingOptions) GetEnableDecodingColorDepth10HevcOk() (*bool, bool)`

GetEnableDecodingColorDepth10HevcOk returns a tuple with the EnableDecodingColorDepth10Hevc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDecodingColorDepth10Hevc

`func (o *EncodingOptions) SetEnableDecodingColorDepth10Hevc(v bool)`

SetEnableDecodingColorDepth10Hevc sets EnableDecodingColorDepth10Hevc field to given value.

### HasEnableDecodingColorDepth10Hevc

`func (o *EncodingOptions) HasEnableDecodingColorDepth10Hevc() bool`

HasEnableDecodingColorDepth10Hevc returns a boolean if a field has been set.

### GetEnableDecodingColorDepth10Vp9

`func (o *EncodingOptions) GetEnableDecodingColorDepth10Vp9() bool`

GetEnableDecodingColorDepth10Vp9 returns the EnableDecodingColorDepth10Vp9 field if non-nil, zero value otherwise.

### GetEnableDecodingColorDepth10Vp9Ok

`func (o *EncodingOptions) GetEnableDecodingColorDepth10Vp9Ok() (*bool, bool)`

GetEnableDecodingColorDepth10Vp9Ok returns a tuple with the EnableDecodingColorDepth10Vp9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDecodingColorDepth10Vp9

`func (o *EncodingOptions) SetEnableDecodingColorDepth10Vp9(v bool)`

SetEnableDecodingColorDepth10Vp9 sets EnableDecodingColorDepth10Vp9 field to given value.

### HasEnableDecodingColorDepth10Vp9

`func (o *EncodingOptions) HasEnableDecodingColorDepth10Vp9() bool`

HasEnableDecodingColorDepth10Vp9 returns a boolean if a field has been set.

### GetEnableEnhancedNvdecDecoder

`func (o *EncodingOptions) GetEnableEnhancedNvdecDecoder() bool`

GetEnableEnhancedNvdecDecoder returns the EnableEnhancedNvdecDecoder field if non-nil, zero value otherwise.

### GetEnableEnhancedNvdecDecoderOk

`func (o *EncodingOptions) GetEnableEnhancedNvdecDecoderOk() (*bool, bool)`

GetEnableEnhancedNvdecDecoderOk returns a tuple with the EnableEnhancedNvdecDecoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEnhancedNvdecDecoder

`func (o *EncodingOptions) SetEnableEnhancedNvdecDecoder(v bool)`

SetEnableEnhancedNvdecDecoder sets EnableEnhancedNvdecDecoder field to given value.

### HasEnableEnhancedNvdecDecoder

`func (o *EncodingOptions) HasEnableEnhancedNvdecDecoder() bool`

HasEnableEnhancedNvdecDecoder returns a boolean if a field has been set.

### GetPreferSystemNativeHwDecoder

`func (o *EncodingOptions) GetPreferSystemNativeHwDecoder() bool`

GetPreferSystemNativeHwDecoder returns the PreferSystemNativeHwDecoder field if non-nil, zero value otherwise.

### GetPreferSystemNativeHwDecoderOk

`func (o *EncodingOptions) GetPreferSystemNativeHwDecoderOk() (*bool, bool)`

GetPreferSystemNativeHwDecoderOk returns a tuple with the PreferSystemNativeHwDecoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferSystemNativeHwDecoder

`func (o *EncodingOptions) SetPreferSystemNativeHwDecoder(v bool)`

SetPreferSystemNativeHwDecoder sets PreferSystemNativeHwDecoder field to given value.

### HasPreferSystemNativeHwDecoder

`func (o *EncodingOptions) HasPreferSystemNativeHwDecoder() bool`

HasPreferSystemNativeHwDecoder returns a boolean if a field has been set.

### GetEnableIntelLowPowerH264HwEncoder

`func (o *EncodingOptions) GetEnableIntelLowPowerH264HwEncoder() bool`

GetEnableIntelLowPowerH264HwEncoder returns the EnableIntelLowPowerH264HwEncoder field if non-nil, zero value otherwise.

### GetEnableIntelLowPowerH264HwEncoderOk

`func (o *EncodingOptions) GetEnableIntelLowPowerH264HwEncoderOk() (*bool, bool)`

GetEnableIntelLowPowerH264HwEncoderOk returns a tuple with the EnableIntelLowPowerH264HwEncoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntelLowPowerH264HwEncoder

`func (o *EncodingOptions) SetEnableIntelLowPowerH264HwEncoder(v bool)`

SetEnableIntelLowPowerH264HwEncoder sets EnableIntelLowPowerH264HwEncoder field to given value.

### HasEnableIntelLowPowerH264HwEncoder

`func (o *EncodingOptions) HasEnableIntelLowPowerH264HwEncoder() bool`

HasEnableIntelLowPowerH264HwEncoder returns a boolean if a field has been set.

### GetEnableIntelLowPowerHevcHwEncoder

`func (o *EncodingOptions) GetEnableIntelLowPowerHevcHwEncoder() bool`

GetEnableIntelLowPowerHevcHwEncoder returns the EnableIntelLowPowerHevcHwEncoder field if non-nil, zero value otherwise.

### GetEnableIntelLowPowerHevcHwEncoderOk

`func (o *EncodingOptions) GetEnableIntelLowPowerHevcHwEncoderOk() (*bool, bool)`

GetEnableIntelLowPowerHevcHwEncoderOk returns a tuple with the EnableIntelLowPowerHevcHwEncoder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIntelLowPowerHevcHwEncoder

`func (o *EncodingOptions) SetEnableIntelLowPowerHevcHwEncoder(v bool)`

SetEnableIntelLowPowerHevcHwEncoder sets EnableIntelLowPowerHevcHwEncoder field to given value.

### HasEnableIntelLowPowerHevcHwEncoder

`func (o *EncodingOptions) HasEnableIntelLowPowerHevcHwEncoder() bool`

HasEnableIntelLowPowerHevcHwEncoder returns a boolean if a field has been set.

### GetEnableHardwareEncoding

`func (o *EncodingOptions) GetEnableHardwareEncoding() bool`

GetEnableHardwareEncoding returns the EnableHardwareEncoding field if non-nil, zero value otherwise.

### GetEnableHardwareEncodingOk

`func (o *EncodingOptions) GetEnableHardwareEncodingOk() (*bool, bool)`

GetEnableHardwareEncodingOk returns a tuple with the EnableHardwareEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHardwareEncoding

`func (o *EncodingOptions) SetEnableHardwareEncoding(v bool)`

SetEnableHardwareEncoding sets EnableHardwareEncoding field to given value.

### HasEnableHardwareEncoding

`func (o *EncodingOptions) HasEnableHardwareEncoding() bool`

HasEnableHardwareEncoding returns a boolean if a field has been set.

### GetAllowHevcEncoding

`func (o *EncodingOptions) GetAllowHevcEncoding() bool`

GetAllowHevcEncoding returns the AllowHevcEncoding field if non-nil, zero value otherwise.

### GetAllowHevcEncodingOk

`func (o *EncodingOptions) GetAllowHevcEncodingOk() (*bool, bool)`

GetAllowHevcEncodingOk returns a tuple with the AllowHevcEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHevcEncoding

`func (o *EncodingOptions) SetAllowHevcEncoding(v bool)`

SetAllowHevcEncoding sets AllowHevcEncoding field to given value.

### HasAllowHevcEncoding

`func (o *EncodingOptions) HasAllowHevcEncoding() bool`

HasAllowHevcEncoding returns a boolean if a field has been set.

### GetAllowAv1Encoding

`func (o *EncodingOptions) GetAllowAv1Encoding() bool`

GetAllowAv1Encoding returns the AllowAv1Encoding field if non-nil, zero value otherwise.

### GetAllowAv1EncodingOk

`func (o *EncodingOptions) GetAllowAv1EncodingOk() (*bool, bool)`

GetAllowAv1EncodingOk returns a tuple with the AllowAv1Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAv1Encoding

`func (o *EncodingOptions) SetAllowAv1Encoding(v bool)`

SetAllowAv1Encoding sets AllowAv1Encoding field to given value.

### HasAllowAv1Encoding

`func (o *EncodingOptions) HasAllowAv1Encoding() bool`

HasAllowAv1Encoding returns a boolean if a field has been set.

### GetEnableSubtitleExtraction

`func (o *EncodingOptions) GetEnableSubtitleExtraction() bool`

GetEnableSubtitleExtraction returns the EnableSubtitleExtraction field if non-nil, zero value otherwise.

### GetEnableSubtitleExtractionOk

`func (o *EncodingOptions) GetEnableSubtitleExtractionOk() (*bool, bool)`

GetEnableSubtitleExtractionOk returns a tuple with the EnableSubtitleExtraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSubtitleExtraction

`func (o *EncodingOptions) SetEnableSubtitleExtraction(v bool)`

SetEnableSubtitleExtraction sets EnableSubtitleExtraction field to given value.

### HasEnableSubtitleExtraction

`func (o *EncodingOptions) HasEnableSubtitleExtraction() bool`

HasEnableSubtitleExtraction returns a boolean if a field has been set.

### GetHardwareDecodingCodecs

`func (o *EncodingOptions) GetHardwareDecodingCodecs() []string`

GetHardwareDecodingCodecs returns the HardwareDecodingCodecs field if non-nil, zero value otherwise.

### GetHardwareDecodingCodecsOk

`func (o *EncodingOptions) GetHardwareDecodingCodecsOk() (*[]string, bool)`

GetHardwareDecodingCodecsOk returns a tuple with the HardwareDecodingCodecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareDecodingCodecs

`func (o *EncodingOptions) SetHardwareDecodingCodecs(v []string)`

SetHardwareDecodingCodecs sets HardwareDecodingCodecs field to given value.

### HasHardwareDecodingCodecs

`func (o *EncodingOptions) HasHardwareDecodingCodecs() bool`

HasHardwareDecodingCodecs returns a boolean if a field has been set.

### SetHardwareDecodingCodecsNil

`func (o *EncodingOptions) SetHardwareDecodingCodecsNil(b bool)`

 SetHardwareDecodingCodecsNil sets the value for HardwareDecodingCodecs to be an explicit nil

### UnsetHardwareDecodingCodecs
`func (o *EncodingOptions) UnsetHardwareDecodingCodecs()`

UnsetHardwareDecodingCodecs ensures that no value is present for HardwareDecodingCodecs, not even an explicit nil
### GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions

`func (o *EncodingOptions) GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions() []string`

GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions returns the AllowOnDemandMetadataBasedKeyframeExtractionForExtensions field if non-nil, zero value otherwise.

### GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsOk

`func (o *EncodingOptions) GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsOk() (*[]string, bool)`

GetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsOk returns a tuple with the AllowOnDemandMetadataBasedKeyframeExtractionForExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions

`func (o *EncodingOptions) SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions(v []string)`

SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions sets AllowOnDemandMetadataBasedKeyframeExtractionForExtensions field to given value.

### HasAllowOnDemandMetadataBasedKeyframeExtractionForExtensions

`func (o *EncodingOptions) HasAllowOnDemandMetadataBasedKeyframeExtractionForExtensions() bool`

HasAllowOnDemandMetadataBasedKeyframeExtractionForExtensions returns a boolean if a field has been set.

### SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsNil

`func (o *EncodingOptions) SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsNil(b bool)`

 SetAllowOnDemandMetadataBasedKeyframeExtractionForExtensionsNil sets the value for AllowOnDemandMetadataBasedKeyframeExtractionForExtensions to be an explicit nil

### UnsetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions
`func (o *EncodingOptions) UnsetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions()`

UnsetAllowOnDemandMetadataBasedKeyframeExtractionForExtensions ensures that no value is present for AllowOnDemandMetadataBasedKeyframeExtractionForExtensions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


