/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.11
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


type ChannelsAPI interface {

	/*
	GetAllChannelFeatures Get all channel features.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAllChannelFeaturesRequest
	*/
	GetAllChannelFeatures(ctx context.Context) ApiGetAllChannelFeaturesRequest

	// GetAllChannelFeaturesExecute executes the request
	//  @return []ChannelFeatures
	GetAllChannelFeaturesExecute(r ApiGetAllChannelFeaturesRequest) ([]ChannelFeatures, *http.Response, error)

	/*
	GetChannelFeatures Get channel features.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId Channel id.
	@return ApiGetChannelFeaturesRequest
	*/
	GetChannelFeatures(ctx context.Context, channelId string) ApiGetChannelFeaturesRequest

	// GetChannelFeaturesExecute executes the request
	//  @return ChannelFeatures
	GetChannelFeaturesExecute(r ApiGetChannelFeaturesRequest) (*ChannelFeatures, *http.Response, error)

	/*
	GetChannelItems Get channel items.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channelId Channel Id.
	@return ApiGetChannelItemsRequest
	*/
	GetChannelItems(ctx context.Context, channelId string) ApiGetChannelItemsRequest

	// GetChannelItemsExecute executes the request
	//  @return BaseItemDtoQueryResult
	GetChannelItemsExecute(r ApiGetChannelItemsRequest) (*BaseItemDtoQueryResult, *http.Response, error)

	/*
	GetChannels Gets available channels.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetChannelsRequest
	*/
	GetChannels(ctx context.Context) ApiGetChannelsRequest

	// GetChannelsExecute executes the request
	//  @return BaseItemDtoQueryResult
	GetChannelsExecute(r ApiGetChannelsRequest) (*BaseItemDtoQueryResult, *http.Response, error)

	/*
	GetLatestChannelItems Gets latest channel items.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetLatestChannelItemsRequest
	*/
	GetLatestChannelItems(ctx context.Context) ApiGetLatestChannelItemsRequest

	// GetLatestChannelItemsExecute executes the request
	//  @return BaseItemDtoQueryResult
	GetLatestChannelItemsExecute(r ApiGetLatestChannelItemsRequest) (*BaseItemDtoQueryResult, *http.Response, error)
}

// ChannelsAPIService ChannelsAPI service
type ChannelsAPIService service

type ApiGetAllChannelFeaturesRequest struct {
	ctx context.Context
	ApiService ChannelsAPI
}

func (r ApiGetAllChannelFeaturesRequest) Execute() ([]ChannelFeatures, *http.Response, error) {
	return r.ApiService.GetAllChannelFeaturesExecute(r)
}

/*
GetAllChannelFeatures Get all channel features.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAllChannelFeaturesRequest
*/
func (a *ChannelsAPIService) GetAllChannelFeatures(ctx context.Context) ApiGetAllChannelFeaturesRequest {
	return ApiGetAllChannelFeaturesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ChannelFeatures
func (a *ChannelsAPIService) GetAllChannelFeaturesExecute(r ApiGetAllChannelFeaturesRequest) ([]ChannelFeatures, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ChannelFeatures
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsAPIService.GetAllChannelFeatures")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Channels/Features"

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

type ApiGetChannelFeaturesRequest struct {
	ctx context.Context
	ApiService ChannelsAPI
	channelId string
}

func (r ApiGetChannelFeaturesRequest) Execute() (*ChannelFeatures, *http.Response, error) {
	return r.ApiService.GetChannelFeaturesExecute(r)
}

/*
GetChannelFeatures Get channel features.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId Channel id.
 @return ApiGetChannelFeaturesRequest
*/
func (a *ChannelsAPIService) GetChannelFeatures(ctx context.Context, channelId string) ApiGetChannelFeaturesRequest {
	return ApiGetChannelFeaturesRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
	}
}

// Execute executes the request
//  @return ChannelFeatures
func (a *ChannelsAPIService) GetChannelFeaturesExecute(r ApiGetChannelFeaturesRequest) (*ChannelFeatures, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChannelFeatures
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsAPIService.GetChannelFeatures")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Channels/{channelId}/Features"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

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

type ApiGetChannelItemsRequest struct {
	ctx context.Context
	ApiService ChannelsAPI
	channelId string
	folderId *string
	userId *string
	startIndex *int32
	limit *int32
	sortOrder *[]SortOrder
	filters *[]ItemFilter
	sortBy *[]ItemSortBy
	fields *[]ItemFields
}

// Optional. Folder Id.
func (r ApiGetChannelItemsRequest) FolderId(folderId string) ApiGetChannelItemsRequest {
	r.folderId = &folderId
	return r
}

// Optional. User Id.
func (r ApiGetChannelItemsRequest) UserId(userId string) ApiGetChannelItemsRequest {
	r.userId = &userId
	return r
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r ApiGetChannelItemsRequest) StartIndex(startIndex int32) ApiGetChannelItemsRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return.
func (r ApiGetChannelItemsRequest) Limit(limit int32) ApiGetChannelItemsRequest {
	r.limit = &limit
	return r
}

// Optional. Sort Order - Ascending,Descending.
func (r ApiGetChannelItemsRequest) SortOrder(sortOrder []SortOrder) ApiGetChannelItemsRequest {
	r.sortOrder = &sortOrder
	return r
}

// Optional. Specify additional filters to apply.
func (r ApiGetChannelItemsRequest) Filters(filters []ItemFilter) ApiGetChannelItemsRequest {
	r.filters = &filters
	return r
}

// Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime.
func (r ApiGetChannelItemsRequest) SortBy(sortBy []ItemSortBy) ApiGetChannelItemsRequest {
	r.sortBy = &sortBy
	return r
}

// Optional. Specify additional fields of information to return in the output.
func (r ApiGetChannelItemsRequest) Fields(fields []ItemFields) ApiGetChannelItemsRequest {
	r.fields = &fields
	return r
}

func (r ApiGetChannelItemsRequest) Execute() (*BaseItemDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetChannelItemsExecute(r)
}

/*
GetChannelItems Get channel items.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channelId Channel Id.
 @return ApiGetChannelItemsRequest
*/
func (a *ChannelsAPIService) GetChannelItems(ctx context.Context, channelId string) ApiGetChannelItemsRequest {
	return ApiGetChannelItemsRequest{
		ApiService: a,
		ctx: ctx,
		channelId: channelId,
	}
}

// Execute executes the request
//  @return BaseItemDtoQueryResult
func (a *ChannelsAPIService) GetChannelItemsExecute(r ApiGetChannelItemsRequest) (*BaseItemDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsAPIService.GetChannelItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Channels/{channelId}/Items"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", url.PathEscape(parameterValueToString(r.channelId, "channelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.folderId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "folderId", r.folderId, "form", "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.sortOrder != nil {
		t := *r.sortOrder
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sortOrder", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sortOrder", t, "form", "multi")
		}
	}
	if r.filters != nil {
		t := *r.filters
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filters", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filters", t, "form", "multi")
		}
	}
	if r.sortBy != nil {
		t := *r.sortBy
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", t, "form", "multi")
		}
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "form", "multi")
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

type ApiGetChannelsRequest struct {
	ctx context.Context
	ApiService ChannelsAPI
	userId *string
	startIndex *int32
	limit *int32
	supportsLatestItems *bool
	supportsMediaDeletion *bool
	isFavorite *bool
}

// User Id to filter by. Use System.Guid.Empty to not filter by user.
func (r ApiGetChannelsRequest) UserId(userId string) ApiGetChannelsRequest {
	r.userId = &userId
	return r
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r ApiGetChannelsRequest) StartIndex(startIndex int32) ApiGetChannelsRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return.
func (r ApiGetChannelsRequest) Limit(limit int32) ApiGetChannelsRequest {
	r.limit = &limit
	return r
}

// Optional. Filter by channels that support getting latest items.
func (r ApiGetChannelsRequest) SupportsLatestItems(supportsLatestItems bool) ApiGetChannelsRequest {
	r.supportsLatestItems = &supportsLatestItems
	return r
}

// Optional. Filter by channels that support media deletion.
func (r ApiGetChannelsRequest) SupportsMediaDeletion(supportsMediaDeletion bool) ApiGetChannelsRequest {
	r.supportsMediaDeletion = &supportsMediaDeletion
	return r
}

// Optional. Filter by channels that are favorite.
func (r ApiGetChannelsRequest) IsFavorite(isFavorite bool) ApiGetChannelsRequest {
	r.isFavorite = &isFavorite
	return r
}

func (r ApiGetChannelsRequest) Execute() (*BaseItemDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetChannelsExecute(r)
}

/*
GetChannels Gets available channels.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetChannelsRequest
*/
func (a *ChannelsAPIService) GetChannels(ctx context.Context) ApiGetChannelsRequest {
	return ApiGetChannelsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BaseItemDtoQueryResult
func (a *ChannelsAPIService) GetChannelsExecute(r ApiGetChannelsRequest) (*BaseItemDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsAPIService.GetChannels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Channels"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.supportsLatestItems != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supportsLatestItems", r.supportsLatestItems, "form", "")
	}
	if r.supportsMediaDeletion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supportsMediaDeletion", r.supportsMediaDeletion, "form", "")
	}
	if r.isFavorite != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isFavorite", r.isFavorite, "form", "")
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

type ApiGetLatestChannelItemsRequest struct {
	ctx context.Context
	ApiService ChannelsAPI
	userId *string
	startIndex *int32
	limit *int32
	filters *[]ItemFilter
	fields *[]ItemFields
	channelIds *[]string
}

// Optional. User Id.
func (r ApiGetLatestChannelItemsRequest) UserId(userId string) ApiGetLatestChannelItemsRequest {
	r.userId = &userId
	return r
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r ApiGetLatestChannelItemsRequest) StartIndex(startIndex int32) ApiGetLatestChannelItemsRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return.
func (r ApiGetLatestChannelItemsRequest) Limit(limit int32) ApiGetLatestChannelItemsRequest {
	r.limit = &limit
	return r
}

// Optional. Specify additional filters to apply.
func (r ApiGetLatestChannelItemsRequest) Filters(filters []ItemFilter) ApiGetLatestChannelItemsRequest {
	r.filters = &filters
	return r
}

// Optional. Specify additional fields of information to return in the output.
func (r ApiGetLatestChannelItemsRequest) Fields(fields []ItemFields) ApiGetLatestChannelItemsRequest {
	r.fields = &fields
	return r
}

// Optional. Specify one or more channel id&#39;s, comma delimited.
func (r ApiGetLatestChannelItemsRequest) ChannelIds(channelIds []string) ApiGetLatestChannelItemsRequest {
	r.channelIds = &channelIds
	return r
}

func (r ApiGetLatestChannelItemsRequest) Execute() (*BaseItemDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetLatestChannelItemsExecute(r)
}

/*
GetLatestChannelItems Gets latest channel items.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetLatestChannelItemsRequest
*/
func (a *ChannelsAPIService) GetLatestChannelItems(ctx context.Context) ApiGetLatestChannelItemsRequest {
	return ApiGetLatestChannelItemsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BaseItemDtoQueryResult
func (a *ChannelsAPIService) GetLatestChannelItemsExecute(r ApiGetLatestChannelItemsRequest) (*BaseItemDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChannelsAPIService.GetLatestChannelItems")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Channels/Items/Latest"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.filters != nil {
		t := *r.filters
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "filters", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "filters", t, "form", "multi")
		}
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "form", "multi")
		}
	}
	if r.channelIds != nil {
		t := *r.channelIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "channelIds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "channelIds", t, "form", "multi")
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
