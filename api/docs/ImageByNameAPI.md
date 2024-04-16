# \ImageByNameAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGeneralImage**](ImageByNameAPI.md#GetGeneralImage) | **Get** /Images/General/{name}/{type} | Get General Image.
[**GetGeneralImages**](ImageByNameAPI.md#GetGeneralImages) | **Get** /Images/General | Get all general images.
[**GetMediaInfoImage**](ImageByNameAPI.md#GetMediaInfoImage) | **Get** /Images/MediaInfo/{theme}/{name} | Get media info image.
[**GetMediaInfoImages**](ImageByNameAPI.md#GetMediaInfoImages) | **Get** /Images/MediaInfo | Get all media info images.
[**GetRatingImage**](ImageByNameAPI.md#GetRatingImage) | **Get** /Images/Ratings/{theme}/{name} | Get rating image.
[**GetRatingImages**](ImageByNameAPI.md#GetRatingImages) | **Get** /Images/Ratings | Get all general images.



## GetGeneralImage

> *os.File GetGeneralImage(ctx, name, type_).Execute()

Get General Image.

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
	name := "name_example" // string | The name of the image.
	type_ := "type__example" // string | Image Type (primary, backdrop, logo, etc).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageByNameAPI.GetGeneralImage(context.Background(), name, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageByNameAPI.GetGeneralImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGeneralImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageByNameAPI.GetGeneralImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the image. | 
**type_** | **string** | Image Type (primary, backdrop, logo, etc). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGeneralImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/octet-stream, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGeneralImages

> []ImageByNameInfo GetGeneralImages(ctx).Execute()

Get all general images.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageByNameAPI.GetGeneralImages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageByNameAPI.GetGeneralImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGeneralImages`: []ImageByNameInfo
	fmt.Fprintf(os.Stdout, "Response from `ImageByNameAPI.GetGeneralImages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGeneralImagesRequest struct via the builder pattern


### Return type

[**[]ImageByNameInfo**](ImageByNameInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaInfoImage

> *os.File GetMediaInfoImage(ctx, theme, name).Execute()

Get media info image.

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
	theme := "theme_example" // string | The theme to get the image from.
	name := "name_example" // string | The name of the image.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageByNameAPI.GetMediaInfoImage(context.Background(), theme, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageByNameAPI.GetMediaInfoImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaInfoImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageByNameAPI.GetMediaInfoImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**theme** | **string** | The theme to get the image from. | 
**name** | **string** | The name of the image. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaInfoImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/octet-stream, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaInfoImages

> []ImageByNameInfo GetMediaInfoImages(ctx).Execute()

Get all media info images.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageByNameAPI.GetMediaInfoImages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageByNameAPI.GetMediaInfoImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaInfoImages`: []ImageByNameInfo
	fmt.Fprintf(os.Stdout, "Response from `ImageByNameAPI.GetMediaInfoImages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaInfoImagesRequest struct via the builder pattern


### Return type

[**[]ImageByNameInfo**](ImageByNameInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRatingImage

> *os.File GetRatingImage(ctx, theme, name).Execute()

Get rating image.

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
	theme := "theme_example" // string | The theme to get the image from.
	name := "name_example" // string | The name of the image.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageByNameAPI.GetRatingImage(context.Background(), theme, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageByNameAPI.GetRatingImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRatingImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageByNameAPI.GetRatingImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**theme** | **string** | The theme to get the image from. | 
**name** | **string** | The name of the image. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRatingImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/octet-stream, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRatingImages

> []ImageByNameInfo GetRatingImages(ctx).Execute()

Get all general images.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageByNameAPI.GetRatingImages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageByNameAPI.GetRatingImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRatingImages`: []ImageByNameInfo
	fmt.Fprintf(os.Stdout, "Response from `ImageByNameAPI.GetRatingImages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRatingImagesRequest struct via the builder pattern


### Return type

[**[]ImageByNameInfo**](ImageByNameInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

