package screen

import (
	"fmt"
	"github.com/emoji-udp-server/contracts"
	"log"
	"sync"
)

type Screen struct {
	mtx sync.Mutex
}

func (r *Screen) Print(cmd string) {
	r.mtx.Lock()
	fmt.Println(cmd)
	r.mtx.Unlock()
}

func Create() contracts.UI {
	log.Println("INFO screen.Create")
	return &Screen{}
}
