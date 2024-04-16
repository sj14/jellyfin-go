# \ItemRefreshAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RefreshItem**](ItemRefreshAPI.md#RefreshItem) | **Post** /Items/{itemId}/Refresh | Refreshes metadata for an item.



## RefreshItem

> RefreshItem(ctx, itemId).MetadataRefreshMode(metadataRefreshMode).ImageRefreshMode(imageRefreshMode).ReplaceAllMetadata(replaceAllMetadata).ReplaceAllImages(replaceAllImages).Execute()

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
	metadataRefreshMode := openapiclient.MetadataRefreshMode("None") // MetadataRefreshMode | (Optional) Specifies the metadata refresh mode. (optional) (default to "None")
	imageRefreshMode := openapiclient.MetadataRefreshMode("None") // MetadataRefreshMode | (Optional) Specifies the image refresh mode. (optional) (default to "None")
	replaceAllMetadata := true // bool | (Optional) Determines if metadata should be replaced. Only applicable if mode is FullRefresh. (optional) (default to false)
	replaceAllImages := true // bool | (Optional) Determines if images should be replaced. Only applicable if mode is FullRefresh. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemRefreshAPI.RefreshItem(context.Background(), itemId).MetadataRefreshMode(metadataRefreshMode).ImageRefreshMode(imageRefreshMode).ReplaceAllMetadata(replaceAllMetadata).ReplaceAllImages(replaceAllImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemRefreshAPI.RefreshItem``: %v\n", err)
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

 **metadataRefreshMode** | [**MetadataRefreshMode**](MetadataRefreshMode.md) | (Optional) Specifies the metadata refresh mode. | [default to &quot;None&quot;]
 **imageRefreshMode** | [**MetadataRefreshMode**](MetadataRefreshMode.md) | (Optional) Specifies the image refresh mode. | [default to &quot;None&quot;]
 **replaceAllMetadata** | **bool** | (Optional) Determines if metadata should be replaced. Only applicable if mode is FullRefresh. | [default to false]
 **replaceAllImages** | **bool** | (Optional) Determines if images should be replaced. Only applicable if mode is FullRefresh. | [default to false]

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

