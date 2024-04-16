/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.8.13
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
)


type ItemUpdateAPI interface {

	/*
	GetMetadataEditorInfo Gets metadata editor info for an item.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param itemId The item id.
	@return ApiGetMetadataEditorInfoRequest
	*/
	GetMetadataEditorInfo(ctx context.Context, itemId string) ApiGetMetadataEditorInfoRequest

	// GetMetadataEditorInfoExecute executes the request
	//  @return MetadataEditorInfo
	GetMetadataEditorInfoExecute(r ApiGetMetadataEditorInfoRequest) (*MetadataEditorInfo, *http.Response, error)

	/*
	UpdateItem Updates an item.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param itemId The item id.
	@return ApiUpdateItemRequest
	*/
	UpdateItem(ctx context.Context, itemId string) ApiUpdateItemRequest

	// UpdateItemExecute executes the request
	UpdateItemExecute(r ApiUpdateItemRequest) (*http.Response, error)

	/*
	UpdateItemContentType Updates an item's content type.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param itemId The item id.
	@return ApiUpdateItemContentTypeRequest
	*/
	UpdateItemContentType(ctx context.Context, itemId string) ApiUpdateItemContentTypeRequest

	// UpdateItemContentTypeExecute executes the request
	UpdateItemContentTypeExecute(r ApiUpdateItemContentTypeRequest) (*http.Response, error)
}

// ItemUpdateAPIService ItemUpdateAPI service
type ItemUpdateAPIService service

type ApiGetMetadataEditorInfoRequest struct {
	ctx context.Context
	ApiService ItemUpdateAPI
	itemId string
}

func (r ApiGetMetadataEditorInfoRequest) Execute() (*MetadataEditorInfo, *http.Response, error) {
	return r.ApiService.GetMetadataEditorInfoExecute(r)
}

/*
GetMetadataEditorInfo Gets metadata editor info for an item.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId The item id.
 @return ApiGetMetadataEditorInfoRequest
*/
func (a *ItemUpdateAPIService) GetMetadataEditorInfo(ctx context.Context, itemId string) ApiGetMetadataEditorInfoRequest {
	return ApiGetMetadataEditorInfoRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
//  @return MetadataEditorInfo
func (a *ItemUpdateAPIService) GetMetadataEditorInfoExecute(r ApiGetMetadataEditorInfoRequest) (*MetadataEditorInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MetadataEditorInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ItemUpdateAPIService.GetMetadataEditorInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Items/{itemId}/MetadataEditor"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CustomAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type ApiUpdateItemRequest struct {
	ctx context.Context
	ApiService ItemUpdateAPI
	itemId string
	baseItemDto *BaseItemDto
}

// The new item properties.
func (r ApiUpdateItemRequest) BaseItemDto(baseItemDto BaseItemDto) ApiUpdateItemRequest {
	r.baseItemDto = &baseItemDto
	return r
}

func (r ApiUpdateItemRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateItemExecute(r)
}

/*
UpdateItem Updates an item.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId The item id.
 @return ApiUpdateItemRequest
*/
func (a *ItemUpdateAPIService) UpdateItem(ctx context.Context, itemId string) ApiUpdateItemRequest {
	return ApiUpdateItemRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
func (a *ItemUpdateAPIService) UpdateItemExecute(r ApiUpdateItemRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ItemUpdateAPIService.UpdateItem")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Items/{itemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.baseItemDto == nil {
		return nil, reportError("baseItemDto is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.baseItemDto
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CustomAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateItemContentTypeRequest struct {
	ctx context.Context
	ApiService ItemUpdateAPI
	itemId string
	contentType *string
}

// The content type of the item.
func (r ApiUpdateItemContentTypeRequest) ContentType(contentType string) ApiUpdateItemContentTypeRequest {
	r.contentType = &contentType
	return r
}

func (r ApiUpdateItemContentTypeRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateItemContentTypeExecute(r)
}

/*
UpdateItemContentType Updates an item's content type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId The item id.
 @return ApiUpdateItemContentTypeRequest
*/
func (a *ItemUpdateAPIService) UpdateItemContentType(ctx context.Context, itemId string) ApiUpdateItemContentTypeRequest {
	return ApiUpdateItemContentTypeRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
func (a *ItemUpdateAPIService) UpdateItemContentTypeExecute(r ApiUpdateItemContentTypeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ItemUpdateAPIService.UpdateItemContentType")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Items/{itemId}/ContentType"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "contentType", r.contentType, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["CustomAuthentication"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
