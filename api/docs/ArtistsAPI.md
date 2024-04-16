# \ArtistsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlbumArtists**](ArtistsAPI.md#GetAlbumArtists) | **Get** /Artists/AlbumArtists | Gets all album artists from a given item, folder, or the entire library.
[**GetArtistByName**](ArtistsAPI.md#GetArtistByName) | **Get** /Artists/{name} | Gets an artist by name.
[**GetArtists**](ArtistsAPI.md#GetArtists) | **Get** /Artists | Gets all artists from a given item, folder, or the entire library.



## GetAlbumArtists

> BaseItemDtoQueryResult GetAlbumArtists(ctx).MinCommunityRating(minCommunityRating).StartIndex(startIndex).Limit(limit).SearchTerm(searchTerm).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).Filters(filters).IsFavorite(isFavorite).MediaTypes(mediaTypes).Genres(genres).GenreIds(genreIds).OfficialRatings(officialRatings).Tags(tags).Years(years).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).UserId(userId).NameStartsWithOrGreater(nameStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).SortBy(sortBy).SortOrder(sortOrder).EnableImages(enableImages).EnableTotalRecordCount(enableTotalRecordCount).Execute()

Gets all album artists from a given item, folder, or the entire library.

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
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	searchTerm := "searchTerm_example" // string | Optional. Search term. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Specify this to localize the search to a specific item or folder. Omit to use the root. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	excludeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be filtered out based on item type. This allows multiple, comma delimited. (optional)
	includeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. (optional)
	filters := []openapiclient.ItemFilter{openapiclient.ItemFilter("IsFolder")} // []ItemFilter | Optional. Specify additional filters to apply. (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	mediaTypes := []string{"Inner_example"} // []string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	genres := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimited. (optional)
	genreIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on genre id. This allows multiple, pipe delimited. (optional)
	officialRatings := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimited. (optional)
	tags := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimited. (optional)
	years := []int32{int32(123)} // []int32 | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimited. (optional)
	enableUserData := true // bool | Optional, include user data. (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered to include only those containing the specified person ids. (optional)
	personTypes := []string{"Inner_example"} // []string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited. (optional)
	studios := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimited. (optional)
	studioIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on studio id. This allows multiple, pipe delimited. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)
	sortBy := []string{"Inner_example"} // []string | Optional. Specify one or more sort orders, comma delimited. (optional)
	sortOrder := []openapiclient.SortOrder{openapiclient.SortOrder("Ascending")} // []SortOrder | Sort Order - Ascending,Descending. (optional)
	enableImages := true // bool | Optional, include image information in output. (optional) (default to true)
	enableTotalRecordCount := true // bool | Total record count. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtistsAPI.GetAlbumArtists(context.Background()).MinCommunityRating(minCommunityRating).StartIndex(startIndex).Limit(limit).SearchTerm(searchTerm).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).Filters(filters).IsFavorite(isFavorite).MediaTypes(mediaTypes).Genres(genres).GenreIds(genreIds).OfficialRatings(officialRatings).Tags(tags).Years(years).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).UserId(userId).NameStartsWithOrGreater(nameStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).SortBy(sortBy).SortOrder(sortOrder).EnableImages(enableImages).EnableTotalRecordCount(enableTotalRecordCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.GetAlbumArtists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlbumArtists`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.GetAlbumArtists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlbumArtistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **searchTerm** | **string** | Optional. Search term. | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **excludeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be filtered out based on item type. This allows multiple, comma delimited. | 
 **includeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. | 
 **filters** | [**[]ItemFilter**](ItemFilter.md) | Optional. Specify additional filters to apply. | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **mediaTypes** | **[]string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **genres** | **[]string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimited. | 
 **genreIds** | **[]string** | Optional. If specified, results will be filtered based on genre id. This allows multiple, pipe delimited. | 
 **officialRatings** | **[]string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimited. | 
 **tags** | **[]string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimited. | 
 **years** | **[]int32** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimited. | 
 **enableUserData** | **bool** | Optional, include user data. | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **[]string** | Optional. If specified, results will be filtered to include only those containing the specified person ids. | 
 **personTypes** | **[]string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited. | 
 **studios** | **[]string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimited. | 
 **studioIds** | **[]string** | Optional. If specified, results will be filtered based on studio id. This allows multiple, pipe delimited. | 
 **userId** | **string** | User id. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 
 **sortBy** | **[]string** | Optional. Specify one or more sort orders, comma delimited. | 
 **sortOrder** | [**[]SortOrder**](SortOrder.md) | Sort Order - Ascending,Descending. | 
 **enableImages** | **bool** | Optional, include image information in output. | [default to true]
 **enableTotalRecordCount** | **bool** | Total record count. | [default to true]

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


## GetArtistByName

> BaseItemDto GetArtistByName(ctx, name).UserId(userId).Execute()

Gets an artist by name.

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
	name := "name_example" // string | Studio name.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user id, and attach user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtistsAPI.GetArtistByName(context.Background(), name).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.GetArtistByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtistByName`: BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.GetArtistByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Studio name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtistByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data. | 

### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtists

> BaseItemDtoQueryResult GetArtists(ctx).MinCommunityRating(minCommunityRating).StartIndex(startIndex).Limit(limit).SearchTerm(searchTerm).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).Filters(filters).IsFavorite(isFavorite).MediaTypes(mediaTypes).Genres(genres).GenreIds(genreIds).OfficialRatings(officialRatings).Tags(tags).Years(years).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).UserId(userId).NameStartsWithOrGreater(nameStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).SortBy(sortBy).SortOrder(sortOrder).EnableImages(enableImages).EnableTotalRecordCount(enableTotalRecordCount).Execute()

Gets all artists from a given item, folder, or the entire library.

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
	minCommunityRating := float64(1.2) // float64 | Optional filter by minimum community rating. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	searchTerm := "searchTerm_example" // string | Optional. Search term. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Specify this to localize the search to a specific item or folder. Omit to use the root. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	excludeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be filtered out based on item type. This allows multiple, comma delimited. (optional)
	includeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. (optional)
	filters := []openapiclient.ItemFilter{openapiclient.ItemFilter("IsFolder")} // []ItemFilter | Optional. Specify additional filters to apply. (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. (optional)
	mediaTypes := []string{"Inner_example"} // []string | Optional filter by MediaType. Allows multiple, comma delimited. (optional)
	genres := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimited. (optional)
	genreIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on genre id. This allows multiple, pipe delimited. (optional)
	officialRatings := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimited. (optional)
	tags := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimited. (optional)
	years := []int32{int32(123)} // []int32 | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimited. (optional)
	enableUserData := true // bool | Optional, include user data. (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)
	person := "person_example" // string | Optional. If specified, results will be filtered to include only those containing the specified person. (optional)
	personIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered to include only those containing the specified person ids. (optional)
	personTypes := []string{"Inner_example"} // []string | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited. (optional)
	studios := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimited. (optional)
	studioIds := []string{"Inner_example"} // []string | Optional. If specified, results will be filtered based on studio id. This allows multiple, pipe delimited. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id. (optional)
	nameStartsWithOrGreater := "nameStartsWithOrGreater_example" // string | Optional filter by items whose name is sorted equally or greater than a given input string. (optional)
	nameStartsWith := "nameStartsWith_example" // string | Optional filter by items whose name is sorted equally than a given input string. (optional)
	nameLessThan := "nameLessThan_example" // string | Optional filter by items whose name is equally or lesser than a given input string. (optional)
	sortBy := []string{"Inner_example"} // []string | Optional. Specify one or more sort orders, comma delimited. (optional)
	sortOrder := []openapiclient.SortOrder{openapiclient.SortOrder("Ascending")} // []SortOrder | Sort Order - Ascending,Descending. (optional)
	enableImages := true // bool | Optional, include image information in output. (optional) (default to true)
	enableTotalRecordCount := true // bool | Total record count. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtistsAPI.GetArtists(context.Background()).MinCommunityRating(minCommunityRating).StartIndex(startIndex).Limit(limit).SearchTerm(searchTerm).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).Filters(filters).IsFavorite(isFavorite).MediaTypes(mediaTypes).Genres(genres).GenreIds(genreIds).OfficialRatings(officialRatings).Tags(tags).Years(years).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Person(person).PersonIds(personIds).PersonTypes(personTypes).Studios(studios).StudioIds(studioIds).UserId(userId).NameStartsWithOrGreater(nameStartsWithOrGreater).NameStartsWith(nameStartsWith).NameLessThan(nameLessThan).SortBy(sortBy).SortOrder(sortOrder).EnableImages(enableImages).EnableTotalRecordCount(enableTotalRecordCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.GetArtists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtists`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.GetArtists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetArtistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minCommunityRating** | **float64** | Optional filter by minimum community rating. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **searchTerm** | **string** | Optional. Search term. | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **excludeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be filtered out based on item type. This allows multiple, comma delimited. | 
 **includeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. | 
 **filters** | [**[]ItemFilter**](ItemFilter.md) | Optional. Specify additional filters to apply. | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. | 
 **mediaTypes** | **[]string** | Optional filter by MediaType. Allows multiple, comma delimited. | 
 **genres** | **[]string** | Optional. If specified, results will be filtered based on genre. This allows multiple, pipe delimited. | 
 **genreIds** | **[]string** | Optional. If specified, results will be filtered based on genre id. This allows multiple, pipe delimited. | 
 **officialRatings** | **[]string** | Optional. If specified, results will be filtered based on OfficialRating. This allows multiple, pipe delimited. | 
 **tags** | **[]string** | Optional. If specified, results will be filtered based on tag. This allows multiple, pipe delimited. | 
 **years** | **[]int32** | Optional. If specified, results will be filtered based on production year. This allows multiple, comma delimited. | 
 **enableUserData** | **bool** | Optional, include user data. | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 
 **person** | **string** | Optional. If specified, results will be filtered to include only those containing the specified person. | 
 **personIds** | **[]string** | Optional. If specified, results will be filtered to include only those containing the specified person ids. | 
 **personTypes** | **[]string** | Optional. If specified, along with Person, results will be filtered to include only those containing the specified person and PersonType. Allows multiple, comma-delimited. | 
 **studios** | **[]string** | Optional. If specified, results will be filtered based on studio. This allows multiple, pipe delimited. | 
 **studioIds** | **[]string** | Optional. If specified, results will be filtered based on studio id. This allows multiple, pipe delimited. | 
 **userId** | **string** | User id. | 
 **nameStartsWithOrGreater** | **string** | Optional filter by items whose name is sorted equally or greater than a given input string. | 
 **nameStartsWith** | **string** | Optional filter by items whose name is sorted equally than a given input string. | 
 **nameLessThan** | **string** | Optional filter by items whose name is equally or lesser than a given input string. | 
 **sortBy** | **[]string** | Optional. Specify one or more sort orders, comma delimited. | 
 **sortOrder** | [**[]SortOrder**](SortOrder.md) | Sort Order - Ascending,Descending. | 
 **enableImages** | **bool** | Optional, include image information in output. | [default to true]
 **enableTotalRecordCount** | **bool** | Total record count. | [default to true]

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

