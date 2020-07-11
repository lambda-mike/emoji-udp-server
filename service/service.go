package service

import (
	// TODO remove?
	"fmt"
	"github.com/emoji-udp-server/contracts"
)

type mock struct{}

func (m *mock) Handle(cmd string) {
	fmt.Println("I am Mock service; handling cmd: ", cmd)
}

func CreateMock() contracts.CmdHandler {
	return &mock{}
}
