# \FilterAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQueryFilters**](FilterAPI.md#GetQueryFilters) | **Get** /Items/Filters2 | Gets query filters.
[**GetQueryFiltersLegacy**](FilterAPI.md#GetQueryFiltersLegacy) | **Get** /Items/Filters | Gets legacy query filters.



## GetQueryFilters

> QueryFilters GetQueryFilters(ctx).UserId(userId).ParentId(parentId).IncludeItemTypes(includeItemTypes).IsAiring(isAiring).IsMovie(isMovie).IsSports(isSports).IsKids(isKids).IsNews(isNews).IsSeries(isSeries).Recursive(recursive).Execute()

Gets query filters.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. User id. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Specify this to localize the search to a specific item or folder. Omit to use the root. (optional)
	includeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. (optional)
	isAiring := true // bool | Optional. Is item airing. (optional)
	isMovie := true // bool | Optional. Is item movie. (optional)
	isSports := true // bool | Optional. Is item sports. (optional)
	isKids := true // bool | Optional. Is item kids. (optional)
	isNews := true // bool | Optional. Is item news. (optional)
	isSeries := true // bool | Optional. Is item series. (optional)
	recursive := true // bool | Optional. Search recursive. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterAPI.GetQueryFilters(context.Background()).UserId(userId).ParentId(parentId).IncludeItemTypes(includeItemTypes).IsAiring(isAiring).IsMovie(isMovie).IsSports(isSports).IsKids(isKids).IsNews(isNews).IsSeries(isSeries).Recursive(recursive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.GetQueryFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueryFilters`: QueryFilters
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.GetQueryFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Optional. User id. | 
 **parentId** | **string** | Optional. Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **includeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. | 
 **isAiring** | **bool** | Optional. Is item airing. | 
 **isMovie** | **bool** | Optional. Is item movie. | 
 **isSports** | **bool** | Optional. Is item sports. | 
 **isKids** | **bool** | Optional. Is item kids. | 
 **isNews** | **bool** | Optional. Is item news. | 
 **isSeries** | **bool** | Optional. Is item series. | 
 **recursive** | **bool** | Optional. Search recursive. | 

### Return type

[**QueryFilters**](QueryFilters.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueryFiltersLegacy

> QueryFiltersLegacy GetQueryFiltersLegacy(ctx).UserId(userId).ParentId(parentId).IncludeItemTypes(includeItemTypes).MediaTypes(mediaTypes).Execute()

Gets legacy query filters.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. User id. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Parent id. (optional)
	includeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. (optional)
	mediaTypes := []string{"Inner_example"} // []string | Optional. Filter by MediaType. Allows multiple, comma delimited. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterAPI.GetQueryFiltersLegacy(context.Background()).UserId(userId).ParentId(parentId).IncludeItemTypes(includeItemTypes).MediaTypes(mediaTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterAPI.GetQueryFiltersLegacy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueryFiltersLegacy`: QueryFiltersLegacy
	fmt.Fprintf(os.Stdout, "Response from `FilterAPI.GetQueryFiltersLegacy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQueryFiltersLegacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Optional. User id. | 
 **parentId** | **string** | Optional. Parent id. | 
 **includeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. | 
 **mediaTypes** | **[]string** | Optional. Filter by MediaType. Allows multiple, comma delimited. | 

### Return type

[**QueryFiltersLegacy**](QueryFiltersLegacy.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

