package sample

type ActionCommand struct {
	steps []Command
}

func (a *ActionCommand) Execute(pos int) int {
	for _, step := range a.steps {
		pos = step.Execute(pos)
	}
	return pos
}

func (a *ActionCommand) AddStep(step Command) {
	a.steps = append(a.steps, step)
}

func (a *ActionCommand) Undo() {
	sz := len(a.steps)
	if sz == 0 {
		return
	}
	a.steps = a.steps[:sz-1]
}
