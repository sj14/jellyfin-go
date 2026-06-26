# \LibraryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteItem**](LibraryAPI.md#DeleteItem) | **Delete** /Items/{itemId} | Deletes an item from the library and filesystem.
[**DeleteItems**](LibraryAPI.md#DeleteItems) | **Delete** /Items | Deletes items from the library and filesystem.
[**GetAncestors**](LibraryAPI.md#GetAncestors) | **Get** /Items/{itemId}/Ancestors | Gets all parents of an item.
[**GetDownload**](LibraryAPI.md#GetDownload) | **Get** /Items/{itemId}/Download | Downloads item media.
[**GetFile**](LibraryAPI.md#GetFile) | **Get** /Items/{itemId}/File | Get the original file of an item.
[**GetIntros**](LibraryAPI.md#GetIntros) | **Get** /Items/{itemId}/Intros | Gets intros to play before the main media item plays.
[**GetItem**](LibraryAPI.md#GetItem) | **Get** /Items/{itemId} | Gets an item from a user&#39;s library.
[**GetItemCollections**](LibraryAPI.md#GetItemCollections) | **Get** /Items/{itemId}/Collections | Gets the collections that include the specified item.
[**GetItemCounts**](LibraryAPI.md#GetItemCounts) | **Get** /Items/Counts | Get item counts.
[**GetItems**](LibraryAPI.md#GetItems) | **Get** /Items | Gets items based on a query.
[**GetLatestMedia**](LibraryAPI.md#GetLatestMedia) | **Get** /Items/Latest | Gets latest media.
[**GetLibraryOptionsInfo**](LibraryAPI.md#GetLibraryOptionsInfo) | **Get** /Libraries/AvailableOptions | Gets the library options info.
[**GetLocalTrailers**](LibraryAPI.md#GetLocalTrailers) | **Get** /Items/{itemId}/LocalTrailers | Gets local trailers for an item.
[**GetMediaFolders**](LibraryAPI.md#GetMediaFolders) | **Get** /Library/MediaFolders | Gets all user media folders.
[**GetPhysicalPaths**](LibraryAPI.md#GetPhysicalPaths) | **Get** /Library/PhysicalPaths | Gets a list of physical paths from virtual folders.
[**GetResumeItems**](LibraryAPI.md#GetResumeItems) | **Get** /UserItems/Resume | Gets items based on a query.
[**GetRootFolder**](LibraryAPI.md#GetRootFolder) | **Get** /Items/Root | Gets the root folder from a user&#39;s library.
[**GetSimilarAlbums**](LibraryAPI.md#GetSimilarAlbums) | **Get** /Albums/{itemId}/Similar | Gets similar items.
[**GetSimilarArtists**](LibraryAPI.md#GetSimilarArtists) | **Get** /Artists/{itemId}/Similar | Gets similar items.
[**GetSimilarItems**](LibraryAPI.md#GetSimilarItems) | **Get** /Items/{itemId}/Similar | Gets similar items.
[**GetSimilarMovies**](LibraryAPI.md#GetSimilarMovies) | **Get** /Movies/{itemId}/Similar | Gets similar items.
[**GetSimilarShows**](LibraryAPI.md#GetSimilarShows) | **Get** /Shows/{itemId}/Similar | Gets similar items.
[**GetSimilarTrailers**](LibraryAPI.md#GetSimilarTrailers) | **Get** /Trailers/{itemId}/Similar | Gets similar items.
[**GetSpecialFeatures**](LibraryAPI.md#GetSpecialFeatures) | **Get** /Items/{itemId}/SpecialFeatures | Gets special features for an item.
[**GetThemeMedia**](LibraryAPI.md#GetThemeMedia) | **Get** /Items/{itemId}/ThemeMedia | Get theme songs and videos for an item.
[**GetThemeSongs**](LibraryAPI.md#GetThemeSongs) | **Get** /Items/{itemId}/ThemeSongs | Get theme songs for an item.
[**GetThemeVideos**](LibraryAPI.md#GetThemeVideos) | **Get** /Items/{itemId}/ThemeVideos | Get theme videos for an item.
[**PostAddedMovies**](LibraryAPI.md#PostAddedMovies) | **Post** /Library/Movies/Added | Reports that new movies have been added by an external source.
[**PostAddedSeries**](LibraryAPI.md#PostAddedSeries) | **Post** /Library/Series/Added | Reports that new episodes of a series have been added by an external source.
[**PostUpdatedMedia**](LibraryAPI.md#PostUpdatedMedia) | **Post** /Library/Media/Updated | Reports that new movies have been added by an external source.
[**PostUpdatedMovies**](LibraryAPI.md#PostUpdatedMovies) | **Post** /Library/Movies/Updated | Reports that new movies have been added by an external source.
[**PostUpdatedSeries**](LibraryAPI.md#PostUpdatedSeries) | **Post** /Library/Series/Updated | Reports that new episodes of a series have been added by an external source.
[**RefreshItem**](LibraryAPI.md#RefreshItem) | **Post** /Items/{itemId}/Refresh | Refreshes metadata for an item.
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: video/*, audio/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: video/*, audio/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntros

> BaseItemDtoQueryResult GetIntros(ctx, itemId).UserId(userId).Execute()

Gets intros to play before the main media item plays.

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
	resp, r, err := apiClient.LibraryAPI.GetIntros(context.Background(), itemId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetIntros``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntros`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetIntros`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntrosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User id. | 

### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItem

> BaseItemDto GetItem(ctx, itemId).UserId(userId).Execute()

Gets an item from a user's library.

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
	resp, r, err := apiClient.LibraryAPI.GetItem(context.Background(), itemId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItem`: BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User id. | 

### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemCollections

> BaseItemDtoQueryResult GetItemCollections(ctx, itemId).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).Execute()

Gets the collections that include the specified item.

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
	startIndex := int32(56) // int32 | Optional. The index of the first record in the output. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetItemCollections(context.Background(), itemId).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetItemCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemCollections`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetItemCollections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data. | 
 **startIndex** | **int32** | Optional. The index of the first record in the output. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. | 

### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItems

> BaseItemDtoQueryResult GetItems(ctx).UserId(userId).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).IndexNumber(indexNumber).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHd(isHd).Is4K(is4K).LocationTypes(locationTypes).ExcludeLocationTypes(excludeLocationTypes).IsMissing(isMissing).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).Filters(filters).IsFavorite(isFavorite).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).Years(years).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).Artists(artists).ExcludeArtistIds(excludeArtistIds).ArtistIds(artistIds).AlbumArtistIds(albumArtistIds).ContributingArtistIds(contributingArtistIds).Albums(albums).AlbumIds(albumIds).Ids(ids).VideoTypes(videoTypes).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).CollapseBoxSetItems(collapseBoxSetItems).MinWidth(minWidth).MinHeight(minHeight).MaxWidth(maxWidth).MaxHeight(maxHeight).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).StudioIds(studioIds).GenreIds(genreIds).AudioLanguages(audioLanguages).SubtitleLanguages(subtitleLanguages).EnableTotalRecordCount(enableTotalRecordCount).EnableImages(enableImages).Execute()

Gets items based on a query.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id supplied as query parameter; this is required when not using an API key. (optional)
	maxOfficialRating := "maxOfficialRating_example" // string | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). (optional)
	hasThemeSong := true // bool | Optional filter by items with theme songs. (optional)
	hasThemeVideo := true // bool | Optional filter by items with theme videos. (optional)
	hasSubtitles := true // bool | Optional filter by items with subtitles. (optional)
	hasSpecialFeature := true // bool | Optional filter by items with special features. (optional)
	hasTrailer := true // bool | Optional filter by items with trailers. (optional)
	adjacentTo := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Return items that are siblings of a supplied item. (optional)
	indexNumber := int32(56) // int32 | Optional filter by index number. (optional)
	parentIndexNumber := int32(56) // int32 | Optional filter by parent index number. (optional)
	hasParentalRating := true // bool | Optional filter by items that have or do not have a parental rating. (optional)
	isHd := true // bool | Optional filter by items that are HD or not. (optional)
	is4K := true // bool | Optional filter by items that are 4K or not. (optional)
	locationTypes := []openapiclient.LocationType{openapiclient.LocationType("FileSystem")} // []LocationType | Optional. If specified, results will be filtered based on LocationType. This allows multiple, comma delimited. (optional)
	excludeLocationTypes := []openapiclient.LocationType{openapiclient.LocationType("FileSystem")} // []LocationType | Optional. If specified, results will be filtered based on the LocationType. This allows multiple, comma delimited. (optional)
	isMissing := true // bool | Optional filter by items that are missing episodes or not. (optional)
	isUnaired := true // bool | Optional filter by items that are unaired episodes or not. (optional)
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	minCriticRating := float64(1.2) // float64 | Optional filter by minimum critic rating. (optional)
	minPremiereDate := time.Now() // time.Time | Optional. The minimum premiere date. Format = ISO. (optional)
	minDateLastSaved := time.Now() // time.Time | Optional. The minimum last saved date. Format = ISO. (optional)
	minDateLastSavedForUser := time.Now() // time.Time | Optional. The minimum last saved date for the current user. Format = ISO. (optional)
	maxPremiereDate := time.Now() // time.Time | Optional. The maximum premiere date. Format = ISO. (optional)
	hasOverview := true // bool | Optional filter by items that have an overview or not. (optional)
	hasImdbId := true // bool | Optional filter by items that have an IMDb id or not. (optional)
	hasTmdbId := true // bool | Optional filter by items that have a TMDb id or not. (optional)
	hasTvdbId := true // bool | Optional filter by items that have a TVDb id or not. (optional)
	isMovie := true // bool | Optional filter for live tv movies. (optional)
	isSeries := true // bool | Optional filter for live tv series. (optional)
	isNews := true // bool | Optional filter for live tv news. (optional)
	isKids := true // bool | Optional filter for live tv kids. (optional)
	isSports := true // bool | Optional filter for live tv sports. (optional)
	excludeItemIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered by excluding item ids. This allows multiple, comma delimited. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	recursive := true // bool | When searching within folders, this determines whether or not the search will be recursive. true/false. (optional)
	searchTerm := "searchTerm_example" // string | Optional. Filter based on a search term. (optional)
	sortOrder := []openapiclient.SortOrder{openapiclient.SortOrder("Ascending")} // []SortOrder | Sort Order - Ascending, Descending. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Specify this to localize the search to a specific item or folder. Omit to use the root. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. (optional)
	excludeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. (optional)
	includeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be filtered based on the item type. This allows multiple, comma delimited. (optional)
	filters := []openapiclient.ItemFilter{openapiclient.ItemFilter("IsFolder")} // []ItemFilter | Optional. Specify additional filters to apply. This allows multiple, comma delimited. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes. (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	mediaTypes := []openapiclient.MediaType{openapiclient.MediaType("Unknown")} // []MediaType | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	imageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. (optional)
	sortBy := []openapiclient.ItemSortBy{openapiclient.ItemSortBy("Default")} // []ItemSortBy | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. (optional)
	isPlayed := true // bool | Optional filter by items that are played, or not. (optional)
	genres := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimited. (optional)
	officialRatings := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimited. (optional)
	tags := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimited. (optional)
	years := []int32{int32(123)} // []int32 | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimited. (optional)
	enableUserData := true // bool | Optional, include user data. (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered to include only those containing the specified person id. (optional)
	personTypes := []string{"Inner_example"} // []string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited. (optional)
	studios := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimited. (optional)
	artists := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on artists. This allows multiple, pipe delimited. (optional)
	excludeArtistIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on artist id. This allows multiple, pipe delimited. (optional)
	artistIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered to include only those containing the specified artist id. (optional)
	albumArtistIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered to include only those containing the specified album artist id. (optional)
	contributingArtistIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered to include only those containing the specified contributing artist id. (optional)
	albums := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimited. (optional)
	albumIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on album id. This allows multiple, pipe delimited. (optional)
	ids := []string{"Inner_example"} // []string | Optional. If specific items are needed, specify a list of item id's to retrieve. This allows multiple, comma delimited. (optional)
	videoTypes := []openapiclient.VideoType{openapiclient.VideoType("VideoFile")} // []VideoType | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimited. (optional)
	minOfficialRating := "minOfficialRating_example" // string | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). (optional)
	isLocked := true // bool | Optional filter by items that are locked. (optional)
	isPlaceHolder := true // bool | Optional filter by items that are placeholders. (optional)
	hasOfficialRating := true // bool | Optional filter by items that have official ratings. (optional)
	collapseBoxSetItems := true // bool | Whether or not to hide items behind their boxsets. (optional)
	minWidth := int32(56) // int32 | Optional. Filter by the minimum width of the item. (optional)
	minHeight := int32(56) // int32 | Optional. Filter by the minimum height of the item. (optional)
	maxWidth := int32(56) // int32 | Optional. Filter by the maximum width of the item. (optional)
	maxHeight := int32(56) // int32 | Optional. Filter by the maximum height of the item. (optional)
	is3D := true // bool | Optional filter by items that are 3D, or not. (optional)
	seriesStatus := []openapiclient.SeriesStatus{openapiclient.SeriesStatus("Continuing")} // []SeriesStatus | Optional filter by Series Status. Allows multiple, comma delimited. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)
	studioIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on studio id. This allows multiple, pipe delimited. (optional)
	genreIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on genre id. This allows multiple, pipe delimited. (optional)
	audioLanguages := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on audio language. This allows multiple, comma delimited values. (optional)
	subtitleLanguages := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on subtitle language. This allows multiple, comma delimited values. (optional)
	enableTotalRecordCount := true // bool | Optional. Enable the total record count. (optional) (default to true)
	enableImages := true // bool | Optional, include image information in output. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetItems(context.Background()).UserId(userId).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).IndexNumber(indexNumber).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHd(isHd).Is4K(is4K).LocationTypes(locationTypes).ExcludeLocationTypes(excludeLocationTypes).IsMissing(isMissing).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).Filters(filters).IsFavorite(isFavorite).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).Years(years).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).Artists(artists).ExcludeArtistIds(excludeArtistIds).ArtistIds(artistIds).AlbumArtistIds(albumArtistIds).ContributingArtistIds(contributingArtistIds).Albums(albums).AlbumIds(albumIds).Ids(ids).VideoTypes(videoTypes).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).CollapseBoxSetItems(collapseBoxSetItems).MinWidth(minWidth).MinHeight(minHeight).MaxWidth(maxWidth).MaxHeight(maxHeight).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).StudioIds(studioIds).GenreIds(genreIds).AudioLanguages(audioLanguages).SubtitleLanguages(subtitleLanguages).EnableTotalRecordCount(enableTotalRecordCount).EnableImages(enableImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItems`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The user id supplied as query parameter; this is required when not using an API key. | 
 **maxOfficialRating** | **string** | Optional filter by maximum official rating (PG, PG-13, TV-MA, etc). | 
 **hasThemeSong** | **bool** | Optional filter by items with theme songs. | 
 **hasThemeVideo** | **bool** | Optional filter by items with theme videos. | 
 **hasSubtitles** | **bool** | Optional filter by items with subtitles. | 
 **hasSpecialFeature** | **bool** | Optional filter by items with special features. | 
 **hasTrailer** | **bool** | Optional filter by items with trailers. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **indexNumber** | **int32** | Optional filter by index number. | 
 **parentIndexNumber** | **int32** | Optional filter by parent index number. | 
 **hasParentalRating** | **bool** | Optional filter by items that have or do not have a parental rating. | 
 **isHd** | **bool** | Optional filter by items that are HD or not. | 
 **is4K** | **bool** | Optional filter by items that are 4K or not. | 
 **locationTypes** | [**[]LocationType**](LocationType.md) | Optional. If specified, results will be filtered based on LocationType. This allows multiple, comma delimited. | 
 **excludeLocationTypes** | [**[]LocationType**](LocationType.md) | Optional. If specified, results will be filtered based on the LocationType. This allows multiple, comma delimited. | 
 **isMissing** | **bool** | Optional filter by items that are missing episodes or not. | 
 **isUnaired** | **bool** | Optional filter by items that are unaired episodes or not. | 
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **minCriticRating** | **float64** | Optional filter by minimum critic rating. | 
 **minPremiereDate** | **time.Time** | Optional. The minimum premiere date. Format &#x3D; ISO. | 
 **minDateLastSaved** | **time.Time** | Optional. The minimum last saved date. Format &#x3D; ISO. | 
 **minDateLastSavedForUser** | **time.Time** | Optional. The minimum last saved date for the current user. Format &#x3D; ISO. | 
 **maxPremiereDate** | **time.Time** | Optional. The maximum premiere date. Format &#x3D; ISO. | 
 **hasOverview** | **bool** | Optional filter by items that have an overview or not. | 
 **hasImdbId** | **bool** | Optional filter by items that have an IMDb id or not. | 
 **hasTmdbId** | **bool** | Optional filter by items that have a TMDb id or not. | 
 **hasTvdbId** | **bool** | Optional filter by items that have a TVDb id or not. | 
 **isMovie** | **bool** | Optional filter for live tv movies. | 
 **isSeries** | **bool** | Optional filter for live tv series. | 
 **isNews** | **bool** | Optional filter for live tv news. | 
 **isKids** | **bool** | Optional filter for live tv kids. | 
 **isSports** | **bool** | Optional filter for live tv sports. | 
 **excludeItemIds** | **[]string** | Optional. If specified, results will be filtered by excluding item ids. This allows multiple, comma delimited. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **recursive** | **bool** | When searching within folders, this determines whether or not the search will be recursive. true/false. | 
 **searchTerm** | **string** | Optional. Filter based on a search term. | 
 **sortOrder** | [**[]SortOrder**](SortOrder.md) | Sort Order - Ascending, Descending. | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. | 
 **excludeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. | 
 **includeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be filtered based on the item type. This allows multiple, comma delimited. | 
 **filters** | [**[]ItemFilter**](ItemFilter.md) | Optional. Specify additional filters to apply. This allows multiple, comma delimited. Options: IsFolder, IsNotFolder, IsUnplayed, IsPlayed, IsFavorite, IsResumable, Likes, Dislikes. | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **mediaTypes** | [**[]MediaType**](MediaType.md) | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **imageTypes** | [**[]ImageType**](ImageType.md) | Optional. If specified, results will be filtered based on those containing image types. This allows multiple, comma delimited. | 
 **sortBy** | [**[]ItemSortBy**](ItemSortBy.md) | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. | 
 **isPlayed** | **bool** | Optional filter by items that are played, or not. | 
 **genres** | **[]string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimited. | 
 **officialRatings** | **[]string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimited. | 
 **tags** | **[]string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimited. | 
 **years** | **[]int32** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimited. | 
 **enableUserData** | **bool** | Optional, include user data. | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **[]string** | Optional. If specified, results will be filtered to include only those containing the specified person id. | 
 **personTypes** | **[]string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited. | 
 **studios** | **[]string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimited. | 
 **artists** | **[]string** | Optional. If specified, results will be filtered based on artists. This allows multiple, pipe delimited. | 
 **excludeArtistIds** | **[]string** | Optional. If specified, results will be filtered based on artist id. This allows multiple, pipe delimited. | 
 **artistIds** | **[]string** | Optional. If specified, results will be filtered to include only those containing the specified artist id. | 
 **albumArtistIds** | **[]string** | Optional. If specified, results will be filtered to include only those containing the specified album artist id. | 
 **contributingArtistIds** | **[]string** | Optional. If specified, results will be filtered to include only those containing the specified contributing artist id. | 
 **albums** | **[]string** | Optional. If specified, results will be filtered based on album. This allows multiple, pipe delimited. | 
 **albumIds** | **[]string** | Optional. If specified, results will be filtered based on album id. This allows multiple, pipe delimited. | 
 **ids** | **[]string** | Optional. If specific items are needed, specify a list of item id&#39;s to retrieve. This allows multiple, comma delimited. | 
 **videoTypes** | [**[]VideoType**](VideoType.md) | Optional filter by VideoType (videofile, dvd, bluray, iso). Allows multiple, comma delimited. | 
 **minOfficialRating** | **string** | Optional filter by minimum official rating (PG, PG-13, TV-MA, etc). | 
 **isLocked** | **bool** | Optional filter by items that are locked. | 
 **isPlaceHolder** | **bool** | Optional filter by items that are placeholders. | 
 **hasOfficialRating** | **bool** | Optional filter by items that have official ratings. | 
 **collapseBoxSetItems** | **bool** | Whether or not to hide items behind their boxsets. | 
 **minWidth** | **int32** | Optional. Filter by the minimum width of the item. | 
 **minHeight** | **int32** | Optional. Filter by the minimum height of the item. | 
 **maxWidth** | **int32** | Optional. Filter by the maximum width of the item. | 
 **maxHeight** | **int32** | Optional. Filter by the maximum height of the item. | 
 **is3D** | **bool** | Optional filter by items that are 3D, or not. | 
 **seriesStatus** | [**[]SeriesStatus**](SeriesStatus.md) | Optional filter by Series Status. Allows multiple, comma delimited. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 
 **studioIds** | **[]string** | Optional. If specified, results will be filtered based on studio id. This allows multiple, pipe delimited. | 
 **genreIds** | **[]string** | Optional. If specified, results will be filtered based on genre id. This allows multiple, pipe delimited. | 
 **audioLanguages** | **[]string** | Optional. If specified, results will be filtered based on audio language. This allows multiple, comma delimited values. | 
 **subtitleLanguages** | **[]string** | Optional. If specified, results will be filtered based on subtitle language. This allows multiple, comma delimited values. | 
 **enableTotalRecordCount** | **bool** | Optional. Enable the total record count. | [default to true]
 **enableImages** | **bool** | Optional, include image information in output. | [default to true]

### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestMedia

> []BaseItemDto GetLatestMedia(ctx).UserId(userId).ParentId(parentId).Fields(fields).IncludeItemTypes(includeItemTypes).IsPlayed(isPlayed).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Limit(limit).GroupItems(groupItems).Execute()

Gets latest media.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Specify this to localize the search to a specific item or folder. Omit to use the root. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	includeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. (optional)
	isPlayed := true // bool | Filter by items that are played, or not. (optional)
	enableImages := true // bool | Optional. include image information in output. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. the max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional. include user data. (optional)
	limit := int32(56) // int32 | Return item limit. (optional) (default to 20)
	groupItems := true // bool | Whether or not to group items into a parent container. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetLatestMedia(context.Background()).UserId(userId).ParentId(parentId).Fields(fields).IncludeItemTypes(includeItemTypes).IsPlayed(isPlayed).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Limit(limit).GroupItems(groupItems).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetLatestMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestMedia`: []BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetLatestMedia`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User id. | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **includeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. | 
 **isPlayed** | **bool** | Filter by items that are played, or not. | 
 **enableImages** | **bool** | Optional. include image information in output. | 
 **imageTypeLimit** | **int32** | Optional. the max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 
 **enableUserData** | **bool** | Optional. include user data. | 
 **limit** | **int32** | Return item limit. | [default to 20]
 **groupItems** | **bool** | Whether or not to group items into a parent container. | [default to true]

### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalTrailers

> []BaseItemDto GetLocalTrailers(ctx, itemId).UserId(userId).Execute()

Gets local trailers for an item.

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
	resp, r, err := apiClient.LibraryAPI.GetLocalTrailers(context.Background(), itemId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetLocalTrailers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalTrailers`: []BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetLocalTrailers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalTrailersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User id. | 

### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResumeItems

> BaseItemDtoQueryResult GetResumeItems(ctx).UserId(userId).StartIndex(startIndex).Limit(limit).SearchTerm(searchTerm).ParentId(parentId).Fields(fields).MediaTypes(mediaTypes).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).EnableTotalRecordCount(enableTotalRecordCount).EnableImages(enableImages).ExcludeActiveSessions(excludeActiveSessions).Execute()

Gets items based on a query.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id. (optional)
	startIndex := int32(56) // int32 | The start index. (optional)
	limit := int32(56) // int32 | The item limit. (optional)
	searchTerm := "searchTerm_example" // string | The search term. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Specify this to localize the search to a specific item or folder. Omit to use the root. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. (optional)
	mediaTypes := []openapiclient.MediaType{openapiclient.MediaType("Unknown")} // []MediaType | Optional. Filter by MediaType. Allows multiple, comma delimited. (optional)
	enableUserData := true // bool | Optional. Include user data. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. The max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)
	excludeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. (optional)
	includeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be filtered based on the item type. This allows multiple, comma delimited. (optional)
	enableTotalRecordCount := true // bool | Optional. Enable the total record count. (optional) (default to true)
	enableImages := true // bool | Optional. Include image information in output. (optional) (default to true)
	excludeActiveSessions := true // bool | Optional. Whether to exclude the currently active sessions. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetResumeItems(context.Background()).UserId(userId).StartIndex(startIndex).Limit(limit).SearchTerm(searchTerm).ParentId(parentId).Fields(fields).MediaTypes(mediaTypes).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).EnableTotalRecordCount(enableTotalRecordCount).EnableImages(enableImages).ExcludeActiveSessions(excludeActiveSessions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetResumeItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResumeItems`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetResumeItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResumeItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The user id. | 
 **startIndex** | **int32** | The start index. | 
 **limit** | **int32** | The item limit. | 
 **searchTerm** | **string** | The search term. | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines. | 
 **mediaTypes** | [**[]MediaType**](MediaType.md) | Optional. Filter by MediaType. Allows multiple, comma delimited. | 
 **enableUserData** | **bool** | Optional. Include user data. | 
 **imageTypeLimit** | **int32** | Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 
 **excludeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. | 
 **includeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be filtered based on the item type. This allows multiple, comma delimited. | 
 **enableTotalRecordCount** | **bool** | Optional. Enable the total record count. | [default to true]
 **enableImages** | **bool** | Optional. Include image information in output. | [default to true]
 **excludeActiveSessions** | **bool** | Optional. Whether to exclude the currently active sessions. | [default to false]

### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRootFolder

> BaseItemDto GetRootFolder(ctx).UserId(userId).Execute()

Gets the root folder from a user's library.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAPI.GetRootFolder(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetRootFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRootFolder`: BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetRootFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRootFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User id. | 

### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecialFeatures

> []BaseItemDto GetSpecialFeatures(ctx, itemId).UserId(userId).Execute()

Gets special features for an item.

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
	resp, r, err := apiClient.LibraryAPI.GetSpecialFeatures(context.Background(), itemId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.GetSpecialFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpecialFeatures`: []BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `LibraryAPI.GetSpecialFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecialFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User id. | 

### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
- **Accept**: text/html

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
- **Accept**: text/html

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
- **Accept**: text/html

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
- **Accept**: text/html

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
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshItem

> RefreshItem(ctx, itemId).MetadataRefreshMode(metadataRefreshMode).ImageRefreshMode(imageRefreshMode).ReplaceAllMetadata(replaceAllMetadata).ReplaceAllImages(replaceAllImages).RegenerateTrickplay(regenerateTrickplay).Execute()

Refreshes metadata for an item.

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
	metadataRefreshMode := "metadataRefreshMode_example" // MetadataRefreshMode | (Optional) Specifies the metadata refresh mode. (optional) (default to "None")
	imageRefreshMode := "imageRefreshMode_example" // MetadataRefreshMode | (Optional) Specifies the image refresh mode. (optional) (default to "None")
	replaceAllMetadata := true // bool | (Optional) Determines if metadata should be replaced. Only applicable if mode is FullRefresh. (optional) (default to false)
	replaceAllImages := true // bool | (Optional) Determines if images should be replaced. Only applicable if mode is FullRefresh. (optional) (default to false)
	regenerateTrickplay := true // bool | (Optional) Determines if trickplay images should be replaced. Only applicable if mode is FullRefresh. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryAPI.RefreshItem(context.Background(), itemId).MetadataRefreshMode(metadataRefreshMode).ImageRefreshMode(imageRefreshMode).ReplaceAllMetadata(replaceAllMetadata).ReplaceAllImages(replaceAllImages).RegenerateTrickplay(regenerateTrickplay).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAPI.RefreshItem``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRefreshItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metadataRefreshMode** | **MetadataRefreshMode** | (Optional) Specifies the metadata refresh mode. | [default to &quot;None&quot;]
 **imageRefreshMode** | **MetadataRefreshMode** | (Optional) Specifies the image refresh mode. | [default to &quot;None&quot;]
 **replaceAllMetadata** | **bool** | (Optional) Determines if metadata should be replaced. Only applicable if mode is FullRefresh. | [default to false]
 **replaceAllImages** | **bool** | (Optional) Determines if images should be replaced. Only applicable if mode is FullRefresh. | [default to false]
 **regenerateTrickplay** | **bool** | (Optional) Determines if trickplay images should be replaced. Only applicable if mode is FullRefresh. | [default to false]

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
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

