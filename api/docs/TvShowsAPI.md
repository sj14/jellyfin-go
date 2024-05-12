# \TvShowsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEpisodes**](TvShowsAPI.md#GetEpisodes) | **Get** /Shows/{seriesId}/Episodes | Gets episodes for a tv season.
[**GetNextUp**](TvShowsAPI.md#GetNextUp) | **Get** /Shows/NextUp | Gets a list of next up episodes.
[**GetSeasons**](TvShowsAPI.md#GetSeasons) | **Get** /Shows/{seriesId}/Seasons | Gets seasons for a tv series.
[**GetUpcomingEpisodes**](TvShowsAPI.md#GetUpcomingEpisodes) | **Get** /Shows/Upcoming | Gets a list of upcoming episodes.



## GetEpisodes

> BaseItemDtoQueryResult GetEpisodes(ctx, seriesId).UserId(userId).Fields(fields).Season(season).SeasonId(seasonId).IsMissing(isMissing).AdjacentTo(adjacentTo).StartItemId(startItemId).StartIndex(startIndex).Limit(limit).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).SortBy(sortBy).Execute()

Gets episodes for a tv season.

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
	seriesId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The series id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. (optional)
	season := int32(56) // int32 | Optional filter by season number. (optional)
	seasonId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by season id. (optional)
	isMissing := true // bool | Optional. Filter by items that are missing episodes or not. (optional)
	adjacentTo := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Return items that are siblings of a supplied item. (optional)
	startItemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Skip through the list until a given item is found. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	enableImages := true // bool | Optional, include image information in output. (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional. Include user data. (optional)
	sortBy := "sortBy_example" // ItemSortBy | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvShowsAPI.GetEpisodes(context.Background(), seriesId).UserId(userId).Fields(fields).Season(season).SeasonId(seasonId).IsMissing(isMissing).AdjacentTo(adjacentTo).StartItemId(startItemId).StartIndex(startIndex).Limit(limit).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).SortBy(sortBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvShowsAPI.GetEpisodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEpisodes`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `TvShowsAPI.GetEpisodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seriesId** | **string** | The series id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEpisodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | The user id. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 
 **season** | **int32** | Optional filter by season number. | 
 **seasonId** | **string** | Optional. Filter by season id. | 
 **isMissing** | **bool** | Optional. Filter by items that are missing episodes or not. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **startItemId** | **string** | Optional. Skip through the list until a given item is found. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **enableImages** | **bool** | Optional, include image information in output. | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 
 **enableUserData** | **bool** | Optional. Include user data. | 
 **sortBy** | **ItemSortBy** | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. | 

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


## GetNextUp

> BaseItemDtoQueryResult GetNextUp(ctx).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).SeriesId(seriesId).ParentId(parentId).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).NextUpDateCutoff(nextUpDateCutoff).EnableTotalRecordCount(enableTotalRecordCount).DisableFirstEpisode(disableFirstEpisode).EnableResumable(enableResumable).EnableRewatching(enableRewatching).Execute()

Gets a list of next up episodes.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id of the user to get the next up episodes for. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	seriesId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by series id. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Specify this to localize the search to a specific item or folder. Omit to use the root. (optional)
	enableImages := true // bool | Optional. Include image information in output. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. The max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional. Include user data. (optional)
	nextUpDateCutoff := time.Now() // time.Time | Optional. Starting date of shows to show in Next Up section. (optional)
	enableTotalRecordCount := true // bool | Whether to enable the total records count. Defaults to true. (optional) (default to true)
	disableFirstEpisode := true // bool | Whether to disable sending the first episode in a series as next up. (optional) (default to false)
	enableResumable := true // bool | Whether to include resumable episodes in next up results. (optional) (default to true)
	enableRewatching := true // bool | Whether to include watched episodes in next up results. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvShowsAPI.GetNextUp(context.Background()).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).SeriesId(seriesId).ParentId(parentId).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).NextUpDateCutoff(nextUpDateCutoff).EnableTotalRecordCount(enableTotalRecordCount).DisableFirstEpisode(disableFirstEpisode).EnableResumable(enableResumable).EnableRewatching(enableRewatching).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvShowsAPI.GetNextUp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNextUp`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `TvShowsAPI.GetNextUp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNextUpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The user id of the user to get the next up episodes for. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **seriesId** | **string** | Optional. Filter by series id. | 
 **parentId** | **string** | Optional. Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **enableImages** | **bool** | Optional. Include image information in output. | 
 **imageTypeLimit** | **int32** | Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 
 **enableUserData** | **bool** | Optional. Include user data. | 
 **nextUpDateCutoff** | **time.Time** | Optional. Starting date of shows to show in Next Up section. | 
 **enableTotalRecordCount** | **bool** | Whether to enable the total records count. Defaults to true. | [default to true]
 **disableFirstEpisode** | **bool** | Whether to disable sending the first episode in a series as next up. | [default to false]
 **enableResumable** | **bool** | Whether to include resumable episodes in next up results. | [default to true]
 **enableRewatching** | **bool** | Whether to include watched episodes in next up results. | [default to false]

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


## GetSeasons

> BaseItemDtoQueryResult GetSeasons(ctx, seriesId).UserId(userId).Fields(fields).IsSpecialSeason(isSpecialSeason).IsMissing(isMissing).AdjacentTo(adjacentTo).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Execute()

Gets seasons for a tv series.

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
	seriesId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The series id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. (optional)
	isSpecialSeason := true // bool | Optional. Filter by special season. (optional)
	isMissing := true // bool | Optional. Filter by items that are missing episodes or not. (optional)
	adjacentTo := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Return items that are siblings of a supplied item. (optional)
	enableImages := true // bool | Optional. Include image information in output. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. The max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional. Include user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvShowsAPI.GetSeasons(context.Background(), seriesId).UserId(userId).Fields(fields).IsSpecialSeason(isSpecialSeason).IsMissing(isMissing).AdjacentTo(adjacentTo).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvShowsAPI.GetSeasons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeasons`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `TvShowsAPI.GetSeasons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**seriesId** | **string** | The series id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSeasonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | The user id. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls. | 
 **isSpecialSeason** | **bool** | Optional. Filter by special season. | 
 **isMissing** | **bool** | Optional. Filter by items that are missing episodes or not. | 
 **adjacentTo** | **string** | Optional. Return items that are siblings of a supplied item. | 
 **enableImages** | **bool** | Optional. Include image information in output. | 
 **imageTypeLimit** | **int32** | Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 
 **enableUserData** | **bool** | Optional. Include user data. | 

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


## GetUpcomingEpisodes

> BaseItemDtoQueryResult GetUpcomingEpisodes(ctx).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).ParentId(parentId).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Execute()

Gets a list of upcoming episodes.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id of the user to get the upcoming episodes for. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Specify this to localize the search to a specific item or folder. Omit to use the root. (optional)
	enableImages := true // bool | Optional. Include image information in output. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. The max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional. Include user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TvShowsAPI.GetUpcomingEpisodes(context.Background()).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).ParentId(parentId).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TvShowsAPI.GetUpcomingEpisodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUpcomingEpisodes`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `TvShowsAPI.GetUpcomingEpisodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUpcomingEpisodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | The user id of the user to get the upcoming episodes for. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **parentId** | **string** | Optional. Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **enableImages** | **bool** | Optional. Include image information in output. | 
 **imageTypeLimit** | **int32** | Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 
 **enableUserData** | **bool** | Optional. Include user data. | 

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

