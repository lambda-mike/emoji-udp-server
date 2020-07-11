package server

import (
	"fmt"
	"github.com/emoji-udp-server/contracts"
)

type mock struct {
	handler contracts.CmdHandler
}

func (m *mock) addHandler(h contracts.CmdHandler) {
	// TODO
	m.handler = h
	m.handler.Handle("3 :thumbsup:")
}

func CreateMock(h contracts.CmdHandler) contracts.CmdServer {
	fmt.Println("mock factory fn")
	m := mock{}
	m.addHandler(h)
	return m
}

// TODO factory fn:
//func Create
