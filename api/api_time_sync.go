/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type TimeSyncAPI interface {

	/*
	GetUtcTime Gets the current UTC time.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetUtcTimeRequest
	*/
	GetUtcTime(ctx context.Context) ApiGetUtcTimeRequest

	// GetUtcTimeExecute executes the request
	//  @return UtcTimeResponse
	GetUtcTimeExecute(r ApiGetUtcTimeRequest) (*UtcTimeResponse, *http.Response, error)
}

// TimeSyncAPIService TimeSyncAPI service
type TimeSyncAPIService service

type ApiGetUtcTimeRequest struct {
	ctx context.Context
	ApiService TimeSyncAPI
}

func (r ApiGetUtcTimeRequest) Execute() (*UtcTimeResponse, *http.Response, error) {
	return r.ApiService.GetUtcTimeExecute(r)
}

/*
GetUtcTime Gets the current UTC time.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUtcTimeRequest
*/
func (a *TimeSyncAPIService) GetUtcTime(ctx context.Context) ApiGetUtcTimeRequest {
	return ApiGetUtcTimeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UtcTimeResponse
func (a *TimeSyncAPIService) GetUtcTimeExecute(r ApiGetUtcTimeRequest) (*UtcTimeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UtcTimeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimeSyncAPIService.GetUtcTime")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/GetUtcTime"

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
