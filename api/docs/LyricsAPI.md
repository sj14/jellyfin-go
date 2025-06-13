# \LyricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLyrics**](LyricsAPI.md#DeleteLyrics) | **Delete** /Audio/{itemId}/Lyrics | Deletes an external lyric file.
[**DownloadRemoteLyrics**](LyricsAPI.md#DownloadRemoteLyrics) | **Post** /Audio/{itemId}/RemoteSearch/Lyrics/{lyricId} | Downloads a remote lyric.
[**GetLyrics**](LyricsAPI.md#GetLyrics) | **Get** /Audio/{itemId}/Lyrics | Gets an item&#39;s lyrics.
[**GetRemoteLyrics**](LyricsAPI.md#GetRemoteLyrics) | **Get** /Providers/Lyrics/{lyricId} | Gets the remote lyrics.
[**SearchRemoteLyrics**](LyricsAPI.md#SearchRemoteLyrics) | **Get** /Audio/{itemId}/RemoteSearch/Lyrics | Search remote lyrics.
[**UploadLyrics**](LyricsAPI.md#UploadLyrics) | **Post** /Audio/{itemId}/Lyrics | Upload an external lyric file.



## DeleteLyrics

> DeleteLyrics(ctx, itemId).Execute()

Deletes an external lyric file.

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
	r, err := apiClient.LyricsAPI.DeleteLyrics(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LyricsAPI.DeleteLyrics``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteLyricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadRemoteLyrics

> LyricDto DownloadRemoteLyrics(ctx, itemId, lyricId).Execute()

Downloads a remote lyric.

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
	lyricId := "lyricId_example" // string | The lyric id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LyricsAPI.DownloadRemoteLyrics(context.Background(), itemId, lyricId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LyricsAPI.DownloadRemoteLyrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadRemoteLyrics`: LyricDto
	fmt.Fprintf(os.Stdout, "Response from `LyricsAPI.DownloadRemoteLyrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**lyricId** | **string** | The lyric id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRemoteLyricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LyricDto**](LyricDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLyrics

> LyricDto GetLyrics(ctx, itemId).Execute()

Gets an item's lyrics.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LyricsAPI.GetLyrics(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LyricsAPI.GetLyrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLyrics`: LyricDto
	fmt.Fprintf(os.Stdout, "Response from `LyricsAPI.GetLyrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLyricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LyricDto**](LyricDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteLyrics

> LyricDto GetRemoteLyrics(ctx, lyricId).Execute()

Gets the remote lyrics.

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
	lyricId := "lyricId_example" // string | The remote provider item id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LyricsAPI.GetRemoteLyrics(context.Background(), lyricId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LyricsAPI.GetRemoteLyrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteLyrics`: LyricDto
	fmt.Fprintf(os.Stdout, "Response from `LyricsAPI.GetRemoteLyrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lyricId** | **string** | The remote provider item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteLyricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LyricDto**](LyricDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRemoteLyrics

> []RemoteLyricInfoDto SearchRemoteLyrics(ctx, itemId).Execute()

Search remote lyrics.

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
	resp, r, err := apiClient.LyricsAPI.SearchRemoteLyrics(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LyricsAPI.SearchRemoteLyrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchRemoteLyrics`: []RemoteLyricInfoDto
	fmt.Fprintf(os.Stdout, "Response from `LyricsAPI.SearchRemoteLyrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchRemoteLyricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RemoteLyricInfoDto**](RemoteLyricInfoDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLyrics

> LyricDto UploadLyrics(ctx, itemId).FileName(fileName).Body(body).Execute()

Upload an external lyric file.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The item the lyric belongs to.
	fileName := "fileName_example" // string | Name of the file being uploaded.
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LyricsAPI.UploadLyrics(context.Background(), itemId).FileName(fileName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LyricsAPI.UploadLyrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadLyrics`: LyricDto
	fmt.Fprintf(os.Stdout, "Response from `LyricsAPI.UploadLyrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item the lyric belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadLyricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileName** | **string** | Name of the file being uploaded. | 
 **body** | ***os.File** |  | 

### Return type

[**LyricDto**](LyricDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

