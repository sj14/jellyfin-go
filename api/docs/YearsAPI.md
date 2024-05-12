# \YearsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetYear**](YearsAPI.md#GetYear) | **Get** /Years/{year} | Gets a year.
[**GetYears**](YearsAPI.md#GetYears) | **Get** /Years | Get years.



## GetYear

> BaseItemDto GetYear(ctx, year).UserId(userId).Execute()

Gets a year.

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
	year := int32(56) // int32 | The year.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user id, and attach user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.YearsAPI.GetYear(context.Background(), year).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `YearsAPI.GetYear``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetYear`: BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `YearsAPI.GetYear`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**year** | **int32** | The year. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetYearRequest struct via the builder pattern


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


## GetYears

> BaseItemDtoQueryResult GetYears(ctx).StartIndex(startIndex).Limit(limit).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).MediaTypes(mediaTypes).SortBy(sortBy).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Recursive(recursive).EnableImages(enableImages).Execute()

Get years.

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
	startIndex := int32(56) // int32 | Skips over a given number of items within the results. Use for paging. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	sortOrder := []openapiclient.SortOrder{openapiclient.SortOrder("Ascending")} // []SortOrder | Sort Order - Ascending,Descending. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Specify this to localize the search to a specific item or folder. Omit to use the root. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	excludeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be excluded based on item type. This allows multiple, comma delimited. (optional)
	includeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be included based on item type. This allows multiple, comma delimited. (optional)
	mediaTypes := []openapiclient.MediaType{openapiclient.MediaType("Unknown")} // []MediaType | Optional. Filter by MediaType. Allows multiple, comma delimited. (optional)
	sortBy := []openapiclient.ItemSortBy{openapiclient.ItemSortBy("Default")} // []ItemSortBy | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. (optional)
	enableUserData := true // bool | Optional. Include user data. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. The max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User Id. (optional)
	recursive := true // bool | Search recursively. (optional) (default to true)
	enableImages := true // bool | Optional. Include image information in output. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.YearsAPI.GetYears(context.Background()).StartIndex(startIndex).Limit(limit).SortOrder(sortOrder).ParentId(parentId).Fields(fields).ExcludeItemTypes(excludeItemTypes).IncludeItemTypes(includeItemTypes).MediaTypes(mediaTypes).SortBy(sortBy).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).UserId(userId).Recursive(recursive).EnableImages(enableImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `YearsAPI.GetYears``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetYears`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `YearsAPI.GetYears`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetYearsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **int32** | Skips over a given number of items within the results. Use for paging. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **sortOrder** | [**[]SortOrder**](SortOrder.md) | Sort Order - Ascending,Descending. | 
 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **excludeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be excluded based on item type. This allows multiple, comma delimited. | 
 **includeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be included based on item type. This allows multiple, comma delimited. | 
 **mediaTypes** | [**[]MediaType**](MediaType.md) | Optional. Filter by MediaType. Allows multiple, comma delimited. | 
 **sortBy** | [**[]ItemSortBy**](ItemSortBy.md) | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. | 
 **enableUserData** | **bool** | Optional. Include user data. | 
 **imageTypeLimit** | **int32** | Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 
 **userId** | **string** | User Id. | 
 **recursive** | **bool** | Search recursively. | [default to true]
 **enableImages** | **bool** | Optional. Include image information in output. | [default to true]

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

