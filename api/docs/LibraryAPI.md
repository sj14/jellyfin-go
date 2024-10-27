# \LibraryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteItem**](LibraryAPI.md#DeleteItem) | **Delete** /Items/{itemId} | Deletes an item from the library and filesystem.
[**DeleteItems**](LibraryAPI.md#DeleteItems) | **Delete** /Items | Deletes items from the library and filesystem.
[**GetAncestors**](LibraryAPI.md#GetAncestors) | **Get** /Items/{itemId}/Ancestors | Gets all parents of an item.
[**GetCriticReviews**](LibraryAPI.md#GetCriticReviews) | **Get** /Items/{itemId}/CriticReviews | Gets critic review for an item.
[**GetDownload**](LibraryAPI.md#GetDownload) | **Get** /Items/{itemId}/Download | Downloads item media.
[**GetFile**](LibraryAPI.md#GetFile) | **Get** /Items/{itemId}/File | Get the original file of an item.
[**GetItemCounts**](LibraryAPI.md#GetItemCounts) | **Get** /Items/Counts | Get item counts.
[**GetLibraryOptionsInfo**](LibraryAPI.md#GetLibraryOptionsInfo) | **Get** /Libraries/AvailableOptions | Gets the library options info.
[**GetMediaFolders**](LibraryAPI.md#GetMediaFolders) | **Get** /Library/MediaFolders | Gets all user media folders.
[**GetPhysicalPaths**](LibraryAPI.md#GetPhysicalPaths) | **Get** /Library/PhysicalPaths | Gets a list of physical paths from virtual folders.
[**GetSimilarAlbums**](LibraryAPI.md#GetSimilarAlbums) | **Get** /Albums/{itemId}/Similar | Gets similar items.
[**GetSimilarArtists**](LibraryAPI.md#GetSimilarArtists) | **Get** /Artists/{itemId}/Similar | Gets similar items.
[**GetSimilarItems**](LibraryAPI.md#GetSimilarItems) | **Get** /Items/{itemId}/Similar | Gets similar items.
[**GetSimilarMovies**](LibraryAPI.md#GetSimilarMovies) | **Get** /Movies/{itemId}/Similar | Gets similar items.
[**GetSimilarShows**](LibraryAPI.md#GetSimilarShows) | **Get** /Shows/{itemId}/Similar | Gets similar items.
[**GetSimilarTrailers**](LibraryAPI.md#GetSimilarTrailers) | **Get** /Trailers/{itemId}/Similar | Gets similar items.
[**GetThemeMedia**](LibraryAPI.md#GetThemeMedia) | **Get** /Items/{itemId}/ThemeMedia | Get theme songs and videos for an item.
[**GetThemeSongs**](LibraryAPI.md#GetThemeSongs) | **Get** /Items/{itemId}/ThemeSongs | Get theme songs for an item.
[**GetThemeVideos**](LibraryAPI.md#GetThemeVideos) | **Get** /Items/{itemId}/ThemeVideos | Get theme videos for an item.
[**PostAddedMovies**](LibraryAPI.md#PostAddedMovies) | **Post** /Library/Movies/Added | Reports that new movies have been added by an external source.
[**PostAddedSeries**](LibraryAPI.md#PostAddedSeries) | **Post** /Library/Series/Added | Reports that new episodes of a series have been added by an external source.
[**PostUpdatedMedia**](LibraryAPI.md#PostUpdatedMedia) | **Post** /Library/Media/Updated | Reports that new movies have been added by an external source.
[**PostUpdatedMovies**](LibraryAPI.md#PostUpdatedMovies) | **Post** /Library/Movies/Updated | Reports that new movies have been added by an external source.
[**PostUpdatedSeries**](LibraryAPI.md#PostUpdatedSeries) | **Post** /Library/Series/Updated | Reports that new episodes of a series have been added by an external source.
[**RefreshLibrary**](LibraryAPI.md#RefreshLibrary) | **Post** /Library/Refresh | Starts a library scan.



## DeleteItem

> DeleteItem(ctx, itemId).Execute()

Deletes an item from the library and filesystem.

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
	r, err := apiClient.LibraryAPI.DeleteItem(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.DeleteItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteItemRequest struct via the builder pattern


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


## DeleteItems

> DeleteItems(ctx).Ids(ids).Execute()

Deletes items from the library and filesystem.

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
	ids := []string{"Inner_example"} // []string | The item ids. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.DeleteItems(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.DeleteItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | The item ids. | 

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


## GetAncestors

> []BaseItemDto GetAncestors(ctx, itemId).UserId(userId).Execute()

Gets all parents of an item.

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
	resp, r, err := apiClient.LibraryAPI.GetAncestors(context.Background(), itemId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetAncestors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAncestors`: []BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetAncestors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAncestorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data. | 

### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCriticReviews

> BaseItemDtoQueryResult GetCriticReviews(ctx, itemId).Execute()

Gets critic review for an item.

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
	itemId := "itemId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetCriticReviews(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetCriticReviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCriticReviews`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetCriticReviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCriticReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetDownload

> *os.File GetDownload(ctx, itemId).Execute()

Downloads item media.

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
	resp, r, err := apiClient.LibraryAPI.GetDownload(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDownload`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/*, audio/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFile

> *os.File GetFile(ctx, itemId).Execute()

Get the original file of an item.

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
	resp, r, err := apiClient.LibraryAPI.GetFile(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFile`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: video/*, audio/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemCounts

> ItemCounts GetItemCounts(ctx).UserId(userId).IsFavorite(isFavorite).Execute()

Get item counts.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Get counts from a specific user's library. (optional)
	isFavorite := true // bool | Optional. Get counts of favorite items. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetItemCounts(context.Background()).UserId(userId).IsFavorite(isFavorite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetItemCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemCounts`: ItemCounts
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetItemCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Optional. Get counts from a specific user&#39;s library. | 
 **isFavorite** | **bool** | Optional. Get counts of favorite items. | 

### Return type

[**ItemCounts**](ItemCounts.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLibraryOptionsInfo

> LibraryOptionsResultDto GetLibraryOptionsInfo(ctx).LibraryContentType(libraryContentType).IsNewLibrary(isNewLibrary).Execute()

Gets the library options info.

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
	libraryContentType := "libraryContentType_example" // CollectionType | Library content type. (optional)
	isNewLibrary := true // bool | Whether this is a new library. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetLibraryOptionsInfo(context.Background()).LibraryContentType(libraryContentType).IsNewLibrary(isNewLibrary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetLibraryOptionsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryOptionsInfo`: LibraryOptionsResultDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetLibraryOptionsInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryOptionsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryContentType** | **CollectionType** | Library content type. | 
 **isNewLibrary** | **bool** | Whether this is a new library. | [default to false]

### Return type

[**LibraryOptionsResultDto**](LibraryOptionsResultDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaFolders

> BaseItemDtoQueryResult GetMediaFolders(ctx).IsHidden(isHidden).Execute()

Gets all user media folders.

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
	isHidden := true // bool | Optional. Filter by folders that are marked hidden, or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetMediaFolders(context.Background()).IsHidden(isHidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetMediaFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaFolders`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetMediaFolders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaFoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isHidden** | **bool** | Optional. Filter by folders that are marked hidden, or not. | 

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


## GetPhysicalPaths

> []string GetPhysicalPaths(ctx).Execute()

Gets a list of physical paths from virtual folders.

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
	resp, r, err := apiClient.LibraryAPI.GetPhysicalPaths(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetPhysicalPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPhysicalPaths`: []string
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetPhysicalPaths`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhysicalPathsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarAlbums

> BaseItemDtoQueryResult GetSimilarAlbums(ctx, itemId).ExcludeArtistIds(excludeArtistIds).UserId(userId).Limit(limit).Fields(fields).Execute()

Gets similar items.

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
	excludeArtistIds := []string{"Inner_example"} // []string | Exclude artist ids. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user id, and attach user data. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetSimilarAlbums(context.Background(), itemId).ExcludeArtistIds(excludeArtistIds).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetSimilarAlbums``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimilarAlbums`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetSimilarAlbums`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarAlbumsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeArtistIds** | **[]string** | Exclude artist ids. | 
 **userId** | **string** | Optional. Filter by user id, and attach user data. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 

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


## GetSimilarArtists

> BaseItemDtoQueryResult GetSimilarArtists(ctx, itemId).ExcludeArtistIds(excludeArtistIds).UserId(userId).Limit(limit).Fields(fields).Execute()

Gets similar items.

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
	excludeArtistIds := []string{"Inner_example"} // []string | Exclude artist ids. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user id, and attach user data. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetSimilarArtists(context.Background(), itemId).ExcludeArtistIds(excludeArtistIds).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetSimilarArtists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimilarArtists`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetSimilarArtists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarArtistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeArtistIds** | **[]string** | Exclude artist ids. | 
 **userId** | **string** | Optional. Filter by user id, and attach user data. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 

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


## GetSimilarItems

> BaseItemDtoQueryResult GetSimilarItems(ctx, itemId).ExcludeArtistIds(excludeArtistIds).UserId(userId).Limit(limit).Fields(fields).Execute()

Gets similar items.

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
	excludeArtistIds := []string{"Inner_example"} // []string | Exclude artist ids. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user id, and attach user data. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetSimilarItems(context.Background(), itemId).ExcludeArtistIds(excludeArtistIds).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetSimilarItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimilarItems`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetSimilarItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeArtistIds** | **[]string** | Exclude artist ids. | 
 **userId** | **string** | Optional. Filter by user id, and attach user data. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 

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


## GetSimilarMovies

> BaseItemDtoQueryResult GetSimilarMovies(ctx, itemId).ExcludeArtistIds(excludeArtistIds).UserId(userId).Limit(limit).Fields(fields).Execute()

Gets similar items.

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
	excludeArtistIds := []string{"Inner_example"} // []string | Exclude artist ids. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user id, and attach user data. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetSimilarMovies(context.Background(), itemId).ExcludeArtistIds(excludeArtistIds).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetSimilarMovies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimilarMovies`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetSimilarMovies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarMoviesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeArtistIds** | **[]string** | Exclude artist ids. | 
 **userId** | **string** | Optional. Filter by user id, and attach user data. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 

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


## GetSimilarShows

> BaseItemDtoQueryResult GetSimilarShows(ctx, itemId).ExcludeArtistIds(excludeArtistIds).UserId(userId).Limit(limit).Fields(fields).Execute()

Gets similar items.

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
	excludeArtistIds := []string{"Inner_example"} // []string | Exclude artist ids. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user id, and attach user data. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetSimilarShows(context.Background(), itemId).ExcludeArtistIds(excludeArtistIds).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetSimilarShows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimilarShows`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetSimilarShows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarShowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeArtistIds** | **[]string** | Exclude artist ids. | 
 **userId** | **string** | Optional. Filter by user id, and attach user data. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 

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


## GetSimilarTrailers

> BaseItemDtoQueryResult GetSimilarTrailers(ctx, itemId).ExcludeArtistIds(excludeArtistIds).UserId(userId).Limit(limit).Fields(fields).Execute()

Gets similar items.

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
	excludeArtistIds := []string{"Inner_example"} // []string | Exclude artist ids. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user id, and attach user data. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetSimilarTrailers(context.Background(), itemId).ExcludeArtistIds(excludeArtistIds).UserId(userId).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetSimilarTrailers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimilarTrailers`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetSimilarTrailers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarTrailersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeArtistIds** | **[]string** | Exclude artist ids. | 
 **userId** | **string** | Optional. Filter by user id, and attach user data. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 

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


## GetThemeMedia

> AllThemeMediaResult GetThemeMedia(ctx, itemId).UserId(userId).InheritFromParent(inheritFromParent).SortBy(sortBy).SortOrder(sortOrder).Execute()

Get theme songs and videos for an item.

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
	inheritFromParent := true // bool | Optional. Determines whether or not parent items should be searched for theme media. (optional) (default to false)
	sortBy := []openapiclient.ItemSortBy{openapiclient.ItemSortBy("Default")} // []ItemSortBy | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. (optional)
	sortOrder := []openapiclient.SortOrder{openapiclient.SortOrder("Ascending")} // []SortOrder | Optional. Sort Order - Ascending, Descending. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetThemeMedia(context.Background(), itemId).UserId(userId).InheritFromParent(inheritFromParent).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetThemeMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThemeMedia`: AllThemeMediaResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetThemeMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThemeMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data. | 
 **inheritFromParent** | **bool** | Optional. Determines whether or not parent items should be searched for theme media. | [default to false]
 **sortBy** | [**[]ItemSortBy**](ItemSortBy.md) | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. | 
 **sortOrder** | [**[]SortOrder**](SortOrder.md) | Optional. Sort Order - Ascending, Descending. | 

### Return type

[**AllThemeMediaResult**](AllThemeMediaResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThemeSongs

> ThemeMediaResult GetThemeSongs(ctx, itemId).UserId(userId).InheritFromParent(inheritFromParent).SortBy(sortBy).SortOrder(sortOrder).Execute()

Get theme songs for an item.

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
	inheritFromParent := true // bool | Optional. Determines whether or not parent items should be searched for theme media. (optional) (default to false)
	sortBy := []openapiclient.ItemSortBy{openapiclient.ItemSortBy("Default")} // []ItemSortBy | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. (optional)
	sortOrder := []openapiclient.SortOrder{openapiclient.SortOrder("Ascending")} // []SortOrder | Optional. Sort Order - Ascending, Descending. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetThemeSongs(context.Background(), itemId).UserId(userId).InheritFromParent(inheritFromParent).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetThemeSongs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThemeSongs`: ThemeMediaResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetThemeSongs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThemeSongsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data. | 
 **inheritFromParent** | **bool** | Optional. Determines whether or not parent items should be searched for theme media. | [default to false]
 **sortBy** | [**[]ItemSortBy**](ItemSortBy.md) | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. | 
 **sortOrder** | [**[]SortOrder**](SortOrder.md) | Optional. Sort Order - Ascending, Descending. | 

### Return type

[**ThemeMediaResult**](ThemeMediaResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetThemeVideos

> ThemeMediaResult GetThemeVideos(ctx, itemId).UserId(userId).InheritFromParent(inheritFromParent).SortBy(sortBy).SortOrder(sortOrder).Execute()

Get theme videos for an item.

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
	inheritFromParent := true // bool | Optional. Determines whether or not parent items should be searched for theme media. (optional) (default to false)
	sortBy := []openapiclient.ItemSortBy{openapiclient.ItemSortBy("Default")} // []ItemSortBy | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. (optional)
	sortOrder := []openapiclient.SortOrder{openapiclient.SortOrder("Ascending")} // []SortOrder | Optional. Sort Order - Ascending, Descending. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetThemeVideos(context.Background(), itemId).UserId(userId).InheritFromParent(inheritFromParent).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetThemeVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetThemeVideos`: ThemeMediaResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetThemeVideos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetThemeVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data. | 
 **inheritFromParent** | **bool** | Optional. Determines whether or not parent items should be searched for theme media. | [default to false]
 **sortBy** | [**[]ItemSortBy**](ItemSortBy.md) | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. | 
 **sortOrder** | [**[]SortOrder**](SortOrder.md) | Optional. Sort Order - Ascending, Descending. | 

### Return type

[**ThemeMediaResult**](ThemeMediaResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAddedMovies

> PostAddedMovies(ctx).TmdbId(tmdbId).ImdbId(imdbId).Execute()

Reports that new movies have been added by an external source.

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
	tmdbId := "tmdbId_example" // string | The tmdbId. (optional)
	imdbId := "imdbId_example" // string | The imdbId. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.PostAddedMovies(context.Background()).TmdbId(tmdbId).ImdbId(imdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.PostAddedMovies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAddedMoviesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tmdbId** | **string** | The tmdbId. | 
 **imdbId** | **string** | The imdbId. | 

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


## PostAddedSeries

> PostAddedSeries(ctx).TvdbId(tvdbId).Execute()

Reports that new episodes of a series have been added by an external source.

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
	tvdbId := "tvdbId_example" // string | The tvdbId. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.PostAddedSeries(context.Background()).TvdbId(tvdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.PostAddedSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAddedSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tvdbId** | **string** | The tvdbId. | 

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


## PostUpdatedMedia

> PostUpdatedMedia(ctx).MediaUpdateInfoDto(mediaUpdateInfoDto).Execute()

Reports that new movies have been added by an external source.

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
	mediaUpdateInfoDto := *openapiclient.NewMediaUpdateInfoDto() // MediaUpdateInfoDto | The update paths.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.PostUpdatedMedia(context.Background()).MediaUpdateInfoDto(mediaUpdateInfoDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.PostUpdatedMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUpdatedMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaUpdateInfoDto** | [**MediaUpdateInfoDto**](MediaUpdateInfoDto.md) | The update paths. | 

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


## PostUpdatedMovies

> PostUpdatedMovies(ctx).TmdbId(tmdbId).ImdbId(imdbId).Execute()

Reports that new movies have been added by an external source.

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
	tmdbId := "tmdbId_example" // string | The tmdbId. (optional)
	imdbId := "imdbId_example" // string | The imdbId. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.PostUpdatedMovies(context.Background()).TmdbId(tmdbId).ImdbId(imdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.PostUpdatedMovies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUpdatedMoviesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tmdbId** | **string** | The tmdbId. | 
 **imdbId** | **string** | The imdbId. | 

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


## PostUpdatedSeries

> PostUpdatedSeries(ctx).TvdbId(tvdbId).Execute()

Reports that new episodes of a series have been added by an external source.

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
	tvdbId := "tvdbId_example" // string | The tvdbId. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.PostUpdatedSeries(context.Background()).TvdbId(tvdbId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.PostUpdatedSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUpdatedSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tvdbId** | **string** | The tvdbId. | 

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


## RefreshLibrary

> RefreshLibrary(ctx).Execute()

Starts a library scan.

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
	r, err := apiClient.LibraryAPI.RefreshLibrary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.RefreshLibrary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRefreshLibraryRequest struct via the builder pattern


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

