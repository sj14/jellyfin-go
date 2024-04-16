# \RemoteImageAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadRemoteImage**](RemoteImageAPI.md#DownloadRemoteImage) | **Post** /Items/{itemId}/RemoteImages/Download | Downloads a remote image for an item.
[**GetRemoteImageProviders**](RemoteImageAPI.md#GetRemoteImageProviders) | **Get** /Items/{itemId}/RemoteImages/Providers | Gets available remote image providers for an item.
[**GetRemoteImages**](RemoteImageAPI.md#GetRemoteImages) | **Get** /Items/{itemId}/RemoteImages | Gets available remote images for an item.



## DownloadRemoteImage

> DownloadRemoteImage(ctx, itemId).Type_(type_).ImageUrl(imageUrl).Execute()

Downloads a remote image for an item.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item Id.
	type_ := openapiclient.ImageType("Primary") // ImageType | The image type.
	imageUrl := "imageUrl_example" // string | The image url. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RemoteImageAPI.DownloadRemoteImage(context.Background(), itemId).Type_(type_).ImageUrl(imageUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteImageAPI.DownloadRemoteImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRemoteImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**ImageType**](ImageType.md) | The image type. | 
 **imageUrl** | **string** | The image url. | 

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


## GetRemoteImageProviders

> []ImageProviderInfo GetRemoteImageProviders(ctx, itemId).Execute()

Gets available remote image providers for an item.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item Id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteImageAPI.GetRemoteImageProviders(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteImageAPI.GetRemoteImageProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteImageProviders`: []ImageProviderInfo
	fmt.Fprintf(os.Stdout, "Response from `RemoteImageAPI.GetRemoteImageProviders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteImageProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ImageProviderInfo**](ImageProviderInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteImages

> RemoteImageResult GetRemoteImages(ctx, itemId).Type_(type_).StartIndex(startIndex).Limit(limit).ProviderName(providerName).IncludeAllLanguages(includeAllLanguages).Execute()

Gets available remote images for an item.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item Id.
	type_ := openapiclient.ImageType("Primary") // ImageType | The image type. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	providerName := "providerName_example" // string | Optional. The image provider to use. (optional)
	includeAllLanguages := true // bool | Optional. Include all languages. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteImageAPI.GetRemoteImages(context.Background(), itemId).Type_(type_).StartIndex(startIndex).Limit(limit).ProviderName(providerName).IncludeAllLanguages(includeAllLanguages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteImageAPI.GetRemoteImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteImages`: RemoteImageResult
	fmt.Fprintf(os.Stdout, "Response from `RemoteImageAPI.GetRemoteImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | [**ImageType**](ImageType.md) | The image type. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **providerName** | **string** | Optional. The image provider to use. | 
 **includeAllLanguages** | **bool** | Optional. Include all languages. | [default to false]

### Return type

[**RemoteImageResult**](RemoteImageResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

