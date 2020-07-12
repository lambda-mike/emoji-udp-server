package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	//54321
	port := 54321
	udpAddr := net.UDPAddr{IP: []byte{127, 0, 0, 1}, Port: port, Zone: ""}
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
