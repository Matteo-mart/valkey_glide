package serveur

import (
	"fmt"
	"syscall"

	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

func Shutdown(port uint32) error {
	connections, err := net.Connections("tcp")
	if err != nil {
		return fmt.Errorf("impossible de lister les connexions : %v", err)
	}

	found := false
	for _, conn := range connections {
		if conn.Laddr.Port == port && conn.Pid != 0 {
			found = true
			fmt.Printf("Processus trouvé (PID: %d) sur le port %d. Tentative d'arrêt...\n", conn.Pid, port)

			p, err := process.NewProcess(conn.Pid)
			if err != nil {
				return fmt.Errorf("erreur lors de l'accès au processus %d : %v", conn.Pid, err)
			}

			err = p.SendSignal(syscall.SIGKILL)
			if err != nil {
				return fmt.Errorf("impossible de tuer le processus %d : %v", conn.Pid, err)
			}

			fmt.Printf("Processus %d terminé avec succès.\n", conn.Pid)
		}
	}

	if !found {
		fmt.Printf("Aucun processus n'utilise le port %d.\n", port)
	}

	return nil
}
