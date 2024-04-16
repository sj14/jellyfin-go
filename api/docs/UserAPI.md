# \UserAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateUser**](UserAPI.md#AuthenticateUser) | **Post** /Users/{userId}/Authenticate | Authenticates a user.
[**AuthenticateUserByName**](UserAPI.md#AuthenticateUserByName) | **Post** /Users/AuthenticateByName | Authenticates a user by name.
[**AuthenticateWithQuickConnect**](UserAPI.md#AuthenticateWithQuickConnect) | **Post** /Users/AuthenticateWithQuickConnect | Authenticates a user with quick connect.
[**CreateUserByName**](UserAPI.md#CreateUserByName) | **Post** /Users/New | Creates a user.
[**DeleteUser**](UserAPI.md#DeleteUser) | **Delete** /Users/{userId} | Deletes a user.
[**ForgotPassword**](UserAPI.md#ForgotPassword) | **Post** /Users/ForgotPassword | Initiates the forgot password process for a local user.
[**ForgotPasswordPin**](UserAPI.md#ForgotPasswordPin) | **Post** /Users/ForgotPassword/Pin | Redeems a forgot password pin.
[**GetCurrentUser**](UserAPI.md#GetCurrentUser) | **Get** /Users/Me | Gets the user based on auth token.
[**GetPublicUsers**](UserAPI.md#GetPublicUsers) | **Get** /Users/Public | Gets a list of publicly visible users for display on a login screen.
[**GetUserById**](UserAPI.md#GetUserById) | **Get** /Users/{userId} | Gets a user by Id.
[**GetUsers**](UserAPI.md#GetUsers) | **Get** /Users | Gets a list of users.
[**UpdateUser**](UserAPI.md#UpdateUser) | **Post** /Users/{userId} | Updates a user.
[**UpdateUserConfiguration**](UserAPI.md#UpdateUserConfiguration) | **Post** /Users/{userId}/Configuration | Updates a user configuration.
[**UpdateUserEasyPassword**](UserAPI.md#UpdateUserEasyPassword) | **Post** /Users/{userId}/EasyPassword | Updates a user&#39;s easy password.
[**UpdateUserPassword**](UserAPI.md#UpdateUserPassword) | **Post** /Users/{userId}/Password | Updates a user&#39;s password.
[**UpdateUserPolicy**](UserAPI.md#UpdateUserPolicy) | **Post** /Users/{userId}/Policy | Updates a user policy.



## AuthenticateUser

> AuthenticationResult AuthenticateUser(ctx, userId).Pw(pw).Password(password).Execute()

Authenticates a user.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.
	pw := "pw_example" // string | The password as plain text.
	password := "password_example" // string | The password sha1-hash. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.AuthenticateUser(context.Background(), userId).Pw(pw).Password(password).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.AuthenticateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticateUser`: AuthenticationResult
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.AuthenticateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pw** | **string** | The password as plain text. | 
 **password** | **string** | The password sha1-hash. | 

### Return type

[**AuthenticationResult**](AuthenticationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticateUserByName

> AuthenticationResult AuthenticateUserByName(ctx).AuthenticateUserByName(authenticateUserByName).Execute()

Authenticates a user by name.

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
	authenticateUserByName := *openapiclient.NewAuthenticateUserByName() // AuthenticateUserByName | The M:Jellyfin.Api.Controllers.UserController.AuthenticateUserByName(Jellyfin.Api.Models.UserDtos.AuthenticateUserByName) request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.AuthenticateUserByName(context.Background()).AuthenticateUserByName(authenticateUserByName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.AuthenticateUserByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticateUserByName`: AuthenticationResult
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.AuthenticateUserByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateUserByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticateUserByName** | [**AuthenticateUserByName**](AuthenticateUserByName.md) | The M:Jellyfin.Api.Controllers.UserController.AuthenticateUserByName(Jellyfin.Api.Models.UserDtos.AuthenticateUserByName) request. | 

### Return type

[**AuthenticationResult**](AuthenticationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthenticateWithQuickConnect

> AuthenticationResult AuthenticateWithQuickConnect(ctx).QuickConnectDto(quickConnectDto).Execute()

Authenticates a user with quick connect.

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
	quickConnectDto := *openapiclient.NewQuickConnectDto("Secret_example") // QuickConnectDto | The Jellyfin.Api.Models.UserDtos.QuickConnectDto request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.AuthenticateWithQuickConnect(context.Background()).QuickConnectDto(quickConnectDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.AuthenticateWithQuickConnect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AuthenticateWithQuickConnect`: AuthenticationResult
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.AuthenticateWithQuickConnect`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateWithQuickConnectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quickConnectDto** | [**QuickConnectDto**](QuickConnectDto.md) | The Jellyfin.Api.Models.UserDtos.QuickConnectDto request. | 

### Return type

[**AuthenticationResult**](AuthenticationResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserByName

> UserDto CreateUserByName(ctx).CreateUserByName(createUserByName).Execute()

Creates a user.

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
	createUserByName := *openapiclient.NewCreateUserByName() // CreateUserByName | The create user by name request body.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CreateUserByName(context.Background()).CreateUserByName(createUserByName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CreateUserByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserByName`: UserDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CreateUserByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserByName** | [**CreateUserByName**](CreateUserByName.md) | The create user by name request body. | 

### Return type

[**UserDto**](UserDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId).Execute()

Deletes a user.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.DeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForgotPassword

> ForgotPasswordResult ForgotPassword(ctx).ForgotPasswordDto(forgotPasswordDto).Execute()

Initiates the forgot password process for a local user.

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
	forgotPasswordDto := *openapiclient.NewForgotPasswordDto("EnteredUsername_example") // ForgotPasswordDto | The forgot password request containing the entered username.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.ForgotPassword(context.Background()).ForgotPasswordDto(forgotPasswordDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ForgotPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForgotPassword`: ForgotPasswordResult
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ForgotPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForgotPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forgotPasswordDto** | [**ForgotPasswordDto**](ForgotPasswordDto.md) | The forgot password request containing the entered username. | 

### Return type

[**ForgotPasswordResult**](ForgotPasswordResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ForgotPasswordPin

> PinRedeemResult ForgotPasswordPin(ctx).ForgotPasswordPinDto(forgotPasswordPinDto).Execute()

Redeems a forgot password pin.

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
	forgotPasswordPinDto := *openapiclient.NewForgotPasswordPinDto("Pin_example") // ForgotPasswordPinDto | The forgot password pin request containing the entered pin.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.ForgotPasswordPin(context.Background()).ForgotPasswordPinDto(forgotPasswordPinDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ForgotPasswordPin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ForgotPasswordPin`: PinRedeemResult
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ForgotPasswordPin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiForgotPasswordPinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forgotPasswordPinDto** | [**ForgotPasswordPinDto**](ForgotPasswordPinDto.md) | The forgot password pin request containing the entered pin. | 

### Return type

[**PinRedeemResult**](PinRedeemResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUser

> UserDto GetCurrentUser(ctx).Execute()

Gets the user based on auth token.

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
	resp, r, err := apiClient.UserAPI.GetCurrentUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUser`: UserDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserRequest struct via the builder pattern


### Return type

[**UserDto**](UserDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicUsers

> []UserDto GetPublicUsers(ctx).Execute()

Gets a list of publicly visible users for display on a login screen.

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
	resp, r, err := apiClient.UserAPI.GetPublicUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetPublicUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicUsers`: []UserDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetPublicUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicUsersRequest struct via the builder pattern


### Return type

[**[]UserDto**](UserDto.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserById

> UserDto GetUserById(ctx, userId).Execute()

Gets a user by Id.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserById(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserById`: UserDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserDto**](UserDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> []UserDto GetUsers(ctx).IsHidden(isHidden).IsDisabled(isDisabled).Execute()

Gets a list of users.

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
	isHidden := true // bool | Optional filter by IsHidden=true or false. (optional)
	isDisabled := true // bool | Optional filter by IsDisabled=true or false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUsers(context.Background()).IsHidden(isHidden).IsDisabled(isDisabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: []UserDto
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isHidden** | **bool** | Optional filter by IsHidden&#x3D;true or false. | 
 **isDisabled** | **bool** | Optional filter by IsDisabled&#x3D;true or false. | 

### Return type

[**[]UserDto**](UserDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UpdateUser(ctx, userId).UserDto(userDto).Execute()

Updates a user.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.
	userDto := *openapiclient.NewUserDto() // UserDto | The updated user model.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateUser(context.Background(), userId).UserDto(userDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userDto** | [**UserDto**](UserDto.md) | The updated user model. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserConfiguration

> UpdateUserConfiguration(ctx, userId).UserConfiguration(userConfiguration).Execute()

Updates a user configuration.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.
	userConfiguration := *openapiclient.NewUserConfiguration() // UserConfiguration | The new user configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateUserConfiguration(context.Background(), userId).UserConfiguration(userConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUserConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userConfiguration** | [**UserConfiguration**](UserConfiguration.md) | The new user configuration. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserEasyPassword

> UpdateUserEasyPassword(ctx, userId).UpdateUserEasyPassword(updateUserEasyPassword).Execute()

Updates a user's easy password.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.
	updateUserEasyPassword := *openapiclient.NewUpdateUserEasyPassword() // UpdateUserEasyPassword | The M:Jellyfin.Api.Controllers.UserController.UpdateUserEasyPassword(System.Guid,Jellyfin.Api.Models.UserDtos.UpdateUserEasyPassword) request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateUserEasyPassword(context.Background(), userId).UpdateUserEasyPassword(updateUserEasyPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUserEasyPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserEasyPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserEasyPassword** | [**UpdateUserEasyPassword**](UpdateUserEasyPassword.md) | The M:Jellyfin.Api.Controllers.UserController.UpdateUserEasyPassword(System.Guid,Jellyfin.Api.Models.UserDtos.UpdateUserEasyPassword) request. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPassword

> UpdateUserPassword(ctx, userId).UpdateUserPassword(updateUserPassword).Execute()

Updates a user's password.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.
	updateUserPassword := *openapiclient.NewUpdateUserPassword() // UpdateUserPassword | The M:Jellyfin.Api.Controllers.UserController.UpdateUserPassword(System.Guid,Jellyfin.Api.Models.UserDtos.UpdateUserPassword) request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateUserPassword(context.Background(), userId).UpdateUserPassword(updateUserPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUserPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserPassword** | [**UpdateUserPassword**](UpdateUserPassword.md) | The M:Jellyfin.Api.Controllers.UserController.UpdateUserPassword(System.Guid,Jellyfin.Api.Models.UserDtos.UpdateUserPassword) request. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPolicy

> UpdateUserPolicy(ctx, userId).UserPolicy(userPolicy).Execute()

Updates a user policy.

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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.
	userPolicy := *openapiclient.NewUserPolicy() // UserPolicy | The new user policy.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateUserPolicy(context.Background(), userId).UserPolicy(userPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUserPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userPolicy** | [**UserPolicy**](UserPolicy.md) | The new user policy. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

