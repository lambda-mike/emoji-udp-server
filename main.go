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
	// TODO move to config lib?
	// TODO ParseCmdLineFlags
	// -h param is handled by the flag lib automatically
	isRaw := flag.Bool("r", false, "disable the translation from keyword to emoji")
	n := flag.Int("n", 1, "cmd input number multiplier")
	sep := flag.String("s", "", "emojis separator")
	flag.Parse()
	conf, err := config.Create(*isRaw, *n, *sep)
	if err != nil {
		log.Panicln("ERR Could not create Config struct from command-line arguments, got error: ", err)
	}
	log.Println("INFO Config: ", conf)
	port, err := config.ParsePort(config.ReadPortFromEnv())
	if err != nil {
		log.Panicln("ERR Did you add", config.EMOJI_PORT, "env var? Got err:", err)
	}
	// TODO use real UDP server as CmdServer
	ui := screen.Create()
	emojiService := service.Create(ui, conf)
	udpServer := server.CreateUDPServer(emojiService)
	udpServer.Listen(port)
}
