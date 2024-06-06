/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)


type ActivityLogAPI interface {

	/*
	GetLogEntries Gets activity log entries.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetLogEntriesRequest
	*/
	GetLogEntries(ctx context.Context) ApiGetLogEntriesRequest

	// GetLogEntriesExecute executes the request
	//  @return ActivityLogEntryQueryResult
	GetLogEntriesExecute(r ApiGetLogEntriesRequest) (*ActivityLogEntryQueryResult, *http.Response, error)
}

// ActivityLogAPIService ActivityLogAPI service
type ActivityLogAPIService service

type ApiGetLogEntriesRequest struct {
	ctx context.Context
	ApiService ActivityLogAPI
	startIndex *int32
	limit *int32
	minDate *time.Time
	hasUserId *bool
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r ApiGetLogEntriesRequest) StartIndex(startIndex int32) ApiGetLogEntriesRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return.
func (r ApiGetLogEntriesRequest) Limit(limit int32) ApiGetLogEntriesRequest {
	r.limit = &limit
	return r
}

// Optional. The minimum date. Format &#x3D; ISO.
func (r ApiGetLogEntriesRequest) MinDate(minDate time.Time) ApiGetLogEntriesRequest {
	r.minDate = &minDate
	return r
}

// Optional. Filter log entries if it has user id, or not.
func (r ApiGetLogEntriesRequest) HasUserId(hasUserId bool) ApiGetLogEntriesRequest {
	r.hasUserId = &hasUserId
	return r
}

func (r ApiGetLogEntriesRequest) Execute() (*ActivityLogEntryQueryResult, *http.Response, error) {
	return r.ApiService.GetLogEntriesExecute(r)
}

/*
GetLogEntries Gets activity log entries.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetLogEntriesRequest
*/
func (a *ActivityLogAPIService) GetLogEntries(ctx context.Context) ApiGetLogEntriesRequest {
	return ApiGetLogEntriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ActivityLogEntryQueryResult
func (a *ActivityLogAPIService) GetLogEntriesExecute(r ApiGetLogEntriesRequest) (*ActivityLogEntryQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ActivityLogEntryQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ActivityLogAPIService.GetLogEntries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/System/ActivityLog/Entries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.minDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "minDate", r.minDate, "")
	}
	if r.hasUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hasUserId", r.hasUserId, "")
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
