# \ActivityLogAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogEntries**](ActivityLogAPI.md#GetLogEntries) | **Get** /System/ActivityLog/Entries | Gets activity log entries.



## GetLogEntries

> ActivityLogEntryQueryResult GetLogEntries(ctx).StartIndex(startIndex).Limit(limit).MinDate(minDate).HasUserId(hasUserId).Execute()

Gets activity log entries.

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
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	minDate := time.Now() // time.Time | Optional. The minimum date. Format = ISO. (optional)
	hasUserId := true // bool | Optional. Filter log entries if it has user id, or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityLogAPI.GetLogEntries(context.Background()).StartIndex(startIndex).Limit(limit).MinDate(minDate).HasUserId(hasUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityLogAPI.GetLogEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogEntries`: ActivityLogEntryQueryResult
	fmt.Fprintf(os.Stdout, "Response from `ActivityLogAPI.GetLogEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **minDate** | **time.Time** | Optional. The minimum date. Format &#x3D; ISO. | 
 **hasUserId** | **bool** | Optional. Filter log entries if it has user id, or not. | 

### Return type

[**ActivityLogEntryQueryResult**](ActivityLogEntryQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

