package main

import (
	"context"
	"fmt"

	glide "github.com/valkey-io/valkey-glide/go/v2"
)

func setKey(ctx context.Context, client glide.Client, key, value string) error {
	_, err := client.Set(ctx, key, value)
	if err != nil {
		return fmt.Errorf("erreur Set [%s]: %w", key, err)
	}
	fmt.Printf("Set réussi -> %s = %s\n", key, value)
	return nil
}

func getKey(ctx context.Context, client glide.Client, key string) (string, error) {
	value, err := client.Get(ctx, key)
	if err != nil {
		return "", fmt.Errorf("erreur Get [%s]: %w", key, err)
	}
	if value.IsNil() {
		fmt.Printf("Clé introuvable: %s\n", key)
		return "", nil
	}
	fmt.Printf("Get réussi -> %s = %s\n", key, value.Value())
	return value.Value(), nil
}

func setMultipleKeys(ctx context.Context, client glide.Client, keyValues map[string]string) error {
	_, err := client.MSet(ctx, keyValues)
	if err != nil {
		return fmt.Errorf("erreur MSet: %w", err)
	}
	fmt.Printf("MSet réussi -> %d clés insérées\n", len(keyValues))
	return nil
}

func getMultipleKeys(ctx context.Context, client glide.Client, keys []string) error {
	values, err := client.MGet(ctx, keys)
	if err != nil {
		return fmt.Errorf("erreur MGet: %w", err)
	}
	fmt.Println("MGet réussi:")
	for i, val := range values {
		if val.IsNil() {
			fmt.Printf("%s -> introuvable\n", keys[i])
		} else {
			fmt.Printf("%s -> %s\n", keys[i], val.Value())
		}
	}
	return nil
}

func deleteKeys(ctx context.Context, client glide.Client, keys []string) error {
	count, err := client.Del(ctx, keys)
	if err != nil {
		return fmt.Errorf("erreur Del: %w", err)
	}
	fmt.Printf("%d clé(s) supprimée(s)\n", count)
	return nil
}
