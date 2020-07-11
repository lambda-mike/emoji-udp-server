package contracts

type Cmd struct {
	N     int
	Emoji string
}

type CmdHandler interface {
	Handle(cmd string)
}

type CmdServer interface {
	Listen(port int)
}

type CmdParser interface {
	Parse(cmd string) (Cmd, error)
}
