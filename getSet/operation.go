package getset

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"time"

	glide "github.com/valkey-io/valkey-glide/go/v2"
)

// User structure pour la gestion JSON
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func SetKey(ctx context.Context, client glide.Client, key, value string, ttlSeconds int64) error {
	_, err := client.Set(ctx, key, value)
	if err != nil {
		return fmt.Errorf("erreur Set [%s]: %w", key, err)
	}

	if ttlSeconds > 0 {
		_, err = client.Expire(ctx, key, time.Duration(ttlSeconds)*time.Second)
		if err != nil {
			return fmt.Errorf("erreur Expire [%s]: %w", key, err)
		}
	}

	log.Printf("Set réussi -> %s (TTL: %ds)\n", key, ttlSeconds)
	return nil
}

/*
setMultipleKeysWithTTL :
*/
func SetMultipleKeysWithTTL(ctx context.Context, client glide.Client, keyValues map[string]string, ttlSeconds int64) error {
	for k, v := range keyValues {
		if err := SetKey(ctx, client, k, v, ttlSeconds); err != nil {
			return err
		}
	}
	return nil
}

/*
getKey :
*/
func GetKey(ctx context.Context, client glide.Client, key string) (string, error) {
	value, err := client.Get(ctx, key)
	if err != nil {
		return "", fmt.Errorf("erreur technique Get [%s]: %w", key, err)
	}

	if value.IsNil() {
		slog.Debug("Clé inexistante", "key", key)
		return "", nil
	}

	return value.Value(), nil
}

/*
getMultipleKeys :
*/
func GetMultipleKeys(ctx context.Context, client glide.Client, keys []string) ([]string, error) {
	values, err := client.MGet(ctx, keys)
	if err != nil {
		return nil, fmt.Errorf("erreur MGet: %w", err)
	}

	results := make([]string, len(values))
	for i, val := range values {
		if !val.IsNil() {
			results[i] = val.Value()
		}
	}
	return results, nil
}

/*
deleteKeys :
*/
func DeleteKeys(ctx context.Context, client glide.Client, keys []string) (int64, error) {
	count, err := client.Del(ctx, keys)
	if err != nil {
		return 0, fmt.Errorf("erreur suppression: %w", err)
	}
	return count, nil
}

/*
setUser :
*/
func SetUser(ctx context.Context, client glide.Client, key string, user User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("erreur sérialisation user [%d]: %w", user.ID, err)
	}

	return SetKey(ctx, client, key, string(data), 3600)
}

/*
getUser :
*/
func GetUser(ctx context.Context, client glide.Client, key string) (*User, error) {
	result, err := client.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("erreur Get user [%s]: %w", key, err)
	}

	if result.IsNil() {
		return nil, nil
	}

	var user User
	if err := json.Unmarshal([]byte(result.Value()), &user); err != nil {
		return nil, fmt.Errorf("erreur désérialisation user [%s]: %w", key, err)
	}

	return &user, nil
}
