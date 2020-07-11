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
	Parse(rawCmd string) (Cmd, error)
}

type CmdTransformer interface {
	Transform(cmd Cmd) (Cmd, error)
	// TODO Chain
}

type ResponseBuilder interface {
	Build(cmd Cmd) (string, error)
}

type UI interface {
	Print(cmd string)
}
