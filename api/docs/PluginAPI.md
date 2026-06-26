# \PluginAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPackageInstallation**](PluginAPI.md#CancelPackageInstallation) | **Delete** /Packages/Installing/{packageId} | Cancels a package installation.
[**DisablePlugin**](PluginAPI.md#DisablePlugin) | **Post** /Plugins/{pluginId}/{version}/Disable | Disable a plugin.
[**EnablePlugin**](PluginAPI.md#EnablePlugin) | **Post** /Plugins/{pluginId}/{version}/Enable | Enables a disabled plugin.
[**GetConfigurationPages**](PluginAPI.md#GetConfigurationPages) | **Get** /web/ConfigurationPages | Gets the configuration pages.
[**GetDashboardConfigurationPage**](PluginAPI.md#GetDashboardConfigurationPage) | **Get** /web/ConfigurationPage | Gets a dashboard configuration page.
[**GetPackageInfo**](PluginAPI.md#GetPackageInfo) | **Get** /Packages/{name} | Gets a package by name or assembly GUID.
[**GetPackages**](PluginAPI.md#GetPackages) | **Get** /Packages | Gets available packages.
[**GetPluginConfiguration**](PluginAPI.md#GetPluginConfiguration) | **Get** /Plugins/{pluginId}/Configuration | Gets plugin configuration.
[**GetPluginImage**](PluginAPI.md#GetPluginImage) | **Get** /Plugins/{pluginId}/{version}/Image | Gets a plugin&#39;s image.
[**GetPluginManifest**](PluginAPI.md#GetPluginManifest) | **Post** /Plugins/{pluginId}/Manifest | Gets a plugin&#39;s manifest.
[**GetPlugins**](PluginAPI.md#GetPlugins) | **Get** /Plugins | Gets a list of currently installed plugins.
[**GetRepositories**](PluginAPI.md#GetRepositories) | **Get** /Repositories | Gets all package repositories.
[**InstallPackage**](PluginAPI.md#InstallPackage) | **Post** /Packages/Installed/{name} | Installs a package.
[**SetRepositories**](PluginAPI.md#SetRepositories) | **Post** /Repositories | Sets the enabled and existing package repositories.
[**UninstallPlugin**](PluginAPI.md#UninstallPlugin) | **Delete** /Plugins/{pluginId} | Uninstalls a plugin.
[**UninstallPluginByVersion**](PluginAPI.md#UninstallPluginByVersion) | **Delete** /Plugins/{pluginId}/{version} | Uninstalls a plugin by version.
[**UpdatePluginConfiguration**](PluginAPI.md#UpdatePluginConfiguration) | **Post** /Plugins/{pluginId}/Configuration | Updates plugin configuration.



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
	r, err := apiClient.PluginAPI.CancelPackageInstallation(context.Background(), packageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.CancelPackageInstallation``: %v\n", err)
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
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisablePlugin

> DisablePlugin(ctx, pluginId, version).Execute()

Disable a plugin.

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
	pluginId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Plugin id.
	version := "version_example" // string | Plugin version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginAPI.DisablePlugin(context.Background(), pluginId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.DisablePlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin id. | 
**version** | **string** | Plugin version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisablePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnablePlugin

> EnablePlugin(ctx, pluginId, version).Execute()

Enables a disabled plugin.

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
	pluginId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Plugin id.
	version := "version_example" // string | Plugin version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginAPI.EnablePlugin(context.Background(), pluginId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.EnablePlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin id. | 
**version** | **string** | Plugin version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnablePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigurationPages

> []ConfigurationPageInfo GetConfigurationPages(ctx).EnableInMainMenu(enableInMainMenu).Execute()

Gets the configuration pages.

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
	enableInMainMenu := true // bool | Whether to enable in the main menu. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginAPI.GetConfigurationPages(context.Background()).EnableInMainMenu(enableInMainMenu).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetConfigurationPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigurationPages`: []ConfigurationPageInfo
	fmt.Fprintf(os.Stdout, "Response from `PluginAPI.GetConfigurationPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableInMainMenu** | **bool** | Whether to enable in the main menu. | 

### Return type

[**[]ConfigurationPageInfo**](ConfigurationPageInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardConfigurationPage

> *os.File GetDashboardConfigurationPage(ctx).Name(name).Execute()

Gets a dashboard configuration page.

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
	name := "name_example" // string | The name of the page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginAPI.GetDashboardConfigurationPage(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetDashboardConfigurationPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardConfigurationPage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `PluginAPI.GetDashboardConfigurationPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardConfigurationPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the page. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html, application/x-javascript, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

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
	resp, r, err := apiClient.PluginAPI.GetPackageInfo(context.Background(), name).AssemblyGuid(assemblyGuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetPackageInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackageInfo`: PackageInfo
	fmt.Fprintf(os.Stdout, "Response from `PluginAPI.GetPackageInfo`: %v\n", resp)
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
	resp, r, err := apiClient.PluginAPI.GetPackages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackages`: []PackageInfo
	fmt.Fprintf(os.Stdout, "Response from `PluginAPI.GetPackages`: %v\n", resp)
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginConfiguration

> map[string]interface{} GetPluginConfiguration(ctx, pluginId).Execute()

Gets plugin configuration.

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
	pluginId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Plugin id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginAPI.GetPluginConfiguration(context.Background(), pluginId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetPluginConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPluginConfiguration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PluginAPI.GetPluginConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginImage

> *os.File GetPluginImage(ctx, pluginId, version).Execute()

Gets a plugin's image.

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
	pluginId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Plugin id.
	version := "version_example" // string | Plugin version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginAPI.GetPluginImage(context.Background(), pluginId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetPluginImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPluginImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `PluginAPI.GetPluginImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin id. | 
**version** | **string** | Plugin version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPluginManifest

> GetPluginManifest(ctx, pluginId).Execute()

Gets a plugin's manifest.

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
	pluginId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Plugin id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginAPI.GetPluginManifest(context.Background(), pluginId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetPluginManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlugins

> []PluginInfo GetPlugins(ctx).Execute()

Gets a list of currently installed plugins.

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
	resp, r, err := apiClient.PluginAPI.GetPlugins(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlugins`: []PluginInfo
	fmt.Fprintf(os.Stdout, "Response from `PluginAPI.GetPlugins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginsRequest struct via the builder pattern


### Return type

[**[]PluginInfo**](PluginInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
	resp, r, err := apiClient.PluginAPI.GetRepositories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositories`: []RepositoryInfo
	fmt.Fprintf(os.Stdout, "Response from `PluginAPI.GetRepositories`: %v\n", resp)
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
	r, err := apiClient.PluginAPI.InstallPackage(context.Background(), name).AssemblyGuid(assemblyGuid).Version(version).RepositoryUrl(repositoryUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.InstallPackage``: %v\n", err)
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

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
	r, err := apiClient.PluginAPI.SetRepositories(context.Background()).RepositoryInfo(repositoryInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.SetRepositories``: %v\n", err)
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
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallPlugin

> UninstallPlugin(ctx, pluginId).Execute()

Uninstalls a plugin.

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
	pluginId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Plugin id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginAPI.UninstallPlugin(context.Background(), pluginId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.UninstallPlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallPluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UninstallPluginByVersion

> UninstallPluginByVersion(ctx, pluginId, version).Execute()

Uninstalls a plugin by version.

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
	pluginId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Plugin id.
	version := "version_example" // string | Plugin version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginAPI.UninstallPluginByVersion(context.Background(), pluginId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.UninstallPluginByVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin id. | 
**version** | **string** | Plugin version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallPluginByVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePluginConfiguration

> UpdatePluginConfiguration(ctx, pluginId).Execute()

Updates plugin configuration.



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
	pluginId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Plugin id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PluginAPI.UpdatePluginConfiguration(context.Background(), pluginId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.UpdatePluginConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePluginConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

