package main

import (
	"context"
	"fmt"
	"time"

	glide "github.com/valkey-io/valkey-glide/go/v2"
	"github.com/valkey-io/valkey-glide/go/v2/config"
)

func connection() (*glide.Client, error) {
	// Configuration
	host := "localhost"
	port := 6379

	// Building the configuration
	cfg := config.NewClientConfiguration().
		WithAddress(&config.NodeAddress{Host: host, Port: port}).
		WithRequestTimeout(5 * time.Second).
		WithUseTLS(false).
		WithDatabaseId(0)

	// Initialize the client
	client, err := glide.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create Valkey client: %w", err)
	}

	ctx := context.Background()
	_, err = client.Set(ctx, "connection_test", "ok")
	if err != nil {
		return nil, fmt.Errorf("client created but host unreachable: %w", err)
	}

	return client, nil
}
