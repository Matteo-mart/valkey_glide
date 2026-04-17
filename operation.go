package main

import (
	"context"
	"fmt"
	"time"

	glide "github.com/valkey-io/valkey-glide/go/v2"
)

func setKey(ctx context.Context, client glide.Client, key, value string, ttlSeconds int64) error {
	_, err := client.Set(ctx, key, value)
	if err != nil {
		return fmt.Errorf("Erreur Set [%s]: %w", key, err)
	}

	if ttlSeconds > 0 {
		_, err = client.Expire(ctx, key, time.Duration(ttlSeconds)*time.Second)
		if err != nil {
			return fmt.Errorf("Erreur Expire [%s]: %w", key, err)
		}
	}

	fmt.Printf("Set réussi -> %s = %s (TTL: %ds)\n", key, value, ttlSeconds)
	return nil
}

func setMultipleKeysWithTTL(ctx context.Context, client glide.Client, keyValues map[string]string, ttlSeconds int64) error {
	for k, v := range keyValues {
		err := setKey(ctx, client, k, v, ttlSeconds)
		if err != nil {
			return err
		}
	}
	return nil
}

func getKey(ctx context.Context, client glide.Client, key string) (string, error) {
	value, err := client.Get(ctx, key)
	if err != nil {
		return "", fmt.Errorf("Erreur Get [%s]: %w", key, err)
	}
	if value.IsNil() {
		fmt.Printf("Clé introuvable: %s\n", key)
		return "", nil
	}
	fmt.Printf("Get réussi -> %s = %s\n", key, value.Value())
	return value.Value(), nil
}

func getMultipleKeys(ctx context.Context, client glide.Client, keys []string) error {
	values, err := client.MGet(ctx, keys)
	if err != nil {
		return fmt.Errorf("Erreur MGet: %w", err)
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
		return fmt.Errorf("Erreur Del: %w", err)
	}
	fmt.Printf("%d clé(s) supprimée(s)\n", count)
	return nil
}
