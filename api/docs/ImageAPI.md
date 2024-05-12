# \ImageAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCustomSplashscreen**](ImageAPI.md#DeleteCustomSplashscreen) | **Delete** /Branding/Splashscreen | Delete a custom splashscreen.
[**DeleteItemImage**](ImageAPI.md#DeleteItemImage) | **Delete** /Items/{itemId}/Images/{imageType} | Delete an item&#39;s image.
[**DeleteItemImageByIndex**](ImageAPI.md#DeleteItemImageByIndex) | **Delete** /Items/{itemId}/Images/{imageType}/{imageIndex} | Delete an item&#39;s image.
[**DeleteUserImage**](ImageAPI.md#DeleteUserImage) | **Delete** /UserImage | Delete the user&#39;s image.
[**GetArtistImage**](ImageAPI.md#GetArtistImage) | **Get** /Artists/{name}/Images/{imageType}/{imageIndex} | Get artist image by name.
[**GetGenreImage**](ImageAPI.md#GetGenreImage) | **Get** /Genres/{name}/Images/{imageType} | Get genre image by name.
[**GetGenreImageByIndex**](ImageAPI.md#GetGenreImageByIndex) | **Get** /Genres/{name}/Images/{imageType}/{imageIndex} | Get genre image by name.
[**GetItemImage**](ImageAPI.md#GetItemImage) | **Get** /Items/{itemId}/Images/{imageType} | Gets the item&#39;s image.
[**GetItemImage2**](ImageAPI.md#GetItemImage2) | **Get** /Items/{itemId}/Images/{imageType}/{imageIndex}/{tag}/{format}/{maxWidth}/{maxHeight}/{percentPlayed}/{unplayedCount} | Gets the item&#39;s image.
[**GetItemImageByIndex**](ImageAPI.md#GetItemImageByIndex) | **Get** /Items/{itemId}/Images/{imageType}/{imageIndex} | Gets the item&#39;s image.
[**GetItemImageInfos**](ImageAPI.md#GetItemImageInfos) | **Get** /Items/{itemId}/Images | Get item image infos.
[**GetMusicGenreImage**](ImageAPI.md#GetMusicGenreImage) | **Get** /MusicGenres/{name}/Images/{imageType} | Get music genre image by name.
[**GetMusicGenreImageByIndex**](ImageAPI.md#GetMusicGenreImageByIndex) | **Get** /MusicGenres/{name}/Images/{imageType}/{imageIndex} | Get music genre image by name.
[**GetPersonImage**](ImageAPI.md#GetPersonImage) | **Get** /Persons/{name}/Images/{imageType} | Get person image by name.
[**GetPersonImageByIndex**](ImageAPI.md#GetPersonImageByIndex) | **Get** /Persons/{name}/Images/{imageType}/{imageIndex} | Get person image by name.
[**GetSplashscreen**](ImageAPI.md#GetSplashscreen) | **Get** /Branding/Splashscreen | Generates or gets the splashscreen.
[**GetStudioImage**](ImageAPI.md#GetStudioImage) | **Get** /Studios/{name}/Images/{imageType} | Get studio image by name.
[**GetStudioImageByIndex**](ImageAPI.md#GetStudioImageByIndex) | **Get** /Studios/{name}/Images/{imageType}/{imageIndex} | Get studio image by name.
[**GetUserImage**](ImageAPI.md#GetUserImage) | **Get** /UserImage | Get user profile image.
[**HeadArtistImage**](ImageAPI.md#HeadArtistImage) | **Head** /Artists/{name}/Images/{imageType}/{imageIndex} | Get artist image by name.
[**HeadGenreImage**](ImageAPI.md#HeadGenreImage) | **Head** /Genres/{name}/Images/{imageType} | Get genre image by name.
[**HeadGenreImageByIndex**](ImageAPI.md#HeadGenreImageByIndex) | **Head** /Genres/{name}/Images/{imageType}/{imageIndex} | Get genre image by name.
[**HeadItemImage**](ImageAPI.md#HeadItemImage) | **Head** /Items/{itemId}/Images/{imageType} | Gets the item&#39;s image.
[**HeadItemImage2**](ImageAPI.md#HeadItemImage2) | **Head** /Items/{itemId}/Images/{imageType}/{imageIndex}/{tag}/{format}/{maxWidth}/{maxHeight}/{percentPlayed}/{unplayedCount} | Gets the item&#39;s image.
[**HeadItemImageByIndex**](ImageAPI.md#HeadItemImageByIndex) | **Head** /Items/{itemId}/Images/{imageType}/{imageIndex} | Gets the item&#39;s image.
[**HeadMusicGenreImage**](ImageAPI.md#HeadMusicGenreImage) | **Head** /MusicGenres/{name}/Images/{imageType} | Get music genre image by name.
[**HeadMusicGenreImageByIndex**](ImageAPI.md#HeadMusicGenreImageByIndex) | **Head** /MusicGenres/{name}/Images/{imageType}/{imageIndex} | Get music genre image by name.
[**HeadPersonImage**](ImageAPI.md#HeadPersonImage) | **Head** /Persons/{name}/Images/{imageType} | Get person image by name.
[**HeadPersonImageByIndex**](ImageAPI.md#HeadPersonImageByIndex) | **Head** /Persons/{name}/Images/{imageType}/{imageIndex} | Get person image by name.
[**HeadStudioImage**](ImageAPI.md#HeadStudioImage) | **Head** /Studios/{name}/Images/{imageType} | Get studio image by name.
[**HeadStudioImageByIndex**](ImageAPI.md#HeadStudioImageByIndex) | **Head** /Studios/{name}/Images/{imageType}/{imageIndex} | Get studio image by name.
[**HeadUserImage**](ImageAPI.md#HeadUserImage) | **Head** /UserImage | Get user profile image.
[**PostUserImage**](ImageAPI.md#PostUserImage) | **Post** /UserImage | Sets the user image.
[**SetItemImage**](ImageAPI.md#SetItemImage) | **Post** /Items/{itemId}/Images/{imageType} | Set item image.
[**SetItemImageByIndex**](ImageAPI.md#SetItemImageByIndex) | **Post** /Items/{itemId}/Images/{imageType}/{imageIndex} | Set item image.
[**UpdateItemImageIndex**](ImageAPI.md#UpdateItemImageIndex) | **Post** /Items/{itemId}/Images/{imageType}/{imageIndex}/Index | Updates the index for an item image.
[**UploadCustomSplashscreen**](ImageAPI.md#UploadCustomSplashscreen) | **Post** /Branding/Splashscreen | Uploads a custom splashscreen.  The body is expected to the image contents base64 encoded.



## DeleteCustomSplashscreen

> DeleteCustomSplashscreen(ctx).Execute()

Delete a custom splashscreen.

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
	r, err := apiClient.ImageAPI.DeleteCustomSplashscreen(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.DeleteCustomSplashscreen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomSplashscreenRequest struct via the builder pattern


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


## DeleteItemImage

> DeleteItemImage(ctx, itemId, imageType).ImageIndex(imageIndex).Execute()

Delete an item's image.

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
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | The image index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageAPI.DeleteItemImage(context.Background(), itemId, imageType).ImageIndex(imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.DeleteItemImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 
**imageType** | **ImageType** | Image type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **imageIndex** | **int32** | The image index. | 

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


## DeleteItemImageByIndex

> DeleteItemImageByIndex(ctx, itemId, imageType, imageIndex).Execute()

Delete an item's image.

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
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | The image index.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageAPI.DeleteItemImageByIndex(context.Background(), itemId, imageType, imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.DeleteItemImageByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | The image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteItemImageByIndexRequest struct via the builder pattern


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


## DeleteUserImage

> DeleteUserImage(ctx).UserId(userId).Execute()

Delete the user's image.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User Id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageAPI.DeleteUserImage(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.DeleteUserImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User Id. | 

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


## GetArtistImage

> *os.File GetArtistImage(ctx, name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Get artist image by name.

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
	name := "name_example" // string | Artist name.
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | Image index.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetArtistImage(context.Background(), name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetArtistImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtistImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetArtistImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Artist name. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtistImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenreImage

> *os.File GetGenreImage(ctx, name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()

Get genre image by name.

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
	name := "name_example" // string | Genre name.
	imageType := "imageType_example" // ImageType | Image type.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)
	imageIndex := int32(56) // int32 | Image index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetGenreImage(context.Background(), name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetGenreImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenreImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetGenreImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Genre name. | 
**imageType** | **ImageType** | Image type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenreImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 
 **imageIndex** | **int32** | Image index. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGenreImageByIndex

> *os.File GetGenreImageByIndex(ctx, name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Get genre image by name.

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
	name := "name_example" // string | Genre name.
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | Image index.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetGenreImageByIndex(context.Background(), name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetGenreImageByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenreImageByIndex`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetGenreImageByIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Genre name. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenreImageByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemImage

> *os.File GetItemImage(ctx, itemId, imageType).MaxWidth(maxWidth).MaxHeight(maxHeight).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Tag(tag).Format(format).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()

Gets the item's image.

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
	imageType := "imageType_example" // ImageType | Image type.
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Optional. The MediaBrowser.Model.Drawing.ImageFormat of the returned image. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)
	imageIndex := int32(56) // int32 | Image index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetItemImage(context.Background(), itemId, imageType).MaxWidth(maxWidth).MaxHeight(maxHeight).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Tag(tag).Format(format).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetItemImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetItemImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 
**imageType** | **ImageType** | Image type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Optional. The MediaBrowser.Model.Drawing.ImageFormat of the returned image. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 
 **imageIndex** | **int32** | Image index. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemImage2

> *os.File GetItemImage2(ctx, itemId, imageType, maxWidth, maxHeight, tag, format, percentPlayed, unplayedCount, imageIndex).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Gets the item's image.

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
	imageType := "imageType_example" // ImageType | Image type.
	maxWidth := int32(56) // int32 | The maximum image width to return.
	maxHeight := int32(56) // int32 | The maximum image height to return.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers.
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png.
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay.
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render.
	imageIndex := int32(56) // int32 | Image index.
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetItemImage2(context.Background(), itemId, imageType, maxWidth, maxHeight, tag, format, percentPlayed, unplayedCount, imageIndex).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetItemImage2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemImage2`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetItemImage2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 
**imageType** | **ImageType** | Image type. | 
**maxWidth** | **int32** | The maximum image width to return. | 
**maxHeight** | **int32** | The maximum image height to return. | 
**tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
**format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
**percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
**unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemImage2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------









 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemImageByIndex

> *os.File GetItemImageByIndex(ctx, itemId, imageType, imageIndex).MaxWidth(maxWidth).MaxHeight(maxHeight).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Tag(tag).Format(format).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Gets the item's image.

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
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | Image index.
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Optional. The MediaBrowser.Model.Drawing.ImageFormat of the returned image. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetItemImageByIndex(context.Background(), itemId, imageType, imageIndex).MaxWidth(maxWidth).MaxHeight(maxHeight).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Tag(tag).Format(format).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetItemImageByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemImageByIndex`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetItemImageByIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemImageByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Optional. The MediaBrowser.Model.Drawing.ImageFormat of the returned image. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetItemImageInfos

> []ImageInfo GetItemImageInfos(ctx, itemId).Execute()

Get item image infos.

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
	resp, r, err := apiClient.ImageAPI.GetItemImageInfos(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetItemImageInfos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemImageInfos`: []ImageInfo
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetItemImageInfos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemImageInfosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ImageInfo**](ImageInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMusicGenreImage

> *os.File GetMusicGenreImage(ctx, name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()

Get music genre image by name.

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
	name := "name_example" // string | Music genre name.
	imageType := "imageType_example" // ImageType | Image type.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)
	imageIndex := int32(56) // int32 | Image index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetMusicGenreImage(context.Background(), name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetMusicGenreImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMusicGenreImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetMusicGenreImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Music genre name. | 
**imageType** | **ImageType** | Image type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMusicGenreImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 
 **imageIndex** | **int32** | Image index. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMusicGenreImageByIndex

> *os.File GetMusicGenreImageByIndex(ctx, name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Get music genre image by name.

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
	name := "name_example" // string | Music genre name.
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | Image index.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetMusicGenreImageByIndex(context.Background(), name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetMusicGenreImageByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMusicGenreImageByIndex`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetMusicGenreImageByIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Music genre name. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMusicGenreImageByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonImage

> *os.File GetPersonImage(ctx, name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()

Get person image by name.

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
	name := "name_example" // string | Person name.
	imageType := "imageType_example" // ImageType | Image type.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)
	imageIndex := int32(56) // int32 | Image index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetPersonImage(context.Background(), name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetPersonImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPersonImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetPersonImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Person name. | 
**imageType** | **ImageType** | Image type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 
 **imageIndex** | **int32** | Image index. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonImageByIndex

> *os.File GetPersonImageByIndex(ctx, name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Get person image by name.

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
	name := "name_example" // string | Person name.
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | Image index.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetPersonImageByIndex(context.Background(), name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetPersonImageByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPersonImageByIndex`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetPersonImageByIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Person name. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonImageByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSplashscreen

> *os.File GetSplashscreen(ctx).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).Width(width).Height(height).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Quality(quality).Execute()

Generates or gets the splashscreen.

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
	tag := "tag_example" // string | Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Apply a foreground layer on top of the image. (optional)
	quality := int32(56) // int32 | Quality setting, from 0-100. (optional) (default to 90)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetSplashscreen(context.Background()).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).Width(width).Height(height).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Quality(quality).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetSplashscreen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSplashscreen`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetSplashscreen`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSplashscreenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | **string** | Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Blur image. | 
 **backgroundColor** | **string** | Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Apply a foreground layer on top of the image. | 
 **quality** | **int32** | Quality setting, from 0-100. | [default to 90]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioImage

> *os.File GetStudioImage(ctx, name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()

Get studio image by name.

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
	name := "name_example" // string | Studio name.
	imageType := "imageType_example" // ImageType | Image type.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)
	imageIndex := int32(56) // int32 | Image index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetStudioImage(context.Background(), name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetStudioImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetStudioImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Studio name. | 
**imageType** | **ImageType** | Image type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 
 **imageIndex** | **int32** | Image index. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStudioImageByIndex

> *os.File GetStudioImageByIndex(ctx, name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Get studio image by name.

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
	name := "name_example" // string | Studio name.
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | Image index.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetStudioImageByIndex(context.Background(), name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetStudioImageByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStudioImageByIndex`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetStudioImageByIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Studio name. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStudioImageByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserImage

> *os.File GetUserImage(ctx).UserId(userId).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()

Get user profile image.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id. (optional)
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)
	imageIndex := int32(56) // int32 | Image index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetUserImage(context.Background()).UserId(userId).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetUserImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetUserImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User id. | 
 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 
 **imageIndex** | **int32** | Image index. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadArtistImage

> *os.File HeadArtistImage(ctx, name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Get artist image by name.

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
	name := "name_example" // string | Artist name.
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | Image index.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.HeadArtistImage(context.Background(), name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.HeadArtistImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadArtistImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.HeadArtistImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Artist name. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadArtistImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadGenreImage

> *os.File HeadGenreImage(ctx, name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()

Get genre image by name.

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
	name := "name_example" // string | Genre name.
	imageType := "imageType_example" // ImageType | Image type.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)
	imageIndex := int32(56) // int32 | Image index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.HeadGenreImage(context.Background(), name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.HeadGenreImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadGenreImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.HeadGenreImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Genre name. | 
**imageType** | **ImageType** | Image type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadGenreImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 
 **imageIndex** | **int32** | Image index. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadGenreImageByIndex

> *os.File HeadGenreImageByIndex(ctx, name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Get genre image by name.

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
	name := "name_example" // string | Genre name.
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | Image index.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.HeadGenreImageByIndex(context.Background(), name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.HeadGenreImageByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadGenreImageByIndex`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.HeadGenreImageByIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Genre name. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadGenreImageByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadItemImage

> *os.File HeadItemImage(ctx, itemId, imageType).MaxWidth(maxWidth).MaxHeight(maxHeight).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Tag(tag).Format(format).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()

Gets the item's image.

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
	imageType := "imageType_example" // ImageType | Image type.
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Optional. The MediaBrowser.Model.Drawing.ImageFormat of the returned image. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)
	imageIndex := int32(56) // int32 | Image index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.HeadItemImage(context.Background(), itemId, imageType).MaxWidth(maxWidth).MaxHeight(maxHeight).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Tag(tag).Format(format).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.HeadItemImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadItemImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.HeadItemImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 
**imageType** | **ImageType** | Image type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadItemImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Optional. The MediaBrowser.Model.Drawing.ImageFormat of the returned image. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 
 **imageIndex** | **int32** | Image index. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadItemImage2

> *os.File HeadItemImage2(ctx, itemId, imageType, maxWidth, maxHeight, tag, format, percentPlayed, unplayedCount, imageIndex).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Gets the item's image.

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
	imageType := "imageType_example" // ImageType | Image type.
	maxWidth := int32(56) // int32 | The maximum image width to return.
	maxHeight := int32(56) // int32 | The maximum image height to return.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers.
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png.
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay.
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render.
	imageIndex := int32(56) // int32 | Image index.
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.HeadItemImage2(context.Background(), itemId, imageType, maxWidth, maxHeight, tag, format, percentPlayed, unplayedCount, imageIndex).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.HeadItemImage2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadItemImage2`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.HeadItemImage2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 
**imageType** | **ImageType** | Image type. | 
**maxWidth** | **int32** | The maximum image width to return. | 
**maxHeight** | **int32** | The maximum image height to return. | 
**tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
**format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
**percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
**unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadItemImage2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------









 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadItemImageByIndex

> *os.File HeadItemImageByIndex(ctx, itemId, imageType, imageIndex).MaxWidth(maxWidth).MaxHeight(maxHeight).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Tag(tag).Format(format).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Gets the item's image.

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
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | Image index.
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Optional. The MediaBrowser.Model.Drawing.ImageFormat of the returned image. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.HeadItemImageByIndex(context.Background(), itemId, imageType, imageIndex).MaxWidth(maxWidth).MaxHeight(maxHeight).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Tag(tag).Format(format).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.HeadItemImageByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadItemImageByIndex`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.HeadItemImageByIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadItemImageByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Optional. The MediaBrowser.Model.Drawing.ImageFormat of the returned image. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadMusicGenreImage

> *os.File HeadMusicGenreImage(ctx, name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()

Get music genre image by name.

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
	name := "name_example" // string | Music genre name.
	imageType := "imageType_example" // ImageType | Image type.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)
	imageIndex := int32(56) // int32 | Image index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.HeadMusicGenreImage(context.Background(), name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.HeadMusicGenreImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadMusicGenreImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.HeadMusicGenreImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Music genre name. | 
**imageType** | **ImageType** | Image type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadMusicGenreImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 
 **imageIndex** | **int32** | Image index. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadMusicGenreImageByIndex

> *os.File HeadMusicGenreImageByIndex(ctx, name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Get music genre image by name.

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
	name := "name_example" // string | Music genre name.
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | Image index.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.HeadMusicGenreImageByIndex(context.Background(), name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.HeadMusicGenreImageByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadMusicGenreImageByIndex`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.HeadMusicGenreImageByIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Music genre name. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadMusicGenreImageByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadPersonImage

> *os.File HeadPersonImage(ctx, name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()

Get person image by name.

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
	name := "name_example" // string | Person name.
	imageType := "imageType_example" // ImageType | Image type.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)
	imageIndex := int32(56) // int32 | Image index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.HeadPersonImage(context.Background(), name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.HeadPersonImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadPersonImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.HeadPersonImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Person name. | 
**imageType** | **ImageType** | Image type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadPersonImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 
 **imageIndex** | **int32** | Image index. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadPersonImageByIndex

> *os.File HeadPersonImageByIndex(ctx, name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Get person image by name.

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
	name := "name_example" // string | Person name.
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | Image index.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.HeadPersonImageByIndex(context.Background(), name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.HeadPersonImageByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadPersonImageByIndex`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.HeadPersonImageByIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Person name. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadPersonImageByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadStudioImage

> *os.File HeadStudioImage(ctx, name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()

Get studio image by name.

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
	name := "name_example" // string | Studio name.
	imageType := "imageType_example" // ImageType | Image type.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)
	imageIndex := int32(56) // int32 | Image index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.HeadStudioImage(context.Background(), name, imageType).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.HeadStudioImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadStudioImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.HeadStudioImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Studio name. | 
**imageType** | **ImageType** | Image type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadStudioImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 
 **imageIndex** | **int32** | Image index. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadStudioImageByIndex

> *os.File HeadStudioImageByIndex(ctx, name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()

Get studio image by name.

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
	name := "name_example" // string | Studio name.
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | Image index.
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.HeadStudioImageByIndex(context.Background(), name, imageType, imageIndex).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.HeadStudioImageByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadStudioImageByIndex`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.HeadStudioImageByIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Studio name. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadStudioImageByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadUserImage

> *os.File HeadUserImage(ctx).UserId(userId).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()

Get user profile image.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id. (optional)
	tag := "tag_example" // string | Optional. Supply the cache tag from the item object to receive strong caching headers. (optional)
	format := "format_example" // ImageFormat | Determines the output format of the image - original,gif,jpg,png. (optional)
	maxWidth := int32(56) // int32 | The maximum image width to return. (optional)
	maxHeight := int32(56) // int32 | The maximum image height to return. (optional)
	percentPlayed := float64(1.2) // float64 | Optional. Percent to render for the percent played overlay. (optional)
	unplayedCount := int32(56) // int32 | Optional. Unplayed count overlay to render. (optional)
	width := int32(56) // int32 | The fixed image width to return. (optional)
	height := int32(56) // int32 | The fixed image height to return. (optional)
	quality := int32(56) // int32 | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. (optional)
	fillWidth := int32(56) // int32 | Width of box to fill. (optional)
	fillHeight := int32(56) // int32 | Height of box to fill. (optional)
	blur := int32(56) // int32 | Optional. Blur image. (optional)
	backgroundColor := "backgroundColor_example" // string | Optional. Apply a background color for transparent images. (optional)
	foregroundLayer := "foregroundLayer_example" // string | Optional. Apply a foreground layer on top of the image. (optional)
	imageIndex := int32(56) // int32 | Image index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.HeadUserImage(context.Background()).UserId(userId).Tag(tag).Format(format).MaxWidth(maxWidth).MaxHeight(maxHeight).PercentPlayed(percentPlayed).UnplayedCount(unplayedCount).Width(width).Height(height).Quality(quality).FillWidth(fillWidth).FillHeight(fillHeight).Blur(blur).BackgroundColor(backgroundColor).ForegroundLayer(foregroundLayer).ImageIndex(imageIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.HeadUserImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HeadUserImage`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.HeadUserImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHeadUserImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User id. | 
 **tag** | **string** | Optional. Supply the cache tag from the item object to receive strong caching headers. | 
 **format** | **ImageFormat** | Determines the output format of the image - original,gif,jpg,png. | 
 **maxWidth** | **int32** | The maximum image width to return. | 
 **maxHeight** | **int32** | The maximum image height to return. | 
 **percentPlayed** | **float64** | Optional. Percent to render for the percent played overlay. | 
 **unplayedCount** | **int32** | Optional. Unplayed count overlay to render. | 
 **width** | **int32** | The fixed image width to return. | 
 **height** | **int32** | The fixed image height to return. | 
 **quality** | **int32** | Optional. Quality setting, from 0-100. Defaults to 90 and should suffice in most cases. | 
 **fillWidth** | **int32** | Width of box to fill. | 
 **fillHeight** | **int32** | Height of box to fill. | 
 **blur** | **int32** | Optional. Blur image. | 
 **backgroundColor** | **string** | Optional. Apply a background color for transparent images. | 
 **foregroundLayer** | **string** | Optional. Apply a foreground layer on top of the image. | 
 **imageIndex** | **int32** | Image index. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/*, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUserImage

> PostUserImage(ctx).UserId(userId).Body(body).Execute()

Sets the user image.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User Id. (optional)
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageAPI.PostUserImage(context.Background()).UserId(userId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.PostUserImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUserImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | User Id. | 
 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: image/*
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetItemImage

> SetItemImage(ctx, itemId, imageType).Body(body).Execute()

Set item image.

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
	imageType := "imageType_example" // ImageType | Image type.
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageAPI.SetItemImage(context.Background(), itemId, imageType).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.SetItemImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 
**imageType** | **ImageType** | Image type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetItemImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: image/*
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetItemImageByIndex

> SetItemImageByIndex(ctx, itemId, imageType, imageIndex).Body(body).Execute()

Set item image.

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
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | (Unused) Image index.
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageAPI.SetItemImageByIndex(context.Background(), itemId, imageType, imageIndex).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.SetItemImageByIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | (Unused) Image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetItemImageByIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: image/*
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateItemImageIndex

> UpdateItemImageIndex(ctx, itemId, imageType, imageIndex).NewIndex(newIndex).Execute()

Updates the index for an item image.

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
	imageType := "imageType_example" // ImageType | Image type.
	imageIndex := int32(56) // int32 | Old image index.
	newIndex := int32(56) // int32 | New image index.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageAPI.UpdateItemImageIndex(context.Background(), itemId, imageType, imageIndex).NewIndex(newIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.UpdateItemImageIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** | Item id. | 
**imageType** | **ImageType** | Image type. | 
**imageIndex** | **int32** | Old image index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateItemImageIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **newIndex** | **int32** | New image index. | 

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


## UploadCustomSplashscreen

> UploadCustomSplashscreen(ctx).Body(body).Execute()

Uploads a custom splashscreen.  The body is expected to the image contents base64 encoded.

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
	body := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageAPI.UploadCustomSplashscreen(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.UploadCustomSplashscreen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadCustomSplashscreenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: image/*
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

