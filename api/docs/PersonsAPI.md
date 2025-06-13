# \PersonsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPerson**](PersonsAPI.md#GetPerson) | **Get** /Persons/{name} | Get person by name.
[**GetPersons**](PersonsAPI.md#GetPersons) | **Get** /Persons | Gets all persons.



## GetPerson

> BaseItemDto GetPerson(ctx, name).UserId(userId).Execute()

Get person by name.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sj14/jellyfin-go/api"
)

func main() {
	name := "name_example" // string | Person name.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. Filter by user id, and attach user data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonsAPI.GetPerson(context.Background(), name).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonsAPI.GetPerson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPerson`: BaseItemDto
	fmt.Fprintf(os.Stdout, "Response from `PersonsAPI.GetPerson`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Person name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **string** | Optional. Filter by user id, and attach user data. | 

### Return type

[**BaseItemDto**](BaseItemDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersons

> BaseItemDtoQueryResult GetPersons(ctx).Limit(limit).SearchTerm(searchTerm).Fields(fields).Filters(filters).IsFavorite(isFavorite).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).ExcludePersonTypes(excludePersonTypes).PersonTypes(personTypes).AppearsInItemId(appearsInItemId).UserId(userId).EnableImages(enableImages).Execute()

Gets all persons.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/sj14/jellyfin-go/api"
)

func main() {
	limit := int32(56) // int32 | Optional. The maximum number of records to return. (optional)
	searchTerm := "searchTerm_example" // string | The search term. (optional)
	fields := []openapiclient.ItemFields{openapiclient.ItemFields("AirTime")} // []ItemFields | Optional. Specify additional fields of information to return in the output. (optional)
	filters := []openapiclient.ItemFilter{openapiclient.ItemFilter("IsFolder")} // []ItemFilter | Optional. Specify additional filters to apply. (optional)
	isFavorite := true // bool | Optional filter by items that are marked as favorite, or not. userId is required. (optional)
	enableUserData := true // bool | Optional, include user data. (optional)
	imageTypeLimit := int32(56) // int32 | Optional, the max number of images to return, per image type. (optional)
	enableImageTypes := []openapiclient.ImageType{openapiclient.ImageType("Primary")} // []ImageType | Optional. The image types to include in the output. (optional)
	excludePersonTypes := []string{"Inner_example"} // []string | Optional. If specified results will be filtered to exclude those containing the specified PersonType. Allows multiple, comma-delimited. (optional)
	personTypes := []string{"Inner_example"} // []string | Optional. If specified results will be filtered to include only those containing the specified PersonType. Allows multiple, comma-delimited. (optional)
	appearsInItemId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Optional. If specified, person results will be filtered on items related to said persons. (optional)
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | User id. (optional)
	enableImages := true // bool | Optional, include image information in output. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PersonsAPI.GetPersons(context.Background()).Limit(limit).SearchTerm(searchTerm).Fields(fields).Filters(filters).IsFavorite(isFavorite).EnableUserData(enableUserData).ImageTypeLimit(imageTypeLimit).EnableImageTypes(enableImageTypes).ExcludePersonTypes(excludePersonTypes).PersonTypes(personTypes).AppearsInItemId(appearsInItemId).UserId(userId).EnableImages(enableImages).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PersonsAPI.GetPersons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPersons`: BaseItemDtoQueryResult
	fmt.Fprintf(os.Stdout, "Response from `PersonsAPI.GetPersons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Optional. The maximum number of records to return. | 
 **searchTerm** | **string** | The search term. | 
 **fields** | [**[]ItemFields**](ItemFields.md) | Optional. Specify additional fields of information to return in the output. | 
 **filters** | [**[]ItemFilter**](ItemFilter.md) | Optional. Specify additional filters to apply. | 
 **isFavorite** | **bool** | Optional filter by items that are marked as favorite, or not. userId is required. | 
 **enableUserData** | **bool** | Optional, include user data. | 
 **imageTypeLimit** | **int32** | Optional, the max number of images to return, per image type. | 
 **enableImageTypes** | [**[]ImageType**](ImageType.md) | Optional. The image types to include in the output. | 
 **excludePersonTypes** | **[]string** | Optional. If specified results will be filtered to exclude those containing the specified PersonType. Allows multiple, comma-delimited. | 
 **personTypes** | **[]string** | Optional. If specified results will be filtered to include only those containing the specified PersonType. Allows multiple, comma-delimited. | 
 **appearsInItemId** | **string** | Optional. If specified, person results will be filtered on items related to said persons. | 
 **userId** | **string** | User id. | 
 **enableImages** | **bool** | Optional, include image information in output. | [default to true]

### Return type

[**BaseItemDtoQueryResult**](BaseItemDtoQueryResult.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

