package server

import (
	"bufio"
	"fmt"
	"github.com/emoji-udp-server/contracts"
	"os"
)

type mock struct {
	handler contracts.CmdHandler
}

func (m *mock) addHandler(h contracts.CmdHandler) {
	// TODO
	m.handler = h
	m.handler.Handle("3 :thumbsup:")
}

func (m *mock) Listen(port int) {
	fmt.Println("Mock server pretending to listen on port", port)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		fmt.Println("Scanned:", cmd)
		m.handler.Handle(cmd)
	}
}

func CreateMock(h contracts.CmdHandler) contracts.CmdServer {
	fmt.Println("mock factory fn")
	m := mock{}
	m.addHandler(h)
	return &m
}

// TODO factory fn:
//func Create
