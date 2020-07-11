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

func CreateParser() contracts.CmdParser {
	return &parser{}
}

type multiplier struct {
	n int
}

func (p *multiplier) Transform(cmd contracts.Cmd) (contracts.Cmd, error) {
	// TODO what if integer overflows?
	cmd.N *= p.n
	return cmd, nil
}

func CreateMultiplier(n int) contracts.CmdTransformer {
	log.Println("INFO cmd.CreateResponseBuilder")
	m := multiplier{}
	m.n = n
	return &m
}

type IdentityTranslator struct{}

func (p *IdentityTranslator) Transform(cmd contracts.Cmd) (contracts.Cmd, error) {
	return cmd, nil
}

type memoryTableTranslator struct {
	table map[string]string
}

func (m *memoryTableTranslator) Transform(cmd contracts.Cmd) (contracts.Cmd, error) {
	emoji, ok := m.table[cmd.Emoji]
	if !ok {
		msg := "Could not translate emoji: " + cmd.Emoji
		log.Println("WARN", msg)
		return cmd, errors.New(msg)
	}
	cmd.Emoji = emoji
	return cmd, nil
}

func CreateMemoryTableTranslator() *memoryTableTranslator {
	m := memoryTableTranslator{}
	m.table = map[string]string{
		":thumbsup:":   "👍",
		":thumbsdown:": "👎",
		":ok:":         "👌",
		":crossed:":    "🤞",
	}
	return &m
}

func CreateTranslator(raw bool) contracts.CmdTransformer {
	log.Println("INFO cmd.CreateTranslator")
	if raw {
		return &IdentityTranslator{}
	} else {
		return CreateMemoryTableTranslator()
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
