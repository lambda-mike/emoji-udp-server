package service

import (
	"fmt"
	"github.com/emoji-udp-server/cmd"
	"github.com/emoji-udp-server/config"
	"github.com/emoji-udp-server/contracts"
	"log"
	"strings"
)

type EmojiService struct {
	ui           contracts.UI
	parser       contracts.CmdParser
	transformers []contracts.CmdTransformer
	respBuilder  contracts.ResponseBuilder
}

func (es *EmojiService) Handle(req string) {
	log.Println("INFO EmojiService handling cmd:", req)
	command, err := es.parseRequest(req)
	if err != nil {
		es.handleErr(req, err)
		return
	}
	log.Println("INFO cmd", command, "err", err)
	log.Println("TODO transform cmd (2 transformers composed)")
	cmdTran, err := es.transformCmd(command, err)
	if err != nil {
		es.handleErr(req, err)
		return
	}
	log.Println("INFO cmd", cmdTran, "err", err)
	log.Println("TODO build response")
	resp := es.respBuilder.Build(cmdTran)
	es.ui.Print(resp)
}

func (es *EmojiService) handleErr(req string, err error) {
	log.Println("DBG handleErr", req, err)
	if err != nil {
		resp := CreateUnknownCmdMsg(req)
		es.ui.Print(resp)
	}
}

func Create(
	ui contracts.UI,
	cfg config.Config,
) contracts.CmdHandler {
	log.Println("INFO service.Create")
	es := EmojiService{}
	es.ui = ui
	es.parser = cmd.CreateParser()
	transformers := make([]contracts.CmdTransformer, 0, 2)
	transformers = append(transformers, cmd.CreateMultiplier(cfg.N))
	transformers = append(transformers, cmd.CreateTranslator(cfg.Raw))
	es.transformers = transformers
	es.respBuilder = CreateResponseBuilder(cfg.Separator)
	return &es
}

func (es *EmojiService) parseRequest(req string) (contracts.Cmd, error) {
	return es.parser.Parse(req)
}

func (es *EmojiService) transformCmd(c contracts.Cmd, err error) (contracts.Cmd, error) {
	if err != nil {
		return c, err
	}
	for _, t := range es.transformers {
		c, err = t.Transform(c)
		if err != nil {
			break
		}
	}
	return c, err
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
	// emojis will contain extra sep at the end
	emojis := strings.Repeat(cmd.Emoji+e.separator, cmd.N)
	if cmd.N > 0 {
		limit := len(emojis) - len(e.separator)
		// Remove extra separator at the end
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

func CreateUnknownCmdMsg(cmd string) string {
	return fmt.Sprintf("Unknown command: {%s}", cmd)
}
