# \LibraryStructureAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMediaPath**](LibraryStructureAPI.md#AddMediaPath) | **Post** /Library/VirtualFolders/Paths | Add a media path to a library.
[**AddVirtualFolder**](LibraryStructureAPI.md#AddVirtualFolder) | **Post** /Library/VirtualFolders | Adds a virtual folder.
[**GetVirtualFolders**](LibraryStructureAPI.md#GetVirtualFolders) | **Get** /Library/VirtualFolders | Gets all virtual folders.
[**RemoveMediaPath**](LibraryStructureAPI.md#RemoveMediaPath) | **Delete** /Library/VirtualFolders/Paths | Remove a media path.
[**RemoveVirtualFolder**](LibraryStructureAPI.md#RemoveVirtualFolder) | **Delete** /Library/VirtualFolders | Removes a virtual folder.
[**RenameVirtualFolder**](LibraryStructureAPI.md#RenameVirtualFolder) | **Post** /Library/VirtualFolders/Name | Renames a virtual folder.
[**UpdateLibraryOptions**](LibraryStructureAPI.md#UpdateLibraryOptions) | **Post** /Library/VirtualFolders/LibraryOptions | Update library options.
[**UpdateMediaPath**](LibraryStructureAPI.md#UpdateMediaPath) | **Post** /Library/VirtualFolders/Paths/Update | Updates a media path.



## AddMediaPath

> AddMediaPath(ctx).MediaPathDto(mediaPathDto).RefreshLibrary(refreshLibrary).Execute()

Add a media path to a library.

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
	mediaPathDto := *openapiclient.NewMediaPathDto("Name_example") // MediaPathDto | The media path dto.
	refreshLibrary := true // bool | Whether to refresh the library. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureAPI.AddMediaPath(context.Background()).MediaPathDto(mediaPathDto).RefreshLibrary(refreshLibrary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureAPI.AddMediaPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddMediaPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mediaPathDto** | [**MediaPathDto**](MediaPathDto.md) | The media path dto. | 
 **refreshLibrary** | **bool** | Whether to refresh the library. | [default to false]

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


## AddVirtualFolder

> AddVirtualFolder(ctx).Name(name).CollectionType(collectionType).Paths(paths).RefreshLibrary(refreshLibrary).AddVirtualFolderDto(addVirtualFolderDto).Execute()

Adds a virtual folder.

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
	name := "name_example" // string | The name of the virtual folder. (optional)
	collectionType := openapiclient.CollectionTypeOptions("Movies") // CollectionTypeOptions | The type of the collection. (optional)
	paths := []string{"Inner_example"} // []string | The paths of the virtual folder. (optional)
	refreshLibrary := true // bool | Whether to refresh the library. (optional) (default to false)
	addVirtualFolderDto := *openapiclient.NewAddVirtualFolderDto() // AddVirtualFolderDto | The library options. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureAPI.AddVirtualFolder(context.Background()).Name(name).CollectionType(collectionType).Paths(paths).RefreshLibrary(refreshLibrary).AddVirtualFolderDto(addVirtualFolderDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureAPI.AddVirtualFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddVirtualFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the virtual folder. | 
 **collectionType** | [**CollectionTypeOptions**](CollectionTypeOptions.md) | The type of the collection. | 
 **paths** | **[]string** | The paths of the virtual folder. | 
 **refreshLibrary** | **bool** | Whether to refresh the library. | [default to false]
 **addVirtualFolderDto** | [**AddVirtualFolderDto**](AddVirtualFolderDto.md) | The library options. | 

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


## GetVirtualFolders

> []VirtualFolderInfo GetVirtualFolders(ctx).Execute()

Gets all virtual folders.

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
	resp, r, err := apiClient.LibraryStructureAPI.GetVirtualFolders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureAPI.GetVirtualFolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVirtualFolders`: []VirtualFolderInfo
	fmt.Fprintf(os.Stdout, "Response from `LibraryStructureAPI.GetVirtualFolders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualFoldersRequest struct via the builder pattern


### Return type

[**[]VirtualFolderInfo**](VirtualFolderInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMediaPath

> RemoveMediaPath(ctx).Name(name).Path(path).RefreshLibrary(refreshLibrary).Execute()

Remove a media path.

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
	name := "name_example" // string | The name of the library. (optional)
	path := "path_example" // string | The path to remove. (optional)
	refreshLibrary := true // bool | Whether to refresh the library. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureAPI.RemoveMediaPath(context.Background()).Name(name).Path(path).RefreshLibrary(refreshLibrary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureAPI.RemoveMediaPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMediaPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the library. | 
 **path** | **string** | The path to remove. | 
 **refreshLibrary** | **bool** | Whether to refresh the library. | [default to false]

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


## RemoveVirtualFolder

> RemoveVirtualFolder(ctx).Name(name).RefreshLibrary(refreshLibrary).Execute()

Removes a virtual folder.

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
	name := "name_example" // string | The name of the folder. (optional)
	refreshLibrary := true // bool | Whether to refresh the library. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureAPI.RemoveVirtualFolder(context.Background()).Name(name).RefreshLibrary(refreshLibrary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureAPI.RemoveVirtualFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveVirtualFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the folder. | 
 **refreshLibrary** | **bool** | Whether to refresh the library. | [default to false]

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


## RenameVirtualFolder

> RenameVirtualFolder(ctx).Name(name).NewName(newName).RefreshLibrary(refreshLibrary).Execute()

Renames a virtual folder.

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
	name := "name_example" // string | The name of the virtual folder. (optional)
	newName := "newName_example" // string | The new name. (optional)
	refreshLibrary := true // bool | Whether to refresh the library. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureAPI.RenameVirtualFolder(context.Background()).Name(name).NewName(newName).RefreshLibrary(refreshLibrary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureAPI.RenameVirtualFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenameVirtualFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The name of the virtual folder. | 
 **newName** | **string** | The new name. | 
 **refreshLibrary** | **bool** | Whether to refresh the library. | [default to false]

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


## UpdateLibraryOptions

> UpdateLibraryOptions(ctx).UpdateLibraryOptionsDto(updateLibraryOptionsDto).Execute()

Update library options.

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
	updateLibraryOptionsDto := *openapiclient.NewUpdateLibraryOptionsDto() // UpdateLibraryOptionsDto | The library name and options. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureAPI.UpdateLibraryOptions(context.Background()).UpdateLibraryOptionsDto(updateLibraryOptionsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureAPI.UpdateLibraryOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLibraryOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLibraryOptionsDto** | [**UpdateLibraryOptionsDto**](UpdateLibraryOptionsDto.md) | The library name and options. | 

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


## UpdateMediaPath

> UpdateMediaPath(ctx).UpdateMediaPathRequestDto(updateMediaPathRequestDto).Execute()

Updates a media path.

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
	updateMediaPathRequestDto := *openapiclient.NewUpdateMediaPathRequestDto("Name_example", *openapiclient.NewMediaPathInfo()) // UpdateMediaPathRequestDto | The name of the library and path infos.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LibraryStructureAPI.UpdateMediaPath(context.Background()).UpdateMediaPathRequestDto(updateMediaPathRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryStructureAPI.UpdateMediaPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMediaPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMediaPathRequestDto** | [**UpdateMediaPathRequestDto**](UpdateMediaPathRequestDto.md) | The name of the library and path infos. | 

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

