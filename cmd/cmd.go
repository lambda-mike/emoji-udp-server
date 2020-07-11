package cmd

import (
	"errors"
	"github.com/emoji-udp-server/contracts"
	"log"
)

type Parser struct{}

func (p *Parser) Parse(cmd string) (contracts.Cmd, error) {
	return contracts.Cmd{}, errors.New("TODO Parse Cmd")
}

type Multiplyer struct{}

func (p *Multiplyer) Transform(cmd contracts.Cmd) contracts.Cmd {
	// TODO
	return cmd
}

type IdentityTranslator struct{}

func (p *IdentityTranslator) Transform(cmd contracts.Cmd) contracts.Cmd {
	return cmd
}

// TODO
// MemoryTableTranslator

type ResponseBuilder struct{}

func (p *ResponseBuilder) Build(cmd contracts.Cmd) (string, error) {
	// TODO
	return "TODO", nil
}

func CreateResponseBuilder(
	separator string,
	t contracts.CmdTransformer,
) contracts.CmdResponseBuilder {
	// TODO
	log.Println("cmd.CreateResponseBuilder")
	return nil
}
