# \HlsSegmentAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHlsAudioSegmentLegacyAac**](HlsSegmentAPI.md#GetHlsAudioSegmentLegacyAac) | **Get** /Audio/{itemId}/hls/{segmentId}/stream.aac | Gets the specified audio segment for an audio item.
[**GetHlsAudioSegmentLegacyMp3**](HlsSegmentAPI.md#GetHlsAudioSegmentLegacyMp3) | **Get** /Audio/{itemId}/hls/{segmentId}/stream.mp3 | Gets the specified audio segment for an audio item.
[**GetHlsPlaylistLegacy**](HlsSegmentAPI.md#GetHlsPlaylistLegacy) | **Get** /Videos/{itemId}/hls/{playlistId}/stream.m3u8 | Gets a hls video playlist.
[**GetHlsVideoSegmentLegacy**](HlsSegmentAPI.md#GetHlsVideoSegmentLegacy) | **Get** /Videos/{itemId}/hls/{playlistId}/{segmentId}.{segmentContainer} | Gets a hls video segment.
[**StopEncodingProcess**](HlsSegmentAPI.md#StopEncodingProcess) | **Delete** /Videos/ActiveEncodings | Stops an active encoding.



## GetHlsAudioSegmentLegacyAac

> *os.File GetHlsAudioSegmentLegacyAac(ctx, itemId, segmentId).Execute()

Gets the specified audio segment for an audio item.

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
	itemId := "itemId_example" // string | The item id.
	segmentId := "segmentId_example" // string | The segment id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HlsSegmentAPI.GetHlsAudioSegmentLegacyAac(context.Background(), itemId, segmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HlsSegmentAPI.GetHlsAudioSegmentLegacyAac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHlsAudioSegmentLegacyAac`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `HlsSegmentAPI.GetHlsAudioSegmentLegacyAac`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**segmentId** | **string** | The segment id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHlsAudioSegmentLegacyAacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHlsAudioSegmentLegacyMp3

> *os.File GetHlsAudioSegmentLegacyMp3(ctx, itemId, segmentId).Execute()

Gets the specified audio segment for an audio item.

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
	itemId := "itemId_example" // string | The item id.
	segmentId := "segmentId_example" // string | The segment id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HlsSegmentAPI.GetHlsAudioSegmentLegacyMp3(context.Background(), itemId, segmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HlsSegmentAPI.GetHlsAudioSegmentLegacyMp3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHlsAudioSegmentLegacyMp3`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `HlsSegmentAPI.GetHlsAudioSegmentLegacyMp3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**segmentId** | **string** | The segment id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHlsAudioSegmentLegacyMp3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHlsPlaylistLegacy

> *os.File GetHlsPlaylistLegacy(ctx, itemId, playlistId).Execute()

Gets a hls video playlist.

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
	itemId := "itemId_example" // string | The video id.
	playlistId := "playlistId_example" // string | The playlist id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HlsSegmentAPI.GetHlsPlaylistLegacy(context.Background(), itemId, playlistId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HlsSegmentAPI.GetHlsPlaylistLegacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHlsPlaylistLegacy`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `HlsSegmentAPI.GetHlsPlaylistLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The video id. | 
**playlistId** | **string** | The playlist id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHlsPlaylistLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-mpegURL

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHlsVideoSegmentLegacy

> *os.File GetHlsVideoSegmentLegacy(ctx, itemId, playlistId, segmentId, segmentContainer).Execute()

Gets a hls video segment.

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
	itemId := "itemId_example" // string | The item id.
	playlistId := "playlistId_example" // string | The playlist id.
	segmentId := "segmentId_example" // string | The segment id.
	segmentContainer := "segmentContainer_example" // string | The segment container.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HlsSegmentAPI.GetHlsVideoSegmentLegacy(context.Background(), itemId, playlistId, segmentId, segmentContainer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HlsSegmentAPI.GetHlsVideoSegmentLegacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHlsVideoSegmentLegacy`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `HlsSegmentAPI.GetHlsVideoSegmentLegacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**playlistId** | **string** | The playlist id. | 
**segmentId** | **string** | The segment id. | 
**segmentContainer** | **string** | The segment container. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHlsVideoSegmentLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopEncodingProcess

> StopEncodingProcess(ctx).DeviceId(deviceId).PlaySessionId(playSessionId).Execute()

Stops an active encoding.

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
	deviceId := "deviceId_example" // string | The device id of the client requesting. Used to stop encoding processes when needed.
	playSessionId := "playSessionId_example" // string | The play session id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HlsSegmentAPI.StopEncodingProcess(context.Background()).DeviceId(deviceId).PlaySessionId(playSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HlsSegmentAPI.StopEncodingProcess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopEncodingProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | The device id of the client requesting. Used to stop encoding processes when needed. | 
 **playSessionId** | **string** | The play session id. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

