# \PackageAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPackageInstallation**](PackageAPI.md#CancelPackageInstallation) | **Delete** /Packages/Installing/{packageId} | Cancels a package installation.
[**GetPackageInfo**](PackageAPI.md#GetPackageInfo) | **Get** /Packages/{name} | Gets a package by name or assembly GUID.
[**GetPackages**](PackageAPI.md#GetPackages) | **Get** /Packages | Gets available packages.
[**GetRepositories**](PackageAPI.md#GetRepositories) | **Get** /Repositories | Gets all package repositories.
[**InstallPackage**](PackageAPI.md#InstallPackage) | **Post** /Packages/Installed/{name} | Installs a package.
[**SetRepositories**](PackageAPI.md#SetRepositories) | **Post** /Repositories | Sets the enabled and existing package repositories.



## CancelPackageInstallation

> CancelPackageInstallation(ctx, packageId).Execute()

Cancels a package installation.

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
	packageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Installation Id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PackageAPI.CancelPackageInstallation(context.Background(), packageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageAPI.CancelPackageInstallation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | **string** | Installation Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPackageInstallationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackageInfo

> PackageInfo GetPackageInfo(ctx, name).AssemblyGuid(assemblyGuid).Execute()

Gets a package by name or assembly GUID.

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
	name := "name_example" // string | The name of the package.
	assemblyGuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The GUID of the associated assembly. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageAPI.GetPackageInfo(context.Background(), name).AssemblyGuid(assemblyGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageAPI.GetPackageInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackageInfo`: PackageInfo
	fmt.Fprintf(os.Stdout, "Response from `PackageAPI.GetPackageInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the package. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackageInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assemblyGuid** | **string** | The GUID of the associated assembly. | 

### Return type

[**PackageInfo**](PackageInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackages

> []PackageInfo GetPackages(ctx).Execute()

Gets available packages.

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
	resp, r, err := apiClient.PackageAPI.GetPackages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageAPI.GetPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackages`: []PackageInfo
	fmt.Fprintf(os.Stdout, "Response from `PackageAPI.GetPackages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackagesRequest struct via the builder pattern


### Return type

[**[]PackageInfo**](PackageInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositories

> []RepositoryInfo GetRepositories(ctx).Execute()

Gets all package repositories.

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
	resp, r, err := apiClient.PackageAPI.GetRepositories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageAPI.GetRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositories`: []RepositoryInfo
	fmt.Fprintf(os.Stdout, "Response from `PackageAPI.GetRepositories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoriesRequest struct via the builder pattern


### Return type

[**[]RepositoryInfo**](RepositoryInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallPackage

> InstallPackage(ctx, name).AssemblyGuid(assemblyGuid).Version(version).RepositoryUrl(repositoryUrl).Execute()

Installs a package.

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
	name := "name_example" // string | Package name.
	assemblyGuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | GUID of the associated assembly. (optional)
	version := "version_example" // string | Optional version. Defaults to latest version. (optional)
	repositoryUrl := "repositoryUrl_example" // string | Optional. Specify the repository to install from. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PackageAPI.InstallPackage(context.Background(), name).AssemblyGuid(assemblyGuid).Version(version).RepositoryUrl(repositoryUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageAPI.InstallPackage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Package name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assemblyGuid** | **string** | GUID of the associated assembly. | 
 **version** | **string** | Optional version. Defaults to latest version. | 
 **repositoryUrl** | **string** | Optional. Specify the repository to install from. | 

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


## SetRepositories

> SetRepositories(ctx).RepositoryInfo(repositoryInfo).Execute()

Sets the enabled and existing package repositories.

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
	repositoryInfo := []openapiclient.RepositoryInfo{*openapiclient.NewRepositoryInfo()} // []RepositoryInfo | The list of package repositories.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PackageAPI.SetRepositories(context.Background()).RepositoryInfo(repositoryInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageAPI.SetRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repositoryInfo** | [**[]RepositoryInfo**](RepositoryInfo.md) | The list of package repositories. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

