# \VideosAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAlternateSources**](VideosAPI.md#DeleteAlternateSources) | **Delete** /Videos/{itemId}/AlternateSources | Removes alternate video sources.
[**GetAdditionalPart**](VideosAPI.md#GetAdditionalPart) | **Get** /Videos/{itemId}/AdditionalParts | Gets additional parts for a video.
[**GetVideoStream**](VideosAPI.md#GetVideoStream) | **Get** /Videos/{itemId}/stream | Gets a video stream.
[**GetVideoStreamByContainer**](VideosAPI.md#GetVideoStreamByContainer) | **Get** /Videos/{itemId}/stream.{container} | Gets a video stream.
[**HeadVideoStream**](VideosAPI.md#HeadVideoStream) | **Head** /Videos/{itemId}/stream | Gets a video stream.
[**HeadVideoStreamByContainer**](VideosAPI.md#HeadVideoStreamByContainer) | **Head** /Videos/{itemId}/stream.{container} | Gets a video stream.
[**MergeVersions**](VideosAPI.md#MergeVersions) | **Post** /Videos/MergeVersions | Merges videos into a single record.



## DeleteAlternateSources

> DeleteAlternateSources(ctx, itemId).Execute()

Removes alternate video sources.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideosAPI.DeleteAlternateSources(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.DeleteAlternateSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlternateSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdditionalPart

> BaseItemDtoQueryResult GetAdditionalPart(ctx, itemId).UserId(userId).Execute()

Gets additional parts for a video.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user id, and attach user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.GetAdditionalPart(context.Background(), itemId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.GetAdditionalPart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdditionalPart`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.GetAdditionalPart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdditionalPartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data. | 

### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoStream

> *os.File GetVideoStream(ctx, itemId).Container(container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).Execute()

Gets a video stream.

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
	container := "container_example" // string | The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. (optional)
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
	audioCodec := "audioCodec_example" // string | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. Options: aac, mp3, vorbis, wma. (optional)
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
	maxWidth := int32(56) // int32 | Optional. The maximum horizontal resolution of the encoded video. (optional)
	maxHeight := int32(56) // int32 | Optional. The maximum vertical resolution of the encoded video. (optional)
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
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. Options: h265, h264, mpeg4, theora, vp8, vp9, vpx (deprecated), wmv. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.GetVideoStream(context.Background(), itemId).Container(container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.GetVideoStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideoStream`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.GetVideoStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **container** | **string** | The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. | 
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
 **audioCodec** | **string** | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
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
 **maxWidth** | **int32** | Optional. The maximum horizontal resolution of the encoded video. | 
 **maxHeight** | **int32** | Optional. The maximum vertical resolution of the encoded video. | 
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
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h265, h264, mpeg4, theora, vp8, vp9, vpx (deprecated), wmv. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideoStreamByContainer

> *os.File GetVideoStreamByContainer(ctx, itemId, container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).Execute()

Gets a video stream.

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
	container := "container_example" // string | The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv.
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
	audioCodec := "audioCodec_example" // string | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. Options: aac, mp3, vorbis, wma. (optional)
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
	maxWidth := int32(56) // int32 | Optional. The maximum horizontal resolution of the encoded video. (optional)
	maxHeight := int32(56) // int32 | Optional. The maximum vertical resolution of the encoded video. (optional)
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
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. Options: h265, h264, mpeg4, theora, vp8, vp9, vpx (deprecated), wmv. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.GetVideoStreamByContainer(context.Background(), itemId, container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.GetVideoStreamByContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVideoStreamByContainer`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.GetVideoStreamByContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**container** | **string** | The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideoStreamByContainerRequest struct via the builder pattern


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
 **audioCodec** | **string** | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
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
 **maxWidth** | **int32** | Optional. The maximum horizontal resolution of the encoded video. | 
 **maxHeight** | **int32** | Optional. The maximum vertical resolution of the encoded video. | 
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
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h265, h264, mpeg4, theora, vp8, vp9, vpx (deprecated), wmv. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadVideoStream

> *os.File HeadVideoStream(ctx, itemId).Container(container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).Execute()

Gets a video stream.

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
	container := "container_example" // string | The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. (optional)
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
	audioCodec := "audioCodec_example" // string | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. Options: aac, mp3, vorbis, wma. (optional)
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
	maxWidth := int32(56) // int32 | Optional. The maximum horizontal resolution of the encoded video. (optional)
	maxHeight := int32(56) // int32 | Optional. The maximum vertical resolution of the encoded video. (optional)
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
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. Options: h265, h264, mpeg4, theora, vp8, vp9, vpx (deprecated), wmv. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.HeadVideoStream(context.Background(), itemId).Container(container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.HeadVideoStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadVideoStream`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.HeadVideoStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadVideoStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **container** | **string** | The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. | 
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
 **audioCodec** | **string** | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
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
 **maxWidth** | **int32** | Optional. The maximum horizontal resolution of the encoded video. | 
 **maxHeight** | **int32** | Optional. The maximum vertical resolution of the encoded video. | 
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
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h265, h264, mpeg4, theora, vp8, vp9, vpx (deprecated), wmv. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadVideoStreamByContainer

> *os.File HeadVideoStreamByContainer(ctx, itemId, container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).Execute()

Gets a video stream.

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
	container := "container_example" // string | The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv.
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
	audioCodec := "audioCodec_example" // string | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url's extension. Options: aac, mp3, vorbis, wma. (optional)
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
	maxWidth := int32(56) // int32 | Optional. The maximum horizontal resolution of the encoded video. (optional)
	maxHeight := int32(56) // int32 | Optional. The maximum vertical resolution of the encoded video. (optional)
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
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url's extension. Options: h265, h264, mpeg4, theora, vp8, vp9, vpx (deprecated), wmv. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideosAPI.HeadVideoStreamByContainer(context.Background(), itemId, container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.HeadVideoStreamByContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadVideoStreamByContainer`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `VideosAPI.HeadVideoStreamByContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**container** | **string** | The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadVideoStreamByContainerRequest struct via the builder pattern


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
 **audioCodec** | **string** | Optional. Specify a audio codec to encode to, e.g. mp3. If omitted the server will auto-select using the url&#39;s extension. Options: aac, mp3, vorbis, wma. | 
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
 **maxWidth** | **int32** | Optional. The maximum horizontal resolution of the encoded video. | 
 **maxHeight** | **int32** | Optional. The maximum vertical resolution of the encoded video. | 
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
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. If omitted the server will auto-select using the url&#39;s extension. Options: h265, h264, mpeg4, theora, vp8, vp9, vpx (deprecated), wmv. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeVersions

> MergeVersions(ctx).Ids(ids).Execute()

Merges videos into a single record.

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
	ids := []string{"Inner_example"} // []string | Item id list. This allows multiple, comma delimited.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VideosAPI.MergeVersions(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideosAPI.MergeVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMergeVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | Item id list. This allows multiple, comma delimited. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

