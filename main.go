package main

import (
	"flag"
	"github.com/emoji-udp-server/config"
	"github.com/emoji-udp-server/screen"
	"github.com/emoji-udp-server/server"
	"github.com/emoji-udp-server/service"
	"log"
)

func main() {
	isRaw := flag.Bool("r", false, "disable the translation from keyword to emoji")
	n := flag.Int("n", 1, "cmd input number multiplier")
	sep := flag.String("s", "", "emojis separator")
	// -h param is handled by flag lib itself automatically
	flag.Parse()
	conf, err := config.Create(*isRaw, *n, *sep)
	if err != nil {
		log.Panicln("ERR Could not create Config struct from command-line arguments, got error: ", err)
	}
	// TODO remove
	log.Println("INFO Config: ", conf)
	// TODO get port from env var
	port := 54321
	// TODO use real UDP server as CmdServer
	ui := screen.Create()
	emojiService := service.Create(ui)
	udpServer := server.CreateUDPServer(emojiService)
	udpServer.Listen(port)
}
