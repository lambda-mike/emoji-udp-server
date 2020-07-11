package contracts

type CmdHandler interface {
	Handle(cmd string)
}

type CmdProducer interface {
	AddHandler(h CmdHandler)
}
