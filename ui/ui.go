package ui

import "fmt"

type Responder struct{}

func (r *Responder) Respond(cmd string) {
	fmt.Println(cmd)
}
