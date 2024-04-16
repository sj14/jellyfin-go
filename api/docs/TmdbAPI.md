# \TmdbAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TmdbClientConfiguration**](TmdbAPI.md#TmdbClientConfiguration) | **Get** /Tmdb/ClientConfiguration | Gets the TMDb image configuration options.



## TmdbClientConfiguration

> ConfigImageTypes TmdbClientConfiguration(ctx).Execute()

Gets the TMDb image configuration options.

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
	resp, r, err := apiClient.TmdbAPI.TmdbClientConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TmdbAPI.TmdbClientConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TmdbClientConfiguration`: ConfigImageTypes
	fmt.Fprintf(os.Stdout, "Response from `TmdbAPI.TmdbClientConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTmdbClientConfigurationRequest struct via the builder pattern


### Return type

[**ConfigImageTypes**](ConfigImageTypes.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

