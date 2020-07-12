package main

import (
	"github.com/emoji-udp-server/config"
	"github.com/emoji-udp-server/screen"
	"github.com/emoji-udp-server/server"
	"github.com/emoji-udp-server/service"
	"log"
)

func main() {
	conf, err := config.ParseCmdLineFlags()
	if err != nil {
		log.Panicln("ERR Could not create Config struct from command-line arguments, got error: ", err)
	}
	log.Println("INFO Config: ", conf)
	port, err := config.ParsePort(config.ReadPortFromEnv())
	if err != nil {
		log.Panicln("ERR Did you add", config.EMOJI_PORT, "env var? Got err:", err)
	}
	ui := screen.Create()
	emojiService := service.Create(ui, conf)
	udpServer := server.CreateUDPServer(port, emojiService)
	udpServer.Listen()
}
