package main

import (
	"context"
	"fmt"
)

func main() {
	client, err := connection()
	if err != nil {
		fmt.Printf("Il y a une erreur: %v\n", err)
		return
	}
	defer client.Close()

	resp, err := client.Ping(context.Background())
	if err != nil {
		fmt.Printf("Il y a une erreur: %v\n", err)
		return
	}

	fmt.Printf("Connecté, Réponse du serveur: %s\n", resp)
}
