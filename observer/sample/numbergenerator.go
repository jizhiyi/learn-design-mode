package sample

type NumberGenerator interface {
	AddObserver(observer Observer)
	GetNumber() int
	NotifyObserver()
	Execute()
}

type NumberGeneratorBase struct {
	NumberGenerator
	observers []Observer
}

func (n *NumberGeneratorBase) AddObserver(observer Observer) {
	n.observers = append(n.observers, observer)
}

func (n *NumberGeneratorBase) NotifyObserver() {
	for _, observer := range n.observers {
		observer.Updata(n)
	}
}
