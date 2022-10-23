package sample

import "fmt"

type Worker struct {
	curState state
	hour     int
}

func (w *Worker) SetHour(hour int) {
	w.hour = hour
	doClock(w, hour)
}

func (w *Worker) setState(s state) {
	w.curState = s
}

func (w *Worker) DoThing() {
	fmt.Printf("Now [%d], DoThing: ", w.hour)
	w.curState.doThing()
}
