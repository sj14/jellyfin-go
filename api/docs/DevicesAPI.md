# \DevicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDevice**](DevicesAPI.md#DeleteDevice) | **Delete** /Devices | Deletes a device.
[**GetDeviceInfo**](DevicesAPI.md#GetDeviceInfo) | **Get** /Devices/Info | Get info for a device.
[**GetDeviceOptions**](DevicesAPI.md#GetDeviceOptions) | **Get** /Devices/Options | Get options for a device.
[**GetDevices**](DevicesAPI.md#GetDevices) | **Get** /Devices | Get Devices.
[**UpdateDeviceOptions**](DevicesAPI.md#UpdateDeviceOptions) | **Post** /Devices/Options | Update device options.



## DeleteDevice

> DeleteDevice(ctx).Id(id).Execute()

Deletes a device.

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
	id := "id_example" // string | Device Id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.DeleteDevice(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DeleteDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Device Id. | 

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


## GetDeviceInfo

> DeviceInfo GetDeviceInfo(ctx).Id(id).Execute()

Get info for a device.

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
	id := "id_example" // string | Device Id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDeviceInfo(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDeviceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceInfo`: DeviceInfo
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDeviceInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Device Id. | 

### Return type

[**DeviceInfo**](DeviceInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceOptions

> DeviceOptions GetDeviceOptions(ctx).Id(id).Execute()

Get options for a device.

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
	id := "id_example" // string | Device Id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDeviceOptions(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDeviceOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceOptions`: DeviceOptions
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDeviceOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Device Id. | 

### Return type

[**DeviceOptions**](DeviceOptions.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevices

> DeviceInfoQueryResult GetDevices(ctx).SupportsSync(supportsSync).UserId(userId).Execute()

Get Devices.

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
	supportsSync := true // bool | Gets or sets a value indicating whether [supports synchronize]. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Gets or sets the user identifier. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevicesAPI.GetDevices(context.Background()).SupportsSync(supportsSync).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevices`: DeviceInfoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportsSync** | **bool** | Gets or sets a value indicating whether [supports synchronize]. | 
 **userId** | **string** | Gets or sets the user identifier. | 

### Return type

[**DeviceInfoQueryResult**](DeviceInfoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceOptions

> UpdateDeviceOptions(ctx).Id(id).DeviceOptionsDto(deviceOptionsDto).Execute()

Update device options.

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
	id := "id_example" // string | Device Id.
	deviceOptionsDto := *openapiclient.NewDeviceOptionsDto() // DeviceOptionsDto | Device Options.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevicesAPI.UpdateDeviceOptions(context.Background()).Id(id).DeviceOptionsDto(deviceOptionsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.UpdateDeviceOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | Device Id. | 
 **deviceOptionsDto** | [**DeviceOptionsDto**](DeviceOptionsDto.md) | Device Options. | 

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

