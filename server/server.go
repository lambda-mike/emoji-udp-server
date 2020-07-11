package server

import (
	"fmt"
	"github.com/emoji-udp-server/contracts"
)

type mock struct{}

func (m *mock) AddHandler(h contracts.CmdHandler) {
	// TODO
	h.Handle("3 :thumbsup:")
}

func CreateMock() contracts.CmdProducer {
	fmt.Println("mock factory fn")
	return &mock{}
}
