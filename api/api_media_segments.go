/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.5
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
	"reflect"
)


type MediaSegmentsAPI interface {

	/*
	GetItemSegments Gets all media segments based on an itemId.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param itemId The ItemId.
	@return ApiGetItemSegmentsRequest
	*/
	GetItemSegments(ctx context.Context, itemId string) ApiGetItemSegmentsRequest

	// GetItemSegmentsExecute executes the request
	//  @return MediaSegmentDtoQueryResult
	GetItemSegmentsExecute(r ApiGetItemSegmentsRequest) (*MediaSegmentDtoQueryResult, *http.Response, error)
}

// MediaSegmentsAPIService MediaSegmentsAPI service
type MediaSegmentsAPIService service

type ApiGetItemSegmentsRequest struct {
	ctx context.Context
	ApiService MediaSegmentsAPI
	itemId string
	includeSegmentTypes *[]MediaSegmentType
}

// Optional filter of requested segment types.
func (r ApiGetItemSegmentsRequest) IncludeSegmentTypes(includeSegmentTypes []MediaSegmentType) ApiGetItemSegmentsRequest {
	r.includeSegmentTypes = &includeSegmentTypes
	return r
}

func (r ApiGetItemSegmentsRequest) Execute() (*MediaSegmentDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetItemSegmentsExecute(r)
}

/*
GetItemSegments Gets all media segments based on an itemId.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId The ItemId.
 @return ApiGetItemSegmentsRequest
*/
func (a *MediaSegmentsAPIService) GetItemSegments(ctx context.Context, itemId string) ApiGetItemSegmentsRequest {
	return ApiGetItemSegmentsRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
	}
}

// Execute executes the request
//  @return MediaSegmentDtoQueryResult
func (a *MediaSegmentsAPIService) GetItemSegmentsExecute(r ApiGetItemSegmentsRequest) (*MediaSegmentDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MediaSegmentDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaSegmentsAPIService.GetItemSegments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/MediaSegments/{itemId}"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeSegmentTypes != nil {
		t := *r.includeSegmentTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "includeSegmentTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "includeSegmentTypes", t, "form", "multi")
		}
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
