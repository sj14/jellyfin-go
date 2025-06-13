# \ChannelsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllChannelFeatures**](ChannelsAPI.md#GetAllChannelFeatures) | **Get** /Channels/Features | Get all channel features.
[**GetChannelFeatures**](ChannelsAPI.md#GetChannelFeatures) | **Get** /Channels/{channelId}/Features | Get channel features.
[**GetChannelItems**](ChannelsAPI.md#GetChannelItems) | **Get** /Channels/{channelId}/Items | Get channel items.
[**GetChannels**](ChannelsAPI.md#GetChannels) | **Get** /Channels | Gets available channels.
[**GetLatestChannelItems**](ChannelsAPI.md#GetLatestChannelItems) | **Get** /Channels/Items/Latest | Gets latest channel items.



## GetAllChannelFeatures

> []ChannelFeatures GetAllChannelFeatures(ctx).Execute()

Get all channel features.

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
	resp, r, err := apiClient.ChannelsAPI.GetAllChannelFeatures(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetAllChannelFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllChannelFeatures`: []ChannelFeatures
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetAllChannelFeatures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllChannelFeaturesRequest struct via the builder pattern


### Return type

[**[]ChannelFeatures**](ChannelFeatures.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelFeatures

> ChannelFeatures GetChannelFeatures(ctx, channelId).Execute()

Get channel features.

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
	channelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Channel id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.GetChannelFeatures(context.Background(), channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannelFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelFeatures`: ChannelFeatures
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannelFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | Channel id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ChannelFeatures**](ChannelFeatures.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannelItems

> BaseItemDtoQueryResult GetChannelItems(ctx, channelId).FolderId(folderId).UserId(userId).StartIndex(startIndex).Limit(limit).SortOrder(sortOrder).Filters(filters).SortBy(sortBy).Fields(fields).Execute()

Get channel items.

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
	channelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Channel Id.
	folderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Folder Id. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. User Id. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	sortOrder := []openapiclient.SortOrder{openapiclient.SortOrder("Ascending")} // []SortOrder | Optional. Sort Order - Ascending,Descending. (optional)
	filters := []openapiclient.ItemFilter{openapiclient.ItemFilter("IsFolder")} // []ItemFilter | Optional. Specify additional filters to apply. (optional)
	sortBy := []openapiclient.ItemSortBy{openapiclient.ItemSortBy("Default")} // []ItemSortBy | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.GetChannelItems(context.Background(), channelId).FolderId(folderId).UserId(userId).StartIndex(startIndex).Limit(limit).SortOrder(sortOrder).Filters(filters).SortBy(sortBy).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannelItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannelItems`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannelItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | Channel Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **folderId** | **string** | Optional. Folder Id. | 
 **userId** | **string** | Optional. User Id. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **sortOrder** | [**[]SortOrder**](SortOrder.md) | Optional. Sort Order - Ascending,Descending. | 
 **filters** | [**[]ItemFilter**](ItemFilter.md) | Optional. Specify additional filters to apply. | 
 **sortBy** | [**[]ItemSortBy**](ItemSortBy.md) | Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. | 

### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannels

> BaseItemDtoQueryResult GetChannels(ctx).UserId(userId).StartIndex(startIndex).Limit(limit).SupportsLatestItems(supportsLatestItems).SupportsMediaDeletion(supportsMediaDeletion).IsFavorite(isFavorite).Execute()

Gets available channels.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User Id to filter by. Use System.Guid.Empty to not filter by user. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	supportsLatestItems := true // bool | Optional. Filter by channels that support getting latest items. (optional)
	supportsMediaDeletion := true // bool | Optional. Filter by channels that support media deletion. (optional)
	isFavorite := true // bool | Optional. Filter by channels that are favorite. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.GetChannels(context.Background()).UserId(userId).StartIndex(startIndex).Limit(limit).SupportsLatestItems(supportsLatestItems).SupportsMediaDeletion(supportsMediaDeletion).IsFavorite(isFavorite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChannels`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User Id to filter by. Use System.Guid.Empty to not filter by user. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **supportsLatestItems** | **bool** | Optional. Filter by channels that support getting latest items. | 
 **supportsMediaDeletion** | **bool** | Optional. Filter by channels that support media deletion. | 
 **isFavorite** | **bool** | Optional. Filter by channels that are favorite. | 

### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestChannelItems

> BaseItemDtoQueryResult GetLatestChannelItems(ctx).UserId(userId).StartIndex(startIndex).Limit(limit).Filters(filters).Fields(fields).ChannelIds(channelIds).Execute()

Gets latest channel items.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. User Id. (optional)
	startIndex := int32(56) // int32 | Optional. The record index to start at. All items with a lower index will be dropped from the results. (optional)
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	filters := []openapiclient.ItemFilter{openapiclient.ItemFilter("IsFolder")} // []ItemFilter | Optional. Specify additional filters to apply. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	channelIds := []string{"Inner_example"} // []string | Optional. Specify one or more channel id's, comma delimited. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelsAPI.GetLatestChannelItems(context.Background()).UserId(userId).StartIndex(startIndex).Limit(limit).Filters(filters).Fields(fields).ChannelIds(channelIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelsAPI.GetLatestChannelItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestChannelItems`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `ChannelsAPI.GetLatestChannelItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestChannelItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | Optional. User Id. | 
 **startIndex** | **int32** | Optional. The record index to start at. All items with a lower index will be dropped from the results. | 
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **filters** | [**[]ItemFilter**](ItemFilter.md) | Optional. Specify additional filters to apply. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **channelIds** | **[]string** | Optional. Specify one or more channel id&#39;s, comma delimited. | 

### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

