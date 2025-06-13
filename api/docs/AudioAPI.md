# \AudioAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAudioStream**](AudioAPI.md#GetAudioStream) | **Get** /Audio/{itemId}/stream | Gets an audio stream.
[**GetAudioStreamByContainer**](AudioAPI.md#GetAudioStreamByContainer) | **Get** /Audio/{itemId}/stream.{container} | Gets an audio stream.
[**HeadAudioStream**](AudioAPI.md#HeadAudioStream) | **Head** /Audio/{itemId}/stream | Gets an audio stream.
[**HeadAudioStreamByContainer**](AudioAPI.md#HeadAudioStreamByContainer) | **Head** /Audio/{itemId}/stream.{container} | Gets an audio stream.



## GetAudioStream

> *os.File GetAudioStream(ctx, itemId).Container(container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()

Gets an audio stream.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sj14/jellyfin-go/api"
)

func main() {
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The item id.
	container := "container_example" // string | The audio container. (optional)
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. (optional)
	params := "params_example" // string | The streaming parameters. (optional)
	tag := "tag_example" // string | The tag. (optional)
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	playSessionId := "playSessionId_example" // string | The play session id. (optional)
	segmentContainer := "segmentContainer_example" // string | The segment container. (optional)
	segmentLength := int32(56) // int32 | The segment length. (optional)
	minSegments := int32(56) // int32 | The minimum number of segments. (optional)
	mediaSourceId := "mediaSourceId_example" // string | The media version id, if playing an alternate version. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify an audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	allowVideoStreamCopy := true // bool | Whether or not to allow copying of the video stream url. (optional)
	allowAudioStreamCopy := true // bool | Whether or not to allow copying of the audio stream url. (optional)
	breakOnNonKeyFrames := true // bool | Optional. Whether to break on non key frames. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100. (optional)
	maxAudioBitDepth := int32(56) // int32 | Optional. The maximum audio bit depth. (optional)
	audioBitRate := int32(56) // int32 | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. (optional)
	audioChannels := int32(56) // int32 | Optional. Specify a specific number of audio channels to encode to, e.g. 2. (optional)
	maxAudioChannels := int32(56) // int32 | Optional. Specify a maximum number of audio channels to encode to, e.g. 2. (optional)
	profile := "profile_example" // string | Optional. Specify a specific an encoder profile (varies by encoder), e.g. main, baseline, high. (optional)
	level := "level_example" // string | Optional. Specify a level for the encoder profile (varies by encoder), e.g. 3, 3.1. (optional)
	framerate := float32(3.4) // float32 | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	maxFramerate := float32(3.4) // float32 | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	copyTimestamps := true // bool | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms. (optional)
	width := int32(56) // int32 | Optional. The fixed horizontal resolution of the encoded video. (optional)
	height := int32(56) // int32 | Optional. The fixed vertical resolution of the encoded video. (optional)
	videoBitRate := int32(56) // int32 | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. (optional)
	subtitleStreamIndex := int32(56) // int32 | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. (optional)
	subtitleMethod := "subtitleMethod_example" // SubtitleDeliveryMethod | Optional. Specify the subtitle delivery method. (optional)
	maxRefFrames := int32(56) // int32 | Optional. (optional)
	maxVideoBitDepth := int32(56) // int32 | Optional. The maximum video bit depth. (optional)
	requireAvc := true // bool | Optional. Whether to require avc. (optional)
	deInterlace := true // bool | Optional. Whether to deinterlace the video. (optional)
	requireNonAnamorphic := true // bool | Optional. Whether to require a non anamorphic stream. (optional)
	transcodingMaxAudioChannels := int32(56) // int32 | Optional. The maximum number of audio channels to transcode. (optional)
	cpuCoreLimit := int32(56) // int32 | Optional. The limit of how many cpu cores to use. (optional)
	liveStreamId := "liveStreamId_example" // string | The live stream id. (optional)
	enableMpegtsM2TsMode := true // bool | Optional. Whether to enable the MpegtsM2Ts mode. (optional)
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)
	enableAudioVbrEncoding := true // bool | Optional. Whether to enable Audio Encoding. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.GetAudioStream(context.Background(), itemId).Container(container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.GetAudioStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAudioStream`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.GetAudioStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **container** | **string** | The audio container. | 
 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. | 
 **params** | **string** | The streaming parameters. | 
 **tag** | **string** | The tag. | 
 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **playSessionId** | **string** | The play session id. | 
 **segmentContainer** | **string** | The segment container. | 
 **segmentLength** | **int32** | The segment length. | 
 **minSegments** | **int32** | The minimum number of segments. | 
 **mediaSourceId** | **string** | The media version id, if playing an alternate version. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify an audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **allowVideoStreamCopy** | **bool** | Whether or not to allow copying of the video stream url. | 
 **allowAudioStreamCopy** | **bool** | Whether or not to allow copying of the audio stream url. | 
 **breakOnNonKeyFrames** | **bool** | Optional. Whether to break on non key frames. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100. | 
 **maxAudioBitDepth** | **int32** | Optional. The maximum audio bit depth. | 
 **audioBitRate** | **int32** | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **int32** | Optional. Specify a specific number of audio channels to encode to, e.g. 2. | 
 **maxAudioChannels** | **int32** | Optional. Specify a maximum number of audio channels to encode to, e.g. 2. | 
 **profile** | **string** | Optional. Specify a specific an encoder profile (varies by encoder), e.g. main, baseline, high. | 
 **level** | **string** | Optional. Specify a level for the encoder profile (varies by encoder), e.g. 3, 3.1. | 
 **framerate** | **float32** | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **float32** | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **bool** | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **width** | **int32** | Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **int32** | Optional. The fixed vertical resolution of the encoded video. | 
 **videoBitRate** | **int32** | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **int32** | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | **SubtitleDeliveryMethod** | Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **int32** | Optional. | 
 **maxVideoBitDepth** | **int32** | Optional. The maximum video bit depth. | 
 **requireAvc** | **bool** | Optional. Whether to require avc. | 
 **deInterlace** | **bool** | Optional. Whether to deinterlace the video. | 
 **requireNonAnamorphic** | **bool** | Optional. Whether to require a non anamorphic stream. | 
 **transcodingMaxAudioChannels** | **int32** | Optional. The maximum number of audio channels to transcode. | 
 **cpuCoreLimit** | **int32** | Optional. The limit of how many cpu cores to use. | 
 **liveStreamId** | **string** | The live stream id. | 
 **enableMpegtsM2TsMode** | **bool** | Optional. Whether to enable the MpegtsM2Ts mode. | 
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 
 **enableAudioVbrEncoding** | **bool** | Optional. Whether to enable Audio Encoding. | [default to true]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/*, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudioStreamByContainer

> *os.File GetAudioStreamByContainer(ctx, itemId, container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()

Gets an audio stream.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sj14/jellyfin-go/api"
)

func main() {
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The item id.
	container := "container_example" // string | The audio container.
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. (optional)
	params := "params_example" // string | The streaming parameters. (optional)
	tag := "tag_example" // string | The tag. (optional)
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	playSessionId := "playSessionId_example" // string | The play session id. (optional)
	segmentContainer := "segmentContainer_example" // string | The segment container. (optional)
	segmentLength := int32(56) // int32 | The segment length. (optional)
	minSegments := int32(56) // int32 | The minimum number of segments. (optional)
	mediaSourceId := "mediaSourceId_example" // string | The media version id, if playing an alternate version. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify an audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	allowVideoStreamCopy := true // bool | Whether or not to allow copying of the video stream url. (optional)
	allowAudioStreamCopy := true // bool | Whether or not to allow copying of the audio stream url. (optional)
	breakOnNonKeyFrames := true // bool | Optional. Whether to break on non key frames. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100. (optional)
	maxAudioBitDepth := int32(56) // int32 | Optional. The maximum audio bit depth. (optional)
	audioBitRate := int32(56) // int32 | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. (optional)
	audioChannels := int32(56) // int32 | Optional. Specify a specific number of audio channels to encode to, e.g. 2. (optional)
	maxAudioChannels := int32(56) // int32 | Optional. Specify a maximum number of audio channels to encode to, e.g. 2. (optional)
	profile := "profile_example" // string | Optional. Specify a specific an encoder profile (varies by encoder), e.g. main, baseline, high. (optional)
	level := "level_example" // string | Optional. Specify a level for the encoder profile (varies by encoder), e.g. 3, 3.1. (optional)
	framerate := float32(3.4) // float32 | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	maxFramerate := float32(3.4) // float32 | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	copyTimestamps := true // bool | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms. (optional)
	width := int32(56) // int32 | Optional. The fixed horizontal resolution of the encoded video. (optional)
	height := int32(56) // int32 | Optional. The fixed vertical resolution of the encoded video. (optional)
	videoBitRate := int32(56) // int32 | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. (optional)
	subtitleStreamIndex := int32(56) // int32 | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. (optional)
	subtitleMethod := "subtitleMethod_example" // SubtitleDeliveryMethod | Optional. Specify the subtitle delivery method. (optional)
	maxRefFrames := int32(56) // int32 | Optional. (optional)
	maxVideoBitDepth := int32(56) // int32 | Optional. The maximum video bit depth. (optional)
	requireAvc := true // bool | Optional. Whether to require avc. (optional)
	deInterlace := true // bool | Optional. Whether to deinterlace the video. (optional)
	requireNonAnamorphic := true // bool | Optional. Whether to require a non anamorphic stream. (optional)
	transcodingMaxAudioChannels := int32(56) // int32 | Optional. The maximum number of audio channels to transcode. (optional)
	cpuCoreLimit := int32(56) // int32 | Optional. The limit of how many cpu cores to use. (optional)
	liveStreamId := "liveStreamId_example" // string | The live stream id. (optional)
	enableMpegtsM2TsMode := true // bool | Optional. Whether to enable the MpegtsM2Ts mode. (optional)
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)
	enableAudioVbrEncoding := true // bool | Optional. Whether to enable Audio Encoding. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.GetAudioStreamByContainer(context.Background(), itemId, container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.GetAudioStreamByContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAudioStreamByContainer`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.GetAudioStreamByContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**container** | **string** | The audio container. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudioStreamByContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. | 
 **params** | **string** | The streaming parameters. | 
 **tag** | **string** | The tag. | 
 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **playSessionId** | **string** | The play session id. | 
 **segmentContainer** | **string** | The segment container. | 
 **segmentLength** | **int32** | The segment length. | 
 **minSegments** | **int32** | The minimum number of segments. | 
 **mediaSourceId** | **string** | The media version id, if playing an alternate version. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify an audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **allowVideoStreamCopy** | **bool** | Whether or not to allow copying of the video stream url. | 
 **allowAudioStreamCopy** | **bool** | Whether or not to allow copying of the audio stream url. | 
 **breakOnNonKeyFrames** | **bool** | Optional. Whether to break on non key frames. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100. | 
 **maxAudioBitDepth** | **int32** | Optional. The maximum audio bit depth. | 
 **audioBitRate** | **int32** | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **int32** | Optional. Specify a specific number of audio channels to encode to, e.g. 2. | 
 **maxAudioChannels** | **int32** | Optional. Specify a maximum number of audio channels to encode to, e.g. 2. | 
 **profile** | **string** | Optional. Specify a specific an encoder profile (varies by encoder), e.g. main, baseline, high. | 
 **level** | **string** | Optional. Specify a level for the encoder profile (varies by encoder), e.g. 3, 3.1. | 
 **framerate** | **float32** | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **float32** | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **bool** | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **width** | **int32** | Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **int32** | Optional. The fixed vertical resolution of the encoded video. | 
 **videoBitRate** | **int32** | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **int32** | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | **SubtitleDeliveryMethod** | Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **int32** | Optional. | 
 **maxVideoBitDepth** | **int32** | Optional. The maximum video bit depth. | 
 **requireAvc** | **bool** | Optional. Whether to require avc. | 
 **deInterlace** | **bool** | Optional. Whether to deinterlace the video. | 
 **requireNonAnamorphic** | **bool** | Optional. Whether to require a non anamorphic stream. | 
 **transcodingMaxAudioChannels** | **int32** | Optional. The maximum number of audio channels to transcode. | 
 **cpuCoreLimit** | **int32** | Optional. The limit of how many cpu cores to use. | 
 **liveStreamId** | **string** | The live stream id. | 
 **enableMpegtsM2TsMode** | **bool** | Optional. Whether to enable the MpegtsM2Ts mode. | 
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 
 **enableAudioVbrEncoding** | **bool** | Optional. Whether to enable Audio Encoding. | [default to true]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/*, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadAudioStream

> *os.File HeadAudioStream(ctx, itemId).Container(container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()

Gets an audio stream.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sj14/jellyfin-go/api"
)

func main() {
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The item id.
	container := "container_example" // string | The audio container. (optional)
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. (optional)
	params := "params_example" // string | The streaming parameters. (optional)
	tag := "tag_example" // string | The tag. (optional)
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	playSessionId := "playSessionId_example" // string | The play session id. (optional)
	segmentContainer := "segmentContainer_example" // string | The segment container. (optional)
	segmentLength := int32(56) // int32 | The segment length. (optional)
	minSegments := int32(56) // int32 | The minimum number of segments. (optional)
	mediaSourceId := "mediaSourceId_example" // string | The media version id, if playing an alternate version. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify an audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	allowVideoStreamCopy := true // bool | Whether or not to allow copying of the video stream url. (optional)
	allowAudioStreamCopy := true // bool | Whether or not to allow copying of the audio stream url. (optional)
	breakOnNonKeyFrames := true // bool | Optional. Whether to break on non key frames. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100. (optional)
	maxAudioBitDepth := int32(56) // int32 | Optional. The maximum audio bit depth. (optional)
	audioBitRate := int32(56) // int32 | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. (optional)
	audioChannels := int32(56) // int32 | Optional. Specify a specific number of audio channels to encode to, e.g. 2. (optional)
	maxAudioChannels := int32(56) // int32 | Optional. Specify a maximum number of audio channels to encode to, e.g. 2. (optional)
	profile := "profile_example" // string | Optional. Specify a specific an encoder profile (varies by encoder), e.g. main, baseline, high. (optional)
	level := "level_example" // string | Optional. Specify a level for the encoder profile (varies by encoder), e.g. 3, 3.1. (optional)
	framerate := float32(3.4) // float32 | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	maxFramerate := float32(3.4) // float32 | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	copyTimestamps := true // bool | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms. (optional)
	width := int32(56) // int32 | Optional. The fixed horizontal resolution of the encoded video. (optional)
	height := int32(56) // int32 | Optional. The fixed vertical resolution of the encoded video. (optional)
	videoBitRate := int32(56) // int32 | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. (optional)
	subtitleStreamIndex := int32(56) // int32 | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. (optional)
	subtitleMethod := "subtitleMethod_example" // SubtitleDeliveryMethod | Optional. Specify the subtitle delivery method. (optional)
	maxRefFrames := int32(56) // int32 | Optional. (optional)
	maxVideoBitDepth := int32(56) // int32 | Optional. The maximum video bit depth. (optional)
	requireAvc := true // bool | Optional. Whether to require avc. (optional)
	deInterlace := true // bool | Optional. Whether to deinterlace the video. (optional)
	requireNonAnamorphic := true // bool | Optional. Whether to require a non anamorphic stream. (optional)
	transcodingMaxAudioChannels := int32(56) // int32 | Optional. The maximum number of audio channels to transcode. (optional)
	cpuCoreLimit := int32(56) // int32 | Optional. The limit of how many cpu cores to use. (optional)
	liveStreamId := "liveStreamId_example" // string | The live stream id. (optional)
	enableMpegtsM2TsMode := true // bool | Optional. Whether to enable the MpegtsM2Ts mode. (optional)
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)
	enableAudioVbrEncoding := true // bool | Optional. Whether to enable Audio Encoding. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.HeadAudioStream(context.Background(), itemId).Container(container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.HeadAudioStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadAudioStream`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.HeadAudioStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadAudioStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **container** | **string** | The audio container. | 
 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. | 
 **params** | **string** | The streaming parameters. | 
 **tag** | **string** | The tag. | 
 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **playSessionId** | **string** | The play session id. | 
 **segmentContainer** | **string** | The segment container. | 
 **segmentLength** | **int32** | The segment length. | 
 **minSegments** | **int32** | The minimum number of segments. | 
 **mediaSourceId** | **string** | The media version id, if playing an alternate version. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify an audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **allowVideoStreamCopy** | **bool** | Whether or not to allow copying of the video stream url. | 
 **allowAudioStreamCopy** | **bool** | Whether or not to allow copying of the audio stream url. | 
 **breakOnNonKeyFrames** | **bool** | Optional. Whether to break on non key frames. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100. | 
 **maxAudioBitDepth** | **int32** | Optional. The maximum audio bit depth. | 
 **audioBitRate** | **int32** | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **int32** | Optional. Specify a specific number of audio channels to encode to, e.g. 2. | 
 **maxAudioChannels** | **int32** | Optional. Specify a maximum number of audio channels to encode to, e.g. 2. | 
 **profile** | **string** | Optional. Specify a specific an encoder profile (varies by encoder), e.g. main, baseline, high. | 
 **level** | **string** | Optional. Specify a level for the encoder profile (varies by encoder), e.g. 3, 3.1. | 
 **framerate** | **float32** | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **float32** | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **bool** | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **width** | **int32** | Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **int32** | Optional. The fixed vertical resolution of the encoded video. | 
 **videoBitRate** | **int32** | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **int32** | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | **SubtitleDeliveryMethod** | Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **int32** | Optional. | 
 **maxVideoBitDepth** | **int32** | Optional. The maximum video bit depth. | 
 **requireAvc** | **bool** | Optional. Whether to require avc. | 
 **deInterlace** | **bool** | Optional. Whether to deinterlace the video. | 
 **requireNonAnamorphic** | **bool** | Optional. Whether to require a non anamorphic stream. | 
 **transcodingMaxAudioChannels** | **int32** | Optional. The maximum number of audio channels to transcode. | 
 **cpuCoreLimit** | **int32** | Optional. The limit of how many cpu cores to use. | 
 **liveStreamId** | **string** | The live stream id. | 
 **enableMpegtsM2TsMode** | **bool** | Optional. Whether to enable the MpegtsM2Ts mode. | 
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 
 **enableAudioVbrEncoding** | **bool** | Optional. Whether to enable Audio Encoding. | [default to true]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/*, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadAudioStreamByContainer

> *os.File HeadAudioStreamByContainer(ctx, itemId, container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()

Gets an audio stream.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sj14/jellyfin-go/api"
)

func main() {
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The item id.
	container := "container_example" // string | The audio container.
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. (optional)
	params := "params_example" // string | The streaming parameters. (optional)
	tag := "tag_example" // string | The tag. (optional)
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	playSessionId := "playSessionId_example" // string | The play session id. (optional)
	segmentContainer := "segmentContainer_example" // string | The segment container. (optional)
	segmentLength := int32(56) // int32 | The segment length. (optional)
	minSegments := int32(56) // int32 | The minimum number of segments. (optional)
	mediaSourceId := "mediaSourceId_example" // string | The media version id, if playing an alternate version. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify an audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	allowVideoStreamCopy := true // bool | Whether or not to allow copying of the video stream url. (optional)
	allowAudioStreamCopy := true // bool | Whether or not to allow copying of the audio stream url. (optional)
	breakOnNonKeyFrames := true // bool | Optional. Whether to break on non key frames. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100. (optional)
	maxAudioBitDepth := int32(56) // int32 | Optional. The maximum audio bit depth. (optional)
	audioBitRate := int32(56) // int32 | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. (optional)
	audioChannels := int32(56) // int32 | Optional. Specify a specific number of audio channels to encode to, e.g. 2. (optional)
	maxAudioChannels := int32(56) // int32 | Optional. Specify a maximum number of audio channels to encode to, e.g. 2. (optional)
	profile := "profile_example" // string | Optional. Specify a specific an encoder profile (varies by encoder), e.g. main, baseline, high. (optional)
	level := "level_example" // string | Optional. Specify a level for the encoder profile (varies by encoder), e.g. 3, 3.1. (optional)
	framerate := float32(3.4) // float32 | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	maxFramerate := float32(3.4) // float32 | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. (optional)
	copyTimestamps := true // bool | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms. (optional)
	width := int32(56) // int32 | Optional. The fixed horizontal resolution of the encoded video. (optional)
	height := int32(56) // int32 | Optional. The fixed vertical resolution of the encoded video. (optional)
	videoBitRate := int32(56) // int32 | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. (optional)
	subtitleStreamIndex := int32(56) // int32 | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. (optional)
	subtitleMethod := "subtitleMethod_example" // SubtitleDeliveryMethod | Optional. Specify the subtitle delivery method. (optional)
	maxRefFrames := int32(56) // int32 | Optional. (optional)
	maxVideoBitDepth := int32(56) // int32 | Optional. The maximum video bit depth. (optional)
	requireAvc := true // bool | Optional. Whether to require avc. (optional)
	deInterlace := true // bool | Optional. Whether to deinterlace the video. (optional)
	requireNonAnamorphic := true // bool | Optional. Whether to require a non anamorphic stream. (optional)
	transcodingMaxAudioChannels := int32(56) // int32 | Optional. The maximum number of audio channels to transcode. (optional)
	cpuCoreLimit := int32(56) // int32 | Optional. The limit of how many cpu cores to use. (optional)
	liveStreamId := "liveStreamId_example" // string | The live stream id. (optional)
	enableMpegtsM2TsMode := true // bool | Optional. Whether to enable the MpegtsM2Ts mode. (optional)
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)
	enableAudioVbrEncoding := true // bool | Optional. Whether to enable Audio Encoding. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AudioAPI.HeadAudioStreamByContainer(context.Background(), itemId, container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AudioAPI.HeadAudioStreamByContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadAudioStreamByContainer`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `AudioAPI.HeadAudioStreamByContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**container** | **string** | The audio container. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadAudioStreamByContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. | 
 **params** | **string** | The streaming parameters. | 
 **tag** | **string** | The tag. | 
 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **playSessionId** | **string** | The play session id. | 
 **segmentContainer** | **string** | The segment container. | 
 **segmentLength** | **int32** | The segment length. | 
 **minSegments** | **int32** | The minimum number of segments. | 
 **mediaSourceId** | **string** | The media version id, if playing an alternate version. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify an audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **allowVideoStreamCopy** | **bool** | Whether or not to allow copying of the video stream url. | 
 **allowAudioStreamCopy** | **bool** | Whether or not to allow copying of the audio stream url. | 
 **breakOnNonKeyFrames** | **bool** | Optional. Whether to break on non key frames. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100. | 
 **maxAudioBitDepth** | **int32** | Optional. The maximum audio bit depth. | 
 **audioBitRate** | **int32** | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **audioChannels** | **int32** | Optional. Specify a specific number of audio channels to encode to, e.g. 2. | 
 **maxAudioChannels** | **int32** | Optional. Specify a maximum number of audio channels to encode to, e.g. 2. | 
 **profile** | **string** | Optional. Specify a specific an encoder profile (varies by encoder), e.g. main, baseline, high. | 
 **level** | **string** | Optional. Specify a level for the encoder profile (varies by encoder), e.g. 3, 3.1. | 
 **framerate** | **float32** | Optional. A specific video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **maxFramerate** | **float32** | Optional. A specific maximum video framerate to encode to, e.g. 23.976. Generally this should be omitted unless the device has specific requirements. | 
 **copyTimestamps** | **bool** | Whether or not to copy timestamps when transcoding with an offset. Defaults to false. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **width** | **int32** | Optional. The fixed horizontal resolution of the encoded video. | 
 **height** | **int32** | Optional. The fixed vertical resolution of the encoded video. | 
 **videoBitRate** | **int32** | Optional. Specify a video bitrate to encode to, e.g. 500000. If omitted this will be left to encoder defaults. | 
 **subtitleStreamIndex** | **int32** | Optional. The index of the subtitle stream to use. If omitted no subtitles will be used. | 
 **subtitleMethod** | **SubtitleDeliveryMethod** | Optional. Specify the subtitle delivery method. | 
 **maxRefFrames** | **int32** | Optional. | 
 **maxVideoBitDepth** | **int32** | Optional. The maximum video bit depth. | 
 **requireAvc** | **bool** | Optional. Whether to require avc. | 
 **deInterlace** | **bool** | Optional. Whether to deinterlace the video. | 
 **requireNonAnamorphic** | **bool** | Optional. Whether to require a non anamorphic stream. | 
 **transcodingMaxAudioChannels** | **int32** | Optional. The maximum number of audio channels to transcode. | 
 **cpuCoreLimit** | **int32** | Optional. The limit of how many cpu cores to use. | 
 **liveStreamId** | **string** | The live stream id. | 
 **enableMpegtsM2TsMode** | **bool** | Optional. Whether to enable the MpegtsM2Ts mode. | 
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 
 **enableAudioVbrEncoding** | **bool** | Optional. Whether to enable Audio Encoding. | [default to true]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/*, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

