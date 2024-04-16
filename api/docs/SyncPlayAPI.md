# \SyncPlayAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyncPlayBuffering**](SyncPlayAPI.md#SyncPlayBuffering) | **Post** /SyncPlay/Buffering | Notify SyncPlay group that member is buffering.
[**SyncPlayCreateGroup**](SyncPlayAPI.md#SyncPlayCreateGroup) | **Post** /SyncPlay/New | Create a new SyncPlay group.
[**SyncPlayGetGroups**](SyncPlayAPI.md#SyncPlayGetGroups) | **Get** /SyncPlay/List | Gets all SyncPlay groups.
[**SyncPlayJoinGroup**](SyncPlayAPI.md#SyncPlayJoinGroup) | **Post** /SyncPlay/Join | Join an existing SyncPlay group.
[**SyncPlayLeaveGroup**](SyncPlayAPI.md#SyncPlayLeaveGroup) | **Post** /SyncPlay/Leave | Leave the joined SyncPlay group.
[**SyncPlayMovePlaylistItem**](SyncPlayAPI.md#SyncPlayMovePlaylistItem) | **Post** /SyncPlay/MovePlaylistItem | Request to move an item in the playlist in SyncPlay group.
[**SyncPlayNextItem**](SyncPlayAPI.md#SyncPlayNextItem) | **Post** /SyncPlay/NextItem | Request next item in SyncPlay group.
[**SyncPlayPause**](SyncPlayAPI.md#SyncPlayPause) | **Post** /SyncPlay/Pause | Request pause in SyncPlay group.
[**SyncPlayPing**](SyncPlayAPI.md#SyncPlayPing) | **Post** /SyncPlay/Ping | Update session ping.
[**SyncPlayPreviousItem**](SyncPlayAPI.md#SyncPlayPreviousItem) | **Post** /SyncPlay/PreviousItem | Request previous item in SyncPlay group.
[**SyncPlayQueue**](SyncPlayAPI.md#SyncPlayQueue) | **Post** /SyncPlay/Queue | Request to queue items to the playlist of a SyncPlay group.
[**SyncPlayReady**](SyncPlayAPI.md#SyncPlayReady) | **Post** /SyncPlay/Ready | Notify SyncPlay group that member is ready for playback.
[**SyncPlayRemoveFromPlaylist**](SyncPlayAPI.md#SyncPlayRemoveFromPlaylist) | **Post** /SyncPlay/RemoveFromPlaylist | Request to remove items from the playlist in SyncPlay group.
[**SyncPlaySeek**](SyncPlayAPI.md#SyncPlaySeek) | **Post** /SyncPlay/Seek | Request seek in SyncPlay group.
[**SyncPlaySetIgnoreWait**](SyncPlayAPI.md#SyncPlaySetIgnoreWait) | **Post** /SyncPlay/SetIgnoreWait | Request SyncPlay group to ignore member during group-wait.
[**SyncPlaySetNewQueue**](SyncPlayAPI.md#SyncPlaySetNewQueue) | **Post** /SyncPlay/SetNewQueue | Request to set new playlist in SyncPlay group.
[**SyncPlaySetPlaylistItem**](SyncPlayAPI.md#SyncPlaySetPlaylistItem) | **Post** /SyncPlay/SetPlaylistItem | Request to change playlist item in SyncPlay group.
[**SyncPlaySetRepeatMode**](SyncPlayAPI.md#SyncPlaySetRepeatMode) | **Post** /SyncPlay/SetRepeatMode | Request to set repeat mode in SyncPlay group.
[**SyncPlaySetShuffleMode**](SyncPlayAPI.md#SyncPlaySetShuffleMode) | **Post** /SyncPlay/SetShuffleMode | Request to set shuffle mode in SyncPlay group.
[**SyncPlayStop**](SyncPlayAPI.md#SyncPlayStop) | **Post** /SyncPlay/Stop | Request stop in SyncPlay group.
[**SyncPlayUnpause**](SyncPlayAPI.md#SyncPlayUnpause) | **Post** /SyncPlay/Unpause | Request unpause in SyncPlay group.



## SyncPlayBuffering

> SyncPlayBuffering(ctx).BufferRequestDto(bufferRequestDto).Execute()

Notify SyncPlay group that member is buffering.

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
	bufferRequestDto := *openapiclient.NewBufferRequestDto() // BufferRequestDto | The player status.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayBuffering(context.Background()).BufferRequestDto(bufferRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayBuffering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayBufferingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bufferRequestDto** | [**BufferRequestDto**](BufferRequestDto.md) | The player status. | 

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


## SyncPlayCreateGroup

> SyncPlayCreateGroup(ctx).NewGroupRequestDto(newGroupRequestDto).Execute()

Create a new SyncPlay group.

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
	newGroupRequestDto := *openapiclient.NewNewGroupRequestDto() // NewGroupRequestDto | The settings of the new group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayCreateGroup(context.Background()).NewGroupRequestDto(newGroupRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayCreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newGroupRequestDto** | [**NewGroupRequestDto**](NewGroupRequestDto.md) | The settings of the new group. | 

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


## SyncPlayGetGroups

> []GroupInfoDto SyncPlayGetGroups(ctx).Execute()

Gets all SyncPlay groups.

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
	resp, r, err := apiClient.SyncPlayAPI.SyncPlayGetGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayGetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncPlayGetGroups`: []GroupInfoDto
	fmt.Fprintf(os.Stdout, "Response from `SyncPlayAPI.SyncPlayGetGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayGetGroupsRequest struct via the builder pattern


### Return type

[**[]GroupInfoDto**](GroupInfoDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncPlayJoinGroup

> SyncPlayJoinGroup(ctx).JoinGroupRequestDto(joinGroupRequestDto).Execute()

Join an existing SyncPlay group.

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
	joinGroupRequestDto := *openapiclient.NewJoinGroupRequestDto() // JoinGroupRequestDto | The group to join.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayJoinGroup(context.Background()).JoinGroupRequestDto(joinGroupRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayJoinGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayJoinGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **joinGroupRequestDto** | [**JoinGroupRequestDto**](JoinGroupRequestDto.md) | The group to join. | 

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


## SyncPlayLeaveGroup

> SyncPlayLeaveGroup(ctx).Execute()

Leave the joined SyncPlay group.

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
	r, err := apiClient.SyncPlayAPI.SyncPlayLeaveGroup(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayLeaveGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayLeaveGroupRequest struct via the builder pattern


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


## SyncPlayMovePlaylistItem

> SyncPlayMovePlaylistItem(ctx).MovePlaylistItemRequestDto(movePlaylistItemRequestDto).Execute()

Request to move an item in the playlist in SyncPlay group.

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
	movePlaylistItemRequestDto := *openapiclient.NewMovePlaylistItemRequestDto() // MovePlaylistItemRequestDto | The new position for the item.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayMovePlaylistItem(context.Background()).MovePlaylistItemRequestDto(movePlaylistItemRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayMovePlaylistItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayMovePlaylistItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movePlaylistItemRequestDto** | [**MovePlaylistItemRequestDto**](MovePlaylistItemRequestDto.md) | The new position for the item. | 

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


## SyncPlayNextItem

> SyncPlayNextItem(ctx).NextItemRequestDto(nextItemRequestDto).Execute()

Request next item in SyncPlay group.

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
	nextItemRequestDto := *openapiclient.NewNextItemRequestDto() // NextItemRequestDto | The current item information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayNextItem(context.Background()).NextItemRequestDto(nextItemRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayNextItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayNextItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextItemRequestDto** | [**NextItemRequestDto**](NextItemRequestDto.md) | The current item information. | 

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


## SyncPlayPause

> SyncPlayPause(ctx).Execute()

Request pause in SyncPlay group.

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
	r, err := apiClient.SyncPlayAPI.SyncPlayPause(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayPause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayPauseRequest struct via the builder pattern


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


## SyncPlayPing

> SyncPlayPing(ctx).PingRequestDto(pingRequestDto).Execute()

Update session ping.

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
	pingRequestDto := *openapiclient.NewPingRequestDto() // PingRequestDto | The new ping.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayPing(context.Background()).PingRequestDto(pingRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pingRequestDto** | [**PingRequestDto**](PingRequestDto.md) | The new ping. | 

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


## SyncPlayPreviousItem

> SyncPlayPreviousItem(ctx).PreviousItemRequestDto(previousItemRequestDto).Execute()

Request previous item in SyncPlay group.

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
	previousItemRequestDto := *openapiclient.NewPreviousItemRequestDto() // PreviousItemRequestDto | The current item information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayPreviousItem(context.Background()).PreviousItemRequestDto(previousItemRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayPreviousItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayPreviousItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **previousItemRequestDto** | [**PreviousItemRequestDto**](PreviousItemRequestDto.md) | The current item information. | 

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


## SyncPlayQueue

> SyncPlayQueue(ctx).QueueRequestDto(queueRequestDto).Execute()

Request to queue items to the playlist of a SyncPlay group.

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
	queueRequestDto := *openapiclient.NewQueueRequestDto() // QueueRequestDto | The items to add.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayQueue(context.Background()).QueueRequestDto(queueRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queueRequestDto** | [**QueueRequestDto**](QueueRequestDto.md) | The items to add. | 

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


## SyncPlayReady

> SyncPlayReady(ctx).ReadyRequestDto(readyRequestDto).Execute()

Notify SyncPlay group that member is ready for playback.

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
	readyRequestDto := *openapiclient.NewReadyRequestDto() // ReadyRequestDto | The player status.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayReady(context.Background()).ReadyRequestDto(readyRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayReady``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayReadyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **readyRequestDto** | [**ReadyRequestDto**](ReadyRequestDto.md) | The player status. | 

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


## SyncPlayRemoveFromPlaylist

> SyncPlayRemoveFromPlaylist(ctx).RemoveFromPlaylistRequestDto(removeFromPlaylistRequestDto).Execute()

Request to remove items from the playlist in SyncPlay group.

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
	removeFromPlaylistRequestDto := *openapiclient.NewRemoveFromPlaylistRequestDto() // RemoveFromPlaylistRequestDto | The items to remove.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlayRemoveFromPlaylist(context.Background()).RemoveFromPlaylistRequestDto(removeFromPlaylistRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayRemoveFromPlaylist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayRemoveFromPlaylistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **removeFromPlaylistRequestDto** | [**RemoveFromPlaylistRequestDto**](RemoveFromPlaylistRequestDto.md) | The items to remove. | 

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


## SyncPlaySeek

> SyncPlaySeek(ctx).SeekRequestDto(seekRequestDto).Execute()

Request seek in SyncPlay group.

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
	seekRequestDto := *openapiclient.NewSeekRequestDto() // SeekRequestDto | The new playback position.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlaySeek(context.Background()).SeekRequestDto(seekRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlaySeek``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlaySeekRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seekRequestDto** | [**SeekRequestDto**](SeekRequestDto.md) | The new playback position. | 

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


## SyncPlaySetIgnoreWait

> SyncPlaySetIgnoreWait(ctx).IgnoreWaitRequestDto(ignoreWaitRequestDto).Execute()

Request SyncPlay group to ignore member during group-wait.

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
	ignoreWaitRequestDto := *openapiclient.NewIgnoreWaitRequestDto() // IgnoreWaitRequestDto | The settings to set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlaySetIgnoreWait(context.Background()).IgnoreWaitRequestDto(ignoreWaitRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlaySetIgnoreWait``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlaySetIgnoreWaitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ignoreWaitRequestDto** | [**IgnoreWaitRequestDto**](IgnoreWaitRequestDto.md) | The settings to set. | 

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


## SyncPlaySetNewQueue

> SyncPlaySetNewQueue(ctx).PlayRequestDto(playRequestDto).Execute()

Request to set new playlist in SyncPlay group.

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
	playRequestDto := *openapiclient.NewPlayRequestDto() // PlayRequestDto | The new playlist to play in the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlaySetNewQueue(context.Background()).PlayRequestDto(playRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlaySetNewQueue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlaySetNewQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playRequestDto** | [**PlayRequestDto**](PlayRequestDto.md) | The new playlist to play in the group. | 

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


## SyncPlaySetPlaylistItem

> SyncPlaySetPlaylistItem(ctx).SetPlaylistItemRequestDto(setPlaylistItemRequestDto).Execute()

Request to change playlist item in SyncPlay group.

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
	setPlaylistItemRequestDto := *openapiclient.NewSetPlaylistItemRequestDto() // SetPlaylistItemRequestDto | The new item to play.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlaySetPlaylistItem(context.Background()).SetPlaylistItemRequestDto(setPlaylistItemRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlaySetPlaylistItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlaySetPlaylistItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setPlaylistItemRequestDto** | [**SetPlaylistItemRequestDto**](SetPlaylistItemRequestDto.md) | The new item to play. | 

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


## SyncPlaySetRepeatMode

> SyncPlaySetRepeatMode(ctx).SetRepeatModeRequestDto(setRepeatModeRequestDto).Execute()

Request to set repeat mode in SyncPlay group.

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
	setRepeatModeRequestDto := *openapiclient.NewSetRepeatModeRequestDto() // SetRepeatModeRequestDto | The new repeat mode.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlaySetRepeatMode(context.Background()).SetRepeatModeRequestDto(setRepeatModeRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlaySetRepeatMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlaySetRepeatModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setRepeatModeRequestDto** | [**SetRepeatModeRequestDto**](SetRepeatModeRequestDto.md) | The new repeat mode. | 

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


## SyncPlaySetShuffleMode

> SyncPlaySetShuffleMode(ctx).SetShuffleModeRequestDto(setShuffleModeRequestDto).Execute()

Request to set shuffle mode in SyncPlay group.

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
	setShuffleModeRequestDto := *openapiclient.NewSetShuffleModeRequestDto() // SetShuffleModeRequestDto | The new shuffle mode.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SyncPlayAPI.SyncPlaySetShuffleMode(context.Background()).SetShuffleModeRequestDto(setShuffleModeRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlaySetShuffleMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlaySetShuffleModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setShuffleModeRequestDto** | [**SetShuffleModeRequestDto**](SetShuffleModeRequestDto.md) | The new shuffle mode. | 

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


## SyncPlayStop

> SyncPlayStop(ctx).Execute()

Request stop in SyncPlay group.

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
	r, err := apiClient.SyncPlayAPI.SyncPlayStop(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayStopRequest struct via the builder pattern


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


## SyncPlayUnpause

> SyncPlayUnpause(ctx).Execute()

Request unpause in SyncPlay group.

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
	r, err := apiClient.SyncPlayAPI.SyncPlayUnpause(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyncPlayAPI.SyncPlayUnpause``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncPlayUnpauseRequest struct via the builder pattern


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

