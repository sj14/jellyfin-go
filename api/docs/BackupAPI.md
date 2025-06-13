# \BackupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBackup**](BackupAPI.md#CreateBackup) | **Post** /Backup/Create | Creates a new Backup.
[**GetBackup**](BackupAPI.md#GetBackup) | **Get** /Backup/Manifest | Gets the descriptor from an existing archive is present.
[**ListBackups**](BackupAPI.md#ListBackups) | **Get** /Backup | Gets a list of all currently present backups in the backup directory.
[**StartRestoreBackup**](BackupAPI.md#StartRestoreBackup) | **Post** /Backup/Restore | Restores to a backup by restarting the server and applying the backup.



## CreateBackup

> BackupManifestDto CreateBackup(ctx).BackupOptionsDto(backupOptionsDto).Execute()

Creates a new Backup.

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
	backupOptionsDto := *openapiclient.NewBackupOptionsDto() // BackupOptionsDto | The backup options. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.CreateBackup(context.Background()).BackupOptionsDto(backupOptionsDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.CreateBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBackup`: BackupManifestDto
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.CreateBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupOptionsDto** | [**BackupOptionsDto**](BackupOptionsDto.md) | The backup options. | 

### Return type

[**BackupManifestDto**](BackupManifestDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackup

> BackupManifestDto GetBackup(ctx).Path(path).Execute()

Gets the descriptor from an existing archive is present.

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
	path := "path_example" // string | The data to start a restore process.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.GetBackup(context.Background()).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.GetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackup`: BackupManifestDto
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.GetBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | The data to start a restore process. | 

### Return type

[**BackupManifestDto**](BackupManifestDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBackups

> []BackupManifestDto ListBackups(ctx).Execute()

Gets a list of all currently present backups in the backup directory.

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
	resp, r, err := apiClient.BackupAPI.ListBackups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.ListBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBackups`: []BackupManifestDto
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.ListBackups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBackupsRequest struct via the builder pattern


### Return type

[**[]BackupManifestDto**](BackupManifestDto.md)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartRestoreBackup

> StartRestoreBackup(ctx).BackupRestoreRequestDto(backupRestoreRequestDto).Execute()

Restores to a backup by restarting the server and applying the backup.

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
	backupRestoreRequestDto := *openapiclient.NewBackupRestoreRequestDto() // BackupRestoreRequestDto | The data to start a restore process.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupAPI.StartRestoreBackup(context.Background()).BackupRestoreRequestDto(backupRestoreRequestDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.StartRestoreBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartRestoreBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **backupRestoreRequestDto** | [**BackupRestoreRequestDto**](BackupRestoreRequestDto.md) | The data to start a restore process. | 

### Return type

 (empty response body)

### Authorization

[CustomAuthentication](../README.md#CustomAuthentication)

### HTTP request headers

- **Content-Type**: application/json, text/json, application/*+json
- **Accept**: application/json, application/json; profile=CamelCase, application/json; profile=PascalCase, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

