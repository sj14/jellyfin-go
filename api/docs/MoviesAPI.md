# \MoviesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMovieRecommendations**](MoviesAPI.md#GetMovieRecommendations) | **Get** /Movies/Recommendations | Gets movie recommendations.



## GetMovieRecommendations

> []RecommendationDto GetMovieRecommendations(ctx).UserId(userId).ParentId(parentId).Fields(fields).CategoryLimit(categoryLimit).ItemLimit(itemLimit).Execute()

Gets movie recommendations.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user id, and attach user data. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Specify this to localize the search to a specific item or folder. Omit to use the root. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. The fields to return. (optional)
	categoryLimit := int32(56) // int32 | The max number of categories to return. (optional) (default to 5)
	itemLimit := int32(56) // int32 | The max number of items to return per category. (optional) (default to 8)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MoviesAPI.GetMovieRecommendations(context.Background()).UserId(userId).ParentId(parentId).Fields(fields).CategoryLimit(categoryLimit).ItemLimit(itemLimit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MoviesAPI.GetMovieRecommendations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMovieRecommendations`: []RecommendationDto
	fmt.Fprintf(os.Stdout, "Response from `MoviesAPI.GetMovieRecommendations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMovieRecommendationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Optional. Filter by user id, and attach user data. | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. The fields to return. | 
 **categoryLimit** | **int32** | The max number of categories to return. | [default to 5]
 **itemLimit** | **int32** | The max number of items to return per category. | [default to 8]

### Return type

[**[]RecommendationDto**](RecommendationDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

