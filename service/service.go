package service

import (
	"github.com/emoji-udp-server/config"
	"github.com/emoji-udp-server/contracts"
	"log"
)

type EmojiService struct {
	ui contracts.UI
}

func (es *EmojiService) Handle(cmd string) {
	log.Println("INFO EmojiService handling cmd:", cmd)
	log.Println("TODO parse cmd")
	log.Println("TODO transform cmd (2 transformers composed)")
	log.Println("TODO build response")
	log.Println("TODO print response")
	es.ui.Print(cmd)
}

func Create(
	ui contracts.UI,
	cfg config.Config,
) contracts.CmdHandler {
	log.Println("INFO service.Create")
	es := EmojiService{}
	es.ui = ui
	return &es
}

type EmojiConcatenator struct{}

func (e *EmojiConcatenator) Build(cmd contracts.Cmd) (string, error) {
	// TODO
	return "TODO", nil
}

func CreateResponseBuilder(separator string) contracts.ResponseBuilder {
	// TODO
	log.Println("INFO cmd.CreateResponseBuilder")
	return &EmojiConcatenator{}
}
