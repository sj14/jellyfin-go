# \EnvironmentAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDefaultDirectoryBrowser**](EnvironmentAPI.md#GetDefaultDirectoryBrowser) | **Get** /Environment/DefaultDirectoryBrowser | Get Default directory browser.
[**GetDirectoryContents**](EnvironmentAPI.md#GetDirectoryContents) | **Get** /Environment/DirectoryContents | Gets the contents of a given directory in the file system.
[**GetDrives**](EnvironmentAPI.md#GetDrives) | **Get** /Environment/Drives | Gets available drives from the server&#39;s file system.
[**GetNetworkShares**](EnvironmentAPI.md#GetNetworkShares) | **Get** /Environment/NetworkShares | Gets network paths.
[**GetParentPath**](EnvironmentAPI.md#GetParentPath) | **Get** /Environment/ParentPath | Gets the parent path of a given path.
[**ValidatePath**](EnvironmentAPI.md#ValidatePath) | **Post** /Environment/ValidatePath | Validates path.



## GetDefaultDirectoryBrowser

> DefaultDirectoryBrowserInfoDto GetDefaultDirectoryBrowser(ctx).Execute()

Get Default directory browser.

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
	resp, r, err := apiClient.EnvironmentAPI.GetDefaultDirectoryBrowser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.GetDefaultDirectoryBrowser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultDirectoryBrowser`: DefaultDirectoryBrowserInfoDto
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.GetDefaultDirectoryBrowser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultDirectoryBrowserRequest struct via the builder pattern


### Return type

[**DefaultDirectoryBrowserInfoDto**](DefaultDirectoryBrowserInfoDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDirectoryContents

> []FileSystemEntryInfo GetDirectoryContents(ctx).Path(path).IncludeFiles(includeFiles).IncludeDirectories(includeDirectories).Execute()

Gets the contents of a given directory in the file system.

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
	path := "path_example" // string | The path.
	includeFiles := true // bool | An optional filter to include or exclude files from the results. true/false. (optional) (default to false)
	includeDirectories := true // bool | An optional filter to include or exclude folders from the results. true/false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.GetDirectoryContents(context.Background()).Path(path).IncludeFiles(includeFiles).IncludeDirectories(includeDirectories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.GetDirectoryContents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDirectoryContents`: []FileSystemEntryInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.GetDirectoryContents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDirectoryContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | The path. | 
 **includeFiles** | **bool** | An optional filter to include or exclude files from the results. true/false. | [default to false]
 **includeDirectories** | **bool** | An optional filter to include or exclude folders from the results. true/false. | [default to false]

### Return type

[**[]FileSystemEntryInfo**](FileSystemEntryInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDrives

> []FileSystemEntryInfo GetDrives(ctx).Execute()

Gets available drives from the server's file system.

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
	resp, r, err := apiClient.EnvironmentAPI.GetDrives(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.GetDrives``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDrives`: []FileSystemEntryInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.GetDrives`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDrivesRequest struct via the builder pattern


### Return type

[**[]FileSystemEntryInfo**](FileSystemEntryInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkShares

> []FileSystemEntryInfo GetNetworkShares(ctx).Execute()

Gets network paths.

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
	resp, r, err := apiClient.EnvironmentAPI.GetNetworkShares(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.GetNetworkShares``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkShares`: []FileSystemEntryInfo
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.GetNetworkShares`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSharesRequest struct via the builder pattern


### Return type

[**[]FileSystemEntryInfo**](FileSystemEntryInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParentPath

> string GetParentPath(ctx).Path(path).Execute()

Gets the parent path of a given path.

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
	path := "path_example" // string | The path.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.GetParentPath(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.GetParentPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParentPath`: string
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.GetParentPath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetParentPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | The path. | 

### Return type

**string**

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePath

> ValidatePath(ctx).ValidatePathDto(validatePathDto).Execute()

Validates path.

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
	validatePathDto := *openapiclient.NewValidatePathDto() // ValidatePathDto | Validate request object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EnvironmentAPI.ValidatePath(context.Background()).ValidatePathDto(validatePathDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.ValidatePath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidatePathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validatePathDto** | [**ValidatePathDto**](ValidatePathDto.md) | Validate request object. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

