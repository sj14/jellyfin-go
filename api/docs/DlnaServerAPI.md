# \DlnaServerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConnectionManager**](DlnaServerAPI.md#GetConnectionManager) | **Get** /Dlna/{serverId}/ConnectionManager | Gets Dlna media receiver registrar xml.
[**GetConnectionManager2**](DlnaServerAPI.md#GetConnectionManager2) | **Get** /Dlna/{serverId}/ConnectionManager/ConnectionManager | Gets Dlna media receiver registrar xml.
[**GetConnectionManager3**](DlnaServerAPI.md#GetConnectionManager3) | **Get** /Dlna/{serverId}/ConnectionManager/ConnectionManager.xml | Gets Dlna media receiver registrar xml.
[**GetContentDirectory**](DlnaServerAPI.md#GetContentDirectory) | **Get** /Dlna/{serverId}/ContentDirectory | Gets Dlna content directory xml.
[**GetContentDirectory2**](DlnaServerAPI.md#GetContentDirectory2) | **Get** /Dlna/{serverId}/ContentDirectory/ContentDirectory | Gets Dlna content directory xml.
[**GetContentDirectory3**](DlnaServerAPI.md#GetContentDirectory3) | **Get** /Dlna/{serverId}/ContentDirectory/ContentDirectory.xml | Gets Dlna content directory xml.
[**GetDescriptionXml**](DlnaServerAPI.md#GetDescriptionXml) | **Get** /Dlna/{serverId}/description | Get Description Xml.
[**GetDescriptionXml2**](DlnaServerAPI.md#GetDescriptionXml2) | **Get** /Dlna/{serverId}/description.xml | Get Description Xml.
[**GetIcon**](DlnaServerAPI.md#GetIcon) | **Get** /Dlna/icons/{fileName} | Gets a server icon.
[**GetIconId**](DlnaServerAPI.md#GetIconId) | **Get** /Dlna/{serverId}/icons/{fileName} | Gets a server icon.
[**GetMediaReceiverRegistrar**](DlnaServerAPI.md#GetMediaReceiverRegistrar) | **Get** /Dlna/{serverId}/MediaReceiverRegistrar | Gets Dlna media receiver registrar xml.
[**GetMediaReceiverRegistrar2**](DlnaServerAPI.md#GetMediaReceiverRegistrar2) | **Get** /Dlna/{serverId}/MediaReceiverRegistrar/MediaReceiverRegistrar | Gets Dlna media receiver registrar xml.
[**GetMediaReceiverRegistrar3**](DlnaServerAPI.md#GetMediaReceiverRegistrar3) | **Get** /Dlna/{serverId}/MediaReceiverRegistrar/MediaReceiverRegistrar.xml | Gets Dlna media receiver registrar xml.
[**ProcessConnectionManagerControlRequest**](DlnaServerAPI.md#ProcessConnectionManagerControlRequest) | **Post** /Dlna/{serverId}/ConnectionManager/Control | Process a connection manager control request.
[**ProcessContentDirectoryControlRequest**](DlnaServerAPI.md#ProcessContentDirectoryControlRequest) | **Post** /Dlna/{serverId}/ContentDirectory/Control | Process a content directory control request.
[**ProcessMediaReceiverRegistrarControlRequest**](DlnaServerAPI.md#ProcessMediaReceiverRegistrarControlRequest) | **Post** /Dlna/{serverId}/MediaReceiverRegistrar/Control | Process a media receiver registrar control request.



## GetConnectionManager

> *os.File GetConnectionManager(ctx, serverId).Execute()

Gets Dlna media receiver registrar xml.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.GetConnectionManager(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.GetConnectionManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionManager`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.GetConnectionManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionManager2

> *os.File GetConnectionManager2(ctx, serverId).Execute()

Gets Dlna media receiver registrar xml.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.GetConnectionManager2(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.GetConnectionManager2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionManager2`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.GetConnectionManager2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionManager2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionManager3

> *os.File GetConnectionManager3(ctx, serverId).Execute()

Gets Dlna media receiver registrar xml.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.GetConnectionManager3(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.GetConnectionManager3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectionManager3`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.GetConnectionManager3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionManager3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentDirectory

> *os.File GetContentDirectory(ctx, serverId).Execute()

Gets Dlna content directory xml.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.GetContentDirectory(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.GetContentDirectory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentDirectory`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.GetContentDirectory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentDirectoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentDirectory2

> *os.File GetContentDirectory2(ctx, serverId).Execute()

Gets Dlna content directory xml.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.GetContentDirectory2(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.GetContentDirectory2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentDirectory2`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.GetContentDirectory2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentDirectory2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentDirectory3

> *os.File GetContentDirectory3(ctx, serverId).Execute()

Gets Dlna content directory xml.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.GetContentDirectory3(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.GetContentDirectory3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContentDirectory3`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.GetContentDirectory3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentDirectory3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDescriptionXml

> *os.File GetDescriptionXml(ctx, serverId).Execute()

Get Description Xml.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.GetDescriptionXml(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.GetDescriptionXml``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDescriptionXml`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.GetDescriptionXml`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDescriptionXmlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDescriptionXml2

> *os.File GetDescriptionXml2(ctx, serverId).Execute()

Get Description Xml.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.GetDescriptionXml2(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.GetDescriptionXml2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDescriptionXml2`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.GetDescriptionXml2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDescriptionXml2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIcon

> *os.File GetIcon(ctx, fileName).Execute()

Gets a server icon.

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
	fileName := "fileName_example" // string | The icon filename.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.GetIcon(context.Background(), fileName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.GetIcon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIcon`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.GetIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileName** | **string** | The icon filename. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIconRequest struct via the builder pattern


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


## GetIconId

> *os.File GetIconId(ctx, serverId, fileName).Execute()

Gets a server icon.

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
	serverId := "serverId_example" // string | Server UUID.
	fileName := "fileName_example" // string | The icon filename.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.GetIconId(context.Background(), serverId, fileName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.GetIconId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIconId`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.GetIconId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 
**fileName** | **string** | The icon filename. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIconIdRequest struct via the builder pattern


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


## GetMediaReceiverRegistrar

> *os.File GetMediaReceiverRegistrar(ctx, serverId).Execute()

Gets Dlna media receiver registrar xml.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.GetMediaReceiverRegistrar(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.GetMediaReceiverRegistrar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaReceiverRegistrar`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.GetMediaReceiverRegistrar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaReceiverRegistrarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaReceiverRegistrar2

> *os.File GetMediaReceiverRegistrar2(ctx, serverId).Execute()

Gets Dlna media receiver registrar xml.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.GetMediaReceiverRegistrar2(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.GetMediaReceiverRegistrar2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaReceiverRegistrar2`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.GetMediaReceiverRegistrar2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaReceiverRegistrar2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaReceiverRegistrar3

> *os.File GetMediaReceiverRegistrar3(ctx, serverId).Execute()

Gets Dlna media receiver registrar xml.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.GetMediaReceiverRegistrar3(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.GetMediaReceiverRegistrar3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaReceiverRegistrar3`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.GetMediaReceiverRegistrar3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaReceiverRegistrar3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessConnectionManagerControlRequest

> *os.File ProcessConnectionManagerControlRequest(ctx, serverId).Execute()

Process a connection manager control request.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.ProcessConnectionManagerControlRequest(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.ProcessConnectionManagerControlRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessConnectionManagerControlRequest`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.ProcessConnectionManagerControlRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessConnectionManagerControlRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessContentDirectoryControlRequest

> *os.File ProcessContentDirectoryControlRequest(ctx, serverId).Execute()

Process a content directory control request.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.ProcessContentDirectoryControlRequest(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.ProcessContentDirectoryControlRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessContentDirectoryControlRequest`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.ProcessContentDirectoryControlRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessContentDirectoryControlRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessMediaReceiverRegistrarControlRequest

> *os.File ProcessMediaReceiverRegistrarControlRequest(ctx, serverId).Execute()

Process a media receiver registrar control request.

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
	serverId := "serverId_example" // string | Server UUID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DlnaServerAPI.ProcessMediaReceiverRegistrarControlRequest(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DlnaServerAPI.ProcessMediaReceiverRegistrarControlRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessMediaReceiverRegistrarControlRequest`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DlnaServerAPI.ProcessMediaReceiverRegistrarControlRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | Server UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessMediaReceiverRegistrarControlRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

