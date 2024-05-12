# \TrickplayAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrickplayHlsPlaylist**](TrickplayAPI.md#GetTrickplayHlsPlaylist) | **Get** /Videos/{itemId}/Trickplay/{width}/tiles.m3u8 | Gets an image tiles playlist for trickplay.
[**GetTrickplayTileImage**](TrickplayAPI.md#GetTrickplayTileImage) | **Get** /Videos/{itemId}/Trickplay/{width}/{index}.jpg | Gets a trickplay tile image.



## GetTrickplayHlsPlaylist

> *os.File GetTrickplayHlsPlaylist(ctx, itemId, width).MediaSourceId(mediaSourceId).Execute()

Gets an image tiles playlist for trickplay.

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
	width := int32(56) // int32 | The width of a single tile.
	mediaSourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The media version id, if using an alternate version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrickplayAPI.GetTrickplayHlsPlaylist(context.Background(), itemId, width).MediaSourceId(mediaSourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrickplayAPI.GetTrickplayHlsPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrickplayHlsPlaylist`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TrickplayAPI.GetTrickplayHlsPlaylist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**width** | **int32** | The width of a single tile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrickplayHlsPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mediaSourceId** | **string** | The media version id, if using an alternate version. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-mpegURL, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrickplayTileImage

> *os.File GetTrickplayTileImage(ctx, itemId, width, index).MediaSourceId(mediaSourceId).Execute()

Gets a trickplay tile image.

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
	width := int32(56) // int32 | The width of a single tile.
	index := int32(56) // int32 | The index of the desired tile.
	mediaSourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The media version id, if using an alternate version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TrickplayAPI.GetTrickplayTileImage(context.Background(), itemId, width, index).MediaSourceId(mediaSourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TrickplayAPI.GetTrickplayTileImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrickplayTileImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TrickplayAPI.GetTrickplayTileImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 
**width** | **int32** | The width of a single tile. | 
**index** | **int32** | The index of the desired tile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrickplayTileImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mediaSourceId** | **string** | The media version id, if using an alternate version. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

