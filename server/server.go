package server

import "fmt"

type CmdProducer interface {
	AddHandler(func(cmd string))
}

type mock struct{}

func (m *mock) AddHandler(h func(cmd string)) {
	// TODO
	h("3 :thumbsup:")
}

func CreateMock() CmdProducer {
	fmt.Println("mock factory fn")
	return &mock{}
}
