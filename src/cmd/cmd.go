package cmd

import (
	"errors"
	"fmt"
	"github.com/emoji-udp-server/contracts"
	"log"
)

const MaxUint = ^uint(0)

var TooBigNumErr = errors.New("Given numbers are too big")

type parser struct{}

func (p *parser) Parse(rawCmd string) (contracts.Cmd, error) {
	var (
		n     uint
		emoji string
	)
	_, err := fmt.Sscanf(rawCmd, "%d %s", &n, &emoji)
	cmd := contracts.Cmd{}
	if err != nil {
		log.Println("WARN Parse error", err)
		return cmd, err
	}
	cmd.N = n
	cmd.Emoji = emoji
	return cmd, nil
}

func CreateParser() contracts.CmdParser {
	return &parser{}
}

type multiplier struct {
	n uint
}

// TODO agree with product team on reasonable max N
func (p *multiplier) Transform(cmd contracts.Cmd) (contracts.Cmd, error) {
	// Mutliplication will overflow silently for big numbers
	// Let's find out the upper bound and error explicitly if that would be the case
	limit := MaxUint / p.n
	if cmd.N > limit {
		return cmd, TooBigNumErr
	}
	cmd.N *= p.n
	return cmd, nil
}

func CreateMultiplier(n uint) contracts.CmdTransformer {
	log.Println("INFO cmd.CreateMultiplier")
	m := multiplier{}
	m.n = n
	return &m
}

type memoryTableTranslator struct {
	table               map[string]string
	translationDisabled bool
}

func (m *memoryTableTranslator) Transform(cmd contracts.Cmd) (contracts.Cmd, error) {
	emoji, ok := m.table[cmd.Emoji]
	if !ok {
		msg := "Could not translate emoji: " + cmd.Emoji
		log.Println("WARN", msg)
		return cmd, errors.New(msg)
	}
	if !m.translationDisabled {
		cmd.Emoji = emoji
	}
	return cmd, nil
}

func CreateTranslator(raw bool) *memoryTableTranslator {
	log.Println("INFO cmd.CreateTranslator")
	m := memoryTableTranslator{}
	m.translationDisabled = raw
	m.table = map[string]string{
		":thumbsup:":   "👍",
		":thumbsdown:": "👎",
		":ok:":         "👌",
		":crossed:":    "🤞",
	}
	return &m
}
