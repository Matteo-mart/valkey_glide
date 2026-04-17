package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	glide "github.com/valkey-io/valkey-glide/go/v2"
	"github.com/valkey-io/valkey-glide/go/v2/config"
)

func connection() (glide.Client, error) {
	//recup la config via la variable d'environnement
	host := os.Getenv("VALKEY_HOST")
	if host == "" {
		host = "localhost" // Valeur par défaut
	}

	portValkey := os.Getenv("VALKEY_PORT")
	port, err := strconv.Atoi(portValkey)
	if err != nil {
		port = 6379 // Valeur par défaut
	}

	cfg := config.NewClientConfiguration().
		WithAddress(&config.NodeAddress{Host: host, Port: port}).
		WithRequestTimeout(5 * time.Second).
		WithUseTLS(false).
		WithDatabaseId(0)

	client, err := glide.NewClient(cfg)
	if err != nil {
		// On retourne l'erreur pour que le main puisse la gérer
		fmt.Errorf("erreur création Valkey Client: %w", err)
	}

	ctx := context.Background()
	_, err = client.Set(ctx, "connection_test", "ok") //
	if err != nil {
		// Nettoyage avant de quitter
		client.Close()
		fmt.Errorf("client créé mais Host inaccessible: %w", err)
	}

	log.Printf("Connexion réussie sur %s:%d", host, port)
	return *client, nil
}
