package service

import (
	"fmt"
	"github.com/emoji-udp-server/cmd"
	"github.com/emoji-udp-server/config"
	"github.com/emoji-udp-server/contracts"
	"log"
)

const MetricsCmd string = ":metrics:"

type EmojiService struct {
	metrics      contracts.MetricsProvider
	parser       contracts.CmdParser
	respBuilder  contracts.ResponseBuilder
	transformers []contracts.CmdTransformer
	ui           contracts.UI
}

func (es *EmojiService) Handle(req string) {
	log.Println("INFO EmojiService handling request:", req)
	if req == MetricsCmd {
		resp := es.metrics.GetReport()
		es.ui.Print(resp)
		return
	}
	command, err := es.parseRequest(req)
	if err != nil {
		es.handleErr(req, err)
		return
	}
	log.Println("DBG cmd", command, "err", err)
	cmdTran, err := es.transformCmd(command, err)
	if err != nil {
		es.handleErr(req, err)
		return
	}
	log.Println("DBG cmd", cmdTran, "err", err)
	resp := es.respBuilder.Build(cmdTran)
	es.ui.Print(resp)
	es.metrics.IncValid()
}

func (es *EmojiService) handleErr(req string, err error) {
	log.Println("WARN handleErr", req, err)
	if err == nil {
		return
	}
	es.metrics.IncInvalid()
	if err == cmd.TooBigNumErr {
		resp :=
			"N in request: '" + req + "' is too big! It could cause integer overflow!"
		es.ui.Print(resp)
		return
	}
	resp := CreateUnknownCmdMsg(req)
	es.ui.Print(resp)
}

func Create(
	cfg config.Config,
	metrics contracts.MetricsProvider,
	ui contracts.UI,
) contracts.CmdHandler {
	log.Println("INFO service.Create")
	es := EmojiService{}
	es.metrics = metrics
	es.parser = cmd.CreateParser()
	es.respBuilder = CreateResponseBuilder(cfg.Separator)
	transformers := make([]contracts.CmdTransformer, 0, 2)
	transformers = append(transformers, cmd.CreateMultiplier(cfg.N))
	transformers = append(transformers, cmd.CreateTranslator(cfg.Raw))
	es.transformers = transformers
	es.ui = ui
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
	emojis := ""
	if cmd.N > 0 {
		// emojis will contain extra sep at the end
		var i uint
		for i = 0; i < cmd.N; i++ {
			emojis += cmd.Emoji + e.separator
		}
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
