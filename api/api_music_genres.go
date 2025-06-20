/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
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


type MusicGenresAPI interface {

	/*
	GetMusicGenre Gets a music genre, by name.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param genreName The genre name.
	@return ApiGetMusicGenreRequest
	*/
	GetMusicGenre(ctx context.Context, genreName string) ApiGetMusicGenreRequest

	// GetMusicGenreExecute executes the request
	//  @return BaseItemDto
	GetMusicGenreExecute(r ApiGetMusicGenreRequest) (*BaseItemDto, *http.Response, error)

	/*
	GetMusicGenres Gets all music genres from a given item, folder, or the entire library.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetMusicGenresRequest

	Deprecated
	*/
	GetMusicGenres(ctx context.Context) ApiGetMusicGenresRequest

	// GetMusicGenresExecute executes the request
	//  @return BaseItemDtoQueryResult
	// Deprecated
	GetMusicGenresExecute(r ApiGetMusicGenresRequest) (*BaseItemDtoQueryResult, *http.Response, error)
}

// MusicGenresAPIService MusicGenresAPI service
type MusicGenresAPIService service

type ApiGetMusicGenreRequest struct {
	ctx context.Context
	ApiService MusicGenresAPI
	genreName string
	userId *string
}

// Optional. Filter by user id, and attach user data.
func (r ApiGetMusicGenreRequest) UserId(userId string) ApiGetMusicGenreRequest {
	r.userId = &userId
	return r
}

func (r ApiGetMusicGenreRequest) Execute() (*BaseItemDto, *http.Response, error) {
	return r.ApiService.GetMusicGenreExecute(r)
}

/*
GetMusicGenre Gets a music genre, by name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param genreName The genre name.
 @return ApiGetMusicGenreRequest
*/
func (a *MusicGenresAPIService) GetMusicGenre(ctx context.Context, genreName string) ApiGetMusicGenreRequest {
	return ApiGetMusicGenreRequest{
		ApiService: a,
		ctx: ctx,
		genreName: genreName,
	}
}

// Execute executes the request
//  @return BaseItemDto
func (a *MusicGenresAPIService) GetMusicGenreExecute(r ApiGetMusicGenreRequest) (*BaseItemDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MusicGenresAPIService.GetMusicGenre")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/MusicGenres/{genreName}"
	localVarPath = strings.Replace(localVarPath, "{"+"genreName"+"}", url.PathEscape(parameterValueToString(r.genreName, "genreName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase", "text/html"}

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

type ApiGetMusicGenresRequest struct {
	ctx context.Context
	ApiService MusicGenresAPI
	startIndex *int32
	limit *int32
	searchTerm *string
	parentId *string
	fields *[]ItemFields
	excludeItemTypes *[]BaseItemKind
	includeItemTypes *[]BaseItemKind
	isFavorite *bool
	imageTypeLimit *int32
	enableImageTypes *[]ImageType
	userId *string
	nameStartsWithOrGreater *string
	nameStartsWith *string
	nameLessThan *string
	sortBy *[]ItemSortBy
	sortOrder *[]SortOrder
	enableImages *bool
	enableTotalRecordCount *bool
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r ApiGetMusicGenresRequest) StartIndex(startIndex int32) ApiGetMusicGenresRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return.
func (r ApiGetMusicGenresRequest) Limit(limit int32) ApiGetMusicGenresRequest {
	r.limit = &limit
	return r
}

// The search term.
func (r ApiGetMusicGenresRequest) SearchTerm(searchTerm string) ApiGetMusicGenresRequest {
	r.searchTerm = &searchTerm
	return r
}

// Specify this to localize the search to a specific item or folder. Omit to use the root.
func (r ApiGetMusicGenresRequest) ParentId(parentId string) ApiGetMusicGenresRequest {
	r.parentId = &parentId
	return r
}

// Optional. Specify additional fields of information to return in the output.
func (r ApiGetMusicGenresRequest) Fields(fields []ItemFields) ApiGetMusicGenresRequest {
	r.fields = &fields
	return r
}

// Optional. If specified, results will be filtered out based on item type. This allows multiple, comma delimited.
func (r ApiGetMusicGenresRequest) ExcludeItemTypes(excludeItemTypes []BaseItemKind) ApiGetMusicGenresRequest {
	r.excludeItemTypes = &excludeItemTypes
	return r
}

// Optional. If specified, results will be filtered in based on item type. This allows multiple, comma delimited.
func (r ApiGetMusicGenresRequest) IncludeItemTypes(includeItemTypes []BaseItemKind) ApiGetMusicGenresRequest {
	r.includeItemTypes = &includeItemTypes
	return r
}

// Optional filter by items that are marked as favorite, or not.
func (r ApiGetMusicGenresRequest) IsFavorite(isFavorite bool) ApiGetMusicGenresRequest {
	r.isFavorite = &isFavorite
	return r
}

// Optional, the max number of images to return, per image type.
func (r ApiGetMusicGenresRequest) ImageTypeLimit(imageTypeLimit int32) ApiGetMusicGenresRequest {
	r.imageTypeLimit = &imageTypeLimit
	return r
}

// Optional. The image types to include in the output.
func (r ApiGetMusicGenresRequest) EnableImageTypes(enableImageTypes []ImageType) ApiGetMusicGenresRequest {
	r.enableImageTypes = &enableImageTypes
	return r
}

// User id.
func (r ApiGetMusicGenresRequest) UserId(userId string) ApiGetMusicGenresRequest {
	r.userId = &userId
	return r
}

// Optional filter by items whose name is sorted equally or greater than a given input string.
func (r ApiGetMusicGenresRequest) NameStartsWithOrGreater(nameStartsWithOrGreater string) ApiGetMusicGenresRequest {
	r.nameStartsWithOrGreater = &nameStartsWithOrGreater
	return r
}

// Optional filter by items whose name is sorted equally than a given input string.
func (r ApiGetMusicGenresRequest) NameStartsWith(nameStartsWith string) ApiGetMusicGenresRequest {
	r.nameStartsWith = &nameStartsWith
	return r
}

// Optional filter by items whose name is equally or lesser than a given input string.
func (r ApiGetMusicGenresRequest) NameLessThan(nameLessThan string) ApiGetMusicGenresRequest {
	r.nameLessThan = &nameLessThan
	return r
}

// Optional. Specify one or more sort orders, comma delimited.
func (r ApiGetMusicGenresRequest) SortBy(sortBy []ItemSortBy) ApiGetMusicGenresRequest {
	r.sortBy = &sortBy
	return r
}

// Sort Order - Ascending,Descending.
func (r ApiGetMusicGenresRequest) SortOrder(sortOrder []SortOrder) ApiGetMusicGenresRequest {
	r.sortOrder = &sortOrder
	return r
}

// Optional, include image information in output.
func (r ApiGetMusicGenresRequest) EnableImages(enableImages bool) ApiGetMusicGenresRequest {
	r.enableImages = &enableImages
	return r
}

// Optional. Include total record count.
func (r ApiGetMusicGenresRequest) EnableTotalRecordCount(enableTotalRecordCount bool) ApiGetMusicGenresRequest {
	r.enableTotalRecordCount = &enableTotalRecordCount
	return r
}

func (r ApiGetMusicGenresRequest) Execute() (*BaseItemDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetMusicGenresExecute(r)
}

/*
GetMusicGenres Gets all music genres from a given item, folder, or the entire library.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMusicGenresRequest

Deprecated
*/
func (a *MusicGenresAPIService) GetMusicGenres(ctx context.Context) ApiGetMusicGenresRequest {
	return ApiGetMusicGenresRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BaseItemDtoQueryResult
// Deprecated
func (a *MusicGenresAPIService) GetMusicGenresExecute(r ApiGetMusicGenresRequest) (*BaseItemDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MusicGenresAPIService.GetMusicGenres")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/MusicGenres"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.searchTerm != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchTerm", r.searchTerm, "form", "")
	}
	if r.parentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parentId", r.parentId, "form", "")
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
	if r.excludeItemTypes != nil {
		t := *r.excludeItemTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "excludeItemTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "excludeItemTypes", t, "form", "multi")
		}
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
	if r.isFavorite != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isFavorite", r.isFavorite, "form", "")
	}
	if r.imageTypeLimit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "imageTypeLimit", r.imageTypeLimit, "form", "")
	}
	if r.enableImageTypes != nil {
		t := *r.enableImageTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "enableImageTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "enableImageTypes", t, "form", "multi")
		}
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
	}
	if r.nameStartsWithOrGreater != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nameStartsWithOrGreater", r.nameStartsWithOrGreater, "form", "")
	}
	if r.nameStartsWith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nameStartsWith", r.nameStartsWith, "form", "")
	}
	if r.nameLessThan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nameLessThan", r.nameLessThan, "form", "")
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
	if r.enableImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableImages", r.enableImages, "form", "")
	} else {
		var defaultValue bool = true
		r.enableImages = &defaultValue
	}
	if r.enableTotalRecordCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableTotalRecordCount", r.enableTotalRecordCount, "form", "")
	} else {
		var defaultValue bool = true
		r.enableTotalRecordCount = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase", "text/html"}

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
