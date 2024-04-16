# \SuggestionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSuggestions**](SuggestionsAPI.md#GetSuggestions) | **Get** /Users/{userId}/Suggestions | Gets suggestions.



## GetSuggestions

> BaseItemDtoQueryResult GetSuggestions(ctx, userId).MediaType(mediaType).Type_(type_).StartIndex(startIndex).Limit(limit).EnableTotalRecordCount(enableTotalRecordCount).Execute()

Gets suggestions.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.
	mediaType := []string{"Inner_example"} // []string | The media types. (optional)
	type_ := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | The type. (optional)
	startIndex := int32(56) // int32 | Optional. The start index. (optional)
	limit := int32(56) // int32 | Optional. The limit. (optional)
	enableTotalRecordCount := true // bool | Whether to enable the total record count. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SuggestionsAPI.GetSuggestions(context.Background(), userId).MediaType(mediaType).Type_(type_).StartIndex(startIndex).Limit(limit).EnableTotalRecordCount(enableTotalRecordCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SuggestionsAPI.GetSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuggestions`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `SuggestionsAPI.GetSuggestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mediaType** | **[]string** | The media types. | 
 **type_** | [**[]BaseItemKind**](BaseItemKind.md) | The type. | 
 **startIndex** | **int32** | Optional. The start index. | 
 **limit** | **int32** | Optional. The limit. | 
 **enableTotalRecordCount** | **bool** | Whether to enable the total record count. | [default to false]

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

