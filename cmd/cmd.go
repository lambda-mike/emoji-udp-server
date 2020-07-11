package cmd

import (
	"errors"
	"github.com/emoji-udp-server/contracts"
)

type Parser struct{}

func (p *Parser) Parse(cmd string) (contracts.Cmd, error) {
	return contracts.Cmd{}, errors.New("TODO Parse Cmd")
}
