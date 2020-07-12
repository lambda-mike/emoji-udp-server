package server

import (
	"bufio"
	"github.com/emoji-udp-server/contracts"
	"log"
	"net"
	"os"
)

type UDPServer struct {
	conn    *net.UDPConn
	handler contracts.CmdHandler
}

func CreateUDPServer(port int, h contracts.CmdHandler) contracts.CmdServer {
	log.Println("INFO CreateUDPServer")
	udpAddr := net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: port, Zone: ""}
	log.Println("INFO udpAddr:", udpAddr)
	serv := UDPServer{}
	conn, err := net.ListenUDP("udp", &udpAddr)
	if err != nil {
		log.Panicln("Could not set up UDP server with port:", port)
	}
	serv.conn = conn
	serv.handler = h
	return &serv
}

func (s *UDPServer) Listen() {
	log.Println("INFO UDP server is listenning...", s.conn.LocalAddr())
	defer s.conn.Close()
	// 1KB buffer
	buf := make([]byte, 1024)
	for {
		n, addr, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			log.Panicln("ERR When reading from UDP", err)
		}
		log.Println("INFO Received", n, "bytes from addr", addr)
		req := buf[:n]
		s.handler.Handle(string(req))
	}
}

type MockServer struct {
	handler contracts.CmdHandler
}

func (m *MockServer) Listen() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		log.Println("INFO Scanned:", cmd)
		m.handler.Handle(cmd)
	}
}

func CreateMockServer(port int, h contracts.CmdHandler) contracts.CmdServer {
	log.Println("INFO CreateUDPServer; port:", port)
	serv := MockServer{}
	serv.handler = h
	return &serv
}
