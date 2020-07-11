package server

import (
	"bufio"
	"github.com/emoji-udp-server/contracts"
	"log"
	"os"
)

type UDPServer struct {
	handler contracts.CmdHandler
}

func (m *UDPServer) Listen(port int) {
	log.Println("INFO Mock server pretending to listen on port", port)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		log.Println("INFO Scanned:", cmd)
		m.handler.Handle(cmd)
	}
}

func CreateUDPServer(h contracts.CmdHandler) contracts.CmdServer {
	log.Println("INFO CreateUDPServer")
	serv := UDPServer{}
	serv.handler = h
	return &serv
}
