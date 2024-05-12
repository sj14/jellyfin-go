# \ItemLookupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplySearchCriteria**](ItemLookupAPI.md#ApplySearchCriteria) | **Post** /Items/RemoteSearch/Apply/{itemId} | Applies search criteria to an item and refreshes metadata.
[**GetBookRemoteSearchResults**](ItemLookupAPI.md#GetBookRemoteSearchResults) | **Post** /Items/RemoteSearch/Book | Get book remote search.
[**GetBoxSetRemoteSearchResults**](ItemLookupAPI.md#GetBoxSetRemoteSearchResults) | **Post** /Items/RemoteSearch/BoxSet | Get box set remote search.
[**GetExternalIdInfos**](ItemLookupAPI.md#GetExternalIdInfos) | **Get** /Items/{itemId}/ExternalIdInfos | Get the item&#39;s external id info.
[**GetMovieRemoteSearchResults**](ItemLookupAPI.md#GetMovieRemoteSearchResults) | **Post** /Items/RemoteSearch/Movie | Get movie remote search.
[**GetMusicAlbumRemoteSearchResults**](ItemLookupAPI.md#GetMusicAlbumRemoteSearchResults) | **Post** /Items/RemoteSearch/MusicAlbum | Get music album remote search.
[**GetMusicArtistRemoteSearchResults**](ItemLookupAPI.md#GetMusicArtistRemoteSearchResults) | **Post** /Items/RemoteSearch/MusicArtist | Get music artist remote search.
[**GetMusicVideoRemoteSearchResults**](ItemLookupAPI.md#GetMusicVideoRemoteSearchResults) | **Post** /Items/RemoteSearch/MusicVideo | Get music video remote search.
[**GetPersonRemoteSearchResults**](ItemLookupAPI.md#GetPersonRemoteSearchResults) | **Post** /Items/RemoteSearch/Person | Get person remote search.
[**GetSeriesRemoteSearchResults**](ItemLookupAPI.md#GetSeriesRemoteSearchResults) | **Post** /Items/RemoteSearch/Series | Get series remote search.
[**GetTrailerRemoteSearchResults**](ItemLookupAPI.md#GetTrailerRemoteSearchResults) | **Post** /Items/RemoteSearch/Trailer | Get trailer remote search.



## ApplySearchCriteria

> ApplySearchCriteria(ctx, itemId).RemoteSearchResult(remoteSearchResult).ReplaceAllImages(replaceAllImages).Execute()

Applies search criteria to an item and refreshes metadata.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.
	remoteSearchResult := *openapiclient.NewRemoteSearchResult() // RemoteSearchResult | The remote search result.
	replaceAllImages := true // bool | Optional. Whether or not to replace all images. Default: True. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ItemLookupAPI.ApplySearchCriteria(context.Background(), itemId).RemoteSearchResult(remoteSearchResult).ReplaceAllImages(replaceAllImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.ApplySearchCriteria``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplySearchCriteriaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteSearchResult** | [**RemoteSearchResult**](RemoteSearchResult.md) | The remote search result. | 
 **replaceAllImages** | **bool** | Optional. Whether or not to replace all images. Default: True. | [default to true]

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


## GetBookRemoteSearchResults

> []RemoteSearchResult GetBookRemoteSearchResults(ctx).BookInfoRemoteSearchQuery(bookInfoRemoteSearchQuery).Execute()

Get book remote search.

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
	bookInfoRemoteSearchQuery := *openapiclient.NewBookInfoRemoteSearchQuery() // BookInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetBookRemoteSearchResults(context.Background()).BookInfoRemoteSearchQuery(bookInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetBookRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBookRemoteSearchResults`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetBookRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBookRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookInfoRemoteSearchQuery** | [**BookInfoRemoteSearchQuery**](BookInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoxSetRemoteSearchResults

> []RemoteSearchResult GetBoxSetRemoteSearchResults(ctx).BoxSetInfoRemoteSearchQuery(boxSetInfoRemoteSearchQuery).Execute()

Get box set remote search.

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
	boxSetInfoRemoteSearchQuery := *openapiclient.NewBoxSetInfoRemoteSearchQuery() // BoxSetInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetBoxSetRemoteSearchResults(context.Background()).BoxSetInfoRemoteSearchQuery(boxSetInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetBoxSetRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBoxSetRemoteSearchResults`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetBoxSetRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBoxSetRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **boxSetInfoRemoteSearchQuery** | [**BoxSetInfoRemoteSearchQuery**](BoxSetInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalIdInfos

> []ExternalIdInfo GetExternalIdInfos(ctx, itemId).Execute()

Get the item's external id info.

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
	itemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Item id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetExternalIdInfos(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetExternalIdInfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExternalIdInfos`: []ExternalIdInfo
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetExternalIdInfos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalIdInfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ExternalIdInfo**](ExternalIdInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMovieRemoteSearchResults

> []RemoteSearchResult GetMovieRemoteSearchResults(ctx).MovieInfoRemoteSearchQuery(movieInfoRemoteSearchQuery).Execute()

Get movie remote search.

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
	movieInfoRemoteSearchQuery := *openapiclient.NewMovieInfoRemoteSearchQuery() // MovieInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetMovieRemoteSearchResults(context.Background()).MovieInfoRemoteSearchQuery(movieInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetMovieRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMovieRemoteSearchResults`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetMovieRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMovieRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **movieInfoRemoteSearchQuery** | [**MovieInfoRemoteSearchQuery**](MovieInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMusicAlbumRemoteSearchResults

> []RemoteSearchResult GetMusicAlbumRemoteSearchResults(ctx).AlbumInfoRemoteSearchQuery(albumInfoRemoteSearchQuery).Execute()

Get music album remote search.

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
	albumInfoRemoteSearchQuery := *openapiclient.NewAlbumInfoRemoteSearchQuery() // AlbumInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetMusicAlbumRemoteSearchResults(context.Background()).AlbumInfoRemoteSearchQuery(albumInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetMusicAlbumRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMusicAlbumRemoteSearchResults`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetMusicAlbumRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMusicAlbumRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **albumInfoRemoteSearchQuery** | [**AlbumInfoRemoteSearchQuery**](AlbumInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMusicArtistRemoteSearchResults

> []RemoteSearchResult GetMusicArtistRemoteSearchResults(ctx).ArtistInfoRemoteSearchQuery(artistInfoRemoteSearchQuery).Execute()

Get music artist remote search.

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
	artistInfoRemoteSearchQuery := *openapiclient.NewArtistInfoRemoteSearchQuery() // ArtistInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetMusicArtistRemoteSearchResults(context.Background()).ArtistInfoRemoteSearchQuery(artistInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetMusicArtistRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMusicArtistRemoteSearchResults`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetMusicArtistRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMusicArtistRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistInfoRemoteSearchQuery** | [**ArtistInfoRemoteSearchQuery**](ArtistInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMusicVideoRemoteSearchResults

> []RemoteSearchResult GetMusicVideoRemoteSearchResults(ctx).MusicVideoInfoRemoteSearchQuery(musicVideoInfoRemoteSearchQuery).Execute()

Get music video remote search.

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
	musicVideoInfoRemoteSearchQuery := *openapiclient.NewMusicVideoInfoRemoteSearchQuery() // MusicVideoInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetMusicVideoRemoteSearchResults(context.Background()).MusicVideoInfoRemoteSearchQuery(musicVideoInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetMusicVideoRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMusicVideoRemoteSearchResults`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetMusicVideoRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMusicVideoRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **musicVideoInfoRemoteSearchQuery** | [**MusicVideoInfoRemoteSearchQuery**](MusicVideoInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonRemoteSearchResults

> []RemoteSearchResult GetPersonRemoteSearchResults(ctx).PersonLookupInfoRemoteSearchQuery(personLookupInfoRemoteSearchQuery).Execute()

Get person remote search.

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
	personLookupInfoRemoteSearchQuery := *openapiclient.NewPersonLookupInfoRemoteSearchQuery() // PersonLookupInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetPersonRemoteSearchResults(context.Background()).PersonLookupInfoRemoteSearchQuery(personLookupInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetPersonRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPersonRemoteSearchResults`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetPersonRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **personLookupInfoRemoteSearchQuery** | [**PersonLookupInfoRemoteSearchQuery**](PersonLookupInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSeriesRemoteSearchResults

> []RemoteSearchResult GetSeriesRemoteSearchResults(ctx).SeriesInfoRemoteSearchQuery(seriesInfoRemoteSearchQuery).Execute()

Get series remote search.

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
	seriesInfoRemoteSearchQuery := *openapiclient.NewSeriesInfoRemoteSearchQuery() // SeriesInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetSeriesRemoteSearchResults(context.Background()).SeriesInfoRemoteSearchQuery(seriesInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetSeriesRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSeriesRemoteSearchResults`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetSeriesRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSeriesRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **seriesInfoRemoteSearchQuery** | [**SeriesInfoRemoteSearchQuery**](SeriesInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrailerRemoteSearchResults

> []RemoteSearchResult GetTrailerRemoteSearchResults(ctx).TrailerInfoRemoteSearchQuery(trailerInfoRemoteSearchQuery).Execute()

Get trailer remote search.

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
	trailerInfoRemoteSearchQuery := *openapiclient.NewTrailerInfoRemoteSearchQuery() // TrailerInfoRemoteSearchQuery | Remote search query.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ItemLookupAPI.GetTrailerRemoteSearchResults(context.Background()).TrailerInfoRemoteSearchQuery(trailerInfoRemoteSearchQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ItemLookupAPI.GetTrailerRemoteSearchResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrailerRemoteSearchResults`: []RemoteSearchResult
	fmt.Fprintf(os.Stdout, "Response from `ItemLookupAPI.GetTrailerRemoteSearchResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrailerRemoteSearchResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trailerInfoRemoteSearchQuery** | [**TrailerInfoRemoteSearchQuery**](TrailerInfoRemoteSearchQuery.md) | Remote search query. | 

### Return type

[**[]RemoteSearchResult**](RemoteSearchResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

