# \PlaylistsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToPlaylist**](PlaylistsAPI.md#AddToPlaylist) | **Post** /Playlists/{playlistId}/Items | Adds items to a playlist.
[**CreatePlaylist**](PlaylistsAPI.md#CreatePlaylist) | **Post** /Playlists | Creates a new playlist.
[**GetPlaylistItems**](PlaylistsAPI.md#GetPlaylistItems) | **Get** /Playlists/{playlistId}/Items | Gets the original items of a playlist.
[**MoveItem**](PlaylistsAPI.md#MoveItem) | **Post** /Playlists/{playlistId}/Items/{itemId}/Move/{newIndex} | Moves a playlist item.
[**RemoveFromPlaylist**](PlaylistsAPI.md#RemoveFromPlaylist) | **Delete** /Playlists/{playlistId}/Items | Removes items from a playlist.



## AddToPlaylist

> AddToPlaylist(ctx, playlistId).Ids(ids).UserId(userId).Execute()

Adds items to a playlist.

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
	playlistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The playlist id.
	ids := []string{"Inner_example"} // []string | Item id, comma delimited. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The userId. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistsAPI.AddToPlaylist(context.Background(), playlistId).Ids(ids).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.AddToPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** | The playlist id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddToPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **[]string** | Item id, comma delimited. | 
 **userId** | **string** | The userId. | 

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


## CreatePlaylist

> PlaylistCreationResult CreatePlaylist(ctx).Name(name).Ids(ids).UserId(userId).MediaType(mediaType).CreatePlaylistDto(createPlaylistDto).Execute()

Creates a new playlist.



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
	name := "name_example" // string | The playlist name. (optional)
	ids := []string{"Inner_example"} // []string | The item ids. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id. (optional)
	mediaType := "mediaType_example" // string | The media type. (optional)
	createPlaylistDto := *openapiclient.NewCreatePlaylistDto() // CreatePlaylistDto | The create playlist payload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistsAPI.CreatePlaylist(context.Background()).Name(name).Ids(ids).UserId(userId).MediaType(mediaType).CreatePlaylistDto(createPlaylistDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.CreatePlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlaylist`: PlaylistCreationResult
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.CreatePlaylist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The playlist name. | 
 **ids** | **[]string** | The item ids. | 
 **userId** | **string** | The user id. | 
 **mediaType** | **string** | The media type. | 
 **createPlaylistDto** | [**CreatePlaylistDto**](CreatePlaylistDto.md) | The create playlist payload. | 

### Return type

[**PlaylistCreationResult**](PlaylistCreationResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylistItems

> BaseItemDtoQueryResult GetPlaylistItems(ctx, playlistId).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Execute()

Gets the original items of a playlist.

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
	playlistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The playlist id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id.
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	enableImages := true // bool | Optional. Include image information in output. (optional)
	enableUserData := true // bool | Optional. Include user data. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. The max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistsAPI.GetPlaylistItems(context.Background(), playlistId).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.GetPlaylistItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaylistItems`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.GetPlaylistItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** | The playlist id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User id. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **enableImages** | **bool** | Optional. Include image information in output. | 
 **enableUserData** | **bool** | Optional. Include user data. | 
 **imageTypeLimit** | **int32** | Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 

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


## MoveItem

> MoveItem(ctx, playlistId, itemId, newIndex).Execute()

Moves a playlist item.

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
	playlistId := "playlistId_example" // string | The playlist id.
	itemId := "itemId_example" // string | The item id.
	newIndex := int32(56) // int32 | The new index.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistsAPI.MoveItem(context.Background(), playlistId, itemId, newIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.MoveItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** | The playlist id. | 
**itemId** | **string** | The item id. | 
**newIndex** | **int32** | The new index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveItemRequest struct via the builder pattern


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


## RemoveFromPlaylist

> RemoveFromPlaylist(ctx, playlistId).EntryIds(entryIds).Execute()

Removes items from a playlist.

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
	playlistId := "playlistId_example" // string | The playlist id.
	entryIds := []string{"Inner_example"} // []string | The item ids, comma delimited. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistsAPI.RemoveFromPlaylist(context.Background(), playlistId).EntryIds(entryIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.RemoveFromPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** | The playlist id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFromPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entryIds** | **[]string** | The item ids, comma delimited. | 

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

