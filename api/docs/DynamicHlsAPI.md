# \DynamicHlsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHlsAudioSegment**](DynamicHlsAPI.md#GetHlsAudioSegment) | **Get** /Audio/{itemId}/hls1/{playlistId}/{segmentId}.{container} | Gets a video stream using HTTP live streaming.
[**GetHlsVideoSegment**](DynamicHlsAPI.md#GetHlsVideoSegment) | **Get** /Videos/{itemId}/hls1/{playlistId}/{segmentId}.{container} | Gets a video stream using HTTP live streaming.
[**GetLiveHlsStream**](DynamicHlsAPI.md#GetLiveHlsStream) | **Get** /Videos/{itemId}/live.m3u8 | Gets a hls live stream.
[**GetMasterHlsAudioPlaylist**](DynamicHlsAPI.md#GetMasterHlsAudioPlaylist) | **Get** /Audio/{itemId}/master.m3u8 | Gets an audio hls playlist stream.
[**GetMasterHlsVideoPlaylist**](DynamicHlsAPI.md#GetMasterHlsVideoPlaylist) | **Get** /Videos/{itemId}/master.m3u8 | Gets a video hls playlist stream.
[**GetVariantHlsAudioPlaylist**](DynamicHlsAPI.md#GetVariantHlsAudioPlaylist) | **Get** /Audio/{itemId}/main.m3u8 | Gets an audio stream using HTTP live streaming.
[**GetVariantHlsVideoPlaylist**](DynamicHlsAPI.md#GetVariantHlsVideoPlaylist) | **Get** /Videos/{itemId}/main.m3u8 | Gets a video stream using HTTP live streaming.
[**HeadMasterHlsAudioPlaylist**](DynamicHlsAPI.md#HeadMasterHlsAudioPlaylist) | **Head** /Audio/{itemId}/master.m3u8 | Gets an audio hls playlist stream.
[**HeadMasterHlsVideoPlaylist**](DynamicHlsAPI.md#HeadMasterHlsVideoPlaylist) | **Head** /Videos/{itemId}/master.m3u8 | Gets a video hls playlist stream.



## GetHlsAudioSegment

> *os.File GetHlsAudioSegment(ctx, itemId, playlistId, segmentId, container).RuntimeTicks(runtimeTicks).ActualSegmentLengthTicks(actualSegmentLengthTicks).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).MaxStreamingBitrate(maxStreamingBitrate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()

Gets a video stream using HTTP live streaming.

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
	playlistId := "playlistId_example" // string | The playlist id.
	segmentId := int32(56) // int32 | The segment id.
	container := "container_example" // string | The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv.
	runtimeTicks := int64(789) // int64 | The position of the requested segment in ticks.
	actualSegmentLengthTicks := int64(789) // int64 | The length of the requested segment in ticks.
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
	audioCodec := "audioCodec_example" // string | Optional. Specify an audio codec to encode to, e.g. mp3. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	allowVideoStreamCopy := true // bool | Whether or not to allow copying of the video stream url. (optional)
	allowAudioStreamCopy := true // bool | Whether or not to allow copying of the audio stream url. (optional)
	breakOnNonKeyFrames := true // bool | Optional. Whether to break on non key frames. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100. (optional)
	maxAudioBitDepth := int32(56) // int32 | Optional. The maximum audio bit depth. (optional)
	maxStreamingBitrate := int32(56) // int32 | Optional. The maximum streaming bitrate. (optional)
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
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)
	enableAudioVbrEncoding := true // bool | Optional. Whether to enable Audio Encoding. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicHlsAPI.GetHlsAudioSegment(context.Background(), itemId, playlistId, segmentId, container).RuntimeTicks(runtimeTicks).ActualSegmentLengthTicks(actualSegmentLengthTicks).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).MaxStreamingBitrate(maxStreamingBitrate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicHlsAPI.GetHlsAudioSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHlsAudioSegment`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DynamicHlsAPI.GetHlsAudioSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**playlistId** | **string** | The playlist id. | 
**segmentId** | **int32** | The segment id. | 
**container** | **string** | The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHlsAudioSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **runtimeTicks** | **int64** | The position of the requested segment in ticks. | 
 **actualSegmentLengthTicks** | **int64** | The length of the requested segment in ticks. | 
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
 **audioCodec** | **string** | Optional. Specify an audio codec to encode to, e.g. mp3. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **allowVideoStreamCopy** | **bool** | Whether or not to allow copying of the video stream url. | 
 **allowAudioStreamCopy** | **bool** | Whether or not to allow copying of the audio stream url. | 
 **breakOnNonKeyFrames** | **bool** | Optional. Whether to break on non key frames. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100. | 
 **maxAudioBitDepth** | **int32** | Optional. The maximum audio bit depth. | 
 **maxStreamingBitrate** | **int32** | Optional. The maximum streaming bitrate. | 
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
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. | 
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

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/*, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHlsVideoSegment

> *os.File GetHlsVideoSegment(ctx, itemId, playlistId, segmentId, container).RuntimeTicks(runtimeTicks).ActualSegmentLengthTicks(actualSegmentLengthTicks).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).AlwaysBurnInSubtitleWhenTranscoding(alwaysBurnInSubtitleWhenTranscoding).Execute()

Gets a video stream using HTTP live streaming.

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
	playlistId := "playlistId_example" // string | The playlist id.
	segmentId := int32(56) // int32 | The segment id.
	container := "container_example" // string | The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv.
	runtimeTicks := int64(789) // int64 | The position of the requested segment in ticks.
	actualSegmentLengthTicks := int64(789) // int64 | The length of the requested segment in ticks.
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. (optional)
	params := "params_example" // string | The streaming parameters. (optional)
	tag := "tag_example" // string | The tag. (optional)
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	playSessionId := "playSessionId_example" // string | The play session id. (optional)
	segmentContainer := "segmentContainer_example" // string | The segment container. (optional)
	segmentLength := int32(56) // int32 | The desired segment length. (optional)
	minSegments := int32(56) // int32 | The minimum number of segments. (optional)
	mediaSourceId := "mediaSourceId_example" // string | The media version id, if playing an alternate version. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify an audio codec to encode to, e.g. mp3. (optional)
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
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)
	enableAudioVbrEncoding := true // bool | Optional. Whether to enable Audio Encoding. (optional) (default to true)
	alwaysBurnInSubtitleWhenTranscoding := true // bool | Whether to always burn in subtitles when transcoding. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicHlsAPI.GetHlsVideoSegment(context.Background(), itemId, playlistId, segmentId, container).RuntimeTicks(runtimeTicks).ActualSegmentLengthTicks(actualSegmentLengthTicks).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).AlwaysBurnInSubtitleWhenTranscoding(alwaysBurnInSubtitleWhenTranscoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicHlsAPI.GetHlsVideoSegment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHlsVideoSegment`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DynamicHlsAPI.GetHlsVideoSegment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**playlistId** | **string** | The playlist id. | 
**segmentId** | **int32** | The segment id. | 
**container** | **string** | The video container. Possible values are: ts, webm, asf, wmv, ogv, mp4, m4v, mkv, mpeg, mpg, avi, 3gp, wmv, wtv, m2ts, mov, iso, flv. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHlsVideoSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **runtimeTicks** | **int64** | The position of the requested segment in ticks. | 
 **actualSegmentLengthTicks** | **int64** | The length of the requested segment in ticks. | 
 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. | 
 **params** | **string** | The streaming parameters. | 
 **tag** | **string** | The tag. | 
 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **playSessionId** | **string** | The play session id. | 
 **segmentContainer** | **string** | The segment container. | 
 **segmentLength** | **int32** | The desired segment length. | 
 **minSegments** | **int32** | The minimum number of segments. | 
 **mediaSourceId** | **string** | The media version id, if playing an alternate version. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify an audio codec to encode to, e.g. mp3. | 
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
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 
 **enableAudioVbrEncoding** | **bool** | Optional. Whether to enable Audio Encoding. | [default to true]
 **alwaysBurnInSubtitleWhenTranscoding** | **bool** | Whether to always burn in subtitles when transcoding. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/*, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLiveHlsStream

> *os.File GetLiveHlsStream(ctx, itemId).Container(container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).MaxWidth(maxWidth).MaxHeight(maxHeight).EnableSubtitlesInManifest(enableSubtitlesInManifest).EnableAudioVbrEncoding(enableAudioVbrEncoding).AlwaysBurnInSubtitleWhenTranscoding(alwaysBurnInSubtitleWhenTranscoding).Execute()

Gets a hls live stream.

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
	audioCodec := "audioCodec_example" // string | Optional. Specify an audio codec to encode to, e.g. mp3. (optional)
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
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)
	maxWidth := int32(56) // int32 | Optional. The max width. (optional)
	maxHeight := int32(56) // int32 | Optional. The max height. (optional)
	enableSubtitlesInManifest := true // bool | Optional. Whether to enable subtitles in the manifest. (optional)
	enableAudioVbrEncoding := true // bool | Optional. Whether to enable Audio Encoding. (optional) (default to true)
	alwaysBurnInSubtitleWhenTranscoding := true // bool | Whether to always burn in subtitles when transcoding. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicHlsAPI.GetLiveHlsStream(context.Background(), itemId).Container(container).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).MaxWidth(maxWidth).MaxHeight(maxHeight).EnableSubtitlesInManifest(enableSubtitlesInManifest).EnableAudioVbrEncoding(enableAudioVbrEncoding).AlwaysBurnInSubtitleWhenTranscoding(alwaysBurnInSubtitleWhenTranscoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicHlsAPI.GetLiveHlsStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLiveHlsStream`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DynamicHlsAPI.GetLiveHlsStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveHlsStreamRequest struct via the builder pattern


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
 **audioCodec** | **string** | Optional. Specify an audio codec to encode to, e.g. mp3. | 
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
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 
 **maxWidth** | **int32** | Optional. The max width. | 
 **maxHeight** | **int32** | Optional. The max height. | 
 **enableSubtitlesInManifest** | **bool** | Optional. Whether to enable subtitles in the manifest. | 
 **enableAudioVbrEncoding** | **bool** | Optional. Whether to enable Audio Encoding. | [default to true]
 **alwaysBurnInSubtitleWhenTranscoding** | **bool** | Whether to always burn in subtitles when transcoding. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-mpegURL, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMasterHlsAudioPlaylist

> *os.File GetMasterHlsAudioPlaylist(ctx, itemId).MediaSourceId(mediaSourceId).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).MaxStreamingBitrate(maxStreamingBitrate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAdaptiveBitrateStreaming(enableAdaptiveBitrateStreaming).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()

Gets an audio hls playlist stream.

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
	mediaSourceId := "mediaSourceId_example" // string | The media version id, if playing an alternate version.
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. (optional)
	params := "params_example" // string | The streaming parameters. (optional)
	tag := "tag_example" // string | The tag. (optional)
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	playSessionId := "playSessionId_example" // string | The play session id. (optional)
	segmentContainer := "segmentContainer_example" // string | The segment container. (optional)
	segmentLength := int32(56) // int32 | The segment length. (optional)
	minSegments := int32(56) // int32 | The minimum number of segments. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify an audio codec to encode to, e.g. mp3. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	allowVideoStreamCopy := true // bool | Whether or not to allow copying of the video stream url. (optional)
	allowAudioStreamCopy := true // bool | Whether or not to allow copying of the audio stream url. (optional)
	breakOnNonKeyFrames := true // bool | Optional. Whether to break on non key frames. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100. (optional)
	maxAudioBitDepth := int32(56) // int32 | Optional. The maximum audio bit depth. (optional)
	maxStreamingBitrate := int32(56) // int32 | Optional. The maximum streaming bitrate. (optional)
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
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)
	enableAdaptiveBitrateStreaming := true // bool | Enable adaptive bitrate streaming. (optional) (default to false)
	enableAudioVbrEncoding := true // bool | Optional. Whether to enable Audio Encoding. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicHlsAPI.GetMasterHlsAudioPlaylist(context.Background(), itemId).MediaSourceId(mediaSourceId).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).MaxStreamingBitrate(maxStreamingBitrate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAdaptiveBitrateStreaming(enableAdaptiveBitrateStreaming).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicHlsAPI.GetMasterHlsAudioPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMasterHlsAudioPlaylist`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DynamicHlsAPI.GetMasterHlsAudioPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMasterHlsAudioPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaSourceId** | **string** | The media version id, if playing an alternate version. | 
 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. | 
 **params** | **string** | The streaming parameters. | 
 **tag** | **string** | The tag. | 
 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **playSessionId** | **string** | The play session id. | 
 **segmentContainer** | **string** | The segment container. | 
 **segmentLength** | **int32** | The segment length. | 
 **minSegments** | **int32** | The minimum number of segments. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify an audio codec to encode to, e.g. mp3. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **allowVideoStreamCopy** | **bool** | Whether or not to allow copying of the video stream url. | 
 **allowAudioStreamCopy** | **bool** | Whether or not to allow copying of the audio stream url. | 
 **breakOnNonKeyFrames** | **bool** | Optional. Whether to break on non key frames. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100. | 
 **maxAudioBitDepth** | **int32** | Optional. The maximum audio bit depth. | 
 **maxStreamingBitrate** | **int32** | Optional. The maximum streaming bitrate. | 
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
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 
 **enableAdaptiveBitrateStreaming** | **bool** | Enable adaptive bitrate streaming. | [default to false]
 **enableAudioVbrEncoding** | **bool** | Optional. Whether to enable Audio Encoding. | [default to true]

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-mpegURL, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMasterHlsVideoPlaylist

> *os.File GetMasterHlsVideoPlaylist(ctx, itemId).MediaSourceId(mediaSourceId).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAdaptiveBitrateStreaming(enableAdaptiveBitrateStreaming).EnableTrickplay(enableTrickplay).EnableAudioVbrEncoding(enableAudioVbrEncoding).AlwaysBurnInSubtitleWhenTranscoding(alwaysBurnInSubtitleWhenTranscoding).Execute()

Gets a video hls playlist stream.

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
	mediaSourceId := "mediaSourceId_example" // string | The media version id, if playing an alternate version.
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. (optional)
	params := "params_example" // string | The streaming parameters. (optional)
	tag := "tag_example" // string | The tag. (optional)
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	playSessionId := "playSessionId_example" // string | The play session id. (optional)
	segmentContainer := "segmentContainer_example" // string | The segment container. (optional)
	segmentLength := int32(56) // int32 | The segment length. (optional)
	minSegments := int32(56) // int32 | The minimum number of segments. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify an audio codec to encode to, e.g. mp3. (optional)
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
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)
	enableAdaptiveBitrateStreaming := true // bool | Enable adaptive bitrate streaming. (optional) (default to false)
	enableTrickplay := true // bool | Enable trickplay image playlists being added to master playlist. (optional) (default to true)
	enableAudioVbrEncoding := true // bool | Whether to enable Audio Encoding. (optional) (default to true)
	alwaysBurnInSubtitleWhenTranscoding := true // bool | Whether to always burn in subtitles when transcoding. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicHlsAPI.GetMasterHlsVideoPlaylist(context.Background(), itemId).MediaSourceId(mediaSourceId).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAdaptiveBitrateStreaming(enableAdaptiveBitrateStreaming).EnableTrickplay(enableTrickplay).EnableAudioVbrEncoding(enableAudioVbrEncoding).AlwaysBurnInSubtitleWhenTranscoding(alwaysBurnInSubtitleWhenTranscoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicHlsAPI.GetMasterHlsVideoPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMasterHlsVideoPlaylist`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DynamicHlsAPI.GetMasterHlsVideoPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMasterHlsVideoPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaSourceId** | **string** | The media version id, if playing an alternate version. | 
 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. | 
 **params** | **string** | The streaming parameters. | 
 **tag** | **string** | The tag. | 
 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **playSessionId** | **string** | The play session id. | 
 **segmentContainer** | **string** | The segment container. | 
 **segmentLength** | **int32** | The segment length. | 
 **minSegments** | **int32** | The minimum number of segments. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify an audio codec to encode to, e.g. mp3. | 
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
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 
 **enableAdaptiveBitrateStreaming** | **bool** | Enable adaptive bitrate streaming. | [default to false]
 **enableTrickplay** | **bool** | Enable trickplay image playlists being added to master playlist. | [default to true]
 **enableAudioVbrEncoding** | **bool** | Whether to enable Audio Encoding. | [default to true]
 **alwaysBurnInSubtitleWhenTranscoding** | **bool** | Whether to always burn in subtitles when transcoding. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-mpegURL, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariantHlsAudioPlaylist

> *os.File GetVariantHlsAudioPlaylist(ctx, itemId).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).MaxStreamingBitrate(maxStreamingBitrate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()

Gets an audio stream using HTTP live streaming.

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
	audioCodec := "audioCodec_example" // string | Optional. Specify an audio codec to encode to, e.g. mp3. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	allowVideoStreamCopy := true // bool | Whether or not to allow copying of the video stream url. (optional)
	allowAudioStreamCopy := true // bool | Whether or not to allow copying of the audio stream url. (optional)
	breakOnNonKeyFrames := true // bool | Optional. Whether to break on non key frames. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100. (optional)
	maxAudioBitDepth := int32(56) // int32 | Optional. The maximum audio bit depth. (optional)
	maxStreamingBitrate := int32(56) // int32 | Optional. The maximum streaming bitrate. (optional)
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
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)
	enableAudioVbrEncoding := true // bool | Optional. Whether to enable Audio Encoding. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicHlsAPI.GetVariantHlsAudioPlaylist(context.Background(), itemId).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).MaxStreamingBitrate(maxStreamingBitrate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicHlsAPI.GetVariantHlsAudioPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVariantHlsAudioPlaylist`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DynamicHlsAPI.GetVariantHlsAudioPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariantHlsAudioPlaylistRequest struct via the builder pattern


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
 **audioCodec** | **string** | Optional. Specify an audio codec to encode to, e.g. mp3. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **allowVideoStreamCopy** | **bool** | Whether or not to allow copying of the video stream url. | 
 **allowAudioStreamCopy** | **bool** | Whether or not to allow copying of the audio stream url. | 
 **breakOnNonKeyFrames** | **bool** | Optional. Whether to break on non key frames. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100. | 
 **maxAudioBitDepth** | **int32** | Optional. The maximum audio bit depth. | 
 **maxStreamingBitrate** | **int32** | Optional. The maximum streaming bitrate. | 
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
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. | 
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

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-mpegURL, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariantHlsVideoPlaylist

> *os.File GetVariantHlsVideoPlaylist(ctx, itemId).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).AlwaysBurnInSubtitleWhenTranscoding(alwaysBurnInSubtitleWhenTranscoding).Execute()

Gets a video stream using HTTP live streaming.

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
	audioCodec := "audioCodec_example" // string | Optional. Specify an audio codec to encode to, e.g. mp3. (optional)
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
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)
	enableAudioVbrEncoding := true // bool | Optional. Whether to enable Audio Encoding. (optional) (default to true)
	alwaysBurnInSubtitleWhenTranscoding := true // bool | Whether to always burn in subtitles when transcoding. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicHlsAPI.GetVariantHlsVideoPlaylist(context.Background(), itemId).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).MediaSourceId(mediaSourceId).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAudioVbrEncoding(enableAudioVbrEncoding).AlwaysBurnInSubtitleWhenTranscoding(alwaysBurnInSubtitleWhenTranscoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicHlsAPI.GetVariantHlsVideoPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVariantHlsVideoPlaylist`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DynamicHlsAPI.GetVariantHlsVideoPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariantHlsVideoPlaylistRequest struct via the builder pattern


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
 **audioCodec** | **string** | Optional. Specify an audio codec to encode to, e.g. mp3. | 
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
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 
 **enableAudioVbrEncoding** | **bool** | Optional. Whether to enable Audio Encoding. | [default to true]
 **alwaysBurnInSubtitleWhenTranscoding** | **bool** | Whether to always burn in subtitles when transcoding. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-mpegURL, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadMasterHlsAudioPlaylist

> *os.File HeadMasterHlsAudioPlaylist(ctx, itemId).MediaSourceId(mediaSourceId).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).MaxStreamingBitrate(maxStreamingBitrate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAdaptiveBitrateStreaming(enableAdaptiveBitrateStreaming).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()

Gets an audio hls playlist stream.

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
	mediaSourceId := "mediaSourceId_example" // string | The media version id, if playing an alternate version.
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. (optional)
	params := "params_example" // string | The streaming parameters. (optional)
	tag := "tag_example" // string | The tag. (optional)
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	playSessionId := "playSessionId_example" // string | The play session id. (optional)
	segmentContainer := "segmentContainer_example" // string | The segment container. (optional)
	segmentLength := int32(56) // int32 | The segment length. (optional)
	minSegments := int32(56) // int32 | The minimum number of segments. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify an audio codec to encode to, e.g. mp3. (optional)
	enableAutoStreamCopy := true // bool | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. (optional)
	allowVideoStreamCopy := true // bool | Whether or not to allow copying of the video stream url. (optional)
	allowAudioStreamCopy := true // bool | Whether or not to allow copying of the audio stream url. (optional)
	breakOnNonKeyFrames := true // bool | Optional. Whether to break on non key frames. (optional)
	audioSampleRate := int32(56) // int32 | Optional. Specify a specific audio sample rate, e.g. 44100. (optional)
	maxAudioBitDepth := int32(56) // int32 | Optional. The maximum audio bit depth. (optional)
	maxStreamingBitrate := int32(56) // int32 | Optional. The maximum streaming bitrate. (optional)
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
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)
	enableAdaptiveBitrateStreaming := true // bool | Enable adaptive bitrate streaming. (optional) (default to false)
	enableAudioVbrEncoding := true // bool | Optional. Whether to enable Audio Encoding. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicHlsAPI.HeadMasterHlsAudioPlaylist(context.Background(), itemId).MediaSourceId(mediaSourceId).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).MaxStreamingBitrate(maxStreamingBitrate).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAdaptiveBitrateStreaming(enableAdaptiveBitrateStreaming).EnableAudioVbrEncoding(enableAudioVbrEncoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicHlsAPI.HeadMasterHlsAudioPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadMasterHlsAudioPlaylist`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DynamicHlsAPI.HeadMasterHlsAudioPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadMasterHlsAudioPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaSourceId** | **string** | The media version id, if playing an alternate version. | 
 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. | 
 **params** | **string** | The streaming parameters. | 
 **tag** | **string** | The tag. | 
 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **playSessionId** | **string** | The play session id. | 
 **segmentContainer** | **string** | The segment container. | 
 **segmentLength** | **int32** | The segment length. | 
 **minSegments** | **int32** | The minimum number of segments. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify an audio codec to encode to, e.g. mp3. | 
 **enableAutoStreamCopy** | **bool** | Whether or not to allow automatic stream copy if requested values match the original source. Defaults to true. | 
 **allowVideoStreamCopy** | **bool** | Whether or not to allow copying of the video stream url. | 
 **allowAudioStreamCopy** | **bool** | Whether or not to allow copying of the audio stream url. | 
 **breakOnNonKeyFrames** | **bool** | Optional. Whether to break on non key frames. | 
 **audioSampleRate** | **int32** | Optional. Specify a specific audio sample rate, e.g. 44100. | 
 **maxAudioBitDepth** | **int32** | Optional. The maximum audio bit depth. | 
 **maxStreamingBitrate** | **int32** | Optional. The maximum streaming bitrate. | 
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
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 
 **enableAdaptiveBitrateStreaming** | **bool** | Enable adaptive bitrate streaming. | [default to false]
 **enableAudioVbrEncoding** | **bool** | Optional. Whether to enable Audio Encoding. | [default to true]

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-mpegURL, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadMasterHlsVideoPlaylist

> *os.File HeadMasterHlsVideoPlaylist(ctx, itemId).MediaSourceId(mediaSourceId).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAdaptiveBitrateStreaming(enableAdaptiveBitrateStreaming).EnableTrickplay(enableTrickplay).EnableAudioVbrEncoding(enableAudioVbrEncoding).AlwaysBurnInSubtitleWhenTranscoding(alwaysBurnInSubtitleWhenTranscoding).Execute()

Gets a video hls playlist stream.

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
	mediaSourceId := "mediaSourceId_example" // string | The media version id, if playing an alternate version.
	static := true // bool | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. (optional)
	params := "params_example" // string | The streaming parameters. (optional)
	tag := "tag_example" // string | The tag. (optional)
	deviceProfileId := "deviceProfileId_example" // string | Optional. The dlna device profile id to utilize. (optional)
	playSessionId := "playSessionId_example" // string | The play session id. (optional)
	segmentContainer := "segmentContainer_example" // string | The segment container. (optional)
	segmentLength := int32(56) // int32 | The segment length. (optional)
	minSegments := int32(56) // int32 | The minimum number of segments. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	audioCodec := "audioCodec_example" // string | Optional. Specify an audio codec to encode to, e.g. mp3. (optional)
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
	videoCodec := "videoCodec_example" // string | Optional. Specify a video codec to encode to, e.g. h264. (optional)
	subtitleCodec := "subtitleCodec_example" // string | Optional. Specify a subtitle codec to encode to. (optional)
	transcodeReasons := "transcodeReasons_example" // string | Optional. The transcoding reason. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. (optional)
	videoStreamIndex := int32(56) // int32 | Optional. The index of the video stream to use. If omitted the first video stream will be used. (optional)
	context := "context_example" // EncodingContext | Optional. The MediaBrowser.Model.Dlna.EncodingContext. (optional)
	streamOptions := map[string]string{"key": "Inner_example"} // map[string]string | Optional. The streaming options. (optional)
	enableAdaptiveBitrateStreaming := true // bool | Enable adaptive bitrate streaming. (optional) (default to false)
	enableTrickplay := true // bool | Enable trickplay image playlists being added to master playlist. (optional) (default to true)
	enableAudioVbrEncoding := true // bool | Whether to enable Audio Encoding. (optional) (default to true)
	alwaysBurnInSubtitleWhenTranscoding := true // bool | Whether to always burn in subtitles when transcoding. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DynamicHlsAPI.HeadMasterHlsVideoPlaylist(context.Background(), itemId).MediaSourceId(mediaSourceId).Static(static).Params(params).Tag(tag).DeviceProfileId(deviceProfileId).PlaySessionId(playSessionId).SegmentContainer(segmentContainer).SegmentLength(segmentLength).MinSegments(minSegments).DeviceId(deviceId).AudioCodec(audioCodec).EnableAutoStreamCopy(enableAutoStreamCopy).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).BreakOnNonKeyFrames(breakOnNonKeyFrames).AudioSampleRate(audioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).AudioBitRate(audioBitRate).AudioChannels(audioChannels).MaxAudioChannels(maxAudioChannels).Profile(profile).Level(level).Framerate(framerate).MaxFramerate(maxFramerate).CopyTimestamps(copyTimestamps).StartTimeTicks(startTimeTicks).Width(width).Height(height).MaxWidth(maxWidth).MaxHeight(maxHeight).VideoBitRate(videoBitRate).SubtitleStreamIndex(subtitleStreamIndex).SubtitleMethod(subtitleMethod).MaxRefFrames(maxRefFrames).MaxVideoBitDepth(maxVideoBitDepth).RequireAvc(requireAvc).DeInterlace(deInterlace).RequireNonAnamorphic(requireNonAnamorphic).TranscodingMaxAudioChannels(transcodingMaxAudioChannels).CpuCoreLimit(cpuCoreLimit).LiveStreamId(liveStreamId).EnableMpegtsM2TsMode(enableMpegtsM2TsMode).VideoCodec(videoCodec).SubtitleCodec(subtitleCodec).TranscodeReasons(transcodeReasons).AudioStreamIndex(audioStreamIndex).VideoStreamIndex(videoStreamIndex).Context(context).StreamOptions(streamOptions).EnableAdaptiveBitrateStreaming(enableAdaptiveBitrateStreaming).EnableTrickplay(enableTrickplay).EnableAudioVbrEncoding(enableAudioVbrEncoding).AlwaysBurnInSubtitleWhenTranscoding(alwaysBurnInSubtitleWhenTranscoding).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicHlsAPI.HeadMasterHlsVideoPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadMasterHlsVideoPlaylist`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DynamicHlsAPI.HeadMasterHlsVideoPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadMasterHlsVideoPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaSourceId** | **string** | The media version id, if playing an alternate version. | 
 **static** | **bool** | Optional. If true, the original file will be streamed statically without any encoding. Use either no url extension or the original file extension. true/false. | 
 **params** | **string** | The streaming parameters. | 
 **tag** | **string** | The tag. | 
 **deviceProfileId** | **string** | Optional. The dlna device profile id to utilize. | 
 **playSessionId** | **string** | The play session id. | 
 **segmentContainer** | **string** | The segment container. | 
 **segmentLength** | **int32** | The segment length. | 
 **minSegments** | **int32** | The minimum number of segments. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **audioCodec** | **string** | Optional. Specify an audio codec to encode to, e.g. mp3. | 
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
 **videoCodec** | **string** | Optional. Specify a video codec to encode to, e.g. h264. | 
 **subtitleCodec** | **string** | Optional. Specify a subtitle codec to encode to. | 
 **transcodeReasons** | **string** | Optional. The transcoding reason. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to use. If omitted the first audio stream will be used. | 
 **videoStreamIndex** | **int32** | Optional. The index of the video stream to use. If omitted the first video stream will be used. | 
 **context** | **EncodingContext** | Optional. The MediaBrowser.Model.Dlna.EncodingContext. | 
 **streamOptions** | **map[string]string** | Optional. The streaming options. | 
 **enableAdaptiveBitrateStreaming** | **bool** | Enable adaptive bitrate streaming. | [default to false]
 **enableTrickplay** | **bool** | Enable trickplay image playlists being added to master playlist. | [default to true]
 **enableAudioVbrEncoding** | **bool** | Whether to enable Audio Encoding. | [default to true]
 **alwaysBurnInSubtitleWhenTranscoding** | **bool** | Whether to always burn in subtitles when transcoding. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-mpegURL, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

