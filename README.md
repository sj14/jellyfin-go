# Jellyfin API Go Client

Auto-generated Go client package for [Jellyfin](https://jellyfin.org/). See the [API Documentation](/api/README.md).

The Jellyfin `ServerConfiguration` is not yet generated due to the conflicting struct from the generator.

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
