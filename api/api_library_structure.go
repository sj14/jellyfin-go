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
	"reflect"
)


type LibraryStructureAPI interface {

	/*
	AddMediaPath Add a media path to a library.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddMediaPathRequest
	*/
	AddMediaPath(ctx context.Context) ApiAddMediaPathRequest

	// AddMediaPathExecute executes the request
	AddMediaPathExecute(r ApiAddMediaPathRequest) (*http.Response, error)

	/*
	AddVirtualFolder Adds a virtual folder.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAddVirtualFolderRequest
	*/
	AddVirtualFolder(ctx context.Context) ApiAddVirtualFolderRequest

	// AddVirtualFolderExecute executes the request
	AddVirtualFolderExecute(r ApiAddVirtualFolderRequest) (*http.Response, error)

	/*
	GetVirtualFolders Gets all virtual folders.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetVirtualFoldersRequest
	*/
	GetVirtualFolders(ctx context.Context) ApiGetVirtualFoldersRequest

	// GetVirtualFoldersExecute executes the request
	//  @return []VirtualFolderInfo
	GetVirtualFoldersExecute(r ApiGetVirtualFoldersRequest) ([]VirtualFolderInfo, *http.Response, error)

	/*
	RemoveMediaPath Remove a media path.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRemoveMediaPathRequest
	*/
	RemoveMediaPath(ctx context.Context) ApiRemoveMediaPathRequest

	// RemoveMediaPathExecute executes the request
	RemoveMediaPathExecute(r ApiRemoveMediaPathRequest) (*http.Response, error)

	/*
	RemoveVirtualFolder Removes a virtual folder.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRemoveVirtualFolderRequest
	*/
	RemoveVirtualFolder(ctx context.Context) ApiRemoveVirtualFolderRequest

	// RemoveVirtualFolderExecute executes the request
	RemoveVirtualFolderExecute(r ApiRemoveVirtualFolderRequest) (*http.Response, error)

	/*
	RenameVirtualFolder Renames a virtual folder.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRenameVirtualFolderRequest
	*/
	RenameVirtualFolder(ctx context.Context) ApiRenameVirtualFolderRequest

	// RenameVirtualFolderExecute executes the request
	RenameVirtualFolderExecute(r ApiRenameVirtualFolderRequest) (*http.Response, error)

	/*
	UpdateLibraryOptions Update library options.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdateLibraryOptionsRequest
	*/
	UpdateLibraryOptions(ctx context.Context) ApiUpdateLibraryOptionsRequest

	// UpdateLibraryOptionsExecute executes the request
	UpdateLibraryOptionsExecute(r ApiUpdateLibraryOptionsRequest) (*http.Response, error)

	/*
	UpdateMediaPath Updates a media path.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiUpdateMediaPathRequest
	*/
	UpdateMediaPath(ctx context.Context) ApiUpdateMediaPathRequest

	// UpdateMediaPathExecute executes the request
	UpdateMediaPathExecute(r ApiUpdateMediaPathRequest) (*http.Response, error)
}

// LibraryStructureAPIService LibraryStructureAPI service
type LibraryStructureAPIService service

type ApiAddMediaPathRequest struct {
	ctx context.Context
	ApiService LibraryStructureAPI
	mediaPathDto *MediaPathDto
	refreshLibrary *bool
}

// The media path dto.
func (r ApiAddMediaPathRequest) MediaPathDto(mediaPathDto MediaPathDto) ApiAddMediaPathRequest {
	r.mediaPathDto = &mediaPathDto
	return r
}

// Whether to refresh the library.
func (r ApiAddMediaPathRequest) RefreshLibrary(refreshLibrary bool) ApiAddMediaPathRequest {
	r.refreshLibrary = &refreshLibrary
	return r
}

func (r ApiAddMediaPathRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddMediaPathExecute(r)
}

/*
AddMediaPath Add a media path to a library.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddMediaPathRequest
*/
func (a *LibraryStructureAPIService) AddMediaPath(ctx context.Context) ApiAddMediaPathRequest {
	return ApiAddMediaPathRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *LibraryStructureAPIService) AddMediaPathExecute(r ApiAddMediaPathRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LibraryStructureAPIService.AddMediaPath")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Library/VirtualFolders/Paths"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mediaPathDto == nil {
		return nil, reportError("mediaPathDto is required and must be specified")
	}

	if r.refreshLibrary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "refreshLibrary", r.refreshLibrary, "")
	} else {
		var defaultValue bool = false
		r.refreshLibrary = &defaultValue
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
	localVarPostBody = r.mediaPathDto
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

type ApiAddVirtualFolderRequest struct {
	ctx context.Context
	ApiService LibraryStructureAPI
	name *string
	collectionType *CollectionTypeOptions
	paths *[]string
	refreshLibrary *bool
	addVirtualFolderDto *AddVirtualFolderDto
}

// The name of the virtual folder.
func (r ApiAddVirtualFolderRequest) Name(name string) ApiAddVirtualFolderRequest {
	r.name = &name
	return r
}

// The type of the collection.
func (r ApiAddVirtualFolderRequest) CollectionType(collectionType CollectionTypeOptions) ApiAddVirtualFolderRequest {
	r.collectionType = &collectionType
	return r
}

// The paths of the virtual folder.
func (r ApiAddVirtualFolderRequest) Paths(paths []string) ApiAddVirtualFolderRequest {
	r.paths = &paths
	return r
}

// Whether to refresh the library.
func (r ApiAddVirtualFolderRequest) RefreshLibrary(refreshLibrary bool) ApiAddVirtualFolderRequest {
	r.refreshLibrary = &refreshLibrary
	return r
}

// The library options.
func (r ApiAddVirtualFolderRequest) AddVirtualFolderDto(addVirtualFolderDto AddVirtualFolderDto) ApiAddVirtualFolderRequest {
	r.addVirtualFolderDto = &addVirtualFolderDto
	return r
}

func (r ApiAddVirtualFolderRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddVirtualFolderExecute(r)
}

/*
AddVirtualFolder Adds a virtual folder.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddVirtualFolderRequest
*/
func (a *LibraryStructureAPIService) AddVirtualFolder(ctx context.Context) ApiAddVirtualFolderRequest {
	return ApiAddVirtualFolderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *LibraryStructureAPIService) AddVirtualFolderExecute(r ApiAddVirtualFolderRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LibraryStructureAPIService.AddVirtualFolder")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Library/VirtualFolders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.collectionType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "collectionType", r.collectionType, "")
	}
	if r.paths != nil {
		t := *r.paths
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "paths", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "paths", t, "multi")
		}
	}
	if r.refreshLibrary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "refreshLibrary", r.refreshLibrary, "")
	} else {
		var defaultValue bool = false
		r.refreshLibrary = &defaultValue
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
	localVarPostBody = r.addVirtualFolderDto
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

type ApiGetVirtualFoldersRequest struct {
	ctx context.Context
	ApiService LibraryStructureAPI
}

func (r ApiGetVirtualFoldersRequest) Execute() ([]VirtualFolderInfo, *http.Response, error) {
	return r.ApiService.GetVirtualFoldersExecute(r)
}

/*
GetVirtualFolders Gets all virtual folders.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetVirtualFoldersRequest
*/
func (a *LibraryStructureAPIService) GetVirtualFolders(ctx context.Context) ApiGetVirtualFoldersRequest {
	return ApiGetVirtualFoldersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []VirtualFolderInfo
func (a *LibraryStructureAPIService) GetVirtualFoldersExecute(r ApiGetVirtualFoldersRequest) ([]VirtualFolderInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []VirtualFolderInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LibraryStructureAPIService.GetVirtualFolders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Library/VirtualFolders"

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

type ApiRemoveMediaPathRequest struct {
	ctx context.Context
	ApiService LibraryStructureAPI
	name *string
	path *string
	refreshLibrary *bool
}

// The name of the library.
func (r ApiRemoveMediaPathRequest) Name(name string) ApiRemoveMediaPathRequest {
	r.name = &name
	return r
}

// The path to remove.
func (r ApiRemoveMediaPathRequest) Path(path string) ApiRemoveMediaPathRequest {
	r.path = &path
	return r
}

// Whether to refresh the library.
func (r ApiRemoveMediaPathRequest) RefreshLibrary(refreshLibrary bool) ApiRemoveMediaPathRequest {
	r.refreshLibrary = &refreshLibrary
	return r
}

func (r ApiRemoveMediaPathRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveMediaPathExecute(r)
}

/*
RemoveMediaPath Remove a media path.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRemoveMediaPathRequest
*/
func (a *LibraryStructureAPIService) RemoveMediaPath(ctx context.Context) ApiRemoveMediaPathRequest {
	return ApiRemoveMediaPathRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *LibraryStructureAPIService) RemoveMediaPathExecute(r ApiRemoveMediaPathRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LibraryStructureAPIService.RemoveMediaPath")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Library/VirtualFolders/Paths"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.path != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "path", r.path, "")
	}
	if r.refreshLibrary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "refreshLibrary", r.refreshLibrary, "")
	} else {
		var defaultValue bool = false
		r.refreshLibrary = &defaultValue
	}
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

type ApiRemoveVirtualFolderRequest struct {
	ctx context.Context
	ApiService LibraryStructureAPI
	name *string
	refreshLibrary *bool
}

// The name of the folder.
func (r ApiRemoveVirtualFolderRequest) Name(name string) ApiRemoveVirtualFolderRequest {
	r.name = &name
	return r
}

// Whether to refresh the library.
func (r ApiRemoveVirtualFolderRequest) RefreshLibrary(refreshLibrary bool) ApiRemoveVirtualFolderRequest {
	r.refreshLibrary = &refreshLibrary
	return r
}

func (r ApiRemoveVirtualFolderRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveVirtualFolderExecute(r)
}

/*
RemoveVirtualFolder Removes a virtual folder.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRemoveVirtualFolderRequest
*/
func (a *LibraryStructureAPIService) RemoveVirtualFolder(ctx context.Context) ApiRemoveVirtualFolderRequest {
	return ApiRemoveVirtualFolderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *LibraryStructureAPIService) RemoveVirtualFolderExecute(r ApiRemoveVirtualFolderRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LibraryStructureAPIService.RemoveVirtualFolder")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Library/VirtualFolders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.refreshLibrary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "refreshLibrary", r.refreshLibrary, "")
	} else {
		var defaultValue bool = false
		r.refreshLibrary = &defaultValue
	}
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

type ApiRenameVirtualFolderRequest struct {
	ctx context.Context
	ApiService LibraryStructureAPI
	name *string
	newName *string
	refreshLibrary *bool
}

// The name of the virtual folder.
func (r ApiRenameVirtualFolderRequest) Name(name string) ApiRenameVirtualFolderRequest {
	r.name = &name
	return r
}

// The new name.
func (r ApiRenameVirtualFolderRequest) NewName(newName string) ApiRenameVirtualFolderRequest {
	r.newName = &newName
	return r
}

// Whether to refresh the library.
func (r ApiRenameVirtualFolderRequest) RefreshLibrary(refreshLibrary bool) ApiRenameVirtualFolderRequest {
	r.refreshLibrary = &refreshLibrary
	return r
}

func (r ApiRenameVirtualFolderRequest) Execute() (*http.Response, error) {
	return r.ApiService.RenameVirtualFolderExecute(r)
}

/*
RenameVirtualFolder Renames a virtual folder.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRenameVirtualFolderRequest
*/
func (a *LibraryStructureAPIService) RenameVirtualFolder(ctx context.Context) ApiRenameVirtualFolderRequest {
	return ApiRenameVirtualFolderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *LibraryStructureAPIService) RenameVirtualFolderExecute(r ApiRenameVirtualFolderRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LibraryStructureAPIService.RenameVirtualFolder")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Library/VirtualFolders/Name"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "")
	}
	if r.newName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "newName", r.newName, "")
	}
	if r.refreshLibrary != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "refreshLibrary", r.refreshLibrary, "")
	} else {
		var defaultValue bool = false
		r.refreshLibrary = &defaultValue
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateLibraryOptionsRequest struct {
	ctx context.Context
	ApiService LibraryStructureAPI
	updateLibraryOptionsDto *UpdateLibraryOptionsDto
}

// The library name and options.
func (r ApiUpdateLibraryOptionsRequest) UpdateLibraryOptionsDto(updateLibraryOptionsDto UpdateLibraryOptionsDto) ApiUpdateLibraryOptionsRequest {
	r.updateLibraryOptionsDto = &updateLibraryOptionsDto
	return r
}

func (r ApiUpdateLibraryOptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateLibraryOptionsExecute(r)
}

/*
UpdateLibraryOptions Update library options.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateLibraryOptionsRequest
*/
func (a *LibraryStructureAPIService) UpdateLibraryOptions(ctx context.Context) ApiUpdateLibraryOptionsRequest {
	return ApiUpdateLibraryOptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *LibraryStructureAPIService) UpdateLibraryOptionsExecute(r ApiUpdateLibraryOptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LibraryStructureAPIService.UpdateLibraryOptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Library/VirtualFolders/LibraryOptions"

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/json; profile=CamelCase", "application/json; profile=PascalCase"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateLibraryOptionsDto
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateMediaPathRequest struct {
	ctx context.Context
	ApiService LibraryStructureAPI
	updateMediaPathRequestDto *UpdateMediaPathRequestDto
}

// The name of the library and path infos.
func (r ApiUpdateMediaPathRequest) UpdateMediaPathRequestDto(updateMediaPathRequestDto UpdateMediaPathRequestDto) ApiUpdateMediaPathRequest {
	r.updateMediaPathRequestDto = &updateMediaPathRequestDto
	return r
}

func (r ApiUpdateMediaPathRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateMediaPathExecute(r)
}

/*
UpdateMediaPath Updates a media path.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateMediaPathRequest
*/
func (a *LibraryStructureAPIService) UpdateMediaPath(ctx context.Context) ApiUpdateMediaPathRequest {
	return ApiUpdateMediaPathRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *LibraryStructureAPIService) UpdateMediaPathExecute(r ApiUpdateMediaPathRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LibraryStructureAPIService.UpdateMediaPath")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Library/VirtualFolders/Paths/Update"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updateMediaPathRequestDto == nil {
		return nil, reportError("updateMediaPathRequestDto is required and must be specified")
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
	localVarPostBody = r.updateMediaPathRequestDto
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
