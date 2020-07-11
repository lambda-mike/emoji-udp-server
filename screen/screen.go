package screen

import (
	"fmt"
	"github.com/emoji-udp-server/contracts"
)

type Screen struct{}

func (r *Screen) Print(cmd string) {
	fmt.Println(cmd)
}

func Create() contracts.UI {
	return &Screen{}
}
