# \QuickConnectAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authorize**](QuickConnectAPI.md#Authorize) | **Post** /QuickConnect/Authorize | Authorizes a pending quick connect request.
[**Connect**](QuickConnectAPI.md#Connect) | **Get** /QuickConnect/Connect | Attempts to retrieve authentication information.
[**GetEnabled**](QuickConnectAPI.md#GetEnabled) | **Get** /QuickConnect/Enabled | Gets the current quick connect state.
[**Initiate**](QuickConnectAPI.md#Initiate) | **Get** /QuickConnect/Initiate | Initiate a new quick connect request.



## Authorize

> bool Authorize(ctx).Code(code).Execute()

Authorizes a pending quick connect request.

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
	code := "code_example" // string | Quick connect code to authorize.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuickConnectAPI.Authorize(context.Background()).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuickConnectAPI.Authorize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Authorize`: bool
	fmt.Fprintf(os.Stdout, "Response from `QuickConnectAPI.Authorize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | Quick connect code to authorize. | 

### Return type

**bool**

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Connect

> QuickConnectResult Connect(ctx).Secret(secret).Execute()

Attempts to retrieve authentication information.

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
	secret := "secret_example" // string | Secret previously returned from the Initiate endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuickConnectAPI.Connect(context.Background()).Secret(secret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuickConnectAPI.Connect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Connect`: QuickConnectResult
	fmt.Fprintf(os.Stdout, "Response from `QuickConnectAPI.Connect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secret** | **string** | Secret previously returned from the Initiate endpoint. | 

### Return type

[**QuickConnectResult**](QuickConnectResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnabled

> bool GetEnabled(ctx).Execute()

Gets the current quick connect state.

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
	resp, r, err := apiClient.QuickConnectAPI.GetEnabled(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuickConnectAPI.GetEnabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnabled`: bool
	fmt.Fprintf(os.Stdout, "Response from `QuickConnectAPI.GetEnabled`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnabledRequest struct via the builder pattern


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Initiate

> QuickConnectResult Initiate(ctx).Execute()

Initiate a new quick connect request.

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
	resp, r, err := apiClient.QuickConnectAPI.Initiate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuickConnectAPI.Initiate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Initiate`: QuickConnectResult
	fmt.Fprintf(os.Stdout, "Response from `QuickConnectAPI.Initiate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateRequest struct via the builder pattern


### Return type

[**QuickConnectResult**](QuickConnectResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

