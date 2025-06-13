# \VideoAttachmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttachment**](VideoAttachmentsAPI.md#GetAttachment) | **Get** /Videos/{videoId}/{mediaSourceId}/Attachments/{index} | Get video attachment.



## GetAttachment

> *os.File GetAttachment(ctx, videoId, mediaSourceId, index).Execute()

Get video attachment.

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
	videoId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Video ID.
	mediaSourceId := "mediaSourceId_example" // string | Media Source ID.
	index := int32(56) // int32 | Attachment Index.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VideoAttachmentsAPI.GetAttachment(context.Background(), videoId, mediaSourceId, index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VideoAttachmentsAPI.GetAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachment`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `VideoAttachmentsAPI.GetAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | Video ID. | 
**mediaSourceId** | **string** | Media Source ID. | 
**index** | **int32** | Attachment Index. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

