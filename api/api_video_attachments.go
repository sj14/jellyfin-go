/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
)


type VideoAttachmentsAPI interface {

	/*
	GetAttachment Get video attachment.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param videoId Video ID.
	@param mediaSourceId Media Source ID.
	@param index Attachment Index.
	@return ApiGetAttachmentRequest
	*/
	GetAttachment(ctx context.Context, videoId string, mediaSourceId string, index int32) ApiGetAttachmentRequest

	// GetAttachmentExecute executes the request
	//  @return *os.File
	GetAttachmentExecute(r ApiGetAttachmentRequest) (*os.File, *http.Response, error)
}

// VideoAttachmentsAPIService VideoAttachmentsAPI service
type VideoAttachmentsAPIService service

type ApiGetAttachmentRequest struct {
	ctx context.Context
	ApiService VideoAttachmentsAPI
	videoId string
	mediaSourceId string
	index int32
}

func (r ApiGetAttachmentRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.GetAttachmentExecute(r)
}

/*
GetAttachment Get video attachment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param videoId Video ID.
 @param mediaSourceId Media Source ID.
 @param index Attachment Index.
 @return ApiGetAttachmentRequest
*/
func (a *VideoAttachmentsAPIService) GetAttachment(ctx context.Context, videoId string, mediaSourceId string, index int32) ApiGetAttachmentRequest {
	return ApiGetAttachmentRequest{
		ApiService: a,
		ctx: ctx,
		videoId: videoId,
		mediaSourceId: mediaSourceId,
		index: index,
	}
}

// Execute executes the request
//  @return *os.File
func (a *VideoAttachmentsAPIService) GetAttachmentExecute(r ApiGetAttachmentRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VideoAttachmentsAPIService.GetAttachment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Videos/{videoId}/{mediaSourceId}/Attachments/{index}"
	localVarPath = strings.Replace(localVarPath, "{"+"videoId"+"}", url.PathEscape(parameterValueToString(r.videoId, "videoId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mediaSourceId"+"}", url.PathEscape(parameterValueToString(r.mediaSourceId, "mediaSourceId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"index"+"}", url.PathEscape(parameterValueToString(r.index, "index")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/octet-stream", "application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
