# \UserLibraryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUserItemRating**](UserLibraryAPI.md#DeleteUserItemRating) | **Delete** /Users/{userId}/Items/{itemId}/Rating | Deletes a user&#39;s saved personal rating for an item.
[**GetIntros**](UserLibraryAPI.md#GetIntros) | **Get** /Users/{userId}/Items/{itemId}/Intros | Gets intros to play before the main media item plays.
[**GetItem**](UserLibraryAPI.md#GetItem) | **Get** /Users/{userId}/Items/{itemId} | Gets an item from a user&#39;s library.
[**GetLatestMedia**](UserLibraryAPI.md#GetLatestMedia) | **Get** /Users/{userId}/Items/Latest | Gets latest media.
[**GetLocalTrailers**](UserLibraryAPI.md#GetLocalTrailers) | **Get** /Users/{userId}/Items/{itemId}/LocalTrailers | Gets local trailers for an item.
[**GetRootFolder**](UserLibraryAPI.md#GetRootFolder) | **Get** /Users/{userId}/Items/Root | Gets the root folder from a user&#39;s library.
[**GetSpecialFeatures**](UserLibraryAPI.md#GetSpecialFeatures) | **Get** /Users/{userId}/Items/{itemId}/SpecialFeatures | Gets special features for an item.
[**MarkFavoriteItem**](UserLibraryAPI.md#MarkFavoriteItem) | **Post** /Users/{userId}/FavoriteItems/{itemId} | Marks an item as a favorite.
[**UnmarkFavoriteItem**](UserLibraryAPI.md#UnmarkFavoriteItem) | **Delete** /Users/{userId}/FavoriteItems/{itemId} | Unmarks item as a favorite.
[**UpdateUserItemRating**](UserLibraryAPI.md#UpdateUserItemRating) | **Post** /Users/{userId}/Items/{itemId}/Rating | Updates a user&#39;s rating for an item.



## DeleteUserItemRating

> UserItemDataDto DeleteUserItemRating(ctx, userId, itemId).Execute()

Deletes a user's saved personal rating for an item.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id.
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryAPI.DeleteUserItemRating(context.Background(), userId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryAPI.DeleteUserItemRating``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserItemRating`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryAPI.DeleteUserItemRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id. | 
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserItemRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIntros

> BaseItemDtoQueryResult GetIntros(ctx, userId, itemId).Execute()

Gets intros to play before the main media item plays.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id.
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryAPI.GetIntros(context.Background(), userId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryAPI.GetIntros``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntros`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryAPI.GetIntros`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id. | 
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntrosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItem

> BaseItemDto GetItem(ctx, userId, itemId).Execute()

Gets an item from a user's library.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id.
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryAPI.GetItem(context.Background(), userId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryAPI.GetItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItem`: BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryAPI.GetItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id. | 
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestMedia

> []BaseItemDto GetLatestMedia(ctx, userId).ParentId(parentId).Fields(fields).IncludeItemTypes(includeItemTypes).IsPlayed(isPlayed).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Limit(limit).GroupItems(groupItems).Execute()

Gets latest media.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id.
	parentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Specify this to localize the search to a specific item or folder. Omit to use the root. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	includeItemTypes := []openapiclient.BaseItemKind{openapiclient.BaseItemKind("AggregateFolder")} // []BaseItemKind | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. (optional)
	isPlayed := true // bool | Filter by items that are played, or not. (optional)
	enableImages := true // bool | Optional. include image information in output. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. the max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)
	enableUserData := true // bool | Optional. include user data. (optional)
	limit := int32(56) // int32 | Return item limit. (optional) (default to 20)
	groupItems := true // bool | Whether or not to group items into a parent container. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryAPI.GetLatestMedia(context.Background(), userId).ParentId(parentId).Fields(fields).IncludeItemTypes(includeItemTypes).IsPlayed(isPlayed).EnableImages(enableImages).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).EnableUserData(enableUserData).Limit(limit).GroupItems(groupItems).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryAPI.GetLatestMedia``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestMedia`: []BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryAPI.GetLatestMedia`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestMediaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parentId** | **string** | Specify this to localize the search to a specific item or folder. Omit to use the root. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **includeItemTypes** | [**[]BaseItemKind**](BaseItemKind.md) | Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited. | 
 **isPlayed** | **bool** | Filter by items that are played, or not. | 
 **enableImages** | **bool** | Optional. include image information in output. | 
 **imageTypeLimit** | **int32** | Optional. the max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 
 **enableUserData** | **bool** | Optional. include user data. | 
 **limit** | **int32** | Return item limit. | [default to 20]
 **groupItems** | **bool** | Whether or not to group items into a parent container. | [default to true]

### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalTrailers

> []BaseItemDto GetLocalTrailers(ctx, userId, itemId).Execute()

Gets local trailers for an item.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id.
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryAPI.GetLocalTrailers(context.Background(), userId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryAPI.GetLocalTrailers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalTrailers`: []BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryAPI.GetLocalTrailers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id. | 
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalTrailersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRootFolder

> BaseItemDto GetRootFolder(ctx, userId).Execute()

Gets the root folder from a user's library.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryAPI.GetRootFolder(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryAPI.GetRootFolder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRootFolder`: BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryAPI.GetRootFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRootFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpecialFeatures

> []BaseItemDto GetSpecialFeatures(ctx, userId, itemId).Execute()

Gets special features for an item.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id.
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryAPI.GetSpecialFeatures(context.Background(), userId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryAPI.GetSpecialFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpecialFeatures`: []BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryAPI.GetSpecialFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id. | 
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSpecialFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkFavoriteItem

> UserItemDataDto MarkFavoriteItem(ctx, userId, itemId).Execute()

Marks an item as a favorite.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id.
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryAPI.MarkFavoriteItem(context.Background(), userId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryAPI.MarkFavoriteItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkFavoriteItem`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryAPI.MarkFavoriteItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id. | 
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkFavoriteItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnmarkFavoriteItem

> UserItemDataDto UnmarkFavoriteItem(ctx, userId, itemId).Execute()

Unmarks item as a favorite.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id.
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryAPI.UnmarkFavoriteItem(context.Background(), userId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryAPI.UnmarkFavoriteItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnmarkFavoriteItem`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryAPI.UnmarkFavoriteItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id. | 
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmarkFavoriteItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserItemRating

> UserItemDataDto UpdateUserItemRating(ctx, userId, itemId).Likes(likes).Execute()

Updates a user's rating for an item.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id.
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.
	likes := true // bool | Whether this M:Jellyfin.Api.Controllers.UserLibraryController.UpdateUserItemRating(System.Guid,System.Guid,System.Nullable{System.Boolean}) is likes. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserLibraryAPI.UpdateUserItemRating(context.Background(), userId, itemId).Likes(likes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserLibraryAPI.UpdateUserItemRating``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserItemRating`: UserItemDataDto
	fmt.Fprintf(os.Stdout, "Response from `UserLibraryAPI.UpdateUserItemRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User id. | 
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserItemRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **likes** | **bool** | Whether this M:Jellyfin.Api.Controllers.UserLibraryController.UpdateUserItemRating(System.Guid,System.Guid,System.Nullable{System.Boolean}) is likes. | 

### Return type

[**UserItemDataDto**](UserItemDataDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

