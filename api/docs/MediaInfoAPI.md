# \MediaInfoAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseLiveStream**](MediaInfoAPI.md#CloseLiveStream) | **Post** /LiveStreams/Close | Closes a media source.
[**GetBitrateTestBytes**](MediaInfoAPI.md#GetBitrateTestBytes) | **Get** /Playback/BitrateTest | Tests the network with a request with the size of the bitrate.
[**GetPlaybackInfo**](MediaInfoAPI.md#GetPlaybackInfo) | **Get** /Items/{itemId}/PlaybackInfo | Gets live playback media info for an item.
[**GetPostedPlaybackInfo**](MediaInfoAPI.md#GetPostedPlaybackInfo) | **Post** /Items/{itemId}/PlaybackInfo | Gets live playback media info for an item.
[**OpenLiveStream**](MediaInfoAPI.md#OpenLiveStream) | **Post** /LiveStreams/Open | Opens a media source.



## CloseLiveStream

> CloseLiveStream(ctx).LiveStreamId(liveStreamId).Execute()

Closes a media source.

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
	liveStreamId := "liveStreamId_example" // string | The livestream id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MediaInfoAPI.CloseLiveStream(context.Background()).LiveStreamId(liveStreamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaInfoAPI.CloseLiveStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloseLiveStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **liveStreamId** | **string** | The livestream id. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBitrateTestBytes

> *os.File GetBitrateTestBytes(ctx).Size(size).Execute()

Tests the network with a request with the size of the bitrate.

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
	size := int32(56) // int32 | The bitrate. Defaults to 102400. (optional) (default to 102400)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaInfoAPI.GetBitrateTestBytes(context.Background()).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaInfoAPI.GetBitrateTestBytes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBitrateTestBytes`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `MediaInfoAPI.GetBitrateTestBytes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBitrateTestBytesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | The bitrate. Defaults to 102400. | [default to 102400]

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaybackInfo

> PlaybackInfoResponse GetPlaybackInfo(ctx, itemId).UserId(userId).Execute()

Gets live playback media info for an item.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaInfoAPI.GetPlaybackInfo(context.Background(), itemId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaInfoAPI.GetPlaybackInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaybackInfo`: PlaybackInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `MediaInfoAPI.GetPlaybackInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaybackInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | The user id. | 

### Return type

[**PlaybackInfoResponse**](PlaybackInfoResponse.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostedPlaybackInfo

> PlaybackInfoResponse GetPostedPlaybackInfo(ctx, itemId).UserId(userId).MaxStreamingBitrate(maxStreamingBitrate).StartTimeTicks(startTimeTicks).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).MaxAudioChannels(maxAudioChannels).MediaSourceId(mediaSourceId).LiveStreamId(liveStreamId).AutoOpenLiveStream(autoOpenLiveStream).EnableDirectPlay(enableDirectPlay).EnableDirectStream(enableDirectStream).EnableTranscoding(enableTranscoding).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).PlaybackInfoDto(playbackInfoDto).Execute()

Gets live playback media info for an item.



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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id. (optional)
	maxStreamingBitrate := int32(56) // int32 | The maximum streaming bitrate. (optional)
	startTimeTicks := int64(789) // int64 | The start time in ticks. (optional)
	audioStreamIndex := int32(56) // int32 | The audio stream index. (optional)
	subtitleStreamIndex := int32(56) // int32 | The subtitle stream index. (optional)
	maxAudioChannels := int32(56) // int32 | The maximum number of audio channels. (optional)
	mediaSourceId := "mediaSourceId_example" // string | The media source id. (optional)
	liveStreamId := "liveStreamId_example" // string | The livestream id. (optional)
	autoOpenLiveStream := true // bool | Whether to auto open the livestream. (optional)
	enableDirectPlay := true // bool | Whether to enable direct play. Default: true. (optional)
	enableDirectStream := true // bool | Whether to enable direct stream. Default: true. (optional)
	enableTranscoding := true // bool | Whether to enable transcoding. Default: true. (optional)
	allowVideoStreamCopy := true // bool | Whether to allow to copy the video stream. Default: true. (optional)
	allowAudioStreamCopy := true // bool | Whether to allow to copy the audio stream. Default: true. (optional)
	playbackInfoDto := *openapiclient.NewPlaybackInfoDto() // PlaybackInfoDto | The playback info. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaInfoAPI.GetPostedPlaybackInfo(context.Background(), itemId).UserId(userId).MaxStreamingBitrate(maxStreamingBitrate).StartTimeTicks(startTimeTicks).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).MaxAudioChannels(maxAudioChannels).MediaSourceId(mediaSourceId).LiveStreamId(liveStreamId).AutoOpenLiveStream(autoOpenLiveStream).EnableDirectPlay(enableDirectPlay).EnableDirectStream(enableDirectStream).EnableTranscoding(enableTranscoding).AllowVideoStreamCopy(allowVideoStreamCopy).AllowAudioStreamCopy(allowAudioStreamCopy).PlaybackInfoDto(playbackInfoDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaInfoAPI.GetPostedPlaybackInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostedPlaybackInfo`: PlaybackInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `MediaInfoAPI.GetPostedPlaybackInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostedPlaybackInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | The user id. | 
 **maxStreamingBitrate** | **int32** | The maximum streaming bitrate. | 
 **startTimeTicks** | **int64** | The start time in ticks. | 
 **audioStreamIndex** | **int32** | The audio stream index. | 
 **subtitleStreamIndex** | **int32** | The subtitle stream index. | 
 **maxAudioChannels** | **int32** | The maximum number of audio channels. | 
 **mediaSourceId** | **string** | The media source id. | 
 **liveStreamId** | **string** | The livestream id. | 
 **autoOpenLiveStream** | **bool** | Whether to auto open the livestream. | 
 **enableDirectPlay** | **bool** | Whether to enable direct play. Default: true. | 
 **enableDirectStream** | **bool** | Whether to enable direct stream. Default: true. | 
 **enableTranscoding** | **bool** | Whether to enable transcoding. Default: true. | 
 **allowVideoStreamCopy** | **bool** | Whether to allow to copy the video stream. Default: true. | 
 **allowAudioStreamCopy** | **bool** | Whether to allow to copy the audio stream. Default: true. | 
 **playbackInfoDto** | [**PlaybackInfoDto**](PlaybackInfoDto.md) | The playback info. | 

### Return type

[**PlaybackInfoResponse**](PlaybackInfoResponse.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenLiveStream

> LiveStreamResponse OpenLiveStream(ctx).OpenToken(openToken).UserId(userId).PlaySessionId(playSessionId).MaxStreamingBitrate(maxStreamingBitrate).StartTimeTicks(startTimeTicks).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).MaxAudioChannels(maxAudioChannels).ItemId(itemId).EnableDirectPlay(enableDirectPlay).EnableDirectStream(enableDirectStream).AlwaysBurnInSubtitleWhenTranscoding(alwaysBurnInSubtitleWhenTranscoding).OpenLiveStreamDto(openLiveStreamDto).Execute()

Opens a media source.

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
	openToken := "openToken_example" // string | The open token. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id. (optional)
	playSessionId := "playSessionId_example" // string | The play session id. (optional)
	maxStreamingBitrate := int32(56) // int32 | The maximum streaming bitrate. (optional)
	startTimeTicks := int64(789) // int64 | The start time in ticks. (optional)
	audioStreamIndex := int32(56) // int32 | The audio stream index. (optional)
	subtitleStreamIndex := int32(56) // int32 | The subtitle stream index. (optional)
	maxAudioChannels := int32(56) // int32 | The maximum number of audio channels. (optional)
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The item id. (optional)
	enableDirectPlay := true // bool | Whether to enable direct play. Default: true. (optional)
	enableDirectStream := true // bool | Whether to enable direct stream. Default: true. (optional)
	alwaysBurnInSubtitleWhenTranscoding := true // bool | Always burn-in subtitle when transcoding. (optional)
	openLiveStreamDto := *openapiclient.NewOpenLiveStreamDto() // OpenLiveStreamDto | The open live stream dto. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaInfoAPI.OpenLiveStream(context.Background()).OpenToken(openToken).UserId(userId).PlaySessionId(playSessionId).MaxStreamingBitrate(maxStreamingBitrate).StartTimeTicks(startTimeTicks).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).MaxAudioChannels(maxAudioChannels).ItemId(itemId).EnableDirectPlay(enableDirectPlay).EnableDirectStream(enableDirectStream).AlwaysBurnInSubtitleWhenTranscoding(alwaysBurnInSubtitleWhenTranscoding).OpenLiveStreamDto(openLiveStreamDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaInfoAPI.OpenLiveStream``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OpenLiveStream`: LiveStreamResponse
	fmt.Fprintf(os.Stdout, "Response from `MediaInfoAPI.OpenLiveStream`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenLiveStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openToken** | **string** | The open token. | 
 **userId** | **string** | The user id. | 
 **playSessionId** | **string** | The play session id. | 
 **maxStreamingBitrate** | **int32** | The maximum streaming bitrate. | 
 **startTimeTicks** | **int64** | The start time in ticks. | 
 **audioStreamIndex** | **int32** | The audio stream index. | 
 **subtitleStreamIndex** | **int32** | The subtitle stream index. | 
 **maxAudioChannels** | **int32** | The maximum number of audio channels. | 
 **itemId** | **string** | The item id. | 
 **enableDirectPlay** | **bool** | Whether to enable direct play. Default: true. | 
 **enableDirectStream** | **bool** | Whether to enable direct stream. Default: true. | 
 **alwaysBurnInSubtitleWhenTranscoding** | **bool** | Always burn-in subtitle when transcoding. | 
 **openLiveStreamDto** | [**OpenLiveStreamDto**](OpenLiveStreamDto.md) | The open live stream dto. | 

### Return type

[**LiveStreamResponse**](LiveStreamResponse.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

