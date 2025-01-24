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
)


type DisplayPreferencesAPI interface {

	/*
	GetDisplayPreferences Get Display Preferences.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param displayPreferencesId Display preferences id.
	@return ApiGetDisplayPreferencesRequest
	*/
	GetDisplayPreferences(ctx context.Context, displayPreferencesId string) ApiGetDisplayPreferencesRequest

	// GetDisplayPreferencesExecute executes the request
	//  @return DisplayPreferencesDto
	GetDisplayPreferencesExecute(r ApiGetDisplayPreferencesRequest) (*DisplayPreferencesDto, *http.Response, error)

	/*
	UpdateDisplayPreferences Update Display Preferences.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param displayPreferencesId Display preferences id.
	@return ApiUpdateDisplayPreferencesRequest
	*/
	UpdateDisplayPreferences(ctx context.Context, displayPreferencesId string) ApiUpdateDisplayPreferencesRequest

	// UpdateDisplayPreferencesExecute executes the request
	UpdateDisplayPreferencesExecute(r ApiUpdateDisplayPreferencesRequest) (*http.Response, error)
}

// DisplayPreferencesAPIService DisplayPreferencesAPI service
type DisplayPreferencesAPIService service

type ApiGetDisplayPreferencesRequest struct {
	ctx context.Context
	ApiService DisplayPreferencesAPI
	displayPreferencesId string
	client *string
	userId *string
}

// Client.
func (r ApiGetDisplayPreferencesRequest) Client(client string) ApiGetDisplayPreferencesRequest {
	r.client = &client
	return r
}

// User id.
func (r ApiGetDisplayPreferencesRequest) UserId(userId string) ApiGetDisplayPreferencesRequest {
	r.userId = &userId
	return r
}

func (r ApiGetDisplayPreferencesRequest) Execute() (*DisplayPreferencesDto, *http.Response, error) {
	return r.ApiService.GetDisplayPreferencesExecute(r)
}

/*
GetDisplayPreferences Get Display Preferences.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param displayPreferencesId Display preferences id.
 @return ApiGetDisplayPreferencesRequest
*/
func (a *DisplayPreferencesAPIService) GetDisplayPreferences(ctx context.Context, displayPreferencesId string) ApiGetDisplayPreferencesRequest {
	return ApiGetDisplayPreferencesRequest{
		ApiService: a,
		ctx: ctx,
		displayPreferencesId: displayPreferencesId,
	}
}

// Execute executes the request
//  @return DisplayPreferencesDto
func (a *DisplayPreferencesAPIService) GetDisplayPreferencesExecute(r ApiGetDisplayPreferencesRequest) (*DisplayPreferencesDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DisplayPreferencesDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DisplayPreferencesAPIService.GetDisplayPreferences")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/DisplayPreferences/{displayPreferencesId}"
	localVarPath = strings.Replace(localVarPath, "{"+"displayPreferencesId"+"}", url.PathEscape(parameterValueToString(r.displayPreferencesId, "displayPreferencesId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.client == nil {
		return localVarReturnValue, nil, reportError("client is required and must be specified")
	}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "client", r.client, "form", "")
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

type ApiUpdateDisplayPreferencesRequest struct {
	ctx context.Context
	ApiService DisplayPreferencesAPI
	displayPreferencesId string
	client *string
	displayPreferencesDto *DisplayPreferencesDto
	userId *string
}

// Client.
func (r ApiUpdateDisplayPreferencesRequest) Client(client string) ApiUpdateDisplayPreferencesRequest {
	r.client = &client
	return r
}

// New Display Preferences object.
func (r ApiUpdateDisplayPreferencesRequest) DisplayPreferencesDto(displayPreferencesDto DisplayPreferencesDto) ApiUpdateDisplayPreferencesRequest {
	r.displayPreferencesDto = &displayPreferencesDto
	return r
}

// User Id.
func (r ApiUpdateDisplayPreferencesRequest) UserId(userId string) ApiUpdateDisplayPreferencesRequest {
	r.userId = &userId
	return r
}

func (r ApiUpdateDisplayPreferencesRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateDisplayPreferencesExecute(r)
}

/*
UpdateDisplayPreferences Update Display Preferences.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param displayPreferencesId Display preferences id.
 @return ApiUpdateDisplayPreferencesRequest
*/
func (a *DisplayPreferencesAPIService) UpdateDisplayPreferences(ctx context.Context, displayPreferencesId string) ApiUpdateDisplayPreferencesRequest {
	return ApiUpdateDisplayPreferencesRequest{
		ApiService: a,
		ctx: ctx,
		displayPreferencesId: displayPreferencesId,
	}
}

// Execute executes the request
func (a *DisplayPreferencesAPIService) UpdateDisplayPreferencesExecute(r ApiUpdateDisplayPreferencesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DisplayPreferencesAPIService.UpdateDisplayPreferences")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/DisplayPreferences/{displayPreferencesId}"
	localVarPath = strings.Replace(localVarPath, "{"+"displayPreferencesId"+"}", url.PathEscape(parameterValueToString(r.displayPreferencesId, "displayPreferencesId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.client == nil {
		return nil, reportError("client is required and must be specified")
	}
	if r.displayPreferencesDto == nil {
		return nil, reportError("displayPreferencesDto is required and must be specified")
	}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "client", r.client, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "text/json", "application/*+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.displayPreferencesDto
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
