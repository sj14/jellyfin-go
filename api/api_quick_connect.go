/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.0
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


type QuickConnectAPI interface {

	/*
	AuthorizeQuickConnect Authorizes a pending quick connect request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthorizeQuickConnectRequest
	*/
	AuthorizeQuickConnect(ctx context.Context) ApiAuthorizeQuickConnectRequest

	// AuthorizeQuickConnectExecute executes the request
	//  @return bool
	AuthorizeQuickConnectExecute(r ApiAuthorizeQuickConnectRequest) (bool, *http.Response, error)

	/*
	GetQuickConnectEnabled Gets the current quick connect state.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetQuickConnectEnabledRequest
	*/
	GetQuickConnectEnabled(ctx context.Context) ApiGetQuickConnectEnabledRequest

	// GetQuickConnectEnabledExecute executes the request
	//  @return bool
	GetQuickConnectEnabledExecute(r ApiGetQuickConnectEnabledRequest) (bool, *http.Response, error)

	/*
	GetQuickConnectState Attempts to retrieve authentication information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetQuickConnectStateRequest
	*/
	GetQuickConnectState(ctx context.Context) ApiGetQuickConnectStateRequest

	// GetQuickConnectStateExecute executes the request
	//  @return QuickConnectResult
	GetQuickConnectStateExecute(r ApiGetQuickConnectStateRequest) (*QuickConnectResult, *http.Response, error)

	/*
	InitiateQuickConnect Initiate a new quick connect request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiInitiateQuickConnectRequest
	*/
	InitiateQuickConnect(ctx context.Context) ApiInitiateQuickConnectRequest

	// InitiateQuickConnectExecute executes the request
	//  @return QuickConnectResult
	InitiateQuickConnectExecute(r ApiInitiateQuickConnectRequest) (*QuickConnectResult, *http.Response, error)
}

// QuickConnectAPIService QuickConnectAPI service
type QuickConnectAPIService service

type ApiAuthorizeQuickConnectRequest struct {
	ctx context.Context
	ApiService QuickConnectAPI
	code *string
	userId *string
}

// Quick connect code to authorize.
func (r ApiAuthorizeQuickConnectRequest) Code(code string) ApiAuthorizeQuickConnectRequest {
	r.code = &code
	return r
}

// The user the authorize. Access to the requested user is required.
func (r ApiAuthorizeQuickConnectRequest) UserId(userId string) ApiAuthorizeQuickConnectRequest {
	r.userId = &userId
	return r
}

func (r ApiAuthorizeQuickConnectRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.AuthorizeQuickConnectExecute(r)
}

/*
AuthorizeQuickConnect Authorizes a pending quick connect request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthorizeQuickConnectRequest
*/
func (a *QuickConnectAPIService) AuthorizeQuickConnect(ctx context.Context) ApiAuthorizeQuickConnectRequest {
	return ApiAuthorizeQuickConnectRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return bool
func (a *QuickConnectAPIService) AuthorizeQuickConnectExecute(r ApiAuthorizeQuickConnectRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuickConnectAPIService.AuthorizeQuickConnect")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/QuickConnect/Authorize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.code == nil {
		return localVarReturnValue, nil, reportError("code is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "code", r.code, "form", "")
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
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
		if localVarHTTPResponse.StatusCode == 403 {
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

type ApiGetQuickConnectEnabledRequest struct {
	ctx context.Context
	ApiService QuickConnectAPI
}

func (r ApiGetQuickConnectEnabledRequest) Execute() (bool, *http.Response, error) {
	return r.ApiService.GetQuickConnectEnabledExecute(r)
}

/*
GetQuickConnectEnabled Gets the current quick connect state.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetQuickConnectEnabledRequest
*/
func (a *QuickConnectAPIService) GetQuickConnectEnabled(ctx context.Context) ApiGetQuickConnectEnabledRequest {
	return ApiGetQuickConnectEnabledRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return bool
func (a *QuickConnectAPIService) GetQuickConnectEnabledExecute(r ApiGetQuickConnectEnabledRequest) (bool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  bool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuickConnectAPIService.GetQuickConnectEnabled")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/QuickConnect/Enabled"

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

type ApiGetQuickConnectStateRequest struct {
	ctx context.Context
	ApiService QuickConnectAPI
	secret *string
}

// Secret previously returned from the Initiate endpoint.
func (r ApiGetQuickConnectStateRequest) Secret(secret string) ApiGetQuickConnectStateRequest {
	r.secret = &secret
	return r
}

func (r ApiGetQuickConnectStateRequest) Execute() (*QuickConnectResult, *http.Response, error) {
	return r.ApiService.GetQuickConnectStateExecute(r)
}

/*
GetQuickConnectState Attempts to retrieve authentication information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetQuickConnectStateRequest
*/
func (a *QuickConnectAPIService) GetQuickConnectState(ctx context.Context) ApiGetQuickConnectStateRequest {
	return ApiGetQuickConnectStateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QuickConnectResult
func (a *QuickConnectAPIService) GetQuickConnectStateExecute(r ApiGetQuickConnectStateRequest) (*QuickConnectResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuickConnectResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuickConnectAPIService.GetQuickConnectState")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/QuickConnect/Connect"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.secret == nil {
		return localVarReturnValue, nil, reportError("secret is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "secret", r.secret, "form", "")
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

type ApiInitiateQuickConnectRequest struct {
	ctx context.Context
	ApiService QuickConnectAPI
}

func (r ApiInitiateQuickConnectRequest) Execute() (*QuickConnectResult, *http.Response, error) {
	return r.ApiService.InitiateQuickConnectExecute(r)
}

/*
InitiateQuickConnect Initiate a new quick connect request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInitiateQuickConnectRequest
*/
func (a *QuickConnectAPIService) InitiateQuickConnect(ctx context.Context) ApiInitiateQuickConnectRequest {
	return ApiInitiateQuickConnectRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QuickConnectResult
func (a *QuickConnectAPIService) InitiateQuickConnectExecute(r ApiInitiateQuickConnectRequest) (*QuickConnectResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QuickConnectResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QuickConnectAPIService.InitiateQuickConnect")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/QuickConnect/Initiate"

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
