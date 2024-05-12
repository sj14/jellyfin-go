# \DisplayPreferencesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDisplayPreferences**](DisplayPreferencesAPI.md#GetDisplayPreferences) | **Get** /DisplayPreferences/{displayPreferencesId} | Get Display Preferences.
[**UpdateDisplayPreferences**](DisplayPreferencesAPI.md#UpdateDisplayPreferences) | **Post** /DisplayPreferences/{displayPreferencesId} | Update Display Preferences.



## GetDisplayPreferences

> DisplayPreferencesDto GetDisplayPreferences(ctx, displayPreferencesId).Client(client).UserId(userId).Execute()

Get Display Preferences.

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
	displayPreferencesId := "displayPreferencesId_example" // string | Display preferences id.
	client := "client_example" // string | Client.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DisplayPreferencesAPI.GetDisplayPreferences(context.Background(), displayPreferencesId).Client(client).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisplayPreferencesAPI.GetDisplayPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDisplayPreferences`: DisplayPreferencesDto
	fmt.Fprintf(os.Stdout, "Response from `DisplayPreferencesAPI.GetDisplayPreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**displayPreferencesId** | **string** | Display preferences id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDisplayPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **client** | **string** | Client. | 
 **userId** | **string** | User id. | 

### Return type

[**DisplayPreferencesDto**](DisplayPreferencesDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDisplayPreferences

> UpdateDisplayPreferences(ctx, displayPreferencesId).Client(client).DisplayPreferencesDto(displayPreferencesDto).UserId(userId).Execute()

Update Display Preferences.

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
	displayPreferencesId := "displayPreferencesId_example" // string | Display preferences id.
	client := "client_example" // string | Client.
	displayPreferencesDto := *openapiclient.NewDisplayPreferencesDto() // DisplayPreferencesDto | New Display Preferences object.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User Id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DisplayPreferencesAPI.UpdateDisplayPreferences(context.Background(), displayPreferencesId).Client(client).DisplayPreferencesDto(displayPreferencesDto).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DisplayPreferencesAPI.UpdateDisplayPreferences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**displayPreferencesId** | **string** | Display preferences id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDisplayPreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **client** | **string** | Client. | 
 **displayPreferencesDto** | [**DisplayPreferencesDto**](DisplayPreferencesDto.md) | New Display Preferences object. | 
 **userId** | **string** | User Id. | 

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

