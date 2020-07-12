package service

import (
	"errors"
	"fmt"
	"github.com/emoji-udp-server/cmd"
	"github.com/emoji-udp-server/config"
	"github.com/emoji-udp-server/contracts"
	"log"
	"strings"
)

type EmojiService struct {
	ui     contracts.UI
	parser contracts.CmdParser
}

func (es *EmojiService) Handle(req string) {
	log.Println("INFO EmojiService handling cmd:", req)
	command, err := es.parseRequest(req)
	log.Println("INFO cmd", command, "err", err)
	log.Println("TODO transform cmd (2 transformers composed)")
	log.Println("TODO build response")
	log.Println("TODO print response")
	es.ui.Print(req)
}

func Create(
	ui contracts.UI,
	cfg config.Config,
) contracts.CmdHandler {
	log.Println("INFO service.Create")
	es := EmojiService{}
	es.ui = ui
	es.parser = cmd.CreateParser()
	return &es
}

func (es *EmojiService) parseRequest(req string) (contracts.Cmd, error) {
	return es.parser.Parse(req)
}

type EmojiConcatenator struct {
	separator string
}

func (e *EmojiConcatenator) Build(cmd contracts.Cmd) string {
	if cmd.Emoji == "" {
		return ""
	}
	if cmd.N < 0 {
		return ""
	}
	emojis := strings.Repeat(cmd.Emoji+e.separator, cmd.N)
	if cmd.N > 0 {
		limit := len(emojis) - len(e.separator)
		emojis = emojis[:limit]
	}
	return emojis
}

func CreateResponseBuilder(s string) contracts.ResponseBuilder {
	log.Println("INFO cmd.CreateResponseBuilder")
	ec := EmojiConcatenator{}
	ec.separator = s
	return &ec
}

func UnknownCmdErr(cmd string) error {
	msg := fmt.Sprintf("Unknown command: {%s}", cmd)
	return errors.New(msg)
}
