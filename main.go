package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	// Connexion
	client, err := connection()
	if err != nil {
		fmt.Printf("Connexion échouée: %v\n", err)
		return
	}
	defer client.Close()

	// Ping
	resp, err := client.Ping(ctx)
	if err != nil {
		fmt.Printf("Ping échoué: %v\n", err)
		return
	}
	fmt.Printf("Serveur connecté -> %s\n\n", resp)

	// Set / Get
	if err := setKey(ctx, *client, "user:1000", "matteo_martinez", 60); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if _, err := getKey(ctx, *client, "user:1000"); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// MSet / MGet
	keyValues := map[string]string{
		"user:1000": "matteo",
		"user:1001": "martinez",
	}

	if err := setMultipleKeysWithTTL(ctx, *client, keyValues, 60); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	keys := []string{"user:1000", "user:1001"}
	if err := getMultipleKeys(ctx, *client, keys); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// Delete
	if err := deleteKeys(ctx, *client, keys); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}
