# \QuickConnectAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthorizeQuickConnect**](QuickConnectAPI.md#AuthorizeQuickConnect) | **Post** /QuickConnect/Authorize | Authorizes a pending quick connect request.
[**GetQuickConnectEnabled**](QuickConnectAPI.md#GetQuickConnectEnabled) | **Get** /QuickConnect/Enabled | Gets the current quick connect state.
[**GetQuickConnectState**](QuickConnectAPI.md#GetQuickConnectState) | **Get** /QuickConnect/Connect | Attempts to retrieve authentication information.
[**InitiateQuickConnect**](QuickConnectAPI.md#InitiateQuickConnect) | **Post** /QuickConnect/Initiate | Initiate a new quick connect request.



## AuthorizeQuickConnect

> bool AuthorizeQuickConnect(ctx).Code(code).UserId(userId).Execute()

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user the authorize. Access to the requested user is required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.QuickConnectAPI.AuthorizeQuickConnect(context.Background()).Code(code).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuickConnectAPI.AuthorizeQuickConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthorizeQuickConnect`: bool
	fmt.Fprintf(os.Stdout, "Response from `QuickConnectAPI.AuthorizeQuickConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthorizeQuickConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | Quick connect code to authorize. | 
 **userId** | **string** | The user the authorize. Access to the requested user is required. | 

### Return type

**bool**

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuickConnectEnabled

> bool GetQuickConnectEnabled(ctx).Execute()

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
	resp, r, err := apiClient.QuickConnectAPI.GetQuickConnectEnabled(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuickConnectAPI.GetQuickConnectEnabled``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuickConnectEnabled`: bool
	fmt.Fprintf(os.Stdout, "Response from `QuickConnectAPI.GetQuickConnectEnabled`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuickConnectEnabledRequest struct via the builder pattern


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuickConnectState

> QuickConnectResult GetQuickConnectState(ctx).Secret(secret).Execute()

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
	resp, r, err := apiClient.QuickConnectAPI.GetQuickConnectState(context.Background()).Secret(secret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuickConnectAPI.GetQuickConnectState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuickConnectState`: QuickConnectResult
	fmt.Fprintf(os.Stdout, "Response from `QuickConnectAPI.GetQuickConnectState`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuickConnectStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **secret** | **string** | Secret previously returned from the Initiate endpoint. | 

### Return type

[**QuickConnectResult**](QuickConnectResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitiateQuickConnect

> QuickConnectResult InitiateQuickConnect(ctx).Execute()

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
	resp, r, err := apiClient.QuickConnectAPI.InitiateQuickConnect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuickConnectAPI.InitiateQuickConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitiateQuickConnect`: QuickConnectResult
	fmt.Fprintf(os.Stdout, "Response from `QuickConnectAPI.InitiateQuickConnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateQuickConnectRequest struct via the builder pattern


### Return type

[**QuickConnectResult**](QuickConnectResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

