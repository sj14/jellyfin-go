# \SessionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToSession**](SessionAPI.md#AddUserToSession) | **Post** /Sessions/{sessionId}/User/{userId} | Adds an additional user to a session.
[**DisplayContent**](SessionAPI.md#DisplayContent) | **Post** /Sessions/{sessionId}/Viewing | Instructs a session to browse to an item or view.
[**GetAuthProviders**](SessionAPI.md#GetAuthProviders) | **Get** /Auth/Providers | Get all auth providers.
[**GetPasswordResetProviders**](SessionAPI.md#GetPasswordResetProviders) | **Get** /Auth/PasswordResetProviders | Get all password reset providers.
[**GetSessions**](SessionAPI.md#GetSessions) | **Get** /Sessions | Gets a list of sessions.
[**Play**](SessionAPI.md#Play) | **Post** /Sessions/{sessionId}/Playing | Instructs a session to play an item.
[**PostCapabilities**](SessionAPI.md#PostCapabilities) | **Post** /Sessions/Capabilities | Updates capabilities for a device.
[**PostFullCapabilities**](SessionAPI.md#PostFullCapabilities) | **Post** /Sessions/Capabilities/Full | Updates capabilities for a device.
[**RemoveUserFromSession**](SessionAPI.md#RemoveUserFromSession) | **Delete** /Sessions/{sessionId}/User/{userId} | Removes an additional user from a session.
[**ReportSessionEnded**](SessionAPI.md#ReportSessionEnded) | **Post** /Sessions/Logout | Reports that a session has ended.
[**ReportViewing**](SessionAPI.md#ReportViewing) | **Post** /Sessions/Viewing | Reports that a session is viewing an item.
[**SendFullGeneralCommand**](SessionAPI.md#SendFullGeneralCommand) | **Post** /Sessions/{sessionId}/Command | Issues a full general command to a client.
[**SendGeneralCommand**](SessionAPI.md#SendGeneralCommand) | **Post** /Sessions/{sessionId}/Command/{command} | Issues a general command to a client.
[**SendMessageCommand**](SessionAPI.md#SendMessageCommand) | **Post** /Sessions/{sessionId}/Message | Issues a command to a client to display a message to the user.
[**SendPlaystateCommand**](SessionAPI.md#SendPlaystateCommand) | **Post** /Sessions/{sessionId}/Playing/{command} | Issues a playstate command to a client.
[**SendSystemCommand**](SessionAPI.md#SendSystemCommand) | **Post** /Sessions/{sessionId}/System/{command} | Issues a system command to a client.



## AddUserToSession

> AddUserToSession(ctx, sessionId, userId).Execute()

Adds an additional user to a session.

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
	sessionId := "sessionId_example" // string | The session id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionAPI.AddUserToSession(context.Background(), sessionId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.AddUserToSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The session id. | 
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisplayContent

> DisplayContent(ctx, sessionId).ItemType(itemType).ItemId(itemId).ItemName(itemName).Execute()

Instructs a session to browse to an item or view.

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
	sessionId := "sessionId_example" // string | The session Id.
	itemType := "itemType_example" // BaseItemKind | The type of item to browse to.
	itemId := "itemId_example" // string | The Id of the item.
	itemName := "itemName_example" // string | The name of the item.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionAPI.DisplayContent(context.Background(), sessionId).ItemType(itemType).ItemId(itemId).ItemName(itemName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.DisplayContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The session Id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisplayContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **itemType** | **BaseItemKind** | The type of item to browse to. | 
 **itemId** | **string** | The Id of the item. | 
 **itemName** | **string** | The name of the item. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthProviders

> []NameIdPair GetAuthProviders(ctx).Execute()

Get all auth providers.

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
	resp, r, err := apiClient.SessionAPI.GetAuthProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.GetAuthProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuthProviders`: []NameIdPair
	fmt.Fprintf(os.Stdout, "Response from `SessionAPI.GetAuthProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthProvidersRequest struct via the builder pattern


### Return type

[**[]NameIdPair**](NameIdPair.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPasswordResetProviders

> []NameIdPair GetPasswordResetProviders(ctx).Execute()

Get all password reset providers.

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
	resp, r, err := apiClient.SessionAPI.GetPasswordResetProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.GetPasswordResetProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPasswordResetProviders`: []NameIdPair
	fmt.Fprintf(os.Stdout, "Response from `SessionAPI.GetPasswordResetProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordResetProvidersRequest struct via the builder pattern


### Return type

[**[]NameIdPair**](NameIdPair.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessions

> []SessionInfo GetSessions(ctx).ControllableByUserId(controllableByUserId).DeviceId(deviceId).ActiveWithinSeconds(activeWithinSeconds).Execute()

Gets a list of sessions.

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
	controllableByUserId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Filter by sessions that a given user is allowed to remote control. (optional)
	deviceId := "deviceId_example" // string | Filter by device Id. (optional)
	activeWithinSeconds := int32(56) // int32 | Optional. Filter by sessions that were active in the last n seconds. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionAPI.GetSessions(context.Background()).ControllableByUserId(controllableByUserId).DeviceId(deviceId).ActiveWithinSeconds(activeWithinSeconds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.GetSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSessions`: []SessionInfo
	fmt.Fprintf(os.Stdout, "Response from `SessionAPI.GetSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **controllableByUserId** | **string** | Filter by sessions that a given user is allowed to remote control. | 
 **deviceId** | **string** | Filter by device Id. | 
 **activeWithinSeconds** | **int32** | Optional. Filter by sessions that were active in the last n seconds. | 

### Return type

[**[]SessionInfo**](SessionInfo.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Play

> Play(ctx, sessionId).PlayCommand(playCommand).ItemIds(itemIds).StartPositionTicks(startPositionTicks).MediaSourceId(mediaSourceId).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).StartIndex(startIndex).Execute()

Instructs a session to play an item.

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
	sessionId := "sessionId_example" // string | The session id.
	playCommand := "playCommand_example" // PlayCommand | The type of play command to issue (PlayNow, PlayNext, PlayLast). Clients who have not yet implemented play next and play last may play now.
	itemIds := []string{"Inner_example"} // []string | The ids of the items to play, comma delimited.
	startPositionTicks := int64(789) // int64 | The starting position of the first item. (optional)
	mediaSourceId := "mediaSourceId_example" // string | Optional. The media source id. (optional)
	audioStreamIndex := int32(56) // int32 | Optional. The index of the audio stream to play. (optional)
	subtitleStreamIndex := int32(56) // int32 | Optional. The index of the subtitle stream to play. (optional)
	startIndex := int32(56) // int32 | Optional. The start index. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionAPI.Play(context.Background(), sessionId).PlayCommand(playCommand).ItemIds(itemIds).StartPositionTicks(startPositionTicks).MediaSourceId(mediaSourceId).AudioStreamIndex(audioStreamIndex).SubtitleStreamIndex(subtitleStreamIndex).StartIndex(startIndex).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.Play``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The session id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **playCommand** | **PlayCommand** | The type of play command to issue (PlayNow, PlayNext, PlayLast). Clients who have not yet implemented play next and play last may play now. | 
 **itemIds** | **[]string** | The ids of the items to play, comma delimited. | 
 **startPositionTicks** | **int64** | The starting position of the first item. | 
 **mediaSourceId** | **string** | Optional. The media source id. | 
 **audioStreamIndex** | **int32** | Optional. The index of the audio stream to play. | 
 **subtitleStreamIndex** | **int32** | Optional. The index of the subtitle stream to play. | 
 **startIndex** | **int32** | Optional. The start index. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCapabilities

> PostCapabilities(ctx).Id(id).PlayableMediaTypes(playableMediaTypes).SupportedCommands(supportedCommands).SupportsMediaControl(supportsMediaControl).SupportsPersistentIdentifier(supportsPersistentIdentifier).Execute()

Updates capabilities for a device.

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
	id := "id_example" // string | The session id. (optional)
	playableMediaTypes := []openapiclient.MediaType{openapiclient.MediaType("Unknown")} // []MediaType | A list of playable media types, comma delimited. Audio, Video, Book, Photo. (optional)
	supportedCommands := []openapiclient.GeneralCommandType{openapiclient.GeneralCommandType("MoveUp")} // []GeneralCommandType | A list of supported remote control commands, comma delimited. (optional)
	supportsMediaControl := true // bool | Determines whether media can be played remotely.. (optional) (default to false)
	supportsPersistentIdentifier := true // bool | Determines whether the device supports a unique identifier. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionAPI.PostCapabilities(context.Background()).Id(id).PlayableMediaTypes(playableMediaTypes).SupportedCommands(supportedCommands).SupportsMediaControl(supportsMediaControl).SupportsPersistentIdentifier(supportsPersistentIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.PostCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | The session id. | 
 **playableMediaTypes** | [**[]MediaType**](MediaType.md) | A list of playable media types, comma delimited. Audio, Video, Book, Photo. | 
 **supportedCommands** | [**[]GeneralCommandType**](GeneralCommandType.md) | A list of supported remote control commands, comma delimited. | 
 **supportsMediaControl** | **bool** | Determines whether media can be played remotely.. | [default to false]
 **supportsPersistentIdentifier** | **bool** | Determines whether the device supports a unique identifier. | [default to true]

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFullCapabilities

> PostFullCapabilities(ctx).ClientCapabilitiesDto(clientCapabilitiesDto).Id(id).Execute()

Updates capabilities for a device.

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
	clientCapabilitiesDto := *openapiclient.NewClientCapabilitiesDto() // ClientCapabilitiesDto | The MediaBrowser.Model.Session.ClientCapabilities.
	id := "id_example" // string | The session id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionAPI.PostFullCapabilities(context.Background()).ClientCapabilitiesDto(clientCapabilitiesDto).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.PostFullCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostFullCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientCapabilitiesDto** | [**ClientCapabilitiesDto**](ClientCapabilitiesDto.md) | The MediaBrowser.Model.Session.ClientCapabilities. | 
 **id** | **string** | The session id. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserFromSession

> RemoveUserFromSession(ctx, sessionId, userId).Execute()

Removes an additional user from a session.

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
	sessionId := "sessionId_example" // string | The session id.
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The user id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionAPI.RemoveUserFromSession(context.Background(), sessionId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.RemoveUserFromSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The session id. | 
**userId** | **string** | The user id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportSessionEnded

> ReportSessionEnded(ctx).Execute()

Reports that a session has ended.

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
	r, err := apiClient.SessionAPI.ReportSessionEnded(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ReportSessionEnded``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReportSessionEndedRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportViewing

> ReportViewing(ctx).ItemId(itemId).SessionId(sessionId).Execute()

Reports that a session is viewing an item.

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
	itemId := "itemId_example" // string | The item id.
	sessionId := "sessionId_example" // string | The session id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionAPI.ReportViewing(context.Background()).ItemId(itemId).SessionId(sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.ReportViewing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportViewingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemId** | **string** | The item id. | 
 **sessionId** | **string** | The session id. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendFullGeneralCommand

> SendFullGeneralCommand(ctx, sessionId).GeneralCommand(generalCommand).Execute()

Issues a full general command to a client.

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
	sessionId := "sessionId_example" // string | The session id.
	generalCommand := *openapiclient.NewGeneralCommand() // GeneralCommand | The MediaBrowser.Model.Session.GeneralCommand.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionAPI.SendFullGeneralCommand(context.Background(), sessionId).GeneralCommand(generalCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.SendFullGeneralCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The session id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendFullGeneralCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generalCommand** | [**GeneralCommand**](GeneralCommand.md) | The MediaBrowser.Model.Session.GeneralCommand. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendGeneralCommand

> SendGeneralCommand(ctx, sessionId, command).Execute()

Issues a general command to a client.

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
	sessionId := "sessionId_example" // string | The session id.
	command := "command_example" // GeneralCommandType | The command to send.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionAPI.SendGeneralCommand(context.Background(), sessionId, command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.SendGeneralCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The session id. | 
**command** | **GeneralCommandType** | The command to send. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendGeneralCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessageCommand

> SendMessageCommand(ctx, sessionId).MessageCommand(messageCommand).Execute()

Issues a command to a client to display a message to the user.

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
	sessionId := "sessionId_example" // string | The session id.
	messageCommand := *openapiclient.NewMessageCommand("Text_example") // MessageCommand | The MediaBrowser.Model.Session.MessageCommand object containing Header, Message Text, and TimeoutMs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionAPI.SendMessageCommand(context.Background(), sessionId).MessageCommand(messageCommand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.SendMessageCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The session id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendMessageCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messageCommand** | [**MessageCommand**](MessageCommand.md) | The MediaBrowser.Model.Session.MessageCommand object containing Header, Message Text, and TimeoutMs. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendPlaystateCommand

> SendPlaystateCommand(ctx, sessionId, command).SeekPositionTicks(seekPositionTicks).ControllingUserId(controllingUserId).Execute()

Issues a playstate command to a client.

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
	sessionId := "sessionId_example" // string | The session id.
	command := "command_example" // PlaystateCommand | The MediaBrowser.Model.Session.PlaystateCommand.
	seekPositionTicks := int64(789) // int64 | The optional position ticks. (optional)
	controllingUserId := "controllingUserId_example" // string | The optional controlling user id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionAPI.SendPlaystateCommand(context.Background(), sessionId, command).SeekPositionTicks(seekPositionTicks).ControllingUserId(controllingUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.SendPlaystateCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The session id. | 
**command** | **PlaystateCommand** | The MediaBrowser.Model.Session.PlaystateCommand. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendPlaystateCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **seekPositionTicks** | **int64** | The optional position ticks. | 
 **controllingUserId** | **string** | The optional controlling user id. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSystemCommand

> SendSystemCommand(ctx, sessionId, command).Execute()

Issues a system command to a client.

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
	sessionId := "sessionId_example" // string | The session id.
	command := "command_example" // GeneralCommandType | The command to send.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SessionAPI.SendSystemCommand(context.Background(), sessionId, command).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.SendSystemCommand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | The session id. | 
**command** | **GeneralCommandType** | The command to send. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendSystemCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

