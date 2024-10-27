# \UniversalAudioAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUniversalAudioStream**](UniversalAudioAPI.md#GetUniversalAudioStream) | **Get** /Audio/{itemId}/universal | Gets an audio stream.
[**HeadUniversalAudioStream**](UniversalAudioAPI.md#HeadUniversalAudioStream) | **Head** /Audio/{itemId}/universal | Gets an audio stream.



## GetUniversalAudioStream

> *os.File GetUniversalAudioStream(ctx, itemId).Container(container).MediaSourceId(mediaSourceId).DeviceId(deviceId).UserId(userId).AudioCodec(audioCodec).MaxAudioChannels(maxAudioChannels).TranscodingAudioChannels(transcodingAudioChannels).MaxStreamingBitrate(maxStreamingBitrate).AudioBitRate(audioBitRate).StartTimeTicks(startTimeTicks).TranscodingContainer(transcodingContainer).TranscodingProtocol(transcodingProtocol).MaxAudioSampleRate(maxAudioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).EnableRemoteMedia(enableRemoteMedia).EnableAudioVbrEncoding(enableAudioVbrEncoding).BreakOnNonKeyFrames(breakOnNonKeyFrames).EnableRedirection(enableRedirection).Execute()

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
	container := []string{"Inner_example"} // []string | Optional. The audio container. (optional)
	mediaSourceId := "mediaSourceId_example" // string | The media version id, if playing an alternate version. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. The user id. (optional)
	audioCodec := "audioCodec_example" // string | Optional. The audio codec to transcode to. (optional)
	maxAudioChannels := int32(56) // int32 | Optional. The maximum number of audio channels. (optional)
	transcodingAudioChannels := int32(56) // int32 | Optional. The number of how many audio channels to transcode to. (optional)
	maxStreamingBitrate := int32(56) // int32 | Optional. The maximum streaming bitrate. (optional)
	audioBitRate := int32(56) // int32 | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms. (optional)
	transcodingContainer := "transcodingContainer_example" // string | Optional. The container to transcode to. (optional)
	transcodingProtocol := "transcodingProtocol_example" // MediaStreamProtocol | Optional. The transcoding protocol. (optional)
	maxAudioSampleRate := int32(56) // int32 | Optional. The maximum audio sample rate. (optional)
	maxAudioBitDepth := int32(56) // int32 | Optional. The maximum audio bit depth. (optional)
	enableRemoteMedia := true // bool | Optional. Whether to enable remote media. (optional)
	enableAudioVbrEncoding := true // bool | Optional. Whether to enable Audio Encoding. (optional) (default to true)
	breakOnNonKeyFrames := true // bool | Optional. Whether to break on non key frames. (optional) (default to false)
	enableRedirection := true // bool | Whether to enable redirection. Defaults to true. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniversalAudioAPI.GetUniversalAudioStream(context.Background(), itemId).Container(container).MediaSourceId(mediaSourceId).DeviceId(deviceId).UserId(userId).AudioCodec(audioCodec).MaxAudioChannels(maxAudioChannels).TranscodingAudioChannels(transcodingAudioChannels).MaxStreamingBitrate(maxStreamingBitrate).AudioBitRate(audioBitRate).StartTimeTicks(startTimeTicks).TranscodingContainer(transcodingContainer).TranscodingProtocol(transcodingProtocol).MaxAudioSampleRate(maxAudioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).EnableRemoteMedia(enableRemoteMedia).EnableAudioVbrEncoding(enableAudioVbrEncoding).BreakOnNonKeyFrames(breakOnNonKeyFrames).EnableRedirection(enableRedirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniversalAudioAPI.GetUniversalAudioStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUniversalAudioStream`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `UniversalAudioAPI.GetUniversalAudioStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUniversalAudioStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **container** | **[]string** | Optional. The audio container. | 
 **mediaSourceId** | **string** | The media version id, if playing an alternate version. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **userId** | **string** | Optional. The user id. | 
 **audioCodec** | **string** | Optional. The audio codec to transcode to. | 
 **maxAudioChannels** | **int32** | Optional. The maximum number of audio channels. | 
 **transcodingAudioChannels** | **int32** | Optional. The number of how many audio channels to transcode to. | 
 **maxStreamingBitrate** | **int32** | Optional. The maximum streaming bitrate. | 
 **audioBitRate** | **int32** | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **transcodingContainer** | **string** | Optional. The container to transcode to. | 
 **transcodingProtocol** | **MediaStreamProtocol** | Optional. The transcoding protocol. | 
 **maxAudioSampleRate** | **int32** | Optional. The maximum audio sample rate. | 
 **maxAudioBitDepth** | **int32** | Optional. The maximum audio bit depth. | 
 **enableRemoteMedia** | **bool** | Optional. Whether to enable remote media. | 
 **enableAudioVbrEncoding** | **bool** | Optional. Whether to enable Audio Encoding. | [default to true]
 **breakOnNonKeyFrames** | **bool** | Optional. Whether to break on non key frames. | [default to false]
 **enableRedirection** | **bool** | Whether to enable redirection. Defaults to true. | [default to true]

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadUniversalAudioStream

> *os.File HeadUniversalAudioStream(ctx, itemId).Container(container).MediaSourceId(mediaSourceId).DeviceId(deviceId).UserId(userId).AudioCodec(audioCodec).MaxAudioChannels(maxAudioChannels).TranscodingAudioChannels(transcodingAudioChannels).MaxStreamingBitrate(maxStreamingBitrate).AudioBitRate(audioBitRate).StartTimeTicks(startTimeTicks).TranscodingContainer(transcodingContainer).TranscodingProtocol(transcodingProtocol).MaxAudioSampleRate(maxAudioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).EnableRemoteMedia(enableRemoteMedia).EnableAudioVbrEncoding(enableAudioVbrEncoding).BreakOnNonKeyFrames(breakOnNonKeyFrames).EnableRedirection(enableRedirection).Execute()

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
	container := []string{"Inner_example"} // []string | Optional. The audio container. (optional)
	mediaSourceId := "mediaSourceId_example" // string | The media version id, if playing an alternate version. (optional)
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. The user id. (optional)
	audioCodec := "audioCodec_example" // string | Optional. The audio codec to transcode to. (optional)
	maxAudioChannels := int32(56) // int32 | Optional. The maximum number of audio channels. (optional)
	transcodingAudioChannels := int32(56) // int32 | Optional. The number of how many audio channels to transcode to. (optional)
	maxStreamingBitrate := int32(56) // int32 | Optional. The maximum streaming bitrate. (optional)
	audioBitRate := int32(56) // int32 | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. (optional)
	startTimeTicks := int64(789) // int64 | Optional. Specify a starting offset, in ticks. 1 tick = 10000 ms. (optional)
	transcodingContainer := "transcodingContainer_example" // string | Optional. The container to transcode to. (optional)
	transcodingProtocol := "transcodingProtocol_example" // MediaStreamProtocol | Optional. The transcoding protocol. (optional)
	maxAudioSampleRate := int32(56) // int32 | Optional. The maximum audio sample rate. (optional)
	maxAudioBitDepth := int32(56) // int32 | Optional. The maximum audio bit depth. (optional)
	enableRemoteMedia := true // bool | Optional. Whether to enable remote media. (optional)
	enableAudioVbrEncoding := true // bool | Optional. Whether to enable Audio Encoding. (optional) (default to true)
	breakOnNonKeyFrames := true // bool | Optional. Whether to break on non key frames. (optional) (default to false)
	enableRedirection := true // bool | Whether to enable redirection. Defaults to true. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UniversalAudioAPI.HeadUniversalAudioStream(context.Background(), itemId).Container(container).MediaSourceId(mediaSourceId).DeviceId(deviceId).UserId(userId).AudioCodec(audioCodec).MaxAudioChannels(maxAudioChannels).TranscodingAudioChannels(transcodingAudioChannels).MaxStreamingBitrate(maxStreamingBitrate).AudioBitRate(audioBitRate).StartTimeTicks(startTimeTicks).TranscodingContainer(transcodingContainer).TranscodingProtocol(transcodingProtocol).MaxAudioSampleRate(maxAudioSampleRate).MaxAudioBitDepth(maxAudioBitDepth).EnableRemoteMedia(enableRemoteMedia).EnableAudioVbrEncoding(enableAudioVbrEncoding).BreakOnNonKeyFrames(breakOnNonKeyFrames).EnableRedirection(enableRedirection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UniversalAudioAPI.HeadUniversalAudioStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadUniversalAudioStream`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `UniversalAudioAPI.HeadUniversalAudioStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadUniversalAudioStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **container** | **[]string** | Optional. The audio container. | 
 **mediaSourceId** | **string** | The media version id, if playing an alternate version. | 
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **userId** | **string** | Optional. The user id. | 
 **audioCodec** | **string** | Optional. The audio codec to transcode to. | 
 **maxAudioChannels** | **int32** | Optional. The maximum number of audio channels. | 
 **transcodingAudioChannels** | **int32** | Optional. The number of how many audio channels to transcode to. | 
 **maxStreamingBitrate** | **int32** | Optional. The maximum streaming bitrate. | 
 **audioBitRate** | **int32** | Optional. Specify an audio bitrate to encode to, e.g. 128000. If omitted this will be left to encoder defaults. | 
 **startTimeTicks** | **int64** | Optional. Specify a starting offset, in ticks. 1 tick &#x3D; 10000 ms. | 
 **transcodingContainer** | **string** | Optional. The container to transcode to. | 
 **transcodingProtocol** | **MediaStreamProtocol** | Optional. The transcoding protocol. | 
 **maxAudioSampleRate** | **int32** | Optional. The maximum audio sample rate. | 
 **maxAudioBitDepth** | **int32** | Optional. The maximum audio bit depth. | 
 **enableRemoteMedia** | **bool** | Optional. Whether to enable remote media. | 
 **enableAudioVbrEncoding** | **bool** | Optional. Whether to enable Audio Encoding. | [default to true]
 **breakOnNonKeyFrames** | **bool** | Optional. Whether to break on non key frames. | [default to false]
 **enableRedirection** | **bool** | Whether to enable redirection. Defaults to true. | [default to true]

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

