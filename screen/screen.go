package screen

import (
	"fmt"
	"github.com/emoji-udp-server/contracts"
	"log"
)

type Screen struct{}

func (r *Screen) Print(cmd string) {
	fmt.Println(cmd)
}

func Create() contracts.UI {
	log.Println("INFO screen.Create")
	return &Screen{}
}
