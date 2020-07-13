package main

import (
	"bytes"
	"github.com/emoji-udp-server/config"
	"github.com/emoji-udp-server/metrics"
	"github.com/emoji-udp-server/screen"
	"github.com/emoji-udp-server/server"
	"github.com/emoji-udp-server/service"
	"io"
	"log"
	"os"
)

func main() {
	initLogging()
	conf := config.ParseCmdLineFlags()
	log.Println("INFO Config: ", conf)
	port, err := config.ParsePort(config.ReadPortFromEnv())
	if err != nil {
		log.Panicln("ERR Did you add", config.EMOJI_PORT, "env var? Got err:", err)
	}
	mp := metrics.Create()
	ui := screen.Create()
	emojiService := service.Create(conf, mp, ui)
	udpServer, _ := server.CreateUDPServer(port, emojiService)
	udpServer.Listen()
}

func initLogging() {
	var sink io.Writer
	sink, err := os.Create("emoji.log")
	if err != nil {
		var buf bytes.Buffer
		log.SetOutput(&buf)
	}
	log.SetOutput(sink)
}
