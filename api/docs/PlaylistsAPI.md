# \PlaylistsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddItemToPlaylist**](PlaylistsAPI.md#AddItemToPlaylist) | **Post** /Playlists/{playlistId}/Items | Adds items to a playlist.
[**CreatePlaylist**](PlaylistsAPI.md#CreatePlaylist) | **Post** /Playlists | Creates a new playlist.
[**GetPlaylistItems**](PlaylistsAPI.md#GetPlaylistItems) | **Get** /Playlists/{playlistId}/Items | Gets the original items of a playlist.
[**GetPlaylistUser**](PlaylistsAPI.md#GetPlaylistUser) | **Get** /Playlists/{playlistId}/Users/{userId} | Get a playlist user.
[**GetPlaylistUsers**](PlaylistsAPI.md#GetPlaylistUsers) | **Get** /Playlists/{playlistId}/Users | Get a playlist&#39;s users.
[**MoveItem**](PlaylistsAPI.md#MoveItem) | **Post** /Playlists/{playlistId}/Items/{itemId}/Move/{newIndex} | Moves a playlist item.
[**RemoveItemFromPlaylist**](PlaylistsAPI.md#RemoveItemFromPlaylist) | **Delete** /Playlists/{playlistId}/Items | Removes items from a playlist.
[**RemoveUserFromPlaylist**](PlaylistsAPI.md#RemoveUserFromPlaylist) | **Delete** /Playlists/{playlistId}/Users/{userId} | Remove a user from a playlist&#39;s users.
[**UpdatePlaylist**](PlaylistsAPI.md#UpdatePlaylist) | **Post** /Playlists/{playlistId} | Updates a playlist.
[**UpdatePlaylistUser**](PlaylistsAPI.md#UpdatePlaylistUser) | **Post** /Playlists/{playlistId}/Users/{userId} | Modify a user of a playlist&#39;s users.



## AddItemToPlaylist

> AddItemToPlaylist(ctx, playlistId).Ids(ids).UserId(userId).Execute()

Adds items to a playlist.

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
	playlistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The playlist id.
	ids := []string{"Inner_example"} // []string | Item id, comma delimited. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The userId. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistsAPI.AddItemToPlaylist(context.Background(), playlistId).Ids(ids).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.AddItemToPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** | The playlist id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddItemToPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **[]string** | Item id, comma delimited. | 
 **userId** | **string** | The userId. | 

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


## CreatePlaylist

> PlaylistCreationResult CreatePlaylist(ctx).Name(name).Ids(ids).UserId(userId).MediaType(mediaType).CreatePlaylistDto(createPlaylistDto).Execute()

Creates a new playlist.



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
	name := "name_example" // string | The playlist name. (optional)
	ids := []string{"Inner_example"} // []string | The item ids. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id. (optional)
	mediaType := "mediaType_example" // MediaType | The media type. (optional)
	createPlaylistDto := *openapiclient.NewCreatePlaylistDto() // CreatePlaylistDto | The create playlist payload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistsAPI.CreatePlaylist(context.Background()).Name(name).Ids(ids).UserId(userId).MediaType(mediaType).CreatePlaylistDto(createPlaylistDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.CreatePlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlaylist`: PlaylistCreationResult
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.CreatePlaylist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The playlist name. | 
 **ids** | **[]string** | The item ids. | 
 **userId** | **string** | The user id. | 
 **mediaType** | **MediaType** | The media type. | 
 **createPlaylistDto** | [**CreatePlaylistDto**](CreatePlaylistDto.md) | The create playlist payload. | 

### Return type

[**PlaylistCreationResult**](PlaylistCreationResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylistItems

> BaseItemDtoQueryResult GetPlaylistItems(ctx, playlistId).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Execute()

Gets the original items of a playlist.

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
	playlistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The playlist id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	enableImages := true // bool | Optional. Include image information in output. (optional)
	enableUserData := true // bool | Optional. Include user data. (optional)
	imageTypeLimit := int32(56) // int32 | Optional. The max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistsAPI.GetPlaylistItems(context.Background(), playlistId).UserId(userId).StartIndex(startIndex).Limit(limit).Fields(fields).EnableImages(enableImages).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.GetPlaylistItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaylistItems`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.GetPlaylistItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** | The playlist id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | User id. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **enableImages** | **bool** | Optional. Include image information in output. | 
 **enableUserData** | **bool** | Optional. Include user data. | 
 **imageTypeLimit** | **int32** | Optional. The max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 

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


## GetPlaylistUser

> PlaylistUserPermissions GetPlaylistUser(ctx, playlistId, userId).Execute()

Get a playlist user.

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
	playlistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The playlist id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistsAPI.GetPlaylistUser(context.Background(), playlistId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.GetPlaylistUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaylistUser`: PlaylistUserPermissions
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.GetPlaylistUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** | The playlist id. | 
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PlaylistUserPermissions**](PlaylistUserPermissions.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlaylistUsers

> []PlaylistUserPermissions GetPlaylistUsers(ctx, playlistId).Execute()

Get a playlist's users.

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
	playlistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The playlist id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaylistsAPI.GetPlaylistUsers(context.Background(), playlistId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.GetPlaylistUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlaylistUsers`: []PlaylistUserPermissions
	fmt.Fprintf(os.Stdout, "Response from `PlaylistsAPI.GetPlaylistUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** | The playlist id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlaylistUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PlaylistUserPermissions**](PlaylistUserPermissions.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveItem

> MoveItem(ctx, playlistId, itemId, newIndex).Execute()

Moves a playlist item.

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
	playlistId := "playlistId_example" // string | The playlist id.
	itemId := "itemId_example" // string | The item id.
	newIndex := int32(56) // int32 | The new index.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistsAPI.MoveItem(context.Background(), playlistId, itemId, newIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.MoveItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** | The playlist id. | 
**itemId** | **string** | The item id. | 
**newIndex** | **int32** | The new index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveItemRequest struct via the builder pattern


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


## RemoveItemFromPlaylist

> RemoveItemFromPlaylist(ctx, playlistId).EntryIds(entryIds).Execute()

Removes items from a playlist.

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
	playlistId := "playlistId_example" // string | The playlist id.
	entryIds := []string{"Inner_example"} // []string | The item ids, comma delimited. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistsAPI.RemoveItemFromPlaylist(context.Background(), playlistId).EntryIds(entryIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.RemoveItemFromPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** | The playlist id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveItemFromPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entryIds** | **[]string** | The item ids, comma delimited. | 

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


## RemoveUserFromPlaylist

> RemoveUserFromPlaylist(ctx, playlistId, userId).Execute()

Remove a user from a playlist's users.

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
	playlistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The playlist id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistsAPI.RemoveUserFromPlaylist(context.Background(), playlistId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.RemoveUserFromPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** | The playlist id. | 
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromPlaylistRequest struct via the builder pattern


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


## UpdatePlaylist

> UpdatePlaylist(ctx, playlistId).UpdatePlaylistDto(updatePlaylistDto).Execute()

Updates a playlist.

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
	playlistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The playlist id.
	updatePlaylistDto := *openapiclient.NewUpdatePlaylistDto() // UpdatePlaylistDto | The Jellyfin.Api.Models.PlaylistDtos.UpdatePlaylistDto id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistsAPI.UpdatePlaylist(context.Background(), playlistId).UpdatePlaylistDto(updatePlaylistDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.UpdatePlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** | The playlist id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePlaylistDto** | [**UpdatePlaylistDto**](UpdatePlaylistDto.md) | The Jellyfin.Api.Models.PlaylistDtos.UpdatePlaylistDto id. | 

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


## UpdatePlaylistUser

> UpdatePlaylistUser(ctx, playlistId, userId).UpdatePlaylistUserDto(updatePlaylistUserDto).Execute()

Modify a user of a playlist's users.

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
	playlistId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The playlist id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.
	updatePlaylistUserDto := *openapiclient.NewUpdatePlaylistUserDto() // UpdatePlaylistUserDto | The Jellyfin.Api.Models.PlaylistDtos.UpdatePlaylistUserDto.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaylistsAPI.UpdatePlaylistUser(context.Background(), playlistId, userId).UpdatePlaylistUserDto(updatePlaylistUserDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaylistsAPI.UpdatePlaylistUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playlistId** | **string** | The playlist id. | 
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlaylistUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePlaylistUserDto** | [**UpdatePlaylistUserDto**](UpdatePlaylistUserDto.md) | The Jellyfin.Api.Models.PlaylistDtos.UpdatePlaylistUserDto. | 

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

