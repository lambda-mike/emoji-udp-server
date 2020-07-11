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

type CmdTransformer interface {
	Transform(cmd Cmd) Cmd
	// TODO Chain
}

type CmdResponseBuilder interface {
	Build(cmd Cmd) (string, error)
}
