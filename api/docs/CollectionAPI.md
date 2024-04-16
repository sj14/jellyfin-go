# \CollectionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToCollection**](CollectionAPI.md#AddToCollection) | **Post** /Collections/{collectionId}/Items | Adds items to a collection.
[**CreateCollection**](CollectionAPI.md#CreateCollection) | **Post** /Collections | Creates a new collection.
[**RemoveFromCollection**](CollectionAPI.md#RemoveFromCollection) | **Delete** /Collections/{collectionId}/Items | Removes items from a collection.



## AddToCollection

> AddToCollection(ctx, collectionId).Ids(ids).Execute()

Adds items to a collection.

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
	collectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The collection id.
	ids := []string{"Inner_example"} // []string | Item ids, comma delimited.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionAPI.AddToCollection(context.Background(), collectionId).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAPI.AddToCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The collection id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddToCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **[]string** | Item ids, comma delimited. | 

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


## CreateCollection

> CollectionCreationResult CreateCollection(ctx).Name(name).Ids(ids).ParentId(parentId).IsLocked(isLocked).Execute()

Creates a new collection.

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
	name := "name_example" // string | The name of the collection. (optional)
	ids := []string{"Inner_example"} // []string | Item Ids to add to the collection. (optional)
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Create the collection within a specific folder. (optional)
	isLocked := true // bool | Whether or not to lock the new collection. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionAPI.CreateCollection(context.Background()).Name(name).Ids(ids).ParentId(parentId).IsLocked(isLocked).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAPI.CreateCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCollection`: CollectionCreationResult
	fmt.Fprintf(os.Stdout, "Response from `CollectionAPI.CreateCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the collection. | 
 **ids** | **[]string** | Item Ids to add to the collection. | 
 **parentId** | **string** | Optional. Create the collection within a specific folder. | 
 **isLocked** | **bool** | Whether or not to lock the new collection. | [default to false]

### Return type

[**CollectionCreationResult**](CollectionCreationResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFromCollection

> RemoveFromCollection(ctx, collectionId).Ids(ids).Execute()

Removes items from a collection.

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
	collectionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The collection id.
	ids := []string{"Inner_example"} // []string | Item ids, comma delimited.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionAPI.RemoveFromCollection(context.Background(), collectionId).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionAPI.RemoveFromCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The collection id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFromCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **[]string** | Item ids, comma delimited. | 

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

