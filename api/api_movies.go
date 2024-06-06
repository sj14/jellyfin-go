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
	"reflect"
)


type MoviesAPI interface {

	/*
	GetMovieRecommendations Gets movie recommendations.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetMovieRecommendationsRequest
	*/
	GetMovieRecommendations(ctx context.Context) ApiGetMovieRecommendationsRequest

	// GetMovieRecommendationsExecute executes the request
	//  @return []RecommendationDto
	GetMovieRecommendationsExecute(r ApiGetMovieRecommendationsRequest) ([]RecommendationDto, *http.Response, error)
}

// MoviesAPIService MoviesAPI service
type MoviesAPIService service

type ApiGetMovieRecommendationsRequest struct {
	ctx context.Context
	ApiService MoviesAPI
	userId *string
	parentId *string
	fields *[]ItemFields
	categoryLimit *int32
	itemLimit *int32
}

// Optional. Filter by user id, and attach user data.
func (r ApiGetMovieRecommendationsRequest) UserId(userId string) ApiGetMovieRecommendationsRequest {
	r.userId = &userId
	return r
}

// Specify this to localize the search to a specific item or folder. Omit to use the root.
func (r ApiGetMovieRecommendationsRequest) ParentId(parentId string) ApiGetMovieRecommendationsRequest {
	r.parentId = &parentId
	return r
}

// Optional. The fields to return.
func (r ApiGetMovieRecommendationsRequest) Fields(fields []ItemFields) ApiGetMovieRecommendationsRequest {
	r.fields = &fields
	return r
}

// The max number of categories to return.
func (r ApiGetMovieRecommendationsRequest) CategoryLimit(categoryLimit int32) ApiGetMovieRecommendationsRequest {
	r.categoryLimit = &categoryLimit
	return r
}

// The max number of items to return per category.
func (r ApiGetMovieRecommendationsRequest) ItemLimit(itemLimit int32) ApiGetMovieRecommendationsRequest {
	r.itemLimit = &itemLimit
	return r
}

func (r ApiGetMovieRecommendationsRequest) Execute() ([]RecommendationDto, *http.Response, error) {
	return r.ApiService.GetMovieRecommendationsExecute(r)
}

/*
GetMovieRecommendations Gets movie recommendations.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMovieRecommendationsRequest
*/
func (a *MoviesAPIService) GetMovieRecommendations(ctx context.Context) ApiGetMovieRecommendationsRequest {
	return ApiGetMovieRecommendationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []RecommendationDto
func (a *MoviesAPIService) GetMovieRecommendationsExecute(r ApiGetMovieRecommendationsRequest) ([]RecommendationDto, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []RecommendationDto
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MoviesAPIService.GetMovieRecommendations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Movies/Recommendations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "")
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
	if r.categoryLimit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "categoryLimit", r.categoryLimit, "")
	} else {
		var defaultValue int32 = 5
		r.categoryLimit = &defaultValue
	}
	if r.itemLimit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemLimit", r.itemLimit, "")
	} else {
		var defaultValue int32 = 8
		r.itemLimit = &defaultValue
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
