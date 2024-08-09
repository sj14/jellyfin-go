/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.9
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


type YearsAPI interface {

	/*
	GetYear Gets a year.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param year The year.
	@return ApiGetYearRequest
	*/
	GetYear(ctx context.Context, year int32) ApiGetYearRequest

	// GetYearExecute executes the request
	//  @return BaseItemDto
	GetYearExecute(r ApiGetYearRequest) (*BaseItemDto, *http.Response, error)

	/*
	GetYears Get years.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetYearsRequest
	*/
	GetYears(ctx context.Context) ApiGetYearsRequest

	// GetYearsExecute executes the request
	//  @return BaseItemDtoQueryResult
	GetYearsExecute(r ApiGetYearsRequest) (*BaseItemDtoQueryResult, *http.Response, error)
}

// YearsAPIService YearsAPI service
type YearsAPIService service

type ApiGetYearRequest struct {
	ctx context.Context
	ApiService YearsAPI
	year int32
	userId *string
}

// Optional. Filter by user id, and attach user data.
func (r ApiGetYearRequest) UserId(userId string) ApiGetYearRequest {
	r.userId = &userId
	return r
}

func (r ApiGetYearRequest) Execute() (*BaseItemDto, *http.Response, error) {
	return r.ApiService.GetYearExecute(r)
}

/*
GetYear Gets a year.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param year The year.
 @return ApiGetYearRequest
*/
func (a *YearsAPIService) GetYear(ctx context.Context, year int32) ApiGetYearRequest {
	return ApiGetYearRequest{
		ApiService: a,
		ctx: ctx,
		year: year,
	}
}

// Execute executes the request
//  @return BaseItemDto
func (a *YearsAPIService) GetYearExecute(r ApiGetYearRequest) (*BaseItemDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "YearsAPIService.GetYear")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Years/{year}"
	localVarPath = strings.Replace(localVarPath, "{"+"year"+"}", url.PathEscape(parameterValueToString(r.year, "year")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiGetYearsRequest struct {
	ctx context.Context
	ApiService YearsAPI
	startIndex *int32
	limit *int32
	sortOrder *[]SortOrder
	parentId *string
	fields *[]ItemFields
	excludeItemTypes *[]BaseItemKind
	includeItemTypes *[]BaseItemKind
	mediaTypes *[]MediaType
	sortBy *[]ItemSortBy
	enableUserData *bool
	imageTypeLimit *int32
	enableImageTypes *[]ImageType
	userId *string
	recursive *bool
	enableImages *bool
}

// Skips over a given number of items within the results. Use for paging.
func (r ApiGetYearsRequest) StartIndex(startIndex int32) ApiGetYearsRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return.
func (r ApiGetYearsRequest) Limit(limit int32) ApiGetYearsRequest {
	r.limit = &limit
	return r
}

// Sort Order - Ascending,Descending.
func (r ApiGetYearsRequest) SortOrder(sortOrder []SortOrder) ApiGetYearsRequest {
	r.sortOrder = &sortOrder
	return r
}

// Specify this to localize the search to a specific item or folder. Omit to use the root.
func (r ApiGetYearsRequest) ParentId(parentId string) ApiGetYearsRequest {
	r.parentId = &parentId
	return r
}

// Optional. Specify additional fields of information to return in the output.
func (r ApiGetYearsRequest) Fields(fields []ItemFields) ApiGetYearsRequest {
	r.fields = &fields
	return r
}

// Optional. If specified, results will be excluded based on item type. This allows multiple, comma delimited.
func (r ApiGetYearsRequest) ExcludeItemTypes(excludeItemTypes []BaseItemKind) ApiGetYearsRequest {
	r.excludeItemTypes = &excludeItemTypes
	return r
}

// Optional. If specified, results will be included based on item type. This allows multiple, comma delimited.
func (r ApiGetYearsRequest) IncludeItemTypes(includeItemTypes []BaseItemKind) ApiGetYearsRequest {
	r.includeItemTypes = &includeItemTypes
	return r
}

// Optional. Filter by MediaType. Allows multiple, comma delimited.
func (r ApiGetYearsRequest) MediaTypes(mediaTypes []MediaType) ApiGetYearsRequest {
	r.mediaTypes = &mediaTypes
	return r
}

// Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime.
func (r ApiGetYearsRequest) SortBy(sortBy []ItemSortBy) ApiGetYearsRequest {
	r.sortBy = &sortBy
	return r
}

// Optional. Include user data.
func (r ApiGetYearsRequest) EnableUserData(enableUserData bool) ApiGetYearsRequest {
	r.enableUserData = &enableUserData
	return r
}

// Optional. The max number of images to return, per image type.
func (r ApiGetYearsRequest) ImageTypeLimit(imageTypeLimit int32) ApiGetYearsRequest {
	r.imageTypeLimit = &imageTypeLimit
	return r
}

// Optional. The image types to include in the output.
func (r ApiGetYearsRequest) EnableImageTypes(enableImageTypes []ImageType) ApiGetYearsRequest {
	r.enableImageTypes = &enableImageTypes
	return r
}

// User Id.
func (r ApiGetYearsRequest) UserId(userId string) ApiGetYearsRequest {
	r.userId = &userId
	return r
}

// Search recursively.
func (r ApiGetYearsRequest) Recursive(recursive bool) ApiGetYearsRequest {
	r.recursive = &recursive
	return r
}

// Optional. Include image information in output.
func (r ApiGetYearsRequest) EnableImages(enableImages bool) ApiGetYearsRequest {
	r.enableImages = &enableImages
	return r
}

func (r ApiGetYearsRequest) Execute() (*BaseItemDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetYearsExecute(r)
}

/*
GetYears Get years.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetYearsRequest
*/
func (a *YearsAPIService) GetYears(ctx context.Context) ApiGetYearsRequest {
	return ApiGetYearsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BaseItemDtoQueryResult
func (a *YearsAPIService) GetYearsExecute(r ApiGetYearsRequest) (*BaseItemDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "YearsAPIService.GetYears")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Years"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.sortOrder != nil {
		t := *r.sortOrder
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sortOrder", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sortOrder", t, "multi")
		}
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
	if r.mediaTypes != nil {
		t := *r.mediaTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "mediaTypes", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "mediaTypes", t, "multi")
		}
	}
	if r.sortBy != nil {
		t := *r.sortBy
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", t, "multi")
		}
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
	if r.recursive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recursive", r.recursive, "")
	} else {
		var defaultValue bool = true
		r.recursive = &defaultValue
	}
	if r.enableImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableImages", r.enableImages, "")
	} else {
		var defaultValue bool = true
		r.enableImages = &defaultValue
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
