# \StartupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteWizard**](StartupAPI.md#CompleteWizard) | **Post** /Startup/Complete | Completes the startup wizard.
[**GetFirstUser**](StartupAPI.md#GetFirstUser) | **Get** /Startup/User | Gets the first user.
[**GetFirstUser2**](StartupAPI.md#GetFirstUser2) | **Get** /Startup/FirstUser | Gets the first user.
[**GetStartupConfiguration**](StartupAPI.md#GetStartupConfiguration) | **Get** /Startup/Configuration | Gets the initial startup wizard configuration.
[**SetRemoteAccess**](StartupAPI.md#SetRemoteAccess) | **Post** /Startup/RemoteAccess | Sets remote access and UPnP.
[**UpdateInitialConfiguration**](StartupAPI.md#UpdateInitialConfiguration) | **Post** /Startup/Configuration | Sets the initial startup wizard configuration.
[**UpdateStartupUser**](StartupAPI.md#UpdateStartupUser) | **Post** /Startup/User | Sets the user name and password.



## CompleteWizard

> CompleteWizard(ctx).Execute()

Completes the startup wizard.

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
	r, err := apiClient.StartupAPI.CompleteWizard(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StartupAPI.CompleteWizard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteWizardRequest struct via the builder pattern


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


## GetFirstUser

> StartupUserDto GetFirstUser(ctx).Execute()

Gets the first user.

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
	resp, r, err := apiClient.StartupAPI.GetFirstUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StartupAPI.GetFirstUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirstUser`: StartupUserDto
	fmt.Fprintf(os.Stdout, "Response from `StartupAPI.GetFirstUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirstUserRequest struct via the builder pattern


### Return type

[**StartupUserDto**](StartupUserDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirstUser2

> StartupUserDto GetFirstUser2(ctx).Execute()

Gets the first user.

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
	resp, r, err := apiClient.StartupAPI.GetFirstUser2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StartupAPI.GetFirstUser2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirstUser2`: StartupUserDto
	fmt.Fprintf(os.Stdout, "Response from `StartupAPI.GetFirstUser2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirstUser2Request struct via the builder pattern


### Return type

[**StartupUserDto**](StartupUserDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStartupConfiguration

> StartupConfigurationDto GetStartupConfiguration(ctx).Execute()

Gets the initial startup wizard configuration.

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
	resp, r, err := apiClient.StartupAPI.GetStartupConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StartupAPI.GetStartupConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStartupConfiguration`: StartupConfigurationDto
	fmt.Fprintf(os.Stdout, "Response from `StartupAPI.GetStartupConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStartupConfigurationRequest struct via the builder pattern


### Return type

[**StartupConfigurationDto**](StartupConfigurationDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRemoteAccess

> SetRemoteAccess(ctx).StartupRemoteAccessDto(startupRemoteAccessDto).Execute()

Sets remote access and UPnP.

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
	startupRemoteAccessDto := *openapiclient.NewStartupRemoteAccessDto(false, false) // StartupRemoteAccessDto | The startup remote access dto.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StartupAPI.SetRemoteAccess(context.Background()).StartupRemoteAccessDto(startupRemoteAccessDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StartupAPI.SetRemoteAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRemoteAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startupRemoteAccessDto** | [**StartupRemoteAccessDto**](StartupRemoteAccessDto.md) | The startup remote access dto. | 

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


## UpdateInitialConfiguration

> UpdateInitialConfiguration(ctx).StartupConfigurationDto(startupConfigurationDto).Execute()

Sets the initial startup wizard configuration.

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
	startupConfigurationDto := *openapiclient.NewStartupConfigurationDto() // StartupConfigurationDto | The updated startup configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StartupAPI.UpdateInitialConfiguration(context.Background()).StartupConfigurationDto(startupConfigurationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StartupAPI.UpdateInitialConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInitialConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startupConfigurationDto** | [**StartupConfigurationDto**](StartupConfigurationDto.md) | The updated startup configuration. | 

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


## UpdateStartupUser

> UpdateStartupUser(ctx).StartupUserDto(startupUserDto).Execute()

Sets the user name and password.

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
	startupUserDto := *openapiclient.NewStartupUserDto() // StartupUserDto | The DTO containing username and password. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.StartupAPI.UpdateStartupUser(context.Background()).StartupUserDto(startupUserDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StartupAPI.UpdateStartupUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStartupUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startupUserDto** | [**StartupUserDto**](StartupUserDto.md) | The DTO containing username and password. | 

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

