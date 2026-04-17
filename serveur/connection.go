package serveur

import (
	"context"
	"log"
	"log/slog"
	"os"
	"strconv"
	"time"

	glide "github.com/valkey-io/valkey-glide/go/v2"
	"github.com/valkey-io/valkey-glide/go/v2/config"
)

/*
Initialisation, connection avec valkey et création d'un client
*/
func Connection() (*glide.Client, error) {
	logger := slog.Default()
	//recup la config via la variable d'environnement
	host := os.Getenv("VALKEY_HOST")
	if host == "" {
		// Valeur par défault
		host = "localhost"
	}

	portValkey := os.Getenv("VALKEY_PORT")
	port, err := strconv.Atoi(portValkey)
	if err != nil {
		// Valeur par défault
		port = 6379
	}

	cfg := config.NewClientConfiguration().
		WithAddress(&config.NodeAddress{Host: host, Port: port}).
		WithRequestTimeout(5 * time.Second).
		WithUseTLS(false).
		WithDatabaseId(0)

	client, err := glide.NewClient(cfg)
	if err != nil {
		logger.Error("erreur création Valkey Client:", "error", err)
	}

	ctx := context.Background()
	// Test connexion
	_, err = client.Set(ctx, "connection_test", "ok")
	if err != nil {
		// Nettoyage avant de quitter
		client.Close()
		logger.Error("client créé mais Host inaccessible: ", "error", err)
	}

	log.Printf("\nConnexion réussie sur %s:%d", host, port)
	return client, nil

}
