package serveur

import (
	"fmt"
	"syscall"

	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

func Shutdown(port uint32) error {

	found := false

	connections, err := net.Connections("tcp")
	if err != nil {
		return fmt.Errorf("impossible de lister : %v", err)
	}

	killedPids := make(map[int32]bool)

	for _, conn := range connections {
		if conn.Laddr.Port == port && conn.Pid != 0 {
			if killedPids[conn.Pid] {
				continue
			}

			p, err := process.NewProcess(conn.Pid)
			if err != nil {
				continue
			}

			if err := p.SendSignal(syscall.SIGKILL); err == nil {
				fmt.Printf("Processus %d (port %d) terminé.\n", conn.Pid, port)
				killedPids[conn.Pid] = true
			}
		}
	}

	if !found {
		fmt.Printf("Aucun processus n'utilise le port %d.\n", port)
	}

	return nil
}
