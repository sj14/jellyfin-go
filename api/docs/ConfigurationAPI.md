# \ConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration**](ConfigurationAPI.md#GetConfiguration) | **Get** /System/Configuration | Gets application configuration.
[**GetDefaultMetadataOptions**](ConfigurationAPI.md#GetDefaultMetadataOptions) | **Get** /System/Configuration/MetadataOptions/Default | Gets a default MetadataOptions object.
[**GetNamedConfiguration**](ConfigurationAPI.md#GetNamedConfiguration) | **Get** /System/Configuration/{key} | Gets a named configuration.
[**UpdateConfiguration**](ConfigurationAPI.md#UpdateConfiguration) | **Post** /System/Configuration | Updates application configuration.
[**UpdateMediaEncoderPath**](ConfigurationAPI.md#UpdateMediaEncoderPath) | **Post** /System/MediaEncoder/Path | Updates the path to the media encoder.
[**UpdateNamedConfiguration**](ConfigurationAPI.md#UpdateNamedConfiguration) | **Post** /System/Configuration/{key} | Updates named configuration.



## GetConfiguration

> ServerConfiguration GetConfiguration(ctx).Execute()

Gets application configuration.

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
	resp, r, err := apiClient.ConfigurationAPI.GetConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration`: ServerConfiguration
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


### Return type

[**ServerConfiguration**](ServerConfiguration.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultMetadataOptions

> MetadataOptions GetDefaultMetadataOptions(ctx).Execute()

Gets a default MetadataOptions object.

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
	resp, r, err := apiClient.ConfigurationAPI.GetDefaultMetadataOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetDefaultMetadataOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultMetadataOptions`: MetadataOptions
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetDefaultMetadataOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultMetadataOptionsRequest struct via the builder pattern


### Return type

[**MetadataOptions**](MetadataOptions.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamedConfiguration

> *os.File GetNamedConfiguration(ctx, key).Execute()

Gets a named configuration.

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
	key := "key_example" // string | Configuration key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetNamedConfiguration(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetNamedConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamedConfiguration`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetNamedConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Configuration key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamedConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> UpdateConfiguration(ctx).ServerConfiguration(serverConfiguration).Execute()

Updates application configuration.

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
	serverConfiguration := *openapiclient.NewServerConfiguration() // ServerConfiguration | Configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationAPI.UpdateConfiguration(context.Background()).ServerConfiguration(serverConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.UpdateConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverConfiguration** | [**ServerConfiguration**](ServerConfiguration.md) | Configuration. | 

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


## UpdateMediaEncoderPath

> UpdateMediaEncoderPath(ctx).MediaEncoderPathDto(mediaEncoderPathDto).Execute()

Updates the path to the media encoder.

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
	mediaEncoderPathDto := *openapiclient.NewMediaEncoderPathDto() // MediaEncoderPathDto | Media encoder path form body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationAPI.UpdateMediaEncoderPath(context.Background()).MediaEncoderPathDto(mediaEncoderPathDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.UpdateMediaEncoderPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMediaEncoderPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaEncoderPathDto** | [**MediaEncoderPathDto**](MediaEncoderPathDto.md) | Media encoder path form body. | 

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


## UpdateNamedConfiguration

> UpdateNamedConfiguration(ctx, key).Body(body).Execute()

Updates named configuration.

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
	key := "key_example" // string | Configuration key.
	body := interface{}(987) // interface{} | Configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationAPI.UpdateNamedConfiguration(context.Background(), key).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.UpdateNamedConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Configuration key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNamedConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** | Configuration. | 

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

