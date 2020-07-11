package cmd

import (
	"errors"
	"github.com/emoji-udp-server/contracts"
	"log"
)

type parser struct{}

func (p *parser) Parse(cmd string) (contracts.Cmd, error) {
	return contracts.Cmd{}, errors.New("TODO Parse Cmd")
}

type multiplier struct {
	n int
}

func (p *multiplier) Transform(cmd contracts.Cmd) contracts.Cmd {
	cmd.N *= p.n
	return cmd
}

func CreateMultiplier(n int) contracts.CmdTransformer {
	log.Println("INFO cmd.CreateResponseBuilder")
	m := multiplier{}
	m.n = n
	return &m
}

type IdentityTranslator struct{}

func (p *IdentityTranslator) Transform(cmd contracts.Cmd) contracts.Cmd {
	return cmd
}

// TODO add lookup table
type memoryTableTranslator struct{}

func (m *memoryTableTranslator) Transform(cmd contracts.Cmd) contracts.Cmd {
	// TODO
	return cmd
}

func CreateTranslator(raw bool) contracts.CmdTransformer {
	log.Println("INFO cmd.CreateTranslator")
	if raw {
		return &IdentityTranslator{}
	} else {
		return &memoryTableTranslator{}
	}
}

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
	log.Println("INFO cmd.CreateResponseBuilder")
	return nil
}
