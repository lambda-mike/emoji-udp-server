package service

import (
	"github.com/emoji-udp-server/contracts"
)

type EmojiService struct {
	ui contracts.UI
}

func (es *EmojiService) Handle(cmd string) {
	es.ui.Print("EmojiService handling cmd: " + cmd)
}

func Create(ui contracts.UI) contracts.CmdHandler {
	es := EmojiService{}
	es.ui = ui
	return &es
}
