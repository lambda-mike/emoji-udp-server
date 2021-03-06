package contracts

type Cmd struct {
	N     uint
	Emoji string
}

type CmdHandler interface {
	Handle(cmd string)
}

type CmdServer interface {
	Listen()
}

type CmdParser interface {
	Parse(rawCmd string) (Cmd, error)
}

type CmdTransformer interface {
	Transform(cmd Cmd) (Cmd, error)
}

type ResponseBuilder interface {
	Build(cmd Cmd) string
}

type UI interface {
	Print(cmd string)
}

type MetricsProvider interface {
	GetReport() string
	IncValid()
	IncInvalid()
}
