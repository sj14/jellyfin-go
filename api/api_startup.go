/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.0
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


type StartupAPI interface {

	/*
	CompleteWizard Completes the startup wizard.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCompleteWizardRequest
	*/
	CompleteWizard(ctx context.Context) ApiCompleteWizardRequest

	// CompleteWizardExecute executes the request
	CompleteWizardExecute(r ApiCompleteWizardRequest) (*http.Response, error)

	/*
	GetFirstUser Gets the first user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetFirstUserRequest
	*/
	GetFirstUser(ctx context.Context) ApiGetFirstUserRequest

	// GetFirstUserExecute executes the request
	//  @return StartupUserDto
	GetFirstUserExecute(r ApiGetFirstUserRequest) (*StartupUserDto, *http.Response, error)

	/*
	GetFirstUser2 Gets the first user.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetFirstUser2Request
	*/
	GetFirstUser2(ctx context.Context) ApiGetFirstUser2Request

	// GetFirstUser2Execute executes the request
	//  @return StartupUserDto
	GetFirstUser2Execute(r ApiGetFirstUser2Request) (*StartupUserDto, *http.Response, error)

	/*
	GetStartupConfiguration Gets the initial startup wizard configuration.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetStartupConfigurationRequest
	*/
	GetStartupConfiguration(ctx context.Context) ApiGetStartupConfigurationRequest

	// GetStartupConfigurationExecute executes the request
	//  @return StartupConfigurationDto
	GetStartupConfigurationExecute(r ApiGetStartupConfigurationRequest) (*StartupConfigurationDto, *http.Response, error)

	/*
	SetRemoteAccess Sets remote access and UPnP.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSetRemoteAccessRequest
	*/
	SetRemoteAccess(ctx context.Context) ApiSetRemoteAccessRequest

	// SetRemoteAccessExecute executes the request
	SetRemoteAccessExecute(r ApiSetRemoteAccessRequest) (*http.Response, error)

	/*
	UpdateInitialConfiguration Sets the initial startup wizard configuration.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdateInitialConfigurationRequest
	*/
	UpdateInitialConfiguration(ctx context.Context) ApiUpdateInitialConfigurationRequest

	// UpdateInitialConfigurationExecute executes the request
	UpdateInitialConfigurationExecute(r ApiUpdateInitialConfigurationRequest) (*http.Response, error)

	/*
	UpdateStartupUser Sets the user name and password.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdateStartupUserRequest
	*/
	UpdateStartupUser(ctx context.Context) ApiUpdateStartupUserRequest

	// UpdateStartupUserExecute executes the request
	UpdateStartupUserExecute(r ApiUpdateStartupUserRequest) (*http.Response, error)
}

// StartupAPIService StartupAPI service
type StartupAPIService service

type ApiCompleteWizardRequest struct {
	ctx context.Context
	ApiService StartupAPI
}

func (r ApiCompleteWizardRequest) Execute() (*http.Response, error) {
	return r.ApiService.CompleteWizardExecute(r)
}

/*
CompleteWizard Completes the startup wizard.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCompleteWizardRequest
*/
func (a *StartupAPIService) CompleteWizard(ctx context.Context) ApiCompleteWizardRequest {
	return ApiCompleteWizardRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *StartupAPIService) CompleteWizardExecute(r ApiCompleteWizardRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StartupAPIService.CompleteWizard")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Startup/Complete"

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetFirstUserRequest struct {
	ctx context.Context
	ApiService StartupAPI
}

func (r ApiGetFirstUserRequest) Execute() (*StartupUserDto, *http.Response, error) {
	return r.ApiService.GetFirstUserExecute(r)
}

/*
GetFirstUser Gets the first user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFirstUserRequest
*/
func (a *StartupAPIService) GetFirstUser(ctx context.Context) ApiGetFirstUserRequest {
	return ApiGetFirstUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StartupUserDto
func (a *StartupAPIService) GetFirstUserExecute(r ApiGetFirstUserRequest) (*StartupUserDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StartupUserDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StartupAPIService.GetFirstUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Startup/User"

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

type ApiGetFirstUser2Request struct {
	ctx context.Context
	ApiService StartupAPI
}

func (r ApiGetFirstUser2Request) Execute() (*StartupUserDto, *http.Response, error) {
	return r.ApiService.GetFirstUser2Execute(r)
}

/*
GetFirstUser2 Gets the first user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetFirstUser2Request
*/
func (a *StartupAPIService) GetFirstUser2(ctx context.Context) ApiGetFirstUser2Request {
	return ApiGetFirstUser2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StartupUserDto
func (a *StartupAPIService) GetFirstUser2Execute(r ApiGetFirstUser2Request) (*StartupUserDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StartupUserDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StartupAPIService.GetFirstUser2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Startup/FirstUser"

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

type ApiGetStartupConfigurationRequest struct {
	ctx context.Context
	ApiService StartupAPI
}

func (r ApiGetStartupConfigurationRequest) Execute() (*StartupConfigurationDto, *http.Response, error) {
	return r.ApiService.GetStartupConfigurationExecute(r)
}

/*
GetStartupConfiguration Gets the initial startup wizard configuration.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetStartupConfigurationRequest
*/
func (a *StartupAPIService) GetStartupConfiguration(ctx context.Context) ApiGetStartupConfigurationRequest {
	return ApiGetStartupConfigurationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StartupConfigurationDto
func (a *StartupAPIService) GetStartupConfigurationExecute(r ApiGetStartupConfigurationRequest) (*StartupConfigurationDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StartupConfigurationDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StartupAPIService.GetStartupConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Startup/Configuration"

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

type ApiSetRemoteAccessRequest struct {
	ctx context.Context
	ApiService StartupAPI
	startupRemoteAccessDto *StartupRemoteAccessDto
}

// The startup remote access dto.
func (r ApiSetRemoteAccessRequest) StartupRemoteAccessDto(startupRemoteAccessDto StartupRemoteAccessDto) ApiSetRemoteAccessRequest {
	r.startupRemoteAccessDto = &startupRemoteAccessDto
	return r
}

func (r ApiSetRemoteAccessRequest) Execute() (*http.Response, error) {
	return r.ApiService.SetRemoteAccessExecute(r)
}

/*
SetRemoteAccess Sets remote access and UPnP.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSetRemoteAccessRequest
*/
func (a *StartupAPIService) SetRemoteAccess(ctx context.Context) ApiSetRemoteAccessRequest {
	return ApiSetRemoteAccessRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *StartupAPIService) SetRemoteAccessExecute(r ApiSetRemoteAccessRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StartupAPIService.SetRemoteAccess")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Startup/RemoteAccess"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.startupRemoteAccessDto == nil {
		return nil, reportError("startupRemoteAccessDto is required and must be specified")
	}

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
	localVarPostBody = r.startupRemoteAccessDto
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

type ApiUpdateInitialConfigurationRequest struct {
	ctx context.Context
	ApiService StartupAPI
	startupConfigurationDto *StartupConfigurationDto
}

// The updated startup configuration.
func (r ApiUpdateInitialConfigurationRequest) StartupConfigurationDto(startupConfigurationDto StartupConfigurationDto) ApiUpdateInitialConfigurationRequest {
	r.startupConfigurationDto = &startupConfigurationDto
	return r
}

func (r ApiUpdateInitialConfigurationRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateInitialConfigurationExecute(r)
}

/*
UpdateInitialConfiguration Sets the initial startup wizard configuration.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateInitialConfigurationRequest
*/
func (a *StartupAPIService) UpdateInitialConfiguration(ctx context.Context) ApiUpdateInitialConfigurationRequest {
	return ApiUpdateInitialConfigurationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *StartupAPIService) UpdateInitialConfigurationExecute(r ApiUpdateInitialConfigurationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StartupAPIService.UpdateInitialConfiguration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Startup/Configuration"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.startupConfigurationDto == nil {
		return nil, reportError("startupConfigurationDto is required and must be specified")
	}

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
	localVarPostBody = r.startupConfigurationDto
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

type ApiUpdateStartupUserRequest struct {
	ctx context.Context
	ApiService StartupAPI
	startupUserDto *StartupUserDto
}

// The DTO containing username and password.
func (r ApiUpdateStartupUserRequest) StartupUserDto(startupUserDto StartupUserDto) ApiUpdateStartupUserRequest {
	r.startupUserDto = &startupUserDto
	return r
}

func (r ApiUpdateStartupUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateStartupUserExecute(r)
}

/*
UpdateStartupUser Sets the user name and password.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateStartupUserRequest
*/
func (a *StartupAPIService) UpdateStartupUser(ctx context.Context) ApiUpdateStartupUserRequest {
	return ApiUpdateStartupUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *StartupAPIService) UpdateStartupUserExecute(r ApiUpdateStartupUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StartupAPIService.UpdateStartupUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Startup/User"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.startupUserDto
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
