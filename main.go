package main

// TODO use flags to parse cmd arguments

import (
	"fmt"
	"github.com/emoji-udp-server/config"
	"log"
)

func main() {
	isRaw := false
	n := 1
	sep := ""
	conf, err := config.Create(isRaw, n, sep)
	if err != nil {
		log.Panicln("Could not create Config struct, got error: ", err)
	}
	fmt.Println("Welcome to Emoji UDP Server!")
	log.Println("INFO: Config: ", conf)
}
