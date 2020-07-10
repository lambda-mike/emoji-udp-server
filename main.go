package main

import (
	"flag"
	"fmt"
	"github.com/emoji-udp-server/config"
	"log"
)

func main() {
	isRaw := flag.Bool("r", false, "disable the translation from keyword to emoji")
	n := flag.Int("n", 1, "cmd input number multiplier")
	sep := flag.String("s", "", "emojis separator")
	flag.Parse()
	conf, err := config.Create(*isRaw, *n, *sep)
	if err != nil {
		log.Panicln("ERR: Could not create Config struct from command-line arguments, got error: ", err)
	}
	fmt.Println("Welcome to Emoji UDP Server!")
	log.Println("INFO: Config: ", conf)
}
