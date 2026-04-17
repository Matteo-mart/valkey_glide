package main

import (
	"context"
	"fmt"
	"time"

	glide "github.com/valkey-io/valkey-glide/go/v2"
	"github.com/valkey-io/valkey-glide/go/v2/config"
)

func connection() (*glide.Client, error) {
	host := "localhost"
	port := 6379

	cfg := config.NewClientConfiguration().
		WithAddress(&config.NodeAddress{Host: host, Port: port}).
		WithRequestTimeout(10 * time.Second).
		WithUseTLS(false).
		WithDatabaseId(0)

	keyValue := map[string]string{
		"user:1000": "matteo",
		"user:1001": "martinez",
	}

	client, err := glide.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("Erreur nouveau client: %w\n", err)
	}
	// defer client.Close()

	_, err = client.Set(context.Background(), "user:1000", "matteo_martinez")
	if err != nil {
		return nil, fmt.Errorf("Erreur Set: %v\n", err)
	}
	fmt.Println("Set réussi")

	value, err := client.Get(context.Background(), "user:1000")
	if value.IsNil() {
		fmt.Println("Clé introuvable")
	} else {
		fmt.Printf("Get Réussi: %s\n", value.Value())
	}

	_, err = client.MSet(context.Background(), keyValue)
	if err != nil {
		return nil, fmt.Errorf("Erreur MSET: %v\n", err)
	}
	fmt.Println("MSet réussi")

	keys := []string{"user:1000", "user:1001"}
	values, err := client.MGet(context.Background(), keys)
	if err != nil {
		return nil, fmt.Errorf("Erreur MGet\n")
	}
	fmt.Printf("MGet réussi: %v\n", values)

	deleteCount, err := client.Del(context.Background(), keys)
	if err != nil {
		return nil, fmt.Errorf("Erreur Del: %v\n", err)
	}
	fmt.Printf("Clés %d supprimés\n", deleteCount)

	return client, nil
}
