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
	"reflect"
)


type FilterAPI interface {

	/*
	GetQueryFilters Gets query filters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetQueryFiltersRequest
	*/
	GetQueryFilters(ctx context.Context) ApiGetQueryFiltersRequest

	// GetQueryFiltersExecute executes the request
	//  @return QueryFilters
	GetQueryFiltersExecute(r ApiGetQueryFiltersRequest) (*QueryFilters, *http.Response, error)

	/*
	GetQueryFiltersLegacy Gets legacy query filters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetQueryFiltersLegacyRequest
	*/
	GetQueryFiltersLegacy(ctx context.Context) ApiGetQueryFiltersLegacyRequest

	// GetQueryFiltersLegacyExecute executes the request
	//  @return QueryFiltersLegacy
	GetQueryFiltersLegacyExecute(r ApiGetQueryFiltersLegacyRequest) (*QueryFiltersLegacy, *http.Response, error)
}

// FilterAPIService FilterAPI service
type FilterAPIService service

type ApiGetQueryFiltersRequest struct {
	ctx context.Context
	ApiService FilterAPI
	userId *string
	parentId *string
	includeItemTypes *[]BaseItemKind
	isAiring *bool
	isMovie *bool
	isSports *bool
	isKids *bool
	isNews *bool
	isSeries *bool
	recursive *bool
}

// Optional. User id.
func (r ApiGetQueryFiltersRequest) UserId(userId string) ApiGetQueryFiltersRequest {
	r.userId = &userId
	return r
}

// Optional. Specify this to localize the search to a specific item or folder. Omit to use the root.
func (r ApiGetQueryFiltersRequest) ParentId(parentId string) ApiGetQueryFiltersRequest {
	r.parentId = &parentId
	return r
}

// Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited.
func (r ApiGetQueryFiltersRequest) IncludeItemTypes(includeItemTypes []BaseItemKind) ApiGetQueryFiltersRequest {
	r.includeItemTypes = &includeItemTypes
	return r
}

// Optional. Is item airing.
func (r ApiGetQueryFiltersRequest) IsAiring(isAiring bool) ApiGetQueryFiltersRequest {
	r.isAiring = &isAiring
	return r
}

// Optional. Is item movie.
func (r ApiGetQueryFiltersRequest) IsMovie(isMovie bool) ApiGetQueryFiltersRequest {
	r.isMovie = &isMovie
	return r
}

// Optional. Is item sports.
func (r ApiGetQueryFiltersRequest) IsSports(isSports bool) ApiGetQueryFiltersRequest {
	r.isSports = &isSports
	return r
}

// Optional. Is item kids.
func (r ApiGetQueryFiltersRequest) IsKids(isKids bool) ApiGetQueryFiltersRequest {
	r.isKids = &isKids
	return r
}

// Optional. Is item news.
func (r ApiGetQueryFiltersRequest) IsNews(isNews bool) ApiGetQueryFiltersRequest {
	r.isNews = &isNews
	return r
}

// Optional. Is item series.
func (r ApiGetQueryFiltersRequest) IsSeries(isSeries bool) ApiGetQueryFiltersRequest {
	r.isSeries = &isSeries
	return r
}

// Optional. Search recursive.
func (r ApiGetQueryFiltersRequest) Recursive(recursive bool) ApiGetQueryFiltersRequest {
	r.recursive = &recursive
	return r
}

func (r ApiGetQueryFiltersRequest) Execute() (*QueryFilters, *http.Response, error) {
	return r.ApiService.GetQueryFiltersExecute(r)
}

/*
GetQueryFilters Gets query filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetQueryFiltersRequest
*/
func (a *FilterAPIService) GetQueryFilters(ctx context.Context) ApiGetQueryFiltersRequest {
	return ApiGetQueryFiltersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QueryFilters
func (a *FilterAPIService) GetQueryFiltersExecute(r ApiGetQueryFiltersRequest) (*QueryFilters, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QueryFilters
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilterAPIService.GetQueryFilters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Items/Filters2"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.parentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parentId", r.parentId, "form", "")
	}
	if r.includeItemTypes != nil {
		t := *r.includeItemTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "includeItemTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "includeItemTypes", t, "form", "multi")
		}
	}
	if r.isAiring != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isAiring", r.isAiring, "form", "")
	}
	if r.isMovie != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isMovie", r.isMovie, "form", "")
	}
	if r.isSports != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isSports", r.isSports, "form", "")
	}
	if r.isKids != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isKids", r.isKids, "form", "")
	}
	if r.isNews != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isNews", r.isNews, "form", "")
	}
	if r.isSeries != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isSeries", r.isSeries, "form", "")
	}
	if r.recursive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recursive", r.recursive, "form", "")
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

type ApiGetQueryFiltersLegacyRequest struct {
	ctx context.Context
	ApiService FilterAPI
	userId *string
	parentId *string
	includeItemTypes *[]BaseItemKind
	mediaTypes *[]MediaType
}

// Optional. User id.
func (r ApiGetQueryFiltersLegacyRequest) UserId(userId string) ApiGetQueryFiltersLegacyRequest {
	r.userId = &userId
	return r
}

// Optional. Parent id.
func (r ApiGetQueryFiltersLegacyRequest) ParentId(parentId string) ApiGetQueryFiltersLegacyRequest {
	r.parentId = &parentId
	return r
}

// Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited.
func (r ApiGetQueryFiltersLegacyRequest) IncludeItemTypes(includeItemTypes []BaseItemKind) ApiGetQueryFiltersLegacyRequest {
	r.includeItemTypes = &includeItemTypes
	return r
}

// Optional. Filter by MediaType. Allows multiple, comma delimited.
func (r ApiGetQueryFiltersLegacyRequest) MediaTypes(mediaTypes []MediaType) ApiGetQueryFiltersLegacyRequest {
	r.mediaTypes = &mediaTypes
	return r
}

func (r ApiGetQueryFiltersLegacyRequest) Execute() (*QueryFiltersLegacy, *http.Response, error) {
	return r.ApiService.GetQueryFiltersLegacyExecute(r)
}

/*
GetQueryFiltersLegacy Gets legacy query filters.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetQueryFiltersLegacyRequest
*/
func (a *FilterAPIService) GetQueryFiltersLegacy(ctx context.Context) ApiGetQueryFiltersLegacyRequest {
	return ApiGetQueryFiltersLegacyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QueryFiltersLegacy
func (a *FilterAPIService) GetQueryFiltersLegacyExecute(r ApiGetQueryFiltersLegacyRequest) (*QueryFiltersLegacy, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QueryFiltersLegacy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilterAPIService.GetQueryFiltersLegacy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Items/Filters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.parentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parentId", r.parentId, "form", "")
	}
	if r.includeItemTypes != nil {
		t := *r.includeItemTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "includeItemTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "includeItemTypes", t, "form", "multi")
		}
	}
	if r.mediaTypes != nil {
		t := *r.mediaTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "mediaTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "mediaTypes", t, "form", "multi")
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
