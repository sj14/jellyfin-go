# Jellyfin API Go Client

Auto-generated Go client package for [Jellyfin](https://jellyfin.org/). See the [API Documentation](/api/README.md). I mainly built it for usage in [jellyctl](https://github.com/sj14/jellyctl).

## Example

Generate an API token in the Jellyfin WebUI at Administration -> Overview -> Advanced -> API Token.

Minimal example to get started:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jellyfin "github.com/sj14/jellyfin-go/api"
)

func main() {
	config := &jellyfin.Configuration{
		Servers:       jellyfin.ServerConfigurations{{URL: "http://<ADDRESS>:<PORT>"}},
		DefaultHeader: map[string]string{"Authorization": `MediaBrowser Token="<API_TOKEN>"`},
	}
	client := jellyfin.NewAPIClient(config)

	result, resp, err := client.ActivityLogAPI.GetLogEntries(context.Background()).Execute()
	if err != nil {
		log.Printf("Error when calling `GetLogEntries``: %v\n", err)
		log.Printf("Full HTTP response: %v\n", resp)
		os.Exit(1)
	}

	for _, i := range result.Items {
		fmt.Println(*i.Name)
	}
}
```

## Development

The `configuration.go` file was adjusted (renamed the struct `ServerConfiguration` to `APIServerConfiguration`) and ignored from further generations to avoid conflicts with the same struct form the `model_server_configuration.go` file. In case the `configuration.go` needs to be regenerated, the struct needs to be renamed again.
