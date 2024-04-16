# \UserViewsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGroupingOptions**](UserViewsAPI.md#GetGroupingOptions) | **Get** /Users/{userId}/GroupingOptions | Get user view grouping options.
[**GetUserViews**](UserViewsAPI.md#GetUserViews) | **Get** /Users/{userId}/Views | Get user views.



## GetGroupingOptions

> []SpecialViewOptionDto GetGroupingOptions(ctx, userId).Execute()

Get user view grouping options.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserViewsAPI.GetGroupingOptions(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserViewsAPI.GetGroupingOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroupingOptions`: []SpecialViewOptionDto
	fmt.Fprintf(os.Stdout, "Response from `UserViewsAPI.GetGroupingOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupingOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SpecialViewOptionDto**](SpecialViewOptionDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserViews

> BaseItemDtoQueryResult GetUserViews(ctx, userId).IncludeExternalContent(includeExternalContent).PresetViews(presetViews).IncludeHidden(includeHidden).Execute()

Get user views.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id.
	includeExternalContent := true // bool | Whether or not to include external views such as channels or live tv. (optional)
	presetViews := []string{"Inner_example"} // []string | Preset views. (optional)
	includeHidden := true // bool | Whether or not to include hidden content. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserViewsAPI.GetUserViews(context.Background(), userId).IncludeExternalContent(includeExternalContent).PresetViews(presetViews).IncludeHidden(includeHidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserViewsAPI.GetUserViews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserViews`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `UserViewsAPI.GetUserViews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeExternalContent** | **bool** | Whether or not to include external views such as channels or live tv. | 
 **presetViews** | **[]string** | Preset views. | 
 **includeHidden** | **bool** | Whether or not to include hidden content. | [default to false]

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

