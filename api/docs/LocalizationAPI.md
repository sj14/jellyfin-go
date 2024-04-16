# \LocalizationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCountries**](LocalizationAPI.md#GetCountries) | **Get** /Localization/Countries | Gets known countries.
[**GetCultures**](LocalizationAPI.md#GetCultures) | **Get** /Localization/Cultures | Gets known cultures.
[**GetLocalizationOptions**](LocalizationAPI.md#GetLocalizationOptions) | **Get** /Localization/Options | Gets localization options.
[**GetParentalRatings**](LocalizationAPI.md#GetParentalRatings) | **Get** /Localization/ParentalRatings | Gets known parental ratings.



## GetCountries

> []CountryInfo GetCountries(ctx).Execute()

Gets known countries.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalizationAPI.GetCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationAPI.GetCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountries`: []CountryInfo
	fmt.Fprintf(os.Stdout, "Response from `LocalizationAPI.GetCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


### Return type

[**[]CountryInfo**](CountryInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCultures

> []CultureDto GetCultures(ctx).Execute()

Gets known cultures.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalizationAPI.GetCultures(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationAPI.GetCultures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCultures`: []CultureDto
	fmt.Fprintf(os.Stdout, "Response from `LocalizationAPI.GetCultures`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCulturesRequest struct via the builder pattern


### Return type

[**[]CultureDto**](CultureDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalizationOptions

> []LocalizationOption GetLocalizationOptions(ctx).Execute()

Gets localization options.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalizationAPI.GetLocalizationOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationAPI.GetLocalizationOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocalizationOptions`: []LocalizationOption
	fmt.Fprintf(os.Stdout, "Response from `LocalizationAPI.GetLocalizationOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalizationOptionsRequest struct via the builder pattern


### Return type

[**[]LocalizationOption**](LocalizationOption.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParentalRatings

> []ParentalRating GetParentalRatings(ctx).Execute()

Gets known parental ratings.

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocalizationAPI.GetParentalRatings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocalizationAPI.GetParentalRatings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParentalRatings`: []ParentalRating
	fmt.Fprintf(os.Stdout, "Response from `LocalizationAPI.GetParentalRatings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetParentalRatingsRequest struct via the builder pattern


### Return type

[**[]ParentalRating**](ParentalRating.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

