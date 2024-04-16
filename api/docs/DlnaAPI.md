# \DlnaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfile**](DlnaAPI.md#CreateProfile) | **Post** /Dlna/Profiles | Creates a profile.
[**DeleteProfile**](DlnaAPI.md#DeleteProfile) | **Delete** /Dlna/Profiles/{profileId} | Deletes a profile.
[**GetDefaultProfile**](DlnaAPI.md#GetDefaultProfile) | **Get** /Dlna/Profiles/Default | Gets the default profile.
[**GetProfile**](DlnaAPI.md#GetProfile) | **Get** /Dlna/Profiles/{profileId} | Gets a single profile.
[**GetProfileInfos**](DlnaAPI.md#GetProfileInfos) | **Get** /Dlna/ProfileInfos | Get profile infos.
[**UpdateProfile**](DlnaAPI.md#UpdateProfile) | **Post** /Dlna/Profiles/{profileId} | Updates a profile.



## CreateProfile

> CreateProfile(ctx).DeviceProfile(deviceProfile).Execute()

Creates a profile.

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
	deviceProfile := *openapiclient.NewDeviceProfile() // DeviceProfile | Device profile. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaAPI.CreateProfile(context.Background()).DeviceProfile(deviceProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaAPI.CreateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceProfile** | [**DeviceProfile**](DeviceProfile.md) | Device profile. | 

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


## DeleteProfile

> DeleteProfile(ctx, profileId).Execute()

Deletes a profile.

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
	profileId := "profileId_example" // string | Profile id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaAPI.DeleteProfile(context.Background(), profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaAPI.DeleteProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Profile id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProfileRequest struct via the builder pattern


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


## GetDefaultProfile

> DeviceProfile GetDefaultProfile(ctx).Execute()

Gets the default profile.

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
	resp, r, err := apiClient.DlnaAPI.GetDefaultProfile(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaAPI.GetDefaultProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultProfile`: DeviceProfile
	fmt.Fprintf(os.Stdout, "Response from `DlnaAPI.GetDefaultProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultProfileRequest struct via the builder pattern


### Return type

[**DeviceProfile**](DeviceProfile.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfile

> DeviceProfile GetProfile(ctx, profileId).Execute()

Gets a single profile.

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
	profileId := "profileId_example" // string | Profile Id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaAPI.GetProfile(context.Background(), profileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaAPI.GetProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfile`: DeviceProfile
	fmt.Fprintf(os.Stdout, "Response from `DlnaAPI.GetProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Profile Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceProfile**](DeviceProfile.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProfileInfos

> []DeviceProfileInfo GetProfileInfos(ctx).Execute()

Get profile infos.

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
	resp, r, err := apiClient.DlnaAPI.GetProfileInfos(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaAPI.GetProfileInfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProfileInfos`: []DeviceProfileInfo
	fmt.Fprintf(os.Stdout, "Response from `DlnaAPI.GetProfileInfos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileInfosRequest struct via the builder pattern


### Return type

[**[]DeviceProfileInfo**](DeviceProfileInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> UpdateProfile(ctx, profileId).DeviceProfile(deviceProfile).Execute()

Updates a profile.

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
	profileId := "profileId_example" // string | Profile id.
	deviceProfile := *openapiclient.NewDeviceProfile() // DeviceProfile | Device profile. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DlnaAPI.UpdateProfile(context.Background(), profileId).DeviceProfile(deviceProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaAPI.UpdateProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profileId** | **string** | Profile id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceProfile** | [**DeviceProfile**](DeviceProfile.md) | Device profile. | 

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

