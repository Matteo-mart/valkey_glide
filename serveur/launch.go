package serveur

import (
	"log"
	"os/exec"
)

func Launch() {
	log.Println("Lancement du serveur en arrière-plan...")

	cmd := exec.Command("valkey-server")

	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Start()
	if err != nil {
		log.Printf("Erreur lors du lancement : %v\n", err)
		return
	}

	err = cmd.Process.Release()
	if err != nil {
		log.Printf("Erreur lors du détachement du processus : %v\n", err)
		return
	}

	log.Printf("Serveur Valkey lancé en arrière-plan (PID: %d)\n", cmd.Process.Pid)
}
