# \SubtitleAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSubtitle**](SubtitleAPI.md#DeleteSubtitle) | **Delete** /Videos/{itemId}/Subtitles/{index} | Deletes an external subtitle file.
[**DownloadRemoteSubtitles**](SubtitleAPI.md#DownloadRemoteSubtitles) | **Post** /Items/{itemId}/RemoteSearch/Subtitles/{subtitleId} | Downloads a remote subtitle.
[**GetFallbackFont**](SubtitleAPI.md#GetFallbackFont) | **Get** /FallbackFont/Fonts/{name} | Gets a fallback font file.
[**GetFallbackFontList**](SubtitleAPI.md#GetFallbackFontList) | **Get** /FallbackFont/Fonts | Gets a list of available fallback font files.
[**GetRemoteSubtitles**](SubtitleAPI.md#GetRemoteSubtitles) | **Get** /Providers/Subtitles/Subtitles/{id} | Gets the remote subtitles.
[**GetSubtitle**](SubtitleAPI.md#GetSubtitle) | **Get** /Videos/{routeItemId}/{routeMediaSourceId}/Subtitles/{routeIndex}/Stream.{routeFormat} | Gets subtitles in a specified format.
[**GetSubtitlePlaylist**](SubtitleAPI.md#GetSubtitlePlaylist) | **Get** /Videos/{itemId}/{mediaSourceId}/Subtitles/{index}/subtitles.m3u8 | Gets an HLS subtitle playlist.
[**GetSubtitleWithTicks**](SubtitleAPI.md#GetSubtitleWithTicks) | **Get** /Videos/{routeItemId}/{routeMediaSourceId}/Subtitles/{routeIndex}/{routeStartPositionTicks}/Stream.{routeFormat} | Gets subtitles in a specified format.
[**SearchRemoteSubtitles**](SubtitleAPI.md#SearchRemoteSubtitles) | **Get** /Items/{itemId}/RemoteSearch/Subtitles/{language} | Search remote subtitles.
[**UploadSubtitle**](SubtitleAPI.md#UploadSubtitle) | **Post** /Videos/{itemId}/Subtitles | Upload an external subtitle file.



## DeleteSubtitle

> DeleteSubtitle(ctx, itemId, index).Execute()

Deletes an external subtitle file.

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
	index := int32(56) // int32 | The index of the subtitle file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleAPI.DeleteSubtitle(context.Background(), itemId, index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.DeleteSubtitle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**index** | **int32** | The index of the subtitle file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubtitleRequest struct via the builder pattern


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


## DownloadRemoteSubtitles

> DownloadRemoteSubtitles(ctx, itemId, subtitleId).Execute()

Downloads a remote subtitle.

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
	subtitleId := "subtitleId_example" // string | The subtitle id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleAPI.DownloadRemoteSubtitles(context.Background(), itemId, subtitleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.DownloadRemoteSubtitles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**subtitleId** | **string** | The subtitle id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRemoteSubtitlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetFallbackFont

> *os.File GetFallbackFont(ctx, name).Execute()

Gets a fallback font file.

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
	name := "name_example" // string | The name of the fallback font file to get.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubtitleAPI.GetFallbackFont(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.GetFallbackFont``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFallbackFont`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SubtitleAPI.GetFallbackFont`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the fallback font file to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFallbackFontRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: font/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFallbackFontList

> []FontFile GetFallbackFontList(ctx).Execute()

Gets a list of available fallback font files.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubtitleAPI.GetFallbackFontList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.GetFallbackFontList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFallbackFontList`: []FontFile
	fmt.Fprintf(os.Stdout, "Response from `SubtitleAPI.GetFallbackFontList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFallbackFontListRequest struct via the builder pattern


### Return type

[**[]FontFile**](FontFile.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteSubtitles

> *os.File GetRemoteSubtitles(ctx, id).Execute()

Gets the remote subtitles.

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
	id := "id_example" // string | The item id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubtitleAPI.GetRemoteSubtitles(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.GetRemoteSubtitles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteSubtitles`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SubtitleAPI.GetRemoteSubtitles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteSubtitlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubtitle

> *os.File GetSubtitle(ctx, routeItemId, routeMediaSourceId, routeIndex, routeFormat).ItemId(itemId).MediaSourceId(mediaSourceId).Index(index).Format(format).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).AddVttTimeMap(addVttTimeMap).StartPositionTicks(startPositionTicks).Execute()

Gets subtitles in a specified format.

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
	routeItemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The (route) item id.
	routeMediaSourceId := "routeMediaSourceId_example" // string | The (route) media source id.
	routeIndex := int32(56) // int32 | The (route) subtitle stream index.
	routeFormat := "routeFormat_example" // string | The (route) format of the returned subtitle.
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The item id. (optional)
	mediaSourceId := "mediaSourceId_example" // string | The media source id. (optional)
	index := int32(56) // int32 | The subtitle stream index. (optional)
	format := "format_example" // string | The format of the returned subtitle. (optional)
	endPositionTicks := int64(789) // int64 | Optional. The end position of the subtitle in ticks. (optional)
	copyTimestamps := true // bool | Optional. Whether to copy the timestamps. (optional) (default to false)
	addVttTimeMap := true // bool | Optional. Whether to add a VTT time map. (optional) (default to false)
	startPositionTicks := int64(789) // int64 | The start position of the subtitle in ticks. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubtitleAPI.GetSubtitle(context.Background(), routeItemId, routeMediaSourceId, routeIndex, routeFormat).ItemId(itemId).MediaSourceId(mediaSourceId).Index(index).Format(format).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).AddVttTimeMap(addVttTimeMap).StartPositionTicks(startPositionTicks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.GetSubtitle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubtitle`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SubtitleAPI.GetSubtitle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeItemId** | **string** | The (route) item id. | 
**routeMediaSourceId** | **string** | The (route) media source id. | 
**routeIndex** | **int32** | The (route) subtitle stream index. | 
**routeFormat** | **string** | The (route) format of the returned subtitle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubtitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **itemId** | **string** | The item id. | 
 **mediaSourceId** | **string** | The media source id. | 
 **index** | **int32** | The subtitle stream index. | 
 **format** | **string** | The format of the returned subtitle. | 
 **endPositionTicks** | **int64** | Optional. The end position of the subtitle in ticks. | 
 **copyTimestamps** | **bool** | Optional. Whether to copy the timestamps. | [default to false]
 **addVttTimeMap** | **bool** | Optional. Whether to add a VTT time map. | [default to false]
 **startPositionTicks** | **int64** | The start position of the subtitle in ticks. | [default to 0]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubtitlePlaylist

> *os.File GetSubtitlePlaylist(ctx, itemId, index, mediaSourceId).SegmentLength(segmentLength).Execute()

Gets an HLS subtitle playlist.

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
	index := int32(56) // int32 | The subtitle stream index.
	mediaSourceId := "mediaSourceId_example" // string | The media source id.
	segmentLength := int32(56) // int32 | The subtitle segment length.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubtitleAPI.GetSubtitlePlaylist(context.Background(), itemId, index, mediaSourceId).SegmentLength(segmentLength).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.GetSubtitlePlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubtitlePlaylist`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SubtitleAPI.GetSubtitlePlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**index** | **int32** | The subtitle stream index. | 
**mediaSourceId** | **string** | The media source id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubtitlePlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **segmentLength** | **int32** | The subtitle segment length. | 

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


## GetSubtitleWithTicks

> *os.File GetSubtitleWithTicks(ctx, routeItemId, routeMediaSourceId, routeIndex, routeStartPositionTicks, routeFormat).ItemId(itemId).MediaSourceId(mediaSourceId).Index(index).StartPositionTicks(startPositionTicks).Format(format).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).AddVttTimeMap(addVttTimeMap).Execute()

Gets subtitles in a specified format.

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
	routeItemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The (route) item id.
	routeMediaSourceId := "routeMediaSourceId_example" // string | The (route) media source id.
	routeIndex := int32(56) // int32 | The (route) subtitle stream index.
	routeStartPositionTicks := int64(789) // int64 | The (route) start position of the subtitle in ticks.
	routeFormat := "routeFormat_example" // string | The (route) format of the returned subtitle.
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The item id. (optional)
	mediaSourceId := "mediaSourceId_example" // string | The media source id. (optional)
	index := int32(56) // int32 | The subtitle stream index. (optional)
	startPositionTicks := int64(789) // int64 | The start position of the subtitle in ticks. (optional)
	format := "format_example" // string | The format of the returned subtitle. (optional)
	endPositionTicks := int64(789) // int64 | Optional. The end position of the subtitle in ticks. (optional)
	copyTimestamps := true // bool | Optional. Whether to copy the timestamps. (optional) (default to false)
	addVttTimeMap := true // bool | Optional. Whether to add a VTT time map. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubtitleAPI.GetSubtitleWithTicks(context.Background(), routeItemId, routeMediaSourceId, routeIndex, routeStartPositionTicks, routeFormat).ItemId(itemId).MediaSourceId(mediaSourceId).Index(index).StartPositionTicks(startPositionTicks).Format(format).EndPositionTicks(endPositionTicks).CopyTimestamps(copyTimestamps).AddVttTimeMap(addVttTimeMap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.GetSubtitleWithTicks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubtitleWithTicks`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `SubtitleAPI.GetSubtitleWithTicks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeItemId** | **string** | The (route) item id. | 
**routeMediaSourceId** | **string** | The (route) media source id. | 
**routeIndex** | **int32** | The (route) subtitle stream index. | 
**routeStartPositionTicks** | **int64** | The (route) start position of the subtitle in ticks. | 
**routeFormat** | **string** | The (route) format of the returned subtitle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubtitleWithTicksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **itemId** | **string** | The item id. | 
 **mediaSourceId** | **string** | The media source id. | 
 **index** | **int32** | The subtitle stream index. | 
 **startPositionTicks** | **int64** | The start position of the subtitle in ticks. | 
 **format** | **string** | The format of the returned subtitle. | 
 **endPositionTicks** | **int64** | Optional. The end position of the subtitle in ticks. | 
 **copyTimestamps** | **bool** | Optional. Whether to copy the timestamps. | [default to false]
 **addVttTimeMap** | **bool** | Optional. Whether to add a VTT time map. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRemoteSubtitles

> []RemoteSubtitleInfo SearchRemoteSubtitles(ctx, itemId, language).IsPerfectMatch(isPerfectMatch).Execute()

Search remote subtitles.

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
	language := "language_example" // string | The language of the subtitles.
	isPerfectMatch := true // bool | Optional. Only show subtitles which are a perfect match. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubtitleAPI.SearchRemoteSubtitles(context.Background(), itemId, language).IsPerfectMatch(isPerfectMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.SearchRemoteSubtitles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchRemoteSubtitles`: []RemoteSubtitleInfo
	fmt.Fprintf(os.Stdout, "Response from `SubtitleAPI.SearchRemoteSubtitles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**language** | **string** | The language of the subtitles. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchRemoteSubtitlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isPerfectMatch** | **bool** | Optional. Only show subtitles which are a perfect match. | 

### Return type

[**[]RemoteSubtitleInfo**](RemoteSubtitleInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadSubtitle

> UploadSubtitle(ctx, itemId).UploadSubtitleDto(uploadSubtitleDto).Execute()

Upload an external subtitle file.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The item the subtitle belongs to.
	uploadSubtitleDto := *openapiclient.NewUploadSubtitleDto("Language_example", "Format_example", false, "Data_example") // UploadSubtitleDto | The request body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubtitleAPI.UploadSubtitle(context.Background(), itemId).UploadSubtitleDto(uploadSubtitleDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.UploadSubtitle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item the subtitle belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadSubtitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uploadSubtitleDto** | [**UploadSubtitleDto**](UploadSubtitleDto.md) | The request body. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

