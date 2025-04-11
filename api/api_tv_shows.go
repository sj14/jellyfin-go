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
	"strings"
	"reflect"
	"time"
)


type TvShowsAPI interface {

	/*
	GetEpisodes Gets episodes for a tv season.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param seriesId The series id.
	@return ApiGetEpisodesRequest
	*/
	GetEpisodes(ctx context.Context, seriesId string) ApiGetEpisodesRequest

	// GetEpisodesExecute executes the request
	//  @return BaseItemDtoQueryResult
	GetEpisodesExecute(r ApiGetEpisodesRequest) (*BaseItemDtoQueryResult, *http.Response, error)

	/*
	GetNextUp Gets a list of next up episodes.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetNextUpRequest
	*/
	GetNextUp(ctx context.Context) ApiGetNextUpRequest

	// GetNextUpExecute executes the request
	//  @return BaseItemDtoQueryResult
	GetNextUpExecute(r ApiGetNextUpRequest) (*BaseItemDtoQueryResult, *http.Response, error)

	/*
	GetSeasons Gets seasons for a tv series.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param seriesId The series id.
	@return ApiGetSeasonsRequest
	*/
	GetSeasons(ctx context.Context, seriesId string) ApiGetSeasonsRequest

	// GetSeasonsExecute executes the request
	//  @return BaseItemDtoQueryResult
	GetSeasonsExecute(r ApiGetSeasonsRequest) (*BaseItemDtoQueryResult, *http.Response, error)

	/*
	GetUpcomingEpisodes Gets a list of upcoming episodes.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetUpcomingEpisodesRequest
	*/
	GetUpcomingEpisodes(ctx context.Context) ApiGetUpcomingEpisodesRequest

	// GetUpcomingEpisodesExecute executes the request
	//  @return BaseItemDtoQueryResult
	GetUpcomingEpisodesExecute(r ApiGetUpcomingEpisodesRequest) (*BaseItemDtoQueryResult, *http.Response, error)
}

// TvShowsAPIService TvShowsAPI service
type TvShowsAPIService service

type ApiGetEpisodesRequest struct {
	ctx context.Context
	ApiService TvShowsAPI
	seriesId string
	userId *string
	fields *[]ItemFields
	season *int32
	seasonId *string
	isMissing *bool
	adjacentTo *string
	startItemId *string
	startIndex *int32
	limit *int32
	enableImages *bool
	imageTypeLimit *int32
	enableImageTypes *[]ImageType
	enableUserData *bool
	sortBy *ItemSortBy
}

// The user id.
func (r ApiGetEpisodesRequest) UserId(userId string) ApiGetEpisodesRequest {
	r.userId = &userId
	return r
}

// Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls.
func (r ApiGetEpisodesRequest) Fields(fields []ItemFields) ApiGetEpisodesRequest {
	r.fields = &fields
	return r
}

// Optional filter by season number.
func (r ApiGetEpisodesRequest) Season(season int32) ApiGetEpisodesRequest {
	r.season = &season
	return r
}

// Optional. Filter by season id.
func (r ApiGetEpisodesRequest) SeasonId(seasonId string) ApiGetEpisodesRequest {
	r.seasonId = &seasonId
	return r
}

// Optional. Filter by items that are missing episodes or not.
func (r ApiGetEpisodesRequest) IsMissing(isMissing bool) ApiGetEpisodesRequest {
	r.isMissing = &isMissing
	return r
}

// Optional. Return items that are siblings of a supplied item.
func (r ApiGetEpisodesRequest) AdjacentTo(adjacentTo string) ApiGetEpisodesRequest {
	r.adjacentTo = &adjacentTo
	return r
}

// Optional. Skip through the list until a given item is found.
func (r ApiGetEpisodesRequest) StartItemId(startItemId string) ApiGetEpisodesRequest {
	r.startItemId = &startItemId
	return r
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r ApiGetEpisodesRequest) StartIndex(startIndex int32) ApiGetEpisodesRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return.
func (r ApiGetEpisodesRequest) Limit(limit int32) ApiGetEpisodesRequest {
	r.limit = &limit
	return r
}

// Optional, include image information in output.
func (r ApiGetEpisodesRequest) EnableImages(enableImages bool) ApiGetEpisodesRequest {
	r.enableImages = &enableImages
	return r
}

// Optional, the max number of images to return, per image type.
func (r ApiGetEpisodesRequest) ImageTypeLimit(imageTypeLimit int32) ApiGetEpisodesRequest {
	r.imageTypeLimit = &imageTypeLimit
	return r
}

// Optional. The image types to include in the output.
func (r ApiGetEpisodesRequest) EnableImageTypes(enableImageTypes []ImageType) ApiGetEpisodesRequest {
	r.enableImageTypes = &enableImageTypes
	return r
}

// Optional. Include user data.
func (r ApiGetEpisodesRequest) EnableUserData(enableUserData bool) ApiGetEpisodesRequest {
	r.enableUserData = &enableUserData
	return r
}

// Optional. Specify one or more sort orders, comma delimited. Options: Album, AlbumArtist, Artist, Budget, CommunityRating, CriticRating, DateCreated, DatePlayed, PlayCount, PremiereDate, ProductionYear, SortName, Random, Revenue, Runtime.
func (r ApiGetEpisodesRequest) SortBy(sortBy ItemSortBy) ApiGetEpisodesRequest {
	r.sortBy = &sortBy
	return r
}

func (r ApiGetEpisodesRequest) Execute() (*BaseItemDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetEpisodesExecute(r)
}

/*
GetEpisodes Gets episodes for a tv season.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param seriesId The series id.
 @return ApiGetEpisodesRequest
*/
func (a *TvShowsAPIService) GetEpisodes(ctx context.Context, seriesId string) ApiGetEpisodesRequest {
	return ApiGetEpisodesRequest{
		ApiService: a,
		ctx: ctx,
		seriesId: seriesId,
	}
}

// Execute executes the request
//  @return BaseItemDtoQueryResult
func (a *TvShowsAPIService) GetEpisodesExecute(r ApiGetEpisodesRequest) (*BaseItemDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TvShowsAPIService.GetEpisodes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Shows/{seriesId}/Episodes"
	localVarPath = strings.Replace(localVarPath, "{"+"seriesId"+"}", url.PathEscape(parameterValueToString(r.seriesId, "seriesId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
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
	if r.season != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "season", r.season, "form", "")
	}
	if r.seasonId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "seasonId", r.seasonId, "form", "")
	}
	if r.isMissing != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isMissing", r.isMissing, "form", "")
	}
	if r.adjacentTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adjacentTo", r.adjacentTo, "form", "")
	}
	if r.startItemId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startItemId", r.startItemId, "form", "")
	}
	if r.startIndex != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startIndex", r.startIndex, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.enableImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableImages", r.enableImages, "form", "")
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
	if r.enableUserData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableUserData", r.enableUserData, "form", "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", r.sortBy, "form", "")
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

type ApiGetNextUpRequest struct {
	ctx context.Context
	ApiService TvShowsAPI
	userId *string
	startIndex *int32
	limit *int32
	fields *[]ItemFields
	seriesId *string
	parentId *string
	enableImages *bool
	imageTypeLimit *int32
	enableImageTypes *[]ImageType
	enableUserData *bool
	nextUpDateCutoff *time.Time
	enableTotalRecordCount *bool
	disableFirstEpisode *bool
	enableResumable *bool
	enableRewatching *bool
}

// The user id of the user to get the next up episodes for.
func (r ApiGetNextUpRequest) UserId(userId string) ApiGetNextUpRequest {
	r.userId = &userId
	return r
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r ApiGetNextUpRequest) StartIndex(startIndex int32) ApiGetNextUpRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return.
func (r ApiGetNextUpRequest) Limit(limit int32) ApiGetNextUpRequest {
	r.limit = &limit
	return r
}

// Optional. Specify additional fields of information to return in the output.
func (r ApiGetNextUpRequest) Fields(fields []ItemFields) ApiGetNextUpRequest {
	r.fields = &fields
	return r
}

// Optional. Filter by series id.
func (r ApiGetNextUpRequest) SeriesId(seriesId string) ApiGetNextUpRequest {
	r.seriesId = &seriesId
	return r
}

// Optional. Specify this to localize the search to a specific item or folder. Omit to use the root.
func (r ApiGetNextUpRequest) ParentId(parentId string) ApiGetNextUpRequest {
	r.parentId = &parentId
	return r
}

// Optional. Include image information in output.
func (r ApiGetNextUpRequest) EnableImages(enableImages bool) ApiGetNextUpRequest {
	r.enableImages = &enableImages
	return r
}

// Optional. The max number of images to return, per image type.
func (r ApiGetNextUpRequest) ImageTypeLimit(imageTypeLimit int32) ApiGetNextUpRequest {
	r.imageTypeLimit = &imageTypeLimit
	return r
}

// Optional. The image types to include in the output.
func (r ApiGetNextUpRequest) EnableImageTypes(enableImageTypes []ImageType) ApiGetNextUpRequest {
	r.enableImageTypes = &enableImageTypes
	return r
}

// Optional. Include user data.
func (r ApiGetNextUpRequest) EnableUserData(enableUserData bool) ApiGetNextUpRequest {
	r.enableUserData = &enableUserData
	return r
}

// Optional. Starting date of shows to show in Next Up section.
func (r ApiGetNextUpRequest) NextUpDateCutoff(nextUpDateCutoff time.Time) ApiGetNextUpRequest {
	r.nextUpDateCutoff = &nextUpDateCutoff
	return r
}

// Whether to enable the total records count. Defaults to true.
func (r ApiGetNextUpRequest) EnableTotalRecordCount(enableTotalRecordCount bool) ApiGetNextUpRequest {
	r.enableTotalRecordCount = &enableTotalRecordCount
	return r
}

// Whether to disable sending the first episode in a series as next up.
func (r ApiGetNextUpRequest) DisableFirstEpisode(disableFirstEpisode bool) ApiGetNextUpRequest {
	r.disableFirstEpisode = &disableFirstEpisode
	return r
}

// Whether to include resumable episodes in next up results.
func (r ApiGetNextUpRequest) EnableResumable(enableResumable bool) ApiGetNextUpRequest {
	r.enableResumable = &enableResumable
	return r
}

// Whether to include watched episodes in next up results.
func (r ApiGetNextUpRequest) EnableRewatching(enableRewatching bool) ApiGetNextUpRequest {
	r.enableRewatching = &enableRewatching
	return r
}

func (r ApiGetNextUpRequest) Execute() (*BaseItemDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetNextUpExecute(r)
}

/*
GetNextUp Gets a list of next up episodes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNextUpRequest
*/
func (a *TvShowsAPIService) GetNextUp(ctx context.Context) ApiGetNextUpRequest {
	return ApiGetNextUpRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BaseItemDtoQueryResult
func (a *TvShowsAPIService) GetNextUpExecute(r ApiGetNextUpRequest) (*BaseItemDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TvShowsAPIService.GetNextUp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Shows/NextUp"

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
	if r.seriesId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "seriesId", r.seriesId, "form", "")
	}
	if r.parentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parentId", r.parentId, "form", "")
	}
	if r.enableImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableImages", r.enableImages, "form", "")
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
	if r.enableUserData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableUserData", r.enableUserData, "form", "")
	}
	if r.nextUpDateCutoff != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nextUpDateCutoff", r.nextUpDateCutoff, "form", "")
	}
	if r.enableTotalRecordCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableTotalRecordCount", r.enableTotalRecordCount, "form", "")
	} else {
		var defaultValue bool = true
		r.enableTotalRecordCount = &defaultValue
	}
	if r.disableFirstEpisode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "disableFirstEpisode", r.disableFirstEpisode, "form", "")
	} else {
		var defaultValue bool = false
		r.disableFirstEpisode = &defaultValue
	}
	if r.enableResumable != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableResumable", r.enableResumable, "form", "")
	} else {
		var defaultValue bool = true
		r.enableResumable = &defaultValue
	}
	if r.enableRewatching != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableRewatching", r.enableRewatching, "form", "")
	} else {
		var defaultValue bool = false
		r.enableRewatching = &defaultValue
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

type ApiGetSeasonsRequest struct {
	ctx context.Context
	ApiService TvShowsAPI
	seriesId string
	userId *string
	fields *[]ItemFields
	isSpecialSeason *bool
	isMissing *bool
	adjacentTo *string
	enableImages *bool
	imageTypeLimit *int32
	enableImageTypes *[]ImageType
	enableUserData *bool
}

// The user id.
func (r ApiGetSeasonsRequest) UserId(userId string) ApiGetSeasonsRequest {
	r.userId = &userId
	return r
}

// Optional. Specify additional fields of information to return in the output. This allows multiple, comma delimited. Options: Budget, Chapters, DateCreated, Genres, HomePageUrl, IndexOptions, MediaStreams, Overview, ParentId, Path, People, ProviderIds, PrimaryImageAspectRatio, Revenue, SortName, Studios, Taglines, TrailerUrls.
func (r ApiGetSeasonsRequest) Fields(fields []ItemFields) ApiGetSeasonsRequest {
	r.fields = &fields
	return r
}

// Optional. Filter by special season.
func (r ApiGetSeasonsRequest) IsSpecialSeason(isSpecialSeason bool) ApiGetSeasonsRequest {
	r.isSpecialSeason = &isSpecialSeason
	return r
}

// Optional. Filter by items that are missing episodes or not.
func (r ApiGetSeasonsRequest) IsMissing(isMissing bool) ApiGetSeasonsRequest {
	r.isMissing = &isMissing
	return r
}

// Optional. Return items that are siblings of a supplied item.
func (r ApiGetSeasonsRequest) AdjacentTo(adjacentTo string) ApiGetSeasonsRequest {
	r.adjacentTo = &adjacentTo
	return r
}

// Optional. Include image information in output.
func (r ApiGetSeasonsRequest) EnableImages(enableImages bool) ApiGetSeasonsRequest {
	r.enableImages = &enableImages
	return r
}

// Optional. The max number of images to return, per image type.
func (r ApiGetSeasonsRequest) ImageTypeLimit(imageTypeLimit int32) ApiGetSeasonsRequest {
	r.imageTypeLimit = &imageTypeLimit
	return r
}

// Optional. The image types to include in the output.
func (r ApiGetSeasonsRequest) EnableImageTypes(enableImageTypes []ImageType) ApiGetSeasonsRequest {
	r.enableImageTypes = &enableImageTypes
	return r
}

// Optional. Include user data.
func (r ApiGetSeasonsRequest) EnableUserData(enableUserData bool) ApiGetSeasonsRequest {
	r.enableUserData = &enableUserData
	return r
}

func (r ApiGetSeasonsRequest) Execute() (*BaseItemDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetSeasonsExecute(r)
}

/*
GetSeasons Gets seasons for a tv series.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param seriesId The series id.
 @return ApiGetSeasonsRequest
*/
func (a *TvShowsAPIService) GetSeasons(ctx context.Context, seriesId string) ApiGetSeasonsRequest {
	return ApiGetSeasonsRequest{
		ApiService: a,
		ctx: ctx,
		seriesId: seriesId,
	}
}

// Execute executes the request
//  @return BaseItemDtoQueryResult
func (a *TvShowsAPIService) GetSeasonsExecute(r ApiGetSeasonsRequest) (*BaseItemDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TvShowsAPIService.GetSeasons")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Shows/{seriesId}/Seasons"
	localVarPath = strings.Replace(localVarPath, "{"+"seriesId"+"}", url.PathEscape(parameterValueToString(r.seriesId, "seriesId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userId", r.userId, "form", "")
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
	if r.isSpecialSeason != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isSpecialSeason", r.isSpecialSeason, "form", "")
	}
	if r.isMissing != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "isMissing", r.isMissing, "form", "")
	}
	if r.adjacentTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "adjacentTo", r.adjacentTo, "form", "")
	}
	if r.enableImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableImages", r.enableImages, "form", "")
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
	if r.enableUserData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableUserData", r.enableUserData, "form", "")
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

type ApiGetUpcomingEpisodesRequest struct {
	ctx context.Context
	ApiService TvShowsAPI
	userId *string
	startIndex *int32
	limit *int32
	fields *[]ItemFields
	parentId *string
	enableImages *bool
	imageTypeLimit *int32
	enableImageTypes *[]ImageType
	enableUserData *bool
}

// The user id of the user to get the upcoming episodes for.
func (r ApiGetUpcomingEpisodesRequest) UserId(userId string) ApiGetUpcomingEpisodesRequest {
	r.userId = &userId
	return r
}

// Optional. The record index to start at. All items with a lower index will be dropped from the results.
func (r ApiGetUpcomingEpisodesRequest) StartIndex(startIndex int32) ApiGetUpcomingEpisodesRequest {
	r.startIndex = &startIndex
	return r
}

// Optional. The maximum number of records to return.
func (r ApiGetUpcomingEpisodesRequest) Limit(limit int32) ApiGetUpcomingEpisodesRequest {
	r.limit = &limit
	return r
}

// Optional. Specify additional fields of information to return in the output.
func (r ApiGetUpcomingEpisodesRequest) Fields(fields []ItemFields) ApiGetUpcomingEpisodesRequest {
	r.fields = &fields
	return r
}

// Optional. Specify this to localize the search to a specific item or folder. Omit to use the root.
func (r ApiGetUpcomingEpisodesRequest) ParentId(parentId string) ApiGetUpcomingEpisodesRequest {
	r.parentId = &parentId
	return r
}

// Optional. Include image information in output.
func (r ApiGetUpcomingEpisodesRequest) EnableImages(enableImages bool) ApiGetUpcomingEpisodesRequest {
	r.enableImages = &enableImages
	return r
}

// Optional. The max number of images to return, per image type.
func (r ApiGetUpcomingEpisodesRequest) ImageTypeLimit(imageTypeLimit int32) ApiGetUpcomingEpisodesRequest {
	r.imageTypeLimit = &imageTypeLimit
	return r
}

// Optional. The image types to include in the output.
func (r ApiGetUpcomingEpisodesRequest) EnableImageTypes(enableImageTypes []ImageType) ApiGetUpcomingEpisodesRequest {
	r.enableImageTypes = &enableImageTypes
	return r
}

// Optional. Include user data.
func (r ApiGetUpcomingEpisodesRequest) EnableUserData(enableUserData bool) ApiGetUpcomingEpisodesRequest {
	r.enableUserData = &enableUserData
	return r
}

func (r ApiGetUpcomingEpisodesRequest) Execute() (*BaseItemDtoQueryResult, *http.Response, error) {
	return r.ApiService.GetUpcomingEpisodesExecute(r)
}

/*
GetUpcomingEpisodes Gets a list of upcoming episodes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUpcomingEpisodesRequest
*/
func (a *TvShowsAPIService) GetUpcomingEpisodes(ctx context.Context) ApiGetUpcomingEpisodesRequest {
	return ApiGetUpcomingEpisodesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BaseItemDtoQueryResult
func (a *TvShowsAPIService) GetUpcomingEpisodesExecute(r ApiGetUpcomingEpisodesRequest) (*BaseItemDtoQueryResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BaseItemDtoQueryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TvShowsAPIService.GetUpcomingEpisodes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Shows/Upcoming"

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
	if r.parentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parentId", r.parentId, "form", "")
	}
	if r.enableImages != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableImages", r.enableImages, "form", "")
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
	if r.enableUserData != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "enableUserData", r.enableUserData, "form", "")
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
