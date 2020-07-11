package contracts

type CmdHandler interface {
	Handle(cmd string)
}

type CmdServer interface{}
