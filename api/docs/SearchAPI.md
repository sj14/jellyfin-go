# \SearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSearchHints**](SearchAPI.md#GetSearchHints) | **Get** /Search/Hints | Gets the search hint result.



## GetSearchHints

> SearchHintResult GetSearchHints(ctx).SearchTerm(searchTerm).StartIndex(startIndex).Limit(limit).UserId(userId).IncludeItemTypes(includeItemTypes).ExcludeItemTypes(excludeItemTypes).MediaTypes(mediaTypes).ParentId(parentId).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).IncludePeople(includePeople).IncludeMedia(includeMedia).IncludeGenres(includeGenres).IncludeStudios(includeStudios).IncludeArtists(includeArtists).Execute()

Gets the search hint result.

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
	searchTerm := "searchTerm_example" // string | The search term to filter on.
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Supply a user id to search within a user's library or omit to search all. (optional)
	includeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | If specified, only results with the specified item types are returned. This allows multiple, comma delimited. (optional)
	excludeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | If specified, results with these item types are filtered out. This allows multiple, comma delimited. (optional)
	mediaTypes := []openapiclient.MediaType{openapiclient.MediaType("Unknown")} // []MediaType | If specified, only results with the specified media types are returned. This allows multiple, comma delimited. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | If specified, only children of the parent are returned. (optional)
	isMovie := true // bool | Optional filter for movies. (optional)
	isSeries := true // bool | Optional filter for series. (optional)
	isNews := true // bool | Optional filter for news. (optional)
	isKids := true // bool | Optional filter for kids. (optional)
	isSports := true // bool | Optional filter for sports. (optional)
	includePeople := true // bool | Optional filter whether to include people. (optional) (default to true)
	includeMedia := true // bool | Optional filter whether to include media. (optional) (default to true)
	includeGenres := true // bool | Optional filter whether to include genres. (optional) (default to true)
	includeStudios := true // bool | Optional filter whether to include studios. (optional) (default to true)
	includeArtists := true // bool | Optional filter whether to include artists. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.GetSearchHints(context.Background()).SearchTerm(searchTerm).StartIndex(startIndex).Limit(limit).UserId(userId).IncludeItemTypes(includeItemTypes).ExcludeItemTypes(excludeItemTypes).MediaTypes(mediaTypes).ParentId(parentId).IsMovie(isMovie).IsSeries(isSeries).IsNews(isNews).IsKids(isKids).IsSports(isSports).IncludePeople(includePeople).IncludeMedia(includeMedia).IncludeGenres(includeGenres).IncludeStudios(includeStudios).IncludeArtists(includeArtists).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetSearchHints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSearchHints`: SearchHintResult
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetSearchHints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchHintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string** | The search term to filter on. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **userId** | **string** | Optional. Supply a user id to search within a user&#39;s library or omit to search all. | 
 **includeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | If specified, only results with the specified item types are returned. This allows multiple, comma delimited. | 
 **excludeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | If specified, results with these item types are filtered out. This allows multiple, comma delimited. | 
 **mediaTypes** | [**[]MediaType**](MediaType.md) | If specified, only results with the specified media types are returned. This allows multiple, comma delimited. | 
 **parentId** | **string** | If specified, only children of the parent are returned. | 
 **isMovie** | **bool** | Optional filter for movies. | 
 **isSeries** | **bool** | Optional filter for series. | 
 **isNews** | **bool** | Optional filter for news. | 
 **isKids** | **bool** | Optional filter for kids. | 
 **isSports** | **bool** | Optional filter for sports. | 
 **includePeople** | **bool** | Optional filter whether to include people. | [default to true]
 **includeMedia** | **bool** | Optional filter whether to include media. | [default to true]
 **includeGenres** | **bool** | Optional filter whether to include genres. | [default to true]
 **includeStudios** | **bool** | Optional filter whether to include studios. | [default to true]
 **includeArtists** | **bool** | Optional filter whether to include artists. | [default to true]

### Return type

[**SearchHintResult**](SearchHintResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

