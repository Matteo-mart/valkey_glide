package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	getset "valkey_glide/getSet"
	"valkey_glide/serveur"
)

func main() {

	logger := slog.Default()
	serveur.Launch()

	ctx := context.Background()

	// Initialisation de la connexion
	client, err := serveur.Connection()
	if err != nil {
		logger.Error("Impossible de démarrer l'application", "error", err)
		return
	}
	defer client.Close()

	// Test de Ping
	if resp, err := client.Ping(ctx); err != nil {
		logger.Error("Ping échoué", "error", err)
	} else {
		logger.Info("Serveur connecté", "response", resp)
	}

	// Opérations de base
	if err := getset.SetKey(ctx, *client, "user:1000", "matteo_martinez", 60); err != nil {
		logger.Error("Erreur SetKey", "error", err)
	}

	if val, err := getset.GetKey(ctx, *client, "user:1000"); err != nil {
		logger.Error("Erreur GetKey", "error", err)
	} else {
		logger.Info("Valeur récupérée", "key", "user:1000", "value", val)
	}

	// Gestion d'objets complexes (JSON)
	user := getset.User{ID: 1001, Email: "matteo.martinez@example.com"}
	if err := getset.SetUser(ctx, *client, "user:obj:1001", user); err != nil {
		logger.Error("Erreur lors de la sauvegarde user", "user_id", user.ID, "error", err)
	}

	// Opérations groupées
	keyValues := map[string]string{
		"temp:1": "data1",
		"temp:2": "data2",
	}
	if err := getset.SetMultipleKeysWithTTL(ctx, *client, keyValues, 120); err != nil {
		logger.Error("Erreur lors de l'insertion multiple", "error", err)
	}

	// Nettoyage
	if count, err := getset.DeleteKeys(ctx, *client, []string{"temp:1", "temp:2"}); err != nil {
		logger.Error("Erreur lors du nettoyage", "error", err)
	} else {
		logger.Info("Nettoyage effectué", "deleted_count", count)
	}

	logger.Info("Toutes les opérations terminées")

	port := uint32(6379)
	err = serveur.Shutdown(port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erreur : %v\n", err)
		os.Exit(1)
	}
}
