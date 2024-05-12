# \ItemsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetItemUserData**](ItemsAPI.md#GetItemUserData) | **Get** /UserItems/{itemId}/UserData | Get Item User Data.
[**GetItems**](ItemsAPI.md#GetItems) | **Get** /Items | Gets items based on a query.
[**GetResumeItems**](ItemsAPI.md#GetResumeItems) | **Get** /UserItems/Resume | Gets items based on a query.
[**UpdateItemUserData**](ItemsAPI.md#UpdateItemUserData) | **Post** /UserItems/{itemId}/UserData | Update Item User Data.



## GetItemUserData

> UserItemDataDto GetItemUserData(ctx, itemId).UserId(userId).Execute()

Get Item User Data.

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
	resp, r, err := apiClient.ItemsAPI.GetItemUserData(context.Background(), itemId).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemsAPI.GetItemUserData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemUserData`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `ItemsAPI.GetItemUserData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemUserDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | The user id. | 

### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItems

> BaseItemDtoQueryResult GetItems(ctx).UserId(userId).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHd(isHd).Is4K(is4K).LocationTypes(locationTypes).ExcludeLocationTypes(excludeLocationTypes).IsMissing(isMissing).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).Filters(filters).IsFavorite(isFavorite).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).Years(years).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).Artists(artists).ExcludeArtistIds(excludeArtistIds).ArtistIds(artistIds).AlbumArtistIds(albumArtistIds).ContributingArtistIds(contributingArtistIds).Albums(albums).AlbumIds(albumIds).Ids(ids).VideoTypes(videoTypes).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).CollapseBoxSetItems(collapseBoxSetItems).MinWidth(minWidth).MinHeight(minHeight).MaxWidth(maxWidth).MaxHeight(maxHeight).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).StudioIds(studioIds).GenreIds(genreIds).EnableTotalRecordCount(enableTotalRecordCount).EnableImages(enableImages).Execute()

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
	enableTotalRecordCount := true // bool | Optional. Enable the total record count. (optional) (default to true)
	enableImages := true // bool | Optional, include image information in output. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemsAPI.GetItems(context.Background()).UserId(userId).MaxOfficialRating(maxOfficialRating).HasThemeSong(hasThemeSong).HasThemeVideo(hasThemeVideo).HasSubtitles(hasSubtitles).HasSpecialFeature(hasSpecialFeature).HasTrailer(hasTrailer).AdjacentTo(adjacentTo).ParentIndexNumber(parentIndexNumber).HasParentalRating(hasParentalRating).IsHd(isHd).Is4K(is4K).LocationTypes(locationTypes).ExcludeLocationTypes(excludeLocationTypes).IsMissing(isMissing).IsUnaired(isUnaired).MinCommunityRating(minCommunityRating).MinCriticRating(minCriticRating).MinPremiereDate(minPremiereDate).MinDateLastSaved(minDateLastSaved).MinDateLastSavedForUser(minDateLastSavedForUser).MaxPremiereDate(maxPremiereDate).HasOverview(hasOverview).HasImdbId(hasImdbId).HasTmdbId(hasTmdbId).HasTvdbId(hasTvdbId).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).ExcludeItemIds(excludeItemIds).StartIndex(startIndex).Limit(limit).Recursive(recursive).SearchTerm(searchTerm).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).Filters(filters).IsFavorite(isFavorite).MediaTypes(mediaTypes).ImageTypes(imageTypes).SortBy(sortBy).IsPlayed(isPlayed).Genres(genres).OfficialRatings(officialRatings).Tags(tags).Years(years).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).Artists(artists).ExcludeArtistIds(excludeArtistIds).ArtistIds(artistIds).AlbumArtistIds(albumArtistIds).ContributingArtistIds(contributingArtistIds).Albums(albums).AlbumIds(albumIds).Ids(ids).VideoTypes(videoTypes).MinOfficialRating(minOfficialRating).IsLocked(isLocked).IsPlaceHolder(isPlaceHolder).HasOfficialRating(hasOfficialRating).CollapseBoxSetItems(collapseBoxSetItems).MinWidth(minWidth).MinHeight(minHeight).MaxWidth(maxWidth).MaxHeight(maxHeight).Is3D(is3D).SeriesStatus(seriesStatus).NameStartsWithOrGreater(nameStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).StudioIds(studioIds).GenreIds(genreIds).EnableTotalRecordCount(enableTotalRecordCount).EnableImages(enableImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemsAPI.GetItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItems`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `ItemsAPI.GetItems`: %v\n", resp)
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
 **enableTotalRecordCount** | **bool** | Optional. Enable the total record count. | [default to true]
 **enableImages** | **bool** | Optional, include image information in output. | [default to true]

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
	resp, r, err := apiClient.ItemsAPI.GetResumeItems(context.Background()).UserId(userId).StartIndex(startIndex).Limit(limit).SearchTerm(searchTerm).ParentId(parentId).Fields(fields).MediaTypes(mediaTypes).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).EnableTotalRecordCount(enableTotalRecordCount).EnableImages(enableImages).ExcludeActiveSessions(excludeActiveSessions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemsAPI.GetResumeItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResumeItems`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `ItemsAPI.GetResumeItems`: %v\n", resp)
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItemUserData

> UserItemDataDto UpdateItemUserData(ctx, itemId).UpdateUserItemDataDto(updateUserItemDataDto).UserId(userId).Execute()

Update Item User Data.

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
	updateUserItemDataDto := *openapiclient.NewUpdateUserItemDataDto() // UpdateUserItemDataDto | New user data object.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemsAPI.UpdateItemUserData(context.Background(), itemId).UpdateUserItemDataDto(updateUserItemDataDto).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemsAPI.UpdateItemUserData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateItemUserData`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `ItemsAPI.UpdateItemUserData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemUserDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserItemDataDto** | [**UpdateUserItemDataDto**](UpdateUserItemDataDto.md) | New user data object. | 
 **userId** | **string** | The user id. | 

### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

