package sample

type Actor struct {
	initPos int
	curPos  int
	Command
}

func NewActor(command Command) *Actor {
	a := &Actor{initPos: 0}
	a.Command = command
	return a
}

func (a *Actor) DoAction() {
	a.curPos = a.Command.Execute(a.initPos)
}

func (a *Actor) GetCurPos() int {
	return a.curPos
}
