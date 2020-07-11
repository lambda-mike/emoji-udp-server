package service

import (
	"github.com/emoji-udp-server/contracts"
	"log"
)

type EmojiService struct {
	ui contracts.UI
}

func (es *EmojiService) Handle(cmd string) {
	log.Println("INFO EmojiService handling cmd:", cmd)
	es.ui.Print(cmd)
}

func Create(ui contracts.UI) contracts.CmdHandler {
	log.Println("INFO service.Create")
	es := EmojiService{}
	es.ui = ui
	return &es
}
