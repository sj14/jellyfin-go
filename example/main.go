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
