/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.7
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


type StudiosAPI interface {

	/*
	GetStudio Gets a studio by name.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param name Studio name.
	@return ApiGetStudioRequest
	*/
	GetStudio(ctx context.Context, name string) ApiGetStudioRequest

	// GetStudioExecute executes the request
	//  @return BaseItemDto
	GetStudioExecute(r ApiGetStudioRequest) (*BaseItemDto, *http.Response, error)

	/*
	GetStudios Gets all studios from a given item, folder, or the entire library.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetStudiosRequest
	*/
	GetStudios(ctx context.Context) ApiGetStudiosRequest

	// GetStudiosExecute executes the request
	//  @return BaseItemDtoQueryResult
	GetStudiosExecute(r ApiGetStudiosRequest) (*BaseItemDtoQueryResult, *http.Response, error)
}

// StudiosAPIService StudiosAPI service
type StudiosAPIService service

type ApiGetStudioRequest struct {
	ctx context.Context
	ApiService StudiosAPI
	name string
	userId *string
}

// Optional. Filter by user id, and attach user data.
func (r ApiGetStudioRequest) UserId(userId string) ApiGetStudioRequest {
	r.userId = &userId
	return r
}

func (r ApiGetStudioRequest) Execute() (*BaseItemDto, *http.Response, error) {
	return r.ApiService.GetStudioExecute(r)
}

/*
GetStudio Gets a studio by name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param name Studio name.
 @return ApiGetStudioRequest
*/
func (a *StudiosAPIService) GetStudio(ctx context.Context, name string) ApiGetStudioRequest {
	return ApiGetStudioRequest{
		ApiService: a,
		ctx: ctx,
		name: name,
	}
}

// Execute executes the request
//  @return BaseItemDto
func (a *StudiosAPIService) GetStudioExecute(r ApiGetStudioRequest) (*BaseItemDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiosAPIService.GetStudio")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Studios/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
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

type ApiGetStudiosRequest struct {
	ctx context.Context
	ApiService StudiosAPI
	startIndex *int32
	limit *int32
	searchTerm *string
	parentId *string
	fields *[]ItemFields
	excludeItemTypes *[]BaseItemKind
	includeItemTypes *[]BaseItemKind
	isFavorite *bool
	enableUserData *bool
	imageTypeLimit *int32
	enableImageTypes *[]ImageType
	userId *string
	nameStartsWithOrGreater *string
	nameStartsWith *string
	nameLessThan *string
	enableImages *bool
	enableTotalRecordCount *bool
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r ApiGetStudiosRequest) StartIndex(startIndex int32) ApiGetStudiosRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return.
func (r ApiGetStudiosRequest) Limit(limit int32) ApiGetStudiosRequest {
	r.limit = &limit
	return r
}

// Optional. Search term.
func (r ApiGetStudiosRequest) SearchTerm(searchTerm string) ApiGetStudiosRequest {
	r.searchTerm = &searchTerm
	return r
}

// Specify this to localize the search to a specific item or folder. Omit to use the root.
func (r ApiGetStudiosRequest) ParentId(parentId string) ApiGetStudiosRequest {
	r.parentId = &parentId
	return r
}

// Optional. Specify additional fields of information to return in the output.
func (r ApiGetStudiosRequest) Fields(fields []ItemFields) ApiGetStudiosRequest {
	r.fields = &fields
	return r
}

// Optional. If specified, results will be filtered out based on item type. This allows multiple, comma delimited.
func (r ApiGetStudiosRequest) ExcludeItemTypes(excludeItemTypes []BaseItemKind) ApiGetStudiosRequest {
	r.excludeItemTypes = &excludeItemTypes
	return r
}

// Optional. If specified, results will be filtered based on item type. This allows multiple, comma delimited.
func (r ApiGetStudiosRequest) IncludeItemTypes(includeItemTypes []BaseItemKind) ApiGetStudiosRequest {
	r.includeItemTypes = &includeItemTypes
	return r
}

// Optional filter by items that are marked as favorite, or not.
func (r ApiGetStudiosRequest) IsFavorite(isFavorite bool) ApiGetStudiosRequest {
	r.isFavorite = &isFavorite
	return r
}

// Optional, include user data.
func (r ApiGetStudiosRequest) EnableUserData(enableUserData bool) ApiGetStudiosRequest {
	r.enableUserData = &enableUserData
	return r
}

// Optional, the max number of images to return, per image type.
func (r ApiGetStudiosRequest) ImageTypeLimit(imageTypeLimit int32) ApiGetStudiosRequest {
	r.imageTypeLimit = &imageTypeLimit
	return r
}

// Optional. The image types to include in the output.
func (r ApiGetStudiosRequest) EnableImageTypes(enableImageTypes []ImageType) ApiGetStudiosRequest {
	r.enableImageTypes = &enableImageTypes
	return r
}

// User id.
func (r ApiGetStudiosRequest) UserId(userId string) ApiGetStudiosRequest {
	r.userId = &userId
	return r
}

// Optional filter by items whose name is sorted equally or greater than a given input string.
func (r ApiGetStudiosRequest) NameStartsWithOrGreater(nameStartsWithOrGreater string) ApiGetStudiosRequest {
	r.nameStartsWithOrGreater = &nameStartsWithOrGreater
	return r
}

// Optional filter by items whose name is sorted equally than a given input string.
func (r ApiGetStudiosRequest) NameStartsWith(nameStartsWith string) ApiGetStudiosRequest {
	r.nameStartsWith = &nameStartsWith
	return r
}

// Optional filter by items whose name is equally or lesser than a given input string.
func (r ApiGetStudiosRequest) NameLessThan(nameLessThan string) ApiGetStudiosRequest {
	r.nameLessThan = &nameLessThan
	return r
}

// Optional, include image information in output.
func (r ApiGetStudiosRequest) EnableImages(enableImages bool) ApiGetStudiosRequest {
	r.enableImages = &enableImages
	return r
}

// Total record count.
func (r ApiGetStudiosRequest) EnableTotalRecordCount(enableTotalRecordCount bool) ApiGetStudiosRequest {
	r.enableTotalRecordCount = &enableTotalRecordCount
	return r
}

func (r ApiGetStudiosRequest) Execute() (*BaseItemDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetStudiosExecute(r)
}

/*
GetStudios Gets all studios from a given item, folder, or the entire library.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetStudiosRequest
*/
func (a *StudiosAPIService) GetStudios(ctx context.Context) ApiGetStudiosRequest {
	return ApiGetStudiosRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BaseItemDtoQueryResult
func (a *StudiosAPIService) GetStudiosExecute(r ApiGetStudiosRequest) (*BaseItemDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StudiosAPIService.GetStudios")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Studios"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.searchTerm != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "searchTerm", r.searchTerm, "")
	}
	if r.parentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parentId", r.parentId, "")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeItemTypes != nil {
		t := *r.excludeItemTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "excludeItemTypes", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "excludeItemTypes", t, "multi")
		}
	}
	if r.includeItemTypes != nil {
		t := *r.includeItemTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "includeItemTypes", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "includeItemTypes", t, "multi")
		}
	}
	if r.isFavorite != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isFavorite", r.isFavorite, "")
	}
	if r.enableUserData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableUserData", r.enableUserData, "")
	}
	if r.imageTypeLimit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "imageTypeLimit", r.imageTypeLimit, "")
	}
	if r.enableImageTypes != nil {
		t := *r.enableImageTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "enableImageTypes", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "enableImageTypes", t, "multi")
		}
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
	}
	if r.nameStartsWithOrGreater != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nameStartsWithOrGreater", r.nameStartsWithOrGreater, "")
	}
	if r.nameStartsWith != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nameStartsWith", r.nameStartsWith, "")
	}
	if r.nameLessThan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nameLessThan", r.nameLessThan, "")
	}
	if r.enableImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableImages", r.enableImages, "")
	} else {
		var defaultValue bool = true
		r.enableImages = &defaultValue
	}
	if r.enableTotalRecordCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableTotalRecordCount", r.enableTotalRecordCount, "")
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
