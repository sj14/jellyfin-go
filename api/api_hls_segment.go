/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.1
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


type HlsSegmentAPI interface {

	/*
	GetHlsAudioSegmentLegacyAac Gets the specified audio segment for an audio item.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param itemId The item id.
	@param segmentId The segment id.
	@return ApiGetHlsAudioSegmentLegacyAacRequest
	*/
	GetHlsAudioSegmentLegacyAac(ctx context.Context, itemId string, segmentId string) ApiGetHlsAudioSegmentLegacyAacRequest

	// GetHlsAudioSegmentLegacyAacExecute executes the request
	//  @return *os.File
	GetHlsAudioSegmentLegacyAacExecute(r ApiGetHlsAudioSegmentLegacyAacRequest) (*os.File, *http.Response, error)

	/*
	GetHlsAudioSegmentLegacyMp3 Gets the specified audio segment for an audio item.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param itemId The item id.
	@param segmentId The segment id.
	@return ApiGetHlsAudioSegmentLegacyMp3Request
	*/
	GetHlsAudioSegmentLegacyMp3(ctx context.Context, itemId string, segmentId string) ApiGetHlsAudioSegmentLegacyMp3Request

	// GetHlsAudioSegmentLegacyMp3Execute executes the request
	//  @return *os.File
	GetHlsAudioSegmentLegacyMp3Execute(r ApiGetHlsAudioSegmentLegacyMp3Request) (*os.File, *http.Response, error)

	/*
	GetHlsPlaylistLegacy Gets a hls video playlist.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param itemId The video id.
	@param playlistId The playlist id.
	@return ApiGetHlsPlaylistLegacyRequest
	*/
	GetHlsPlaylistLegacy(ctx context.Context, itemId string, playlistId string) ApiGetHlsPlaylistLegacyRequest

	// GetHlsPlaylistLegacyExecute executes the request
	//  @return *os.File
	GetHlsPlaylistLegacyExecute(r ApiGetHlsPlaylistLegacyRequest) (*os.File, *http.Response, error)

	/*
	GetHlsVideoSegmentLegacy Gets a hls video segment.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param itemId The item id.
	@param playlistId The playlist id.
	@param segmentId The segment id.
	@param segmentContainer The segment container.
	@return ApiGetHlsVideoSegmentLegacyRequest
	*/
	GetHlsVideoSegmentLegacy(ctx context.Context, itemId string, playlistId string, segmentId string, segmentContainer string) ApiGetHlsVideoSegmentLegacyRequest

	// GetHlsVideoSegmentLegacyExecute executes the request
	//  @return *os.File
	GetHlsVideoSegmentLegacyExecute(r ApiGetHlsVideoSegmentLegacyRequest) (*os.File, *http.Response, error)

	/*
	StopEncodingProcess Stops an active encoding.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiStopEncodingProcessRequest
	*/
	StopEncodingProcess(ctx context.Context) ApiStopEncodingProcessRequest

	// StopEncodingProcessExecute executes the request
	StopEncodingProcessExecute(r ApiStopEncodingProcessRequest) (*http.Response, error)
}

// HlsSegmentAPIService HlsSegmentAPI service
type HlsSegmentAPIService service

type ApiGetHlsAudioSegmentLegacyAacRequest struct {
	ctx context.Context
	ApiService HlsSegmentAPI
	itemId string
	segmentId string
}

func (r ApiGetHlsAudioSegmentLegacyAacRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.GetHlsAudioSegmentLegacyAacExecute(r)
}

/*
GetHlsAudioSegmentLegacyAac Gets the specified audio segment for an audio item.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId The item id.
 @param segmentId The segment id.
 @return ApiGetHlsAudioSegmentLegacyAacRequest
*/
func (a *HlsSegmentAPIService) GetHlsAudioSegmentLegacyAac(ctx context.Context, itemId string, segmentId string) ApiGetHlsAudioSegmentLegacyAacRequest {
	return ApiGetHlsAudioSegmentLegacyAacRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
		segmentId: segmentId,
	}
}

// Execute executes the request
//  @return *os.File
func (a *HlsSegmentAPIService) GetHlsAudioSegmentLegacyAacExecute(r ApiGetHlsAudioSegmentLegacyAacRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HlsSegmentAPIService.GetHlsAudioSegmentLegacyAac")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Audio/{itemId}/hls/{segmentId}/stream.aac"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"segmentId"+"}", url.PathEscape(parameterValueToString(r.segmentId, "segmentId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"audio/*"}

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

type ApiGetHlsAudioSegmentLegacyMp3Request struct {
	ctx context.Context
	ApiService HlsSegmentAPI
	itemId string
	segmentId string
}

func (r ApiGetHlsAudioSegmentLegacyMp3Request) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.GetHlsAudioSegmentLegacyMp3Execute(r)
}

/*
GetHlsAudioSegmentLegacyMp3 Gets the specified audio segment for an audio item.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId The item id.
 @param segmentId The segment id.
 @return ApiGetHlsAudioSegmentLegacyMp3Request
*/
func (a *HlsSegmentAPIService) GetHlsAudioSegmentLegacyMp3(ctx context.Context, itemId string, segmentId string) ApiGetHlsAudioSegmentLegacyMp3Request {
	return ApiGetHlsAudioSegmentLegacyMp3Request{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
		segmentId: segmentId,
	}
}

// Execute executes the request
//  @return *os.File
func (a *HlsSegmentAPIService) GetHlsAudioSegmentLegacyMp3Execute(r ApiGetHlsAudioSegmentLegacyMp3Request) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HlsSegmentAPIService.GetHlsAudioSegmentLegacyMp3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Audio/{itemId}/hls/{segmentId}/stream.mp3"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"segmentId"+"}", url.PathEscape(parameterValueToString(r.segmentId, "segmentId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"audio/*"}

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

type ApiGetHlsPlaylistLegacyRequest struct {
	ctx context.Context
	ApiService HlsSegmentAPI
	itemId string
	playlistId string
}

func (r ApiGetHlsPlaylistLegacyRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.GetHlsPlaylistLegacyExecute(r)
}

/*
GetHlsPlaylistLegacy Gets a hls video playlist.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId The video id.
 @param playlistId The playlist id.
 @return ApiGetHlsPlaylistLegacyRequest
*/
func (a *HlsSegmentAPIService) GetHlsPlaylistLegacy(ctx context.Context, itemId string, playlistId string) ApiGetHlsPlaylistLegacyRequest {
	return ApiGetHlsPlaylistLegacyRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
		playlistId: playlistId,
	}
}

// Execute executes the request
//  @return *os.File
func (a *HlsSegmentAPIService) GetHlsPlaylistLegacyExecute(r ApiGetHlsPlaylistLegacyRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HlsSegmentAPIService.GetHlsPlaylistLegacy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Videos/{itemId}/hls/{playlistId}/stream.m3u8"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"playlistId"+"}", url.PathEscape(parameterValueToString(r.playlistId, "playlistId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/x-mpegURL"}

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

type ApiGetHlsVideoSegmentLegacyRequest struct {
	ctx context.Context
	ApiService HlsSegmentAPI
	itemId string
	playlistId string
	segmentId string
	segmentContainer string
}

func (r ApiGetHlsVideoSegmentLegacyRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.GetHlsVideoSegmentLegacyExecute(r)
}

/*
GetHlsVideoSegmentLegacy Gets a hls video segment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param itemId The item id.
 @param playlistId The playlist id.
 @param segmentId The segment id.
 @param segmentContainer The segment container.
 @return ApiGetHlsVideoSegmentLegacyRequest
*/
func (a *HlsSegmentAPIService) GetHlsVideoSegmentLegacy(ctx context.Context, itemId string, playlistId string, segmentId string, segmentContainer string) ApiGetHlsVideoSegmentLegacyRequest {
	return ApiGetHlsVideoSegmentLegacyRequest{
		ApiService: a,
		ctx: ctx,
		itemId: itemId,
		playlistId: playlistId,
		segmentId: segmentId,
		segmentContainer: segmentContainer,
	}
}

// Execute executes the request
//  @return *os.File
func (a *HlsSegmentAPIService) GetHlsVideoSegmentLegacyExecute(r ApiGetHlsVideoSegmentLegacyRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HlsSegmentAPIService.GetHlsVideoSegmentLegacy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Videos/{itemId}/hls/{playlistId}/{segmentId}.{segmentContainer}"
	localVarPath = strings.Replace(localVarPath, "{"+"itemId"+"}", url.PathEscape(parameterValueToString(r.itemId, "itemId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"playlistId"+"}", url.PathEscape(parameterValueToString(r.playlistId, "playlistId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"segmentId"+"}", url.PathEscape(parameterValueToString(r.segmentId, "segmentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"segmentContainer"+"}", url.PathEscape(parameterValueToString(r.segmentContainer, "segmentContainer")), -1)

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
	localVarHTTPHeaderAccepts := []string{"video/*", "application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

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

type ApiStopEncodingProcessRequest struct {
	ctx context.Context
	ApiService HlsSegmentAPI
	deviceId *string
	playSessionId *string
}

// The device id of the client requesting. Used to stop encoding processes when needed.
func (r ApiStopEncodingProcessRequest) DeviceId(deviceId string) ApiStopEncodingProcessRequest {
	r.deviceId = &deviceId
	return r
}

// The play session id.
func (r ApiStopEncodingProcessRequest) PlaySessionId(playSessionId string) ApiStopEncodingProcessRequest {
	r.playSessionId = &playSessionId
	return r
}

func (r ApiStopEncodingProcessRequest) Execute() (*http.Response, error) {
	return r.ApiService.StopEncodingProcessExecute(r)
}

/*
StopEncodingProcess Stops an active encoding.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStopEncodingProcessRequest
*/
func (a *HlsSegmentAPIService) StopEncodingProcess(ctx context.Context) ApiStopEncodingProcessRequest {
	return ApiStopEncodingProcessRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *HlsSegmentAPIService) StopEncodingProcessExecute(r ApiStopEncodingProcessRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HlsSegmentAPIService.StopEncodingProcess")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Videos/ActiveEncodings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceId == nil {
		return nil, reportError("deviceId is required and must be specified")
	}
	if r.playSessionId == nil {
		return nil, reportError("playSessionId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "deviceId", r.deviceId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "playSessionId", r.playSessionId, "form", "")
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
