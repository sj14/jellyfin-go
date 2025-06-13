# \PlaystateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarkPlayedItem**](PlaystateAPI.md#MarkPlayedItem) | **Post** /UserPlayedItems/{itemId} | Marks an item as played for user.
[**MarkUnplayedItem**](PlaystateAPI.md#MarkUnplayedItem) | **Delete** /UserPlayedItems/{itemId} | Marks an item as unplayed for user.
[**OnPlaybackProgress**](PlaystateAPI.md#OnPlaybackProgress) | **Post** /PlayingItems/{itemId}/Progress | Reports a session&#39;s playback progress.
[**OnPlaybackStart**](PlaystateAPI.md#OnPlaybackStart) | **Post** /PlayingItems/{itemId} | Reports that a session has begun playing an item.
[**OnPlaybackStopped**](PlaystateAPI.md#OnPlaybackStopped) | **Delete** /PlayingItems/{itemId} | Reports that a session has stopped playing an item.
[**PingPlaybackSession**](PlaystateAPI.md#PingPlaybackSession) | **Post** /Sessions/Playing/Ping | Pings a playback session.
[**ReportPlaybackProgress**](PlaystateAPI.md#ReportPlaybackProgress) | **Post** /Sessions/Playing/Progress | Reports playback progress within a session.
[**ReportPlaybackStart**](PlaystateAPI.md#ReportPlaybackStart) | **Post** /Sessions/Playing | Reports playback has started within a session.
[**ReportPlaybackStopped**](PlaystateAPI.md#ReportPlaybackStopped) | **Post** /Sessions/Playing/Stopped | Reports playback has stopped within a session.



## MarkPlayedItem

> UserItemDataDto MarkPlayedItem(ctx, itemId).UserId(userId).DatePlayed(datePlayed).Execute()

Marks an item as played for user.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/sj14/jellyfin-go/api"
)

func main() {
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id. (optional)
	datePlayed := time.Now() // time.Time | Optional. The date the item was played. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaystateAPI.MarkPlayedItem(context.Background(), itemId).UserId(userId).DatePlayed(datePlayed).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateAPI.MarkPlayedItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkPlayedItem`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `PlaystateAPI.MarkPlayedItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkPlayedItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User id. | 
 **datePlayed** | **time.Time** | Optional. The date the item was played. | 

### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkUnplayedItem

> UserItemDataDto MarkUnplayedItem(ctx, itemId).UserId(userId).Execute()

Marks an item as unplayed for user.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaystateAPI.MarkUnplayedItem(context.Background(), itemId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateAPI.MarkUnplayedItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkUnplayedItem`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `PlaystateAPI.MarkUnplayedItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkUnplayedItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User id. | 

### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OnPlaybackProgress

> OnPlaybackProgress(ctx, itemId).MediaSourceId(mediaSourceId).PositionTicks(positionTicks).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).VolumeLevel(volumeLevel).PlayMethod(playMethod).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).RepeatMode(repeatMode).IsPaused(isPaused).IsMuted(isMuted).Execute()

Reports a session's playback progress.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.
	mediaSourceId := "mediaSourceId_example" // string | The id of the MediaSource. (optional)
	positionTicks := int64(789) // int64 | Optional. The current position, in ticks. 1 tick = 10000 ms. (optional)
	audioStreamIndex := int32(56) // int32 | The audio stream index. (optional)
	subtitleStreamIndex := int32(56) // int32 | The subtitle stream index. (optional)
	volumeLevel := int32(56) // int32 | Scale of 0-100. (optional)
	playMethod := "playMethod_example" // PlayMethod | The play method. (optional)
	liveStreamId := "liveStreamId_example" // string | The live stream id. (optional)
	playSessionId := "playSessionId_example" // string | The play session id. (optional)
	repeatMode := "repeatMode_example" // RepeatMode | The repeat mode. (optional)
	isPaused := true // bool | Indicates if the player is paused. (optional) (default to false)
	isMuted := true // bool | Indicates if the player is muted. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateAPI.OnPlaybackProgress(context.Background(), itemId).MediaSourceId(mediaSourceId).PositionTicks(positionTicks).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).VolumeLevel(volumeLevel).PlayMethod(playMethod).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).RepeatMode(repeatMode).IsPaused(isPaused).IsMuted(isMuted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateAPI.OnPlaybackProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPlaybackProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaSourceId** | **string** | The id of the MediaSource. | 
 **positionTicks** | **int64** | Optional. The current position, in ticks. 1 tick &#x3D; 10000 ms. | 
 **audioStreamIndex** | **int32** | The audio stream index. | 
 **subtitleStreamIndex** | **int32** | The subtitle stream index. | 
 **volumeLevel** | **int32** | Scale of 0-100. | 
 **playMethod** | **PlayMethod** | The play method. | 
 **liveStreamId** | **string** | The live stream id. | 
 **playSessionId** | **string** | The play session id. | 
 **repeatMode** | **RepeatMode** | The repeat mode. | 
 **isPaused** | **bool** | Indicates if the player is paused. | [default to false]
 **isMuted** | **bool** | Indicates if the player is muted. | [default to false]

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


## OnPlaybackStart

> OnPlaybackStart(ctx, itemId).MediaSourceId(mediaSourceId).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).PlayMethod(playMethod).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).CanSeek(canSeek).Execute()

Reports that a session has begun playing an item.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.
	mediaSourceId := "mediaSourceId_example" // string | The id of the MediaSource. (optional)
	audioStreamIndex := int32(56) // int32 | The audio stream index. (optional)
	subtitleStreamIndex := int32(56) // int32 | The subtitle stream index. (optional)
	playMethod := "playMethod_example" // PlayMethod | The play method. (optional)
	liveStreamId := "liveStreamId_example" // string | The live stream id. (optional)
	playSessionId := "playSessionId_example" // string | The play session id. (optional)
	canSeek := true // bool | Indicates if the client can seek. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateAPI.OnPlaybackStart(context.Background(), itemId).MediaSourceId(mediaSourceId).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).PlayMethod(playMethod).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).CanSeek(canSeek).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateAPI.OnPlaybackStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPlaybackStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaSourceId** | **string** | The id of the MediaSource. | 
 **audioStreamIndex** | **int32** | The audio stream index. | 
 **subtitleStreamIndex** | **int32** | The subtitle stream index. | 
 **playMethod** | **PlayMethod** | The play method. | 
 **liveStreamId** | **string** | The live stream id. | 
 **playSessionId** | **string** | The play session id. | 
 **canSeek** | **bool** | Indicates if the client can seek. | [default to false]

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


## OnPlaybackStopped

> OnPlaybackStopped(ctx, itemId).MediaSourceId(mediaSourceId).NextMediaType(nextMediaType).PositionTicks(positionTicks).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).Execute()

Reports that a session has stopped playing an item.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.
	mediaSourceId := "mediaSourceId_example" // string | The id of the MediaSource. (optional)
	nextMediaType := "nextMediaType_example" // string | The next media type that will play. (optional)
	positionTicks := int64(789) // int64 | Optional. The position, in ticks, where playback stopped. 1 tick = 10000 ms. (optional)
	liveStreamId := "liveStreamId_example" // string | The live stream id. (optional)
	playSessionId := "playSessionId_example" // string | The play session id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateAPI.OnPlaybackStopped(context.Background(), itemId).MediaSourceId(mediaSourceId).NextMediaType(nextMediaType).PositionTicks(positionTicks).LiveStreamId(liveStreamId).PlaySessionId(playSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateAPI.OnPlaybackStopped``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOnPlaybackStoppedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaSourceId** | **string** | The id of the MediaSource. | 
 **nextMediaType** | **string** | The next media type that will play. | 
 **positionTicks** | **int64** | Optional. The position, in ticks, where playback stopped. 1 tick &#x3D; 10000 ms. | 
 **liveStreamId** | **string** | The live stream id. | 
 **playSessionId** | **string** | The play session id. | 

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


## PingPlaybackSession

> PingPlaybackSession(ctx).PlaySessionId(playSessionId).Execute()

Pings a playback session.

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
	playSessionId := "playSessionId_example" // string | Playback session id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateAPI.PingPlaybackSession(context.Background()).PlaySessionId(playSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateAPI.PingPlaybackSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPingPlaybackSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playSessionId** | **string** | Playback session id. | 

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


## ReportPlaybackProgress

> ReportPlaybackProgress(ctx).PlaybackProgressInfo(playbackProgressInfo).Execute()

Reports playback progress within a session.

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
	playbackProgressInfo := *openapiclient.NewPlaybackProgressInfo() // PlaybackProgressInfo | The playback progress info. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateAPI.ReportPlaybackProgress(context.Background()).PlaybackProgressInfo(playbackProgressInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateAPI.ReportPlaybackProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportPlaybackProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playbackProgressInfo** | [**PlaybackProgressInfo**](PlaybackProgressInfo.md) | The playback progress info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportPlaybackStart

> ReportPlaybackStart(ctx).PlaybackStartInfo(playbackStartInfo).Execute()

Reports playback has started within a session.

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
	playbackStartInfo := *openapiclient.NewPlaybackStartInfo() // PlaybackStartInfo | The playback start info. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateAPI.ReportPlaybackStart(context.Background()).PlaybackStartInfo(playbackStartInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateAPI.ReportPlaybackStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportPlaybackStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playbackStartInfo** | [**PlaybackStartInfo**](PlaybackStartInfo.md) | The playback start info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportPlaybackStopped

> ReportPlaybackStopped(ctx).PlaybackStopInfo(playbackStopInfo).Execute()

Reports playback has stopped within a session.

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
	playbackStopInfo := *openapiclient.NewPlaybackStopInfo() // PlaybackStopInfo | The playback stop info. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaystateAPI.ReportPlaybackStopped(context.Background()).PlaybackStopInfo(playbackStopInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaystateAPI.ReportPlaybackStopped``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportPlaybackStoppedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playbackStopInfo** | [**PlaybackStopInfo**](PlaybackStopInfo.md) | The playback stop info. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

