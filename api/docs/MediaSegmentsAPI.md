# \MediaSegmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetItemSegments**](MediaSegmentsAPI.md#GetItemSegments) | **Get** /MediaSegments/{itemId} | Gets all media segments based on an itemId.



## GetItemSegments

> MediaSegmentDtoQueryResult GetItemSegments(ctx, itemId).IncludeSegmentTypes(includeSegmentTypes).Execute()

Gets all media segments based on an itemId.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ItemId.
	includeSegmentTypes := []openapiclient.MediaSegmentType{openapiclient.MediaSegmentType("Unknown")} // []MediaSegmentType | Optional filter of requested segment types. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaSegmentsAPI.GetItemSegments(context.Background(), itemId).IncludeSegmentTypes(includeSegmentTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaSegmentsAPI.GetItemSegments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemSegments`: MediaSegmentDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `MediaSegmentsAPI.GetItemSegments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | The ItemId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemSegmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSegmentTypes** | [**[]MediaSegmentType**](MediaSegmentType.md) | Optional filter of requested segment types. | 

### Return type

[**MediaSegmentDtoQueryResult**](MediaSegmentDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

