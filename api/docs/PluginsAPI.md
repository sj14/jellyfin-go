# \PluginsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisablePlugin**](PluginsAPI.md#DisablePlugin) | **Post** /Plugins/{pluginId}/{version}/Disable | Disable a plugin.
[**EnablePlugin**](PluginsAPI.md#EnablePlugin) | **Post** /Plugins/{pluginId}/{version}/Enable | Enables a disabled plugin.
[**GetPluginConfiguration**](PluginsAPI.md#GetPluginConfiguration) | **Get** /Plugins/{pluginId}/Configuration | Gets plugin configuration.
[**GetPluginImage**](PluginsAPI.md#GetPluginImage) | **Get** /Plugins/{pluginId}/{version}/Image | Gets a plugin&#39;s image.
[**GetPluginManifest**](PluginsAPI.md#GetPluginManifest) | **Post** /Plugins/{pluginId}/Manifest | Gets a plugin&#39;s manifest.
[**GetPlugins**](PluginsAPI.md#GetPlugins) | **Get** /Plugins | Gets a list of currently installed plugins.
[**UninstallPlugin**](PluginsAPI.md#UninstallPlugin) | **Delete** /Plugins/{pluginId} | Uninstalls a plugin.
[**UninstallPluginByVersion**](PluginsAPI.md#UninstallPluginByVersion) | **Delete** /Plugins/{pluginId}/{version} | Uninstalls a plugin by version.
[**UpdatePluginConfiguration**](PluginsAPI.md#UpdatePluginConfiguration) | **Post** /Plugins/{pluginId}/Configuration | Updates plugin configuration.



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
	r, err := apiClient.PluginsAPI.DisablePlugin(context.Background(), pluginId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.DisablePlugin``: %v\n", err)
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

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
	r, err := apiClient.PluginsAPI.EnablePlugin(context.Background(), pluginId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.EnablePlugin``: %v\n", err)
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

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
	resp, r, err := apiClient.PluginsAPI.GetPluginConfiguration(context.Background(), pluginId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPluginConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPluginConfiguration`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPluginConfiguration`: %v\n", resp)
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

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
	resp, r, err := apiClient.PluginsAPI.GetPluginImage(context.Background(), pluginId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPluginImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPluginImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPluginImage`: %v\n", resp)
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
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

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
	r, err := apiClient.PluginsAPI.GetPluginManifest(context.Background(), pluginId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPluginManifest``: %v\n", err)
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

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
	resp, r, err := apiClient.PluginsAPI.GetPlugins(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlugins`: []PluginInfo
	fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPlugins`: %v\n", resp)
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

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
	r, err := apiClient.PluginsAPI.UninstallPlugin(context.Background(), pluginId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.UninstallPlugin``: %v\n", err)
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

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
	r, err := apiClient.PluginsAPI.UninstallPluginByVersion(context.Background(), pluginId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.UninstallPluginByVersion``: %v\n", err)
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

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
	r, err := apiClient.PluginsAPI.UpdatePluginConfiguration(context.Background(), pluginId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.UpdatePluginConfiguration``: %v\n", err)
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
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

