package server

import (
	"bufio"
	"fmt"
	"github.com/emoji-udp-server/contracts"
	"os"
)

type UDPServer struct {
	handler contracts.CmdHandler
}

func (m *UDPServer) Listen(port int) {
	fmt.Println("Mock server pretending to listen on port", port)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		fmt.Println("Scanned:", cmd)
		m.handler.Handle(cmd)
	}
}

func CreateUDPServer(h contracts.CmdHandler) contracts.CmdServer {
	fmt.Println("CreateUDPServer")
	serv := UDPServer{}
	serv.handler = h
	return &serv
}
