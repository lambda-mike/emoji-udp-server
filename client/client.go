package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udpAddr := net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: 54321, Zone: ""}
	conn, err := net.DialUDP("udp", nil /*(laddr)*/, &udpAddr)
	defer conn.Close()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("Cmd:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := scanner.Text()
		log.Println("INFO sending:", cmd)
		conn.Write([]byte(cmd))
		fmt.Println("Cmd:")
	}
}
